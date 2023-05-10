package gouuidtools

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
)

type UUID struct {
	v      [16]byte
	Object any
}

func NewUUIDFromString(val string) (*UUID, error) {

	if len(val) > 128 {
		return nil, errors.New("unacceptable uuid string format")
	}

	var val_new string

	for _, i := range val {
		i_str := string(i)
		switch i_str {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f",
			"A", "B", "C", "D", "E", "F":
			val_new += i_str
		}
	}

	if len(val_new) != 32 {
		return nil, errors.New("unacceptable uuid string format")
	}

	val_new = strings.ToLower(val_new)

	var ret [16]byte

	counter := 0
	for len(val_new) != 0 {
		var scanned byte
		fmt.Sscanf(val_new[:2], "%x", &scanned)
		ret[counter] = scanned
		counter++
		val_new = val_new[2:]
	}

	return &UUID{v: ret}, nil
}

func NewUUIDFromByteSlice(val []byte) (*UUID, error) {
	if len(val) != 16 {
		return nil, errors.New("invalid byte slice")
	}

	var ret [16]byte
	for i := 0; i != 16; i++ {
		ret[i] = val[i]
	}

	return &UUID{v: ret}, nil
}

func NewUUIDFromByteArray(val [16]byte) *UUID {
	return &UUID{v: val}
}

func NewUUIDFromRandom() (*UUID, error) {
	ret := []byte{}
	for len(ret) < 16 {
		buf := make([]byte, 16)
		x, err := rand.Read(buf)
		if err != nil {
			return nil, err
		}
		ret = append(ret, buf[0:x]...)
	}
	return NewUUIDFromByteSlice(ret)
}

func (self *UUID) Equal(val *UUID) bool {
	for i := 0; i != 16; i++ {
		if self.v[i] != val.v[i] {
			return false
		}
	}
	return true
}

func (self *UUID) EqualByteSlice(val []byte) bool {
	if len(val) != 16 {
		return false
	}
	for i := 0; i != 16; i++ {
		if self.v[i] != val[i] {
			return false
		}
	}
	return true
}

func (self *UUID) EqualByteArray(val [16]byte) bool {
	for i := 0; i != 16; i++ {
		if self.v[i] != val[i] {
			return false
		}
	}
	return true
}

func IsNil(val *UUID) bool {
	for i := 0; i != 16; i++ {
		if val.v[i] != 0 {
			return false
		}
	}
	return true
}

func (self *UUID) IsNil(val *UUID) bool {
	return IsNil(self)
}

func (self *UUID) format(minuses bool) string {
	var ret string
	for i := 0; i != 16; i++ {
		ret += fmt.Sprintf("%x", self.v[i])
		if minuses {
			switch i {
			case 4, 6, 8, 10:
				ret += "-"
			}
		}
	}
	return ret
}

func (self *UUID) Format() string {
	return self.format(true)
}

func (self *UUID) FormatNoMinuses() string {
	return self.format(false)
}

func (self *UUID) ByteArray() [16]byte {
	return self.v
}

func (self *UUID) ByteSlice() []byte {
	ret := []byte{}
	for _, i := range self.v {
		ret = append(ret, i)
	}
	return ret
}
