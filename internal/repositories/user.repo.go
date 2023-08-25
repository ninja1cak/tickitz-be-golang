package repositories

import (
	"errors"
	"fmt"
	"log"
	"ninja1cak/coffeshop-be/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.user (
		email_user, 
		password_user, 
		role,
		status
		)
	VALUES(
		:email_user, 
		:password_user, 
		:role,
		'pending'
	)`

	_, err := r.NamedExec(query, data)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"user_email_key\"" {
			return "", errors.New("Email already registered.")
		}
		return "", err
	}

	return "user success created", nil
}

func (r *RepoUser) GetUser(user_id string) (interface{}, error) {
	var queryId string = ""
	var err error
	if user_id != "" {
		queryId = "WHERE id_user= $1"
	}

	query := fmt.Sprintf(`SELECT 
	id_user,
	email_user,
	first_name,
	last_name,
	phone_number,
	url_photo_user
FROM
	public.user %s`, queryId)

	data := []models.User{}

	if user_id == "" {

		if err = r.Select(&data, query); err != nil {
			return "", err
		}
	} else {
		if err = r.Select(&data, query, user_id); err != nil {
			return "", err
		}
	}
	log.Println(data)

	return data, nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	set := ""

	if data.Password_user != "" {
		set += "password_user = :password_user,"
	}

	if data.First_name != nil {
		set += "first_name = :first_name,"
	}

	if data.Last_name != nil {
		set += "last_name = :last_name,"
	}

	if *data.Url_photo_user != "" {
		set += "url_photo_user = :url_photo_user,"
	}

	set += "updated_at = NOW()"

	query := fmt.Sprintf(`UPDATE public.user
	SET
		%s
	WHERE
		email_user = :email_user
		`, set)

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "update user success", nil
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.user
	WHERE
		email_user = :email_user
		`

	res, err := r.NamedExec(query, data)
	rowsChange, err := res.RowsAffected()
	if rowsChange == 0 {
		return "data not found", nil
	}
	if err != nil {
		return "", err
	}

	return "Delete user success", nil
}

func (r *RepoUser) GetAuthData(email string) (*models.User, error) {
	var result models.User
	query := `SELECT id_user, email_user, password_user, role, status from public.user where email_user = ?`

	if err := r.Get(&result, r.Rebind(query), email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil

}

func (r *RepoUser) UpdateStatusUser(email string) (string, error) {

	query := `UPDATE public.user
	SET
		status = 'active'
	WHERE
		email_user = $1
		`

	result := r.MustExec(query, email)
	log.Println(result)
	return "Verify Success", nil
}
