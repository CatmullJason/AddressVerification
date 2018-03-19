package functions

import (
	"AddressVerification/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressVerification(t *testing.T) {

	testAddress := models.RequestAddress{
		//ID:          "adr_4ff39618e9f5429298ef69b2d436884b",
		//Object:      "Address",
		Mode:            "test",
		Street1:         "5200 Skidmore Dr.",
		Street2:         "",
		City:            "Idaho Falls",
		State:           "",
		Zip:             "83406",
		Country:         "US",
		Residential:     false,
		CarrierFacility: "",
		Name:            "",
		Company:         "EasyPost",
		Phone:           "",
		Email:           "",
		FederalTaxID:    "",
		StateTaxID:      "",
		VerifyStrict:    []string{"delivery"},
	}

	verifyResponse, err := AddressVerify(testAddress)

	fmt.Println(err)

	assert.Nil(t, err)
	assert.True(t, verifyResponse.Success)
}
