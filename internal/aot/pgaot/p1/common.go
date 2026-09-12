package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_select_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(1) < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v15 + int32(4)
	goto L3
L2:
	;
	v22 = int32(0)
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = F_exprType(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	m.G0 = v13 + int32(32)
	return v224
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v216
	v224 = v213
	goto L4
L6:
	;
	if l3 == int32(0) {
		v224 = v24
		goto L4
	} else {
		goto L56
	}
L7:
	;
	v70 = F_getBaseType(m, v24)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L18
	}
L8:
	;
	return int32(0)
L9:
	;
	if v24 == int32(705) {
		v67 = v22
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v22 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v35 = (v22 - v32) >> (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 <= v35 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v43 = v35
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = v48 + v43<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = F_exprType(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L6
L15:
	;
	if v53 != v24 {
		v67 = v51
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v57 = v43 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 < v58 {
		v43 = v57
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v70
	F_get_type_category_preferred(m, v70, v13+int32(27), v13+int32(26))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v67 == int32(0) {
		v186 = v70
		v187 = v23
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v224 = int32(0)
	goto L4
L21:
	;
	if v186 == int32(705) {
		goto L52
	} else {
		goto L53
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v84 = (v67 - v81) >> (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v85 <= v84 {
		v186 = v70
		v187 = v23
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v93 = v84
	v95 = v70
	v96 = v23
	goto L24
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v93<<(uint(int32(2))%32))))
	v102 = F_exprType(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L26
	}
L25:
	;
	v186 = v172
	v187 = v173
	goto L21
L26:
	;
	v104 = F_getBaseType(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v104
	if v104 == int32(705) {
		v172 = v95
		v173 = v96
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v175 = v93 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v175 < v176 {
		v93 = v175
		v95 = v172
		v96 = v173
		goto L24
	} else {
		goto L51
	}
L29:
	;
	if v104 == v95 {
		v172 = v95
		v173 = v96
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_get_type_category_preferred(m, v104, v13+int32(19), v13+int32(18))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	if v95 != int32(705) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+19)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	if v118 != v119 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v104
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)) = uint8(v168)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v170)
	v172 = v104
	v173 = v101
	goto L28
L35:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)))
	if v148 != 0 {
		v172 = v95
		v173 = v96
		goto L28
	} else {
		goto L46
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v130 = F_format_type_be(m, v95)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v132 = F_format_type_be(m, v104)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(459307), v13)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v140 = F_exprLocation(m, v101)
	mBase = m.M
	F_parser_errposition(m, l0, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(499785), int32(1422), int32(366759))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
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
	v155 = F_can_coerce_type(m, int32(1), v13+int32(28), v13+int32(20), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	if v155 == int32(0) {
		v172 = v95
		v173 = v96
		goto L28
	} else {
		goto L48
	}
L48:
	;
	v165 = F_can_coerce_type(m, int32(1), v13+int32(20), v13+int32(28), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	if v165 != 0 {
		v172 = v95
		v173 = v96
		goto L28
	} else {
		goto L50
	}
L50:
	;
	goto L34
L51:
	;
	goto L25
L52:
	;
	v191 = int32(25)
	goto L54
L53:
	;
	v191 = v186
	goto L54
L54:
	;
	if l3 == int32(0) {
		v224 = v191
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v213 = v191
	v216 = v187
	goto L5
L56:
	;
	v213 = v24
	v216 = v23
	goto L5
}
