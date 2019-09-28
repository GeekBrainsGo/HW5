package server

import (
	"encoding/json"
	"go_basics/packages/ormblog/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Server stands for server struct.
type Server struct {
	lg            *logrus.Logger
	db            *gorm.DB
	rootDir       string
	templatesDir  string
	indexTemplate string
	Page          models.Page
}

// New creates new server.
func New(lg *logrus.Logger, rootDir string, db *gorm.DB) *Server {
	return &Server{
		lg:            lg,
		db:            db,
		rootDir:       rootDir,
		templatesDir:  "/templates",
		indexTemplate: "index.html",
	}
}

// Start starts new server.
func (s *Server) Start(addr string) error {
	r := chi.NewRouter()
	s.bindRoutes(r)
	s.lg.Debug("server is started ...")
	return http.ListenAndServe(addr, r)
}

// SendErr sends and log error to user.
func (s *Server) SendErr(w http.ResponseWriter, err error, code int, obj ...interface{}) {
	s.lg.WithField("data", obj).WithError(err).Error("server error")
	w.WriteHeader(code)
	errModel := models.ErrorModel{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}
	data, _ := json.Marshal(errModel)
	w.Write(data)
}

// SendInternalErr sends 500 error.
func (s *Server) SendInternalErr(w http.ResponseWriter, err error, obj ...interface{}) {
	s.SendErr(w, err, http.StatusInternalServerError, obj)
}
