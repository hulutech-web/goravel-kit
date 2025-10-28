package seeders

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	users := []models.User{}
	err := facades.Orm().Factory().Count(10).Create(&users)
	if err != nil {
		return err
	}
	user1 := models.User{}
	facades.Orm().Query().Where("id", 1).First(&user1)
	user1.Username = "admin"
	if facades.Hash().NeedsRehash("admin888") {
		//Hash加密
		user1.Password, _ = facades.Hash().Make("admin888")
		_, err := facades.Orm().Query().Where("id=?", user1.ID).Update(&user1)
		if err != nil {
			return err
		}
	}
	user2 := models.User{}
	facades.Orm().Query().Where("id", 2).First(&user2)
	user2.Username = "test"
	if facades.Hash().NeedsRehash("admin888") {
		//Hash加密
		user2.Password, _ = facades.Hash().Make("admin888")
		_, err := facades.Orm().Query().Where("id=?", user2.ID).Update(&user2)
		if err != nil {
			return err
		}
	}
	return nil
}
