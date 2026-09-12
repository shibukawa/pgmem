package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_polymorphic_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 <= int32(3830) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v149
L2:
	;
	v137 = F_format_type_be(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L48
	} else {
		goto L49
	}
L3:
	;
	if base.Ui32(int32(2)) < base.Ui32(l0-int32(_a_F_check_valid_polymorphic_signature_0)) {
		v149 = v4
		goto L1
	} else {
		goto L41
	}
L4:
	;
	switch l0 - int32(2277) {
	case 0, 6:
		goto L7
	case 1, 2, 3, 4, 5:
		goto L3
	default:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	switch l0 - int32(_a_F_check_valid_polymorphic_signature_1) {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		goto L26
	}
L7:
	;
	v21 = int32(_a_F_check_valid_polymorphic_signature_2)
	if l2 <= int32(0) {
		v135 = v21
		goto L2
	} else {
		goto L11
	}
L8:
	;
	if l0 == int32(2776) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l0 != int32(3500) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v28 = v4
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+v28<<(uint(int32(2))%32))))
	if v35 <= int32(3499) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v135 = v21
	goto L2
L14:
	;
	v49 = v28 + int32(1)
	if v49 != l2 {
		v28 = v49
		goto L12
	} else {
		goto L23
	}
L15:
	;
	switch v35 - int32(2277) {
	case 0, 6:
		v149 = v4
		goto L1
	case 1, 2, 3, 4, 5:
		goto L14
	default:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v35 == int32(3500) {
		v149 = v4
		goto L1
	} else {
		goto L20
	}
L18:
	;
	if v35 != int32(2776) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v149 = v4
	goto L1
L20:
	;
	if v35 == int32(3831) {
		v149 = v4
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v35 == int32(_a_F_check_valid_polymorphic_signature_1) {
		v149 = v4
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	goto L13
L24:
	;
	v79 = int32(_a_F_check_valid_polymorphic_signature_3)
	if l2 <= int32(0) {
		v135 = v79
		goto L2
	} else {
		goto L35
	}
L25:
	;
	v57 = int32(_a_F_check_valid_polymorphic_signature_4)
	if l2 <= int32(0) {
		v135 = v57
		goto L2
	} else {
		goto L29
	}
L26:
	;
	if l0 == int32(_a_F_check_valid_polymorphic_signature_5) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if l0 != int32(3831) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v63 = v4
	goto L30
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1+v63<<(uint(int32(2))%32))))
	if v71 == int32(3831) {
		v149 = v4
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v135 = v57
	goto L2
L32:
	;
	if v71 == int32(_a_F_check_valid_polymorphic_signature_1) {
		v149 = v4
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v77 = v63 + int32(1)
	if v77 != l2 {
		v63 = v77
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v85 = v4
	goto L36
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85<<(uint(int32(2))%32))))
	if v93 == int32(_a_F_check_valid_polymorphic_signature_6) {
		v149 = v4
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v135 = v79
	goto L2
L38:
	;
	if v93 == int32(_a_F_check_valid_polymorphic_signature_5) {
		v149 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v99 = v85 + int32(1)
	if v99 != l2 {
		v85 = v99
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v105 = int32(_a_F_check_valid_polymorphic_signature_7)
	if l2 <= int32(0) {
		v135 = v105
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v111 = v4
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1+v111<<(uint(int32(2))%32))))
	if base.Ui32(v119-int32(_a_F_check_valid_polymorphic_signature_0)) < base.Ui32(int32(4)) {
		v149 = v4
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v135 = v105
	goto L2
L45:
	;
	if v119 == int32(_a_F_check_valid_polymorphic_signature_6) {
		v149 = v4
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v127 = v111 + int32(1)
	if v127 != l2 {
		v111 = v127
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	return int32(0)
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v137
	v142 = F_psprintf(m, v135, v11)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v149 = v142
	goto L1
}
func F_validOperatorName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	v2 = int32(0)
	if l0&int32(3) == v2 {
		v28 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v298
L2:
	;
	if base.Ui32(v61+int32(-64)) < base.Ui32(int32(-63)) {
		v298 = v2
		goto L1
	} else {
		goto L19
	}
L3:
	;
	v61 = v53 - l0
	goto L2
L4:
	;
	v32 = v28
	goto L13
L5:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v61 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v17 = l0
	goto L9
L9:
	;
	v21 = v17 + int32(1)
	if v21&int32(3) == int32(0) {
		v28 = v21
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v53 = v21
	goto L3
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 != 0 {
		v17 = v21
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 == v41 {
		v32 = v32 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v47 = v32
	goto L16
L15:
	;
	goto L14
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		v47 = v47 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v53 = v47
	goto L3
L18:
	;
	goto L17
L19:
	;
	v66 = int32(_a_F_validOperatorName_0)
	v70 = m.G0
	v72 = v70 - int32(32)
	v73 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+24)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v72)+16)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[0])))
	if v81 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v149 != v61 {
		v298 = v2
		goto L1
	} else {
		goto L41
	}
L21:
	;
	v149 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[1])))
	if v85 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v89 = l0
	goto L27
