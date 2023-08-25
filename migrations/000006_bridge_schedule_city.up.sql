-- public.bridge_schedule_city definition

-- Drop table

-- DROP TABLE public.bridge_schedule_city;

CREATE TABLE public.bridge_schedule_city (
	id_bsc serial4 NOT NULL,
	id_schedule int4 NULL,
	id_city int4 NULL,
	CONSTRAINT bsc_pkey PRIMARY KEY (id_bsc)
);


-- public.bridge_schedule_city foreign keys

ALTER TABLE public.bridge_schedule_city ADD CONSTRAINT bsc_fk FOREIGN KEY (id_schedule) REFERENCES public.schedule(id_schedule) ON DELETE CASCADE;
ALTER TABLE public.bridge_schedule_city ADD CONSTRAINT bsc_fk1 FOREIGN KEY (id_city) REFERENCES public.city(id_city) ON DELETE CASCADE;