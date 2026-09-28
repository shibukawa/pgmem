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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	v9 = m.G0
	v11 = v9 - int32(288)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v13
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_0), v11+int32(272))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_1), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if int32(0) < v24 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_2), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v32
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_3), v11+int32(256))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	switch v46 {
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
	v288 = v32 + int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if v288 < v289 {
		v32 = v288
		goto L7
	} else {
		goto L88
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v46
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_4), v11+int32(32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L87
	}
L12:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+240)) = v266
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_5), v11+int32(240))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L86
	}
L13:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+228)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v209
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_6), v11+int32(224))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L66
	}
L14:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v159
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_7), v11+int32(176))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L57
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v11)+148)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v47
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_8), v11+int32(144))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+16)))
	if v59 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_9), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+17)))
	if v66 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_10), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_11), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if v107 != 0 {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v79
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_12), v11+int32(128))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if int32(0) <= v86 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v86
	if v89 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_13), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v93 = int32(_a_F_plpgsql_dumptree_14)
	goto L35
L34:
	;
	v93 = int32(_a_F_plpgsql_dumptree_15)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v93
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_16), v11+int32(112))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if int32(0) <= v108 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v150 == int32(0) {
		goto L10
	} else {
		goto L55
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v108
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_17), v11+int32(96))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_18), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v122
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_12), v11+int32(80))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	if int32(0) <= v129 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v129
	if v132 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_13), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L54
	}
L50:
	;
	v136 = int32(_a_F_plpgsql_dumptree_14)
	goto L52
L51:
	;
	v136 = int32(_a_F_plpgsql_dumptree_15)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v136
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_16), v11-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v150
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_19), v11+int32(48))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L10
L57:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if int32(0) < v166 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_13), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v179 = v173 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179+v180)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183+v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v182
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_20), v11+int32(160))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v194 = v173 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if v194 < v195 {
		v173 = v194
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
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+16)))
	if v218 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_9), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+17)))
	if v225 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_10), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v232 == int32(0) {
		goto L10
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_11), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v240
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_12), v11+int32(208))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	if int32(0) <= v247 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v247
	if v250 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_13), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	v254 = int32(_a_F_plpgsql_dumptree_14)
	goto L83
