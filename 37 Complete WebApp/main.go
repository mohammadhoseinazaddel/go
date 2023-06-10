package main

import (
	"TopLearn/WebApp/datalayer"
	"TopLearn/WebApp/restapi"
	"log"
)

func main() {
	db, err := datalayer.CreateDBConnection("root:1234@/go_blog?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	restapi.RunApi("localhost:8383", *db)

}
