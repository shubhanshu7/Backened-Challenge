package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    string        `bson:"user_id"`
	ProductID string        `bson:"product_id"`
	Quantity  int           `bson:"quantity"`
	Status    string        `bson:"status"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GetOrders(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()

	var orders []Order
	user := r.Context().Value("user").(*Claims)
	collection := session.DB("ecommerce").C("orders")
	err := collection.Find(bson.M{"user_id": user.Username}).All(&orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(*Claims)
	order.UserID = user.Username
	order.ID = bson.NewObjectId()

	session := getSession()
	defer session.Close()

	collection := session.DB("ecommerce").C("orders")
	err = collection.Insert(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
