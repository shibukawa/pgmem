package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_FindDeletedTupleInLocalRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[0]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+41)))
	if v21 != int32(1) {
		v420 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v420
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[1])))
	if v25&int32(1) == int32(0) {
		v420 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[2]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(3) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v290 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v290
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v290)
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
	v297 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L84
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L80
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L11
	} else {
		goto L76
	}
L7:
	;
	if v126 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v126 = v35
	goto L7
L9:
	;
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[3]))
	v41 = F_LWLockAcquire(m, v37+int32(_a_F_FindDeletedTupleInLocalRel_0), int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[2]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v49 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[4]))
	if v57 <= v49 {
		v100 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v100 == int32(0) {
		goto L6
	} else {
		goto L26
	}
L14:
	;
	goto L13
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[5]))
	v69 = v49
	goto L16
L16:
	;
	v75 = v61 + int32(16) + v69<<(uint(int32(7))%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)))
	if v76 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v100 = int32(0)
	goto L14
L18:
	;
	v93 = v69 + int32(1)
	if v93 != v57 {
		v69 = v93
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v79 == int32(4) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	if v82 != v48 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
	if base.B2i32(v84 != v49)|base.B2i32(int32(3) != v79) != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v100 = v75
	goto L14
L25:
	;
	goto L17
L26:
	;
	v109 = base.AtomicRmwXchg32(m, v100, int32(56), int32(1))
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_s_lock(m, v100+int32(56), int32(_a_F_FindDeletedTupleInLocalRel_1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)+72))
	v116 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+56)), uint32(v116))
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[3]))
	F_LWLockRelease(m, v120+int32(_a_F_FindDeletedTupleInLocalRel_0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v126 = v115
	goto L7
L32:
	;
	v420 = int32(0)
	goto L1
L33:
	;
	goto L34
L34:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v134 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	if v134 == int32(0) {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+20)))
	v141 = int32(768)
	if v140&v141 != v141 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v146 = v145
	goto L40
L39:
	;
	v146 = int32(2)
	goto L40
L40:
	;
	F_ReleaseCatCache(m, v134)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v149 = int32(3)
	if base.B2i32(base.Ui32(v126) < base.Ui32(v149))|base.B2i32(base.Ui32(v146) < base.Ui32(v149)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v160 = m.G0
	v162 = v160 - int32(1792)
	m.G0 = v162
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v164
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v164)
	v170 = F_GetRelationIdentityOrPK(m, l0)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L48
	}
L43:
	;
	if v146-v126 < int32(0) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(v126) <= base.Ui32(v146) {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L4
L47:
	;
	goto L42
L48:
	;
	v173 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v176 = F_index_open(m, l1, int32(3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	v180 = F_build_replindex_scan_key(m, v162, v176, l2)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v182 = int32(0)
	v184 = F_index_beginscan(m, l0, v176, int32(_a_F_FindDeletedTupleInLocalRel_2), int32(0), v180, v182, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v186 = int32(0)
	F_index_rescan(m, v184, v162, v180, v186, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v191 = F_index_getnext_slot(m, v184, int32(1), v173)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	if v191 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v195 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_index_endscan(m, v184)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L11
	} else {
		goto L73
	}
L58:
	;
	if base.B2i32(l1 == v170) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L57
L60:
	;
	v229 = F_index_getnext_slot(m, v184, int32(1), v173)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L71
	}
L61:
	;
	if v195 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v224 = v195
	goto L63
L63:
	;
	F_update_most_recent_deletion_info(m, v173, v126, l3, l5, l4)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L70
	}
L64:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v216 = F_palloc0_mul(m, int32(4), v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L67
	}
L65:
	;
	v218 = v195
	goto L66
L66:
	;
	v220 = F_tuples_equal(m, v173, l2, v218, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L68
	}
L67:
	;
	v218 = v216
	goto L66
L68:
	;
	if v220 == int32(0) {
		v227 = v218
		goto L60
	} else {
		goto L69
	}
L69:
	;
	v224 = v218
	goto L63
L70:
	;
	v227 = v224
	goto L60
L71:
	;
	if v229 != 0 {
		v195 = v227
		goto L58
	} else {
		goto L72
	}
L72:
	;
	goto L59
L73:
	;
	F_relation_close(m, v176, int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_ExecDropSingleTupleTableSlot(m, v173)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	m.G0 = v162 + int32(1792)
	v420 = base.B2i32(v252 != int64(0))
	goto L1
L76:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_FindDeletedTupleInLocalRel_3), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_FindDeletedTupleInLocalRel_4), int32(3351), int32(_a_F_FindDeletedTupleInLocalRel_5))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
	F_errmsg_internal(m, int32(_a_F_FindDeletedTupleInLocalRel_6), v17)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_FindDeletedTupleInLocalRel_4), int32(3277), int32(_a_F_FindDeletedTupleInLocalRel_7))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v420 = base.B2i32(v397 != int64(0))
	goto L1
L84:
	;
	if v297 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v302 = F_RelationGetIndexAttrBitmap(m, l0, int32(1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L88
	}
L86:
	;
	v304 = v297
	goto L87
L87:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v308 = F_palloc0_mul(m, int32(4), v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L11
	} else {
		goto L89
	}
L88:
	;
	v304 = v302
	goto L87
L89:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[6]))
	if v311 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L11
	} else {
		goto L113
	}
L91:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FindDeletedTupleInLocalRel[7])))
	if v313&int32(1) == int32(0) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v319 = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v325 = m.T0[v324].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_FindDeletedTupleInLocalRel_2), v319, v319, v319, int32(449))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L11
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v328 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	v330 = int32(0)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+188))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+16))
	m.T0[v337].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v325, v330, v330, v330, v330, v330)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+40)) = v341
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+188))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+20))
	v347 = m.T0[v346].(func(*base.Module, int32, int32, int32) int32)(m, v325, int32(1), v328)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	if v347 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	goto L102
L100:
	;
	goto L101
L101:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+188))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	m.T0[v392].(func(*base.Module, int32))(m, v325)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L11
	} else {
		goto L111
	}
L102:
	;
	v363 = F_tuples_equal(m, v328, l2, v308, v304)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	if v363 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_update_most_recent_deletion_info(m, v328, v126, l3, l5, l4)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L11
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+40)) = v368
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+188))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+20))
	v374 = m.T0[v373].(func(*base.Module, int32, int32, int32) int32)(m, v325, int32(1), v328)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L11
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	if v374 != 0 {
		goto L102
	} else {
		goto L110
	}
L110:
	;
	goto L103
L111:
	;
	F_ExecDropSingleTupleTableSlot(m, v328)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	v397 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	goto L83
L113:
	;
	F_errmsg_internal(m, int32(_a_F_FindDeletedTupleInLocalRel_8), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_FindDeletedTupleInLocalRel_9), int32(931), int32(_a_F_FindDeletedTupleInLocalRel_10))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FindLockCycleRecurseMember(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v233 int32
	_ = v233
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int64
	_ = v624
	var v626 int64
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v660 int64
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v697 int32
	_ = v697
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)))
	if v17 == int32(1) {
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
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_FindLockCycleRecurseMember[0])))
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v250 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[1]))
	if v250 < v252 {
		goto L63
	} else {
		goto L64
	}
L6:
	;
	v36 = v16 + int32(24)
	if v32 == v36 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v40 = l2 + int32(1)
	v54 = v32
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(16))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+364))
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v233 != v36 {
		v54 = v233
		goto L8
	} else {
		goto L59
	}
L11:
	;
	v60 = v59
	goto L13
L12:
	;
	v60 = v58
	goto L13
L13:
	;
	if base.B2i32(v60 == l1)|base.B2i32(v38 <= int32(0)) != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(8))))
	v75 = int32(1)
	goto L15
L15:
	;
	v85 = int32(1) << (uint(v75) % 32)
	if v85&v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v58)+364))
	if v98 != 0 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v89 = v67 & v85
	goto L19
L18:
	;
	v89 = int32(0)
	goto L19
L19:
	;
	if v89 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v93 = v75 + int32(1)
	if v93 <= v38 {
		v75 = v93
		goto L15
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L16
L23:
	;
	goto L10
L24:
	;
	if v192 != 0 {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v99 = v98
	goto L27
L26:
	;
	v99 = v58
	goto L27
L27:
	;
	v100 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v100 < v104 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v107 = v100
	goto L31
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v104 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v102+v104<<(uint(int32(2))%32)))) = v99
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v99)+392))
	if v141 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v102+v107<<(uint(int32(2))%32))))
	if v99 == v117 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v124 = v107 + int32(1)
	if v124 != v104 {
		v107 = v124
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v192 = int32(0)
	goto L24
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v40
	v192 = int32(1)
	goto L24
L39:
	;
	goto L32
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v99)+372))
	if v148 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v144 = F_FindLockCycleRecurseMember(m, v99, v99, v40, l3, l4)
	mBase = m.M
	if v144 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v192 = int32(1)
	goto L24
L43:
	;
	v192 = int32(0)
	goto L24
L44:
	;
	v152 = v99 + int32(368)
	if v148 == v152 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v154 = v148
	goto L46
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v161 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v175 != v152 {
		v154 = v175
		goto L46
	} else {
		goto L53
	}
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if v164 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v168 = v154 - int32(376)
	if v168 == v99 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v170 = F_FindLockCycleRecurseMember(m, v168, v99, v40, l3, l4)
	mBase = m.M
	if v170 == int32(0) {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v192 = int32(1)
	goto L24
L53:
	;
	goto L47
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v197 = v194 + l2*int32(24)
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+8)) = v198
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+20)) = v204
	return int32(1)
L55:
	;
	goto L56
L56:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[6]))
	if l0 != v209 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+36)))
	if v211&int32(1) == int32(0) {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[7])) = v58
	goto L10
L59:
	;
	goto L9
L60:
	;
	return v697
L61:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v657 = v654 + l2*int32(24)
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v657)+8)) = v658
	v660 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v657))) = v660
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(v657)+16)) = v662
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v657)+20)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v667 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v666*v667))) = l1
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v671*v667)+4)) = v379
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v676*v667)+8)) = v16
	v681 = int32(1)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v682 + v681
	v697 = v681
	goto L60
L62:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	if v482 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L63:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[8]))
	v263 = v250
	goto L66
L64:
	;
	goto L65
L65:
	;
	v296 = v16 + int32(32)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v298 == int32(0) {
		v337 = l0
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v274 = v256 + v263*int32(12)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v275 == v16 {
		goto L62
	} else {
		goto L68
	}
L67:
	;
	goto L65
L68:
	;
	v278 = v263 + int32(1)
	if v278 != v252 {
		v263 = v278
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v345 = int32(0)
	if base.B2i32(v297 == v345)|base.B2i32(v297 == v296) != 0 {
		v697 = v345
		goto L60
	} else {
		goto L79
	}
L71:
	;
	v301 = int32(0)
	if base.B2i32(v297 == v301)|base.B2i32(v297 == v296) != 0 {
		v337 = v301
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v312 = v297
	v313 = v301
	goto L73
L73:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v312-int32(24))))
	if v325 == l1 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v337 = v327
	goto L70
L75:
	;
	v327 = v312 - int32(388)
	goto L77
L76:
	;
	v327 = v313
	goto L77
L77:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v328 != v296 {
		v312 = v328
		v313 = v327
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v351 = l2 + int32(1)
	v357 = v297
	goto L80
L80:
	;
	v368 = v357 - int32(388)
	if v368 == v337 {
		v697 = v345
		goto L60
	} else {
		goto L82
	}
L81:
	;
	v697 = v345
	goto L60
L82:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	if int32(base.Ui32(v31)>>(uint(v370)%32))&int32(1) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	if v480 != v296 {
		v357 = v480
		goto L80
	} else {
		goto L120
	}
L84:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v357-int32(24))))
	if v378 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v379 = v378
	goto L87
L86:
	;
	v379 = v368
	goto L87
L87:
	;
	if v379 == l1 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v368)+364))
	if v384 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v478 != 0 {
		goto L61
	} else {
		goto L119
	}
L90:
	;
	v385 = v384
	goto L92
L91:
	;
	v385 = v368
	goto L92
L92:
	;
	v386 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v386 < v390 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v393 = v386
	goto L96
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v390 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v388+v390<<(uint(int32(2))%32)))) = v385
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v385)+392))
	if v427 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L96:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v388+v393<<(uint(int32(2))%32))))
	if v385 == v403 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	if v393 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v410 = v393 + int32(1)
	if v410 != v390 {
		v393 = v410
		goto L96
	} else {
		goto L104
	}
L101:
	;
	v478 = int32(0)
	goto L89
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v351
	v478 = int32(1)
	goto L89
L104:
	;
	goto L97
