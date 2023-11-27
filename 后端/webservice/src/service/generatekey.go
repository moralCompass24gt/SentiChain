package service
// 生成公私钥和账户地址，并返回私钥的二进制数据用户创建连接
import (
    "crypto/ecdsa"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

func GenerateKeies()(string,string,string){
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key binary:",privateKeyBytes)
    fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
    fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:])  // publicKey in hex without "0x"

    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Println("address: ", address)  // account address
	return hexutil.Encode(privateKeyBytes)[2:],hexutil.Encode(publicKeyBytes)[4:],address
}