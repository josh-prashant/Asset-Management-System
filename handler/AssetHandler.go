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

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset service.Asset
	err := json.NewDecoder(r.Body).Decode(&asset)
	if err != nil {
		api.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	if asset.Available == 0 {
		asset.Available = asset.TotalCount
	}
	err = asset.Create()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, "Asset created successfully", w)
	fmt.Println("Asset created successfully")
}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["assetId"]
	i, err := strconv.Atoi(empId)
	if err != nil {
		return
	}
	err = service.DeleteAsset(i)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, "Asset deleted successfully", w)
}

func EditAsset(w http.ResponseWriter, r *http.Request) {
	var asset service.Asset
	json.NewDecoder(r.Body).Decode(&asset)

	asset, err := asset.Update()
	if err != nil {
		api.Response(http.StatusInternalServerError, "Update Failed", w)
		return
	}
	res := make(map[string]service.Asset)
	res["Asset updated successfully"] = asset
	api.Response(http.StatusOK, res, w)

}

func ReadAllAsset(rw http.ResponseWriter, req *http.Request) {
	users, err := service.GetAllAssets()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), rw)
		return
	}
	api.Response(http.StatusOK, users, rw)
	fmt.Println("ReadAll asset successful")
}
