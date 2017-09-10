package models

import (
	"fmt"
	"errors"
	"net/url"
	"encoding/json"

	"github.com/berryhill/gf-api/api/db"

	"gopkg.in/mgo.v2/bson"
)


type Product struct {
	ProductId		*bson.ObjectId          `json:"product_id"`
	Name 			string        			`json:"name"`
	Brand			string					`json:"brand"`
	Price			string					`json:"price"`
	StandardPrice	string					`json:"standard_price"`
	ItemIds			[]*bson.ObjectId		`json:"product_ids"`
	items			[]*Item					`json:"items"`
	BestDealId		*bson.ObjectId          `json:"best_deal_id"`
	bestDeal		*Item					`json:"best_deal"`
	PercentDiscount	string					`json:"percent_discount"`
	Managed			bool                	`json:"managed"`
}

func NewProduct(name string, brand string) *Product {

	p := new(Product)
	product_id := bson.NewObjectId()
	p.Name = name
	p.Brand = brand
	p.ProductId = &product_id
	p.Managed = false

	return p
}


func (p *Product) Create() error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("products")

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

	found = false; result := Product{}
	err = collection.Find(bson.M{
			"title": title,
			"brand": brand,
			"url": url}).One(&result)

	// TODO: Need to compare error to "not found"

	if err != nil {
		found = true
		p.Create()
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
	fmt.Println(p.Price)
	fmt.Println(p.StandardPrice)
	fmt.Println(p.PercentDiscount)

	fmt.Println()
}

func GetProducts(
	product_type string, query_params url.Values, page int, per_page int) (
	products []*Product, err error) {

	// TODO: Implement a product collection to check if product exists
	// TODO: Implement kabab case in the URI.. currently '../fly_rods'


	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C("products")

	params_exist := false
	for _ = range query_params {
		params_exist = true
	}

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
