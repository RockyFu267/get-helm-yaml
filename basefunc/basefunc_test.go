package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_GetALLRelease(t *testing.T) {
	res, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res[0])

}

func Test_getReleaseList(t *testing.T) {
	resTmp, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	res := GetReleaseList(resTmp)

	fmt.Println(res)
	fmt.Println(res[0].Name)
	fmt.Println(res[1].NameSpaces)
	fmt.Println(res[2].TimeStamp)
}

func Test_GetReleaseConfig(t *testing.T) {
	resTmp01, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	resTmp02 := GetReleaseList(resTmp01)
	res, err := GetReleaseConfig(resTmp02[0])
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(res)
	var str string
	for _, v := range res {
		str = str + "\n" + v
	}
	fmt.Println(str)

}

func Test_GetReleaseConfigWithOutValue(t *testing.T) {
	resTmp01, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	resTmp02 := GetReleaseList(resTmp01)
	res, err := GetReleaseConfigWithOutValue(resTmp02[0])
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(res)
	var str string
	for _, v := range res {
		str = str + "\n" + v
	}
	fmt.Println(str)

}
func Test_GetALLReleaseALL(t *testing.T) {
	resTmp01, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	resTmp02 := GetReleaseList(resTmp01)
	res, err := GetALLReleaseALL(resTmp02)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res[0].Name)

}

func Test_WriteFile(t *testing.T) {
	resTmp01, err := GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	resTmp02 := GetReleaseList(resTmp01)
	resTmp03, err := GetALLReleaseALL(resTmp02)
	if err != nil {
		log.Println(err)
		return
	}

	err = WriteFile(resTmp03[0], "/Users/fuao/Desktop/code/github/gethelmyaml")
	if err != nil {
		log.Println(err)
		return
	}
}
