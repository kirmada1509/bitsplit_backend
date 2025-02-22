package models

import "time"

type User struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	FirebaseUID string `json:"firebase_uid"`
}

type Group struct {
	GroupID      string
	GroupName    string
	OwnerID      string
	OwnerName    string
	BillAmount  float64
	MembersCount int
	UnpaidCount  int
	Currency     string
	Description  string // If nullable, use *string
	CreatedAt    time.Time
}

type GroupUser struct {
	UserID        string  `json:"user_id"`
	UserName      string  `json:"user_name"`
	GroupID       string  `json:"group_id"`
	GroupName     string  `json:"group_name"`
	Role          string  `json:"role"`
	PaymentStatus string  `json:"payment_status"`
	BillAmount   float64 `json:"bill_amount"`
}
