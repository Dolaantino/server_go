package apiserver

import (
	"io"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil{
		return err
	}
	s.configureRouter()
	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func(s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
