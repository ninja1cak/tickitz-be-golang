package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type RepoMovie struct {
	*sqlx.DB
}

func NewMovie(db *sqlx.DB) *RepoMovie {
	return &RepoMovie{db}
}

func (r *RepoMovie) CreateMovie(data *models.Movie) (*config.Result, error) {
	data.Casts_movie = strings.Join(data.CastsArr_movie, ",")

	query := `INSERT INTO public.movie (
		title_movie, 
		director_movie,
		duration_movie,
		casts_movie,
		synopsis_movie,
		release_date_movie,
		url_image_movie
		)
	VALUES(
		:title_movie, 
		:director_movie,
		:duration_movie,
		array[:casts_movie],
		:synopsis_movie,
		:release_date_movie,
		:url_image_movie
	)`

	_, err := r.NamedExec(query, data)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"product_product_slug_key\"" {
			return nil, errors.New("Duplicate Products")
		}
		return nil, err
	}

	var movie models.Movie
	r.Get(&movie, r.Rebind(`SELECT id_movie FROM public.movie WHERE title_movie = ?`), data.Title_movie)

	if err != nil {
		return nil, err
	}
	var genre models.Genre

	for i := 0; i < len(data.Genre.GenreArr); i++ {
		r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
		if genre.Id_genre != 0 {
			r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, movie.Id_movie, genre.Id_genre)
		} else {
			r.MustExec(`INSERT INTO public.genre(name_genre) VALUES ($1)`, data.Genre.GenreArr[i])
			r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
			r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, movie.Id_movie, genre.Id_genre)

		}
		genre.Id_genre = 0

	}

	query = `INSERT INTO public.schedule (
		id_movie,
		date_start,
		date_end,
		price_seat
		)
	VALUES(
		$1,
		$2,
		$3,
		$4
	)`
	log.Println(movie.Id_movie, data.Schedule.Date_start, data.Schedule.Date_end, data.Schedule.Price_seat)
	r.MustExec(query, movie.Id_movie, data.Schedule.Date_start, data.Schedule.Date_end, data.Schedule.Price_seat)

	var cinema models.Cinema
	var schedule models.Schedule
	var time models.Time
	var city models.City
	r.Get(&schedule, r.Rebind(`SELECT id_schedule FROM public.schedule WHERE id_movie = ?`), movie.Id_movie)

	for i := 0; i < len(data.Schedule.CinemaArr_name); i++ {
		r.Get(&cinema, r.Rebind(`SELECT id_cinema FROM public.cinema WHERE cinema_name = ?`), data.Cinema.CinemaArr_name[i])
		if cinema.Id_cinema == 0 {
			return nil, errors.New("Cinema not found")
		}
		r.MustExec(`INSERT INTO public.bridge_schedule_cinema(id_schedule, id_cinema) VALUES($1, $2)`, schedule.Id_schedule, cinema.Id_cinema)

	}

	for i := 0; i < len(data.TimeArr); i++ {
		r.Get(&time, r.Rebind(`SELECT id_time FROM public.time WHERE time = ?`), data.TimeArr[i])

		if time.Id_time != 0 {
			r.MustExec(`INSERT INTO public.bridge_schedule_time(id_schedule, id_time) VALUES ($1, $2)`, schedule.Id_schedule, time.Id_time)
		} else {
			r.MustExec(`INSERT INTO public.time(time) VALUES ($1)`, data.TimeArr[i])
			r.Get(&time, r.Rebind(`SELECT id_time FROM public.time WHERE time = ?`), data.TimeArr[i])
			r.MustExec(`INSERT INTO public.bridge_schedule_time(id_schedule, id_time) VALUES ($1, $2)`, schedule.Id_schedule, time.Id_time)
		}
		time.Id_time = 0

	}

	for i := 0; i < len(data.CityArr); i++ {
		r.Get(&city, r.Rebind(`SELECT id_city FROM public.city WHERE city = ?`), data.CityArr[i])
		log.Println(data.CityArr[i], city.Id_city)
		if city.Id_city != 0 {
			r.MustExec(`INSERT INTO public.bridge_schedule_city(id_schedule, id_city) VALUES ($1, $2)`, schedule.Id_schedule, city.Id_city)
		} else {
			r.MustExec(`INSERT INTO public.city(city) VALUES ($1)`, data.CityArr[i])
			r.Get(&city, r.Rebind(`SELECT id_city FROM public.city WHERE city = ?`), data.CityArr[i])
			r.MustExec(`INSERT INTO public.bridge_schedule_city(id_schedule, id_city) VALUES ($1, $2)`, schedule.Id_schedule, city.Id_city)
		}
		city.Id_city = 0

	}

	return &config.Result{Data: "Movie success created"}, nil
}

