package table

//-------------------------------------//
//SQL struct
//-------------------------------------//

type Users struct {
	ID   string `gorm:"primaryKey"`
	Age  int
	Name string
}

type UserInfo struct {
	ID       string `gorm:"primaryKey"`
	Password string
	Token    string
}
