package schema

import "time"

type Personas struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	UserId    string    `json"user_id"`
	FirstName string    `json"first_name"`
	LastName  string    `json"last_name"`
	Room      string    `json"room"`
	VillageId string    `json"village_id"`
	Status    string    `json"status"`
	RoleId    string    `json"role_id"`
	CreateAt  time.Time `json"create_at"`
}

type Users struct {
	Id       string `json"id"`
	UserId   string `json"user_id"`
	Password string `json"password"`
}

type Roles struct {
	Id       string `json"id"`
	RoleId   string `json"role_id"`
	RoleDesc string `json"role_desc"`
}

type Villages struct {
	Id          string    `json"id"`
	VillageId   string    `json"village_id"`
	villageName string    `json"village_name"`
	villageDesc string    `json"village_desc"`
	CreateAt    time.Time `json"create_at"`
}
