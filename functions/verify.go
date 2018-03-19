package functions

import (
	"AddressVerification/models"
	"encoding/json"
	"fmt"
)

/********************************************************
* AddressVerify : Creates a request with accompanying
* headers, body, URI, method, etc. and then calls SendRequest
* to receive a response from a specified API
*********************************************************/
func AddressVerify(address models.Address) (models.Verification, error) {

	//Create request
	request := models.Request{}
	request.AddHeader("Content-Type", "application/json")
	request.AddHeader("Authorization", "Basic VnJvODROdkNuMFJIclNwZnBRdm0wUTo=")
	request.AddBody(address)
	request.SetUri("https://api.easypost.com/v2/addresses")
	request.SetMethod("POST")

	//Send request
	responseBody, err := SendRequest(request)

	//Initialize response address
	responseAddress := models.Address{}

	err = json.Unmarshal(responseBody, &responseAddress)
	if err != nil {
		fmt.Println(err)
		return models.Verification{}, fmt.Errorf("Unrecognized response from easypost %s", err.Error())
	}

	return responseAddress.Verifications.Delivery, err
}
