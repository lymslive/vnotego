package notebook

import (
	"fmt"
	"github.com/lymslive/vnotego/date"
	"testing"
	"time"
)

func TestNoteID(t *testing.T) {
	var nid = NewNoteID()
	fmt.Printf("today first noteid: %v\n", nid)
	fmt.Printf("and its type: %T\n", nid)

	year, month, day := time.Now().Date()
	if year != int(nid.Year) || int(month) != int(nid.Month) || day != int(nid.Day) {
		t.Fatal("fail to create NoteID from today")
	}

	date.SepField(date.SEP_PATH)
	days, _ := date.EndDay(int(nid.Year), int(nid.Month))
	for i := 0; i < days; i++ {
		fmt.Printf("%s/%s\n", nid.Date, nid)
		nid = nid.Next()
	}
}
