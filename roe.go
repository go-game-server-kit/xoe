package xoe

func Roe(err error, args ...Args) bool {
	if err != nil {
		Report("Roe", err, args)
		return true
	}
	return false
}

func Roe1With[T1 any](t1 T1, err error) func(args ...Args) T1 {
	return func(args ...Args) T1 {
		if err != nil {
			Report("Roe1", err, args)
		}
		return t1
	}
}

func Roe1[T1 any](t1 T1, err error) T1 {
	return Roe1With(t1, err)(WithStackSkip(1))
}

func Roe2With[T1 any, T2 any](t1 T1, t2 T2, err error) func(args ...Args) (T1, T2) {
	return func(args ...Args) (T1, T2) {
		if err != nil {
			Report("Roe2", err, args)
		}
		return t1, t2
	}
}

func Roe2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	return Roe2With(t1, t2, err)(WithStackSkip(1))
}

func Roe3With[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) func(args ...Args) (T1, T2, T3) {
	return func(args ...Args) (T1, T2, T3) {
		if err != nil {
			Report("Roe3", err, args)
		}
		return t1, t2, t3
	}
}

func Roe3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	return Roe3With(t1, t2, t3, err)(WithStackSkip(1))
}

func Roe4With[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) func(args ...Args) (T1, T2, T3, T4) {
	return func(args ...Args) (T1, T2, T3, T4) {
		if err != nil {
			Report("Roe4", err, args)
		}
		return t1, t2, t3, t4
	}
}

func Roe4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	return Roe4With(t1, t2, t3, t4, err)(WithStackSkip(1))
}

func Roe5With[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) func(args ...Args) (T1, T2, T3, T4, T5) {
	return func(args ...Args) (T1, T2, T3, T4, T5) {
		if err != nil {
			Report("Roe5", err, args)
		}
		return t1, t2, t3, t4, t5
	}
}

func Roe5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	return Roe5With(t1, t2, t3, t4, t5, err)(WithStackSkip(1))
}

func Roe6With[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) func(args ...Args) (T1, T2, T3, T4, T5, T6) {
	return func(args ...Args) (T1, T2, T3, T4, T5, T6) {
		if err != nil {
			Report("Roe6", err, args)
		}
		return t1, t2, t3, t4, t5, t6
	}
}

func Roe6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	return Roe6With(t1, t2, t3, t4, t5, t6, err)(WithStackSkip(1))
}
