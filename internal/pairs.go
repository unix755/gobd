package internal

import (
	"bufio"
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

type Pair struct {
	OS   string
	ARCH string
}

func GetAllPairs() (pairs []Pair, err error) {
	// go tool dist list 获取所有支持的系统/架构对
	bytes, err := exec.Command("go", "tool", "dist", "list").Output()
	if err != nil {
		return nil, err
	}
	// 按行读取支持的所有所有系统/架构对
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		// 每一行按/分割提取系统与架构
		pair := strings.Split(scanner.Text(), "/")
		if len(pair) != 2 {
			return nil, errors.New("invalid pair format")
		}
		pairs = append(pairs, Pair{OS: pair[0], ARCH: pair[1]})
	}
	// 检查按行读取是否出错
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return pairs, nil
}

func GetMainPairs() (p []Pair) {
	var goos = []string{"darwin", "freebsd", "linux", "windows"}
	var goarch = []string{"amd64", "arm64", "386", "arm"}
	allPairs, err := GetAllPairs()
	if err != nil {
		return nil
	}

	for _, pair := range allPairs {
		if in(goos, pair.OS) && in(goarch, pair.ARCH) {
			p = append(p, pair)
		}
	}
	return p
}

func GetSelectedPairs(OS string, ARCH string) (p []Pair) {
	allPairs, err := GetAllPairs()
	if err != nil {
		return nil
	}

	if OS != "" && ARCH != "" {
		for _, pair := range allPairs {
			if in([]string{OS}, pair.OS) && in([]string{ARCH}, pair.ARCH) {
				p = append(p, pair)
			}
		}
		return p
	}

	if OS == "" && ARCH != "" {
		for _, pair := range allPairs {
			if in([]string{ARCH}, pair.ARCH) {
				p = append(p, pair)
			}
		}
		return p
	}

	if OS != "" && ARCH == "" {
		for _, pair := range allPairs {
			if in([]string{OS}, pair.OS) {
				p = append(p, pair)
			}
		}
		return p
	}

	// OS == "" && ARCH == ""
	for _, pair := range allPairs {
		if in([]string{runtime.GOOS}, pair.OS) && in([]string{runtime.GOARCH}, pair.ARCH) {
			p = append(p, pair)
		}
	}
	return p
}

func in(ss []string, t string) bool {
	for _, s := range ss {
		if t == s {
			return true
		}
	}
	return false
}
