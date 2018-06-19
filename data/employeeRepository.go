package data

import (
	"squeezecnn/models"
	r "gopkg.in/gorethink/gorethink.v4"
	"squeezecnn/common"
	"log"
	"fmt"
	"strings"
	"errors"
	"encoding/json"
)

func AddEmployee(employee models.Employee,session *r.Session) error {
	jsonEmployee, erreur := json.Marshal(employee)
	if erreur != nil {
		panic(erreur)
	}
	var mapEmployee map[string]interface{}
	erreur = json.Unmarshal(jsonEmployee,&mapEmployee)
	if erreur != nil {
		panic(erreur)
	}

	err := r.Table(common.EmployeeDbStruct.Table).Insert(r.Expr(mapEmployee)).Exec(session)
	return err
}

func FindEmployee(employee models.Employee, session *r.Session) (e models.Employee,err error) {
	cursor, err := r.Table(common.EmployeeDbStruct.Table).Run(session)
	defer cursor.Close()
	if err != nil {
		return models.Employee{},err
	}

	var row map[string]interface{}
	for cursor.Next(&row) {
		log.Print(row)
		nom := fmt.Sprintf("%s-%s",row["firstName"],row["lastName"])
		nomDetected := fmt.Sprintf("%s-%s",employee.FirstName,employee.LastName)
		jsonEmployee,_:= json.Marshal(row)
		var employeeRow models.Employee
		json.Unmarshal(jsonEmployee,&employeeRow)
		if strings.ToLower(nomDetected) == strings.ToLower(nom) {
			return employeeRow,nil
		}
	}
	return models.Employee{},errors.New("Unexpected error")
}

func RemoveEmployee(employee models.Employee, session *r.Session) error {
	cursor,err := r.Table(common.EmployeeDbStruct.Table).Run(session)
	if err != nil {
		return err
	}
	defer cursor.Close()

	var row map[string]interface{}
	var employeeId string
	for cursor.Next(&row) {
		nom := fmt.Sprintf("%s-%s",row["firstName"],row["lastName"])
		nomToDelete := fmt.Sprintf("%s-%s",employee.FirstName,employee.LastName)
		if strings.ToLower(nom) == strings.ToLower(nomToDelete) {
			if t, ok := row["id"].(string); ok {
				employeeId = t
				break
			} else {
				panic(ok)
			}
		}
	}

	/* Deleting user */
	err = r.Table(common.EmployeeDbStruct.Table).Get(employeeId).Delete().Exec(session)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(employee models.Employee, session *r.Session) (e models.Employee,err error) {
 return
}


