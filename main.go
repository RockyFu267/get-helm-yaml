package main

import (
	"get-helm-yaml/basefunc"
	"log"
	"os"
)

func main() {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println("Get pwdPATH ERROR: ", err)
		return
	}
	pwdPath = pwdPath + "/Charts"
	err = basefunc.CreateDir(pwdPath)
	if err != nil {
		log.Println(err)
		return
	}

	resTmp, err := basefunc.GetALLRelease()
	if err != nil {
		log.Println(err)
		return
	}

	resList := basefunc.GetReleaseList(resTmp)

	resAll, err := basefunc.GetALLReleaseALL(resList)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range resAll {
		err := basefunc.WriteFile(v, pwdPath)
		if err != nil {
			log.Println(err, v.Name)
			continue
		}
	}

}
