package debug

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	data "github.com/Tyler-Arciniaga/DevLog/internal/data"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func DebugAdd(title string, tag string) {
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

func DebugList(tag string) {
	currLog, _, _ := GetCurrDebugLog()
	fmt.Print("Debug Tracker 🐞\n")
	fmt.Printf("'Just go out and kill a few beasts...' - Gehrman\n\n")

	fmt.Println("ID   	Status  Tag		Task")
	fmt.Println("---  	------  ----	------------------")

	if tag == "" {
		for _, task := range currLog {
			FormatTask(task)
		}
	} else {
		for _, task := range currLog {
			if strings.EqualFold(tag, task.Tag) {
				FormatTask(task)
			}
		}
	}

}

func DebugSquash(bugID string) {
	//extract debug log from debug.json
	currLog, logPath, _ := GetCurrDebugLog()

	//locate index of bug with specified id
	var indexOfBug int
	for i, task := range currLog {
		if strconv.Itoa(task.Id) == bugID {
			indexOfBug = i
			break
		}
	}
	currLog[indexOfBug].Caught = true

	//repackage debug log into json
	packagedLog, e := json.MarshalIndent(currLog, " ", " ")
	checkErr(e)

	//overwrite updated debug log into debug.json file
	e = os.WriteFile(logPath, packagedLog, 0644)
	checkErr(e)

	fmt.Printf("Bug %s squashed!\n", bugID)

}

func FormatTask(task data.DebugData) {
	var status string
	if !task.Caught {
		status = "❌"
	} else {
		status = "✅"
	}
	var tag string
	if task.Tag == "" {
		tag = " NA"
	} else {
		tag = task.Tag
	}

	fmt.Printf(" %d  	 [%s]	%s	%s\n", task.Id, status, tag, task.Title)
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
