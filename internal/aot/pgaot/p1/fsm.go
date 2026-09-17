package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fsm_page_contents(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_superuser(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = F_get_page_from_raw(m, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	m.G0 = v8 + int32(48)
	return v68
L8:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v68 = int32(0)
	goto L7
L10:
	;
	goto L11
L11:
	;
	F_initStringInfo(m, v8+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = int32(0)
	goto L13
L13:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+(v17+int32(28))))))
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
	F_appendStringInfo(m, v8+int32(32), int32(_a_F_fsm_page_contents_0), v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v32
	F_appendStringInfo(m, v8+int32(32), int32(_a_F_fsm_page_contents_1), v8+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v49 = v32 + int32(1)
	if v49 != int32(_a_F_fsm_page_contents_2) {
		v32 = v49
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	goto L14
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	v61 = F_cstring_to_text_with_len(m, v59, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v68 = v61
	goto L7
L22:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_fsm_page_contents_3), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_fsm_page_contents_4), int32(46), int32(_a_F_fsm_page_contents_5))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fsm_set_avail(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v3 = l2
	v9 = l0 + int32(28)
	v11 = l1 + int32(4095)
	v12 = v9 + v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v3)
	v23 = v11
	goto L4
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.Ui32(v15) < base.Ui32(v3) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v27 = int32(1)
	v28 = v23 - v27
	v29 = int32(2)
	v30 = base.I32_div_s(v28, v29)
	v32 = v30 << (uint(v27) % 32)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v32)+1)))
	v36 = v32 + v29
	if base.Ui32(v36) <= base.Ui32(int32(_a_F_fsm_set_avail_0)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.Ui32(v55) < base.Ui32(v3) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v40 = v34 & int32(255)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v36))))
	if base.Ui32(v42) < base.Ui32(v40) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v45 = v34
	goto L8
L8:
	;
	v47 = v30 + v9
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != v45&int32(255) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v44 = v40
	goto L11
L10:
	;
	v44 = v42
	goto L11
L11:
	;
	v45 = v44
	goto L8
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v45)
	if int32(1) < v28 {
		v23 = v30
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L5
L15:
	;
	goto L14
L16:
	;
	v61 = int32(4094)
	goto L19
L17:
	;
	goto L18
L18:
	;
	return int32(1)
L19:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v61) {
		v82 = int32(0)
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v83 = v61 + v9
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != v82&int32(255) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v71 = v61 << (uint(int32(1)) % 32)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(29)+v71))))
	if v61 == int32(4081) {
		v82 = v73
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v9)+2)))
	if base.Ui32(v77) < base.Ui32(v73) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = v73
	goto L26
L25:
	;
	v79 = v77
	goto L26
L26:
	;
	v82 = v79
	goto L21
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v82)
	goto L29
L28:
	;
	goto L29
L29:
	;
	if v61 != 0 {
		v61 = v61 - int32(1)
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L20
}
