package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
)

func TryParseResponseBodyAndWriteResponse(w http.ResponseWriter, response *http.Response, responseType string) {
	responseBody := (*response).Body
	defer response.Body.Close()
	log.Printf("Response body: %s\n", responseBody)

	parsedResponse := getResponseStruct(responseType)

	err := json.NewDecoder(responseBody).Decode(&parsedResponse)
	if err != nil {
		log.Println("Can't parse response body from user service.")
		WriteResponse(w, models.StatusResponse{
			Name:    "",
			Message: "Can't parse response body from user service.",
		}, http.StatusInternalServerError)
		return
	}
	WriteResponse(w, parsedResponse, response.StatusCode)
}

func getResponseStruct(responseType string) any {
	if responseType == models.IsAdminResponseType {
		return models.IsAdminResponse{}
	}
	if responseType == models.AdminType {
		return models.Admin{}
	}
	if responseType == models.AddItemResponseType {
		return models.ItemResponse{}
	}
	if responseType == models.DeliverItemResponseType {
		return models.ItemDeliveredResponse{}
	}
	if responseType == models.PurchaseItemResponseType {
		return models.ItemBoughtResponse{}
	}
	if responseType == models.UpdatePriceResponseType {
		return models.ItemWithNewPriceResponse{}
	}
	if responseType == models.ItemResponseType {
		return models.ItemResponse{}
	}
	if responseType == models.OrderResponseType {
		return models.OrderStatusResponse{}
	}
	return models.StatusResponse{}
}
