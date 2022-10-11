package service

import (
	"com/josh/asset/db"
	"errors"
	"fmt"
)

type EmpAssets struct {
	EmpId     int
	AssetName string
	AssetType string
	AssetId   int
	// Quantity  int
	Desc string
}

type Result struct {
	Name string
	Type string
	Id   int
}

// Admin can access the list of employees and their allocated assets list.

func GetAssetByEmpId(empId int) ([]EmpAssets, error) {
	db := db.GetDB()
	var assets []EmpAssets
	var result []Result
	// db.First(&assets, empId)
	err := db.Raw("select a.name ,a.type,a.id from assets as a where a.id in"+
		"(Select al.assetid from asset_allocation as al where al.emp_id=?)",
		empId).Scan(&result).Error

	if err != nil {
		return assets, errors.New("Err to retrive employee assets")
	}

	for _, ass := range result {
		assets = append(assets, EmpAssets{EmpId: empId, AssetName: ass.Name, AssetType: ass.Type, AssetId: ass.Id})
	}

	fmt.Println(assets)
	return assets, nil
}

func AssignAsset(ass AssetAlllocation) error {
	// addAssetQuery := `INSERT INTO  AssetAllocation ( EmpId, AssetId) VALUES (?,?) `

	db := db.GetDB()
	err := db.Create(ass).Error
	// err := db.QueryRow(addAssetQuery, ass.EmpId, ass.AssetId).Err()
	if err != nil {
		return err
	}
	fmt.Println("AssetAlllocation:Asset assigned to emp successfully")
	return nil

}
