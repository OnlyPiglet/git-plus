/*
Copyright © 2024 OnlyPiglet <jackwuchenghao4@gmail.com>
*/

package main

import (
	"fmt"
	"github.com/OnlyPiglet/git-plus/pkg/ecmd"
	"github.com/OnlyPiglet/git-plus/pkg/user"
	giturl "github.com/kubescape/go-git-url"
	"os"
)

func main() {

	err := ecmd.Exec("git", "version")

	if err != nil {
		os.Stderr.WriteString("dependency missing:\n\tplease install git first, https://git-scm.com/")
		return
	}

	os.Stdout.WriteString("^ v ^ FreeLife ^ v ^\n\n")
	args := os.Args
	lens := len(args)
	if lens <= 1 {
		return
	}
	scope := args[1]
	switch scope {
	case "adduser":
		if lens <= 4 {
			os.Stderr.WriteString("params invalid:\n\texample: git-plus adduser github.com foo foo@example.com")
			return
		}
		host := args[2]
		name := args[3]
		email := args[4]
		err := user.AddUser(host, name, email)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("adduser failed,reason is %s\n", err.Error()))
			return
		}
	case "deluser":
		if lens <= 3 {
			os.Stderr.WriteString("params invalid:\n\texample: git-plus deluser github.com foo")
			return
		}
		host := args[2]
		name := args[3]
		err := user.DelUser(host, name)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("adduser failed,reason is %s\n", err.Error()))
			return
		}
	case "listuser":
		if lens <= 2 {
			os.Stderr.WriteString("params invalid:\n\texample: git-plus listuser github.com")
			return
		}
		host := args[2]
		err := user.ListUser(host)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("listuser failed,reason is %s\n", err.Error()))
			return
		}
	case "clone":
		url, err := giturl.NewGitURL(args[2])
		if err != nil {
			args = args[1:]
			if err := ecmd.Exec("git", args...); err != nil {
				//panic(err)
				return
			}
		}
		phost := url.GetHostName()
		repoName := url.GetRepoName()
		u, err := user.GetUser(phost)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("clone %s failed,reason is %s\n", repoName, err.Error()))
		}
		args = args[1:]
		if err := ecmd.Exec("git", args...); err != nil {
			return
		}
		os.Chdir(repoName)
		if err := ecmd.Exec("git", "config", "--local", "user.name", u.Name); err != nil {
			return
			//panic(err)
		}
		if err := ecmd.Exec("git", "config", "--local", "user.email", u.Email); err != nil {
			return
			//panic(err)
		}
	default:
		args = args[1:]
		if err := ecmd.Exec("git", args...); err != nil {
			return
			//panic(err)
		}
	}
}
