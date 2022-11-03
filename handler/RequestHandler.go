package handler

import (
	api "com/josh/asset/api"
	"com/josh/asset/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var request service.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	request.Status = service.PENDING
	request.Date = time.Now().String()
	err = request.CreateRequest()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, "Asset Request created successfully", w)
	fmt.Println("Asset Request created successfully")
}

// requestId and Response as param
func UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	var request service.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	fmt.Println("In UpdateRequestStatus Request", request)
	update, err := request.UpdateRequestStatus()
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	fmt.Println("Before update")
	if update && request.Status == service.APPROVE {
		updatedReq, err := service.GetRequest(request.ReqId)
		fmt.Println("updated", updatedReq)
		asset, err := service.GetAssetByAssetId(updatedReq.AssetId)
		if err != nil {
			api.Response(http.StatusInternalServerError, err.Error(), w)
			return
		}
		asset.Available = asset.Available - 1
		asset.Allocate = asset.Allocate + 1
		asset.Update()
	}
	if update {
		api.Response(http.StatusOK, "Asset Request updated successfully", w)
		fmt.Println("Asset Request updated successfully")
	}
}

func GetAllRequestsByEmpId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empId"]
	i, err := strconv.Atoi(empId)
	if err != nil {
		return
	}
	requests, err := service.GetAllRequestsByEmpId(i)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, requests, w)
}

func FetchAllRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	if status == "" {
		requests, err := service.ReadAllRequests()
		if err != nil {
			api.Response(http.StatusInternalServerError, err.Error(), w)
			return
		}
		api.Response(http.StatusOK, requests, w)
		return
	} else {
		requests, err := service.FetchAllRequests(service.Status(status))
		if err != nil {
			api.Response(http.StatusInternalServerError, err.Error(), w)
			return
		}
		api.Response(http.StatusOK, requests, w)
	}
}
func GetRequestStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqId := vars["reqId"]
	i, err := strconv.Atoi(reqId)
	if err != nil {
		return
	}
	request, err := service.GetRequest(i)
	if err != nil {
		api.Response(http.StatusInternalServerError, err.Error(), w)
		return
	}
	api.Response(http.StatusOK, request, w)
}
