package models

import (
	"fmt"
	"errors"
	"encoding/json"

	"github.com/berryhill/web-scrapper/db"

	"gopkg.in/mgo.v2/bson"
)


type Product struct {
	ProductId		*bson.ObjectId          `json:"product_id"`
	Active			bool          			`json:"active"`
	Url 			string        			`json:"url"`
	Image 			string        			`json:"image"`
	Type 			string        			`json:"type"`
	Brand			string        			`json:"brand"`
	Name 			string        			`json:"name"`
	Title			string                	`json:"title"`
	Price 			string        			`json:"price"`
	Retailer		string                  `json:"retailer"`
	Details			[]string				`json:"details"`
}

func NewProduct() *Product {

	p := new(Product)
	product_id := bson.NewObjectId()
	p.ProductId = &product_id
	p.Active = true

	return p
}


func (p *Product) Create(db_col string) error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C(db_col)

	err := collection.Insert(p)
	if err != nil {
		errors.New("Error inserting product into DB")
	}

	return err
}

func GetAllProducts(product string) (products []*Product, err error) {

	session := db.Session.Clone()
	defer session.Close()

	// TODO: Implement a product collection to check if product exists
	
	collection := session.DB("test").C(product)

	err = collection.Find(nil).All(&products)
	if err != nil {
		// TODO: Log error
	}

	return products, err
}

func (p *Product) MarshalJson() ([]byte, error) {

	json, _ := json.Marshal(p)

	return json, nil
}


func (p *Product) Handle(
	name string, title string, brand string, url string, db_col string) (
	found bool, err error) {

	// TODO: Improve product validation with details

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C(db_col)

	found = false; result := Product{}
	err = collection.Find(bson.M{
			"title": title,
			"brand": brand,
			"url": url}).One(&result)

	// TODO: Need to compare error to "not found"

	if err != nil {
		found = true
		p.Create(db_col)
	} else {
		p.Print()
	}

	return found, err
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

