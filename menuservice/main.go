package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MenuItem struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

func ReadMenuData() []MenuItem {
	jsonFile, err := os.Open("menu.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	menubytes, _ := ioutil.ReadAll(jsonFile)

	var menu []MenuItem
	json.Unmarshal(menubytes, &menu)
	return menu

}

func Menu(w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(ReadMenuData())
	w.Write(data)
}

func Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Menu Service is up and running</h1>"))
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/menu", Menu)

	http.ListenAndServe(":7777", nil)
}
