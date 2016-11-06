package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 数据库连接常量
const (
	USER      = ""
	PASSWD    = ""
	HOST      = ""
	PORT      = ""
	DBNAME    = ""
	DBCHARSET = "UTF8"

	GAME_DBNAME = "xiangya_weixin"
)

// 微信游戏的数据库
func GameConnction() (*sql.DB, error) {
	db, err := sql.Open("mysql", USER+":"+PASSWD+"@tcp("+HOST+":"+PORT+")/"+GAME_DBNAME+"?charset="+DBCHARSET)
	if err != nil {
		log.Fatal("Wechat Service : connectdb err : ", err)
		return db, err
	}
	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(10)
	db.Ping()
	return db, err
}

// 初始化象伢问答数据连接
func GetConnction() (*sql.DB, error) {
	db, err := sql.Open("mysql", USER+":"+PASSWD+"@tcp("+HOST+":"+PORT+")/"+DBNAME+"?charset="+DBCHARSET)
	if err != nil {
		log.Fatal("Wechat Service : connectdb err : ", err)
		return db, err
	}
	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(10)
	db.Ping()
	return db, err
}

// 关闭数据库
func CloseDb(db *sql.DB) {
	db.Close()
}

// 数据查询
func Select(db *sql.DB, sql string) (*sql.Rows, error) {
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("Wechat Service : Query sql err :", err)
	}
	return rows, err
}

// 数据库插入操作
func Insert(db *sql.DB, sql string) bool {
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal("Wechat Service : insert action err : ", err)
		return false
	}
	defer stmt.Close()
	return true
}
