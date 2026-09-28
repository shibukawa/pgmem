package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpfamilyIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14219(m, l0, l1, int32(41), int32(_a_F_OpfamilyIsVisibleExt_0), int32(2352), int32(_a_F_OpfamilyIsVisibleExt_1), int32(42))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_get_opfamily_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(42), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_opfamily_name_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_opfamily_name_1), int32(1562), int32(_a_F_get_opfamily_name_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v34 = F_pstrdup(m, v29+v30+int32(8))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v34
				}
			}
		}
	}
}
func F_opfamily_can_sort_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	v11 = int64(0)
	v13 = F_SearchSysCacheList(m, int32(13), int32(1), int64(403), v11, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+80))
	if v36 != l0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v47 = v26 + int32(1)
	if v47 != v17 {
		v26 = v47
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+84))
	if v38 != l1 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	return base.B2i32(v40 != int32(0))
L12:
	;
	goto L7
L13:
	;
	return int32(0)
}
