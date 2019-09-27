package server

import (
	"encoding/json"
	"io/ioutil"
	"myblog/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// getTemplateHandler - возвращает шаблон
func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {

	blogs, err := models.GetAllBlogItems(serv.db)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Title = "Мой блог"
	serv.Page.Data = blogs
	serv.Page.Command = "index"

	if err := serv.dictionary["BLOGS"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

func getBlogIDFromRequest(r *http.Request) (uint, error) {

	bl := chi.URLParam(r, "id")
	// blogID, _ := strconv.ParseInt(bl, 10, 64)
	tempInt, err := strconv.Atoi(bl)
	if err != nil {
		return 0, err
	}
	blogID := uint(tempInt)

	return blogID, err
}

// editBlogHandler - редактирование блога
func (serv *Server) editBlogHandler(w http.ResponseWriter, r *http.Request) {

	blogID, err := getBlogIDFromRequest(r)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	blog, err := models.GetBlogItem(serv.db, blogID)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Title = "Редактирование"
	serv.Page.Data = blog
	serv.Page.Command = "edit"

	if err := serv.dictionary["BLOG"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// putBlogHandler - обновляет блог
func (serv *Server) putBlogHandler(w http.ResponseWriter, r *http.Request) {

	blogID, err := getBlogIDFromRequest(r)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)

	blog := models.BlogItem{}
	_ = json.Unmarshal(data, &blog)

	blog.ID = blogID

	if err := blog.UpdateBlog(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}

// addGetBlogHandler - добавление блога
func (serv *Server) addGetBlogHandler(w http.ResponseWriter, r *http.Request) {

	serv.Page.Title = "Добавление блога"
	serv.Page.Data = models.BlogItem{}
	serv.Page.Command = "new"

	if err := serv.dictionary["BLOG"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// addBlogHandler - добавляет блог
func (serv *Server) addBlogHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	blog := models.BlogItem{}
	_ = json.Unmarshal(data, &blog)

	if err := blog.AddBlog(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}

// deleteBlogHandler - удаляет блог
func (serv *Server) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {

	blogID, err := getBlogIDFromRequest(r)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	blog := models.BlogItem{}
	blog.ID = blogID

	if err := blog.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}
