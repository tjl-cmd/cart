package main

import (
	"cart/common"
	"cart/domain/repository"
	service2 "cart/domain/service"
	"cart/handler"
	pb "cart/proto"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "go.micro.cart.service"
	version = "latest"
	qps = 100
)

func main() {
	// 配置中心
	srv, db := Init(service, version, "127.0.0.1", 8500, []string{"127.0.0.1:8500"}...)
	// 配置链路追踪
	t, io, err := common.NewJaegerTracer(service, "localhost:6831")
	if err != nil {
		log.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	rp := repository.NewCartRepository(db)
	err = rp.InitTable()
	// 初始化数据库表
	if err != nil {
		log.Error("初始化数据库表失败")
	}
	cartDataService := service2.NewCartDataService(rp)
	// Register handler
	pb.RegisterCartHandler(srv.Server(), &handler.Cart{CateDatService: cartDataService})
	defer db.Close()

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func Init(service, version, address string, port int64, consulAddress ...string) (micro.Service, *gorm.DB) {
	// 配置中心
	consulConfig, err := common.GetConsulConfig(address, port, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = consulAddress
	})

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address("127.0.0.1:8087"),
		// 添加consul，作为注册中心
		micro.Registry(consulRegistry),
		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(qps)),
	)

	// 获取mysql配置,路径中不带前缀
	sql := common.GetMysqlFromConsul(consulConfig, "mysql")
	if sql.User == "" || sql.Host == "" || sql.Port == 0 || sql.Pwd == "" || sql.Database == "" {
		log.Error("初始化配置失败")
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", sql.User, sql.Pwd, sql.Host, sql.Port, sql.Database)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Error(err)
	}
	db.SingularTable(true)

	srv.Init()
	return srv, db
}
