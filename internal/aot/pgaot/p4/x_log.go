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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(1168)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	if v11 <= int32(0) {
		v108 = int32(1)
		m.G0 = v8 + int32(1168)
		return v108
	} else {
		if v11 != int32(2) {
			v17 = int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[199]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+440))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+440)) = v17
			if v20 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[199]))
				F_s_lock(m, v24+int32(440), int32(498539), int32(6414), int32(353961))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _consts[199]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+440)) = int32(0)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+316))
					if v38 == int32(1) {
						v108 = v17
						m.G0 = v8 + int32(1168)
						return v108
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(373087)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
						v53 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(32))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v61 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
							mBase = m.M
							if v61 == int32(0) {
								v108 = int32(1)
								m.G0 = v8 + int32(1168)
								return v108
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(22921)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
								v74 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(16))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = int32(0)
									v83 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), v76)
									mBase = m.M
									if v83 == int32(0) {
										v108 = v76
										m.G0 = v8 + int32(1168)
										return v108
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(373087)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										v93 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											v102 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
											mBase = m.M
											if v102 == int32(0) {
												v108 = int32(1)
												m.G0 = v8 + int32(1168)
												return v108
											} else {
												F_XLogArchiveNotify(m, l0)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = int32(0)
													m.G0 = v8 + int32(1168)
													return v108
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
				v35 = *(*int32)(unsafe.Add(mBase, _consts[199]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+440)) = int32(0)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+316))
				if v38 == int32(1) {
					v108 = v17
					m.G0 = v8 + int32(1168)
					return v108
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(373087)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
					v53 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(32))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v61 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
						mBase = m.M
						if v61 == int32(0) {
							v108 = int32(1)
							m.G0 = v8 + int32(1168)
							return v108
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(22921)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							v74 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(16))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = int32(0)
								v83 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), v76)
								mBase = m.M
								if v83 == int32(0) {
									v108 = v76
									m.G0 = v8 + int32(1168)
									return v108
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(373087)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									v93 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										v102 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
										mBase = m.M
										if v102 == int32(0) {
											v108 = int32(1)
											m.G0 = v8 + int32(1168)
											return v108
										} else {
											F_XLogArchiveNotify(m, l0)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = int32(0)
												m.G0 = v8 + int32(1168)
												return v108
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(373087)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
			v53 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(32))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v61 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
				mBase = m.M
				if v61 == int32(0) {
					v108 = int32(1)
					m.G0 = v8 + int32(1168)
					return v108
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(22921)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v74 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8+int32(16))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = int32(0)
						v83 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), v76)
						mBase = m.M
						if v83 == int32(0) {
							v108 = v76
							m.G0 = v8 + int32(1168)
							return v108
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(373087)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v93 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(175729), v8)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v102 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
								mBase = m.M
								if v102 == int32(0) {
									v108 = int32(1)
									m.G0 = v8 + int32(1168)
									return v108
								} else {
									F_XLogArchiveNotify(m, l0)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v108 = int32(0)
										m.G0 = v8 + int32(1168)
										return v108
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
	var v47 int32
	_ = v47
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
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
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v17 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v255
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
	v23 = base.B2i32(v21 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v23)
	if v21 != int32(2) {
		v255 = int32(0)
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_s_lock(m, v31+int32(440), int32(498539), int32(2990), int32(322497))
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
	v41 = int32(4411280)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v42)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v45
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+440)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v49 & int64(-8192)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v42)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+280)) = v53
	v55 = int32(4411320)
	*(*int64)(unsafe.Add(mBase, _consts[276])) = v53
	v58 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+272)) = v59
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v59
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v65 = *(*int64)(unsafe.Add(mBase, _consts[276]))
	v66 = base.B2i32(base.Ui64(v65) < base.Ui64(v63))
	if v66 == v47 {
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
	v70 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+440)) = int32(1)
	if v71 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v93 = v63
	v94 = v65
	goto L13
L13:
	;
	if base.Ui64(v93) <= base.Ui64(v94) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_s_lock(m, v75+int32(440), int32(498539), int32(3001), int32(322497))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v84)+440)) = int32(0)
	v90 = *(*int64)(unsafe.Add(mBase, _consts[276]))
	v93 = v85
	v94 = v90
	goto L13
L17:
	;
	goto L16
