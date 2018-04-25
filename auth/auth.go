package auth

import (
	"time"
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"encoding/json"
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"fmt"
)

const (
	codeRedirect = 302
	keyToken ="oauth2_token"
	keyNextPage = "next"
)

var (
	// PathLogin is the path to handle OAuth 2.0 login.
	PathLogin = "/login"
	// PathLogout is the path to handle OAuth 2.0 logout.
	PathLogout = "/logout"
	// PathCallback is the path to handle callback from OAuth 2.0 backend
	// to exchange credentials
	PathCallback = "/auth/google/callback"
	// PathError os the path to handle error case.
	PathError = "/unauthorized"
)

type userInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}
// Tokens represent a contains that contains user's OAuth 2.0 access and refresh tokens.
type Tokens interface {
	Access() string
	Refresh() string
	Expired() bool
	ExpiryTime() time.Time
}

type token struct {
	oauth2.Token
}

// Access returns the access token
func (t *token) Access() string {
	return t.AccessToken
}
// Refresh returns the refresh token
func (t *token) Refresh() string {
	return t.RefreshToken
}
// Expired returns whether the access token is expired or not
func (t *token) Expired() bool {
	if t == nil {
		return true
	}
	return !t.Token.Valid()
}
// ExpiryTime returns the expiry time of the user's access token
func (t *token) ExpiryTime() time.Time {
	return t.Expiry
}
// String return  the string representation of the token
func (t *token) String () string  {
	return fmt.Sprintf("tokens: %s expire at : %s", t.Access(), t.ExpiryTime())
}
// Google returns a new Google OAuth 2.0 backend endpoint
func Google(conf *oauth2.Config) gin.HandlerFunc  {
	return NewOAuth2Provider(conf)
}

// NewOAuth2provider returns a generic OAuth 2.0 backend endpoint
func NewOAuth2Provider(conf *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			switch c.Request.URL.Path {
			case PathLogin:
				login(conf, c)
			case PathLogout:
				logout(c)
			case PathCallback:
				handleOAuth2Callback(conf, c)
			}
		}
		s := sessions.Default(c)
		tk := unmarshalToken(s)
		if tk != nil {
			if tk.Expired() && tk.Refresh() == "" {
				s.Delete(keyToken)
				s.Save()
				tk = nil
			}
		}
	}
}
// Handler that redirects user to the login page, if user is not logged in
// simple usage:
// m.Get("/login_required",oauth2.LoginRequired, func () ... {})
var LoginRequired = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		token := unmarshalToken(s)
		if token == nil || token.Expired() {
			next := url.QueryEscape(c.Request.URL.RequestURI())
			http.Redirect(c.Writer, c.Request, PathLogin+ "?next="+next,codeRedirect)
		}
	}
}()

func login(f *oauth2.Config, c *gin.Context)  {
	s := sessions.Default(c)
	next := extractPath(c.Request.URL.Query().Get(keyNextPage))
	if s.Get(keyToken) == nil {
		// User is not logged in
		if next == "" {
			next = "/"
		}
		http.Redirect(c.Writer, c.Request, f.AuthCodeURL(next), codeRedirect)
		return
	}
	// No need to login , redirect to the next page.
	http.Redirect(c.Writer, c.Request, next, codeRedirect)
}



func logout(c *gin.Context)  {
	s := sessions.Default(c)
	next := extractPath(c.Request.URL.Query().Get(keyNextPage))
	s.Delete(keyToken)
	s.Save()
	http.Redirect(c.Writer, c.Request, next, codeRedirect)
}

func handleOAuth2Callback(f *oauth2.Config, c *gin.Context)  {
	s := sessions.Default(c)
	next := extractPath(c.Request.URL.Query().Get("state"))
	code := c.Request.URL.Query().Get("code")
	t, err := f.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("exchange oauth token failed:", err)
		http.Redirect(c.Writer, c.Request, PathError, codeRedirect)
		return
	}
	val, _ := json.Marshal(t)
	s.Set(keyToken, val)
	s.Save()
	http.Redirect(c.Writer, c.Request, next, codeRedirect)
}

func unmarshalToken(s sessions.Session) *token {
	if s.Get(keyToken) == nil {
		return nil
	}
	data := s.Get(keyToken).([]byte)
	var tk oauth2.Token
	json.Unmarshal(data, &tk)
	return &token{tk}
}

func extractPath(next string) string  {
	n, err := url.Parse(next)
	if err !=nil {
		return "/"
	}
	return n.Path
}

func GoogleAuthConfig(keyPath string, debug bool) *oauth2.Config  {
	jsonKey, err := ioutil.ReadFile(keyPath)
	if err !=nil {
		log.Fatal(err)
	}
	conf, _ := google.ConfigFromJSON(jsonKey, "profile")
	if debug {
		conf.RedirectURL = "http://localhost:8888/auth/google/callback"
	}
	return conf
}

func GoogleAuthFromConfig(keyPath string, debug bool) gin.HandlerFunc  {
	return Google(GoogleAuthConfig(keyPath, debug))
}