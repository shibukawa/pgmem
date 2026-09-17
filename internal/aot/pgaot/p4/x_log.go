package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XLogArchiveCheckDone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(1168)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveCheckDone[0]))
	if v11 <= int32(0) {
		v97 = int32(1)
		m.G0 = v8 + int32(1168)
		return v97
	} else {
		if v11 != int32(2) {
			v17 = int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveCheckDone[1]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+440))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+440)) = v17
			if v20 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveCheckDone[1]))
				F_s_lock(m, v24+int32(440), int32(_a_F_XLogArchiveCheckDone_0), int32(_a_F_XLogArchiveCheckDone_1), int32(_a_F_XLogArchiveCheckDone_2))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveCheckDone[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+440)) = int32(0)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+316))
					if v38 == int32(1) {
						v97 = v17
						m.G0 = v8 + int32(1168)
						return v97
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_XLogArchiveCheckDone_3)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
						v47 = v8 + int32(144)
						v52 = F_pg_snprintf(m, v47, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(32))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v58 = F___fstatat(m, int32(-100), v47, v8+int32(48), int32(0))
							mBase = m.M
							if v58 == int32(0) {
								v97 = int32(1)
								m.G0 = v8 + int32(1168)
								return v97
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_XLogArchiveCheckDone_5)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
								v66 = v8 + int32(144)
								v71 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(16))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = int32(0)
									v75 = v8 + int32(48)
									v78 = F___fstatat(m, int32(-100), v66, v75, v73)
									mBase = m.M
									if v78 == int32(0) {
										v97 = v73
										m.G0 = v8 + int32(1168)
										return v97
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_XLogArchiveCheckDone_3)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										v86 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											v91 = F___fstatat(m, int32(-100), v66, v75, int32(0))
											mBase = m.M
											if v91 == int32(0) {
												v97 = int32(1)
												m.G0 = v8 + int32(1168)
												return v97
											} else {
												F_XLogArchiveNotify(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													v97 = int32(0)
													m.G0 = v8 + int32(1168)
													return v97
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
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveCheckDone[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+440)) = int32(0)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+316))
				if v38 == int32(1) {
					v97 = v17
					m.G0 = v8 + int32(1168)
					return v97
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_XLogArchiveCheckDone_3)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
					v47 = v8 + int32(144)
					v52 = F_pg_snprintf(m, v47, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(32))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v58 = F___fstatat(m, int32(-100), v47, v8+int32(48), int32(0))
						mBase = m.M
						if v58 == int32(0) {
							v97 = int32(1)
							m.G0 = v8 + int32(1168)
							return v97
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_XLogArchiveCheckDone_5)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							v66 = v8 + int32(144)
							v71 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(16))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = int32(0)
								v75 = v8 + int32(48)
								v78 = F___fstatat(m, int32(-100), v66, v75, v73)
								mBase = m.M
								if v78 == int32(0) {
									v97 = v73
									m.G0 = v8 + int32(1168)
									return v97
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_XLogArchiveCheckDone_3)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									v86 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v91 = F___fstatat(m, int32(-100), v66, v75, int32(0))
										mBase = m.M
										if v91 == int32(0) {
											v97 = int32(1)
											m.G0 = v8 + int32(1168)
											return v97
										} else {
											F_XLogArchiveNotify(m, l0)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												v97 = int32(0)
												m.G0 = v8 + int32(1168)
												return v97
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_XLogArchiveCheckDone_3)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
			v47 = v8 + int32(144)
			v52 = F_pg_snprintf(m, v47, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(32))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v58 = F___fstatat(m, int32(-100), v47, v8+int32(48), int32(0))
				mBase = m.M
				if v58 == int32(0) {
					v97 = int32(1)
					m.G0 = v8 + int32(1168)
					return v97
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_XLogArchiveCheckDone_5)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v66 = v8 + int32(144)
					v71 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8+int32(16))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = int32(0)
						v75 = v8 + int32(48)
						v78 = F___fstatat(m, int32(-100), v66, v75, v73)
						mBase = m.M
						if v78 == int32(0) {
							v97 = v73
							m.G0 = v8 + int32(1168)
							return v97
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_XLogArchiveCheckDone_3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v86 = F_pg_snprintf(m, v66, int32(1024), int32(_a_F_XLogArchiveCheckDone_4), v8)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v91 = F___fstatat(m, int32(-100), v66, v75, int32(0))
								mBase = m.M
								if v91 == int32(0) {
									v97 = int32(1)
									m.G0 = v8 + int32(1168)
									return v97
								} else {
									F_XLogArchiveNotify(m, l0)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v97 = int32(0)
										m.G0 = v8 + int32(1168)
										return v97
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
func F_XLogBackgroundFlush(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[1])))
	if v17 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v252
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
	v23 = base.B2i32(v21 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[1])) = uint8(v23)
	if v21 != int32(2) {
		v252 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+440)) = int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+308))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	F_s_lock(m, v31+int32(440), int32(_a_F_XLogBackgroundFlush_0), int32(2990), int32(_a_F_XLogBackgroundFlush_1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v41 = int32(_a_F_XLogBackgroundFlush_2)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v42)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v42)+440)) = int32(0)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v49 & int64(-8192)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v42)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+280)) = v53
	v55 = int32(_a_F_XLogBackgroundFlush_3)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2])) = v53
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+272)) = v59
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[3])) = v59
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2]))
	v66 = base.B2i32(base.Ui64(v65) < base.Ui64(v63))
	if base.Ui64(v65) < base.Ui64(v63) {
		v109 = v63
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	v113 = m.G0
	v114 = int32(16)
	v115 = v113 - v114
	m.G0 = v115
	F_gettimeofday(m, v115)
	mBase = m.M
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v115)))
	v119 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+8)))
	m.G0 = v115 + v114
	v127 = v119 + v118*int64(1000000) - int64(946684800000000)
	goto L21
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+440)) = int32(1)
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	F_s_lock(m, v73+int32(440), int32(_a_F_XLogBackgroundFlush_0), int32(3001), int32(_a_F_XLogBackgroundFlush_1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v83)+440)) = v81
	v89 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2]))
	if base.Ui64(v89) < base.Ui64(v84) {
		v109 = v84
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[11]))
	if v92 < int32(0) {
		v252 = v81
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[12]))
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[3]))
	v102 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[13])))
	v103 = base.I64_div_u_s(v98-int64(1), v102)
	if v96 == v103 {
		v252 = v81
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v252 = v81
	goto L1
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[4]))
	if v129 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v166 = int32(_a_F_XLogBackgroundFlush_4)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[7])) = v168 + int32(1)
	v172 = F_WaitXLogInsertionsToFinish(m, v109)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L35
	}
