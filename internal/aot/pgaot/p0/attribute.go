package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAttributeCompression(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 == v3 {
		v124 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L14
	} else {
		goto L42
	}
L3:
	;
	m.G0 = v7 + int32(32)
	return v124
L4:
	;
	v11 = int32(97796)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v15 == int32(0) {
		v34 = v14
		v35 = v15
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v35-v34 == int32(0) {
		v124 = v3
		goto L3
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v14 != v15 {
		v34 = v14
		v35 = v15
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v19 = l1
	v20 = v11
	goto L9
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v23
		v35 = v24
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v34 = v23
	v35 = v24
	goto L6
L11:
	;
	v27 = int32(1)
	if v23 == v24 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v39 = F_get_typstorage(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v39 == int32(112) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v46 = int32(7994)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[298])))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v50 == int32(0) {
		v69 = v49
		v70 = v50
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v121 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L18:
	;
	if v70-v69 == int32(0) {
		v121 = int32(112)
		goto L17
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	if v49 != v50 {
		v69 = v49
		v70 = v50
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v54 = l1
	v55 = v46
	goto L22
L22:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v58
		v70 = v59
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v69 = v58
	v70 = v59
	goto L19
L24:
	;
	v62 = int32(1)
	if v58 == v59 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v74 = int32(0)
	v75 = int32(545409)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[299])))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v79 == v74 {
		v98 = v78
		v99 = v79
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v99-v98 != 0 {
		v121 = v74
		goto L17
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	if v78 != v79 {
		v98 = v78
		v99 = v79
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v83 = l1
	v84 = v75
	goto L31
L31:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v87
		v99 = v88
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v98 = v87
	v99 = v88
	goto L28
L33:
	;
	v91 = int32(1)
	if v87 == v88 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(439619), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	F_errdetail(m, int32(558463), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(491425), int32(292), int32(418823))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v124 = v121
	goto L3
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v136 = F_format_type_be(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v136
	F_errmsg(m, int32(268244), v7)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(489510), int32(22067), int32(268392))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	F_errmsg(m, int32(701765), v7+int32(16))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(489510), int32(22073), int32(268392))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
