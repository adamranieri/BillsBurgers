package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type PromoCode struct {
	Code     string  `json:"code"`
	Uses     int     `json:"uses"`
	Discount float32 `json:"discount"`
}

type RedemptionTicket struct {
	Code string `json:"code"`
}

func Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Promo Service up and running</h1>"))
}

func ReadPromoData() []PromoCode {
	jsonFile, err := os.Open("promo-codes.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	promobytes, _ := ioutil.ReadAll(jsonFile)

	var promos []PromoCode
	json.Unmarshal(promobytes, &promos)
	return promos

}

func Redeem(w http.ResponseWriter, req *http.Request) {
	var ticket RedemptionTicket
	json.NewDecoder(req.Body).Decode(&ticket)
	promos := ReadPromoData()
	fmt.Println(ticket.Code)

	for _, promo := range promos {
		fmt.Println(promo, ticket.Code)
		if promo.Code == ticket.Code {

			if promo.Uses > 0 {
				promo.Uses = promo.Uses - 1
				j, _ := json.Marshal(promos)
				file, _ := os.Create("promo-codes.json")
				defer file.Close()

				file.Truncate(0)
				file.Seek(0, 0)
				file.WriteString(string(j))
				data, _ := json.Marshal(promo)

				w.Write(data)
				return
			} else {
				w.WriteHeader(422)
				w.Write([]byte("promo code has no uses left"))
				return
			}
		}
	}
	w.WriteHeader(404)
	w.Write([]byte("promo code not found"))
}

func main() {

	http.HandleFunc("/redeem", Redeem)
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8888", nil)
}
