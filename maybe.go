package goehm

type Maybe[T any] func(func(some T), func())

func Some[T any](some T) (Maybe[T]) {
	return Maybe[T](func(withSome func(some_ T), withNone func()){
		withSome(some)
	})
}

func None[T any]() (Maybe[T]) {
	return Maybe[T](func(withSome func(some_ T), withNone func()){
		withNone()
	})
}

//func IsSome[T any](maybe Maybe[T]) bool {
	//maybe(func(some T) { return true }, func() { return false} )
//}
