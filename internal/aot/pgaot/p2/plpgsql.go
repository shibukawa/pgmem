package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_append_source_text(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
	F_appendBinaryStringInfo(m, l0, v6+l1, l2-l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_plpgsql_dumptree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	v9 = m.G0
	v11 = v9 - int32(288)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v13
	F_pg_printf(m, int32(768436), v11+int32(272))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_pg_printf(m, int32(768577), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if int32(0) < v27 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pg_printf(m, int32(768357), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v36
	F_pg_printf(m, int32(760024), v11+int32(256))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	switch v49 {
	case 0, 4:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	default:
		goto L11
	}
L10:
	;
	v317 = v36 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if v317 < v318 {
		v36 = v317
		goto L7
	} else {
		goto L88
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v49
	F_pg_printf(m, int32(765511), v11+int32(32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L87
	}
L12:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+240)) = v295
	F_pg_printf(m, int32(765713), v11+int32(240))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L86
	}
L13:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+228)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v230
	F_pg_printf(m, int32(760902), v11+int32(224))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L66
	}
L14:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v180
	F_pg_printf(m, int32(177229), v11+int32(176))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L57
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v11)+148)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v50
	F_pg_printf(m, int32(765537), v11+int32(144))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	if v62 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_pg_printf(m, int32(766315), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+17)))
	if v69 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pg_printf(m, int32(766438), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v76 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_pg_printf(m, int32(758397), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	if v119 != 0 {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v84
	F_pg_printf(m, int32(701129), v11+int32(128))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if int32(0) <= v92 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v92
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	F_pg_printf(m, int32(771523), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v102 = int32(690199)
	goto L35
L34:
	;
	v102 = int32(771673)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v102
	F_pg_printf(m, int32(179917), v11+int32(112))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	goto L27
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v42)+32))
	if int32(0) <= v120 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v171 == int32(0) {
		goto L10
	} else {
		goto L55
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v120
	F_pg_printf(m, int32(765416), v11+int32(96))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_pg_printf(m, int32(758560), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v136
	F_pg_printf(m, int32(701129), v11+int32(80))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	if int32(0) <= v144 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v144
	if v147 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	F_pg_printf(m, int32(771523), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L54
	}
L50:
	;
	v154 = int32(690199)
	goto L52
L51:
	;
	v154 = int32(771673)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v154
	F_pg_printf(m, int32(179917), v11-int32(-64))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	goto L40
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v171
	F_pg_printf(m, int32(765614), v11+int32(48))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L10
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	if int32(0) < v187 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v194 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pg_printf(m, int32(771523), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v200 = v194 << (uint(int32(2)) % 32)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v42)+32))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200+v201)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204+v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v203
	F_pg_printf(m, int32(482058), v11+int32(160))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v215 = v194 + int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	if v215 < v216 {
		v194 = v215
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L10
L66:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	if v239 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_pg_printf(m, int32(766315), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+17)))
	if v246 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pg_printf(m, int32(766438), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v253 == int32(0) {
		goto L10
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	F_pg_printf(m, int32(758397), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v263
	F_pg_printf(m, int32(701129), v11+int32(208))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	if int32(0) <= v271 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v271
	if v274 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	F_pg_printf(m, int32(771523), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	v281 = int32(690199)
	goto L83
L82:
	;
	v281 = int32(771673)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v281
	F_pg_printf(m, int32(179917), v11+int32(192))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	goto L10
L86:
	;
	goto L10
L87:
	;
	goto L10
L88:
	;
	goto L8
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1316])) = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v339
	F_pg_printf(m, int32(560188), v11+int32(16))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	F_dump_block(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v350
	F_pg_printf(m, int32(771222), v11)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v358 = F_fflush(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	m.G0 = v11 + int32(288)
	return
}
func F_plpgsql_estate_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l1)+528)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+61)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v25)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+63)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v27)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v20
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v20
	if l2 != 0 {
		v39 = *(*int32)(unsafe.Add(mBase, _consts[261]))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
		v42 = v41
		v43 = v39
	} else {
		v42 = v6
		v43 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+476))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+504))
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v49
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v54
	v57 = F_makeParamList(m, v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v57
		*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(7201)
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = int32(7202)
		*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = l0
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = int32(4781)
		v72 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v72
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v72
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v74)+28)) = v77
		v80 = *(*int32)(unsafe.Add(mBase, _consts[1309]))
		if v80 == v72 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(103079215120)
			v92 = F_hash_create(m, int32(148262), int32(16), v11, int32(40))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1309])) = v92
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
					v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v100
					v105 = F_hash_create(m, int32(407827), int32(16), v11, int32(1064))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						v127 = v105
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
						v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
						if l4 != 0 {
							v135 = l4
						} else {
							v135 = v130
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
						v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v139 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
						v141 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
							if v154 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
								if v173 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										m.G0 = v11 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, _consts[1312]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v110
					v114 = *(*int32)(unsafe.Add(mBase, _consts[1313]))
					if v114 != 0 {
						v127 = v114
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
						v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
						if l4 != 0 {
							v135 = l4
						} else {
							v135 = v130
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
						v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v139 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
						v141 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
							if v154 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
								if v173 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										m.G0 = v11 + int32(48)
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
						v124 = F_hash_create(m, int32(407854), int32(16), v11, int32(40))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1313])) = v124
							v127 = v124
							*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
							v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
							v131 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
							if l4 != 0 {
								v135 = l4
							} else {
								v135 = v130
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
							v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v139 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
							v141 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
							*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
							F_plpgsql_create_econtext(m, l0)
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return
							} else {
								v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
								if v154 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
									*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
									*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
									*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
									*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
									v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
									if v173 == int32(0) {
										m.G0 = v11 + int32(48)
										return
									} else {
										m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return
										} else {
											m.G0 = v11 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			if l3 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
				v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v100
				v105 = F_hash_create(m, int32(407827), int32(16), v11, int32(1064))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v127 = v105
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
					v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
					if l4 != 0 {
						v135 = l4
					} else {
						v135 = v130
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
					v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v139 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
					v141 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
					F_plpgsql_create_econtext(m, l0)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
						if v154 == int32(0) {
							m.G0 = v11 + int32(48)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
							v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
							if v173 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				}
			} else {
				v110 = *(*int32)(unsafe.Add(mBase, _consts[1312]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v110
				v114 = *(*int32)(unsafe.Add(mBase, _consts[1313]))
				if v114 != 0 {
					v127 = v114
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
					v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
					if l4 != 0 {
						v135 = l4
					} else {
						v135 = v130
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
					v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v139 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
					v141 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
					F_plpgsql_create_econtext(m, l0)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
						if v154 == int32(0) {
							m.G0 = v11 + int32(48)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
							v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
							if v173 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
					v124 = F_hash_create(m, int32(407854), int32(16), v11, int32(40))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1313])) = v124
						v127 = v124
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v127
						v130 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v131
						if l4 != 0 {
							v135 = l4
						} else {
							v135 = v130
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v135
						v138 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v139 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v139
						v141 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v141
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v138
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v139
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v139
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v141
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							v153 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
							if v154 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(7203)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(7204)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = int32(7205)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(7206)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(7200)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
								if v173 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v173].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										m.G0 = v11 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_plpgsql_exec_get_datum_type_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v13 {
	case 0, 4:
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v79
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81
		v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
		v88 = v83
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
		m.G0 = v11 + int32(32)
		return
	default:
		F_errstart_cold(m, int32(21), int32(569208))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v65
			F_errmsg_internal(m, int32(494524), v11)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				F_errfinish(m, int32(512305), int32(5601), int32(247032))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 2:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v15 == int32(2249) {
				v22 = v14 + int32(36)
			} else {
				v22 = l1 + int32(28)
			}
		} else {
			v22 = l1 + int32(28)
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
		v88 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
		m.G0 = v11 + int32(32)
		return
	case 3:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32))))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
		if v34 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v33)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
				v40 = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+48))
				if v41 != v42 {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v47 = F_expanded_record_lookup_field(m, v40, v44, l1+int32(32))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 == int32(0) {
							F_errstart_cold(m, int32(21), int32(569208))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102
									F_errmsg(m, int32(737066), v11+int32(16))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return
									} else {
										F_errfinish(m, int32(512305), int32(5590), int32(247032))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v52
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v56
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
							v88 = v58
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
							m.G0 = v11 + int32(32)
							return
						}
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					v88 = v58
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
					m.G0 = v11 + int32(32)
					return
				}
			}
		} else {
			v40 = v34
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+48))
			if v41 != v42 {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v47 = F_expanded_record_lookup_field(m, v40, v44, l1+int32(32))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					if v47 == int32(0) {
						F_errstart_cold(m, int32(21), int32(569208))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v103
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102
								F_errmsg(m, int32(737066), v11+int32(16))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return
								} else {
									F_errfinish(m, int32(512305), int32(5590), int32(247032))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v56
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
						v88 = v58
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
						m.G0 = v11 + int32(32)
						return
					}
				}
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v56
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				v88 = v58
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88
				m.G0 = v11 + int32(32)
				return
			}
		}
	}
}
func F_plpgsql_exec_trigger(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	F_plpgsql_estate_setup(m, v11+int32(32), l0, v3, v3, v3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(7200)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(360800)
	v27 = int32(4529176)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v11 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v28
	v35 = v11 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v35
	F_copy_plpgsql_datums(m, v35, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+484))
	v43 = int32(2)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42<<(uint(v43)%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+480))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41+v47<<(uint(v43)%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v55 = F_make_expanded_record_from_tupdesc(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v59 = F_make_expanded_record_from_exprecord(m, v55, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v62&int32(4) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l1 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L7:
	;
	switch v62&int32(3) - int32(1) {
	case 0:
		v146 = v59
		goto L8
	case 1:
		goto L11
	case 2:
		goto L10
	default:
		goto L9
	}
L8:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v148 = int32(0)
	F_expanded_record_set_tuple(m, v146, v147, v148, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v146 = v145
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L26
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v73 = int32(0)
	F_expanded_record_set_tuple(m, v71, v72, v73, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v46)+36))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v79 = int32(0)
	F_expanded_record_set_tuple(m, v77, v78, v79, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v83 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v86 != int32(1) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v89&int32(24) != int32(8) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v94 <= int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v102 = int32(0)
	v104 = v94
	goto L18
L18:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(110)+v104<<(uint(int32(4))%32)+v102*int32(100)))))
	if v114 != int32(115) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L6
