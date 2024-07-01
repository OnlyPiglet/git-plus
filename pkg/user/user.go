package user

import (
	"errors"
	yaml "gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

var configPath = ""

func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath = filepath.Join(dir, ".git-plus-config")
	if _, err = os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err := os.Create(configPath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

type HostConfig struct {
	Host string `json:"host" yaml:"host"`
	User []*User
}
type User struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

func parseHostConfig() ([]HostConfig, error) {
	fd, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	body, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	var cf []HostConfig
	err = yaml.Unmarshal(body, &cf)
	if err != nil {
		return nil, err
	}
	return cf, nil
}

func flushHostConfig(body []byte) error {
	wfd, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer wfd.Close()
	_, err = wfd.Write(body)
	return err
}

func AddUser(host, name, email string) error {
	cf, err := parseHostConfig()
	if err != nil {
		return err
	}
	cf = addUser(cf, host, name, email)
	abody, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}
	err = flushHostConfig(abody)
	if err != nil {
		return err
	}
	return err
}

func addUser(cf []HostConfig, host, name, email string) []HostConfig {
	for i, hc := range cf {
		if hc.Host == host {
			for _, user := range hc.User {
				if user.Name == name {
					user.Email = email
					return cf
				}
			}
			cf[i].User = append(hc.User, &User{Name: name, Email: email})
			return cf
		}
	}

	if cf == nil {
		return []HostConfig{
			{Host: host, User: []*User{{Name: name, Email: email}}},
		}
	}
	return append(cf, HostConfig{Host: host, User: []*User{{Name: name, Email: email}}})
}

func DelUser(host, name string) error {
	cf, err := parseHostConfig()
	if err != nil {
		return err
	}
	cf = delUser(cf, host, name)
	abody, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}
	err = flushHostConfig(abody)
	if err != nil {
		return err
	}
	return err
}

func delUser(cf []HostConfig, host, name string) []HostConfig {
	for i, hc := range cf {
		if hc.Host == host {
			for j, user := range hc.User {
				if user.Name == name {
					cf[i].User = append(hc.User[:j], hc.User[j+1:]...)
				}
			}
		}
	}
	return cf
}

func ListUser(host string) error {
	cf, err := parseHostConfig()
	if err != nil {
		return err
	}
	for _, config := range cf {
		if config.Host == host {
			body, err := yaml.Marshal(config)
			if err != nil {
				return err
			}
			os.Stdout.Write(body)
			return nil
		}
	}
	return nil
}

func GetUser(host string) (*User, error) {
	cf, err := parseHostConfig()
	if err != nil {
		return nil, err
	}
	for _, config := range cf {
		if config.Host == host {
			if config.User == nil || len(config.User) == 0 {
				return nil, errors.New("no user of " + host)
			} else {
				return config.User[0], nil
			}
		}
	}
	return nil, errors.New("no user of " + host)
}
