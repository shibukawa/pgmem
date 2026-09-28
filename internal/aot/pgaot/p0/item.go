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
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != int32(2) {
		v29 = F_palloc(m, int32(24))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_makeItemUnary[0]))
			if v32 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
					return v29
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
				return v29
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			v29 = F_palloc(m, int32(24))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_makeItemUnary[0]))
				if v32 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
						return v29
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
					return v29
				}
			}
		} else {
			v8 = F_palloc(m, int32(24))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_makeItemUnary[0]))
				if v13 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2)
						v20 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
						v21 = F_DirectFunctionCall1Coll(m, int32(1546), int32(0), v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							v24 = F_pg_detoast_datum(m, base.I32_wrap_i64(v21))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v24
								return v8
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2)
					v20 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
					v21 = F_DirectFunctionCall1Coll(m, int32(1546), int32(0), v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = F_pg_detoast_datum(m, base.I32_wrap_i64(v21))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v24
							return v8
						}
					}
				}
			}
		}
	}
}
