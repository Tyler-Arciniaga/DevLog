package data

type DevLogData struct {
	Notes     []NoteData
	DebugList []DebugData
}

type NoteData struct {
	Timestamp int    `json:"timestamp"`
	Title     string `json:"title"`
	Tag       string `json:"tag"`
}

type DebugData struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Tag    string `json:"tag"`
	Status bool   `json:"status"`
}
