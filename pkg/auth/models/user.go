package models

type User struct {
	UserID     uint
	UserName   string
	Email      string
	Mobile     string
	Password   string
	AvatarID   string
	Provider   string
	LastLogin  string
	Permission bool
	Interests  []uint
	Referel    string
}

type Interests struct {
	Interest   string
	InterestID string
}

type Connection struct {
	UserID          uint
	ConnectedUserID uint
	ConnectionDate  string
}
