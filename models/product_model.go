package models

import (
	"fmt"
	"errors"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	"github.com/berryhill/web-scrapper/db"
)


type Product struct {
	ProductId		*bson.ObjectId          `json:"product_id"`
	Active			bool          			`json:"active"`
	Url 			string        			`json:"url"`
	Image 			string        			`json:"image"`
	Type 			string        			`json:"type"`
	Brand			string        			`json:"brand"`
	Name 			string        			`json:"name"`
	Price 			string        			`json:"price"`
	Retailer		string                  `json:"retailer"`
	Details			map[string]string		`json:"details"`
}

func NewProduct() *Product {

	p := new(Product)
	product_id := bson.NewObjectId()
	p.ProductId = &product_id
	p.Active = true

	return p
}

func (p *Product) MarshalJson() ([]byte, error) {

	json, _ := json.Marshal(p)

	return json, nil
}


func (p *Product) Handle(name string, brand string) (found bool, err error) {

	// TODO: Handle Details
	// TODO: Improve product validation with details

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("fly_rods")

	found = false; result := Product{}
	err = collection.Find(
		bson.M{"name": name, "brand": brand}).One(&result)

	// TODO: Need to compare error to "not found"

	if err != nil {
		found = true
		p.create()
	}

	return found, err
}


func (p *Product) create() error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("fly_rods")

	err := collection.Insert(p)
	if err != nil {
		errors.New("Error inserting Product into DB")
	}

	return err
}

func (p *Product) Print() {

	if p.Name == "" {
		return
	}

	fmt.Println(p.Active)
	fmt.Println(p.Url)
	fmt.Println(p.Image)
	fmt.Println(p.Type)
	fmt.Println(p.Brand)
	fmt.Println(p.Name)
	fmt.Println(p.Price)
	fmt.Println()
}

