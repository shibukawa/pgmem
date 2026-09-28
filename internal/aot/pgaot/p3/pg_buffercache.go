package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_mark_dirty_all(m *base.Module, l0 int32) int64 {
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
	var v61 int32
	_ = v61
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
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
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
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L60
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
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L57
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
	v42 = v10 + int32(-44)
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	v46 = v10 + int32(-40)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v43
	v50 = v10 + int32(-48)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[0]))
	if v43 < v54 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v61 = int32(1)
	goto L12
L10:
	;
	goto L11
L11:
	;
	m.G0 = v39 + int32(32)
	v230 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v230
	v232 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+24)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v232
	v234 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v241 = F_heap_form_tuple(m, v236, v10+int32(-32), v10+int32(-36))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L55
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[1]))
	v71 = v68 + v61*int32(56)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[2]))
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
	if v81&int64(16777216) != v78 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[3]))
	F_ResourceOwnerEnlarge(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v214 = v61 + int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[0]))
	if v214 <= v216 {
		v61 = v214
		goto L12
	} else {
		goto L54
	}
L21:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v94 = int64(4194304)
	v96 = base.AtomicRmwOr64(m, v77, int32(0), v94)
	if v96&v94 != int64(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v109 = v96
	goto L26
L24:
	;
	goto L25
L25:
	;
	v195 = F_MarkDirtyUnpinnedBufferInternal(m, v61, v71-int32(56), v39+int32(8))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L47
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = int32(_a_F_pg_buffercache_mark_dirty_all_0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(_a_F_pg_buffercache_mark_dirty_all_1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(_a_F_pg_buffercache_mark_dirty_all_2)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
	v118 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v118
	if v109&int64(4194304) != v118 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	goto L31
L29:
	;
	goto L30
L30:
	;
	v157 = int32(_a_F_pg_buffercache_mark_dirty_all_3)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[4]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(8))+8))
	if v160 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v39+int32(8))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	v137 = int64(0)
	v140 = base.AtomicRmwCmpxchg64(m, v77, int32(0), v137, v137)
	if v140&int64(4194304) != v137 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v177 = int64(4194304)
	v179 = base.AtomicRmwOr64(m, v77, int32(0), v177)
	if v179&v177 != int64(0) {
		v109 = v179
		goto L26
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty_all[4])) = v175
	goto L36
L38:
	;
	if int32(999) < v158 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v158 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v165 = int32(900)
	if v165 <= v158 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v168 = v165
	goto L44
L43:
	;
	v168 = v158
	goto L44
L44:
	;
	v175 = v168 + int32(100)
	goto L37
L45:
	;
	v175 = v158 - int32(1)
	goto L37
L46:
	;
	goto L27
L47:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v197 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v198 = v46
	goto L50
L49:
	;
	v198 = v50
	goto L50
L50:
	;
	if v195 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v199 = v42
	goto L53
L52:
	;
	v199 = v198
	goto L53
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v200 + int32(1)
	goto L20
L54:
	;
	goto L13
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v244 = F_HeapTupleHeaderGetDatum(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	m.G0 = v12 - int32(-64)
	return v244
L57:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_mark_dirty_all_4), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_all_5), int32(909), int32(_a_F_pg_buffercache_mark_dirty_all_6))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_buffercache_mark_dirty_all_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_mark_dirty_all_7), v12)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_all_5), int32(691), int32(_a_F_pg_buffercache_mark_dirty_all_8))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_numa_pages(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pg_buffercache_os_pages_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pg_buffercache_pages(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v250 int64
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = F_get_call_result_type(m, l0, int32(0), v15+int32(92))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L61
	}
L2:
	;
	return int64(0)
L3:
	;
	if v21 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if base.Ui32(v28-int32(10)) <= base.Ui32(int32(-3)) {
		goto L1
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
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L58
	}
L7:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v36 = int32(_a_F_pg_buffercache_pages_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0])) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = F_CreateTemplateTupleDesc(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v44, int32(1), int32(_a_F_pg_buffercache_pages_1), int32(23), int32(-1), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v44, int32(2), int32(_a_F_pg_buffercache_pages_2), int32(26), int32(-1), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v44, int32(3), int32(_a_F_pg_buffercache_pages_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v44, int32(4), int32(_a_F_pg_buffercache_pages_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v44, int32(5), int32(_a_F_pg_buffercache_pages_5), int32(21), int32(-1), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v44, int32(6), int32(_a_F_pg_buffercache_pages_6), int32(20), int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v44, int32(7), int32(_a_F_pg_buffercache_pages_7), int32(16), int32(-1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v44, int32(8), int32(_a_F_pg_buffercache_pages_8), int32(21), int32(-1), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v43 == int32(9) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_TupleDescInitEntry(m, v44, int32(9), int32(_a_F_pg_buffercache_pages_9), int32(23), int32(-1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v111 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v111 < v120 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	v198 = F_BlessTupleDesc(m, v44)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L41
	}
L23:
	;
	v124 = v44 + int32(28)
	v131 = v111
	v132 = v120
	v134 = v111
	goto L27
L24:
	;
	v188 = v111
	v195 = v120
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v188
	goto L22
L26:
	;
	v188 = v182
	v195 = v161
	goto L25
L27:
	;
	v140 = v124 + v120<<(uint(int32(3))%32) + v131*int32(100)
	v143 = v124 + v131<<(uint(int32(3))%32)
	if v120 != v132 {
		v161 = v132
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v182 = v120
	goto L26
L29:
	;
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+2)))
	if v162 <= int32(0) {
		v182 = v131
		goto L26
	} else {
		goto L37
	}
L30:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+7)))
	if v145 != int32(118) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v161 = v131
	goto L29
