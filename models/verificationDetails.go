package models

type VerificationDetails struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	TimeZone  string  `json:"time_zone"`
}
