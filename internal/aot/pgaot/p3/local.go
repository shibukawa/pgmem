package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LocalProcessControlFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = F_palloc(m, int32(296))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_LocalProcessControlFile[0])) = v3
		F_ReadControlFile(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_StartLocalBufferIO(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l0 + int32(36)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 != int32(-1) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v16
		if l1 != 0 {
			v27 = int32(0)
			m.G0 = v7 + int32(16)
			return v27
		} else {
			F_pgaio_wref_wait(m, v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v22&int32(16777216) != 0 {
					v27 = int32(0)
				} else {
					v27 = int32(1)
				}
				m.G0 = v7 + int32(16)
				return v27
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v22&int32(16777216) != 0 {
			v27 = int32(0)
		} else {
			v27 = int32(1)
		}
		m.G0 = v7 + int32(16)
		return v27
	}
}
func F_local_buffer_write_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v14 = v7 + int32(8)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_write_error_callback[0]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_GetRelationPath(m, v14, v15, v16, v17, v19, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14
				F_errcontext_msg(m, int32(_a_F_local_buffer_write_error_callback_0), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v7 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v7 + int32(80)
		return
	}
}
func F_read_local_xlog_page(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l1 + base.I64_extend_i32_s(l2)
	goto L2
L1:
	;
	if base.Ui64(v85) < base.Ui64(l1-int64(-8192)) {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_local_xlog_page[0])))
	if v27 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	v85 = v84
	v86 = v74
	goto L1
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L18
	}
L5:
	;
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	v35 = base.B2i32(v33 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_read_local_xlog_page[0])) = uint8(v35)
	v37 = v35
	goto L8
L7:
	;
	v37 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v41 = v12 + int32(4)
	v43 = int32(_a_F_read_local_xlog_page_0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v45 = int64(0)
	v48 = base.AtomicRmwCmpxchg64(m, v44, int32(280), v45, v45)
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[2])) = v48
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v56 = base.AtomicRmwCmpxchg64(m, v52, int32(272), v45, v45)
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[3])) = v56
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v66 = F_GetXLogReplayRecPtr(m, v12+int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v70 = v63
	goto L4
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v60
	goto L15
L14:
	;
	goto L15
L15:
	;
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[2]))
	goto L12
L16:
	;
	return int32(0)
L17:
	;
	v70 = v66
	goto L4
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v74 == v75 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.Ui64(v15) <= base.Ui64(v70) {
		v85 = v70
		v86 = v71
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[4]))
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	goto L2
L26:
	;
	goto L25
L27:
	;
	m.G0 = v12 + int32(48)
	return v103
L28:
	;
	if base.Ui64(v85) < base.Ui64(v15) {
		v103 = int32(-1)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	v95 = int32(_a_F_read_local_xlog_page_1)
	goto L30
L30:
	;
	v97 = v12 + int32(8)
	v98 = F_WALRead(m, l0, l4, l1, v95, v86, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	v95 = base.I32_wrap_i64(v85 - l1)
	goto L30
L32:
	;
	if v98 != 0 {
		v103 = v95
		goto L27
	} else {
		goto L33
	}
L33:
	;
	F_WALReadRaiseError(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v103 = v95
	goto L27
}
func F_update_local_synced_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v337 int64
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[0]))
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v17)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v19)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	if base.Ui64(v22) <= base.Ui64(v21) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L21
	} else {
		goto L92
	}
L8:
	;
	m.G0 = v13 + int32(144)
	return v308
L9:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	if base.Ui64(v90) < base.Ui64(v89) {
		v112 = v88
		goto L31
	} else {
		goto L32
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v25))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v24)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v40 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	if v43 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	if v37 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L14:
	;
	v37 = base.B2i32(base.Ui32(v24) < base.Ui32(v25))
	goto L13
L15:
	;
	goto L16
L16:
	;
	v37 = int32(base.Ui32(v24-v25) >> (uint(int32(31)) % 32))
	goto L13
L17:
	;
	goto L12
L18:
	;
	v46 = int32(15)
	goto L20
L19:
	;
	v46 = int32(14)
	goto L20
L20:
	;
	v48 = F_errstart(m, v46, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v52
	F_errmsg(m, int32(_a_F_update_local_synced_slot_0), v13+int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if l3 == int32(0) {
		v308 = v40
		goto L8
	} else {
		goto L29
	}
L26:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v62
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+16)) = uint32(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v60
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v59)
	v67 = int64(32)
	v68 = int64(base.Ui64(v61) >> (uint(v67) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+12)) = uint32(v68)
	v71 = int64(base.Ui64(v59) >> (uint(v67) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v71)
	F_errdetail(m, int32(_a_F_update_local_synced_slot_1), v13)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(220), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v86)
	v308 = v40
	goto L8
