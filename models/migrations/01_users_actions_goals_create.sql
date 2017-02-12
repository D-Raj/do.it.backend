CREATE TABLE user_sources (
  id INT NOT NULL AUTO_INCREMENT,
  name varchar(50) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY `user_sources_name` (`name`)
);

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  external_id VARCHAR(255) NOT NULL,
  user_source_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_source_id) REFERENCES user_sources(id),
  UNIQUE KEY users_external_id_source (external_id, user_source_id)
);

CREATE TABLE goals (
  id INT NOT NULL AUTO_INCREMENT NOT NULL,
  value VARCHAR(255) NOT NULL,
  weight INT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY goals_user_id_value (value, weight)
);

CREATE TABLE actions (
  id INT NOT NULL AUTO_INCREMENT NOT NULL,
  user_id INT NOT NULL,
  goal_id INT NOT NULL,
  day INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (goal_id) REFERENCES goals(id),
  UNIQUE KEY days_user_id_goal_id_day (user_id, goal_id, day)
);

CREATE TABLE users_goals (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  goal_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (goal_id) REFERENCES goals(id),
  UNIQUE KEY users_goals_user_id_goal_id (user_id, goal_id)
);