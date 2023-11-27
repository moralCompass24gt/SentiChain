package service
/*
对外暴露的网络接口服务
*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	// "crypto/rand"
	// "encoding/hex"
	"strconv"
	"fmt"
)
type goodsDetail struct{
	extent  float64 `json:“extent”`
	price 	float64 `json:”price“`
	story  string `json:“story”`
	pic  string `json:“pic”`
	name string `json:“name”`
	description string `json:“description”`
}
type updateChainInfo struct{
	infoType string `json:“type”`
	infoDetail string `json:“detail”`
}
type exchangeInfo  struct{
	price int `json:“price”`
	buyerAddress string `json:“buyaddress”`//买家公钥地址
	sellerAddress string `json:“selleraddress”`//卖家公钥地址
	cid int `json:“commodityid”`//商品id
}
func UploadPicTest(c *gin.Context){
	pic,err := c.FormFile("pic") //获取前端传入的图片对象
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})//传入的不是一张图片
			return 
		}
		fileType := pic.Header["Content-Type"][0]
		if fileType != "image/png" && fileType != "image/jpg" && fileType != "image/jpeg" && fileType != "image/gif" {
			log.Println("上传图片非jpg,png,jpeg,gif,请重新上传！")
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "上传失败,图片非jpg,png,jpeg,gif,请重新上传！",
			})
			return
		} else if pic.Size/1024 > 2000 {
			log.Println("上传图片大于2M,请重新上传")
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "上传图片大于2M,请重新上传",
			})
			return
		} else {
			f, err := pic.Open()
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":"上传失败，打开文件失败！",
					"ERROR-CONTROLLER": err,
				})
			}
			defer f.Close()
			picaddress := Upload(pic.Filename, f)//返回图片地址
			fmt.Println(picaddress)
			if picaddress == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "上传失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":    "上传成功！",
				// "result": picaddress,
			})
		}
}
func ReleaseGoods(c *gin.Context){
	/*
	前端POST数据，要通过智能合约和数据库进行
	接口参数:用户公钥地址，磨损程度，价格，故事(需要上链)，图片(需要上链),商品名称，商品描述
	将商品上链并将相关信息存储到数据库中
	流程:前端将图片等数据-post->后端->将图片存入腾讯云COS服务器中并获取图片的url->将图片url等信息存入数据库中->通过智能合约将商品信息存储到链上->返回信息给后端
	返回结果:根据statuscode和返回json数据来判断
	*/
	pic,err := c.FormFile("pic") //获取前端传入的图片对象
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})//传入的不是一张图片
		return 
	}
	fileType := pic.Header["Content-Type"][0]
	if fileType != "image/png" && fileType != "image/jpg" && fileType != "image/jpeg" && fileType != "image/gif" {
		log.Println("上传图片非jpg,png,jpeg,gif,请重新上传！")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "上传失败,图片非jpg,png,jpeg,gif,请重新上传！",
		})
		return
	} else if pic.Size/1024 > 2000 {
		log.Println("上传图片大于2M,请重新上传")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "上传图片大于2M,请重新上传",
		})
		return
	} else {
		f, err := pic.Open()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":"上传失败，打开文件失败！",
				"ERROR-CONTROLLER": err,
			})
		}
		defer f.Close()
		picaddress := Upload(pic.Filename, f)//返回图片地址
		fmt.Println(picaddress)
		if picaddress == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "图片上传失败",
			})
			return
		}
		extent,err := strconv.ParseFloat(c.PostForm("extent"),64)//前端请求需要使用multipart/form-data类型和传入对应的json参数名字
		price,err := strconv.ParseFloat(c.PostForm("price"),64)
		story := c.PostForm("story")
		description := c.PostForm("description")
		name := c.PostForm("name")
		useraddress := c.PostForm("address")//用户公钥地址,需要判断是否合法
		exit := UserAddressValidation(useraddress)
		if exit==false{
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "用户公钥地址不存在发布失败",
			})
			return 
		}
		// storyPublickey,_ := generateHexPublicKey()
		commodity := Commodity{
			// CommodityAddress:hexPublicKey,
			Extent:extent,
			ExpectPrice:price,
			Pic:picaddress,
			CommodityDescription:description,
			CommodityName:name,
			StoryDetail :story,
			PublicAddress:useraddress,
		}
		insertstory := Story{
			StoryDetail :story,
		}
		newcommodity,err2 := InsertIntoCommodity(commodity,insertstory) //往数据库中插入数据,同时插入商品表和故事表
		if err2==false{
			c.JSON(http.StatusInternalServerError,gin.H{
				"msg": "数据库更新失败",
			})
			return 
		}
		err3 := CreateGoods(newcommodity,useraddress) //通过合约来将商品放到链上
		err4 := AddBalance(useraddress,10) //发布商品成功则增加10积分
		if err3==nil && err4==true{
			c.JSON(http.StatusOK,gin.H{
				"msg": "发布成功",
			})
			return 
		}else{
			c.JSON(http.StatusInternalServerError,gin.H{
				"msg": err3,
			})
			return 
		}
	}
}
func ConfirmOrderBySeller(c *gin.Context){
	/*
	前端POST数据，要通过智能合约和数据库进行
	接口参数:价格(需要上链)，买家与卖家的公钥地址(需要上链),商品公钥地址
	返回结果:boolean(确认成功或失败)
	*/
	price := c.PostForm("price")
	cid := c.PostForm("commodityid")
	buyerAddress := c.PostForm("buyeraddress")
	sellerAddress := c.PostForm("selleraddress")
	buyerexit := UserAddressValidation(buyerAddress)
	sellererexit := UserAddressValidation(sellerAddress)
	id,_ :=strconv.Atoi(cid)
	commodityexit := CommodityAddressValidation(id)
	if buyerexit == false || commodityexit == false || sellererexit==false{
		c.JSON(http.StatusNotFound,gin.H{"msg": "用户公钥地址或商品id不存在",})
		return 
	}else{
		numprice,_ := strconv.Atoi(price)
		err := ExchangeGoods(numprice,buyerAddress,sellerAddress,id)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"msg": "交易失败",
			})
			return 
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg": "交易成功",
			})
			return 
		}
	}
	
}
func SearchIntergration(c *gin.Context){
	/*
	要通过智能合约进行
	接口参数:用户公钥地址
	返回结果:返回存储在链上的用户积分
	*/
	useraddress := c.Query("address")
	exit := UserAddressValidation(useraddress)
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"用户公钥地址不存在",
		})
		return
	}
	res,err := ReturnBalance(useraddress)
	fmt.Println(res)
	fmt.Println(err)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"balance":res,
		})
	}
}
func ReturnSellGoodsItem(c *gin.Context){
	/*
	要通过智能合约进行
	接口参数:用户公钥地址
	返回结果:返回用户卖出的商品列表
	*/
	useraddress := c.Query("address")
	exit := UserAddressValidation(useraddress)
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"用户公钥地址不存在",
		})
		return
	}
	// res := ReturnIndexBySeller(useraddress)
	// cidlist := convertAndRemoveDuplicates(res)//将返回的结果进行去重并转换类型

	cidlist2,err := SearchTradeOfSellInfo(useraddress)
	// fmt.Println(cidlist)
	list,err :=SelectFromCommditiesReturnSlice(cidlist2)
	// fmt.Printf("返回的结果为:%v\n",list)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"commoditieslist":list,
		})
	}
}
func ReturnBuyGoodsItem(c *gin.Context){
	/*
	要通过智能合约进行
	接口参数:用户公钥地址
	返回结果:返回用户已买的商品列表
	*/
	useraddress := c.Query("address")
	exit := UserAddressValidation(useraddress)
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"用户公钥地址不存在",
		})
		return
	}
	// res := ReturnIndexByBuyer(useraddress)//返回的是big.Int类型的切片
	// cidlist := convertAndRemoveDuplicates(res)

	cidlist2,err := SearchTradeOfBuyInfo(useraddress)
	list,err :=SelectFromCommditiesReturnSlice(cidlist2)
	// fmt.Printf("返回的结果为:%v\n",list)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"commoditieslist":list,
		})
	}

}
func Skip2GoodsDetail(c *gin.Context){
	/*
	接口参数:故事id
	返回结果:商品的详情信息
	*/
	storyid := c.Query("storyid")
	id,err := strconv.Atoi(storyid)
	exit := StoryAddressValidation(id)
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"故事公钥地址不存在",
		})
		return
	}
	res,err := SelectFromStoriesByCommodityId(id)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"not found",
		})
	}else{
		c.JSON(http.StatusOK,res)
	}
}
func ReturnStoryItem(c *gin.Context){
	/*
	接口参数:无
	返回结果:返回故事列表
	*/
	// res,err := SelectFromStoriesReturnSlice()
	db := Init()
	var result []struct{
		Story
		PublicAddress string
	}
	err:=db.Table("stories").Select("stories.*,commodities.public_address").Joins("LEFT JOIN commodities ON stories.c_id = commodities.commodity_id").Scan(&result).Error
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"storieslist":result,
		})
	}
}
func ReturnGoodsItem(c *gin.Context){
	/*
	接口参数:无
	返回结果:返回商品列表，flag为false的物品,交易成功将标志位设为true(false-->想要出售，true-->暂时不想出售)
	*/
	res,err := SelectFromCommditiesReturnSlice(nil)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"mes":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"commoditieslist":res,
		})
	}
}
func ReturnGoodsDetail(c *gin.Context){
	/*
	接口参数:商品id
	返回结果:返回商品详细信息
	*/
	commodityid := c.Query("commodityid")//从请求参数中获取商品信息
	res,_ :=strconv.Atoi(commodityid)
	exit := CommodityAddressValidation(res)
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"商品id不存在",
		})
		return
	}
	data,err := SelectFromCommdities(res)//根据商品私钥信息返回商品详细信息
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		c.JSON(http.StatusOK,data)
	}
}
func ReturnAllStories(c *gin.Context){
	//返回存储在链上的所有故事
	res,err := IndexStoryById()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		if res==nil{
			c.JSON(http.StatusNotFound,gin.H{
				"msg":"not found",
			})
		}
		fmt.Println(res)
		c.JSON(http.StatusOK,gin.H{
			"story":res,
		})
	}
}
func ReturnGoodsNotSell(c *gin.Context){
	res,err := IndexGoodsOnList()
	// fmt.Println(res)
	cidlist := convertAndRemoveDuplicates(res)
	list,err := SelectFromCommditiesReturnSlice(cidlist)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		if res==nil{
			c.JSON(http.StatusNotFound,gin.H{
				"msg":"not found",
			})
		}else{
			c.JSON(http.StatusOK ,gin.H{
				"commoditylist":list,
			})
		}
	}
}
func ReturnUserReleaseGoodsHistory(c *gin.Context){
	// 查询用户所发布的所有商品信息
	useraddress := c.Query("address")
	// res,err := IndexGoodsOnList()
	// // fmt.Println(res)
	// cidlist := convertAndRemoveDuplicates(res)
	exit := UserAddressValidation(useraddress)//用户公钥地址校验
	if exit==false{
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"用户公钥地址不存在",
		})
		return
	}
	list,err := SearchUserGoodsDetail(useraddress)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err,
		})
	}else{
		if list==nil{
			c.JSON(http.StatusNotFound,gin.H{
				"msg":"not found",
			})
		}else{
			c.JSON(http.StatusOK ,gin.H{
				"commoditylist":list,
			})
		}
	}
}