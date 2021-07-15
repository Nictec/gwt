package gwt

import (
	"github.com/Nictec/gwt/handler"
	"github.com/Nictec/gwt/logger"
	"github.com/Nictec/gwt/urls"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
	"net/http"
)

type Server struct{
	Name string
	Host string
	Urls []urls.Path
	Middleware []negroni.Handler
	SecretKey string
	Session *sessions.CookieStore
	Debug bool
}

func NewServer(name string, host string, urls []urls.Path, secretKey string, debug bool) Server {
	return Server{Name: name,Host: host, Urls: urls, SecretKey: secretKey, Debug: debug}
}

func (s *Server) contextHandler(next handler.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request){
		ctx := handler.Context{
			W:       w,
			Request: r,
			Session: s.Session,
			Debug:   s.Debug,
		}
		next(ctx)
	}
	return fn
}

func (s *Server) Run() error{
	// set up the router
	router := mux.NewRouter()
	for _, route := range s.Urls{
		router.HandleFunc(route.Path, s.contextHandler(route.Handler))
	}

	// csrf protection
	CSRF := csrf.Protect([]byte(s.SecretKey))

	// set up session
	s.Session = sessions.NewCookieStore([]byte(s.SecretKey))

	//set up logging
	logger.Init()

	//set up middleware
	n := negroni.Classic()
	for _, mw := range s.Middleware{
		n.Use(mw)
	}
	n.UseHandler(CSRF(router))

	//start the server
	logger.Info(s.Name + " is running on " + s.Host)
	err := http.ListenAndServe(s.Host, n)

	return err
}