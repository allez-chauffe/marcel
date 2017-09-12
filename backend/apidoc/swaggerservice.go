package apidoc

import "net/http"

// swagger:route GET /swagger.json getSwaggerConfig
//
// Gets swagger.json file, generated from these API
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger.json")
}