L82:
	;
	v254 = int32(_a_F_plpgsql_dumptree_15)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v254
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_16), v11+int32(192))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_dumptree[0])) = int32(0)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v307
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_21), v11+int32(16))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	F_dump_block(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v317
	F_pg_printf(m, int32(_a_F_plpgsql_dumptree_22), v11)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_dumptree[1]))
	v324 = F_fflush(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
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
	var v18 int64
	_ = v18
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l1)+528)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v16)
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+61)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v25)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+63)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v27)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v18
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v18
	if l2 != 0 {
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[0]))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
		v42 = v41
		v43 = v39
	} else {
		v42 = v6
		v43 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+476))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+504))
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v49
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v54
	v57 = F_makeParamList(m, v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v57
		*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(_a_F_plpgsql_estate_setup_0)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = l0
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = int32(_a_F_plpgsql_estate_setup_1)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v68 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v68
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = int32(_a_F_plpgsql_estate_setup_2)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v68
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = v77
		v80 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[2]))
		if v80 == v68 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(103079215120)
			v89 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_3), int64(16), v11, int32(40))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[2])) = v89
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(137438953488)
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v96
					v101 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_4), int64(16), v11, int32(1064))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						v117 = v101
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							if v144 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
								if v157 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
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
					v104 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[5]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v104
					v107 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6]))
					if v107 != 0 {
						v117 = v107
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							if v144 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
								if v157 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										m.G0 = v11 + int32(48)
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(137438953488)
						v114 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_10), int64(16), v11, int32(40))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6])) = v114
							v117 = v114
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
							v121 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
							if l4 != 0 {
								v125 = l4
							} else {
								v125 = v120
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
							v129 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
							v131 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
							*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
							F_plpgsql_create_econtext(m, l0)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								if v144 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
									*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
									*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
									*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
									*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
									if v157 == int32(0) {
										m.G0 = v11 + int32(48)
										return
									} else {
										m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(137438953488)
				v96 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v96
				v101 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_4), int64(16), v11, int32(1064))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					v117 = v101
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
					v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
					if l4 != 0 {
						v125 = l4
					} else {
						v125 = v120
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					v129 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
					*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
					*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
					F_plpgsql_create_econtext(m, l0)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						if v144 == int32(0) {
							m.G0 = v11 + int32(48)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
							v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
							if v157 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
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
				v104 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[5]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v104
				v107 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6]))
				if v107 != 0 {
					v117 = v107
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
					v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
					if l4 != 0 {
						v125 = l4
					} else {
						v125 = v120
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					v129 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
					*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
					*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
					F_plpgsql_create_econtext(m, l0)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						if v144 == int32(0) {
							m.G0 = v11 + int32(48)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
							*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
							v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
							if v157 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(137438953488)
					v114 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_10), int64(16), v11, int32(40))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6])) = v114
						v117 = v114
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v131
						F_plpgsql_create_econtext(m, l0)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[4]))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							if v144 == int32(0) {
								m.G0 = v11 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = int32(_a_F_plpgsql_estate_setup_5)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = int32(_a_F_plpgsql_estate_setup_6)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = int32(_a_F_plpgsql_estate_setup_7)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(_a_F_plpgsql_estate_setup_8)
								*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(_a_F_plpgsql_estate_setup_9)
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
								if v157 == int32(0) {
									m.G0 = v11 + int32(48)
									return
								} else {
									m.T0[v157].(func(*base.Module, int32, int32))(m, l0, l1)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v12 {
	case 0, 4:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v75
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
		v81 = v77
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
		m.G0 = v10 + int32(32)
		return
	default:
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_info_0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v62
			F_errmsg_internal(m, int32(_a_F_plpgsql_exec_get_datum_type_info_1), v10)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_info_2), int32(_a_F_plpgsql_exec_get_datum_type_info_3), int32(_a_F_plpgsql_exec_get_datum_type_info_4))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 2:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v13 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v14 == int32(2249) {
				v21 = v13 + int32(36)
			} else {
				v21 = l1 + int32(28)
			}
		} else {
			v21 = l1 + int32(28)
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
		v81 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
		m.G0 = v10 + int32(32)
		return
	case 3:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32))))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
		if v33 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
				v39 = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+48))
				if v40 != v41 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v46 = F_expanded_record_lookup_field(m, v39, v43, l1+int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						if v46 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_info_0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v94
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v93
									F_errmsg(m, int32(_a_F_plpgsql_exec_get_datum_type_info_5), v10+int32(16))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_info_2), int32(_a_F_plpgsql_exec_get_datum_type_info_6), int32(_a_F_plpgsql_exec_get_datum_type_info_4))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
							v81 = v57
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
							m.G0 = v10 + int32(32)
							return
						}
					}
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					v81 = v57
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
					m.G0 = v10 + int32(32)
					return
				}
			}
		} else {
			v39 = v33
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+48))
			if v40 != v41 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v46 = F_expanded_record_lookup_field(m, v39, v43, l1+int32(32))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					if v46 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_info_0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v94
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v93
								F_errmsg(m, int32(_a_F_plpgsql_exec_get_datum_type_info_5), v10+int32(16))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_info_2), int32(_a_F_plpgsql_exec_get_datum_type_info_6), int32(_a_F_plpgsql_exec_get_datum_type_info_4))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
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
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
						v81 = v57
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
						m.G0 = v10 + int32(32)
						return
					}
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				v81 = v57
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v81
				m.G0 = v10 + int32(32)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int64
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
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
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	v15 = v12 + int32(40)
	F_plpgsql_estate_setup(m, v15, l0, v3, v3, v3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_plpgsql_exec_trigger_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(_a_F_plpgsql_exec_trigger_1)
	v28 = int32(_a_F_plpgsql_exec_trigger_2)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0])) = v12 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v15
	F_copy_plpgsql_datums(m, v15, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+484))
	v40 = int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38+v39<<(uint(v40)%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+480))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38+v44<<(uint(v40)%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+52))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v52 = F_make_expanded_record_from_tupdesc(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v56 = F_make_expanded_record_from_exprecord(m, v52, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v59&int32(4) == int32(0) {
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
	switch v59&int32(3) - int32(1) {
	case 0:
		v142 = v56
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
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v144 = int32(0)
	F_expanded_record_set_tuple(m, v142, v143, v144, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	v142 = v141
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_trigger_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L26
	}
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v70 = int32(0)
	F_expanded_record_set_tuple(m, v68, v69, v70, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v76 = int32(0)
	F_expanded_record_set_tuple(m, v74, v75, v76, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	if v80 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+17)))
	if v83 != int32(1) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v86&int32(24) != int32(8) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v91 <= int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v97 = int32(0)
	v98 = v91
	goto L18
L18:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v98<<(uint(int32(3))%32)+v97*int32(100))+118)))
	if v110 != int32(115) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L6
