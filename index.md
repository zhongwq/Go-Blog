---
layout: default
---

# 项目文档

### 仓库地址
[前端](https://github.com/GoProjectGroupForEducation/MinimalBlog-Vue)
[后端](https://github.com/GoProjectGroupForEducation/Go-Blog)

### 安装指南
##### 前端

前端部分需要先把项目先clone下来, 执行命令如下

```bash
git clone https://github.com/GoProjectGroupForEducation/MinimalBlog-Vue.git
cd MinimalBlog-Vue
npm install
npm run dev
```

##### 后端
后段部分直接通过执行以下go get命令进行安装

```bash
go get github.com/GoProjectGroupForEducation/Go-Blog
```

执行完成后，我们就可以通过`./Go-Blog`打开服务了


### 项目博客来源

组内成员个人平时所写的博客

### API设计



### 前端使用效果截图

未登陆，主页（可以通过点击标签查看当前标签下的所有文章，博客列表可以点击自己写的博客才有的list（右下角3个点来点击Edit进行编辑））

![屏幕快照 2018-12-16 上午8.33.35](https://lh3.googleusercontent.com/-4sjTduxwQG0/XBWekE47hkI/AAAAAAAAAIc/G48UQHfIR-wHq7NqxUwQdaUB70qG3COpQCHMYCw/I/%255BUNSET%255D)


登陆后，主页

![屏幕快照 2018-12-16 上午8.38.58](https://lh3.googleusercontent.com/-xSRmfihJeok/XBWesmo5afI/AAAAAAAAAIk/uFWgzFjYU_gvx3vFpsOCqGfvf1Dy5ppVwCHMYCw/I/%255BUNSET%255D)

AddPost页面(拖入图片即可自动发送图片到后端，传回图片静态文件地址，自动以md格式添加到拖入光标处)

![屏幕快照 2018-12-15 下午10.32.08](https://lh3.googleusercontent.com/-CupAiaWyYBQ/XBWe4VApyFI/AAAAAAAAAIo/pHg_jpSTSYkx1zRwlhP1NiQMwPrS1BIlgCHMYCw/I/%255BUNSET%255D)

用户个人主页

![localhost_8080_ -3-](https://lh3.googleusercontent.com/-Xz4--W5QDmg/XBWfTZ_9inI/AAAAAAAAAI0/fhOS6EwqWwU2ZGNs_qWkmwa5JZs-0TSiwCHMYCw/I/localhost_8080_%2B%25283%2529.png)

博客页面，可以通过点击标签查看当前标签下的所有文章,登陆后可以进行评论

未登陆
![localhost_8080_ -4-](https://lh3.googleusercontent.com/-p8tVLsW1w-k/XBWgFYmSRzI/AAAAAAAAAJI/Xia0AyAwFNY7-VpBxfmAj5ftxgwSf622wCHMYCw/I/localhost_8080_%2B%25284%2529.png)


登陆后

![localhost_8080_ -2-](https://lh3.googleusercontent.com/-YlcnxakPeTc/XBWfbufp-8I/AAAAAAAAAI4/DeNnqYI8lAIGgUN1Wmcij_xktqSL74blQCHMYCw/I/localhost_8080_%2B%25282%2529.png)

评论也可以编辑，和博客一样点击3个点的列表点击Edit就可编辑，不过Comment较为简单，所以只弹出一个弹窗提供更改

![屏幕快照 2018-12-16 上午8.47.03](https://lh3.googleusercontent.com/-5TDr8fshcZo/XBWgrO2tG-I/AAAAAAAAAJQ/ywREL_-ke5o5XaypORpwiYdMt72mIBsAACHMYCw/I/%255BUNSET%255D)


