package export

import (
	"os"
	"strings"
	"time"

	"text/template"

	data "github.com/Tyler-Arciniaga/DevLog/internal/data"
	debug "github.com/Tyler-Arciniaga/DevLog/internal/debug"
	notes "github.com/Tyler-Arciniaga/DevLog/internal/note"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Export(format string, sinceDate string) {
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
	sLine = "# Your DevLog (" + time.Now().Format(time.DateOnly) + ")\n## 📝 Notes\n"
	newExport.WriteString(sLine)

	//create func map for template to use in pipelining
	funcMap := template.FuncMap{
		"formatTimeStamp": formatTimeStamp,
	}

	//instantiate new template object
	newMD := template.New("newMD").Funcs(funcMap)

	//loop through notes log and parse accordingly, then write into string builder
	newMD = template.Must(newMD.Parse("{{range .Notes}}{{if gt .Timestamp $.StartDate}}- [{{.Timestamp | formatTimeStamp}}] {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}\n{{end}}{{end}}\n"))
	newMD.Execute(&newExport, exportData)

	//header of debug tasks
	sLine = "## 🐛 Debug Tasks\n"
	newExport.WriteString(sLine)

	//loop through debug log and parse accordingly
	newMD = template.Must(newMD.Parse("{{range .DebugList}}- {{if .Caught -}} ✅{{else}}❌{{end}} {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}\n{{end}}"))
	newMD.Execute(&newExport, exportData)

	//convert final string builder into string type
	FinalExport := newExport.String()

	//create export path and write file out
	//(overwrites any pre-exisitng export briefs made in same day)
	writePath := makeExportPath("md")
	e := os.WriteFile(writePath, []byte(FinalExport), 0644)
	checkErr(e)

}

func makeExportPath(format string) string {
	dir, _ := os.Getwd()
	var end string
	if format == "html" {
		end = ".html"
	} else {
		end = ".md"
	}
	return dir + "/DevLogExport-" + time.Now().Format(time.DateOnly) + end

}

func formatTimeStamp(unixTime int) string {
	return time.Unix(int64(unixTime), 0).String()[:19]
}
