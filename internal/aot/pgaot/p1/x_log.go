package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetXLogReceiptTime(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, _c_F_GetXLogReceiptTime[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReceiptTime[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v7 == int32(3)))
	return
}
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
		F_s_lock(m, v11+int32(96), int32(_a_F_GetXLogReplayRecPtr_0), int32(_a_F_GetXLogReplayRecPtr_1), int32(_a_F_GetXLogReplayRecPtr_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(0)
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
			if l0 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
			} else {
			}
			return v25
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogReplayRecPtr[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(0)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
		if l0 != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
		} else {
		}
		return v25
	}
}
func F_WaitXLogInsertionsToFinish(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v61 int64
	_ = v61
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v122 int64
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v167 int64
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int64
	_ = v292
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v310 int64
	_ = v310
	var v323 int64
	_ = v323
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v330 int64
	_ = v330
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	v1 = l0
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[0]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[1]))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+264))
	if base.Ui64(v1) <= base.Ui64(v23) {
		v330 = v23
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L9
	} else {
		goto L77
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v330
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(1)
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_s_lock(m, v22, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1528), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(0)
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v39 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[2])))
	v40 = base.I64_div_u_s(v37, v39)
	v42 = v37 - v40*v39
	if base.Ui64(v42) <= base.Ui64(int64(8151)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	return int64(0)
L10:
	;
	goto L8
L11:
	;
	v72 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[3])))
	v76 = v40*v72 + v70&int64(4294967295)
	if base.Ui64(v1) <= base.Ui64(v76) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v45 = int64(0)
	if v42 == v45 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v52 = v42 - int64(8152)
	v53 = int64(8168)
	v54 = base.I64_div_u_s(v52, v53)
	v56 = v54 << (uint(int64(13)) % 64)
	v61 = v52 - v54*v53
	if v61 == int64(0) {
		v70 = v56 - int64(-8192)
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v50 = v45
	goto L17
L16:
	;
	v50 = v42 + int64(40)
	goto L17
L17:
	;
	v70 = v50
	goto L11
L18:
	;
	v70 = v61 + v56 + int64(8216)
	goto L11
L19:
	;
	v104 = v76
	v113 = int32(0)
	goto L27
L20:
	;
	v100 = v1
	goto L19
L21:
	;
	goto L22
L22:
	;
	v80 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if v80 == int32(0) {
		v100 = v76
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+12)) = uint32(v76)
	v85 = int64(32)
	v86 = int64(base.Ui64(v76) >> (uint(v85) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+8)) = uint32(v86)
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)) = uint32(v1)
	v90 = int64(base.Ui64(v1) >> (uint(v85) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17))) = uint32(v90)
	F_errmsg(m, int32(_a_F_WaitXLogInsertionsToFinish_2), v17)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1545), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v100 = v76
	goto L19
L27:
	;
	v115 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v115
	v122 = v115
	goto L30
L28:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[1]))
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v305)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v305)+264)) = v306
	if base.Ui64(v299) <= base.Ui64(v306) {
		goto L65
	} else {
		goto L66
	}
L29:
	;
	v301 = v113 + int32(1)
	if v301 != int32(8) {
		v104 = v299
		v113 = v301
		goto L27
	} else {
		goto L64
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[4]))
	v136 = v135 + v113<<(uint(int32(7))%32)
	v138 = v136 + int32(16)
	v140 = v17 + int32(24)
	v141 = int32(0)
	v142 = int32(_a_F_WaitXLogInsertionsToFinish_3)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v144 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[0]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v150&int32(_a_F_WaitXLogInsertionsToFinish_4) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v292 == int64(0) {
		v299 = v104
		goto L29
	} else {
		goto L60
	}
L32:
	;
	if int32(0) < v242 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	v158 = v141
	goto L36
L34:
	;
	v227 = v141
	goto L35
L35:
	;
	v242 = v227
	v246 = int32(1)
	goto L32
L36:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v167
	if v167 != v122 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v227 = v203
	goto L35
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140))) = v167
	v242 = v158
	v246 = int32(0)
	goto L32
L39:
	;
	goto L40
L40:
	;
	F_LWLockQueueSelf(m, v136, int32(2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v177 = v175 | int32(1073741824)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v177
	v181 = v177 & int32(_a_F_WaitXLogInsertionsToFinish_4)
	if v181 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136))))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193 | int32(16777216)
	v203 = v158
	goto L48
L43:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v182
	if v182 == v122 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_LWLockDequeueSelf(m, v136)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L47
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140))) = v182
	goto L45
L47:
	;
	v242 = v158
	v246 = base.B2i32(v181 == int32(0))
	goto L32
