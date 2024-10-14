package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 下载图片并返回image.Image
func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法获取图片: %v", err)
	}
	defer resp.Body.Close()

	// 读取图片数据
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取图片数据失败: %v", err)
	}

	return imgData, nil
}

// 计算图片的MD5哈希
func calculateMD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
	url1 := "https://api.faviconkit.com/google.com/128" // 替换为实际链接
	url2 := "https://api.faviconkit.com/google.com/128" // 替换为实际链接

	// 比较两个图片链接是否相同
	img1, err := downloadImage(url1)
	if err != nil {
		return
	}
	img2, err := downloadImage(url2)
	if err != nil {
		return
	}

	// 计算哈希值进行比较
	hash1 := calculateMD5(img1)
	hash2 := calculateMD5(img2)

	if hash1 == hash2 {
		fmt.Println("两个Favicon图片相同")
	} else {
		fmt.Println("两个Favicon图片不相同")
	}

	// 保存图片2
	// 保存图片数据到文件
	fileName := fmt.Sprintf("%s_%s.png", "google.com", "128")
	// savePath := filepath.Join("favicon", fileName)
	err = os.WriteFile(fileName, img2, 0644)
	if err != nil {
		fmt.Printf("保存图片失败: %v", err)
	}
}
