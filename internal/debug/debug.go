package debug

import (
	"encoding/json"
	"fmt"
	"os"

	data "github.com/Tyler-Arciniaga/DevLog/internal/data"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func AddDebug(title string, tag string) {
	currLog, logPath, lastID := GetCurrDebugLog()

	//create a new debug task object
	newDebugNote := &data.DebugData{
		Id:     lastID + 1,
		Title:  title,
		Tag:    tag,
		Caught: false,
	}

	//create an updated debug log with new task appended on to list of old tasks
	updatedDebugLog := []data.DebugData{*newDebugNote}
	updatedDebugLog = append(updatedDebugLog, currLog...)

	//marshall Go styled slice back into JSON
	packagedLog, e := json.MarshalIndent(updatedDebugLog, " ", " ")
	checkErr(e)

	//overwrite previous debug log with newly updated version
	e = os.WriteFile(logPath, packagedLog, 0644)
	checkErr(e)

	fmt.Println("New debug task added to your debug log!")

}

func GetCurrDebugLog() ([]data.DebugData, string, int) {
	//read and store json content from log.json
	dir, _ := os.Getwd()
	logPath := dir + "/.devlog/debug.json"

	dat, e := os.ReadFile(logPath)
	checkErr(e)

	//unmarshal json into Go styled slice
	var currentLog []data.DebugData
	e = json.Unmarshal(dat, &currentLog)
	checkErr(e)

	//get id of most recent debug task in current log
	var lastId int
	if len(currentLog) == 0 {
		lastId = 0
	} else {
		lastId = currentLog[0].Id
	}

	return currentLog, logPath, lastId
}
