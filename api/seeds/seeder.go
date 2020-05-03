package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mertatayurt/go_api/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Merti Atayurti",
		Email:    "atayurtmert@gmail.com",
		Password: "SecretPassword",
	},
	models.User{
		Nickname: "Elon Musk",
		Email:    "elon_musk_34@hotmail.com",
		Password: "AmGonnaGoToMarsOneDay",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Is this real ?",
		Content: "Maybe",
	},
	models.Post{
		Title:   "Who are the celebrities living in the moon ?",
		Content: "Moby, Lady Gaga",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatal("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatal("cannot migrate: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal("cannot add foreign key: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal("cannot seed user to table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal("cannot seed post to table: %v", err)
		}
	}
}
