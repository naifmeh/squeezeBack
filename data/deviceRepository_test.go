package data

import (
	"testing"
	r "gopkg.in/gorethink/gorethink.v4"
	"fmt"
	"encoding/json"
	"squeezecnn/models"
)
func TestAuthenticate(t *testing.T) {
	mock := r.NewMock()
	mock.On(r.Table("devices")).Return([]interface{}{
		map[string]interface{}{"deviceMac":"A8:D5:E7:FF:E8:E2", "deviceName":"Raspberry"},
		map[string]interface{}{"deviceMac":"A8:D5:A7:FF:C8:E2", "deviceName":"RaspberryPi"},
	}, nil)

	cursor, err:= r.Table("devices").Run(mock)

	if err != nil {
		t.Errorf("err is %v", err)
	}

	var row map[string]interface{}
	for cursor.Next(&row) {
		fmt.Println(row)
		jsonString, _ := json.Marshal(row)
		var device models.Device
		json.Unmarshal(jsonString,&device)
		fmt.Println(device)
	}


	if err !=nil {
		t.Errorf("err is %v",err)
	}



	mock.AssertExpectations(t)
}