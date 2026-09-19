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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	v14 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(240)
	m.G0 = v20
	v24 = F_GetSysCacheOid(m, int32(15), l0, l5, l1, v14)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(240)
	return v268
L2:
	;
	return int32(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l12 != 0 {
		v268 = v14
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v109 = F_table_open(m, int32(3456), int32(6))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(3456)
	F_checkMembershipInCurrentExtension(m, v20+int32(100))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L26
	}
L11:
	;
	v39 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v39 == int32(0) {
		v268 = v14
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L25
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_1), v20+int32(32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(l5) <= base.Ui32(int32(41)) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_5), v20+int32(48))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_c_F_CollationCreate[0])))
	v60 = v58
	goto L23
L22:
	;
	v60 = int32(_a_F_CollationCreate_4)
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L15
L25:
	;
	v268 = v14
	goto L1
L26:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L38
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_6), v20-int32(-64))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(l5) <= base.Ui32(int32(41)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L28
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_7), v20+int32(80))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_c_F_CollationCreate[0])))
	v94 = v92
	goto L36
L35:
	;
	v94 = int32(_a_F_CollationCreate_4)
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
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_CollationCreate[1]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	goto L43
L41:
	;
	v118 = int32(-1)
	goto L42
L42:
	;
	v120 = F_GetSysCacheOid(m, int32(15), l0, v118, l1, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v118 = v116
	goto L42
L44:
	;
	if v120 != 0 {
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
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v174
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(0)
	v179 = v20 + int32(112)
	v181 = F_strncpy(m, v179, l0, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+63)) = uint8(v174)
	goto L66
L48:
	;
	F_relation_close(m, v109, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v268 = v14
	goto L1
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(3456)
	F_checkMembershipInCurrentExtension(m, v20+int32(100))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L62
	}
L55:
	;
	F_relation_close(m, v109, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v139 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v139 == int32(0) {
		v268 = v14
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_1), v20)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(160), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v268 = v14
	goto L1
L62:
	;
	F_errcode(m, int32(_a_F_CollationCreate_0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	F_errmsg(m, int32(_a_F_CollationCreate_6), v20+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CollationCreate_2), int32(167), int32(_a_F_CollationCreate_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v186 = F_GetNewOidWithIndex(m, v109, int32(3085), int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+216)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+204)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = v179
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
	v195 = F_cstring_to_text(m, l6)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+183)) = uint8(v198)
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+220)) = v195
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
	v200 = F_cstring_to_text(m, l7)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+184)) = uint8(v203)
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v200
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
	v205 = F_cstring_to_text(m, l8)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+185)) = uint8(v208)
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v205
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
	v210 = F_cstring_to_text(m, l9)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+186)) = uint8(v213)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v210
	goto L83
L88:
	;
	v224 = F_heap_form_tuple(m, v173, v20+int32(192), v20+int32(176))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L93
	}
L89:
	;
	v215 = F_cstring_to_text(m, l10)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+187)) = uint8(v218)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v215
	goto L88
L93:
	;
	F_CatalogTupleInsert(m, v109, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v228 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = int32(2615)
	v239 = v20 + int32(100)
	F_recordDependencyOn(m, v239, v20+int32(88), int32(110))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_recordDependencyOnOwner(m, int32(3456), v186, l2)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_recordDependencyOnCurrentExtension(m, v239, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_CollationCreate[2]))
	if v252 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v254 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3456), v186, v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v224)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	F_relation_close(m, v109, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	v268 = v186
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v12 = F_SearchSysCache1(m, v8, l0)
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
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_CollationIsVisible_0), v9)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_F_CollationIsVisible_1), int32(2434), int32(_a_F_CollationIsVisible_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
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
	v35 = v31 + v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	if v36 != int32(11) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_ReleaseCatCache(m, v12)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L11:
	;
	v39 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_CollationIsVisible[0]))
	if v41 == v39 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v86 = F_CollationGetCollid(m, v35+int32(4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L14:
	;
	if v80 == int32(0) {
		v89 = v39
		goto L10
	} else {
		goto L27
	}
L15:
	;
	v80 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v48 <= int32(0) {
		v74 = v39
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v80 = v74
	goto L14
L19:
	;
	v51 = int32(0)
	if v51 < v48 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = v48
	goto L22
L21:
	;
	v54 = v51
	goto L22
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v57 = int32(0)
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55+v57<<(uint(int32(2))%32))))
	v66 = base.B2i32(v65 == v36)
	if v65 == v36 {
		v74 = v66
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v74 = v66
	goto L18
L25:
	;
	v68 = v57 + int32(1)
	if v68 != v54 {
		v57 = v68
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
	v89 = base.B2i32(v86 == l0)
	goto L10
L29:
	;
	m.G0 = v9 + int32(16)
	return v89
}
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v150 int32
	_ = v150
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	goto L1
L1:
	;
	F_DeconstructQualifiedName(m, l0, v12+int32(12), v12+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L48
	}
L5:
	;
	m.G0 = v12 + int32(16)
	return v150
L6:
	;
	if v67 == int32(0) {
		goto L4
	} else {
		goto L47
	}
L7:
	;
	v26 = F_LookupExplicitNamespace(m, v25, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L26
	}
L10:
	;
	v28 = int32(0)
	if v26|base.B2i32(l1 == v28) == v28 {
		v150 = v3
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v36 = F_GetSysCacheOid(m, int32(15), v34, v16, v26, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v36 != 0 {
		v150 = v36
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v40 = F_SearchSysCache3(m, int32(15), v34, int32(-1), v26)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+76)))
	if v45 == int32(105) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = v3
	goto L17
L17:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L18:
	;
	F_ReleaseCatCache(m, v40)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L24
	}
L19:
	;
	goto L22
L20:
	;
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v63 = v62
	goto L18
L22:
	;
	if base.B2i32(base.Ui32(v16) < base.Ui32(int32(35)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v16))%64)&int64(34357509982) != int64(0)) == int32(0) {
		v63 = int32(0)
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v67 = v63
	goto L17
L25:
	;
	v150 = v67
	goto L5
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[1]))
	if v73 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L46
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v76 <= int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v85 = v3
	goto L30
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v85<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[2]))
	if v93 == v95 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	v131 = v85 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v131 < v132 {
		v85 = v131
		goto L30
	} else {
		goto L45
	}
L33:
	;
	v99 = F_GetSysCacheOid(m, int32(15), v79, v16, v93, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v99 != 0 {
		v150 = v99
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v103 = F_SearchSysCache3(m, int32(15), v79, int32(-1), v93)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v103 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
	v109 = v107 + v108
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+76)))
	if v110 != int32(105) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	F_ReleaseCatCache(m, v103)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	if base.B2i32(base.Ui32(v16) < base.Ui32(int32(35)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v16))%64)&int64(34357509982) != int64(0)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	F_ReleaseCatCache(m, v103)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	if v125 != 0 {
		v150 = v125
		goto L5
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	goto L31
L46:
	;
	v150 = int32(0)
	goto L5
L47:
	;
	v150 = v67
	goto L5
L48:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v177 = F_NameListToString(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_get_collation_oid[0]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v177
	F_errmsg(m, int32(_a_F_get_collation_oid_0), v12)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_get_collation_oid_1), int32(4020), int32(_a_F_get_collation_oid_2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