L48:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+74)))
	if v214 != 0 {
		v203 = v203 + int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v219&int32(_a_F_WaitXLogInsertionsToFinish_4) != 0 {
		v158 = v203
		goto L36
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L37
L52:
	;
	v258 = v242
	goto L55
L53:
	;
	goto L54
L54:
	;
	v286 = int32(_a_F_WaitXLogInsertionsToFinish_3)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitXLogInsertionsToFinish[5])) = v288 - int32(1)
	if v246 != 0 {
		v299 = v104
		goto L29
	} else {
		goto L58
	}
L55:
	;
	v268 = int32(1)
	if base.Ui32(v268) < base.Ui32(v258) {
		v258 = v258 - v268
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	goto L56
L58:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	if base.Ui64(v292) < base.Ui64(v100) {
		v122 = v292
		goto L30
	} else {
		goto L59
	}
L59:
	;
	goto L31
L60:
	;
	if base.Ui64(v292) < base.Ui64(v104) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v297 = v292
	goto L63
L62:
	;
	v297 = v104
	goto L63
L63:
	;
	v299 = v297
	goto L29
L64:
	;
	goto L28
L65:
	;
	v330 = v306
	goto L4
L66:
	;
	goto L67
L67:
	;
	v310 = v306
	goto L68
L68:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v305)+264))
	v324 = base.B2i32(v310 == v323)
	if v310 == v323 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v330 = v323
	goto L4
L70:
	;
	v325 = v299
	goto L72
L71:
	;
	v325 = v323
	goto L72
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v305)+264)) = v325
	if v310 == v323 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v330 = v299
	goto L4
L74:
	;
	goto L75
