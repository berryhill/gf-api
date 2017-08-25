package models

import (
	"errors"

	"github.com/berryhill/gf-api/db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)


type Retailer struct {
	Name 			string           	`json:"name"`
	BaseUrl			string        		`json:"base_url"`
	Products		map[string]string   `json:"product_urls"`
}

func NewRetailer(
	name string, base_url string,
	product_url string, product_name string)*Retailer {

	r := new(Retailer)
	r.Name = name
	r.BaseUrl = base_url
	products := make(map[string]string)
	products[product_name] = product_url
	r.Products = products

	return r
}

func (r *Retailer) Create() error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("retailers")

	err := collection.Insert(r)
	if err != nil {
		errors.New("Error inserting Product into DB")
	}

	return err
}

func (r *Retailer) Get(name string) (*Retailer, error) {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("retailers")

	result := Retailer{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return &result, nil
	}

	return &result, nil
}

func (r *Retailer) Print() {

	fmt.Println(r.Name)
	fmt.Println(r.BaseUrl)
	fmt.Println(r.Products)
	fmt.Println()
}
