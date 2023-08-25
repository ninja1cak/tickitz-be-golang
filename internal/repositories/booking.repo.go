package repositories

import (
	"fmt"
	"log"
	"math"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type RepoBooking struct {
	*sqlx.DB
}

func NewBooking(db *sqlx.DB) *RepoBooking {
	return &RepoBooking{db}
}

func (r *RepoBooking) CreateBooking(data *models.Booking) (*config.Result, error) {
	log.Println(data.Seats_booking, data.Id_user)
	var cinema models.Cinema
	r.Get(&cinema, r.Rebind(`SELECT id_cinema FROM public.cinema WHERE cinema_name = ?`), data.CinemaArr_name[0])

	query := `INSERT INTO public.booking (
		id_schedule,
		id_user, 
		id_cinema,
		seats_booking,
		total_prices_booking,
		watch_date,
		watch_time,
		payment_method
		)
	VALUES(
		$1,
		$2,
		$3,
		array[$4],
		$5,
		$6,
		$7,
		$8
	)`

	r.MustExec(query, data.Id_schedule, data.Id_user, cinema.Id_cinema, data.Seats_booking, data.Total_prices_booking, data.Watch_date, data.Watch_time, data.Payment_method)

	return &config.Result{Data: "Booking success created"}, nil
}

func (r *RepoBooking) GetBookingByUser(limit string, page string, id_user string) (*config.Result, error) {
	var offset int
	var next int
	var prev int

	counts := struct {
		Count int `db:"total"`
	}{}

	qCount := fmt.Sprintf(`select count(id_booking) total 
		from public.booking WHERE id_user = %s`, id_user)
	err := r.Get(&counts, qCount)

	log.Println(counts)
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

	query := fmt.Sprintf(`
	select 
		id_user,
		m.id_movie,
		m.title_movie ,
		array_to_string(seats_booking, ',') seats_booking,
		total_prices_booking,
		watch_date,
		watch_time,
		cinema_name,
		cinema_logo_url
		
	from public.booking b 
	join public.schedule s on  b.id_schedule = s.id_schedule  
	join public.movie m on m.id_movie = s.id_movie 
	join public.cinema c on b.id_cinema = c.id_cinema where id_user = %s
    LIMIT %s OFFSET %v `, id_user, limit, offset)

	var data []models.Booking
	err = r.Select(&data, query)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: data, Meta: meta}, nil
}

// func (r *RepoBooking) UpdateMovie(data *models.Movie) (string, error) {
// 	set := ""

// 	if data.Title_movie != "" {
// 		set += "title_movie = :title_movie,"
// 	}

// 	if data.Director_movie != "" {
// 		set += "director_movie = :director_movie,"
// 	}

// 	if len(data.CastsArr_movie) != 0 {
// 		data.Casts_movie = strings.Join(data.CastsArr_movie, ",")
// 		set += "casts_movie = array[:casts_movie],"
// 	}

// 	if data.Synopsis_movie != "" {
// 		set += "synopsis_movie = :synopsis_movie,"
// 	}

// 	if data.Duration_movie != "" {
// 		set += "duration_movie = :duration_movie,"
// 	}

// 	if data.Release_date_movie.String() != "" {
// 		set += "release_date_movie = :release_date_movie,"
// 	}

// 	if *data.Url_image_movie != "" {
// 		set += "url_image_movie = :url_image_movie,"
// 	}

// 	if set != "" {
// 		set += "updated_at = NOW()"

// 		query := fmt.Sprintf(`UPDATE public.movie
// 		SET
// 			%s
// 		WHERE
// 			id_movie = :id_movie
// 			`, set)

// 		_, err := r.NamedExec(query, data)
// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	var genre models.Genre
// 	if len(data.Genre.GenreArr) > 0 {
// 		_, err := r.NamedExec(`DELETE FROM public.bridge_movie_genre WHERE id_movie = :id_movie`, data)
// 		if err != nil {
// 			return "", err
// 		}
// 		for i := 0; i < len(data.Genre.GenreArr); i++ {
// 			r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
// 			if genre.Id_genre != 0 {
// 				r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, data.Id_movie, genre.Id_genre)
// 			} else {
// 				r.MustExec(`INSERT INTO public.genre(name_genre) VALUES ($1)`, data.Genre.GenreArr[i])
// 				r.Get(&genre, r.Rebind(`SELECT id_genre FROM public.genre WHERE name_genre = ?`), data.GenreArr[i])
// 				r.MustExec(`INSERT INTO public.bridge_movie_genre(id_movie, id_genre) VALUES ($1, $2)`, data.Id_movie, genre.Id_genre)
// 			}
// 			genre.Id_genre = 0
// 		}
// 	}

// 	// if len(dataSize.ProductSizeSlice) != 0 {
// 	// 	dataSize.Product_id = data.Product_id
// 	// 	dataSize.Product_size = strings.Join(dataSize.ProductSizeSlice[0:1], "")
// 	// 	dataSize.Product_price = dataSize.ProductPriceSlice[0]
// 	// 	_, err := r.NamedExec(`UPDATE public.product_size
// 	// 	SET
// 	// 		product_price = :product_price,
// 	// 		updated_at = NOW()
// 	// 	WHERE
// 	// 		product_size = :product_size AND product_id = :product_id
// 	// 		`, dataSize)

// 	// 	if err != nil {
// 	// 		return "", err
// 	// 	}

// 	// }
// 	return "update success", nil
// }

// func (r *RepoBooking) DeleteMovie(data *models.Movie) (string, error) {

// 	_, err := r.NamedExec(`DELETE FROM public.schedule WHERE id_movie = :id_movie`, data)
// 	if err != nil {
// 		return "", err
// 	}
// 	_, err = r.NamedExec(`DELETE FROM public.bridge_movie_genre WHERE id_movie = :id_movie`, data)
// 	if err != nil {
// 		return "", err
// 	}

// 	_, err = r.NamedExec(`DELETE FROM public.movie WHERE id_movie = :id_movie`, data)

// 	if err != nil {
// 		return "", err
// 	}

// 	return "delete success ", nil
// }
