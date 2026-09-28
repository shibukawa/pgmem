package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CollationCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	v14 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(288)
	m.G0 = v22
	v25 = base.I64_extend_i32_u(l0)
	v26 = base.I64_extend_i32_s(l5)
	v27 = base.I64_extend_i32_u(l1)
	v29 = F_GetSysCacheOid(m, int32(15), v25, v26, v27, int64(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(288)
	return v294
L2:
	;
	return int32(0)
L3:
	;
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l12 != 0 {
		v294 = v14
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v124 = F_table_open(m, int32(3456), int32(6))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L39
	}
L7:
	;
	if l11 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = int32(3456)
	F_checkMembershipInCurrentExtension(m, v22+int32(100))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L26
	}
L11:
	;
	v44 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v44 == int32(0) {
		v294 = v14
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if l5 == int32(-1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(104), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L25
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_1), v22+int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if base.B2i32(l5 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l5)) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_5), v22+int32(48))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v70 = int32(_a_F_CollationCreate_4)
	goto L23
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_c_F_CollationCreate[0])))
	v70 = v69
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L15
L25:
	;
	v294 = v14
	goto L1
L26:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if l5 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(114), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L38
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_6), v22-int32(-64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.B2i32(l5 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l5)) != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L28
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_7), v22+int32(80))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v109 = int32(_a_F_CollationCreate_4)
	goto L36
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_c_F_CollationCreate[0])))
	v109 = v108
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L28
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	if l5 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CollationCreate[1]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	goto L43
L41:
	;
	v134 = int64(-1)
	goto L42
L42:
	;
	v136 = F_GetSysCacheOid(m, int32(15), v25, v134, v27, int64(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v134 = base.I64_extend_i32_s(v131)
	goto L42
L44:
	;
	if v136 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if l12 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v124)+52))
	v190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v22)+176)) = int64(0)
	v195 = v22 + int32(112)
	v197 = F_strncpy(m, v195, l0, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+63)) = uint8(v190)
	goto L66
L48:
	;
	F_relation_close(m, v124, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if l11 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v294 = v14
	goto L1
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = int32(3456)
	F_checkMembershipInCurrentExtension(m, v22+int32(100))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L62
	}
L55:
	;
	F_relation_close(m, v124, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v155 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v155 == int32(0) {
		v294 = v14
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_1), v22)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(160), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v294 = v14
	goto L1
L62:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_6), v22+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(167), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v202 = F_GetNewOidWithIndex(m, v124, int32(3085), int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+224)) = base.I64_extend_i32_s(l3)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+216)) = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+208)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v22)+200)) = base.I64_extend_i32_u(v195)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+192)) = base.I64_extend_i32_u(v202)
	if l6 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if l7 != 0 {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v216 = F_cstring_to_text(m, l6)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+183)) = uint8(v220)
	goto L68
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+248)) = base.I64_extend_i32_u(v216)
	goto L68
L73:
	;
	if l8 != 0 {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v222 = F_cstring_to_text(m, l7)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+184)) = uint8(v226)
	goto L73
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+256)) = base.I64_extend_i32_u(v222)
	goto L73
L78:
	;
	if l9 != 0 {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v228 = F_cstring_to_text(m, l8)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+185)) = uint8(v232)
	goto L78
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+264)) = base.I64_extend_i32_u(v228)
	goto L78
L83:
	;
	if l10 != 0 {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v234 = F_cstring_to_text(m, l9)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+186)) = uint8(v238)
	goto L83
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+272)) = base.I64_extend_i32_u(v234)
	goto L83
L88:
	;
	v250 = F_heap_form_tuple(m, v189, v22+int32(192), v22+int32(176))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L93
	}
L89:
	;
	v240 = F_cstring_to_text(m, l10)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+187)) = uint8(v244)
	goto L88
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+280)) = base.I64_extend_i32_u(v240)
	goto L88
L93:
	;
	F_CatalogTupleInsert(m, v124, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v22)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = int32(2615)
	v265 = v22 + int32(100)
	F_recordDependencyOn(m, v265, v22+int32(88), int32(110))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_recordDependencyOnOwner(m, int32(3456), v202, l2)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_recordDependencyOnCurrentExtension(m, v265, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_CollationCreate[2]))
	if v278 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v280 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3456), v202, v280, v280)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v250)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	F_relation_close(m, v124, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	v294 = v202
	goto L1
}
func F_CollationIsVisible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v13 = F_SearchSysCache1(m, v8, base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_CollationIsVisible_0), v9)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_F_CollationIsVisible_1), int32(2503), int32(_a_F_CollationIsVisible_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	v36 = v32 + v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v37 != int32(11) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_ReleaseCatCache(m, v13)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L29
	}
