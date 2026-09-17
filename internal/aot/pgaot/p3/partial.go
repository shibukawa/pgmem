package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_partial_path(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_add_partial_path[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_pfree(m, l1)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L65
	}
L7:
	;
	v20 = v3
	v22 = v17
	v27 = v3
	goto L10
L8:
	;
	v182 = v3
	v185 = int32(0)
	goto L9
L9:
	;
	v186 = F_list_insert_nth(m, v185, v182, l1)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L64
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v20 < v30 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v182 = v168
	v185 = v171
	goto L9
L12:
	;
	v32 = int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v20<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v33 == v39 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v168 = v27
	goto L14
L14:
	;
	goto L11
L15:
	;
	if v153 != 0 {
		v20 = v151 + int32(1)
		v22 = v153
		v27 = v156
		goto L10
	} else {
		goto L63
	}
L16:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v38)+56))
	if base.F64_ge(v142, v143) != 0 {
		goto L58
	} else {
		goto L59
	}
L17:
	;
	if v99 == int32(3) {
		v137 = v32
		goto L16
	} else {
		goto L41
	}
L18:
	;
	v99 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v48 = int32(0)
	goto L23
L21:
	;
	if v88 != 0 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	v83 = int32(0)
	if v70 != 0 {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v52 = int32(0)
	if v33 == v52 {
		v62 = v52
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v99 = int32(3)
	goto L17
L25:
	;
	if v39 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v56 <= v48 {
		v62 = int32(0)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v62 = v58 + v48<<(uint(int32(2))%32)
	goto L25
L28:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if base.B2i32(v62 == v68)|base.B2i32(v70 == v68) != 0 {
		goto L22
	} else {
		goto L33
	}
L29:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v48 < v63 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v65 = int32(0)
	v88 = base.B2i32(v62 == v65)
	v90 = v65
	goto L21
L32:
	;
	goto L31
L33:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70+v48<<(uint(int32(2))%32))))
	if v78 == v80 {
		v48 = v48 + int32(1)
		goto L23
	} else {
		goto L34
	}
L34:
	;
	goto L24
L35:
	;
	v87 = int32(2)
	goto L37
L36:
	;
	v87 = v83
	goto L37
L37:
	;
	v88 = base.B2i32(v62 == v83)
	v90 = v87
	goto L21
L38:
	;
	v92 = v90
	goto L40
L39:
	;
	v92 = int32(1)
	goto L40
L40:
	;
	v99 = v92
	goto L17
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v102 != v103 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = F_list_delete_nth_cell(m, v129, v20)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L56
	}
L43:
	;
	if v102 <= v103 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v38)+56))
	if base.F64_gt(v107, base.F64_mul(v108, float64(1.01))) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v137 = int32(0)
	goto L16
L47:
	;
	v137 = base.B2i32(v99 == int32(1))
	goto L16
L48:
	;
	goto L49
L49:
	;
	if base.F64_gt(v108, base.F64_mul(v107, float64(1.01))) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	switch v99 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L6
	default:
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v99 == int32(2) {
		v137 = v32
		goto L16
	} else {
		goto L55
	}
L53:
	;
	if base.F64_gt(v108, base.F64_mul(v107, float64(1.0000000001))) != 0 {
		goto L42
	} else {
		goto L54
	}
L54:
	;
	v137 = int32(0)
	goto L16
L55:
	;
	goto L42
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v130
	F_pfree(m, v38)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v151 = v20 - int32(1)
	v153 = v130
	v156 = v27
	goto L15
