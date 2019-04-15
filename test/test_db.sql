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

INSERT INTO `tasks` (title, description, tags, position, completed, list_id) VALUES
("task1", "description task1","tag1", 2, 0, 1),
("task2", "description task2","tag2", 32, 0, 1),
("task2", "description task2","", 42, 0, 1);