L105:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v385)+372))
	if v434 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v430 = F_FindLockCycleRecurseMember(m, v385, v385, v351, l3, l4)
	mBase = m.M
	if v430 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v478 = int32(1)
	goto L89
L108:
	;
	v478 = int32(0)
	goto L89
L109:
	;
	v438 = v385 + int32(368)
	if v434 == v438 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v440 = v434
	goto L111
L111:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	if v447 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L108
L113:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v461 != v438 {
		v440 = v461
		goto L111
	} else {
		goto L118
	}
L114:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	if v450 == int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v454 = v440 - int32(376)
	if v454 == v385 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v456 = F_FindLockCycleRecurseMember(m, v454, v385, v351, l3, l4)
	mBase = m.M
	if v456 == int32(0) {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v478 = int32(1)
	goto L89
L118:
	;
	goto L112
L119:
	;
	goto L83
L120:
	;
	goto L81
L121:
	;
	return int32(0)
L122:
	;
	goto L123
L123:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v489 = l2 + int32(1)
	v490 = int32(0)
	v497 = v490
	goto L124
L124:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v487+v497<<(uint(int32(2))%32))))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+364))
	if v511 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v623 = v620 + l2*int32(24)
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v623)+8)) = v624
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v623))) = v626
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+16)) = v628
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+20)) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v633 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v632*v633))) = l1
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v637*v633)+4)) = v512
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v642*v633)+8)) = v16
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v648 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v647 + v648
	return v648
L126:
	;
	v512 = v511
	goto L128
L127:
	;
	v512 = v510
	goto L128
L128:
	;
	if v512 == l1 {
		v697 = v490
		goto L60
	} else {
		goto L129
	}
L129:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v510)+400))
	if int32(base.Ui32(v31)>>(uint(v514)%32))&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L125
L131:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v510)+364))
	if v521 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	goto L133
L133:
	;
	v617 = v497 + int32(1)
	if v617 != v482 {
		v497 = v617
		goto L124
	} else {
		goto L165
	}
L134:
	;
	if v615 != 0 {
		goto L130
	} else {
		goto L164
	}
L135:
	;
	v522 = v521
	goto L137
L136:
	;
	v522 = v510
	goto L137
L137:
	;
	v523 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v523 < v527 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v530 = v523
	goto L141
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v527 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v525+v527<<(uint(int32(2))%32)))) = v522
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v522)+392))
	if v564 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L141:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v525+v530<<(uint(int32(2))%32))))
	if v522 == v540 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	goto L140
L143:
	;
	if v530 != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	v547 = v530 + int32(1)
	if v547 != v527 {
		v530 = v547
		goto L141
	} else {
		goto L149
	}
L146:
	;
	v615 = int32(0)
	goto L134
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v489
	v615 = int32(1)
	goto L134
L149:
	;
	goto L142
L150:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v522)+372))
	if v571 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v567 = F_FindLockCycleRecurseMember(m, v522, v522, v489, l3, l4)
	mBase = m.M
	if v567 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v615 = int32(1)
	goto L134
L153:
	;
	v615 = int32(0)
	goto L134
L154:
	;
	v575 = v522 + int32(368)
	if v571 == v575 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v577 = v571
	goto L156
L156:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v577)+16))
	if v584 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L153
L158:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v598 != v575 {
		v577 = v598
		goto L156
	} else {
		goto L163
	}
L159:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v577)+8))
	if v587 == int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v591 = v577 - int32(376)
	if v591 == v522 {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v593 = F_FindLockCycleRecurseMember(m, v591, v522, v489, l3, l4)
	mBase = m.M
	if v593 == int32(0) {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	v615 = int32(1)
	goto L134
L163:
	;
	goto L157
L164:
	;
	goto L133
L165:
	;
	v697 = v490
	goto L60
}
func F_ForceSyncCommit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ForceSyncCommit[0])) = uint8(v2)
	return
}
func F_ForgetBackgroundWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ForgetBackgroundWorker[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1488))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+192)))
	if v17&v6 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v20 + int32(1)
	} else {
	}
	v24 = int32(0)
	v27 = base.AtomicRmwOr32(m, v24, int32(_a_F_ForgetBackgroundWorker_0), v24)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v11*int32(1488)+v6))) = uint8(v24)
	v32 = F_errstart(m, int32(14), v24)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return
	} else {
		if v32 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_ForgetBackgroundWorker_1), v7)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ForgetBackgroundWorker_2), int32(466), int32(_a_F_ForgetBackgroundWorker_3))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1496))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1500))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1496))
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
					F_pfree(m, l0)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1496))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1500))
			*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v44
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1496))
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
			F_pfree(m, l0)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_ForgetDatabaseSyncRequests(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(4294967295)
	v14 = int32(_a_F_ForgetDatabaseSyncRequests_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+10)) = uint16(v14)
	v20 = F_RegisterSyncRequest(m, v5+int32(8), int32(3), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v5 + int32(32)
		return
	}
}
func F_FreeDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 {
	case 0:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v34 = F_fclose(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = v34
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	case 1:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = F_pgl_pclose(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v36 = v6
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	case 2:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v13 = F_close(m, v12)
		mBase = m.M
		F_emscripten_builtin_free(m, v10)
		mBase = m.M
		v36 = v13
		v37 = int32(_a_F_FreeDesc_0)
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
		v41 = v39 - int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
		v47 = v44 + v41*int32(12)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
		return v36
	case 3:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_pgaio_closing_fd(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = F_close(m, v18)
			mBase = m.M
			v36 = v19
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_FreeDesc_1), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_FreeDesc_2), int32(2813), int32(_a_F_FreeDesc_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_FreeSnapshotBuilder(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 != 0 {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+30)))
		if v7 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_FreeSnapshotBuilder_0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FreeSnapshotBuilder_1), int32(348), int32(_a_F_FreeSnapshotBuilder_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
			v12 = v10 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v12
			if v12 == int32(0) {
				F_pfree(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					F_MemoryContextDelete(m, v5)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				F_MemoryContextDelete(m, v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int64)(unsafe.Add(mBase, _c_F_FreeSpaceMapVacuum[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v8
	v14 = F_fsm_vacuum_page(m, l0, v5, int32(0), int32(-1), v5+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int64
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v613 int32
	_ = v613
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v728 int32
	_ = v728
	var v752 int32
	_ = v752
	var v760 int32
	_ = v760
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v798 int32
	_ = v798
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v930 int32
	_ = v930
	var v945 int32
	_ = v945
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1012 int32
	_ = v1012
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1194 int32
	_ = v1194
	var v1231 int32
	_ = v1231
	var v1258 int32
	_ = v1258
	var v1268 int32
	_ = v1268
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1318 int32
	_ = v1318
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1577 int32
	_ = v1577
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1660 int32
	_ = v1660
	var v1709 int32
	_ = v1709
	var v1730 int32
	_ = v1730
	var v1744 int32
	_ = v1744
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1769 int32
	_ = v1769
	var v1797 int32
	_ = v1797
	var v1806 int32
	_ = v1806
	v9 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(32)
	m.G0 = v38
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v9
	F_DeconstructQualifiedName(m, l0, v38+int32(12), v38+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v38 + int32(32)
	return v1806
L4:
	;
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+8)))
	v69 = int64(0)
	v71 = F_SearchSysCacheList(m, int32(46), int32(1), v68, v69, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v51 | int32(1)
	v55 = F_LookupExplicitNamespace(m, v50, l6)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	if v55 == int32(0) {
		v1806 = v9
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v59 | int32(2)
	v65 = v55
	goto L4
L10:
	;
	v65 = v9
	goto L4
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if int32(0) < v73 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v86 = v9
	v100 = v9
	v108 = v9
	goto L15
L13:
	;
	v1769 = v9
	goto L14
L14:
	;
	F_ReleaseCatCacheList(m, v71)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L246
	}
L15:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v71-int32(-64)+v108<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+72))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
	v119 = v117 + v118
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+104)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v121 | int32(4)
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v1769 = v1730
	goto L14
L17:
	;
	v1758 = v108 + int32(1)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if v1758 < v1759 {
		v86 = v1730
		v100 = v1744
		v108 = v1758
		goto L15
	} else {
		goto L245
	}
L18:
	;
	v1730 = v86
	v1744 = v1709
	goto L17
L19:
	;
	v225 = v119 + int32(136)
	v227 = v116 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v121 | int32(12)
	if l5 == int32(0) {
		v256 = v120
		v257 = v225
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+68))
	if v126 != v65 {
		v1709 = v100
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_FuncnameGetCandidates[0]))
	if v129 == int32(0) {
		v1709 = v100
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v195 = int32(0)
	goto L19
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v132 <= int32(0) {
		v1709 = v100
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v135 = int32(0)
	if v135 < v132 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v138 = v132
	goto L28
L27:
	;
	v138 = v135
	goto L28
L28:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)+68))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_FuncnameGetCandidates[1]))
	v151 = int32(0)
	goto L29
L29:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v139+v151<<(uint(int32(2))%32))))
	if base.B2i32(v183 == v141)&base.B2i32(v141 != v143) != 0 {
		v195 = v151
		goto L19
	} else {
		goto L31
	}
L30:
	;
	v1709 = v100
	goto L18
L31:
	;
	v187 = v151 + int32(1)
	if v187 != v138 {
		v151 = v187
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v867
	if l1 < v256 {
		goto L126
	} else {
		goto L127
	}
L34:
	;
	v752 = v716 | int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v752
	if base.B2i32(v256 <= v293)|base.B2i32(v256 <= l1) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L35:
	;
	v570 = v499
	v573 = v504
	v578 = v504
	goto L88
L36:
	;
	v536 = l4 & base.B2i32(l1 < v256)
	if v536 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L37:
	;
	v532 = int32(0)
	v533 = v508
	v534 = v100
	goto L36
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L79
	}
L39:
	;
	if l2 != 0 {
		goto L48
	} else {
		goto L49
	}
L40:
	;
	v237 = F_SysCacheGetAttr(m, int32(46), v227, int32(21), v38+int32(28))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+28)))
	if v239 != 0 {
		v256 = v120
		v257 = v225
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v241 = F_pg_detoast_datum(m, base.I32_wrap_i64(v237))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if v243 != int32(1) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	if v246 < int32(0) {
		goto L38
	} else {
		goto L45
	}
L45:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	if v249 != 0 {
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	if v250 != int32(26) {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v256 = v246
	v257 = v241 + int32(24)
	goto L39
L48:
	;
	v259 = l4 & base.B2i32(l1 < v256)
	if v259 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v508 = int32(0)
	if base.B2i32(l3 == v508)|base.B2i32(l1 < v256) != 0 {
		goto L37
	} else {
		goto L78
	}
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v264 | int32(16)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v116)+72))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+22)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v275 = F_SysCacheGetAttr(m, int32(47), v227, int32(23), v38+int32(19))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L57
	}
L52:
	;
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+106)))
	if v256 <= l1+v260 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if l1 != v256 {
		v1709 = v100
		goto L18
	} else {
		goto L56
	}
L55:
	;
	v1709 = v100
	goto L18
L56:
	;
	goto L51
L57:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+19)))
	if v277 != 0 {
		v1709 = v100
		goto L18
	} else {
		goto L58
	}
L58:
	;
	v284 = F_get_func_arg_info(m, v227, v38+int32(28), v38+int32(24), v38+int32(20))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v287 = F_palloc_mul(m, int32(4), v256)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v290 = F_palloc0_mul(m, int32(1), v284)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v292 = int32(0)
	v293 = l1 - v270
	if v293 <= v292 {
		v499 = v292
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v500 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v297 = v293 & int32(3)
	v298 = int32(0)
	if base.Ui32(v270-l1) <= base.Ui32(int32(-4)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v306 = v298
	v321 = int32(0)
	goto L67
L65:
	;
	v382 = v298
	goto L66
L66:
	;
	v417 = v382
	v425 = v298
	goto L71
L67:
	;
	v341 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v287+v306<<(uint(v341)%32)))) = v306
	v346 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v306+v290))) = uint8(v346)
	v349 = v306 | v346
	*(*int32)(unsafe.Add(mBase, uint32(v287+v349<<(uint(v341)%32)))) = v349
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v290))) = uint8(v346)
	v358 = v306 | v341
	*(*int32)(unsafe.Add(mBase, uint32(v287+v358<<(uint(v341)%32)))) = v358
	*(*uint8)(unsafe.Add(mBase, uint32(v358+v290))) = uint8(v346)
	v367 = v306 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v287+v367<<(uint(v341)%32)))) = v367
	*(*uint8)(unsafe.Add(mBase, uint32(v367+v290))) = uint8(v346)
	v375 = int32(4)
	v376 = v306 + v375
	v378 = v321 + v375
	if v378 != v293&int32(2147483644) {
		v306 = v376
		v321 = v378
		goto L67
	} else {
		goto L69
	}
