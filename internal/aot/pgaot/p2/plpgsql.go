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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
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
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[0]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v54
	v57 = F_makeParamList(m, v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v57
		*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(_a_F_plpgsql_estate_setup_0)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = l0
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = int32(_a_F_plpgsql_estate_setup_1)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v68 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v68
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = int32(_a_F_plpgsql_estate_setup_2)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v68
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = v77
		v80 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[2]))
		if v80 == v68 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(103079215120)
			v89 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_3), int32(16), v11, int32(40))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[2])) = v89
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v96
					v101 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_4), int32(16), v11, int32(1064))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						v117 = v101
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v104
					v107 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6]))
					if v107 != 0 {
						v117 = v107
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
						v114 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_10), int32(16), v11, int32(40))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6])) = v114
							v117 = v114
							*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
							v121 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
							if l4 != 0 {
								v125 = l4
							} else {
								v125 = v120
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
							v129 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
							v131 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
							*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
				v96 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v96
				v101 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_4), int32(16), v11, int32(1064))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					v117 = v101
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
					v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
					if l4 != 0 {
						v125 = l4
					} else {
						v125 = v120
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					v129 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v104
				v107 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6]))
				if v107 != 0 {
					v117 = v107
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
					v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
					if l4 != 0 {
						v125 = l4
					} else {
						v125 = v120
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
					v129 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
					v131 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
					*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953488)
					v114 = F_hash_create(m, int32(_a_F_plpgsql_estate_setup_10), int32(16), v11, int32(40))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[6])) = v114
						v117 = v114
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[3]))
						v121 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v121
						if l4 != 0 {
							v125 = l4
						} else {
							v125 = v120
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_estate_setup[1]))
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v129
						v131 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v129
						*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v131
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
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
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
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v14 = v11 + int32(32)
	F_plpgsql_estate_setup(m, v14, l0, v3, v3, v3)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_plpgsql_exec_trigger_0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(_a_F_plpgsql_exec_trigger_1)
	v27 = int32(_a_F_plpgsql_exec_trigger_2)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0])) = v11 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v14
	F_copy_plpgsql_datums(m, v14, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+484))
	v39 = int32(2)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(v39)%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+480))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37+v43<<(uint(v39)%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v51 = F_make_expanded_record_from_tupdesc(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v55 = F_make_expanded_record_from_exprecord(m, v51, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+36)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v58&int32(4) == int32(0) {
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
	switch v58&int32(3) - int32(1) {
	case 0:
		v140 = v55
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
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v142 = int32(0)
	F_expanded_record_set_tuple(m, v140, v141, v142, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v140 = v139
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_trigger_3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L26
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v69 = int32(0)
	F_expanded_record_set_tuple(m, v67, v68, v69, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v75 = int32(0)
	F_expanded_record_set_tuple(m, v73, v74, v75, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v79 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+17)))
	if v82 != int32(1) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v85&int32(24) != int32(8) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v90 <= int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v96 = int32(0)
	v97 = v90
	goto L18
L18:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v97<<(uint(int32(4))%32)+v96*int32(100))+110)))
	if v108 != int32(115) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L6
L20:
	;
	if v123 < v124 {
		v96 = v123
		v97 = v124
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v123 = v96 + int32(1)
	v124 = v97
	goto L20
L22:
	;
	goto L23
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v114 = int32(1)
	v115 = v96 + v114
	v116 = int32(0)
	F_expanded_record_set_field_internal(m, v113, v115, v116, v114, v116, v116)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v123 = v115
	v124 = v122
	goto L20
L25:
	;
	goto L19
L26:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_trigger_4), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1030), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(_a_F_plpgsql_exec_trigger_7)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v325+v326<<(uint(int32(2))%32))))
	v331 = int32(0)
	F_assign_simple_var(m, v11+int32(32), v330, v331, v331, v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L84
	}
L31:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v156 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v158 = F_palloc(m, int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v237 != 0 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v158)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+4)) = v164
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v168)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v158)+16)) = base.F64_convert_i64_s(v169)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v172
	if v161 == int32(0) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	if v178 == int32(0) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)+36))
	if v181 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_register_ENR(m, v231, v158)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L57
	}
L39:
	;
	v182 = int32(0)
	if v181 == v182 {
		v219 = v182
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v226 = F_palloc0(m, int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L56
	}
L42:
	;
	if v219 != 0 {
		goto L30
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v187 == int32(0) {
		v219 = v182
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v190 <= int32(0) {
		v219 = v182
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v193 = int32(0)
	if v193 < v190 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v196 = v190
	goto L49
L48:
	;
	v196 = v193
	goto L49
L49:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v199 = int32(0)
	goto L50
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v197+v199<<(uint(int32(2))%32))))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = F_strcmp(m, v208, v161)
	mBase = m.M
	if v209 == int32(0) {
		v219 = v207
		goto L43
	} else {
		goto L52
	}
L51:
	;
	v219 = int32(0)
	goto L43
L52:
	;
	v213 = v199 + int32(1)
	if v213 != v196 {
		v199 = v213
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+36))
	if v223 != 0 {
		v231 = v223
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L41
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+36)) = v226
	v231 = v226
	goto L38
L57:
	;
	goto L34
L58:
	;
	v239 = F_palloc(m, int32(32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
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
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v245
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v249)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v239)+16)) = base.F64_convert_i64_s(v250)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = v253
	if v242 == int32(0) {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	if v259 == int32(0) {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+36))
	if v262 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_register_ENR(m, v312, v239)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L83
	}
