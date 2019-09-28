package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"serv/models"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
)

// HandleGetIndexHtml - возвращает главную страницу - index.html
func (serv *Server) HandleGetIndexHtml(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(serv.rootDir + "/static/index.html")
	data, _ := ioutil.ReadAll(file)

	posts := models.GetAllPosts(serv.db)

	serv.Posts = posts

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", serv)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

// HandleGetPostHtml - возвращает страницу конкретного поста - post.html
func (serv *Server) HandleGetPostHtml(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")

	postID, _ := strconv.Atoi(postIDStr)

	post := models.PostItem{}
	serv.db.Find(&post, "id = ?", postID)

	file, _ := os.Open(serv.rootDir + "/static/post.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

// HandleGetEditHtml - возвращает страницу редактирования поста - edit.html
func (serv *Server) HandleGetEditHtml(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")

	postID, _ := strconv.Atoi(postIDStr)

	post := models.PostItem{}
	serv.db.Find(&post, "id = ?", postID)

	file, _ := os.Open(serv.rootDir + "/static/edit.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

// HandleGetNewHtml - возвращает страницу создания нового поста - new.html
func (serv *Server) HandleGetNewHtml(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(serv.rootDir + "/static/new.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", serv)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

// postNewHandler - добавление нового поста
func (serv *Server) postNewHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	post := models.PostItem{}
	_ = json.Unmarshal(data, &post)

	post.Insert(serv.db)

	data, _ = json.Marshal(post)
	w.Write(data)
}

// deletePostHandler - удаляем пост
func (serv *Server) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	postID_, _ := strconv.Atoi(postID)

	post := models.PostItem{ID: uint(postID_)}

	post.Delete(serv.db)

}

// putPostHandler - обновляем пост
func (serv *Server) putPostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	fmt.Println(postID)

	postID_, _ := strconv.Atoi(postID)

	data, _ := ioutil.ReadAll(r.Body)

	post := models.PostItem{}
	_ = json.Unmarshal(data, &post)
	post.ID = uint(postID_)

	data, _ = json.Marshal(post)
	w.Write(data)

	post.Update(serv.db)

}
