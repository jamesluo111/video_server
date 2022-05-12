package dbops

import "fmt"

//从video删除表中删除指定记录

func GetDeleteVideoRecord(count int) ([]string, error) {
	stsmOut, err := dbCoon.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	var ids []string
	if err != nil {
		return ids, err
	}
	idsRow, err := stsmOut.Query(count)
	if err != nil {
		return ids, err
	}

	for idsRow.Next() {
		var id string
		if err := idsRow.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stsmOut.Close()
	return ids, nil
}

func DeleteVideoRecord(id string) error {
	stsmDel, err := dbCoon.Prepare("DELETE FROM video_del_rec WHERE video_id = ?")
	if err != nil {
		fmt.Println("=========删除视频失败！========")
		return err
	}

	_, err = stsmDel.Exec(id)
	if err != nil {
		fmt.Println("=========删除视频失败！！========")
		return err
	}
	defer stsmDel.Close()
	return nil
}
