package model

import (
	"fmt"
	Serviceimplementation "project4/serviceimplentation"
	"project6/Notification/dao"
	"project6/Notification/entities"
	Services "project6/Notification/services"
)

type UserModel struct {
}

func GetNotificationMode(notificationMode string) Services.Notification {

	switch notificationMode {
	case "CallNotificationMode":
		return &Serviceimplementation.CallNotification{}

	case "EmailNotificationMode ":
		return &Serviceimplementation.EmailNotification{}

	case "MessageNotificationMode":
		return &Serviceimplementation.MessageNotification{}

	default:
		fmt.Println("Invalid mode chosen")
		return nil

	}
}

func (*UserModel) FindAll() ([]entities.User, error) {
	db, err := dao.GetDB()
	if err != nil {

		fmt.Println(err)
		return nil, err
	} else {
		rows, err2 := db.Query("select * from user")
		if err2 != nil {
			return nil, err2
		} else {
			var users []entities.User
			for rows.Next() {
				var User entities.User
				rows.Scan(&User.Id, &User.Name, &User.Email, &User.Mobile)
				users = append(users, User)
			}
			return users, nil
		}
	}
}

func (*UserModel) Create(Users *entities.User) bool {
	fmt.Println("hello")
	db, err := dao.GetDB()

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		result, err2 := db.Exec("insert into user(name, email, mobile) values(?,?,?)", Users.Name, Users.Email, Users.Mobile)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
