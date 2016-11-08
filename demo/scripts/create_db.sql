drop table users CASCADE;
drop table api_keys CASCADE;

create table users (
  id       serial primary key,
  username varchar(255),
  password varchar(255)
);

create table api_keys (
  token   varchar(255),
  user_id integer references users(id)
);

INSERT INTO users (username, password) VALUES ('billy', 'pass');
INSERT INTO users (username, password) VALUES ('chris', 'pass');
INSERT INTO users (username, password) VALUES ('jeff', 'pass');
