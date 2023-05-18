package basefunc

import (
	"get-helm-yaml/basecmd"
	"io"
	"log"
	"os"
)

// WriteFile 将结果写入新文件
func WriteFile(inputRes ReleaseInfo, path string) error {
	chartYamlFile := path + "/" + inputRes.Name + "---" + inputRes.NameSpaces + "-" + inputRes.TimeStamp + ".yaml"

	//文件是否存在
	checkDir, err := CheckDir(chartYamlFile)
	if err != nil {
		log.Println("check dir: ", err)
		//***结束***
		return err
	}
	var file *os.File
	if checkDir {
		log.Println("check dir&file: ", err)
		//***结束***
		return err
	} else {
		//不存在 创建文件
		file, err = os.Create(chartYamlFile)
		if err != nil {
			log.Println("创建失败", err)
			return err
		}
	}
	defer file.Close()
	//写入文件
	_, err = io.WriteString(file, inputRes.ConfigYaml)
	if err != nil {
		log.Println("写入错误：", err)
		return err
	}
	log.Println("写入成功")

	return nil
}

// CreateDir path是绝对路径
func CreateDir(input string) error {
	_, err := basecmd.CmdAndChangeDirToResAllInOne("./", "mkdir -p "+input)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CheckDir 检查文件或者目录是否存在
func CheckDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
