DROP TABLE public.bridge_movie_genre;

CREATE TABLE public.bridge_movie_genre (
	id_bridge_movie_genre serial4 NOT NULL,
	id_movie int4 NULL,
	id_genre int4 NULL,
	CONSTRAINT bridge_movie_genre_pkey PRIMARY KEY (id_bridge_movie_genre)
);


-- public.bridge_movie_genre foreign keys

ALTER TABLE public.bridge_movie_genre ADD CONSTRAINT fk_id_genre FOREIGN KEY (id_genre) REFERENCES public.genre(id_genre) DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE public.bridge_movie_genre ADD CONSTRAINT fk_id_movie FOREIGN KEY (id_movie) REFERENCES public.movie(id_movie) DEFERRABLE INITIALLY DEFERRED;