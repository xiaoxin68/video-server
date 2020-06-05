package handlers

import (
	"video-server/api/database"
	"video-server/api/model"
	"video-server/api/utils"
)

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtInts, err := database.DBConn.Prepare("INSERT INTO comments(id,video_id,author_id,content) values (?,?,?,?)")
	if err != err {
		return nil
	}

	_, err = stmtInts.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtInts.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*model.Comment, error) {
	stmtOut, err := database.DBConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*model.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, nil
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &model.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}

	defer stmtOut.Close()
	return res, nil
}
