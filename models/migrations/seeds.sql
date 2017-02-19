INSERT INTO user_sources (name) VALUES ('GOOGLE');

INSERT INTO users (external_id, user_source_id, name, email)
      VALUES ('117597133628919936066', 1, 'Brent Hamilton', 'jus2funky@gmail.com');

INSERT INTO goals (name) VALUES ('exercise');
INSERT INTO goals (name) VALUES ('good food');
INSERT INTO goals (name) VALUES ('!degenerecy');
INSERT INTO goals (name) VALUES ('productivity');

INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 1, 1486800000, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 2, 1486800000, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 1, 1486886400, 1, false);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 2, 1486886400, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 3, 1486886400, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 1, 1486972800, 1, false);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 2, 1486972800, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 4, 1486972800, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 2, 1487059200, 1, true);
INSERT INTO actions (user_id, goal_id, day, weight, achieved) VALUES (1, 4, 1487059200, 1, false);