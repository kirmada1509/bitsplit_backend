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
    GID string `json:"tid"`
    OWNER_ID string `json:"owner_id"`
}