L23:
	;
	v139 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2]))
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[6]))
	goto L28
L24:
	;
	v131 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[5]))
	if v131 != int64(0) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[5])) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v109
	goto L22
L27:
	;
	goto L26
L28:
	;
	if base.I64_extend_i32_s(v141)*int64(1000) <= v127-v131 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[5])) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v109
	goto L22
L30:
	;
	goto L31
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[4]))
	v152 = int64(13)
	if v151 <= base.I32_wrap_i64(int64(base.Ui64(v109)>>(uint(v152)%64))-int64(base.Ui64(v139)>>(uint(v152)%64))) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[5])) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v109
	goto L22
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	goto L22
L35:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[8]))
	v179 = F_LWLockAcquire(m, v175+int32(1024), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v181 = int32(_a_F_XLogBackgroundFlush_2)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+280)) = v183
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2])) = v183
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v188)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v188)+272)) = v189
	*(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[3])) = v189
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	if base.Ui64(v193) <= base.Ui64(v189) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[8]))
	F_LWLockRelease(m, v206+int32(1024))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L43
	}
L38:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	v197 = *(*int64)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[2]))
	if base.Ui64(v195) <= base.Ui64(v197) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v199
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v201
	F_XLogWrite(m, v12, v29, v66)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	v211 = int32(_a_F_XLogBackgroundFlush_4)
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[7]))
	v214 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[7])) = v213 - v214
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[1])))
	if v219 == v214 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[0]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+316))
	v226 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[1])) = uint8(base.B2i32(v225 != v226))
	v231 = base.B2i32(v225 == v226)
	goto L46
L45:
	;
	v231 = v214
	goto L46
L46:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[9])))
	if v233 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v246 = int32(1)
	F_AdvanceXLInsertBuffer(m, int64(0), v29, v246)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L51
	}
