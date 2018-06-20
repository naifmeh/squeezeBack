package models

import "time"

type (
	Employee struct {
		id string `json:"id, omitempty"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
		RegisteredOn time.Time `json: "registeredOn,omitempty""`
		Authorized bool `json:"authorized"`
		Frequency int `json:"frequency""`
		AuthStarting int64 `json:"authStarting"`
		AuthEnding int64 `json:"authEnding""`
	}
	Device struct {
		DeviceName string `json:"deviceName"`
		DeviceMac string `json: "deviceMac, omitempty"`
	}
)

