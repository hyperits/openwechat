package openwechat

import (
	"fmt"

	"github.com/hyperits/openwechat/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repositories TODO
type Repositories struct {
	db      *gorm.DB
	RotRepo *repo.RotBaseRepo
}

// ToDSNString TODO
func ToDSNString(user, passwd, ip, dataBase string, port int) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2fShanghai", user, passwd, ip, port, dataBase)
}

// NewRepositories TODO
func NewRepositories(info *ConfInfo) *Repositories {
	db, err := gorm.Open(mysql.Open(ToDSNString(info.User, info.Passw, info.IP, "wechatRot", info.Port)))
	if err != nil {
		return nil
	}
	return &Repositories{db: db,
		RotRepo: repo.NewRotBaseRepo(db),
	}
}
