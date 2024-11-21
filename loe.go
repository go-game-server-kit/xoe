package xoe

func Loe(err error, args ...Args) bool {
	if err != nil {
		Log("Loe", err, args)
		return true
	}
	return false
}

func LoeFn(fn func() error, args ...Args) bool {
	err := fn()
	if err != nil {
		Log("LoeFn", err, args)
		return true
	}
	return false
}

func Loe1With[T1 any](t1 T1, err error) func(args ...Args) T1 {
	return func(args ...Args) T1 {
		if err != nil {
			Log("Loe1", err, args)
		}
		return t1
	}
}

func Loe1[T1 any](t1 T1, err error) T1 {
	return Loe1With(t1, err)(WithStackSkip(1))
}

func Loe2With[T1 any, T2 any](t1 T1, t2 T2, err error) func(args ...Args) (T1, T2) {
	return func(args ...Args) (T1, T2) {
		if err != nil {
			Log("Loe2", err, args)
		}
		return t1, t2
	}
}

func Loe2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	return Loe2With(t1, t2, err)(WithStackSkip(1))
}

func Loe3With[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) func(args ...Args) (T1, T2, T3) {
	return func(args ...Args) (T1, T2, T3) {
		if err != nil {
			Log("Loe3", err, args)
		}
		return t1, t2, t3
	}
}

func Loe3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	return Loe3With(t1, t2, t3, err)(WithStackSkip(1))
}

func Loe4With[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) func(args ...Args) (T1, T2, T3, T4) {
	return func(args ...Args) (T1, T2, T3, T4) {
		if err != nil {
			Log("Loe4", err, args)
		}
		return t1, t2, t3, t4
	}
}

func Loe4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	return Loe4With(t1, t2, t3, t4, err)(WithStackSkip(1))
}

func Loe5With[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) func(args ...Args) (T1, T2, T3, T4, T5) {
	return func(args ...Args) (T1, T2, T3, T4, T5) {
		if err != nil {
			Log("Loe5", err, args)
		}
		return t1, t2, t3, t4, t5
	}
}

func Loe5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	return Loe5With(t1, t2, t3, t4, t5, err)(WithStackSkip(1))
}

func Loe6With[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) func(args ...Args) (T1, T2, T3, T4, T5, T6) {
	return func(args ...Args) (T1, T2, T3, T4, T5, T6) {
		if err != nil {
			Log("Loe6", err, args)
		}
		return t1, t2, t3, t4, t5, t6
	}
}

func Loe6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	return Loe6With(t1, t2, t3, t4, t5, t6, err)(WithStackSkip(1))
}
