package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/xiayulin123/GO_MongoDB/controllers"
)

func main(){
	r := httprouter.New()
	url := controllers.NewUserController(getSession())
	r.GET("/user/:id", url.GetUser)
	r.POST("/user", url.CreateUser)
	r.DELETE("/user/:id", url.DeleteUser)

	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil{
		panic(err)
	}
	return session
}