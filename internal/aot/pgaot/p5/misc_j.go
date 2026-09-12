package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_jointree_contains_lateral_outer_refs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v5 {
		v201 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L57
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return v201
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v15 - int32(63) {
	case 0:
		v201 = v5
		goto L2
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L1
	}
L4:
	;
	v201 = int32(1)
	goto L2
L5:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = int32(0)
	v118 = l2 | base.B2i32(v115 != v116)
	if v115 != 0 {
		goto L33
	} else {
		goto L34
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l2 == int32(0) {
		v201 = v5
		goto L2
	} else {
		goto L16
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = v5
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v28<<(uint(int32(2))%32))))
	v37 = F_jointree_contains_lateral_outer_refs(m, l0, v36, l2, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	if v37 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v42 = v28 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v42 < v43 {
		v28 = v42
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = F_pull_varnos_of_level(m, l0, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(0)
	if v56 == v58 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v111 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L19:
	;
	v111 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if l3 == int32(0) {
		v102 = v58
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v111 = v102
	goto L18
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v68 < v67 {
		v102 = v58
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(1)
	if v67 <= v70 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v73 = v70
	goto L27
L26:
	;
	v73 = v67
	goto L27
L27:
	;
	v74 = int32(8)
	v79 = int32(0)
	goto L28
L28:
	;
	v86 = v79 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v56+v74+v86)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+(l3+v74))))
	v93 = v88 & (v90 ^ int32(-1))
	v95 = base.B2i32(v93 == int32(0))
	if v93 != 0 {
		v102 = v95
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v102 = v95
	goto L22
L30:
	;
	v97 = v79 + int32(1)
	if v97 != v73 {
		v79 = v97
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v201 = v5
	goto L2
L33:
	;
	v120 = v116
	goto L35
L34:
	;
	v120 = l3
	goto L35
L35:
	;
	v121 = F_jointree_contains_lateral_outer_refs(m, l0, v114, v118, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	if v121 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v124 = F_jointree_contains_lateral_outer_refs(m, l0, v123, v118, v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	if v124 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v118 == int32(0) {
		v201 = v5
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v129 = F_pull_varnos_of_level(m, l0, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v131 = int32(0)
	if v129 == v131 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v184 != 0 {
		v201 = v5
		goto L2
	} else {
		goto L56
	}
L43:
	;
	v184 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	if v120 == int32(0) {
		v175 = v131
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v184 = v175
	goto L42
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v141 < v140 {
		v175 = v131
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v143 = int32(1)
	if v140 <= v143 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v146 = v143
	goto L51
L50:
	;
	v146 = v140
	goto L51
L51:
	;
	v147 = int32(8)
	v152 = int32(0)
	goto L52
L52:
	;
	v159 = v152 << (uint(int32(2)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v129+v147+v159)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+(v120+v147))))
	v166 = v161 & (v163 ^ int32(-1))
	v168 = base.B2i32(v166 == int32(0))
	if v166 != 0 {
		v175 = v168
		goto L46
	} else {
		goto L54
	}
L53:
	;
	v175 = v168
	goto L46
L54:
	;
	v170 = v152 + int32(1)
	if v170 != v146 {
		v152 = v170
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L4
L57:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v210
	F_errmsg_internal(m, int32(485936), v11)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(499340), int32(2397), int32(156913))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonpath_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_jsonPathFromCstring(m, v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
