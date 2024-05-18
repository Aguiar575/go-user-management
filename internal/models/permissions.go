package models

type UserPermission struct {
	ID        int  `json:"id"`
	UserID    int  `json:"userId"`
	ContextID int  `json:"contextId"`
	Read      bool `json:"read"`
	Write     bool `json:"write"`
}
