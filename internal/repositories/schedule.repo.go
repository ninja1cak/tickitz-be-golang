package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type RepoSchedule struct {
	*sqlx.DB
}

func NewSchedule(db *sqlx.DB) *RepoSchedule {
	return &RepoSchedule{db}
}

func (r *RepoSchedule) GetSchedule(limit string, page string, sortLocation string, sortTime string, sortIdMovie string, sortByDate string) (*config.Result, error) {
	var offset int
	var next int
	var prev int
	var qSortLocation string = ""
	var qSortTime string = ""
	var qSortIdMovie string = ""
	var qSortByDate string = ""

	if sortLocation != "" {
		sortLocation = "%" + sortLocation + "%"
		qSortLocation = fmt.Sprintf("AND city ILIKE '%s'", sortLocation)
	}

	if sortTime != "" {
		sortTime = "%" + sortTime + "%"
		qSortTime = fmt.Sprintf("AND time ILIKE '%s'", sortTime)
	}

	if sortIdMovie != "" {
		qSortIdMovie = fmt.Sprintf("AND m.id_movie = %s", sortIdMovie)
	}

	if sortByDate != "" {
		qSortByDate = fmt.Sprintf("AND date_start <= '%s' AND date_end >= '%s'", sortByDate, sortByDate)
	}

	counts := struct {
		Count int `db:"total"`
	}{}

	qCount := fmt.Sprintf(`select count(id_movie) total
	from
		(	
	select		
		m.id_movie,
		s.id_schedule,
		title_movie,
		string_agg(distinct name_genre, ', ') name_genre ,
		date_start,
		date_end,
		price_seat,
		string_agg(distinct time, ', ' ) time,
		city,
		cinema_name,
		cinema_logo_url 
	from
		movie m 
	join bridge_movie_genre bmg on bmg.id_movie = m.id_movie 
	join genre g on g.id_genre = bmg.id_genre 
	join schedule s on s.id_movie = m.id_movie
	join bridge_schedule_time bst on bst.id_schedule = s.id_schedule 
	join "time" t on t.id_time = bst.id_time 
	join bridge_schedule_city bsc on bsc.id_schedule = s.id_schedule 
	join city c on c.id_city = bsc.id_city 
	join bridge_schedule_cinema bsc2 on bsc2.id_schedule  = s.id_schedule 
	join cinema c2 on c2.id_cinema = bsc2.id_cinema where true %s %s %s %s group by s.id_schedule,m.id_movie, m.title_movie, s.date_start, s.date_end, price_seat, cinema_name, cinema_logo_url, city
	) ta`, qSortLocation, qSortTime, qSortIdMovie, qSortByDate)
	err := r.Get(&counts, qCount)

	if err != nil {
		return nil, err
	}
	if counts.Count == 0 {
		return nil, errors.New("data not found")
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
			s.id_schedule,
			title_movie,
			string_agg(distinct name_genre, ', ') name_genre ,
			date_start,
			date_end,
			price_seat,
			string_agg(distinct time, ', ' ) time,
			city,
			cinema_name,
			cinema_logo_url 
		from
			movie m 
		join bridge_movie_genre bmg on bmg.id_movie = m.id_movie 
		join genre g on g.id_genre = bmg.id_genre 
		join schedule s on s.id_movie = m.id_movie
		join bridge_schedule_time bst on bst.id_schedule = s.id_schedule 
		join "time" t on t.id_time = bst.id_time 
		join bridge_schedule_city bsc on bsc.id_schedule = s.id_schedule 
		join city c on c.id_city = bsc.id_city 
		join bridge_schedule_cinema bsc2 on bsc2.id_schedule  = s.id_schedule 
		join cinema c2 on c2.id_cinema = bsc2.id_cinema %s %s %s %s group by s.id_schedule, m.id_movie, m.title_movie, s.date_start, s.date_end, price_seat, cinema_name, cinema_logo_url, city 
		limit %s offset %v `, qSortLocation, qSortTime, qSortIdMovie, qSortByDate, limit, offset)

	var data []models.Movie
	err = r.Select(&data, query)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: data, Meta: meta}, nil
}

func (r *RepoSchedule) GetCity() (*config.Result, error) {
	var city []models.City

	err := r.Select(&city, `select city from public.city`)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: city}, nil
}

func (r *RepoSchedule) GetTime() (*config.Result, error) {
	var time []models.Time

	err := r.Select(&time, `select time from public.time`)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: time}, nil
}

func (r *RepoSchedule) GetCinema() (*config.Result, error) {
	var cinema []models.Cinema

	err := r.Select(&cinema, `select cinema_name from public.cinema`)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: cinema}, nil
}
