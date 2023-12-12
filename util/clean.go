package util

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Clean(c *cli.Context) error {
	// 构建命令

	dir := GetWXAppletPath() // 当前目录
	//fmt.Printf(dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("读取目录 %s 出错：%v\n", dir, err)
		//return
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "wx") {
			err := os.Chdir(dir) // 切换当前工作目录到要删除的文件夹所在的目录
			if err != nil {
				fmt.Printf("切换到目录 %s 出错：%v\n", dir, err)
				continue
			}

			err = os.RemoveAll(entry.Name())
			if err != nil {
				fmt.Printf("删除目录 %s 出错：%v\n", entry.Name(), err)
			} else {
				fmt.Printf("已删除目录：%s\n", entry.Name())
			}

			err = os.Chdir("..") // 切换回上级目录
			if err != nil {
				fmt.Printf("切换回上级目录出错：%v\n", err)
			}
		}
	}

	fmt.Println("完成")
	return nil
}