L68:
	;
	if v297 == int32(0) {
		v499 = v293
		goto L62
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v382 = v376
	goto L66
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287+v417<<(uint(int32(2))%32)))) = v417
	v457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v417+v290))) = uint8(v457)
	v462 = v425 + v457
	if v462 != v297 {
		v417 = v417 + v457
		v425 = v462
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v499 = v293
	goto L62
L73:
	;
	goto L72
L74:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v716 = v503
	v728 = v499
	goto L34
L75:
	;
	goto L76
L76:
	;
	v504 = int32(0)
	if v504 < v284 {
		goto L35
	} else {
		goto L77
	}
L77:
	;
	v1709 = v100
	goto L18
L78:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v119)+88))
	v515 = base.B2i32(v513 != int32(0))
	v532 = v513
	v533 = v515
	v534 = v100 | v515
	goto L36
L79:
	;
	F_errmsg_internal(m, int32(_a_F_FuncnameGetCandidates_0), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_FuncnameGetCandidates_1), int32(1306), int32(_a_F_FuncnameGetCandidates_2))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v549 = int32(0)
	if base.B2i32((v536|(base.B2i32(l1 == v256)|v533))&int32(1) == v549)&base.B2i32(v549 <= l1) != 0 {
		v1709 = v543
		goto L18
	} else {
		goto L87
	}
L83:
	;
	v543 = v534
	goto L82
L84:
	;
	goto L85
L85:
	;
	v540 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+106)))
	if v256 <= l1+v540 {
		v543 = int32(1)
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v1709 = v534
	goto L18
L87:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v857 = v532
	v866 = int32(0)
	v867 = v554 | int32(16)
	v869 = v533
	v879 = v543
	v884 = v536
	goto L33
L88:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v594 = int32(0)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v597+v573<<(uint(int32(2))%32))))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v605 = v594
	v613 = v594
	goto L90
L89:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v710 | int32(32)
	if v705&int32(1) != 0 {
		v1709 = v100
		goto L18
	} else {
		goto L111
	}
L90:
	;
	if l5|base.B2i32(v593 == v594) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v695 = v613 + v290
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	v697 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v697)
	*(*int32)(unsafe.Add(mBase, uint32(v287+v570<<(uint(int32(2))%32)))) = v613
	v704 = v570 + v697
	v705 = v578 | v696
	v707 = v573 + v697
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v707 < v708 {
		v570 = v704
		v573 = v707
		v578 = v705
		goto L88
	} else {
		goto L110
	}
L92:
	;
	goto L91
L93:
	;
	v693 = v605 + int32(1)
	if v693 != v284 {
		v605 = v693
		v613 = v690
		goto L90
	} else {
		goto L109
	}
L94:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605+v593))))
	v645 = v643 - int32(98)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v645))|base.B2i32(int32(1)<<(uint(v645)%32)&int32(_a_F_FuncnameGetCandidates_3) == int32(0)) != 0 {
		v690 = v613
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v602+v605<<(uint(int32(2))%32))))
	if v659 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	if base.B2i32(v662 == int32(0))|base.B2i32(v662 != v665) != 0 {
		v683 = v662
		v684 = v665
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v690 = v613 + int32(1)
	goto L93
L101:
	;
	if v683-v684 == int32(0) {
		goto L92
	} else {
		goto L108
	}
L102:
	;
	goto L101
L103:
	;
	v668 = v659
	v669 = v601
	goto L104
L104:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+1)))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+1)))
	if v673 == int32(0) {
		v683 = v673
		v684 = v672
		goto L102
	} else {
		goto L106
	}
L105:
	;
	v683 = v673
	v684 = v672
	goto L102
L106:
	;
	v676 = int32(1)
	if v673 == v672 {
		v668 = v668 + v676
		v669 = v669 + v676
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	goto L100
L109:
	;
	v1709 = v100
	goto L18
L110:
	;
	goto L89
L111:
	;
	v716 = v710
	v728 = v704
	goto L34
L112:
	;
	v760 = int32(*(*int16)(unsafe.Add(mBase, uint32(v268+v269)+106)))
	v772 = v293
	v774 = v728
	goto L115
L113:
	;
	v813 = v752
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v813 | int32(128)
	if l3 != 0 {
		goto L122
	} else {
		goto L123
	}
L115:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772+v290))))
	if v798 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v813 = v812
	goto L114
L117:
	;
	if v772 < v256-v760 {
		v1709 = v100
		goto L18
	} else {
		goto L120
	}
L118:
	;
	v808 = v774
	goto L119
L119:
	;
	v810 = v772 + int32(1)
	if v810 != v256 {
		v772 = v810
		v774 = v808
		goto L115
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287+v774<<(uint(int32(2))%32)))) = v772
	v808 = v774 + int32(1)
	goto L119
L121:
	;
	goto L116
L122:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v119)+88))
	if v851 != 0 {
		v1709 = v100
		goto L18
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v855 = int32(0)
	v857 = v855
	v866 = v287
	v867 = v813 | int32(384)
	v869 = v855
	v879 = int32(1)
	v884 = v259
	goto L33
L125:
	;
	goto L124
L126:
	;
	v894 = v256
	goto L128
L127:
	;
	v894 = l1
	goto L128
L128:
	;
	v896 = v894 << (uint(int32(2)) % 32)
	v899 = F_palloc(m, v896+int32(32))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899)+4)) = v195
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v899)+28)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v899)+16)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v899)+12)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v899)+8)) = v902
	if v866 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v869 != 0 {
		goto L147
	} else {
		goto L148
	}
L131:
	;
	if v256 <= int32(0) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v1088 = v256 << (uint(int32(2)) % 32)
	if v1088 == int32(0) {
		goto L130
	} else {
		goto L145
	}
L134:
	;
	v910 = v256 & int32(3)
	v912 = v899 + int32(32)
	v913 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v256) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v930 = v913
	v945 = int32(0)
	goto L138
L136:
	;
	v1012 = v913
	goto L137
L137:
	;
	v1047 = v1012
	v1052 = v913
	goto L142
L138:
	;
	v955 = int32(2)
	v956 = v930 << (uint(v955) % 32)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v866+v956)))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v257+v959<<(uint(v955)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v956))) = v963
	v965 = int32(4)
	v966 = v956 | v965
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v866+v966)))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v257+v969<<(uint(v955)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v966))) = v973
	v976 = v956 | int32(8)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v866+v976)))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v257+v979<<(uint(v955)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v976))) = v983
	v986 = v956 | int32(12)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v866+v986)))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v257+v989<<(uint(v955)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v986))) = v993
	v996 = v930 + v965
	v998 = v945 + v965
	if v998 != v256&int32(2147483644) {
		v930 = v996
		v945 = v998
		goto L138
	} else {
		goto L140
	}
L139:
	;
	if v910 == int32(0) {
		goto L130
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v1012 = v996
	goto L137
L142:
	;
	v1072 = int32(2)
	v1073 = v1047 << (uint(v1072) % 32)
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v866+v1073)))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v257+v1076<<(uint(v1072)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v1073))) = v1080
	v1082 = int32(1)
	v1085 = v1052 + v1082
	if v1085 != v910 {
		v1047 = v1047 + v1082
		v1052 = v1085
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L130
L144:
	;
	goto L143
L145:
	;
	base.MemoryCopy(m, v899+int32(32), v257, v1088)
	goto L130
L146:
	;
	v1307 = int32(0)
	if v884 != 0 {
		goto L160
	} else {
		goto L161
	}
L147:
	;
	v1129 = v894 - v256
	v1130 = int32(1)
	v1131 = v1129 + v1130
	*(*int32)(unsafe.Add(mBase, uint32(v899)+20)) = v1131
	v1134 = v899 + int32(32)
	v1136 = v256 - v1130
	v1139 = v1131 & int32(7)
	if v1139 != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899)+20)) = int32(0)
	goto L146
L150:
	;
	v1149 = int32(0)
	v1150 = v1136
	goto L153
L151:
	;
	v1194 = v1136
	goto L152
L152:
	;
	if base.Ui32(v1129) < base.Ui32(int32(7)) {
		goto L146
	} else {
		goto L156
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134+v1150<<(uint(int32(2))%32)))) = v857
	v1179 = int32(1)
	v1180 = v1150 + v1179
	v1182 = v1149 + v1179
	if v1182 != v1139 {
		v1149 = v1182
		v1150 = v1180
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v1194 = v1180
	goto L152
L155:
	;
	goto L154
L156:
	;
	v1231 = v1194
	goto L157
L157:
	;
	v1258 = v1134 + v1231<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1258))) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+28)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+24)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+20)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+16)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+12)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+8)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+4)) = v857
	v1268 = v1231 + int32(8)
	if v1268 != v894 {
		v1231 = v1268
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L146
L159:
	;
	goto L158
L160:
	;
	v1310 = v256 - l1
	goto L162
L161:
	;
	v1310 = v1307
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899)+24)) = v1310
	if v86 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if (v879^int32(1))&base.B2i32(v65 != int32(0)) != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v1660 = v1307
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v1660
	v1730 = v899
	v1744 = v879
	goto L17
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v86
	v1730 = v899
	v1744 = v879
	goto L17
L167:
	;
	goto L168
L168:
	;
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+53)))
	if (v1318^int32(-1)|v879)&int32(1) == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+4))
	if v1546 == v195 {
		goto L225
	} else {
		goto L226
	}
L170:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	if v1326 != v894 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1399 = v899 + int32(32)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v899)+16))
	v1401 = v1400 - v1310
	v1403 = v1401 << (uint(int32(2)) % 32)
	v1404 = v86
	goto L195
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v86
	v1730 = v899
	v1744 = v879
	goto L17
L174:
	;
	goto L175
L175:
	;
	v1329 = int32(32)
	v1330 = v899 + v1329
	v1332 = v86 + v1329
	if base.Ui32(int32(4)) <= base.Ui32(v896) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	if v1394 == int32(0) {
		v1511 = v86
		goto L169
	} else {
		goto L194
	}
L177:
	;
	v1394 = int32(0)
	goto L176
L178:
	;
	v1368 = v1363
	v1369 = v1364
	v1370 = v1365
	goto L188
L179:
	;
	if (v1330|v1332)&int32(3) != 0 {
		v1363 = v1330
		v1364 = v1332
		v1365 = v896
		goto L178
	} else {
		goto L182
	}
L180:
	;
	v1356 = v1330
	v1357 = v1332
	v1358 = v896
	goto L181
L181:
	;
	if v1358 == int32(0) {
		goto L177
	} else {
		goto L187
	}
L182:
	;
	v1340 = v1330
	v1341 = v1332
	v1342 = v896
	goto L183
L183:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1340)))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1341)))
	if v1345 != v1346 {
		v1363 = v1340
		v1364 = v1341
		v1365 = v1342
		goto L178
	} else {
		goto L185
	}
L184:
	;
	v1356 = v1351
	v1357 = v1349
	v1358 = v1353
	goto L181
L185:
	;
	v1348 = int32(4)
	v1349 = v1341 + v1348
	v1351 = v1340 + v1348
	v1353 = v1342 - v1348
	if base.Ui32(int32(3)) < base.Ui32(v1353) {
		v1340 = v1351
		v1341 = v1349
		v1342 = v1353
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v1363 = v1356
	v1364 = v1357
	v1365 = v1358
	goto L178
L188:
	;
	v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368))))
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369))))
	if v1373 == v1374 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v1394 = v1373 - v1374
	goto L176
L190:
	;
	v1376 = int32(1)
	v1381 = v1370 - v1376
	if v1381 != 0 {
		v1368 = v1368 + v1376
		v1369 = v1369 + v1376
		v1370 = v1381
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	goto L177
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v86
	v1730 = v899
	v1744 = v879
	goto L17
L195:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+16))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+24))
	if v1439-v1440 == v1401 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v86
	v1730 = v899
	v1744 = v879
	goto L17
L197:
	;
	v1444 = v1404 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v1403) {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	goto L199
L199:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1404)))
	if v1509 != 0 {
		v1404 = v1509
		goto L195
	} else {
		goto L219
	}
L200:
	;
	if v1506 == int32(0) {
		v1511 = v1404
		goto L169
	} else {
		goto L218
	}
L201:
	;
	v1506 = int32(0)
	goto L200
L202:
	;
	v1480 = v1475
	v1481 = v1476
	v1482 = v1477
	goto L212
L203:
	;
	if (v1399|v1444)&int32(3) != 0 {
		v1475 = v1399
		v1476 = v1444
		v1477 = v1403
		goto L202
	} else {
		goto L206
	}
L204:
	;
	v1468 = v1399
	v1469 = v1444
	v1470 = v1403
	goto L205
