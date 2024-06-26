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
	"odyssey.lms/internal/middleware"
	"odyssey.lms/web"
)

func init() {
	db.CheckDBSettings()
	db.RunMigrations()
	auth.CreateDefaultAdminUser()
}

func RunApplication() {

	http.HandleFunc("POST /api/auth/sign-in", handler.SignIn)
	http.Handle("GET /api/auth/is-signed-in", middleware.Authed(http.HandlerFunc(handler.IsSignedIn)))

	http.HandleFunc("GET /api/user", handler.GetUsers)
	http.HandleFunc("POST /api/user", handler.CreateUser)
	http.HandleFunc("DELETE /api/user/{id}", handler.DeleteUser)
	http.HandleFunc("POST /api/user/activate/{id}", handler.ActivateUser)
	http.HandleFunc("POST /api/user/deactivate/{id}", handler.DeactivateUser)

	http.HandleFunc("GET /api/event", handler.GetEvents)

	http.Handle("GET /api/support-ticket", middleware.Authed(http.HandlerFunc(handler.GetSupportTickets)))
	http.Handle("GET /api/support-ticket/self", middleware.Authed(http.HandlerFunc(handler.GetSupportTicketsSelf)))
	http.Handle("GET /api/support-ticket/{id}", middleware.Authed(http.HandlerFunc(handler.GetSupportTicketById)))
	http.Handle("POST /api/support-ticket", middleware.Authed(http.HandlerFunc(handler.CreateSupportTicket)))
	http.Handle("POST /api/support-ticket/{id}/message", middleware.Authed(http.HandlerFunc(handler.CreateSupportTicketMessage)))
	http.Handle("POST /api/support-ticket/{id}/resolve", middleware.Authed(http.HandlerFunc(handler.ResolveTicket)))

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
