package api

import (
	"auth-service/logspb"
	"auth-service/utils"
	"fmt"
	"log"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
)

func (server *Server) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := server.writeJson(w, status, env, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func (server *Server) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	server.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (server *Server) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	server.background(func() {
		err := server.emitter.SendToLogService(&logspb.Log{
			ErrorMessage: err.Error(),
			StackTrace:   fmt.Sprintf("request_method: %s, request_url: %s", r.Method, r.URL.String()),
		})
		if err != nil {
			log.Println(err)
		}
	})
	log.Println("error: ", err)

	message := "the server encountered a problem and could not process your request"
	server.errorResponse(w, r, http.StatusInternalServerError, message)
}
func (server *Server) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	server.errorResponse(w, r, http.StatusUnauthorized, message)
}
func (server *Server) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	server.errorResponse(w, r, http.StatusNotFound, message)
}

func (server *Server) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	message := "authentication is impossible for this user and browsers will not propose a new attempt."
	server.errorResponse(w, r, http.StatusForbidden, message)
}

func (server *Server) unauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid credentials"
	server.errorResponse(w, r, http.StatusUnauthorized, message)
}
func (serve *Server) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	serve.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
func (server *Server) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	server.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
func (server *Server) validationErrorResponse(w http.ResponseWriter, r *http.Request, err error, userValidator *utils.UserValidtor) {
	validationErrors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		validationErrors[e.Field()] = e.Translate(userValidator.Trans)
	}
	server.failedValidationResponse(w, r, validationErrors)
}

func (server *Server) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	server.errorResponse(w, r, http.StatusUnauthorized, message)
}
func (server *Server) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	server.errorResponse(w, r, http.StatusUnauthorized, message)
}
