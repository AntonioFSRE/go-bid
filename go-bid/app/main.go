package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_bidHttpDelivery "github.com/AntonioFSRE/go-bid/bid/delivery/http"
	_bidHttpDeliveryMiddleware "github.com/AntonioFSRE/go-bid/bid/delivery/http/middleware"
	_bidRepo "github.com/AntonioFSRE/go-bid/bid/repository/postgres"
	_bidUcase "github.com/AntonioFSRE/go-bid/bid/usecase"
	_userRepo "github.com/AntonioFSRE/go-bid/user/repository/postgres"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Europe/Zagreb")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`postgres`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _bidHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	userRepo := _userRepo.NewPostgresUserRepository(dbConn)
	b := _bidRepo.NewPostgresBidRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	u := _bidUcase.NewBidUsecase(b, userRepo, timeoutContext)
	_bidHttpDelivery.NewBidHandler(e, u)

	log.Fatal(e.Start(viper.GetString("server.address")))
}