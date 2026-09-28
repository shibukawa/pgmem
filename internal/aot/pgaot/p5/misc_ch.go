package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckArchiveTimeout(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[0]))
	if v12 <= int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[1])))
		if v17 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[2]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+308))
			v25 = base.B2i32(v23 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[1])) = uint8(v25)
			v27 = v25
		} else {
			v27 = int32(0)
		}
		if v27 != 0 {
			m.G0 = v9 + int32(16)
			return
		} else {
			v28 = F_time(m)
			mBase = m.M
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[0]))
			v32 = *(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3]))
			if base.I32_wrap_i64(v28-v32) < v30 {
				m.G0 = v9 + int32(16)
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[4]))
				v43 = F_LWLockAcquire(m, v39+int32(1024), int32(1))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[2]))
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+240))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v46)+248))
					*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v48
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[4]))
					F_LWLockRelease(m, v51+int32(1024))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = int32(_a_F_CheckArchiveTimeout_0)
						v58 = *(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3]))
						if v47 < v58 {
							v60 = v58
						} else {
							v60 = v47
						}
						*(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3])) = v60
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[0]))
						if base.I32_wrap_i64(v28-v60) < v63 {
							m.G0 = v9 + int32(16)
							return
						} else {
							v67 = F_GetLastImportantRecPtr(m)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(v67) <= base.Ui64(v69) {
									*(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3])) = v28
									m.G0 = v9 + int32(16)
									return
								} else {
									v72 = F_RequestXLogSwitch(m, int32(1))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[5]))
										if v72&base.I64_extend_i32_s(v75-int32(1)) == int64(0) {
											*(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3])) = v28
											m.G0 = v9 + int32(16)
											return
										} else {
											v84 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												if v84 == int32(0) {
													*(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3])) = v28
													m.G0 = v9 + int32(16)
													return
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[0]))
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v89
													F_errmsg_internal(m, int32(_a_F_CheckArchiveTimeout_1), v9)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_CheckArchiveTimeout_2), int32(756), int32(_a_F_CheckArchiveTimeout_3))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, _c_F_CheckArchiveTimeout[3])) = v28
															m.G0 = v9 + int32(16)
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
}
func F_CheckDeadLockAlert(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[1])) = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[2]))
	v12 = base.AtomicRmwOr32(m, v1, int32(_a_F_CheckDeadLockAlert_0), v1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0])) = v3
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_F_CheckDeadLockAlert_0), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v20 == v16 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v23 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[3]))
	if v27 == v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[4]))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v57 = F_pgmem_kill(m, v23, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v31 + int32(16)
	goto L1
L10:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+15)) = uint8(v37)
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[5]))
	v45 = F_write(m, v41, v31+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v45 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0]))
	if v49 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_CheckDim_2(m *base.Module, l0 int32) {
	var v10 int32
	_ = v10
	Fn14208(m, l0, int32(79), int32(_a_F_CheckDim_2_0), int32(_a_F_CheckDim_2_1), int32(1000000000), int32(74), int32(_a_F_CheckDim_2_2), int32(1000000001))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_CheckDuplicateColumnOrPathNames(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = v3
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v29<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v35 == int32(4) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v239 = v29 + int32(1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v239 < v240 {
		v29 = v239
		goto L4
	} else {
		goto L59
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v141 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v40 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	F_CheckDuplicateColumnOrPathNames(m, l0, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L36
	}
L13:
	;
	v124 = F_lappend(m, v40, v39)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L29
	} else {
		goto L35
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v52 = int32(0)
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v46+v52<<(uint(int32(2))%32))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.B2i32(v64 == int32(0))|base.B2i32(v64 != v67) != 0 {
		v85 = v64
		v86 = v67
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	if v85-v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	v70 = v39
	v71 = v61
	goto L21
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v75
		v86 = v74
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v85 = v75
	v86 = v74
	goto L19
L23:
	;
	v78 = int32(1)
	if v75 == v74 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v89 = v52 + int32(1)
	if v89 != v43 {
		v52 = v89
		goto L16
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L17
L28:
	;
	goto L13
L29:
	;
	return
L30:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v99
	F_errmsg(m, int32(_a_F_CheckDuplicateColumnOrPathNames_0), v13)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	F_parser_errposition(m, v104, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_CheckDuplicateColumnOrPathNames_1), int32(190), int32(_a_F_CheckDuplicateColumnOrPathNames_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124
	goto L12
L36:
	;
	goto L6
L37:
	;
	v225 = F_lappend(m, v141, v140)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L29
	} else {
		goto L58
	}
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v144 <= int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v153 = int32(0)
	goto L40
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v147+v153<<(uint(int32(2))%32))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if base.B2i32(v165 == int32(0))|base.B2i32(v165 != v168) != 0 {
		v186 = v165
		v187 = v168
		goto L43
	} else {
		goto L44
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L29
	} else {
		goto L53
	}
L42:
	;
	if v186-v187 != 0 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	goto L42
L44:
	;
	v171 = v140
	v172 = v162
	goto L45
L45:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v176 == int32(0) {
		v186 = v176
		v187 = v175
		goto L43
	} else {
		goto L47
	}
L46:
	;
	v186 = v176
	v187 = v175
	goto L43
L47:
	;
	v179 = int32(1)
	if v176 == v175 {
		v171 = v171 + v179
		v172 = v172 + v179
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v190 = v153 + int32(1)
	if v190 != v144 {
		v153 = v190
		goto L40
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L41
L52:
	;
	goto L37
L53:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v199
	F_errmsg(m, int32(_a_F_CheckDuplicateColumnOrPathNames_0), v13+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	F_parser_errposition(m, v206, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L29
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CheckDuplicateColumnOrPathNames_1), int32(203), int32(_a_F_CheckDuplicateColumnOrPathNames_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L29
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v225
	goto L6
L59:
	;
	goto L5
}
func F_CheckPubRelationColumnList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l1 == v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L17
	} else {
		goto L24
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	m.G0 = v11 + int32(48)
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = int32(0)
	if v18 < v15 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = v15
	goto L8
L7:
	;
	v21 = v18
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v29 = v5
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22+v29<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v35 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L3
L11:
	;
	v44 = v29 + int32(1)
	if v44 != v21 {
		v29 = v44
		goto L9
	} else {
		goto L16
	}
L12:
	;
	if l2 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if l3 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
	if v40 == int32(112) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	goto L10
L17:
	;
	return
L18:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	v67 = F_get_namespace_name(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v70 + int32(4)
	F_errmsg(m, int32(_a_F_CheckPubRelationColumnList_0), v11)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v81 = F_errdetail(m, int32(_a_F_CheckPubRelationColumnList_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_CheckPubRelationColumnList_2), int32(813), int32(_a_F_CheckPubRelationColumnList_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
	v98 = F_get_namespace_name(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v101 + int32(4)
	F_errmsg(m, int32(_a_F_CheckPubRelationColumnList_0), v11+int32(32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_CheckPubRelationColumnList_4)
	v117 = F_errdetail(m, int32(_a_F_CheckPubRelationColumnList_5), v11+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CheckPubRelationColumnList_2), int32(828), int32(_a_F_CheckPubRelationColumnList_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CheckSelectLocking(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v11 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v62 = l1 - int32(1)
					if base.Ui32(v62) <= base.Ui32(int32(3)) {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
						v69 = v67
					} else {
						v69 = int32(_a_F_CheckSelectLocking_0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v69
					F_errmsg(m, int32(_a_F_CheckSelectLocking_1), v6+int32(80))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3413), int32(_a_F_CheckSelectLocking_3))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v12 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v89 = l1 - int32(1)
						if base.Ui32(v89) <= base.Ui32(int32(3)) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
							v96 = v94
						} else {
							v96 = int32(_a_F_CheckSelectLocking_0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v96
						F_errmsg(m, int32(_a_F_CheckSelectLocking_4), v6-int32(-64))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3420), int32(_a_F_CheckSelectLocking_3))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				if v13 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v89 = l1 - int32(1)
							if base.Ui32(v89) <= base.Ui32(int32(3)) {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
								v96 = v94
							} else {
								v96 = int32(_a_F_CheckSelectLocking_0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v96
							F_errmsg(m, int32(_a_F_CheckSelectLocking_4), v6-int32(-64))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3420), int32(_a_F_CheckSelectLocking_3))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v14 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v116 = l1 - int32(1)
								if base.Ui32(v116) <= base.Ui32(int32(3)) {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
									v123 = v121
								} else {
									v123 = int32(_a_F_CheckSelectLocking_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v123
								F_errmsg(m, int32(_a_F_CheckSelectLocking_5), v6+int32(48))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3427), int32(_a_F_CheckSelectLocking_3))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
						if v15 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									v143 = l1 - int32(1)
									if base.Ui32(v143) <= base.Ui32(int32(3)) {
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v143<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
										v150 = v148
									} else {
										v150 = int32(_a_F_CheckSelectLocking_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v150
									F_errmsg(m, int32(_a_F_CheckSelectLocking_6), v6)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3434), int32(_a_F_CheckSelectLocking_3))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
							if v18 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return
									} else {
										v168 = l1 - int32(1)
										if base.Ui32(v168) <= base.Ui32(int32(3)) {
											v173 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
											v175 = v173
										} else {
											v175 = int32(_a_F_CheckSelectLocking_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v175
										F_errmsg(m, int32(_a_F_CheckSelectLocking_7), v6+int32(16))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3441), int32(_a_F_CheckSelectLocking_3))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
								if v21 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return
										} else {
											v195 = l1 - int32(1)
											if base.Ui32(v195) <= base.Ui32(int32(3)) {
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v195<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
												v202 = v200
											} else {
												v202 = int32(_a_F_CheckSelectLocking_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v202
											F_errmsg(m, int32(_a_F_CheckSelectLocking_8), v6+int32(32))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3448), int32(_a_F_CheckSelectLocking_3))
												mBase = m.M
												v213 = m.ExcPending
												if v213 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									m.G0 = v6 + int32(112)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v35 = l1 - int32(1)
				if base.Ui32(v35) <= base.Ui32(int32(3)) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
					v42 = v40
				} else {
					v42 = int32(_a_F_CheckSelectLocking_0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v42
				F_errmsg(m, int32(_a_F_CheckSelectLocking_9), v6+int32(96))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3406), int32(_a_F_CheckSelectLocking_3))
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
}
func F_charin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v7 = F_strlen(m, v5)
	mBase = m.M
	if base.B2i32(v7 != int32(4))|base.B2i32(v6&int32(255) != int32(92)) != 0 {
		v41 = v6
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		if v15&int32(248) != int32(48) {
			v41 = int32(92)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
			if v21&int32(248) != int32(48) {
				v41 = int32(92)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
				if v28&int32(248) != int32(48) {
					v41 = int32(92)
				} else {
					v41 = v15<<(uint(int32(6))%32) + v21<<(uint(int32(3))%32) + v28 + int32(80)
				}
			}
		}
	}
	return base.I64_extend8_s(base.I64_extend_i32_u(v41))
}
func F_checkInsertTargets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return v266
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v19 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+48))
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v205)+120)))
	if v206 <= int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L60
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L31
	} else {
		goto L55
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L31
	} else {
		goto L50
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L31
	} else {
		goto L45
	}
L8:
	;
	v28 = v4
	v30 = v4
	v32 = v4
	goto L11
L9:
	;
	goto L10
L10:
	;
	v266 = l1
	goto L1
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v28<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+120)))
	if v41 < v45 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L10
