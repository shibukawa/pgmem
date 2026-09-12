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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
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
	return v276
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
		v276 = v14
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v115 = F_table_open(m, int32(3456), int32(6))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
	v79 = m.ExcPending
	if v79 != 0 {
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
		v276 = v14
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(290948))
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
	F_errfinish(m, int32(496039), int32(104), int32(355412))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L25
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l0
	F_errmsg(m, int32(332979), v20+int32(32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = l0
	F_errmsg(m, int32(333154), v20+int32(48))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_consts[257])))
	v63 = v62
	goto L23
L22:
	;
	v63 = int32(756936)
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L15
L25:
	;
	v276 = v14
	goto L1
L26:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	F_errfinish(m, int32(496039), int32(114), int32(355412))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L38
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = l0
	F_errmsg(m, int32(116386), v20-int32(-64))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = l0
	F_errmsg(m, int32(116769), v20+int32(80))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(3))%32))+uint32(_consts[257])))
	v100 = v99
	goto L36
L35:
	;
	v100 = int32(756936)
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
	v121 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	goto L43
L41:
	;
	v124 = int32(-1)
	goto L42
L42:
	;
	v126 = F_GetSysCacheOid(m, int32(15), l0, v124, l1, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v124 = v122
	goto L42
L44:
	;
	if v126 != 0 {
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(0)
	v187 = F_strncpy(m, v20+int32(112), l0, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+63)) = uint8(v180)
	goto L66
L48:
	;
	F_sequence_close(m, v115, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v276 = v14
	goto L1
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(3456)
	F_checkMembershipInCurrentExtension(m, v20+int32(100))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L62
	}
L55:
	;
	F_sequence_close(m, v115, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v145 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v145 == int32(0) {
		v276 = v14
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg(m, int32(332979), v20)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(496039), int32(160), int32(355412))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v276 = v14
	goto L1
L62:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	F_errmsg(m, int32(116386), v20+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(496039), int32(167), int32(355412))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
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
	v192 = F_GetNewOidWithIndex(m, v115, int32(3085), int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = v20 + int32(112)
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
	v203 = F_cstring_to_text(m, l6)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+183)) = uint8(v206)
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+220)) = v203
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
	v208 = F_cstring_to_text(m, l7)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v211 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+184)) = uint8(v211)
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v208
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
	v213 = F_cstring_to_text(m, l8)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+185)) = uint8(v216)
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v213
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
	v218 = F_cstring_to_text(m, l9)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+186)) = uint8(v221)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v218
	goto L83
L88:
	;
	v232 = F_heap_form_tuple(m, v179, v20+int32(192), v20+int32(176))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L93
	}
L89:
	;
	v223 = F_cstring_to_text(m, l10)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+187)) = uint8(v226)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v223
	goto L88
L93:
	;
	F_CatalogTupleInsert(m, v115, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v236 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = int32(2615)
	F_recordDependencyOn(m, v20+int32(100), v20+int32(88), int32(110))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_recordDependencyOnOwner(m, int32(3456), v192, l2)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_recordDependencyOnCurrentExtension(m, v20+int32(100), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v262 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v264 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3456), v192, v264, v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v232)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	F_sequence_close(m, v115, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	v276 = v192
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
	var v73 int32
	_ = v73
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
	F_errmsg_internal(m, int32(45921), v9)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(499246), int32(2434), int32(64377))
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
	v41 = *(*int32)(unsafe.Add(mBase, _consts[250]))
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
		v73 = v39
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v80 = v73
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
		v73 = v66
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v73 = v66
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
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
	var v151 int32
	_ = v151
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[251]))
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
		goto L56
	}
L5:
	;
	m.G0 = v12 + int32(16)
	return v151
L6:
	;
	if v65 == int32(0) {
		goto L4
	} else {
		goto L55
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
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L31
	}
L10:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v33 = F_GetSysCacheOid(m, int32(15), v31, v16, v26, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	if v26 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v151 = v3
	goto L5
L14:
	;
	if v33 != 0 {
		v151 = v33
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v37 = F_SearchSysCache3(m, int32(15), v31, int32(-1), v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v37 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+76)))
	if v42 == int32(105) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v65 = v3
	goto L19
L19:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L20:
	;
	F_ReleaseCatCache(m, v37)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32(v16) < base.Ui32(int32(35)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v61 = v60
	goto L20
L24:
	;
	if v56 == int32(0) {
		v61 = int32(0)
		goto L20
	} else {
		goto L28
	}
L25:
	;
	v56 = base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v16))%64)&int64(34357509982) != int64(0))
	goto L27
L26:
	;
	v56 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L23
L29:
	;
	v65 = v61
	goto L19
L30:
	;
	v151 = v65
	goto L5
L31:
	;
	v70 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v72 == v70 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L54
	}
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v75 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v81 = v70
	goto L35
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v81<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	if v92 == v94 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L32
L37:
	;
	v131 = v81 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v131 < v132 {
		v81 = v131
		goto L35
	} else {
		goto L53
	}
L38:
	;
	v98 = F_GetSysCacheOid(m, int32(15), v78, v16, v92, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v98 != 0 {
		v151 = v98
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v102 = F_SearchSysCache3(m, int32(15), v78, int32(-1), v92)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v102 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
	v108 = v106 + v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+76)))
	if v109 != int32(105) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	F_ReleaseCatCache(m, v102)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L51
	}
L44:
	;
	if base.Ui32(v16) < base.Ui32(int32(35)) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v122 != 0 {
		goto L43
	} else {
		goto L49
	}
L46:
	;
	v122 = base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v16))%64)&int64(34357509982) != int64(0))
	goto L48
L47:
	;
	v122 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	F_ReleaseCatCache(m, v102)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	if v125 != 0 {
		v151 = v125
		goto L5
	} else {
		goto L52
	}
L52:
	;
	goto L37
L53:
	;
	goto L36
L54:
	;
	v151 = int32(0)
	goto L5
L55:
	;
	v151 = v65
	goto L5
L56:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v177 = F_NameListToString(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v177
	F_errmsg(m, int32(72000), v12)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(499246), int32(4020), int32(434043))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
