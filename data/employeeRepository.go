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
	"encoding/base64"
	"image"
	"bytes"
	"os"
	"image/jpeg"
	"bufio"
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

func UpdateEmployee(employee models.Employee, session *r.Session) (e models.Employee,err error) {
	jsonEmployee, err := json.Marshal(employee)

	if err != nil {
		return models.Employee{}, err
	}

	var mapEmployee map[string]interface{}
	err = json.Unmarshal(jsonEmployee,&mapEmployee)
	if err != nil {
		return models.Employee{}, err
	}
	cursor, err := r.Table(common.EmployeeDbStruct.Table).Run(session)
	defer cursor.Close()
	var row map[string]interface{}
	var id string
	for cursor.Next(&row) {
		if t, ok := row["email"].(string); ok {
			if t == employee.Email {
				id = row["id"].(string)
				break
			}
		}
	}

	cursor, err = r.Table(common.EmployeeDbStruct.Table).Get(id).Update(mapEmployee).Run(session)
	defer cursor.Close()
	if err != nil {
		return models.Employee{}, err
	}

	return employee,nil
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
			fmt.Printf(" beg : %d , end %d",employeeRow.AuthStarting,employeeRow.AuthEnding)
			if CompareAuthorizationTime(employeeRow.AuthStarting,employeeRow.AuthEnding) == false {
				return models.Employee{}, errors.New("Not Authorized for now")
			}
			if employeeRow.Authorized == false {
				return models.Employee{}, errors.New("Not authorized")
			}
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

func GetAllEmployee(session *r.Session) (e []models.Employee, err error){
	cursor, err := r.Table(common.EmployeeDbStruct.Table).Run(session)
	if err != nil {
		return nil,err
	}
	defer cursor.Close()

	var row map[string]interface{}

	for cursor.Next(&row) {
		jsonEmployee, _ := json.Marshal(row)
		var employeeRow models.Employee
		json.Unmarshal(jsonEmployee,&employeeRow)
		e = append(e, employeeRow)
	}

	return e,nil
}

func SaveEmployeeImage(imageface models.ImageFace) error {

	data, err := base64.StdEncoding.DecodeString(imageface.Data)
	if err != nil {
		return err
	}

	img,_,_ := image.Decode(bytes.NewReader(data))
	path := common.GetTrainingPath()
	path += imageface.Name
	err = os.Mkdir(path,os.FileMode(0777))
	path += "/"
	path += imageface.Filename
	out, err := os.Create(path)

	if err != nil  {
		return err
	}

	err = jpeg.Encode(out, img, nil)

	return err

}

func ReadLogFile() (str []string,err error) {
	file, err := os.Open("./employeesLogs")
	if err != nil {
		return nil,err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return str,nil
}



