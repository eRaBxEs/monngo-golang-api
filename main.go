package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/user/erabxes/mongo-golang/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	// this function helps to simply connect to mongodb
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
