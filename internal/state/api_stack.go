package state

import "github.com/bbdLe/iLua/internal/log"

func (self *luaState) GetTop() int {
	return self.stack.top
}

func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

func (self *luaState) Rotate(idx, n int) {
	t := self.GetTop() - 1
	p := self.AbsIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(t, m)
	self.stack.reverse(m+1, p)
	self.stack.reverse(t, p)
}

func (self *luaState) SetTop(idx int) {
	newTop := self.AbsIndex(idx)
	if newTop < 0 {
		log.Logger.Fatal("stack underflow!")
	}

	n := self.GetTop() - newTop
	if n > 0 {
		self.Pop(n)
	} else {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
