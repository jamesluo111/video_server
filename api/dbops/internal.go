package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbCoon.Prepare(`INSERT INTO sessions (session_id, ttl, login_name) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbCoon.Prepare(`SELECT ttl, login_name FROM sessions WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if ttlInt, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = ttlInt
		ss.UserName = uname
	} else {
		return nil, err
	}
	stmtOut.Close()
	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stsmOut, err := dbCoon.Prepare(`SELECT * FROM sessions`)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stsmOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
			log.Printf("retrieve all session %s", er)
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{TTL: ttl, UserName: login_name}
			m.Store(id, ss)
		}
	}
	stsmOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stsmIns, err := dbCoon.Prepare(`DELETE FROM sessions WHERE id = ?`)
	if err != nil {
		return err
	}
	if _, err = stsmIns.Query(sid); err != nil {
		return err
	}
	stsmIns.Close()
	return nil
}
