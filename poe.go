package xoe

func Poe(err error, args ...Args) bool {
	if err != nil {
		Report("Poe", err, args)
		panic(err)
	}
	return false
}

func PoeFn(fn func() error, args ...Args) bool {
	err := fn()
	if err != nil {
		Report("PoeFn", err, args)
		panic(err)
	}
	return false
}

func Poe1With[T1 any](t1 T1, err error) func(args ...Args) T1 {
	return func(args ...Args) T1 {
		if err != nil {
			Report("Poe1", err, args)
			panic(err)
		}
		return t1
	}
}

func Poe1[T1 any](t1 T1, err error) T1 {
	return Poe1With(t1, err)(WithStackSkip(1))
}

func Poe2With[T1 any, T2 any](t1 T1, t2 T2, err error) func(args ...Args) (T1, T2) {
	return func(args ...Args) (T1, T2) {
		if err != nil {
			Report("Poe2", err, args)
			panic(err)
		}
		return t1, t2
	}
}

func Poe2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	return Poe2With(t1, t2, err)(WithStackSkip(1))
}

func Poe3With[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) func(args ...Args) (T1, T2, T3) {
	return func(args ...Args) (T1, T2, T3) {
		if err != nil {
			Report("Poe3", err, args)
			panic(err)
		}
		return t1, t2, t3
	}
}

func Poe3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	return Poe3With(t1, t2, t3, err)(WithStackSkip(1))
}

func Poe4With[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) func(args ...Args) (T1, T2, T3, T4) {
	return func(args ...Args) (T1, T2, T3, T4) {
		if err != nil {
			Report("Poe4", err, args)
			panic(err)
		}
		return t1, t2, t3, t4
	}
}

func Poe4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	return Poe4With(t1, t2, t3, t4, err)(WithStackSkip(1))
}

func Poe5With[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) func(args ...Args) (T1, T2, T3, T4, T5) {
	return func(args ...Args) (T1, T2, T3, T4, T5) {
		if err != nil {
			Report("Poe5", err, args)
			panic(err)
		}
		return t1, t2, t3, t4, t5
	}
}

func Poe5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	return Poe5With(t1, t2, t3, t4, t5, err)(WithStackSkip(1))
}

func Poe6With[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) func(args ...Args) (T1, T2, T3, T4, T5, T6) {
	return func(args ...Args) (T1, T2, T3, T4, T5, T6) {
		if err != nil {
			Report("Poe6", err, args)
			panic(err)
		}
		return t1, t2, t3, t4, t5, t6
	}
}

func Poe6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	return Poe6With(t1, t2, t3, t4, t5, t6, err)(WithStackSkip(1))
}
