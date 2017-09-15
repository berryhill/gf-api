package models

import (
	"errors"
	"encoding/json"

	"github.com/berryhill/gf-api/api/db"
)


type ProductType struct {
	Name 			string        `json:"name"`
	//Url 			string        `json:"url"`
	//DbCollection	string    	  `json:"db_collection"`
}

func NewProductType(
	name string, url string, db_collection string) *ProductType {

	pt := new(ProductType)
	pt.Name = name
	//pt.Url = url
	//pt.DbCollection = db_collection

	return pt
}

func (pt *ProductType) Create() error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("product_types")

	err := collection.Insert(pt)
	if err != nil {
		errors.New("Error inserting product-type into DB")
	}

	return err
}
func GetProductTypes() (product_types []*map[string]interface{}, err error) {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("product_types")

	err = collection.Find(nil).All(&product_types)
	if err != nil {
		// TODO: Log error
	}

	return product_types, err
}

func (pt *ProductType) MarshalJson() ([]byte, error) {

	json, _ := json.Marshal(pt)

	return json, nil
}