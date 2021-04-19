package server

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type TokenServer struct {
	token      string
	tokenMutex sync.Mutex
	Log        LogFunc `envconfig:"-"`

	*http.Server
}

type LogFunc func(pattern string, args ...interface{})

type Config struct {
	IsWithToken bool    `envconfig:"default=false"`
	Address     string  `envconfig:"default=127.0.0.1:5000"`
	Log         LogFunc `envconfig:"-"`
}

func New(c *Config) *TokenServer {
	ts := &TokenServer{
		tokenMutex: sync.Mutex{},
		Log:        c.Log,
		token:      "",
	}

	server := &http.Server{
		Addr: c.Address,
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				c.Log("Unexpected token server call with method %s", r.Method)
				rw.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			// ioutil.ReadAll(r.Body)
			var token string
			if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
				c.Log("Could not decode request body: %+v", err)
				rw.Write([]byte("Unexpected body format"))
				rw.WriteHeader(http.StatusBadRequest)
			}
			ts.tokenMutex.Lock()
			defer ts.tokenMutex.Lock()
			ts.token = token
			c.Log("Token successfully provided and set")
			rw.WriteHeader(http.StatusOK)
		}),
	}

	ts.Server = server
	return ts
}

func (s *TokenServer) IsTokenSet() (string, bool) {
	s.tokenMutex.Lock()
	defer s.tokenMutex.Unlock()
	return s.token, len(s.token) > 0
}

func WaitForToken(ts *TokenServer) string {
	go func() {
		if err := ts.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ts.Log("Unexpected server listen err: %+v", err)
			return
		}
		ts.Log("Token server closed")
	}()
	for {
		ts.Log("Waiting for token to be set")
		if token, isSet := ts.IsTokenSet(); isSet {
			ts.Log("Token was found")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			if err := ts.Server.Shutdown(ctx); err != nil {
				ts.Log("Unexpected err: %s", err.Error())
			}
			return token
		}
		time.Sleep(time.Second * 2)
	}
}