L13:
	;
	if v100 == int32(0) {
		goto L7
	} else {
		goto L30
	}
L14:
	;
	v100 = v51 + int32(1)
	goto L13
L15:
	;
	goto L14
L16:
	;
	v51 = v41
	goto L19
L17:
	;
	goto L18
L18:
	;
	goto L26
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v60 = v53 + v54<<(uint(int32(3))%32) + v51*int32(100)
	v63 = F_namestrcmp(m, v60+int32(32), v40)
	mBase = m.M
	if v63 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	if v66 != int32(1) {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v70 = v51 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+120)))
	if v70 < v72 {
		v51 = v70
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L20
L26:
	;
	v100 = int32(0)
	goto L13
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v104 = F_bms_is_member(m, v100, v30)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v103 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v119 = F_lappend_int(m, v118, v100)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L31
	} else {
		goto L43
	}
L34:
	;
	if v104 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v104 != 0 {
		goto L5
	} else {
		goto L41
	}
L37:
	;
	v110 = F_bms_is_member(m, v100, v32)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v110 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v112 = F_bms_add_member(m, v30, v100)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v116 = v112
	v117 = v32
	goto L33
L41:
	;
	v114 = F_bms_add_member(m, v32, v100)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v116 = v30
	v117 = v114
	goto L33
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v119
	v123 = v28 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v123 < v124 {
		v28 = v123
		v30 = v116
		v32 = v117
		goto L11
	} else {
		goto L44
	}
