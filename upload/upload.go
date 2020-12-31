/*
 * 文件上传
 */

package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	//创建OSSClient实例
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "xxxx", "xxxx")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//获取存储空间
	bucket, err := client.Bucket("nice-lingke")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//上传本地文件
	err = bucket.PutObjectFromFile("path/file.txt", "file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

/*
	Applications/Adobe Photoshop CC 2017 $
	iskyMbp: /Applications/Adobe Lightroom CC $
*/
