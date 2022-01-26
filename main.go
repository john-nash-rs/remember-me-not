package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

type Project struct {
	Id             string
	Name           string
	Directory      string
	RunInstruction string
	Href           string
}

type ProjectList struct {
	Projects []*Project
}

func loadConfig() ProjectList {
	var projects ProjectList
	configFile, err := ioutil.ReadFile("./config/project_information.json")
	check(err)
	json.Unmarshal(configFile, &projects)
	return projects

}

func home(writer http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(writer, projects)
}

func run(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	scriptFileName := "./scripts/" + id + ".sh"
	out, err := exec.Command("/bin/sh", scriptFileName).Output()
	check(err)
	writer.Write([]byte(projectMap[id].Name + " has been executed successfully"))
	fmt.Println(string(out))
}

var projects ProjectList

func main() {
	projects = loadConfig()
	generateScripts()
	http.HandleFunc("/", home)
	http.HandleFunc("/run", run)
	http.ListenAndServe(":1010", nil)
}

var scriptMap map[string]string
var projectMap map[string]Project

func generateScripts() {
	scriptMap = make(map[string]string)
	projectMap = make(map[string]Project)
	for index, project := range projects.Projects {
		project.Href = "http://localhost:1010/run?id=" + project.Id
		scriptFileName := "./scripts/" + project.Id + ".sh"
		scriptFile, err := os.Create(scriptFileName)
		script := []byte("cd " + project.Directory + "\n" + project.RunInstruction)
		err = ioutil.WriteFile(scriptFileName, []byte(script), 0777)
		check(err)
		scriptMap[project.Id] = scriptFile.Name()
		projectMap[project.Id] = *project
		defer scriptFile.Close()
		fmt.Println(index, project.Directory, scriptMap)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
