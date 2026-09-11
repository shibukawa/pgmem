package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_out_grouping_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v13
	goto L2
L1:
	;
	return v112
L2:
	;
	if v14 <= v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v112 = int32(0)
	goto L1
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	if base.Ui32(v33) < base.Ui32(int32(192)) {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l3 < v90 {
		goto L20
	} else {
		goto L21
	}
L8:
	;
	v37 = v23 + int32(1)
	if v37 == v14 {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v15))))
	v42 = v40 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v33) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v15))))
	v58 = v56 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v33) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v46 = v23 + int32(2)
	if v46 != v14 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v90 = v33<<(uint(int32(6))%32)&int32(1984) | v42
	v91 = int32(2)
	goto L7
L14:
	;
	goto L13
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v62))))
	v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(1835008) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
	v91 = int32(4)
	goto L7
L16:
	;
	v62 = v23 + int32(3)
	if v62 != v14 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v90 = v33<<(uint(int32(12))%32)&int32(61440) | v42<<(uint(int32(6))%32) | v58
	v91 = int32(3)
	goto L7
L19:
	;
	goto L18
L20:
	;
	v108 = v91 + v23
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	if l4 != 0 {
		v23 = v108
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v95 = v90 - l2
	if v95 < int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v95)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) != 0 {
		v112 = v91
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	goto L3
}
func F_record_out(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v28 = F_lookup_rowtype_tupdesc(m, v26, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v20
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v33
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+36)) = uint16(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(base.Ui32(v31) >> (uint(int32(2)) % 32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v43 == v33 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v63 == v26 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v54 = F_MemoryContextAlloc(m, v49, v30*int32(44)+int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v46 != v30 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v62 = v43
	v63 = v48
	goto L5
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = int64(0)
	v62 = v59
	v63 = int32(0)
	goto L5
L10:
	;
	v108 = F_palloc(m, v30<<(uint(int32(2))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v65 == v27 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v68 = v30 * int32(44)
	v70 = v68 + int32(12)
	if v62&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v26
	goto L10
L16:
	;
	v96 = F__emscripten_memset_bulkmem(m, v62, base.I32_extend8_s(int32(0)), v70)
	mBase = m.M
	goto L24
L17:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v70) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v62+v70) <= base.Ui32(v62) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v82 = v62 + v68 + int32(12)
	v84 = v62 + int32(4)
	if base.Ui32(v84) < base.Ui32(v82) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = v82
	goto L22
L21:
	;
	v86 = v84
	goto L22
L22:
	;
	v93 = F__emscripten_memset_bulkmem(m, v62, base.I32_extend8_s(int32(0)), (v62^int32(-1)+v86)&int32(-4)+int32(4))
	mBase = m.M
	goto L23
L23:
	;
	goto L15
L24:
	;
	goto L15
L25:
	;
	v110 = F_palloc(m, v30)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v17+int32(28), v28, v108, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_initStringInfo(m, v17+int32(12))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(40))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(0) < v30 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v129 = int32(0)
	v136 = v129
	v137 = v129
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(41))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L79
	}
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v151 = v28 + int32(20) + v145<<(uint(int32(4))%32) + v137*int32(100)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+91)))
	if v152 != 0 {
		v332 = v136
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v342 = v137 + int32(1)
	if v342 != v30 {
		v136 = v332
		v137 = v342
		goto L33
	} else {
		goto L78
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)+68))
	if v136&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(44))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v161 = int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v110))))
	if v163 != 0 {
		v332 = v161
		goto L35
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v166 = v62 + int32(12) + v137*int32(44)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v153 != v167 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_getTypeOutputInfo(m, v153, v166+int32(4), v166+int32(12))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v108+v137<<(uint(int32(2))%32))))
	v189 = F_OutputFunctionCall(m, v166+int32(16), v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+20))
	F_fmgr_info_cxt(m, v175, v166+int32(16), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v153
	goto L44
L47:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v193 = v191
	v195 = v189
	goto L50
L48:
	;
	v242 = v189
	goto L59
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v212 <= v213+int32(1) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	switch v193 & int32(255) {
	case 0:
		goto L52
	default:
		goto L53
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L49
	}
L51:
	;
	if v191 != 0 {
		v237 = int32(0)
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v209 = v195 + int32(1)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v193 = v210
	v195 = v209
	goto L50
L54:
	;
	goto L49
L55:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(34))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v225 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v223+v213))) = uint8(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v228 = int32(1)
	v229 = v227 + v228
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231+v229))) = uint8(v233)
	v237 = v228
	goto L48
L58:
	;
	v237 = int32(1)
	goto L48
L59:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v253 = base.I32_extend8_s(v252)
	if v252 == int32(34) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v317 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v315+v261))) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v321 = v319 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+v321))) = uint8(v325)
	v332 = v161
	goto L35
L61:
	;
	goto L60
L62:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v291 <= v292+int32(1) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v270 <= v271+int32(1) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	if v252 == int32(92) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	if v252 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if v237 == int32(0) {
		v332 = v161
		goto L35
	} else {
		goto L67
	}
L67:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v261+int32(1) < v260 {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(34))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v332 = v161
	goto L35
L70:
	;
	F_appendStringInfoChar(m, v17+int32(12), v253)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v279+v271))) = uint8(v253)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v284 = v282 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286+v284))) = uint8(v288)
	goto L62
L73:
	;
	goto L62
L74:
	;
	F_appendStringInfoChar(m, v17+int32(12), v253)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v302+v292))) = uint8(v253)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v306 = int32(1)
	v307 = v305 + v306
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309+v307))) = uint8(v311)
	v242 = v242 + v306
	goto L59
L77:
	;
	v242 = v242 + int32(1)
	goto L59
L78:
	;
	goto L34
L79:
	;
	F_pfree(m, v108)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v110)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if int32(0) <= v367 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_DecrTupleDescRefCount(m, v28)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	m.G0 = v17 + int32(48)
	return v372
L85:
	;
	goto L84
}
