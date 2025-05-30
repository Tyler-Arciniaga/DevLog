package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	data "github.com/Tyler-Arciniaga/DevLog/internal/data"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Note(arg0 string, arg1 string) {
	if arg0 == "add" {
		fmt.Println("Here")
		NoteAdd(arg0, arg1)
	}
	//Fix this function, right now its pointless
}

func NoteAdd(title string, tag string) {

	currentLog, logPath := GetCurrLog()
	//create new note from command arguements
	newNote := &data.NoteData{
		Timestamp: int(time.Now().Unix()),
		Title:     title,
		Tag:       tag,
	}

	//create new slice with new note
	updatedLog := []data.NoteData{*newNote}

	//append slice with note to current log slice
	updatedLog = append(updatedLog, currentLog...)

	//marshall updated log slice back into JSON
	packagedNotes, e := json.MarshalIndent(updatedLog, " ", " ")
	checkErr(e)

	//write back into log.json file (ovewrites previous log with newly updated log)
	e = os.WriteFile(logPath, packagedNotes, 0644)
	checkErr(e)
	fmt.Println("Note added to your log!")
}

func NoteList(tag string) {
	currentLog, _ := GetCurrLog()

	if tag == "" {
		for _, note := range currentLog {
			FormatNote(note)
		}
	} else {
		for _, note := range currentLog {
			if strings.EqualFold(tag, note.Tag) {
				FormatNote(note)
			}
		}
	}

}

func FormatNote(note data.NoteData) {
	//convert back to local time layout form Unix epoch time
	nonUnixTime := time.Unix(int64(note.Timestamp), 0)

	//print note information
	fmt.Printf("[%s]\n", nonUnixTime.String()[:19])
	fmt.Printf("	Tag: 	%s\n", note.Tag)
	fmt.Printf("	Title:	%s\n\n", note.Title)
}

func GetCurrLog() ([]data.NoteData, string) {
	//read and store json content from log.json
	dir, _ := os.Getwd()
	logPath := dir + "/.devlog/log.json"

	dat, e := os.ReadFile(logPath)
	checkErr(e)

	//unmarshal json into Go styled slice
	var currentLog []data.NoteData
	e = json.Unmarshal(dat, &currentLog)
	checkErr(e)

	return currentLog, logPath
}
