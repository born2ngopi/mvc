package Controllers

import (
	"net/http"
)

type ControllersRepo interface{
	Index(w http.ResponseWriter, r *http.Request)	
}

type controller struct {}

func Init() ControllersRepo {
	return &controller{}
}

func (*controller)Index(w http.ResponseWriter, r *http.Request) {
	
}