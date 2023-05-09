package utils

import(
    "time"
    "fmt"
)

type Timer struct {
	st  *time.Time
	msg string
}

func NewTimer(msg string) *Timer {
	t := &Timer{}
	defaultMsg := "计时器耗时为:"
	if msg == "" {
		t.msg = defaultMsg
	} else {
		t.msg = msg
	}
	return t
}

func (t *Timer) Start() {
	now := time.Now()
	t.st = &now
}

func (t *Timer) End() {
	if t.st == nil {
		fmt.Println("计时器未设置开始时间点")
	} else {
		fmt.Printf("%v%v\n", t.msg, time.Since(*t.st))
	}

}
