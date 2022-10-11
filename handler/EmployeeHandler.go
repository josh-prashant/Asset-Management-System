package handler

import (
	api "com/josh/asset/api"
	"com/josh/asset/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var emp service.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		api.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	err = emp.Create()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, "Employee created successfully", w)
	fmt.Println("Employee created successfully")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empId"]
	i, err := strconv.Atoi(empId)
	if err != nil {
		return
	}
	err = service.Delete(i)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, "Employee deleted successfully", w)
}

func ReadAll(rw http.ResponseWriter, req *http.Request) {
	users, err := service.ReadAll()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), rw)
		return
	}
	api.Response(http.StatusOK, users, rw)
	fmt.Println("ReadAll emp successful")
}

func Update(w http.ResponseWriter, r *http.Request) {
	var emp service.Employee
	json.NewDecoder(r.Body).Decode(&emp)

	emp, err := service.Update(emp)
	if err != nil {
		api.Response(http.StatusInternalServerError, "Update Failed", w)
		return
	}
	res := make(map[string]service.Employee)
	res["Employee updated successfully"] = emp
	api.Response(http.StatusInternalServerError, res, w)
}

// asset
func EmployeeAssetAllocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empId"]
	i, err := strconv.Atoi(empId)
	if err != nil {
		return
	}
	assets, err := service.GetAssetByEmpId(i)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, assets, w)

}

func AssignAsset(w http.ResponseWriter, r *http.Request) {

	var empAsset service.AssetAlllocation
	err := json.NewDecoder(r.Body).Decode(&empAsset)
	if err != nil {
		api.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	if empAsset.Assetid == 0 || empAsset.EmpId == 0 {
		api.Response(http.StatusBadRequest, "Invalid/Incomplete request body", w)
		return
	}
	err = service.AssignAsset(empAsset)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}

	asset, err := service.GetAssetByAssetId(empAsset.Assetid)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	asset.Available = asset.Available - 1
	asset.Allocate = asset.Allocate + 1
	asset.Update()

}
