package main

import (
	"log"
	"main/datalayer"
	"main/restapi"
)

func main() {
	db, err := datalayer.CreateDBConnection("root:1234@/go_blog?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	restapi.RunApi("localhost:8383", *db)

}
