package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistGetLeafTupleSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = F_heap_compute_data_size(m, l0, l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = v5 + int32(23)
		if base.Ui32(v10) <= base.Ui32(int32(16)) {
			v13 = int32(16)
		} else {
			v13 = v10
		}
		return v13 & int32(-8)
	}
}
func F_SpGistPageAddNewItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v16 = l0 + v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v176
L2:
	;
	v155 = int32(0)
	v157 = F_PageAddItemExtended(m, l0, l1, l2, v155, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L32
	} else {
		goto L42
	}
L3:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v22 = v20 - v21
	v23 = int32(0)
	if v23 < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if base.Ui32(v26+int32(16)) < base.Ui32((l2+int32(7))&int32(-8)) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	v26 = v22
	goto L7
L6:
	;
	v26 = v23
	goto L7
L7:
	;
	goto L4
L8:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v36) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(base.Ui32(v36+int32(262120)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v44 = int32(0)
	goto L11
L11:
	;
	v46 = v44 & int32(65535)
	goto L13
L12:
	;
	v143 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v143)
	goto L2
L13:
	;
	if l3 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v111 = v66 & int32(65535)
	F_PageIndexTupleDelete(m, l0, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	if base.Ui32(v46) < base.Ui32(v60) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	if v57 != 0 {
		v60 = v57
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v60 = int32(1)
	goto L15
L19:
	;
	goto L18
L20:
	;
	goto L14
L21:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L30
	}
L22:
	;
	v66 = v60
	goto L23
L23:
	;
	v73 = v66 & int32(65535)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0+v79&int32(32767))))
	v84 = int32(3)
	if v83&v84 != v84 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v73 != 0 {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v89 = v66 + int32(1)
	if base.Ui32(v89&int32(65535)) <= base.Ui32(v46) {
		v66 = v89
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L21
L29:
	;
	goto L21
L30:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	if v105 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v108)
	goto L13
L32:
	;
	return int32(0)
L33:
	;
	v117 = F_PageAddItemExtended(m, l0, l1, l2, v111, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v117 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v121 = v119 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v121)
	if l3 == int32(0) {
		v176 = v117
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v126 = v117 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v126)
	v176 = v117
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	F_errmsg_internal(m, int32(398409), v13+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(482816), int32(1274), int32(284910))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	if v157 != 0 {
		v176 = v157
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(398409), v13)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(482816), int32(1286), int32(284910))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
