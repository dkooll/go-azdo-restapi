## Purpose
This is a mini project that demonstrates the use of Go programming language to interact with the Azure DevOps API. The purpose of this project is to learn and understand some basic Go concepts.

## Index
* [Description](#description)
* [Prerequisites](#prerequisites)
* [Overview](#overview)
    * [Function Flow](#function-flow)
* [Details](#details)
    * [Functions](#functions)
    * [Testing](#testing)
* [Notes](#Notes)
* [Concepts Learned](#concepts-learned)

## Description
This code creates an agent pool in Azure DevOps. It first checks if it already exists by sending a GET request to the Azure DevOps API, then it will send a POST request to create the pool if it does not exist.

The program uses structs and the json package to map the response to struct fields. It also uses the HTTP package to make the GET and POST requests to the API. The program defines some constants, such as the pool name and organization name, which are used in the API requests.

## Prerequisites
Before running this code, you will need to have the following:

- A Microsoft Azure DevOps account
- A Personal Access Token (PAT) with the "Agent Pools (read and manage)" scope, to authenticate with the Azure DevOps API.
- Go programming language installed on your computer
- Basic knowledge of Go programming
- Set an environment variable named `AZURE_TOKEN` with your PAT as its value.

## Overview
### Function Flow

![Rest Api](/images/rest-api.png "rest api")

## Details
### Functions

The program starts by importing several Go packages that are used throughout the code, including `net/http` for making HTTP requests and `encoding/json` for working with JSON data. The program also defines some constants, such as the pool name and organization name, which are used in the API requests.

The program defines two structs: "Response" and "Pools". These structs are used to unmarshal the response body from the API when querying the list of pools. The structs are defined with json tags which allow the json package to map the json response to struct fields.

The function [checkAgentPoolExists](https://github.com/dkooll/go-azdo-restapi/blob/09ae17b16452ae1e85d0298f6e63eca4e4095b14/main.go#L28) is used to check if the agent pool already exists or not. This function sends a GET request to the Azure DevOps API to check if the specified pool name already exists. The url of the API is constructed by concatenating the organization name and the API endpoint. The API endpoint is used to retrieve the list of pools in the organization. The function also sets the headers for the request, including the "Authorization" header. This is used to pass the token to the API.

The function creates a new http.Client, then it uses the client to execute the GET request. The response body is read and unmarshaled into the struct Response. Then the function loops through the pools in the response and checks if the name of the pool is equal to the specified pool name. If a pool is found with the specified name, the function returns true, otherwise it returns false.

The function [createAgentPool](https://github.com/dkooll/go-azdo-restapi/blob/09ae17b16452ae1e85d0298f6e63eca4e4095b14/main.go#L62) is used to create the agent pool. This function first calls the checkAgentPoolExists function to check if the agent pool already exists. If the pool already exists, the function returns an error message. If the pool does not exist, the function creates a new http.Client, then it uses json package to marshal the pool struct into bytes.

It creates a new http.Request with the POST method, the url of the API and the bytes buffer of the marshaled pool struct. The headers for the request are set as the previous function, including the "Authorization" header. Then the function uses the client to execute the POST request. The response status code is checked, if it's not OK, the function returns an error message with the status.

In the main function, the program creates a new struct of Pools with the specified pool name, AutoProvision set to true and IsHosted set to false, then it calls the createAgentPool function. If the function returns an error, the program will print the error message, otherwise it will print a message that the pool has been created.

### Testing

The first function, [TestCreateAgentPoolSuccess](https://github.com/dkooll/go-azdo-restapi/blob/07af1300b6929ae3160640f9e1558861c0f818cf/agentpool_test.go#L9) is using a mock server. The httptest package is used to create a test server and the http.HandlerFunc function is used to handle the requests made to the server. The test server checks the request method and returns a response accordingly, with a status code of 200 OK for a POST request and 409 Conflict for a GET request. The global variable "client" is overridden with the test server's client so that the createAgentPool function makes its request to the test server instead of a real server. The test checks that the createAgentPool function returns an error with the message "error: 409 Conflict" and fails if it does not.

The [TestCreateAgentPoolInvalidInput](https://github.com/dkooll/go-azdo-restapi/blob/07af1300b6929ae3160640f9e1558861c0f818cf/agentpool_test.go#L39) test case verifies that the createAgentPool function properly handles invalid input. It creates a Pools struct with an empty pool name and calls the createAgentPool function with it. The test case asserts that an error is returned and that the error message is the expected message "error: Invalid input: pool name cannot be empty".

## Notes

- This code is provided as-is, and is intended for educational purposes only. It is not production-ready and may not be suitable for your specific use case.

- The Azure DevOps API has a rate limit of 30 requests per minute. If you exceed this limit, your requests will be blocked for a period of time.

- The Azure DevOps API is versioned, please make sure to use the correct API version for your organization.
- The API endpoint and payload for creating a agent pool might change with different API version.
- The code in this repository uses the basic authentication method, it's recommended to use the OAuth authentication method for production use.

## Concepts Learned

- Using linked structs
- Making GET and POST requests to the azure devops api
- Using the json package to marshal structs into json and unmarshal json response into structs
- Accessing environment variables using the os package
- Process function inputs and return values