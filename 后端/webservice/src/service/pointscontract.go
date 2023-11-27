package service
import (
    "fmt"
    "log"
	"fintech/points"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)
func DeployPointsContractOnce(){//只需调用一次
	fmt.Println("-------------------starting deploy points contract-----------------------")
	client := InitContractClient()
	address, receipt, _, err := points.DeployEshouPoint(client.GetTransactOpts(), client)//部署合约
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("points contract address: ", address.Hex()) // the address should be saved, will use in next example
	fmt.Println("points transaction hash: ", receipt.TransactionHash)
	fmt.Println("-------------------End deploying points contract-----------------------")
 }
 func PointsSessionObjectReturn()*points.EshouPointSession{//返回一个跟合约通信的对象
	client := InitContractClient()
	contractAddress :=common.HexToAddress("0x0102e8B6fC8cdF9626fDdC1C3Ea8C1E79b3FCE94")//传入积分合约的部署地址
	instance,err :=points.NewEshouPoint(contractAddress,client)
	if err !=nil{
		log.Fatal(err)
	}
	pointsSession :=&points.EshouPointSession{Contract:instance,CallOpts:*client.GetCallOpts(),TransactOpts:*client.GetTransactOpts()}//创建一个合约通信对象
	return pointsSession
}
func ReturnBalance(address string)(*big.Int,error){
	//返回账户余额
	session := PointsSessionObjectReturn()
	useraddress := common.HexToAddress(address)
	res,err := session.GetBalance(useraddress)
	if err!=nil{
		// log.Fatal(err)
		return nil,err
	}
	return res,nil

}
func IsAccountExist(address string)(bool){
	session := PointsSessionObjectReturn()
	useraddress := common.HexToAddress(address)
	res,err := session.IsAccountExist(useraddress)
	if err!=nil{
		log.Fatal(err)
		return false
	}else{
		return res
	}
}
func AddBalance(address string,points int)(bool){
	session := PointsSessionObjectReturn()
	useraddress := common.HexToAddress(address)
	p := big.NewInt(int64(points))//int型与*big.int的转换
	res,_,_,err := session.AddBalance(useraddress,p)
	if err!=nil{
		log.Fatal(err)
		return false
	}else{
		return res
	}
}
func AddAccount(address string)(error){
	//将用户的公钥地址注册到区块链上
	session := PointsSessionObjectReturn()
	useraddress := common.HexToAddress(address)
	_,_,err := session.AddAccount(useraddress)
	if err!=nil{
		return err
	}else{
		return nil
	}
}