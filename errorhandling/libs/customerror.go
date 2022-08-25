package libs

import (
	"errors"
	"fmt"
)

func CustomError() {
	var error1 = fmt.Errorf("커스텀 에러 발생합니다")
	fmt.Println(error1)
	var error2 = errors.New("커스텀 에러 두번째 발생합니다")
	fmt.Println(error2)
}
