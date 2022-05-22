package helper

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func Dprintln(m string) {
	if os.Getenv("CPM_DEV_MODE") == "true" {
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		output := fmt.Sprintf("%s:%d %s - %s", frame.File, frame.Line, frame.Function, m)
		log.Println(output)
	} else {
		return
	}

}

func DprintlnSlice(m []string) {
	if os.Getenv("CPM_DEV_MODE") == "true" {
		output := fmt.Sprintf("%#v\n", m)
		Dprintln(output)
	} else {
		return
	}

}
