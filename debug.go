package dry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"sync"
)

// PrettyPrintAsJSON marshalles input as indented JSON
// and calles fmt.Println with the result.
// If indent arguments are given, they are joined into
// a string and used as JSON line indent.
// If no indet argument is given, two spaces will be used
// to indent JSON lines.
func PrettyPrintAsJSON(input interface{}, indent ...string) error {
	var indentStr string
	if len(indent) == 0 {
		indentStr = "  "
	} else {
		indentStr = strings.Join(indent, "")
	}
	data, err := json.MarshalIndent(input, "", indentStr)
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(data))
	return err
}

// Nop is a dummy function that can be called in source files where
// other debug functions are constantly added and removed.
// That way import "github.com/ungerik/go-quick" won't cause an error when
// no other debug function is currently used.
// Arbitrary objects can be passed as arguments to avoid "declared and not used"
// error messages when commenting code out and in.
// The result is a nil interface{} dummy value.
func Nop(dummiesIn ...interface{}) (dummyOut interface{}) {
	return nil
}

func StackTrace(skipFrames int) string {
	buf := new(bytes.Buffer) // the returned data
	var lastFile string
	for i := 3; ; i++ { // print 3 level of callstack.
		contin := fprintStackTraceLine(i, &lastFile, buf)
		if !contin {
			break
		}
	}
	return buf.String()
}

func StackTraceLine(skipFrames int) string {
	var buf bytes.Buffer
	var lastFile string
	fprintStackTraceLine(skipFrames, &lastFile, &buf)
	return buf.String()
}

// i, 是caller的层级，具体怎么定义的就不知道了
// 试着打印 stacktrace 的信息，成功的话会返回true, lastFile参数有可能会设置为对应的文件名
// 所有的结果会储存到 bytes.Buffer 中
func fprintStackTraceLine(i int, lastFile *string, buf *bytes.Buffer) bool {
	var lines [][]byte

	pc, file, line, ok := runtime.Caller(i)
	if !ok {
		return false
	}

	// Print this much at least.  If we can't find the source, it won't show.
	fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	if file != *lastFile {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return true
		}
		lines = bytes.Split(data, []byte{'\n'})
		*lastFile = file
	}
	line-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	return true
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	sep       = []byte("/")
)

// 从行数获得源代码中的第几行的信息
// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.Trim(lines[n], " \t")
}

// 通过pc(program counter)获得函数名称的信息
// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod

	if period := bytes.LastIndex(name, sep); period >= 0 {
		name = name[period+1:]
	}

	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// DebugMutex wraps a sync.Mutex and adds debug output
type DebugMutex struct {
	m sync.Mutex
}

func (self *DebugMutex) Lock() {
	fmt.Println("Mutex.Lock()\n" + StackTraceLine(3))
	self.m.Lock()
}

func (self *DebugMutex) Unlock() {
	fmt.Println("Mutex.Unlock()\n" + StackTraceLine(3))
	self.m.Unlock()
}

// DebugRWMutex wraps a sync.RWMutex and adds debug output
type DebugRWMutex struct {
	m sync.RWMutex
}

func (self *DebugRWMutex) RLock() {
	fmt.Println("RWMutex.RLock()\n" + StackTraceLine(3))
	self.m.RLock()
}

func (self *DebugRWMutex) RUnlock() {
	fmt.Println("RWMutex.RUnlock()\n" + StackTraceLine(3))
	self.m.RUnlock()
}

func (self *DebugRWMutex) Lock() {
	fmt.Println("RWMutex.Lock()\n" + StackTraceLine(3))
	self.m.Lock()
}

func (self *DebugRWMutex) Unlock() {
	fmt.Println("RWMutex.Unlock()\n" + StackTraceLine(3))
	self.m.Unlock()
}

func (self *DebugRWMutex) RLocker() sync.Locker {
	return self.m.RLocker()
}
