package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    UID   string `json:"uid"`
}

type Group struct{
    ID int `json:"id"`
    Name string `json:"name"`
    GID string `json:"gid"`
    OWNER_ID string `json:"owner_id"`
}

type GroupUser struct{
    ID int `json:"id"`
    UID string `json:"uid"`
    USER_NAME string `json:"user_name"`
    GID string `json:"gid"`
    STATUS string `json:"status"`
    IS_OWNER bool `json:"is_owner"`
}
