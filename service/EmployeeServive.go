package service

import (
	"com/josh/asset/db"
	"errors"
	"strconv"

	"fmt"
)

var (
	readAllEmployeesQuery = `SELECT * FROM Employee;`
	addEmployeeQuery      = `INSERT INTO  Employee ( FirstName, LastName,Mobile) VALUES (?,?,?) `
	deleteEmployeeQuery   = `DELETE FROM Employee WHERE EmpId = ?;`
	updateEmployeeQuery   = `update Employee SET FirstName = ? ,LastName=?,Mobile=? where EmpId=?`
)

type EmployeeService interface {
	Create() error
	ReadAll() ([]Employee, error)
	Update(emp Employee) (Employee error)
	Delete(empId int) error
}

func (data Employee) Create() error {
	db := db.GetDB()
	// id := 0
	// err := db.QueryRow(addEmployeeQuery, data.FirstName, data.LastName, data.Mobile).Err()
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	fmt.Println("EmployeeService:Employee created successfully", data)
	return nil
}

func ReadAll() ([]Employee, error) {
	db := db.GetDB()
	var users []Employee
	fmt.Println("befor fetch all", db)
	// db.Get(readAllEmployeesQuery)
	db.Find(&users)
	// rows, err := db.Query(readAllEmployeesQuery)
	// if err != nil {
	// 	return users, errors.New("Err to retrive users")
	// }

	// defer rows.Close()
	// var ur Employee
	// for rows.Next() {
	// 	err = rows.Scan(&ur.Id, &ur.FirstName, &ur.LastName, &ur.Mobile)
	// 	if err != nil {
	// 		return users, errors.New("Err to retrive users")
	// 	}
	// 	//  fmt.Println(ur)
	// 	users = append(users, ur)
	// }
	fmt.Println(users)
	return users, nil
}

func Delete(empId int) error {
	db := db.GetDB()

	var emp Employee
	result := db.Where("emp_id = ?", empId).First(&emp).Delete(emp)

	if result.Error != nil {
		return errors.New("Error while delete employee " + strconv.Itoa(empId))
	}
	if result.RowsAffected == 0 {
		return errors.New("Invalid Employee Id")
	}
	fmt.Println("EmployeeService:Employee Deleted successfully")
	return nil
}

// not working
func Update(emp Employee) (Employee, error) {
	db := db.GetDB()

	db = db.Debug().Model(&Employee{}).Where("emp_id = ?", emp.EmpId).Take(&Employee{}).UpdateColumns(
		map[string]interface{}{
			"first_name": emp.FirstName,
			"last_name":  emp.LastName,
			"mobile":     emp.Mobile,
		},
	)

	// res, err := db.Exec(updateEmployeeQuery, emp.FirstName, emp.LastName, emp.Mobile, emp.Id)
	fmt.Println("emp", emp)

	// db.Table("employees").Where("emp_id = ?", emp.EmpId).First(&emp)
	// err := db.Table("employees").Update(emp).Error
	// fmt.Printf(err.Error())
	// if err != nil {
	// 	return Employee{}, err
	// }
	// cnt, err := res.RowsAffected()
	// if err != nil {
	// 	return Employee{}, errors.New("Update query error")
	// }
	// if cnt == 0 {
	// 	return Employee{}, errors.New("Invalid Employee Id")
	// }
	fmt.Println("EmployeeService:Employee Updated successfully")

	return emp, nil
}

func GetEmployeeByEmail(email string) (Employee, error) {
	var emp Employee
	db := db.GetDB()
	err := db.Model(emp).Where("email=?", email).Scan(&emp).Error
	if err != nil {
		return Employee{}, err
	}
	fmt.Println("EmployeeService:Employee Found ")
	return emp, nil
}
