package service

import (
	"com/josh/asset/db"
	"errors"
	"fmt"
)

var (
	readAllAssetQuery = `SELECT * FROM Asset;`
	addAssetQuery     = `INSERT INTO  Asset ( Name, TotalCount,Type,Expiry,Available) VALUES (?,?,?,?,?) `
	deleteAssetQuery  = `DELETE FROM Asset WHERE Id = ?;`
	updateAssetQuery  = `update Employee SET Name = ? ,TotalCount=?,Available=?
						,Type = ? ,Allocate = ?, Expiry = ? where Id=?`
	readAssetQuery = `SELECT * FROM Asset WHERE Id = ?;`
)

type AssetService interface {
	Create() error
	Update() (Asset error)
	Delete(assetId int) error
	ReadAll() ([]Asset, error)
}

// AssetImpl
func (asset *Asset) Create() error {
	db := db.GetDB()
	err := db.Create(asset).Error
	if err != nil {
		return err
	}
	fmt.Println("AssetService:Asset created successfully")
	return nil
}

func (asset *Asset) Update() (Asset, error) {
	db := db.GetDB()
	db.Save(asset)
	// res, err := db.Exec(updateEmployeeQuery, asset.Name, asset.TotalCount,
	// 	asset.Available, asset.Type, asset.Allocate, asset.Expiry, asset.Id)
	// if err != nil {
	// 	return Asset{}, err
	// }
	// cnt, err := res.RowsAffected()
	// if err != nil {
	// 	return Asset{}, errors.New("Update query error")
	// }
	// if cnt == 0 {
	// 	return Asset{}, errors.New("Invalid Asset Id")
	// }
	fmt.Println("AssetService:Asset Updated successfully")
	return *asset, nil
}

func DeleteAsset(assetId int) error {
	db := db.GetDB()

	var asset Asset
	db.First(&asset, assetId)
	db.Delete(&asset)

	// res, err := db.Exec(deleteAssetQuery, assetId)
	// if err != nil {
	// 	return err
	// }
	// cnt, err := res.RowsAffected()
	// if err != nil {
	// 	return errors.New("Error while delete Asset " + strconv.Itoa(assetId))
	// }
	// if cnt == 0 {
	// 	return errors.New("Invalid Asset Id")
	// }
	// fmt.Println("AssetService:Asset Deleted successfully")
	return nil
}

func GetAllAssets() ([]Asset, error) {
	db := db.GetDB()
	var assets []Asset
	err := db.Find(&assets).Error
	// rows, err := db.Query(readAllAssetQuery)
	if err != nil {
		return assets, errors.New("Err to retrive users")
	}

	// defer rows.Close()
	// var asset Asset
	// for rows.Next() {
	// 	err = rows.Scan(&asset.Id, &asset.Name, &asset.TotalCount,
	// 		&asset.Available, &asset.Allocate, &asset.Type, &asset.Expiry)
	// 	if err != nil {
	// 		return assets, errors.New("Err to retrive Assets")
	// 	}
	// 	//  fmt.Println(ur)
	// 	assets = append(assets, asset)
	// }
	fmt.Println(assets)
	return assets, nil
}

func GetAssetByAssetId(assetId int) (Asset, error) {
	var asset Asset
	db := db.GetDB()
	// res, err := db.Query(readAssetQuery, assetId)
	err := db.First(&asset, assetId).Error
	if err != nil {
		return Asset{}, err
	}

	fmt.Println("AssetService:Employee Found ")
	// res.Scan(&asset.Id, &asset.Name, &asset.TotalCount,
	// 	&asset.Available, &asset.Allocate, &asset.Type, &asset.Expiry)
	return asset, nil
}
