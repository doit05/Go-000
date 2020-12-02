package services

import "log"

type usr interface {
  GetUsers(userIdList []int64) ([]User, error)
}

func getUsers(uids []int64) {
  userList, err := usrInstance.GetUsers(uids)
  if err != nil {
    log.Println(err)
  }
  
  //todo 
  fmt.Println(userList)

}