L48:
	;
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[9])) = uint8(v237)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBackgroundFlush[10]))
	if v240 <= v237 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	F_WalSndWakeup(m, int32(1), v231)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v252 = v246
	goto L1
}
func F_XLogFileClose(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[0]))
	if int32(0) < v9 {
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileClose[1])))
		if v13&int32(2) != 0 {
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[2]))
			v18 = int64(0)
			v21 = F_posix_fadvise(m, v17, v18, v18, int32(4))
			mBase = m.M
		}
	}
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[2]))
	v24 = F_close(m, v23)
	mBase = m.M
	if v24 != 0 {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[3]))
		v28 = v6 + int32(16)
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[4]))
		v32 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileClose[5]))
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[6]))
		F_XLogFileName(m, v28, v30, v32, v34)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[3])) = v26
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28
					F_errmsg(m, int32(_a_F_XLogFileClose_0), v6)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_XLogFileClose_1), int32(3661), int32(_a_F_XLogFileClose_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[2])) = int32(-1)
		v57 = int32(_a_F_XLogFileClose_3)
		v59 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[7]))
		*(*int32)(unsafe.Add(mBase, _c_F_XLogFileClose[7])) = v59 - int32(1)
		m.G0 = v6 + int32(80)
		return
	}
}
func F_XLogFileInit(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	v13 = F_XLogFileInitInternal(m, l0, l1, v7+int32(1071), v7+int32(32))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v13 {
			v81 = v13
			m.G0 = v7 + int32(1072)
			return v81
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInit[0]))
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInit[1]))
			if v26 != int32(14) {
				v29 = int32(_a_F_XLogFileInit_0)
			} else {
				v29 = int32(0)
			}
			v30 = v20 << (uint(int32(13)) % 32) & v29
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInit[2])))
			if v32 != int32(1) {
				v54 = v30
				v57 = v7 + int32(32)
				v60 = F_BasicOpenFile(m, v57, v54|int32(_a_F_XLogFileInit_1))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					if int32(0) <= v60 {
						v81 = v60
						m.G0 = v7 + int32(1072)
						return v81
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
								F_errmsg(m, int32(_a_F_XLogFileInit_2), v7+int32(16))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_XLogFileInit_3), int32(3396), int32(_a_F_XLogFileInit_4))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInit[3]))
				switch v36 {
				case 0, 1, 3:
					v54 = v30
					v57 = v7 + int32(32)
					v60 = F_BasicOpenFile(m, v57, v54|int32(_a_F_XLogFileInit_1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v60 {
							v81 = v60
							m.G0 = v7 + int32(1072)
							return v81
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
									F_errmsg(m, int32(_a_F_XLogFileInit_2), v7+int32(16))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_XLogFileInit_3), int32(3396), int32(_a_F_XLogFileInit_4))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				case 2:
					v54 = v30 | int32(_a_F_XLogFileInit_5)
					v57 = v7 + int32(32)
					v60 = F_BasicOpenFile(m, v57, v54|int32(_a_F_XLogFileInit_1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v60 {
							v81 = v60
							m.G0 = v7 + int32(1072)
							return v81
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
									F_errmsg(m, int32(_a_F_XLogFileInit_2), v7+int32(16))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_XLogFileInit_3), int32(3396), int32(_a_F_XLogFileInit_4))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				case 4:
					v54 = v30 | int32(_a_F_XLogFileInit_6)
					v57 = v7 + int32(32)
					v60 = F_BasicOpenFile(m, v57, v54|int32(_a_F_XLogFileInit_1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v60 {
							v81 = v60
							m.G0 = v7 + int32(1072)
							return v81
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
									F_errmsg(m, int32(_a_F_XLogFileInit_2), v7+int32(16))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_XLogFileInit_3), int32(3396), int32(_a_F_XLogFileInit_4))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36
						F_errmsg_internal(m, int32(_a_F_XLogFileInit_7), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_XLogFileInit_3), int32(_a_F_XLogFileInit_8), int32(_a_F_XLogFileInit_9))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogFileOpen(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	v14 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileOpen[0])))
	v15 = base.I64_div_u_s(int64(4294967296), v14)
	v16 = base.I64_div_u_s(l0, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v16)
	v19 = l0 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v19)
	v27 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(_a_F_XLogFileOpen_0), v9+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileOpen[1]))
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileOpen[2]))
		if v38 != int32(14) {
			v41 = int32(_a_F_XLogFileOpen_1)
		} else {
			v41 = int32(0)
		}
		v42 = v32 << (uint(int32(13)) % 32) & v41
		v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileOpen[3])))
		if v44 != int32(1) {
			v66 = v42
			v69 = v9 + int32(48)
			v72 = F_BasicOpenFile(m, v69, v66|int32(_a_F_XLogFileOpen_2))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				if v72 < int32(0) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v69
							F_errmsg(m, int32(_a_F_XLogFileOpen_3), v9+int32(16))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_XLogFileOpen_4), int32(3628), int32(_a_F_XLogFileOpen_5))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					m.G0 = v9 + int32(1072)
					return v72
				}
			}
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileOpen[4]))
			switch v48 {
			case 0, 1, 3:
				v66 = v42
				v69 = v9 + int32(48)
				v72 = F_BasicOpenFile(m, v69, v66|int32(_a_F_XLogFileOpen_2))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v69
								F_errmsg(m, int32(_a_F_XLogFileOpen_3), v9+int32(16))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_XLogFileOpen_4), int32(3628), int32(_a_F_XLogFileOpen_5))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v9 + int32(1072)
						return v72
					}
				}
			case 2:
				v66 = v42 | int32(_a_F_XLogFileOpen_6)
				v69 = v9 + int32(48)
				v72 = F_BasicOpenFile(m, v69, v66|int32(_a_F_XLogFileOpen_2))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v69
								F_errmsg(m, int32(_a_F_XLogFileOpen_3), v9+int32(16))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_XLogFileOpen_4), int32(3628), int32(_a_F_XLogFileOpen_5))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v9 + int32(1072)
						return v72
					}
				}
			case 4:
				v66 = v42 | int32(_a_F_XLogFileOpen_7)
				v69 = v9 + int32(48)
				v72 = F_BasicOpenFile(m, v69, v66|int32(_a_F_XLogFileOpen_2))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v69
								F_errmsg(m, int32(_a_F_XLogFileOpen_3), v9+int32(16))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_XLogFileOpen_4), int32(3628), int32(_a_F_XLogFileOpen_5))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v9 + int32(1072)
						return v72
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
					F_errmsg_internal(m, int32(_a_F_XLogFileOpen_8), v9)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_XLogFileOpen_4), int32(_a_F_XLogFileOpen_9), int32(_a_F_XLogFileOpen_10))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogFileRead(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v10 = m.G0
	v12 = v10 - int32(1264)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = l1
	v17 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[0])))
	v18 = base.I64_div_u_s(int64(4294967296), v17)
	v19 = base.I64_div_u_s(l0, v18)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+84)) = uint32(v19)
	v22 = l0 - v18*v19
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+88)) = uint32(v22)
	v25 = v12 + int32(1200)
	v30 = F_pg_snprintf(m, v25, int32(64), int32(_a_F_XLogFileRead_0), v12+int32(80))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		if l2 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v25
			v38 = v12 + int32(1120)
			v43 = F_pg_snprintf(m, v38, int32(80), int32(_a_F_XLogFileRead_1), v12+int32(48))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = F_strlen(m, v38)
				mBase = m.M
				v47 = v12 + int32(96)
				v50 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[0])))
				v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileRead[1])))
				v53 = F_RestoreArchivedFile(m, v47, v25, int32(_a_F_XLogFileRead_2), v50, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					if v53 == int32(0) {
						v144 = int32(-1)
						m.G0 = v12 + int32(1264)
						return v144
					} else {
						F_KeepFileRestoredFromArchive(m, v47, v25)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v25
							v64 = F_pg_snprintf(m, v47, int32(1024), int32(_a_F_XLogFileRead_3), v12+int32(32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v90 = F_BasicOpenFile(m, v12+int32(96), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									if int32(0) <= v90 {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[2])) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(1200)
										v100 = v12 + int32(1120)
										v103 = F_pg_snprintf(m, v100, int32(80), int32(_a_F_XLogFileRead_4), v12)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v105 = F_strlen(m, v100)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[3])) = l2
											*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[4])) = l2
											if l2 == int32(3) {
												v144 = v90
											} else {
												v116 = m.G0
												v117 = int32(16)
												v118 = v116 - v117
												m.G0 = v118
												F_gettimeofday(m, v118)
												mBase = m.M
												v121 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
												v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v118)+8)))
												m.G0 = v118 + v117
												*(*int64)(unsafe.Add(mBase, _c_F_XLogFileRead[5])) = v122 + v121*int64(1000000) - int64(946684800000000)
												v144 = v90
											}
											m.G0 = v12 + int32(1264)
											return v144
										}
									} else {
										if l3 == int32(0) {
											F_errstart_cold(m, int32(23), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(96)
													F_errmsg(m, int32(_a_F_XLogFileRead_5), v12+int32(16))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_XLogFileRead_6), int32(_a_F_XLogFileRead_7), int32(_a_F_XLogFileRead_8))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v135 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[6]))
											if v135 != int32(44) {
												F_errstart_cold(m, int32(23), int32(0))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(96)
														F_errmsg(m, int32(_a_F_XLogFileRead_5), v12+int32(16))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_XLogFileRead_6), int32(_a_F_XLogFileRead_7), int32(_a_F_XLogFileRead_8))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v144 = int32(-1)
												m.G0 = v12 + int32(1264)
												return v144
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
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l1
			v69 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[0])))
			v70 = base.I64_div_u_s(int64(4294967296), v69)
			v71 = base.I64_div_u_s(l0, v70)
			*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v71)
			v74 = l0 - v70*v71
			*(*uint32)(unsafe.Add(mBase, uint32(v12)+72)) = uint32(v74)
			v82 = F_pg_snprintf(m, v12+int32(96), int32(1024), int32(_a_F_XLogFileRead_9), v12-int32(-64))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				v90 = F_BasicOpenFile(m, v12+int32(96), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					if int32(0) <= v90 {
						*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[2])) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(1200)
						v100 = v12 + int32(1120)
						v103 = F_pg_snprintf(m, v100, int32(80), int32(_a_F_XLogFileRead_4), v12)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v105 = F_strlen(m, v100)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[3])) = l2
							*(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[4])) = l2
							if l2 == int32(3) {
								v144 = v90
							} else {
								v116 = m.G0
								v117 = int32(16)
								v118 = v116 - v117
								m.G0 = v118
								F_gettimeofday(m, v118)
								mBase = m.M
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
								v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v118)+8)))
								m.G0 = v118 + v117
								*(*int64)(unsafe.Add(mBase, _c_F_XLogFileRead[5])) = v122 + v121*int64(1000000) - int64(946684800000000)
								v144 = v90
							}
							m.G0 = v12 + int32(1264)
							return v144
						}
					} else {
						if l3 == int32(0) {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(96)
									F_errmsg(m, int32(_a_F_XLogFileRead_5), v12+int32(16))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_XLogFileRead_6), int32(_a_F_XLogFileRead_7), int32(_a_F_XLogFileRead_8))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v135 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileRead[6]))
							if v135 != int32(44) {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(96)
										F_errmsg(m, int32(_a_F_XLogFileRead_5), v12+int32(16))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_XLogFileRead_6), int32(_a_F_XLogFileRead_7), int32(_a_F_XLogFileRead_8))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v144 = int32(-1)
								m.G0 = v12 + int32(1264)
								return v144
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogGetLastRemovedSegno(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_XLogGetLastRemovedSegno[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+440)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_XLogGetLastRemovedSegno[0]))
		F_s_lock(m, v9+int32(440), int32(_a_F_XLogGetLastRemovedSegno_0), int32(3760), int32(_a_F_XLogGetLastRemovedSegno_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogGetLastRemovedSegno[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_XLogGetLastRemovedSegno[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
		return v23
	}
}
func F_XLogPrefetchReconfigure(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(_a_F_XLogPrefetchReconfigure_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchReconfigure[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchReconfigure[0])) = v3 + int32(1)
	return
}
func F_XLogPrefetcherNextBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int64
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int64
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int64
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int64
	_ = v420
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int64
	_ = v484
	var v485 int32
	_ = v485
	var v487 int64
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v540 int64
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int64
	_ = v603
	var v604 int64
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int64
	_ = v613
	var v619 int32
	_ = v619
	var v620 int64
	_ = v620
	var v627 int32
	_ = v627
	var v628 int64
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = l0 + int32(36)
	v23 = l0 + int32(84)
	v25 = l0 + int32(28)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)+32))
	v32 = v27
	v36 = v26
	goto L4
