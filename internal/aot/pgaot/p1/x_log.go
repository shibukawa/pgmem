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
	v4 = *(*int64)(unsafe.Add(mBase, _consts[260]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[261]))
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[259]))
		F_s_lock(m, v11+int32(96), int32(513789), int32(4586), int32(216592))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[259]))
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
		v22 = *(*int32)(unsafe.Add(mBase, _consts[259]))
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
	var v103 int64
	_ = v103
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v121 int64
	_ = v121
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
	var v206 int32
	_ = v206
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
	var v245 int32
	_ = v245
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
	var v312 int64
	_ = v312
	var v323 int64
	_ = v323
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v329 int64
	_ = v329
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+264))
	if base.Ui64(v1) <= base.Ui64(v23) {
		v329 = v23
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
	return v329
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
	F_s_lock(m, v22, int32(520929), int32(1528), int32(337063))
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
	v39 = int64(*(*int32)(unsafe.Add(mBase, _consts[221])))
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
	v72 = int64(*(*int32)(unsafe.Add(mBase, _consts[138])))
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
	v103 = v76
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
	F_errmsg(m, int32(538946), v17)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(520929), int32(1545), int32(337063))
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
	v121 = v115
	goto L30
L28:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[30]))
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
		v103 = v299
		v113 = v301
		goto L27
	} else {
		goto L64
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v136 = v135 + v113<<(uint(int32(7))%32)
	v138 = v136 + int32(16)
	v140 = v17 + int32(24)
	v141 = int32(0)
	v142 = int32(4548540)
	v144 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v144 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v150&int32(262144) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v292 == int64(0) {
		v299 = v103
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
	v245 = int32(1)
	goto L32
L36:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v167
	if v167 != v121 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v227 = v206
	goto L35
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140))) = v167
	v242 = v158
	v245 = int32(0)
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
	v181 = v177 & int32(262144)
	if v181 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136))))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193 | int32(16777216)
	v206 = v158
	goto L48
L43:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v182
	if v182 == v121 {
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
	v245 = base.B2i32(v181 == int32(0))
	goto L32
L48:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+74)))
	if v214 != 0 {
		v206 = v206 + int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v219&int32(262144) != 0 {
		v158 = v206
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
	v286 = int32(4548540)
	v288 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v288 - int32(1)
	if v245 != 0 {
		v299 = v103
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
		v121 = v292
		goto L30
	} else {
		goto L59
	}
L59:
	;
	goto L31
L60:
	;
	if base.Ui64(v292) < base.Ui64(v103) {
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
	v297 = v103
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
	v329 = v306
	goto L4
L66:
	;
	goto L67
L67:
	;
	v312 = v306
	goto L68
L68:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v305)+264))
	v324 = base.B2i32(v323 == v312)
	if v323 == v312 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v329 = v323
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
	if v323 == v312 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v329 = v299
	goto L4
L74:
	;
	goto L75
