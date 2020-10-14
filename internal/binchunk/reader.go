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
