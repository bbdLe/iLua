package binchunk

import (
	"encoding/binary"
	"math"

	"github.com/bbdLe/iLua/internal/log"
)

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *reader) readUint32() uint32 {
	n := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return n
}

func (self *reader) readUint64() uint64 {
	n := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return n
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	if size == 0 {
		return ""
	}
	if size == 0xFF {
		size = uint(self.readUint64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes)
}

func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE {
		log.Logger.Fatal("not a precompiled chunk!")
	}
	if self.readByte() != LUAC_VERSION {
		log.Logger.Fatal("version mismatch!")
	}
	if self.readByte() != LUAC_FORMAT {
		log.Logger.Fatal("format mismatch!")
	}
	if string(self.readBytes(6)) != LUAC_DATA {
		log.Logger.Fatal("corrupted!")
	}
	if self.readByte() != CINT_SIZE {
		log.Logger.Fatal("int size mismatch!")
	}
	if self.readByte() != CSIZET_SIZE {
		log.Logger.Fatal("size_t size mismatch!")
	}
	if self.readByte() != INSTRUCTION_SIZE {
		log.Logger.Fatal("instruction size mismatch!")
	}
	if self.readByte() != LUA_INTEGER_SIZE {
		log.Logger.Fatal("lua integer size mismatch!")
	}
	if self.readByte() != LUA_NUMBER_SIZE {
		log.Logger.Fatal("lua number size mismatch!")
	}
	if self.readLuaInteger() != LUAC_INT {
		log.Logger.Fatal("endianness mismatch!")
	}
	if self.readLuaNumber() != LUAC_NUM {
		log.Logger.Fatal("float format mismatch!")
	}
	log.Logger.Info("check success")
}

func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {
		source = parentSource
	}
	return &Prototype{
		Source:         source,
		LineDefine:     self.readUint32(),
		LastLineDefine: self.readUint32(),
		NumParams:      self.readByte(),
		IsVararg:       self.readByte(),
		MaxStackSize:   self.readByte(),
		Code:           self.readCode(),
		Constants:      self.readConstants(),
		Upvalues:       self.readUpValues(),
		Protos:         self.readProtos(source),
		LineInfo:       self.readLineInfo(),
		LocVars:        self.readLocVars(),
		UpvalueNames:   self.readUpValueNames(),
	}
}

func (self *reader) readCode() []uint32 {
	codes := make([]uint32, self.readUint32())
	for i := range codes {
		codes[i] = self.readUint32()
	}
	return codes
}

func (self *reader) readConstants() []interface{} {
	constants := make([]interface{}, self.readUint32())
	for i := range constants {
		constants[i] = self.readConstant()
	}

	return constants
}

func (self *reader) readConstant() interface{} {
	switch self.readByte() {
	case TAG_NIL:
		return nil
	case TAG_BOOLEAN:
		return self.readByte() != 0
	case TAG_INTEGER:
		return self.readLuaInteger()
	case TAG_NUMBER:
		return self.readLuaNumber()
	case TAG_LONG_STR, TAG_SHORT_STR:
		return self.readString()
	default:
		log.Logger.Fatal("unknown constant type!")
		return nil
	}
}

func (self *reader) readUpValues() []Upvalue {
	upvalues := make([]Upvalue, self.readUint32())
	for i := range upvalues {
		upvalues[i] = Upvalue{
			Instack: self.readByte(),
			Idx:     self.readByte(),
		}
	}
	return upvalues
}

func (self *reader) readProtos(source string) []*Prototype {
	protos := make([]*Prototype, self.readUint32())
	for i := range protos {
		protos[i] = self.readProto(source)
	}
	return protos
}

func (self *reader) readLineInfo() []uint32 {
	lineInfos := make([]uint32, self.readUint32())
	for i := range lineInfos {
		lineInfos[i] = self.readUint32()
	}
	return lineInfos
}

func (self *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, self.readUint32())
	for i := range locVars {
		locVars[i] = LocVar{
			VarName: self.readString(),
			StartPC: self.readUint32(),
			EndPC:   self.readUint32(),
		}
	}
	return locVars
}

func (self *reader) readUpValueNames() []string {
	names := make([]string, self.readUint32())
	for i := range names {
		names[i] = self.readString()
	}
	return names
}
