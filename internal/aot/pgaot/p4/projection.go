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
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = F_palloc0(m, int32(76))
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(1632087572864)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l2
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v32
	v36 = F_expr_setup_walker(m, l0, v17)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = v20 + int32(4)
	F_ExecPushExprSetupSteps(m, v39, v17)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v212 = v6
		v216 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v219 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L6:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 <= v44 {
		v212 = v6
		v216 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v55 = v44
	v56 = int32(0)
	v64 = v6
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v56<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v212 = v197
	v216 = v179
	goto L5
L10:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v181 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v180 + v181
	v186 = v174 + v180*int32(40)
	v187 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+4)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v179
	v197 = base.I32_extend16_s(v177) - v181
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v186)+32)) = v187
	v202 = v56 + v181
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v202 < v203 {
		v55 = v175
		v56 = v202
		v64 = v179
		goto L8
	} else {
		goto L51
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v172
	v174 = v172
	v175 = v113
	v177 = v116
	v179 = v115
	goto L10
L12:
	;
	F_ExecInitExprRec(m, v72, v39, v20+int32(12), v20+int32(9))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v75 != int32(6) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+8)))
	if v78 <= int32(0) {
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v81 < v78 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	switch v97 + int32(2) {
	case 0:
		goto L24
	case 1:
		v113 = int32(18)
		goto L22
	default:
		goto L23
	}
L19:
	;
	v88 = l4 + v81<<(uint(int32(4))%32) + v78*int32(100)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+11)))
	if v89 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(80))+68))
	if v90 != v93 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v115 = v78 - int32(1)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+8)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v117 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
	switch v101 {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	default:
		v113 = v55
		goto L22
	}
L24:
	;
	v113 = int32(19)
	goto L22
L25:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
	v110 = v108 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v110)
	v113 = int32(22)
	goto L22
L26:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
	v105 = v103 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v105)
	v113 = int32(21)
	goto L22
L27:
	;
	v113 = int32(20)
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(16)
	v123 = F_palloc(m, int32(640))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v125 != v117 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v172 = v123
	goto L11
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v174 = v127
	v175 = v113
	v177 = v116
	v179 = v115
	goto L10
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v117 << (uint(int32(1)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v134 = F_repalloc(m, v131, v117*int32(80))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v172 = v134
	goto L11
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v143 = F_exprType(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v145 = F_get_typlen(m, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v145 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v149 = int32(24)
	goto L41
L40:
	;
	v149 = int32(23)
	goto L41
L41:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+8)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v151 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v170
	v174 = v170
	v175 = v149
	v177 = v150
	v179 = v64
	goto L10
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(16)
	v157 = F_palloc(m, int32(640))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v159 != v151 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v170 = v157
	goto L42
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v174 = v161
	v175 = v149
	v177 = v150
	v179 = v64
	goto L10
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v151 << (uint(int32(1)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v168 = F_repalloc(m, v165, v151*int32(80))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v170 = v168
	goto L42
L51:
	;
	goto L9
L52:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v242 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v241 + v242
	v247 = v240 + v241*int32(40)
	v248 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v247)+4)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v247)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v247)+24)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v247)+20)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v247)+16)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v247)+32)) = v248
	v260 = F_jit_compile_expr(m, v39)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L62
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v238
	v240 = v238
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(16)
	v225 = F_palloc(m, int32(640))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v227 != v219 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v238 = v225
	goto L53
L58:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v240 = v229
	goto L52
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v219 << (uint(int32(1)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v236 = F_repalloc(m, v233, v219*int32(80))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v238 = v236
	goto L53
L62:
	;
	if v260 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_ExecReadyInterpretedExpr(m, v39)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
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