L20:
	;
	if v125 < v126 {
		v97 = v125
		v98 = v126
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v125 = v97 + int32(1)
	v126 = v98
	goto L20
L22:
	;
	goto L23
L23:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	v116 = int32(1)
	v117 = v97 + v116
	v120 = int32(0)
	F_expanded_record_set_field_internal(m, v115, v117, int64(0), v116, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v125 = v117
	v126 = v124
	goto L20
L25:
	;
	goto L19
L26:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_trigger_4), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1031), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(_a_F_plpgsql_exec_trigger_7)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328+v329<<(uint(int32(2))%32))))
	v335 = int32(0)
	F_assign_simple_var(m, v12+int32(40), v333, int64(0), v335, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L84
	}
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v159 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v161 = F_palloc(m, int32(32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v240 != 0 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v161)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = v167
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v171)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v161)+16)) = base.F64_convert_i64_s(v172)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v175
	if v164 == int32(0) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	if v181 == int32(0) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+36))
	if v184 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_register_ENR(m, v234, v161)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L57
	}
L39:
	;
	v185 = int32(0)
	if v184 == v185 {
		v222 = v185
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v229 = F_palloc0(m, int32(4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L56
	}
L42:
	;
	if v222 != 0 {
		goto L30
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v190 == int32(0) {
		v222 = v185
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v193 <= int32(0) {
		v222 = v185
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v196 = int32(0)
	if v196 < v193 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v199 = v193
	goto L49
L48:
	;
	v199 = v196
	goto L49
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v202 = int32(0)
	goto L50
L50:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v200+v202<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v212 = F_strcmp(m, v211, v164)
	mBase = m.M
	if v212 == int32(0) {
		v222 = v210
		goto L43
	} else {
		goto L52
	}
L51:
	;
	v222 = int32(0)
	goto L43
L52:
	;
	v216 = v202 + int32(1)
	if v216 != v199 {
		v202 = v216
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+36))
	if v226 != 0 {
		v234 = v226
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L41
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v229
	v234 = v229
	goto L38
L57:
	;
	goto L34
L58:
	;
	v242 = F_palloc(m, int32(32))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v242)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v248
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v242)+16)) = base.F64_convert_i64_s(v253)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = v256
	if v245 == int32(0) {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	if v262 == int32(0) {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262)+36))
	if v265 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_register_ENR(m, v315, v242)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L83
	}
L65:
	;
	v266 = int32(0)
	if v265 == v266 {
		v303 = v266
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v310 = F_palloc0(m, int32(4))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L82
	}
L68:
	;
	if v303 != 0 {
		goto L30
	} else {
		goto L80
	}
L69:
	;
	goto L68
L70:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	if v271 == int32(0) {
		v303 = v266
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v274 <= int32(0) {
		v303 = v266
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v277 = int32(0)
	if v277 < v274 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v280 = v274
	goto L75
L74:
	;
	v280 = v277
	goto L75
L75:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v283 = int32(0)
	goto L76
L76:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v281+v283<<(uint(int32(2))%32))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v293 = F_strcmp(m, v292, v245)
	mBase = m.M
	if v293 == int32(0) {
		v303 = v291
		goto L69
	} else {
		goto L78
	}
L77:
	;
	v303 = int32(0)
	goto L69
L78:
	;
	v297 = v283 + int32(1)
	if v297 != v280 {
		v283 = v297
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+36))
	if v307 != 0 {
		v315 = v307
		goto L64
	} else {
		goto L81
	}
L81:
	;
	goto L67
L82:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+36)) = v310
	v315 = v310
	goto L64
L83:
	;
	goto L60
L84:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v341 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[3]))
	if v380 != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v346
	v378 = v346
	goto L85
L87:
	;
	goto L88
L88:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v348 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	if v370 == int32(0) {
		v378 = v369
		goto L85
	} else {
		goto L95
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v353
	v368 = v341
	v369 = v353
	goto L89
L91:
	;
	goto L92
L92:
	;
	m.T0[v348].(func(*base.Module, int32, int32))(m, v12+int32(40), l0)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v362 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v364
	if v361 == v362 {
		v378 = v364
		goto L85
	} else {
		goto L94
	}
L94:
	;
	v368 = v361
	v369 = v364
	goto L89
L95:
	;
	m.T0[v370].(func(*base.Module, int32, int32))(m, v12+int32(40), v369)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v378 = v369
	goto L85
L97:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v384 = v12 + int32(40)
	v385 = F_exec_stmt_block(m, v384, v378)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	if v389 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = int32(0)
	if v385 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	if v392 == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	m.T0[v392].(func(*base.Module, int32, int32))(m, v384, v378)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_trigger_3))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L158
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(_a_F_plpgsql_exec_trigger_8)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+77)))
	if v404 == int32(1) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = int32(0)
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_trigger_3))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L154
	}