L20:
	;
	if v129 < v130 {
		v102 = v129
		v104 = v130
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v129 = v102 + int32(1)
	v130 = v104
	goto L20
L22:
	;
	goto L23
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v120 = int32(1)
	v121 = v102 + v120
	v122 = int32(0)
	F_expanded_record_set_field_internal(m, v119, v121, v122, v120, v122, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v129 = v121
	v130 = v128
	goto L20
L25:
	;
	goto L19
L26:
	;
	F_errmsg_internal(m, int32(551505), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(512305), int32(1030), int32(228336))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	goto L6
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(12491)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v332+v333<<(uint(int32(2))%32))))
	v338 = int32(0)
	F_assign_simple_var(m, v11+int32(32), v337, v338, v338, v338)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L84
	}
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v162 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v164 = F_palloc(m, int32(32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v243 != 0 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v170
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v174)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v164)+16)) = base.F64_convert_i64_s(v175)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v178
	if v167 == int32(0) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if v184 == int32(0) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v184)+36))
	if v187 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_register_ENR(m, v237, v164)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L57
	}
L39:
	;
	v188 = int32(0)
	if v187 == v188 {
		v224 = v188
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v232 = F_palloc0(m, int32(4))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L56
	}
L42:
	;
	if v224 != 0 {
		goto L30
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v193 == int32(0) {
		v224 = v188
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v196 <= int32(0) {
		v224 = v188
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v202 = v196
	goto L49
L48:
	;
	v202 = v199
	goto L49
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v205 = int32(0)
	goto L50
L50:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v203+v205<<(uint(int32(2))%32))))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v215 = F_strcmp(m, v214, v167)
	mBase = m.M
	if v215 == int32(0) {
		v224 = v213
		goto L43
	} else {
		goto L52
	}
