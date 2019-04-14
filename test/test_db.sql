USE todo_db;

INSERT INTO `users` (username) VALUES
("antonio"),
("marcos"),
("matias");

INSERT INTO `lists` (title, user_id) VALUES
("antonio tasklist 1", (SELECT id FROM users WHERE username = 'antonio')),
("antonio tasklist 2", (SELECT id FROM users WHERE username = 'antonio')),
("marcos tasklist 1", (SELECT id FROM users WHERE username = 'marcos')),
("matias tasklist 1", (SELECT id FROM users WHERE username = 'matias'));

