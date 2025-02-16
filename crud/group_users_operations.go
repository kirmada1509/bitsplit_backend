package crud

import (
	"bitsplit_backend/models"
	"fmt"
)

//add user to group

func (crud CRUD) AddUserToGroup(groupUser models.GroupUser) error {
	query := `INSERT INTO group_users (user_id, user_name, group_id, status, is_owner) VALUES (?, ?, ?, ?, ?)`
	_, err := crud.DB.Exec(query, groupUser.UID, groupUser.USER_NAME, groupUser.GID, groupUser.STATUS, groupUser.IS_OWNER)

	return err
}


func (crud CRUD) GetAllGroupUsers() ([]models.GroupUser, error){
	query := `SELECT * FROM group_users`
	rows, err := crud.DB.Query(query)
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

func (crud CRUD) GetGroupsUserIsIn(userID string) ([]models.Group, error) {
    ///TODO: checking for name, change it to based on id
	query := `
		SELECT g.id, g.name, g.GID, g.owner_id 
		FROM groups g
		JOIN group_users gu ON g.GID = gu.group_id
		WHERE gu.user_name = ?;
	`

	rows, err := crud.DB.Query(query, userID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group

	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.GID, &group.OWNER_ID); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}
