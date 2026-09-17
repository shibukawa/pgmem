package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_overlap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_ArrayGetNItemsSafe(m, v9, l0+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = F_ArrayGetNItemsSafe(m, v16, l1+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = (v24<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L6
L5:
	;
	v31 = v21
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v32 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = (v35<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L9
L8:
	;
	v42 = v32
	goto L9
L9:
	;
	v43 = int32(0)
	if base.B2i32(v12 <= v43)|base.B2i32(v19 <= v43) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v50 = int32(0)
	v52 = v50
	v53 = v50
	goto L12
L12:
	;
	v60 = int32(2)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0+v31+v53<<(uint(v60)%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1+v42+v52<<(uint(v60)%32))))
	if v63 < v67 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L10
L14:
	;
	if v12 <= v77 {
		goto L10
	} else {
		goto L21
	}
L15:
	;
	v76 = v52
	v77 = v53 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v63 == v67 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	v76 = v52 + int32(1)
	v77 = v53
	goto L14
L21:
	;
	if v76 < v19 {
		v52 = v76
		v53 = v77
		goto L12
	} else {
		goto L22
	}
L22:
	;
	goto L13
}
func F_make_inner_pathkeys_for_merge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v4 = int32(0)
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v15 = v14
	goto L3
L2:
	;
	v15 = v4
	goto L3
L3:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L43
	} else {
		goto L59
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L43
	} else {
		goto L56
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v208 = v4
	goto L8
L8:
	;
	return v208
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v26 = v15
	v27 = v4
	v29 = v4
	v30 = v4
	v31 = v4
	goto L12
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v29<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v208 = v191
	goto L8
L14:
	;
	v44 = v40
	goto L17
L15:
	;
	goto L16
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v44
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	if v55 != 0 {
		v44 = v55
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	goto L18
L20:
	;
	v74 = v70
	goto L23
L21:
	;
	goto L22
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+120)))
	if v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v74
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+56))
	if v85 != 0 {
		v74 = v85
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	goto L24
L26:
	;
	v102 = int32(104)
	goto L28
L27:
	;
	v102 = int32(100)
	goto L28
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38+v102)))
	if v101 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v107 = int32(100)
	goto L31
L30:
	;
	v107 = int32(104)
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v38+v107)))
	if v30 != v109 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	v127 = v26
	v128 = v30
	v129 = v31
	goto L34
L34:
	;
	if v109 != v104 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v109 != v114 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v117 = v26 + int32(4)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v117) < base.Ui32(v119+v120<<(uint(int32(2))%32)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v125 = v117
	goto L39
L38:
	;
	v125 = int32(0)
	goto L39
L39:
	;
	v127 = v125
	v128 = v109
	v129 = v113
	goto L34
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+16)))
	v134 = F_make_canonical_pathkey(m, l0, v104, v131, v132, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v138 = v129
	goto L42
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+40)))
	if v140 != 0 {
		v191 = v27
		goto L45
	} else {
		goto L46
	}
L43:
	;
	return int32(0)
L44:
	;
	v138 = v134
	goto L42
L45:
	;
	v199 = v29 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v199 < v200 {
		v26 = v127
		v27 = v191
		v29 = v199
		v30 = v128
		v31 = v129
		goto L12
	} else {
		goto L55
	}
L46:
	;
	if v27 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v183 = F_lappend(m, v27, v138)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L43
	} else {
		goto L54
	}
L48:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v143 <= int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v151 = int32(0)
	goto L50
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v146+v151<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v139 == v165 {
		v191 = v27
		goto L45
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	v168 = v151 + int32(1)
	if v143 != v168 {
		v151 = v168
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v191 = v183
	goto L45
L55:
	;
	goto L13
L56:
	;
	F_errmsg_internal(m, int32(_a_F_make_inner_pathkeys_for_merge_0), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_make_inner_pathkeys_for_merge_1), int32(1897), int32(_a_F_make_inner_pathkeys_for_merge_2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L43
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errmsg_internal(m, int32(_a_F_make_inner_pathkeys_for_merge_3), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_make_inner_pathkeys_for_merge_1), int32(1902), int32(_a_F_make_inner_pathkeys_for_merge_2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L43
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
