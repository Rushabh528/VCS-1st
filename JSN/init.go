package JSN

import (
	"log"
	"os"
	"path/filepath"
)

func Init() {
	wd, _ := os.Getwd()

	headDirPath0 := filepath.Join(wd, "files/.jsn")

	err := os.Mkdir(headDirPath0, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory %s: %v", headDirPath0, err)
	}
	headPath := filepath.Join(headDirPath0, "HEAD")
	headDirPath2 := filepath.Join(headDirPath0, "objects")
	err0 := os.WriteFile(headPath, []byte{}, os.ModePerm)
	err1 := os.Mkdir(headDirPath2, os.ModePerm)
	if err0 != nil {
		log.Fatal(err0)
	}
	if err1 != nil {
		log.Fatal(err1)
	}
	return
}
