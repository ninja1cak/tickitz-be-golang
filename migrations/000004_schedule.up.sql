-- public.schedule definition

-- Drop table

-- DROP TABLE public.schedule;

CREATE TABLE public.schedule (
	id_schedule serial4 NOT NULL,
	id_movie int4 NULL,
	date_start date NULL,
	date_end date NULL,
	cinema_address text NULL,
	price_seat int4 NULL,
	CONSTRAINT schedule_pkey PRIMARY KEY (id_schedule)
);


-- public.schedule foreign keys

ALTER TABLE public.schedule ADD CONSTRAINT fk_id_movie FOREIGN KEY (id_movie) REFERENCES public.movie(id_movie);