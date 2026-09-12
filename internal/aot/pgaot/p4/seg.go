package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 float32
	_ = v31
	var v32 float32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 float32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_palloc(m, int32(40))
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
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	switch v17 - int32(60) {
	case 0, 2:
		goto L4
	case 1:
		v29 = v13
		goto L3
	default:
		goto L5
	}
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	v31 = *(*float32)(unsafe.Add(mBase, uint32(v11)))
	v32 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.F32_ne(v31, v32) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
	v26 = F_pg_sprintf(m, v13, int32(503012), v9+int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	if v17 != int32(126) {
		v29 = v13
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v29 = v26 + v13
	goto L3
L8:
	;
	m.G0 = v9 + int32(32)
	return v13
L9:
	;
	if v30&int32(255) != int32(45) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v34 != v30&int32(255) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+8)))
	v39 = F_restore(m, v29, v31, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+8)))
	v46 = F_restore(m, v29, v31, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v54 = v29
	goto L15
L15:
	;
	v57 = F_pg_sprintf(m, v54, int32(660895), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v48 = v46 + v29
	v51 = F_pg_sprintf(m, v48, int32(746695), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = v51 + v48
	goto L15
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v59 == int32(45) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v62 = v54 + v57
	v65 = F_pg_sprintf(m, v62, int32(746695), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v67 = v65 + v62
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	switch v68 - int32(60) {
	case 0, 2:
		goto L22
	default:
		goto L23
	}
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+9)))
	v83 = F_restore(m, v80, v81, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = base.I32_extend8_s(v68)
	v77 = F_pg_sprintf(m, v67, int32(503012), v9)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v71 != int32(126) {
		v80 = v67
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v80 = v77 + v67
	goto L21
L26:
	;
	goto L8
}
func F_seg_right(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)+4))
	return base.F32_gt(v3, v5)
}
func F_seg_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(6670), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F_seg_yy_scan_string(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = F_strlen(m, l0)
	v4 = F_seg_yy_scan_bytes(m, l0, v3, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_seg_yy_switch_to_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	F_seg_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v8 == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8+v11<<(uint(int32(2))%32))))
			if v15 == l0 {
			} else {
				if v15 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21<<(uint(int32(2))%32))))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v28
					v30 = v20
					v32 = v21
				} else {
					v30 = v8
					v32 = v11
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32)))) = l0
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			}
		}
		return
	}
}