L51:
	;
	v224 = int32(0)
	goto L43
L52:
	;
	v219 = v205 + int32(1)
	if v219 != v202 {
		v205 = v219
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+36))
	if v229 != 0 {
		v237 = v229
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L41
L56:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	*(*int32)(unsafe.Add(mBase, uint32(v235)+36)) = v232
	v237 = v232
	goto L38
L57:
	;
	goto L34
L58:
	;
	v245 = F_palloc(m, int32(32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L30
L61:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v245)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v251
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v255)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v245)+16)) = base.F64_convert_i64_s(v256)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = v259
	if v248 == int32(0) {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if v265 == int32(0) {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+36))
	if v268 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_register_ENR(m, v318, v245)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L83
	}
L65:
	;
	v269 = int32(0)
	if v268 == v269 {
		v305 = v269
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v313 = F_palloc0(m, int32(4))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L82
	}
L68:
	;
	if v305 != 0 {
		goto L30
	} else {
		goto L80
	}
L69:
	;
	goto L68
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v274 == int32(0) {
		v305 = v269
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v277 <= int32(0) {
		v305 = v269
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v280 = int32(0)
	if v280 < v277 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v283 = v277
	goto L75
L74:
	;
	v283 = v280
	goto L75
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v286 = int32(0)
	goto L76
L76:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v284+v286<<(uint(int32(2))%32))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v296 = F_strcmp(m, v295, v248)
	mBase = m.M
	if v296 == int32(0) {
		v305 = v294
		goto L69
	} else {
		goto L78
	}
L77:
	;
	v305 = int32(0)
	goto L69
L78:
	;
	v300 = v286 + int32(1)
	if v300 != v283 {
		v286 = v300
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+36))
	if v310 != 0 {
		v318 = v310
		goto L64
	} else {
		goto L81
	}
