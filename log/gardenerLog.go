package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LogSkills interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type LogBase struct {
	FilePath  []string
	LogChanel map[string]chan string
}

var GardenerLog LogSkills
var logBase LogBase

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}
	dir = strings.Replace(dir, "\\", "/", -1)

	logBase.LogChanel = make(map[string]chan string)
	GardenerLog = logBase
	go logBase.initLogFile(dir+"/debug.log", "debug")
	go logBase.initLogFile(dir+"/info.log", "info")
	go logBase.initLogFile(dir+"/error.log", "error")
}

func (p LogBase) initLogFile(path string, index string) {
	// openFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 777)
	// defer openFile.Close()
	// if err != nil {

	// }
	logChan := make(chan string, 100)
	p.LogChanel[index] = logChan

	for {
		logData := <-logChan
		fmt.Println(logData)
	}
}

func (p LogBase) Debug(format string, args ...interface{}) {
	debugStr := logStr(format, args...)
	p.LogChanel["debug"] <- debugStr
}
func (p LogBase) Info(format string, args ...interface{}) {
	infoStr := logStr(format, args...)
	p.LogChanel["info"] <- infoStr
}
func (p LogBase) Error(format string, args ...interface{}) {
	errorStr := logStr(format, args...)
	p.LogChanel["error"] <- errorStr
}

func logStr(format string, args ...interface{}) string {
	end := len(format)
	var buffer []byte
	paramPos := 0
	for i := 0; i < end; {
		lasti := i
		for format[i] != '{' {
			i++
		}

		if i+1 < end && format[i+1] == '}' {
			buffer = append(buffer, format[lasti:i]...)
			buffer = appendArg(buffer, paramPos, args)
			i += 2
		}
	}

	return string(buffer)
}

func appendArg(buffer []byte, pos int, args []interface{}) []byte {
	if pos >= len(args) {
		buffer = append(buffer, "{}"...)
		return buffer
	}

	arg := args[pos]

	if arg == nil {
		buffer = append(buffer, "nil"...)
		return buffer
	}

	fmt.Printf("v1 type:%T\n", arg)
	var temp string
	switch f := arg.(type) {
	case bool:
		if f {
			buffer = append(buffer, "true"...)
		} else {
			buffer = append(buffer, "false"...)
		}
	case float32:
		temp = strconv.FormatFloat(float64(f), 'E', -1, 32)
		buffer = append(buffer, temp...)
	case float64:
		temp = strconv.FormatFloat(f, 'E', -1, 64)
		buffer = append(buffer, temp...)
	case complex64:
	case complex128:
	case int:
		temp = strconv.FormatInt(int64(f), 10)
		buffer = append(buffer, temp...)
	case int8:
		temp = strconv.FormatInt(int64(f), 10)
		buffer = append(buffer, temp...)
	case int16:
		temp = strconv.FormatInt(int64(f), 10)
		buffer = append(buffer, temp...)
	case int32:
		temp = strconv.FormatInt(int64(f), 10)
		buffer = append(buffer, temp...)
	case int64:
		temp = strconv.FormatInt(int64(f), 10)
		buffer = append(buffer, temp...)
	case uint:
		temp = strconv.FormatUint(uint64(f), 10)
		buffer = append(buffer, temp...)
	case uint8:
		temp = strconv.FormatUint(uint64(f), 10)
		buffer = append(buffer, temp...)
	case uint16:
		temp = strconv.FormatUint(uint64(f), 10)
		buffer = append(buffer, temp...)
	case uint32:
		temp = strconv.FormatUint(uint64(f), 10)
		buffer = append(buffer, temp...)
	case uint64:
		temp = strconv.FormatUint(uint64(f), 10)
		buffer = append(buffer, temp...)
	case uintptr:
	case string:
		buffer = append(buffer, f...)
	case []byte:
		for _, c := range f {
			buffer = append(buffer, c)
		}
	}
	return buffer
}
