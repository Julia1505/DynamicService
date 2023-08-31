

CREATE TABLE IF NOT EXISTS users (
    user_id integer,
    PRIMARY KEY(user_id)
);

CREATE TABLE IF NOT EXISTS segment (
    segment_id serial PRIMARY KEY,
    segment_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS segment_user (
    user_id int REFERENCES users (user_id) ON UPDATE CASCADE ON DELETE CASCADE,
    segment_id int REFERENCES segment (segment_id) ON UPDATE CASCADE,
    CONSTRAINT segment_user_pkey PRIMARY KEY (user_id, segment_id)
);