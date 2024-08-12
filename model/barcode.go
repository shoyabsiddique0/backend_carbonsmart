package model

type Barcode struct {
	BarcodeId string `json:"barcode_id" bson:"barcode_id"`
	Ean       string `json:"ean" bson:"ean"`
	Name      string `json:"name" bson:"name"`
	ImageLink string `json:"image_link" bson:"image_link"`
	UserId    string `json:"user_id" bson:"user_id"`
}
