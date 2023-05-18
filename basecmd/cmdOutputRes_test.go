package basecmd

import (
	"fmt"
	"log"
	"testing"
)

func Test_CmdAndChangeDirToResAllInOne(t *testing.T) {
	res, err := CmdAndChangeDirToResAllInOne("/Users/fuao/Downloads/", "ls -l")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}