L110:
	;
	v407 = int32(0)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+64)))
	if v408 != 0 {
		v489 = v407
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v495 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L112:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v409&int32(4) == int32(0) {
		v489 = v407
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
	v415 = base.I32_wrap_i64(v414)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	if v416 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v452 = F_pg_detoast_datum(m, v415)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L131
	}
L115:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	if v419&int32(254) != int32(2) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v414))+2))
	goto L117
L117:
	;
	v426 = F_expanded_record_get_tuple(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425)+44))
	if v428 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v431 = F_expanded_record_fetch_tupdesc(m, v425)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	v433 = v428
	goto L121
L121:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+52))
	if v433 == v435 {
		v444 = v426
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v433 = v431
	goto L121
L123:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v444 == v446 {
		v489 = v444
		goto L111
	} else {
		goto L128
	}
L124:
	;
	v438 = F_convert_tuples_by_position(m, v433, v435, int32(_a_F_plpgsql_exec_trigger_9))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v438 == int32(0) {
		v444 = v426
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v442 = F_execute_attr_map_tuple(m, v426, v438)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v444 = v442
	goto L123
L128:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v444 == v448 {
		v489 = v444
		goto L111
	} else {
		goto L129
	}
L129:
	;
	v450 = F_SPI_copytuple(m, v444)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v489 = v450
	goto L111
L131:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v452
	v456 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v456
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v456)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(base.Ui32(v454) >> (uint(int32(2)) % 32))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v467 = F_lookup_rowtype_tupdesc(m, v465, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+52))
	v472 = F_convert_tuples_by_position(m, v467, v470, int32(_a_F_plpgsql_exec_trigger_9))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	if v472 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v476 = F_execute_attr_map_tuple(m, v12+int32(8), v472)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	v480 = v12 + int32(8)
	goto L136
L136:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	if int32(0) <= v481 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v480 = v476
	goto L136
L138:
	;
	F_DecrTupleDescRefCount(m, v467)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v486 = F_SPI_copytuple(m, v480)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v489 = v486
	goto L111
L143:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[4]))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	F_pfree(m, v507)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L147
	}
L144:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	if v498 == int32(0) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	m.T0[v498].(func(*base.Module, int32, int32))(m, v12+int32(40), l0)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[4])) = v508
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	F_FreeExprContext(m, v513, int32(1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v517 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v519 == v517 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0])) = v534
	m.G0 = v12 + int32(192)
	return v489
L150:
	;
	F_SPI_freetuptable(m, v519)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v524 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v524
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	if v526 == v524 {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526)+20))
	F_MemoryContextReset(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	F_errcode(m, int32(83887490))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_trigger_10), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1060), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_trigger_11), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1068), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_additem[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13
	v16 = v8 + int32(12)
	if (l2^v16)&int32(3) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_additem[0])) = v8
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
	v59 = v52 + v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v64 = int32(-2139062144)
	if (int32(16843008)-v61|v61)&v64 == v64 {
		v52 = v59
		v53 = v61
		v54 = v57
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v69 = v59
	v70 = v61
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_init[0])) = int32(0)
	return
}
func F_plpgsql_ns_push(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = l0
	goto L3
L2:
	;
	v5 = int32(_a_F_plpgsql_ns_push_0)
	goto L3
L3:
	;
	v6 = F_strlen(m, v5)
	mBase = m.M
	v9 = F_palloc(m, v6+int32(13))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_push[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
	v18 = v9 + int32(12)
	if (v5^v18)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_push[0])) = v9
	return
L7:
	;
	goto L6
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v72)
	if v72&int32(255) == int32(0) {
		goto L7
	} else {
		goto L23
	}
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v71 = v5
	v72 = v24
	v73 = v18
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v5&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = v5
	v30 = v18
	goto L15
L13:
	;
	v42 = v5
	v44 = v18
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 != v49 {
		v71 = v42
		v72 = v46
		v73 = v44
		goto L8
	} else {
		goto L19
	}
L15:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v31)
	if v31 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	v42 = v38
	v44 = v36
	goto L14
