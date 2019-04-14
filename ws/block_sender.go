package ws

import (
	"common"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	abci "github.com/tendermint/tendermint/abci/types"
	"net/http"
)

type BlockInfo struct {
	Time   int64  ` json:"time"`
	Hash   []byte `json:"hash"`
	Height int64  `json:"height"`
	Tx     int64  `json:"tx"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewBlockInfoFromReq(req abci.RequestBeginBlock) BlockInfo {
	return BlockInfo{
		Time:   req.Header.Time.Unix(),
		Hash:   req.Hash,
		Height: req.Header.Height,
		Tx:     req.Header.NumTxs,
	}
}

type WsBlockSender struct {
	addr string
	Hub  *Hub
	conn *websocket.Conn
	send chan BlockInfo
}

func NewWsBlockSender(addr string) *WsBlockSender {
	return &WsBlockSender{
		addr: addr,
		send: make(chan BlockInfo, 10),
		Hub:  NewHub(),
	}
}

func (ws WsBlockSender) Start() error {
	go func() {
		go ws.Hub.run()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			common.Log.Event("EventConnectBlockWs", common.Printf("Income connection ws fro blocks"))
			serveWs(ws.Hub, w, r)
		})

		http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("OK"))
		})

		err := http.ListenAndServe(ws.addr, nil)
		if err != nil {
			fmt.Printf("ListenAndServe: %v", err)
		}
	}()

	return nil
}

func (ws WsBlockSender) Stop() error {
	err := ws.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsBlockSender) Send(info BlockInfo) error {
	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}

	ws.Hub.broadcast <- bytes

	return nil
}
