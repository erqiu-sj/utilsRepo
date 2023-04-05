package utils

import (
	"encoding/json"
	"io/ioutil"
)

type ReadPackageJson struct {
	Path string // package.json file path
}

func ReadPackageJsonHelper[T any](opt ReadPackageJson, structs T) T {
	curdFileUtils := CurdFile{}
	file := curdFileUtils.OpenFile(opt.Path, ErrorLogOpt{ErrorType: ErrorTypeConstant.Error, Msg: "package.json file is not open"})
	byteFile, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(byteFile, &structs)
	return structs
}
