package runner

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFiles(dir string) ([]string, error) {
	// 定义结果切片
	contents := []string{}

	// 遍历文件夹
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 跳过目录
		if info.IsDir() {
			return nil
		}
		//fmt.Println(path)
		// 读取文件内容
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// 将文件内容添加到结果切片
		contents = append(contents, string(content))

		return nil
	})
	if err != nil {
		return nil, err
	}

	return contents, nil
}