L18:
	;
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	if v98 < v96 {
		v255 = v96
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v116 = m.G0
	v117 = int32(16)
	v118 = v116 - v117
	m.G0 = v118
	F___gettimeofday(m, v118)
	mBase = m.M
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v118)+8)))
	m.G0 = v118 + v117
	v130 = v122 + v121*int64(1000000) - int64(946684800000000)
	goto L24
L21:
	;
	v102 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	v104 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v108 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
	v109 = base.I64_div_u_s(v104-int64(1), v108)
	if v102 == v109 {
		v255 = v96
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v255 = v96
	goto L1
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	if v132 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v169 = int32(4510372)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v171 + int32(1)
	v175 = F_WaitXLogInsertionsToFinish(m, v93)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L38
	}
L26:
	;
	v142 = *(*int64)(unsafe.Add(mBase, _consts[276]))
	v144 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	goto L31
L27:
	;
	v134 = *(*int64)(unsafe.Add(mBase, _consts[279]))
	if v134 != int64(0) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int64)(unsafe.Add(mBase, _consts[279])) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v93
	goto L25
L30:
	;
	goto L29
L31:
	;
	if base.I64_extend_i32_s(v144)*int64(1000) <= v130-v134 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, _consts[279])) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v93
	goto L25
L33:
	;
	goto L34
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v155 = int64(13)
	if v154 <= base.I32_wrap_i64(int64(base.Ui64(v93)>>(uint(v155)%64))-int64(base.Ui64(v142)>>(uint(v155)%64))) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, _consts[279])) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v93
	goto L25
L36:
	;
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	goto L25
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v182 = F_LWLockAcquire(m, v178+int32(1024), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v184 = int32(4411280)
	v185 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v185)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v185)+280)) = v186
	*(*int64)(unsafe.Add(mBase, _consts[276])) = v186
	v191 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v191)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v191)+272)) = v192
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v192
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	if base.Ui64(v196) <= base.Ui64(v192) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v209+int32(1024))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L46
	}
L41:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	v200 = *(*int64)(unsafe.Add(mBase, _consts[276]))
	if base.Ui64(v198) <= base.Ui64(v200) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v204
	F_XLogWrite(m, v12, v29, v66)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	goto L40
L46:
	;
	v214 = int32(4510372)
	v216 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v217 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v216 - v217
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v222 == v217 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+316))
	v229 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(base.B2i32(v228 != v229))
	v234 = base.B2i32(v228 == v229)
	goto L49
L48:
	;
	v234 = v217
	goto L49
L49:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _consts[281])))
	if v236 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v249 = int32(1)
	F_AdvanceXLInsertBuffer(m, int64(0), v29, v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L54
	}
