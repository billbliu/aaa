package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 下载指定domain和size的favicon图片
func downloadFavicon(domain string, size string, savePath string) error {
	// 构造FaviconKit API URL
	url := fmt.Sprintf("https://api.faviconkit.com/%s/%s", domain, size)

	// 发起GET请求获取图片
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("无法获取favicon: %v", err)
	}
	defer response.Body.Close()

	// 检查请求是否成功
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("获取favicon失败，状态码: %d", response.StatusCode)
	}

	// 创建本地文件用于保存图片
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("无法创建文件: %v", err)
	}
	defer file.Close()

	// 将下载的内容写入本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	fmt.Println("Favicon已成功下载到:", savePath)
	return nil
}

func main() {
	// 需要下载的domain和size
	domain := "google.com"
	size := "128" // 尺寸可以根据需求调整

	// 保存路径
	savePath := filepath.Join("favicon", fmt.Sprintf("%s_%s.png", domain, size))

	// 确保保存目录存在
	os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

	// 下载并保存favicon
	err := downloadFavicon(domain, size, savePath)
	if err != nil {
		fmt.Println("下载失败:", err)
	}
}
