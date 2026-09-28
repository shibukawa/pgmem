package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_mark_dirty_relation(m *base.Module, l0 int32) int64 {
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
	var v73 int32
	_ = v73
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
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
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
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L72
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
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L69
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
	v53 = v15 + int32(36)
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54
	v57 = v15 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v54
	v61 = v15 + v49
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v54
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[0]))
	if v54 < v65 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v73 = int32(1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	m.G0 = v50 + int32(32)
	F_relation_close(m, v42, int32(1))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L66
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[1]))
	v85 = v82 + v73*int32(56)
	v87 = v85 - int32(32)
	v88 = int64(0)
	v91 = base.AtomicRmwCmpxchg64(m, v87, int32(0), v88, v88)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[2]))
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
	v280 = v73 + int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[0]))
	if v280 <= v282 {
		v73 = v280
		goto L15
	} else {
		goto L65
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
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[3]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(_a_F_pg_buffercache_mark_dirty_relation_0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = int32(_a_F_pg_buffercache_mark_dirty_relation_1)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = int32(_a_F_pg_buffercache_mark_dirty_relation_2)
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
	v193 = int32(_a_F_pg_buffercache_mark_dirty_relation_3)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[4]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_relation[4])) = v211
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
	v250 = F_MarkDirtyUnpinnedBufferInternal(m, v73, v101, v50+int32(8))
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
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v252 + int32(1)
	goto L21
L60:
	;
	goto L61
L61:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+8)))
	if v256 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v259 + int32(1)
	goto L21
L63:
	;
	goto L64
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v263 + int32(1)
	goto L21
L65:
	;
	goto L16
L66:
	;
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+36)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v302
	v304 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v304
	v306 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v313 = F_heap_form_tuple(m, v308, v15+int32(48), v15+int32(44))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v313)+16))
	v316 = F_HeapTupleHeaderGetDatum(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	m.G0 = v15 + int32(80)
	return v316
L69:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_mark_dirty_relation_4), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_relation_5), int32(863), int32(_a_F_pg_buffercache_mark_dirty_relation_6))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_pg_buffercache_mark_dirty_relation_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_mark_dirty_relation_7), v15+int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_relation_5), int32(691), int32(_a_F_pg_buffercache_mark_dirty_relation_8))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_pg_buffercache_mark_dirty_relation_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_mark_dirty_relation_9), v15)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_relation_5), int32(875), int32(_a_F_pg_buffercache_mark_dirty_relation_6))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