L44:
	;
	goto L12
L45:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v146 + int32(4)
	F_errmsg(m, int32(_a_F_checkInsertTargets_0), v15)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1074), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L31
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v40
	F_errmsg(m, int32(_a_F_checkInsertTargets_3), v15+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1089), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v40
	F_errmsg(m, int32(_a_F_checkInsertTargets_3), v15+int32(32))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1100), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L31
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
	v212 = v4
	v213 = v4
	goto L61
L61:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+52))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v229 = v222 + v223<<(uint(int32(3))%32) + v212*int32(100)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+119)))
	if v230 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v266 = v259
	goto L1
L63:
	;
	if v258 != v206 {
		v212 = v258
		v213 = v259
		goto L61
	} else {
		goto L71
	}
L64:
	;
	v258 = v212 + int32(1)
	v259 = v213
	goto L63
L65:
	;
	goto L66
L66:
	;
	v236 = F_palloc0(m, int32(20))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = int32(81)
	v242 = F_pstrdup(m, v229+int32(32))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v242
	v249 = F_lappend(m, v213, v236)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v253 = v212 + int32(1)
	v254 = F_lappend_int(m, v251, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v254
	v258 = v253
	v259 = v249
	goto L63
L71:
	;
	goto L62
}
func F_check_acl(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(1033) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v5 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_acl_0), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_acl_1), int32(613), int32(_a_F_check_acl_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v8 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_check_acl_3), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_acl_1), int32(617), int32(_a_F_check_acl_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_check_acl_4), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_acl_1), int32(609), int32(_a_F_check_acl_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
}
func F_check_amop_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v19 = v17 + v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
			if l1 != v20 {
				v31 = int32(0)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+76)))
				if v22 != int32(98) {
					v31 = int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
					if v25 != l2 {
						v31 = int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
						if v28 == l3 {
							v31 = int32(1)
						} else {
							v31 = int32(0)
						}
					}
				}
			}
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v31
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_check_amop_signature_0), v9)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_check_amop_signature_1), int32(214), int32(_a_F_check_amop_signature_2))
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
func F_check_autovacuum_work_mem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v4 == int32(-1))|base.B2i32(int32(63) < v4) == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(64)
	} else {
	}
	return int32(1)
}
func F_check_backtrace_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_strlen(m, v12)
	mBase = m.M
	v14 = int32(_a_F_check_backtrace_functions_0)
	v18 = m.G0
	v20 = v18 - int32(32)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v21
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_backtrace_functions[0])))
	if v29 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v13 != v97 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v97 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_backtrace_functions[1])))
	if v33 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v37 = v12
	goto L8
L6:
	;
	goto L7
L7:
	;
	v47 = v14
	v48 = v29
	goto L11
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v43 == v29 {
		v37 = v37 + int32(1)
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v97 = v37 - v12
	goto L1
L10:
	;
	goto L9
L11:
	;
	v55 = v20 + int32(base.Ui32(v48)>>(uint(int32(3))%32))&int32(28)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56 | v57<<(uint(v48)%32)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v61 != 0 {
		v47 = v47 + v57
		v48 = v61
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v64 == int32(0) {
		v87 = v12
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v97 = v87 - v12
	goto L1
L15:
	;
	v68 = v12
	v69 = v64
	goto L16
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(base.Ui32(v69)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v77)>>(uint(v69)%32))&int32(1) == int32(0) {
		v87 = v68
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v87 = v85
	goto L14
L18:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v85 = v68 + int32(1)
	if v83 != 0 {
		v68 = v85
		v69 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[3])) = v101
	v106 = F_format_elog_string(m, int32(_a_F_check_backtrace_functions_1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v113 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[4])) = v106
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(1)
L26:
	;
	goto L27
L27:
	;
	v122 = F_guc_malloc(m, v13+int32(2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	if v122 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L31
L31:
	;
	if v13 <= int32(0) {
		v204 = v4
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v213 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v204+v122))) = uint16(v213)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v122
	return int32(1)
L33:
	;
	if v13 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v139 = v4
	v142 = v4
	v143 = v4
	goto L37
L35:
	;
	v183 = v4
	v186 = v4
	goto L36
L36:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v186))))
	switch v193 - int32(9) {
	case 0, 1, 23:
		v204 = v183
		goto L32
	default:
		goto L48
	case 35:
		v196 = v4
		goto L47
	}
L37:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v142))))
	switch v150 - int32(9) {
	case 0, 1, 23:
		v159 = v139
		goto L39
	default:
		goto L41
	case 35:
		v153 = int32(0)
		goto L40
	}
L38:
	;
	if v13&int32(1) == int32(0) {
		v204 = v172
		goto L32
	} else {
		goto L46
	}
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v142)+1)))
	switch v163 - int32(9) {
	case 0, 1, 23:
		v172 = v159
		goto L42
	default:
		goto L44
	case 35:
		v166 = int32(0)
		goto L43
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v139+v122))) = uint8(v153)
	v159 = v139 + int32(1)
	goto L39
L41:
	;
	v153 = v150
	goto L40
