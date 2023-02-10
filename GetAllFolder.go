package GetAllFolder

import (
	"github.com/zhangyiming748/log"
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
		filename := strings.Join([]string{dirname, fi.Name()}, "/") //拼写当前文件夹中所有的文件地址
		// log.Info.Println(filename)                                  //打印文件地址
		if fi.IsDir() { //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(fi.Name(), "h265") {
				continue
			}
			if strings.Contains(filename, "/.") {
				log.Info.Printf("跳过隐藏文件夹:%v\n", fi.Name())
				continue
			}
			log.Info.Printf("获取到的文件夹:%v\n", filename)
			all = append(all, filename)
			folders = append(folders, filename)
			ListFolders(filename) //递归调用
		}
	}
	return all
}
