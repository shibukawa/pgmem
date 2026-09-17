package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_select_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = base.B2i32(v17 < int32(2))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = v20 + int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v24 = F_exprType(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v224
L2:
	;
	if l3 == int32(0) {
		v224 = v209
		goto L1
	} else {
		goto L55
	}
L3:
	;
	v74 = F_getBaseType(m, v24)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L19
	}
L4:
	;
	return int32(0)
L5:
	;
	if v24 == int32(705) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v17 < int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v17 < int32(2) {
		v209 = v24
		v211 = v23
		goto L2
	} else {
		goto L12
	}
L9:
	;
	v31 = int32(0)
	goto L11
L10:
	;
	v31 = v22
	goto L11
L11:
	;
	v67 = v31
	goto L3
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v35 = (v22 - v32) >> (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 <= v35 {
		v209 = v24
		v211 = v23
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v46 = v35
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = v50 + v46<<(uint(int32(2))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = F_exprType(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	v209 = v24
	v211 = v23
	goto L2
L16:
	;
	if v55 != v24 {
		v67 = v53
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v59 = v46 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v59 < v60 {
		v46 = v59
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v74
	F_get_type_category_preferred(m, v74, v15+int32(27), v15+int32(26))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v67 == int32(0) {
		v191 = v74
		v194 = v23
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v224 = int32(0)
	goto L1
L22:
	;
	if v191 == int32(705) {
		goto L52
	} else {
		goto L53
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v88 = (v67 - v85) >> (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v89 <= v88 {
		v191 = v74
		v194 = v23
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v97 = v74
	v98 = v88
	v100 = v23
	goto L25
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v98<<(uint(int32(2))%32))))
	v108 = F_exprType(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	v191 = v177
	v194 = v178
	goto L22
L27:
	;
	v110 = F_getBaseType(m, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v110
	if base.B2i32(v110 == int32(705))|base.B2i32(v110 == v97) != 0 {
		v177 = v97
		v178 = v100
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v182 = v98 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v182 < v183 {
		v97 = v177
		v98 = v182
		v100 = v178
		goto L25
	} else {
		goto L51
	}
L30:
	;
	F_get_type_category_preferred(m, v110, v15+int32(19), v15+int32(18))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v97 != int32(705) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	if v125 != v126 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v110
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v173)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)) = uint8(v175)
	v177 = v110
	v178 = v107
	goto L29
L35:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v155 != 0 {
		v177 = v97
		v178 = v100
		goto L29
	} else {
		goto L46
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v137 = F_format_type_be(m, v97)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v139 = F_format_type_be(m, v110)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	F_errmsg(m, int32(_a_F_select_common_type_0), v15)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v147 = F_exprLocation(m, v107)
	mBase = m.M
	F_parser_errposition(m, l0, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_select_common_type_1), int32(1422), int32(_a_F_select_common_type_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
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
	v158 = v15 + int32(28)
	v160 = v15 + int32(20)
	v162 = F_can_coerce_type(m, int32(1), v158, v160, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v162 == int32(0) {
		v177 = v97
		v178 = v100
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v168 = F_can_coerce_type(m, int32(1), v160, v158, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v168 != 0 {
		v177 = v97
		v178 = v100
		goto L29
	} else {
		goto L50
	}
L50:
	;
	goto L34
L51:
	;
	goto L26
L52:
	;
	v200 = int32(25)
	goto L54
L53:
	;
	v200 = v191
	goto L54
L54:
	;
	v209 = v200
	v211 = v194
	goto L2
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v211
	v224 = v209
	goto L1
}