L81:
	;
	goto L67
L82:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+36)) = v313
	v318 = v313
	goto L64
L83:
	;
	goto L60
L84:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v345 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v384 != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v350
	v382 = v350
	goto L85
L87:
	;
	goto L88
L88:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v352 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	if v374 == int32(0) {
		v382 = v373
		goto L85
	} else {
		goto L95
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v357
	v372 = v345
	v373 = v357
	goto L89
L91:
	;
	goto L92
L92:
	;
	m.T0[v352].(func(*base.Module, int32, int32))(m, v11+int32(32), l0)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v368
	if v365 == v366 {
		v382 = v368
		goto L85
	} else {
		goto L94
	}
L94:
	;
	v372 = v365
	v373 = v368
	goto L89
L95:
	;
	m.T0[v374].(func(*base.Module, int32, int32))(m, v11+int32(32), v373)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v382 = v373
	goto L85
L97:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v389 = F_exec_stmt_block(m, v11+int32(32), v382)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v393 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(0)
	if v389 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if v396 == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	m.T0[v396].(func(*base.Module, int32, int32))(m, v11+int32(32), v382)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L157
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(101507)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+61)))
	if v410 == int32(1) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(0)
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L153
	}
L110:
	;
	v413 = int32(0)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	if v414 != 0 {
		v489 = v413
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[1311]))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	if v494 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L112:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v415&int32(4) == int32(0) {
		v489 = v413
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v421 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v456 = F_pg_detoast_datum(m, v420)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L130
	}
L115:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+1)))
	if v424&int32(254) != int32(2) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v420)+2))
	v430 = F_expanded_record_get_tuple(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v429)+44))
	if v432 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v435 = F_expanded_record_fetch_tupdesc(m, v429)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v437 = v432
	goto L120
L120:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+52))
	if v437 == v439 {
		v448 = v430
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v437 = v435
	goto L120
L122:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v448 == v450 {
		v489 = v448
		goto L111
	} else {
		goto L127
	}
L123:
	;
	v442 = F_convert_tuples_by_position(m, v437, v439, int32(402637))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v442 == int32(0) {
		v448 = v430
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v446 = F_execute_attr_map_tuple(m, v430, v442)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v448 = v446
	goto L122
L127:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v448 == v452 {
		v489 = v448
		goto L111
	} else {
		goto L128
	}
L128:
	;
	v454 = F_SPI_copytuple(m, v448)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v489 = v454
	goto L111
L130:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v456
	v460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v460
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)) = uint16(v460)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(base.Ui32(v458) >> (uint(int32(2)) % 32))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	v471 = F_lookup_rowtype_tupdesc(m, v469, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+52))
	v476 = F_convert_tuples_by_position(m, v471, v474, int32(402637))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v476 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v478 = F_execute_attr_map_tuple(m, v11, v476)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v480 = v11
	goto L135
L135:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	if int32(0) <= v481 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v480 = v478
	goto L135
L137:
	;
	F_DecrTupleDescRefCount(m, v471)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v486 = F_SPI_copytuple(m, v480)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v489 = v486
	goto L111
