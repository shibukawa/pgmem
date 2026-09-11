package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeItemUnary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != int32(2) {
		v28 = F_palloc(m, int32(24))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			if v31 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
					return v28
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
				return v28
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			v28 = F_palloc(m, int32(24))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				if v31 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
						return v28
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
					return v28
				}
			}
		} else {
			v8 = F_palloc(m, int32(24))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				if v13 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2)
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v21 = F_DirectFunctionCall1Coll(m, int32(1421), int32(0), v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							v23 = F_pg_detoast_datum(m, v21)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
								return v8
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v21 = F_DirectFunctionCall1Coll(m, int32(1421), int32(0), v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_pg_detoast_datum(m, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
							return v8
						}
					}
				}
			}
		}
	}
}
