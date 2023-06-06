package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	ID   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "user:1234@/people")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/* r */
	rows, err := db.Query("select * from person")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	people := []person{}
	for rows.Next() {
		p := person{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		people = append(people, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(people)

	row := db.QueryRow("select * from person where id=?", 2)
	p := person{}

	err = row.Scan(&p.ID, &p.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	/*C U D*/

	res, err := db.Exec("INSERT INTO `people`.`person` (`Name`) VALUES (?)", "Ali")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	res, err = db.Exec("UPDATE `people`.`person` SET `Name` = ? WHERE (`id` = ?);", "Ali", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	res, err = db.Exec("DELETE FROM `people`.`person` ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

}
