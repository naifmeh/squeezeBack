package common

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
)

type (
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}
)

type configuration struct {
	Server, Host, DBUser, DBPwd, Database string
}
type dbconfiguration struct {
	Table string
	Fields []string
}

var EmployeeDbStruct, DeviceDbStruct dbconfiguration
var AppConfig configuration

func initConfig() {
	loadAppConfig()
	loadEmployeeConfig()
	loadDeviceConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig] : %s\n", err)
	}
}

func loadEmployeeConfig() {
	file, err := os.Open("common/db.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadDbConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	EmployeeDbStruct = dbconfiguration{}
	err = decoder.Decode(&EmployeeDbStruct)
	if err != nil {
		log.Fatalf("[loadDbConfig]: %s\n", err)
	}
}

func loadDeviceConfig() {
	file, err := os.Open("common/db.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadDbConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	DeviceDbStruct = dbconfiguration{}
	err = decoder.Decode(&DeviceDbStruct)
	if err != nil {
		log.Fatalf("[loadDbConfig]: %s\n", err)
	}
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s", handlerError)
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}