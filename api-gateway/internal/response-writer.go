package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/dto"
)

func TryParseResponseBodyAndWriteResponse(w http.ResponseWriter, response *http.Response, userID int) {
	responceBody := (*response).Body
	defer response.Body.Close()
	log.Printf("Response body: %s\n", responceBody)
	var parsedResponse dto.UserRegistered
	err := json.NewDecoder(responceBody).Decode(&parsedResponse)
	if err != nil {
		log.Println("Can't parse response body from user service.")
		WriteResponse(w, dto.StatusResponse{
			UserID:  userID,
			Message: "Can't parse response body from user service.",
		}, http.StatusInternalServerError)
		return
	}
	WriteResponse(w, parsedResponse, response.StatusCode)
}

func WriteResponse(w http.ResponseWriter, reponseBody interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(reponseBody); err != nil {
		http.Error(w, "Can't parse body", http.StatusInternalServerError)
		return
	}
}