L25:
	;
	goto L26
L26:
	;
	v99 = v66
	v100 = v81
	goto L30
L27:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v95 == v81 {
		v89 = v89 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v149 = v89 - l0
	goto L20
L29:
	;
	goto L28
L30:
	;
	v107 = v72 + int32(base.Ui32(v100)>>(uint(int32(3))%32))&int32(28)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 | v109<<(uint(v100)%32)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v113 != 0 {
		v99 = v99 + v109
		v100 = v113
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v116 == int32(0) {
		v141 = l0
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v149 = v141 - l0
	goto L20
L34:
	;
	v120 = l0
	v121 = v116
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(base.Ui32(v121)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v129)>>(uint(v121)%32))&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = v137
	goto L33
L37:
	;
	v141 = v120
	goto L33
L38:
	;
	goto L39
L39:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v137 = v120 + int32(1)
	if v135 != 0 {
		v120 = v137
		v121 = v135
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v152 = F_strstr(m, l0, int32(_a_F_validOperatorName_1))
	mBase = m.M
	if v152 != 0 {
		v298 = v2
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v154 = F_strstr(m, l0, int32(_a_F_validOperatorName_2))
	mBase = m.M
	if v154 != 0 {
		v298 = v2
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v61) < base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v286 = int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v287 != int32(33) {
		v298 = v286
		goto L1
	} else {
		goto L77
	}
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v61-int32(1)))))
	switch v160 - int32(43) {
	case 0, 2:
		goto L46
	default:
		goto L44
	}
L46:
	;
	v166 = v61 - int32(2)
	goto L47
L47:
	;
	v169 = int32(_a_F_validOperatorName_3)
	v171 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v166))))
	v172 = int32(11)
	goto L52
L48:
	;
	v298 = v2
	goto L1
L49:
	;
	if v275 != 0 {
		goto L44
	} else {
		goto L75
	}
L50:
	;
	v275 = int32(0)
	goto L49
L51:
	;
	v253 = v246
	v255 = v248
	goto L69
L52:
	;
	goto L60
L60:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[2])))
	if v209 == v171&int32(255) {
		v239 = v169
		v241 = v172
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v241 == int32(0) {
		goto L50
	} else {
		goto L68
	}
L62:
	;
	goto L63
L63:
	;
	v219 = v169
	v221 = v172
	goto L64
L64:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v226 = v225 ^ v171&int32(255)*int32(16843009)
	v229 = int32(-2139062144)
	if (int32(16843008)-v226|v226)&v229 != v229 {
		v246 = v219
		v248 = v221
		goto L51
	} else {
		goto L66
	}
L65:
	;
	v239 = v234
	v241 = v236
	goto L61
L66:
	;
	v233 = int32(4)
	v234 = v219 + v233
	v236 = v221 - v233
	if base.Ui32(int32(3)) < base.Ui32(v236) {
		v219 = v234
		v221 = v236
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v246 = v239
	v248 = v241
	goto L51
L69:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v171&int32(255) == v258 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L50
L71:
	;
	v275 = v253
	goto L49
L72:
	;
	goto L73
L73:
	;
	v260 = int32(1)
	v263 = v255 - v260
	if v263 != 0 {
		v253 = v253 + v260
		v255 = v263
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v276 = int32(0)
	if base.B2i32(v166 <= v276) == v276 {
		v166 = v166 - int32(1)
		goto L47
	} else {
		goto L76
	}
L76:
	;
	goto L48
L77:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v290 != int32(61) {
		v298 = v286
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v298 = base.B2i32(v293 != int32(0))
	goto L1
}
