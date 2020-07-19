# StreamingVideo
Streaming video website

### 流媒体点播网站

##### 为什么选择视频网站
* Go是一门网络编程语言
* 视频网站包含Go在实战项目中的绝大部分技能要点
* 优良的native http库以及模板引擎(无需任何第三方框架)

##### 总体架构
* 前后端分离

##### 什么是前后端解耦
* 前后端解耦是时下流行的web网站架构
* 前端页面和服务通过普通的web引擎渲染
* 后端数据通过渲染后的页面脚本调用后处理和呈现

##### 前后端解耦的优势
* 解放生产力，提高合作效率
* 松耦合的架构更灵活，部署更方便，更符合微服务的设计特征
* 性能的提升，可靠性的提升

##### 前后端解耦的缺点
* 工作量大
* 前后端分离带来的团队成本以及学习成本
* 系统复杂度加大

### 后端服务 -- API

##### API
* REST  API
* REST是一种设计风格，不是任何架构标准
* 当今RESTful API通常使用HTTP作为通信协议，JSON作为数据格式

##### API
* 统一接口
* 无状态
* 可缓存
* 分层
* CS模式

##### API设计原则
* 以URL风格设计API
* 通过不同的METHOD(GET,POST,PUT,DELETE)来区分对资源的CRUD
* 返回码符合HTTP资源描述的规定

##### API设计总览

##### API设计：用户
* 创建(注册)用户: URL：/user  Method：POST , SC: 201,400,500
* 用户登录： URL: /user/:username Method: POST , SC:200,400,500
* 获取用户基本信息: URL: /user/:username Method:GET , SC:200,400,401,403,500
* 用户注销： URL: /user/:username Method:DELETE , SC:204,400,401,403,500 

##### API设计： 用户资源
* List all videos: URL:/user/:username/videos Method:GET,SC:200,400,500
* Get one video: URL:/user/:username/videos/:vid-id Method:GET,SC:200,400,500
* Delete one video : URL:/user/:username/videos/:vid-id 
	Method: DELTE , SC: 204,400,401,403,500

##### API设计： 评论
* Show commtents: URL: /videos/:vid-id/comments Method:GET , SC:200,400,500
* Post a comment: URL :/videos/:vid-id/comments Method:POST , SC:201,400,500
* Delete a comment: URL:/videos/:vid-id/comment/:commtent-id
	Method: DELETE , SC:204,400,401,403,500

##### 数据库设计： 用户
* table user

##### 数据库设计： 视频资源
* table videoinfo 

##### 数据库设计： 评论
* comments 

##### Session？
* session
	Session 是服务端使用的一种记录客户状态的机制
	**注:**
	1. Session保存在服务器端,为了获取更高的存取速度，服务器一般把Session放
		在内存里。Session内容应该尽量精简。
	2. 每个用户都有一个独立的Session。
	3. Session在用户第一次访问服务器的时候自动创建
	4. 为防止内存溢出，服务器会把长时间没有活跃的Session从内存删除，这个时
		间就是Session的超时时间。
	
	

* 为什么要用session
* session 和 cookie的区别

### 后端服务 -- Streaming

##### Streaming
* 静态视频，非RTMP
* 独立的服务，可独立部署
* 统一的api格式

##### Stream server
* Streaming
* Upload file

##### Limiter
* share channel instead of share memory 

### Scheduler 
* 什么是scheduler
* 为什么需要scheduler
* scheduler通常做什么

###### Scheduler包含什么
* RESTful的http server
* Timer
* 生产者/消费者模型下的task runner


