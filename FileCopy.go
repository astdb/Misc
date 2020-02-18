// fc copies a specified directory with its contents (file only) from a specified source path to a specified destination path. 
// file copy functions from https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang/21067803#21067803

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
		log.Fatal("Usage: $>fc.exe \"sourcepath\" \"destpath\" \"foldername\"")
	}

	rootDir := strings.TrimSpace(os.Args[1])
	destRoot := strings.TrimSpace(os.Args[2])
	targetDir := strings.TrimSpace(os.Args[3])

	log.Printf("fc(): Copying contents of <%s> directory from <%s> to <%s>...", targetDir, rootDir, destRoot)
	time.Sleep(time.Duration(8) * time.Second)
	log.Println("fc(): starting copy..")

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// fmt.Println(f.Name())

		// each of these is a folder name as well
		// create dir
		if f.Name() == targetDir {
			log.Printf("fc(): found %s in %s\n", targetDir, rootDir)
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

			for _, f2 := range files2 {
				sourceFile := fmt.Sprintf("%s\\%s", fPath, f2.Name())
				destFile := fmt.Sprintf("%s\\%s", path, f2.Name())

				log.Printf("\n\nfc(): copying %s to %s\n", sourceFile, destFile)
				err := CopyFile(sourceFile, destFile)
				if err != nil {
					fmt.Printf("fc(): copy failed %q\n", err)
				} else {
					fmt.Printf("fc(): copy succeeded\n")
				}
			}
		}
	}
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