L1:
	;
	m.G0 = v18 + int32(48)
	return v644
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L16
	} else {
		goto L172
	}
L3:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v619)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v619)+48)) = v620 + int64(1)
	v644 = int32(0)
	goto L1
L4:
	;
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v612)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v612)+32)) = v613 + int64(1)
	v644 = int32(0)
	goto L1
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+120))
	if v46 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v85 = v36
	goto L8
L8:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	if base.Ui64(v86) <= base.Ui64(v28) {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[1]))
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v67 = F_XLogReadAhead(m, v32, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L20
	}
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1256)))
	if v49 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v52 = int32(2)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui64(v28) <= base.Ui64(v53) {
		v644 = v52
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v56 = F_XLogReadAhead(m, v32, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	if v56 != 0 {
		v71 = v56
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+124))
	if v61 == int32(0) {
		v644 = v52
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v64
	v644 = v52
	goto L1
L20:
	;
	if v67 != 0 {
		v71 = v67
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v644 = int32(2)
	goto L1
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71
	v85 = v71
	goto L8
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[2]))
	if int32(0) < v75 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	v644 = int32(0)
	goto L1
L26:
	;
	goto L25
L27:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v85)+72))
	if v253 <= v254 {
		goto L79
	} else {
		goto L80
	}
L28:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+48)))
	v90 = v88 & int32(-16)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+49)))
	switch v91 {
	case 0:
		goto L31
	default:
		goto L27
	case 2:
		goto L29
	case 4:
		goto L30
	}
