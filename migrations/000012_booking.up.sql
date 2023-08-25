-- public.booking definition

-- Drop table

-- DROP TABLE public.booking;

CREATE TABLE public.booking (
	id_booking serial4 NOT NULL,
	id_movie int4 NULL,
	seats_booking _varchar NULL,
	total_prices_booking int4 NULL,
	id_schedule int4 NULL,
	watch_date date NULL,
	payment_method varchar(50) NULL,
	id_user int4 NULL,
	CONSTRAINT booking_pkey PRIMARY KEY (id_booking)
);


-- public.booking foreign keys

ALTER TABLE public.booking ADD CONSTRAINT fk_id_movie FOREIGN KEY (id_movie) REFERENCES public.movie(id_movie);
ALTER TABLE public.booking ADD CONSTRAINT fk_id_schedule FOREIGN KEY (id_schedule) REFERENCES public.schedule(id_schedule);
ALTER TABLE public.booking ADD CONSTRAINT fk_id_user FOREIGN KEY (id_user) REFERENCES public."user"(id_user);