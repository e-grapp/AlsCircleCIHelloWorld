package main

import (
	"fmt"
	"strings"

	"bitbucket.org/ai69/amoy"
	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
)

var (
	AppName    string
	CIBuildNum string
	BuildDate  string
	BuildHost  string
	GoVersion  string
	GitBranch  string
	GitCommit  string
	GitSummary string
)

func buildInfo() string {
	if ystring.IsBlank(BuildDate) {
		return amoy.EmptyStr
	}

	flag := "➣ "
	if yos.IsOnWindows() {
		flag = "> "
	}
	sb := strings.Builder{}
	addNonBlankField := func(name, value string) {
		if ystring.IsNotBlank(value) {
			fmt.Fprintln(&sb, flag+name+":", value)
		}
	}

	addNonBlankField("K-App Name", AppName)
	addNonBlankField(" Build Num", CIBuildNum)
	addNonBlankField("Build Date", BuildDate)
	addNonBlankField("Build Host", BuildHost)
	addNonBlankField("Go Version", GoVersion)
	addNonBlankField("Git Branch", GitBranch)
	addNonBlankField("Git Commit", GitCommit)
	addNonBlankField("GitSummary", GitSummary)
	return sb.String()
}
