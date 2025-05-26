package note

import (
	"testing"
)

func TestNoteAdd(t *testing.T) int {
	got := NoteAdd("d", "c")
	if got {
		return 0
	}
	return 0

}
