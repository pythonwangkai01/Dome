package handlers

import "errors"

func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService Error :" + err.Error())
		//todo
		// log.Info(err)
		panic(err)
	}
}
