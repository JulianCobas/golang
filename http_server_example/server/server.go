package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Start is mine
func (s *Server) Start() (chan bool, error) {
	if s.initialized != true {
		if err := s.Init(); err != nil {
			return nil, err
		}
	}

	s.router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	s.router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		counterChan := s.worker.IncrementCounter()

		s.logger.Infof("Channel waiting...")
		result := <-counterChan
		s.logger.Infof("Channel received...")

		if result {
			s.logger.Infof("Result counter increment: %v", result)
			resp := []byte(strconv.Itoa(s.worker.counter))
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		} else {
			s.logger.Error("Increment counter error")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("409 - Something bad happened!"))
		}

	})

	addr := fmt.Sprintf("%s:%s", s.Viper.GetString("ip"), s.Viper.GetString("port"))
	http.ListenAndServe(addr, s.router)

	return s.exitChan, nil
}

//Init is mine
func (s *Server) Init() error {
	s.logger.Infof("Initializing the server")
	s.exitChan = make(chan bool)
	s.initialized = true

	r := mux.NewRouter()

	s.worker = &Worker{
		counterChan: make(chan bool),
	}

	s.worker.SetLogger(s.logger)
	return nil
}

//IncrementCounter is mine
func (s *Server) incrementCounter() {
	s.worker.IncrementCounter()
	s.worker.counterChan <- true
}
