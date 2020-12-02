package dao

type User struct {
 uid int64
}

func GetUsers(userIdList []int64) ([]User, error) {
	var userList []ImDeviceToken
	err := DbManager.Table("user").Select(columns).Where("uid in (?)", userIdList).Find(&userList).Error
	if err == sql.ErrNoRows {
		return userList,nil
	}
	if err != nil {
		return nil, fmt.Errorf("GetUser fails error : %w", err)
	}
	return userList, err
}