L75:
	;
	if base.Ui64(v323) < base.Ui64(v299) {
		v312 = v323
		goto L68
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	F_errmsg_internal(m, int32(379419), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(520929), int32(1517), int32(337063))
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v3 = m.G0
	v5 = v3 - int32(2208)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = int32(389613)
	v16 = F_pg_snprintf(m, v5+int32(160), int32(1024), int32(185541), v5+int32(48))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v24 = F___fstatat(m, int32(-100), v5+int32(160), v5-int32(-64), int32(0))
		mBase = m.M
		if v24 == int32(0) {
			m.G0 = v5 + int32(2208)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(23166)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = l0
			v36 = F_pg_snprintf(m, v5+int32(1184), int32(1024), int32(185541), v5+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v44 = F___fstatat(m, int32(-100), v5+int32(1184), v5-int32(-64), int32(0))
				mBase = m.M
				if v44 == int32(0) {
					v52 = F_durable_rename(m, v5+int32(1184), v5+int32(160), int32(19))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v5 + int32(2208)
						return
					}
				} else {
					v57 = F_AllocateFile(m, v5+int32(160), int32(33848))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						if v57 == int32(0) {
							v63 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								if v63 == int32(0) {
									m.G0 = v5 + int32(2208)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(160)
										F_errmsg(m, int32(310905), v5)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_errfinish(m, int32(521237), int32(537), int32(389690))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												m.G0 = v5 + int32(2208)
												return
											}
										}
									}
								}
							}
						} else {
							v80 = F_FreeFile(m, v57)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								if v80 == int32(0) {
									m.G0 = v5 + int32(2208)
									return
								} else {
									v86 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 == int32(0) {
											m.G0 = v5 + int32(2208)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v5 + int32(160)
												F_errmsg(m, int32(310860), v5+int32(16))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													F_errfinish(m, int32(521237), int32(545), int32(389690))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														m.G0 = v5 + int32(2208)
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(1184)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(389613)
	v17 = F_pg_snprintf(m, v6+int32(160), int32(1024), int32(185541), v6+int32(48))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v27 = F___fstatat(m, int32(-100), v6+int32(160), v6-int32(-64), int32(0))
		mBase = m.M
		if v27 == int32(0) {
			v91 = v2
			m.G0 = v6 + int32(1184)
			return v91
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(23166)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
			v39 = F_pg_snprintf(m, v6+int32(160), int32(1024), int32(185541), v6+int32(32))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v47 = F___fstatat(m, int32(-100), v6+int32(160), v6-int32(-64), int32(0))
				mBase = m.M
				if v47 == int32(0) {
					v91 = int32(1)
					m.G0 = v6 + int32(1184)
					return v91
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(389613)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
					v59 = F_pg_snprintf(m, v6+int32(160), int32(1024), int32(185541), v6+int32(16))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v67 = F___fstatat(m, int32(-100), v6+int32(160), v6-int32(-64), int32(0))
						mBase = m.M
						if v67 == int32(0) {
							v91 = v2
							m.G0 = v6 + int32(1184)
							return v91
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
							v75 = F_pg_snprintf(m, v6+int32(160), int32(1024), int32(186944), v6)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v83 = F___fstatat(m, int32(-100), v6+int32(160), v6-int32(-64), int32(0))
								mBase = m.M
								if v83 == int32(0) {
									v91 = int32(1)
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, _consts[137]))
									if v87 == int32(44) {
										v91 = v2
									} else {
										v91 = int32(1)
									}
								}
								m.G0 = v6 + int32(1184)
								return v91
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
	v4 = int64(*(*int32)(unsafe.Add(mBase, _consts[221])))
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
	v27 = int64(*(*int32)(unsafe.Add(mBase, _consts[138])))
	return v5*v27 + v25&int64(4294967295)
}
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	if l0 < int32(33) {
		v10 = int32(20)
		if l1 <= v10 {
			v13 = v10
		} else {
			v13 = l1
		}
		v14 = int32(4)
		if l0 <= v14 {
			v17 = v14
		} else {
			v17 = l0
		}
		v19 = *(*int32)(unsafe.Add(mBase, _consts[254]))
		if v19 <= v17 {
			v21 = int32(4449700)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[255]))
			v25 = v17 + int32(1)
			v28 = F_repalloc(m, v23, v25*int32(8260))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[255])) = v28
				v32 = *(*int32)(unsafe.Add(mBase, _consts[254]))
				v34 = int32(8260)
				v35 = (v25 - v32) * v34
				v38 = v28 + v32*v34
				if v38&int32(3) != 0 {
					v70 = F__emscripten_memset_bulkmem(m, v38, base.I32_extend8_s(int32(0)), v35)
					mBase = m.M
				} else {
					if base.Ui32(int32(1024)) < base.Ui32(v35) {
						v70 = F__emscripten_memset_bulkmem(m, v38, base.I32_extend8_s(int32(0)), v35)
						mBase = m.M
					} else {
						if base.Ui32(v35+v38) <= base.Ui32(v38) {
						} else {
							v48 = int32(8260)
							v52 = v17*v48 + v28 + v48
							v54 = v32 * v48
							v57 = v54 + v28 + int32(4)
							if base.Ui32(v57) < base.Ui32(v52) {
								v59 = v52
							} else {
								v59 = v57
							}
							v67 = F__emscripten_memset_bulkmem(m, v38, base.I32_extend8_s(int32(0)), (v28^int32(-1)+v59-v54)&int32(-4)+int32(4))
							mBase = m.M
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, _consts[254])) = v25
				v83 = *(*int32)(unsafe.Add(mBase, _consts[256]))
				if v83 < v13 {
					v86 = *(*int32)(unsafe.Add(mBase, _consts[257]))
					v89 = F_repalloc(m, v86, v13*int32(12))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[256])) = v13
						*(*int32)(unsafe.Add(mBase, _consts[257])) = v89
						return
					}
				} else {
					return
				}
			}
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, _consts[256]))
			if v83 < v13 {
				v86 = *(*int32)(unsafe.Add(mBase, _consts[257]))
				v89 = F_repalloc(m, v86, v13*int32(12))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[256])) = v13
					*(*int32)(unsafe.Add(mBase, _consts[257])) = v89
					return
				}
			} else {
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(482127), int32(0))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				F_errfinish(m, int32(514616), int32(194), int32(438629))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
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
	v23 = F_pg_snprintf(m, l0, int32(64), int32(534000), v10)
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
	var v23 int64
	_ = v23
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
	var v105 int32
	_ = v105
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
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	v5 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_AllocateDir(m, int32(322229))
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
	v17 = F_ReadDir(m, v12, int32(322229))
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
	v23 = v5
	goto L7
