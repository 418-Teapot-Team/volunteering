CREATE TABLE users
(
    id            bigint auto_increment primary key,
    first_name    varchar(255) not null,
    last_name     varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login    datetime     not null
);

CREATE TABLE projects
(
    id          bigint auto_increment primary key,
    user_id     bigint       NOT NULL,
    title       varchar(255) NOT NULL,
    description varchar(255) NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE tasks
(
    id          bigint auto_increment primary key,
    user_id     bigint       NOT NULL,
    project_id  bigint       NOT NULL,
    title       varchar(255) NOT NULL,
    description varchar(255) NULL,
    is_public   bool      default false,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (project_id) REFERENCES projects (id)
);