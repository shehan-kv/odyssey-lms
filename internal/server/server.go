package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
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

	http.Handle("GET /api/course", middleware.Authed(http.HandlerFunc(handler.GetCourses)))
	http.Handle("GET /api/course/enroll", middleware.Authed(http.HandlerFunc(handler.GetEnrolledCourses)))
	http.Handle("GET /api/course/{id}/enroll", middleware.Authed(http.HandlerFunc(handler.GetEnrolledCourse)))
	http.Handle("GET /api/course/{id}/enroll/section", middleware.Authed(http.HandlerFunc(handler.GetEnrolledSections)))
	http.Handle("GET /api/course/{courseId}/enroll/section/{sectionId}", middleware.Authed(http.HandlerFunc(handler.GetEnrolledSection)))
	http.Handle("GET /api/course/{id}", middleware.Authed(http.HandlerFunc(handler.GetCourseById)))
	http.Handle("GET /api/course/category", middleware.Authed(http.HandlerFunc(handler.GetCategories)))
	http.Handle("POST /api/course", middleware.Authed(http.HandlerFunc(handler.CreateCourse)))
	http.Handle("POST /api/course/{id}/enroll", middleware.Authed(http.HandlerFunc(handler.EnrollInCourse)))
	http.Handle("POST /api/course/category", middleware.Authed(http.HandlerFunc(handler.CreateCategory)))
	http.Handle("POST /api/course/{courseId}/enroll/section/{sectionId}/complete", middleware.Authed(http.HandlerFunc(handler.CompleteSection)))

	http.HandleFunc("GET /api/system", handler.GetSystemInfo)

	staticUiFs, err := fs.Sub(web.WebUiFS, "ui/build")
	if err != nil {
		log.Fatalln("[ ERROR ] Error with static file system")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("[ ERROR ] Could not get current working directory")
	}
	uploadsFs := os.DirFS(path.Join(cwd, "uploads"))

	http.Handle("GET /_app/", http.FileServerFS(staticUiFs))
	http.Handle("GET /favicon.png", http.FileServerFS(staticUiFs))
	http.Handle("GET /uploads/", http.StripPrefix("/uploads", http.FileServerFS(uploadsFs)))

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
	err = http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Failed to listen on " + listenOn + colors.Reset)
	}
}
