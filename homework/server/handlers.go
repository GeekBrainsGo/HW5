package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"serv/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/volatiletech/sqlboiler/boil"
)

func (serv *Server) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(filepath.Join(serv.staticDir + "/index.html"))
	data, _ := ioutil.ReadAll(file)
	if blogItems, err := models.Blogs().All(nil, serv.db); err != nil {
		serv.lg.Error("Error getting all posts", err)
	} else {
		indexTemplate := template.Must(template.New("index").Parse(string(data)))
		err := indexTemplate.ExecuteTemplate(w, "index", blogItems)
		if err != nil {
			serv.lg.WithError(err).Error("template")
		}
	}

}

func (serv *Server) handleGetPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(filepath.Join(serv.staticDir + "/post.html"))
	data, _ := ioutil.ReadAll(file)
	postNumberStr := chi.URLParam(r, "id")
	indexTemplate := template.Must(template.New("index").Parse(string(data)))
	postNumber, err := strconv.Atoi(postNumberStr)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	searchedPost, err := models.FindBlog(nil, serv.db, postNumber)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	err = indexTemplate.ExecuteTemplate(w, "index", searchedPost)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) handleGetEditPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(filepath.Join(serv.staticDir + "/edit.html"))
	data, _ := ioutil.ReadAll(file)
	postNumberStr := chi.URLParam(r, "id")
	indexTemplate := template.Must(template.New("index").Parse(string(data)))
	postNumber, err := strconv.Atoi(postNumberStr)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	searchedPost, err := models.FindBlog(nil, serv.db, postNumber)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	err = indexTemplate.ExecuteTemplate(w, "index", searchedPost)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) handlePostEditPost(w http.ResponseWriter, r *http.Request) {
	var post models.Blog
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Update(nil, serv.db, boil.Infer())
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}

}

func (serv *Server) handlePostCreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Blog
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Insert(nil, serv.db, boil.Infer())
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}
}

func (serv *Server) handlePostDeletePost(w http.ResponseWriter, r *http.Request) {
	var post models.Blog
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Delete(nil, serv.db)
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}

}
