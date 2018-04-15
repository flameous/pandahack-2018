package db

import (
	"database/sql"
	"fmt"
	"github.com/flameous/pandahack-2018/types"
	"github.com/pkg/errors"
)

type Database struct {
	*sql.DB
}

func NewDatabase(user, pass, addr, db string) (*Database, error) {
	d, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		user, pass, addr, db,
	))
	if err != nil {
		return nil, err
	}
	return &Database{d}, d.Ping()
}

func (db *Database) GetUserByID(id int) (interface{}, error) {
	var u types.User
	err := db.QueryRow(`SELECT id, first_name, last_name, username, email, password FROM users WHERE id = $1`, id).Scan(
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *Database) GetTaskByID(id int) (interface{}, error) {
	var t types.Task

	err := db.QueryRow(`SELECT id, title FROM tasks WHERE id = $1`, id).Scan(
		&t.Id,
		&t.Title,
	)
	if err != nil {
		return nil, errors.Wrap(err, "group 1")
	}

	rows, err := db.Query(`SELECT id, completed, question, answer FROM personal_tasks
 	WHERE id=ANY(SELECT personal_task_id FROM tasks_personal_tasks WHERE task_id = $1)`, id)
	if err != nil {
		return nil, errors.Wrap(err, "group 2")
	}
	defer rows.Close()

	for rows.Next() {
		var pt types.PersonalTask
		if err = rows.Scan(&pt.Id, &pt.Completed, &pt.Question, &pt.Answer); err != nil {
			return nil, errors.Wrap(err, "group 3")
		}
		t.PersonalTasks = append(t.PersonalTasks, &pt)
	}

	for _, v := range t.PersonalTasks {
		var u types.User
		err = db.QueryRow(`SELECT id, first_name, last_name, username, email, password
		FROM users WHERE id=ANY(SELECT user_id FROM user_personal_tasks WHERE personal_task_id = $1)`,
			v.Id).Scan(
			&u.Id,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Email,
			&u.Password,
		)
		if err != nil {
			return nil, errors.Wrap(err, "group 4")
		}
		v.User = &u
	}
	return t, nil
}

type GetterByID func(id int) (interface{}, error)
