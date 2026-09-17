package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mark_partial_aggref(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l1
	if l1&int32(2) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v7 == int32(2281) {
			v10 = int32(17)
		} else {
			v10 = v7
		}
		if l1&int32(4) != 0 {
			v13 = v10
		} else {
			v13 = v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13
	} else {
	}
	return
}
func F_matchPartialInPendingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v29 int32
	_ = v29
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
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
	goto L3
L3:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v29 = l2
	goto L7
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v29<<(uint(int32(2))%32))))
	v47 = l1 + v44&int32(_a_F_matchPartialInPendingList_0)
	v48 = F_gintuple_get_attrnum(m, l0, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	return int32(0)
L10:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+20)))
	if v48 != v52 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L13
L13:
	;
	v57 = v29 - int32(1)
	v58 = l7 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = F_gintuple_get_key(m, l0, v47, l6+v57)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v57))))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5+v57<<(uint(int32(2))%32)))) = v66
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v69)
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+20)))
	v77 = v75 - int32(1)
	v81 = int32(2)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_matchPartialInPendingList_1)+v77<<(uint(v81)%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l5+v57<<(uint(v81)%32))))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+12)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v92 = F_FunctionCall4Coll(m, l0+int32(_a_F_matchPartialInPendingList_2)+v77*int32(28), v84, v85, v89, v90, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	if v92 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(1)
L23:
	;
	goto L24
L24:
	;
	if v92 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v101 = v29 + int32(1)
	if v101 != l3 {
		v29 = v101
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L8
L28:
	;
	goto L27
}
func F_try_partial_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 float64
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 float64
	_ = v263
	var v270 float64
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	v11 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v11
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(112)
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l6 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v25 = v16 + int32(108)
	if l6 == v23 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v105 = v11
	goto L8
L8:
	;
	if l7 != 0 {
		goto L44
	} else {
		goto L45
	}
L9:
	;
	if v103 != 0 {
		goto L41
	} else {
		goto L42
	}
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v91
	v103 = int32(1)
	goto L9
L11:
	;
	if l6 != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l6 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
	v103 = int32(1)
	goto L9
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
	v103 = int32(1)
	goto L9
L16:
	;
	goto L17
L17:
	;
	if v23 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v43
	v103 = v43
	goto L9
L19:
	;
	goto L20
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v47 = int32(0)
	if v47 < v46 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v50 = v46
	goto L23
L22:
	;
	v50 = v47
	goto L23
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v56 = v22
	goto L24
L24:
	;
	if v56 < v51 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v67 = v63 + v56<<(uint(int32(2))%32)
	goto L28
L27:
	;
	v67 = int32(0)
	goto L28
L28:
	;
	if v56 == v50 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v50
	v103 = base.B2i32(v67 == int32(0))
	goto L9
L30:
	;
	goto L31
L31:
	;
	v73 = base.B2i32(v67 == int32(0))
	if v67 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v56
	v103 = v73
	goto L9
L33:
	;
	goto L34
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v77 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v56
	v103 = v73
	goto L9
L36:
	;
	goto L37
L37:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v77+v56<<(uint(int32(2))%32))))
	if v81 != v85 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v56
	v103 = int32(0)
	goto L9
L39:
	;
	v56 = v56 + int32(1)
	goto L24
L41:
	;
	v104 = v22
	goto L43
L42:
	;
	v104 = l6
	goto L43
L43:
	;
	v105 = v104
	goto L8
