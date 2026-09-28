package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_evict_relation(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v139 int64
	_ = v139
	var v148 int64
	_ = v148
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v231 int64
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+46)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+44)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v2
	v30 = F_get_call_result_type(m, l0, v2, v15+int32(76))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L74
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L70
	}
L3:
	;
	return int64(0)
L4:
	;
	if v30 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v36 = F_superuser(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L67
	}
L8:
	;
	if v36 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = F_relation_open(m, v40, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	if v45 == int32(116) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v48 = m.G0
	v49 = int32(32)
	v50 = v48 - v49
	m.G0 = v50
	v53 = v15 + v49
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54
	v57 = v15 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v54
	v61 = v15 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v54
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[0]))
	if v54 < v65 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v74 = int32(1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	m.G0 = v50 + int32(32)
	F_relation_close(m, v42, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L64
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[1]))
	v85 = v82 + v74*int32(56)
	v87 = v85 - int32(32)
	v88 = int64(0)
	v91 = base.AtomicRmwCmpxchg64(m, v87, int32(0), v88, v88)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[2]))
	if v93 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v91&int64(16777216) == int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v277 = v74 + int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[0]))
	if v277 <= v279 {
		v74 = v277
		goto L15
	} else {
		goto L63
	}
L22:
	;
	v101 = v85 - int32(56)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v102 != v103 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v106 = v85 - int32(52)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v107 != v108 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v111 = v85 - int32(48)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v112 != v113 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[3]))
	F_ResourceOwnerEnlarge(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v121 = int64(4194304)
	v123 = base.AtomicRmwOr64(m, v87, int32(0), v121)
	if v123&v121 != int64(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v139 = v123
	goto L31
L29:
	;
	v231 = v123
	goto L30
L30:
	;
	if v231&int64(16777216) == int64(0) {
		goto L53
	} else {
		goto L54
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(_a_F_pg_buffercache_evict_relation_0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = int32(_a_F_pg_buffercache_evict_relation_1)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = int32(_a_F_pg_buffercache_evict_relation_2)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v148
	if v139&int64(4194304) != v148 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v231 = v215
	goto L30
L33:
	;
	goto L36
L34:
	;
	goto L35
L35:
	;
	v193 = int32(_a_F_pg_buffercache_evict_relation_3)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[4]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(8))+8))
	if v196 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	F_perform_spin_delay(m, v50+int32(8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v170 = int64(0)
	v173 = base.AtomicRmwCmpxchg64(m, v87, int32(0), v170, v170)
	if v173&int64(4194304) != v170 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v213 = int64(4194304)
	v215 = base.AtomicRmwOr64(m, v87, int32(0), v213)
	if v215&v213 != int64(0) {
		v139 = v215
		goto L31
	} else {
		goto L51
	}
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[4])) = v211
	goto L41
L43:
	;
	if int32(999) < v194 {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v194 < int32(11) {
		goto L41
	} else {
		goto L50
	}
L46:
	;
	v201 = int32(900)
	if v201 <= v194 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v204 = v201
	goto L49
L48:
	;
	v204 = v194
	goto L49
L49:
	;
	v211 = v204 + int32(100)
	goto L42
L50:
	;
	v211 = v194 - int32(1)
	goto L42
L51:
	;
	goto L32
L52:
	;
	v250 = F_EvictUnpinnedBufferInternal(m, v101, v50+int32(8))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L58
	}
L53:
	;
	v247 = base.AtomicRmwSub64(m, v87, int32(0), int64(4194304))
	goto L21
L54:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v236 != v237 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v239 != v240 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v242 == v243 {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	if v250 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v252 = v57
	goto L61
L60:
	;
	v252 = v53
	goto L61
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v254 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v253 + v254
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+8)))
	if v257 != v254 {
		goto L21
	} else {
		goto L62
	}
L62:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v260 + int32(1)
	goto L21
L63:
	;
	goto L16
L64:
	;
	v299 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v299
	v301 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+36)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v301
	v303 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v310 = F_heap_form_tuple(m, v305, v15+int32(48), v15+int32(44))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310)+16))
	v313 = F_HeapTupleHeaderGetDatum(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	m.G0 = v15 + int32(80)
	return v313
L67:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_relation_4), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_5), int32(746), int32(_a_F_pg_buffercache_evict_relation_6))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_pg_buffercache_evict_relation_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_relation_7), v15+int32(16))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_5), int32(691), int32(_a_F_pg_buffercache_evict_relation_8))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_pg_buffercache_evict_relation_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_relation_9), v15)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_5), int32(758), int32(_a_F_pg_buffercache_evict_relation_6))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_summary(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 float64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	v2 = int32(0)
	v7 = int64(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v18 = F_get_call_result_type(m, l0, v2, v11+int32(-4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_s(v105)
	if v105 != 0 {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v43 = int32(0)
	v45 = v2
	v46 = v2
	v47 = v2
	v48 = v2
	v49 = v7
	goto L12
L3:
	;
	return int64(0)
L4:
	;
	if v18 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[0]))
	if int32(0) < v25 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v105 = v2
	v109 = v7
	v110 = v7
	v112 = float64(0)
	v113 = int64(0)
	goto L1
L9:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_summary_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_summary_1), int32(587), int32(_a_F_pg_buffercache_summary_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[1]))
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v105 = v85
	v109 = base.I64_extend_i32_s(v86)
	v110 = base.I64_extend_i32_s(v87)
	v112 = base.F64_convert_i64_s(v88)
	v113 = base.I64_extend_i32_u(v93)
	goto L1
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[2]))
	v62 = int64(0)
	v65 = base.AtomicRmwCmpxchg64(m, v58+v43*int32(56), int32(24), v62, v62)
	if v65&int64(16777216) != v62 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v93 = v48 + base.B2i32(v65&int64(262143) != int64(0))
	v95 = v43 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[0]))
	if v95 < v97 {
		v43 = v95
		v45 = v85
		v46 = v86
		v47 = v87
		v48 = v93
		v49 = v88
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v70 = int32(1)
	v85 = v45 + v70
	v86 = int32(base.Ui32(base.I32_wrap_i64(v65))>>(uint(int32(23))%32))&v70 + v46
	v87 = v47
	v88 = int64(base.Ui64(v65)>>(uint(int64(18))%64))&int64(15) + v49
	goto L18
L20:
	;
	goto L21
L21:
	;
	v85 = v45
	v86 = v46
	v87 = v47 + int32(1)
	v88 = v49
	goto L18
L22:
	;
	goto L13
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v133 = F_heap_form_tuple(m, v128, v11+int32(-48), v11+int32(-56))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L27
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = base.F64_div(v112, base.F64_convert_i32_s(v105))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v126)
	goto L23
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v136 = F_HeapTupleHeaderGetDatum(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v13 - int32(-64)
	return v136
}
