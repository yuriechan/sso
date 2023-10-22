package main

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Initialize Viper across the application
	InitializeViper()

	// Initialize Logger across the application
	InitializeZapCustomLogger()

	// Initialize Oauth2 Services
	InitializeOAuthGoogle()

	// Routes for the application
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/login-gl", HandleGoogleLogin)
	http.HandleFunc("/callback-gl", CallBackFromGoogle)

	Log.Info("Started running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil), nil)
}

/*
HandleMain Function renders the index page when the application index route is called
*/
func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(IndexPage))
}

/*
HandleLogin Function
*/
func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		Log.Error("Parse: " + err.Error())
	}
	Log.Info(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	Log.Info(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
