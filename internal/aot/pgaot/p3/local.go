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
		*(*int32)(unsafe.Add(mBase, _consts[36])) = v3
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = *(*int32)(unsafe.Add(mBase, _consts[100]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_GetRelationPath(m, v6+int32(8), v14, v15, v16, v18, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v11
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v6 + int32(8)
				F_errcontext_msg(m, int32(678726), v6)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v6 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v6 + int32(80)
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	v35 = base.B2i32(v33 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v35)
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
	v44 = int32(4360064)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+280)) = v46
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v46
	v51 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+272)) = v52
	*(*int64)(unsafe.Add(mBase, _consts[179])) = v52
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
	v57 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v58
	goto L15
L14:
	;
	goto L15
L15:
	;
	v61 = *(*int64)(unsafe.Add(mBase, _consts[178]))
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
	v77 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	return v103
L29:
	;
	if base.Ui64(v84) < base.Ui64(v15) {
		v103 = int32(-1)
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v94 = int32(8192)
	goto L31
L31:
	;
	v97 = F_WALRead(m, l0, l4, l1, v94, v85, v12+int32(8))
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
		v103 = v94
		goto L28
	} else {
		goto L34
	}
L34:
	;
	F_WALReadRaiseError(m, v12+int32(8))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	v103 = v94
	goto L28
}
func F_update_local_synced_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int64
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v242 int32
	_ = v242
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
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
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
	v10 = m.G0
	v12 = v10 - int32(144)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v16)
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
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v18)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
	if base.Ui64(v21) <= base.Ui64(v20) {
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
		goto L93
	}
L8:
	;
	m.G0 = v12 + int32(144)
	return v309
L9:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v15)+120))
	if base.Ui64(v89) < base.Ui64(v88) {
		v111 = v87
		goto L31
	} else {
		goto L32
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v24))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v23)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v39 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v42 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	if v36 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L14:
	;
	v36 = base.B2i32(base.Ui32(v23) < base.Ui32(v24))
	goto L13
L15:
	;
	goto L16
L16:
	;
	v36 = int32(base.Ui32(v23-v24) >> (uint(int32(31)) % 32))
	goto L13
L17:
	;
	goto L12
L18:
	;
	v45 = int32(15)
	goto L20
L19:
	;
	v45 = int32(14)
	goto L20
L20:
	;
	v47 = F_errstart(m, v45, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v47 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v51
	F_errmsg(m, int32(669333), v12+int32(32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
		v309 = v39
		goto L8
	} else {
		goto L29
	}
L26:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v61
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+16)) = uint32(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v59
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v58)
	v66 = int64(32)
	v67 = int64(base.Ui64(v60) >> (uint(v66) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v67)
	v70 = int64(base.Ui64(v58) >> (uint(v66) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v70)
	F_errdetail(m, int32(549811), v12)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(488687), int32(220), int32(83504))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v85)
	v309 = v39
	goto L8
L30:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if l1 != v195 {
		goto L62
	} else {
		goto L63
	}
L31:
	;
	v112 = m.G0
	v114 = v112 - int32(1152)
	m.G0 = v114
	*(*int32)(unsafe.Add(mBase, uint32(v114)+16)) = int32(116093)
	*(*uint32)(unsafe.Add(mBase, uint32(v114)+24)) = uint32(v111)
	v120 = int64(base.Ui64(v111) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v114)+20)) = uint32(v120)
	v127 = F_pg_sprintf(m, v114+int32(128), int32(232903), v114+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L21
	} else {
		goto L41
	}
L32:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
	if base.Ui64(v91) < base.Ui64(v87) {
		v111 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v94))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v93)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v106 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v106 = base.B2i32(base.Ui32(v94) < base.Ui32(v93))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v106 = base.B2i32(int32(0) < v93-v94)
	goto L34
L38:
	;
	v192 = int32(0)
	goto L30
L39:
	;
	goto L40
L40:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = v110
	goto L31
L41:
	;
	v135 = F___fstatat(m, int32(-100), v114+int32(128), v114+int32(32), int32(0))
	mBase = m.M
	goto L43
L42:
	;
	m.G0 = v114 + int32(1152)
	if v135 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	if v135 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v139 == int32(44) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v114 + int32(128)
	F_errmsg(m, int32(291252), v114)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(488448), int32(2073), int32(114580))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v165 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v165
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
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v186 = F_LogicalSlotAdvanceAndCheckSnapState(m, v185, l2)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L21
	} else {
		goto L58
	}
L53:
	;
	F_s_lock(m, v15, int32(488687), int32(259), int32(83504))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v173
	v175 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v177
	if l2 == v178 {
		v192 = v165
		goto L30
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v183)
	v192 = v165
	goto L30
L58:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v15)+120))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v188 != v189 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v192 = int32(1)
	goto L30
L60:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v290 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v290
	if v289 != 0 {
		goto L87
	} else {
		goto L88
	}
L61:
	;
	v279 = int32(0)
	if v192 == v279 {
		v309 = v279
		goto L8
	} else {
		goto L84
	}
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = F_strncpy(m, v12+int32(80), v236, int32(64))
	mBase = m.M
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+63)) = uint8(v239)
	goto L76
L63:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+136)))
	if v197 != v198 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+202)))
	if v200 != v201 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v205 = v15 + int32(137)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v209 == int32(0) {
		v228 = v208
		v229 = v209
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v229-v228 != 0 {
		goto L62
	} else {
		goto L74
	}
L67:
	;
	goto L66
L68:
	;
	if v208 != v209 {
		v228 = v208
		v229 = v209
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v213 = v203
	v214 = v205
	goto L70
L70:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	if v218 == int32(0) {
		v228 = v217
		v229 = v218
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v228 = v217
	v229 = v218
	goto L67
L72:
	;
	v221 = int32(1)
	if v217 == v218 {
		v213 = v213 + v221
		v214 = v214 + v221
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v15)+128))
	if v231 == v232 {
		goto L61
	} else {
		goto L75
	}
L75:
	;
	goto L62
L76:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v242 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v242
	if v241 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_s_lock(m, v15, int32(488687), int32(297), int32(83504))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L21
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+137)) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v12)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+193)) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+185)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v12)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+177)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+169)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+161)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+153)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+145)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l1
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+136)) = uint8(v267)
	v269 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v269
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+202)) = uint8(v271)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L21
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	if v192 != 0 {
		goto L60
	} else {
		goto L83
	}
L83:
	;
	v309 = v242
	goto L8
L84:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L21
	} else {
		goto L86
	}
L86:
	;
	goto L60
L87:
	;
	F_s_lock(m, v15, int32(488687), int32(333), int32(83504))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L21
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v299 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v298
	F_ReplicationSlotsComputeRequiredXmin(m, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L21
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L21
	} else {
		goto L92
	}
L92:
	;
	v309 = v290
	goto L8
L93:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v322
	F_errmsg_internal(m, int32(84284), v12-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v15)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+60)) = uint32(v330)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v329)
	v333 = int64(32)
	v334 = int64(base.Ui64(v330) >> (uint(v333) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+56)) = uint32(v334)
	v337 = int64(base.Ui64(v329) >> (uint(v333) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)) = uint32(v337)
	F_errdetail_internal(m, int32(628274), v12+int32(48))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(488687), int32(280), int32(83504))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L21
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
