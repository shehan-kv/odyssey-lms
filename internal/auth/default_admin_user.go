package auth

import (
	"context"
	"database/sql"
	"log"
	"math/rand"

	"odyssey.lms/internal/colors"
	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
)

func CreateDefaultAdminUser() {

	ctx := context.Background()
	adminCount, _ := db.QUERY.CountUsersByRole(ctx, "administrator")

	if adminCount != 0 {
		return
	}

	tempPassword := getRandomString(10)
	hashedPassword, err := HashPassword(tempPassword)
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Failed to create admin account" + colors.Reset)
	}

	adminUser, err := db.QUERY.CreateUser(ctx, params.CreateUser{
		FirstName: "Default",
		LastName:  "Administrator",
		Email:     "admin@lms.local",
		Password:  hashedPassword,
		Bio: sql.NullString{
			String: "This is the default admin account of the system, created when other admin " +
				"accounts couldn't be found. Please delete this account once you create an admin account",
			Valid: true,
		},
	})
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Failed to create admin account" + colors.Reset)
	}

	_, err = db.QUERY.AssignUserRole(ctx, params.AssignUserRole{UserID: adminUser.ID, RoleName: "administrator"})
	if err != nil {
		_ = db.QUERY.DeleteUserById(ctx, adminUser.ID)
		log.Fatal(colors.RedBold + "[ ERROR ] Failed to create admin account" + colors.Reset)
	}

	log.Println(colors.Blue + "[ INFO ] Created " + colors.Reset + colors.Yellow + "admin@lms.local " + colors.Reset +
		colors.Blue + "account with password: " + colors.Reset + colors.Yellow + tempPassword + colors.Reset)

}

func getRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	randomLetters := make([]byte, length)
	for i := range randomLetters {
		randomLetters[i] = letters[rand.Intn(len(letters))]
	}

	return string(randomLetters)
}
