package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/spf13/viper"
)

type Project struct {
	Name           string
	Directory      string
	RunInstruction string
}

type ProjectList struct {
	Projects []Project
}

func loadConfig() (projectInformation Project, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("project_information")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&projectInformation)
	return

}

func home(writer http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	data := ProjectList{
		Projects: []Project{
			{Name: "Superset", Directory: "/user", RunInstruction: "Docker compose up"},
		},
	}
	tmpl.Execute(writer, data)
}

func main() {
	projectInformation, err := loadConfig()
	fmt.Printf(projectInformation.Name)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	http.HandleFunc("/home", home)
	http.ListenAndServe(":1010", nil)
}