L29:
	;
	v148 = v90 - int32(16)
	if v148 != 0 {
		goto L48
	} else {
		goto L49
	}
L30:
	;
	if v90 != 0 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	if v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = base.B2i32(v90 != int32(144))
	goto L34
L33:
	;
	v95 = int32(0)
	goto L34
L34:
	;
	if v95 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v86
	goto L27
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85)+64))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v110 = F_hash_search(m, v104, v18+int32(32), int32(1), v18+int32(47))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v110)+16)) = v86
	if v112 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v118 == v116 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v110)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v134 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v122 = v25
	goto L43
L42:
	;
	v122 = v118
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v110)+32)) = v122
	v126 = v110 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v126
	goto L27
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v138 = v25
	goto L46
L45:
	;
	v138 = v134
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v110)+32)) = v138
	v142 = v110 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(0)
	goto L27
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v85)+64))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v200)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v212 = F_hash_search(m, v206, v18+int32(32), int32(1), v18+int32(47))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L16
	} else {
		goto L65
	}
L48:
	;
	if v148 == int32(16) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v85)+64))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v152 != 0 {
		goto L27
	} else {
		goto L54
	}
L51:
	;
	goto L47
L52:
	;
	goto L27
L54:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v163 = F_hash_search(m, v157, v18+int32(32), int32(1), v18+int32(47))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+16)) = v86
	if v165 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v171 == v169 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v163)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v187 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v175 = v25
	goto L61