L5:
	;
	v151 = v5
	goto L6
L6:
	;
	F_FreeDir(m, v12)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v26 = v21 + int32(19)
	v27 = F_strlen(m, v26)
	mBase = m.M
	if v27 != int32(24) {
		v142 = v23
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v151 = v142
	goto L6
L9:
	;
	v145 = F_ReadDir(m, v12, int32(322229))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L38
	}
L10:
	;
	v30 = int32(561281)
	v34 = m.G0
	v36 = v34 - int32(32)
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[229])))
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v113 != int32(24) {
		v142 = v23
		goto L9
	} else {
		goto L32
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
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
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
		v105 = v26
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v113 = v105 - v26
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
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v105 = v101
	goto L24
L28:
	;
	v105 = v84
	goto L24
L29:
	;
	goto L30
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v101 = v84 + int32(1)
	if v99 != 0 {
		v84 = v101
		v85 = v99
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v117 = int64(*(*int32)(unsafe.Add(mBase, _consts[138])))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(24)
	v128 = F_sscanf(m, v26, int32(534000), v9)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if l0 != v130 {
		v142 = v23
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)))
	v135 = base.I64_div_u_s(int64(4294967296), v117)
	v137 = v132 + v133*v135
	if base.Ui64(v23-int64(1)) < base.Ui64(v137) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v141 = v23
	goto L37
L36:
	;
	v141 = v137
	goto L37
L37:
	;
	v142 = v141
	goto L9