L42:
	;
	v173 = int32(2)
	v174 = v142 + v173
	v176 = v143 + v173
	if v176 != v13&int32(2147483646) {
		v139 = v172
		v142 = v174
		v143 = v176
		goto L37
	} else {
		goto L45
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v159+v122))) = uint8(v166)
	v172 = v159 + int32(1)
	goto L42
L44:
	;
	v166 = v163
	goto L43
L45:
	;
	goto L38
L46:
	;
	v183 = v172
	v186 = v174
	goto L36
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v122))) = uint8(v196)
	v204 = v183 + int32(1)
	goto L32
L48:
	;
	v196 = v193
	goto L47
}
func F_check_duplicates_in_publist(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v17
	v25 = v3
	v26 = v3
	goto L4
L4:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v24 <= v30 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	if l1 != 0 {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	v41 = v30
	goto L8
L8:
	;
	if v41 == v25 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31+v41<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.B2i32(v57 == int32(0))|base.B2i32(v57 != v60) != 0 {
		v78 = v57
		v79 = v60
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v78-v79 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v63 = v36
	v64 = v54
	goto L14
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v68 == int32(0) {
		v78 = v68
		v79 = v67
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v78 = v68
	v79 = v67
	goto L12
L16:
	;
	v71 = int32(1)
	if v68 == v67 {
		v63 = v63 + v71
		v64 = v64 + v71
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v82 = v41 + int32(1)
	if v82 == v24 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L9
L21:
	;
	v41 = v82
	goto L8
L22:
	;
	return
L23:
	;
	F_errcode(m, int32(_a_F_check_duplicates_in_publist_0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v54
	F_errmsg(m, int32(_a_F_check_duplicates_in_publist_1), v13)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_check_duplicates_in_publist_2), int32(3532), int32(_a_F_check_duplicates_in_publist_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v113 = F_cstring_to_text(m, v36)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L22
	} else {
		goto L30
	}
L28:
	;
	v119 = v26
	goto L29
L29:
	;
	v121 = v25 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v121 < v122 {
		v24 = v122
		v25 = v121
		v26 = v119
		goto L4
	} else {
		goto L31
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v113)
	v119 = v26 + int32(1)
	goto L29
L31:
	;
	goto L5
}
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) == int32(0) {
		if base.B2i32(l0 != l3)&base.B2i32(int32(0) <= l3) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				v53 = int32(3)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v53)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v55
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v53)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v59
				F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_0), v7+int32(-48))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1807), int32(_a_F_check_encoding_conversion_args_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if base.B2i32(l1 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l1)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
					F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_3), v7+int32(-32))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1809), int32(_a_F_check_encoding_conversion_args_2))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if base.B2i32(l1 != l4)&base.B2i32(int32(0) <= l4) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						v90 = int32(3)
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v90)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v92
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l4<<(uint(v90)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v96
						F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_4), v7+int32(-16))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1813), int32(_a_F_check_encoding_conversion_args_2))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					if l2 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_5), int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1815), int32(_a_F_check_encoding_conversion_args_2))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.G0 = v9 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_6), v9)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1803), int32(_a_F_check_encoding_conversion_args_2))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
func F_check_exclusion_or_unique_constraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v214 int32
	_ = v214
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v433 int32
	_ = v433
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v600 int32
	_ = v600
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v629 int32
	_ = v629
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	v25 = m.G0
	v27 = v25 - int32(2272)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+10)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = int32(100)
	goto L3
L2:
	;
	v35 = int32(112)
	goto L3
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35)))
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = int32(96)
	goto L6
L5:
	;
	v40 = int32(108)
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2+v40)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+124)))
	if v43 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L12
	} else {
		goto L156
	}
L8:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	if v137|base.B2i32(v31 <= int32(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	v47 = v31 - int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v47))))
	if v49 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2+v47<<(uint(int32(1))%32))+12)))
	v63 = v50 + v51<<(uint(int32(3))%32) + v58*int32(100) - int32(72)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v66 = F_lookup_type_cache(m, v64, int32(_a_F_check_exclusion_or_unique_constraint_0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l4+v47<<(uint(int32(3))%32))))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v63)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+536)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v63)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+528)) = v81
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v63)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+520)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v63)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+512)) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v63)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+504)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v63)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+496)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v63)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+488)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+480)) = v93
	switch v74&int32(255) - int32(109) {
	case 0:
		goto L19
	default:
		goto L20
	case 5:
		goto L18
	}
L12:
	;
	return int32(0)
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+300))
	if v70 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = F_get_typtype(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+13)))
	v74 = v73
	goto L11
L17:
	;
	v74 = v71
	goto L11
L18:
	;
	v121 = F_pg_detoast_datum(m, base.I32_wrap_i64(v78))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L26
	}
L19:
	;
	v117 = F_pg_detoast_datum(m, base.I32_wrap_i64(v78))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L24
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v27 + int32(480)
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_1), v27+int32(80))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_2), int32(1199), int32(_a_F_check_exclusion_or_unique_constraint_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v119 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L7
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v121+int32(base.Ui32(v123)>>(uint(int32(2))%32))-int32(1)))))
	goto L27
L27:
	;
	if v129&int32(1) != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L12
	} else {
		goto L153
	}
L30:
	;
	v646 = F_BuildIndexValueDescription(m, l1, l4, l5)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L12
	} else {
		goto L128
	}
L31:
	;
	m.G0 = v27 + int32(2272)
	return v629
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+408)) = int32(4)
	if int32(0) < v31 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v153 = int32(0)
	goto L34
L34:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v153))))
	if v167 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v629 = int32(1)
	goto L31
L36:
	;
	v171 = v153 + int32(1)
	if v31 != v171 {
		v153 = v171
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L32
L40:
	;
	v214 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v283 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L50
	}
L43:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v214))))
	if v235 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v236 = int32(65)
	goto L47
L46:
	;
	v236 = int32(0)
	goto L47
