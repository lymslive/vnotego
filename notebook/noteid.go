package notebook

import "fmt"
import "github.com/lymslive/vnotego/date"

// 日记ID 包括日期、当日序号与私有标记等信息
type NoteID struct {
	date.Date
	Seq     int
	Private bool
}

// 日记ID字符串形式 yyyymmdd_n-
func (self NoteID) String() string {
	var str = fmt.Sprintf("%d_%d", self.IntNum(), self.Seq)
	if self.Private {
		str += "-"
	}
	return str
}

// 下一个日志ID，序号自增
func (self NoteID) Next() NoteID {
	var next NoteID = self
	next.Seq++
	return next
}

// 生成今天的第一个日志ID
func NewNoteID() NoteID {
	return NoteID{date.Today(), 1, false}
}