L38:
	;
	if v145 != 0 {
		v21 = v145
		v23 = v142
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	m.G0 = v9 + int32(32)
	return v151
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
	var v2 int32
	_ = v2
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	v2 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v6 != int32(1) {
		v70 = *(*int64)(unsafe.Add(mBase, _consts[225]))
		if base.Ui64(l0) <= base.Ui64(v70) {
			v87 = v2
		} else {
			v72 = int32(4449456)
			v73 = *(*int32)(unsafe.Add(mBase, _consts[30]))
			v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+280))
			*(*int64)(unsafe.Add(mBase, uint32(v73)+280)) = v74
			v76 = int32(4449496)
			*(*int64)(unsafe.Add(mBase, _consts[225])) = v74
			v79 = *(*int32)(unsafe.Add(mBase, _consts[30]))
			v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+272))
			*(*int64)(unsafe.Add(mBase, uint32(v79)+272)) = v80
			*(*int64)(unsafe.Add(mBase, _consts[226])) = v80
			v85 = *(*int64)(unsafe.Add(mBase, _consts[225]))
			v87 = base.B2i32(base.Ui64(v85) < base.Ui64(l0))
		}
		return v87 & int32(1)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[30]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+316))
		v13 = int32(2)
		*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(base.B2i32(v12 != v13))
		if v12 == v13 {
			v70 = *(*int64)(unsafe.Add(mBase, _consts[225]))
			if base.Ui64(l0) <= base.Ui64(v70) {
				v87 = v2
			} else {
				v72 = int32(4449456)
				v73 = *(*int32)(unsafe.Add(mBase, _consts[30]))
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+280))
				*(*int64)(unsafe.Add(mBase, uint32(v73)+280)) = v74
				v76 = int32(4449496)
				*(*int64)(unsafe.Add(mBase, _consts[225])) = v74
				v79 = *(*int32)(unsafe.Add(mBase, _consts[30]))
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+272))
				*(*int64)(unsafe.Add(mBase, uint32(v79)+272)) = v80
				*(*int64)(unsafe.Add(mBase, _consts[226])) = v80
				v85 = *(*int64)(unsafe.Add(mBase, _consts[225]))
				v87 = base.B2i32(base.Ui64(v85) < base.Ui64(l0))
			}
			return v87 & int32(1)
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, _consts[227]))
			if v19 != int64(0) {
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
				if v23 != int32(1) {
				} else {
					v27 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[228])) = uint8(v27)
				}
			}
			if base.Ui64(l0) <= base.Ui64(v19) {
				v87 = v2
				return v87 & int32(1)
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[228])))
				if v31 != 0 {
					v87 = v2
					return v87 & int32(1)
				} else {
					v32 = int32(1)
					v34 = *(*int32)(unsafe.Add(mBase, _consts[24]))
					v38 = F_LWLockConditionalAcquire(m, v34+int32(1152), v32)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v38 == int32(0) {
							v87 = v32
							return v87 & int32(1)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, _consts[109]))
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+136))
							*(*int64)(unsafe.Add(mBase, _consts[227])) = v47
							v50 = *(*int32)(unsafe.Add(mBase, _consts[24]))
							F_LWLockRelease(m, v50+int32(1152))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v56 = *(*int64)(unsafe.Add(mBase, _consts[227]))
								if v56 == int64(0) {
									v60 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[228])) = uint8(v60)
								} else {
								}
								if base.Ui64(l0) <= base.Ui64(v56) {
									v87 = int32(0)
								} else {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[228])))
									v87 = v65 ^ int32(1)
								}
								return v87 & int32(1)
							}
						}
					}
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
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int64
	_ = v293
	var v314 int32
	_ = v314
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v384 int64
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(112)
	return v453
