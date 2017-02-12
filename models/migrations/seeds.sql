INSERT INTO user_sources (name) VALUES ('GOOGLE');

INSERT INTO users (external_id, user_source_id, name, email)
      VALUES ('117597133628919936066', 1, 'Brent Hamilton', 'jus2funky@gmail.com');

INSERT INTO goals (value) VALUES ('exercise');
INSERT INTO goals (value) VALUES ('good food');
INSERT INTO goals (value) VALUES ('!degenerecy');
INSERT INTO goals (value) VALUES ('productivity');

INSERT INTO users_goals (user_id, goal_id, weight) VALUES (1, 1, 1);
INSERT INTO users_goals (user_id, goal_id, weight) VALUES (1, 2, 1);
INSERT INTO users_goals (user_id, goal_id, weight) VALUES (1, 3, 1);
INSERT INTO users_goals (user_id, goal_id, weight) VALUES (1, 4, 1);

INSERT INTO actions (user_id, goal_id, day) VALUES (1, 1, 1486800000);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 2, 1486800000);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 1, 1486886400);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 2, 1486886400);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 3, 1486886400);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 1, 1486972800);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 2, 1486972800);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 4, 1486972800);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 2, 1487059200);
INSERT INTO actions (user_id, goal_id, day) VALUES (1, 4, 1487059200);