L65:
	;
	v263 = int32(0)
	if v262 == v263 {
		v300 = v263
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v307 = F_palloc0(m, int32(4))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L82
	}
L68:
	;
	if v300 != 0 {
		goto L30
	} else {
		goto L80
	}
L69:
	;
	goto L68
L70:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v268 == int32(0) {
		v300 = v263
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v271 <= int32(0) {
		v300 = v263
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v274 = int32(0)
	if v274 < v271 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v277 = v271
	goto L75
L74:
	;
	v277 = v274
	goto L75
L75:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v280 = int32(0)
	goto L76
L76:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v278+v280<<(uint(int32(2))%32))))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v290 = F_strcmp(m, v289, v242)
	mBase = m.M
	if v290 == int32(0) {
		v300 = v288
		goto L69
	} else {
		goto L78
	}
L77:
	;
	v300 = int32(0)
	goto L69
L78:
	;
	v294 = v280 + int32(1)
	if v294 != v277 {
		v280 = v294
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+36))
	if v304 != 0 {
		v312 = v304
		goto L64
	} else {
		goto L81
	}
L81:
	;
	goto L67
L82:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+36)) = v307
	v312 = v307
	goto L64
L83:
	;
	goto L60
L84:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	if v338 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[3]))
	if v377 != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(0)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v343
	v375 = v343
	goto L85
L87:
	;
	goto L88
L88:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v345 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	if v367 == int32(0) {
		v375 = v366
		goto L85
	} else {
		goto L95
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v350
	v365 = v338
	v366 = v350
	goto L89
L91:
	;
	goto L92
L92:
	;
	m.T0[v345].(func(*base.Module, int32, int32))(m, v11+int32(32), l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v359 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v361
	if v358 == v359 {
		v375 = v361
		goto L85
	} else {
		goto L94
	}
L94:
	;
	v365 = v358
	v366 = v361
	goto L89
L95:
	;
	m.T0[v367].(func(*base.Module, int32, int32))(m, v11+int32(32), v366)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v375 = v366
	goto L85
L97:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v381 = v11 + int32(32)
	v382 = F_exec_stmt_block(m, v381, v375)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v386 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(0)
	if v382 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	if v389 == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	m.T0[v389].(func(*base.Module, int32, int32))(m, v381, v375)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
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
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L157
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(_a_F_plpgsql_exec_trigger_8)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+61)))
	if v401 == int32(1) {
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_trigger_3))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L153
	}
L110:
	;
	v404 = int32(0)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	if v405 != 0 {
		v480 = v404
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[2]))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	if v485 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L112:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v406&int32(4) == int32(0) {
		v480 = v404
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v412 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v447 = F_pg_detoast_datum(m, v411)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L130
	}
L115:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v415&int32(254) != int32(2) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v411)+2))
	v421 = F_expanded_record_get_tuple(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)+44))
	if v423 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v426 = F_expanded_record_fetch_tupdesc(m, v420)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v428 = v423
	goto L120
L120:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+52))
	if v428 == v430 {
		v439 = v421
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v428 = v426
	goto L120
L122:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v439 == v441 {
		v480 = v439
		goto L111
	} else {
		goto L127
	}
L123:
	;
	v433 = F_convert_tuples_by_position(m, v428, v430, int32(_a_F_plpgsql_exec_trigger_9))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v433 == int32(0) {
		v439 = v421
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v437 = F_execute_attr_map_tuple(m, v421, v433)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v439 = v437
	goto L122
L127:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v439 == v443 {
		v480 = v439
		goto L111
	} else {
		goto L128
	}
L128:
	;
	v445 = F_SPI_copytuple(m, v439)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v480 = v445
	goto L111
L130:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v447
	v451 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v451
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)) = uint16(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(base.Ui32(v449) >> (uint(int32(2)) % 32))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v462 = F_lookup_rowtype_tupdesc(m, v460, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
	v467 = F_convert_tuples_by_position(m, v462, v465, int32(_a_F_plpgsql_exec_trigger_9))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v467 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v469 = F_execute_attr_map_tuple(m, v11, v467)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v471 = v11
	goto L135
L135:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	if int32(0) <= v472 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v471 = v469
	goto L135
L137:
	;
	F_DecrTupleDescRefCount(m, v462)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v477 = F_SPI_copytuple(m, v471)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v480 = v477
	goto L111
L142:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[4]))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+8))
	F_pfree(m, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485)+8))
	if v488 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	m.T0[v488].(func(*base.Module, int32, int32))(m, v11+int32(32), l0)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[4])) = v498
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	F_FreeExprContext(m, v503, int32(1))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v507 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+152)) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	if v509 == v507 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_trigger[0])) = v524
	m.G0 = v11 + int32(176)
	return v480
L149:
	;
	F_SPI_freetuptable(m, v509)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v514 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	if v516 == v514 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v516)+20))
	F_MemoryContextReset(m, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
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
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_trigger_10), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1059), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
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
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_trigger_11), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_trigger_5), int32(1067), int32(_a_F_plpgsql_exec_trigger_6))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8<<(uint(int32(2))%32)-int32(4))))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
	if v17 == int32(0) {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v20 != int32(1) {
			v29 = v15
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v23 != int32(3) {
				v29 = v15
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
				v29 = v26 + int32(18)
			}
		}
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
		v31 = v29
		v32 = v30
	} else {
		v31 = v15
		v32 = int32(1)
	}
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v31
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v32)
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
	F_errfinish(m, int32(_a_F_plpgsql_recognize_err_condition_3), int32(2141), int32(_a_F_plpgsql_recognize_err_condition_4))
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
