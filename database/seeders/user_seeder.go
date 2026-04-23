package seeders

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/pkg/helpers"
	"gorm.io/gorm"
)
func RunUserSeeder(db *gorm.DB){
	jsonData, err := os.ReadFile("database/seeders/json/users.json")
	if err != nil{
		log.Println("Gagal membaca file JSON:", err)
		return
	}
	var users []entities.User
	if err := json.Unmarshal(jsonData, &users); err != nil{
		log.Println("Gagal parsing JSON:", err)
	}
	for _, user := range users {
		var existingUser entities.User
		err := db.Where("email = ?", user.Email).First(&existingUser).Error
		
		if err == gorm.ErrRecordNotFound {
			hashedPassword, _ := helpers.HashPassword(user.Password)
			user.Password = hashedPassword

			if err := db.Create(&user).Error; err != nil {
				log.Println("Gagal seed user:", user.Email, err)
			} else {
				log.Println("Berhasil seed user:", user.Email)
			}
		}
	}
}