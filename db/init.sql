-- table users
CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  first_name TEXT,
  last_name  TEXT,
  username   TEXT UNIQUE,
  email      TEXT UNIQUE,
  password   TEXT
);

INSERT INTO
  users (first_name, last_name, username, email, password)
VALUES
  ('Вася', 'Хуяся', 'foobar', 'foo@bar.com', '12345'),
  ('Петя', 'Хуетя', 'foobar1', 'foo@bar.com1', '12345'),
  ('Лиза', 'Хуиза', 'foobar2', 'foo@bar.com2', '12345'),
  ('Маша', 'Хуяша', 'foobar3', 'foo@bar.com3', '12345'),
  ('Qux', 'Quux', 'quxquux', 'qux@quux.com', '654321');

--
--
--
--
--
--
--
-- table tasks
CREATE TABLE tasks (
  id    SERIAL PRIMARY KEY,
  title TEXT
);

-- insert tasks
INSERT INTO
  tasks (title)
VALUES
  ('Задача первая'),
  ('Задача вторая'),
  ('Задача третья'),
  ('Задача четвёртая');

--
--
--
--
--
--
--
--
--
--
CREATE TABLE personal_tasks (
  id        SERIAL PRIMARY KEY,
  completed BOOLEAN,
  question  TEXT,
  answer    TEXT
);

INSERT INTO
  personal_tasks (completed, question, answer)
VALUES
  (FALSE, '4 / 2 = ?', '2'),
  (TRUE, 'Имя первого космонавта', 'юрий'),
  (TRUE, 'Что такое ЕПАМ?', 'галера'),
  (FALSE, 'foo, ?, baz, qux...', 'bar');

--
--
--
--
--
--
--
--
--
--
CREATE TABLE tasks_personal_tasks (
  task_id          INT,
  personal_task_id INT,
  PRIMARY KEY (task_id, personal_task_id)
);

INSERT INTO
  tasks_personal_tasks (task_id, personal_task_id)
VALUES
  (1, 1),
  (1, 2),
  (1, 3),
  (1, 4),
  (2, 1);

--
--
--
--
--
--
--
--
--
--
CREATE TABLE user_personal_tasks (
  user_id          INT,
  personal_task_id INT,
  PRIMARY KEY (user_id, personal_task_id)
);

-- add users to tasks
INSERT INTO
  user_personal_tasks (user_id, personal_task_id)
VALUES
  (1, 1),
  (2, 2),
  (3, 3),
  (4, 4);
