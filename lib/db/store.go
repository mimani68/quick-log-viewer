package db

import (
	"fmt"
	"io/ioutil"
	"time"
)

func Store(value, environment, service string) error {
	filename := fmt.Sprintf("./db/%s-%s-%s-%s.txt", project, environment, service, time.Now().Format("2006-01-02T15:04:05"))
	return ioutil.WriteFile(filename, []byte(value), 0644)
}
