package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Project incoming"))
	// })

	basePath := "/"

	r.Route(basePath, func(root chi.Router) {
		root.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})

		workDir, _ := os.Getwd()
		println("workDir", workDir)
		filesDir := filepath.Join(workDir, "files")
		println("filesDir", filesDir)
		FileServer(root, basePath, "../portfolioClient/build", http.Dir(filesDir))
	})

	http.ListenAndServe(":7777", r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

}
