package init

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func checkExists(name string) {
	_, e := os.Stat(name)
	if e == nil {
		fmt.Printf("%q already exists...", name)
		os.Exit(1)
	}
}

func Init() {
	fmt.Println("Initializing devlog directory...")
	checkExists("./.devlog")

	err := os.Mkdir(".devlog", 0755)
	checkErr(err)

	fmt.Println("Creating log.json...")
	CreateNewFile("log.json")

	fmt.Println("Creating debug.json...")
	CreateNewFile("debug.json")
}

func CreateNewFile(name string) {
	path := "./.devlog/" + name
	checkExists("./devlog/debug.json")
	checkErr(os.WriteFile(path, []byte("[]"), 0644))
}
