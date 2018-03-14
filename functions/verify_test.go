package functions

import (
	"AddressVerification/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressVerification(t *testing.T) {

	testAddress := models.RequestAddress{
		//ID:          "adr_4ff39618e9f5429298ef69b2d436884b",
		//Object:      "Address",
		Mode:    "test",
		Street1: "417 MONTGOMERY ST",
		//Street2: "FLOOR 5",
		City:  "SAN FRANCISCO",
		State: "CA",
		//Zip:     "94104",
		//Country: "US",
		//Residential: false,
		//CarrierFacility: "",
		//Name:            "",
		//Company: "EasyPost",
		//Phone:   "4151234567",
		//Email:           "",
		//FederalTaxID:    "",
		//StateTaxID:      "",
		VerifyStrict: []string{"delivery"},
	}

	verifyResponse, err := AddressVerify(testAddress)

	//fmt.Println(verifyResponse)

	assert.Nil(t, err)
	assert.True(t, verifyResponse.Success)
}
