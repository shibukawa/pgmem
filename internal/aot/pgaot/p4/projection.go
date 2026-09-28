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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = F_palloc0(m, int32(88))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(390)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(386)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = l2
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v34
	v38 = F_expr_setup_walker(m, l0, v17)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = v20 + int32(8)
	F_ExecPushExprSetupSteps(m, v41, v17)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v216 = v6
		v220 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v223 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L6:
	;
	v46 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 <= v46 {
		v216 = v6
		v220 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = v46
	v58 = int32(0)
	v66 = v6
	goto L8
L8:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v58<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v74 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v216 = v201
	v220 = v183
	goto L5
L10:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v185 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v184 + v185
	v190 = v178 + v184*int32(40)
	v191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+4)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+24)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v190)+20)) = v183
	v201 = base.I32_extend16_s(v181) - v185
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v190)+32)) = v191
	v206 = v58 + v185
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v206 < v207 {
		v57 = v179
		v58 = v206
		v66 = v183
		goto L8
	} else {
		goto L51
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v176
	v178 = v176
	v179 = v115
	v181 = v118
	v183 = v117
	goto L10
L12:
	;
	F_ExecInitExprRec(m, v74, v41, v20+int32(16), v20+int32(13))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v77 != int32(6) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+8)))
	if v80 <= int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if l4 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v83 < v80 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	switch v99 + int32(2) {
	case 0:
		goto L24
	case 1:
		v115 = int32(18)
		goto L22
	default:
		goto L23
	}
L19:
	;
	v90 = l4 + v83<<(uint(int32(3))%32) + v80*int32(100)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+19)))
	if v91 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90-int32(72))+68))
	if v92 != v95 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v117 = v80 - int32(1)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+8)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v119 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	switch v103 {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	default:
		v115 = v57
		goto L22
	}
L24:
	;
	v115 = int32(19)
	goto L22
L25:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)))
	v112 = v110 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)) = uint8(v112)
	v115 = int32(22)
	goto L22
L26:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)))
	v107 = v105 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)) = uint8(v107)
	v115 = int32(21)
	goto L22
L27:
	;
	v115 = int32(20)
	goto L22
L28:
	;
	v122 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v122
	v126 = F_palloc_mul(m, int32(40), v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v128 != v119 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v176 = v126
	goto L11
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v178 = v130
	v179 = v115
	v181 = v118
	v183 = v117
	goto L10
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v119 << (uint(int32(1)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v137 = F_repalloc(m, v134, v119*int32(80))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v176 = v137
	goto L11
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v146 = F_exprType(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v148 = F_get_typlen(m, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v148 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v152 = int32(24)
	goto L41
L40:
	;
	v152 = int32(23)
	goto L41
L41:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+8)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v154 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v174
	v178 = v174
	v179 = v152
	v181 = v153
	v183 = v66
	goto L10
L43:
	;
	v157 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v157
	v161 = F_palloc_mul(m, int32(40), v157)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v163 != v154 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v174 = v161
	goto L42
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v178 = v165
	v179 = v152
	v181 = v153
	v183 = v66
	goto L10
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v154 << (uint(int32(1)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v172 = F_repalloc(m, v169, v154*int32(80))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v174 = v172
	goto L42
L51:
	;
	goto L9
L52:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v247 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v246 + v247
	v252 = v245 + v246*int32(40)
	v253 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v252)+4)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v252)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v252)+24)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v252)+16)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v252)+32)) = v253
	v265 = F_jit_compile_expr(m, v41)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L62
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v243
	v245 = v243
	goto L52
L54:
	;
	v226 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v226
	v230 = F_palloc_mul(m, int32(40), v226)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v232 != v223 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v243 = v230
	goto L53
L58:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v245 = v234
	goto L52
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v223 << (uint(int32(1)) % 32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v241 = F_repalloc(m, v238, v223*int32(80))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v243 = v241
	goto L53
L62:
	;
	if v265 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_ExecReadyInterpretedExpr(m, v41)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	m.G0 = v17 + int32(16)
	return v20
L66:
	;
	goto L65
}
