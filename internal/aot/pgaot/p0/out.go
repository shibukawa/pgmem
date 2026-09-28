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
func F_record_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
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
	var v338 int32
	_ = v338
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
	var v372 int64
	_ = v372
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v63 = v2
	goto L5
L10:
	;
	v112 = F_palloc_mul(m, int32(8), v30)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v70 = v30 * int32(44)
	v72 = v70 + int32(12)
	if v62&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v72)) == int32(0) {
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
	if v72 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v72 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v84 = v62 + v70 + int32(12)
	v86 = v62 + int32(4)
	if base.Ui32(v86) < base.Ui32(v84) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v88 = v84
	goto L22
L21:
	;
	v88 = v86
	goto L22
L22:
	;
	v93 = (v62^int32(-1)+v88)&int32(-4) + int32(4)
	if v93 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.MemoryFill(m, v62, int32(0), v93)
	goto L15
L24:
	;
	base.MemoryFill(m, v62, int32(0), v72)
	goto L15
L25:
	;
	v115 = F_palloc_mul(m, int32(1), v30)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v17+int32(28), v28, v112, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v120 = v17 + int32(12)
	F_initStringInfo(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_appendStringInfoChar(m, v120, int32(40))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
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
	v136 = int32(0)
	v142 = v2
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
	v151 = v28 + v145<<(uint(int32(3))%32) + v136*int32(100)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+119)))
	if v152 != 0 {
		v338 = v142
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v342 = v136 + int32(1)
	if v342 != v30 {
		v136 = v342
		v142 = v338
		goto L33
	} else {
		goto L78
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)+96))
	if v142 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(44))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v159 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v115))))
	if v161 != 0 {
		v338 = v159
		goto L35
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v164 = v62 + int32(12) + v136*int32(44)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v153 != v165 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_getTypeOutputInfo(m, v153, v164+int32(4), v164+int32(12))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v112+v136<<(uint(int32(3))%32))))
	v187 = F_OutputFunctionCall(m, v164+int32(16), v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	F_fmgr_info_cxt(m, v173, v164+int32(16), v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v153
	goto L44
L47:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v191 = v187
	v194 = v189
	goto L50
L48:
	;
	v239 = v187
	goto L59
L49:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v210 <= v211+int32(1) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	switch v194 & int32(255) {
	case 0:
		goto L52
	default:
		goto L53
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L49
	}
L51:
	;
	if v189 != 0 {
		v235 = int32(0)
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	v191 = v191 + int32(1)
	v194 = v206
	goto L50
L54:
	;
	goto L49
L55:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(34))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v223 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v221+v211))) = uint8(v223)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v226 = int32(1)
	v227 = v225 + v226
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v229+v227))) = uint8(v231)
	v235 = v226
	goto L48
L58:
	;
	v235 = int32(1)
	goto L48
L59:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v251 = base.I32_extend8_s(v250)
	if base.B2i32(v250 == int32(34))|base.B2i32(v250 == int32(92)) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v317 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v315+v262))) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v321 = v319 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+v321))) = uint8(v325)
	v338 = v159
	goto L35
L61:
	;
	goto L60
L62:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v292 <= v293+int32(1) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	if v250 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v271 <= v272+int32(1) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if v235 == int32(0) {
		v338 = v159
		goto L35
	} else {
		goto L67
	}
L67:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v262+int32(1) < v261 {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	F_appendStringInfoChar(m, v17+int32(12), int32(34))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v338 = v159
	goto L35
L70:
	;
	F_appendStringInfoChar(m, v17+int32(12), v251)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v280+v272))) = uint8(v251)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v285 = v283 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v287+v285))) = uint8(v289)
	goto L62
L73:
	;
	goto L62
L74:
	;
	F_appendStringInfoChar(m, v17+int32(12), v251)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v301+v293))) = uint8(v251)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v306 = v304 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v308+v306))) = uint8(v310)
	goto L76
L76:
	;
	v239 = v239 + int32(1)
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
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v115)
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
	v372 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+12)))
	m.G0 = v17 + int32(48)
	return v372
L85:
	;
	goto L84
}
