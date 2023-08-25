-- public.cinema definition

-- Drop table

-- DROP TABLE public.cinema;

CREATE TABLE public.cinema (
	id_cinema serial4 NOT NULL,
	cinema_name varchar(255) NULL,
	cinema_logo_url text NULL,
	CONSTRAINT cinema_pkey PRIMARY KEY (id_cinema)
);

