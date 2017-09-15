package models

import (
	"fmt"
	"errors"
	"net/url"
	"encoding/json"

	"github.com/berryhill/gf-api/api/db"

	"gopkg.in/mgo.v2"
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
	Managed			bool                	`json:"managed"`
}

func NewProduct() *Product {

	p := new(Product)
	product_id := bson.NewObjectId()
	p.ProductId = &product_id
	p.Active = true
	p.Managed = false

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

	index := mgo.Index{
		Key: []string{"$text:title", "$text:brand", "$text:url"},
	}

	err = collection.EnsureIndex(index)


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

	fmt.Println(p.Name)
	fmt.Println(p.Brand)
	fmt.Println(p.Type)
	fmt.Println(p.Price)
	fmt.Println(p.Active)
	fmt.Println(p.Url)
	fmt.Println(p.Image)
	fmt.Println()
}

func GetProducts(
	product string, query_params url.Values, page int, per_page int) (
	products []*Product, err error) {

	// TODO: Implement a product collection to check if product exists
	// TODO: Implement kabab case in the URI.. currently '../fly_rods'


	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C(product)

	index := mgo.Index{
		Key: []string{"$text:name", "$text:brand"},
	}

	err = collection.EnsureIndex(index)

	params_exist := false
	for _ = range query_params {
		params_exist = true
	}

	//collection.Find(bson.M{"abc": &bson.RegEx{Pattern: "efg", Options: "i"}})

	if params_exist {
		for key, value := range query_params {
			if key == "search" {
				fmt.Println(value[0])
				err = collection.Find(
					bson.M{
						"$text": bson.M{"$search": value[0]}}).All(&products)
				if err != nil {
					// TODO: Log error
				}
			}
		}
	} else {
		err = collection.Find(nil).All(&products)
		if err != nil {
			// TODO: Log error
		}
	}

	return products, err
}
