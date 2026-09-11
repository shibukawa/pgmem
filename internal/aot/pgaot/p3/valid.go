package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_extension_name(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
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
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	if l0&int32(3) == int32(0) {
		v31 = l0
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L36
	} else {
		goto L52
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L36
	} else {
		goto L47
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L36
	} else {
		goto L42
	}
L4:
	;
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v64 = v56 - l0
	goto L4
L6:
	;
	v35 = v31
	goto L15
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v64 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v20 = l0
	goto L11
L11:
	;
	v24 = v20 + int32(1)
	if v24&int32(3) == int32(0) {
		v31 = v24
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v56 = v24
	goto L5
L13:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != 0 {
		v20 = v24
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v44 = int32(-2139062144)
	if (int32(16843008)-v41|v41)&v44 == v44 {
		v35 = v35 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v50 = v35
	goto L18
L17:
	;
	goto L16
L18:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		v50 = v50 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v56 = v50
	goto L5
L20:
	;
	goto L19
L21:
	;
	v66 = F_strstr(m, l0, int32(_a_F_check_valid_extension_name_0))
	mBase = m.M
	if v66 != 0 {
		goto L3
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v67 == int32(45) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v64-int32(1)))))
	if v73 == int32(45) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v77 = l0
	goto L28
L27:
	;
	if v87 != 0 {
		goto L1
	} else {
		goto L35
	}
L28:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	v87 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v79 == int32(47) {
		v87 = v77
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v77 = v77 + int32(1)
	goto L28
L35:
	;
	m.G0 = v6 - int32(-64)
	return
L36:
	;
	return
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_check_valid_extension_name_1), v6)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errdetail(m, int32(_a_F_check_valid_extension_name_2), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_check_valid_extension_name_3), int32(372), int32(_a_F_check_valid_extension_name_4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L36
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
	F_errmsg(m, int32(_a_F_check_valid_extension_name_1), v4+int32(-16))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	F_errdetail(m, int32(_a_F_check_valid_extension_name_5), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_check_valid_extension_name_3), int32(381), int32(_a_F_check_valid_extension_name_4))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L36
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L36
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg(m, int32(_a_F_check_valid_extension_name_1), v4+int32(-48))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	F_errdetail(m, int32(_a_F_check_valid_extension_name_6), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_check_valid_extension_name_3), int32(393), int32(_a_F_check_valid_extension_name_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L36
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	F_errmsg(m, int32(_a_F_check_valid_extension_name_1), v4+int32(-32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	F_errdetail(m, int32(_a_F_check_valid_extension_name_7), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L36
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_check_valid_extension_name_3), int32(403), int32(_a_F_check_valid_extension_name_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L36
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
