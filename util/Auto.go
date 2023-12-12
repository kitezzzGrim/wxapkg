package util

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func List_1() ([]string, error) {
	warn_.Println("[*] only support WeChat version 3.8.*")
	root := GetWXAppletPath()
	info_.Printf("Mac本机当前微信小程序缓存保存路径为: %s\n", root)

	stat, err := os.Stat(root)
	if err != nil { // 判断小程序缓存路径是否存在
		err_.Println("[-] not support WeChat version")
		return nil, err
	}

	if !stat.IsDir() {
		err_.Println("[-] not support WeChat version")
		return nil, nil
	}

	foundPaths := []string{} // 保存包含 "__APP__" 字符串的路径的切片
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "__APP__") {
			fmt.Printf("发现小程序主包地址: %s\n", path)
			foundPaths = append(foundPaths, path) // 将包含 "__APP__" 字符串的路径添加到切片中
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历出错:", err)
	} else if len(foundPaths) == 0 {
		fmt.Println("不存在包含 '__APP__' 字符串的路径,确认客户端是否打开小程序过")
	}

	return foundPaths, nil
}

func Auto(c *cli.Context) error {
	foundPaths, _ := List_1()

	if len(foundPaths) == 0 {
		return nil
	}
	fmt.Println("正在对该小程序主包进行解包、JS格式化...")
	//unpackArgs := []string{os.Args[0], "unpack", "--in", foundPaths[0], "--format", "-v"}

	unpackArgs := []string{os.Args[0], "unpack", "--in", "", "--format", "-v", "-out", ""}
	for i, foundPath := range foundPaths {
		unpackArgs[3] = foundPath
		unpackArgs[7] = fmt.Sprintf("out_%d", i+1)

		// 执行解包命令
		unpackCmd := exec.Command(unpackArgs[0], unpackArgs[1:]...)
		//unpackCmd.Stdout = os.Stdout
		//unpackCmd.Stderr = os.Stderr

		err := unpackCmd.Run()
		if err != nil {
			fmt.Println("解包命令执行失败:", err)
			return err
		} else {
			fmt.Printf("第%d个小程序解包完成\n", i+1)
		}
	}

	// 执行解包命令
	//unpackCmd := exec.Command(unpackArgs[0], unpackArgs[1:]...)
	//unpackCmd.Stdout = os.Stdout
	//unpackCmd.Stderr = os.Stderr
	//
	//err := unpackCmd.Run()
	//if err != nil {
	//	fmt.Println("解包命令执行失败:", err)
	//	return err
	//} else {
	//	fmt.Println("解包命令执行完成")
	//	fmt.Println("")
	//}

	//fmt.Println("按回车键确认是否进行清理小程序路径缓存文件...")
	//fmt.Scanln()
	//
	//fmt.Println("正在执行clean命令...")
	//cleanArgs := []string{os.Args[0], "clean"}
	//
	//// 执行clean命令
	//cleanCmd := exec.Command(cleanArgs[0], cleanArgs[1:]...)
	//cleanCmd.Stdout = os.Stdout
	//cleanCmd.Stderr = os.Stderr
	//
	//err = cleanCmd.Run()
	//if err != nil {
	//	fmt.Println("clean命令执行失败:", err)
	//	return err
	//} else {
	//	fmt.Println("clean命令执行完成")
	//	fmt.Println("")
	//}

	return nil
}

//执行unpack命令 具体为 unpack --in foundPaths[0] --format -v
