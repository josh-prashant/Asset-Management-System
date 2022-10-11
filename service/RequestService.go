package service

import (
	"com/josh/asset/db"
	"errors"
)

type RequestService interface {
	CreateRequest()
	UpdateRequestStatus() (bool, error)
	CheckRequestStatus(requestId int) Request //optional
	GetAllRequestsByEmpId(empId int) []Request
	FetchAllRequests(status Status) []Request //admin
}

type Status string

const (
	ACCEPT  Status = "ACCEPT"
	PENDING Status = "PENDING"
	REJECT  Status = "REJECT"
)

func (req Request) CreateRequest() error {
	db := db.GetDB()
	err := db.Create(req).Error
	if err != nil {
		return errors.New("Request failed")
	}
	return nil

}

// need to improve
func (input Request) UpdateRequestStatus() (bool, error) {
	db := db.GetDB()
	result := db.Model(Request{}).Where("req_id = ?", input.ReqId).Update("status", input.Status)

	if result.Error != nil {
		return false, errors.New("Request failed")
	}
	if result.RowsAffected == 0 {
		return false, errors.New("Status not updated")

	}
	return true, nil

}
func (req Request) CheckRequestStatus(requestId int) Request {
	return req
}

func GetAllRequestsByEmpId(empId int) ([]Request, error) {
	var requests []Request
	err := db.GetDB().Model(Request{}).Where("emp_id=?", empId).Scan(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func ReadAllRequests() (requests []Request, err error) {
	db := db.GetDB()
	db.Find(&requests)
	return
}

func FetchAllRequests(status Status) (requests []Request, err error) {
	db := db.GetDB()
	err = db.Where("status = ?", status).Scan(&requests).Error
	return
}
