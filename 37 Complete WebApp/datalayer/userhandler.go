package datalayer

import (
	"database/sql"
	"log"
)

type User struct {
	ID       uint
	Password string
	Email    string
	Name     string
}

func (handler *SQLHandler) GetUsers() ([]User, error) {
	rows, err := handler.db.Query("SELECT * FROM `go_blog`.`users`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		g, err := getRowsDataUser(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, g)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (handler *SQLHandler) GetUserById(userId int) (User, error) {
	row := handler.db.QueryRow("select * from users where Id=?", userId)

	return getRowDataUser(row)
}

func (handler *SQLHandler) GetUserByEmail(email string) (User, error) {
	row := handler.db.QueryRow("select * from users where email=?", email)

	return getRowDataUser(row)
}

func getRowDataUser(row *sql.Row) (User, error) {
	g := User{}

	err := row.Scan(&g.ID, &g.Name, &g.Email, &g.Password)
	if err != nil {
		return g, err
	}
	return g, nil
}

func getRowsDataUser(row *sql.Rows) (User, error) {
	g := User{}

	err := row.Scan(&g.ID, &g.Name, &g.Email, &g.Password)
	if err != nil {
		return g, err
	}
	return g, nil
}
