// fc copies a specified set of directories with its contents (file only) from a specified source path to a specified destination path.
// filecopy functions from https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang/21067803#21067803

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
	if len(os.Args) < 4 {
		log.Fatal("Usage: $>fc.exe \"sourcepath\" \"destpath\" \"foldername1, foldername2, foldername3\"")
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
	// log.Println(" Orders API Called")

	// read command line args
	rootDir := strings.TrimSpace(os.Args[1])
	destRoot := strings.TrimSpace(os.Args[2])
	targetDirs := strings.Split(strings.TrimSpace(os.Args[3]), ",")

	log.Printf("fc(): Copying contents of <%v> directory/directories from <%s> to <%s>...", targetDirs, rootDir, destRoot)
	time.Sleep(time.Duration(4) * time.Second)
	log.Println("fc(): starting copy..")

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	// for each subfolder in the root dir
	for _, f := range files {
		fmt.Println(f.Name())

		// check if this subfolder is in the list of foldernames in the input: if so, copy to destination
		if toBeCopied(f.Name(), targetDirs) {
			log.Printf("------------------------------------------------------ \nfc(): found <%s> folder in root directory <%s>, copying..\n", f.Name(), rootDir)
			path := fmt.Sprintf("%s\\%s", destRoot, f.Name())
			log.Printf("\n\nfc(): creating destination folder %s..\n", path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}

			fPath := fmt.Sprintf("%s\\%s", rootDir, f.Name())
			files2, err2 := ioutil.ReadDir(fPath)
			if err2 != nil {
				log.Fatal(err2)
			}

			fileNum := len(files2)
			fileCount := 1

			for _, f2 := range files2 {
				sourceFile := fmt.Sprintf("%s\\%s", fPath, f2.Name())
				destFile := fmt.Sprintf("%s\\%s", path, f2.Name())

				log.Printf("\n\tfc(): copying <%s> to <%s> (file %d of %d)\n", sourceFile, destFile, fileCount, fileNum)
				err := CopyFile(sourceFile, destFile)
				if err != nil {
					fmt.Printf("\n\tfc(): ***************** copy failed %q *********************** \n", err)
				} else {
					fmt.Printf("\tfc(): copy succeeded.\n")
					fileCount++
				}
			}
		}
	}
}

func toBeCopied(folderName string, folders []string) bool {
	// log.Printf("toBeCopied(): checking if %s is in %v..", folderName, folders)
	for _, folder := range folders {
		if strings.TrimSpace(folderName) == strings.TrimSpace(folder) {
			// log.Printf("toBeCopied(): %s is in the list of directories to be copied.", folderName)
			return true
		}
	}

	return false
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
