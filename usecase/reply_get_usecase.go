package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

// func GetReply(db *sql.DB, parent_id string, posted_by string) ([]model.ReplyResGet, error) {
// func GetReply(db *sql.DB, parent_id string) ([]model.ReplyResGet, error) {
// 	// users, err := dao.GetReplyById(db, parent_id, posted_by)
// 	users, err := dao.GetReplyById(db, parent_id)
// 	if err != nil {
// 		log.Printf("fail: GetReplyById, %v\n", err)
// 		return nil, err
// 	}
// 	return users, nil
// }

func GetReply(db *sql.DB, parent_id string, uid string) ([]model.ReplyReturn, error) {
	replies, err := dao.GetAllReplyById(db)
	if err != nil {
		log.Printf("fail: GetAllReply, %v\n", err)
		return nil, err
	}
	likes, err := dao.GetStatusById(db, uid)
	if err != nil {
		log.Printf("fail: GetStatusById, %v\n", err)
		return nil, err
	}

	likesMap := make(map[string]bool)
	for _, id := range likes {
		likesMap[id] = true
	}

	var reply_ret []model.ReplyReturn
	for _, reply := range replies {
		var status int
		if likesMap[reply.Id] {
			status = 1
		} else {
			status = 0
		}
	
		// replyReturn を作成して新しい配列に追加する
		replyReturn := model.ReplyReturn{
			Id:      reply.Id,
			Name:    reply.Name,
			Time:    reply.Time,
			Content: reply.Content,
			Likes:   reply.Likes,
			Status:  status,
			Parent_Id: reply.Parent_Id,
			Display_name: reply.Display_name,
		}
	
		reply_ret = append(reply_ret, replyReturn)
	}
	return reply_ret, nil
}