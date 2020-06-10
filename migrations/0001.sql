CREATE TABLE groups
(
    group_id SERIAL PRIMARY KEY,
    title    VARCHAR(255) NOT NULL
);

CREATE TABLE tasks
(
    task_id  SERIAL PRIMARY KEY,
    title    VARCHAR(255) NOT NULL,
    group_id INTEGER      NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
);

CREATE TABLE timeframes
(
    task_id    int                         NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks (task_id) ON DELETE CASCADE,
    time_start TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    time_end   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

INSERT INTO groups(title) VALUES ('provero4ka')