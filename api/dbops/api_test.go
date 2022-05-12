package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempVid string

func clearTable() {
	dbCoon.Exec("TRUNCATE users")
	dbCoon.Exec("TRUNCATE video_info")
	dbCoon.Exec("TRUNCATE comments")
	dbCoon.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}

//func TestUserWorkFlow(t *testing.T) {
//	t.Run("ADD", testAddUser)
//	t.Run("GET", testGetUser)
//	t.Run("DEL", testDeleteUser)
//	t.Run("REGET", testReGetUser)
//}

func testAddUser(t *testing.T) {
	err := AddUserCredential("luo", "123")
	if err != nil {
		t.Errorf("error add user: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("luo")
	if err != nil {
		t.Errorf("error get user: %v", err)
	}
	fmt.Println("密码为：", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("luo", "123")
	if err != nil {
		t.Errorf("error delete user: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("luo")
	if err != nil {
		t.Errorf("error reget user: %v", err)
	}
	if pwd != "" {
		t.Errorf("error reget user")
	}
}

//func TestVideoWorkFlow(t *testing.T) {
//	clearTable()
//	t.Run("ADDUSER", testAddUser)
//	t.Run("ADD", testAddVideo)
//	t.Run("GET", testGetVideo)
//	t.Run("DEL", testDeleteVideo)
//	t.Run("REGET", testReGetVideo)
//}

func testAddVideo(t *testing.T) {
	video, err := AddNewVideo(1, "shipin")
	if err != nil {
		t.Errorf("error add video: %v", err)
	}
	fmt.Printf("add video success: %v\n", video)
	tempVid = video.Id
}

func testGetVideo(t *testing.T) {
	video, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("error get video: %v", err)
	}
	if video != nil {
		fmt.Printf("get video success: %v\n", video)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("error delete video: %v", err)
	}
	fmt.Println("delete video success")
}

func testReGetVideo(t *testing.T) {
	video, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("error get video: %v", err)
	}
	if video == nil {
		fmt.Printf("reGet video success: %v\n", video)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTable()
	t.Run("ADDUSER", testAddUser)
	t.Run("ADD", testAddComment)
	t.Run("GET", testGetCommentList)
}

func testAddComment(t *testing.T) {
	err := AddNewComment("123", 1, "456")
	if err != nil {
		t.Errorf("add new comment error:%v", err)
	}
}

func testGetCommentList(t *testing.T) {
	vid := "123"
	from := 1514764800
	//将当前时间戳十进制int类型转化为string类型再转化为int类型
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	comments, err := GetCommentList(vid, from, to)
	if err != nil {
		t.Errorf("get commentlist error:%v", err)
	}
	for i, ele := range comments {
		fmt.Printf("i=%d, elements=%v\n", i, ele)
	}
}

func TestNowTime(t *testing.T) {
	nowTime1 := time.Now().UnixMilli()
	nowTime2 := time.Now().UnixNano() / 1000000
	fmt.Println(nowTime1)
	fmt.Println(nowTime2)
}
