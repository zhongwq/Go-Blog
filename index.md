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

若是前后端主机不一致，需要自行配置前端的proxyTable(用于跨域请求)，其在config/index.js里，可改为后端对应的target，默认配置为api请求转发到localhost:8081端口

![屏幕快照 2018-12-16 下午1.05.38](https://lh3.googleusercontent.com/-7m_tNl8Uk54/XBXdP-GR9mI/AAAAAAAAAJ0/XXc1WGSF3bEZo1KHaByTypHIiH6kbCTVQCHMYCw/I/%255BUNSET%255D)

##### 后端
后段部分直接通过执行以下go get命令进行安装

```bash
go get github.com/GoProjectGroupForEducation/Go-Blog
```

执行完成后，我们就可以通过`./Go-Blog`打开服务了

所用博客数据压缩到了data.zip压缩包，可以把其解压到原根目录，替换data.db以及static文件夹

### 项目博客来源

组内成员个人平时所写的博客

### API设计

##### article

- `GET /articles/?pageNum={pageNum}` 根据请求的页数pageNum获取当前页数的文章
- `GET /articles/user/{user_id}` 根据用户id获取该用户的文章
- `GET /articles/tag/{tag_content}` 根据tag获取含有该tag的文章
- `GET /articles/concerning` 获取已关注的文章
- `GET, PUT /articles/{article_id}` 更新或获取含有相应id的文章
- `GET, POST /articles/{article_id}/comments` 更新或获取含有相应id文章的评论
- `GET, PUT /articles/{article_id}/comments/{comment_id}` 更新或获取含有相应id文章的含有对应id的评论

##### user

- `GET /user/allusers` 获取所有用户
- `GET, PUT /user/{user_id}` 获取含有对应id的用户
- `GET /user/{user_id}/follower`  获取含有对应id的用户的follower
- `GET /user/{user_id}/following` 获取含有对应id的用户follow的用户
- `POST /user/login` 登陆请求
- `POST /user/register` 注册请求
- `POST /user/icon/{filename}` 更新用户头像
- `POST /user/follow` follow某用户
- `POST /user/unfollow` 取消follow某用户

##### tag

- `GET /tag` 获取所有已存在的tag

##### Example

- 用户登陆

![](./images/1.png)

- 新建文章

![](./images/2.png)

- 获取分页号为1的文章，data中还含有一个Number字段标识文章的总数量，Articles字段存放分页号为1的文章，一个分页有6篇文章

![](./images/5.png)

- 更新文章

![](./images/3.png)

- 获取id为1的文章

![](./images/4.png)



### 前端使用效果截图

> 前端使用的Font awesome库可能在部分网络环境(墙内)中会有异常，碰到这种情况可以挂个梯子，加载一次，之后就有缓存可以直接访问了

登陆页面
![屏幕快照 2018-12-16 下午1.03.00](https://lh3.googleusercontent.com/-AEKyqYxBfus/XBXclVf9XjI/AAAAAAAAAJg/9tKAHbDS_lYpVt3j6sCsnhNxs0j5rvOaACHMYCw/I/%255BUNSET%255D)


注册页面
![屏幕快照 2018-12-16 下午1.03.06](https://lh3.googleusercontent.com/-1kYzCsLtrMQ/XBXcmwqVOjI/AAAAAAAAAJo/NqchOXhtnM0HlWhUGV46M7h8mXtEDxZKwCHMYCw/I/%255BUNSET%255D)



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

用户还可以更改头像等的个人信息
![屏幕快照 2018-12-16 下午1.01.47](https://lh3.googleusercontent.com/-H8vVIZsLPtI/XBXcWCm3X0I/AAAAAAAAAJc/7OAd4R035v4tPCWJ0fK9TIBp8GdDR_kJgCHMYCw/I/%255BUNSET%255D)



