package datalayer

import (
	"database/sql"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Group struct {
	ID      uint
	Title   string
	EnTitle string
}

func (g Group) Validate() (map[string]string, []string, error) {

	errs := validation.Errors{
		"Title":   validation.Validate(g.Title, validation.Required.Error("عنوان نمیتواند خالی باشد")),
		"EnTitle": validation.Validate(g.EnTitle, validation.Required.Error("عنوان انگلیسی نمیتواند خالی باشد")),
	}

	for key, value := range errs {
		if value == nil {
			delete(errs, key)
		}
	}

	validationErrorsArray := make([]string, len(errs))
	validationErrors := make(map[string]string)

	for key, value := range errs {
		validationErrors[key] = value.Error()
		validationErrorsArray = append(validationErrorsArray, value.Error())
	}

	return validationErrors, validationErrorsArray, errs.Filter()
}
func (handler *SQLHandler) InsertGroups(group Group) error {
	_, err := handler.db.Exec("INSERT INTO `go_blog`.`groups` (`Title`, `EnTitle`) VALUES (?,?) ", group.Title, group.EnTitle)
	if err != nil {
		return err
	}
	return nil
}
func (handler *SQLHandler) UpdateGroup(group Group) error {
	_, err := handler.db.Exec("UPDATE `go_blog`.`groups` SET `Title` = ?, `EnTitle` = ?	WHERE Id=? ",
		group.Title, group.EnTitle, group.ID)
	if err != nil {
		return err
	}
	return nil
}
func (handler *SQLHandler) DeleteGroupById(id int) error {
	_, err := handler.db.Exec("DELETE FROM `go_blog`.`groups` WHERE Id=? ", id)
	if err != nil {
		return err
	}
	return nil
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
	row := handler.db.QueryRow("select * from `go_blog`.`groups` where Id=?", groupId)

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
