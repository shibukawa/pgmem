package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_check_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_check_upd_0), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_RI_FKey_check(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_ri_KeysEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v23 <= v6 {
		v265 = int32(1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
	} else {
		goto L69
	}
L2:
	;
	m.G0 = v20 + int32(80)
	return v265
L3:
	;
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = int32(172)
	goto L6
L5:
	;
	v28 = int32(236)
	goto L6
L6:
	;
	v44 = v6
	goto L7
L7:
	;
	v53 = l3 + v28 + v44<<(uint(int32(1))%32)
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v55 < v54 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v265 = v248
	goto L2
L9:
	;
	F_slot_getsomeattrs_int(m, l1, v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v61 = int32(0)
	v63 = v54 - int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v64))))
	if v66 != 0 {
		v265 = v61
		goto L2
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v63<<(uint(int32(2))%32))))
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v73 < v72 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_slot_getsomeattrs_int(m, l2, v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v78 = v72 - int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79))))
	if v81 != 0 {
		v265 = v61
		goto L2
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v78<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v248 = int32(1)
	v250 = v44 + v248
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v250 < v251 {
		v44 = v250
		goto L7
	} else {
		goto L68
	}
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v91 = v87 + v88<<(uint(int32(4))%32)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+10)))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+8)))
	v94 = F_datum_image_eq(m, v71, v86, v92, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+165)))
	if v96 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v94 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v265 = v61
	goto L2
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v109 = F_attnumTypeId(m, l0, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L31
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v44 == v99-int32(1) {
		v106 = l3 + int32(684)
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v106 = l3 + int32(556) + v44<<(uint(int32(2))%32)
	goto L26
L30:
	;
	goto L29
L31:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v112 = F_attnumCollationId(m, l0, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[0]))
	if v115 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(3023656976388)
	v126 = F_hash_create(m, int32(_a_F_ri_KeysEqual_0), int32(64), v20+int32(32), int32(40))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	v156 = v115
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v107
	v164 = F_hash_search(m, v156, v20+int32(32), int32(1), v20+int32(31))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L12
	} else {
		goto L40
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[1])) = v126
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1501), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(51539607560)
	v142 = F_hash_create(m, int32(_a_F_ri_KeysEqual_1), int32(256), v20+int32(32), int32(40))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[2])) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(292057776136)
	v153 = F_hash_create(m, int32(_a_F_ri_KeysEqual_2), int32(256), v20+int32(32), int32(40))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[0])) = v153
	v156 = v153
	goto L35
L40:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+31)))
	if v166 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v164)+44))
	if v220 != 0 {
		goto L61
	} else {
		goto L62
	}
L42:
	;
	v174 = F_get_opcode(m, v107)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L47
	}
L43:
	;
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)) = uint8(v169)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
	if v171&int32(1) != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[3]))
	F_fmgr_info_cxt(m, v174, v164+int32(12), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	F_op_input_types(m, v107, v20+int32(24), v20+int32(20))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v109 == v188 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)) = uint8(v217)
	goto L41
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+44)) = int32(0)
	goto L50
L52:
	;
	v193 = F_find_coercion_pathway(m, v188, v109, int32(0), v20+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v193-int32(3)) <= base.Ui32(int32(-3)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v200 = F_IsBinaryCoercible(m, v109, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v204 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L57:
	;
	if v200 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_ri_KeysEqual[3]))
	F_fmgr_info_cxt(m, v204, v164+int32(40), v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	goto L50
L61:
	;
	v222 = v164 + int32(40)
	v223 = int32(0)
	v226 = F_FunctionCall3Coll(m, v222, v223, v86, int32(-1), v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L64
	}
L62:
	;
	v233 = v71
	v234 = v86
	goto L63
L63:
	;
	v238 = F_FunctionCall2Coll(m, v164+int32(12), v112, v234, v233)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L12
	} else {
		goto L66
	}
L64:
	;
	v228 = int32(0)
	v231 = F_FunctionCall3Coll(m, v222, v228, v71, int32(-1), v228)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	v233 = v231
	v234 = v226
	goto L63
