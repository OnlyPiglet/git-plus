# git-plus

## 简介

如果你也有Github和Gitlab账号切换困扰,那么git-plus可以解决这个问题

### 用法

#### 新增用户

git-plus adduser github.com foo foo@example.com

git-plus adduser gitlab.com bar bar@example.com

#### 删除用户

git-plus deluser gitlab.com bar

#### 查看用户

git-plus listuser gitlab.com

#### clone 项目

git-plus clone git@github.com:OnlyPiglet/watchdog.git

git-plus clone https://github.com/OnlyPiglet/watchdog.git


### 安装

#### go 安装
go install github.com:OnlyPiglet/git-plus@latest