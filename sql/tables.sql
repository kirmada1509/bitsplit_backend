-- CREATE TABLE users (
--     name TEXT NOT NULL,
--     email TEXT UNIQUE NOT NULL,
--     uid TEXT UNIQUE NOT NULL,
-- );


-- CREATE TABLE groups(
--     id INTEGER PRIMARY KEY AUTOINCREMENT,
--     name TEXT NOT NULL,
--     GID TEXT NOT NULL,
--     owner_id TEXT NOT NULL,
--     FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
-- );

-- CREATE TABLE auth(
--     id INTEGER PRIMARY KEY AUTOINCREMENT,
--     name TEXT NOT NULL,
--     email TEXT NOT NULL,
--     owner_id TEXT NOT NULL,
--     password TEXT NOT NULL,
--     FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
-- );


-- CREATE TABLE group_users (
--     id INTEGER PRIMARY KEY AUTOINCREMENT,
--     user_id TEXT NOT NULL,
--     user_name TEXT NOT NULL,
--     group_id TEXT NOT NULL,
--     status TEXT,
--     is_owner BOOLEAN NOT NULL DEFAULT FALSE,
--     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
--     FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
-- );


CREATE TABLE USERS (
    user_id TEXT PRIMARY KEY,
    user_name TEXT NOT NULL,
    email TEXT NOT NULL,
    firebase_uid TEXT UNIQUE NOT NULL
);

CREATE TABLE GROUPS (
    group_id TEXT PRIMARY KEY,
    group_name TEXT NOT NULL,
    owner_id TEXT NOT NULL,
    owner_name TEXT NOT NULL,
    bill_amount REAL DEFAULT 0,
    members_count INTEGER DEFAULT 0,
    unpaid_count INTEGER DEFAULT 0,
    currency TEXT DEFAULT 'INR',
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES USERS(user_id) ON DELETE CASCADE
);

CREATE TABLE GROUP_USERS (
    user_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    group_id TEXT NOT NULL,
    group_name TEXT NOT NULL,
    role TEXT CHECK(role IN ('owner', 'member')) DEFAULT 'member',
    payment_status TEXT CHECK(payment_status IN ('paid', 'unpaid')) DEFAULT 'unpaid',
    bill_amount REAL DEFAULT 0,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES USERS(user_id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id) ON DELETE CASCADE
);