func (r *RepoMovie) GetMovie(limit string, page string, search string, sort string) (*config.Result, error) {
	var offset int
	var next int
	var prev int
	var qSearch string = ""
	var qSort string = ""

	if search != "" {
		search = "%" + search + "%"
		qSearch = fmt.Sprintf("AND title_movie ILIKE '%s'", search)
	}

	if sort != "" {
		sort = "%" + sort + "%"
		qSort = fmt.Sprintf("AND name_genre ILIKE '%s'", sort)
	}
	counts := struct {
		Count int `db:"total"`
	}{}

	qCount := fmt.Sprintf(`select count(id_movie) total from (select
			m.id_movie,
			m.title_movie,
			string_agg(name_genre, ', ' order by name_genre) name_genre
		from movie m 
		join bridge_movie_genre bmg on m.id_movie = bmg.id_movie 
		join genre g ON g.id_genre = bmg.id_genre group by m.id_movie) p WHERE true %s %s`, qSearch, qSort)
	err := r.Get(&counts, qCount)

	if err != nil {
		return nil, err
	}
	if counts.Count == 0 {
		return &config.Result{Message: "data not found"}, nil
	}
	lim, _ := strconv.Atoi(limit)
	pag, _ := strconv.Atoi(page)
	offset = lim * (pag - 1)

	if float64(pag) == math.Ceil(float64(counts.Count)/float64(lim)) {
		next = 0
	} else {
		next = pag + 1
	}
	if pag == 1 {
		prev = 0
	} else {
		prev = pag - 1
	}

	var meta = config.Meta{
		Next:  next,
		Prev:  prev,
		Total: counts.Count,
	}

	query := fmt.Sprintf(`select
			m.id_movie,
			title_movie, 
			director_movie,
			duration_movie,
			array_to_string(casts_movie, ', ') casts_movie,
			synopsis_movie,
			release_date_movie,
			url_image_movie,
			string_agg(name_genre, ', ' order by name_genre) name_genre
		from movie m 
		join bridge_movie_genre bmg on m.id_movie = bmg.id_movie 
		join genre g ON g.id_genre = bmg.id_genre %s %s
		group by m.id_movie LIMIT %s OFFSET %v `, qSearch, qSort, limit, offset)

	var data []models.Movie
	err = r.Select(&data, query)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: data, Meta: meta}, nil
}

func (r *RepoMovie) UpdateMovie(data *models.Movie) (string, error) {
	set := ""

	if data.Title_movie != "" {
		set += "title_movie = :title_movie,"
	}

	if data.Director_movie != "" {
		set += "director_movie = :director_movie,"
	}

	if len(data.CastsArr_movie) != 0 {
		data.Casts_movie = strings.Join(data.CastsArr_movie, ",")
		set += "casts_movie = array[:casts_movie],"
	}

	if data.Synopsis_movie != "" {
		set += "synopsis_movie = :synopsis_movie,"
	}

	if data.Duration_movie != "" {
		set += "duration_movie = :duration_movie,"
	}

	if data.Release_date_movie.String() != "" {
		set += "release_date_movie = :release_date_movie,"
	}

	if *data.Url_image_movie != "" {
		set += "url_image_movie = :url_image_movie,"
	}

	if set != "" {
		set += "updated_at = NOW()"

		query := fmt.Sprintf(`UPDATE public.movie
		SET
			%s
		WHERE
			id_movie = :id_movie
			`, set)

		_, err := r.NamedExec(query, data)
		if err != nil {
			return "", err
		}
	}

	var genre models.Genre
	if len(data.Genre.GenreArr) > 0 {
		_, err := r.NamedExec(`DELETE FROM public.bridge_movie_genre WHERE id_movie = :id_movie`, data)
		if err != nil {
			return "", err
		}
		for i := 0; i < len(data.Genre.GenreArr); i++ {
			r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
			if genre.Id_genre != 0 {
				r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, data.Id_movie, genre.Id_genre)
			} else {
				r.MustExec(`INSERT INTO public.genre(name_genre) VALUES ($1)`, data.Genre.GenreArr[i])
				r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
				r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, data.Id_movie, genre.Id_genre)
			}
			genre.Id_genre = 0
		}
	}

	// if len(dataSize.ProductSizeSlice) != 0 {
	// 	dataSize.Product_id = data.Product_id
	// 	dataSize.Product_size = strings.Join(dataSize.ProductSizeSlice[0:1], "")
	// 	dataSize.Product_price = dataSize.ProductPriceSlice[0]
	// 	_, err := r.NamedExec(`UPDATE public.product_size
	// 	SET
	// 		product_price = :product_price,
	// 		updated_at = NOW()
	// 	WHERE
	// 		product_size = :product_size AND product_id = :product_id
	// 		`, dataSize)

	// 	if err != nil {
	// 		return "", err
	// 	}

	// }
	return "update success", nil
}

func (r *RepoMovie) DeleteMovie(data *models.Movie) (string, error) {

	_, err := r.NamedExec(`DELETE FROM public.schedule WHERE id_movie = :id_movie`, data)
	if err != nil {
		return "", err
	}
	_, err = r.NamedExec(`DELETE FROM public.bridge_movie_genre WHERE id_movie = :id_movie`, data)
	if err != nil {
		return "", err
	}

	_, err = r.NamedExec(`DELETE FROM public.movie WHERE id_movie = :id_movie`, data)

	if err != nil {
		return "", err
	}

	return "delete success ", nil
}
