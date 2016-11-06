-- 抽奖游戏
CREATE TABLE `game_gift` (
      `id` int(10) NOT NULL AUTO_INCREMENT,
      `gift_name` varchar(50) NOT NULL,
      `gift_num` int(10) DEFAULT '0',
      `has_num` int(10) DEFAULT '0',
      `gift_pic` varchar(200) DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ; 

-- 获奖用户
CREATE TABLE `game_usergift` (
      `id` int(10) NOT NULL AUTO_INCREMENT,
      `username` varchar(50) DEFAULT NULL,
      `wxopenid` varchar(60) NOT NULL,
      `gift_name` varchar(50) NOT NULL,
      `add_time` varchar(20) NOT NULL,
      `get_code` varchar(10) NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ; 

-- 用户信息
CREATE TABLE `wx_user_message`(
    `id` int(10) NOT NULL AUTO_INCREMENT , 
)

