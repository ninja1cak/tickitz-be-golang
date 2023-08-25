-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id_user int4 NOT NULL DEFAULT nextval('users_id_user_seq'::regclass),
	email_user varchar(255) NOT NULL,
	password_user text NOT NULL,
	username text NULL,
	"role" varchar(255) NULL,
	status text NULL,
	first_name varchar(50) NULL,
	last_name varchar(50) NULL,
	phone_number varchar(50) NULL,
	url_photo_user text NULL,
	updated_at timestamp NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT email UNIQUE (email_user),
	CONSTRAINT users_pkey PRIMARY KEY (id_user),
	CONSTRAINT users_username_key UNIQUE (username)
);