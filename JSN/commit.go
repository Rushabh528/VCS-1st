package JSN

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Commit() {

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "files/")
	objpath := filepath.Join(path, ".jsn/objects/")

	var treePart string = ""
	var tree string = ""
	currentTime := time.Now()
	name := os.Getenv("AUTHOR_NAME")
	email := os.Getenv("AUTHOR_EMAIL")
	var commit string = ""
	var commitPart string = ""
	var IDs []string
	var compressedContents [][]byte

	list, err := Ls()
	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
		os.Exit(1)
	}
	for _, file := range list {
		size, err := GetFileSize1(path + "\\" + file)
		if err != nil {
			fmt.Printf("Error getting file size for %s: %v\n", file, err)
			os.Exit(1)
		}
		details, err := os.ReadFile(path + "\\" + file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			os.Exit(1)
		}
		content := fmt.Sprintf("blob %d\x00%s", size, string(details))
		ID := Sha1(content)
		compressedContent := Compress([]byte(content))
		if err != nil {
			fmt.Printf("Error compressing content for %s: %v\n", file, err)
			os.Exit(1)
		}
		// fmt.Println(ID, string(compressedContent))
		IDs = append(IDs, ID)
		compressedContents = append(compressedContents, []byte(compressedContent))

		treePart += fmt.Sprintf("100644 %s\x00%s", file, ID)
	}
	tree = fmt.Sprintf("tree %d\x00%s", len(treePart), treePart)
	compressedTreeContent := Compress([]byte(tree))
	treeID := Sha1(tree)
	IDs = append(IDs, treeID)
	compressedContents = append(compressedContents, []byte(compressedTreeContent))

	// fmt.Println(tree)

	commitPart = fmt.Sprintf("tree %s\nauthor %s <%s> %s\ncommitter %s <%s> %s", treeID, name, email, currentTime, name, email, currentTime)
	commit = fmt.Sprintf("commit %d\x00%s", len(commitPart), commitPart)
	compressedCommitContent := Compress([]byte(commit))

	commitId := Sha1(commit)
	IDs = append(IDs, commitId)
	compressedContents = append(compressedContents, []byte(compressedCommitContent))

	// fmt.Println(commit)
	// fmt.Println(commitId)
	fmt.Println(IDs)
	for i, ID := range IDs {
		err := os.Mkdir(filepath.Join(objpath, ID[:2]), os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
		}
		objFile := filepath.Join(objpath, ID[:2], ID[2:])
		if err := os.WriteFile(objFile, compressedContents[i], os.ModePerm); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}
}
