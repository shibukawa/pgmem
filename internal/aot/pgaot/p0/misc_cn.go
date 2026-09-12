package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CNStoBIG5(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	v3 = int32(0)
	v5 = l0 & int32(_a_F_CNStoBIG5_0)
	switch l1 - int32(149) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L12
	}
L1:
	;
	return v317 & int32(_a_F_CNStoBIG5_1)
L2:
	;
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314))))
	v317 = v315
	goto L1
L3:
	;
	if v5 != int32(_a_F_CNStoBIG5_2) {
		v317 = v3
		goto L1
	} else {
		goto L96
	}
L4:
	;
	v314 = int32(_a_F_CNStoBIG5_3)
	goto L2
L5:
	;
	v314 = int32(_a_F_CNStoBIG5_4)
	goto L2
L6:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[0])))
	v317 = v308
	goto L1
L7:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[1])))
	v317 = v306
	goto L1
L8:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[2])))
	v317 = v304
	goto L1
L9:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[3])))
	v317 = v302
	goto L1
L10:
	;
	v180 = int32(47)
	v182 = int32(23)
	v183 = int32(0)
	goto L64
L11:
	;
	v49 = int32(24)
	v51 = int32(12)
	v52 = int32(0)
	goto L30
L12:
	;
	switch l1 - int32(246) {
	case 0:
		goto L13
	case 1:
		goto L14
	default:
		v317 = v3
		goto L1
	}
L13:
	;
	if v5 <= int32(_a_F_CNStoBIG5_5) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	switch v5 - int32(_a_F_CNStoBIG5_6) {
	case 0:
		v314 = int32(_a_F_CNStoBIG5_7)
		goto L2
	case 1:
		goto L5
	case 2, 3, 4, 5, 6:
		v317 = v3
		goto L1
	case 7:
		goto L4
	default:
		goto L3
	}
L15:
	;
	if v5 == int32(_a_F_CNStoBIG5_8) {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v5 <= int32(_a_F_CNStoBIG5_9) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v5 == int32(_a_F_CNStoBIG5_10) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	if v5 != int32(_a_F_CNStoBIG5_11) {
		v317 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[4])))
	v317 = v24
	goto L1
L21:
	;
	if v5 == int32(_a_F_CNStoBIG5_12) {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v5 == int32(_a_F_CNStoBIG5_13) {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	if v5 != int32(_a_F_CNStoBIG5_14) {
		v317 = v3
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[5])))
	v317 = v32
	goto L1
L26:
	;
	if v5 != int32(_a_F_CNStoBIG5_15) {
		v317 = v3
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[6])))
	v317 = v38
	goto L1
L28:
	;
	v317 = v167 & int32(_a_F_CNStoBIG5_1)
	goto L1
L29:
	;
	goto L28
L30:
	;
	v57 = v51 << (uint(int32(2)) % 32)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_CNStoBIG5[7]))))
	v60 = base.B2i32(base.Ui32(v5) < base.Ui32(v59))
	if base.Ui32(v5) < base.Ui32(v59) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v167 = int32(0)
	goto L29
L32:
	;
	goto L31
L33:
	;
	if base.Ui32(v5) < base.Ui32(v59) {
		goto L55
	} else {
		goto L56
	}
L34:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_CNStoBIG5[8]))))
	if base.Ui32(v61) <= base.Ui32(v5) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_CNStoBIG5[9]))))
	if v63 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v68 = v5 - v59&int32(_a_F_CNStoBIG5_16)
	if base.Ui32(int32(_a_F_CNStoBIG5_17)) <= base.Ui32(v5) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v71 = int32(255)
	v72 = v5 & v71
	v74 = v59 & v71
	v84 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v74))
	if base.Ui32(int32(160)) < base.Ui32(v74) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v111 = int32(255)
	v122 = v63 & v111
	if base.Ui32(int32(160)) < base.Ui32(v122) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v85 = int32(0)
	goto L42
L41:
	;
	v85 = int32(-34)
	goto L42
L42:
	;
	if base.Ui32(int32(160)) < base.Ui32(v74) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v88 = int32(34)
	goto L45
L44:
	;
	v88 = int32(0)
	goto L45
L45:
	;
	if base.Ui32(int32(160)) < base.Ui32(v72) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v91 = v85
	goto L48
L47:
	;
	v91 = v88
	goto L48
L48:
	;
	v96 = int32(33)
	v97 = v72 - v74 + v68>>(uint(int32(8))%32)*int32(157) + v91 + v63&int32(255) - v96
	v98 = int32(94)
	v99 = base.I32_div_s(v97, v98)
	v167 = v97 - v99*v98 + v63&int32(_a_F_CNStoBIG5_16) + v99<<(uint(int32(8))%32) + v96
	goto L29
L49:
	;
	v128 = int32(_a_F_CNStoBIG5_18)
	goto L51
