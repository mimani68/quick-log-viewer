package main

import (
	"fmt"
	"time"

	"app.io/lib/db"
	"github.com/google/uuid"
)

func main() {
	for i := 0; i < 1000; i++ {
		log := "here is local message just for test"
		s := fmt.Sprintf("{\"id\":\"%s\",\"level\":\"INFO\",\"message\":%s}", uuid.New(), log)
		err := db.Store(s, "stage", "ai")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 250)
	}
	// time.Sleep(time.Second * 2)

	// example query
	//
	// 1. Simple open regex          .........  ^f1-.*d3.*
	// 2. Both criteria should met   .........  ^(?=.*4e45)(?=.*INFO).*
	// 3. Three criteria should met  .........  ^(?=.*INFO)(?=.*d5c1)(?=.*test).*
	// 4. Line matching              .........  ^(?=.*e88.*ce).*

	data, errOfRead := db.Read("410ac20d-8825-4251-841c-112abf1c4a44", "stage", "ai")
	if errOfRead != nil {
		panic(errOfRead)
	}
	for _, val := range data {
		fmt.Printf("[INFO] %s \n", val)
	}
}
