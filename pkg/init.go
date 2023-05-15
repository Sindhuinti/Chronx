package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const credentialPath = "cred.json"

var tokenStash string

func GetClient() (*http.Client, error) {
	getOS()
	
	credentialsB, err := os.ReadFile(credentialPath)
	
	
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(credentialsB, calendar.CalendarReadonlyScope)

	config.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/calendar",
		"openid",
	}

	if err != nil {
		return nil, err
	}

	token, err := getToken(config)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)
	fmt.Println("Oauth Success")
	return client, nil

}

func getToken(config *oauth2.Config) (*oauth2.Token, error) {
	

	stashedToken, err := getStashedToken()
	if err != nil {
		token, err := getNewToken(config)
		if err != nil {
			return token, err
		}
		stashToken(token)
		return token, nil
	}
	return stashedToken, nil
}

func getStashedToken() (*oauth2.Token, error) {
	tokenB, err := os.ReadFile(tokenStash)
	if err != nil {
		return &oauth2.Token{}, err
	}

	if len(tokenB) == 0 {
		return &oauth2.Token{}, fmt.Errorf("unable to access token")
	}
	var token oauth2.Token
	_ = json.Unmarshal(tokenB, &token)
	return &token, nil
}

func stashToken(token *oauth2.Token) error {
	tokenB, err := json.Marshal(token)
	if err != nil {
		return err
	}

	return os.WriteFile(tokenStash, tokenB, 0644)
}

func getNewToken(config *oauth2.Config) (*oauth2.Token, error) {

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		return &oauth2.Token{}, fmt.Errorf("unable to read auth code: %s", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		return &oauth2.Token{}, fmt.Errorf("unable to exchange authcode for token: %s", err)
	}
	return token, nil

}

func getOS() {
	OS := runtime.GOOS
	if OS == "windows" {
		tokenStash = os.Getenv("LOCALAPPDATA") + "\\chronxToken.json"
	} else {
		tokenStash = "/tmp/chronxToken.json"
	}

}
