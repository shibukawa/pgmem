package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_evict_all(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v109 int64
	_ = v109
	var v118 int64
	_ = v118
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+28)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v2
	v27 = F_get_call_result_type(m, l0, v2, v10+int32(-4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L57
	}
L2:
	;
	return int64(0)
L3:
	;
	if v27 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = F_superuser(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L54
	}
L7:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v37 = m.G0
	v39 = v37 - int32(32)
	m.G0 = v39
	v42 = v10 + int32(-40)
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	v46 = v10 + int32(-48)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v43
	v50 = v10 + int32(-44)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[0]))
	if v43 < v54 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = int32(1)
	goto L12
L10:
	;
	goto L11
L11:
	;
	m.G0 = v39 + int32(32)
	v235 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+24)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v235
	v237 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v237
	v239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v246 = F_heap_form_tuple(m, v241, v10+int32(-32), v10+int32(-36))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L52
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[1]))
	v71 = v68 + v62*int32(56)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[2]))
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v77 = v71 - int32(32)
	v78 = int64(0)
	v81 = base.AtomicRmwCmpxchg64(m, v77, int32(0), v78, v78)
	if v81&int64(16777216) == v78 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v219 = v62 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[0]))
	if v219 <= v221 {
		v62 = v219
		goto L12
	} else {
		goto L51
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[3]))
	F_ResourceOwnerEnlarge(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v94 = int64(4194304)
	v96 = base.AtomicRmwOr64(m, v77, int32(0), v94)
	if v96&v94 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v109 = v96
	goto L25
L23:
	;
	goto L24
L24:
	;
	v195 = F_EvictUnpinnedBufferInternal(m, v71-int32(56), v39+int32(8))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L46
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = int32(_a_F_pg_buffercache_evict_all_0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(_a_F_pg_buffercache_evict_all_1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(_a_F_pg_buffercache_evict_all_2)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
	v118 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v118
	if v109&int64(4194304) != v118 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	goto L30
L28:
	;
	goto L29
L29:
	;
	v157 = int32(_a_F_pg_buffercache_evict_all_3)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[4]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(8))+8))
	if v160 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	F_perform_spin_delay(m, v39+int32(8))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v137 = int64(0)
	v140 = base.AtomicRmwCmpxchg64(m, v77, int32(0), v137, v137)
	if v140&int64(4194304) != v137 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v177 = int64(4194304)
	v179 = base.AtomicRmwOr64(m, v77, int32(0), v177)
	if v179&v177 != int64(0) {
		v109 = v179
		goto L25
	} else {
		goto L45
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[4])) = v175
	goto L35
L37:
	;
	if int32(999) < v158 {
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v158 < int32(11) {
		goto L35
	} else {
		goto L44
	}
L40:
	;
	v165 = int32(900)
	if v165 <= v158 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = v165
	goto L43
L42:
	;
	v168 = v158
	goto L43
L43:
	;
	v175 = v168 + int32(100)
	goto L36
L44:
	;
	v175 = v158 - int32(1)
	goto L36
L45:
	;
	goto L26
L46:
	;
	if v195 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v197 = v42
	goto L49
L48:
	;
	v197 = v46
	goto L49
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v199 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v198 + v199
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v202 != v199 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v205 + int32(1)
	goto L18
L51:
	;
	goto L13
L52:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v249 = F_HeapTupleHeaderGetDatum(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v12 - int32(-64)
	return v249
L54:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_all_4), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_all_5), int32(793), int32(_a_F_pg_buffercache_evict_all_6))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_buffercache_evict_all_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_all_7), v12)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_all_5), int32(691), int32(_a_F_pg_buffercache_evict_all_8))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_os_pages(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_buffercache_os_pages_internal(m, l0, base.B2i32(v2 != int64(0)))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
