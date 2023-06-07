package datalayer

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID     uint
	Name   string
	Family string
}

type SQLHandler struct {
	db *sql.DB
}

func CreateDBConnection(connString string) (*SQLHandler, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	return &SQLHandler{
		db: db,
	}, nil
}

func (handler *SQLHandler) GetPeople() ([]Person, error) {
	rows, err := handler.db.Query("select * from person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := []Person{}
	for rows.Next() {
		p := Person{}
		err = rows.Scan(&p.ID, &p.Name, &p.Family)
		if err != nil {
			log.Println(err)
			continue
		}
		people = append(people, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return people, nil
}

func (handler *SQLHandler) GetPersonByName(name string) (Person, error) {
	row := handler.db.QueryRow("select * from person where Name=?", name)
	p := Person{}

	err := row.Scan(&p.ID, &p.Name, &p.Family)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (handler *SQLHandler) AddPeople(person Person) error {
	_, err := handler.db.Exec("INSERT INTO `people`.`person` (`Name`,`Family`) VALUES (?,?)", person.Name, person.Family)
	return err
}

func (handler *SQLHandler) UpdatePerson(person Person) error {
	_, err := handler.db.Exec("UPDATE `people`.`person` SET `Name` = ?,`Family`=? WHERE (`id` = ?);", person.Name, person.Family, person.ID)
	return err
}

func (handler *SQLHandler) DeletePerson(person Person) error {
	_, err := handler.db.Exec("DELETE FROM `people`.`person` WHERE (`id` = ?);", person.ID)
	return err
}
