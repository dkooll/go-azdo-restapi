package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	poolName = "Selfhosted"
	orgName  = "cloudnation"
)

type Response struct {
	Count int     `json:"count"`
	Pools []Pools `json:"value"`
}

type Pools struct {
	Name          string `json:"name"`
	AutoProvision bool   `json:"autoProvision"`
	IsHosted      bool   `json:"isHosted"`
}

func checkAgentPoolExists(poolName string) (bool, error) {
	req, err := http.NewRequest("GET", "https://dev.azure.com/"+orgName+"/_apis/distributedtask/pools?api-version=7.0", nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Basic "+os.Getenv("AZURE_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var response Response
	json.Unmarshal(body, &response)

	for _, p := range response.Pools {
		if p.Name == poolName {
			return true, nil
		}
	}

	return false, nil
}

func createAgentPool(poolname Pools) error {
	exists, err := checkAgentPoolExists(poolName)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("pool %s already exists", poolName)
	}

	client := &http.Client{}
	poolBytes, err := json.Marshal(poolname)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://dev.azure.com/"+orgName+"/_apis/distributedtask/pools?api-version=7.0", bytes.NewBuffer(poolBytes))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+os.Getenv("AZURE_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	defer resp.Body.Close()
	return nil
}

func main() {
	pool := Pools{
		Name:          poolName,
		AutoProvision: true,
		IsHosted:      false,
	}

	err := createAgentPool(pool)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Pool %s created\n", poolName)
}