L75:
	;
	if base.Ui64(v323) < base.Ui64(v299) {
		v310 = v323
		goto L68
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	F_errmsg_internal(m, int32(_a_F_WaitXLogInsertionsToFinish_5), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_WaitXLogInsertionsToFinish_0), int32(1517), int32(_a_F_WaitXLogInsertionsToFinish_1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogArchiveForceDone(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v5 = m.G0
	v7 = v5 - int32(2208)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = int32(_a_F_XLogArchiveForceDone_0)
	v13 = v7 + int32(160)
	v18 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_XLogArchiveForceDone_1), v7+int32(48))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = v7 - int32(-64)
		v24 = F___fstatat(m, int32(-100), v13, v21, int32(0))
		mBase = m.M
		if v24 == int32(0) {
			m.G0 = v7 + int32(2208)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(_a_F_XLogArchiveForceDone_2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
			v31 = v7 + int32(1184)
			v36 = F_pg_snprintf(m, v31, int32(1024), int32(_a_F_XLogArchiveForceDone_1), v7+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v40 = F___fstatat(m, int32(-100), v31, v21, int32(0))
				mBase = m.M
				if v40 == int32(0) {
					v44 = F_durable_rename(m, v31, v13, int32(19))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(2208)
						return
					}
				} else {
					v47 = v7 + int32(160)
					v49 = F_AllocateFile(m, v47, int32(_a_F_XLogArchiveForceDone_3))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						if v49 == int32(0) {
							v55 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								if v55 == int32(0) {
									m.G0 = v7 + int32(2208)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v47
										F_errmsg(m, int32(_a_F_XLogArchiveForceDone_4), v7)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(537), int32(_a_F_XLogArchiveForceDone_6))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												m.G0 = v7 + int32(2208)
												return
											}
										}
									}
								}
							}
						} else {
							v70 = F_FreeFile(m, v49)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								if v70 == int32(0) {
									m.G0 = v7 + int32(2208)
									return
								} else {
									v76 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										if v76 == int32(0) {
											m.G0 = v7 + int32(2208)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(160)
												F_errmsg(m, int32(_a_F_XLogArchiveForceDone_7), v7+int32(16))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_XLogArchiveForceDone_5), int32(545), int32(_a_F_XLogArchiveForceDone_6))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														m.G0 = v7 + int32(2208)
														return
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
func F_XLogArchiveIsBusy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(1184)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a_F_XLogArchiveIsBusy_0)
	v14 = v8 + int32(160)
	v19 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(48))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = v8 - int32(-64)
		v27 = F___fstatat(m, int32(-100), v14, v24, int32(0))
		mBase = m.M
		if v27 == int32(0) {
			v73 = v2
			m.G0 = v8 + int32(1184)
			return v73
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_XLogArchiveIsBusy_2)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
			v37 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(32))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v41 = F___fstatat(m, int32(-100), v14, v24, int32(0))
				mBase = m.M
				if v41 == int32(0) {
					v73 = int32(1)
					m.G0 = v8 + int32(1184)
					return v73
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_XLogArchiveIsBusy_0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v51 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_1), v8+int32(16))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = F___fstatat(m, int32(-100), v14, v24, int32(0))
						mBase = m.M
						if v55 == int32(0) {
							v73 = v2
							m.G0 = v8 + int32(1184)
							return v73
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v61 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_XLogArchiveIsBusy_3), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v65 = F___fstatat(m, int32(-100), v14, v24, int32(0))
								mBase = m.M
								if v65 == int32(0) {
									v73 = int32(1)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveIsBusy[0]))
									if v69 == int32(44) {
										v73 = v2
									} else {
										v73 = int32(1)
									}
								}
								m.G0 = v8 + int32(1184)
								return v73
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogBeginRead(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v23
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v27
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v25)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v25)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v23
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v12 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_pfree(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v18 = v10
	goto L8
L8:
	;
	if v18 != 0 {
		v8 = v18
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v18 = v17
	goto L8
L11:
	;
	goto L5
}
func F_XLogBytePosToRecPtr(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	v4 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogBytePosToRecPtr[0])))
	v5 = base.I64_div_u_s(l0, v4)
	v7 = l0 - v5*v4
	if base.Ui64(v7) <= base.Ui64(int64(8151)) {
		v25 = v7 + int64(40)
	} else {
		v13 = v7 - int64(8152)
		v14 = int64(8168)
		v15 = base.I64_div_u_s(v13, v14)
		v25 = v13 - v15*v14 + v15<<(uint(int64(13))%64) + int64(8216)
	}
	v27 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogBytePosToRecPtr[1])))
	return v5*v27 + v25&int64(4294967295)
}
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	if l0 < int32(33) {
		v11 = int32(4)
		if l0 <= v11 {
			v14 = v11
		} else {
			v14 = l0
		}
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0]))
		if v16 <= v14 {
			v18 = int32(_a_F_XLogEnsureRecordSpace_0)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[1]))
			v22 = v14 + int32(1)
			v25 = F_repalloc(m, v20, v22*int32(_a_F_XLogEnsureRecordSpace_1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[1])) = v25
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0]))
				v31 = int32(_a_F_XLogEnsureRecordSpace_1)
				v32 = (v22 - v29) * v31
				v34 = v29 * v31
				v35 = v25 + v34
				if v35&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v32)) == int32(0) {
					if v29 == v22 {
					} else {
						v46 = int32(_a_F_XLogEnsureRecordSpace_1)
						v50 = v14*v46 + v25 + v46
						v53 = v25 + v34 + int32(4)
						if base.Ui32(v53) < base.Ui32(v50) {
							v55 = v50
						} else {
							v55 = v53
						}
						v61 = (v25^int32(-1)+v55-v34)&int32(-4) + int32(4)
						if v61 == int32(0) {
						} else {
							base.MemoryFill(m, v35, int32(0), v61)
						}
					}
				} else {
					if v32 == int32(0) {
					} else {
						base.MemoryFill(m, v35, int32(0), v32)
					}
				}
				*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[0])) = v22
				v81 = int32(20)
				if l1 <= v81 {
					v84 = v81
				} else {
					v84 = l1
				}
				v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2]))
				if v86 < v84 {
					v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3]))
					v92 = F_repalloc(m, v89, v84*int32(12))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2])) = v84
						*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3])) = v92
						return
					}
				} else {
					return
				}
			}
		} else {
			v81 = int32(20)
			if l1 <= v81 {
				v84 = v81
			} else {
				v84 = l1
			}
			v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2]))
			if v86 < v84 {
				v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3]))
				v92 = F_repalloc(m, v89, v84*int32(12))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[2])) = v84
					*(*int32)(unsafe.Add(mBase, _c_F_XLogEnsureRecordSpace[3])) = v92
					return
				}
			} else {
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_XLogEnsureRecordSpace_2), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_XLogEnsureRecordSpace_3), int32(194), int32(_a_F_XLogEnsureRecordSpace_4))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
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
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	v15 = base.I64_div_u_s(int64(4294967296), base.I64_extend_i32_s(l3))
	v16 = base.I64_div_u_s(l2, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v16)
	v19 = l2 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+8)) = uint32(v19)
	v23 = F_pg_snprintf(m, l0, int32(64), int32(_a_F_XLogFileName_0), v10)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		m.G0 = v10 + int32(16)
		return
	}
}
func F_XLogGetOldestSegno(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	v5 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_AllocateDir(m, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = F_ReadDir(m, v12, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v17
	v24 = v5
	goto L7
L5:
	;
	v152 = v5
	goto L6
L6:
	;
	F_FreeDir(m, v12)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v26 = v21 + int32(19)
	v27 = F_strlen(m, v26)
	mBase = m.M
	if v27 != int32(24) {
		v143 = v24
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v143
	goto L6
L9:
	;
	v145 = F_ReadDir(m, v12, int32(_a_F_XLogGetOldestSegno_0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L36
	}
L10:
	;
	v30 = int32(_a_F_XLogGetOldestSegno_1)
	v34 = m.G0
	v36 = v34 - int32(32)
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[0])))
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v113 != int32(24) {
		v143 = v24
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v113 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[1])))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = v26
	goto L18
L16:
	;
	goto L17
