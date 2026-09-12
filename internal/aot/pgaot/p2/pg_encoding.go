package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if l3 == int32(0) {
		v121 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L12
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L46
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L43
	}
L5:
	;
	m.G0 = v9 + int32(48)
	return v121
L6:
	;
	if l1 <= int32(0) {
		v121 = l0
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l2 == l3 {
		v121 = l0
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3*int32(28))+uint32(_consts[1222])))
	v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	goto L16
L12:
	;
	return int32(0)
L13:
	;
	if l1 == v23 {
		v121 = l0
		goto L5
	} else {
		goto L14
	}
L14:
	;
	F_report_invalid_encoding(m, l3, l0+v23, l1-v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	if base.B2i32(v34 == int32(2)) == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v39 = F_FindDefaultConversionProc(m, l2, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if v39 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v51 = F_MemoryContextAllocHuge(m, v46, l1<<(uint(int32(2))%32)|int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v54 = F_OidFunctionCall6Coll(m, v39, l2, l3, l0, v51, l1, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(l1) < base.Ui32(int32(1000001)) {
		v121 = v51
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v51&int32(3) == int32(0) {
		v81 = v51
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v114) {
		goto L1
	} else {
		goto L41
	}
L25:
	;
	v114 = v106 - v51
	goto L24
L26:
	;
	v85 = v81
	goto L35
L27:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v65 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = int32(0)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v70 = v51
	goto L31
L31:
	;
	v74 = v70 + int32(1)
	if v74&int32(3) == int32(0) {
		v81 = v74
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v106 = v74
	goto L25
L33:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v79 != 0 {
		v70 = v74
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v94 = int32(-2139062144)
	if (int32(16843008)-v91|v91)&v94 == v94 {
		v85 = v85 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v100 = v85
	goto L38
L37:
	;
	goto L36
L38:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v104 != 0 {
		v100 = v100 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v106 = v100
	goto L25
L40:
	;
	goto L39
L41:
	;
	v119 = F_repalloc(m, v51, v114+int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v121 = v119
	goto L5
L43:
	;
	F_errmsg_internal(m, int32(248599), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(477194), int32(388), int32(261740))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(l2) <= base.Ui32(int32(41)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if base.Ui32(l3) <= base.Ui32(int32(41)) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(3))%32))+uint32(_consts[359])))
	v158 = v157
	goto L51
L50:
	;
	v158 = int32(722455)
	goto L51
L51:
	;
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v158
	F_errmsg(m, int32(69045), v9)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L56
	}
L53:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(3))%32))+uint32(_consts[359])))
	v168 = v167
	goto L55
L54:
	;
	v168 = int32(722455)
	goto L55
L55:
	;
	goto L52
L56:
	;
	F_errfinish(m, int32(477194), int32(396), int32(261740))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(13006), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errdetail(m, int32(583770), v9+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(477194), int32(412), int32(261740))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(13006), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errdetail(m, int32(583770), v9+int32(32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(477194), int32(440), int32(261740))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
