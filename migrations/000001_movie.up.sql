CREATE TABLE public.movie (
	id_movie serial4 NOT NULL,
	title_movie varchar(255) NOT NULL,
	director_movie varchar(255) NULL,
	casts_movie _text NULL,
	synopsis_movie text NULL,
	duration_movie varchar(255) NULL,
	release_date_movie date NULL,
	url_image_movie text NULL,
	updated_at timestamp NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT movie_pkey PRIMARY KEY (id_movie),
	CONSTRAINT movie_title_movie_key UNIQUE (title_movie)
);