L17:
	;
	v63 = v30
	v64 = v45
	goto L21
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == v45 {
		v53 = v53 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v113 = v53 - v26
	goto L11
L20:
	;
	goto L19
L21:
	;
	v71 = v36 + int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v73<<(uint(v64)%32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v77 != 0 {
		v63 = v63 + v73
		v64 = v77
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v80 == int32(0) {
		v103 = v26
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v113 = v103 - v26
	goto L11
L25:
	;
	v84 = v26
	v85 = v80
	goto L26
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(base.Ui32(v85)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v93)>>(uint(v85)%32))&int32(1) == int32(0) {
		v103 = v84
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v103 = v101
	goto L24
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v101 = v84 + int32(1)
	if v99 != 0 {
		v84 = v101
		v85 = v99
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v117 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogGetOldestSegno[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(24)
	v128 = F_sscanf(m, v26, int32(_a_F_XLogGetOldestSegno_2), v9)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if l0 != v130 {
		v143 = v24
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)))
	v135 = base.I64_div_u_s(int64(4294967296), v117)
	v137 = v132 + v133*v135
	if base.Ui64(v24-int64(1)) < base.Ui64(v137) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = v24
	goto L35
L34:
	;
	v141 = v137
	goto L35
L35:
	;
	v143 = v141
	goto L9
L36:
	;
	if v145 != 0 {
		v21 = v145
		v24 = v143
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	m.G0 = v9 + int32(32)
	return v152
}
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = F_XLogReadBufferForRedoExtended(m, l0, l1, int32(1), int32(0), v6+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		m.G0 = v6 + int32(16)
		return v16
	}
}
func F_XLogNeedsFlush(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])))
	if v6 != int32(1) {
		v79 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
		if base.Ui64(l0) <= base.Ui64(v79) {
			v97 = v3
		} else {
			v81 = int32(_a_F_XLogNeedsFlush_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
			v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+280))
			*(*int64)(unsafe.Add(mBase, uint32(v82)+280)) = v83
			v85 = int32(_a_F_XLogNeedsFlush_1)
			*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])) = v83
			v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
			v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+272))
			*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = v89
			*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = v89
			v94 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
			v97 = base.B2i32(base.Ui64(v94) < base.Ui64(l0))
		}
		return v97
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+316))
		v13 = int32(2)
		*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[0])) = uint8(base.B2i32(v12 != v13))
		if v12 == v13 {
			v79 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
			if base.Ui64(l0) <= base.Ui64(v79) {
				v97 = v3
			} else {
				v81 = int32(_a_F_XLogNeedsFlush_0)
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+280))
				*(*int64)(unsafe.Add(mBase, uint32(v82)+280)) = v83
				v85 = int32(_a_F_XLogNeedsFlush_1)
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1])) = v83
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[2]))
				v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+272))
				*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = v89
				*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[3])) = v89
				v94 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[1]))
				v97 = base.B2i32(base.Ui64(v94) < base.Ui64(l0))
			}
			return v97
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
			if v19 != int64(0) {
				if base.Ui64(l0) <= base.Ui64(v19) {
					v97 = v3
					return v97
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
					if v35&int32(1) != 0 {
						v97 = v3
						return v97
					} else {
						v38 = int32(1)
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
						v44 = F_LWLockConditionalAcquire(m, v40+int32(1152), v38)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v44 == int32(0) {
								v97 = v38
								return v97
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7]))
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+136))
								*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v53
								v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
								F_LWLockRelease(m, v56+int32(1152))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v62 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
									if v62 != int64(0) {
										v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
										v71 = v66
									} else {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v68)
										v71 = v68
									}
									return (v71 ^ int32(1)) & base.B2i32(base.Ui64(v62) < base.Ui64(l0))
								}
							}
						}
					}
				}
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[8])))
				if v23&int32(1) == int32(0) {
					if base.Ui64(l0) <= base.Ui64(v19) {
						v97 = v3
						return v97
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
						if v35&int32(1) != 0 {
							v97 = v3
							return v97
						} else {
							v38 = int32(1)
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
							v44 = F_LWLockConditionalAcquire(m, v40+int32(1152), v38)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v44 == int32(0) {
									v97 = v38
									return v97
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[7]))
									v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+136))
									*(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4])) = v53
									v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[6]))
									F_LWLockRelease(m, v56+int32(1152))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v62 = *(*int64)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[4]))
										if v62 != int64(0) {
											v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])))
											v71 = v66
										} else {
											v68 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v68)
											v71 = v68
										}
										return (v71 ^ int32(1)) & base.B2i32(base.Ui64(v62) < base.Ui64(l0))
									}
								}
							}
						}
					}
				} else {
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_XLogNeedsFlush[5])) = uint8(v29)
					return int32(0)
				}
			}
		}
	}
}
func F_XLogReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v322 int32
	_ = v322
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int64
	_ = v380
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int64
	_ = v439
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l3|base.B2i32(l4 == v6) == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(112)
	return v451
