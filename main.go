package main

import (
	"fmt"
	"net"
	"net/http"
	
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	A2S_Header = []byte{ 0xFF, 0xFF, 0xFF, 0xFF }
)

func main() {
	logger, _ = zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()
	go func() {
		logger.Sugar().Info(http.ListenAndServe("localhost:6060", nil))
	}()
	
	fmt.Println("aaa")

	conn, err := net.Dial("udp", "squaller:8211")
	if err != nil {
		logger.Fatal("failed to open a connection to the server", zap.Error(err))
		return
	}
	defer conn.Close()
	logger.Debug("connected to the server")
	
	// A2S_INFO
	conn.Write(A2S_Header)
	conn.Write([]byte("Source Engine Query\x00"))
	buffer := make([]byte, 1400)
	length, err := conn.Read(buffer)
	if err != nil {
		logger.Error("failed to receive data from sender", zap.Error(err))
	}
	
	for i := 0; i < length; i++ {
		fmt.Printf("%d: %x\n", i, buffer[i]);
	}
}
