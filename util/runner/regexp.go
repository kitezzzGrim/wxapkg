package runner

import (
	"regexp"
)

func CompileRegexes(regexList []string) ([]*regexp.Regexp, error) { // 将字符串列表中的正则表达式编译成一系列的 *regexp.Regexp 对象
	var compiledRegexes []*regexp.Regexp

	for _, r := range regexList {
		re, err := regexp.Compile(r)
		if err != nil {
			return nil, err
		}
		compiledRegexes = append(compiledRegexes, re)

	}
	//fmt.Println(compiledRegexes)
	return compiledRegexes, nil
}