L2:
	;
	if v405 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v27
	v30 = v18 + int32(88)
	v31 = m.G0
	v33 = v31 - int32(32)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	F_ResourceOwnerEnlarge(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v341
	v343 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v343
	v348 = F_smgropen(m, v18+int32(72), int32(-1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L84
	}
L6:
	;
	return int32(0)
L7:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if l4 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v33 + int32(32)
	if v322 != 0 {
		v405 = l4
		goto L2
	} else {
		goto L83
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[1]))
	v54 = v49 + (l4^int32(-1))<<(uint(int32(6))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v55&int32(16777216) == int32(0) {
		v322 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = l4
	v85 = v81 + l4<<(uint(int32(6))%32)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[3]))
	if l4 == v87 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v45 != v60 {
		v322 = v6
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v44 != v62 {
		v322 = v6
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if v43 != v64 {
		v322 = v6
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if l2 != v66 {
		v322 = v6
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if l1 != v68 {
		v322 = v6
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v70 = int32(1)
	F_PinLocalBuffer(m, v54, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v74 = int32(_a_F_XLogReadBufferExtended_0)
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[4])) = v76 + int64(1)
	v322 = v70
	goto L9
L20:
	;
	if v229&int32(16777216) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = int32(_a_F_XLogReadBufferExtended_1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = int32(_a_F_XLogReadBufferExtended_2)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(_a_F_XLogReadBufferExtended_3)
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(0)
	v154 = v85 - int32(40)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = int32(_a_F_XLogReadBufferExtended_4)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v155 | v156
	if v155&v156 != 0 {
		goto L51
	} else {
		goto L52
	}
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v134 <= int32(0) {
		goto L21
	} else {
		goto L50
	}
L23:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_5)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[5]))
	if l4 == v91 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_6)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[6]))
	if l4 == v95 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_7)
	goto L22
L30:
	;
	goto L31
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[7]))
	if l4 == v99 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_8)
	goto L22
L33:
	;
	goto L34
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[8]))
	if l4 == v103 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_9)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[9]))
	if l4 == v107 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_10)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[10]))
	if l4 == v111 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_11)
	goto L22
L42:
	;
	goto L43
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[11]))
	if l4 == v115 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v132 = int32(_a_F_XLogReadBufferExtended_12)
	goto L22
L45:
	;
	goto L46
L46:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[12]))
	if v119 == int32(0) {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[13]))
	v126 = int32(0)
	v128 = F_hash_search(m, v123, v33+int32(8), v126, v126)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v128 == int32(0) {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v132 = v128
	goto L22
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(40))))
	v229 = v139
	v233 = int32(1)
	goto L20
L51:
	;
	goto L54
L52:
	;
	v191 = v155
	goto L53
L53:
	;
	v204 = int32(_a_F_XLogReadBufferExtended_13)
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[14]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(8))+8))
	if v207 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	F_perform_spin_delay(m, v33+int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L56
	}
L55:
	;
	v191 = v180
	goto L53
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v181 = int32(_a_F_XLogReadBufferExtended_4)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v180 | v181
	if v180&v181 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v229 = v191
	v233 = v148
	goto L20
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[14])) = v222
	goto L59
L61:
	;
	if int32(999) < v205 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v205 < int32(11) {
		goto L59
	} else {
		goto L68
	}
L64:
	;
	v212 = int32(900)
	if v212 <= v205 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v215 = v212
	goto L67
L66:
	;
	v215 = v205
	goto L67
L67:
	;
	v222 = v215 + int32(100)
	goto L60
L68:
	;
	v222 = v205 - int32(1)
	goto L60
L69:
	;
	if v233 != 0 {
		v322 = v6
		goto L9
	} else {
		goto L82
	}
L70:
	;
	v244 = v85 + int32(-64)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	if v45 != v245 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(60))))
	if v44 != v249 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(56))))
	if v43 != v253 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(48))))
	if l2 != v257 {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(52))))
	if l1 != v261 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	if v233 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v295 = int32(_a_F_XLogReadBufferExtended_14)
	v297 = *(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[15])) = v297 + int64(1)
	v322 = int32(1)
	goto L9
L77:
	;
	v264 = F_PinBuffer(m, v244, int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v267 = v85 - int32(40)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = (v268 + v269) & int32(-4194305)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(44))))
	v277 = int32(_a_F_XLogReadBufferExtended_15)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[16])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v269
	v285 = v276 + v269
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v285
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[0]))
	F_ResourceOwnerRemember(m, v288, v285, int32(_a_F_XLogReadBufferExtended_16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	goto L76
L81:
	;
	goto L76
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85-int32(40)))) = v229 & int32(-4194305)
	v322 = v6
	goto L9
L83:
	;
	goto L5
L84:
	;
	F_smgrcreate(m, v348, l1, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	v353 = F_smgrnblocks(m, v348, l1)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(v353) <= base.Ui32(l2) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if l3 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v391
	v393 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v393
	v399 = F_ReadBufferWithoutRelcache(m, v18+int32(24), l1, l2, l3, int32(0), int32(1))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L96
	}