L51:
	;
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[281])) = uint8(v240)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[282]))
	if v243 <= v240 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_WalSndWakeup(m, int32(1), v234)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v255 = v249
	goto L1
}
func F_XLogFileClose(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if int32(0) < v8 {
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
		if v12&int32(2) != 0 {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[268]))
			v17 = int64(0)
			v20 = F_posix_fadvise(m, v16, v17, v17, int32(4))
			mBase = m.M
		}
	}
	v22 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	v23 = F_close(m, v22)
	mBase = m.M
	if v23 != 0 {
		v25 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		v29 = *(*int32)(unsafe.Add(mBase, _consts[269]))
		v31 = *(*int64)(unsafe.Add(mBase, _consts[270]))
		v33 = *(*int32)(unsafe.Add(mBase, _consts[271]))
		F_XLogFileName(m, v5+int32(16), v29, v31, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[140])) = v25
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(16)
					F_errmsg(m, int32(299662), v5)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errfinish(m, int32(498539), int32(3661), int32(361573))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
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
		*(*int32)(unsafe.Add(mBase, _consts[268])) = int32(-1)
		v58 = int32(4431708)
		v60 = *(*int32)(unsafe.Add(mBase, _consts[272]))
		*(*int32)(unsafe.Add(mBase, _consts[272])) = v60 - int32(1)
		m.G0 = v5 + int32(80)
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
			v82 = v13
			m.G0 = v7 + int32(1072)
			return v82
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[267]))
			v24 = *(*int32)(unsafe.Add(mBase, _consts[273]))
			v25 = int32(14)
			v29 = v20 << (uint(int32(13)) % 32) & (base.B2i32(v24 != v25) << (uint(v25) % 32))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
			if v31 != int32(1) {
				v53 = v29
				v59 = F_BasicOpenFile(m, v7+int32(32), v53|int32(524290))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					if int32(0) <= v59 {
						v82 = v59
						m.G0 = v7 + int32(1072)
						return v82
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(32)
								F_errmsg(m, int32(298919), v7+int32(16))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498539), int32(3396), int32(100547))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
				v35 = *(*int32)(unsafe.Add(mBase, _consts[275]))
				switch v35 {
				case 0, 1, 3:
					v53 = v29
					v59 = F_BasicOpenFile(m, v7+int32(32), v53|int32(524290))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v59 {
							v82 = v59
							m.G0 = v7 + int32(1072)
							return v82
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(32)
									F_errmsg(m, int32(298919), v7+int32(16))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(498539), int32(3396), int32(100547))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
					v53 = v29 | int32(1052672)
					v59 = F_BasicOpenFile(m, v7+int32(32), v53|int32(524290))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v59 {
							v82 = v59
							m.G0 = v7 + int32(1072)
							return v82
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(32)
									F_errmsg(m, int32(298919), v7+int32(16))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(498539), int32(3396), int32(100547))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
					v53 = v29 | int32(4096)
					v59 = F_BasicOpenFile(m, v7+int32(32), v53|int32(524290))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v59 {
							v82 = v59
							m.G0 = v7 + int32(1072)
							return v82
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(32)
									F_errmsg(m, int32(298919), v7+int32(16))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(498539), int32(3396), int32(100547))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v35
						F_errmsg_internal(m, int32(488434), v7)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498539), int32(8695), int32(103315))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	v14 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
	v15 = base.I64_div_u_s(int64(4294967296), v14)
	v16 = base.I64_div_u_s(l0, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v16)
	v19 = l0 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v19)
	v27 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(510524), v9+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _consts[267]))
		v36 = *(*int32)(unsafe.Add(mBase, _consts[273]))
		v37 = int32(14)
		v41 = v32 << (uint(int32(13)) % 32) & (base.B2i32(v36 != v37) << (uint(v37) % 32))
		v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
		if v43 != int32(1) {
			v65 = v41
			v71 = F_BasicOpenFile(m, v9+int32(48), v65|int32(524290))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				if v71 < int32(0) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
							F_errmsg(m, int32(298919), v9+int32(16))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498539), int32(3628), int32(282062))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
					return v71
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, _consts[275]))
			switch v47 {
			case 0, 1, 3:
				v65 = v41
				v71 = F_BasicOpenFile(m, v9+int32(48), v65|int32(524290))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if v71 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
								F_errmsg(m, int32(298919), v9+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498539), int32(3628), int32(282062))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
						return v71
					}
				}
			case 2:
				v65 = v41 | int32(1052672)
				v71 = F_BasicOpenFile(m, v9+int32(48), v65|int32(524290))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if v71 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
								F_errmsg(m, int32(298919), v9+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498539), int32(3628), int32(282062))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
						return v71
					}
				}
			case 4:
				v65 = v41 | int32(4096)
				v71 = F_BasicOpenFile(m, v9+int32(48), v65|int32(524290))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if v71 < int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
								F_errmsg(m, int32(298919), v9+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498539), int32(3628), int32(282062))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
						return v71
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v47
					F_errmsg_internal(m, int32(488434), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498539), int32(8695), int32(103315))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v9 = m.G0
	v11 = v9 - int32(1264)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l1
	v16 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
	v17 = base.I64_div_u_s(int64(4294967296), v16)
	v18 = base.I64_div_u_s(l0, v17)
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+84)) = uint32(v18)
	v21 = l0 - v17*v18
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+88)) = uint32(v21)
	v29 = F_pg_snprintf(m, v11+int32(1200), int32(64), int32(510531), v11+int32(80))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		if l2 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(1200)
			v44 = F_pg_snprintf(m, v11+int32(1120), int32(80), int32(181551), v11+int32(48))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v48 = F_strlen(m, v11+int32(1120))
				mBase = m.M
				v55 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
				v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[304])))
				v58 = F_RestoreArchivedFile(m, v11+int32(96), v11+int32(1200), int32(536581), v55, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					if v58 == int32(0) {
						v156 = int32(-1)
						m.G0 = v11 + int32(1264)
						return v156
					} else {
						F_KeepFileRestoredFromArchive(m, v11+int32(96), v11+int32(1200))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(1200)
							v77 = F_pg_snprintf(m, v11+int32(96), int32(1024), int32(177132), v11+int32(32))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v102 = F_BasicOpenFile(m, v11+int32(96), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									if int32(0) <= v102 {
										*(*int32)(unsafe.Add(mBase, _consts[305])) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(1200)
										v115 = F_pg_snprintf(m, v11+int32(1120), int32(80), int32(186516), v11)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											v119 = F_strlen(m, v11+int32(1120))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, _consts[306])) = l2
											*(*int32)(unsafe.Add(mBase, _consts[307])) = l2
											if l2 == int32(3) {
												v156 = v102
											} else {
												v130 = m.G0
												v131 = int32(16)
												v132 = v130 - v131
												m.G0 = v132
												F___gettimeofday(m, v132)
												mBase = m.M
												v135 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
												v136 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+8)))
												m.G0 = v132 + v131
												*(*int64)(unsafe.Add(mBase, _consts[308])) = v136 + v135*int64(1000000) - int64(946684800000000)
												v156 = v102
											}
											m.G0 = v11 + int32(1264)
											return v156
										}
									} else {
										if l3 == int32(0) {
											F_errstart_cold(m, int32(23), int32(0))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(96)
													F_errmsg(m, int32(298919), v11+int32(16))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(492731), int32(4307), int32(465159))
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
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
											v149 = *(*int32)(unsafe.Add(mBase, _consts[140]))
											if v149 != int32(44) {
												F_errstart_cold(m, int32(23), int32(0))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(96)
														F_errmsg(m, int32(298919), v11+int32(16))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(492731), int32(4307), int32(465159))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
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
												v156 = int32(-1)
												m.G0 = v11 + int32(1264)
												return v156
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
			*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l1
			v82 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
			v83 = base.I64_div_u_s(int64(4294967296), v82)
			v84 = base.I64_div_u_s(l0, v83)
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+68)) = uint32(v84)
			v87 = l0 - v83*v84
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+72)) = uint32(v87)
			v95 = F_pg_snprintf(m, v11+int32(96), int32(1024), int32(510524), v11-int32(-64))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				v102 = F_BasicOpenFile(m, v11+int32(96), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					if int32(0) <= v102 {
						*(*int32)(unsafe.Add(mBase, _consts[305])) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(1200)
						v115 = F_pg_snprintf(m, v11+int32(1120), int32(80), int32(186516), v11)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v119 = F_strlen(m, v11+int32(1120))
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[306])) = l2
							*(*int32)(unsafe.Add(mBase, _consts[307])) = l2
							if l2 == int32(3) {
								v156 = v102
							} else {
								v130 = m.G0
								v131 = int32(16)
								v132 = v130 - v131
								m.G0 = v132
								F___gettimeofday(m, v132)
								mBase = m.M
								v135 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
								v136 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+8)))
								m.G0 = v132 + v131
								*(*int64)(unsafe.Add(mBase, _consts[308])) = v136 + v135*int64(1000000) - int64(946684800000000)
								v156 = v102
							}
							m.G0 = v11 + int32(1264)
							return v156
						}
					} else {
						if l3 == int32(0) {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(96)
									F_errmsg(m, int32(298919), v11+int32(16))
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(492731), int32(4307), int32(465159))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
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
							v149 = *(*int32)(unsafe.Add(mBase, _consts[140]))
							if v149 != int32(44) {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(96)
										F_errmsg(m, int32(298919), v11+int32(16))
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(492731), int32(4307), int32(465159))
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
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
								v156 = int32(-1)
								m.G0 = v11 + int32(1264)
								return v156
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+440)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[199]))
		F_s_lock(m, v9+int32(440), int32(498539), int32(3760), int32(241006))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[199]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[199]))
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
	v1 = int32(4411584)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	*(*int32)(unsafe.Add(mBase, _consts[300])) = v3 + int32(1)
	return
}
func F_XLogPrefetcherNextBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int64
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int64
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int64
	_ = v468
	var v469 int32
	_ = v469
	var v471 int64
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v532 int32
	_ = v532
	var v534 int64
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v586 int64
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int64
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int64
	_ = v610
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v671 int32
	_ = v671
	var v672 int64
	_ = v672
	var v679 int32
	_ = v679
	var v680 int64
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	v24 = m.G0
	v25 = int32(48)
	v26 = v24 - v25
	m.G0 = v26
	v29 = l0 + int32(36)
	v31 = l0 + int32(84)
	v33 = l0 + int32(28)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+32))
	v57 = v35
	v62 = v34
	goto L4
