package note

import "fmt"

func Note(title string, tag string) {
	fmt.Println(title)
	if tag != "" {
		fmt.Println(tag)
	}
}
