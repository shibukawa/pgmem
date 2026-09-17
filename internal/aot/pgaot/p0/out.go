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
	v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(_a_F_out_grouping_U_0) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
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
	v90 = v33<<(uint(int32(12))%32)&int32(_a_F_out_grouping_U_1) | v42<<(uint(int32(6))%32) | v58
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
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
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
	F_check_stack_depth(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v27 = F_lookup_rowtype_tupdesc(m, v25, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v19
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v32
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+36)) = uint16(v32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(base.Ui32(v30) >> (uint(int32(2)) % 32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v42 == v32 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v62 == v25 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v53 = F_MemoryContextAlloc(m, v48, v29*int32(44)+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v45 != v29 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v61 = v42
	v62 = v47
	goto L5
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v53
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(0)
	v61 = v58
	v62 = v2
	goto L5
L10:
	;
	v112 = F_palloc(m, v29<<(uint(int32(2))%32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 == v26 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v69 = v29 * int32(44)
	v71 = v69 + int32(12)
	if v61&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v71)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v25
	goto L10
L16:
	;
	if v71 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v71 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v83 = v61 + v69 + int32(12)
	v85 = v61 + int32(4)
	if base.Ui32(v85) < base.Ui32(v83) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = v83
	goto L22
L21:
	;
	v87 = v85
	goto L22
L22:
	;
	v92 = (v61^int32(-1)+v87)&int32(-4) + int32(4)
	if v92 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.MemoryFill(m, v61, int32(0), v92)
	goto L15
L24:
	;
	base.MemoryFill(m, v61, int32(0), v71)
	goto L15
L25:
	;
	v114 = F_palloc(m, v29)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v16+int32(28), v27, v112, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v119 = v16 + int32(12)
	F_initStringInfo(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_appendStringInfoChar(m, v119, int32(40))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(0) < v29 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v135 = int32(0)
	v141 = v2
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_appendStringInfoChar(m, v16+int32(12), int32(41))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L79
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v149 = v27 + v143<<(uint(int32(4))%32) + v135*int32(100)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+111)))
	if v150 != 0 {
		v334 = v141
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v337 = v135 + int32(1)
	if v337 != v29 {
		v135 = v337
		v141 = v334
		goto L33
	} else {
		goto L78
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+88))
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_appendStringInfoChar(m, v16+int32(12), int32(44))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v157 = int32(1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v114))))
	if v159 != 0 {
		v334 = v157
		goto L35
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v162 = v61 + int32(12) + v135*int32(44)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v151 != v163 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_getTypeOutputInfo(m, v151, v162+int32(4), v162+int32(12))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v112+v135<<(uint(int32(2))%32))))
	v185 = F_OutputFunctionCall(m, v162+int32(16), v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	F_fmgr_info_cxt(m, v171, v162+int32(16), v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v151
	goto L44
L47:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v189 = v185
	v192 = v187
	goto L50
L48:
	;
	v236 = v185
	goto L59
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v207 <= v208+int32(1) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	switch v192 & int32(255) {
	case 0:
		goto L52
	default:
		goto L53
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L49
	}
L51:
	;
	if v187 != 0 {
		v232 = int32(0)
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v189 = v189 + int32(1)
	v192 = v203
	goto L50
L54:
	;
	goto L49
L55:
	;
	F_appendStringInfoChar(m, v16+int32(12), int32(34))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v220 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v218+v208))) = uint8(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v223 = int32(1)
	v224 = v222 + v223
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v226+v224))) = uint8(v228)
	v232 = v223
	goto L48
L58:
	;
	v232 = int32(1)
	goto L48
L59:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	v247 = base.I32_extend8_s(v246)
	if base.B2i32(v246 == int32(34))|base.B2i32(v246 == int32(92)) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v313 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v311+v258))) = uint8(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v317 = v315 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v319+v317))) = uint8(v321)
	v334 = v157
	goto L35
L61:
	;
	goto L60
L62:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v288 <= v289+int32(1) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	if v246 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v267 <= v268+int32(1) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if v232 == int32(0) {
		v334 = v157
		goto L35
	} else {
		goto L67
	}
L67:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v258+int32(1) < v257 {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	F_appendStringInfoChar(m, v16+int32(12), int32(34))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v334 = v157
	goto L35
L70:
	;
	F_appendStringInfoChar(m, v16+int32(12), v247)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v276+v268))) = uint8(v247)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v281 = v279 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283+v281))) = uint8(v285)
	goto L62
L73:
	;
	goto L62
L74:
	;
	F_appendStringInfoChar(m, v16+int32(12), v247)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v297+v289))) = uint8(v247)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v302 = v300 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v304+v302))) = uint8(v306)
	goto L76
L76:
	;
	v236 = v236 + int32(1)
	goto L59
L77:
	;
	goto L76
L78:
	;
	goto L34
L79:
	;
	F_pfree(m, v112)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v114)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if int32(0) <= v361 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_DecrTupleDescRefCount(m, v27)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	m.G0 = v16 + int32(48)
	return v366
L85:
	;
	goto L84
}
