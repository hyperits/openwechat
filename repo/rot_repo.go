package repo

import (
	"fmt"

	"gorm.io/gorm"
)

// LabelBaseRepo repo
type RotBaseRepo struct { // nolint
	db *gorm.DB
}

// NewLabelBaseRepo 创建实例
func NewRotBaseRepo(db *gorm.DB) *RotBaseRepo {
	return &RotBaseRepo{db}
}

func (r RotBaseRepo) GetRotNic() (string, error) {
	var nic []Nic
	if err := r.db.Table("nic").Find(&nic).Error; err != nil {
		return "", fmt.Errorf("GetRotNic err:%w", err)
	}
	if len(nic) > 0 {
		return nic[0].Nic, nil
	}
	return "", fmt.Errorf("GetRotNic nil")
}

func (r RotBaseRepo) InsertMsg(msg Msg) error {
	if err := r.db.Table("msg").Create(&msg).Error; err != nil {
		return fmt.Errorf("InsertMsg err:%w", err)
	}
	return nil
}

func (r RotBaseRepo) UpdateAccomplish(id int) error {
	if err := r.db.Table("msg").Where("id = ?", id).Updates(map[string]interface{}{
		"is_done": 1,
	}).Error; err != nil {
		return fmt.Errorf("InsertMsg err:%w", err)
	}
	return nil
}
