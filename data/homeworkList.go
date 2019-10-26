package data

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Homework struct {
	ID      string
	Ta      Ta `gorm:"foreignkey:TaRefer"`
	TaRefer int64
}

func NewHomework(ta *Ta) uuid.UUID {
	tuid := uuid.New()
	Db.Create(&Homework{TaRefer: ta.ID, ID: tuid.String()})
	return tuid
}

func HomeWordCheckValid(tid int64, uid string) (bool, error) {

	hw := []Homework{}
	fmt.Println(uid, tid)
	Db.Where("Ta_Refer=? AND ID=?", tid, uid).Find(&hw)
	fmt.Println(hw)
	if len(hw) != 1 {
		return false, nil
	}
	return true, nil

}