L1:
	;
	m.G0 = v26 + int32(48)
	return v698
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L16
	} else {
		goto L167
	}
L3:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v672 = *(*int64)(unsafe.Add(mBase, uint32(v671)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v671)+48)) = v672 + int64(1)
	v698 = int32(0)
	goto L1
L4:
	;
	if v62 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v664)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v664)+32)) = v665 + int64(1)
	v698 = int32(0)
	goto L1
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v35)+120))
	if v80 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v119 = v62
	goto L8
L8:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	if base.Ui64(v120) <= base.Ui64(v36) {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	if v107 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v101 = F_XLogReadAhead(m, v57, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L20
	}
L11:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1256)))
	if v83 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v86 = int32(2)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui64(v36) <= base.Ui64(v87) {
		v698 = v86
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v90 = F_XLogReadAhead(m, v57, int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	if v90 != 0 {
		v105 = v90
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+124))
	if v95 == int32(0) {
		v698 = v86
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v95)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v98
	v698 = v86
	goto L1
L20:
	;
	if v101 != 0 {
		v105 = v101
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v698 = int32(2)
	goto L1
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v119 = v105
	goto L8
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	if int32(0) < v109 {
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
	v698 = int32(0)
	goto L1
L26:
	;
	goto L25
L27:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v119)+72))
	if v292 <= v293 {
		goto L74
	} else {
		goto L75
	}
