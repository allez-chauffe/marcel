package auth

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/commons"
	"github.com/allez-chauffe/marcel/api/db"
)

type Service struct{}

func NewService() *Service {
	return new(Service)
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	cred := s.getCredentials(w, r)

	if cred != nil {
		s.loginWithCredentials(w, cred.Login, cred.Password)
		return
	}

	s.loginWithRefreshToken(w, r)
}

func (s *Service) loginWithCredentials(w http.ResponseWriter, login string, password string) {
	if login == "" || password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := db.Users().GetByLogin(login)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		commons.WriteResponse(w, http.StatusUnauthorized, "")
		return
	}

	ok, err := user.CheckPassword(password)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		commons.WriteResponse(w, http.StatusUnauthorized, "")
		return
	}

	GenerateAuthToken(w, user)
	GenerateRefreshToken(w, user)

	commons.WriteJsonResponse(w, user)
}

func (s *Service) loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshClaims, err := GetRefreshToken(r)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if refreshClaims == nil {
		commons.WriteResponse(w, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	user, err := db.Users().Get(refreshClaims.Subject)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusUnauthorized, "User not found")
		return
	}

	if user.LastDisconnection.Unix() > refreshClaims.IssuedAt {
		DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusUnauthorized, "Refresh token has been invalidated")
		return
	}

	GenerateAuthToken(w, user)

	commons.WriteJsonResponse(w, user)
}

func (s *Service) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if CheckPermissions(r, nil) {
		// If the user is connected, update it in database
		userID := GetAuth(r).Subject

		if err := db.Users().Disconnect(userID); err != nil {
			// Do not return here we want to delete the tokens
			log.Errorf("Error while disconnecting user %s: %s", userID, err)
		}
	}

	DeleteAuthToken(w)
	DeleteRefreshToken(w)
}

func (s *Service) getCredentials(w http.ResponseWriter, r *http.Request) *Credentials {
	credentials := &Credentials{}

	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		return nil
	}

	return credentials
}

func (s *Service) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if auth := GetAuth(r); auth == nil {
		commons.WriteResponse(w, http.StatusForbidden, "")
	}
}

func (s *Service) ValidateAdminHandler(w http.ResponseWriter, r *http.Request) {
	if !CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}
}
