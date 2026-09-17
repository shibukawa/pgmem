package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgPageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v7 = l6
	v8 = l7
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(832)
	m.G0 = v17
	if l3 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L11
	} else {
		goto L36
	}
L2:
	;
	m.G0 = v17 + int32(832)
	return
L3:
	;
	v22 = l3 << (uint(int32(1)) % 32)
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	base.MemoryCopy(m, v17+int32(16), l2, v22)
	goto L6
L5:
	;
	goto L6
L6:
	;
	if int32(2) <= l3 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v47 = int32(0)
	v56 = v9
	goto L16
L8:
	;
	v29 = v17 + int32(16)
	F_pg_qsort(m, v29, l3, int32(2), int32(244))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_PageIndexMultiDelete(m, l1, v17+int32(16), l3)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L14
	}
L11:
	;
	return
L12:
	;
	F_PageIndexMultiDelete(m, l1, v29, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	if l3 != int32(1) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(16)+v56<<(uint(int32(1))%32)))))
	if v64 == v43 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L2
L18:
	;
	v66 = l4
	goto L20
L19:
	;
	v66 = l5
	goto L20
L20:
	;
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v104 = F_PageAddItemExtended(m, l1, v99, int32(base.Ui32(v100)>>(uint(int32(2))%32)), v64, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L11
	} else {
		goto L30
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v67&int32(3) == v66 {
		v99 = v47
		v100 = v67
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v66&int32(3) | int32(64)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+4)))
	v81 = v79 & int32(_a_F_spgPageIndexMultiDelete_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+4)) = uint16(v81)
	if v66 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L24
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v99 = v73
	v100 = v98
	goto L21
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+10)) = uint16(v8)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+8)) = uint16(v7)
	v88 = int32(base.Ui32(v7) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)) = uint16(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v90
	goto L26
L28:
	;
	goto L29
L29:
	;
	v92 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+10)) = uint16(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v92
	goto L26
L30:
	;
	if v104 != v64 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	switch v66 - int32(1) {
	case 0:
		v111 = int32(2)
		goto L33
	default:
		goto L32
	case 2:
		goto L34
	}
L32:
	;
	v122 = v56 + int32(1)
	if v122 != l3 {
		v47 = v99
		v56 = v122
		goto L16
	} else {
		goto L35
	}
L33:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v114 = l1 + v112 + v111
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v117 = v115 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v117)
	goto L32
L34:
	;
	v111 = int32(4)
	goto L33
L35:
	;
	goto L17
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(base.Ui32(v145) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spgPageIndexMultiDelete_1), v17)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_spgPageIndexMultiDelete_2), int32(171), int32(_a_F_spgPageIndexMultiDelete_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_spg_quad_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500888)
	return int32(0)
}
func F_spg_range_quad_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935504119)
	return int32(0)
}
func F_spg_text_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(90194313241)
	return int32(0)
}
func F_spg_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_spg_xlog_startup[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(_a_F_spg_xlog_startup_0), int32(0), int32(_a_F_spg_xlog_startup_1), int32(_a_F_spg_xlog_startup_2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_spg_xlog_startup[1])) = v8
		return
	}
}
