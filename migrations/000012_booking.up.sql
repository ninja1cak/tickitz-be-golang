-- public.booking definition

-- Drop table

-- DROP TABLE public.booking;

CREATE TABLE public.booking (
	id_booking serial4 NOT NULL,
	seats_booking _varchar NULL,
	total_prices_booking int4 NULL,
	id_schedule int4 NULL,
	watch_date date NULL,
	payment_method varchar(50) NULL,
	id_user int4 NULL,
	watch_time varchar(100) NULL,
	id_cinema int4 NULL,
	CONSTRAINT booking_pkey PRIMARY KEY (id_booking)
);


-- public.booking foreign keys

ALTER TABLE public.booking ADD CONSTRAINT booking_fk4 FOREIGN KEY (id_cinema) REFERENCES public.cinema(id_cinema) ON DELETE CASCADE;
ALTER TABLE public.booking ADD CONSTRAINT fk_id_schedule FOREIGN KEY (id_schedule) REFERENCES public.schedule(id_schedule);
ALTER TABLE public.booking ADD CONSTRAINT fk_id_user FOREIGN KEY (id_user) REFERENCES public."user"(id_user);