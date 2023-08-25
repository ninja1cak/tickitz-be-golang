package models

import "time"

type Movie struct {
	Id_movie           int       `db:"id_movie" form:"id_movie" json:"id_movie,omitempty"`
	Title_movie        string    `db:"title_movie" form:"title_movie" json:"title_movie,omitempty"`
	Director_movie     string    `db:"director_movie" form:"director_movie" json:"director_movie,omitempty"`
	Casts_movie        string    `db:"casts_movie" json:"casts_movie,omitempty"`
	CastsArr_movie     []string  `form:"casts_movie" json:"CastsArr_movie,omitempty"`
	Synopsis_movie     string    `db:"synopsis_movie" form:"synopsis_movie" json:"synopsis_movie,omitempty"`
	Duration_movie     string    `db:"duration_movie" form:"duration_movie" json:"duration_movie,omitempty"`
	Release_date_movie time.Time `db:"release_date_movie" form:"release_date_movie" json:"release_date_movie,omitempty"`
	Url_image_movie    *string   `db:"url_image_movie" form:"url_image_movie" json:"url_image_movie,omitempty"`
	Genre
	Schedule
	Created_at time.Time  `db:"created_at" form:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" form:"updated_at" json:"updated_at,omitempty"`
}

type Genre struct {
	Id_genre int      `db:"id_genre" json:"id_genre,omitempty"`
	Genre    string   `db:"name_genre" form:"genre" json:"genre,omitempty"`
	GenreArr []string `form:"genre" json:"genreArr,omitempty"`
}