L205:
	;
	if v1470 == int32(0) {
		goto L201
	} else {
		goto L211
	}
L206:
	;
	v1452 = v1399
	v1453 = v1444
	v1454 = v1403
	goto L207
L207:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1453)))
	if v1457 != v1458 {
		v1475 = v1452
		v1476 = v1453
		v1477 = v1454
		goto L202
	} else {
		goto L209
	}
L208:
	;
	v1468 = v1463
	v1469 = v1461
	v1470 = v1465
	goto L205
L209:
	;
	v1460 = int32(4)
	v1461 = v1453 + v1460
	v1463 = v1452 + v1460
	v1465 = v1454 - v1460
	if base.Ui32(int32(3)) < base.Ui32(v1465) {
		v1452 = v1463
		v1453 = v1461
		v1454 = v1465
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v1475 = v1468
	v1476 = v1469
	v1477 = v1470
	goto L202
L212:
	;
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1480))))
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1481))))
	if v1485 == v1486 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1506 = v1485 - v1486
	goto L200
L214:
	;
	v1488 = int32(1)
	v1493 = v1482 - v1488
	if v1493 != 0 {
		v1480 = v1480 + v1488
		v1481 = v1481 + v1488
		v1482 = v1493
		goto L212
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	goto L213
L217:
	;
	goto L201
L218:
	;
	goto L199
L219:
	;
	goto L196
L220:
	;
	if v86 == v1511 {
		goto L236
	} else {
		goto L237
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1511)+8)) = int32(0)
	F_pfree(m, v899)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L234
	}
L222:
	;
	if int32(0) < v1548 {
		goto L220
	} else {
		goto L233
	}
L223:
	;
	if int32(0) <= v1553 {
		goto L221
	} else {
		goto L232
	}
L224:
	;
	F_pfree(m, v899)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L231
	}
L225:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+20))
	if v869 == int32(0) {
		goto L222
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1553 = v195 - v1546
	if v1553 <= int32(0) {
		goto L223
	} else {
		goto L230
	}
L228:
	;
	if v1548 == int32(0) {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	goto L221
L230:
	;
	goto L224
L231:
	;
	v1709 = v879
	goto L18
L232:
	;
	goto L220
L233:
	;
	goto L221
L234:
	;
	v1709 = v879
	goto L18
L235:
	;
	F_pfree(m, v1511)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L244
	}
L236:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1648 = v1570
	goto L235
L237:
	;
	goto L238
L238:
	;
	v1577 = v86
	goto L240
L239:
	;
	v1648 = v86
	goto L235
L240:
	;
	if v1577 == int32(0) {
		goto L239
	} else {
		goto L242
	}
L241:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	*(*int32)(unsafe.Add(mBase, uint32(v1577))) = v1610
	goto L239
L242:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	if v1511 != v1608 {
		v1577 = v1608
		goto L240
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	v1660 = v1648
	goto L165
L245:
	;
	goto L16
L246:
	;
	v1806 = v1769
	goto L3
}
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	v5 = int32(0)
	if base.B2i32(l3 != int32(_a_F___fstatat_0))|base.B2i32(l0 < v5) == v5 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v13 != 0 {
			v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
			mBase = m.M
			v37 = v33
		} else {
			v14 = m.Env.X__syscall_fstat64(m, l0, l2)
			mBase = m.M
			v37 = v14
		}
	} else {
		if l0 != int32(-100) {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if base.B2i32(l3 == int32(0))&base.B2i32(v19 == int32(47)) != 0 {
				v31 = m.Env.X__syscall_stat64(m, l1, l2)
				mBase = m.M
				v37 = v31
			} else {
				if base.B2i32(l3 != int32(256))|base.B2i32(v19 != int32(47)) != 0 {
					v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v37 = v33
				} else {
					v35 = m.Env.X__syscall_lstat64(m, l1, l2)
					mBase = m.M
					v37 = v35
				}
			}
		} else {
			if l3 == int32(256) {
				v35 = m.Env.X__syscall_lstat64(m, l1, l2)
				mBase = m.M
				v37 = v35
			} else {
				if l3 != 0 {
					v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v37 = v33
				} else {
					v31 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v37 = v31
				}
			}
		}
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v37) {
		*(*int32)(unsafe.Add(mBase, _c_F___fstatat[0])) = int32(0) - v37
		v45 = int32(-1)
	} else {
		v45 = v37
	}
	return v45
}
func F_fastgetattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14262(m, l0, l1, l2, l3, int32(_a_F_fastgetattr_5_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_fdw_handler_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_fdw_handler_out_0), int32(369), int32(_a_F_fdw_handler_out_1), int32(_a_F_fdw_handler_out_2), int32(_a_F_fdw_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_finalize_primnode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(8) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L4
	case 15:
		goto L3
	}
L3:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104<<(uint(int32(2))%32)-int32(4))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_finalize_primnode(m, v111, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L28
	}
L4:
	;
	v97 = F_expression_tree_walker_impl(m, l0, int32(896), l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
	if v26 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_bms_add_member(m, v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18
	return int32(0)
L10:
	;
	if v81 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v81 = v70
	goto L10
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v29 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 != int32(1) {
		v70 = v3
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v35 <= int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(0)
	v46 = v35
	goto L17
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v70 = v66
	goto L11
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v43<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 == v54 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v81 = int32(0)
	goto L10
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v58 = F_equal(m, v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v61 = v46
	goto L21
L21:
	;
	v63 = v43 + int32(1)
	if v63 < v61 {
		v43 = v63
		v46 = v61
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v58 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v61 = v60
	goto L21
L24:
	;
	goto L18
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v86 = F_bms_add_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v86
	goto L4
L27:
	;
	return v97
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = F_finalize_primnode(m, v147, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L36
	}
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v122 = v3
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v122<<(uint(int32(2))%32))))
	v133 = F_bms_del_member(m, v127, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v133
	v137 = v122 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v137 < v138 {
		v122 = v137
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)+64))
	v151 = F_bms_copy(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v153 == int32(0) {
		v181 = v151
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v186 = F_bms_join(m, v185, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L45
	}
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v156 <= int32(0) {
		v181 = v151
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v162 = int32(0)
	v163 = v151
	goto L41
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v162<<(uint(int32(2))%32))))
	v172 = F_bms_del_member(m, v163, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	v181 = v172
	goto L38
L43:
	;
	v175 = v162 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v175 < v176 {
		v162 = v175
		v163 = v172
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v186
	goto L1
}
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	v5 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = v18 - v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = v21 + v18 - int32(1)
	v27 = l2
	v30 = v5
	v33 = v5
	v34 = v5
	goto L1
L1:
	;
	v45 = (v27-v30)>>(uint(int32(1))%32) + v30
	v48 = l1 + v45*int32(20)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v33 < v34 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v128 = int32(0)
	if base.B2i32(v109 == v112)|base.B2i32(v128 < v112) == v128 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	if int32(1) < v109-v112 {
		v27 = v109
		v30 = v112
		v33 = v115
		v34 = v116
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v109 = v27
	v112 = v45
	v115 = v94
	v116 = v34
	goto L3
L5:
	;
	v51 = v33
	goto L7
L6:
	;
	v51 = v34
	goto L7
L7:
	;
	v54 = v49 + (v51 ^ int32(-1))
	if v54 < int32(0) {
		v94 = v51
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v61 = v51
	v63 = v54
	goto L10
L9:
	;
	if int32(0) <= v81 {
		v94 = v61
		goto L4
	} else {
		goto L17
	}
L10:
	;
	if v19 == v18-v61 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v109 = v27
	v112 = v45
	v115 = v49
	v116 = v34
	goto L3
L12:
	;
	v109 = v45
	v112 = v30
	v115 = v33
	v116 = v20
	goto L3
L13:
	;
	goto L14
L14:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24-v61))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v63))))
	v81 = v77 - v80
	if v81 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v82 = int32(1)
	if int32(0) < v63 {
		v61 = v61 + v82
		v63 = v63 - v82
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	v109 = v45
	v112 = v30
	v115 = v33
	v116 = v61
	goto L3
L18:
	;
	goto L2
L19:
	;
	v135 = v109
	v138 = v112
	v141 = v115
	v142 = v116
	goto L22
L20:
	;
	v258 = v112
	v261 = v115
	goto L21
L21:
	;
	v277 = l1 + v258*int32(20)
	goto L42
L22:
	;
	v153 = (v135-v138)>>(uint(int32(1))%32) + v138
	v156 = l1 + v153*int32(20)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v141 < v142 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v258 = v238
	v261 = v241
	goto L21
L24:
	;
	if int32(1) < v235-v238 {
		v135 = v235
		v138 = v238
		v141 = v241
		v142 = v242
		goto L22
	} else {
		goto L41
	}
L25:
	;
	v235 = v135
	v238 = v153
	v241 = v232
	v242 = v142
	goto L24
L26:
	;
	v232 = v202
	goto L25
L27:
	;
	v159 = v141
	goto L29
L28:
	;
	v159 = v142
	goto L29
L29:
	;
	v162 = v157 + (v159 ^ int32(-1))
	if v162 < int32(0) {
		v202 = v159
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v169 = v159
	v171 = v162
	goto L31
L31:
	;
	if v19 == v18-v169 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v232 = v157
	goto L25
L33:
	;
	v235 = v153
	v238 = v138
	v241 = v141
	v242 = v20
	goto L24
L34:
	;
	goto L35
L35:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24-v169))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v171))))
	v189 = v185 - v188
	if v189 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if int32(0) <= v189 {
		v202 = v169
		goto L26
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v192 = int32(1)
	if int32(0) < v171 {
		v169 = v169 + v192
		v171 = v171 - v192
		goto L31
	} else {
		goto L40
	}
L39:
	;
	v235 = v153
	v238 = v138
	v241 = v141
	v242 = v169
	goto L24
L40:
	;
	goto L32
L41:
	;
	goto L23
L42:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	if v261 < v290 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	return int32(0)
L44:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	if v308 != 0 {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v290
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v294 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v294
	v296 = m.T0[l3].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	return v305
L49:
	;
	return int32(0)
L50:
	;
	if v296 == int32(0) {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v302
	goto L48
L52:
	;
	v277 = v277 + v308*int32(20)
	goto L42
L53:
	;
	goto L54
L54:
	;
	goto L43
}
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = F_getBaseType(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = v5
	goto L3
L3:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = v14
	goto L3
L6:
	;
	v19 = F_getBaseType(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v21 = v5
	goto L8
L8:
	;
	if v18 == v21 {
		v152 = int32(2)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v21 = v19
	goto L8
L10:
	;
	m.G0 = v10 + int32(32)
	return v152
L11:
	;
	v24 = int32(0)
	v25 = int32(2281)
	if base.B2i32(v18 == v25)|base.B2i32(v21 == v25) != 0 {
		v152 = v24
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v33 = F_SearchSysCache2(m, int32(12), base.I64_extend_i32_u(v18), base.I64_extend_i32_u(v21))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	if v140 != 0 {
		goto L54
	} else {
		goto L55
	}
L14:
	;
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+16)))
	switch v39 - int32(97) {
	case 0:
		v57 = int32(1)
		goto L19
	default:
		goto L21
	case 4:
		goto L20
	case 8:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v21&int32(-9) == int32(22) {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+17)))
	switch v63 - int32(98) {
	case 0:
		goto L30
	default:
		goto L29
	case 4:
		goto L28
	case 7:
		v88 = int32(4)
		goto L27
	}
L19:
	;
	if base.Ui32(v57) <= base.Ui32(l2) {
		goto L18
	} else {
		goto L25
	}
L20:
	;
	v57 = int32(3)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46
	F_errmsg_internal(m, int32(_a_F_find_coercion_pathway_3), v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_find_coercion_pathway_1), int32(3208), int32(_a_F_find_coercion_pathway_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v140 = v24
	goto L13
L27:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L35
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v85
	v88 = int32(1)
	goto L27
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v140 = int32(2)
	goto L13
L32:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+17)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v73
	F_errmsg_internal(m, int32(_a_F_find_coercion_pathway_0), v10+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_find_coercion_pathway_1), int32(3230), int32(_a_F_find_coercion_pathway_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v140 = v88
	goto L13
L36:
	;
	v112 = int32(0)
	if l2 == v112 {
		v140 = v112
		goto L13
	} else {
		goto L44
	}
L37:
	;
	v95 = F_get_element_type(m, v21)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v95 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v99 = F_get_element_type(m, v18)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v99 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v105 = F_find_coercion_pathway(m, v95, v99, l2, v10+int32(24))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v105 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v140 = int32(3)
	goto L13
L44:
	;
	F_get_type_category_preferred(m, v21, v10+int32(24), v10+int32(31))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v121 == int32(83) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v140 = int32(4)
	goto L13
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(l2) < base.Ui32(int32(3)) {
		v140 = v112
		goto L13
	} else {
		goto L49
	}
L49:
	;
	F_get_type_category_preferred(m, v18, v10+int32(24), v10+int32(31))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v135 == int32(83) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v138 = int32(4)
	goto L53
L52:
	;
	v138 = int32(0)
	goto L53
L53:
	;
	v140 = v138
	goto L13
L54:
	;
	v144 = v140
	goto L56
L55:
	;
	v144 = int32(4)
	goto L56
L56:
	;
	if l2 == int32(2) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v147 = v144
	goto L59
L58:
	;
	v147 = v140
	goto L59
L59:
	;
	v152 = v147
	goto L10
}
func F_find_nonnullable_rels(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_find_nonnullable_rels_walker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[0]))
	v21 = F_hash_search(m, v16, v12+int32(8), v5, v5)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v353
L2:
	;
	if l1 != 0 {
		goto L83
	} else {
		goto L84
	}
L3:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v248 = F_find_option(m, v246, int32(0), l2, l3)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L80
	}
