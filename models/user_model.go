package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    UID   string `json:"uid"`
}

type Group struct{
    ID int `json:"id`
    Name string `json:"name"`
    GID string `json:"gid"`
    OWNER_ID string `json:"owner_id"`
}

type GroupUser struct{
    ID int `json:"id"`
    UID string `json:"uid"`
    GID string `json:"gid"`
    STATUS string `json:"statsu"`
    IS_OWNER bool `json:"is_owner"`
}
