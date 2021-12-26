package logger

import (
	"context"
	"unsafe"
)

// Values create contract store value into context
type Values interface {
	Set(key Flags, value interface{})
	Load(key Flags) (interface{}, bool)
	LoadAndDelete(key Flags) (interface{}, bool)
}

// Set value to keys
func (l *Locker) Set(key Flags, value interface{}) {
	l.data.Store(key, value)
}

// Load value from key
func (l *Locker) Load(key Flags) (interface{}, bool) {
	return l.data.Load(key)
}

// LoadAndDelete from key
func (l *Locker) LoadAndDelete(key Flags) (interface{}, bool) {
	return l.data.LoadAndDelete(key)
}

func getKeyContext(ctx context.Context) Key {
	iCtx := (*iface)(unsafe.Pointer(&ctx))
	valCtx := (*valueCtx)(unsafe.Pointer(iCtx.data))
	return valCtx.key.(Key)
}

func extract(ctx context.Context) (Values, bool) {
	var (
		lock = new(Locker)
		ok   bool
	)

	if ctx == nil {
		return lock, false
	}

	key := getKeyContext(ctx)

	lock, ok = ctx.Value(key).(*Locker)
	return lock, ok
}
