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
	var v12 int64
	_ = v12
	var v15 int32
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
	v12 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
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
						F_sequence_close(m, v19, int32(1))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v17 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v13+int32(32), int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = int32(3)
	F_ScanKeyInit(m, v13+int32(80), v28, v28, int32(63), l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = F_systable_beginscan(m, v17, int32(2656), int32(1), int32(0), int32(2), v13+int32(32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = F_systable_getnext(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v48 = v41
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v39)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(2604)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54)))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v58
	F_performDeletion(m, v13+int32(20), v59, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v67 = F_systable_getnext(m, v39)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v67 != 0 {
		v48 = v67
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_sequence_close(m, v17, int32(3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	m.G0 = v13 + int32(128)
	return
L18:
	;
	if v41 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(470593), v13)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(495115), int32(196), int32(98096))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
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
	var v22 int32
	_ = v22
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
	v22 = v19
	goto L7
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v36
	m.G0 = v8 + int32(32)
	return v19
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		v36 = v22
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v32)
	v36 = v22 + int32(1)
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
	v22 = v22 + int32(1)
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
	F_errmsg(m, int32(402714), int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(4385680)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errdetail(m, int32(649264), v8+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(494035), int32(753), int32(342901))
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
	F_errmsg(m, int32(402714), int32(0))
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
	F_errdetail(m, int32(650682), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(494035), int32(760), int32(342901))
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
