package functions

import (
	"AddressVerification/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressVerification(t *testing.T) {

	testAddress := models.Address{
		//ID:          "adr_4ff39618e9f5429298ef69b2d436884b",
		//Object:      "Address",
		Mode:    "test",
		Street1: "2020 S. Luxury Ln.",
		//Street2:     "Apt. I204",
		City:        "Idaho Falls",
		State:       "ID",
		Zip:         "83642",
		Country:     "US",
		Residential: true,
		//CarrierFacility: "",
		Name:    "",
		Company: "EasyPost",
		//Phone:           "",
		//Email:           "",
		//FederalTaxID:    "",
		//StateTaxID:      "",
		VerifyStrict: []string{"delivery"},
	}

	verifyResponse, err := AddressVerify(testAddress)

	fmt.Println(err)

	assert.Nil(t, err)
	assert.True(t, verifyResponse.Success)
}