L2:
	;
	if v407 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L3:
	;
	v338 = l0 + int32(8)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v339
	v341 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v341
	v346 = F_smgropen(m, v18+int32(72), int32(-1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L84
	}
L4:
	;
	if l4 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v24
	v27 = v18 + int32(88)
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if l4 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v30 + int32(32)
	if v314 != 0 {
		v407 = l4
		goto L2
	} else {
		goto L83
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v51 = v46 + (l4^int32(-1))<<(uint(int32(6))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	if v52&int32(16777216) == int32(0) {
		v314 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = l4
	v82 = v78 + l4<<(uint(int32(6))%32)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	if l4 == v84 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v42 != v57 {
		v314 = v6
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v41 != v59 {
		v314 = v6
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v40 != v61 {
		v314 = v6
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if l2 != v63 {
		v314 = v6
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if l1 != v65 {
		v314 = v6
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(1)
	F_PinLocalBuffer(m, v51, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v71 = int32(4452184)
	v73 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v73 + int64(1)
	v314 = v67
	goto L9
L20:
	;
	if v226&int32(16777216) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = int32(240134)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(517496)
	v145 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v145
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = int64(0)
	v151 = v82 - int32(40)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v152 | v153
	if v152&v153 != 0 {
		goto L51
	} else {
		goto L52
	}
L22:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v131 <= int32(0) {
		goto L21
	} else {
		goto L50
	}
L23:
	;
	v129 = int32(4469728)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	if l4 == v88 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v129 = int32(4469736)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	if l4 == v92 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v129 = int32(4469744)
	goto L22
L30:
	;
	goto L31
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if l4 == v96 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v129 = int32(4469752)
	goto L22
L33:
	;
	goto L34
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	if l4 == v100 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v129 = int32(4469760)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if l4 == v104 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v129 = int32(4469768)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	if l4 == v108 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = int32(4469776)
	goto L22
L42:
	;
	goto L43
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if l4 == v112 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v129 = int32(4469784)
	goto L22
L45:
	;
	goto L46
L46:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	if v116 == int32(0) {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v123 = int32(0)
	v125 = F_hash_search(m, v120, v30+int32(8), v123, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v125 == int32(0) {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v129 = v125
	goto L22
L50:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(40))))
	v226 = v136
	v232 = int32(1)
	goto L20
L51:
	;
	goto L54
L52:
	;
	v188 = v152
	goto L53
L53:
	;
	v201 = int32(4160012)
	v202 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(8))+8))
	if v204 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	F_perform_spin_delay(m, v30+int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L56
	}
L55:
	;
	v188 = v177
	goto L53
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v178 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v177 | v178
	if v177&v178 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v226 = v188
	v232 = v145
	goto L20
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v219
	goto L59
L61:
	;
	if int32(999) < v202 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v202 < int32(11) {
		goto L59
	} else {
		goto L68
	}
L64:
	;
	v209 = int32(900)
	if v209 <= v202 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v212 = v209
	goto L67
L66:
	;
	v212 = v202
	goto L67
L67:
	;
	v219 = v212 + int32(100)
	goto L60
L68:
	;
	v219 = v202 - int32(1)
	goto L60
L69:
	;
	if v232 != 0 {
		v314 = v6
		goto L9
	} else {
		goto L82
	}
L70:
	;
	v241 = v82 + int32(-64)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v42 != v242 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(60))))
	if v41 != v246 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(56))))
	if v40 != v250 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(48))))
	if l2 != v254 {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(52))))
	if l1 != v258 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	if v232 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v291 = int32(4452152)
	v293 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v293 + int64(1)
	v314 = int32(1)
	goto L9
L77:
	;
	v261 = F_PinBuffer(m, v241, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v264 = v82 - int32(40)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = (v265 + v266) & int32(-4194305)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(44))))
	v274 = int32(4469800)
	v275 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	*(*int32)(unsafe.Add(mBase, _consts[278])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v266
	v282 = v273 + v266
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v282
	v285 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v285, v282, int32(1660384))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v82-int32(40)))) = v226 & int32(-4194305)
	v314 = v6
	goto L9
L83:
	;
	goto L3
L84:
	;
	F_smgrcreate(m, v346, l1, int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	v351 = F_smgrnblocks(m, v346, l1)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(v351) <= base.Ui32(l2) {
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
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v393
	v395 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v395
	v401 = F_ReadBufferWithoutRelcache(m, v18+int32(24), l1, l2, l3, int32(0), int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L6
	} else {
		goto L96
	}
L90:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v358
	v360 = int32(0)
	F_log_invalid_page(m, v18+int32(40), l1, l2, v360)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
		v453 = int32(0)
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v453 = v360
	goto L1
L94:
	;
	v370 = v18 + int32(108)
	v371 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v370))) = uint8(v371)
	v373 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v373)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+109)) = uint16(v373)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	*(*int32)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v373
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v384
	v391 = F_ExtendBufferedRelTo(m, v18+int32(56), l1, int32(3), l2+int32(1), l3)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v453 = v391
	goto L1
L96:
	;
	if l3 != 0 {
		v453 = v401
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v407 = v401
	goto L2
L98:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+14)))
	if v436 != 0 {
		v453 = v407
		goto L1
	} else {
		goto L102
	}
