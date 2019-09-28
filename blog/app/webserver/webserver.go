package webserver

import (
	"blog/models"
	"context"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

// WebServer ...
type WebServer struct {
	router   *chi.Mux
	logger   *logrus.Logger
	database *sql.DB
}

func newServer(db *sql.DB) *WebServer {
	serv := &WebServer{
		router:   chi.NewRouter(),
		logger:   logrus.New(),
		database: db,
	}

	serv.configureRouter()

	return serv
}

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseConnectionString)
	if err != nil {
		return err
	}

	defer db.Close()
	serv := newServer(db)
	return http.ListenAndServe(config.BindAddr, serv)
}

func newDB(dsnURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsnURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (serv *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv *WebServer) configureRouter() {
	//routes
	serv.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	serv.router.HandleFunc("/list", serv.postListHandle())

	serv.router.HandleFunc("/view/{postID}", serv.postViewHandle())

	serv.router.HandleFunc("/delete/{postID}", serv.postDeleteHandle())

	serv.router.HandleFunc("/create", serv.postCreateHandle())

}

func (serv *WebServer) postListHandle() http.HandlerFunc {

	type PageModel struct {
		Title string
		Data  interface{}
	}

	return func(w http.ResponseWriter, r *http.Request) {

		boil.SetDB(serv.database)

		ctx := context.Background()

		posts, err := models.Posts().All(ctx, serv.database)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		pageData := PageModel{
			Title: "Posts List",
			Data:  posts,
		}

		templ := template.Must(template.New("page").ParseFiles("./templates/blog/List.tpl", "./templates/common.tpl"))
		err = templ.ExecuteTemplate(w, "page", pageData)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

	}
}

func (serv *WebServer) postViewHandle() http.HandlerFunc {

	type PageModel struct {
		Title string
		Data  interface{}
	}

	return func(w http.ResponseWriter, r *http.Request) {

		postIDStr := chi.URLParam(r, "postID")
		postID, _ := strconv.ParseUint(postIDStr, 10, 64)

		boil.SetDB(serv.database)

		ctx := context.Background()

		post, err := models.FindPost(ctx, serv.database, uint(postID))
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		pageData := PageModel{
			Title: "View Post",
			Data:  post,
		}

		templ := template.Must(template.New("page").ParseFiles("./templates/blog/View.tpl", "./templates/common.tpl"))
		err = templ.ExecuteTemplate(w, "page", pageData)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

	}
}

func (serv *WebServer) postDeleteHandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		postIDStr := chi.URLParam(r, "postID")
		postID, _ := strconv.ParseUint(postIDStr, 10, 64)

		boil.SetDB(serv.database)

		ctx := context.Background()

		_, err := models.Posts(models.PostWhere.ID.EQ(uint(postID))).DeleteAll(ctx, serv.database)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		w.Header().Add("Location", "/list")
		w.WriteHeader(302)
		return

	}
}

func (serv *WebServer) postCreateHandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		boil.SetDB(serv.database)

		ctx := context.Background()

		var newPost models.Post

		newPost.Title.Scan("New Post Title")
		newPost.Short.Scan("Short body")
		newPost.Body.Scan("Content body")

		err := newPost.Insert(ctx, serv.database, boil.Infer())
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		w.Header().Add("Location", "/list")
		w.WriteHeader(302)
		return

	}
}

func (serv *WebServer) errorAPI(w http.ResponseWriter, r *http.Request, code int, err error) {
	serv.respondJSON(w, r, code, map[string]string{"error": err.Error()})
}

func (serv *WebServer) respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (serv *WebServer) respondWhithTemplate(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
