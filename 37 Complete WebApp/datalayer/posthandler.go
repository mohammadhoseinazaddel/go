package datalayer

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	ID         uint
	Title      string
	Image      string
	ShortDesc  string
	LongDesc   string
	CreateDate time.Time
	GroupID    uint
}

func (handler *SQLHandler) GetPosts() ([]Post, error) {
	rows, err := handler.db.Query("select * from posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		p, err := getRowsData(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (handler *SQLHandler) GetPostsByGroupId(groupId int) ([]Post, error) {
	rows, err := handler.db.Query("select * from posts where GroupId=?", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		p, err := getRowsData(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (handler *SQLHandler) GetPostById(postId int) (Post, error) {
	row := handler.db.QueryRow("select * from posts where Id=?", postId)

	return getRowData(row)
}

func getRowData(row *sql.Row) (Post, error) {
	p := Post{}

	err := row.Scan(&p.ID, &p.Title, &p.Image, &p.ShortDesc, &p.LongDesc, &p.CreateDate, &p.GroupID)
	if err != nil {
		return p, err
	}
	return p, nil
}

func getRowsData(row *sql.Rows) (Post, error) {
	p := Post{}

	err := row.Scan(&p.ID, &p.Title, &p.Image, &p.ShortDesc, &p.LongDesc, &p.CreateDate, &p.GroupID)
	if err != nil {
		return p, err
	}
	return p, nil
}
