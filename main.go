package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"goapi/cache/redis"
	"goapi/config"
	"goapi/model"
	"goapi/router"
	"goapi/router/middleware"
	"log"
	"net/http"
	"time"
)

var (
	//命令行参数 使用的配置文件的地址 默认为空
	configFile = pflag.StringP("ConfigFile", "c", "", "config file path and name.")
)

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*configFile); err != nil {
		log.Fatal(err.Error())
	}

	//init db
	model.DB.Init()
	defer model.DB.Close()

	//init Redis
	redis.RH.Init()
	defer redis.RH.Close()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{
		middleware.Requestid(),
	} //指定中间件

	//加载路由
	router.Load(g, middlewares)

	go func() { //监控服务状态
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up : ", err)
		}
		log.Printf("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))

	//log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
	log.Printf(g.Run(viper.GetString("addr")).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		res, err := http.Get(viper.GetString("url") + "/server/health")
		if err == nil && res.StatusCode == http.StatusOK {
			return nil
		}

		log.Printf("Waiting for the router, retry in 1 second.")

		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}

/*
├─ Project Name
│  ├─ config          //配置文件
│     ├── ...
│  ├─ controller      //控制器层
│     ├── ...
│  ├─ service         //业务层
│     ├── ...
│  ├─ repository      //数据库操作层
│     ├── ...
│  ├─ model           //数据库ORM
│     ├── ...
│  ├─ entity          //实体
│     ├── ...
│  ├─ proto           //proto文件
│     ├── ...
│  ├─ router          //路由
│     ├── middleware  //路由中间件
│         ├── ...
│     ├── ...
│  ├─ util            //工具类
│     ├── ...
│  ├─ main.go         //入口文件
*/
