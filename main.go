package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token 	 string `json:"token"`
	Offset	 int64 	`json:"offset"`
}

func login(c *gin.Context) {
	var creds Credentials
	var err error

	if err = c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	// for the sake of time/complexity, use 200 status
	// then redirect on frontend
	if validate(creds) {
		c.JSON(http.StatusOK, gin.H{ })
		return
	}

	// attach a more specific message?
	c.JSON(http.StatusUnauthorized, gin.H{ "error": "Something went wrong" })
}

// obviously not realistic, done for time/convenience
func validate(creds Credentials) bool {
	var valid Credentials
	valid.Username = "c137@onecause.com"
	valid.Password = "#th@nH@rm#y#r!$100%D0p#"
	valid.Token = time.Now().UTC().Format("1504")

	var err error
	var credTime time.Time
	credTime, err = time.Parse("1504", creds.Token)

	if err != nil {
		return false;
	}

	// assuming token is local
	return (credTime.Add(time.Minute * time.Duration(creds.Offset)).Format("1504") == time.Now().UTC().Format("1504") &&
			valid.Username == creds.Username &&
			valid.Password == creds.Password)
}

// if we weren't just redirecting on login
// we'd want session storage of some sort
// one option being "gin-contrib/sessions"
// along with some kind of storage middleware
// also public/private routes, AuthRequired
func main() {
	var router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	router.POST("/login", login)

	router.Run("localhost:8080")
}
