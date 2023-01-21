// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// const orgName = "cloudnation-nl"

// type Response struct {
// 	Count int     `json:"count"`
// 	Pools []Pools `json:"value"`
// }

// type Pools struct {
// 	Name          string `json:"name"`
// 	AutoProvision bool   `json:"autoProvision"`
// 	IsHosted      bool   `json:"isHosted"`
// }

// func getAgentPools() (Response, error) {
// 	res, err := http.NewRequest("GET", "https://dev.azure.com/"+orgName+"/_apis/distributedtask/pools?api-version=7.0", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	res.Header.Add("Authorization", "Basic "+os.Getenv("AZURE_TOKEN"))
// 	res.Header.Add("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var response Response
// 	json.Unmarshal(body, &response)

// 	return response, nil
// }

// func main() {
// 	response, err := getAgentPools()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for i, p := range response.Pools {
// 		fmt.Println("Pool:", (i + 1), ":", p.Name)
// 		fmt.Println("----------------------------------------")
// 		fmt.Println("Auto Provision:", p.AutoProvision)
// 		fmt.Println("Is Hosted:", p.IsHosted)
// 		fmt.Printf("\n\n")
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// type Response struct {
// 	Count int     `json:"count"`
// 	Pools []Pools `json:"value"`
// }

// type Pools struct {
// 	Name          string `json:"name"`
// 	AutoProvision bool   `json:"autoProvision"`
// 	IsHosted      bool   `json:"isHosted"`
// }

// func main() {
// 	res, err := http.NewRequest("GET", "https://dev.azure.com/cloudnation-nl/_apis/distributedtask/pools?api-version=7.0", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	res.Header.Add("Authorization", "Basic "+os.Getenv("AZURE_TOKEN"))
// 	res.Header.Add("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var response Response
// 	json.Unmarshal(body, &response)

// 	for i, p := range response.Pools {
// 		fmt.Println("Pool:", (i + 1), ":", p.Name)
// 		fmt.Println("----------------------------------------")
// 		fmt.Println("Auto Provision:", p.AutoProvision)
// 		fmt.Println("Is Hosted:", p.IsHosted)
// 		fmt.Printf("\n\n")
// 	}
// }
