package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/unix755/xtools/xExec"
)

// GetName 获取编译后的静态文件名,customName为自定义名称(可为空)
func GetName(targetOS string, targetARCH string, customName string) (name string, err error) {
	// 如果使用自定义名称
	if customName != "" {
		name = fmt.Sprintf("%s-%s-%s", customName, targetOS, targetARCH)
	} else {
		// 不使用自定义名称,则获取模块名称
		packageName, err := GetModuleName()
		if err != nil {
			return "", err
		}
		name = fmt.Sprintf("%s-%s-%s", packageName, targetOS, targetARCH)
	}

	// windows 系统则在静态文件名中加入 .exe 后缀
	if targetOS == "windows" {
		name = name + ".exe"
	}
	return name, nil
}

// CleanCache 清除编译缓存
func CleanCache() (err error) {
	return xExec.Run(exec.Command("go", "clean", "-cache"))
}

// Build 编译
func Build(targetOS string, targetARCH string, name string, dir string, noDebug bool, noCgo bool, opts []string, envs []string) (err error) {
	// 编译的第一参数固定为 build
	args := []string{
		"build",
	}

	// 文件名
	name, err = GetName(targetOS, targetARCH, name)
	if err != nil {
		return err
	}

	// 文件输出路径
	if dir != "" {
		args = append(args, "-o", filepath.Join(dir, name))
	} else {
		args = append(args, "-o", name)
	}

	// debug 参数
	if noDebug {
		args = append(args, "-trimpath", "-ldflags", "-s -w")
	}

	// cgo 环境变量
	if noCgo {
		envs = append(envs, "CGO_ENABLED=0")
	}

	// 最终命令的参数及环境变量
	args = append(args, opts...)
	envs = append(envs, fmt.Sprintf("GOOS=%s", targetOS), fmt.Sprintf("GOARCH=%s", targetARCH))

	cmd := exec.Command("go", args...)
	cmd.Env = append(os.Environ(), envs...)
	return xExec.Run(cmd)
}