L66:
	;
	if v238 == int32(0) {
		v265 = v61
		goto L2
	} else {
		goto L67
	}
L67:
	;
	goto L20
L68:
	;
	goto L8
L69:
	;
	v278 = F_format_type_be(m, v109)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v281 = F_format_type_be(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v278
	F_errmsg_internal(m, int32(_a_F_ri_KeysEqual_3), v20)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ri_KeysEqual_4), int32(3190), int32(_a_F_ri_KeysEqual_5))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ri_set(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v341 int32
	_ = v341
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	v18 = m.G0
	v20 = v18 - int32(640)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = F_ri_FetchConstraintInfo(m, v22, v23, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	v29 = F_table_open(m, v27, int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v39 = int32(8)
	goto L7
L6:
	;
	v39 = int32(10)
	goto L7
L7:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(7)
	goto L10
L9:
	;
	v42 = int32(9)
	goto L10
L10:
	;
	v44 = base.B2i32(l2 == int32(2))
	if l2 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v39
	goto L13
L12:
	;
	v45 = v42
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+636)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v20)+632)) = v36
	v49 = v20 + int32(632)
	v52 = F_ri_FetchPreparedPlan(m, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v498 = v52
	goto L17
L16:
	;
	if v44 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v499 = int32(0)
	v503 = F_ri_PerformCheck(m, v25, v49, v498, v29, v32, v31, v499, v499, int32(1), int32(9))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L102
	}
L18:
	;
	F_initStringInfo(m, v20+int32(616))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)+96))
	if v57 != 0 {
		v61 = v57
		v62 = int32(100)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	v61 = v59
	v62 = int32(236)
	goto L18
L22:
	;
	goto L21
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+119)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v70 = F_get_namespace_name(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v72 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+352)) = uint8(v72)
	v79 = v70
	v81 = v20 + int32(352)
	goto L25
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v94 != int32(34) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v111 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+1)) = uint16(v111)
	v114 = v20 + int32(352)
	if v114&int32(3) == int32(0) {
		v138 = v114
		goto L35
	} else {
		goto L36
	}
L27:
	;
	goto L26
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v106)
	v79 = v79 + int32(1)
	v81 = v107
	goto L25
L29:
	;
	if v94 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)) = uint8(v101)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v106 = v103
	v107 = v81 + int32(2)
	goto L28
L32:
	;
	v106 = v94
	v107 = v81 + int32(1)
	goto L28
L33:
	;
	v174 = v171 + (v20 + int32(352))
	v175 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v175)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v179 = v174 + int32(1)
	v180 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v180)
	v186 = v179
	v188 = v177 + int32(4)
	goto L50
L34:
	;
	v171 = v163 - v114
	goto L33
L35:
	;
	v142 = v138
	goto L44
L36:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v122 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v171 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v127 = v114
	goto L40
L40:
	;
	v131 = v127 + int32(1)
	if v131&int32(3) == int32(0) {
		v138 = v131
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v163 = v131
	goto L34
L42:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v136 != 0 {
		v127 = v131
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v151 = int32(-2139062144)
	if (int32(16843008)-v148|v148)&v151 == v151 {
		v142 = v142 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v157 = v142
	goto L47
L46:
	;
	goto L45
L47:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v161 != 0 {
		v157 = v157 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v163 = v157
	goto L34
L49:
	;
	goto L48
L50:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v201 != int32(34) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v218 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+1)) = uint16(v218)
	v220 = int32(_a_F_ri_set_0)
	if v68&int32(255) == int32(112) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v213)
	v186 = v214
	v188 = v188 + int32(1)
	goto L50
L54:
	;
	if v201 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v208 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)) = uint8(v208)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v213 = v210
	v214 = v186 + int32(2)
	goto L53
L57:
	;
	v213 = v201
	v214 = v186 + int32(1)
	goto L53
L58:
	;
	v227 = v220
	goto L60
