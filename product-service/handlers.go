package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Price       float64       `bson:"price"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GetProducts(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()

	var products []Product
	collection := session.DB("ecommerce").C("products")
	err := collection.Find(nil).All(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := getSession()
	defer session.Close()

	product.ID = bson.NewObjectId()
	collection := session.DB("ecommerce").C("products")
	err = collection.Insert(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := getSession()
	defer session.Close()

	collection := session.DB("ecommerce").C("products")
	err = collection.UpdateId(product.ID, product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := getSession()
	defer session.Close()

	collection := session.DB("ecommerce").C("products")
	err = collection.RemoveId(product.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
