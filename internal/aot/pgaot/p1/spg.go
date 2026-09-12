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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
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
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v17 + int32(832)
	return
L3:
	;
	v24 = l3 << (uint(int32(1)) % 32)
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(2) <= l3 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v25 = F__emscripten_memcpy_bulkmem(m, v17+int32(16), l2, v24)
	mBase = m.M
	goto L7
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v49 = int32(0)
	v58 = v9
	goto L17
L9:
	;
	F_pg_qsort(m, v17+int32(16), l3, int32(2), int32(244))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_PageIndexMultiDelete(m, l1, v17+int32(16), l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	F_PageIndexMultiDelete(m, l1, v17+int32(16), l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	if l3 != int32(1) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(16)+v58<<(uint(int32(1))%32)))))
	if v66 == v45 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L2
L19:
	;
	v68 = l4
	goto L21
L20:
	;
	v68 = l5
	goto L21
L21:
	;
	if v49 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v106 = F_PageAddItemExtended(m, l1, v101, int32(base.Ui32(v102)>>(uint(int32(2))%32)), v66, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L31
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v69&int32(3) == v68 {
		v101 = v49
		v102 = v69
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v68&int32(3) | int32(64)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
	v83 = v81 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v83)
	if v68 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v101 = v75
	v102 = v100
	goto L22
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+10)) = uint16(v8)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+8)) = uint16(v7)
	v90 = int32(base.Ui32(v7) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)) = uint16(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v92
	goto L27
L29:
	;
	goto L30
L30:
	;
	v94 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+10)) = uint16(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v94
	goto L27
L31:
	;
	if v106 != v66 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	switch v68 - int32(1) {
	case 0:
		v113 = int32(2)
		goto L34
	default:
		goto L33
	case 2:
		goto L35
	}
L33:
	;
	v124 = v58 + int32(1)
	if v124 != l3 {
		v49 = v101
		v58 = v124
		goto L17
	} else {
		goto L36
	}
L34:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v116 = v114 + (l1 + v113)
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116))))
	v119 = v117 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v119)
	goto L33
L35:
	;
	v113 = int32(4)
	goto L34
L36:
	;
	goto L18
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(base.Ui32(v147) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(416257), v17)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(503829), int32(171), int32(359456))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(61324), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[121])) = v8
		return
	}
}
