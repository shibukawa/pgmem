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
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatGetMetaPageInfo[0]))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(v6^int32(-1))<<(uint(int32(2))%32))))
				v28 = v20
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatGetMetaPageInfo[1]))
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
					F_errmsg_internal(m, int32(_a_F_IvfflatGetMetaPageInfo_0), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_IvfflatGetMetaPageInfo_1), int32(220), int32(_a_F_IvfflatGetMetaPageInfo_2))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13830(m, l0, int32(5), int32(_a_F_IvfflatGetTypeInfo_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
			v13 = int32(_a_F_IvfflatInitRegisterPage_0)
			v15 = int32(0)
			if v15|(v10&int32(3)|int32(1)) == v15 {
				v31 = v10 + v13
				v33 = v10 + int32(4)
				if base.Ui32(v33) < base.Ui32(v31) {
					v35 = v31
				} else {
					v35 = v33
				}
				v40 = (v10^int32(-1)+v35)&int32(-4) + int32(4)
				if v40 == int32(0) {
				} else {
					base.MemoryFill(m, v10, int32(0), v40)
				}
			} else {
				base.MemoryFill(m, v10, int32(0), v13)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+10)) = int32(_a_F_IvfflatInitRegisterPage_1)
			v54 = int32(_a_F_IvfflatInitRegisterPage_2)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+18)) = uint16(v54)
			v60 = int32(_a_F_IvfflatInitRegisterPage_3)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)) = uint16(v60)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)) = uint16(v60)
			v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
			v64 = v10 + v63
			v65 = int32(_a_F_IvfflatInitRegisterPage_4)
			*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v65)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
			return
		}
	}
}
