package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `josn:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	//db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	db.Model(&Tag{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	//db.Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name =? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id =? AND deleted_on = ?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id=? AND deleted_on = ?", id, 0).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id=?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}
