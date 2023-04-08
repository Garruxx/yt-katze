package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Errorf(format string, a ...any) error {
	
	t := time.Now()
	Err := fmt.Errorf(format, a...)
	pcs := make([]uintptr, 100)
	n := runtime.Callers(1, pcs)

	calls := "Call stack:\n"
	for i := 0; i < n; i++ {
		funcPtr := runtime.FuncForPC(pcs[i])
		file, line := funcPtr.FileLine(pcs[i])
		calls += fmt.Sprintf("%s:%d %s\n\n -- ", file, line, funcPtr.Name())
	}

	err := os.MkdirAll("errors", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating log directory:", err)
		return err
	}

	f, err := os.OpenFile("errors/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()

	message := fmt.Sprintf("[%s] %s:%s \n", t.Format("2006-01-02 15:04:05"), calls, Err)
	if _, err = f.WriteString(message); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return Err
}
