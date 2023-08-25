-- public."time" definition

-- Drop table

-- DROP TABLE public."time";

CREATE TABLE public."time" (
	id_time serial4 NOT NULL,
	"time" varchar(255) NULL,
	CONSTRAINT time_pkey PRIMARY KEY (id_time)
);