L17:
	;
	v35 = int32(1)
	v36 = v30 + v35
	v38 = v28 + v35
	if v38&int32(3) != 0 {
		v28 = v38
		v30 = v36
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v54 = v42
	v55 = v46
	v56 = v44
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v55
	v58 = int32(4)
	v59 = v56 + v58
	v61 = v54 + v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 == v66 {
		v54 = v61
		v55 = v63
		v56 = v59
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v71 = v61
	v72 = v63
	v73 = v59
	goto L8
L22:
	;
	goto L21
L23:
	;
	v80 = v71
	v82 = v73
	goto L24
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)) = uint8(v83)
	v85 = int32(1)
	if v83 != 0 {
		v80 = v80 + v85
		v82 = v82 + v85
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(int32(2))%32)-int32(4))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_exec_eval_datum(m, v10, v18, v7+int32(28), v7+int32(24), v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v27 != v28 {
			F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_param_eval_generic_0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v38 = F_format_type_be(m, v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v41 = F_format_type_be(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v41
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v37
							F_errmsg(m, int32(_a_F_plpgsql_param_eval_generic_1), v7)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_param_eval_generic_2), int32(_a_F_plpgsql_param_eval_generic_3), int32(_a_F_plpgsql_param_eval_generic_4))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
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
			m.G0 = v7 + int32(32)
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
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8<<(uint(int32(2))%32)-int32(4))))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
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
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8<<(uint(int32(2))%32)-int32(4))))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	if v17 == int32(0) {
		v21 = base.I32_wrap_i64(v15)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		if v22 != int32(1) {
			v32 = v15
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
			if v25 != int32(3) {
				v32 = v15
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2))
				v32 = base.I64_extend_i32_u(v28 + int32(18))
			}
		}
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
		v34 = v33
		v35 = v32
	} else {
		v34 = int32(1)
		v35 = v15
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v34)
	return
}
func F_plpgsql_recognize_err_condition(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_recognize_err_condition_0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L39
	} else {
		goto L40
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v182
L3:
	;
	v139 = int32(0)
	goto L26
L4:
	;
	v11 = F_strlen(m, l0)
	mBase = m.M
	if v11 != int32(5) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(_a_F_plpgsql_recognize_err_condition_1)
	v18 = m.G0
	v20 = v18 - int32(32)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v21
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_recognize_err_condition[0])))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v97 != int32(5) {
		goto L3
	} else {
		goto L25
	}
L7:
	;
	v97 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_recognize_err_condition[1])))
	if v33 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = l0
	goto L13
L11:
	;
	goto L12
L12:
	;
	v47 = v14
	v48 = v29
	goto L16
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v43 == v29 {
		v37 = v37 + int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v97 = v37 - l0
	goto L6
L15:
	;
	goto L14
L16:
	;
	v55 = v20 + int32(base.Ui32(v48)>>(uint(int32(3))%32))&int32(28)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56 | v57<<(uint(v48)%32)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v61 != 0 {
		v47 = v47 + v57
		v48 = v61
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v64 == int32(0) {
		v87 = l0
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v97 = v87 - l0
	goto L6
L20:
	;
	v68 = l0
	v69 = v64
	goto L21
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(base.Ui32(v69)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v77)>>(uint(v69)%32))&int32(1) == int32(0) {
		v87 = v68
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v87 = v85
	goto L19
L23:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v85 = v68 + int32(1)
	if v83 != 0 {
		v68 = v85
		v69 = v83
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v101 = int32(16)
	v103 = int32(63)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v182 = (v100+v101)&v103 | (v105+v101)&v103<<(uint(int32(6))%32) | (v113+v101)&v103<<(uint(int32(12))%32) | (v121+v101)&v103<<(uint(int32(18))%32) | (v129+v101)&v103<<(uint(int32(24))%32)
	goto L2
L26:
	;
	v143 = v139 << (uint(int32(3)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_plpgsql_recognize_err_condition[2])))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if base.B2i32(v149 == int32(0))|base.B2i32(v149 != v152) != 0 {
		v170 = v149
		v171 = v152
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_plpgsql_recognize_err_condition[3])))
	v182 = v177
	goto L2
L28:
	;
	if v170-v171 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	v155 = l0
	v156 = v146
	goto L31
L31:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	if v160 == int32(0) {
		v170 = v160
		v171 = v159
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v170 = v160
	v171 = v159
	goto L29
L33:
	;
	v163 = int32(1)
	if v160 == v159 {
		v155 = v155 + v163
		v156 = v156 + v163
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v174 = v139 + int32(1)
	if v174 != int32(251) {
		v139 = v174
		goto L26
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L27
L38:
	;
	goto L1
L39:
	;
	return int32(0)
L40:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(_a_F_plpgsql_recognize_err_condition_2), v7)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_plpgsql_recognize_err_condition_3), int32(2164), int32(_a_F_plpgsql_recognize_err_condition_4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
