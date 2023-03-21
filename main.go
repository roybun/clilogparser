package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"regexp"
)

type parseLogfileOps struct {
	Filename     string
	FO           *os.File
	LogLineRegex []string
}

func parseLogFile(ops parseLogfileOps) {
	fileReader, err := os.Open(ops.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()
	reader := bufio.NewReader(fileReader)
	for {
		line, err := reader.ReadBytes('\n')
		linestr := string(line)
		if err != nil && len(linestr) == 0 {
			break
		}
		for _, regex := range ops.LogLineRegex {
			res, err := regexp.MatchString(regex, linestr)
			if err != nil {
				continue
			}
			if res == false {
				continue
			}
			ops.FO.Write([]byte(linestr))
			break
		}
		if err == io.EOF {
			break
		}
	}
}
func combineLog() {
	logFileDir := os.Getenv("LOGFILEDIR")
	logFileRegex := os.Getenv("LOGFILEREGEX")
	logFileOutput := os.Getenv("LOGFILEOUTPUT")
	var ops parseLogfileOps
	logLineRegex, found := os.LookupEnv("LOGLINEREGEX1")
	if found == true {
		ops.LogLineRegex = append(ops.LogLineRegex, logLineRegex)
	}
	logLineRegex, found = os.LookupEnv("LOGLINEREGEX2")
	if found == true {
		ops.LogLineRegex = append(ops.LogLineRegex, logLineRegex)
	}
	logLineRegex, found = os.LookupEnv("LOGLINEREGEX3")
	if found == true {
		ops.LogLineRegex = append(ops.LogLineRegex, logLineRegex)
	}
	if len(ops.LogLineRegex) == 0 {
		log.Fatal("No log line regex found")
	}
	files, err := os.ReadDir(logFileDir)
	if err != nil {
		log.Fatal(err)
	}

	fo, err := os.Create(logFileOutput)
	if err != nil {
		log.Fatal(err)
	}
	ops.FO = fo
	defer fo.Close()
	for _, f := range files {
		res, err := regexp.MatchString(logFileRegex, f.Name())
		if err != nil {
			continue
		}
		if res == true {
			ops.Filename = path.Join(logFileDir, f.Name())
			parseLogFile(ops)
		}
	}

}
func main() {
	combineLog()
	os.Exit(0)
}
