package gouuidtools

import (
	"runtime"

	"github.com/AnimusPEXUS/goreentrantlock"
)

type UUIDRegistry struct {
	ids      [][16]byte
	ids_lock *goreentrantlock.ReentrantMutexCheckable
}

func NewUUIDRegistry() (*UUIDRegistry, error) {
	self := new(UUIDRegistry)
	self.ids_lock = goreentrantlock.NewReentrantMutexCheckable(false)
	return self, nil
}

func (self *UUIDRegistry) GenUUID() (*UUID, error) {

	self.ids_lock.Lock()
	defer self.ids_lock.Unlock()

	var ret *UUID
	var err error

main_loop:
	for true {
		ret, err = NewUUIDFromRandom()
		if err != nil {
			return nil, err
		}

		if self.Registered(ret) {
			continue main_loop
		}
		break
	}

	self.Register(ret)

	return ret, nil
}

func (self *UUIDRegistry) Registered(val *UUID) bool {

	self.ids_lock.Lock()
	defer self.ids_lock.Unlock()

	for _, x := range self.ids {
		if val.EqualByteArray(x) {
			return true
		}
	}
	return false
}

func (self *UUIDRegistry) Register(val *UUID) {
	self.ids_lock.Lock()
	defer self.ids_lock.Unlock()

	defer runtime.SetFinalizer(val, self.Unregister)

	if self.Registered(val) {
		return
	}

	self.ids = append(self.ids, val.ByteArray())

	return

}

func (self *UUIDRegistry) Unregister(val *UUID) {

	self.ids_lock.Lock()
	defer self.ids_lock.Unlock()

	for i := len(self.ids) - 1; i != -1; i-- {
		if val.EqualByteArray(self.ids[i]) {
			self.ids = append(self.ids[:i], self.ids[i+1:]...)
		}
	}
	runtime.SetFinalizer(val, nil)
	return
}
