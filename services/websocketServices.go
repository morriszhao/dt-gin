package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

//userId 与 websocket map、发送数据
var connMap map[int]websocketConn
var locker sync.RWMutex

func addConn(userId int, conn websocketConn) {
	if connMap == nil {
		connMap = make(map[int]websocketConn)
	}

	locker.Lock()
	defer locker.Unlock()
	connMap[userId] = conn
}
func delConn(userId int) {
	locker.Lock()
	defer locker.Unlock()
	delete(connMap, userId)
}

//websocket 链接封装
type websocketConn struct {
	conn *websocket.Conn
	data chan []byte
}

func (ws *websocketConn) Start() {
	go ws.sendMsgToClient()
	go ws.receiveMsgFromClient()
}

func (ws *websocketConn) sendMsg(message string) {
	ws.data <- []byte(message)
}

func (ws *websocketConn) sendMsgToClient() {
	for {
		select {
		case msg := <-ws.data:
			err := ws.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
func (ws *websocketConn) receiveMsgFromClient() {
	for {
		_, data, err := ws.conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println("收到客户端消息：", string(data))
		msg := websocketMessage{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			ws.sendMsg("发送成功， 消息不合法")
			continue
		}

		userId := msg.UserId
		dstId := msg.DstId

		sendConn := connMap[userId]
		if receiveConn, ok := connMap[dstId]; !ok {

			//todo  分布式  广播消息
			msgString := "接收用户：" + strconv.Itoa(dstId) + " 已经下线"
			delConn(dstId)
			sendConn.sendMsg(msgString)
		} else {
			receiveConn.sendMsg(string(data))
			sendConn.sendMsg("发送成功")
		}

	}
}

//websocket 消息体
type websocketMessage struct {
	UserId  int    `json:"user_id,omitempty""` //谁发的
	DstId   int    `json:"dst_id,omitempty"`   //对端用户ID
	Content string `json:"content,omitempty"`  //消息的内容
}
