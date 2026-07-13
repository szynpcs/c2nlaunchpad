package main

import (
	"c2n/config"

	"github.com/joho/godotenv"

	"c2n/api/handler"
	"c2n/database"
	"c2n/repository"
	"c2n/route"
	"c2n/service"

	"c2n/logger"
	"log"
)

func main() {
	// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 加载配置文件
	config.LoadConfig()

	// 初始化日志记录器
	logger.InitLogger()

	// 初始化数据库
	db := database.InitializeDB()
	// 实例化 repository
	productContractRepo := repository.NewProductContractRepository(db)

	// 实例化 service
	productContractService := service.NewProductContractService(productContractRepo)
	//实例化 合约查询初始化
	saleContractService := service.NewSaleContractService(productContractService)
	err := saleContractService.StartSaleFactoryListen()
	if err != nil {
		log.Fatal(err)
	}
	productContractHandler := handler.NewProductContractHandler(productContractService)

	encodeService := service.NewEncodeService()

	// 实例化 handler
	encodeHandler := handler.NewEncodeHandler(encodeService)

	helloHandler := handler.NewHelloHandler()

	r := route.SetupRouter(
		encodeHandler,
		helloHandler,
		productContractHandler,
	)

	r.Run(":" + config.AppConfig.Port) // 启动服务
}
