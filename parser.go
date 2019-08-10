package envparser

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

// Parser represent a config file parser
type Parser interface {
	File(string) string
	Parse([]byte) error
}

// ParseConfig parses a configuration file
func ParseConfig(env string, config Parser) error {
	filename := config.File(env)
	fileBytes, err := parseFile(filename)
	if err != nil {
		return err
	}
	if err := config.Parse(fileBytes); err != nil {
		return err
	}
	return nil
}

// parseFile takes in filename and returns a reader
func parseFile(filename string) ([]byte, error) {
	path, err := getPath(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot get file path, %v", err)
	}
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file, %v", err)
	}
	return fileBytes, nil
}

// getExecPath gets the Path of the executable file
func getExecPath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	execPath, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", err
	}
	return execPath, nil
}

// getPath gets the full path of the file
func getPath(filename string) (string, error) {
	execPath, err := getExecPath()
	if err != nil {
		return "", err
	}
	PathList := []string{filepath.Dir(execPath), filename}
	Path := strings.Join(PathList, "")
	return Path, nil
}