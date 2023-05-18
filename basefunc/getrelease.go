package basefunc

import (
	"get-helm-yaml/basecmd"
	"log"
	"strings"
)

type ReleaseInfo struct {
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	TimeStamp  string `json:"timestamp"`
	ConfigYaml string `json:"configyaml"`
}

// GetALLReleaseALL 获取所有release的所有配置 []ReleaseInfo
func GetALLReleaseALL(input []ReleaseInfo) ([]ReleaseInfo, error) {
	var res []ReleaseInfo

	for k, v := range input {
		//获取数组结果
		configTmp, err := GetReleaseConfig(v)
		if err != nil {
			log.Println(err)
			return res, err
		}
		//生成字符串配置
		var configStr string
		for _, value := range configTmp {
			configStr = configStr + "\n" + value
		}
		//赋值
		var resTmp ReleaseInfo
		resTmp.Name = input[k].Name
		resTmp.NameSpaces = input[k].NameSpaces
		resTmp.TimeStamp = input[k].TimeStamp
		resTmp.ConfigYaml = configStr
		res = append(res, resTmp)
	}

	return res, nil
}

// GetALLReleaseALLWithOutValue 获取所有release的所有配置 []ReleaseInfo
func GetALLReleaseALLWithOutValue(input []ReleaseInfo) ([]ReleaseInfo, error) {
	var res []ReleaseInfo

	for k, v := range input {
		//获取数组结果
		configTmp, err := GetReleaseConfigWithOutValue(v)
		if err != nil {
			log.Println(err)
			return res, err
		}
		//生成字符串配置
		var configStr string
		for _, value := range configTmp {
			configStr = configStr + "\n" + value
		}
		//赋值
		var resTmp ReleaseInfo
		resTmp.Name = input[k].Name
		resTmp.NameSpaces = input[k].NameSpaces
		resTmp.TimeStamp = input[k].TimeStamp
		resTmp.ConfigYaml = configStr
		res = append(res, resTmp)
	}

	return res, nil
}

// GetReleaseConfigWithOutValue 获取对应release的所有配置
func GetReleaseConfigWithOutValue(input ReleaseInfo) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne("./", "helm get all "+input.Name+" -n "+input.NameSpaces)
	if err != nil {
		log.Println(err)
		log.Println("Check kubectl&helm ")
		return resTmp, err
	}
	for k, v := range resTmp {
		if v == "---" {
			resTmp = resTmp[k+1:]
			break
		}
	}
	return resTmp, nil
}

// GetReleaseConfig 获取对应release的所有配置
func GetReleaseConfig(input ReleaseInfo) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne("./", "helm get all "+input.Name+" -n "+input.NameSpaces)
	if err != nil {
		log.Println(err)
		log.Println("Check kubectl&helm ")
		return resTmp, err
	}
	return resTmp, nil
}

// GetALLRelease 获取所有的release-name ns 修改时间
func GetALLRelease() ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne("./", "helm ls --all-namespaces | awk 'NR>1 {print $1, $2, $4} ' ")
	if err != nil {
		log.Println(err)
		log.Println("Check kubectl&helm ")
		return resTmp, err
	}

	return resTmp, nil
}

// GetReleaseList 返回所有relaes的结构体数组
func GetReleaseList(input []string) (res []ReleaseInfo) {
	if len(input) == 0 {
		return res
	}
	for _, v := range input {
		arr := strings.Fields(v)
		var tmp ReleaseInfo
		tmp.Name = arr[0]
		tmp.NameSpaces = arr[1]
		tmp.TimeStamp = arr[2]

		res = append(res, tmp)
	}
	return res
}