L11:
	;
	v40 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_CollationIsVisible[0]))
	if v42 == v40 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v87 = F_CollationGetCollid(m, v36+int32(4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L14:
	;
	if v81 == int32(0) {
		v90 = v40
		goto L10
	} else {
		goto L27
	}
L15:
	;
	v81 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v49 <= int32(0) {
		v75 = v40
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v81 = v75
	goto L14
L19:
	;
	v52 = int32(0)
	if v52 < v49 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v55 = v49
	goto L22
L21:
	;
	v55 = v52
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v58 = int32(0)
	goto L23
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56+v58<<(uint(int32(2))%32))))
	v67 = base.B2i32(v66 == v37)
	if v66 == v37 {
		v75 = v67
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v75 = v67
	goto L18
L25:
	;
	v69 = v58 + int32(1)
	if v69 != v55 {
		v58 = v69
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L13
L28:
	;
	v90 = base.B2i32(v87 == l0)
	goto L10
L29:
	;
	m.G0 = v9 + int32(16)
	return v90
}
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v181 int32
	_ = v181
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	goto L1
L1:
	;
	F_DeconstructQualifiedName(m, l0, v14+int32(12), v14+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L51
	}
L5:
	;
	m.G0 = v14 + int32(16)
	return v181
L6:
	;
	v181 = int32(0)
	goto L5
L7:
	;
	v29 = F_LookupExplicitNamespace(m, v27, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L30
	}
L10:
	;
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(0)
	goto L13
L12:
	;
	v31 = l1
	goto L13
L13:
	;
	if v31 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)))
	v35 = base.I64_extend_i32_u(v29)
	v37 = F_GetSysCacheOid(m, int32(15), v33, base.I64_extend_i32_s(v18), v35, int64(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v37 != 0 {
		v181 = v37
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v41 = F_SearchSysCache3(m, int32(15), v33, int64(-1), v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	if l1|v76 != 0 {
		v181 = v76
		goto L5
	} else {
		goto L29
	}
L18:
	;
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v76 = int32(0)
	goto L17
L20:
	;
	goto L21
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46 + v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+76)))
	if v49 == int32(105) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_ReleaseCatCache(m, v41)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L28
	}
L23:
	;
	v52 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v73 = v72
	goto L22
L26:
	;
	if base.B2i32(base.B2i32(v18 == int32(7))|base.B2i32(base.Ui32(int32(34)) < base.Ui32(v18)) == v52)&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v18))%64)&int64(34357509982) != int64(0)) == int32(0) {
		v73 = v52
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v76 = v73
	goto L17
L29:
	;
	goto L4
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[1]))
	if v82 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L50
	}
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v85 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v89 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)))
	v94 = int32(0)
	goto L34
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v94<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[2]))
	if v105 == v107 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L31
L36:
	;
	v151 = v94 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v151 < v152 {
		v94 = v151
		goto L34
	} else {
		goto L49
	}
L37:
	;
	v110 = base.I64_extend_i32_u(v105)
	v112 = F_GetSysCacheOid(m, int32(15), v89, base.I64_extend_i32_s(v18), v110, int64(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	if v112 != 0 {
		v181 = v112
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v116 = F_SearchSysCache3(m, int32(15), v89, int64(-1), v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v116 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v122 = v120 + v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+76)))
	if v123 != int32(105) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	F_ReleaseCatCache(m, v116)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	if base.B2i32(base.B2i32(v18 == int32(7))|base.B2i32(base.Ui32(int32(34)) < base.Ui32(v18)) == int32(0))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v18))%64)&int64(34357509982) != int64(0)) != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_ReleaseCatCache(m, v116)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L36
L47:
	;
	if v144 != 0 {
		v181 = v144
		goto L5
	} else {
		goto L48
	}
L48:
	;
	goto L36
L49:
	;
	goto L35
L50:
	;
	goto L6
L51:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v212 = F_NameListToString(m, l0)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[0]))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v212
	F_errmsg(m, int32(_a_F_get_collation_oid_0), v14)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_get_collation_oid_1), int32(4089), int32(_a_F_get_collation_oid_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