L47:
	;
	v237 = int32(1)
	v238 = v214 + v237
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+v214<<(uint(v237)%32)))))
	v246 = v214 << (uint(int32(2)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v29+v246)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246+v42)))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(l4+v214<<(uint(int32(3))%32))))
	F_ScanKeyEntryInitialize(m, v27+int32(480)+v214*int32(56), v236, base.I32_extend16_s(v238), v243, int32(0), v248, v250, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	if v238 != v31 {
		v214 = v238
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l6)+152))
	if v285 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v288 = F_MakePerTupleExprContext(m, l6)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L54
	}
L52:
	;
	v290 = v285
	goto L53
L53:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+4)) = v283
	v294 = v283 + int32(32)
	goto L56
L54:
	;
	v290 = v288
	goto L53
L55:
	;
	F_index_endscan(m, v325)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L12
	} else {
		goto L126
	}
L56:
	;
	v319 = int32(0)
	v325 = F_index_beginscan(m, l0, l1, v27+int32(408), v319, v31, v319, v319)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L12
	} else {
		goto L58
	}
L57:
	;
	if l9 == int32(0) {
		goto L30
	} else {
		goto L117
	}
L58:
	;
	v329 = int32(0)
	F_index_rescan(m, v325, v27+int32(480), v31, v329, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	v334 = F_index_getnext_slot(m, v325, int32(1), v283)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	if v334 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v600 = int32(1)
	goto L55
L62:
	;
	goto L63
L63:
	;
	v352 = v319
	goto L64
L64:
	;
	if l3 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v27)+412))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v27)+416))
	if v524 != 0 {
		goto L94
	} else {
		goto L95
	}
L66:
	;
	goto L65
L67:
	;
	v496 = int32(1)
	v498 = F_index_getnext_slot(m, v325, v496, v283)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L91
	}
L68:
	;
	F_FormIndexDatum(m, l2, v283, l6, v27+int32(144), v27+int32(112))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L12
	} else {
		goto L82
	}
L69:
	;
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v365 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v370 = int32(16)
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+2)))
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294))))
	if v368|v369<<(uint(v370)%32) == v373|v374<<(uint(v370)%32) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if v384 == int32(0) {
		goto L68
	} else {
		goto L77
	}
L72:
	;
	goto L71
L73:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+4)))
	if v380 == v381 {
		v384 = int32(1)
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v384 = int32(0)
	goto L72
L76:
	;
	goto L75
L77:
	;
	if v352 == int32(0) {
		v485 = int32(1)
		goto L67
	} else {
		goto L78
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v394 + int32(4)
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_4), v27-int32(-64))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_2), int32(848), int32(_a_F_check_exclusion_or_unique_constraint_5))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+72)))
	if v414 != int32(1) {
		goto L66
	} else {
		goto L83
	}
L83:
	;
	v417 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v419 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418)+10)))
	if v419 <= v417 {
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v433 = v417
	goto L85
L85:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(112)+v433))))
	if v449 != 0 {
		v485 = v352
		goto L67
	} else {
		goto L87
	}
L86:
	;
	goto L66
L87:
	;
	v451 = v433 << (uint(int32(2)) % 32)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v42+v451)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v454+v451)))
	v458 = v433 << (uint(int32(3)) % 32)
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v458+(v27+int32(144)))))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(l4+v458)))
	v465 = F_OidFunctionCall2Coll(m, v453, v456, v462, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	if v465 == int64(0) {
		v485 = v352
		goto L67
	} else {
		goto L89
	}
L89:
	;
	v470 = v433 + int32(1)
	if v419 != v470 {
		v433 = v470
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	if v498 != 0 {
		v352 = v485
		goto L64
	} else {
		goto L92
	}
L92:
	;
	v600 = v496
	goto L55
L93:
	;
	goto L57
L94:
	;
	v526 = v524
	goto L96
L95:
	;
	v526 = v525
	goto L96
L96:
	;
	if v526 == int32(0) {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	if l8 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	F_index_endscan(m, v325)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L12
	} else {
		goto L108
	}
L99:
	;
	if l8 != int32(2) {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v27)+444))
	if v533 == int32(0) {
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v538 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	if base.B2i32(base.Ui32(v526) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v538) < base.Ui32(int32(3))) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v538-v526 < int32(0) {
		goto L98
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if base.Ui32(v526) <= base.Ui32(v538) {
		goto L93
	} else {
		goto L107
	}
L106:
	;
	goto L93
L107:
	;
	goto L98
L108:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v27)+444))
	if v553 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v27)+412))
	F_SpeculativeInsertionWait(m, v554, v553)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v550 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L56
L113:
	;
	v559 = int32(8)
	goto L115
L114:
	;
	v559 = int32(5)
	goto L115
L115:
	;
	F_XactLockTableWait(m, v526, l0, v294, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	goto L56
L117:
	;
	if l10 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v600 = int32(0)
	goto L55
L119:
	;
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l10)+4)) = uint16(v567)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v569
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_check_exclusion_or_unique_constraint[0]))
	if v572 != int32(3) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_check_exclusion_or_unique_constraint[1]))
	if v576 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_exclusion_or_unique_constraint[2])))
	if v578&int32(1) == int32(0) {
		goto L29
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)+60))
	v586 = m.T0[v585].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l10, v583, v283)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L12
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	goto L118
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+4)) = v291
	F_ExecDropSingleTupleTableSlot(m, v283)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	v629 = v600
	goto L31
L128:
	;
	v652 = F_BuildIndexValueDescription(m, l1, v27+int32(144), v27+int32(112))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	F_errcode(m, int32(16908482))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	v661 = int32(0)
	v665 = base.B2i32(v646 != v661) & base.B2i32(v652 != v661)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v668 = v666 + int32(4)
	if l7 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v668
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_6), v27+int32(16))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L12
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v668
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_7), v27+int32(48))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L12
	} else {
		goto L144
	}
L135:
	;
	if v665 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v684+int32(4))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L12
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v646
	v678 = F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_8), v27)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L12
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v682 = F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_9), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L12
	} else {
		goto L141
	}
L140:
	;
	goto L136
L141:
	;
	goto L136
L142:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_2), int32(947), int32(_a_F_check_exclusion_or_unique_constraint_5))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	if v665 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v711+int32(4))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L12
	} else {
		goto L151
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v646
	v705 = F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_10), v27+int32(32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L12
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v709 = F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_11), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L12
	} else {
		goto L150
	}
