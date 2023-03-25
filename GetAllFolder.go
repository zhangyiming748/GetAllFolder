package GetAllFolder

import (
	"golang.org/x/exp/slog"
	"io"
	"os"
	"strings"
)

var (
	all []string
)

func init() {
	logLevel := os.Getenv("LEVEL")
	//var level slog.Level
	var opt slog.HandlerOptions
	switch logLevel {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "GetAllFolder.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer logf.Close() //如果不关闭可能造成内存泄露
	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(logger)
}
func ListFolders(dirname string) []string {
	fileInfos, _ := os.ReadDir(dirname)
	var folders []string
	for _, fi := range fileInfos {
		filename := strings.Join([]string{dirname, fi.Name()}, string(os.PathSeparator)) //拼写当前文件夹中所有的文件地址
		// log.Info.Println(filename)                                  //打印文件地址
		if fi.IsDir() { //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(fi.Name(), "h265") {
				continue
			}
			if strings.Contains(fi.Name(), "bigger") {
				continue
			}
			if strings.Contains(filename, "/.") {
				slog.Info("跳过隐藏文件夹", slog.Any("文件名", fi.Name()))
				continue
			}
			slog.Info("获取到文件夹", slog.Any("文件名", filename))
			all = append(all, filename)
			folders = append(folders, filename)
			ListFolders(filename) //递归调用
		}
	}
	return all
}
