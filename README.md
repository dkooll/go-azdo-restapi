# go-azdo-restapi

This code is written in Go, it creates an agent pool in Azure DevOps. It first checks if the agent pool already exists by sending a GET request to the Azure DevOps API, then it will send a POST request to create the pool if it does not exist.

The program uses structs and json package to map the json response to struct fields. It also uses HTTP package to make the GET and POST requests to the API. The program defines some constants, such as the pool name and organization name, which are used in the API requests.

The program also uses the os package to get the env variable "AZURE_TOKEN" which is used as an Authorization header in the requests.