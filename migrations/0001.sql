CREATE TABLE groups
(
    id SERIAL PRIMARY KEY,
    title    VARCHAR(255) NOT NULL
);

CREATE TABLE tasks
(
    id  SERIAL PRIMARY KEY,
    title    VARCHAR(255) NOT NULL,
    group_id INTEGER      NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
);

CREATE TABLE timeframes
(
    task_id    int                         NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
    time_start TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    time_end   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);