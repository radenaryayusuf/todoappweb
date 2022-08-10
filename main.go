package main

import (
	// "database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/radenaryayusuf/todoappweb/routes"
	// "github.com/radenaryayusuf/todoappweb/internal/models"
)

// type application struct {
// 	errorLog      *log.Logger
// 	infoLog       *log.Logger
// 	tasks         *models.TaskModel
// 	templateCache map[string]*template.Template
// }

// var tpl = template.Must(template.ParseFiles("./../web/index.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	tpl.Execute(w, nil)
// 	// w.Write([]byte("<h1>Hello World!</h1>"))
// }

// func searchHandler(w http.ResponseWriter, r *http.Request) {
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	params := u.Query()
// 	searchQuery := params.Get("q")
// 	page := params.Get("page")
// 	if page == "" {
// 		page = "1"
// 	}

// 	fmt.Println("Search query is: ", searchQuery)
// 	fmt.Println("Page is: ", page)

// }

func main() {
	logger, _ := thoth.Init("log")

	err := godotenv.Load()
	if err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}
	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	// dsn := flag.String("dsn", "user=postgres dbname=todoapp password=P@ssw0rd sslmode=disable", "Psql data source name")

	// flag.Parse()

	// db, err := openDB(*dsn)
	// if err != nil {
	// 	log.Fatal("Error loading DB connection")
	// }
	// defer db.Close()

	if port == "" {
		port = "3000"
	}

	// fs := http.FileServer(http.Dir("./../web/static/css"))
	// fjs := http.FileServer(http.Dir("./../web/static/js"))

	// mux := http.NewServeMux()

	// mux.Handle("/web/static/css/", http.StripPrefix("/web/static/css/", fs))
	// mux.Handle("/web/static/js/", http.StripPrefix("/web/static/js/", fjs))
	// mux.HandleFunc("/search", searchHandler)
	// mux.HandleFunc("/", indexHandler)
	server := http.ListenAndServe(":"+port, routes.Init())

	if server != nil {
		logger.Log(server)
		log.Fatal(server)
	}
	// http.ListenAndServe(":"+port, mux)
}

// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
