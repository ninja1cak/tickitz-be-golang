package models

import "time"

type User struct {
	Id_user        string    `db:"id_user" form:"id_user" valid:"-"`
	Email_user     string    `db:"email_user" form:"email_user" valid:"email"`
	Password_user  string    `db:"password_user" form:"password_user" valid:"minstringlength(4)~Password minimal 4"`
	First_name     *string   `db:"first_name" form:"first_name" valid:"-"`
	Last_name      *string   `db:"last_name" form:"last_name" valid:"-"`
	Phone_number   *string   `db:"phone_number" form:"phone_number" valid:"-"`
	Url_photo_user *string   `db:"url_photo_user" form:"url_photo_user" valid:"-"`
	Status         *string   `db:"status" form:"status" json:"status,omitempty" valid:"-"`
	Role           string    `db:"role" json:"role,omitempty" valid:"-"`
	Created_at     time.Time `db:"created_at" json:"created_at,omitempty" valid:"-"`
	Updated_at     *string   `db:"updated_at" json:"updated_at,omitempty" valid:"-"`
	// File          *multipart.FileHeader `form:"file" binding:"required"`
}