L58:
	;
	if v137 == int32(0) {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v137 == int32(0) {
		goto L6
	} else {
		goto L62
	}
L61:
	;
	v151 = v20
	v153 = v22
	v156 = v20 + int32(1)
	goto L15
L62:
	;
	v151 = v20
	v153 = v22
	v156 = v27
	goto L15
L63:
	;
	v168 = v156
	goto L14
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v186
	return
L65:
	;
	return
}
func F_add_partial_path_precheck(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 float64
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 float64
	_ = v159
	var v167 float64
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v11 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v308
L2:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v127 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v14 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = v5
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v25<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if l3 == v34 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L2
L7:
	;
	v114 = v25 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v114 < v115 {
		v25 = v114
		goto L5
	} else {
		goto L36
	}
L8:
	;
	if v94 == int32(3) {
		goto L7
	} else {
		goto L32
	}
L9:
	;
	v94 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v43 = int32(0)
	goto L14
L12:
	;
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v78 = int32(0)
	if v65 != 0 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v47 = int32(0)
	if l3 == v47 {
		v57 = v47
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v94 = int32(3)
	goto L8
L16:
	;
	if v34 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v51 <= v43 {
		v57 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = v53 + v43<<(uint(int32(2))%32)
	goto L16
L19:
	;
	v63 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if base.B2i32(v57 == v63)|base.B2i32(v65 == v63) != 0 {
		goto L13
	} else {
		goto L24
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v43 < v58 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v60 = int32(0)
	v83 = base.B2i32(v57 == v60)
	v85 = v60
	goto L12
L23:
	;
	goto L22
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65+v43<<(uint(int32(2))%32))))
	if v73 == v75 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L15
L26:
	;
	v82 = int32(2)
	goto L28
L27:
	;
	v82 = v78
	goto L28
L28:
	;
	v83 = base.B2i32(v57 == v78)
	v85 = v82
	goto L12
L29:
	;
	v87 = v85
	goto L31
L30:
	;
	v87 = int32(1)
	goto L31
L31:
	;
	v94 = v87
	goto L8
L32:
	;
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	v101 = int32(0)
	v105 = base.B2i32(base.F64_gt(l2, base.F64_mul(v97, float64(1.01))) == v101) | base.B2i32(v94 == int32(1))
	if v105 == v101 {
		v308 = v105
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v94 == int32(2) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	if base.F64_lt(base.F64_mul(l2, float64(1.01)), v97) != 0 {
		v308 = v105
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L7
L36:
	;
	goto L6
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v132 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(1)
L41:
	;
	goto L42
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v145 = int32(0)
	goto L43
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v145<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+40))
	if l1 != v154 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v308 = v298
	goto L1
L45:
	;
	if v138 != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	if v154 <= l1 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v153)+56))
	if base.F64_le(l2, base.F64_mul(v159, float64(1.01))) == int32(0) {
		goto L45
	} else {
		goto L50
	}
L49:
	;
	return int32(1)
L50:
	;
	return int32(1)
L51:
	;
	v298 = int32(1)
	v300 = v145 + v298
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v300 < v301 {
		v145 = v300
		goto L43
	} else {
		goto L99
	}
L52:
	;
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v153)+48))
	if base.F64_gt(l2, base.F64_mul(v167, float64(1.01))) == int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	if v174 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v177 = int32(0)
	goto L58
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v153)+64))
	v177 = v176
	goto L58
L58:
	;
	if l3 == v177 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v237&int32(-3) != 0 {
		goto L51
	} else {
		goto L83
	}
L60:
	;
	v237 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v186 = int32(0)
	goto L65
L63:
	;
	if v226 != 0 {
		goto L80
	} else {
		goto L81
	}
L64:
	;
	v221 = int32(0)
	if v208 != 0 {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v190 = int32(0)
	if l3 == v190 {
		v200 = v190
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v237 = int32(3)
	goto L59
L67:
	;
	if v177 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v194 <= v186 {
		v200 = int32(0)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v200 = v196 + v186<<(uint(int32(2))%32)
	goto L67
L70:
	;
	v206 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	if base.B2i32(v200 == v206)|base.B2i32(v208 == v206) != 0 {
		goto L64
	} else {
		goto L75
	}
L71:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v186 < v201 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v203 = int32(0)
	v226 = base.B2i32(v200 == v203)
	v228 = v203
	goto L63
L74:
	;
	goto L73
L75:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v208+v186<<(uint(int32(2))%32))))
	if v216 == v218 {
		v186 = v186 + int32(1)
		goto L65
	} else {
		goto L76
	}
L76:
	;
	goto L66
L77:
	;
	v225 = int32(2)
	goto L79
L78:
	;
	v225 = v221
	goto L79
L79:
	;
	v226 = base.B2i32(v200 == v221)
	v228 = v225
	goto L63
L80:
	;
	v230 = v228
	goto L82
L81:
	;
	v230 = int32(1)
	goto L82
L82:
	;
	v237 = v230
	goto L59
