package models

import (
	"fmt"
	"errors"
	//"net/url"
	"encoding/json"

	"github.com/berryhill/gf-api/api/db"

	"gopkg.in/mgo.v2/bson"
)


type Item struct {
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

func NewItem() *Item {

	i := new(Item)
	item_id := bson.NewObjectId()
	i.ProductId = &item_id
	i.Active = true
	i.Managed = false

	return i
}


func (i *Item) Create(db_col string) error {

	session := db.Session.Clone()
	defer session.Close()

	collection := session.DB("test").C(db_col)

	err := collection.Insert(i)
	if err != nil {
		errors.New("Error inserting product into DB")
	}

	return err
}

func (i *Item) MarshalJson() ([]byte, error) {

	json, _ := json.Marshal(i)

	return json, nil
}


func (i *Item) Handle(
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
		i.Create(db_col)
	} else {
		i.Print()
	}

	return found, err
}

func (i *Item) Print() {

	if i.Name == "" {
		return
	}

	fmt.Println(i.Name)
	fmt.Println(i.Brand)
	fmt.Println(i.Type)
	fmt.Println(i.Price)
	fmt.Println(i.Active)
	fmt.Println(i.Url)
	fmt.Println(i.Image)
	fmt.Println()
}

//func GetProducts(
//	product string, query_params url.Values, page int, per_page int) (
//	products []*Product, err error) {
//
//	// TODO: Implement a product collection to check if product exists
//	// TODO: Implement kabab case in the URI.. currently '../fly_rods'
//
//
//	session := db.Session.Clone()
//	defer session.Close()
//
//	collection := session.DB("test").C(product)
//
//	params_exist := false
//	for _ = range query_params {
//		params_exist = true
//	}
//
//	//collection.Find(bson.M{"abc": &bson.RegEx{Pattern: "efg", Options: "i"}})
//
//	if params_exist {
//		for key, value := range query_params {
//			if key == "search" {
//				fmt.Println(value[0])
//				err = collection.Find(
//					bson.M{
//						"$text": bson.M{"$search": value[0]}}).All(&products)
//				if err != nil {
//					// TODO: Log error
//				}
//			}
//		}
//	} else {
//		err = collection.Find(nil).All(&products)
//		if err != nil {
//			// TODO: Log error
//		}
//	}
//
//	return products, err
//}
