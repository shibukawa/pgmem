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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 + int32(-64)
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
		F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v28 = int32(1)
			v31 = F_systable_beginscan(m, v19, int32(2657), v28, int32(0), v28, v9)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = F_systable_getnext(m, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v33 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
						v39 = v35 + v36
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v40
						v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42
					} else {
					}
					F_systable_endscan(m, v31)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_relation_close(m, v19, int32(1))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							m.G0 = v9 - int32(-64)
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(144)
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
	F_ScanKeyInit(m, v18, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(3)
	F_ScanKeyInit(m, v11+int32(88), v27, v27, int32(63), base.I64_extend_i32_s(l1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = F_systable_beginscan(m, v15, int32(2656), int32(1), int32(0), int32(2), v18)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v11 + int32(144)
	return
L6:
	;
	v39 = F_systable_getnext(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = v39
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2604)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)))
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v54
	F_performDeletion(m, v11+int32(20), v55, l3)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v63 = F_systable_getnext(m, v37)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v63 != 0 {
		v48 = v63
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
	v69 = m.ExcPending
	if v69 != 0 {
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
	v74 = m.ExcPending
	if v74 != 0 {
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
	v80 = m.ExcPending
	if v80 != 0 {
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
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_RemoveAttrDefault_1), int32(201), int32(_a_F_RemoveAttrDefault_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	var v64 int32
	_ = v64
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
	var v84 int32
	_ = v84
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
	v64 = F_errdetail(m, int32(_a_F_read_attr_value_2), v8+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_read_attr_value_3), int32(751), int32(_a_F_read_attr_value_4))
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
	v84 = F_errdetail(m, int32(_a_F_read_attr_value_5), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_read_attr_value_3), int32(758), int32(_a_F_read_attr_value_4))
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
