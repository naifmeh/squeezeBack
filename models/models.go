package models

import "time"

type (
	Employee struct {
		id string `json:"id, omitempty"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastname"`
		Email string `json:"email"`
		RegisteredOn time.Time `json: "registeredOn,omitempty""`
		Authorized bool `json:"authorized"`
		Frequency int `json:"frequency""`
		AuthStarting uint16 `json:"authStarting"`
		AuthEnding uint16 `json:"authEnding""`
	}
	Device struct {
		DeviceName string `json:"deviceName"`
		DeviceMac string `json: "deviceMac, omitempty"`
	}
)

