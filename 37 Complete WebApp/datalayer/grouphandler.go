package datalayer

import (
	"database/sql"
	"log"
)

type Group struct {
	ID      uint
	Title   string
	EnTitle string
}

func (handler *SQLHandler) GetGroups() ([]Group, error) {
	rows, err := handler.db.Query("SELECT * FROM `go_blog`.`groups`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := []Group{}
	for rows.Next() {
		g, err := getRowsDataGroup(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		groups = append(groups, g)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (handler *SQLHandler) GetGroupById(groupId int) (Group, error) {
	row := handler.db.QueryRow("select * from groups where Id=?", groupId)

	return getRowDataGroup(row)
}

func getRowDataGroup(row *sql.Row) (Group, error) {
	g := Group{}

	err := row.Scan(&g.ID, &g.Title, &g.EnTitle)
	if err != nil {
		return g, err
	}
	return g, nil
}

func getRowsDataGroup(row *sql.Rows) (Group, error) {
	g := Group{}

	err := row.Scan(&g.ID, &g.Title, &g.EnTitle)
	if err != nil {
		return g, err
	}
	return g, nil
}
