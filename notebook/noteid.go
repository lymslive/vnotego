package notebook

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lymslive/vnotego/date"
)

// 日记ID 包括日期、当日序号与私有标记等信息
type NoteID struct {
	date.Date
	Seq     int
	Private bool
}

const (
	PUBLIC_LABEL  string = "+"
	PRIVATE_LABEL string = "-"

	SEPARAT_DATE string = "_"
)

// 日记ID字符串形式 yyyymmdd_n-
func (self NoteID) String() string {
	var str = fmt.Sprintf("%d_%d", self.IntNum(), self.Seq)
	if self.Private {
		str += PRIVATE_LABEL
	}
	return str
}

// 从一个字符串反构造 NoteID 对象指针
// 如果非法字符串，返回 nil
func Endcode(sid string) (*NoteID, error) {
	// 最后一个字符表示私有
	var private = false
	if sid[len(sid)-1:] == PRIVATE_LABEL {
		private = true
		sid = sid[0 : len(sid)-1]
	}

	// 按下划线 _ 分割成两部分
	var sp = strings.Split(sid, SEPARAT_DATE)
	if len(sp) != 2 {
		return nil, errors.New("bad NoteID format: " + sid)
	}

	sdate, sseq := sp[0], sp[1]
	ndate, err := strconv.Atoi(sdate)
	if err != nil {
		return nil, err
	}
	nseq, err := strconv.Atoi(sseq)
	if err != nil {
		return nil, err
	}

	year, ndate := ndate/10000, ndate%10000
	month, day := ndate/100, ndate%100

	// 构造对象
	var nid = new(NoteID)
	nid.Year = int16(year)
	nid.Month = int8(month)
	nid.Day = int8(day)
	nid.Seq = nseq
	nid.Private = private

	if !nid.Valid() {
		return nil, errors.New(fmt.Sprintf("invalid date: %d-%d-%d", year, month, day))
	}

	return nid, nil
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