L142:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[1314]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	F_pfree(m, v506)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v497 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	m.T0[v497].(func(*base.Module, int32, int32))(m, v11+int32(32), l0)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1314])) = v507
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	F_FreeExprContext(m, v512, int32(1))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v516 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+152)) = v516
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	if v518 == v516 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v533
	m.G0 = v11 + int32(176)
	return v489
L149:
	;
	F_SPI_freetuptable(m, v518)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v523 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v523
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	if v525 == v523 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v525)+20))
	F_MemoryContextReset(m, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	goto L148
L153:
	;
	F_errcode(m, int32(83887490))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(541135), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(512305), int32(1059), int32(228336))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(109062), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(512305), int32(1067), int32(228336))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_ns_additem(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v5 = F_strlen(m, l2)
	mBase = m.M
	v8 = F_palloc(m, v5+int32(13))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	v12 = int32(4629508)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1315]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13
	v16 = v8 + int32(12)
	if (l2^v16)&int32(3) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1315])) = v8
	return
L4:
	;
	goto L3
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v70)
	if v70&int32(255) == int32(0) {
		goto L4
	} else {
		goto L20
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v69 = l2
	v70 = v22
	v71 = v16
	goto L5
L7:
	;
	goto L8
L8:
	;
	if l2&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = l2
	v28 = v16
	goto L12
L10:
	;
	v40 = l2
	v42 = v16
	goto L11
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 != v47 {
		v69 = v40
		v70 = v44
		v71 = v42
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v29)
	if v29 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v40 = v36
	v42 = v34
	goto L11
L14:
	;
	v33 = int32(1)
	v34 = v28 + v33
	v36 = v26 + v33
	if v36&int32(3) != 0 {
		v26 = v36
		v28 = v34
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v52 = v40
	v53 = v44
	v54 = v42
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v53
	v56 = int32(4)
	v57 = v54 + v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v60 = v52 + v56
	v64 = int32(-2139062144)
	if (v58|(int32(16843008)-v58))&v64 == v64 {
		v52 = v60
		v53 = v58
		v54 = v57
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v69 = v60
	v70 = v58
	v71 = v57
	goto L5
L19:
	;
	goto L18
L20:
	;
	v78 = v69
	v80 = v71
	goto L21
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)) = uint8(v81)
	v83 = int32(1)
	if v81 != 0 {
		v78 = v78 + v83
		v80 = v80 + v83
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L4
L23:
	;
	goto L22
}
func F_plpgsql_ns_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1315])) = int32(0)
	return
}
func F_plpgsql_ns_push(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = l0
	goto L3
L2:
	;
	v8 = int32(771673)
	goto L3
L3:
	;
	v9 = F_strlen(m, v8)
	mBase = m.M
	v12 = F_palloc(m, v9+int32(13))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v17 = int32(4629508)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1315]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v19
	v22 = v12 + int32(12)
	if (v8^v22)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1315])) = v12
	return
L7:
	;
	goto L6
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v76)
	if v76&int32(255) == int32(0) {
		goto L7
	} else {
		goto L23
	}
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v75 = v8
	v76 = v28
	v77 = v22
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v8&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = v8
	v34 = v22
	goto L15
L13:
	;
	v46 = v8
	v48 = v22
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 != v53 {
		v75 = v46
		v76 = v50
		v77 = v48
		goto L8
	} else {
		goto L19
	}
L15:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v35)
	if v35 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	v46 = v42
	v48 = v40
	goto L14
L17:
	;
	v39 = int32(1)
	v40 = v34 + v39
	v42 = v32 + v39
	if v42&int32(3) != 0 {
		v32 = v42
		v34 = v40
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v58 = v46
	v59 = v50
	v60 = v48
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v59
	v62 = int32(4)
	v63 = v60 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v66 = v58 + v62
	v70 = int32(-2139062144)
	if (v64|(int32(16843008)-v64))&v70 == v70 {
		v58 = v66
		v59 = v64
		v60 = v63
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v75 = v66
	v76 = v64
	v77 = v63
	goto L8
L22:
	;
	goto L21
L23:
	;
	v84 = v75
	v86 = v77
	goto L24
L24:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)) = uint8(v87)
	v89 = int32(1)
	if v87 != 0 {
		v84 = v84 + v89
		v86 = v86 + v89
		goto L24
	} else {
		goto L26
	}