L90:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v358
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v360
	v362 = int32(0)
	F_log_invalid_page(m, v18+int32(40), l1, l2, v362)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if l3 == int32(4) {
		v451 = int32(0)
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v451 = v362
	goto L1
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v348
	v372 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v372
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+109)) = uint16(v372)
	v376 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+108)) = uint8(v376)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v372)
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v382
	v389 = F_ExtendBufferedRelTo(m, v18+int32(56), l1, int32(3), l2+int32(1), l3)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v451 = v389
	goto L1
L96:
	;
	if l3 != 0 {
		v451 = v399
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v405 = v399
	goto L2
L98:
	;
	v434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433)+14)))
	if v434 != 0 {
		v451 = v405
		goto L1
	} else {
		goto L102
	}
L99:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[17]))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v419+(v405^int32(-1))<<(uint(int32(2))%32))))
	v433 = v425
	goto L98
L100:
	;
	goto L101
L101:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferExtended[18]))
	v433 = v427 + v405<<(uint(int32(13))%32) + int32(-8192)
	goto L98
L102:
	;
	F_ReleaseBuffer(m, v405)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v437
	v439 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v439
	F_log_invalid_page(m, v18+int32(8), l1, l2, int32(1))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v451 = int32(0)
	goto L1
}
func F_XLogReadRecord(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v40 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v4 == v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
	if v15 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v21
	goto L1
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, v4)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v21 = v18
	goto L13
L11:
	;
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v29
	goto L1
L13:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v22 != int32(1) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v25 != 0 {
		v21 = v25
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	goto L1
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v45 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v41 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v43 = F_XLogReadAhead(m, l0, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v79 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v45 == v52 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v56 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v62
	goto L23
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v59 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v45)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L39
	}
L32:
	;
	v62 = v59
	goto L35
L33:
	;
	goto L34
L34:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v70
	goto L23
