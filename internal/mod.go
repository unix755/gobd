package internal

import (
	"encoding/json"
	"net/url"
	"os/exec"
	"path"
)

type Mod struct {
	Module struct {
		Path string `json:"Path"`
	} `json:"Module"`
}

// GetModuleName 获取模块名称
func GetModuleName() (name string, err error) {
	var mod Mod

	// 使用 go mod edit -json 列出模块信息
	output, err := exec.Command("go", "mod", "edit", "-json").Output()
	if err != nil {
		return "", err
	}

	// json解码到mod
	err = json.Unmarshal(output, &mod)
	if err != nil {
		return "", err
	}

	// mod中提取出模块名称
	modURL, err := url.Parse(mod.Module.Path)
	if err != nil {
		return "", err
	}
	return path.Base(modURL.Path), nil
}