L25:
	;
	goto L7
L26:
	;
	goto L25
}
func F_plpgsql_param_eval_generic(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(2))%32)-int32(4))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_exec_eval_datum(m, v11, v19, v8+int32(28), v8+int32(24), v24, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v28 != v29 {
			F_errstart_cold(m, int32(21), int32(569208))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v41 = F_format_type_be(m, v28)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v44 = F_format_type_be(m, v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v41
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v40
							F_errmsg(m, int32(688925), v8)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(512305), int32(6843), int32(500771))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_plpgsql_param_eval_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8<<(uint(int32(2))%32)-int32(4))))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
	return
}
func F_plpgsql_param_eval_var_ro(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8<<(uint(int32(2))%32)-int32(4))))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
	if v16 == int32(0) {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v19 != int32(1) {
			v28 = v15
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v22 != int32(3) {
				v28 = v15
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
				v28 = v25 + int32(18)
			}
		}
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
		v30 = v28
		v31 = v29
	} else {
		v30 = v15
		v31 = v16
	}
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v31)
	return
}
func F_plpgsql_recognize_err_condition(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v220
L3:
	;
	v138 = int32(0)
	v139 = int32(358669)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1305])))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v144 == v138 {
		v163 = v143
		v164 = v144
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v12 = F_strlen(m, l0)
	mBase = m.M
	if v12 != int32(5) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v15 = int32(520760)
	v19 = m.G0
	v21 = v19 - int32(32)
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v22
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1306])))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v98 != int32(5) {
		goto L3
	} else {
		goto L27
	}
L7:
	;
	v98 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1307])))
	if v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = l0
	goto L13
L11:
	;
	goto L12
L12:
	;
	v48 = v15
	v49 = v30
	goto L16
L13:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v44 == v30 {
		v38 = v38 + int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v98 = v38 - l0
	goto L6
L15:
	;
	goto L14
L16:
	;
	v56 = v21 + int32(base.Ui32(v49)>>(uint(int32(3))%32))&int32(28)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 | v58<<(uint(v49)%32)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v62 != 0 {
		v48 = v48 + v58
		v49 = v62
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v65 == int32(0) {
		v90 = l0
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v98 = v90 - l0
	goto L6
L20:
	;
	v69 = l0
	v70 = v65
	goto L21
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(base.Ui32(v70)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v78)>>(uint(v70)%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v90 = v86
	goto L19
L23:
	;
	v90 = v69
	goto L19
L24:
	;
	goto L25
L25:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	v86 = v69 + int32(1)
	if v84 != 0 {
		v69 = v86
		v70 = v84
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v102 = int32(16)
	v104 = int32(63)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v220 = (v101+v102)&v104 | (v106+v102)&v104<<(uint(int32(6))%32) | (v114+v102)&v104<<(uint(int32(12))%32) | (v122+v102)&v104<<(uint(int32(18))%32) | (v130+v102)&v104<<(uint(int32(24))%32)
	goto L2
L28:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v220 = v214
	goto L2
L29:
	;
	if v164-v163 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	goto L29
L31:
	;
	if v143 != v144 {
		v163 = v143
		v164 = v144
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v148 = l0
	v149 = v139
	goto L33
L33:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v152
		v164 = v153
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v163 = v152
	v164 = v153
	goto L30
L35:
	;
	v156 = int32(1)
	if v152 == v153 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v212 = int32(4194736)
	goto L28
L38:
	;
	goto L39
L39:
	;
	v171 = v138
	goto L40
L40:
	;
	v176 = v171 + int32(1)
	v178 = v176 << (uint(int32(3)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1308])))
	if v181 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v212 = v178 + int32(4194736)
	goto L28
L42:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v187 == int32(0) {
		v206 = v186
		v207 = v187
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v207-v206 != 0 {
		v171 = v176
		goto L40
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	if v186 != v187 {
		v206 = v186
		v207 = v187
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v191 = l0
	v192 = v181
	goto L47
L47:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v196 == int32(0) {
		v206 = v195
		v207 = v196
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v206 = v195
	v207 = v196
	goto L44
L49:
	;
	v199 = int32(1)
	if v195 == v196 {
		v191 = v191 + v199
		v192 = v192 + v199
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L41
L52:
	;
	return int32(0)
L53:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(717582), v8)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(507433), int32(2141), int32(255917))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