L32:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+4)))
	if v148 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+6)))
	if v151&int32(6) != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+2)))
	if v154 <= int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+90)))
	if v157 != int32(118) {
		v161 = v120
		goto L29
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+90)))
	if v165 == int32(118) {
		v182 = v131
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+5)))
	v174 = (v134 + v168 - int32(1)) & (int32(0) - v168)
	if int32(_a_F_pg_buffercache_pages_10) < v174 {
		v182 = v131
		goto L26
	} else {
		goto L39
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v143))) = uint16(v174)
	v180 = v131 + int32(1)
	if v180 != v120 {
		v131 = v180
		v132 = v161
		v134 = v174 + v162
		goto L27
	} else {
		goto L40
	}
L40:
	;
	goto L28
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0])) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v198
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[1]))
	if int32(0) < v204 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v208 = v15 + int32(8)
	v212 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	m.G0 = v15 + int32(96)
	return int64(0)
L45:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[2]))
	if v223 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[3]))
	v230 = v227 + v212*int32(56)
	v231 = F_LockBufHdr(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	v234 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v230)+12)))
	v235 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v230)+4)))
	v236 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v230))))
	v237 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v230)+8)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	v241 = base.AtomicRmwSub64(m, v230, int32(24), int64(4194304))
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)) = uint8(v242)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = base.I64_extend_i32_s(v238 + int32(1))
	v250 = int64(50331648)
	if base.B2i32(v233 != int32(-1))&base.B2i32(v231&v250 == v250) == v242 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)) = uint8(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_tuplestore_putvalues(m, v291, v292, v15+int32(16), v15+int32(7))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L56
	}
L53:
	;
	v257 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+3)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v257
	v289 = int32(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v262)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v235
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+9)) = uint16(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+11)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = base.I64_extend16_s(v234)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = base.I64_extend_i32_u(v233)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v231 & int64(262143)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = int64(base.Ui64(v231)>>(uint(int64(23))%64)) & int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = int64(base.Ui64(v231)>>(uint(int64(18))%64)) & int64(15)
	v289 = v262
	goto L52
L56:
	;
	v300 = v212 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[1]))
	if v300 < v302 {
		v212 = v300
		goto L45
	} else {
		goto L57
	}
L57:
	;
	goto L46
L58:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_pages_11), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_pages_12), int32(104), int32(_a_F_pg_buffercache_pages_13))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_pages_14), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_pages_12), int32(108), int32(_a_F_pg_buffercache_pages_13))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_usage_counts(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	v2 = int32(0)
	v6 = int64(0)
	v24 = m.G0
	v26 = v24 - int32(144)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+128)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+120)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+96)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+88)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v2
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[0]))
	if int32(0) < v55 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v62 = v2
	goto L6
L4:
	;
	v160 = v6
	v161 = v6
	v162 = v6
	v163 = v6
	v164 = v6
	v165 = v6
	v166 = v6
	v167 = v6
	v168 = v6
	v169 = v6
	v170 = v6
	v171 = v6
	v172 = v6
	v173 = v6
	v174 = v6
	v175 = v6
	v176 = v6
	v177 = v6
	goto L5
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v174
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(0)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v186 = v26 + int32(16)
	v188 = v26 + int32(12)
	F_tuplestore_putvalues(m, v183, v184, v186, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[1]))
	v86 = int64(0)
	v89 = base.AtomicRmwCmpxchg64(m, v82+v62*int32(56), int32(24), v86, v86)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[2]))
	if v91 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v137 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+132)))
	v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+64)))
	v139 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+96)))
	v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+128)))
	v141 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+60)))
	v142 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+92)))
	v143 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+124)))
	v144 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+56)))
	v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+88)))
	v146 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+120)))
	v147 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+52)))
	v148 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+84)))
	v149 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+116)))
	v150 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+48)))
	v151 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+80)))
	v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+112)))
	v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+68)))
	v154 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+100)))
	v160 = v154
	v161 = v137
	v162 = v138
	v163 = v139
	v164 = v140
	v165 = v141
	v166 = v142
	v167 = v143
	v168 = v144
	v169 = v145
	v170 = v146
	v171 = v147
	v172 = v148
	v173 = v149
	v174 = v150
	v175 = v151
	v176 = v152
	v177 = v153
	goto L5
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v100 = int32(base.Ui32(base.I32_wrap_i64(v89))>>(uint(int32(18))%32)) & int32(15) << (uint(int32(2)) % 32)
	v103 = v100 + (v26 + int32(112))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 + int32(1)
	if v89&int64(8388608) != int64(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v114 = v26 + int32(80) + v100
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v115 + int32(1)
	goto L14
L13:
	;
	goto L14
L14:
	;
	if v89&int64(262143) != int64(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v126 = v26 + int32(48) + v100
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v127 + int32(1)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v133 = v62 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[0]))
	if v133 < v135 {
		v62 = v133
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	F_tuplestore_putvalues(m, v196, v197, v186, v188)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(2)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	F_tuplestore_putvalues(m, v205, v206, v186, v188)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v165
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(3)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	F_tuplestore_putvalues(m, v214, v215, v186, v188)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(4)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	F_tuplestore_putvalues(m, v223, v224, v186, v188)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(5)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	F_tuplestore_putvalues(m, v232, v233, v186, v188)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v26 + int32(144)
	return int64(0)
}
