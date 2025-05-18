CREATE TABLE events (
  id SERIAL PRIMARY KEY NOT NULL,
  list_participant_id INTEGER UNIQUE,
  title VARCHAR NOT NULL,
  description TEXT,
  date_time TIMESTAMP,
  location VARCHAR,
  registration_open TIMESTAMP,
  registration_close TIMESTAMP,
  chat_link VARCHAR,
  feedback_link VARCHAR
);

CREATE TABLE participants (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR NOT NULL,
  study_group VARCHAR,
  phone VARCHAR,
  telegram VARCHAR
);

CREATE TABLE event_list (
  id SERIAL PRIMARY KEY,
  participant_id INTEGER REFERENCES participants(id),
  event_id INTEGER REFERENCES events(id)
);

CREATE TABLE cinema_evening_participants (
  id SERIAL PRIMARY KEY,
  participant_id INTEGER REFERENCES participants(id),
  is_bmstu BOOLEAN,
  FOREIGN KEY (id) REFERENCES events(list_participant_id)
);

CREATE TABLE game_teams (
  id SERIAL PRIMARY KEY,
  captain_id INTEGER REFERENCES participants(id),
  only_bmstu BOOLEAN,
  team_name VARCHAR,
  team_size INTEGER,
  FOREIGN KEY (id) REFERENCES events(list_participant_id)
);

CREATE TABLE dietary_restrictions (
  id SERIAL PRIMARY KEY,
  no_meat BOOLEAN DEFAULT false,
  no_milk BOOLEAN DEFAULT false,
  allergies VARCHAR
);

CREATE TABLE trip_participants (
  id SERIAL PRIMARY KEY,
  participant_id INTEGER REFERENCES participants(id),
  birth_date DATE,
  vk VARCHAR,
  health_notes TEXT,
  restriction_id INTEGER REFERENCES dietary_restrictions(id),
  FOREIGN KEY (id) REFERENCES events(list_participant_id)
);

CREATE TABLE registration_status (
  id SERIAL PRIMARY KEY,
  participant_id INTEGER REFERENCES participants(id),
  confirmed BOOLEAN DEFAULT false,
  confirmation_date TIMESTAMP
);
