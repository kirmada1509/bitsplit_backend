package crud

import (
	"bitsplit_backend/models"
    // "fmt"
)

// createGroup creates a new Group in the database
func (crud CRUD) CreateGroup(Group models.Group) error {
	query := `INSERT INTO groups (name, gid, owner_id) VALUES (?, ?, ?)`
	_, err := crud.DB.Exec(query, Group.Name, Group.GID, Group.OWNER_ID)
	return err
}

func (crud CRUD) GetGroupByID(GroupID string) (models.Group, error) {
	var Group models.Group
	query := `SELECT name, gid, owner_id FROM groups WHERE gid = ?`
	err := crud.DB.QueryRow(query, GroupID).Scan(&Group.GID, &Group.Name, &Group.OWNER_ID)
	return Group, err
}

func (crud CRUD) UpdateGroup(Group models.Group) error {
    query := `UPDATE groups SET name = ?, gid = ?, owner_id = ? WHERE gid = ?`
    _, err := crud.DB.Exec(query, Group.Name, Group.GID, Group.OWNER_ID, Group.GID)
    return err
}

func (crud CRUD) DeleteGroup(GroupID string) error {
    query := `DELETE FROM groups WHERE gid = ?`
    _, err := crud.DB.Exec(query, GroupID)
    return err
}

func (crud CRUD) GetAllGroups() ([]models.Group, error) {
    query := `SELECT id, name, gid, owner_id FROM groups`
    rows, err := crud.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var Groups []models.Group

    for rows.Next() {
        var Group models.Group
        if err := rows.Scan(&Group.ID, &Group.Name, &Group.GID, &Group.OWNER_ID); err != nil {
            return nil, err
        }
        Groups = append(Groups, Group)
    }

    return Groups, nil
}

func (crud CRUD) SearchInGroups(search_query string) ([]models.Group, error) {
	query := `SELECT id, name, gid, owner_id FROM groups WHERE name LIKE ?`
	searchTerm := "%" + search_query + "%"
	rows, err := crud.DB.Query(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var Groups []models.Group
	for rows.Next(){
		var Group models.Group
        if err := rows.Scan(&Group.ID, &Group.Name, &Group.GID, &Group.OWNER_ID); err != nil {
            return nil, err
        }
        Groups = append(Groups, Group)	
	}
	return Groups, nil
}

func (crud CRUD) GetGroupsByOwnerID(owner_id string) ([]models.Group, error) {
	query := `SELECT id, name, gid, owner_id FROM groups WHERE owner_id = ?`
	rows, err := crud.DB.Query(query, owner_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var Groups []models.Group
	for rows.Next(){
		var Group models.Group
        if err := rows.Scan(&Group.ID, &Group.Name, &Group.GID, &Group.OWNER_ID); err != nil {
            return nil, err
        }
        Groups = append(Groups, Group)	
	}
	return Groups, nil
}