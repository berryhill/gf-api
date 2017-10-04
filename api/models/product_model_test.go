package models

import (
	"testing"
	"errors"

	"github.com/berryhill/gf-api/api/db"

	. "github.com/smartystreets/goconvey/convey"
)

const seededProductName = "ShifterSifter"

func createTestProduct() *Product {
	product := NewProduct()
	product.Type = "fly-rods"
	product.Brand = "mattys"
	product.Name = "ShifterSifter"
	product.Title = "ShifterSifter Fly Rod weight 4-6"
	product.Price = "199"
	product.Url = "http://tester.url.com"
	product.Retailer = "Randys Place"
	product.Image = "http://tester.url.com?image=asdfasdfasdfadf"

	return product
}

func TestProduct_FindProductByName(t *testing.T) {

	Convey("Given a Test DB and a Test Product", t, func() {

		db.Connect()
		So(db.Conf.MongoDBHosts, ShouldEqual, "172.17.0.1:27017")

		test_product := createTestProduct()
		So(test_product, ShouldNotBeNil)

		err := test_product.Create("products")
		So(err, ShouldBeNil)

		Convey("When finding Product by Name that exists", func() {

			product_found, err := FindProductByName(
				seededProductName, "products")
			So(err, ShouldBeNil)

			Convey("Then Retrieved Product should equal Test " +
				"Product", func() {

				So(product_found.Name, ShouldResemble, test_product.Name)
				So(product_found.Brand, ShouldResemble, test_product.Brand)
				So(product_found.Image, ShouldResemble, test_product.Image)
			})
		})

		Convey("When finding Product by Name that doesn't exist",
			func() {

				_, err := FindProductByName(
				"random_name", "products")
				So(err, ShouldResemble, errors.New("not found"))
		})

		err = test_product.Delete("products")
		So(err, ShouldBeNil)
	})
}

func TestProduct_Create(t *testing.T) {

	Convey("Given a Test Product", t, func() {
		test_product := createTestProduct()
		So(test_product, ShouldNotBeNil)

		Convey("When the Test Product is saved", func() {
			err := test_product.Create("products")
			So(err, ShouldBeNil)

			Convey("Then the Test Product should be in the DB", func() {
				retrieved_product := new(Product)
				So(retrieved_product, ShouldNotBeNil)

				retrieved_product, err := FindProductByName(
					seededProductName, "products")
				So(err, ShouldBeNil)
				So(retrieved_product.Name, ShouldResemble, test_product.Name)
				So(retrieved_product.Brand, ShouldResemble, test_product.Brand)
				So(retrieved_product.Image, ShouldResemble, test_product.Image)

				err = test_product.Delete("products")
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestProduct_Delete(t *testing.T) {

	Convey("Given a Test DB with a Test Product", t, func() {
		test_product := createTestProduct()
		So(test_product, ShouldNotBeNil)

		err := test_product.Create("products")
		So(err, ShouldBeNil)

		Convey("When the Test Product is Deleted", func() {
			err := test_product.Delete("products")
			So(err, ShouldBeNil)

			Convey("Then the Test Product should be in the DB", func() {

				_, err := FindProductByName(
					seededProductName, "products")
				So(err, ShouldResemble, errors.New("not found"))
			})
		})
	})
}

func TestProduct_Handle(t *testing.T) {

	Convey("Given a Test DB with a Test Product", t, func() {
		test_product := createTestProduct()
		So(test_product, ShouldNotBeNil)

		err := test_product.Create("products")
		So(err, ShouldBeNil)

		Convey("When the same Test Product 'handled'", func() {
			found, _ := test_product.Handle(
				test_product.Name,
				test_product.Title,
				test_product.Brand,
				test_product.Url,
				"products",
			)
			//So(err, ShouldNotBeNil)

			Convey("Then the Handled Test Product should not be in DB",
				func() {

					So(found, ShouldBeTrue)
			})
		})

		Convey("When the Test Product 'handled' that doesn't exist in DB",
			func() {
				err = test_product.Delete("products")
				So(err, ShouldBeNil)

				found, _ := test_product.Handle(
					test_product.Name,
					test_product.Title,
					test_product.Brand,
					test_product.Url,
					"products",
				)
			//So(err, ShouldBeNil)

			Convey("Then the Handled Test Product should not be in DB",
				func() {

					So(found, ShouldBeFalse)
				})
		})

		err = test_product.Delete("products")
		So(err, ShouldBeNil)
	})
}
