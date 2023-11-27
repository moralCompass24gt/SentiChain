package service
//图片上传服务
import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	//"os"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)
//请求域名:http://pic-1304105328.cos.ap-guangzhou.myqcloud.com
func getCos() *cos.Client {
	viper.SetConfigName("conf")//配置文件名字
	viper.SetConfigType("toml")//配置文件类型
	viper.AddConfigPath("/root/webservice/src/conf")           //配置文件搜索路径
	err := viper.ReadInConfig()//开始读取配置文件
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil
	}
	u, _ := url.Parse(fmt.Sprintf(
		viper.GetString("tencentCOS.COS_URL_FORMAT"),
		viper.Get("tencentCOS.COS_BUCKET_NAME"),
		viper.Get("tencentCOS.COS_APP_ID"),
		viper.Get("tencentCOS.COS_REGION"),
	))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("tencentCOS.COS_SECRET_ID"),
			SecretKey: viper.GetString("tencentCOS.COS_SECRET_KEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
				Writer:         nil,
			},
		},
	})
	return c
}
func CheckErr(err error) bool {
	if err == nil {
		return true
	}
	if cos.IsNotFoundError(err) {
		// WARN
		log.Println("WARN: Resource is not existed")
		return false
	} else if e, ok := cos.IsCOSError(err); ok {
		log.Println(fmt.Sprintf("ERROR: Code: %v\n", e.Code))
		log.Println(fmt.Sprintf("ERROR: Message: %v\n", e.Message))
		log.Println(fmt.Sprintf("ERROR: Resource: %v\n", e.Resource))
		log.Println(fmt.Sprintf("ERROR: RequestID: %v\n", e.RequestID))
		return false
		// ERROR
	} else {
		log.Println(fmt.Sprintf("ERROR: %v\n", err))
		return false
		// ERROR
	}
}
func Upload(fileName string, file io.Reader) string {
	c := getCos()
	_, err := c.Object.Put(context.Background(), fileName, file, nil)
	if CheckErr(err) {
		return fmt.Sprintf(
			viper.GetString("tencentCOS.COS_URL_FORMAT")+"/%v",
			viper.GetString("tencentCOS.COS_BUCKET_NAME"),
			viper.GetString("tencentCOS.COS_APP_ID"),
			viper.GetString("tencentCOS.COS_REGION"),
			fileName,
		)
	}
	//imagepath := "http://pic-1304105328.cos.ap-guangzhou.myqcloud.com" +"/"+fileName
	return ""
}