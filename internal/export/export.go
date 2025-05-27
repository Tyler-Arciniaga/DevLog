package export

import (
	"os"
	"strings"
	"time"

	"text/template"

	debug "github.com/Tyler-Arciniaga/DevLog/internal/debug"
	notes "github.com/Tyler-Arciniaga/DevLog/internal/note"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Export(format string, timeStyle string) {
	//extract both note and debug json files
	currLog, _ := notes.GetCurrLog()
	debugLog, _, _ := debug.GetCurrDebugLog()

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
	newMD = template.Must(newMD.Parse("{{range .}}- [{{.Timestamp | formatTimeStamp}}] {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}\n{{end}}\n"))
	newMD.Execute(&newExport, currLog)

	//header of debug tasks
	sLine = "## 🐛 Debug Tasks\n"
	newExport.WriteString(sLine)

	//loop through debug log and parse accordingly
	newMD = template.Must(newMD.Parse("{{range .}}- {{if .Caught -}} ✅{{else}}❌{{end}} {{.Title}} {{if .Tag -}}({{.Tag}}){{end}}\n{{end}}"))
	newMD.Execute(&newExport, debugLog)

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