L99:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v421+(v407^int32(-1))<<(uint(int32(2))%32))))
	v435 = v427
	goto L98
L100:
	;
	goto L101
L101:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v435 = v429 + v407<<(uint(int32(13))%32) + int32(-8192)
	goto L98
L102:
	;
	F_ReleaseBuffer(m, v407)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v441
	F_log_invalid_page(m, v18+int32(8), l1, l2, int32(1))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v453 = int32(0)
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v5 <= l0 {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[254]))
		if v8 <= l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(143308), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errfinish(m, int32(514616), int32(267), int32(237039))
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
			*(*int32)(unsafe.Add(mBase, _consts[258])) = l0 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[255]))
			v18 = v15 + l0*int32(8260)
			v20 = v18 + int32(4)
			if l1 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
				v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
				v68 = v60
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
		v15 = *(*int32)(unsafe.Add(mBase, _consts[255]))
		v18 = v15 + l0*int32(8260)
		v20 = v18 + int32(4)
		if l1 < int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[8]))
			v42 = v29 + (l1^int32(-1))<<(uint(int32(6))%32)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
			v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l1^int32(-1))<<(uint(int32(2))%32))))
			v68 = v60
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[253]))
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
	F_s_lock(m, v10, int32(516797), int32(190), int32(38183))
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
	v60 = *(*int32)(unsafe.Add(mBase, _consts[24]))
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
	v67 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+320)) = uint8(v68)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[24]))
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[683]))
	v16 = *(*int64)(unsafe.Add(mBase, _consts[684]))
	if base.Ui64(v16) <= base.Ui64(v14) {
		m.G0 = v11 + int32(80)
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[253]))
		v21 = *(*int32)(unsafe.Add(mBase, _consts[685]))
		v23 = *(*int64)(unsafe.Add(mBase, _consts[686]))
		F_issue_xlog_fsync(m, v21, v23, l1)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v28 = *(*int64)(unsafe.Add(mBase, _consts[684]))
			*(*int64)(unsafe.Add(mBase, _consts[683])) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1456))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+1456)) = int32(1)
			v34 = v19 + int32(1456)
			if v30 != 0 {
				F_s_lock(m, v34, int32(517549), int32(998), int32(336696))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
					v42 = *(*int64)(unsafe.Add(mBase, _consts[683]))
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
						v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[578])))
						if v52 != int32(1) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
							if v64 == int32(1) {
								v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
								v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
								v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v81 = F_strlen(m, v11+int32(16))
									mBase = m.M
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v83 = int32(0)
										F_XLogWalRcvSendReply(m, v83, v83)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
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
									v83 = int32(0)
									F_XLogWalRcvSendReply(m, v83, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									}
								}
							}
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, _consts[532]))
							if v56 <= int32(0) {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
								if v64 == int32(1) {
									v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
									v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
									v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v81 = F_strlen(m, v11+int32(16))
										mBase = m.M
										if l0 != 0 {
											m.G0 = v11 + int32(80)
											return
										} else {
											v83 = int32(0)
											F_XLogWalRcvSendReply(m, v83, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
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
										v83 = int32(0)
										F_XLogWalRcvSendReply(m, v83, v83)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
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
									v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
									if v64 == int32(1) {
										v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
										v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
										v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v81 = F_strlen(m, v11+int32(16))
											mBase = m.M
											if l0 != 0 {
												m.G0 = v11 + int32(80)
												return
											} else {
												v83 = int32(0)
												F_XLogWalRcvSendReply(m, v83, v83)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													F_XLogWalRcvSendHSFeedback(m, int32(0))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
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
											v83 = int32(0)
											F_XLogWalRcvSendReply(m, v83, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
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
				v42 = *(*int64)(unsafe.Add(mBase, _consts[683]))
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
					v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[578])))
					if v52 != int32(1) {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
						if v64 == int32(1) {
							v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
							v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
							v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v81 = F_strlen(m, v11+int32(16))
								mBase = m.M
								if l0 != 0 {
									m.G0 = v11 + int32(80)
									return
								} else {
									v83 = int32(0)
									F_XLogWalRcvSendReply(m, v83, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
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
								v83 = int32(0)
								F_XLogWalRcvSendReply(m, v83, v83)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_XLogWalRcvSendHSFeedback(m, int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _consts[532]))
						if v56 <= int32(0) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
							if v64 == int32(1) {
								v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
								v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
								v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v81 = F_strlen(m, v11+int32(16))
									mBase = m.M
									if l0 != 0 {
										m.G0 = v11 + int32(80)
										return
									} else {
										v83 = int32(0)
										F_XLogWalRcvSendReply(m, v83, v83)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
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
									v83 = int32(0)
									F_XLogWalRcvSendReply(m, v83, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_XLogWalRcvSendHSFeedback(m, int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
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
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[687])))
								if v64 == int32(1) {
									v68 = *(*int64)(unsafe.Add(mBase, _consts[684]))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v68)
									v71 = int64(base.Ui64(v68) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v71)
									v77 = F_pg_snprintf(m, v11+int32(16), int32(50), int32(539330), v11)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v81 = F_strlen(m, v11+int32(16))
										mBase = m.M
										if l0 != 0 {
											m.G0 = v11 + int32(80)
											return
										} else {
											v83 = int32(0)
											F_XLogWalRcvSendReply(m, v83, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_XLogWalRcvSendHSFeedback(m, int32(0))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
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
										v83 = int32(0)
										F_XLogWalRcvSendReply(m, v83, v83)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_XLogWalRcvSendHSFeedback(m, int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v36 int64
	_ = v36
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v193 int64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[676]))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v22 = m.G0
	v23 = int32(16)
	v24 = v22 - v23
	m.G0 = v24
	F___gettimeofday(m, v24)
	mBase = m.M
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
	m.G0 = v24 + v23
	v36 = v28 + v27*int64(1000000) - int64(946684800000000)
	goto L8
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	if v16 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[682])))
	if v18 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	goto L2
L8:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = *(*int64)(unsafe.Add(mBase, _consts[678]))
	if v36 < v40 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v45 = int64(*(*int32)(unsafe.Add(mBase, _consts[676])))
	if v45 <= int64(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v51 = int64(9223372036854775807)
	goto L15
L14:
	;
	v51 = v36 + v45*int64(1000000)
	goto L15
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = v51
	goto L18
L17:
	;
	v55 = int64(9223372036854775807)
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, _consts[678])) = v55
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _consts[262])))
	if v59 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+96)) = int32(1)
	if v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v84 = int32(1)
	goto L21
