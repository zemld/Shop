package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/constants"
	"github.com/zemld/Shop/api-gateway/domain/dto"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Checks if user is registered.
// @tag.name File operations
// @param user path string true "User which you want to check"
// @produce json
// @success 200 {object} dto.UserRegistered
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/{user} [get]
func CheckUserRegistered(w http.ResponseWriter, r *http.Request) {
	user, err := internal.TryParseURLPathParam(r.URL.Path, constants.UserRegisteredPath, constants.PathParam)
	if err != nil {
		log.Println("Can't parse user from request URL")
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    "",
				Message: "Can't parse user from request URL"},
			http.StatusBadRequest)
		return
	}
	log.Printf("Parsed user: %s\n", user)
	response, err := sendRequestToUserService(r.URL.Path)
	if err != nil {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    user,
			Message: "Can't recieve response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, user)
}

func sendRequestToUserService(path string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://user-service:8081%s", path), nil)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't get response from user service.")
		return nil, err
	}
	log.Printf("Got response from user service: %s\n", response.Body)
	return response, nil
}
