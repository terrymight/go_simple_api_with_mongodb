package controllers

import (
	"net/http"
	"github.com/terrymight/go_simple_api_db/models"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPeople getPeople lists from the slice
func GetPeople(w http.ResponseWriter, req *http.Request ){	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.People)
}

// GetPerson getGetPerson a user, using ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Oops, ", err)
	}
	for _, searchID := range models.People {
		if searchID.ID == id {
			json.NewEncoder(w).Encode(searchID)
		}
	}
}
// CreatePerson CreateGetPerson a user, using ID
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(req)
	var person models.Person

	// The user ID will be generated by the mongodb
	// id,_ := strconv.Atoi(params["id"])
	// person.ID = id
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		fmt.Println("Oops", err)
	}
	
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)

	// we don`t need this since we are dealing with a database
	//models.People = append(models.People, person) 
	json.NewEncoder(w).Encode(result)	
}

// UpdatePerson updates a person info 
func UpdatePerson(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range models.People {
		if item.ID == id {
			models.People = append(models.People[:index], models.People[index+1:]...)

			var person models.Person

			_ = json.NewDecoder(request.Body).Decode(person)
			person.ID = id
			models.People = append(models.People, person)
			json.NewEncoder(response).Encode(&person)      
			return
		}
	}
	json.NewEncoder(response).Encode(models.People)
}
// DeletePerson Delete from slice
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	for index, person := range models.People {
		if person.ID == id {
			models.People = append(models.People[:index], models.People[index+1:]...)
		}
	}
}