L28:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+48)))
	v124 = v122 & int32(-16)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+49)))
	switch v125 {
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
	switch v124&int32(255) - int32(16) {
	case 0:
		goto L48
	default:
		goto L27
	case 16:
		goto L47
	}
L30:
	;
	if v124&int32(255) != 0 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	v127 = v124 & int32(255)
	if v127 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = base.B2i32(v127 != int32(144))
	goto L34
L33:
	;
	v131 = int32(0)
	goto L34
L34:
	;
	if v131 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v120
	goto L27
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)+64))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v137
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v148 = F_hash_search(m, v142, v26+int32(32), int32(1), v26+int32(47))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v120
	if v150 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = v154
	v157 = v148 + int32(28)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v158 == v154 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v170
	v173 = v148 + int32(28)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v174 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v162 = v33
	goto L43
L42:
	;
	v162 = v158
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v157
	goto L27
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v178 = v33
	goto L46
L45:
	;
	v178 = v174
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = int32(0)
	goto L27
L47:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v119)+64))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v238)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v250 = F_hash_search(m, v244, v26+int32(32), int32(1), v26+int32(47))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L16
	} else {
		goto L60
	}
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v119)+64))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	if v190 != 0 {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v191
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v201 = F_hash_search(m, v195, v26+int32(32), int32(1), v26+int32(47))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+16)) = v120
	if v203 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = v207
	v210 = v201 + int32(28)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v211 == v207 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v201)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v201)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v201)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v223
	v226 = v201 + int32(28)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v227 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v215 = v33
	goto L56
L55:
	;
	v215 = v211
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v201)+32)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v210
	goto L27
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v231 = v33
	goto L59
L58:
	;
	v231 = v227
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v201)+32)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v226
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = int32(0)
	goto L27
L60:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+16)) = v120
	if v252 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+24)) = v239
	v258 = v250 + int32(28)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v259 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v250)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v271
	v274 = v250 + int32(28)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v275 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v263 = v33
	goto L66
L65:
	;
	v263 = v259
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v250)+32)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v258
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v258
	goto L27
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v279 = v33
	goto L69
L68:
	;
	v279 = v275
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v250)+32)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v274
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v250)+24))
	if base.Ui32(v284) < base.Ui32(v239) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v286 = v284
	goto L72
L71:
	;
	v286 = v239
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+24)) = v286
	goto L27
