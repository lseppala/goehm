package goehm

//type Maybe[T any] forall[T, any]
//type forall[T any, R any] func(func(some T) R, func() R) R


//func Some[T any](some T) (Maybe[T]) {
	//return Maybe[T](
		//forall[T, R](
			//func(withSome func(some_ T) R, withNone func()) R {
				//return withSome(some)
			//},
		//),
	//)
//}

//func None[T any]() (Maybe[T]) {
	//return Maybe[T](func(withSome func(some_ T), withNone func()){
		//withNone()
	//})
//}

//func IsSome[T any](maybe Maybe[T]) bool {
	//return maybe(func(some T) { return true }, func() { return false} )
//}


type Maybe[T any] forall[T, any]
type forall[T any, R any] struct{}


func (forall forall[T, R]) Foo(withSome func(some T) R, none R) R {
	return none
}
