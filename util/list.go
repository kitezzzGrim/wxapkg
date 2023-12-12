package util

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func List(c *cli.Context) error {
	warn_.Println("[*] only support WeChat version 3.8.*")
	root := GetWXAppletPath()
	info_.Printf("list %s\n", root)

	stat, err := os.Stat(root)
	if err != nil { // 判断小程序缓存路径是否存在
		err_.Println("[-] not support WeChat version")
		return err
	}

	if !stat.IsDir() {
		err_.Println("[-] not support WeChat version")
		return nil
	}

	found_app := false // 标志变量，初始值为 false
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "__APP__") {
			fmt.Println(path)
			found_app = true // 将标志变量设置为 true

		}
		return nil
	})

	//fmt.Println(found_app)
	if err != nil {
		fmt.Println("遍历出错:", err)
	} else if !found_app {
		fmt.Println("不存在包含 '__APP__' 字符串的路径,确认客户端是否打开小程序过")
	}
	return nil
}
