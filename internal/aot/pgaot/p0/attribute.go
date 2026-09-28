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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 == v3 {
		v139 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L13
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L42
	}
L3:
	;
	m.G0 = v7 + int32(32)
	return v139
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
		v139 = v3
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
	v46 = m.G0
	v48 = v46 - int32(32)
	m.G0 = v48
	v50 = int32(_a_F_GetAttributeCompression_1)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetAttributeCompression[1])))
	if base.B2i32(v53 == int32(0))|base.B2i32(v53 != v56) != 0 {
		v74 = v53
		v75 = v56
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v108 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L36
	}
L18:
	;
	if v74-v75 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	v59 = l1
	v60 = v50
	goto L21
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v64 == int32(0) {
		v74 = v64
		v75 = v63
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v74 = v64
	v75 = v63
	goto L19
L23:
	;
	v67 = int32(1)
	if v64 == v63 {
		v59 = v59 + v67
		v60 = v60 + v67
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v77 = int32(_a_F_GetAttributeCompression_2)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetAttributeCompression[2])))
	if base.B2i32(v80 == int32(0))|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v108 = int32(112)
	goto L27
L27:
	;
	m.G0 = v48 + int32(32)
	goto L16
L28:
	;
	if v101-v102 == int32(0) {
		goto L17
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v86 = l1
	v87 = v77
	goto L31
L31:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v101 = v91
	v102 = v90
	goto L29
L33:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v108 = int32(0)
	goto L27
L36:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(_a_F_GetAttributeCompression_2)
	F_errmsg(m, int32(_a_F_GetAttributeCompression_3), v48+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(_a_F_GetAttributeCompression_2)
	v129 = F_errdetail(m, int32(_a_F_GetAttributeCompression_4), v48)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_5), int32(292), int32(_a_F_GetAttributeCompression_6))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
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
	v139 = v108
	goto L3
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v151 = F_format_type_be(m, l0)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v151
	F_errmsg(m, int32(_a_F_GetAttributeCompression_7), v7)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_8), int32(_a_F_GetAttributeCompression_9), int32(_a_F_GetAttributeCompression_10))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
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
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	F_errmsg(m, int32(_a_F_GetAttributeCompression_11), v7+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_GetAttributeCompression_8), int32(_a_F_GetAttributeCompression_12), int32(_a_F_GetAttributeCompression_10))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
