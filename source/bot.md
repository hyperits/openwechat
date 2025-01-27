#  Bot对象

`Bot`对象负责处理网络请求和消息回调以及登录登出的用户行为，一个`Bot`对应一个登录的微信号。



### 创建Bot对象

在登录微信之前需要创建一个`Bot`对象。

```go
bot := openwechat.DefaultBot()
```

使用默认的构造方法`DefaultBot`来创建一个`Bot`对象。



### 登陆二维码回调

但仅仅是这样我们依然无法登录, 我们平常登录微信都需要用手机扫描二维码登录，所以我们得知道需要扫描哪张二维码，最后还需要为它绑定一个登录二维码的回调函数。

```go
// 注册登陆二维码回调
bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
```

`PrintlnQrcodeUrl`这个函数做的事情很简单，就是将我们需要扫码登录的二维码链接打印打控制台上，这样我们就知道去扫描哪张二维码登录了。

可以自定义`UUIDCallback`来实现自己的逻辑。

如：将登录的二维码打印到控制台。

```go
package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/hyperits/openwechat"
)

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func main() {
	bot := openwechat.DefaultBot()
	bot.UUIDCallback = ConsoleQrCode
	bot.Login()
}
```

虽然最终打印的结果肉眼看上去有点不尽人意，但手机也还能够识别...



### 登录



#### 普通登录

上面的准备工作做完了，下面就可以登录，直接调用`Bot.Login`即可。

```go
bot.Login() 
```

登录会返回一个`error`，即登录失败的原因。



#### 热登录

每次执行普通登录都需要扫码，在调试一些功能的时候需要反复编译，这样会很麻烦。

热登录可以只用扫码一次，后面在单位时间内重启程序也不会再要求扫码

```go
// 创建热存储容器对象
reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

// 执行热登录
bot.HotLogin(reloadStorage)
```

`HotLogin`需要接受一个`热存储容器对象`来调用。`热存储容器`用来保存登录的会话信息，本质是一个接口类型

```go
// 热登陆存储接口
type HotReloadStorage io.ReadWriteCloser
```

`NewJsonFileHotReloadStorage`简单实现了该接口，它采用`JSON`的方式存储会话信息。

实现这个接口，实现你自己的存储方式。



### 扫码回调

在pc端微信上我们打开手机扫码进行登录的时候，只扫描二维码，但不点击确认，微信上也能够显示当前扫码用户的头像，并提示用户登录确认。

通过对`bot`对象绑定扫码回调即可实现对应的功能。

```go
bot.ScanCallBack = func(body []byte) { fmt.Println(string(body)) }
```

用户扫码后，body里面会携带用户的头像信息。

**注**：绑定扫码回调须在登录前执行。



### 登录回调

对`bot`对象绑定登录

```go
bot.LoginCallBack = func(body []byte) {
		fmt.Println(string(body))
		// to do your business
}
```

登录回调的参数就是当前客户端需要跳转的链接，可以不用关心它。

登录回调函数可以当做一个信号处理，表示当前扫码登录的用户已经确认登录。



### 桌面模式

`DefaultBot`默认是与网页版微信进行交互，部分用户的网页版wx可能已经被限制登录了。

这时候可以尝试使用`桌面模式`进行登录。

```go
bot := openwechat.DefaultBot(openwechat.Desktop)
```

别的逻辑不用改，直接在创建bot的时候加一个参数就行了。

如果桌面模式还登录不上，请检查你的微信号是不是刚刚申请。



### 消息处理

在用户登录后需要实时接受微信发送过来的消息。

很简单，给`BOT`对象绑定一个消息回调函数就行了。

```go
// 注册消息处理函数
bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
}
```

所有接受的消息都通过`Bot.MessageHandler`来处理。

基于这个回调函数，可以对消息进行多样化处理

```go
dispatcher := openwechat.NewMessageMatchDispatcher()

// 只处理消息类型为文本类型的消息
dispatcher.OnText(func(msg *Message){
	msg.ReplyText("hello")
})

// 注册消息回调函数
bot.MessageHandler = openwechat.DispatchMessage(dispatcher)
```

`openwechat.DispatchMessage`会将消息转发给`dispatcher`对象处理



#### MessageMatchDispatcher

##### 构造方法

```go
openwechat.NewMessageMatchDispatcher()
```

##### 注册消息处理函数

```go
// 注册消息处理函数
func (m *MessageMatchDispatcher) RegisterHandler(matchFunc matchFunc, handlers ...MessageContextHandler)


// 消息匹配函数
type matchFunc func(*Message) bool


// 消息处理函数
type MessageContextHandler func(ctx *MessageContext)
```

`matchFunc`：接受当前收到的消息对象，并返回`bool`值，返回`true`则表示处理当前的消息

`RegisterHandler`：接受一个`matchFunc`和不定长的消息处理函数，如果`matchFunc`返回为`true`，则表示运行对应的处理函数组。



##### OnText

注册处理消息类型为文本类型的消息

```go
func (m *MessageMatchDispatcher) OnText(handlers ...MessageContextHandler)
```

##### OnImage

注册处理消息类型为图片类型的消息

```golang
func (m *MessageMatchDispatcher) OnImage(handlers ...MessageContextHandler)
```

##### OnVoice

注册处理消息类型为语言类型的消息

```go
func (m *MessageMatchDispatcher) OnVoice(handlers ...MessageContextHandler)
```

##### [更多请点击查看源码](https://github.com/hyperits/openwechat/blob/main/message_handle.go)



### 获取登录后的用户

```go
self, err := bot.GetCurrentUser()
```

**注**：该方法在登录成功后调用

[详见`Self`对象](./user.md)





### 阻塞主程序

```go
bot.Block()
```

该方法会一直阻塞，直到用户主动退出或者网络请求发生错误

