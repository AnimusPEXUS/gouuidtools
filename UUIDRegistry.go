package gouuidtools

import (
	"runtime"
	"sync"

	"github.com/AnimusPEXUS/golockerreentrancycontext"
)

type UUIDRegistry struct {
	ids      [][16]byte
	ids_lock *sync.Mutex
}

func NewUUIDRegistry() (*UUIDRegistry, error) {
	self := new(UUIDRegistry)
	self.ids_lock = &sync.Mutex{}
	return self, nil
}

func (self *UUIDRegistry) GenUUID() (*UUID, error) {
	return self.GenUUID_lrc(nil)
}

func (self *UUIDRegistry) GenUUID_lrc(
	lrc *golockerreentrancycontext.LockerReentrancyContext,
) (*UUID, error) {
	if lrc == nil {
		lrc = new(golockerreentrancycontext.LockerReentrancyContext)
	}
	lrc.LockMutex(self.ids_lock)
	defer lrc.UnlockMutex(self.ids_lock)

	var ret *UUID
	var err error

main_loop:
	for true {
		ret, err = NewUUIDFromRandom()
		if err != nil {
			return nil, err
		}

		if self.Registered_lrc(ret, lrc) {
			continue main_loop
		}
		break
	}

	self.Register_lrc(ret, lrc)

	return ret, nil
}

func (self *UUIDRegistry) Registered(val *UUID) bool {
	return self.Registered_lrc(val, nil)
}

func (self *UUIDRegistry) Registered_lrc(
	val *UUID,
	lrc *golockerreentrancycontext.LockerReentrancyContext,
) bool {
	if lrc == nil {
		lrc = new(golockerreentrancycontext.LockerReentrancyContext)
	}
	lrc.LockMutex(self.ids_lock)
	defer lrc.UnlockMutex(self.ids_lock)

	for _, x := range self.ids {
		if val.EqualByteArray(x) {
			return true
		}
	}
	return false
}

func (self *UUIDRegistry) Register(val *UUID) {
	self.Register_lrc(val, nil)
}

func (self *UUIDRegistry) Register_lrc(
	val *UUID,
	lrc *golockerreentrancycontext.LockerReentrancyContext,
) {
	if lrc == nil {
		lrc = new(golockerreentrancycontext.LockerReentrancyContext)
	}
	lrc.LockMutex(self.ids_lock)
	defer lrc.UnlockMutex(self.ids_lock)

	defer runtime.SetFinalizer(val, self.Unregister)

	if self.Registered_lrc(val, lrc) {
		return
	}

	self.ids = append(self.ids, val.ByteArray())

	return
}

func (self *UUIDRegistry) Unregister(val *UUID) {
	self.Unregister_lrc(val, nil)
}

func (self *UUIDRegistry) Unregister_lrc(
	val *UUID,
	lrc *golockerreentrancycontext.LockerReentrancyContext,
) {
	if lrc == nil {
		lrc = new(golockerreentrancycontext.LockerReentrancyContext)
	}
	lrc.LockMutex(self.ids_lock)
	defer lrc.UnlockMutex(self.ids_lock)

	for i := len(self.ids) - 1; i != -1; i-- {
		if val.EqualByteArray(self.ids[i]) {
			self.ids = append(self.ids[:i], self.ids[i+1:]...)
		}
	}
	runtime.SetFinalizer(val, nil)
	return
}