L60:
	;
	v175 = v171
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v163)+32)) = v175
	v179 = v163 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v179
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v179
	goto L27
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v191 = v25
	goto L64
L63:
	;
	v191 = v187
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v163)+32)) = v191
	v195 = v163 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v195
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = int32(0)
	goto L27
L65:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+16)) = v86
	if v214 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+24)) = v201
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v219 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v212)+28))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v212)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v235 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v223 = v25
	goto L71
L70:
	;
	v223 = v219
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v212)+32)) = v223
	v227 = v212 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v227
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v227
	goto L27
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v239 = v25
	goto L74
L73:
	;
	v239 = v235
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v212)+32)) = v239
	v243 = v212 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v212)+24))
	if base.Ui32(v246) < base.Ui32(v201) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v248 = v246
	goto L77
L76:
	;
	v248 = v201
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+24)) = v248
	goto L27
L78:
	;
	goto L5
L79:
	;
	v261 = v253
	goto L82
L80:
	;
	goto L81
L81:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+124))
	if v600 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L82:
	;
	v273 = int32(1)
	v274 = v261 + v273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v274
	v278 = v85 + int32(76) + v261*int32(52)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v279 == v273 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L81
L84:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	if v284 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v581 = v274
	goto L86
L86:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v85)+72))
	if v581 <= v582 {
		v261 = v581
		goto L82
	} else {
		goto L168
	}
L87:
	;
	v644 = int32(0)
	goto L1
L88:
	;
	goto L89
L89:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+29)))
	if v286 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v291)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v291)+40)) = v292 + int64(1)
	v644 = int32(0)
	goto L1
L91:
	;
	goto L92
L92:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+28)))
	if v296&int32(64) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+24)) = v302 + int64(1)
	v644 = int32(0)
	goto L1
L94:
	;
	goto L95
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v278)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v312 = int32(0)
	if base.B2i32(v311 == v312)|base.B2i32(v25 == v311) == v312 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v321 = int32(0)
	v323 = F_hash_search(m, v318, v18+int32(32), v321, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L16
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v340 = v278 + int32(4)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v341 != v342 {
		goto L106
	} else {
		goto L107
	}
L99:
	;
	if v323 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323)+24))
	if base.Ui32(v325) <= base.Ui32(v306) {
		goto L78
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v327 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v327
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v336 = F_hash_search(m, v331, v18+int32(32), v327, v327)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	if v336 != 0 {
		goto L78
	} else {
		goto L105
	}
L105:
	;
	goto L98
L106:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v341 != v353 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v344 != v345 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v347 != v348 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v350 == v351 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	goto L106
L111:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v341 != v364 {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v355 != v356 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v358 != v359 {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v361 == v362 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v341 != v375 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v366 != v367 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v369 != v370 {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v372 == v373 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	goto L116
L121:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v389 = v21 + v386*int32(12)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v340)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+8)) = v390
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	*(*int64)(unsafe.Add(mBase, uint32(v389))) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v394<<(uint(int32(2))%32)))) = v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v404 = base.I32_rem_s(v400+int32(1), int32(4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v340)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v406
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v408
	v413 = F_smgropen(m, v18+int32(16), int32(-1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L16
	} else {
		goto L126
	}
L122:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v377 != v378 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v380 != v381 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v383 == v384 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	v416 = F_smgrexists(m, v413, int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	if v416 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v340)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v431 = F_hash_search(m, v425, v18+int32(32), int32(1), v18+int32(47))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L16
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v480 = F_smgrnblocks(m, v413, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L16
	} else {
		goto L142
	}
L131:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v431)+16)) = v420
	if v433 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v473)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v473)+32)) = v474 + int64(1)
	v644 = int32(0)
	goto L1
L133:
	;
	v437 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v431)+24)) = v437
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v439 == v437 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v431)+28))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v431)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+4)) = v451
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v431)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v453
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v455 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v443 = v25
	goto L138
L137:
	;
	v443 = v439
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v431)+32)) = v443
	v447 = v431 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v447
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v447
	goto L132
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v459 = v25
	goto L141
