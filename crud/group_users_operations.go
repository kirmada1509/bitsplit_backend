package crud

import (
	"bitsplit_backend/models"
)

//add user to group

func (crud CRUD) AddUserToGroup(groupUser models.GroupUser) error {
	query := `INSERT INTO group_users (user_id, user_name, group_id, status, is_owner) VALUES (?, ?, ?, ?, ?)`
	_, err := crud.DB.Exec(query, groupUser.UID, groupUser.USER_NAME, groupUser.GID, groupUser.STATUS, groupUser.IS_OWNER)

	return err
}


func (crud CRUD) GetUsersInGroup(groupId string) ([]models.GroupUser, error){
	query := `SELECT * FROM group_users WHERE group_id = ?`
	rows, err := crud.DB.Query(query, groupId)
    if err != nil {
        print(err.Error())
        return nil, err
    }
    defer rows.Close()

    var users []models.GroupUser

    for rows.Next() {
        var user models.GroupUser
        if err := rows.Scan(&user.ID, &user.UID, &user.USER_NAME, &user.GID, &user.STATUS, &user.IS_OWNER); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}