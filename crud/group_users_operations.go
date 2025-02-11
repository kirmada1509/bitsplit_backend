package crud

import (
	"bitsplit_backend/models"
)


//add user to group

func (crud CRUD) AddUserToGroup(groupUser models.GroupUser) error {
	query := `INSERT INTO group_users (user_id, group_id, status, is_owner) VALUES (?, ?, ?, ?)`
	_, err := crud.DB.Exec(query, groupUser.UID, groupUser.GID, groupUser.STATUS, groupUser.IS_OWNER)
	return err
}