L140:
	;
	v459 = v455
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v431)+32)) = v459
	v463 = v431 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v463
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v431)+24)) = int32(0)
	goto L132
L142:
	;
	if base.Ui32(v480) <= base.Ui32(v478) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v340)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v485
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v495 = F_hash_search(m, v489, v18+int32(32), int32(1), v18+int32(47))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L16
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	F_PrefetchSharedBuffer(m, v18+int32(32), v413, v546, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L16
	} else {
		goto L160
	}
L146:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v495)+16)) = v484
	if v497 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v539)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v539)+32)) = v540 + int64(1)
	v644 = int32(0)
	goto L1
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495)+24)) = v483
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v502 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v516
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v518 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v506 = v25
	goto L153
L152:
	;
	v506 = v502
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v495)+32)) = v506
	v510 = v495 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = v510
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v510
	goto L147
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25
	v522 = v25
	goto L156
L155:
	;
	v522 = v518
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v495)+32)) = v522
	v526 = v495 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v526
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v526
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v495)+24))
	if base.Ui32(v529) < base.Ui32(v483) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v531 = v529
	goto L159
L158:
	;
	v531 = v483
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495)+24)) = v531
	goto L147
L160:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	if v550 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v553)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v553)+16)) = v554 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v550
	v644 = int32(0)
	goto L1
L162:
	;
	goto L163
L163:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+36)))
	if v559&int32(1) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[0]))
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v563)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v563)+8)) = v564 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = int32(0)
	v644 = int32(1)
	goto L1
L165:
	;
	goto L166
L166:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPrefetcherNextBlock[3])))
	if v572&int32(1) == int32(0) {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v581 = v577
	goto L86
L168:
	;
	goto L83
L169:
	;
	v607 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v607
	v32 = v599
	v36 = v607
	goto L4
L170:
	;
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v600)+16))
	v604 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v603 != v604 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v644 = int32(2)
	goto L1
