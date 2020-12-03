package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirectoryExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateFileName(name string) string {
	timeStamp := time.Now().Unix()
	return fmt.Sprintf("%d_%s.sql", timeStamp, name)
}

func GetVersionByFileName(fileName string) string {
	return strings.Split(fileName, "_")[0]
}

func SortFilesByNameAsc(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	return files
}

func CreateMigrationFile(source, name string) (string, error) {
	fileName := CreateFileName(name)
	fullPath := path.Join(source, fileName)
	err := os.MkdirAll(source, 0755)
	if err != nil {
		return "", err
	}
	comment := "-- insert SQL script for update you database"
	err = ioutil.WriteFile(fullPath, []byte(comment), 0644)
	return fullPath, err
}