L44:
	;
	v108 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if l7 == v109 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v165 = int32(0)
	goto L46
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	F_initial_cost_mergejoin(m, l0, v16+int32(8), l8, l5, l2, l3, v105, v165, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	if v162 != 0 {
		goto L65
	} else {
		goto L66
	}
L48:
	;
	v162 = int32(1)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v118 = v108
	goto L52
L51:
	;
	v162 = v154
	goto L47
L52:
	;
	v122 = int32(0)
	if l7 == v122 {
		v132 = v122
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v154 = int32(0)
	goto L51
L54:
	;
	if v109 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v126 <= v118 {
		v132 = int32(0)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v132 = v128 + v118<<(uint(int32(2))%32)
	goto L54
L57:
	;
	v138 = base.B2i32(v132 == int32(0))
	if v132 == int32(0) {
		v154 = v138
		goto L51
	} else {
		goto L62
	}
L58:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v118 < v133 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v162 = base.B2i32(v132 == int32(0))
	goto L47
L61:
	;
	goto L60
L62:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v141 == int32(0) {
		v154 = v138
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v118<<(uint(int32(2))%32)+v141)))
	if v148 == v150 {
		v118 = v118 + int32(1)
		goto L52
	} else {
		goto L64
	}
L64:
	;
	goto L53
L65:
	;
	v163 = v108
	goto L67
L66:
	;
	v163 = l7
	goto L67
L67:
	;
	v165 = v163
	goto L46
L68:
	;
	return
L69:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v171 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v177 == v171 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v317 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L71:
	;
	v317 = v302
	goto L70
L72:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v234 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L73:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v180 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v191 = v171
	goto L75
L75:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v191<<(uint(int32(2))%32))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+64))
	v201 = F_compare_pathkeys(m, l4, v200)
	mBase = m.M
	if v201 == int32(3) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L72
L77:
	;
	v221 = v191 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v221 < v222 {
		v191 = v221
		goto L75
	} else {
		goto L82
	}
L78:
	;
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v199)+56))
	v208 = int32(0)
	v212 = base.B2i32(base.F64_gt(v170, base.F64_mul(v204, float64(1.01))) == v208) | base.B2i32(v201 == int32(1))
	if v212 == v208 {
		v302 = v212
		goto L71
	} else {
		goto L79
	}
L79:
	;
	if v201 == int32(2) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	if base.F64_lt(base.F64_mul(v170, float64(1.01)), v204) != 0 {
		v302 = v212
		goto L71
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	goto L76
L83:
	;
	v317 = int32(1)
	goto L70
L84:
	;
	goto L85
L85:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v238 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v317 = int32(1)
	goto L70
L87:
	;
	goto L88
L88:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v250 = int32(0)
	goto L89
L89:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254+v250<<(uint(int32(2))%32))))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+40))
	if v169 != v259 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v302 = v292
	goto L71
L91:
	;
	if v243 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	if v259 <= v169 {
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v263 = *(*float64)(unsafe.Add(mBase, uint32(v258)+56))
	if base.F64_le(v170, base.F64_mul(v263, float64(1.01))) == int32(0) {
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v317 = int32(1)
	goto L70
L96:
	;
	v317 = int32(1)
	goto L70
L97:
	;
	v292 = int32(1)
	v294 = v250 + v292
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v294 < v295 {
		v250 = v294
		goto L89
	} else {
		goto L110
	}
L98:
	;
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v258)+48))
	if base.F64_gt(v170, base.F64_mul(v270, float64(1.01))) == int32(0) {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	if v277 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	v280 = int32(0)
	goto L104
L103:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)+64))
	v280 = v279
	goto L104
L104:
	;
	v281 = F_compare_pathkeys(m, l4, v280)
	mBase = m.M
	if v281&int32(-3) != 0 {
		goto L97
	} else {
		goto L105
	}
L105:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	if v285 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v288 = v286
	goto L108
L107:
	;
	v288 = int32(0)
	goto L108
L108:
	;
	v289 = F_bms_equal(m, int32(0), v288)
	mBase = m.M
	if v289 != 0 {
		v302 = int32(0)
		goto L71
	} else {
		goto L109
	}
L109:
	;
	goto L97
L110:
	;
	goto L90
L111:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	v325 = F_create_mergejoin_path(m, l0, l1, l8, v16+int32(8), l9, l2, l3, v322, l4, int32(0), l5, v105, v165, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L68
	} else {
		goto L112
	}
L112:
	;
	F_add_partial_path(m, l1, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L68
	} else {
		goto L113
	}
L113:
	;
	goto L1
}
