package controllers

import (
	"net/http"
	"encoding/json"
	"squeezecnn/common"
	"squeezecnn/data"
)

func TrainNetwork(w http.ResponseWriter,r *http.Request) {
	var networkResource NetworkResource

	err := json.NewDecoder(r.Body).Decode(&networkResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid order",
			500,
		)
		return
	}

	network := networkResource.Data
	err = data.TrainNetwork(network)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Could not train network",
			500,
		)
	} else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
	}

}
