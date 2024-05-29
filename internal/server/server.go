package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/colors"
	"odyssey.lms/internal/db"
	"odyssey.lms/internal/handler"
	"odyssey.lms/web"
)

func init() {
	db.CheckDBSettings()
	db.RunMigrations()
	auth.CreateDefaultAdminUser()
}

func RunApplication() {

	http.HandleFunc("POST /api/auth/sign-in", handler.SignIn)

	staticUiFs, _ := fs.Sub(web.WebUiFS, "ui/build")

	http.Handle("GET /_app/", http.FileServerFS(staticUiFs))
	http.Handle("GET /favicon.png", http.FileServerFS(staticUiFs))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticUiFs, "index.html")
	})

	listenOn := os.Getenv("LISTEN_ON")
	if listenOn == "" {
		listenOn = ":8080"
	}

	log.Println(colors.GreenBold + "[ INFO ] Server listening on " + listenOn + colors.Reset)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Failed to listen on " + listenOn + colors.Reset)
	}
}
