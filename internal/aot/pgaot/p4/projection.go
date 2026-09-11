package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAssignScanProjectionInfoWithVarno(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	F_ExecConditionalAssignProjectionInfo(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_ExecBuildProjectionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = F_palloc0(m, int32(76))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(1632087572864)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v33
	v37 = F_expr_setup_walker(m, l0, v18)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = v21 + int32(4)
	F_ExecPushExprSetupSteps(m, v40, v18)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v214 = v6
		v216 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v222 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L6:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v46 <= v45 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v214 = v6
	v216 = v6
	goto L5
L8:
	;
	goto L9
L9:
	;
	v58 = v45
	v59 = int32(0)
	v65 = v6
	goto L10
L10:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v59<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v76 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v214 = v197
	v216 = v180
	goto L5
L12:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v182 + v183
	v188 = v176 + v182*int32(40)
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+4)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v188)+24)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = v180
	v197 = base.I32_extend16_s(v179) - v183
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+32)) = v189
	v204 = v59 + v183
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v204 < v205 {
		v58 = v177
		v59 = v204
		v65 = v180
		goto L10
	} else {
		goto L53
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v174
	v176 = v174
	v177 = v115
	v179 = v118
	v180 = v117
	goto L12
L14:
	;
	F_ExecInitExprRec(m, v76, v40, v21+int32(12), v21+int32(9))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L38
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v79 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+8)))
	if v82 <= int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if l4 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v85 < v82 {
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	switch v99 + int32(2) {
	case 0:
		goto L26
	case 1:
		v115 = int32(18)
		goto L24
	default:
		goto L25
	}
L21:
	;
	v92 = l4 - int32(80) + v85<<(uint(int32(4))%32) + v82*int32(100)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+91)))
	if v93 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+68))
	if v94 != v95 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v117 = v82 - int32(1)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+8)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v119 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v76)+32))
	switch v103 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		v115 = v58
		goto L24
	}
L26:
	;
	v115 = int32(19)
	goto L24
L27:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
	v112 = v110 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v112)
	v115 = int32(22)
	goto L24
L28:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
	v107 = v105 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v107)
	v115 = int32(21)
	goto L24
L29:
	;
	v115 = int32(20)
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = int32(16)
	v125 = F_palloc(m, int32(640))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v127 != v119 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v174 = v125
	goto L13
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v176 = v129
	v177 = v115
	v179 = v118
	v180 = v117
	goto L12
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v119 << (uint(int32(1)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v136 = F_repalloc(m, v133, v119*int32(80))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v174 = v136
	goto L13
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v145 = F_exprType(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v147 = F_get_typlen(m, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v147 == int32(-1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v151 = int32(24)
	goto L43
L42:
	;
	v151 = int32(23)
	goto L43
L43:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+8)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v153 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v172
	v176 = v172
	v177 = v151
	v179 = v152
	v180 = v65
	goto L12
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = int32(16)
	v159 = F_palloc(m, int32(640))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v161 != v153 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v172 = v159
	goto L44
L49:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v176 = v163
	v177 = v151
	v179 = v152
	v180 = v65
	goto L12
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v153 << (uint(int32(1)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v170 = F_repalloc(m, v167, v153*int32(80))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v172 = v170
	goto L44
L53:
	;
	goto L11
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v245 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v244 + v245
	v250 = v243 + v244*int32(40)
	v251 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v245
	*(*int64)(unsafe.Add(mBase, uint32(v250)+24)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v250)+20)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v250)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v250)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+32)) = v251
	v263 = F_jit_compile_expr(m, v40)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L64
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v241
	v243 = v241
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = int32(16)
	v228 = F_palloc(m, int32(640))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v230 != v222 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v241 = v228
	goto L55
L60:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v243 = v232
	goto L54
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v222 << (uint(int32(1)) % 32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v239 = F_repalloc(m, v236, v222*int32(80))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v241 = v239
	goto L55
L64:
	;
	if v263 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_ExecReadyInterpretedExpr(m, v40)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	m.G0 = v18 + int32(16)
	return v21
L68:
	;
	goto L67
}
func F_prepare_projection_slot(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = int32(2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+l2<<(uint(v10)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v15&v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ExecStoreAllNullTuple(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	return
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v25 < v24 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_slot_getsomeattrs_int(m, l1, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v32 = v20
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v33 <= int32(0) {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v29 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v32 = v29
	goto L11
L14:
	;
	v37 = int32(0)
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v37<<(uint(int32(2))%32))))
	v47 = F_bms_is_member(m, v46, v13)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L1
L17:
	;
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v53 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v46-v53))) = uint8(v53)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v58 = v37 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v58 < v59 {
		v37 = v58
		goto L15
	} else {
		goto L21
	}
L21:
	;
	goto L16
}
