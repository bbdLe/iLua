package state

import (
	"github.com/bbdLe/iLua/internal/api"
	"github.com/bbdLe/iLua/internal/log"
)

type luaValue interface{}

func typeOf(val luaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64, float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		log.Logger.Fatal("type is not support")
		return api.LUA_TNONE
	}
}