L50:
	;
	v128 = int32(_a_F_CNStoBIG5_19)
	goto L51
L51:
	;
	v129 = v5&v111 - v59&v111 + int32(base.Ui32(v68)>>(uint(int32(8))%32))*int32(94) + v122 + v128
	v131 = int32(157)
	v132 = base.I32_div_s(base.I32_extend16_s(v129), v131)
	v135 = v129 - v132*v131
	if int32(62) < base.I32_extend16_s(v135) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = int32(98)
	goto L54
L53:
	;
	v147 = int32(64)
	goto L54
L54:
	;
	v167 = v135 + v63&int32(_a_F_CNStoBIG5_16) + v132<<(uint(int32(8))%32) + v147
	goto L29
L55:
	;
	v151 = v52
	goto L57
L56:
	;
	v151 = v51 + int32(1)
	goto L57
L57:
	;
	if base.Ui32(v5) < base.Ui32(v59) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = v51 - int32(1)
	goto L60
L59:
	;
	v154 = v49
	goto L60
L60:
	;
	if v151 <= v154 {
		v49 = v154
		v51 = (v151 + v154) >> (uint(int32(1)) % 32)
		v52 = v151
		goto L30
	} else {
		goto L61
	}
L61:
	;
	goto L32
L62:
	;
	v317 = v298 & int32(_a_F_CNStoBIG5_1)
	goto L1
L63:
	;
	goto L62
L64:
	;
	v188 = v182 << (uint(int32(2)) % 32)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[10]))))
	v191 = base.B2i32(base.Ui32(v5) < base.Ui32(v190))
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v298 = int32(0)
	goto L63
L66:
	;
	goto L65
L67:
	;
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L89
	} else {
		goto L90
	}
L68:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[11]))))
	if base.Ui32(v192) <= base.Ui32(v5) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[12]))))
	if v194 == int32(0) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v199 = v5 - v190&int32(_a_F_CNStoBIG5_16)
	if base.Ui32(int32(_a_F_CNStoBIG5_17)) <= base.Ui32(v5) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = int32(255)
	v203 = v5 & v202
	v205 = v190 & v202
	v215 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v205))
	if base.Ui32(int32(160)) < base.Ui32(v205) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v242 = int32(255)
	v253 = v194 & v242
	if base.Ui32(int32(160)) < base.Ui32(v253) {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v216 = int32(0)
	goto L76
L75:
	;
	v216 = int32(-34)
	goto L76
L76:
	;
	if base.Ui32(int32(160)) < base.Ui32(v205) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v219 = int32(34)
	goto L79
L78:
	;
	v219 = int32(0)
	goto L79
L79:
	;
	if base.Ui32(int32(160)) < base.Ui32(v203) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v222 = v216
	goto L82
L81:
	;
	v222 = v219
	goto L82
L82:
	;
	v227 = int32(33)
	v228 = v203 - v205 + v199>>(uint(int32(8))%32)*int32(157) + v222 + v194&int32(255) - v227
	v229 = int32(94)
	v230 = base.I32_div_s(v228, v229)
	v298 = v228 - v230*v229 + v194&int32(_a_F_CNStoBIG5_16) + v230<<(uint(int32(8))%32) + v227
	goto L63
L83:
	;
	v259 = int32(_a_F_CNStoBIG5_18)
	goto L85
L84:
	;
	v259 = int32(_a_F_CNStoBIG5_19)
	goto L85
L85:
	;
	v260 = v5&v242 - v190&v242 + int32(base.Ui32(v199)>>(uint(int32(8))%32))*int32(94) + v253 + v259
	v262 = int32(157)
	v263 = base.I32_div_s(base.I32_extend16_s(v260), v262)
	v266 = v260 - v263*v262
	if int32(62) < base.I32_extend16_s(v266) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v278 = int32(98)
	goto L88
L87:
	;
	v278 = int32(64)
	goto L88
L88:
	;
	v298 = v266 + v194&int32(_a_F_CNStoBIG5_16) + v263<<(uint(int32(8))%32) + v278
	goto L63
L89:
	;
	v282 = v183
	goto L91
L90:
	;
	v282 = v182 + int32(1)
	goto L91
L91:
	;
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v285 = v182 - int32(1)
	goto L94
L93:
	;
	v285 = v180
	goto L94
L94:
	;
	if v282 <= v285 {
		v180 = v285
		v182 = (v282 + v285) >> (uint(int32(1)) % 32)
		v183 = v282
		goto L64
	} else {
		goto L95
	}
L95:
	;
	goto L66
L96:
	;
	v314 = int32(_a_F_CNStoBIG5_20)
	goto L2
}
func F_cntsize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34 + v35&int32(4095) + int32(1)
	goto L3
L7:
	;
	v22 = int32(0)
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v22<<(uint(int32(2))%32))))
	F_cntsize(m, v27, l1, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v31 = v22 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 < v32 {
		v22 = v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
