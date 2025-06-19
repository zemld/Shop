package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
	"github.com/zemld/Shop/payment-service/domain/constants"
	"github.com/zemld/Shop/payment-service/domain/models"
	"github.com/zemld/Shop/payment-service/mq"
)

func PayForOrder() {
	nc, err := nats.Connect(mq.URL)
	if err != nil {
		log.Fatalf("Error connecting mq: %v", err)
	}
	_, err = nc.Subscribe(mq.Storage, func(msg *nats.Msg) {
		log.Printf("Recieved message: %s", string(msg.Data))
		var orderMsg models.OrderMsg
		if err := json.Unmarshal(msg.Data, &orderMsg); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}
		log.Printf("Processing order for user %s with items: %v", orderMsg.User, orderMsg.Items)
		order := models.Order{
			User:  orderMsg.User,
			Items: orderMsg.Items,
		}
		orderCost := getOrderCost(order)
		query := make(map[string]string)
		query["name"] = order.User
		userBalanceResponse, err := SendRequestToUserService(constants.GET, constants.GetUserBalance, query)
		if err != nil {
			log.Printf("Error sending request to user service: %v", err)
			return
		}
		if userBalanceResponse.StatusCode != http.StatusOK {
			log.Printf("Error getting user balance: %s", userBalanceResponse.Status)
			return
		}
		var user models.User
		if err := json.NewDecoder(userBalanceResponse.Body).Decode(&user); err != nil {
			log.Printf("Error decoding user response: %v", err)
			return
		}
		log.Printf("User %s has balance %f, order cost is %f", user.Name, user.Balance, orderCost)
		if user.Balance < orderCost {
			orderMsg.Message = "Not enough money"
			encodedResult, _ := json.Marshal(orderMsg)
			log.Printf("Insufficient balance for user %s: %f < %f", user.Name, user.Balance, orderCost)
			if err := nc.Publish(mq.PaymentCancel, encodedResult); err != nil {
				log.Printf("Error publishing cancel message: %v", err)
				return
			}
		}
		query["balance"] = fmt.Sprintf("%f", user.Balance-orderCost)
		updateBalanceResponse, err := SendRequestToUserService(constants.POST, constants.UpdateUserBalance, query)
		if err != nil {
			log.Printf("Error updating user balance: %v", err)
			return
		}
		if updateBalanceResponse.StatusCode != http.StatusOK {
			log.Printf("Error updating user balance: %s", updateBalanceResponse.Status)
			return
		}
		if err := nc.Publish(mq.OrderHandled, msg.Data); err != nil {
			log.Printf("Error publishing message: %v", err)
			return
		}
	})
	if err != nil {
		log.Fatalf("Error mq subscription: %v", err)
	}
}

func getOrderCost(order models.Order) float64 {
	var cost float64
	for _, item := range order.Items {
		cost += item.Price * float64(item.Amount)
	}
	return cost
}