L30:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if l1 != v194 {
		goto L62
	} else {
		goto L63
	}
L31:
	;
	v113 = m.G0
	v115 = v113 - int32(1152)
	m.G0 = v115
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = int32(_a_F_update_local_synced_slot_4)
	*(*uint32)(unsafe.Add(mBase, uint32(v115)+24)) = uint32(v112)
	v121 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v115)+20)) = uint32(v121)
	v124 = v115 + int32(128)
	v128 = F_pg_sprintf(m, v124, int32(_a_F_update_local_synced_slot_5), v115+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L21
	} else {
		goto L41
	}
L32:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	if base.Ui64(v92) < base.Ui64(v88) {
		v112 = v88
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v95))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v94)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v107 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v107 = base.B2i32(base.Ui32(v95) < base.Ui32(v94))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v107 = base.B2i32(int32(0) < v94-v95)
	goto L34
L38:
	;
	v190 = int32(0)
	goto L30
L39:
	;
	goto L40
L40:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v112 = v111
	goto L31
L41:
	;
	v134 = F___fstatat(m, int32(-100), v124, v115+int32(32), int32(0))
	mBase = m.M
	goto L43
L42:
	;
	m.G0 = v115 + int32(1152)
	if v134 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	if v134 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_update_local_synced_slot[1]))
	if v138 == int32(44) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v124
	F_errmsg(m, int32(_a_F_update_local_synced_slot_6), v115)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_7), int32(2073), int32(_a_F_update_local_synced_slot_8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v161 = int32(1)
	v164 = base.AtomicRmwXchg32(m, v16, int32(0), v161)
	if v164 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v184 = F_LogicalSlotAdvanceAndCheckSnapState(m, v183, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L21
	} else {
		goto L58
	}
L53:
	;
	F_s_lock(m, v16, int32(_a_F_update_local_synced_slot_2), int32(259), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v174
	v176 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v176))
	if l2 == v176 {
		v190 = v161
		goto L30
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v181 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v181)
	v190 = v161
	goto L30
L58:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v186 != v187 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v190 = int32(1)
	goto L30
L60:
	;
	v288 = int32(1)
	v291 = base.AtomicRmwXchg32(m, v16, int32(0), v288)
	if v291 != 0 {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	v280 = int32(0)
	if v190 == v280 {
		v308 = v280
		goto L8
	} else {
		goto L83
	}
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = F_strncpy(m, v13+int32(80), v236, int32(64))
	mBase = m.M
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+63)) = uint8(v239)
	goto L75
L63:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+136)))
	if v196 != v197 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+202)))
	if v199 != v200 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = v16 + int32(137)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if base.B2i32(v207 == int32(0))|base.B2i32(v207 != v210) != 0 {
		v228 = v207
		v229 = v210
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v228-v229 != 0 {
		goto L62
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	v213 = v202
	v214 = v204
	goto L69
L69:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	if v218 == int32(0) {
		v228 = v218
		v229 = v217
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v228 = v218
	v229 = v217
	goto L67
L71:
	;
	v221 = int32(1)
	if v218 == v217 {
		v213 = v213 + v221
		v214 = v214 + v221
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)+128))
	if v231 == v232 {
		goto L61
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v241 = int32(1)
	v244 = base.AtomicRmwXchg32(m, v16, int32(0), v241)
	if v244 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_s_lock(m, v16, int32(_a_F_update_local_synced_slot_2), int32(297), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L21
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v13)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+193)) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+185)) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+177)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v13)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+169)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+161)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v13)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+153)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+145)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+137)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = l1
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+136)) = uint8(v267)
	v269 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v269
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+202)) = uint8(v271)
	v273 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v273))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L21
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	if v190 != 0 {
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v308 = v241
	goto L8
L83:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	goto L60
L86:
	;
	F_s_lock(m, v16, int32(_a_F_update_local_synced_slot_2), int32(333), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L21
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v297
	v299 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v299))
	F_ReplicationSlotsComputeRequiredXmin(m, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L21
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v308 = v288
	goto L8
L92:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v322
	F_errmsg_internal(m, int32(_a_F_update_local_synced_slot_9), v13-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L21
	} else {
		goto L93
	}
L93:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+60)) = uint32(v330)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+52)) = uint32(v329)
	v333 = int64(32)
	v334 = int64(base.Ui64(v330) >> (uint(v333) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+56)) = uint32(v334)
	v337 = int64(base.Ui64(v329) >> (uint(v333) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+48)) = uint32(v337)
	F_errdetail_internal(m, int32(_a_F_update_local_synced_slot_10), v13+int32(48))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(280), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
