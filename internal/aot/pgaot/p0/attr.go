package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAttrDefaultColumnAddress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetAttrDefaultColumnAddress[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_GetAttrDefaultColumnAddress[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v15
	v19 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), l1)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v27 = int32(1)
			v30 = F_systable_beginscan(m, v19, int32(2657), v27, int32(0), v27, v9)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = F_systable_getnext(m, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
						v38 = v34 + v35
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
						v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
					} else {
					}
					F_systable_endscan(m, v30)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_relation_close(m, v19, int32(1))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					}
				}
			}
		}
	}
}
func F_RemoveAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v15 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = v11 + int32(32)
	F_ScanKeyInit(m, v18, int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = int32(3)
	F_ScanKeyInit(m, v11+int32(80), v26, v26, int32(63), l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_systable_beginscan(m, v15, int32(2656), int32(1), int32(0), int32(2), v18)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v11 + int32(128)
	return
L6:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = v37
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2604)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48)))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v52
	F_performDeletion(m, v11+int32(20), v53, l3)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v61 = F_systable_getnext(m, v35)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v61 != 0 {
		v46 = v61
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l2 == int32(0) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_RemoveAttrDefault_0), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_RemoveAttrDefault_1), int32(196), int32(_a_F_RemoveAttrDefault_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_read_attr_value(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 == l1&int32(255) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L20
	}
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 != int32(61) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v19 = v10 + int32(2)
	v21 = v19
	goto L7
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v36
	m.G0 = v8 + int32(32)
	return v19
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v25 == int32(0) {
		v36 = v21
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v32)
	v36 = v21 + int32(1)
	goto L6
L9:
	;
	if v25 != int32(44) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v21 = v21 + int32(1)
	goto L7
L11:
	;
	goto L12
L12:
	;
	goto L8
L13:
	;
	return int32(0)
L14:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errmsg(m, int32(_a_F_read_attr_value_0), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	F_sanitize_char_2(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_read_attr_value_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errdetail(m, int32(_a_F_read_attr_value_2), v8+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_read_attr_value_3), int32(753), int32(_a_F_read_attr_value_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_read_attr_value_0), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	F_errdetail(m, int32(_a_F_read_attr_value_5), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_read_attr_value_3), int32(760), int32(_a_F_read_attr_value_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
