package main

import (
	"github.com/bbdLe/iLua/internal/binchunk"
	"github.com/bbdLe/iLua/internal/log"
	"go.uber.org/zap"
	"io/ioutil"
)

func main() {
	log.Logger.Info("lua start")
	data, err := ioutil.ReadFile("hello_world.out")
	if err != nil {
		log.Logger.Fatal("read file fail", zap.Error(err))
	}
	binchunk.Undump(data)
}
