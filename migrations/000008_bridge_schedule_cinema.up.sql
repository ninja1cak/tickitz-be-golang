-- public.bridge_schedule_cinema definition

-- Drop table

-- DROP TABLE public.bridge_schedule_cinema;

CREATE TABLE public.bridge_schedule_cinema (
	id_bscinema serial4 NOT NULL,
	id_schedule int4 NULL,
	id_cinema int4 NULL,
	CONSTRAINT bscinema_pkey PRIMARY KEY (id_bscinema)
);


-- public.bridge_schedule_cinema foreign keys

ALTER TABLE public.bridge_schedule_cinema ADD CONSTRAINT bscinema_fk1 FOREIGN KEY (id_cinema) REFERENCES public.cinema(id_cinema) ON DELETE CASCADE;
ALTER TABLE public.bridge_schedule_cinema ADD CONSTRAINT bsinemac_fk FOREIGN KEY (id_schedule) REFERENCES public.schedule(id_schedule) ON DELETE CASCADE;