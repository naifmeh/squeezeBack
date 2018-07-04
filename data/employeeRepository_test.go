package data

import (
	"testing"
	r "gopkg.in/gorethink/gorethink.v4"
	"time"
	"fmt"
	"strings"
)
func TestRemoveEmployee(t *testing.T) {
	mock := r.NewMock()
	mock.On(r.Table("employees")).Return([]interface{}{
		map[string]interface{}{
			"id":           "abcefg5963",
			"firstName":    "Naif",
			"lastName":     "Mehanna",
			"email":        "naif.meh@gmail.com",
			"registeredOn": time.Now().String(),
			"authorized":   "true",
			"frequency":    "0",
			"authStarting": "1554521050",
			"authEnding":   "589852547881"},
	}, nil)

	cursor, err := r.Table("employees").Run(mock)
	if err != nil {
		t.Errorf("err is %v", err)
	}

	var row map[string]interface{}
	var employeeId string
	employeeTest := "naif-mehanna"
	for cursor.Next(&row) {
		fmt.Println(row)
		employee := fmt.Sprintf("%s-%s", row["firstName"], row["lastName"])
		if strings.ToLower(employee) == strings.ToLower(employeeTest) {
			if t, ok := row["id"].(string); ok {
				employeeId = t
			}
			fmt.Println(employeeId)
		}
	}


	mock.AssertExpectations(t)
}

func TestUpdateEmployee(t *testing.T) {
	mock := r.NewMock()
	mock.On(r.Table("employees")).Return([]interface{}{
		map[string]interface{}{
			"id":           "abcefg5963",
			"firstName":    "Naif",
			"lastName":     "Mehanna",
			"email":        "naif.meh@gmail.com",
			"registeredOn": time.Now().String(),
			"authorized":   "true",
			"frequency":    "0",
			"authStarting": "1554521050",
			"authEnding":   "589852547881"},
	}, nil)

	/*employeeId := "abcefg5963"
	cursor, err := r.Table("employees").Get(employeeId).Update("{ \"firstName\":\"Naifo\"}").Run(mock)
	if err != nil {
		t.Errorf("err is %v", err)
	}

	var row map[string]interface{}
	for cursor.Next(&row) {
		fmt.Print(row)
	}*/
	//mock.AssertExpectations(t)
}

func TestGetAllEmployee(t *testing.T) {
	mock := r.NewMock()
	mock.On(r.Table("employees")).Return([]interface{}{
		map[string]interface{}{
			"id":           "abcefg5963",
			"firstName":    "Naif",
			"lastName":     "Mehanna",
			"email":        "naif.meh@gmail.com",
			"registeredOn": time.Now().String(),
			"authorized":   "true",
			"frequency":    "0",
			"authStarting": "1554521050",
			"authEnding":   "589852547881"},
		map[string]interface{}{
			"id":           "abcefg5963de",
			"firstName":    "Naifs",
			"lastName":     "Solotov",
			"email":        "naif.meh@gmail.com",
			"registeredOn": time.Now().String(),
			"authorized":   "true",
			"frequency":    "0",
			"authStarting": "1554521050",
			"authEnding":   "589852547881"},
	},nil)

	cursor, err := r.Table("employees").Run(mock)
	defer cursor.Close()
	if err != nil {
		t.Errorf("Could not load : %v", err)
	}
	var row map[string]interface{}

	for cursor.Next(&row) {
		fmt.Println(row)
	}

	mock.AssertExpectations(t)
}