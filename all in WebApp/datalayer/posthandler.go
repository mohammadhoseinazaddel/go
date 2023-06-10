package datalayer

import (
	"database/sql"
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
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

func (p Post) Validate() (map[string]string, []string, error) {
	// return validation.ValidateStruct(&p,
	// 	validation.Field(&p.Title, validation.Required.Error("عنوان نمیتواند خالی باشد"),
	// 		validation.Length(5, 20).Error("باید بین 5 و 20 کاراکتر باشد.")),
	// 	validation.Field(&p.ShortDesc, validation.Required),
	// 	validation.Field(&p.LongDesc, validation.Required),
	// )

	validationErrors := make(map[string]string)
	errs := validation.Errors{
		"Title":     validation.Validate(p.Title, validation.Required.Error("عنوان نمیتواند خالی باشد")),
		"ShortDesc": validation.Validate(p.Title, validation.Required.Error("توضیح کوتاه نمیتواند خالی باشد")),
		"LongDesc":  validation.Validate(p.Title, validation.Required.Error("توضیح نمیتواند خالی باشد")),
	}

	for key, value := range errs {
		if value == nil {
			delete(errs, key)
		}
	}

	validationErrorsArray := make([]string, len(errs))

	for key, value := range errs {
		validationErrors[key] = value.Error()
		validationErrorsArray = append(validationErrorsArray, value.Error())
	}

	return validationErrors, validationErrorsArray, errs.Filter()
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

func (handler *SQLHandler) InsertPosts(post Post) error {
	_, err := handler.db.Exec("INSERT INTO `go_blog`.`posts` (`Title`, `Image`, `ShortDesc`, `LongDesc`, `CreateDate`, `GroupId`) VALUES (?,?,?,?,?,?) ", post.Title, post.Image, post.ShortDesc, post.LongDesc, post.CreateDate, post.GroupID)
	if err != nil {
		return err
	}
	return nil
}

func (handler *SQLHandler) UpdatePosts(post Post) error {
	_, err := handler.db.Exec("Update `go_blog`.`posts` set `Title`=?, `Image`=?, `ShortDesc`=?, `LongDesc`=?, `CreateDate`=?, `GroupId`=? where `Id`=? ", post.Title, post.Image, post.ShortDesc, post.LongDesc, post.CreateDate, post.GroupID, post.ID)
	if err != nil {
		return err
	}
	return nil
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
