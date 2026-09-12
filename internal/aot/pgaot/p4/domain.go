package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitDomainConstraintRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	v4 = l3
	v5 = int32(0)
	v9 = F_lookup_type_cache(m, l0, int32(8192))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(1620)
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v20 = l1 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v20
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+308))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v30 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v29 + v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v35 != v30 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v96 = v5
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
	return
L8:
	;
	goto L9
L9:
	;
	v39 = int32(4520272)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
	if v34 == int32(0) {
		v87 = v5
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v40
	v96 = v87
	goto L6
L11:
	;
	v46 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v47 <= v46 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v87 = v5
	goto L10
L13:
	;
	goto L14
L14:
	;
	v50 = v46
	v54 = v5
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v50<<(uint(int32(2))%32))))
	v63 = F_palloc0(m, int32(20))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v87 = v77
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(393)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = F_ExecInitExpr(m, v71, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v74
	v77 = F_lappend(m, v54, v63)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v80 = v50 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v80 < v81 {
		v50 = v80
		v54 = v77
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
}
func F_domain_check_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	v2 = l1
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v21 = l2 + int32(44)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+314)))
	if v23&int32(8) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+308))
	if v31 == v32 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)))
	if v26 != int32(100) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_load_domaintype_info(m, v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v156 == int32(0) {
		v300 = v19
		goto L31
	} else {
		goto L32
	}
L7:
	;
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v40 = v38 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v40
	if v40 <= v34 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v48 = v32
	goto L10
L10:
	;
	if v48 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	F_MemoryContextDelete(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+308))
	v48 = v47
	goto L10
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v48
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v52 + v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)))
	if v57 != v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v136
	goto L6
L17:
	;
	v136 = v56
	goto L16
L18:
	;
	goto L19
L19:
	;
	v60 = int32(4520272)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v63
	if v56 == int32(0) {
		v120 = v5
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v61
	v136 = v120
	goto L16
L21:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v68 <= v67 {
		v120 = v5
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v80 = v5
	v81 = v67
	goto L23
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v81<<(uint(int32(2))%32))))
	v91 = F_palloc0(m, int32(20))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v120 = v105
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(393)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v99
	v102 = F_ExecInitExpr(m, v99, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v102
	v105 = F_lappend(m, v80, v91)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v108 = v81 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v108 < v109 {
		v80 = v105
		v81 = v108
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	m.G0 = v17 + int32(48)
	return
L30:
	;
	F_ReScanExprContext(m, v316)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L74
	}
L31:
	;
	if v300 == int32(0) {
		goto L29
	} else {
		goto L73
	}
L32:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v159 <= int32(0) {
		v300 = v19
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v170 = v19
	v174 = v5
	goto L34
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v174<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	switch v181 {
	case 0:
		goto L39
	case 1:
		goto L38
	default:
		goto L37
	}
L35:
	;
	v300 = v287
	goto L31
L36:
	;
	v289 = v174 + int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v289 < v290 {
		v170 = v287
		v174 = v289
		goto L34
	} else {
		goto L72
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L69
	}
L38:
	;
	if v170 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	if v2 == int32(0) {
		v287 = v170
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v184 = F_errsave_start(m, l3)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v184 == int32(0) {
		v300 = v170
		goto L31
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v192 = F_format_type_be(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v192
	F_errmsg(m, int32(158583), v17+int32(16))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_errdatatype(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errsave_finish(m, l3, int32(495008), int32(160), int32(64932))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v300 = v170
	goto L31
L48:
	;
	v210 = int32(4520272)
	v211 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v213
	v215 = F_CreateStandaloneExprContext(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v221 = v170
	goto L50
L50:
	;
	if v2 != 0 {
		v236 = l0
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v211
	*(*int32)(unsafe.Add(mBase, uint32(l2)+76)) = v215
	v221 = v215
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v221)+52)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+48)) = v236
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v240 = F_ExecCheck(m, v239, v221)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L59
	}
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+8)))
	if v223 != int32(65535) {
		v236 = l0
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v226 != int32(1) {
		v235 = l0
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v236 = v235
	goto L52
L56:
	;
	goto L55
L57:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v229 != int32(3) {
		v235 = l0
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
	v235 = v232 + int32(18)
	goto L56
L59:
	;
	if v240 != 0 {
		v287 = v221
		goto L36
	} else {
		goto L60
	}
L60:
	;
	v242 = F_errsave_start(m, l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v242 == int32(0) {
		v316 = v221
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v250 = F_format_type_be(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v250
	F_errmsg(m, int32(700243), v17+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_errdatatype(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_err_generic_string(m, int32(110), v260)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, l3, int32(495008), int32(200), int32(64932))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v316 = v221
	goto L30
L69:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v276
	F_errmsg_internal(m, int32(485588), v17)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(495008), int32(207), int32(64932))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	goto L35
L73:
	;
	v316 = v300
	goto L30
L74:
	;
	goto L29
}
