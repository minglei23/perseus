package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open(
		"mysql", MysqlUser+":"+MysqlPassword+"@tcp("+MysqlAddress+":"+MysqlPort+")/"+MysqlDB,
	)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
}

// POINTS API

func GetPoints(id int) (points int, err error) {
	query := "SELECT points FROM user_info WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&points)
	return points, err
}

func GetIfUserCheckedToday(id int) (checked bool, err error) {
	var count int
	query := "SELECT COUNT(*) FROM activities WHERE user_id = ? AND type = 1 AND DATE(time) = CURDATE()"
	err = db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertActivities(id int, activity int, points int) error {
	query := "INSERT INTO activities (user_id, type, points) VALUES (?, ?, ?)"
	_, err := db.Exec(query, id, activity, points)
	return err
}

func UpdateUserPoints(id int, points int) error {
	query := "UPDATE user_info SET points = ? WHERE id = ?"
	_, err := db.Exec(query, points, id)
	return err
}

// ADMIN API

func InsertVideo(name string, videoType int, totalNumber int, baseUrl string) (id int, err error) {
	query := "INSERT INTO video_info (name, type, total_number, base_url) values (?, ?, ?, ?)"
	result, err := db.Exec(query, name, videoType, totalNumber, baseUrl)
	if err != nil {
		return id, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		return id, err
	}
	return int(ID), nil
}

// LOGIN API

func GetUserIdByEmailAndPassword(email, password string) (id int, activated, vip bool, err error) {
	query := "SELECT id, activated, vip FROM user_info WHERE email = ? AND password = ?"
	err = db.QueryRow(query, email, password).Scan(&id, &activated, &vip)
	if err == sql.ErrNoRows {
		// email or password is not correct
		return -1, false, false, nil
	}
	return id, activated, vip, err
}

// REGISTER API

func EmailExist(email string) (bool, error) {
	var id int
	query := "SELECT id FROM user_info WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}

func InsertUser(password, email string) (id int, err error) {
	query := "INSERT INTO user_info (password, email, activated, vip, points) values (?, ?, FALSE, FALSE, 0)"
	result, err := db.Exec(query, password, email)
	if err != nil {
		return id, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		return id, err
	}
	return int(ID), nil
}

// RESET API

func GetUserIdByEmail(email string) (id int, err error) {
	query := "SELECT id FROM user_info WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&id)
	if err == sql.ErrNoRows {
		// email is not correct
		return -1, nil
	}
	return id, err
}

func UpdateUserPassword(id, password string) error {
	query := "UPDATE user_info SET password = ? WHERE id = ?"
	_, err := db.Exec(query, password, id)
	return err
}

// VERIFY API

func ActivateUser(id string) error {
	query := "UPDATE user_info SET activated = TRUE WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

// VIDEO API

type Video struct {
	ID          int
	Name        string
	Type        int
	TotalNumber int
	BaseURL     string
	Hot         bool
}

type History struct {
	ID          int
	Name        string
	Episode     int
	TotalNumber int
	BaseURL     string
	Hot         bool
}

func buildVideoList(rows *sql.Rows) ([]Video, error) {
	defer rows.Close()
	var videos []Video
	for rows.Next() {
		var v Video
		if err := rows.Scan(&v.ID, &v.Name, &v.Type, &v.TotalNumber, &v.BaseURL, &v.Hot); err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return videos, nil
}

func buildHostoryList(rows *sql.Rows) ([]History, error) {
	defer rows.Close()
	var histories []History
	for rows.Next() {
		var v History
		if err := rows.Scan(&v.ID, &v.Name, &v.Episode, &v.TotalNumber, &v.BaseURL, &v.Hot); err != nil {
			return nil, err
		}
		histories = append(histories, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return histories, nil
}

func GetVideoList() ([]Video, error) {
	rows, err := db.Query("SELECT id, name, type, total_number, base_url, hot FROM video_info")
	if err != nil {
		return nil, err
	}
	return buildVideoList(rows)
}

func GetUserLike(userId int) ([]Video, error) {
	rows, err := db.Query("SELECT video_info.id, video_info.name, video_info.type, video_info.total_number, video_info.base_url, video_info.hot FROM user_like INNER JOIN video_info ON user_like.video_id = video_info.id WHERE user_like.user_id = ? ORDER BY user_like.watch_time DESC", userId)
	if err != nil {
		return nil, err
	}
	return buildVideoList(rows)
}

func GetUserHistoryLstMonth(userId int) ([]History, error) {
	after := time.Now().AddDate(0, -1, 0)
	rows, err := db.Query("SELECT video_info.id, video_info.name, user_history.episode, video_info.total_number, video_info.base_url, video_info.hot FROM user_history INNER JOIN video_info ON user_history.video_id = video_info.id WHERE user_history.user_id = ? AND user_history.watch_time > ? ORDER BY user_history.watch_time DESC", userId, after)
	if err != nil {
		return nil, err
	}
	return buildHostoryList(rows)
}

func InsertUserLike(userId int, videoId int) error {
	_, err := db.Exec("DELETE FROM user_like WHERE user_id = ? AND video_id = ?", userId, videoId)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO user_like (user_id, video_id) VALUES (?, ?)", userId, videoId)
	return err
}

func DeleteUserLike(userId int, videoId int) error {
	_, err := db.Exec("DELETE FROM user_like WHERE user_id = ? AND video_id = ?", userId, videoId)
	return err
}

func InsertUserHistory(userId int, videoId int, episode int) error {
	_, err := db.Exec("DELETE FROM user_history WHERE user_id = ? AND video_id = ?", userId, videoId)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO user_history (user_id, video_id, episode) VALUES (?, ?, ?)", userId, videoId, episode)
	return err
}
