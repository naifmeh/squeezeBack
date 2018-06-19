package controllers

import (
	"net/http"
	"encoding/json"
	"squeezecnn/common"
	"squeezecnn/data"
	"squeezecnn/models"
	"log"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource DeviceResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w,
			err,
			"Invalid device data",
			500)
		return
	}
	device := &dataResource.Data

	context := GetContext()
	log.Print(device)
	erreur := data.RegisterDevice(device,context.RethinkSession)
	if erreur != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	} else {
		if j, err := json.Marshal(DeviceResource{Data: *device}); err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
			return
		} else {
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(j)
		}
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var dataResource AuthResource
	var token string

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login data",
			500,
		)
		return
	}

	authModel := dataResource.Data
	log.Print(authModel)
	authDevice := models.Device{
		DeviceName: authModel.DeviceName,
		DeviceMac: authModel.DeviceMac,
	}

	context := GetContext()
	if device,err := data.Authenticate(authDevice,context.RethinkSession); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid device",
			401,
		)
		return
	} else { // successful device
		token, err = common.GenerateJWT(device.DeviceName,device.DeviceMac,"admin")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type","application/json")
		authorizedDevice := AuthDeviceModel{
			Device: device,
			Token:token,
		}

		j, err := json.Marshal(AuthDeviceResource{Data: authorizedDevice})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Unexpected error",
				500,
			)
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
