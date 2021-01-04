package main

import (
	"sort"
	"strings"
)

type FileSystem struct {
	Name    string
	Files   map[string]*FileSystem
	Content string
}

func Constructor() FileSystem {
	fs := FileSystem{
		Name:  "/",
		Files: map[string]*FileSystem{},
	}

	return fs
}

func (this *FileSystem) Ls(path string) []string {
	currDirectory := this
	folders := strings.Split(path, "/")

	for i := range folders {
		if folders[i] != "" {
			currDirectory = currDirectory.Files[folders[i]]
		}
	}

	files := []string{}

	if currDirectory.Content != "" {
		files = append(files, currDirectory.Name)
	}
	for _, v := range currDirectory.Files {
		files = append(files, v.Name)
	}

	sort.Strings(files)

	return files
}

func (this *FileSystem) Mkdir(path string) {
	currDirectory := this

	folders := strings.Split(path, "/")
	for i := range folders {
		if folders[i] != "" {
			if _, ok := currDirectory.Files[folders[i]]; !ok {
				// if folder doesnt exist, create it
				newFs := Constructor()
				newFs.Name = folders[i]
				currDirectory.Files[folders[i]] = &newFs
			}

			currDirectory = currDirectory.Files[folders[i]]
		}

	}

}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	this.Mkdir(filePath)
	currDirectory := this

	folders := strings.Split(filePath, "/")
	for i := range folders {
		if folders[i] != "" {
			currDirectory = currDirectory.Files[folders[i]]
		}
	}

	currDirectory.Content += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	currDirectory := this

	folders := strings.Split(filePath, "/")
	for i := range folders {
		if folders[i] != "" {
			currDirectory = currDirectory.Files[folders[i]]
		}
	}
	return currDirectory.Content
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
