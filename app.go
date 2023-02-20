package main

//const fetchCableDiameter = () => {
//	return fetch('https://takehome-frontend.oden-qa.network/?metric=cable-diameter')
//}
//
//
import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"encoding/json"

	// "fmt"
	//"io"
	// "io/fs"
	// "log"
	// "mime"
	// "os"
	//"path/filepath"
	//"strings"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))
var UI embed.FS

// func init() {
// 	var err error
// 	uiFS, err = fs.Sub(app.UI, "_ui/build")
// 	if err != nil {
// 		log.Fatal("failed to get ui fs", err)
// 	}
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/health", handleHealth)
// 	mux.HandleFunc("/api", handleApi)
// 	// ...
// 	mux.HandleFunc("/", handleStatic)
// 	// ...
// 	log.Println("starting server...")
// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		log.Println("server failed:", err)
// 	}
// }


// func handleHealth(w http.ResponseWriter, r *http.Request) {
// 	// TODO: implement
// }

// func handleApi(w http.ResponseWriter, r *http.Request) {
// 	// TODO: implement
// }

// func handleStatic(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	path := filepath.Clean(r.URL.Path)
// 	if path == "/" { // Add other paths that you route on the UI-side here
// 		path = "index.html"
// 	}
// 	path = strings.TrimPrefix(path, "/")

// 	file, err := uiFS.Open(path)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			log.Println("file", path, "not found:", err)
// 			http.NotFound(w, r)
// 			return
// 		}
// 		log.Println("file", path, "cannot be read:", err)
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	contentType := mime.TypeByExtension(filepath.Ext(path))
// 	w.Header().Set("Content-Type", contentType)
// 	if strings.HasPrefix(path, "static/") {
// 		w.Header().Set("Cache-Control", "public, max-age=31536000")
// 	}
// 	stat, err := file.Stat()
// 	if err == nil && stat.Size() > 0 {
// 		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
// 	}

// 	n, _ := io.Copy(w, file)
// 	log.Println("file", path, "copied", n, "bytes")
// }


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	// http.HandleFunc("/react", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != "GET" {
	// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 		return
	// 	}
	
	// 	path := filepath.Clean(r.URL.Path)
	// 	if path == "/" { // Add other paths that you route on the UI side here
	// 		path = "index.html"
	// 	}
	// 	path = strings.TrimPrefix(path, "/")
	
	// 	file, err := uiFS.Open(path)
	// 	if err != nil {
	// 		if os.IsNotExist(err) {
	// 			log.Println("file", path, "not found:", err)
	// 			http.NotFound(w, r)
	// 			return
	// 		}
	// 		log.Println("file", path, "cannot be read:", err)
	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 		return
	// 	}
	
	// 	contentType := mime.TypeByExtension(filepath.Ext(path))
	// 	w.Header().Set("Content-Type", contentType)
	// 	if strings.HasPrefix(path, "static/") {
	// 		w.Header().Set("Cache-Control", "public, max-age=31536000")
	// 	}
	// 	stat, err := file.Stat()
	// 	if err == nil && stat.Size() > 0 {
	// 		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	// 	}
	
	// 	n, _ := io.Copy(w, file)
	// 	log.Println("file", path, "copied", n, "bytes")
	// })

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