L149:
	;
	goto L145
L150:
	;
	goto L145
L151:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_2), int32(958), int32(_a_F_check_exclusion_or_unique_constraint_5))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_12), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_13), int32(1355), int32(_a_F_check_exclusion_or_unique_constraint_14))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L12
	} else {
		goto L157
	}
L157:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v741 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v27 + int32(480)
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_15), v27+int32(96))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_2), int32(1207), int32(_a_F_check_exclusion_or_unique_constraint_3))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v11 - int32(9) {
	case 0:
		goto L9
	default:
		goto L2
	case 2:
		goto L8
	case 6:
		goto L7
	case 8, 9, 10:
		goto L6
	case 11:
		goto L5
	case 19:
		goto L4
	case 28:
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v123
L2:
	;
	v123 = int32(0)
	goto L1
L3:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v86 == int32(0) {
		goto L2
	} else {
		goto L38
	}
L4:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_getTypeInputInfo(m, v60, v9+int32(12), v9+int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L29
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v47 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v34 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L15
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v22, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L13
	}
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v14, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v15 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v123 = int32(1)
	goto L1
L13:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v123 = int32(1)
	goto L1
L15:
	;
	if v29 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v123 = int32(1)
	goto L1
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = F_get_opcode(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v41 = v34
	goto L19
L19:
	;
	v42 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v41, l2)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
	v41 = v38
	goto L19
L21:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v123 = int32(1)
	goto L1
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = F_get_opcode(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	v54 = v47
	goto L25
L25:
	;
	v55 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v54, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v51
	v54 = v51
	goto L25
L27:
	;
	if v55 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v123 = int32(1)
	goto L1
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v68 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v67, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v123 = int32(1)
	goto L1
L32:
	;
	goto L33
L33:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = F_exprType(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_getTypeOutputInfo(m, v72, v9+int32(12), v9+int32(7))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v81 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v80, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	if v81 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v123 = int32(1)
	goto L1
L38:
	;
	v89 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v90 <= v89 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v94 = v89
	goto L40
L40:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v94<<(uint(int32(2))%32))))
	v105 = F_get_opcode(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L42
	}
L41:
	;
	goto L2
L42:
	;
	v107 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v105, l2)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	if v107 != 0 {
		v123 = int32(1)
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v110 = v94 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v110 < v111 {
		v94 = v110
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
}
func F_check_mcvlist_array(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != int32(1) {
		v16 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v89 = v4
				m.G0 = v8 + int32(48)
				return v89
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_check_mcvlist_array[0])))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v27
					F_errmsg(m, int32(_a_F_check_mcvlist_array_0), v8+int32(32))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v83 = v4
						v84 = int32(781)
						F_errfinish(m, int32(_a_F_check_mcvlist_array_1), v84, int32(_a_F_check_mcvlist_array_2))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v89 = v83
							m.G0 = v8 + int32(48)
							return v89
						}
					}
				}
			}
		}
	} else {
		v37 = F_array_contains_nulls(m, l0)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			if v37 != 0 {
				v41 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v89 = v4
						m.G0 = v8 + int32(48)
						return v89
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_check_mcvlist_array[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v50
							F_errmsg(m, int32(_a_F_check_mcvlist_array_3), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v83 = v4
								v84 = int32(790)
								F_errfinish(m, int32(_a_F_check_mcvlist_array_1), v84, int32(_a_F_check_mcvlist_array_2))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = v83
									m.G0 = v8 + int32(48)
									return v89
								}
							}
						}
					}
				}
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v57 == l2 {
					v89 = int32(1)
					m.G0 = v8 + int32(48)
					return v89
				} else {
					v59 = int32(0)
					v62 = F_errstart(m, int32(19), v59)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						if v62 == int32(0) {
							v89 = v59
							m.G0 = v8 + int32(48)
							return v89
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_check_mcvlist_array[0])))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v71
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_check_mcvlist_array[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v74
								F_errmsg(m, int32(_a_F_check_mcvlist_array_4), v8+int32(16))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v83 = v59
									v84 = int32(800)
									F_errfinish(m, int32(_a_F_check_mcvlist_array_1), v84, int32(_a_F_check_mcvlist_array_2))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v89 = v83
										m.G0 = v8 + int32(48)
										return v89
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
func F_check_nested_generated_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v91 = v3
		m.G0 = v9 + int32(16)
		return v91
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 == int32(6) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32)-int32(4))))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			if v25 == int32(0) {
				v91 = v3
				m.G0 = v9 + int32(16)
				return v91
			} else {
				v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
				if int32(0) < v28 {
					v31 = F_get_attgenerated(m, v25, v28)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							v91 = v3
							m.G0 = v9 + int32(16)
							return v91
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v45 = F_get_attname(m, v25, v28, int32(0))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
										F_errmsg(m, int32(_a_F_check_nested_generated_walker_0), v9)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											v53 = F_errdetail(m, int32(_a_F_check_nested_generated_walker_1), int32(0))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												F_parser_errposition(m, l1, v55)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_check_nested_generated_walker_2), int32(3243), int32(_a_F_check_nested_generated_walker_3))
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
						}
					}
				} else {
					if v28 != 0 {
						v91 = v3
						m.G0 = v9 + int32(16)
						return v91
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_check_nested_generated_walker_4), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v76 = F_errdetail(m, int32(_a_F_check_nested_generated_walker_5), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										F_parser_errposition(m, l1, v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_nested_generated_walker_2), int32(3250), int32(_a_F_check_nested_generated_walker_3))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
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
			}
		} else {
			v87 = F_expression_tree_walker_impl(m, l0, int32(499), l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v91 = v87
				m.G0 = v9 + int32(16)
				return v91
			}
		}
	}
}
func F_check_publications_origin_tables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = int32(25)
	if l4 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L27
	} else {
		goto L101
	}
L2:
	;
	m.G0 = v12 + int32(96)
	return
