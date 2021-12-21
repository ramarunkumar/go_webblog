package main

import (
	_ "github.com/lib/pq"
)

func initializeRoutes() {

	r.Use(setUserStatus())

	r.GET("/", showIndexPage)

	u := r.Group("/u")
	{
		u.GET("/login", ensureNotLoggedIn(), showLoginPage)

		u.POST("/login", login)

		u.GET("/logout", ensureLoggedIn(), logout)

		u.GET("/signup", showRegistrationPage)

		u.POST("/signup", register)

		u.GET("/dashboard", dashboard)

	}

	articleRoutes := r.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)

		articleRoutes.GET("/create", showArticleCreationPage)

		articleRoutes.POST("/create", createArticle)

		articleRoutes.POST("/update", update)

	}

}
