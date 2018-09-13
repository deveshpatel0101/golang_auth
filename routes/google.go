package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"

	"github.com/tkanos/gonfig"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"golang.org/x/oauth2"
)

var googleOauthConfig *oauth2.Config
var oauthStateString string
var configuration struct {
	ClientID     string
	ClientSecret string
}

func init() {
	err := gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		fmt.Println("Error while reading configuration file.")
		configuration.ClientID = os.Getenv("CLIENT_ID")
		configuration.ClientSecret = os.Getenv("CLIENT_SECRET")
	}
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/google/callback",
		ClientID:     configuration.ClientID,
		ClientSecret: configuration.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
}

// Google will allow users to login using google
func Google(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	uuid, err := uuid.NewV4()
	if err != nil {
		oauthStateString = "something"
	} else {
		oauthStateString = uuid.String()
	}
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

// Callback route after google redirect
func Callback(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	content, err := getGoogleInfo(req.FormValue("state"), req.FormValue("code"))
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusInternalServerError)
		return
	}

	// Convert json to go struct
	var uj models.GoogleUser
	err = json.Unmarshal(content, &uj)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusInternalServerError)
		return
	}

	// Convert google user to user db
	ui := convertGoogleUser(uj)
	err = controllers.CreateUser(ui)
	if !(err == nil || err.Error() == "user already exists") {
		http.Redirect(w, req, "/user/login", http.StatusInternalServerError)
		return
	}

	// Get user info
	fu, err := controllers.GetUserByEmail(ui.Email)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusInternalServerError)
		return
	}

	// Create session
	us, err := controllers.CreateSession(fu)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sid",
		Value:    flash.Encode([]byte(us.UUID)),
		HttpOnly: true,
		Path:     "/",
	})
	http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
}

func getGoogleInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func convertGoogleUser(u models.GoogleUser) models.UserDB {
	newUser := models.UserDB{
		Fname:    u.Fname,
		Lname:    u.Lname,
		Email:    u.Email,
		Picture:  u.Picture,
		GoogleID: u.GoogleID,
		UserType: "google",
	}
	return newUser
}
