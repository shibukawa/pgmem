package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l1 + base.I64_extend_i32_s(l2)
	goto L2
L1:
	;
	if base.Ui64(v84) < base.Ui64(l1-int64(-8192)) {
		goto L29
	} else {
		goto L30
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
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	v84 = v83
	v85 = v72
	goto L1
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	v44 = int32(_a_F_read_local_xlog_page_0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+280)) = v46
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[2])) = v46
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+272)) = v52
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[3])) = v52
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
	v64 = F_GetXLogReplayRecPtr(m, v12+int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v68 = v61
	goto L4
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[1]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v58
	goto L15
L14:
	;
	goto L15
L15:
	;
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page[2]))
	goto L12
L16:
	;
	return int32(0)
L17:
	;
	v68 = v64
	goto L4
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v72 == v73 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.Ui64(v15) <= base.Ui64(v68) {
		v84 = v68
		v85 = v69
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
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page[4]))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
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
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L2
L28:
	;
	m.G0 = v12 + int32(48)
	return v102
L29:
	;
	if base.Ui64(v84) < base.Ui64(v15) {
		v102 = int32(-1)
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v94 = int32(_a_F_read_local_xlog_page_1)
	goto L31
L31:
	;
	v96 = v12 + int32(8)
	v97 = F_WALRead(m, l0, l4, l1, v94, v85, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L33
	}
L32:
	;
	v94 = base.I32_wrap_i64(v84 - l1)
	goto L31
L33:
	;
	if v97 != 0 {
		v102 = v94
		goto L28
	} else {
		goto L34
	}
L34:
	;
	F_WALReadRaiseError(m, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	v102 = v94
	goto L28
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
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v336 int64
	_ = v336
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L21
	} else {
		goto L92
	}
L8:
	;
	m.G0 = v13 + int32(144)
	return v307
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
		v307 = v40
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
	v307 = v40
	goto L8
L30:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if l1 != v193 {
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
	v189 = int32(0)
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
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v162
	if v161 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v183 = F_LogicalSlotAdvanceAndCheckSnapState(m, v182, l2)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v174
	if l2 == v175 {
		v189 = v162
		goto L30
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v180)
	v189 = v162
	goto L30
L58:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v185 != v186 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v189 = int32(1)
	goto L30
L60:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v288 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v288
	if v287 != 0 {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	v278 = int32(0)
	if v189 == v278 {
		v307 = v278
		goto L8
	} else {
		goto L83
	}
L62:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v237 = F_strncpy(m, v13+int32(80), v235, int32(64))
	mBase = m.M
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+63)) = uint8(v238)
	goto L75
L63:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+136)))
	if v195 != v196 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+202)))
	if v198 != v199 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v203 = v16 + int32(137)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if base.B2i32(v206 == int32(0))|base.B2i32(v206 != v209) != 0 {
		v227 = v206
		v228 = v209
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v227-v228 != 0 {
		goto L62
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	v212 = v201
	v213 = v203
	goto L69
L69:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
	if v217 == int32(0) {
		v227 = v217
		v228 = v216
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v227 = v217
	v228 = v216
	goto L67
L71:
	;
	v220 = int32(1)
	if v217 == v216 {
		v212 = v212 + v220
		v213 = v213 + v220
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v16)+128))
	if v230 == v231 {
		goto L61
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v241 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v241
	if v240 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_s_lock(m, v16, int32(_a_F_update_local_synced_slot_2), int32(297), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L21
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v13)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+193)) = v249
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+185)) = v251
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+177)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v13)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+169)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+161)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v13)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+153)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+145)) = v261
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+137)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = l1
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+136)) = uint8(v266)
	v268 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v268
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+202)) = uint8(v270)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	if v189 != 0 {
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v307 = v241
	goto L8
L83:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	v295 = m.ExcPending
	if v295 != 0 {
		goto L21
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v296
	F_ReplicationSlotsComputeRequiredXmin(m, v297)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	v304 = m.ExcPending
	if v304 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v307 = v288
	goto L8
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v321
	F_errmsg_internal(m, int32(_a_F_update_local_synced_slot_9), v13-int32(-64))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L21
	} else {
		goto L93
	}
L93:
	;
	v328 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+60)) = uint32(v329)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+52)) = uint32(v328)
	v332 = int64(32)
	v333 = int64(base.Ui64(v329) >> (uint(v332) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+56)) = uint32(v333)
	v336 = int64(base.Ui64(v328) >> (uint(v332) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+48)) = uint32(v336)
	F_errdetail_internal(m, int32(_a_F_update_local_synced_slot_10), v13+int32(48))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_update_local_synced_slot_2), int32(280), int32(_a_F_update_local_synced_slot_3))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
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
