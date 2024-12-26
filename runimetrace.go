package runtimeTrace

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)
func CallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	return f.Name()
}
func Trace(a int) { 
	fmt.Printf("Enter in to function: %q\n", CallerName(a))

}
func Untrace(a int) { 
	fmt.Printf("Exit from function %q\n", CallerName(a))
}


// openFile opens a file and returns a file pointer.
//This is the first function
func openFile(filename string) (*os.File, error) {
  Trace(1)
      defer Untrace(1)
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    return file, nil
}

// processFile processes the opened file.
func processFile(file *os.File) error {
    // Do some operations on the file.
    Trace(1)
	    defer Untrace(1)
    content, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }
    fmt.Println(string(content))

    return nil
}

// closeFile closes the file.
func closeFile(file *os.File) {
  Trace(1)
      defer Untrace(1)
    if file == nil {
        file.Close()
    }

}

func main() {
  Trace(1)
    fileName := "example.txt"

    file, err := openFile(fileName)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    err = processFile(file)
    if err != nil {
        log.Fatalf("Error processing file: %v", err)
    }

    defer closeFile(file)
    fmt.Println("File processed successfully.")
}