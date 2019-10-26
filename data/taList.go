package data

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Ta struct {
	ID       int64  `gorm:"AUTO_INCREMENT;primary_key"`
	Name     string `gorm:"type:varchar(100); not null; unique"`
	PassWord string `gorm:"type:varchar(100); not null"`
}

func Register(name, passwd string) bool {
	tas := []Ta{}
	Db.Where("Name=?", name).Find(&tas)
	if len(tas) > 0 {
		return false
	}
	nta := Ta{Name: name, PassWord: passwd}
	Db.Create(&nta)
	return true
}

func Login(name, passwd string) bool {
	tas := []Ta{}
	Db.Where("Name = ? AND pass_word = ?", name, passwd).Find(&tas) // 这里的变量要对应数据库的，也就是不能PassWord而是pass_word
	if len(tas) != 1 {
		return false
	}
	return true
}

func PublichHomework(name, passwd string) string {
	tas := []Ta{}
	Db.Where("Name = ? AND Pass_Word = ?", name, passwd).Find(&tas)
	if len(tas) != 1 {
		return ""
	}
	tuid := NewHomework(&tas[0])
	return tuid.String()
}

func GetHomework(name, passwd, uid string) bool {
	ta := Ta{}
	Db.Where("name=? AND pass_word = ?", name, passwd).First(&ta)
	tid := ta.ID
	if tid == 0 {
		return false
	}
	bol, _ := HomeWordCheckValid(tid, uid)
	return bol
}
