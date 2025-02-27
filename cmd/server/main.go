package main

import (
	"example/user/hello/config"
	"example/user/hello/internal/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	//Connect to MongoDB
	/*client := utils.ConnectDB()
	defer client.Disconnect(nil)*/

	router := gin.Default()
	router.Use(cors.Default())

	cors := cors.DefaultConfig()
	cors.AllowAllOrigins = true
	cors.AllowHeaders = append(cors.AllowHeaders, "Authorization", "X-Api-Key")

	// Swagger
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/experience", controllers.GetExperienceHandler)

	port := os.Getenv("PORT")

	router.Run(":" + port)

	/*http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})

	http.ListenAndServe(":8080", nil)*/

}
