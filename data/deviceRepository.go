package data

import (
	"squeezecnn/models"
	r "gopkg.in/gorethink/gorethink.v4"
	"squeezecnn/common"
	"errors"
	"encoding/json"
	"log"
	"strings"
)

func RegisterDevice(device *models.Device, session *r.Session) error {
	jsonDevice, erreur := json.Marshal(device)
	if erreur != nil {
		panic(erreur)
	}
	var mapDevice map[string]interface{}
	err2 := json.Unmarshal(jsonDevice,&mapDevice)
	if err2 != nil {
		panic(err2)
	}
	_,err := r.Table(common.DeviceDbStruct.Table).Insert(r.Expr(mapDevice)).Run(session)

	if err != nil {
		log.Fatalf("[RegisterDevice] : %v",err)
		return err
	}
	return nil
}

func Authenticate(device models.Device, session *r.Session) (d models.Device, err error) {
	res,err := r.Table(common.DeviceDbStruct.Table).Run(session)
	if err != nil {
		return models.Device{},err
	}
	defer res.Close()
	var row map[string]interface{}
	i := 0
	for res.Next(&row) {
		i++
		log.Printf("[Authenticate] Debug : %v",row)
		if t,ok := row["DeviceMac"].(string); ok {
			if strings.ToLower(t) == strings.ToLower(device.DeviceMac) {
				return device,nil
			}
		} else {
			return models.Device{},errors.New("deviceMac not parsed")
		}

	}
	if i == 0{
		return models.Device{},errors.New("empty database")
	} else {
		return models.Device{},errors.New("device not found")
	}

}