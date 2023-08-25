-- public.city definition

-- Drop table

-- DROP TABLE public.city;

CREATE TABLE public.city (
	id_city serial4 NOT NULL,
	city varchar(255) NULL,
	CONSTRAINT city_pkey PRIMARY KEY (id_city)
);