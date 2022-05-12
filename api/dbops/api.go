package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbCoon.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbCoon.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != err {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func GetUser(username string) (int, error) {
	stmtOut, err := dbCoon.Prepare("SELECT id FROM users WHERE login_name = ?")
	if err != err {
		log.Printf("%s", err)
		return 0, err
	}
	var id int
	stmtOut.QueryRow(username).Scan(&id)
	stmtOut.Close()
	return id, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbCoon.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != err {
		log.Printf("%s", err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uuid
	videoId := utils.NewUuid()
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbCoon.Prepare(`INSERT INTO video_info (id, author_id, name, display_ctime) value (?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(videoId, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: videoId, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	//获取sql模板
	stmtOut, err := dbCoon.Prepare(`SELECT author_id, name, display_ctime FROM video_info WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	//定义接收参数
	var authorId int
	var name string
	var displayCtime string
	err = stmtOut.QueryRow(vid).Scan(&authorId, &name, &displayCtime)
	//判断是否报错并且获取的数据是否为空
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: authorId, Name: name, DisplayCtime: displayCtime}
	defer stmtOut.Close()
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtIns, err := dbCoon.Prepare(`DELETE FROM video_info WHERE id = ?`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}

func AddNewComment(vid string, aid int, content string) error {
	uuidString := utils.NewUuid()
	stmtIns, err := dbCoon.Prepare(`INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(uuidString, vid, aid, content)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentList(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbCoon.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
											INNER JOIN users ON comments.author_id = users.id
											WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
											ORDER BY comments.time DESC`)
	if err != nil {
		return nil, err
	}

	var commentArr []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return commentArr, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return commentArr, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		commentArr = append(commentArr, c)
	}
	defer stmtOut.Close()

	return commentArr, nil
}

func GetVideoList(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbCoon.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info
											INNER JOIN users ON video_info.author_id = users.id
											WHERE users.login_name = ?`)
	if err != nil {
		return nil, err
	}
	var videoInfo []*defs.VideoInfo
	rows, err := stmtOut.Query(uname)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, display string
		var authorId int
		rows.Scan(&id, &authorId, &name, &display)
		c := &defs.VideoInfo{Id: id, AuthorId: authorId, Name: name, DisplayCtime: display}
		videoInfo = append(videoInfo, c)
	}
	return videoInfo, nil
}
