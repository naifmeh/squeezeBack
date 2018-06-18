package data

import (
	"squeezecnn/models"
	r "gopkg.in/gorethink/gorethink.v4"
	"squeezecnn/common"
)

func RegisterDevice(device *models.Device, session *r.Session) error {
	_,err := r.Table(common.DeviceDbStruct.Table).Insert(device,r.InsertOpts{
		Conflict: func(id, oldDoc, newDoc r.Term) interface{} {
			return newDoc.Merge(map[string]interface{}{
				"count": oldDoc.Add(newDoc.Field("count")),
			})
		}, //TODO: A changer
	}).Run(session)

	if err != nil {
		return err
	}
	return nil
}

func Authenticate(device *models.Device, session *r.Session) (d models.Device,err error) {
	return
	//TODO: Stopped at page 200, verifiy auth of device then check
}