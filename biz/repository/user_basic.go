package repository

import (
	"chat/biz/domain"
	"chat/common"
	"chat/utils"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type userBasicRepo struct {
	db *gorm.DB
}

func (u userBasicRepo) GetList(ctx context.Context, name, phone, email string, pageIdx, pageSize int) *[]domain.UserBasic {
	db := u.db.WithContext(ctx)
	data := make([]domain.UserBasic, pageSize)
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if phone != "" {
		db = db.Where("phone like ?", "%"+phone+"%")
	}
	if email != "" {
		db = db.Where("email like ?", "%"+email+"%")
	}
	if pageIdx < 0 {
		db.Find(&data)
	}
	db.Limit(pageSize).Offset((pageIdx - 1) * pageSize).Find(&data)
	return &data

}
func (u userBasicRepo) FindByNameAndPwd(ctx context.Context, name string, password string) domain.UserBasic {
	db := u.db.WithContext(ctx)
	user := domain.UserBasic{}
	db.Where("name = ? and pass_word=?", name, password).First(&user)

	//token 加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	db.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}
func (u userBasicRepo) FindByName(ctx context.Context, name string) domain.UserBasic {
	db := u.db.WithContext(ctx)
	user := domain.UserBasic{}
	db.Where("name = ?", name).First(&user)
	return user
}
func (u userBasicRepo) FindByPhone(ctx context.Context, phone string) domain.UserBasic {
	db := u.db.WithContext(ctx)
	user := domain.UserBasic{}
	db.Where("Phone = ?", phone).First(&user)
	return user
}
func (u userBasicRepo) FindByEmail(ctx context.Context, email string) domain.UserBasic {
	db := u.db.WithContext(ctx)
	user := domain.UserBasic{}
	db.Where("email = ?", email).First(&user)
	return user

}
func (u userBasicRepo) Create(ctx context.Context, user domain.UserBasic) error {
	db := u.db.WithContext(ctx)
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf(common.DBError, err)
	}
	return nil
}
func (u userBasicRepo) Delete(ctx context.Context, id uint) error {
	db := u.db.WithContext(ctx)
	if err := db.Table((&domain.UserBasic{}).TableName()).Where("id =?", id).Error; err != nil {
		return fmt.Errorf(common.DBError, err)
	}
	return nil
}
func (u userBasicRepo) Update(ctx context.Context, user domain.UserBasic) error {
	db := u.db.WithContext(ctx)
	if err := db.Updates(&user).Error; err != nil {
		return fmt.Errorf(common.DBError, err)
	}
	return nil
}
func (u userBasicRepo) FindByID(ctx context.Context, id uint) domain.UserBasic {
	user := domain.UserBasic{}
	db := u.db.WithContext(ctx)
	db.Where("id =?", id).Find(&user)
	return user
}
