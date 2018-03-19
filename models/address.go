package models

type Address struct {
	ID              string        `json:"id"`
	Object          string        `json:"object"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	Mode            string        `json:"mode"`
	Street1         string        `json:"street1"`
	Street2         string        `json:"street2"`
	City            string        `json:"city"`
	State           string        `json:"state"`
	Zip             string        `json:"zip"`
	Country         string        `json:"country"`
	Residential     bool          `json:"residential"`
	CarrierFacility string        `json:"carrier_facility"`
	Name            string        `json:"name"`
	Company         string        `json:"company"`
	Phone           string        `json:"phone"`
	Email           string        `json:"email"`
	FederalTaxID    string        `json:"federal_tax_id"`
	StateTaxID      string        `json:"state_tax_id"`
	Verifications   Verifications `json:"verifications"`
	Verify          []string      `json:"verify"`
	VerifyStrict    []string      `json:"verify_strict"`
}
