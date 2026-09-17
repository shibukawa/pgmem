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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 == v3 {
		v127 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L44
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L39
	}
L3:
	;
	m.G0 = v7 + int32(32)
	return v127
L4:
	;
	v11 = int32(_a_F_GetAttributeCompression_0)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetAttributeCompression[0])))
	if base.B2i32(v14 == int32(0))|base.B2i32(v14 != v17) != 0 {
		v35 = v14
		v36 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v35-v36 == int32(0) {
		v127 = v3
		goto L3
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v20 = l1
	v21 = v11
	goto L8
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v25
		v36 = v24
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v35 = v25
	v36 = v24
	goto L6
L10:
	;
	v28 = int32(1)
	if v25 == v24 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v40 = F_get_typstorage(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v40 == int32(112) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(_a_F_GetAttributeCompression_1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetAttributeCompression[1])))
	if base.B2i32(v50 == int32(0))|base.B2i32(v50 != v53) != 0 {
		v71 = v50
		v72 = v53
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v124 == int32(0) {
		goto L1
	} else {
		goto L38
	}
L17:
	;
	if v71-v72 == int32(0) {
		v124 = int32(112)
		goto L16
	} else {
		goto L24
	}
L18:
	;
	goto L17
L19:
	;
	v56 = l1
	v57 = v47
	goto L20
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v61 == int32(0) {
		v71 = v61
		v72 = v60
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v71 = v61
	v72 = v60
	goto L18
L22:
	;
	v64 = int32(1)
	if v61 == v60 {
		v56 = v56 + v64
		v57 = v57 + v64
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v76 = int32(0)
	v77 = int32(_a_F_GetAttributeCompression_2)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetAttributeCompression[2])))
	if base.B2i32(v80 == v76)|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v101-v102 != 0 {
		v124 = v76
		goto L16
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v86 = l1
	v87 = v77
	goto L28
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v101 = v91
	v102 = v90
	goto L26
L30:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_GetAttributeCompression_3), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_errdetail(m, int32(_a_F_GetAttributeCompression_4), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_5), int32(292), int32(_a_F_GetAttributeCompression_6))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v127 = v124
	goto L3
L39:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v139 = F_format_type_be(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v139
	F_errmsg(m, int32(_a_F_GetAttributeCompression_7), v7)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_8), int32(_a_F_GetAttributeCompression_9), int32(_a_F_GetAttributeCompression_10))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	F_errmsg(m, int32(_a_F_GetAttributeCompression_11), v7+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_8), int32(_a_F_GetAttributeCompression_12), int32(_a_F_GetAttributeCompression_10))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
