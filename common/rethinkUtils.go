package common

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
	"time"
)

var session *r.Session

func GetSession() *r.Session {
	if session == nil {
		var err error
		session, err = r.Connect(r.ConnectOpts {
			Address: AppConfig.Host,
			Database : AppConfig.Database,
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
		})
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: AppConfig.Host,
		Database : AppConfig.Database,
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout: 60 * time.Second,
	})

	if err != nil {
		log.Fatalf("[creqteDbSession]: %s \n", err)
	}
}

func addIndexes() {
	session := GetSession()
	defer session.Close()
	errEmployee := r.Table(EmployeeDbStruct.Table).IndexCreate("email").Exec(session)
	if errEmployee != nil {
		log.Fatalln("[addIndexes]: %s\n", errEmployee)
	}

	errDevice := r.Table(DeviceDbStruct.Table).IndexCreate("deviceMac").Exec(session)
	if errDevice != nil {
		log.Fatalln("[addIndexes]: %s\n", errDevice)
	}
}