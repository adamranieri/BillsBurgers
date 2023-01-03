package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type FoodOrder struct {
	PromoCode    string        `json:"promoCode"`
	FoodRequests []FoodRequest `json:"items"`
}

type RedemptionTicket struct {
	Code string `json:"code"`
}

type MenuItem struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

func Process(ctx *gin.Context) {

}

func CalculateCost(order FoodOrder) float32 {
	return 0
}

func FindCost(name string) (float32, error) {
	response, _ := http.Get("http://localhost:7777/menu")
	var menuitems []MenuItem
	json.NewDecoder(response.Body).Decode(&menuitems)
	defer response.Body.Close()

	for _, item := range menuitems {
		if item.Name == name {
			return item.Price, nil
		}
	}

	return 0, errors.New("Item was not found")

}

func RedeemCode(code string) (string, error) {

	ticket := RedemptionTicket{Code: code}
	data, _ := json.Marshal(ticket)
	response, _ := http.Post("http://localhost:8888/redeem", "application/json", bytes.NewBuffer(data))
	switch response.StatusCode {
	case 200:
		return "redeemed", nil
	case 422:
		return "depleted", nil
	case 404:
		return "not found", nil
	default:
		return "", errors.New("Service in unreachable or encountered internal error")
	}

}

func main() {

	// app := gin.Default()

	// app.Run()

}
