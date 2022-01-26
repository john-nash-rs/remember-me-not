package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Project struct {
	Name           string
	Directory      string
	RunInstruction string
}

type ProjectList struct {
	Projects []*Project
}

func loadConfig() ProjectList {
	var projectInformation ProjectList
	configFile, err := ioutil.ReadFile("./config/project_information.json")
	fmt.Println(string(configFile))
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(configFile, &projectInformation)
	fmt.Println(projectInformation)
	return projectInformation

}

func home(writer http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	data := ProjectList{
		Projects: []*Project{
			{Name: "Superset", Directory: "/user", RunInstruction: "Docker compose up"},
		},
	}
	tmpl.Execute(writer, data)
}

func main() {
	projectInformation := loadConfig()
	fmt.Println(projectInformation.Projects[0].Name)
	http.HandleFunc("/home", home)
	http.ListenAndServe(":1010", nil)
}