L73:
	;
	goto L5
L74:
	;
	v299 = v292
	goto L77
L75:
	;
	goto L76
L76:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+124))
	if v652 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L77:
	;
	v320 = int32(1)
	v321 = v299 + v320
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v321
	v325 = v119 + int32(76) + v299*int32(52)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v326 == v320 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L76
L79:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v329
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	if v331 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v624 = v321
	goto L81
L81:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v119)+72))
	if v624 <= v626 {
		v299 = v624
		goto L77
	} else {
		goto L163
	}
L82:
	;
	v698 = int32(0)
	goto L1
L83:
	;
	goto L84
L84:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+29)))
	if v333 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v338)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v338)+40)) = v339 + int64(1)
	v698 = int32(0)
	goto L1
L86:
	;
	goto L87
L87:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+28)))
	if v343&int32(64) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v348)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+24)) = v349 + int64(1)
	v698 = int32(0)
	goto L1
L89:
	;
	goto L90
L90:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v354
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v325)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v358 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v384 = v325 + int32(4)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v385 != v386 {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	if v358 == v33 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v365 = int32(0)
	v367 = F_hash_search(m, v362, v26+int32(32), v365, v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	if v367 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v367)+24))
	if base.Ui32(v369) <= base.Ui32(v353) {
		goto L73
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v371
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v380 = F_hash_search(m, v375, v26+int32(32), v371, v371)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L16
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	if v380 != 0 {
		goto L73
	} else {
		goto L100
	}
L100:
	;
	goto L91
L101:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(88))))
	if v385 != v397 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v388 != v389 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v391 != v392 {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v394 == v395 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(92))))
	if v385 != v408 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(56))))
	if v399 != v400 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v402 != v403 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25)))
	if v405 == v406 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	goto L106
L111:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(96))))
	if v385 != v421 {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(68))))
	if v410 != v411 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64))))
	if v413 != v416 {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(60))))
	if v418 == v419 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v433 = int32(12)
	v435 = v29 + v432*v433
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v435))) = v436
	v439 = v325 + v433
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	*(*int32)(unsafe.Add(mBase, uint32(v435)+8)) = v440
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31+v442<<(uint(int32(2))%32)))) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v452 = base.I32_rem_s(v448+int32(1), int32(4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v454
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v456
	v461 = F_smgropen(m, v26+int32(16), int32(-1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L16
	} else {
		goto L121
	}
L117:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(80))))
	if v423 != v424 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v426 != v427 {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72))))
	if v429 == v430 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	goto L116
L121:
	;
	v464 = F_smgrexists(m, v461, int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	if v464 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v469
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v479 = F_hash_search(m, v473, v26+int32(32), int32(1), v26+int32(47))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L16
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	v527 = F_smgrnblocks(m, v461, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L16
	} else {
		goto L137
	}
L126:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v479)+16)) = v468
	if v481 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v520)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v520)+32)) = v521 + int64(1)
	v698 = int32(0)
	goto L1
L128:
	;
	v485 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v479)+24)) = v485
	v488 = v479 + int32(28)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v489 == v485 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v479)+28))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v479)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+4)) = v499
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v479)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v501
	v504 = v479 + int32(28)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v505 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v493 = v33
	goto L133
L132:
	;
	v493 = v489
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v479)+32)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v488
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v488
	goto L127
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v509 = v33
	goto L136
L135:
	;
	v509 = v505
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v479)+32)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = v504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v479)+24)) = int32(0)
	goto L127
L137:
	;
	if base.Ui32(v527) <= base.Ui32(v525) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v532
	v534 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v542 = F_hash_search(m, v536, v26+int32(32), int32(1), v26+int32(47))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L16
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	F_PrefetchSharedBuffer(m, v26+int32(32), v461, v592, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L16
	} else {
		goto L155
	}
L141:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+47)))
	*(*int64)(unsafe.Add(mBase, uint32(v542)+16)) = v531
	if v544 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v585)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v585)+32)) = v586 + int64(1)
	v698 = int32(0)
	goto L1
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+24)) = v530
	v550 = v542 + int32(28)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v551 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v542)+28))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v542)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v542)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v563
	v566 = v542 + int32(28)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v567 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v555 = v33
	goto L148
L147:
	;
	v555 = v551
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v542)+32)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v550
	goto L142
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v33
	v571 = v33
	goto L151
