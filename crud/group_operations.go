package crud

import (
	"bitsplit_backend/models"
	// "fmt"
)

func (crud CRUD) CreateGroup(Group models.Group) error {
	query := `INSERT INTO GROUPS (group_id, group_name, owner_id, owner_name, bill_amount, members_count, unpaid_count, currency, description, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := crud.DB.Exec(query, Group.GroupID, Group.GroupName, Group.OwnerID, Group.OwnerName, Group.BillAmount, Group.MembersCount, Group.UnpaidCount, Group.Currency, Group.Description, Group.CreatedAt)
	return err
}

func (crud CRUD) GetGroupByID(GroupID string) (models.Group, error) {
	var group models.Group
	query := `SELECT group_id, group_name, owner_id, owner_name, bill_amount, 
						members_count, unpaid_count, currency, description, created_at 
			  FROM groups WHERE group_id = ?`
	err := crud.DB.QueryRow(query, GroupID).Scan(
		&group.GroupID, &group.GroupName, &group.OwnerID, &group.OwnerName,
		&group.BillAmount, &group.MembersCount, &group.UnpaidCount,
		&group.Currency, &group.Description, &group.CreatedAt,
	)
	return group, err
}

func (crud CRUD) GetAllGroups() ([]models.Group, error) {
	query := `SELECT group_id, group_name, owner_id, owner_name, bill_amount, 
						members_count, unpaid_count, currency, description, created_at 
			  FROM groups`

	rows, err := crud.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		var description *string

		err := rows.Scan(
			&group.GroupID, &group.GroupName, &group.OwnerID, &group.OwnerName,
			&group.BillAmount, &group.MembersCount, &group.UnpaidCount,
			&group.Currency, &description, &group.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description != nil {
			group.Description = *description
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (crud CRUD) SearchInGroups(searchQuery string) ([]models.Group, error) {
	query := `SELECT group_id, group_name, owner_id, owner_name, bill_amount, 
						members_count, unpaid_count, currency, description, created_at 
			  FROM groups
			  WHERE group_name LIKE ? OR owner_name LIKE ? OR description LIKE ?`

	searchPattern := "%" + searchQuery + "%"

	rows, err := crud.DB.Query(query, searchPattern, searchPattern, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		var description *string

		err := rows.Scan(
			&group.GroupID, &group.GroupName, &group.OwnerID, &group.OwnerName,
			&group.BillAmount, &group.MembersCount, &group.UnpaidCount,
			&group.Currency, &description, &group.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description != nil {
			group.Description = *description
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (crud CRUD) GetGroupsByOwnerID(ownerID string) ([]models.Group, error) {
	query := `SELECT group_id, group_name, owner_id, owner_name, bill_amount, 
						members_count, unpaid_count, currency, description, created_at 
			  FROM groups
			  WHERE owner_id = ?`

	rows, err := crud.DB.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		var description *string

		err := rows.Scan(
			&group.GroupID, &group.GroupName, &group.OwnerID, &group.OwnerName,
			&group.BillAmount, &group.MembersCount, &group.UnpaidCount,
			&group.Currency, &description, &group.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description != nil {
			group.Description = *description
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
