CREATE TABLE IF NOT EXISTS user_password
(
    id bigint NOT NULL PRIMARY KEY,
    password varchar(500) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(100) NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100) NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);