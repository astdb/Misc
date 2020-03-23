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
		log.Fatal("Usage: $>filecomp.exe \"sourcepath\" \"destpath\"")
	}

	// write log (https://stackoverflow.com/a/51628140)
	logFilename := fmt.Sprintf("filecompare_log_%s.log", time.Now().Format("2006-01-02_15-04-05"))
	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating logfile: %v", err))
	}
	defer logFile.Close()

	wrt := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(wrt)

	// read command line args
	localRoot := strings.TrimSpace(os.Args[1])	// local root directory
	destRoot := strings.TrimSpace(os.Args[2])	// destination root directory

	// get list of files on local root
	log.Printf("Getting list of local root direcotries...")
	rootFolderList, err := getFilesList(localRoot)
	if err != nil {
		log.Fatal(err)
	}

	// get map of folders on destinantion root
	log.Printf("Getting list of destination root directories...")
	destFolderMap, err := getFilesMap(destRoot)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("%v\n", rootFolderList)
	// log.Println("---------------------")
	// log.Printf("%v\n", destFolderMap)


	// compare source/destination lists
	res1 := compare(rootFolderList, destFolderMap)

	if len(res1) != 0 {
		// some folders on source root weren't found on destination root
		log.Println("MISMATCH: The following directories exist on the source root but not in destination root:")
		for _, dirName := range res1 {
			log.Printf("\t%s\n", dirName)
		}

		return
	} else {
		log.Printf("All direcories in source root <%s> were found on destination root <%s>. Comparing individual directory content..\n", localRoot, destRoot)
	}

	// traverse through list of source root folders, and repeat compare
	for _, sourceDir := range rootFolderList {
		var thisDirPath string
		localRootRunes := []rune(strings.TrimSpace(localRoot))
		if localRootRunes[len(localRootRunes)-1] == '\\' {
			thisDirPath = fmt.Sprintf("%s%s", localRoot, sourceDir)
		} else {
			thisDirPath = fmt.Sprintf("%s\\%s", localRoot, sourceDir)
		}

		var destDirPath string
		destRootRunes := []rune(strings.TrimSpace(destRoot))
		if destRootRunes[len(destRootRunes)-1] == '\\' {
			destDirPath = fmt.Sprintf("%s%s", destRoot, sourceDir)
		} else {
			destDirPath = fmt.Sprintf("%s\\%s", destRoot, sourceDir)
		}

		log.Printf("\n------------------------\nComparing source subfolder %s contents to destination subfolder %s...\n", thisDirPath, destDirPath)
		
		sourceFileList, err := getFilesList(thisDirPath)
		if err != nil {
			log.Println("Possible mismatch: error getting files list for source directory: %s\nMoving onto next..", thisDirPath)
			log.Printf("%v\n", err)
			continue
		}		

		destFilesMap, err := getFilesMap(destDirPath)
		if err != nil {
			log.Println("Possible mismatch: error getting files map for destination directory: %s\nMoving onto next..", destDirPath)
			log.Printf("%v\n", err)
			continue
		}

		res2 := compare(sourceFileList, destFilesMap)
		if len(res2) != 0 {
			log.Println("MISMATCH: The following files exist on the source root but not in destination root:")
			for _, dirName := range res2 {
				log.Printf("\t%s\n", dirName)
			}
		} else {
			log.Printf("SUCCESS: All files in source path <%s> were found on destination path <%s>.\n\n", thisDirPath, destDirPath)
		}
	}

}

// given a list of files on source and a map of files (keyed by filename) in destination
// check if all source content is found on destination. 
func compare(srcList []string, destList map[string]int) []string {
	// ensure every item in srcList is available as a key in destList
	filesNotInDest := []string{}
	for _, fileName := range srcList {
		_, exists := destList[strings.TrimSpace(fileName)]
		if !exists {
			// log.Printf("File or folder mismatch: %s not found in destination.", strings.TrimSpace(fileName))
			filesNotInDest = append(filesNotInDest, strings.TrimSpace(fileName))
		}
	}

	return filesNotInDest
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
		_, keyFound := destFiles[strings.TrimSpace(f.Name())]
		if keyFound {
			log.Printf("getFilesMap(): destFiles key <%s> already found.\n", )
		} else {
			destFiles[f.Name()] = 1
		}
	}

	return destFiles, nil
}
