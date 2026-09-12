package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EndReplicationCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), l0, v59+int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v59 = v51 - l0
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = l0
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	return
L19:
	;
	return
}
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	v4 = l3
	v5 = l4
	v6 = l5
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(2224)
	m.G0 = v16
	v19 = F_ReplicationSlotValidateName(m, l0, int32(21))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v5 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L104
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L100
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L96
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L92
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L88
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L84
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v51 = F_LWLockAcquire(m, v47+int32(4608), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v25 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+316))
	v33 = base.B2i32(v31 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v33)
	v35 = v33
	goto L14
L13:
	;
	v35 = int32(0)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[517])))
	if v37 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l2 != int32(2) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[517])))
	if v43 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v58 = F_LWLockAcquire(m, v54+int32(4736), int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[543]))
	if v61 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v65+int32(4736))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v78 = v7
	v80 = v7
	goto L27
L26:
	;
	goto L3
L27:
	;
	v87 = v71 + v80*int32(288)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v126+int32(4736))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L48
	}
L29:
	;
	v92 = v87 + int32(24)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	if v78 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v116-v115 == int32(0) {
		goto L6
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v100 = l0
	v101 = v92
	goto L36
L36:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L33
	} else {
		goto L38
	}
L37:
	;
	v115 = v104
	v116 = v105
	goto L33
L38:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L31
L41:
	;
	v120 = v78
	goto L43
L42:
	;
	v120 = v87
	goto L43
L43:
	;
	if v88 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v121 = v78
	goto L46
L45:
	;
	v121 = v120
	goto L46
L46:
	;
	v123 = v80 + int32(1)
	if v123 != v61 {
		v78 = v121
		v80 = v123
		goto L27
	} else {
		goto L47
	}
L47:
	;
	goto L28
L48:
	;
	if v121 == int32(0) {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v138 = F__emscripten_memset_bulkmem(m, v121+int32(24), base.I32_extend8_s(int32(0)), int32(184))
	mBase = m.M
	goto L50
L50:
	;
	v140 = F_strncpy(m, v138, l0, int32(64))
	mBase = m.M
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+63)) = uint8(v141)
	goto L51
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+136)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+92)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+202)) = uint8(v5)
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+128)) = v148
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+201)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+236)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v121)+16)) = v148
	v155 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)) = uint16(v155)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+244)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v121)+252)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v121)+260)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v121)+268)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v121)+276)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v121)+284)) = v155
	if l1 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v170 = v144
	goto L54
L53:
	;
	v170 = v155
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+88)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = int32(83530)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v138
	v180 = F_pg_sprintf(m, v16+int32(176), int32(174083), v16-int32(-64))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = int32(83530)
	v190 = F_pg_sprintf(m, v16+int32(1200), int32(231633), v16+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v198 = F___fstatat(m, int32(-100), v16+int32(1200), v16+int32(80), int32(0))
	mBase = m.M
	goto L58
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v212 = F_mkdir(m, v16+int32(1200), v211)
	mBase = m.M
	goto L62
L58:
	;
	if v198 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	if v199&int32(61440) != int32(16384) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v206 = F_rmtree(m, v16+int32(1200))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	if v212 < int32(0) {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_fsync_fname(m, v16+int32(1200), int32(1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+13)) = uint8(v220)
	F_SaveSlotToPath(m, v121, v16+int32(1200), int32(21))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v231 = F_rename(m, v16+int32(1200), v16+int32(176))
	mBase = m.M
	if v231 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v232 = int32(4465060)
	v234 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v235 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v234 + v235
	F_fsync_fname(m, v16+int32(176), v235)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_fsync_fname(m, int32(83530), int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v247 = int32(4465060)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v249 - int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v258 = F_LWLockAcquire(m, v254+int32(4736), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v260 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)) = uint8(v260)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v260
	if v262 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_s_lock(m, v121, int32(484206), int32(476), int32(349323))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v271
	*(*int32)(unsafe.Add(mBase, _consts[512])) = v121
	v278 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v278+int32(4736))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v121)+88))
	if v283 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v290 = base.I32_div_s(v121-v287, int32(288))
	goto L78
L76:
	;
	goto L77
L77:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v306+int32(4608))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L82
	}
L78:
	;
	v293 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v290), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v301 = F__emscripten_memset_bulkmem(m, v295+int32(24), base.I32_extend8_s(int32(0)), int32(72))
	mBase = m.M
	goto L80
L80:
	;
	F_pgstat_unlock_entry(m, v293)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	F_ConditionVariableBroadcast(m, v121+int32(224))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	m.G0 = v16 + int32(2224)
	return
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(23489), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(484206), int32(377), int32(349323))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(84060), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(484206), int32(389), int32(349323))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg(m, int32(113984), v16)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(484206), int32(414), int32(349323))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(1200)
	F_errmsg(m, int32(291622), v16+int32(16))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(484206), int32(2288), int32(309380))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v16 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v16 + int32(1200)
	F_errmsg(m, int32(292340), v16+int32(32))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(484206), int32(2300), int32(309380))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(353933), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errhint(m, int32(637563), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(484206), int32(425), int32(349323))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
