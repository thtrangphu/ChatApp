package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	"github.com/mekanican/chat-backend/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

var googleOAuthConfig *oauth2.Config = nil

func setConfig() {
	if googleOAuthConfig == nil {
		googleOAuthConfig = &oauth2.Config{
			ClientID:     config.GetString("GCLOUD_ID"),
			ClientSecret: config.GetString("GCLOUD_SECRET"),
			RedirectURL:  "http://localhost:8000/auth/google/callback",
			Scopes:       []string{"email"},
			Endpoint:     google.Endpoint,
		}
	}
}

func CreateState() (string, error) {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes) //  Read 16 random bytes
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randBytes), nil
}

func GetAuthCodeURL(state string) string {
	setConfig()
	return googleOAuthConfig.AuthCodeURL(state)
}

// Return email of user
func HandleCallback(code string) (*GoogleResponse, error) {
	setConfig()
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	oauthGoogleUrlAPI := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	client := googleOAuthConfig.Client(context.Background(), token)
	response, err := client.Get(oauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	contents := new(GoogleResponse)
	err = json.NewDecoder(response.Body).Decode(&contents)

	if err != nil {
		return nil, err
	}

	return contents, nil
}
