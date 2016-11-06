package game

// 抽奖游戏

import (
	"dao"
	"fmt"
	"math/rand"
	"time"
	"util"
)

const (
	GAME_GIFT_CHANCE = 4
	GAME_WINNER      = 1
)

//游戏礼品
type GameGift struct {
	Id       int
	GiftName string
	PicUrl   string
	GetCode  string
	GiftNum  int
}

//定义中奖概率
func getRand(i int) int {
	if i <= 0 {
		i = 2
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := (r.Intn(100) % i) + 1
	return num

}

func isHasGift() bool {
	db, err := dao.GameConnction()
	querySql := "select count(1) from game_gift where gift_num != 0"
	rows, err := dao.Select(db, querySql)
	var hasGift int
	for rows.Next() {
		err = rows.Scan(&hasGift)
		if err != nil {
			fmt.Println(err)
		}
	}
	dao.CloseDb(db)
	if hasGift > 0 {
		return true
	} else {
		return false
	}
}

//开始游戏
func PlayGiftgame(wxopenId string) GameGift {

	//判断是否还有奖品
	if !isHasGift() {
		return GameGift{GiftNum: 2}
	}

	giftcode := getRand(GAME_GIFT_CHANCE)

	if giftcode > GAME_WINNER {
		return queryGift(giftcode, wxopenId)
	}

	return GameGift{GiftNum: 0}
}

// 查询得到礼
func queryGift(giftcode int, wxopenId string) GameGift {

	db, err := dao.GameConnction()
	var id int
	var picUrl string
	var giftName string

	querySql := "select id , gift_name , gift_pic from game_gift where gift_num != 0"
	rows, err := dao.Select(db, querySql)

	if err != nil {
		fmt.Println("err ", err)
		dao.CloseDb(db)
		return GameGift{GiftNum: 2}
	}

	var count int
	mapGift := make(map[int]GameGift)

	for rows.Next() {
		count++
		err = rows.Scan(&id, &giftName, &picUrl)
		gift := GameGift{Id: id, GiftName: giftName, PicUrl: picUrl}
		mapGift[count] = gift
	}

	if count == 0 {
		dao.CloseDb(db)
		return GameGift{GiftNum: 2}
	}

	r := getRand(count)
	getGift := mapGift[r]

	//礼物数量减1
	updatesql := "update game_gift set gift_num=gift_num-1 , has_num=has_num+1 where id=?"
	stmt, _ := db.Prepare(updatesql)
	stmt.Exec(getGift.Id)

	//礼品保存
	getCode := string(util.Krand(6, util.KC_RAND_KIND_ALL))

	insertSql := "insert into game_usergift(wxopenid,gift_name,add_time,get_code) value(?,?,now(),?)"
	stmt, _ = db.Prepare(insertSql)
	stmt.Exec(wxopenId, getGift.GiftName, getCode)

	dao.CloseDb(db)
	getGift.GetCode = getCode
	getGift.GiftNum = 1

	return getGift

}
