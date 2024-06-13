package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

// func SearchUser(db *sql.DB, name string) ([]model.UserResForHTTPGet, error) {
// 	users, err := dao.GetUserByName(db, name)
// func SearchUser(db *sql.DB, uid string) ([]model.UserResForHTTPGet, error) {
func SearchUser(db *sql.DB, uid string) ([]model.UserResForHTTPGet, error) {
	// users, err := dao.GetUserByName(db, uid)
	users, err := dao.GetUserById(db, uid)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return nil, err
	}
	// return users, nil
	return users, nil
}