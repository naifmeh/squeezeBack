package controllers

import (
	"net/http"
	"encoding/json"
	"squeezecnn/common"
	"squeezecnn/data"
	"time"
)

func RegisterEmployee(w http.ResponseWriter, r *http.Request) {
	var dataResource EmployeeResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}

	employee := dataResource.Data
	employee.RegisteredOn = time.Now()
	context := GetContext()
	err = data.AddEmployee(employee,context.RethinkSession)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Problem adding user",
			500,
		)
		return
	}

	if j, err := json.Marshal(EmployeeResource{Data : employee}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Unexpected error",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}

func AuthorizeEmployee(w http.ResponseWriter, r *http.Request) {
	var dataEmployeeName AuthEmployeeResource

	err := json.NewDecoder(r.Body).Decode(&dataEmployeeName)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid name data",
			500,
		)
		return
	}

	employee := dataEmployeeName.Data
	context := GetContext()

	returnedEmployee, err := data.FindEmployee(employee,context.RethinkSession)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"User either not found or not authorized",
			403,
		)
		return
	}

	if j, err := json.Marshal(EmployeeResource{Data : returnedEmployee}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Unexpected error",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
