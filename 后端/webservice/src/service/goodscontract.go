package service
import (
	"encoding/hex"
	"context"
    "fmt"
    "log"
    "github.com/FISCO-BCOS/go-sdk/client"
	"fintech/goods"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"errors"
	// "reflect"
	// "strconv"
)
func InitContractClient()*client.Client{
	privateKey,_ := hex.DecodeString("1503e2135a3ed614128dbdcb107e7995250170dce46346fce2b192fdd011d237")
	config := &client.Config{
	   IsSMCrypto:  false,
	   GroupID:     "group0",
	   PrivateKey:  privateKey,
	   Host:        "127.0.0.1",
	   Port:        20200,
	   TLSCaFile:   "/root/webservice/src/conf/ca.crt", 
	   TLSKeyFile:  "/root/webservice/src/conf/sdk.key",
	   TLSCertFile: "/root/webservice/src/conf/sdk.crt",
   }//配置
   client, err := client.DialContext(context.Background(), config)//与区块链建立连接
   if err != nil {
	   log.Fatal(err)
   }
   return client
}
func DeployGoodsContractOnce(){//只需调用一次
   fmt.Println("-------------------starting deploy goods contract-----------------------")
   client := InitContractClient()
//    fmt.Println(client)
   address, receipt, _, err := goods.DeployEshouGoods(client.GetTransactOpts(), client)//部署合约
   if err != nil {
	   log.Fatal(err)
   }
   fmt.Println("goods contract address: ", address.Hex()) // the address should be saved, will use in next example
   fmt.Println("goods transaction hash: ", receipt.TransactionHash)
   fmt.Println("-------------------End deploying goods contract-----------------------")
}
func GoodsSessionObjectReturn()*goods.EshouGoodsSession{//返回一个跟合约通信的对象
	client := InitContractClient()
	contractAddress :=common.HexToAddress("0xCcEeF68C9b4811b32c75df284a1396C7C5509561")//传入商品合约的部署地址
	instance,err :=goods.NewEshouGoods(contractAddress,client)//加载商品合约
	if err !=nil{
		log.Fatal(err)
	}
	goodsSession :=&goods.EshouGoodsSession{Contract:instance,CallOpts:*client.GetCallOpts(),TransactOpts:*client.GetTransactOpts()}//创建一个合约通信对象
	return goodsSession
}
func ReturnIndexByBuyer(address string)[]*big.Int{//调用合约返回用户购买过的商品
	session :=GoodsSessionObjectReturn()
	buyerAddress := common.HexToAddress(address)
	res,err := session.IndexByBuyer(buyerAddress)
	// fmt.Println(res)
	// t := reflect.TypeOf(res)
	// fmt.Printf("Type of %v : %v\n",res,t)
	if err!=nil{
		log.Fatal(err)
	}
	return res
}
/*
bind.CallOpts:*client.GetCallOpts()
*bind.TransactOpts:client.GetTransactOpts()
bind.ContractBackend:Client 连接对象
*/
func ReturnIndexBySeller(address string)[]*big.Int{//调用合约返回用户卖过的商品
	session :=GoodsSessionObjectReturn()
	buyerAddress := common.HexToAddress(address)
	res,err := session.IndexBySeller(buyerAddress)
	// fmt.Println(res)
	// t := reflect.TypeOf(res)
	// fmt.Printf("Type of %v : %v\n",res,t)
	if err!=nil{
		log.Fatal(err)
	}
	return res
}
func CreateGoods(commodity Commodity,address string)error{
	//调用合约往区块链上存储商品
	session :=GoodsSessionObjectReturn()
	// client := InitContractClient()
	hextoaddress := common.HexToAddress(address)
	_,_,err := session.CreateGoods(big.NewInt(int64(commodity.CommodityID)),commodity.StoryDetail,commodity.Pic,commodity.CommodityName,commodity.CommodityDescription,hextoaddress)
	fmt.Printf("商品id为%v的商品,存在的标志位为%v\n",commodity.CommodityID,IsGoodExist(commodity.CommodityID))
	if err!=nil{
		log.Fatal(err)
		return err
	}else{
		return nil
	}
}
func ExchangeGoods(price int ,buyeraddress string,selleraddres string,cid int)error{
	//调用合约对双方进行交易
	// client := InitContractClient()
	session :=GoodsSessionObjectReturn()
	address := common.HexToAddress(buyeraddress)
	dealprice:=big.NewInt(int64(price))
	gid := big.NewInt(int64(cid))//int型与*big.int的转换
	check :=BeforeExchangeGoods(selleraddres,cid)//交易前的检查
	if check==false{
		return errors.New("commodity not on the blockchain or seller does not have the commodity")
	}
	_,_,err := session.CreateTrade(gid,address,dealprice)//将交易信息上链
	if err!=nil{
		log.Fatal(err)
		return err
	}else{
		err =AfterExchangeGoods(buyeraddress,cid)//修改商品的owner的公钥地址
		if err!=nil{
			log.Fatal(err)
			return err
		}
		err = InsertIntoTrade(selleraddres,buyeraddress,cid,price)//往数据库中增加交易信息
		if err!=nil{
			return err
		}
		return nil
	}
}
func IndexStoryById()([]string,error){
	// 返回链上的所有故事
	session :=GoodsSessionObjectReturn()
	storylist,err := session.IndexStoryById()
	if err!=nil{
		return nil,err
	}else{
		return storylist,nil
	}
}
func IsGoodExist(goodid int)(bool){
	//账本上是否有商品存在
	session :=GoodsSessionObjectReturn()
	gid := big.NewInt(int64(goodid))//int型与*big.int的转换
	res,err := session.IsGoodExist(gid)
	if err!=nil{
		log.Fatal(err)
		return false
	}else{
		return res
	}
}
func IndexGoodsOnList()([]*big.Int,error){
	session :=GoodsSessionObjectReturn()
	res,err := session.IndexGoodsonList()
	if err!=nil{
		return nil,err 
	}else{
		return res,nil
	}
}