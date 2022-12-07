package endpoints

import (
	"github.com/53jk1/go-base/jwt"
	"log"
	"net/http"
)

func Validate(writer http.ResponseWriter, request *http.Request) {
	token := request.FormValue("token")
	j := jwt.NewJWT(token)
	valid, err := j.Validate()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	writer.WriteHeader(http.StatusOK)
}