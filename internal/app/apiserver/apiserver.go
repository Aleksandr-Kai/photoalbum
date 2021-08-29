package apiserver

import (
	"io"
	"net/http"
	"photoalbum/internal/app/store"

	"github.com/Aleksandr-Kai/logger"
	"github.com/gorilla/mux"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger logger.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logger.NewLogger(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	s.configureLogger()
	s.logger.LogToConsole(logger.Info, "Starting API Server")

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() {
	level := logger.ParseLevel(s.config.LogLevel)
	s.logger.GlobalLevel(level)
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello!")
	}
}
