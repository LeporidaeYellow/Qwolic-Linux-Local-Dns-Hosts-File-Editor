package web

import (
	"hostsEditor/config"
	"hostsEditor/enums"
	"hostsEditor/service/editor"
	"hostsEditor/types"
	"html/template"
	"log"
	"net/http"
)

const localHostString string = ""

func Home(response http.ResponseWriter, request *http.Request) {
	ts, err := template.ParseFiles("web/view/templates/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(response, types.RenderData{HostLines: editor.GetLines(config.VarFileHost)})
	if err != nil {
		log.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func FileEditionByAdding(response http.ResponseWriter, request *http.Request) {
	ts, err := template.ParseFiles("web/view/templates/html/form.edit.file.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	if request.Method != http.MethodPost {
		ts.Execute(response, types.RenderData{Action: "Adding"})
		return
	}

	// if want to add only localhost resolving domain names
	// then use var localHostString with value `127.0.0.1`
	// and send to web form only domain name
	details := localHostString + request.FormValue("expression")
	ts.Execute(response, types.RenderData{Success: true, Action: "Adding", Data: details})
	editor.MakeFileByDeletionOrAddition(config.VarFileHost, details, enums.ActionAdd)
}

func FileEditionByRemoving(response http.ResponseWriter, request *http.Request) {
	ts, err := template.ParseFiles("web/view/templates/html/form.edit.file.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	if request.Method != http.MethodPost {
		ts.Execute(response, types.RenderData{Action: "Deleting", Data: editor.GetLinesForHtmlView(config.VarFileHost)})
		return
	}

	details := request.FormValue("expression")
	ts.Execute(response, types.RenderData{Success: true, Action: "Deleting", Data: details})
	editor.MakeFileByDeletionOrAddition(config.VarFileHost, details, enums.ActionRemove)
}
