package dbops

func TaskAddVideoDelete(vid string) error {
	stsmIns, err := dbCoon.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = stsmIns.Exec(vid)
	if err != nil {
		return err
	}
	defer stsmIns.Close()
	return nil
}
