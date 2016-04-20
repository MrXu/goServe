package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const SocialFacebook string = "socialFacebook"

const FacebookProfileUrl string = "https://graph.facebook.com/v2.2/me?fields=id,name,email,picture,first_name,last_name"

var passportOauth *oauth2.Config

type FacebookProfile struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Hd        string `json:"hd"`
	Locale    string `json:"locale"`
	Name      string `json:"name"`
	Picture   struct {
		Data struct {
			Url string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}

func ConfigFacebook(r *gin.RouterGroup, clientId string, secretKey string, redirectUrl string) {
	facebookConfig := &oauth2.Config{
		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

		ClientID:     clientId, // change this to yours
		ClientSecret: secretKey,
		RedirectURL:  redirectUrl, // change this to your webserver adddress
		Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}

	passportOauth = facebookConfig

	r.GET("/login", func(c *gin.Context) {
		Login(facebookConfig, c)
	})
	r.GET("/callback", MiddlewareFacebook(), func(c *gin.Context) {
		user, err := GetFacebookProfile(c)
		if user == nil || err != nil {
			c.AbortWithStatus(500)
			return
		}
		
		c.JSON(http.StatusOK, *user)
	})

}

func Routes(oauth *oauth2.Config, r *gin.RouterGroup) {
	passportOauth = oauth

	r.GET("/login", func(c *gin.Context) {
		Login(oauth, c)
	})
}

func Login(oauth *oauth2.Config, c *gin.Context) {
	url := oauth.AuthCodeURL("")
	c.Redirect(http.StatusFound, url)
}

func MiddlewareFacebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		getProfile(c)
	}
}

func GetFacebookProfile(c *gin.Context) (*FacebookProfile, error) {
	user, exists := c.Get(SocialFacebook)
	if !exists {
		return nil, errors.New("GinPassportFacebook namespace key doesn't exist")
	}

	return user.(*FacebookProfile), nil
}

func getProfile(c *gin.Context) {
	c.Request.ParseForm()

	config := passportOauth
	code := c.Request.Form.Get("code")

	t, err := config.Exchange(oauth2.NoContext, code)

	if t == nil {
		c.Redirect(301, "/")
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	client := config.Client(oauth2.NoContext, t)

	resp, err := client.Get(FacebookProfileUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userInformation FacebookProfile
	err = json.Unmarshal(contents, &userInformation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(SocialFacebook, &userInformation)
	c.Next()
}
