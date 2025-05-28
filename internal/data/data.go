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
	Caught bool   `json:"caught"`
}

type ExportData struct {
	StartDate int64
	Author    string
	Notes     []NoteData
	DebugList []DebugData
}
