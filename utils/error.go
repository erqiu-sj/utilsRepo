/*
 * @Author: 邱狮杰
 * @Date: 2023-01-23 00:07:10
 * @LastEditTime: 2023-01-23 00:36:57
 * @Description:
 * @FilePath: /memoRepo/utils/error.go
 */

package utils

type ErrorHooks interface {
	ReportImmediately() string // report now
}

type errorTypeMap struct {
	Warn  string
	Info  string
	Error string
}

var (
	ErrorTypeConstant = errorTypeMap{Error: "error", Warn: "warn", Info: "info"}
)

type ErrorLogOpt struct {
	ErrorType string // "warn" | "error" | "info"  error level
	Msg       string // error message
}

func ErrorLog(ops ErrorLogOpt) {
	// verify error level
	switch ops.ErrorType {
	case ErrorTypeConstant.Error:
		panic(ops.Msg)
	case ErrorTypeConstant.Warn:
		break
	case ErrorTypeConstant.Info:
		break
	}
}
