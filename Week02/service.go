package services

import "log"

type usr interface {

}
func getUsers(uids []int64) {
  userList, err := usrInstance.GetUsers(uids)
  if err != nil {
    log.Println(err)
  }
  
  //todo 
  fmt.Println(userList)

}