L3:
	;
	v70 = v12 + int32(80)
	F_initStringInfo(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v68 = int32(0)
	goto L3
L5:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v21 = l4
	v22 = int32(_a_F_check_publications_origin_tables_0)
	goto L10
L8:
	;
	goto L2
L9:
	;
	if v59 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == v26 {
		v48 = v25
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v59 = int32(0)
	goto L9
L12:
	;
	v50 = int32(1)
	if v48 != 0 {
		v21 = v21 + v50
		v22 = v22 + v50
		goto L10
	} else {
		goto L21
	}
L13:
	;
	if base.Ui32((v25-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = v25 | int32(32)
	goto L16
L15:
	;
	v36 = v25
	goto L16
L16:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v45 = v26 | int32(32)
	goto L19
L18:
	;
	v45 = v26
	goto L19
L19:
	;
	if v36 == v45 {
		v48 = v36
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v59 = v36 - v45
	goto L9
L21:
	;
	goto L11
L22:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if l3 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v68 = int32(1)
	goto L3
L26:
	;
	goto L4
L27:
	;
	return
L28:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_check_publications_origin_tables_1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_GetPublicationsStr(m, l1, v70, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_check_publications_origin_tables_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v82 = int32(0)
	if base.B2i32(v68 == v82)|base.B2i32(l6 <= v82) == v82 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v94 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications_origin_tables[0]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+60))
	v151 = m.T0[v150].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v144, int32(1), v12+int32(76))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L27
	} else {
		goto L49
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5+v94<<(uint(int32(2))%32))))
	v103 = F_get_rel_name(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L27
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	v133 = v94 + int32(1)
	if v133 != l6 {
		v94 = v133
		goto L35
	} else {
		goto L48
	}
L38:
	;
	if v103 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v107 = F_get_rel_namespace(m, v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v109 = F_get_namespace_name(m, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	if v109 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v113 = F_quote_literal_cstr(m, v109)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v115 = F_quote_literal_cstr(m, v103)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v113
	F_appendStringInfo(m, v12+int32(80), int32(_a_F_check_publications_origin_tables_3), v12+int32(48))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v113)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v115)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	goto L37
L48:
	;
	goto L36
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	F_pfree(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v156 != int32(2) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v161 = F_MakeSingleTupleTableSlot(m, v159, int32(_a_F_check_publications_origin_tables_4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L27
	} else {
		goto L52
	}
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v166 = F_tuplestore_gettupleslot(m, v163, int32(1), int32(0), v161)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L27
	} else {
		goto L54
	}
L53:
	;
	F_ExecDropSingleTupleTableSlot(m, v161)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L27
	} else {
		goto L87
	}
L54:
	;
	if v166 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v174 = int32(0)
	goto L56
L56:
	;
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v161)+6)))
	if v180 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v198 == int32(0) {
		goto L53
	} else {
		goto L68
	}
L58:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	m.T0[v185].(func(*base.Module, int32, int32))(m, v161, int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L27
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = F_text_to_cstring(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L27
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	m.T0[v193].(func(*base.Module, int32))(m, v161)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L27
	} else {
		goto L63
	}
L63:
	;
	v196 = F_makeString(m, v190)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	v198 = F_list_append_unique(m, v174, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L27
	} else {
		goto L65
	}
L65:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v203 = F_tuplestore_gettupleslot(m, v200, int32(1), int32(0), v161)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L27
	} else {
		goto L66
	}
L66:
	;
	if v203 != 0 {
		v174 = v198
		goto L56
	} else {
		goto L67
	}
L67:
	;
	goto L57
L68:
	;
	v208 = v12 + int32(60)
	F_initStringInfo(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	F_GetPublicationsStr(m, v198, v208, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	v216 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	if v216 == int32(0) {
		goto L53
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L27
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l7
	if v68 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v226 = int32(_a_F_check_publications_origin_tables_5)
	goto L76
L75:
	;
	v226 = int32(_a_F_check_publications_origin_tables_6)
	goto L76
L76:
	;
	F_errmsg(m, v226, v12+int32(16))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L27
	} else {
		goto L77
	}
L77:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v232
	F_errdetail_plural(m, int32(_a_F_check_publications_origin_tables_7), int32(_a_F_check_publications_origin_tables_8), v231, v12)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L27
	} else {
		goto L78
	}
L78:
	;
	if v68 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v240 = int32(_a_F_check_publications_origin_tables_9)
	goto L81
L80:
	;
	v240 = int32(_a_F_check_publications_origin_tables_10)
	goto L81
L81:
	;
	F_errhint(m, v240, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L27
	} else {
		goto L82
	}
L82:
	;
	if v68 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v247 = int32(3062)
	goto L85
L84:
	;
	v247 = int32(3071)
	goto L85
L85:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_tables_11), v247, int32(_a_F_check_publications_origin_tables_12))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L27
	} else {
		goto L86
	}
