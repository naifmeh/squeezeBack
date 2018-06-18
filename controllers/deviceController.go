package controllers

import (
	"net/http"
	"encoding/json"
	"squeezecnn/common"
	"squeezecnn/data"
	"squeezecnn/models"
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

	context := NewContext()
	defer context.Close()
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
	authDevice := models.Device{
		DeviceName: authModel.DeviceName,
		DeviceMac: authModel.DeviceMac,
	}

	context := NewContext()
	defer context.Close()

}
