package state

type luaState struct {
	stack *luaStack
}

func NewLuaState() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
}
