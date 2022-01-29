package infrastructure

import (
	"errors"
	"prac-orm-transaction/domain/repository"
	"prac-orm-transaction/infrastructure/table"
)

type userPersistence struct {
	mysql mysqlRepository
}

func NewUserPersistence(mysqlConn mysqlRepository) repository.UserRepository {
	return &userPersistence{mysql: mysqlConn}
}

func (m *userPersistence) CreateUsersAccount(id, pass, token string) error {
	usersInfo := table.UserInfo{ID: id, Password: pass, Token: token}
	var dataExistsCheck table.UserInfo

	m.mysql.Client.First(&dataExistsCheck, "id=?", usersInfo.ID)
	if dataExistsCheck.ID != "" {
		return errors.New("this userID is already registered")
	}

	if err:=m.mysql.Client.Create(&usersInfo).Error; err != nil {
		return err
	}

	return nil
}

func (u *userPersistence) RegisterUsersInfo() {
	return
}

//m.mysql.Client.Transaction(func(tx *gorm.DB) error {
//	//if err := tx.Create(&usersInfo).Error; err != nil {
//	//	return err
//	//}
//	tx.First(&dataExistsCheck, "id=?", usersInfo.ID)
//	if dataExistsCheck.ID != "" {
//		return errors.New("this userID is already registered")
//	}
//	tx.Create(&usersInfo)
//	tx.First(&dataExistsCheck, "id=?", usersInfo.ID)
//	if dataExistsCheck.ID == "" {
//		return errors.New("new account data insert is error")
//	}
//	// nilが返却されるとトランザクション内の全処理がコミットされる
//	return nil
//}).Error()
