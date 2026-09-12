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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_ArrayGetNItems(m, v9, l0+int32(16))
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
	v19 = F_ArrayGetNItems(m, v16, l1+int32(16))
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
	if v12 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(0)
	v51 = v49
	v52 = v49
	goto L13
L13:
	;
	v59 = int32(2)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+v31+v52<<(uint(v59)%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+v42+v51<<(uint(v59)%32))))
	if v62 < v66 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L10
L15:
	;
	if v12 <= v76 {
		goto L10
	} else {
		goto L22
	}
L16:
	;
	v75 = v51
	v76 = v52 + int32(1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v62 == v66 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(1)
L20:
	;
	goto L21
L21:
	;
	v75 = v51 + int32(1)
	v76 = v52
	goto L15
L22:
	;
	if v75 < v19 {
		v51 = v75
		v52 = v76
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L14
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
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
	if l1 == int32(0) {
		v233 = v4
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v233
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v233 = v4
	goto L4
L7:
	;
	goto L8
L8:
	;
	v26 = v4
	v27 = v15
	v28 = v4
	v30 = v4
	v31 = v4
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L42
	} else {
		goto L58
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v30<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L42
	} else {
		goto L55
	}
L12:
	;
	v44 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	if v70 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v44
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	if v55 != 0 {
		v44 = v55
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	goto L16
L18:
	;
	v74 = v70
	goto L21
L19:
	;
	goto L20
L20:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+120)))
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v74
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+56))
	if v85 != 0 {
		v74 = v85
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	goto L22
L24:
	;
	v102 = int32(104)
	goto L26
L25:
	;
	v102 = int32(100)
	goto L26
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38+v102)))
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L11
L28:
	;
	v107 = int32(100)
	goto L30
L29:
	;
	v107 = int32(104)
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v38+v107)))
	if v31 != v109 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v27 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	v127 = v27
	v128 = v28
	v129 = v31
	goto L33
L33:
	;
	if v109 != v104 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v109 != v114 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v117 = v27 + int32(4)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v117) < base.Ui32(v119+v120<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v117
	goto L38
L37:
	;
	v125 = int32(0)
	goto L38
L38:
	;
	v127 = v125
	v128 = v113
	v129 = v109
	goto L33
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+16)))
	v134 = F_make_canonical_pathkey(m, l0, v104, v131, v132, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v138 = v128
	goto L41
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+40)))
	if v140 != 0 {
		v190 = v26
		goto L44
	} else {
		goto L45
	}
L42:
	;
	return int32(0)
L43:
	;
	v138 = v134
	goto L41
L44:
	;
	v199 = v30 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v200 <= v199 {
		v233 = v190
		goto L4
	} else {
		goto L54
	}
L45:
	;
	if v26 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v183 = F_lappend(m, v26, v138)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L42
	} else {
		goto L53
	}
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v143 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v151 = int32(0)
	goto L49
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v146+v151<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v139 == v165 {
		v190 = v26
		goto L44
	} else {
		goto L51
	}
L50:
	;
	goto L46
L51:
	;
	v168 = v151 + int32(1)
	if v143 != v168 {
		v151 = v168
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v190 = v183
	goto L44
L54:
	;
	v26 = v190
	v27 = v127
	v28 = v128
	v30 = v199
	v31 = v129
	goto L10
L55:
	;
	F_errmsg_internal(m, int32(169412), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L42
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(512995), int32(1897), int32(416058))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L42
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
	F_errmsg_internal(m, int32(372910), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L42
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(512995), int32(1902), int32(416058))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L42
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