L21:
	;
	if v84&int32(1) == int32(0) {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	F_s_lock(m, v68+int32(96), int32(513789), int32(4556), int32(359104))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	*(*uint8)(unsafe.Add(mBase, _consts[262])) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = int32(0)
	v84 = v79
	goto L21
L25:
	;
	return
L26:
	;
	goto L24
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	if v90 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v119 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L25
	} else {
		goto L33
	}
L29:
	;
	v97 = m.G0
	v99 = v97 - int32(48)
	m.G0 = v99
	F_ComputeXidHorizons(m, v99+int32(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v112
	goto L28
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(28)))) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v107
	m.G0 = v99 + int32(48)
	goto L28
L33:
	;
	v123 = base.I32_wrap_i64(int64(base.Ui64(v119) >> (uint(int64(32)) % 64)))
	v125 = v123 - int32(1)
	v126 = base.I32_wrap_i64(v119)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v126) < base.Ui32(v127) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = v125
	goto L36
L35:
	;
	v129 = v123
	goto L36
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v126) < base.Ui32(v130) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v132 = v125
	goto L39
L38:
	;
	v132 = v123
	goto L39
L39:
	;
	v135 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	if v135 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v132
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v129
	F_errmsg_internal(m, int32(52530), v9)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L25
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v151 = int32(4464192)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v153)
	*(*int32)(unsafe.Add(mBase, _consts[680])) = v153
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v153
	goto L46
