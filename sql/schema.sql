CREATE TABLE puzzles
(
    id serial NOT NULL,
    date date NOT NULL,
    letters character(1)[] NOT NULL,
    center character(1) NOT NULL,
    max_score integer NOT NULL,
    CONSTRAINT puzzles_pkey PRIMARY KEY (id)
);

CREATE TABLE rooms
(
    id serial NOT NULL,
    code character(8) NOT NULL,
    score integer NOT NULL,
    CONSTRAINT rooms_pkey PRIMARY KEY (id),
    CONSTRAINT rooms_code_key UNIQUE (code)
);

CREATE TABLE words
(
    id serial NOT NULL,
    word character varying(32) NOT NULL,
    puzzle_id integer NOT NULL,
    point_value smallint NOT NULL,
    definition text NOT NULL,
    part_of_speech character varying(32) NOT NULL,
    synonym character varying(32) NOT NULL,
    CONSTRAINT words_pkey PRIMARY KEY (id),
    CONSTRAINT words_word_key UNIQUE (word),
    CONSTRAINT words_puzzle_id_fkey FOREIGN KEY (puzzle_id)
        REFERENCES puzzles (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE records
(
    id serial NOT NULL,
    word_id integer NOT NULL,
    room_id integer NOT NULL,
    player_name character varying(64) NOT NULL,
    found_at timestamp with time zone NOT NULL,
    CONSTRAINT records_pkey PRIMARY KEY (id),
    CONSTRAINT records_word_id_room_id_key UNIQUE (word_id, room_id),
    CONSTRAINT records_room_id_fkey FOREIGN KEY (room_id)
        REFERENCES rooms (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT records_word_id_fkey FOREIGN KEY (word_id)
        REFERENCES words (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);