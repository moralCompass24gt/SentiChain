package service
import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)
func generateHexPublicKey() (string, error) {
	// 定义公钥字节长度（示例值，具体根据你的需求选择）
	// 用于生成上传图片的名字
	keyLength := 32

	// 生成随机的公钥字节
	publicKeyBytes := make([]byte, keyLength)
	_, err := rand.Read(publicKeyBytes)
	if err != nil {
		return "", err
	}

	// 将公钥字节转换为十六进制字符串
	hexPublicKey := hex.EncodeToString(publicKeyBytes)

	return hexPublicKey, nil
}
func convertAndRemoveDuplicates(bigIntSlice []*big.Int) []int {
	// 将区块链返回的结果进行转换并做去重
	// 使用 map[int]struct{} 做去重
	uniqueSet := make(map[int]struct{})
	var intSlice []int

	for _, bigInt := range bigIntSlice {
		// 将 *big.Int 转换为 int
		intValue := int(bigInt.Int64())

		// 检查是否已经存在于去重集合中
		if _, ok := uniqueSet[intValue]; !ok {
			// 不存在则添加到结果切片和去重集合中
			intSlice = append(intSlice, intValue)
			uniqueSet[intValue] = struct{}{}
		}
	}

	return intSlice
}