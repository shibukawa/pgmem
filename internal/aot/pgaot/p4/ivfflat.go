package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IvfflatGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v6 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_LockBuffer(m, v6, int32(1))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if v6 < int32(0) {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(v6^int32(-1))<<(uint(int32(2))%32))))
				v28 = v20
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				v28 = v22 + v6<<(uint(int32(13))%32) + int32(-8192)
			}
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
			if v29 == int32(22016423) {
				if l1 != 0 {
					v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+34)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
				} else {
				}
				if l2 != 0 {
					v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+32)))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
				} else {
				}
				F_UnlockReleaseBuffer(m, v6)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(457780), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(518927), int32(220), int32(255405))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_IvfflatGetTypeInfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = int32(4120992)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8*int32(0)<<(uint(int32(2))%32)+int32(20)-int32(4))))
	if v20 == int32(0) {
		v34 = v3
		return v34
	} else {
		v25 = F_index_getprocinfo(m, l0, int32(1), int32(5))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v25 == int32(0) {
				v34 = v3
				return v34
			} else {
				v31 = F_FunctionCall0Coll(m, v25)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = v31
					return v34
				}
			}
		}
	}
}
func F_IvfflatInitRegisterPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v10 = F_GenericXLogRegisterBuffer(m, v5, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v10
			if v10&int32(3) != 0 {
			} else {
			}
			v39 = F___memset(m, v10, int32(0), int32(8192))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v10)+10)) = int32(1572864)
			v45 = int32(8196)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+18)) = uint16(v45)
			v51 = int32(8184)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)) = uint16(v51)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)) = uint16(v51)
			v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
			v55 = v10 + v54
			v56 = int32(65412)
			*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)) = uint16(v56)
			*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(-1)
			return
		}
	}
}
