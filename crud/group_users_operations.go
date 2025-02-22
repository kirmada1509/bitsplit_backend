package crud

import (
	"bitsplit_backend/models"
)

func (crud CRUD) AddUserToGroup(groupUser models.GroupUser) error {
	query := `INSERT INTO GROUP_USERS (UserID, user_name, GroupID, group_name, role, payment_status, bill_amount) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := crud.DB.Exec(query,
		groupUser.UserID, groupUser.UserName, groupUser.GroupID, groupUser.GroupName,
		groupUser.Role, groupUser.PaymentStatus, groupUser.BillAmount)

	return err
}

func (crud CRUD) GetAllGroupUsers() ([]models.GroupUser, error) {
	query := `SELECT UserID, user_name, GroupID, group_name, role, payment_status, bill_amount FROM GROUP_USERS`

	rows, err := crud.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupUsers []models.GroupUser
	for rows.Next() {
		var groupUser models.GroupUser
		err := rows.Scan(
			&groupUser.UserID, &groupUser.UserName, &groupUser.GroupID,
			&groupUser.GroupName, &groupUser.Role, &groupUser.PaymentStatus,
			&groupUser.BillAmount,
		)
		if err != nil {
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groupUsers, nil
}

func (crud CRUD) GetUsersInGroup(GroupID string) ([]models.GroupUser, error) {
	query := `SELECT UserID, user_name, GroupID, group_name, role, payment_status, bill_amount 
	          FROM GROUP_USERS WHERE GroupID = ?`

	rows, err := crud.DB.Query(query, GroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupUsers []models.GroupUser
	for rows.Next() {
		var groupUser models.GroupUser
		err := rows.Scan(
			&groupUser.UserID, &groupUser.UserName, &groupUser.GroupID,
			&groupUser.GroupName, &groupUser.Role, &groupUser.PaymentStatus,
			&groupUser.BillAmount,
		)
		if err != nil {
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groupUsers, nil
}

func (crud CRUD) GetGroupsUserIsIn(UserID string) ([]models.Group, error) {
	query := `SELECT g.GroupID, g.group_name, g.owner_id, g.owner_name, g.bill_amount, 
	                 g.members_count, g.unpaid_count, g.currency, g.description, g.created_at
	          FROM GROUPS g
	          INNER JOIN GROUP_USERS gu ON g.GroupID = gu.GroupID
	          WHERE gu.UserID = ?`

	rows, err := crud.DB.Query(query, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.GroupID, &group.GroupName, &group.OwnerID, &group.OwnerName,
			&group.BillAmount, &group.MembersCount, &group.UnpaidCount,
			&group.Currency, &group.Description, &group.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
