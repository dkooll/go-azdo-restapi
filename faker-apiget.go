// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type Response struct {
// 	Status  string    `json:"status"`
// 	Code    int       `json:"code"`
// 	Total   int       `json:"total"`
// 	Persons []Persons `json:"data"`
// }

// type Persons struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Email     string `json:"email"`
// 	Phone     string `json:"phone"`
// 	Birthday  string `json:"birthday"`
// 	Gender    string `json:"gender"`
// 	Address   Address
// 	Website   string `json:"website"`
// 	Image     string `json:"image"`
// }

// type Address struct {
// 	Street         string  `json:"street"`
// 	Streetname     string  `json:"streetname"`
// 	Buildingnumber string  `json:"buildingnumber"`
// 	City           string  `json:"city"`
// 	Zipcode        string  `json:"zipcode"`
// 	Country        string  `json:"country"`
// 	Countrycode    string  `json:"countrycode"`
// 	Latitude       float64 `json:"latitude"`
// 	Longitude      float64 `json:"longitude"`
// }

// func main() {
// 	res, err := http.Get("https://fakerapi.it/api/v1/persons?_quantity=5&_gender=female&_birthday_start=2005-01-01")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var response Response
// 	json.Unmarshal(body, &response)

// 	for i, p := range response.Persons {
// 		fmt.Println("Person:", (i + 1), ":", p.Firstname, p.Lastname)
// 		fmt.Println("----------------------------------------")
// 		fmt.Println("Email:", p.Email)
// 		fmt.Println("Phone", p.Phone)
// 		fmt.Println("Birthday", p.Birthday)
// 		fmt.Println("Gender", p.Gender)
// 		fmt.Println("Address", p.Address.Street, p.Address.Streetname, p.Address.Buildingnumber)
// 		fmt.Printf("\n\n")
// 	}
// }
