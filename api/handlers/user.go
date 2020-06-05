package handlers

import (
	"database/sql"
	"log"
	"video-server/api/database"
)

func AddUserCredetial(loginName string, pwd string) error {
	stmeIns, err := database.DBConn.Prepare("INSERT INTO users(login_name,pwd) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = stmeIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmeIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := database.DBConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := database.DBConn.Prepare("DELETE FROM users WHERE login_name= ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