L83:
	;
	v240 = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	if v241 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v244 = v242
	goto L86
L85:
	;
	v244 = int32(0)
	goto L86
L86:
	;
	v245 = int32(0)
	if int32(1)|base.B2i32(v244 == v245) != 0 {
		v291 = base.B2i32(v240|v244 == v245)
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v291 != 0 {
		v308 = int32(0)
		goto L1
	} else {
		goto L98
	}
L88:
	;
	goto L87
L89:
	;
	v259 = *(*int32)(unsafe.Add(mBase, 4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v259 != v260 {
		v291 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v262 = int32(1)
	if v259 <= v262 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v265 = v262
	goto L93
L92:
	;
	v265 = v259
	goto L93
L93:
	;
	v266 = int32(8)
	v271 = int32(0)
	goto L94
L94:
	;
	v279 = v271 << (uint(int32(2)) % 32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v266+v279)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v244+v266+v279)))
	v284 = base.B2i32(v281 == v283)
	if v281 != v283 {
		v291 = v284
		goto L88
	} else {
		goto L96
	}
L95:
	;
	v291 = v284
	goto L88
L96:
	;
	v287 = v271 + int32(1)
	if v287 != v265 {
		v271 = v287
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	goto L51
L99:
	;
	goto L44
}
func F_findPartialMatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v17 + int32(4)
	v27 = int32(-1)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v28 == int64(0) {
		v50 = v27
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	v64 = v61
	goto L12
L2:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v50
	goto L1
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v33 = int32(0)
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33*int32(12))+4))
	if v41 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v50 = v27
	goto L2
L6:
	;
	v50 = v33
	goto L2
L7:
	;
	goto L8
L8:
	;
	v45 = v33 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v45)) < base.Ui64(v28) {
		v33 = v45
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	m.G0 = v17 + int32(16)
	return v263
L11:
	;
	if v96 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v64&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v96 = v79
	goto L11
L14:
	;
	v96 = int32(0)
	goto L11
L15:
	;
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v75 = v71 & (v72 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v75
	v79 = v70 + v72*int32(12)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v83 = v80 & (v81 ^ v75)
	if v83 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v86)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v90 != int32(1) {
		v64 = base.B2i32(v83 == int32(0))
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v263 = int32(0)
	goto L10
L22:
	;
	goto L23
L23:
	;
	v101 = v20 - int32(1)
	v105 = v96
	goto L24
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[0]))
	if v117 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v263 = int32(0)
	goto L10
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v125 = F_ExecStoreMinimalTuple(m, v122, v123, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L29
	} else {
		goto L31
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L28
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_MemoryContextReset(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(_a_F_findPartialMatch_0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v129
	if int32(0) <= v101 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v133
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = v17 + int32(4)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)))
	v226 = v223
	goto L54
L34:
	;
	v141 = v101
	goto L37
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v133
	v263 = int32(1)
	goto L10
L37:
	;
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v141<<(uint(int32(1))%32)))))
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v156 < v155 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	F_slot_getsomeattrs_int(m, l1, v155)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L29
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v161 = v155 - int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v162))))
	if v164 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	if int32(0) < v141 {
		v141 = v141 - int32(1)
		goto L37
	} else {
		goto L52
	}
L44:
	;
	v166 = v161 << (uint(int32(2)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166+v167)))
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128)+6)))
	if v170 < v155 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_slot_getsomeattrs_int(m, v128, v155)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v161))))
	if v176 != 0 {
		goto L43
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v127+v141<<(uint(int32(2))%32))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184+v166)))
	v187 = F_FunctionCall2Coll(m, l2+v141*int32(28), v183, v169, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	if v187 == int32(0) {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	goto L38
L53:
	;
	if v258 != 0 {
		v105 = v258
		goto L24
	} else {
		goto L63
	}
L54:
	;
	if v226&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v258 = v241
	goto L53
L56:
	;
	v258 = int32(0)
	goto L53
L57:
	;
	goto L58
L58:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v237 = v233 & (v234 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v237
	v241 = v232 + v234*int32(12)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v245 = v242 & (v243 ^ v237)
	if v245 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)) = uint8(v248)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if v252 != int32(1) {
		v226 = base.B2i32(v245 == int32(0))
		goto L54
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	goto L25
}
