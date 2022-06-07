package files

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Exist(filePath string)(bool) {
	_, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func Read(filePath string) ([]string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines

}

func Parser(textLines []string, args Args, params string)([]string) {
  var txt []string
	var lineParsed string
	for _, line := range textLines {
		lineParsed = Replace(line, args, params)
		txt = append(txt, lineParsed)
	}
	return txt
}


func Write(filePath string, lines []string) {
	fmt.Println(filePath)

	file, _ := os.Create(filePath)
 
	// Create a writer
	w := bufio.NewWriter(file)
 
	for _, line := range lines {
		w.WriteString(line + "\n")
	}
 
	// Very important to invoke after writing a large number of lines
	w.Flush()

}

func Replace(line string, args Args, params string)(string) {
	var contentBasePath string
	var contentDomain string
	var contentParams string
	var contentEvent string
	var finalContent string

	contentBasePath = strings.ReplaceAll(line, "{%BASE_PATH%}", args.BasePath)
	contentDomain = strings.ReplaceAll(contentBasePath, "{%DOMAIN%}", args.Domain)
	contentParams = strings.ReplaceAll(contentDomain, "{%PARAMS%}", params)
	contentEvent = strings.ReplaceAll(contentParams, "{%EVENT_ACTION%}", args.EventName)
	finalContent = strings.ReplaceAll(contentEvent, "{%NAME%}", args.Name)
	
	return finalContent
}

func Append(filePath string, content string) {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    if _, err := f.WriteString("\n"+content+"\n"); err != nil {

        log.Fatal(err)
    }
}

func Copy(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}
