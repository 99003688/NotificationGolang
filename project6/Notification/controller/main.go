package main

import (
	// "net/http"
	// serviceimplentation "project4/serviceImplentation"
	"html/template"
	"net/http"
	"project6/Notification/model"
	// services "project6/Notification/services"
)


func main() {
	

	// http.HandleFunc("/mode",mode)
	// http.HandleFunc("/email, )
	// http.HandleFunc("/message", Notify )
	mux:=http.NewServeMux()
	mux.HandleFunc("/", GetAllUsers)

	mode:= model.GetNotificationMode("EmailNotificationMode")
	// http.HandleFunc("/mode",mode)
	// http.ListenAndServe(":3000", nil)
	mode.Notify()
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {

	var UserModel model.UserModel
	user, _ := UserModel.FindAll()
	data := map[string]interface{}{
		"user": user,
	}
	tmp, _ := template.ParseFiles("view/index.html")
	tmp.Execute(response, data)

}


