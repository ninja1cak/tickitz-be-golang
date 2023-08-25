-- public.bridge_schedule_time definition

-- Drop table

-- DROP TABLE public.bridge_schedule_time;

CREATE TABLE public.bridge_schedule_time (
	id_bst serial4 NOT NULL,
	id_schedule int4 NULL,
	id_time int4 NULL,
	CONSTRAINT bst_pkey PRIMARY KEY (id_bst)
);


-- public.bridge_schedule_time foreign keys

ALTER TABLE public.bridge_schedule_time ADD CONSTRAINT bst_fk FOREIGN KEY (id_schedule) REFERENCES public.schedule(id_schedule) ON DELETE CASCADE;
ALTER TABLE public.bridge_schedule_time ADD CONSTRAINT bst_fk1 FOREIGN KEY (id_time) REFERENCES public."time"(id_time) ON DELETE CASCADE;