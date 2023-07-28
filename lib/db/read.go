package db

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var project = "jinz"

func Read(query, environment, service string) ([]string, error) {
	command := fmt.Sprintf("/bin/egrep '%s' ./db/*.txt ", query)
	cmd := exec.Command("bash", "-c", command)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// return stderr.String(), err
		fmt.Println("Error")
	}
	pattern := fmt.Sprintf("./db/%s-%s-%s-\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.txt:", project, environment, service)
	re := regexp.MustCompile(pattern)
	cleanedString := re.ReplaceAllString(stdout.String(), "")
	result := strings.Split(cleanedString, "\n")
	cleanedResult := make([]string, len(result)-1)
	copy(cleanedResult, result[:len(result)-1])
	return cleanedResult, nil
}
