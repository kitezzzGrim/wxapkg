package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	err_  = color.New(color.FgRed)
	warn_ = color.New(color.FgYellow)
	info_ = color.New(color.FgCyan)
)

const (
	// wechat 3.8.* version
	wxappletPath = "/Users/%s/Library/Containers/com.tencent.xinWeChat/Data/.wxapplet/packages/"
)

func GetWXAppletPath() string {
	return fmt.Sprintf(wxappletPath, os.Getenv("USER"))
}
