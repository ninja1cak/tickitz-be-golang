package models

import "time"

type User struct {
	Id_user        string     `db:"id_user" form:"id_user" valid:"-" json:"id_user,omitempty"`
	Email_user     string     `db:"email_user" form:"email_user" valid:"email" json:"email_user,omitempty"`
	Password_user  string     `db:"password_user" form:"password_user" valid:"minstringlength(4)~Password minimal 4" json:"password_user,omitempty"`
	First_name     *string    `db:"first_name" form:"first_name" valid:"-" json:"first_name,omitempty"`
	Last_name      *string    `db:"last_name" form:"last_name" valid:"-" json:"last_name,omitempty"`
	Phone_number   *string    `db:"phone_number" form:"phone_number" valid:"-" json:"phone_number,omitempty"`
	Url_photo_user *string    `db:"url_photo_user" form:"url_photo_user" valid:"-" json:"url_photo_user,omitempty"`
	Status         *string    `db:"status" form:"status" json:"status,omitempty" valid:"-" json:"status,omitempty"`
	Role           string     `db:"role" json:"role,omitempty" valid:"-" json:"role,omitempty"`
	Created_at     *time.Time `db:"created_at" json:"created_at,omitempty" valid:"-" json:"created_at,omitempty"`
	Updated_at     *string    `db:"updated_at" json:"updated_at,omitempty" valid:"-" json:"updated_at,omitempty"`
	// File          *multipart.FileHeader `form:"file" binding:"required"`
}
