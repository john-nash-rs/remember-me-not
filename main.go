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
	Id              string
	Name            string
	Directory       string
	RunInstruction  string
	StopInstruction string
	HrefRun         string
	HrefStop        string
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

func execute(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	action := req.URL.Query().Get("action")
	var scriptFileName = ""
	if action == "run" {
		scriptFileName = "./scripts/" + id + "_run.sh"
	} else {
		scriptFileName = "./scripts/" + id + "_stop.sh"
	}
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
	http.HandleFunc("/execute", execute)
	http.ListenAndServe(":1010", nil)
}

var projectMap map[string]Project

func generateScripts() {
	projectMap = make(map[string]Project)
	for index, project := range projects.Projects {
		project.HrefRun = "http://localhost:1010/execute?action=run&id=" + project.Id
		project.HrefStop = "http://localhost:1010/execute?action=stop&id=" + project.Id
		scriptFileName := "./scripts/" + project.Id + "_run.sh"
		generate(scriptFileName, project.RunInstruction, project.Directory)
		scriptFileName = "./scripts/" + project.Id + "_stop.sh"
		generate(scriptFileName, project.StopInstruction, project.Directory)
		projectMap[project.Id] = *project
		fmt.Println(index, project.Directory)
	}
}

func generate(scriptFileName string, instruction string, directoryName string) {
	scriptFile, err := os.Create(scriptFileName)
	script := []byte("cd " + directoryName + "\n" + instruction)
	err = ioutil.WriteFile(scriptFileName, []byte(script), 0777)
	defer scriptFile.Close()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
