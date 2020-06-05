package handlers

import (
	"database/sql"
	"time"
	"video-server/api/database"
	"video-server/api/model"
	"video-server/api/utils"
)

func AddNewVedio(aid int, name string) (*model.VideoInfo, error) {
	//create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := database.DBConn.Prepare("INSERT INTO video_info(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &model.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*model.VideoInfo, error) {
	stmtOut, err := database.DBConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")
	if err != nil {
		return nil, err
	}

	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &model.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: dct,
	}

	return res, nil
}

func DeleteVedioInfo(vid string) error {
	stmtDel, err := database.DBConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
