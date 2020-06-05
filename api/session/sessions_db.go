package session

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video-server/api/database"
	"video-server/api/model"
)

//插入
func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := database.DBConn.Prepare("INSERT INTO sessions(session_id,TTL,login_name) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

//根据session_id查询
func RetrieveSession(sid string) (*model.SimpleSession, error) {
	ss := &model.SimpleSession{}
	stmtOut, err := database.DBConn.Prepare("SELECT TTL,login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	var ttl, uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

//查询所有
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}

	stmtOut, err := database.DBConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err != nil {
			ss := &model.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

//删除session
func DeleteSession(sid string) error {
	stmtOut, err := database.DBConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
