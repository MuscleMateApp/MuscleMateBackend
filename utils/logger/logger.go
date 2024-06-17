package logger

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

func LogError(err error) {
	fmt.Println(color.Red("[Error]"), err.Error())
}

func LogInfo(info string) {
	fmt.Println(color.Cyan("[Info]"), info)
}

func LogSuccess(success string) {
	fmt.Println(color.Green("[Success]"), success)
}

func LogWarning(warning string) {
	fmt.Println(color.Yellow("[Warning]"), warning)
}

func LogDebug(debug string) {
	fmt.Println(color.Blue("[Debug]"), debug)
}

func LogFatal(fatal string) {
	fmt.Println(color.Red("[Fatal]"), fatal)
}

func LogPanic(panic string) {
	fmt.Println(color.Red("[Panic]"), panic)
}
