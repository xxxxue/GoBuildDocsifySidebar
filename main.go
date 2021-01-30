package main

import (
	"encoding/json"
	"fmt"
	"GoBuildDocsifySidebar/utils"
	"io/ioutil"
	"path"
	"strings"
)

type Config struct {
	HomePath string
}

var (
	_config           *Config
	_homePath         string
	_jsonConfigPath   = "./Config/Config.json"
	_sidebarFileName  = "_sidebar.md"
	_readmeFileName   = "README.md"
	_ignoreDirList    = []string{".git"}
	_ignoreFileList   = []string{"_sidebar.md", "README.md"}
	_ignoreStringList = []string{".assets"}
	_includeDirList   = []string{}
	_level            = 0
)



func Entry(rootPath string, isHome bool) string {
	var sidebarData = ""
	dir := utils.NewFileUtils(rootPath)

	//{Utils.ReplaceSpace(rootDir.GetDirRelativePath())}
	// 生成目录文件夹
	sidebarData += fmt.Sprintf("%s- [%s](%s)\n",
		GenerateSpace(_level),                                  // 空格
		dir.GetLastLevelPath(),                                 // 文件夹名称
		strings.ReplaceAll(rootPath, _config.HomePath, "")+"/", // 文件夹相对路径
	)

	_level++
	if rootPath == _config.HomePath {
		sidebarData = ""
		_level = 0
	}

	itemList := utils.GetDirItemList(rootPath)

	for _, f := range *itemList {

		fullName := path.Join(rootPath, f.Name())

		fileObj := utils.NewFileUtils(fullName)

		if fileObj.IsFile() {
			isContains := false
			for _, item := range _ignoreFileList {
				if item == f.Name() {
					isContains = true
					break
				}
			}

			if fileObj.GetExtension() == ".md" && !isContains {
				sidebarData += fmt.Sprintf("%s- [%s](%s)\n",
					GenerateSpace(_level),
					strings.ReplaceAll(f.Name(), fileObj.GetExtension(), ""), //file.GetFileNameWithoutExtension()
					strings.ReplaceAll(fullName, _config.HomePath, ""),       //Utils.ReplaceSpace(file.GetFileRelativePath())}
				)
			}
		} else {

			isContains := false
			for _, item := range _ignoreDirList {
				if item == f.Name() {
					isContains = true
					break
				}
			}

			isContainsString := false
			for _, item := range _ignoreStringList {
				if strings.Contains(f.Name(), item) {
					isContainsString = true
					break
				}
			}

			if !isContains && !isContainsString {

				if isHome {
					_includeDirList = append(_includeDirList, fullName)
				}

				sidebarData += Entry(fullName, isHome)
				println(f.Name() + " Done!")
				_level--
			}
		}
	}

	return sidebarData
}

func Build() {
	homeData := Entry(_homePath, true)
	WriteDataToFile(_homePath, homeData)
	println("[home] Done!!")

	for _, item := range _includeDirList {
		_level = 0
		includeData := Entry(item, false)

		parentDir := utils.NewFileUtils(item).GetExceptLastLevelPath()

		if parentDir == _homePath {
			includeData = "- [返回首页](/)\n" + includeData
		} else {
			includeData = fmt.Sprintf("- [返回上一级 [%s]](%s)\n",
				utils.NewFileUtils(parentDir).GetLastLevelPath(),
				strings.ReplaceAll(parentDir, _homePath, ""), //{parentDir.GetDirRelativePath()}
			) + includeData
		}
		WriteDataToFile(item, includeData)
	}
}

// 写入文件
func WriteDataToFile(homePath, data string) {
	sidebarPath := path.Join(homePath, _sidebarFileName)
	readmePath := path.Join(homePath, _readmeFileName)

	utils.NewFileUtils(sidebarPath).WriteData(data)
	utils.NewFileUtils(readmePath).WriteData(data)

}

// 初始化 配置
func InitConfig() {
	data, _ := ioutil.ReadFile(_jsonConfigPath)
	err := json.Unmarshal(data, &_config)
	if err != nil {
		println(err.Error())

	}
	// 对路径中的反斜杠进行处理
	_config.HomePath = strings.ReplaceAll(_config.HomePath, "\\", "/")

	_homePath = _config.HomePath
	//println(fmt.Sprintf("%+v", _config))
}


func Run() {
	Try(func() {
		InitConfig()
		Build()
	}, func(e interface{}) {
		println(e)
		println()
		println("  遇到错误!!!")

	})
	utils.Console{}.ReadLen_string("按回车键退出...")
}

//实现 try catch 例子
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}


func main() {
	Run()
}