L59:
	;
	v227 = int32(_a_F_ri_set_1)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v20 + int32(352)
	F_appendStringInfo(m, v20+int32(616), int32(_a_F_ri_set_2), v20+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v61 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	if int32(0) < v341 {
		goto L81
	} else {
		goto L82
	}
L63:
	;
	if l1 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v243 = int32(_a_F_ri_set_3)
	goto L66
L65:
	;
	v243 = int32(_a_F_ri_set_4)
	goto L66
L66:
	;
	v252 = int32(0)
	v255 = v220
	goto L67
L67:
	;
	v265 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62+v25+v252<<(uint(int32(1))%32)))))
	v266 = F_attnumAttName(m, v29, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v268 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+208)) = uint8(v268)
	v274 = v266
	v276 = v20 + int32(208)
	goto L70
L70:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v289 != int32(34) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v319)
	v274 = v274 + int32(1)
	v276 = v320
	goto L70
L73:
	;
	if v289 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v314 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)) = uint8(v314)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v319 = v316
	v320 = v276 + int32(2)
	goto L72
L76:
	;
	v294 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v276)+1)) = uint16(v294)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v20 + int32(208)
	F_appendStringInfo(m, v20+int32(616), int32(_a_F_ri_set_5), v20+int32(32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v319 = v289
	v320 = v276 + int32(1)
	goto L72
L79:
	;
	v310 = v252 + int32(1)
	if v310 != v61 {
		v252 = v310
		v255 = int32(_a_F_ri_set_6)
		goto L67
	} else {
		goto L80
	}
L80:
	;
	goto L62
L81:
	;
	v359 = int32(0)
	v361 = int32(_a_F_ri_set_7)
	goto L84
L82:
	;
	v461 = v341
	goto L83
L83:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v20)+616))
	v479 = F_ri_PlanCheck(m, v474, v461, v20-int32(-64), v20+int32(632), v29, v32)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L101
	}
L84:
	;
	v370 = v359 << (uint(int32(1)) % 32)
	v372 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25+int32(172)+v370))))
	v373 = F_attnumTypeId(m, v32, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v461 = v455
	goto L83
L86:
	;
	v375 = v370 + (v25 + int32(236))
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v375))))
	v377 = F_attnumTypeId(m, v29, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v375))))
	v380 = F_attnumAttName(m, v29, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v382 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+208)) = uint8(v382)
	v388 = v380
	v390 = v20 + int32(208)
	goto L89
L89:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v403 != int32(34) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v420 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+1)) = uint16(v420)
	v423 = v359 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v423
	v430 = F_pg_sprintf(m, v20+int32(192), int32(_a_F_ri_set_8), v20+int32(16))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L97
	}
L91:
	;
	goto L90
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v416))) = uint8(v415)
	v388 = v388 + int32(1)
	v390 = v416
	goto L89
L93:
	;
	if v403 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v410 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+1)) = uint8(v410)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v415 = v412
	v416 = v390 + int32(2)
	goto L92
L96:
	;
	v415 = v403
	v416 = v390 + int32(1)
	goto L92
L97:
	;
	v433 = v359 << (uint(int32(2)) % 32)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(300)+v433)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v361
	F_appendStringInfo(m, v20+int32(616), int32(_a_F_ri_set_9), v20)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_generate_operator_clause(m, v20+int32(616), v20+int32(192), v373, v435, v20+int32(208), v377)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20-int32(-64)+v433))) = v373
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	if v423 < v455 {
		v359 = v423
		v361 = int32(_a_F_ri_set_10)
		goto L84
	} else {
		goto L100
	}
L100:
	;
	goto L85
L101:
	;
	v498 = v479
	goto L17
L102:
	;
	v505 = F_SPI_finish(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v505 == int32(2) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_sequence_close(m, v29, int32(3))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L112
	}
L107:
	;
	if l1 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_ri_restrict(m, l0, int32(1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	m.G0 = v20 + int32(640)
	return
L111:
	;
	goto L110
L112:
	;
	F_errmsg_internal(m, int32(_a_F_ri_set_11), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_ri_set_12), int32(1348), int32(_a_F_ri_set_13))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
