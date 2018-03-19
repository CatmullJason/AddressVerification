package functions

import (
	"AddressVerification/models"
	"encoding/json"
	"fmt"
)

/********************************************************
* AddressVerify: Receives an address and verifies that
* it is correct by posting to EasyPost's address
* verification API
*********************************************************/
func AddressVerify(address models.RequestAddress) (models.Verification, error) {

	request := models.Request{}
	request.AddHeader("Content-Type", "application/json")
	request.AddHeader("Authorization", "Basic VnJvODROdkNuMFJIclNwZnBRdm0wUTo=")
	request.AddBody(address)
	request.SetUri("https://api.easypost.com/v2/addresses")
	request.SetMethod("POST")

	//Make request
	responseBody, err := SendRequest(request)

	//Initialize response address
	responseAddress := models.ResponseAddress{}

	err = json.Unmarshal(responseBody, &responseAddress)
	if err != nil {
		fmt.Println(err)
		return models.Verification{}, fmt.Errorf("Unrecognized response from easypost %s", err.Error())
	}

	return responseAddress.Verifications.Delivery, err
}
