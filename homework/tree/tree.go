package main

import (
	"os"
//	"path/filepath"
	"io/ioutil"
)

// сюда писать функцию DirTree

func dirTree(out *os.File, path string, printFiles bool) (error) {
	return visitDir(out, path, printFiles, "|" + "\t")
}

func addToOut(out *os.File, name string, level string) {
	out.WriteString(level + name + "\n")
}

func visitDir(out *os.File, path string, printFiles bool, level string) (error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			addToOut(out, file.Name(), level)
			visitDir(out, path + "/" + file.Name(), printFiles, level + "|" + "\t")
		} else if printFiles {
			addToOut(out, file.Name(), level)
		}
	}
//	out.WriteString(filepath.Dir(path))
	return nil	
}