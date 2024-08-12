package repository

import (
	"backend/mod/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BarcodeRepo struct {
	MongoCollection *mongo.Collection
}

var url string = "https://world.openfoodfacts.org/api/v0/product/"

func (r *BarcodeRepo) FetchBarcodeData(productId string) (*model.Barcode, error) {
	req, err := http.NewRequest(http.MethodGet, url+productId, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	var barcode model.Barcode
	var barcode_json model.BarcodeJson
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(body, &barcode_json)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if barcode_json.Status == 0 {
		barcode.Name = "Product Not Found"
		barcode.Ean = productId
		return &barcode, nil
	}
	barcode.BarcodeId = uuid.NewString()
	barcode.Ean = barcode_json.Code
	barcode.Name = barcode_json.Product.ProductName
	barcode.ImageLink = barcode_json.Product.ImageURL
	return &barcode, nil
}

func (r *BarcodeRepo) InsertBarcodeScanHistory(barcode *model.Barcode) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), barcode)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *BarcodeRepo) FindAllBarcodeScanned() ([]*model.Barcode, error) {
	var barcodes []*model.Barcode
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &barcodes)
	if err != nil {
		return nil, err
	}

	return barcodes, nil
}

func (r *BarcodeRepo) FindBarcodeByID(id string) (*model.Barcode, error) {
	var barcode model.Barcode
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "ean", Value: id}}).Decode(&barcode)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("barcode not found")
	} else if err != nil {
		return nil, err
	}

	return &barcode, nil
}

func (r *BarcodeRepo) FindBarcodeByUserID(id string) (*[]model.Barcode, error) {
	var barcode []model.Barcode
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{{Key: "user_id", Value: id}})
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("barcode not found")
	} else if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &barcode)
	if err != nil {
		return nil, err
	}
	return &barcode, nil
}
