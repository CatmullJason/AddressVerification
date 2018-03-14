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

	requestBody, err := json.Marshal(address)
	fmt.Println(string(requestBody))

	//Assign API key, uri, and http method
	apiKey := "VnJvODROdkNuMFJIclNwZnBRdm0wUTo="
	uri := "https://api.easypost.com/v2/addresses"
	method := "POT"

	//Make request
	responseBody, err := NewRequest(requestBody, uri, apiKey, method)

	/*client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.easypost.com/v2/addresses", bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", "Basic VnJvODROdkNuMFJIclNwZnBRdm0wUTo=")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.Verification{}, fmt.Errorf("error contacting easypost api %s", err.Error())
	}*/

	//Initialize response address
	responseAddress := models.ResponseAddress{}

	//read response and unmarshal
	//responseBody, err := ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()

	fmt.Println(string(responseBody))
	err = json.Unmarshal(responseBody, &responseAddress)
	if err != nil {
		fmt.Println(err)
		return models.Verification{}, fmt.Errorf("Unrecognized response from easypost %s", err.Error())
	}

	fmt.Println(responseAddress)
	return responseAddress.Verifications.Delivery, err
}