L4:
	;
	v146 = int32(0)
	v153 = v27
	goto L52
L5:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142 != 0 {
		goto L4
	} else {
		goto L50
	}
L6:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	return int32(0)
L8:
	;
	if v21 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v353 = v30
	goto L1
L12:
	;
	v141 = int32(_a_F_find_option_0)
	goto L5
L13:
	;
	v39 = v28 | int32(32)
	goto L15
L14:
	;
	v39 = v28
	goto L15
L15:
	;
	if v39 != int32(115) {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v43 == int32(0) {
		v141 = int32(_a_F_find_option_1)
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32((v43-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v43 | int32(32)
	goto L20
L19:
	;
	v54 = v43
	goto L20
L20:
	;
	if v54 != int32(111) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)))
	if v58 == int32(0) {
		v141 = int32(_a_F_find_option_2)
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v69 = v58 | int32(32)
	goto L25
L24:
	;
	v69 = v58
	goto L25
L25:
	;
	if v69 != int32(114) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+3)))
	if v73 == int32(0) {
		v141 = int32(_a_F_find_option_3)
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v84 = v73 | int32(32)
	goto L30
L29:
	;
	v84 = v73
	goto L30
L30:
	;
	if v84 != int32(116) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v88 == int32(0) {
		v141 = int32(_a_F_find_option_4)
		goto L5
	} else {
		goto L32
	}
L32:
	;
	if v88 != int32(95) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)))
	if v94 == int32(0) {
		v141 = int32(_a_F_find_option_5)
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v105 = v94 | int32(32)
	goto L37
L36:
	;
	v105 = v94
	goto L37
L37:
	;
	if v105 != int32(109) {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v109 == int32(0) {
		v141 = int32(_a_F_find_option_6)
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = v109 | int32(32)
	goto L42
L41:
	;
	v120 = v109
	goto L42
L42:
	;
	if v120 != int32(101) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+7)))
	if v124 == int32(0) {
		v141 = int32(_a_F_find_option_7)
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if base.Ui32((v124-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v135 = v124 | int32(32)
	goto L47
L46:
	;
	v135 = v124
	goto L47
L47:
	;
	if v135 != int32(109) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	if v138 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v141 = int32(_a_F_find_option_8)
	goto L5
L50:
	;
	v245 = int32(_a_F_find_option_9)
	goto L3
L51:
	;
	v193 = int32(0)
	v200 = v27
	goto L66
L52:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v155 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v146 != int32(10) {
		goto L51
	} else {
		goto L65
	}
L54:
	;
	if v146 == int32(10) {
		goto L51
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	v160 = int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+uint32(_c_F_find_option[1]))))
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = v164 | int32(32)
	goto L60
L59:
	;
	v173 = v164
	goto L60
L60:
	;
	v174 = int32(255)
	if base.Ui32((v155-int32(65))&v174) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v184 = v155 | int32(32)
	goto L63
L62:
	;
	v184 = v155
	goto L63
L63:
	;
	if v173&v174 == v184 {
		v146 = v146 + v160
		v153 = v153 + v160
		goto L52
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	v245 = int32(_a_F_find_option_10)
	goto L3
L66:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v202 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v193 != int32(14) {
		goto L2
	} else {
		goto L79
	}
L68:
	;
	if v193 == int32(14) {
		goto L2
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v207 = int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+uint32(_c_F_find_option[2]))))
	if base.Ui32((v211-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v220 = v211 | int32(32)
	goto L74
L73:
	;
	v220 = v211
	goto L74
L74:
	;
	v221 = int32(255)
	if base.Ui32((v202-int32(65))&v221) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v231 = v202 | int32(32)
	goto L77
L76:
	;
	v231 = v202
	goto L77
L77:
	;
	if v220&v221 == v231 {
		v193 = v193 + v207
		v200 = v200 + v207
		goto L66
	} else {
		goto L78
	}
L78:
	;
	goto L2
L79:
	;
	v245 = int32(_a_F_find_option_11)
	goto L3
L80:
	;
	v353 = v248
	goto L1
L81:
	;
	F_pfree(m, v264)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L118
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+4)) = v264
	v353 = v264
	goto L1
L83:
	;
	v253 = F_assignable_custom_variable_name(m, v27, l2, l3)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v329 = int32(0)
	if l2 != 0 {
		v353 = v329
		goto L1
	} else {
		goto L112
	}
L86:
	;
	if v253 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v353 = int32(0)
	goto L1
L88:
	;
	goto L89
L89:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[3]))
	v264 = F_MemoryContextAllocExtended(m, v261, int32(156), int32(2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	if v264 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v269 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	base.MemoryFill(m, v264, int32(0), int32(156))
	v288 = F_guc_strdup(m, l3, v259)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L7
	} else {
		goto L99
	}
L94:
	;
	if v269 == int32(0) {
		v353 = v258
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(_a_F_find_option_12))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_find_option_13), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(646), int32(_a_F_find_option_15))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	v353 = v258
	goto L1
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v288
	if v288 == int32(0) {
		goto L81
	} else {
		goto L100
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+20)) = int64(12884902532)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = int32(_a_F_find_option_16)
	*(*int64)(unsafe.Add(mBase, uint32(v264)+4)) = int64(201863462918)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+96)) = v264 + int32(152)
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[0]))
	v307 = F_hash_search(m, v303, v264, int32(3), v12+int32(15))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	if v307 != 0 {
		goto L82
	} else {
		goto L102
	}
L102:
	;
	v310 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	if v310 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_errcode(m, int32(_a_F_find_option_12))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v324 == int32(0) {
		goto L81
	} else {
		goto L110
	}
L107:
	;
	F_errmsg(m, int32(_a_F_find_option_13), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(941), int32(_a_F_find_option_17))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	F_pfree(m, v324)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	goto L81
L112:
	;
	v331 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	if v331 == int32(0) {
		v353 = v329
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v338
	F_errmsg(m, int32(_a_F_find_option_18), v12)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(1157), int32(_a_F_find_option_19))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v353 = v329
	goto L1
