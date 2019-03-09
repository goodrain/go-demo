package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	dbinfo_http "github.com/goodrain/go-demo/dbinfo/delivery/http"
	dbinfo_repo "github.com/goodrain/go-demo/dbinfo/repository"
	dbinfo_ucase "github.com/goodrain/go-demo/dbinfo/usecase"
	foobar_http "github.com/goodrain/go-demo/foobar/delivery/http"
	foobar_ucase "github.com/goodrain/go-demo/foobar/usecase"
	"github.com/goodrain/go-demo/middleware"
	proxy_http "github.com/goodrain/go-demo/proxy/delivery/http"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	dbuser := os.Getenv("MYSQL_USER")
	dbpw := os.Getenv("MYSQL_PASS")
	dbhost := os.Getenv("MYSQL_HOST")
	dbport := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpw, dbhost, dbport, dbname)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())
	dbconn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		logrus.Warningf("error opening database: %v", err)
	}
	defer dbconn.Close()

	dbinfoRepo := dbinfo_repo.NewMysqlDBInfoRepository(dbconn)

	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	e.Static("/", "public")

	dbinfoUcaser := dbinfo_ucase.NewDBInfoUsecase(dbinfoRepo)
	dbinfo_http.NewDBInfoHTTPHandler(e, dbinfoUcaser)

	foobarUcaser := foobar_ucase.NewFoobarUsecase()
	foobar_http.NewFoobarHandler(e, foobarUcaser)

	proxy_http.NewProxyHandler(e, nil)

	e.Logger.Fatal(e.Start(":5000"))
}