L150:
	;
	v571 = v567
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v542)+32)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v566
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v566
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v542)+24))
	if base.Ui32(v576) < base.Ui32(v530) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v578 = v576
	goto L154
L153:
	;
	v578 = v530
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+24)) = v578
	goto L142
L155:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v596 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v599)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v599)+16)) = v600 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = v596
	v698 = int32(0)
	goto L1
L157:
	;
	goto L158
L158:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+36)))
	if v605&int32(1) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v609)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v609)+8)) = v610 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = int32(0)
	v698 = int32(1)
	goto L1
L160:
	;
	goto L161
L161:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	if v618&int32(1) == int32(0) {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v624 = v623
	goto L81
L163:
	;
	goto L78
L164:
	;
	v659 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v659
	v57 = v651
	v62 = v659
	goto L4
L165:
	;
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v652)+16))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v655 != v656 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v698 = int32(2)
	goto L1
L167:
	;
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v461)))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v682
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v681
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v680
	F_errmsg_internal(m, int32(47582), v26)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L16
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(495864), int32(796), int32(317200))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L16
	} else {
		goto L169
	}
L169:
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v38 int32
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
	var v74 int32
	_ = v74
	var v76 int64
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
	var v125 int32
	_ = v125
	var v127 int64
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
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
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
		v32 = v25 + l1*int32(52) + int32(76)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
		if v33 != int32(1) {
			v50 = v6
		} else {
			if v17 != 0 {
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v32)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v38
			} else {
			}
			if v19 != 0 {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v40
			} else {
			}
			if v21 != 0 {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v42
			} else {
			}
			v44 = int32(1)
			if v23 == int32(0) {
				v50 = v44
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
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
			v197 = m.ExcPending
			if v197 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(373540), int32(0))
				mBase = m.M
				v201 = m.ExcPending
				if v201 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494341), int32(369), int32(461789))
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
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
				v210 = m.ExcPending
				if v210 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(421947), int32(0))
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494341), int32(371), int32(461789))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
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
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v74
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v13)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v76
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
							v92 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+(v84^int32(-1))<<(uint(int32(2))%32))))
							v106 = v98
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
								v223 = m.ExcPending
								if v223 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(2600))
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int32(0)
									} else {
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227
										F_errmsg_internal(m, int32(206200), v13)
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494341), int32(384), int32(461789))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
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
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v125
					v127 = *(*int64)(unsafe.Add(mBase, uint32(v13)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v127
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
											v152 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
											v166 = v158
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
											v152 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
											v166 = v158
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
									v152 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v152+(v148^int32(-1))<<(uint(int32(2))%32))))
									v166 = v158
								} else {
									v160 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
		v182 = m.ExcPending
		if v182 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
			F_errmsg_internal(m, int32(421892), v13+int32(48))
			mBase = m.M
			v188 = m.ExcPending
			if v188 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494341), int32(359), int32(461789))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
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
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v7 <= v1 {
	} else {
		v11 = v7 & int32(7)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[132]))
		if base.Ui32(int32(8)) <= base.Ui32(v7) {
			v19 = v1
			v20 = int32(0)
			for {
				v24 = int32(8260)
				v27 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+v19*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(1))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(2))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(3))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(4))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(5))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(6))*v24))) = uint8(v27)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+(v19|int32(7))*v24))) = uint8(v27)
				v78 = int32(8)
				v79 = v19 + v78
				v81 = v20 + v78
				if v81 != v7&int32(2147483640) {
					v19 = v79
					v20 = v81
					continue
				} else {
					break
				}
				break
			}
			v83 = v79
		} else {
			v83 = v1
		}
		if v11 == int32(0) {
		} else {
			v91 = v83
			v92 = int32(0)
			for {
				v99 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13+v91*int32(8260)))) = uint8(v99)
				v101 = int32(1)
				v104 = v92 + v101
				if v104 != v11 {
					v91 = v91 + v101
					v92 = v104
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, _consts[230])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[231])) = int32(4411552)
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v118
	*(*int32)(unsafe.Add(mBase, _consts[232])) = v118
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, _consts[233])) = uint8(v118)
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
	v2 = int32(4411556)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[94])))
	v5 = v4 | l0
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v5)
	return
}
