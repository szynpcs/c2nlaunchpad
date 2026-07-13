package service

import (
	"c2n/config"
	"c2n/database"
	"c2n/logger"
	"c2n/repository"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)





func TestCalculateEventSignature(t *testing.T){

// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 初始化日志记录器
	logger.InitLogger()
	// 加载配置文件
	config.LoadConfig()
	// 初始化数据库
	db := database.InitializeDB()
	// 实例化 repository
	productContractRepo := repository.NewProductContractRepository(db)

	// 实例化 service
	productContractService := NewProductContractService(productContractRepo)
	saleContractService:= NewSaleContractService(productContractService)
	signResult:=saleContractService.CalculateEventSignature("213213");
	fmt.Printf(signResult.Hex());
	
}

func TestLoadSale(t *testing.T){
// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 初始化日志记录器
	logger.InitLogger()
	// 加载配置文件
	config.LoadConfig()
	// 初始化数据库
	db := database.InitializeDB()
	// 实例化 repository
	productContractRepo := repository.NewProductContractRepository(db)

	// 实例化 service
	productContractService := NewProductContractService(productContractRepo)
	saleContractService:= NewSaleContractService(productContractService)
	prod,err:=saleContractService.QuerySaleInfo("0x8aCd85898458400f7Db866d53FCFF6f0D49741FF");
	if(err!=nil){
		fmt.Println(err)
	}else{
	fmt.Print(prod)
	}


}


func TestListenSaleChange(t *testing.T){

	// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 初始化日志记录器
	logger.InitLogger()
	// 加载配置文件
	config.LoadConfig()
	// 初始化数据库
	db := database.InitializeDB()
	// 实例化 repository
	productContractRepo := repository.NewProductContractRepository(db)

	// 实例化 service
	productContractService := NewProductContractService(productContractRepo)
	saleContractService:= NewSaleContractService(productContractService)
 
	saleContractService.StartSaleFactoryListen()
	
}

 