package utils

func PanicIf(err error) {

	if err != nil {
		panic(err)
	}
}

func Assert(cond bool, msg string) {

	if !cond {
		panic(msg)
	}
}

func SafeCall(f func(), onPanic func(any)) {
	defer func() {
		if r := recover(); r != nil {
			onPanic(r)
		}
	}()
	f()
}
