package service
/*
调用数据库服务
*/

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"sync"
	"log"
)
var (
	dbInstance *gorm.DB
	once sync.Once
)
type Story struct {
    StoryID       int           `gorm:"primaryKey;AUTO_INCREMENT"`
	StoryDetail            string `gorm:"type:text;not null"`
	CID     int			`gorm:"not null"`
    Time          time.Time     `gorm:"not null;default:current_timestamp;type:timestamp"`
}
type Commodity struct {//添加一个flag 标志
	//Commodity has one Story
    CommodityID     int    `gorm:"primaryKey;AUTO_INCREMENT"`
    Extent           float64 `gorm:"not null"`
    ExpectPrice      float64`gorm:"not null"`
    Pic              string `gorm:"not null"`
	CommodityName string `gorm:"not null"`
	Time          time.Time     `gorm:"not null;default:current_timestamp;type:timestamp"`
	CommodityDescription string `gorm:"not null"`
	StoryDetail			string `gorm:"type:text;not null"`
	PublicAddress string `gorm:"not null "` //用户的公钥作为外键
	IsSold	     bool	`gorm:"default:false"`
}
type User struct {
	//User has many commodities
    UserID     int    `gorm:"primaryKey;AUTO_INCREMENT"`
    PublicAddress string `gorm:"not null "`
    Account           string `gorm:"not null"`//账户名
	PrivateKey string `gorm:"not null"`
}
type Trade struct {
	//User has many commodities
    TradeID     int    `gorm:"primaryKey;AUTO_INCREMENT"`
    SellerAddress string `gorm:"not null "`
    BuyerAddress string `gorm:"not null "`
	Price   int  `gorm:"not null "`
	CommodityID int `gorm:"not null"`
	Time          time.Time     `gorm:"not null;default:current_timestamp;type:timestamp"`
}
func Init()*gorm.DB{
	//连接数据库并初始化(单例模式)
	once.Do(func() {
		username := "root"
		password := "123456"
		host := "0.0.0.0"
		port := 3307
		Dbname :="fintech"
		timeout := "10s"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
		panic("连接数据库失败, error=" + err.Error())
		}
		// fmt.Println(db)
		// 连接成功
		fmt.Println("数据库连接成功")
		dbInstance = db
	})
	//fmt.Println("单例模式成功")
	return dbInstance
	//CreateTable(db)
}
func CreateTable(){//创建表并插入数据
	db := Init()//获取数据库操作指针
	// db.AutoMigrate(&User{},&Commodity{},&Story{})
	// fmt.Println("三个表创建成功")
	story := Story{
		StoryDetail:"这是一个搞笑的故事",
		CID:1,
	}
	db.Create(&story)
	// commodity :=  Commodity{
	// 	Extent:0.3,
	// 	ExpectPrice:1005.6,
	// 	Pic:"https://www.haozekeji.com/d/file/2022/07/20220707140753_6c3d0005528dffecbf91.png.png",
	// 	CommodityName:"原神手办",
	// 	CommodityDescription:"我去,原深",
	// 	StoryDetail:"这是一个搞笑的故事",
	// 	PublicAddress:"b20b7e23ee51e804b4de698d60127abc329a1f53ac701338f40365ec39918c86ef4c63ce2f7ebc7b73db969891318107ceb452ad736ebdb800f94f8b3f8071dc"}
	// user := User{
	// 	PublicAddress:"b20b7e23ee51e804b4de698d60127abc329a1f53ac701338f40365ec39918c86ef4c63ce2f7ebc7b73db969891318107ceb452ad736ebdb800f94f8b3f8071dc",
	// 	Account:"0xF34B5F87497A5092F249b7306F04961E5F65eb24",
	// 	PrivateKey:"1503e2135a3ed614128dbdcb107e7995250170dce46346fce2b192fdd011d237",
	// }
	// db.Create(&user)
	// InsertIntoCommodity(commodity,story,user.PublicAddress)
	// fmt.Println("用户表插入成功")
	// db.Create(&User{
	// 	PublicAddress:"8d74ca8b06ddb9a2834cae6b7c7998a26781c390dbe7b44091edacd5d40aa2bd32013cfdde32f8c5d1c0ae54917567acf24c122eed8bbe336566db6e21911c9a",
	// 	Account:"0x4c44652DcB4729F1b590C68cb21E7cC691D1E281",
	// 	PrivateKey:"88d233812e46a2b3c59ca803e6bf4217836a827c024e12d946d171d8aa6f49fc",
	// })
	// fmt.Println("用户表插入成功")
	// db.Create(&User{
	// 	PublicAddress:"80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646",
	// 	Account:"0x08B61a78167D125839f91972e595AC2A733793a1",
	// 	PrivateKey:"bf1f9e809434d337105afdd1247872151eb2eb5f0ba4866b0c6f2fa4450e100a",
	// })
	// fmt.Println("用户表插入成功")
}
func CreateTradeTable(){
	db := Init()//获取数据库操作指针
	if err:=db.AutoMigrate(&Trade{}).Error;err!=nil{
		fmt.Println(err)
	}
	fmt.Println("创建成功")
}
func RecreateUserTable(){
	db := Init()
	db.AutoMigrate(&User{})
	fmt.Println("User表创建成功")
	for i:=1;i<=3;i++{
		privatekey,publicKey,account :=GenerateKeies()
		db.Create(&User{
			PublicAddress:publicKey,
			Account:account,
			PrivateKey:privatekey})
		fmt.Printf("User表数据成功插入%v\n",i)
	}
}
func UpdateTable(){
	db := Init()
	err :=db.AutoMigrate(&Commodity{}).Error
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Success")
	}
}
func SelectFromCommdities(commodityid int)(Commodity,error){
	db := Init()
	//从Commdities表中查找数据并返回单条数据
	var commodity Commodity
	if err:=db.Take(&commodity,"commodity_id = ?",commodityid).Error;err!=nil{
		return Commodity{},err
	}//查询单条记录
	return commodity,nil
}
func SelectFromCommditiesReturnSlice(condition []int)([]Commodity,error){
	var commoditylist []Commodity
	db := Init()
	if condition==nil{
		//从Commdities表中查找数据并返回所有数据,将is_sold为false的进行返回
		if err:=db.Where("is_sold = ?",false).Find(&commoditylist).Error;err!=nil{
			return []Commodity{},err
		}else{
			return commoditylist,nil
		}
	}else{
		if err:=db.Where("commodity_id in ?",condition).Find(&commoditylist).Error;err!=nil{
			return []Commodity{},err
		}else{
			return commoditylist,nil
		}
	}
}
func SelectFromStories(storyid int)(Story,error){
	//从Stories表中查找数据并返回单条数据
	db := Init()
	var story  Story
	if err:=db.Take(&story,"story_id = ?",storyid).Error;err!=nil{
		return Story{} ,err
	}
	return story,nil
}
// func SelectFromStoriesReturnSlice()([]Story,error){
// 	//从Stories表中查找数据并返回所有数据
// 	db := Init()
// 	if err:=db.Find(&storylist).Error;err!=nil{
// 		return []Story{},err
// 	}
// 	return storylist,nil
// }
func SelectFromStoriesByCommodityId(storyid int)(Commodity,error){
	db := Init()
	//根据故事公钥地址查找商品的详细信息
	var story  Story
	var commodity Commodity
	if err:=db.Where("story_id = ?",storyid).Take(&story).Error;err!=nil{
		return Commodity{},err
	}
	// fmt.Println(story)
	if err:=db.Where("commodity_id = ?",story.CID).Take(&commodity).Error;err!=nil{
		return Commodity{},err
	}
	return commodity,nil
}
func InsertIntoCommodity(commodity Commodity,story Story)(Commodity,bool){
	db := Init()
	//插入数据
	result := db.Create(&commodity)
	if result.Error !=nil{
		return Commodity{},false
	}
	if result.RowsAffected == 0{
		return Commodity{},false
	}
	story.CID = commodity.CommodityID
	result = db.Create(&story)
	if result.Error!=nil{
		return Commodity{},false
	}
	return commodity,true
}
func UserAddressValidation(address string)(bool){
	db:=Init()
	var user User
	if err:=db.Where("public_address = ?",address).Take(&user).Error;err!=nil{
		// log.Fatal(err)
		return false
	}else{
		return true
	}
}
func StoryAddressValidation(storyid int)(bool){
	db := Init()
	var story Story
	if err:=db.Where("story_id = ?",storyid).Take(&story).Error;err!=nil{
		log.Fatal(err)
		return false
	}else{
		return true
	}
}
func CommodityAddressValidation(commodityid int)(bool){
	db := Init()
	var commodity Commodity
	if err:=db.Where("commodity_id = ?",commodityid).Take(&commodity).Error;err!=nil{
		log.Fatal(err)
		return false
	}else{
		return true
	}
}
func BeforeExchangeGoods(sellerAddress string,goodid int)(bool){
	//检查目前链上是否有该商品并且商品的公钥地址为卖家地址
	var commodity Commodity
	db := Init()
	res,_ := IndexGoodsOnList()
	cidlist := convertAndRemoveDuplicates(res)
	if cidlist == nil{
		return false
	}else{
		db.Where("commodity_id = ?",goodid).Take(&commodity)
		if commodity.PublicAddress == sellerAddress{
			return true
		}else{
			return false
		}
	}
}
func AfterExchangeGoods(buyerAddress string,goodid int)(error){
	db := Init()
	//在交易完成后需要将原来商品的公钥地址转换成买家的公钥地址并修改is_sold标志位
	if err:=db.Model(&Commodity{}).Where("commodity_id = ?",goodid).Update("public_address",buyerAddress).Error;err!=nil{
		return err
	}else{
		if err = db.Model(&Commodity{}).Where("commodity_id = ?",goodid).Update("is_sold",true).Error;err!=nil{
			return err
		}else{
			return nil
		}
	}
}
func InsertIntoTrade(sellerAddress string,buyerAddress string,goodid int,price int)(error){
	db :=Init()
	trade := Trade{
		SellerAddress:sellerAddress,
		BuyerAddress:buyerAddress,
		CommodityID:goodid,
		Price:price,
	}
	if err:=db.Create(&trade).Error;err!=nil{
		return err
	}else{
		return nil
	}
}
func SearchTradeOfBuyInfo(address string)([]int,error){
	//返回用户卖过的商品ID
	db :=Init()
	var commodityidlist []int
	if err:=db.Model(&Trade{}).Where("buyer_address = ?",address).Pluck("commodity_id", &commodityidlist).Error; err != nil{
		return nil,err
	}
	return commodityidlist,nil
}
func SearchTradeOfSellInfo(address string)([]int,error){
	//返回用户买过的商品ID
	db :=Init()
	var commodityidlist []int
	if err:=db.Model(&Trade{}).Where("seller_address = ?",address).Pluck("commodity_id", &commodityidlist).Error; err != nil{
		return nil,err
	}
	return commodityidlist,nil
}
func SearchUserGoodsDetail(address string)([]Commodity,error){
	// 获取用户所发布的所有商品信息
	db :=Init()
	var commoditylist []Commodity
	if err:=db.Model(&Commodity{}).Where("public_address = ?",address).Find(&commoditylist).Error;err!=nil{
		return []Commodity{},err
	}
	return commoditylist,nil
}