L172:
	;
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v413)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v629
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v628
	F_errmsg_internal(m, int32(_a_F_XLogPrefetcherNextBlock_0), v18)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L16
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_XLogPrefetcherNextBlock_1), int32(796), int32(_a_F_XLogPrefetcherNextBlock_2))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogReadBufferForRedoExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = v13 + int32(68)
	v19 = v13 - int32(-64)
	v21 = v13 + int32(60)
	v23 = v13 + int32(56)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	if v26 < l1 {
		v50 = v6
	} else {
		v30 = v25 + l1*int32(52)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+76)))
		if v31 != int32(1) {
			v50 = v6
		} else {
			v35 = v30 + int32(76)
			if v17 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v36
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v38
			} else {
			}
			if v19 != 0 {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v40
			} else {
			}
			if v21 != 0 {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v42
			} else {
			}
			v44 = int32(1)
			if v23 == int32(0) {
				v50 = v44
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v47
				v50 = v44
			}
		}
	}
	if v50 != 0 {
		v52 = l2 - int32(1)
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v59 = v56 + l1*int32(52)
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+104)))
		v62 = v60 & int32(64)
		if v62 != 0 {
			v63 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v52))
		} else {
			v63 = int32(0)
		}
		if v63 != 0 {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v198 = m.ExcPending
			if v198 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_XLogReadBufferForRedoExtended_0), int32(0))
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_XLogReadBufferForRedoExtended_1), int32(369), int32(_a_F_XLogReadBufferForRedoExtended_2))
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if base.B2i32(v62 == int32(0))&base.B2i32(base.Ui32(v52) <= base.Ui32(int32(1))) != 0 {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_XLogReadBufferForRedoExtended_3), int32(0))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_XLogReadBufferForRedoExtended_1), int32(371), int32(_a_F_XLogReadBufferForRedoExtended_2))
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+106)))
				if v71 == int32(1) {
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v13)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v74
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v76
					if l3 != 0 {
						v82 = int32(2)
					} else {
						v82 = int32(1)
					}
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
					v84 = F_XLogReadBufferExtended(m, v13+int32(16), v70, v69, v82, v83)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v84
						if v84 < int32(0) {
							v92 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[0]))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+(v84^int32(-1))<<(uint(int32(2))%32))))
							v106 = v98
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[1]))
							v106 = v100 + v84<<(uint(int32(13))%32) + int32(-8192)
						}
						v107 = F_RestoreBlockImage(m, l0, l1, v106)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							if v107 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(2600))
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return int32(0)
									} else {
										v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = v228
										F_errmsg_internal(m, int32(_a_F_XLogReadBufferForRedoExtended_4), v13)
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_XLogReadBufferForRedoExtended_1), int32(384), int32(_a_F_XLogReadBufferForRedoExtended_2))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+14)))
								if v111 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v106))) = base.I64_rotr(v15, int64(32))
								} else {
								}
								v115 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								F_MarkBufferDirty(m, v115)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v118 = int32(2)
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
									if v119 != int32(3) {
										v173 = v118
										m.G0 = v13 + int32(80)
										return v173
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										F_FlushOneBuffer(m, v122)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											v173 = v118
											m.G0 = v13 + int32(80)
											return v173
										}
									}
								}
							}
						}
					}
				} else {
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v13)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v125
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v127
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
					v132 = F_XLogReadBufferExtended(m, v13+int32(32), v70, v69, l2, v131)
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v132
						if v132 == int32(0) {
							v173 = int32(3)
							m.G0 = v13 + int32(80)
							return v173
						} else {
							if base.Ui32(l2-int32(3)) <= base.Ui32(int32(-3)) {
								if l3 != 0 {
									F_LockBufferForCleanup(m, v132)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										v148 = v147
										if v148 < int32(0) {
											v152 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[0]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
											v166 = v158
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[1]))
											v166 = v160 + v148<<(uint(int32(13))%32) + int32(-8192)
										}
										v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166)+4)))
										v168 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166))))
										v173 = base.B2i32(base.Ui64(v15) <= base.Ui64(v167|v168<<(uint(int64(32))%64)))
										m.G0 = v13 + int32(80)
										return v173
									}
								} else {
									F_LockBuffer(m, v132, int32(2))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										v148 = v147
										if v148 < int32(0) {
											v152 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[0]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
											v166 = v158
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[1]))
											v166 = v160 + v148<<(uint(int32(13))%32) + int32(-8192)
										}
										v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166)+4)))
										v168 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166))))
										v173 = base.B2i32(base.Ui64(v15) <= base.Ui64(v167|v168<<(uint(int64(32))%64)))
										m.G0 = v13 + int32(80)
										return v173
									}
								}
							} else {
								v148 = v132
								if v148 < int32(0) {
									v152 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[0]))
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
									v166 = v158
								} else {
									v160 = *(*int32)(unsafe.Add(mBase, _c_F_XLogReadBufferForRedoExtended[1]))
									v166 = v160 + v148<<(uint(int32(13))%32) + int32(-8192)
								}
								v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166)+4)))
								v168 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166))))
								v173 = base.B2i32(base.Ui64(v15) <= base.Ui64(v167|v168<<(uint(int64(32))%64)))
								m.G0 = v13 + int32(80)
								return v173
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v183 = m.ExcPending
		if v183 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
			F_errmsg_internal(m, int32(_a_F_XLogReadBufferForRedoExtended_5), v13+int32(48))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_XLogReadBufferForRedoExtended_1), int32(359), int32(_a_F_XLogReadBufferForRedoExtended_2))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_XLogResetInsertion(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	v1 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_XLogResetInsertion[0]))
	if v8 <= v1 {
	} else {
		v12 = v8 & int32(7)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLogResetInsertion[1]))
		if base.Ui32(int32(8)) <= base.Ui32(v8) {
			v20 = v1
			v24 = v1
			for {
				v27 = v14 + v20*int32(_a_F_XLogResetInsertion_0)
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[2]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[3]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[4]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[5]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[6]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[7]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_XLogResetInsertion[8]))) = uint8(v28)
				*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v28)
				v44 = int32(8)
				v45 = v20 + v44
				v47 = v24 + v44
				if v47 != v8&int32(2147483640) {
					v20 = v45
					v24 = v47
					continue
				} else {
					break
				}
				break
			}
			if v12 == int32(0) {
			} else {
				v52 = v45
				v58 = int32(0)
				v59 = v52
				for {
					v67 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v14+v59*int32(_a_F_XLogResetInsertion_0)))) = uint8(v67)
					v69 = int32(1)
					v72 = v58 + v69
					if v72 != v12 {
						v58 = v72
						v59 = v59 + v69
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v52 = v1
			v58 = int32(0)
			v59 = v52
			for {
				v67 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v14+v59*int32(_a_F_XLogResetInsertion_0)))) = uint8(v67)
				v69 = int32(1)
				v72 = v58 + v69
				if v72 != v12 {
					v58 = v72
					v59 = v59 + v69
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, _c_F_XLogResetInsertion[9])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogResetInsertion[10])) = int32(_a_F_XLogResetInsertion_1)
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogResetInsertion[0])) = v87
	*(*int32)(unsafe.Add(mBase, _c_F_XLogResetInsertion[11])) = v87
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogResetInsertion[12])) = uint8(v87)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogResetInsertion[13])) = uint8(v87)
	return
}
func F_XLogSetRecordFlags(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = int32(_a_F_XLogSetRecordFlags_0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSetRecordFlags[0])))
	v5 = v4 | l0
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSetRecordFlags[0])) = uint8(v5)
	return
}