L35:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)))
	if v63 != int32(1) {
		goto L28
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v66 != 0 {
		v62 = v66
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L23
L40:
	;
	if v103 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v82
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v85 != int32(1) {
		v103 = v82
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v79
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v103 = v101
	goto L40
L44:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v88
	goto L47
L46:
	;
	goto L47
L47:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v91)
	v103 = v91
	goto L40
L48:
	;
	v107 = v103 + int32(32)
	goto L50
L49:
	;
	v107 = int32(0)
	goto L50
L50:
	;
	return v107
}
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v3 = l2
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[0]))
	if v5 <= l0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[1]))
		if v8 <= l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_XLogRegisterBuffer_0), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XLogRegisterBuffer_1), int32(267), int32(_a_F_XLogRegisterBuffer_2))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[0])) = l0 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[2]))
			v18 = v15 + l0*int32(_a_F_XLogRegisterBuffer_3)
			v20 = v18 + int32(4)
			if l1 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[3]))
				v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
				v42 = v36 + l1<<(uint(int32(6))%32) + int32(-64)
			}
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v44
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = v43
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v18+int32(20)))) = v49
			if l1 < int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[5]))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
				v68 = v60
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[6]))
				v68 = v62 + l1<<(uint(int32(13))%32) + int32(-8192)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v68
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(0)
			v73 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v73)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(32)
			return
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[2]))
		v18 = v15 + l0*int32(_a_F_XLogRegisterBuffer_3)
		v20 = v18 + int32(4)
		if l1 < int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[3]))
			v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[4]))
			v42 = v36 + l1<<(uint(int32(6))%32) + int32(-64)
		}
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v44
		*(*int64)(unsafe.Add(mBase, uint32(v20))) = v43
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v47
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v18+int32(20)))) = v49
		if l1 < int32(0) {
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[5]))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
			v68 = v60
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _c_F_XLogRegisterBuffer[6]))
			v68 = v62 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v68
		*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(0)
		v73 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v73)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(32)
		return
	}
}
func F_XLogShutdownWalRcv(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+1456)) = int32(1)
	v10 = v5 + int32(1456)
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v10, int32(_a_F_XLogShutdownWalRcv_0), int32(190), int32(_a_F_XLogShutdownWalRcv_1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	switch v16 - int32(1) {
	case 0:
		goto L9
	case 1, 2, 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L10
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v41 = v5 + int32(12)
	F_ConditionVariablePrepareToSleep(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+1456)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v33 == v31 {
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(5)
	goto L7
L9:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+1456)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v21
	F_ConditionVariableBroadcast(m, v5+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
	goto L6
L11:
	;
	goto L6
L12:
	;
	v37 = F_kill(m, v33, int32(15))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v44 = F_WalRcvRunning(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	F_ConditionVariableSleep(m, v41, int32(134217781))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v52 = F_WalRcvRunning(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v52 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	v64 = F_LWLockAcquire(m, v60+int32(1152), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[2]))
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+320)) = uint8(v68)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_XLogShutdownWalRcv[1]))
	F_LWLockRelease(m, v71+int32(1152))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	return
}
func F_XLogWalRcvFlush(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
	if base.Ui64(v16) <= base.Ui64(v14) {
		m.G0 = v11 + int32(80)
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[2]))
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[3]))
		v23 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[4]))
		F_issue_xlog_fsync(m, v21, v23, l1)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v28 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
			*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0])) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1456))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+1456)) = int32(1)
			v34 = v19 + int32(1456)
			if v30 != 0 {
				F_s_lock(m, v34, int32(_a_F_XLogWalRcvFlush_0), int32(998), int32(_a_F_XLogWalRcvFlush_1))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
					v42 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
					if base.Ui64(v40) < base.Ui64(v42) {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v40
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
					F_WakeupRecovery(m)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5])))
						if v52 != int32(1) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
							if v64 == int32(1) {
								v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
								v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
								v74 = v11 + int32(16)
								v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = F_strlen(m, v74)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v82 = int32(0)
										F_XLogWalRcvSendReply(m, v82, v82)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										}
									}
								}
							} else {
								if l0 != 0 {
									m.G0 = v11 + int32(80)
									return
								} else {
									v82 = int32(0)
									F_XLogWalRcvSendReply(m, v82, v82)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									}
								}
							}
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7]))
							if v56 <= int32(0) {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
								if v64 == int32(1) {
									v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
									v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
									v74 = v11 + int32(16)
									v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = F_strlen(m, v74)
										mBase = m.M
										if l0 != 0 {
											m.G0 = v11 + int32(80)
											return
										} else {
											v82 = int32(0)
											F_XLogWalRcvSendReply(m, v82, v82)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									}
								} else {
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v82 = int32(0)
										F_XLogWalRcvSendReply(m, v82, v82)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										}
									}
								}
							} else {
								F_WalSndWakeup(m, int32(1), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
									if v64 == int32(1) {
										v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
										v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
										v74 = v11 + int32(16)
										v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v79 = F_strlen(m, v74)
											mBase = m.M
											if l0 != 0 {
												m.G0 = v11 + int32(80)
												return
											} else {
												v82 = int32(0)
												F_XLogWalRcvSendReply(m, v82, v82)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_XLogWalRcvSendHSFeedback(m, int32(0))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														m.G0 = v11 + int32(80)
														return
													}
												}
											}
										}
									} else {
										if l0 != 0 {
											m.G0 = v11 + int32(80)
											return
										} else {
											v82 = int32(0)
											F_XLogWalRcvSendReply(m, v82, v82)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
				v42 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[0]))
				if base.Ui64(v40) < base.Ui64(v42) {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v40
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
				F_WakeupRecovery(m)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[5])))
					if v52 != int32(1) {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
						if v64 == int32(1) {
							v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
							v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
							v74 = v11 + int32(16)
							v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v79 = F_strlen(m, v74)
								mBase = m.M
								if l0 != 0 {
									m.G0 = v11 + int32(80)
									return
								} else {
									v82 = int32(0)
									F_XLogWalRcvSendReply(m, v82, v82)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									}
								}
							}
						} else {
							if l0 != 0 {
								m.G0 = v11 + int32(80)
								return
							} else {
								v82 = int32(0)
								F_XLogWalRcvSendReply(m, v82, v82)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[7]))
						if v56 <= int32(0) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
							if v64 == int32(1) {
								v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
								v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
								v74 = v11 + int32(16)
								v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = F_strlen(m, v74)
									mBase = m.M
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v82 = int32(0)
										F_XLogWalRcvSendReply(m, v82, v82)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										}
									}
								}
							} else {
								if l0 != 0 {
									m.G0 = v11 + int32(80)
									return
								} else {
									v82 = int32(0)
									F_XLogWalRcvSendReply(m, v82, v82)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									}
								}
							}
						} else {
							F_WalSndWakeup(m, int32(1), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[6])))
								if v64 == int32(1) {
									v68 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvFlush[1]))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
									v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
									v74 = v11 + int32(16)
									v77 = F_pg_snprintf(m, v74, int32(50), int32(_a_F_XLogWalRcvFlush_2), v11)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = F_strlen(m, v74)
										mBase = m.M
										if l0 != 0 {
											m.G0 = v11 + int32(80)
											return
										} else {
											v82 = int32(0)
											F_XLogWalRcvSendReply(m, v82, v82)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									}
								} else {
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v82 = int32(0)
										F_XLogWalRcvSendReply(m, v82, v82)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
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
func F_XLogWalRcvSendHSFeedback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[1]))
	if v13&base.B2i32(v2 < v15) == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[2])))
	if v22&int32(1) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v28 = m.G0
	v29 = int32(16)
	v30 = v28 - v29
	m.G0 = v30
	F_gettimeofday(m, v30)
	mBase = m.M
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v34 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30)+8)))
	m.G0 = v30 + v29
	v42 = v34 + v33*int64(1000000) - int64(946684800000000)
	goto L6
