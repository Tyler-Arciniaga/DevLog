package export

import (
	"os"
	"strings"
	"text/template"
	"time"

	data "github.com/Tyler-Arciniaga/DevLog/internal/data"
	debug "github.com/Tyler-Arciniaga/DevLog/internal/debug"
	notes "github.com/Tyler-Arciniaga/DevLog/internal/note"
)

func ExportHTML(format string, sinceDate string) {
	var exportData data.ExportData
	//extract both note and debug json files
	currLog, _ := notes.GetCurrLog()
	exportData.Notes = currLog

	if sinceDate == "" {
		exportData.StartDate = 0
	} else {
		noEarlier, err := time.Parse("2006-01-02", sinceDate)
		checkErr(err)
		exportData.StartDate = noEarlier.Unix()
	}

	debugLog, _, _ := debug.GetCurrDebugLog()
	exportData.DebugList = debugLog

	//create string builder object and temp string line objects
	var newExport strings.Builder
	var sLine string

	//header of export briefing
	sLine = "<h1> Your DevLog (" + time.Now().Format(time.DateOnly) + ")</h1><h2>📝 Notes</h2>"
	newExport.WriteString(sLine)

	//create func map for template to use in pipelining
	funcMap := template.FuncMap{
		"formatTimeStamp": formatTimeStamp,
	}

	//instantiate new template object
	newMD := template.New("newMD").Funcs(funcMap)

	//loop through notes log and parse accordingly, then write into string builder
	newMD = template.Must(newMD.Parse("{{range .Notes}}{{if gt .Timestamp $.StartDate}}<p1>- [{{.Timestamp | formatTimeStamp}}] {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}<p1><br>{{end}}{{end}}<br>"))
	newMD.Execute(&newExport, exportData)

	//header of debug tasks
	sLine = "<h2>🐛 Debug Tasks</h2>"
	newExport.WriteString(sLine)

	//loop through debug log and parse accordingly
	newMD = template.Must(newMD.Parse("{{range .}}- <p1>{{if .Caught -}} ✅{{else}}❌{{end}} {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}</p1><br>{{end}}<br>"))
	newMD.Execute(&newExport, debugLog)

	//convert final string builder into string type
	FinalExport := newExport.String()

	//create export path and write file out
	//(overwrites any pre-exisitng export briefs made in same day)
	writePath := makeExportPath("html")
	e := os.WriteFile(writePath, []byte(FinalExport), 0644)
	checkErr(e)

}
