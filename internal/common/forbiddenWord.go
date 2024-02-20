package common

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func WordsFilter(ctx context.Context, data string) error {
	wd, _ := os.Getwd()
	wf := wordsfilter.New()
	root, _ := wf.GenerateWithFile(fmt.Sprintf("%s/doc/%s", wd, "prohibitedWords.txt"))
	cxt := strings.ReplaceAll(data, "\n", "")
	cxt = strings.ReplaceAll(cxt, " ", "")
	temporary := wf.Replace(cxt, root)
	tool.Info(ctx, fmt.Sprintf("添加文案素材:%v,替换展示:%v", cxt, temporary))
	c := []rune(cxt)
	d := []rune(temporary)
	res := ""
	if wf.Contains(cxt, root) || cxt != data {
		isB := 0
		for i, r := range d {
			if r != c[i] {
				if res != "" && isB == 0 {
					res = res + ","
				}
				res = res + string(c[i])
				isB = 1
			} else {
				isB = 0
			}
		}
		if res != "" {
			return errors.New(fmt.Sprintf("脚本含有违禁词【%s】,请进行修改", res))
		}
	}
	return nil
}
