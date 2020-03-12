package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: $>fc.exe \"sourcepath\" \"destpath\"")
	}

	// write log (https://stackoverflow.com/a/51628140)
	logFilename := fmt.Sprintf("fc_log_%s.log", time.Now().Format("2006-01-02_15-04-05"))
	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error creting logfile: %v", err)
	}
	defer logFile.Close()

	wrt := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(wrt)

	// read command line args
	rootDir := strings.TrimSpace(os.Args[1])
	destRoot := strings.TrimSpace(os.Args[2])

	log.Printf("fc(): Comparing contents of <%s> to <%s>...", rootDir, destRoot)
	// time.Sleep(time.Duration(4) * time.Second)
	log.Println("fc(): starting copy..")

	// files, err := ioutil.ReadDir(rootDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // for each subfolder in the root dir
	// foldercount := 0
	// for _, f := range files {
	// 	// check if this folder exists on the destination


	// 	log.Printf("%d folders counted.\n", foldercount)
	// }
}

// return a list of files from a given location
func getFilesList(loc string) ([]string, error) {
	srcFiles := []string{}
	files, err := ioutil.ReadDir(loc)
	if err != nil {
		log.Printf("getFilesList(): error: %v\n", err)
		return srcFiles, err
	}

	for _, f := range files {
		srcFiles = append(srcFiles, f.Name())
	}

	return srcFiles, nil
}

// return a map of files from a given location (keys are filenames)
func getFilesMap(loc string) (map[string]int, error) {
	destFiles := map[string]int{}
	files, err := ioutil.ReadDir(loc)
	if err != nil {
		log.Printf("getFilesList(): error: %v\n", err)
		return destFiles, err
	}

	for _, f := range files {
		// srcFiles = append(srcFiles, f.Name())
		keyFound, _ := destFiles[f.Name()]
		if destFiles {
			log.Printf("getFilesMap(): destFiles key <%s> already found.\n", )
		}
	}

	return destFiles, nil
}

func compare(srcList []string, destList map[string]int) {

}