L5:
	;
	goto L4
L6:
	;
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[3]))
	if v42 < v46 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v51 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[1])))
	if v51 <= int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v57 = int64(9223372036854775807)
	goto L13
L12:
	;
	v57 = v42 + v51*int64(1000000)
	goto L13
L13:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v57
	goto L16
L15:
	;
	v61 = int64(9223372036854775807)
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[3])) = v61
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[4])))
	if v65 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+96)) = int32(1)
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v89 = int32(1)
	goto L19
L19:
	;
	if v89&int32(1) == int32(0) {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	F_s_lock(m, v74+int32(96), int32(_a_F_XLogWalRcvSendHSFeedback_0), int32(_a_F_XLogWalRcvSendHSFeedback_1), int32(_a_F_XLogWalRcvSendHSFeedback_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[5]))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[4])) = uint8(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+96)) = int32(0)
	v89 = v85
	goto L19
L23:
	;
	return
L24:
	;
	goto L22
L25:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[0])))
	if v96 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v124 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v103 = m.G0
	v105 = v103 - int32(48)
	m.G0 = v105
	F_ComputeXidHorizons(m, v105+int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v118
	goto L26
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(28)))) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v113
	m.G0 = v105 + int32(48)
	goto L26
L31:
	;
	v128 = base.I32_wrap_i64(int64(base.Ui64(v124) >> (uint(int64(32)) % 64)))
	v130 = v128 - int32(1)
	v131 = base.I32_wrap_i64(v124)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v131) < base.Ui32(v132) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v134 = v130
	goto L34
L33:
	;
	v134 = v128
	goto L34
L34:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if base.Ui32(v131) < base.Ui32(v135) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v137 = v130
	goto L37
L36:
	;
	v137 = v128
	goto L37
L37:
	;
	v140 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	if v140 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v137
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v134
	F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendHSFeedback_3), v10)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L23
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v156 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v158)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[7])) = v158
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v158
	goto L44
L42:
	;
	F_errfinish(m, int32(_a_F_XLogWalRcvSendHSFeedback_4), int32(1233), int32(_a_F_XLogWalRcvSendHSFeedback_5))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_6), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v170 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v173 = int32(104)
	*(*uint8)(unsafe.Add(mBase, uint32(v169+v171))) = uint8(v173)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v177 + int32(1)
	v184 = m.G0
	v185 = int32(16)
	v186 = v184 - v185
	m.G0 = v186
	F_gettimeofday(m, v186)
	mBase = m.M
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v186)+8)))
	m.G0 = v186 + v185
	v198 = v190 + v189*int64(1000000) - int64(946684800000000)
	goto L46
L46:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendHSFeedback_6), int32(8))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	v203 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v205 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v208 = int64(56)
	v210 = int64(65280)
	v212 = int64(40)
	v215 = int64(16711680)
	v217 = int64(24)
	v219 = int64(4278190080)
	v221 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v204+v206))) = v198<<(uint(v208)%64) | v198&v210<<(uint(v212)%64) | (v198&v215<<(uint(v217)%64) | v198&v219<<(uint(v221)%64)) | (int64(base.Ui64(v198)>>(uint(v221)%64))&v219 | int64(base.Ui64(v198)>>(uint(v217)%64))&v215 | (int64(base.Ui64(v198)>>(uint(v212)%64))&v210 | int64(base.Ui64(v198)>>(uint(v208)%64))))
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v246 + int32(8)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_enlargeStringInfo(m, v203, int32(4))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	v255 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v257 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v262 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v256+v258))) = base.I32_rotr(v250, int32(24))&v262 | base.I32_rotr(v250&v262, int32(8))
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v273 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v272 + v273
	F_enlargeStringInfo(m, v255, v273)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	v280 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v282 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v287 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v281+v283))) = base.I32_rotr(v137, int32(24))&v287 | base.I32_rotr(v137&v287, int32(8))
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v298 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v297 + v298
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_enlargeStringInfo(m, v280, v298)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	v306 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v308 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v313 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v307+v309))) = base.I32_rotr(v301, int32(24))&v313 | base.I32_rotr(v301&v313, int32(8))
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v324 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v323 + v324
	F_enlargeStringInfo(m, v306, v324)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v331 = int32(_a_F_XLogWalRcvSendHSFeedback_6)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v333 = int32(_a_F_XLogWalRcvSendHSFeedback_7)
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v338 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v332+v334))) = base.I32_rotr(v134, int32(24))&v338 | base.I32_rotr(v134&v338, int32(8))
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8]))
	v350 = v348 + int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[8])) = v350
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[9]))
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[6]))
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[10]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+44))
	m.T0[v358].(func(*base.Module, int32, int32, int32))(m, v353, v355, v350)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWalRcvSendHSFeedback[2])) = uint8(base.B2i32(v362|v363 == int32(0)))
	goto L1
}
