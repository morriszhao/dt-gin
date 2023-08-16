package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"morris/im/helper"
	"net/http"
	"strconv"
	"time"
)

type chatServices struct {
	ctx *gin.Context
}

func NewChatServices(c *gin.Context) *chatServices {
	return &chatServices{
		ctx: c,
	}
}

func (cc *chatServices) Chat() error {
	//http协议 提升到  websocket 协议
	conn, err := (&websocket.Upgrader{

		//指定升级 websocket   握手完成的 超时时间
		HandshakeTimeout: time.Second * 1,
		//指定io操作的缓存大小 为0则不限制
		ReadBufferSize:  0,
		WriteBufferSize: 0,

		//http 错误响应函数
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			fmt.Println(reason.Error())
		},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(cc.ctx.Writer, cc.ctx.Request, nil)
	if err != nil {
		helper.RespFail(cc.ctx, helper.SystemError, err.Error())
		return err
	}

	//封装一层的目的 主要是为了   利用goroutine并发的能力、  异步写入、 提供业务goroutine的处理速度
	ws := websocketConn{
		conn: conn,
		data: make(chan []byte, 100),
	}

	//userId与websocket map
	userId := GetUserId(cc.ctx)
	addConn(userId, ws)

	//启动
	ws.Start()

	ws.sendMsg("当前用户id：" + strconv.Itoa(userId))
	return nil
}
