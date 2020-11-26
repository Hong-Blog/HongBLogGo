package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"loginsrv/oauth"
	"net/http"
)

func main() {
	r := gin.Default()

	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	oauth.InitServer(manager)
	oauth.SetAllowGetAccessRequest(true)
	oauth.SetAllowedGrantType(oauth2.PasswordCredentials)
	oauth.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		userID = "testok"
		return
	})
	oauth.SetClientInfoHandler(server.ClientFormHandler)

	auth := r.Group("/oauth2")
	{
		auth.GET("/token", oauth.HandleTokenRequest)
	}

	api := r.Group("/api")
	{
		api.Use(oauth.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get(oauth.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8090")
}