L118:
	;
	v353 = int32(0)
	goto L1
}
func F_find_typed_table_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v14 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v10 + int32(16)
	F_ScanKeyInit(m, v19, int32(5), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_table_beginscan_catalog(m, v14, int32(1), v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+188))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	m.T0[v61].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v29 = F_heap_getnext(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v29 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = int32(0)
	goto L5
L9:
	;
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(0)
	v43 = v29
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45)))
	v48 = F_lappend_oid(m, v37, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v52 = v48
	goto L5
L14:
	;
	v50 = F_heap_getnext(m, v27)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v50 != 0 {
		v37 = v48
		v43 = v50
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(80)
	return v52
L19:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg(m, int32(_a_F_find_typed_table_dependencies_0), v10)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(_a_F_find_typed_table_dependencies_1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_find_typed_table_dependencies_2), int32(_a_F_find_typed_table_dependencies_3), int32(_a_F_find_typed_table_dependencies_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fix_opfuncids(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_fix_opfuncids_walker(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_flatCopyTargetEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	v4 = F_palloc0(m, int32(28))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(62)
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v16
		return v4
	}
}
func F_float48le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	var v22 int64
	_ = v22
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = base.F64_promote_f32(v10)
		v22 = base.I64_extend_i32_u(base.F64_ge(v4, v11) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))))
	} else {
		v22 = int64(1)
	}
	return v22
}
func F_float48ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 float32
	_ = v9
	var v10 float64
	_ = v10
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(9223372036854775807)
	v8 = base.I64_reinterpret_f64(v5) & v7
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = base.F64_promote_f32(v9)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v10)&v7) {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313))))
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v8)) | base.F64_ne(v5, v10))
	}
}
func F_float4_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float32
	_ = v8
	var v9 int32
	_ = v9
	var v10 float32
	_ = v10
	var v11 float32
	_ = v11
	var v12 float32
	_ = v12
	var v13 float32
	_ = v13
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v44 float32
	_ = v44
	var v53 int32
	_ = v53
	v6 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = base.F32_reinterpret_i32(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = base.F32_reinterpret_i32(v9)
	v11 = base.F32_sub(v8, v10)
	v12 = base.F32_abs(v11)
	v13 = math.Float32frombits(uint32(0x7f800000))
	if base.B2i32(base.F32_ne(v12, v13)|base.F32_eq(base.F32_abs(v8), v13) == v6)&base.F32_ne(base.F32_abs(v10), v13) == v6 {
		if base.Ui32(base.I32_reinterpret_f32(v12)) < base.Ui32(int32(2139095041)) {
			v44 = v11
		} else {
			v30 = int32(2147483647)
			v31 = v9 & v30
			if base.Ui32(int32(2139095041)) <= base.Ui32(v7&v30) {
				if base.Ui32(v31) <= base.Ui32(int32(2139095040)) {
					v44 = math.Float32frombits(uint32(0x7f800000))
				} else {
					v44 = float32(0)
				}
			} else {
				if base.Ui32(int32(2139095041)) <= base.Ui32(v31) {
					v44 = math.Float32frombits(uint32(0x7f800000))
				} else {
					v44 = float32(0)
				}
			}
		}
		return base.I64_extend_i32_u(base.I32_reinterpret_f32(v44) & int32(2147483647))
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float4_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 float32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 float64
	_ = v150
	var v154 float32
	_ = v154
	var v155 float64
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v22 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v28 == int32(18) {
				v31 = int32(16)
			} else {
				v31 = int32(0)
			}
			if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v38 = int32(4)
			} else {
				v38 = v31
			}
			v51 = v38
		} else {
			v39 = int32(1)
			if v22&v39 != 0 {
				v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v51-int32(1)) {
			v57 = F_cstring_to_text(m, int32(_a_F_float4_to_char_0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v250 = v57
				m.G0 = v14 + int32(96)
				return base.I64_extend_i32_u(v250)
			}
		} else {
			v59 = base.F32_reinterpret_i32(v16)
			v64 = F_palloc0(m, v51<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int64(0)
			} else {
				v70 = F_NUM_cache(m, v51, v14+int32(60), v18, v14+int32(59))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v72&int32(1024) != 0 {
						v76 = int32(2147483647)
						v77 = base.F32_nearest(v59)
						if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v77)&v76) {
							v84 = v76
						} else {
							v84 = base.I32_trunc_sat_f32_s(v77)
						}
						if base.F32_ge(v77, float32(-2.1474836e+09)) != 0 {
							v88 = v84
						} else {
							v88 = int32(2147483647)
						}
						if base.F32_lt(v77, float32(2.1474836e+09)) != 0 {
							v92 = v88
						} else {
							v92 = int32(2147483647)
						}
						v93 = F_int_to_roman(m, v92)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int64(0)
						} else {
							v219 = v93
							v220 = int32(0)
							v224 = v2
							v230 = v64 + int32(4)
							F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
							mBase = m.M
							v234 = m.ExcPending
							if v234 != 0 {
								return int64(0)
							} else {
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v235 == int32(1) {
									F_pfree(m, v70)
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return int64(0)
									} else {
										v240 = F_strlen(m, v230)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
										v250 = v64
										m.G0 = v14 + int32(96)
										return base.I64_extend_i32_u(v250)
									}
								} else {
									v240 = F_strlen(m, v230)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
									v250 = v64
									m.G0 = v14 + int32(96)
									return base.I64_extend_i32_u(v250)
								}
							}
						}
					} else {
						if v72&int32(_a_F_float4_to_char_1) != 0 {
							if base.B2i32(base.Ui32(v16&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_ne(base.F32_abs(v59), math.Float32frombits(uint32(0x7f800000))) == int32(0) {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v109 = v107 + v108
								v112 = F_palloc(m, v109+int32(7))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									v115 = v109 + int32(6)
									if v115 != 0 {
										base.MemoryFill(m, v112, int32(35), v115)
									} else {
									}
									v118 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v115))) = uint8(v118)
									v122 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v122)
									v125 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v107)+1)) = uint8(v125)
									v219 = v112
									v220 = v118
									v224 = v2
									v230 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int64(0)
									} else {
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v235 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int64(0)
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v250)
											}
										} else {
											v240 = F_strlen(m, v230)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
											v250 = v64
											m.G0 = v14 + int32(96)
											return base.I64_extend_i32_u(v250)
										}
									}
								}
							} else {
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v127
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = base.F64_promote_f32(v59)
								v131 = int32(0)
								v135 = F_psprintf(m, int32(_a_F_float4_to_char_2), v14+int32(32))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return int64(0)
								} else {
									v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
									if v137 != int32(43) {
										v219 = v135
										v220 = v131
										v224 = v2
									} else {
										v140 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v140)
										v219 = v135
										v220 = v131
										v224 = v2
									}
									v230 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int64(0)
									} else {
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v235 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int64(0)
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v250)
											}
										} else {
											v240 = F_strlen(m, v230)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
											v250 = v64
											m.G0 = v14 + int32(96)
											return base.I64_extend_i32_u(v250)
										}
									}
								}
							}
						} else {
							if v72&int32(2048) != 0 {
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v144 + v145
								v150 = F_pow(m, float64(10), base.F64_convert_i32_s(v144))
								mBase = m.M
								v154 = base.F32_mul(v59, base.F32_demote_f64(v150))
							} else {
								v154 = v59
							}
							v155 = base.F64_promote_f32(v154)
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v155)
							v161 = F_psprintf(m, int32(_a_F_float4_to_char_3), v14+int32(16))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int64(0)
							} else {
								v163 = F_strlen(m, v161)
								mBase = m.M
								if base.Ui32(v163) <= base.Ui32(int32(5)) {
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if base.Ui32(v166+v163) < base.Ui32(int32(7)) {
										v174 = v166
									} else {
										v172 = int32(6) - v163
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v172
										v174 = v172
									}
								} else {
									v172 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v172
									v174 = v172
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v155
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v174
								v178 = F_psprintf(m, int32(_a_F_float4_to_char_4), v14)
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int64(0)
								} else {
									v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
									v182 = base.B2i32(v180 == int32(45))
									v183 = v178 + v182
									v184 = int32(46)
									v185 = F___strchrnul(m, v183, v184)
									mBase = m.M
									v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
									if v187 == v184 {
										v191 = v185
									} else {
										v191 = int32(0)
									}
									if v191 != 0 {
										v194 = v191 - v183
									} else {
										v193 = F_strlen(m, v183)
										mBase = m.M
										v194 = v193
									}
									if v180 == int32(45) {
										v197 = int32(45)
									} else {
										v197 = int32(43)
									}
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if base.Ui32(v194) < base.Ui32(v198) {
										v219 = v183
										v220 = v198 - v194
										v224 = v197
										v230 = v64 + int32(4)
										F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int64(0)
										} else {
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v235 == int32(1) {
												F_pfree(m, v70)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return int64(0)
												} else {
													v240 = F_strlen(m, v230)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
													v250 = v64
													m.G0 = v14 + int32(96)
													return base.I64_extend_i32_u(v250)
												}
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v250)
											}
										}
									} else {
										if base.Ui32(v194) <= base.Ui32(v198) {
											v219 = v183
											v220 = int32(0)
											v224 = v197
											v230 = v64 + int32(4)
											F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
											mBase = m.M
											v234 = m.ExcPending
											if v234 != 0 {
												return int64(0)
											} else {
												v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v235 == int32(1) {
													F_pfree(m, v70)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return int64(0)
													} else {
														v240 = F_strlen(m, v230)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
														v250 = v64
														m.G0 = v14 + int32(96)
														return base.I64_extend_i32_u(v250)
													}
												} else {
													v240 = F_strlen(m, v230)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
													v250 = v64
													m.G0 = v14 + int32(96)
													return base.I64_extend_i32_u(v250)
												}
											}
										} else {
											v203 = v174 + v198
											v206 = F_palloc(m, v203+int32(2))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int64(0)
											} else {
												v209 = v203 + int32(1)
												if v209 != 0 {
													base.MemoryFill(m, v206, int32(35), v209)
												} else {
												}
												v212 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v206+v209))) = uint8(v212)
												v217 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v206+v198))) = uint8(v217)
												v219 = v206
												v220 = v212
												v224 = v197
												v230 = v64 + int32(4)
												F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int64(0)
												} else {
													v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v235 == int32(1) {
														F_pfree(m, v70)
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return int64(0)
														} else {
															v240 = F_strlen(m, v230)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
															v250 = v64
															m.G0 = v14 + int32(96)
															return base.I64_extend_i32_u(v250)
														}
													} else {
														v240 = F_strlen(m, v230)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
														v250 = v64
														m.G0 = v14 + int32(96)
														return base.I64_extend_i32_u(v250)
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
			}
		}
	}
}
func F_float4div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v11 float32
	_ = v11
	var v17 float32
	_ = v17
	var v19 float32
	_ = v19
	var v25 float32
	_ = v25
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.B2i32(base.Ui32(base.I32_reinterpret_f32(v5)&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_eq(v11, float32(0)) == int32(0) {
		v17 = base.F32_div(v5, v11)
		v19 = math.Float32frombits(uint32(0x7f800000))
		if base.F32_eq(base.F32_abs(v17), v19)&base.F32_ne(base.F32_abs(v5), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = float32(0)
			if base.B2i32(base.F32_eq(v5, v25)|base.F32_ne(v17, v25) == int32(0))&base.F32_ne(base.F32_abs(v11), math.Float32frombits(uint32(0x7f800000))) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				return base.I64_extend_i32_s(base.I32_reinterpret_f32(v17))
			}
		}
	} else {
		F_float_zero_divide_error(m)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float4in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 float32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v128 float32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 float32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 float32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 float32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 float32
	_ = v520
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v543 float32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v569 float32
	_ = v569
	v9 = float32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = l0
	goto L3
L1:
	;
	m.G0 = v13 - int32(-64)
	return v569
L2:
	;
	v526 = v518
	goto L162
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(base.Ui32(v25-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v25 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v516 = v15 + v514
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v516
	v518 = v516
	v520 = v515
	goto L2
L5:
	;
	v15 = v15 + int32(1)
	goto L3
L6:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	goto L7
L9:
	;
	v37 = F_errsave_start(m, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_float4in_internal[0])) = int32(0)
	v63 = F_strtof(m, v15, v11+int32(-4))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float32(0)
L13:
	;
	if v37 == int32(0) {
		v569 = v9
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v11+int32(-16))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(249), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v569 = v9
	goto L1
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_float4in_internal[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v66|base.B2i32(v67 == v15) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v518 = v67
	v520 = v63
	goto L2
L20:
	;
	goto L21
L21:
	;
	v72 = int32(3)
	v77 = v15
	v78 = int32(_a_F_float4in_internal_3)
	v79 = v72
	goto L23
L22:
	;
	if v124 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v124 = int32(0)
	goto L22
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 == v83 {
		v105 = v82
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v107 = int32(1)
	if v105 != 0 {
		v77 = v77 + v107
		v78 = v78 + v107
		v79 = v79 - v107
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v82-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v82 | int32(32)
	goto L32
L31:
	;
	v93 = v82
	goto L32
L32:
	;
	if base.Ui32((v83-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v83 | int32(32)
	goto L35
L34:
	;
	v102 = v83
	goto L35
L35:
	;
	if v93 == v102 {
		v105 = v93
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v124 = v93 - v102
	goto L22
L37:
	;
	goto L27
L38:
	;
	v514 = v72
	v515 = math.Float32frombits(uint32(0x7fc00000))
	goto L8
L39:
	;
	goto L40
L40:
	;
	v128 = math.Float32frombits(uint32(0x7f800000))
	v129 = int32(8)
	v134 = v15
	v135 = int32(_a_F_float4in_internal_4)
	v136 = v129
	goto L42
L41:
	;
	if v181 == int32(0) {
		v514 = v129
		v515 = v128
		goto L8
	} else {
		goto L57
	}
L42:
	;
	if v136 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v181 = int32(0)
	goto L41
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 == v140 {
		v162 = v139
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v164 = int32(1)
	if v162 != 0 {
		v134 = v134 + v164
		v135 = v135 + v164
		v136 = v136 - v164
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v139-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v150 = v139 | int32(32)
	goto L51
L50:
	;
	v150 = v139
	goto L51
L51:
	;
	if base.Ui32((v140-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = v140 | int32(32)
	goto L54
L53:
	;
	v159 = v140
	goto L54
L54:
	;
	if v150 == v159 {
		v162 = v150
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v181 = v150 - v159
	goto L41
L56:
	;
	goto L46
L57:
	;
	v184 = int32(9)
	v189 = v15
	v190 = int32(_a_F_float4in_internal_5)
	v191 = v184
	goto L59
L58:
	;
	if v236 == int32(0) {
		v514 = v184
		v515 = v128
		goto L8
	} else {
		goto L74
	}
L59:
	;
	if v191 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v236 = int32(0)
	goto L58
L61:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		v217 = v194
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v219 = int32(1)
	if v217 != 0 {
		v189 = v189 + v219
		v190 = v190 + v219
		v191 = v191 - v219
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v194-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v205 = v194 | int32(32)
	goto L68
L67:
	;
	v205 = v194
	goto L68
L68:
	;
	if base.Ui32((v195-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v214 = v195 | int32(32)
	goto L71
L70:
	;
	v214 = v195
	goto L71
L71:
	;
	if v205 == v214 {
		v217 = v205
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v236 = v205 - v214
	goto L58
L73:
	;
	goto L63
L74:
	;
	v243 = v15
	v244 = int32(_a_F_float4in_internal_6)
	v245 = int32(9)
	goto L76
L75:
	;
	if v290 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v245 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v290 = int32(0)
	goto L75
L78:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v248 == v249 {
		v271 = v248
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v273 = int32(1)
	if v271 != 0 {
		v243 = v243 + v273
		v244 = v244 + v273
		v245 = v245 - v273
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v248-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v259 = v248 | int32(32)
	goto L85
L84:
	;
	v259 = v248
	goto L85
L85:
	;
	if base.Ui32((v249-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v268 = v249 | int32(32)
	goto L88
L87:
	;
	v268 = v249
	goto L88
L88:
	;
	if v259 == v268 {
		v271 = v259
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v290 = v259 - v268
	goto L75
L90:
	;
	goto L80
L91:
	;
	v514 = v184
	v515 = math.Float32frombits(uint32(0xff800000))
	goto L8
L92:
	;
	goto L93
L93:
	;
	v294 = int32(3)
	v299 = v15
	v300 = int32(_a_F_float4in_internal_7)
	v301 = v294
	goto L95
L94:
	;
	if v346 == int32(0) {
		v514 = v294
		v515 = v128
		goto L8
	} else {
		goto L110
	}
L95:
	;
	if v301 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v346 = int32(0)
	goto L94
L97:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v304 == v305 {
		v327 = v304
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v329 = int32(1)
	if v327 != 0 {
		v299 = v299 + v329
		v300 = v300 + v329
		v301 = v301 - v329
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v304-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v315 = v304 | int32(32)
	goto L104
L103:
	;
	v315 = v304
	goto L104
L104:
	;
	if base.Ui32((v305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v324 = v305 | int32(32)
	goto L107
L106:
	;
	v324 = v305
	goto L107
L107:
	;
	if v315 == v324 {
		v327 = v315
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v346 = v315 - v324
	goto L94
L109:
	;
	goto L99
L110:
	;
	v349 = int32(4)
	v354 = v15
	v355 = int32(_a_F_float4in_internal_8)
	v356 = v349
	goto L112
L111:
	;
	if v401 == int32(0) {
		v514 = v349
		v515 = v128
		goto L8
	} else {
		goto L127
	}
L112:
	;
	if v356 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v401 = int32(0)
	goto L111
L114:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	if v359 == v360 {
		v382 = v359
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v384 = int32(1)
	if v382 != 0 {
		v354 = v354 + v384
		v355 = v355 + v384
		v356 = v356 - v384
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v359-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v370 = v359 | int32(32)
	goto L121
L120:
	;
	v370 = v359
	goto L121
L121:
	;
	if base.Ui32((v360-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v379 = v360 | int32(32)
	goto L124
L123:
	;
	v379 = v360
	goto L124
L124:
	;
	if v370 == v379 {
		v382 = v370
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v401 = v370 - v379
	goto L111
L126:
	;
	goto L116
L127:
	;
	v409 = v15
	v410 = int32(_a_F_float4in_internal_9)
	v411 = int32(4)
	goto L129
L128:
	;
	if v456 == int32(0) {
		v514 = v349
		v515 = math.Float32frombits(uint32(0xff800000))
		goto L8
	} else {
		goto L144
	}
L129:
	;
	if v411 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v456 = int32(0)
	goto L128
L131:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v414 == v415 {
		v437 = v414
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v439 = int32(1)
	if v437 != 0 {
		v409 = v409 + v439
		v410 = v410 + v439
		v411 = v411 - v439
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v414-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v425 = v414 | int32(32)
	goto L138
L137:
	;
	v425 = v414
	goto L138
L138:
	;
	if base.Ui32((v415-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v434 = v415 | int32(32)
	goto L141
L140:
	;
	v434 = v415
	goto L141
L141:
	;
	if v425 == v434 {
		v437 = v425
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v456 = v425 - v434
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v66 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v463 = base.I32_reinterpret_f32(v63) & int32(2147483647)
	if base.B2i32(v463 != int32(0))&base.B2i32(v463 != int32(2139095040)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v494 = float32(0)
	v495 = F_errsave_start(m, l3)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v518 = v67
	v520 = v63
	goto L2
L149:
	;
	goto L150
L150:
	;
	v469 = F_pstrdup(m, v15)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v469+(v67-v15)))) = uint8(v473)
	v475 = float32(0)
	v476 = F_errsave_start(m, l3)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v476 == int32(0) {
		v569 = v475
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v469
	F_errmsg(m, int32(_a_F_float4in_internal_10), v11+int32(-48))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(329), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v569 = v475
	goto L1
L157:
	;
	if v495 == int32(0) {
		v569 = v494
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v11+int32(-32))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(336), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v569 = v494
	goto L1
L162:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if base.B2i32(base.Ui32(v531-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v531 == int32(32)) != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v543 = float32(0)
	v544 = F_errsave_start(m, l3)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L12
	} else {
		goto L170
	}
L164:
	;
	v526 = v526 + int32(1)
	goto L162
L165:
	;
	if v531 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L163
L167:
	;
	v569 = v520
	goto L1
L168:
	;
	goto L169
L169:
	;
	goto L166
L170:
	;
	if v544 == int32(0) {
		v569 = v543
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v13)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(350), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L174
	}
L174:
	;
	v569 = v543
	goto L1
}
func F_float4up(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	return v2
}
func F_float84gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float64
	_ = v11
	var v22 int64
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v22 = base.I64_extend_i32_u(base.F64_lt(v5, v11) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807))))
	} else {
		v22 = int64(0)
	}
	return v22
}
func F_float8div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v11 float64
	_ = v11
	var v18 float64
	_ = v18
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v49 float64
	_ = v49
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)))|base.F64_ne(v11, float64(0)) == int32(0) {
		v18 = F_float_zero_divide_error_ext(m, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v49 = float64(0)
			return base.I64_reinterpret_f64(v49)
		}
	} else {
		v23 = math.Float64frombits(uint64(0x7ff0000000000000))
		v25 = base.F64_div(v5, v11)
		if base.F64_eq(base.F64_abs(v5), v23)|base.F64_ne(base.F64_abs(v25), v23) == int32(0) {
			v33 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				v49 = float64(0)
				return base.I64_reinterpret_f64(v49)
			}
		} else {
			v35 = float64(0)
			if base.F64_eq(v5, v35)|base.F64_ne(v25, v35)|base.F64_eq(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v49 = v25
				return base.I64_reinterpret_f64(v49)
			} else {
				v45 = F_float_underflow_error_ext(m, int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					v49 = float64(0)
					return base.I64_reinterpret_f64(v49)
				}
			}
		}
	}
}
func F_float8gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float64
	_ = v10
	var v21 int64
	_ = v21
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v21 = base.I64_extend_i32_u(base.F64_lt(v4, v10) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807))))
	} else {
		v21 = int64(0)
	}
	return v21
}
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v132 float64
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v467 int64
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 float64
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 float64
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 float64
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 float64
	_ = v524
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v553 float64
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v580 float64
	_ = v580
	v10 = float64(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = l0
	goto L3
L1:
	;
	m.G0 = v15 - int32(-64)
	return v580
L2:
	;
	v532 = v522
	goto L164
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(base.Ui32(v29-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v29 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v520 = v17 + v518
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v520
	v522 = v520
	v524 = v519
	goto L2
L5:
	;
	v17 = v17 + int32(1)
	goto L3
L6:
	;
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	goto L7
L9:
	;
	v41 = F_errsave_start(m, l4)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_float8in_internal[0])) = int32(0)
	v67 = F_strtod(m, v17, v13+int32(-4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float64(0)
L13:
	;
	if v41 == int32(0) {
		v580 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v13+int32(-16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(455), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v580 = v10
	goto L1
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_float8in_internal[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v70|base.B2i32(v71 == v17) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v522 = v71
	v524 = v67
	goto L2
L20:
	;
	goto L21
L21:
	;
	v76 = int32(3)
	v81 = v17
	v82 = int32(_a_F_float8in_internal_3)
	v83 = v76
	goto L23
L22:
	;
	if v128 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v128 = int32(0)
	goto L22
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == v87 {
		v109 = v86
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v111 = int32(1)
	if v109 != 0 {
		v81 = v81 + v111
		v82 = v82 + v111
		v83 = v83 - v111
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v86-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = v86 | int32(32)
	goto L32
L31:
	;
	v97 = v86
	goto L32
L32:
	;
	if base.Ui32((v87-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = v87 | int32(32)
	goto L35
L34:
	;
	v106 = v87
	goto L35
L35:
	;
	if v97 == v106 {
		v109 = v97
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v128 = v97 - v106
	goto L22
L37:
	;
	goto L27
L38:
	;
	v518 = v76
	v519 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L8
L39:
	;
	goto L40
L40:
	;
	v132 = math.Float64frombits(uint64(0x7ff0000000000000))
	v133 = int32(8)
	v138 = v17
	v139 = int32(_a_F_float8in_internal_4)
	v140 = v133
	goto L42
L41:
	;
	if v185 == int32(0) {
		v518 = v133
		v519 = v132
		goto L8
	} else {
		goto L57
	}
L42:
	;
	if v140 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v185 = int32(0)
	goto L41
L44:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		v166 = v143
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v168 = int32(1)
	if v166 != 0 {
		v138 = v138 + v168
		v139 = v139 + v168
		v140 = v140 - v168
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v143-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v154 = v143 | int32(32)
	goto L51
L50:
	;
	v154 = v143
	goto L51
L51:
	;
	if base.Ui32((v144-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v163 = v144 | int32(32)
	goto L54
L53:
	;
	v163 = v144
	goto L54
L54:
	;
	if v154 == v163 {
		v166 = v154
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v185 = v154 - v163
	goto L41
L56:
	;
	goto L46
L57:
	;
	v188 = int32(9)
	v193 = v17
	v194 = int32(_a_F_float8in_internal_5)
	v195 = v188
	goto L59
L58:
	;
	if v240 == int32(0) {
		v518 = v188
		v519 = v132
		goto L8
	} else {
		goto L74
	}
L59:
	;
	if v195 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v240 = int32(0)
	goto L58
L61:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v198 == v199 {
		v221 = v198
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v223 = int32(1)
	if v221 != 0 {
		v193 = v193 + v223
		v194 = v194 + v223
		v195 = v195 - v223
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v198-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v209 = v198 | int32(32)
	goto L68
L67:
	;
	v209 = v198
	goto L68
L68:
	;
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v218 = v199 | int32(32)
	goto L71
L70:
	;
	v218 = v199
	goto L71
L71:
	;
	if v209 == v218 {
		v221 = v209
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v240 = v209 - v218
	goto L58
L73:
	;
	goto L63
L74:
	;
	v247 = v17
	v248 = int32(_a_F_float8in_internal_6)
	v249 = int32(9)
	goto L76
L75:
	;
	if v294 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v249 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v294 = int32(0)
	goto L75
L78:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v252 == v253 {
		v275 = v252
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v277 = int32(1)
	if v275 != 0 {
		v247 = v247 + v277
		v248 = v248 + v277
		v249 = v249 - v277
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v263 = v252 | int32(32)
	goto L85
L84:
	;
	v263 = v252
	goto L85
L85:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v272 = v253 | int32(32)
	goto L88
L87:
	;
	v272 = v253
	goto L88
L88:
	;
	if v263 == v272 {
		v275 = v263
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v294 = v263 - v272
	goto L75
L90:
	;
	goto L80
L91:
	;
	v518 = v188
	v519 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L8
L92:
	;
	goto L93
L93:
	;
	v298 = int32(3)
	v303 = v17
	v304 = int32(_a_F_float8in_internal_7)
	v305 = v298
	goto L95
L94:
	;
	if v350 == int32(0) {
		v518 = v298
		v519 = v132
		goto L8
	} else {
		goto L110
	}
L95:
	;
	if v305 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v350 = int32(0)
	goto L94
L97:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v308 == v309 {
		v331 = v308
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v333 = int32(1)
	if v331 != 0 {
		v303 = v303 + v333
		v304 = v304 + v333
		v305 = v305 - v333
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v308-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v319 = v308 | int32(32)
	goto L104
L103:
	;
	v319 = v308
	goto L104
L104:
	;
	if base.Ui32((v309-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v328 = v309 | int32(32)
	goto L107
L106:
	;
	v328 = v309
	goto L107
L107:
	;
	if v319 == v328 {
		v331 = v319
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v350 = v319 - v328
	goto L94
L109:
	;
	goto L99
L110:
	;
	v353 = int32(4)
	v358 = v17
	v359 = int32(_a_F_float8in_internal_8)
	v360 = v353
	goto L112
L111:
	;
	if v405 == int32(0) {
		v518 = v353
		v519 = v132
		goto L8
	} else {
		goto L127
	}
L112:
	;
	if v360 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v405 = int32(0)
	goto L111
L114:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	if v363 == v364 {
		v386 = v363
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v388 = int32(1)
	if v386 != 0 {
		v358 = v358 + v388
		v359 = v359 + v388
		v360 = v360 - v388
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v363-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v374 = v363 | int32(32)
	goto L121
L120:
	;
	v374 = v363
	goto L121
L121:
	;
	if base.Ui32((v364-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v383 = v364 | int32(32)
	goto L124
L123:
	;
	v383 = v364
	goto L124
L124:
	;
	if v374 == v383 {
		v386 = v374
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v405 = v374 - v383
	goto L111
L126:
	;
	goto L116
L127:
	;
	v413 = v17
	v414 = int32(_a_F_float8in_internal_9)
	v415 = int32(4)
	goto L129
L128:
	;
	if v460 == int32(0) {
		v518 = v353
		v519 = math.Float64frombits(uint64(0xfff0000000000000))
		goto L8
	} else {
		goto L144
	}
L129:
	;
	if v415 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v460 = int32(0)
	goto L128
L131:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v418 == v419 {
		v441 = v418
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v443 = int32(1)
	if v441 != 0 {
		v413 = v413 + v443
		v414 = v414 + v443
		v415 = v415 - v443
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v418-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v429 = v418 | int32(32)
	goto L138
L137:
	;
	v429 = v418
	goto L138
L138:
	;
	if base.Ui32((v419-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v438 = v419 | int32(32)
	goto L141
L140:
	;
	v438 = v419
	goto L141
L141:
	;
	if v429 == v438 {
		v441 = v429
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v460 = v429 - v438
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v70 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v467 = base.I64_reinterpret_f64(v67) & int64(9223372036854775807)
	if base.B2i32(v467 != int64(0))&base.B2i32(v467 != int64(9218868437227405312)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v498 = float64(0)
	v499 = F_errsave_start(m, l4)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v522 = v71
	v524 = v67
	goto L2
L149:
	;
	goto L150
L150:
	;
	v473 = F_pstrdup(m, v17)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	v477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v473+(v71-v17)))) = uint8(v477)
	v479 = float64(0)
	v480 = F_errsave_start(m, l4)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v480 == int32(0) {
		v580 = v479
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v473
	F_errmsg(m, int32(_a_F_float8in_internal_10), v13+int32(-48))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(531), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v580 = v479
	goto L1
L157:
	;
	if v499 == int32(0) {
		v580 = v498
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v13+int32(-32))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(538), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v580 = v498
	goto L1
L162:
	;
	v553 = float64(0)
	v554 = F_errsave_start(m, l4)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L174
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v532
	v580 = v524
	goto L1
L164:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	if base.B2i32(base.Ui32(v538-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v538 == int32(32)) != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if l1 == int32(0) {
		goto L162
	} else {
		goto L173
	}
L166:
	;
	v532 = v532 + int32(1)
	goto L164
L167:
	;
	if v538 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L165
L169:
	;
	if l1 != 0 {
		goto L163
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	goto L168
L172:
	;
	v580 = v524
	goto L1
L173:
	;
	goto L163
L174:
	;
	if v554 == int32(0) {
		v580 = v553
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v15)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L177
	}
L177:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(552), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L12
	} else {
		goto L178
	}
L178:
	;
	v580 = v553
	goto L1
}
func F_float8le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float64
	_ = v10
	var v21 int64
	_ = v21
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v21 = base.I64_extend_i32_u(base.F64_ge(v4, v10) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))))
	} else {
		v21 = int64(1)
	}
	return v21
}
func F_flt4_mul_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 float32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_cash_mul_float8(m, v2, base.F64_promote_f32(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_flt8_mul_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 float64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_cash_mul_float8(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_fmt_u(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
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
	if base.Ui64(int64(4294967296)) <= base.Ui64(l0) {
		v8 = l0
		v9 = l1
		for {
			v14 = v9 - int32(1)
			v15 = int64(10)
			v16 = base.I64_div_u_s(v8, v15)
			v22 = base.I32_wrap_i64(v8-v16*v15) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v22)
			if base.Ui64(int64(42949672959)) < base.Ui64(v8) {
				v8 = v16
				v9 = v14
				continue
			} else {
				break
			}
			break
		}
		v26 = v16
		v27 = v14
	} else {
		v26 = l0
		v27 = l1
	}
	v31 = base.I32_wrap_i64(v26)
	if base.Ui64(int64(10)) <= base.Ui64(v26) {
		v35 = v27
		v36 = v31
		for {
			v40 = v35 - int32(1)
			v41 = int32(10)
			v42 = base.I32_div_u_s(v36, v41)
			v47 = v36 - v42*v41 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v47)
			if base.Ui32(int32(99)) < base.Ui32(v36) {
				v35 = v40
				v36 = v42
				continue
			} else {
				break
			}
			break
		}
		v52 = v40
		v53 = v42
	} else {
		v52 = v27
		v53 = v31
	}
	if v53 != 0 {
		v57 = v52 - int32(1)
		v59 = v53 | int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v59)
		v61 = v57
	} else {
		v61 = v52
	}
	return v61
}
func F_forkname_chars(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v5 = int32(3)
	v6 = int32(_a_F_forkname_chars_0)
	goto L6
L1:
	;
	return v169
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v167
	v169 = v166
	goto L1
L3:
	;
	if l1 == int32(0) {
		v169 = v162
		goto L1
	} else {
		goto L52
	}
L4:
	;
	if v44-v45 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	goto L7
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[0])))
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v14 = v6
	v15 = l0
	v16 = v5
	v17 = v13
	goto L12
L9:
	;
	v40 = l0
	v44 = int32(0)
	goto L10
L10:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	goto L4
L11:
	;
	v40 = v35
	v44 = v37
	goto L10
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(v17 != v19)|base.B2i32(v19 == int32(0)) != 0 {
		v35 = v15
		v37 = v17
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v35 = v29
	v37 = int32(0)
	goto L11
L14:
	;
	v25 = v16 - int32(1)
	if v25 == int32(0) {
		v35 = v15
		v37 = v17
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v28 = int32(1)
	v29 = v15 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v30 != 0 {
		v14 = v14 + v28
		v15 = v29
		v16 = v25
		v17 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v162 = v5
	v163 = int32(1)
	goto L3
L18:
	;
	goto L19
L19:
	;
	v56 = int32(2)
	v58 = int32(_a_F_forkname_chars_1)
	goto L22
L20:
	;
	if v96-v97 == int32(0) {
		v162 = v56
		v163 = v56
		goto L3
	} else {
		goto L33
	}
L22:
	;
	goto L23
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[1])))
	if v65 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v58
	v67 = l0
	v68 = v56
	v69 = v65
	goto L28
L25:
	;
	v92 = l0
	v96 = int32(0)
	goto L26
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	goto L20
L27:
	;
	v92 = v87
	v96 = v89
	goto L26
L28:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if base.B2i32(v69 != v71)|base.B2i32(v71 == int32(0)) != 0 {
		v87 = v67
		v89 = v69
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v87 = v81
	v89 = int32(0)
	goto L27
L30:
	;
	v77 = v68 - int32(1)
	if v77 == int32(0) {
		v87 = v67
		v89 = v69
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v80 = int32(1)
	v81 = v67 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v82 != 0 {
		v66 = v66 + v80
		v67 = v81
		v68 = v77
		v69 = v82
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v107 = int32(4)
	v108 = int32(_a_F_forkname_chars_2)
	goto L36
L34:
	;
	if v146-v147 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	goto L37
L37:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[2])))
	if v115 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = v108
	v117 = l0
	v118 = v107
	v119 = v115
	goto L42
L39:
	;
	v142 = l0
	v146 = int32(0)
	goto L40
L40:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	goto L34
L41:
	;
	v142 = v137
	v146 = v139
	goto L40
L42:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v119 != v121)|base.B2i32(v121 == int32(0)) != 0 {
		v137 = v117
		v139 = v119
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v137 = v131
	v139 = int32(0)
	goto L41
L44:
	;
	v127 = v118 - int32(1)
	if v127 == int32(0) {
		v137 = v117
		v139 = v119
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v130 = int32(1)
	v131 = v117 + v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v132 != 0 {
		v116 = v116 + v130
		v117 = v131
		v118 = v127
		v119 = v132
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if l1 != 0 {
		v166 = v107
		v167 = int32(3)
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = int32(0)
	if l1 == v159 {
		v169 = v159
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v169 = v107
	goto L1
L51:
	;
	v166 = v159
	v167 = int32(-1)
	goto L2
L52:
	;
	v166 = v162
	v167 = v163
	goto L2
}
func F_free_parsestate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v8-int32(1) < int32(1665) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v13 != 0 {
			F_relation_close(m, v13, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(17039621))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1664)
				F_errmsg(m, int32(_a_F_free_parsestate_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_free_parsestate_1), int32(83), int32(_a_F_free_parsestate_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
func F_free_struct_lconv(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_emscripten_builtin_free(m, v2)
	mBase = m.M
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_emscripten_builtin_free(m, v4)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_emscripten_builtin_free(m, v6)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_emscripten_builtin_free(m, v8)
	mBase = m.M
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_emscripten_builtin_free(m, v10)
	mBase = m.M
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_emscripten_builtin_free(m, v12)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_emscripten_builtin_free(m, v14)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_emscripten_builtin_free(m, v16)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_emscripten_builtin_free(m, v20)
	mBase = m.M
	return
}
func F_freeifaddrs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l0 != 0 {
		v3 = l0
		for {
			v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			F_emscripten_builtin_free(m, v3)
			mBase = m.M
			if v5 != 0 {
				v3 = v5
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_frexp(m *base.Module, l0 float64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 float64
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v70 float64
	_ = v70
	v5 = base.I64_reinterpret_f64(l0)
	v9 = int32(2047)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & v9
	if v10 != v9 {
		if v10 == int32(0) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v58 = l0
				v59 = int32(0)
			} else {
				v19 = base.F64_mul(l0, float64(1.8446744073709552e+19))
				v22 = base.I64_reinterpret_f64(v19)
				v26 = int32(2047)
				v27 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(52))%64))) & v26
				if v27 != v26 {
					if v27 == int32(0) {
						if base.F64_eq(v19, float64(0)) != 0 {
							v41 = v19
							v42 = int32(0)
						} else {
							v37 = F_frexp(m, base.F64_mul(v19, float64(1.8446744073709552e+19)), l1)
							mBase = m.M
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v41 = v37
							v42 = v38 + int32(-64)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
						v54 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27 - int32(1022)
						v52 = base.F64_reinterpret_i64(v22&int64(-9218868437227405313) | int64(4602678819172646912))
						v54 = v52
					}
				} else {
					v52 = v19
					v54 = v52
				}
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v58 = v54
				v59 = v55 + int32(-64)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
			return v58
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10 - int32(1022)
			v70 = base.F64_reinterpret_i64(v5&int64(-9218868437227405313) | int64(4602678819172646912))
			return v70
		}
	} else {
		v70 = l0
		return v70
	}
}
func F_fscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_vfscanf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_fsm_get_avail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1)+uint32(_c_F_fsm_get_avail[0]))))
	return v6
}
func F_fsm_search_avail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v274
L2:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)))
	if base.Ui32(v34) < base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+(l0^int32(-1))<<(uint(int32(2))%32))))
	v33 = v25
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[1]))
	v33 = v27 + l0<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v274 = int32(-1)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v38 = v33 + int32(28)
	v44 = l3
	goto L9
L9:
	;
	v52 = int32(4095)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v55 = v53 + v52
	if base.Ui32(int32(4068)) < base.Ui32(v53) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v274 = int32(-1)
	goto L1
L11:
	;
	if v93 <= int32(4094) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v58 = v52
	goto L14
L13:
	;
	v58 = v55
	goto L14
L14:
	;
	if v58 <= int32(0) {
		v93 = v55
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v66 = v58
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v38))))
	if base.Ui32(l1) <= base.Ui32(v73) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v93 = v86
	goto L11
L18:
	;
	v93 = v66
	goto L11
L19:
	;
	goto L20
L20:
	;
	v75 = int32(1)
	if (v66+int32(2))&(v66+v75) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v66
	goto L23
L22:
	;
	v84 = int32(base.Ui32(v66)>>(uint(v75)%32)) - v75
	goto L23
L23:
	;
	v86 = base.I32_div_s(v84, int32(2))
	if int32(1) < v84 {
		v66 = v86
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	F_MarkBufferDirtyHint(m, l0, int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L45
	} else {
		goto L76
	}
L26:
	;
	v106 = v93
	goto L29
L27:
	;
	v234 = v93
	goto L28
L28:
	;
	v242 = v234 - int32(4095)
	if v44&int32(1) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L29:
	;
	v114 = v106 << (uint(int32(1)) % 32)
	if base.Ui32(v114) <= base.Ui32(int32(_a_F_fsm_search_avail_0)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v234 = v227
	goto L28
L31:
	;
	if v227 < int32(4095) {
		v106 = v227
		goto L29
	} else {
		goto L69
	}
L32:
	;
	v118 = v114 | int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v118))))
	if base.Ui32(l1) <= base.Ui32(v120) {
		v227 = v118
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v124 = v114 + int32(2)
	if base.Ui32(v124) <= base.Ui32(int32(_a_F_fsm_search_avail_0)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v38))))
	if base.Ui32(l1) <= base.Ui32(v128) {
		v227 = v124
		goto L31
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v131 = v14 + int32(36)
	if l0 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v164 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v155
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v154
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28)))) = v160
	goto L40
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[2]))
	v153 = v140 + (l0^int32(-1))*int32(56)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search_avail[3]))
	v148 = int32(56)
	v153 = v147 + l0*v148 - v148
	goto L41
L45:
	;
	return int32(0)
L46:
	;
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v172
	F_errmsg_internal(m, int32(_a_F_fsm_search_avail_1), v14)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v44&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_errfinish(m, int32(_a_F_fsm_search_avail_2), int32(277), int32(_a_F_fsm_search_avail_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	F_UnlockBuffer(m, l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L45
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v196 = int32(4094)
	goto L57
L55:
	;
	F_LockBufferInternal(m, l0, int32(3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v196) {
		v218 = int32(0)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L25
L59:
	;
	v219 = v196 + v38
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v220 != v218&int32(255) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v207 = v196 << (uint(int32(1)) % 32)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(29)+v207))))
	if v196 == int32(4081) {
		v218 = v209
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v38)+2)))
	if base.Ui32(v213) < base.Ui32(v209) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v215 = v209
	goto L64
L63:
	;
	v215 = v213
	goto L64
L64:
	;
	v218 = v215
	goto L59
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v218)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v196 != 0 {
		v196 = v196 - int32(1)
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L58
L69:
	;
	goto L30
L70:
	;
	v248 = v242 & int32(_a_F_fsm_search_avail_4)
	v249 = F_BufferBeginSetHintBits(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L45
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v260 = v242 & int32(_a_F_fsm_search_avail_4)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v260 + l2
	v274 = v260
	goto L1
L73:
	;
	if v249 == int32(0) {
		v274 = v248
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = l2 + v248
	v255 = int32(0)
	F_BufferFinishSetHintBits(m, l0, v255, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L45
	} else {
		goto L75
	}
L75:
	;
	v274 = v248
	goto L1
L76:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(l1) <= base.Ui32(v267) {
		v44 = int32(1)
		goto L9
	} else {
		goto L77
	}
L77:
	;
	goto L10
}
func F_fsync_fname(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fsync_fname[0])))
	if v7 != 0 {
		v8 = int32(21)
	} else {
		v8 = int32(24)
	}
	v9 = F_fsync_fname_ext(m, l0, l1, int32(0), v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
