INSERT INTO user_sources (name) VALUES ('GOOGLE');

INSERT INTO users (external_id, user_source_id, name, email)
      VALUES ('117597133628919936066', 1, 'Brent Hamilton', 'jus2funky@gmail.com');

INSERT INTO goals (value, weight) VALUES ('exercise', 1);
INSERT INTO goals (value, weight) VALUES ('good food', 1);
INSERT INTO goals (value, weight) VALUES ('!degenerecy', 1);
INSERT INTO goals (value, weight) VALUES ('productivity', 1);

INSERT INTO users_goals (user_id, goal_id) VALUES (1, 1);
INSERT INTO users_goals (user_id, goal_id) VALUES (1, 2);
INSERT INTO users_goals (user_id, goal_id) VALUES (1, 3);
INSERT INTO users_goals (user_id, goal_id) VALUES (1, 4);

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