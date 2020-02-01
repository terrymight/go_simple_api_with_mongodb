package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/terrymight/go_simple_api_db/routers"

	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	routers.BasicRoutes(router)
	fmt.Println("starting...")
	http.ListenAndServe(":2000", router)
}