L44:
	;
	F_errfinish(m, int32(517549), int32(1233), int32(333932))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_enlargeStringInfo(m, int32(4464192), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v165 = int32(4464196)
	v166 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v168 = int32(104)
	*(*uint8)(unsafe.Add(mBase, uint32(v164+v166))) = uint8(v168)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v172 + int32(1)
	v179 = m.G0
	v180 = int32(16)
	v181 = v179 - v180
	m.G0 = v181
	F___gettimeofday(m, v181)
	mBase = m.M
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
	m.G0 = v181 + v180
	v193 = v185 + v184*int64(1000000) - int64(946684800000000)
	goto L48
L48:
	;
	F_enlargeStringInfo(m, int32(4464192), int32(8))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	v198 = int32(4464192)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v200 = int32(4464196)
	v201 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v203 = int64(56)
	v205 = int64(65280)
	v207 = int64(40)
	v210 = int64(16711680)
	v212 = int64(24)
	v214 = int64(4278190080)
	v216 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v199+v201))) = v193<<(uint(v203)%64) | v193&v205<<(uint(v207)%64) | (v193&v210<<(uint(v212)%64) | v193&v214<<(uint(v216)%64)) | (int64(base.Ui64(v193)>>(uint(v216)%64))&v214 | int64(base.Ui64(v193)>>(uint(v212)%64))&v210 | (int64(base.Ui64(v193)>>(uint(v207)%64))&v205 | int64(base.Ui64(v193)>>(uint(v203)%64))))
	v241 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v241 + int32(8)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_enlargeStringInfo(m, v198, int32(4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v250 = int32(4464192)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v252 = int32(4464196)
	v253 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v255 = int32(24)
	v257 = int32(65280)
	v259 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v253))) = v245<<(uint(v255)%32) | v245&v257<<(uint(v259)%32) | (int32(base.Ui32(v245)>>(uint(v259)%32))&v257 | int32(base.Ui32(v245)>>(uint(v255)%32)))
	v273 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v274 = int32(4)
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v273 + v274
	F_enlargeStringInfo(m, v250, v274)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	v281 = int32(4464192)
	v282 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v283 = int32(4464196)
	v284 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v286 = int32(24)
	v288 = int32(65280)
	v290 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v282+v284))) = v132<<(uint(v286)%32) | v132&v288<<(uint(v290)%32) | (int32(base.Ui32(v132)>>(uint(v290)%32))&v288 | int32(base.Ui32(v132)>>(uint(v286)%32)))
	v304 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v305 = int32(4)
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v304 + v305
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	F_enlargeStringInfo(m, v281, v305)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	v313 = int32(4464192)
	v314 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v315 = int32(4464196)
	v316 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v318 = int32(24)
	v320 = int32(65280)
	v322 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v314+v316))) = v308<<(uint(v318)%32) | v308&v320<<(uint(v322)%32) | (int32(base.Ui32(v308)>>(uint(v322)%32))&v320 | int32(base.Ui32(v308)>>(uint(v318)%32)))
	v336 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v337 = int32(4)
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v336 + v337
	F_enlargeStringInfo(m, v313, v337)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v344 = int32(4464192)
	v345 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v346 = int32(4464196)
	v347 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v349 = int32(24)
	v351 = int32(65280)
	v353 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v345+v347))) = v129<<(uint(v349)%32) | v129&v351<<(uint(v353)%32) | (int32(base.Ui32(v129)>>(uint(v353)%32))&v351 | int32(base.Ui32(v129)>>(uint(v349)%32)))
	v367 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	v369 = v367 + int32(4)
	*(*int32)(unsafe.Add(mBase, _consts[681])) = v369
	v372 = *(*int32)(unsafe.Add(mBase, _consts[675]))
	v374 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v376 = *(*int32)(unsafe.Add(mBase, _consts[659]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+44))
	m.T0[v377].(func(*base.Module, int32, int32, int32))(m, v372, v374, v369)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*uint8)(unsafe.Add(mBase, _consts[682])) = uint8(base.B2i32(v381|v382 == int32(0)))
	goto L1
}
