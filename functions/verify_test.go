package functions

import (
	"AddressVerification/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressVerification(t *testing.T) {

	SetAPIKey("Basic VnJvODROdkNuMFJIclNwZnBRdm0wUTo=")

	testAddress := models.RequestAddress{
		//ID:          "",
		//Object:  "Address",
		Mode:    "test",
		Street1: "5200 E Skidmore Dr.",
		//Street2:     "",
		City: "Idaho Falls",
		//State:       "",
		Zip:     "83406",
		Country: "US",
		//Residential: false,
		//CarrierFacility: "",
		//Name:    "",
		//Company: "EasyPost",
		//Phone:           "",
		//Email:           "",
		//FederalTaxID:    "",
		//StateTaxID:      "",
		VerifyStrict: []string{"delivery"},
	}

	found, address, err := CreateAddress(testAddress)
	fmt.Println(found)
	fmt.Printf("This is the address object: %v", address)
	fmt.Println(address.Verifications.Delivery.Errors)

	assert.Nil(t, err)
	assert.True(t, found)
}
