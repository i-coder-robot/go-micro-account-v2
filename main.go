package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/i-coder-robot/go-micro-account-v2/config"
	"github.com/i-coder-robot/go-micro-account-v2/handler"
	user "github.com/i-coder-robot/go-micro-account-v2/proto"
	"github.com/i-coder-robot/go-micro-account-v2/repository"
	"github.com/i-coder-robot/go-micro-account-v2/service"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

const (
	DriverName = "mysql"
	IP         = "192.168.43.254"
	//IP         = "192.168.0.106"
	// user:password@tcp(container-name:port)/dbname ※mysql
	// DataSourceName = "root:root@tcp(mysql-container:3306)/foods_srv?charset=utf8&parseTime=True&loc=Local"
	DataSourceName = "root:smartwell@tcp(127.0.0.1:3306)/foods_srv?charset=utf8&parseTime=True&loc=Local"
)

func main() {
	consulConfig,err := config.GetConsulConfig("127.0.0.1",8500,"/micro/config")
	if err != nil {
		panic(err)
	}

	registry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t, io, err := config.NewTracer("go.micro.service.account", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8011"),
		micro.Registry(registry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	//路径不带前缀
	mysql4Consul := config.GetMysql4Consul(consulConfig, "mysql")
	fmt.Println(mysql4Consul.User)
	fmt.Println(mysql4Consul.Pwd)
	fmt.Println(mysql4Consul.Database)
	fmt.Println(mysql4Consul.Host)
	srv.Init()
	var db *gorm.DB
	if mysql4Consul.User=="" && mysql4Consul.Pwd==""{
		db, err = gorm.Open(DriverName, DataSourceName)
	}


	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)

	repo := repository.NewUserInterfaceRepository(db)
	repo.InitTable()
	accountSrv := service.NewUserService(repo)

	// Register handler
	err = user.RegisterAccountHandler(srv.Server(), &handler.User{
		AccountService: accountSrv,
	})

	if err != nil {
		panic(err)
	}

	// Run userService
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
