package usecase

import (
	"backend/mod/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type BarcodeService struct {
	MongoCollection *mongo.Collection
}

func (svc *BarcodeService) GetDetailsHandler(w http.ResponseWriter, r *http.Request) {
	repo := repository.BarcodeRepo{MongoCollection: svc.MongoCollection}
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	id := params["id"]
	user_id := params["user_id"]
	result, err := repo.FindBarcodeByID(id)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Data = result
		return
	}
	fmt.Print(id)
	result, err = repo.FetchBarcodeData(id)
	if err != nil {
		res.Error = err.Error()
		return
	}
	fmt.Println(result)
	if result.Name == "Product Not Found" {
		res.Error = result.Name
		return
	}
	result.UserId = user_id
	inserted, err := repo.InsertBarcodeScanHistory(result)
	if err != nil {
		res.Error = err.Error()
		return
	}
	fmt.Println(inserted)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (svc *BarcodeService) GetProductHistoryHandler(w http.ResponseWriter, req *http.Request) {
	repo := repository.BarcodeRepo{MongoCollection: svc.MongoCollection}
	params := mux.Vars(req)
	id := params["user_id"]
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	history, err := repo.FindBarcodeByUserID(id)
	if err != nil {
		res.Error = err.Error()
		return
	}
	res.Data = history
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
