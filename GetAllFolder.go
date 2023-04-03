package GetAllFolder

import (
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

var (
	all []string
)

func ListFolders(dirname string) []string {
	fileInfos, _ := os.ReadDir(dirname)
	var folders []string
	for _, fi := range fileInfos {
		filename := strings.Join([]string{dirname, fi.Name()}, string(os.PathSeparator)) //拼写当前文件夹中所有的文件地址
		// log.Info.Println(filename)                                  //打印文件地址
		if fi.IsDir() { //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(filename, "/.") {
				slog.Debug("跳过隐藏文件夹", slog.Any("文件名", fi.Name()))
				continue
			}
			slog.Info("获取到文件夹", slog.String("文件名", filename))
			all = append(all, filename)
			folders = append(folders, filename)
			ListFolders(filename) //递归调用
		}
	}
	return all
}
