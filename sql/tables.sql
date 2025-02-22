CREATE TABLE users (
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    uid TEXT UNIQUE NOT NULL,
);


CREATE TABLE groups(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    GID TEXT NOT NULL,
    owner_id TEXT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE auth(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    owner_id TEXT NOT NULL,
    password TEXT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE group_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    group_id TEXT NOT NULL,
    status TEXT,
    is_owner BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);


Table USERS
User id
User name
email
Firebase uid