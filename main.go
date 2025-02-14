package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	"github.com/mahdidl/golang_boilerplate/Router"
)

// @title                       Golang Monolithic Boilerplate
// @version                     1.0
// @description                 Golang Monolithic Boilerplate
// @contact.name                Supnex
// @contact.url                 https://supne.com/
// @contact.email               team@supnex.com
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @host                        localhost:3000
// @BasePath                    /api/v1
func main() {

	// =====================================================
	// Open Config file
	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	// =====================================================
	// Config logger
	Config.NewLogger("Polaris Storage Service", config.DB.Host)

	// =====================================================
	// Open Test Connection
	//Config.MongoDatabaseOpen(Config.DatabaseConfig{
	//	User:         config.DB.User,
	//	Password:     config.DB.Password,
	//	Host:         config.DB.Host,
	//	Name:         config.DB.Name,
	//	MaxIdleConns: config.DB.MaxIdleConns,
	//	MaxOpenConns: config.DB.MaxOpenConns,
	//	DisableTLS:   config.DB.DisableTLS,
	//})

	Config.MongoDatabaseOpen(Config.MongoDB{Url: config.MongoDB.Url, DBname: config.MongoDB.DBname})
	//defer func() {
	//	Config.DB.Close()
	//	fmt.Println("database in disconnected ")
	//}()

	// =====================================================
	// Create new token maker
	token.NewPasetoMaker(config.Token.TokenSymmetricKey)

	// =====================================================
	// Create new token maker
	Config.RedisConnection(config)

	// =====================================================
	// Run server
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20
	app.Static("/assets/", "./public")
	Router.Routes(app)

	errorChannel := make(chan error)
	func() {
		Config.Logger.Infow("Project Running On PORT", config.Api.ApiHost)
		errorChannel <- app.Run(config.Api.ApiHost)
	}()

}
