package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	// "net/http"
	"fintech/service"
)

func main() {
	// router := mux.NewRouter()//开启cors服务
	r := gin.Default()        //默认的路由引擎
	r.Use(cors.Default())
	r.GET("/goodsdetail", service.ReturnGoodsDetail)//localhost:9090/goodsdetail?commodityid=
	r.GET("/goodsitem", service.ReturnGoodsItem)//localhost:9090/goodsitem
	r.GET("/storiesitem", service.ReturnStoryItem)//localhost:9090/storiesitem
	r.GET("/togoodsdetail", service.Skip2GoodsDetail)//localhost:9090/togoodsdetail?storyid=
	r.GET("/returnallstories", service.ReturnAllStories)//localhost:9090/returnallstories 返回存储在链上的所有故事
	r.GET("/returngoodsnotsell", service.ReturnGoodsNotSell)//localhost:9090/returngoodsnotsell 返回链上尚未卖出的商品
	r.GET("/searchscore",service.SearchIntergration)//localhost:9090/searchscore?address=
	r.POST("/releasegoods", service.ReleaseGoods)//localhost:9090/releasegoods
	r.GET("/goodsitembuy",service.ReturnBuyGoodsItem)//localhost:9090/goodsitembuy?address=

	r.GET("/usersgoodsdetail",service.ReturnUserReleaseGoodsHistory)//localhost:9090/sersgoodsdetail?address=
	r.GET("/goodsitemsell",service.ReturnSellGoodsItem)//localhost:9090/goodsitemsell?address=

	r.POST("/confirmorder", service.ConfirmOrderBySeller)//localhost:9090/confirmorder

	//启动服务
	r.Run(":9090")//请求地址 url:9090/servicename

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// // 处理预检请求 
	// router.Methods(http.MethodOptions).Handler(handlers.CORS(originsOk, headersOk, methodsOk)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
	// 	w.WriteHeader(http.StatusOK) 
	// })))
	// // Start the server with CORS middleware
	// http.ListenAndServe(":9090", handlers.CORS(originsOk, headersOk, methodsOk)(router))

	// service.DeployGoodsContractOnce() //部署合约，已经部署完成不需要再调用
	// service.DeployPointsContractOnce()

	// service.CreateTradeTable()

	// err:=service.AddAccount("b20b7e23ee51e804b4de698d60127abc329a1f53ac701338f40365ec39918c86ef4c63ce2f7ebc7b73db969891318107ceb452ad736ebdb800f94f8b3f8071dc")//手动注册
	// err2:=service.AddAccount("8d74ca8b06ddb9a2834cae6b7c7998a26781c390dbe7b44091edacd5d40aa2bd32013cfdde32f8c5d1c0ae54917567acf24c122eed8bbe336566db6e21911c9a")
	// err3:=service.AddAccount("80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646")
	// fmt.Printf("报错信息为%v\n",err)
	// fmt.Printf("报错信息为%v\n",err2)
	// fmt.Printf("报错信息为%v\n",err3)

	// service.UpdateTable()

}