L86:
	;
	goto L53
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	if v262 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_pfree(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L27
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v265 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	F_tuplestore_end(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L27
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	if v268 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	F_FreeTupleDesc(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L27
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_pfree(m, v151)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L27
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	goto L2
L101:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L27
	} else {
		goto L102
	}
L102:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v292
	F_errmsg(m, int32(_a_F_check_publications_origin_tables_13), v12+int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L27
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_tables_11), int32(3019), int32(_a_F_check_publications_origin_tables_12))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L27
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_session_authorization(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v4 {
		v145 = int32(1)
		m.G0 = v11 - int32(-64)
		return v145
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_session_authorization[0])))
		if v18 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[1]))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_session_authorization[2])))
			v129 = v22
			v130 = v24
			v133 = F_guc_malloc(m, int32(8))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return int32(0)
			} else {
				if v133 == int32(0) {
					v145 = v4
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v133))) = v129
					v138 = int32(1)
					v140 = v130 & v138
					*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)) = uint8(v140)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v133
					v145 = v138
				}
				m.G0 = v11 - int32(-64)
				return v145
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[3]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
			if base.B2i32(v27 == int32(2)) == int32(0) {
				v145 = v4
				m.G0 = v11 - int32(-64)
				return v145
			} else {
				v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
				v34 = F_SearchSysCache1(m, int32(10), v33)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						if l2 == int32(12) {
							v42 = int32(1)
							v45 = F_errstart(m, int32(18), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								if v45 == int32(0) {
									v145 = v42
									m.G0 = v11 - int32(-64)
									return v145
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
										F_errmsg(m, int32(_a_F_check_session_authorization_0), v11)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_session_authorization_1), int32(864), int32(_a_F_check_session_authorization_2))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v145 = v42
												m.G0 = v11 - int32(-64)
												return v145
											}
										}
									}
								}
							}
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[5])) = v63
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v66
							v72 = F_format_elog_string(m, int32(_a_F_check_session_authorization_0), v9+int32(-48))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[6])) = v72
								v145 = v4
								m.G0 = v11 - int32(-64)
								return v145
							}
						}
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
						v77 = v75 + v76
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+68)))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
						F_ReleaseCatCache(m, v34)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[7]))
							if v79 == v83 {
								v129 = v79
								v130 = v78
								v133 = F_guc_malloc(m, int32(8))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return int32(0)
								} else {
									if v133 == int32(0) {
										v145 = v4
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v129
										v138 = int32(1)
										v140 = v130 & v138
										*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)) = uint8(v140)
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v133
										v145 = v138
									}
									m.G0 = v11 - int32(-64)
									return v145
								}
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[7]))
								v87 = F_superuser_arg(m, v86)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									if v87 != 0 {
										v129 = v79
										v130 = v78
										v133 = F_guc_malloc(m, int32(8))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											if v133 == int32(0) {
												v145 = v4
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v129
												v138 = int32(1)
												v140 = v130 & v138
												*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)) = uint8(v140)
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v133
												v145 = v138
											}
											m.G0 = v11 - int32(-64)
											return v145
										}
									} else {
										if l2 == int32(12) {
											v91 = int32(1)
											v94 = F_errstart(m, int32(18), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												if v94 == int32(0) {
													v145 = v91
													m.G0 = v11 - int32(-64)
													return v145
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v101
														F_errmsg(m, int32(_a_F_check_session_authorization_3), v9+int32(-32))
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_check_session_authorization_1), int32(890), int32(_a_F_check_session_authorization_2))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return int32(0)
															} else {
																v145 = v91
																m.G0 = v11 - int32(-64)
																return v145
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[8])) = int32(16797828)
											v117 = *(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[5])) = v117
											v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v120
											v126 = F_format_elog_string(m, int32(_a_F_check_session_authorization_4), v9+int32(-16))
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_check_session_authorization[6])) = v126
												v145 = v4
												m.G0 = v11 - int32(-64)
												return v145
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
func F_check_srf_call_placement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v11 - int32(2) {
	case 0, 1:
		v52 = int32(_a_F_check_srf_call_placement_0)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 2, 4, 5, 6, 14, 15, 20, 21, 22, 23, 24, 42:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if base.Ui32(v78) <= base.Ui32(int32(44)) {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_c_F_check_srf_call_placement[0])))
					v85 = v83
				} else {
					v85 = int32(_a_F_check_srf_call_placement_4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v85
				F_errmsg(m, int32(_a_F_check_srf_call_placement_5), v8+int32(16))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2865), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
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
	case 3:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v14 == l1 {
			m.G0 = v8 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_srf_call_placement_6), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v28 = F_exprLocation(m, v27)
						mBase = m.M
						F_parser_errposition(m, l0, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2742), int32(_a_F_check_srf_call_placement_3))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
		}
	case 7, 8, 9, 10, 11:
		v52 = int32(_a_F_check_srf_call_placement_7)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 12, 13, 17, 18, 19, 25:
		v99 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)) = uint8(v99)
		m.G0 = v8 + int32(32)
		return
	case 16:
		v52 = int32(_a_F_check_srf_call_placement_8)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 26, 27:
		v52 = int32(_a_F_check_srf_call_placement_9)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 28, 29:
		v52 = int32(_a_F_check_srf_call_placement_10)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 30:
		v52 = int32(_a_F_check_srf_call_placement_11)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 31:
		v52 = int32(_a_F_check_srf_call_placement_12)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 32:
		v52 = int32(_a_F_check_srf_call_placement_13)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 33:
		v52 = int32(_a_F_check_srf_call_placement_14)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 34:
		v52 = int32(_a_F_check_srf_call_placement_15)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 35:
		v52 = int32(_a_F_check_srf_call_placement_16)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 36:
		v52 = int32(_a_F_check_srf_call_placement_17)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 37:
		v52 = int32(_a_F_check_srf_call_placement_18)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 38:
		v52 = int32(_a_F_check_srf_call_placement_19)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 39:
		v52 = int32(_a_F_check_srf_call_placement_20)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 40:
		v52 = int32(_a_F_check_srf_call_placement_21)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	case 41:
		v52 = int32(_a_F_check_srf_call_placement_22)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
				F_errmsg_internal(m, int32(_a_F_check_srf_call_placement_1), v8)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2858), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
	default:
		m.G0 = v8 + int32(32)
		return
	}
}
func F_check_ssl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_ssl[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_ssl[1])) = v8
		v14 = F_format_elog_string(m, int32(_a_F_check_ssl_0), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_ssl[2])) = v14
			return v4 ^ int32(1)
		}
	} else {
		return v4 ^ int32(1)
	}
}
func F_checkint(m *base.Module, l0 int64) int32 {
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v9 = base.I32_wrap_i64(int64(base.Ui64(l0)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(v9) < base.Ui32(int32(1023)) {
		v33 = int32(0)
	} else {
		if base.Ui32(int32(1075)) < base.Ui32(v9) {
			v33 = int32(2)
		} else {
			v16 = int64(1)
			v20 = v16 << (uint(base.I64_extend_i32_u(int32(1075)-v9)) % 64)
			if (v20-v16)&l0 != int64(0) {
				v33 = int32(0)
			} else {
				if l0&v20 == int64(0) {
					v31 = int32(2)
				} else {
					v31 = int32(1)
				}
				v33 = v31
			}
		}
	}
	return v33
}
