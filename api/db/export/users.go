package export

import (
	"github.com/Zenika/MARCEL/api/db/users"
)

type userPassword struct {
	users.User
	PasswordHash string `json:"passwordHash"`
	PasswordSalt string `json:"passwordSalt"`
}

func Users(withPassword bool, outputFile string) error {
	return export(func() (interface{}, error) {
		users, err := users.List()
		if err != nil {
			return nil, err
		}

		if withPassword {
			var usersPassword = make([]userPassword, 0, len(users))

			for _, user := range users {
				usersPassword = append(usersPassword, userPassword{
					User:         user,
					PasswordHash: user.PasswordHash,
					PasswordSalt: user.PasswordSalt,
				})
			}

			return usersPassword, nil
		}

		return users, nil
	}, outputFile)
}
