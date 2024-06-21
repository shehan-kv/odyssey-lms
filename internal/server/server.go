package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

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
	http.HandleFunc("GET /api/auth/is-signed-in", handler.IsSignedIn)

	http.HandleFunc("GET /api/user", handler.GetUsers)

	http.HandleFunc("GET /api/system", handler.GetSystemInfo)

	staticUiFs, _ := fs.Sub(web.WebUiFS, "ui/build")

	http.Handle("GET /_app/", http.FileServerFS(staticUiFs))
	http.Handle("GET /favicon.png", http.FileServerFS(staticUiFs))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
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
