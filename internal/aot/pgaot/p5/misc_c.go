package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CMPTRGM_SIGNED(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		v23 = v5
		v24 = v6
		if base.I32_extend8_s(v23) < base.I32_extend8_s(v24) {
			v30 = int32(-1)
		} else {
			v30 = int32(1)
		}
		return v30
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v8 != v9 {
			v23 = v8
			v24 = v9
			if base.I32_extend8_s(v23) < base.I32_extend8_s(v24) {
				v30 = int32(-1)
			} else {
				v30 = int32(1)
			}
			return v30
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v11 != v12 {
				if base.I32_extend8_s(v11) < base.I32_extend8_s(v12) {
					v19 = int32(-1)
				} else {
					v19 = int32(1)
				}
				v21 = v19
			} else {
				v21 = int32(0)
			}
			return v21
		}
	}
}
func F_CancelDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v3 = l2
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[1]))
	v17 = F_LWLockAcquire(m, v13+int32(512), v4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[2]))
	v31 = v25
	v32 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[1]))
	F_LWLockRelease(m, v69+int32(512))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v32<<(uint(int32(2))%32))))
	v41 = v31 + v38*int32(640)
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v56 = v32 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v56 < v57 {
		v31 = v53
		v32 = v56
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	if v42 != l0 {
		v53 = v31
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+73)) = uint8(v3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	if v45 == int32(0) {
		v53 = v31
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v49 = F_SendProcSignal(m, v45, l1, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[2]))
	v53 = v52
	goto L8
L15:
	;
	goto L7
L16:
	;
	return
}
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
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+316))
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
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+248))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v46)+256))
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
														F_errfinish(m, int32(_a_F_CheckArchiveTimeout_2), int32(728), int32(_a_F_CheckArchiveTimeout_3))
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
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[1])) = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[2]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9 != 0 {
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v12 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v15 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[3]))
	if v19 == v15 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[4]))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v49 = F_pgmem_kill(m, v15, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v23 + int32(16)
	goto L1
L10:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+15)) = uint8(v29)
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[5]))
	v37 = F_write(m, v33, v23+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v37 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0]))
	if v41 == int32(27) {
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
	Fn13824(m, l0, int32(79), int32(_a_F_CheckDim_2_0), int32(_a_F_CheckDim_2_1), int32(1000000000), int32(74), int32(_a_F_CheckDim_2_2), int32(1000000001))
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
	F_errdetail(m, int32(_a_F_CheckPubRelationColumnList_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_CheckPubRelationColumnList_2), int32(809), int32(_a_F_CheckPubRelationColumnList_3))
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
	F_errdetail(m, int32(_a_F_CheckPubRelationColumnList_5), v11+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CheckPubRelationColumnList_2), int32(824), int32(_a_F_CheckPubRelationColumnList_3))
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
						F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3404), int32(_a_F_CheckSelectLocking_3))
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
							F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3411), int32(_a_F_CheckSelectLocking_3))
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
								F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3411), int32(_a_F_CheckSelectLocking_3))
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
									F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3418), int32(_a_F_CheckSelectLocking_3))
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
										F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3425), int32(_a_F_CheckSelectLocking_3))
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
											F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3432), int32(_a_F_CheckSelectLocking_3))
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
												F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3439), int32(_a_F_CheckSelectLocking_3))
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
					F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3397), int32(_a_F_CheckSelectLocking_3))
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
func F_ConditionVariableSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	v8 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableSignal_0), int32(264), int32(_a_F_ConditionVariableSignal_1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v17 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
	return
L7:
	;
	goto L8
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v25 = v22 + v14*int32(640)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	if v27 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v26 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	v36 = v31
	goto L9
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v27*int32(640))+84)) = v26
	v36 = v27
	goto L9
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+84)) = int64(0)
	v49 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v49))
	v53 = v25 + int32(20)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36
	goto L13
L15:
	;
	goto L16
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v42+v26*int32(640))+88)) = v36
	goto L13
L17:
	;
	return
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v57 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v60 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[1]))
	if v64 == v60 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = m.G0
	v68 = v66 - int32(16)
	m.G0 = v68
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[2]))
	if v71 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v94 = F_pgmem_kill(m, v60, int32(23))
	mBase = m.M
	goto L18
L25:
	;
	m.G0 = v68 + int32(16)
	goto L17
L26:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+15)) = uint8(v74)
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[3]))
	v82 = F_write(m, v78, v68+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v82 {
		goto L25
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[4]))
	if v86 == int32(27) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	if l0 < int32(0) {
		return int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[0]))
		v14 = F_LWLockConditionalAcquire(m, v7+l0<<(uint(int32(6))%32)-int32(16), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_ConditionalLockRelationOid_0) {
			v36 = l0 - int32(_a_F_ConditionalLockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_ConditionalLockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_ConditionalLockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_ConditionalLockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_ConditionalLockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v89
	v93 = int32(0)
	v98 = F_LockAcquireExtended(m, v7+int32(16), l1, v93, int32(1), v7+int32(12), v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		return int32(0)
	} else {
		switch v98 {
		case 0, 3:
			m.G0 = v7 + int32(32)
			return base.B2i32(v98 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v104)+53)) = uint8(v105)
				m.G0 = v7 + int32(32)
				return base.B2i32(v98 != int32(0))
			}
		}
	}
}
func F_ConditionalLockTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v5 = int32(0)
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14 | v15<<(uint(v7)%32)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v20)
	v27 = F_LockAcquireExtended(m, v8, l2, v5, int32(1), v5, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.B2i32(v27 != int32(0))
	}
}
func F_ConversionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_SearchSysCache1(m, int32(20), l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v155
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
	v155 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_ConversionIsVisibleExt_0), v12)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ConversionIsVisibleExt_1), int32(2536), int32(_a_F_ConversionIsVisibleExt_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v41 = v37 + v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	if v42 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v47 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v86 == int32(0) {
		v144 = v45
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v86 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v54 <= int32(0) {
		v80 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = v80
	goto L18
L23:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = v54
	goto L26
L25:
	;
	v60 = v57
	goto L26
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v63 = int32(0)
	goto L27
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63<<(uint(int32(2))%32))))
	v72 = base.B2i32(v71 == v42)
	if v71 == v42 {
		v80 = v72
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v80 = v72
	goto L22
L29:
	;
	v74 = v63 + int32(1)
	if v74 != v60 {
		v63 = v74
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v93 == int32(0) {
		v136 = v3
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v144 = base.B2i32(l0 == v136)
	goto L14
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 <= int32(0) {
		v136 = v3
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v105 = int32(0)
	v107 = v102
	v111 = v96
	goto L36
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v105<<(uint(int32(2))%32))))
	if v107 != v117 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v136 = int32(0)
	goto L33
L38:
	;
	v120 = int32(0)
	v122 = F_GetSysCacheOid(m, int32(18), v41+int32(4), v117, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v127 = v107
	v128 = v111
	goto L40
L40:
	;
	v130 = v105 + int32(1)
	if v130 < v128 {
		v105 = v130
		v107 = v127
		v111 = v128
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v122 != 0 {
		v136 = v122
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v127 = v126
	v128 = v124
	goto L40
L43:
	;
	goto L37
L44:
	;
	v155 = v144
	goto L1
}
func F_CopyGetAttnums(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
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
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v248
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v15 <= int32(0) {
		v248 = v4
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v213 <= int32(0) {
		v248 = v4
		goto L1
	} else {
		goto L67
	}
L5:
	;
	v25 = v4
	v26 = v4
	goto L8
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L34
	} else {
		goto L62
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L34
	} else {
		goto L58
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v35 <= v34 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v158 + int32(4)
	F_errmsg(m, int32(_a_F_CopyGetAttnums_0), v11+int32(-48))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L34
	} else {
		goto L56
	}
L10:
	;
	goto L9
L11:
	;
	v113 = int32(0)
	if v25 == v113 {
		goto L41
	} else {
		goto L42
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v42 = v34
	v43 = v35
	goto L14
L14:
	;
	v53 = l0 + v43<<(uint(int32(4))%32) + v42*int32(100)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+111)))
	if v54 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+90)))
	if v83 != 0 {
		goto L6
	} else {
		goto L32
	}
L16:
	;
	goto L15
L17:
	;
	v58 = v53 + int32(20)
	v60 = v53 + int32(24)
	if v60|v33 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v79 = v43
	goto L19
L19:
	;
	v81 = v42 + int32(1)
	if v81 < v79 {
		v42 = v81
		v43 = v79
		goto L14
	} else {
		goto L31
	}
L20:
	;
	if v75 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L21:
	;
	v66 = int32(-1)
	goto L23
L22:
	;
	v66 = int32(0)
	goto L23
L23:
	;
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(1)
	goto L26
L25:
	;
	v67 = v66
	goto L26
L26:
	;
	v68 = int32(0)
	if base.B2i32(v60 == v68)|base.B2i32(v33 == v68) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v75 = v67
	goto L29
L28:
	;
	v74 = F_strncmp(m, v60, v33, int32(64))
	mBase = m.M
	v75 = v74
	goto L29
L29:
	;
	goto L20
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = v78
	goto L19
L31:
	;
	goto L12
L32:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+74)))
	if v84 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	return int32(0)
L35:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_1), v13)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1045), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	if v151 != 0 {
		goto L7
	} else {
		goto L53
	}
L41:
	;
	v151 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v119 <= int32(0) {
		v145 = v113
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v151 = v145
	goto L40
L45:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v125 = v119
	goto L48
L47:
	;
	v125 = v122
	goto L48
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v128 = int32(0)
	goto L49
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v126+v128<<(uint(int32(2))%32))))
	v137 = base.B2i32(v136 == v84)
	if v136 == v84 {
		v145 = v137
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v145 = v137
	goto L44
L51:
	;
	v139 = v128 + int32(1)
	if v139 != v125 {
		v128 = v139
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v152 = F_lappend_int(m, v25, v84)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v155 = v26 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v155 < v156 {
		v25 = v152
		v26 = v155
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v248 = v152
	goto L1
L56:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1040), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L34
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L34
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_4), v11+int32(-32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L34
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1052), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L34
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(_a_F_CopyGetAttnums_5))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_6), v11+int32(-16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L34
	} else {
		goto L64
	}
L64:
	;
	F_errdetail(m, int32(_a_F_CopyGetAttnums_7), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L34
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1029), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L34
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v222 = v4
	v225 = v4
	goto L68
L68:
	;
	v230 = l0 + int32(20) + v222<<(uint(int32(4))%32)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+9)))
	if v231 != 0 {
		v237 = v225
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v248 = v237
	goto L1
L70:
	;
	v239 = v222 + int32(1)
	if v239 != v213 {
		v222 = v239
		v225 = v237
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+10)))
	if v232 != 0 {
		v237 = v225
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v235 = F_lappend_int(m, v225, v222+int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v237 = v235
	goto L70
L74:
	;
	goto L69
}
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = F_strlen(m, l0)
	mBase = m.M
	if v4 <= int32(100) {
		v7 = F_pstrdup(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v13 = F_pg_mbcliplen(m, l0, v4, int32(100))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = F_palloc(m, v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					base.MemoryCopy(m, v17, l0, v13)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13+v17))) = int32(_a_F_CopyLimitPrintoutLength_0)
				return v17
			}
		}
	}
}
func F_CopyReadAttributesText(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
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
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v20 <= v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v456
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v23 == int32(0) {
		v456 = v2
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = l0 + int32(264)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
	goto L11
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_0), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(1581), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v56 <= v55 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_enlargeStringInfo(m, v47, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v61 = v55
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v63 = v61 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v67 = v64
	v68 = v62
	v72 = v2
	goto L16
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v61 = v60
	goto L14
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v81 <= v72 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v262 - v446
	v456 = v443
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v81 << (uint(int32(1)) % 32)
	v88 = F_repalloc(m, v80, v81<<(uint(int32(3))%32))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v91 = v80
	goto L20
L20:
	;
	v93 = v72 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v91+v93))) = v67
	v96 = int32(0)
	if base.Ui32(v63) <= base.Ui32(v68) {
		v260 = v68
		v261 = v68
		v262 = v67
		v266 = v96
		v270 = v96
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v88
	v91 = v88
	goto L20
L22:
	;
	v273 = v260 - v68
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v273 != v274 {
		goto L72
	} else {
		goto L73
	}
L23:
	;
	v101 = v68
	v103 = v67
	v107 = v96
	goto L24
L24:
	;
	v115 = v101 + int32(1)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v117 = base.B2i32(v116 == v45)
	if v116 == v45 {
		v260 = v101
		v261 = v115
		v262 = v103
		v266 = v107
		v270 = v117
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v260 = v250
	v261 = v250
	v262 = v256
	v266 = v252
	v270 = v117
	goto L22
L26:
	;
	if v116 != int32(92) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v249)
	v256 = v103 + int32(1)
	if base.Ui32(v250) < base.Ui32(v63) {
		v101 = v250
		v103 = v256
		v107 = v252
		goto L24
	} else {
		goto L70
	}
L28:
	;
	v249 = v116
	v250 = v115
	v252 = v107
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(v63) <= base.Ui32(v115) {
		v260 = v101
		v261 = v115
		v262 = v103
		v266 = v107
		v270 = v117
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v122 = v101 + int32(2)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v125 = v123 - int32(48)
	switch v125 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		goto L39
	default:
		v249 = v123
		v250 = v122
		v252 = v107
		goto L27
	case 50:
		goto L37
	case 54:
		goto L36
	case 62:
		goto L35
	case 66:
		goto L34
	case 68:
		goto L33
	case 70:
		goto L32
	case 72:
		goto L38
	}
L32:
	;
	v249 = int32(11)
	v250 = v122
	v252 = v107
	goto L27
L33:
	;
	v249 = int32(9)
	v250 = v122
	v252 = v107
	goto L27
L34:
	;
	v249 = int32(13)
	v250 = v122
	v252 = v107
	goto L27
L35:
	;
	v249 = int32(10)
	v250 = v122
	v252 = v107
	goto L27
L36:
	;
	v249 = int32(12)
	v250 = v122
	v252 = v107
	goto L27
L37:
	;
	v249 = int32(8)
	v250 = v122
	v252 = v107
	goto L27
L38:
	;
	v159 = int32(120)
	if base.Ui32(v63) <= base.Ui32(v122) {
		v249 = v159
		v250 = v122
		v252 = v107
		goto L27
	} else {
		goto L49
	}
L39:
	;
	if base.Ui32(v63) <= base.Ui32(v122) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v249 = v152
	v250 = v153
	v252 = base.B2i32(base.I32_extend8_s(v152) <= int32(0)) | v107
	goto L27
L41:
	;
	v152 = v125
	v153 = v122
	goto L40
L42:
	;
	goto L43
L43:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v127&int32(248) != int32(48) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v152 = v125
	v153 = v122
	goto L40
L45:
	;
	goto L46
L46:
	;
	v132 = int32(3)
	v136 = v127 + v125<<(uint(v132)%32) - int32(48)
	v138 = v101 + v132
	if base.Ui32(v63) <= base.Ui32(v138) {
		v152 = v136
		v153 = v138
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v140&int32(248) != int32(48) {
		v152 = v136
		v153 = v138
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v152 = v140 + v136<<(uint(int32(3))%32) - int32(48)
	v153 = v101 + int32(4)
	goto L40
L49:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	goto L50
L50:
	;
	if base.B2i32(base.Ui32(v161-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v161|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v249 = v159
		v250 = v122
		v252 = v107
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v176 = v161 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v176&int32(255)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if base.Ui32(v161-int32(65)) < base.Ui32(int32(26)) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v190 = v176
	goto L54
L54:
	;
	v192 = v101 + int32(3)
	if base.Ui32(v63) <= base.Ui32(v192) {
		v229 = v190
		v230 = v192
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v190 = v187 - int32(87)
	goto L54
L56:
	;
	v187 = v161 | int32(32)
	goto L58
L57:
	;
	v187 = v161
	goto L58
L58:
	;
	goto L55
L59:
	;
	v249 = v229
	v250 = v230
	v252 = base.B2i32(v229&int32(255) == int32(0)) | int32(base.Ui32(v229&int32(128))>>(uint(int32(7))%32)) | v107
	goto L27
L60:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	goto L61
L61:
	;
	if base.B2i32(base.Ui32(v194-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v194|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v229 = v190
		v230 = v192
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v211 = v194 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v211&int32(255)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if base.Ui32(v194-int32(65)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v225 = v211
	goto L65
L65:
	;
	v229 = v225 + v190<<(uint(int32(4))%32)
	v230 = v101 + int32(4)
	goto L59
L66:
	;
	v225 = v222 - int32(87)
	goto L65
L67:
	;
	v222 = v194 | int32(32)
	goto L69
L68:
	;
	v222 = v194
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L25
L71:
	;
	v440 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v440)
	v442 = int32(1)
	v443 = v72 + v442
	if v270 != 0 {
		v67 = v262 + v442
		v68 = v261
		v72 = v443
		goto L16
	} else {
		goto L119
	}
L72:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v326 != 0 {
		goto L89
	} else {
		goto L90
	}
L73:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v273 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v321 != 0 {
		goto L72
	} else {
		goto L87
	}
L75:
	;
	v321 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v282 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v283 = v68
	v284 = v276
	v285 = v273
	v286 = v282
	goto L82
L79:
	;
	v309 = v276
	v313 = int32(0)
	goto L80
L80:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v321 = v313 - v314
	goto L74
L81:
	;
	v309 = v304
	v313 = v306
	goto L80
L82:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if base.B2i32(v286 != v288)|base.B2i32(v288 == int32(0)) != 0 {
		v304 = v284
		v306 = v286
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v304 = v298
	v306 = int32(0)
	goto L81
L84:
	;
	v294 = v285 - int32(1)
	if v294 == int32(0) {
		v304 = v284
		v306 = v286
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v297 = int32(1)
	v298 = v284 + v297
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v299 != 0 {
		v283 = v283 + v297
		v284 = v298
		v285 = v294
		v286 = v299
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v322+v93))) = int32(0)
	goto L71
L88:
	;
	if v266&int32(1) == int32(0) {
		goto L71
	} else {
		goto L117
	}
L89:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v329 = v327
	goto L91
L90:
	;
	v329 = int32(0)
	goto L91
L91:
	;
	if v329 <= v72 {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v331 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v273 != v334 {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	if v273 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v380 != 0 {
		goto L88
	} else {
		goto L108
	}
L96:
	;
	v380 = int32(0)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v341 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v342 = v68
	v343 = v331
	v344 = v273
	v345 = v341
	goto L103
L100:
	;
	v368 = v331
	v372 = int32(0)
	goto L101
L101:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	v380 = v372 - v373
	goto L95
L102:
	;
	v368 = v363
	v372 = v365
	goto L101
L103:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if base.B2i32(v345 != v347)|base.B2i32(v347 == int32(0)) != 0 {
		v363 = v343
		v365 = v345
		goto L102
	} else {
		goto L105
	}
L104:
	;
	v363 = v357
	v365 = int32(0)
	goto L102
L105:
	;
	v353 = v344 - int32(1)
	if v353 == int32(0) {
		v363 = v343
		v365 = v345
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v356 = int32(1)
	v357 = v343 + v356
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	if v358 != 0 {
		v342 = v342 + v356
		v343 = v357
		v344 = v353
		v345 = v358
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v382+v93)))
	v386 = v384 - int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v381+v386<<(uint(int32(2))%32))))
	if v390 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v391+v386))) = uint8(v393)
	goto L71
L110:
	;
	goto L111
L111:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+52))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_3), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v396 + v397<<(uint(int32(4))%32) + v386*int32(100) + int32(24)
	F_errdetail(m, int32(_a_F_CopyReadAttributesText_4), v18)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(1775), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v431+v93)))
	F_pg_verifymbstr(m, v433, v262-v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	goto L71
L119:
	;
	goto L17
}
func F_CreateDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v25 int64
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	switch l0 - int32(1) {
	case 0:
		return int32(_a_F_CreateDestReceiver_0)
	case 1, 2:
		v7 = F_palloc0(m, int32(60))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(base.B2i32(l0 == int32(2)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(25)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(26)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(27)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(28)
			v25 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+28)) = v25
			*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v25
			return v7
		}
	case 3:
		v107 = int32(_a_F_CreateDestReceiver_1)
		return v107
	case 4:
		return int32(_a_F_CreateDestReceiver_2)
	case 5:
		v37 = F_palloc0(m, int32(56))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(784)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = int32(785)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(786)
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(787)
			return v37
		}
	case 6:
		v51 = F_CreateIntoRelDestReceiver(m, int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			return v51
		}
	case 7:
		v55 = F_palloc(m, int32(32))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = int64(8)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(527)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(528)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(529)
			*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(530)
			return v55
		}
	case 8:
		v71 = F_palloc0(m, int32(28))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = int32(9)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(688)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(689)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(690)
			*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(691)
			return v71
		}
	case 9:
		v85 = F_palloc0(m, int32(40))
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = int32(10)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = int32(562)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(563)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = int32(564)
			*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(565)
			return v85
		}
	case 10:
		v101 = F_CreateTupleQueueDestReceiver(m, int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			return v101
		}
	case 11:
		v105 = F_CreateExplainSerializeDestReceiver(m, int32(0))
		mBase = m.M
		v106 = m.ExcPending
		if v106 != 0 {
			return int32(0)
		} else {
			v107 = v105
			return v107
		}
	default:
		return int32(_a_F_CreateDestReceiver_3)
	}
}
func F_CreateRestartPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	var v266 float64
	_ = v266
	var v268 int32
	_ = v268
	var v270 float64
	_ = v270
	var v277 float64
	_ = v277
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v294 int64
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int64
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	v11 = m.G0
	v13 = v11 - int32(1200)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v19 = base.AtomicRmwXchg32(m, v16, int32(440), int32(1))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	F_s_lock(m, v21+int32(440), int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_1), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+352))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v32)+344))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v32)+336))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v32)+328))
	base.MemoryCopy(m, v13+int32(84), v32+int32(356), int32(76))
	v43 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v32)+440)), uint32(v43))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v47 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v13 + int32(1200)
	return v442
L7:
	;
	if v36 != int64(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	v53 = base.B2i32(v51 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(v53)
	if v51 != int32(2) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v56 = int32(0)
	v59 = F_errstart(m, int32(13), v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v59 == int32(0) {
		v442 = v56
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_16), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_17), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v442 = v56
	goto L6
L16:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L32
	}
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+40))
	if base.Ui64(v76) < base.Ui64(v34) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v78 = int32(0)
	v81 = F_errstart(m, int32(13), v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v81 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v34)
	v85 = int64(base.Ui64(v34) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v85)
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_18), v13)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_19), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if l0&int32(1) == int32(0) {
		v442 = v78
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v108 = F_LWLockAcquire(m, v104+int32(1152), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+16)) = int32(2)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v115, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v119+int32(1152))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v442 = v78
	goto L6
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v127)+152)) = v34
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3])) = v34
	F_WALInsertLockRelease(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v137 = base.AtomicRmwXchg32(m, v134, int32(440), int32(1))
	if v137 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	F_s_lock(m, v139+int32(440), int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_3), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+200)) = v34
	v150 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v148)+440)), uint32(v150))
	v153 = int32(_a_F_CreateRestartPoint_4)
	base.MemoryFill(m, v153, v150, int32(80))
	v161 = m.G0
	v162 = int32(16)
	v163 = v161 - v162
	m.G0 = v163
	F_gettimeofday(m, v163)
	mBase = m.M
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
	v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+8)))
	m.G0 = v163 + v162
	goto L38
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[4])) = v167 + v166*int64(1000000) - int64(946684800000000)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v178 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_LogCheckpointStart(m, l0, int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l0&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = int32(_a_F_CreateRestartPoint_5)
	if l0&int32(2) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	F_CheckPointGuts(m, v34, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L53
	}
L46:
	;
	v192 = int32(_a_F_CreateRestartPoint_6)
	goto L48
L47:
	;
	v192 = int32(_a_F_CreateRestartPoint_7)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v192
	if l0&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v198 = int32(_a_F_CreateRestartPoint_8)
	goto L51
L50:
	;
	v198 = int32(_a_F_CreateRestartPoint_7)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v198
	v201 = v13 + int32(160)
	v206 = F_pg_snprintf(m, v201, int32(128), int32(_a_F_CreateRestartPoint_9), v13+int32(48))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v208 = F_strlen(m, v201)
	mBase = m.M
	goto L45
L53:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v213)+40))
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v220 = F_LWLockAcquire(m, v216+int32(1152), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v223)+40))
	if base.Ui64(v224) < base.Ui64(v34) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+48)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v223)+40)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v223)+32)) = v36
	base.MemoryCopy(m, v223+int32(52), v13+int32(84), int32(76))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v235 != int32(5) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v255+int32(1152))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L65
	}
L58:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v251, v223)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L64
	}
L59:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v223)+136))
	if base.Ui64(v238) < base.Ui64(v35) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+144)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v223)+136)) = v35
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[7])) = v35
	goto L62
L61:
	;
	goto L62
L62:
	;
	if l0&int32(1) == int32(0) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = int32(2)
	goto L58
L64:
	;
	goto L57
L65:
	;
	v261 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	if v214 != int64(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v266 = base.F64_convert_i64_u(v261 - v214)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[9])) = v266
	v268 = int32(_a_F_CreateRestartPoint_10)
	v270 = *(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10]))
	if base.F64_gt(v266, v270) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v282 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v283 = base.I64_div_u_s(v261, v282)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v283
	v285 = int32(0)
	v287 = F_GetWalRcvFlushRecPtr(m, v285, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L72
	}
L69:
	;
	v277 = v266
	goto L71
L70:
	;
	v277 = base.F64_add(base.F64_mul(v270, float64(0.9)), base.F64_mul(v266, float64(0.1)))
	goto L71
L71:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10])) = v277
	goto L68
L72:
	;
	v291 = F_GetXLogReplayRecPtr(m, v13+int32(80))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	if base.Ui64(v291) < base.Ui64(v287) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v294 = v287
	goto L76
L75:
	;
	v294 = v291
	goto L76
L76:
	;
	v296 = v13 + int32(72)
	F_KeepLogSeg(m, v294, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	v301 = int32(0)
	v303 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v300, v301, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	if v303 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v306 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	v308 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v309 = base.I64_div_u_s(v306, v308)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v309
	F_KeepLogSeg(m, v294, v296)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	v314 = v300
	goto L81
L81:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v320 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	v314 = v313
	goto L81
L83:
	;
	v336 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	F_RemoveOldXlogFiles(m, v314-int64(1), v336, v294, v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L87
	}
L84:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v318)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v332
	v334 = v332
	goto L83
L85:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)+316))
	v325 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(base.B2i32(v324 != v325))
	if v324 == v325 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v334 = v330
	goto L83
L87:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+320)))
	if v341 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[12])))
	if v384 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v345 = v294 - int64(1)
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11]))
	if base.Ui64(v345&base.I64_extend_i32_s(v347-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v347), float64(0.75))))) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v359 = base.I64_div_u_s(v345, base.I64_extend_i32_s(v347))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v367 = F_XLogFileInitInternal(m, v359+int64(1), v362, v13+int32(1199), v13+int32(160))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if int32(0) <= v367 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v371 = F_close(m, v367)
	mBase = m.M
	goto L94
L93:
	;
	goto L94
L94:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1199)))
	if v372 != int32(1) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v375 = int32(_a_F_CreateRestartPoint_15)
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14])) = v377 + int32(1)
	goto L88
L96:
	;
	v387 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_LogCheckpointEnd(m, int32(1))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	F_TruncateSUBTRANS(m, v387)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v394 = F_GetLatestXTime(m)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v399 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v400 = int32(15)
	goto L105
L104:
	;
	v400 = int32(13)
	goto L105
L105:
	;
	v402 = F_errstart(m, v400, int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	if v402 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v34)
	v406 = int64(base.Ui64(v34) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+32)) = uint32(v406)
	F_errmsg(m, int32(_a_F_CreateRestartPoint_11), v13+int32(32))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v428 = int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[13]))
	if v430 == int32(0) {
		v442 = v428
		goto L6
	} else {
		goto L117
	}
L110:
	;
	if v394 != int64(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v415 = F_timestamptz_to_str(m, v394)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_13), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v415
	F_errdetail(m, int32(_a_F_CreateRestartPoint_12), v13+int32(16))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	goto L109
L117:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v433 == int32(0) {
		v442 = v428
		goto L6
	} else {
		goto L118
	}
L118:
	;
	F_ExecuteRecoveryCommand(m, v430, int32(_a_F_CreateRestartPoint_14), int32(0), int32(134217729))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v442 = v428
	goto L6
}
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1322 int32
	_ = v1322
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1381 int32
	_ = v1381
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1484 int32
	_ = v1484
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1545 int32
	_ = v1545
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1622 int32
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int64
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1720 int32
	_ = v1720
	var v1743 int32
	_ = v1743
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1766 int32
	_ = v1766
	var v1793 int32
	_ = v1793
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1835 int32
	_ = v1835
	var v1837 int64
	_ = v1837
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(432)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v26 == v4 {
		goto L18
	} else {
		goto L19
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L24
	} else {
		goto L433
	}
L2:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1835
	v1837 = *(*int64)(unsafe.Add(mBase, uint32(v1820)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1837
	m.G0 = v22 + int32(432)
	return
L3:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1086 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L4:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1036 == int32(0) {
		v1072 = v1022
		v1074 = v1024
		v1080 = v514
		goto L3
	} else {
		goto L257
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L24
	} else {
		goto L252
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L24
	} else {
		goto L248
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L24
	} else {
		goto L242
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L24
	} else {
		goto L238
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L24
	} else {
		goto L234
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L24
	} else {
		goto L229
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L24
	} else {
		goto L225
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L24
	} else {
		goto L221
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L24
	} else {
		goto L217
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L24
	} else {
		goto L213
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L24
	} else {
		goto L209
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L24
	} else {
		goto L204
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L24
	} else {
		goto L200
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L24
	} else {
		goto L196
	}
L19:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 != int32(1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v40 = v4
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v40<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != int32(3) {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v126 != 0 {
		goto L49
	} else {
		goto L50
	}
L23:
	;
	v64 = F_relation_openrv(m, v59, int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	v69 = v67 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v69))|base.B2i32(int32(1)<<(uint(v69)%32)&int32(_a_F_CreateStatistics_0) == int32(0)) != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
	v81 = F_object_ownercheck(m, int32(1259), v80, v25)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if v81 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v86)+119)))
	switch v87 - int32(73) {
	case 0, 32:
		v97 = int32(20)
		goto L32
	default:
		goto L33
	case 10:
		goto L37
	case 29:
		goto L34
	case 36:
		goto L35
	case 45:
		goto L36
	}
L29:
	;
	goto L30
L30:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[1])))
	if v106 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	F_aclcheck_error(m, int32(2), v99, v100+int32(4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L24
	} else {
		goto L38
	}
L32:
	;
	v99 = v97
	goto L31
L33:
	;
	v97 = int32(41)
	goto L32
L34:
	;
	v99 = int32(18)
	goto L31
L35:
	;
	v99 = int32(23)
	goto L31
L36:
	;
	v99 = int32(51)
	goto L31
L37:
	;
	v99 = int32(37)
	goto L31
L38:
	;
	goto L30
L39:
	;
	v110 = int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
	if base.Ui32(v111) < base.Ui32(int32(_a_F_CreateStatistics_1)) {
		v120 = v110
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v122 = v40 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v122 < v123 {
		v40 = v122
		goto L21
	} else {
		goto L47
	}
L42:
	;
	if v120 != 0 {
		goto L15
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+68))
	if v115 == int32(99) {
		v120 = v110
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v118 = F_isTempToastNamespace(m, v115)
	mBase = m.M
	v120 = v118
	goto L43
L46:
	;
	goto L41
L47:
	;
	goto L22
L48:
	;
	v439 = F_strncpy(m, v22+int32(304), v421, int32(64))
	mBase = m.M
	v440 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v439)+63)) = uint8(v440)
	goto L113
L49:
	;
	v129 = F_QualifiedNameGetCreationNamespace(m, v126, v22+int32(284))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L24
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+68))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)) = uint8(v135)
	v139 = v133 + int32(4)
	if v132 == v135 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	v421 = v131
	v429 = v129
	goto L48
L53:
	;
	v339 = v22 + int32(304)
	v340 = F_pstrdup(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L24
	} else {
		goto L100
	}
L54:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v142 <= int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v150 = v135
	v151 = int32(0)
	v152 = v142
	goto L56
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+v150<<(uint(int32(2))%32))))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v170 == int32(206) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L53
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if int32(0) < v151 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v313 = v151
	v314 = v152
	goto L60
L60:
	;
	v317 = v150 + int32(1)
	if v317 < v314 {
		v150 = v317
		v151 = v313
		v152 = v314
		goto L56
	} else {
		goto L99
	}
L61:
	;
	v179 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(304)+v151))) = uint8(v179)
	v183 = v151 + int32(1)
	goto L63
L62:
	;
	v183 = v151
	goto L63
L63:
	;
	v186 = v22 + int32(304) + v183
	if v173 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v188 = v173
	goto L66
L65:
	;
	v188 = int32(_a_F_CreateStatistics_2)
	goto L66
L66:
	;
	goto L70
L67:
	;
	v308 = F_strlen(m, v186)
	mBase = m.M
	v309 = v308 + v183
	if int32(63) < v309 {
		goto L53
	} else {
		goto L98
	}
L68:
	;
	v305 = F_strlen(m, v294)
	mBase = m.M
	goto L67
L70:
	;
	goto L71
L71:
	;
	v195 = int32(63)
	if (v186^v188)&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v298)
	goto L68
L73:
	;
	v279 = v274
	v280 = v275
	v281 = v276
	goto L94
L74:
	;
	if v269 == int32(0) {
		v294 = v267
		v295 = v268
		goto L72
	} else {
		goto L93
	}
L75:
	;
	v267 = v188
	v268 = v186
	v269 = v195
	goto L74
L76:
	;
	goto L77
L77:
	;
	v199 = int32(0)
	if base.B2i32(v188&int32(3) == v199)|int32(0) == v199 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v235 == int32(0) {
		v294 = v232
		v295 = v233
		goto L72
	} else {
		goto L87
	}
L79:
	;
	v211 = v188
	v212 = v186
	v213 = v195
	goto L82
L80:
	;
	goto L81
L81:
	;
	v232 = v188
	v233 = v186
	v234 = v195
	v235 = int32(1)
	goto L78
L82:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v215)
	if v215 == int32(0) {
		v274 = v211
		v275 = v212
		v276 = v213
		goto L73
	} else {
		goto L84
	}
L83:
	;
	v232 = v226
	v233 = v220
	v234 = v222
	v235 = v224
	goto L78
L84:
	;
	v219 = int32(1)
	v220 = v212 + v219
	v222 = v213 - v219
	v223 = int32(0)
	v224 = base.B2i32(v222 != v223)
	v226 = v211 + v219
	if v226&int32(3) == v223 {
		v232 = v226
		v233 = v220
		v234 = v222
		v235 = v224
		goto L78
	} else {
		goto L85
	}
L85:
	;
	if v222 != 0 {
		v211 = v226
		v212 = v220
		v213 = v222
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if base.B2i32(v238 == int32(0))|base.B2i32(base.Ui32(v234) < base.Ui32(int32(4))) != 0 {
		v267 = v232
		v268 = v233
		v269 = v234
		goto L74
	} else {
		goto L88
	}
L88:
	;
	v245 = v232
	v246 = v233
	v247 = v234
	goto L89
L89:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v253 = int32(-2139062144)
	if (int32(16843008)-v250|v250)&v253 != v253 {
		v274 = v245
		v275 = v246
		v276 = v247
		goto L73
	} else {
		goto L91
	}
L90:
	;
	v267 = v261
	v268 = v259
	v269 = v263
	goto L74
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v250
	v258 = int32(4)
	v259 = v246 + v258
	v261 = v245 + v258
	v263 = v247 - v258
	if base.Ui32(int32(3)) < base.Ui32(v263) {
		v245 = v261
		v246 = v259
		v247 = v263
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v274 = v267
	v275 = v268
	v276 = v269
	goto L73
L94:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v283)
	if v283 == int32(0) {
		v294 = v279
		v295 = v280
		goto L72
	} else {
		goto L96
	}
L95:
	;
	v294 = v290
	v295 = v288
	goto L72
L96:
	;
	v287 = int32(1)
	v288 = v280 + v287
	v290 = v279 + v287
	v292 = v281 - v287
	if v292 != 0 {
		v279 = v290
		v280 = v288
		v281 = v292
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v313 = v309
	v314 = v312
	goto L60
L99:
	;
	goto L57
L100:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+308)) = uint8(v344)
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+304)) = v347
	v350 = F_makeObjectName(m, v139, v340, v339)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L24
	} else {
		goto L101
	}
L101:
	;
	v352 = int32(0)
	v354 = F_GetSysCacheOid(m, int32(63), v350, v134, v352, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	if v354 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v360 = v350
	v361 = int32(0)
	goto L106
L104:
	;
	v401 = v350
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+284)) = v401
	v421 = v401
	v429 = v134
	goto L48
L106:
	;
	F_pfree(m, v360)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L24
	} else {
		goto L108
	}
L107:
	;
	v401 = v391
	goto L105
L108:
	;
	v378 = v361 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = int32(_a_F_CreateStatistics_3)
	v383 = v22 + int32(304)
	v388 = F_pg_snprintf(m, v383, int32(64), int32(_a_F_CreateStatistics_4), v22+int32(144))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	v391 = F_makeObjectName(m, v139, v340, v383)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	v393 = int32(0)
	v395 = F_GetSysCacheOid(m, int32(63), v391, v134, v393, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	if v395 != 0 {
		v360 = v391
		v361 = v378
		goto L106
	} else {
		goto L112
	}
L112:
	;
	goto L107
L113:
	;
	if l2 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v458 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	v463 = F_SearchSysCacheExists(m, int32(63), v460, v429, v458, v458)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L24
	} else {
		goto L120
	}
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v448 = F_object_aclcheck(m, int32(2615), v429, v446, int64(512))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L116
	}
L116:
	;
	if v448 == int32(0) {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v453 = F_get_namespace_name(m, v429)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L24
	} else {
		goto L118
	}
L118:
	;
	F_aclcheck_error(m, v448, int32(36), v453)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L24
	} else {
		goto L119
	}
L119:
	;
	goto L114
L120:
	;
	if v463 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+25)))
	if v465 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v510 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L124:
	;
	v470 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L24
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L24
	} else {
		goto L135
	}
L127:
	;
	if v470 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_5))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L24
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_relation_close(m, v64, int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L24
	} else {
		goto L134
	}
L131:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v475
	F_errmsg(m, int32(_a_F_CreateStatistics_6), v22+int32(16))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L24
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(209), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L24
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	v1820 = int32(_a_F_CreateStatistics_9)
	goto L2
L135:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_5))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L24
	} else {
		goto L136
	}
L136:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v498
	F_errmsg(m, int32(_a_F_CreateStatistics_10), v22+int32(32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L24
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(216), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v1072 = int32(0)
	v1074 = v458
	v1080 = v4
	goto L3
L140:
	;
	goto L141
L141:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if int32(9) <= v514 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	v517 = int32(0)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v518 <= v517 {
		v1022 = v517
		v1024 = v458
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v527 = v517
	v529 = v458
	v530 = int32(0)
	goto L144
L144:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541+v530<<(uint(int32(2))%32))))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v546 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v1022 = v738
	v1024 = v740
	goto L4
L146:
	;
	v753 = v530 + int32(1)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v753 < v754 {
		v527 = v738
		v529 = v740
		v530 = v753
		goto L144
	} else {
		goto L195
	}
L147:
	;
	v547 = F_SearchSysCacheAttName(m, v125, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L24
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v545)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	if v579 == int32(6) {
		goto L157
	} else {
		goto L158
	}
L150:
	;
	if v547 == int32(0) {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+22)))
	v553 = v551 + v552
	v554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v553)+74)))
	if v554 <= int32(0) {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+90)))
	if v557 == int32(118) {
		goto L11
	} else {
		goto L153
	}
L153:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v553)+68))
	v562 = F_lookup_type_cache(m, v560, int32(2))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L24
	} else {
		goto L154
	}
L154:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v562)+56))
	if v564 == int32(0) {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v553)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(288)+v529<<(uint(int32(1))%32)))) = uint16(v572)
	F_ReleaseCatCache(m, v547)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L24
	} else {
		goto L156
	}
L156:
	;
	v738 = v527
	v740 = v529 + int32(1)
	goto L146
L157:
	;
	v582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v578)+8)))
	if v582 <= int32(0) {
		goto L9
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = int32(0)
	F_pull_varattnos(m, v578, int32(1), v22+int32(240))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L24
	} else {
		goto L165
	}
L160:
	;
	v585 = F_get_attgenerated(m, v125, v582)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L24
	} else {
		goto L161
	}
L161:
	;
	if v585 == int32(118) {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v591 = F_lookup_type_cache(m, v589, int32(2))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L24
	} else {
		goto L163
	}
L163:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v591)+56))
	if v593 == int32(0) {
		goto L7
	} else {
		goto L164
	}
L164:
	;
	v598 = int32(1)
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v578)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(288)+v529<<(uint(v598)%32)))) = uint16(v601)
	v738 = v527
	v740 = v529 + v598
	goto L146
L165:
	;
	v617 = int32(-1)
	goto L167
L166:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v716 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L167:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v22)+240))
	if v632 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L24
	} else {
		goto L184
	}
L169:
	;
	if v688 < int32(0) {
		goto L166
	} else {
		goto L180
	}
L170:
	;
	v688 = base.I32_ctz(v674) | v675<<(uint(int32(5))%32)
	goto L169
L171:
	;
	v688 = int32(-2)
	goto L169
L172:
	;
	v639 = v617 + int32(1)
	v641 = base.I32_div_s(v639, int32(32))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v642 <= v641 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v645 = v632 + int32(8)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v645+v641<<(uint(int32(2))%32))))
	v652 = v649 & (int32(-1) << (uint(v639) % 32))
	if v652 != 0 {
		v674 = v652
		v675 = v641
		goto L170
	} else {
		goto L174
	}
L174:
	;
	v654 = v641 + int32(1)
	if v654 == v642 {
		goto L171
	} else {
		goto L175
	}
L175:
	;
	v657 = v654
	goto L176
L176:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v645+v657<<(uint(int32(2))%32))))
	if v664 != 0 {
		v674 = v664
		v675 = v657
		goto L170
	} else {
		goto L178
	}
L177:
	;
	goto L171
L178:
	;
	v666 = v657 + int32(1)
	if v666 != v642 {
		v657 = v666
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v693 = base.I32_extend16_s(v688 - int32(7))
	if v693 <= int32(0) {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v696 = F_get_attgenerated(m, v125, v693)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L24
	} else {
		goto L182
	}
L182:
	;
	if v696 != int32(118) {
		v617 = v688
		goto L167
	} else {
		goto L183
	}
L183:
	;
	goto L168
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L24
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L24
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(343), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L24
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	v731 = F_lappend(m, v527, v578)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L24
	} else {
		goto L194
	}
L189:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	if v719 < int32(2) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v722 = F_exprType(m, v578)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L24
	} else {
		goto L191
	}
L191:
	;
	v725 = F_lookup_type_cache(m, v722, int32(2))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L24
	} else {
		goto L192
	}
L192:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v725)+56))
	if v727 == int32(0) {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	goto L188
L194:
	;
	v738 = v731
	v740 = v529
	goto L146
L195:
	;
	goto L145
L196:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L24
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_12), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L24
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(106), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L24
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L24
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_12), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L24
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(115), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L24
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L24
	} else {
		goto L205
	}
L205:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v795 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_13), v22)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L24
	} else {
		goto L206
	}
L206:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v803 = int32(*(*int8)(unsafe.Add(mBase, uint32(v802)+119)))
	F_errdetail_relkind_not_supported(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L24
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(135), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L24
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L24
	} else {
		goto L210
	}
L210:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v818 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_14), v22+int32(160))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L24
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(153), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L24
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L24
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(8)
	F_errmsg(m, int32(_a_F_CreateStatistics_15), v22-int32(-64))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L24
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(228), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L24
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L24
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v546
	F_errmsg(m, int32(_a_F_CreateStatistics_16), v22+int32(112))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L24
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(261), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L24
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L24
	} else {
		goto L222
	}
L222:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L24
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(268), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L24
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L24
	} else {
		goto L226
	}
L226:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L24
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(274), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L24
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L24
	} else {
		goto L230
	}
L230:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v553)+68))
	v909 = F_format_type_be(m, v908)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L24
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v546
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v22+int32(128))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L24
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(282), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L24
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L24
	} else {
		goto L235
	}
L235:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L24
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(297), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L24
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L24
	} else {
		goto L239
	}
L239:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L24
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(303), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L24
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L24
	} else {
		goto L243
	}
L243:
	;
	v962 = int32(*(*int16)(unsafe.Add(mBase, uint32(v578)+8)))
	v964 = F_get_attname(m, v125, v962, int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L24
	} else {
		goto L244
	}
L244:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v967 = F_format_type_be(m, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L24
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v967
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v964
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v22+int32(80))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L24
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(311), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L24
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L24
	} else {
		goto L249
	}
L249:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L24
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(337), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L24
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	v1004 = F_format_type_be(m, v722)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L24
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v1004
	F_errmsg(m, int32(_a_F_CreateStatistics_19), v22+int32(96))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L24
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(361), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L24
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	if base.B2i32(v1022 == int32(0))|base.B2i32(v1041 != int32(1)) != 0 {
		v1072 = v1022
		v1074 = v1024
		v1080 = v514
		goto L3
	} else {
		goto L258
	}
L258:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1045 != int32(1) {
		v1072 = v1022
		v1074 = v1024
		v1080 = v514
		goto L3
	} else {
		goto L259
	}
L259:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1048 == int32(0) {
		v1072 = v1022
		v1074 = v1024
		v1080 = v514
		goto L3
	} else {
		goto L260
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L24
	} else {
		goto L261
	}
L261:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L24
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_20), int32(0))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L24
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(381), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L24
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	v1372 = int32(0)
	v1373 = base.B2i32(v1072 == v1372)
	if base.B2i32(v1355 == v1372)|base.B2i32(v1080 < int32(2)) == v1372 {
		goto L341
	} else {
		goto L342
	}
L266:
	;
	v1096 = int32(0)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+12))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	v1100 = int32(_a_F_CreateStatistics_21)
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099))))
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	if base.B2i32(v1103 == v1096)|base.B2i32(v1103 != v1106) != 0 {
		v1124 = v1103
		v1125 = v1106
		goto L275
	} else {
		goto L276
	}
L267:
	;
	v1355 = v1094
	v1367 = v4
	v1369 = v4
	v1370 = v4
	goto L265
L268:
	;
	v1094 = int32(1)
	goto L267
L269:
	;
	goto L270
L270:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+4))
	if int32(0) < v1091 {
		goto L266
	} else {
		goto L271
	}
L271:
	;
	v1094 = int32(1)
	goto L267
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L24
	} else {
		goto L336
	}
L273:
	;
	v1190 = base.B2i32(v1126 == int32(0))
	v1191 = int32(1)
	if v1091 == v1191 {
		v1355 = v1096
		v1367 = v1187
		v1369 = v1188
		v1370 = v1190
		goto L265
	} else {
		goto L300
	}
L274:
	;
	if v1126 == int32(0) {
		v1187 = v4
		v1188 = v4
		goto L273
	} else {
		goto L281
	}
L275:
	;
	v1126 = v1124 - v1125
	goto L274
L276:
	;
	v1109 = v1099
	v1110 = v1100
	goto L277
L277:
	;
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+1)))
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109)+1)))
	if v1114 == int32(0) {
		v1124 = v1114
		v1125 = v1113
		goto L275
	} else {
		goto L279
	}
L278:
	;
	v1124 = v1114
	v1125 = v1113
	goto L275
L279:
	;
	v1117 = int32(1)
	if v1114 == v1113 {
		v1109 = v1109 + v1117
		v1110 = v1110 + v1117
		goto L277
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v1129 = int32(_a_F_CreateStatistics_22)
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099))))
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	if base.B2i32(v1132 == int32(0))|base.B2i32(v1132 != v1135) != 0 {
		v1153 = v1132
		v1154 = v1135
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v1153-v1154 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L283:
	;
	goto L282
L284:
	;
	v1138 = v1099
	v1139 = v1129
	goto L285
L285:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139)+1)))
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138)+1)))
	if v1143 == int32(0) {
		v1153 = v1143
		v1154 = v1142
		goto L283
	} else {
		goto L287
	}
L286:
	;
	v1153 = v1143
	v1154 = v1142
	goto L283
L287:
	;
	v1146 = int32(1)
	if v1143 == v1142 {
		v1138 = v1138 + v1146
		v1139 = v1139 + v1146
		goto L285
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	v1187 = v4
	v1188 = int32(1)
	goto L273
L290:
	;
	goto L291
L291:
	;
	v1159 = int32(_a_F_CreateStatistics_23)
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099))))
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	if base.B2i32(v1162 == int32(0))|base.B2i32(v1162 != v1165) != 0 {
		v1183 = v1162
		v1184 = v1165
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v1183-v1184 != 0 {
		v1322 = v1099
		goto L272
	} else {
		goto L299
	}
L293:
	;
	goto L292
L294:
	;
	v1168 = v1099
	v1169 = v1159
	goto L295
L295:
	;
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169)+1)))
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168)+1)))
	if v1173 == int32(0) {
		v1183 = v1173
		v1184 = v1172
		goto L293
	} else {
		goto L297
	}
L296:
	;
	v1183 = v1173
	v1184 = v1172
	goto L293
L297:
	;
	v1176 = int32(1)
	if v1173 == v1172 {
		v1168 = v1168 + v1176
		v1169 = v1169 + v1176
		goto L295
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	v1187 = int32(1)
	v1188 = v4
	goto L273
L300:
	;
	v1194 = int32(0)
	if v1194 < v1091 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1197 = v1091
	goto L303
L302:
	;
	v1197 = v1194
	goto L303
L303:
	;
	v1202 = v1191
	v1212 = v1187
	v1214 = v1188
	v1215 = v1190
	goto L304
L304:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1097+v1202<<(uint(int32(2))%32))))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+4))
	v1222 = int32(_a_F_CreateStatistics_21)
	v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	if base.B2i32(v1225 == int32(0))|base.B2i32(v1225 != v1228) != 0 {
		v1246 = v1225
		v1247 = v1228
		goto L308
	} else {
		goto L309
	}
L305:
	;
	v1355 = v1096
	v1367 = v1310
	v1369 = v1311
	v1370 = v1312
	goto L265
L306:
	;
	v1314 = v1202 + int32(1)
	if v1197 != v1314 {
		v1202 = v1314
		v1212 = v1310
		v1214 = v1311
		v1215 = v1312
		goto L304
	} else {
		goto L335
	}
L307:
	;
	if v1246-v1247 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L308:
	;
	goto L307
L309:
	;
	v1231 = v1221
	v1232 = v1222
	goto L310
L310:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+1)))
	v1236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+1)))
	if v1236 == int32(0) {
		v1246 = v1236
		v1247 = v1235
		goto L308
	} else {
		goto L312
	}
L311:
	;
	v1246 = v1236
	v1247 = v1235
	goto L308
L312:
	;
	v1239 = int32(1)
	if v1236 == v1235 {
		v1231 = v1231 + v1239
		v1232 = v1232 + v1239
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	v1310 = v1212
	v1311 = v1214
	v1312 = int32(1)
	goto L306
L315:
	;
	goto L316
L316:
	;
	v1252 = int32(_a_F_CreateStatistics_22)
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	if base.B2i32(v1255 == int32(0))|base.B2i32(v1255 != v1258) != 0 {
		v1276 = v1255
		v1277 = v1258
		goto L318
	} else {
		goto L319
	}
L317:
	;
	if v1276-v1277 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L318:
	;
	goto L317
L319:
	;
	v1261 = v1221
	v1262 = v1252
	goto L320
L320:
	;
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1262)+1)))
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261)+1)))
	if v1266 == int32(0) {
		v1276 = v1266
		v1277 = v1265
		goto L318
	} else {
		goto L322
	}
L321:
	;
	v1276 = v1266
	v1277 = v1265
	goto L318
L322:
	;
	v1269 = int32(1)
	if v1266 == v1265 {
		v1261 = v1261 + v1269
		v1262 = v1262 + v1269
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v1310 = v1212
	v1311 = int32(1)
	v1312 = v1215
	goto L306
L325:
	;
	goto L326
L326:
	;
	v1282 = int32(_a_F_CreateStatistics_23)
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	if base.B2i32(v1285 == int32(0))|base.B2i32(v1285 != v1288) != 0 {
		v1306 = v1285
		v1307 = v1288
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if v1306-v1307 != 0 {
		v1322 = v1221
		goto L272
	} else {
		goto L334
	}
L328:
	;
	goto L327
L329:
	;
	v1291 = v1221
	v1292 = v1282
	goto L330
L330:
	;
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+1)))
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+1)))
	if v1296 == int32(0) {
		v1306 = v1296
		v1307 = v1295
		goto L328
	} else {
		goto L332
	}
L331:
	;
	v1306 = v1296
	v1307 = v1295
	goto L328
L332:
	;
	v1299 = int32(1)
	if v1296 == v1295 {
		v1291 = v1291 + v1299
		v1292 = v1292 + v1299
		goto L330
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	v1310 = int32(1)
	v1311 = v1214
	v1312 = v1215
	goto L306
L335:
	;
	goto L305
L336:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L24
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v1322
	F_errmsg(m, int32(_a_F_CreateStatistics_24), v22+int32(48))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L24
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(411), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L24
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_pg_qsort(m, v22+int32(288), v1074, int32(2), int32(571))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L24
	} else {
		goto L347
	}
L341:
	;
	v1381 = int32(1)
	v1392 = v1373
	v1393 = v1381
	v1394 = v1381
	v1395 = v1381
	goto L340
L342:
	;
	goto L343
L343:
	;
	if int32(1) < v1080 {
		v1392 = v1373
		v1393 = v1367
		v1394 = v1369
		v1395 = v1370
		goto L340
	} else {
		goto L344
	}
L344:
	;
	if v1072 == int32(0) {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1389 != int32(1) {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1392 = int32(0)
	v1393 = v1367
	v1394 = v1369
	v1395 = v1370
	goto L340
L347:
	;
	v1402 = int32(1)
	if v1402 < v1074 {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v1596 = F_buildint2vector(m, v22+int32(288), v1074)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L24
	} else {
		goto L380
	}
L349:
	;
	v1478 = v1456
	v1484 = v1457
	goto L364
L350:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L24
	} else {
		goto L360
	}
L351:
	;
	v1409 = v1402
	goto L354
L352:
	;
	goto L353
L353:
	;
	if v1392 != 0 {
		goto L348
	} else {
		goto L358
	}
L354:
	;
	v1428 = v22 + int32(288) + v1409<<(uint(int32(1))%32)
	v1429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1428))))
	v1432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1428-int32(2)))))
	if v1429 == v1432 {
		goto L350
	} else {
		goto L356
	}
L355:
	;
	goto L353
L356:
	;
	v1435 = v1409 + int32(1)
	if v1435 != v1074 {
		v1409 = v1435
		goto L354
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	v1456 = int32(0)
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1457 <= v1456 {
		goto L348
	} else {
		goto L359
	}
L359:
	;
	goto L349
L360:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L24
	} else {
		goto L361
	}
L361:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_25), int32(0))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L24
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(457), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L24
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	v1495 = int32(0)
	if v1495 < v1484 {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L24
	} else {
		goto L376
	}
L366:
	;
	goto L365
L367:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+12))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1498+v1478<<(uint(int32(2))%32))))
	v1508 = v1495
	v1510 = int32(0)
	goto L370
L368:
	;
	v1545 = v1484
	goto L369
L369:
	;
	v1557 = v1478 + int32(1)
	if v1557 < v1545 {
		v1478 = v1557
		v1484 = v1545
		goto L364
	} else {
		goto L375
	}
L370:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+12))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1523+v1508<<(uint(int32(2))%32))))
	v1528 = F_equal(m, v1502, v1527)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L24
	} else {
		goto L372
	}
L371:
	;
	if int32(2) <= v1530 {
		goto L366
	} else {
		goto L374
	}
L372:
	;
	v1530 = v1528 + v1510
	v1532 = v1508 + int32(1)
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1532 < v1533 {
		v1508 = v1532
		v1510 = v1530
		goto L370
	} else {
		goto L373
	}
L373:
	;
	goto L371
L374:
	;
	v1545 = v1533
	goto L369
L375:
	;
	goto L348
L376:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L24
	} else {
		goto L377
	}
L377:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_26), int32(0))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L24
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(492), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L24
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L380:
	;
	if v1395 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	if v1394 != 0 {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	v1606 = v22 + int32(176)
	v1607 = int32(0)
	goto L381
L383:
	;
	goto L384
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = int32(100)
	v1606 = v22 + int32(176) | int32(4)
	v1607 = int32(1)
	goto L381
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1606))) = int32(102)
	v1612 = v1607 + int32(1)
	goto L387
L386:
	;
	v1612 = v1607
	goto L387
L387:
	;
	if v1393 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(176)|v1612<<(uint(int32(2))%32)))) = int32(109)
	v1622 = v1612 + int32(1)
	goto L390
L389:
	;
	v1622 = v1612
	goto L390
L390:
	;
	if v1392 != 0 {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v1652 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L24
	} else {
		goto L400
	}
L392:
	;
	v1627 = F_construct_array_builtin(m, v22+int32(176), v1622, int32(18))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L24
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1630 = v22 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v1630+v1622<<(uint(int32(2))%32)))) = int32(101)
	v1639 = F_construct_array_builtin(m, v1630, v1622+int32(1), int32(18))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L24
	} else {
		goto L396
	}
L395:
	;
	v1647 = v1627
	v1649 = int32(0)
	goto L391
L396:
	;
	v1641 = F_nodeToString(m, v1072)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L24
	} else {
		goto L397
	}
L397:
	;
	v1643 = F_cstring_to_text(m, v1641)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L24
	} else {
		goto L398
	}
L398:
	;
	F_pfree(m, v1641)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L24
	} else {
		goto L399
	}
L399:
	;
	v1647 = v1639
	v1649 = v1643
	goto L391
L400:
	;
	v1654 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+264)) = v1654
	*(*int64)(unsafe.Add(mBase, uint32(v22)+256)) = v1654
	*(*int64)(unsafe.Add(mBase, uint32(v22)+224)) = v1654
	v1660 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+232)) = uint8(v1660)
	v1664 = F_GetNewOidWithIndex(m, v1652, int32(3380), int32(1))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L24
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+260)) = v1596
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v22)+252)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v22)+244)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = v1664
	v1671 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+230)) = uint8(v1671)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(v22)+268)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(v22)+248)) = v22 + int32(304)
	if v1649 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1680 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+232)) = uint8(v1680)
	goto L404
L403:
	;
	goto L404
L404:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+52))
	v1687 = F_heap_form_tuple(m, v1682, v22+int32(240), v22+int32(224))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L24
	} else {
		goto L405
	}
L405:
	;
	F_CatalogTupleInsert(m, v1652, v1687)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L24
	} else {
		goto L406
	}
L406:
	;
	F_pfree(m, v1687)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L24
	} else {
		goto L407
	}
L407:
	;
	F_relation_close(m, v1652, int32(3))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L24
	} else {
		goto L408
	}
L408:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[7]))
	if v1697 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1699 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3381), v1664, v1699, v1699)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L24
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	F_CacheInvalidateRelcache(m, v64)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L24
	} else {
		goto L413
	}
L412:
	;
	goto L411
L413:
	;
	v1705 = int32(0)
	F_relation_close(m, v64, v1705)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L24
	} else {
		goto L414
	}
L414:
	;
	v1709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = int32(3381)
	if v1709 < v1074 {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	if v1392 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L416:
	;
	v1720 = v1705
	goto L419
L417:
	;
	goto L418
L418:
	;
	if v1074 != 0 {
		goto L415
	} else {
		goto L423
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = int32(1259)
	v1743 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22+int32(288)+v1720<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = v1743
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(97))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L24
	} else {
		goto L421
	}
L421:
	;
	v1753 = v1720 + int32(1)
	if v1753 != v1074 {
		v1720 = v1753
		goto L419
	} else {
		goto L422
	}
L422:
	;
	goto L415
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = int32(1259)
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(97))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L24
	} else {
		goto L424
	}
L424:
	;
	goto L415
L425:
	;
	F_recordDependencyOnSingleRelExpr(m, v22+int32(200), v1072, v125, int32(97), int32(0))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L24
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = int32(2615)
	v1800 = v22 + int32(200)
	F_recordDependencyOn(m, v1800, v22+int32(212), int32(110))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L24
	} else {
		goto L429
	}
L428:
	;
	goto L427
L429:
	;
	F_recordDependencyOnOwner(m, int32(3381), v1664, v25)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L24
	} else {
		goto L430
	}
L430:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1809 == int32(0) {
		v1820 = v1800
		goto L2
	} else {
		goto L431
	}
L431:
	;
	F_CreateComments(m, v1664, int32(3381), int32(0), v1809)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L24
	} else {
		goto L432
	}
L432:
	;
	v1820 = v1800
	goto L2
L433:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L24
	} else {
		goto L434
	}
L434:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_27), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L24
	} else {
		goto L435
	}
L435:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(439), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L24
	} else {
		goto L436
	}
L436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_calcstrlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v70 + v71
L2:
	;
	v9 = l0
	v10 = v5
	v12 = v2
	goto L5
L3:
	;
	v60 = v5
	v62 = v2
	goto L4
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v70 = v63&int32(4095) + int32(1)
	v71 = v62
	goto L1
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v14 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v18 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v60 = v55
	v62 = v53
	goto L4
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v50 == int32(1) {
		v70 = v49
		v71 = v12
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v49 = v47 + v48
	goto L7
L9:
	;
	v21 = v13
	v22 = v17
	v24 = v14
	goto L12
L10:
	;
	v37 = v17
	v39 = v14
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v47 = v40&int32(4095) + int32(1)
	v48 = v39
	goto L8
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_calcstrlen(m, v25)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(1) {
		v47 = v26
		v48 = v24
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v37 = v32
	v39 = v30
	goto L11
L14:
	;
	v30 = v26 + v24
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 != int32(1) {
		v21 = v31
		v22 = v32
		v24 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v53 = v49 + v12
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != int32(1) {
		v9 = v54
		v10 = v55
		v12 = v53
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L6
}
func F_calculate_indexes_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	v4 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+116)))
	if v10 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v78 = v4
	goto L3
L3:
	;
	return v78
L4:
	;
	F_list_free(m, v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L18
	}
L5:
	;
	return int64(0)
L6:
	;
	if v13 == int32(0) {
		v68 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		v68 = v4
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(0)
	v25 = v4
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v24<<(uint(int32(2))%32))))
	v36 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v68 = v60
	goto L4
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v40 = F_calculate_relation_size(m, v36, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v44 = F_calculate_relation_size(m, v36, v42, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v48 = F_calculate_relation_size(m, v36, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v52 = F_calculate_relation_size(m, v36, v50, int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v60 = v52 + (v48 + (v44 + (v25 + v40)))
	v62 = v24 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v62 < v63 {
		v24 = v62
		v25 = v60
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	v78 = v68
	goto L3
}
func F_casefold(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(1)
		v14 = v9 + v13
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v19 = v17 & v13
		if v19 != 0 {
			v20 = v14
		} else {
			v20 = v9 + int32(4)
		}
		if v17 == int32(1) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v26 == int32(18) {
				v29 = int32(16)
			} else {
				v29 = int32(0)
			}
			if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v36 = int32(4)
			} else {
				v36 = v29
			}
			v47 = v36
		} else {
			v37 = int32(1)
			if v19 != 0 {
				v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = int32(0)
		v50 = m.G0
		v52 = v50 - int32(16)
		m.G0 = v52
		if v20 == v49 {
			v106 = v49
			m.G0 = v52 + int32(16)
			v153 = F_cstring_to_text(m, v106)
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v106)
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return int32(0)
				} else {
					return v153
				}
			}
		} else {
			if v48 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(_a_F_casefold_0)
						F_errmsg(m, int32(_a_F_casefold_1), v52)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_casefold_2), int32(0))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1847), int32(_a_F_casefold_4))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
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
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_casefold[0]))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
				if v60 != int32(6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_casefold_5), int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1853), int32(_a_F_casefold_4))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
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
					v63 = F_pg_newlocale_from_collation(m, v48)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
						if v65 == int32(1) {
							v68 = F_pnstrdup(m, v20, v47)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
								if v70 == int32(0) {
									v106 = v68
								} else {
									v74 = v70
									v75 = v68
									for {
										if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
											v88 = v74 | int32(32)
										} else {
											v88 = v74
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v88)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
										if v90 != 0 {
											v74 = v90
											v75 = v75 + int32(1)
											continue
										} else {
											break
										}
										break
									}
									v106 = v68
								}
								m.G0 = v52 + int32(16)
								v153 = F_cstring_to_text(m, v106)
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v106)
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int32(0)
									} else {
										return v153
									}
								}
							}
						} else {
							v94 = v47 + int32(1)
							v95 = F_palloc(m, v94)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v97 = F_pg_strfold(m, v95, v94, v20, v47, v63)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									v100 = v97 + int32(1)
									if base.Ui32(v100) <= base.Ui32(v94) {
										v106 = v95
										m.G0 = v52 + int32(16)
										v153 = F_cstring_to_text(m, v106)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v106)
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												return v153
											}
										}
									} else {
										v102 = F_repalloc(m, v95, v100)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v104 = F_pg_strfold(m, v102, v100, v20, v47, v63)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												v106 = v102
												m.G0 = v52 + int32(16)
												v153 = F_cstring_to_text(m, v106)
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v106)
													mBase = m.M
													v156 = m.ExcPending
													if v156 != 0 {
														return int32(0)
													} else {
														return v153
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
func F_charin(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
	v7 = F_strlen(m, v5)
	mBase = m.M
	if base.B2i32(v7 != int32(4))|base.B2i32(v6&int32(255) != int32(92)) != 0 {
		v44 = v6
		return base.I32_extend8_s(v44)
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		if v15&int32(248) != int32(48) {
			return int32(92)
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
			if v22&int32(248) != int32(48) {
				return int32(92)
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
				if v30&int32(248) != int32(48) {
					v44 = int32(92)
				} else {
					v44 = v15<<(uint(int32(6))%32) + v22<<(uint(int32(3))%32) + v30 + int32(80)
				}
				return base.I32_extend8_s(v44)
			}
		}
	}
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
	v60 = v53 + v54<<(uint(int32(4))%32) + v51*int32(100)
	v63 = F_namestrcmp(m, v60+int32(24), v40)
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
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+111)))
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
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1073), int32(_a_F_checkInsertTargets_2))
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
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1088), int32(_a_F_checkInsertTargets_2))
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
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1099), int32(_a_F_checkInsertTargets_2))
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
	v229 = v222 + v223<<(uint(int32(4))%32) + v212*int32(100)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+111)))
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
	v242 = F_pstrdup(m, v229+int32(24))
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
						F_errfinish(m, int32(_a_F_check_acl_1), int32(600), int32(_a_F_check_acl_2))
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
							F_errfinish(m, int32(_a_F_check_acl_1), int32(604), int32(_a_F_check_acl_2))
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
					F_errfinish(m, int32(_a_F_check_acl_1), int32(596), int32(_a_F_check_acl_2))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
			if v19 != l1 {
				v30 = int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+76)))
				if v21 != int32(98) {
					v30 = int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
					if v24 != l2 {
						v30 = int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
						if v27 == l3 {
							v30 = int32(1)
						} else {
							v30 = int32(0)
						}
					}
				}
			}
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v30
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_check_amop_signature_0), v9)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_check_amop_signature_1), int32(214), int32(_a_F_check_amop_signature_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	F_errfinish(m, int32(_a_F_check_duplicates_in_publist_2), int32(2383), int32(_a_F_check_duplicates_in_publist_3))
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
	v118 = v26
	goto L29
L29:
	;
	v120 = v25 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 < v121 {
		v24 = v121
		v25 = v120
		v26 = v118
		goto L4
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(2))%32)))) = v113
	v118 = v26 + int32(1)
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if base.Ui32(l0) < base.Ui32(int32(42)) {
		if base.B2i32(l0 != l3)&base.B2i32(int32(0) <= l3) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v47 = int32(3)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v47)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v49
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v47)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v53
				F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_0), v9)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1806), int32(_a_F_check_encoding_conversion_args_2))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(42)) <= base.Ui32(l1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
					F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_3), v7+int32(-32))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1808), int32(_a_F_check_encoding_conversion_args_2))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
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
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = int32(3)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v82)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v84
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l4<<(uint(v82)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
						F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_4), v7+int32(-48))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1812), int32(_a_F_check_encoding_conversion_args_2))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
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
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_5), int32(0))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1814), int32(_a_F_check_encoding_conversion_args_2))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
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
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
			F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_6), v7+int32(-16))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1802), int32(_a_F_check_encoding_conversion_args_2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v202 int32
	_ = v202
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v333 int32
	_ = v333
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v414 int32
	_ = v414
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v629 int32
	_ = v629
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v657 int32
	_ = v657
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	v24 = m.G0
	v26 = v24 - int32(1888)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+10)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = int32(100)
	goto L3
L2:
	;
	v34 = int32(112)
	goto L3
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+v34)))
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = int32(96)
	goto L6
L5:
	;
	v39 = int32(108)
	goto L6
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2+v39)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+124)))
	if v42 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L11
	} else {
		goto L144
	}
L8:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	if v126|base.B2i32(v30 <= int32(0)) != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v46 = v30 - int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v46))))
	if v48 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2+v46<<(uint(int32(1))%32))+12)))
	v62 = v49 + v50<<(uint(int32(4))%32) + v57*int32(100) - int32(80)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+68))
	v65 = F_lookup_type_cache(m, v63, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+13)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l4+v46<<(uint(int32(2))%32))))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v62)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+408)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v62)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+400)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+392)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v62)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+384)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v62)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+376)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v62)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+368)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v62)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+360)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+352)) = v88
	switch v69 - int32(109) {
	case 0:
		goto L14
	default:
		goto L15
	case 5:
		goto L13
	}
L13:
	;
	v112 = F_pg_detoast_datum(m, v73)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L21
	}
L14:
	;
	v109 = F_pg_detoast_datum(m, v73)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L19
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v26 + int32(352)
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_0), v26+int32(80))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(1165), int32(_a_F_check_exclusion_or_unique_constraint_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v111 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L7
L21:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112+int32(base.Ui32(v114)>>(uint(int32(2))%32))-int32(1)))))
	goto L22
L22:
	;
	if v120&int32(1) != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	m.G0 = v26 + int32(1888)
	return v657
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+280)) = int32(4)
	if int32(0) < v30 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v143 = int32(0)
	goto L27
L27:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v143))))
	if v155 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v657 = int32(1)
	goto L24
L29:
	;
	v159 = v143 + int32(1)
	if v30 != v159 {
		v143 = v159
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L25
L33:
	;
	v202 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v266 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L11
	} else {
		goto L43
	}
L36:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v202))))
	if v221 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v222 = int32(65)
	goto L40
L39:
	;
	v222 = int32(0)
	goto L40
L40:
	;
	v223 = int32(1)
	v224 = v202 + v223
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+v202<<(uint(v223)%32)))))
	v232 = v202 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v28+v232)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232+v41)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l4+v232)))
	F_ScanKeyEntryInitialize(m, v26+int32(352)+v202*int32(48), v222, base.I32_extend16_s(v224), v229, int32(0), v234, v236, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	if v224 != v30 {
		v202 = v224
		goto L36
	} else {
		goto L42
	}
L42:
	;
	goto L37
L43:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l6)+152))
	if v268 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v271 = F_MakePerTupleExprContext(m, l6)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L47
	}
L45:
	;
	v273 = v268
	goto L46
L46:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v273)+4)) = v266
	v277 = v266 + int32(28)
	goto L49
L47:
	;
	v273 = v271
	goto L46
L48:
	;
	F_index_endscan(m, v306)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L11
	} else {
		goto L142
	}
L49:
	;
	v301 = int32(0)
	v306 = F_index_beginscan(m, l0, l1, v26+int32(280), v301, v30, v301)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L11
	} else {
		goto L51
	}
L50:
	;
	if l9 != 0 {
		goto L111
	} else {
		goto L112
	}
L51:
	;
	v310 = int32(0)
	F_index_rescan(m, v306, v26+int32(352), v30, v310, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v315 = F_index_getnext_slot(m, v306, int32(1), v266)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	if v315 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v629 = int32(1)
	goto L48
L55:
	;
	goto L56
L56:
	;
	v333 = v301
	goto L57
L57:
	;
	if l3 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v26)+284))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v26)+288))
	if v499 != 0 {
		goto L87
	} else {
		goto L88
	}
L59:
	;
	goto L58
L60:
	;
	v472 = int32(1)
	v474 = F_index_getnext_slot(m, v306, v472, v266)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L84
	}
L61:
	;
	F_FormIndexDatum(m, l2, v266, l6, v26+int32(144), v26+int32(112))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L11
	} else {
		goto L75
	}
L62:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v345 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v350 = int32(16)
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+2)))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277))))
	if v348|v349<<(uint(v350)%32) == v353|v354<<(uint(v350)%32) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v364 == int32(0) {
		goto L61
	} else {
		goto L70
	}
L65:
	;
	goto L64
L66:
	;
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)))
	if v360 == v361 {
		v364 = int32(1)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v364 = int32(0)
	goto L65
L69:
	;
	goto L68
L70:
	;
	if v333 == int32(0) {
		v462 = int32(1)
		goto L60
	} else {
		goto L71
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v374 + int32(4)
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_3), v26-int32(-64))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(838), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+72)))
	if v394 != int32(1) {
		goto L59
	} else {
		goto L76
	}
L76:
	;
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v399 = int32(*(*int16)(unsafe.Add(mBase, uint32(v398)+10)))
	if v399 <= v397 {
		goto L59
	} else {
		goto L77
	}
L77:
	;
	v414 = v397
	goto L78
L78:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(112)+v414))))
	if v428 != 0 {
		v462 = v333
		goto L60
	} else {
		goto L80
	}
L79:
	;
	goto L59
L80:
	;
	v430 = v414 << (uint(int32(2)) % 32)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v41+v430)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v433+v430)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(144)+v430)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l4+v430)))
	v442 = F_OidFunctionCall2Coll(m, v432, v435, v439, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	if v442 == int32(0) {
		v462 = v333
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v447 = v414 + int32(1)
	if v399 != v447 {
		v414 = v447
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	if v474 != 0 {
		v333 = v462
		goto L57
	} else {
		goto L85
	}
L85:
	;
	v629 = v472
	goto L48
L86:
	;
	goto L50
L87:
	;
	v501 = v499
	goto L89
L88:
	;
	v501 = v500
	goto L89
L89:
	;
	if v501 == int32(0) {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	if l8 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if l8 != int32(2) {
		goto L86
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	F_index_endscan(m, v306)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L102
	}
L94:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	if v506 == int32(0) {
		goto L86
	} else {
		goto L95
	}
L95:
	;
	v509 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v501))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v509)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v522 == int32(0) {
		goto L86
	} else {
		goto L101
	}
L98:
	;
	v522 = base.B2i32(base.Ui32(v509) < base.Ui32(v501))
	goto L97
L99:
	;
	goto L100
L100:
	;
	v522 = int32(base.Ui32(v509-v501) >> (uint(int32(31)) % 32))
	goto L97
L101:
	;
	goto L93
L102:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	if v528 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v26)+284))
	F_SpeculativeInsertionWait(m, v529, v528)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L11
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v525 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L49
L107:
	;
	v534 = int32(8)
	goto L109
L108:
	;
	v534 = int32(5)
	goto L109
L109:
	;
	F_XactLockTableWait(m, v501, l0, v277, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	goto L49
L111:
	;
	if l10 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v542 = F_BuildIndexValueDescription(m, l1, l4, l5)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L11
	} else {
		goto L117
	}
L114:
	;
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l10)+4)) = uint16(v537)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v539
	goto L116
L115:
	;
	goto L116
L116:
	;
	v629 = int32(0)
	goto L48
L117:
	;
	v548 = F_BuildIndexValueDescription(m, l1, v26+int32(144), v26+int32(112))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_errcode(m, int32(16908482))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	v557 = int32(0)
	v561 = base.B2i32(v542 != v557) & base.B2i32(v548 != v557)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v564 = v562 + int32(4)
	if l7 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v564
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_5), v26+int32(16))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L11
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v564
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_6), v26+int32(48))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L11
	} else {
		goto L133
	}
L124:
	;
	if v561 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v580+int32(4))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L11
	} else {
		goto L131
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v542
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_7), v26)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L11
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_8), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L11
	} else {
		goto L130
	}
L129:
	;
	goto L125
L130:
	;
	goto L125
L131:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(918), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	if v561 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v607+int32(4))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L11
	} else {
		goto L140
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v542
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_9), v26+int32(32))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L11
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_10), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L11
	} else {
		goto L139
	}
L138:
	;
	goto L134
L139:
	;
	goto L134
L140:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(929), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+4)) = v274
	F_ExecDropSingleTupleTableSlot(m, v266)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L11
	} else {
		goto L143
	}
L143:
	;
	v657 = v629
	goto L24
L144:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v679 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(352)
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_11), v26+int32(96))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(1173), int32(_a_F_check_exclusion_or_unique_constraint_2))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
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
											F_errdetail(m, int32(_a_F_check_nested_generated_walker_1), int32(0))
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
													F_errfinish(m, int32(_a_F_check_nested_generated_walker_2), int32(3207), int32(_a_F_check_nested_generated_walker_3))
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
									F_errdetail(m, int32(_a_F_check_nested_generated_walker_5), int32(0))
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
											F_errfinish(m, int32(_a_F_check_nested_generated_walker_2), int32(3214), int32(_a_F_check_nested_generated_walker_3))
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
			v87 = F_expression_tree_walker_impl(m, l0, int32(464), l1)
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
	var v33 int32
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
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2681), int32(_a_F_check_srf_call_placement_3))
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
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
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v28 = F_exprLocation(m, v27)
						mBase = m.M
						F_parser_errposition(m, l0, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2555), int32(_a_F_check_srf_call_placement_3))
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
	case 7, 8, 12, 13, 17, 18, 19, 25:
		v99 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v99)
		m.G0 = v8 + int32(32)
		return
	case 9, 10, 11:
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2674), int32(_a_F_check_srf_call_placement_3))
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
func F_cidin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_uint32in_subr(m, v2, int32(0), int32(_a_F_cidin_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_clause_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_clause_selectivity_ext(m, l0, l1, l2, l3, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return float64(0)
	} else {
		return v7
	}
}
func F_clean_ipv6_addr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l0 != int32(10) {
	} else {
		v5 = int32(37)
		v6 = F___strchrnul(m, l1, v5)
		mBase = m.M
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v8 == v5 {
			v12 = v6
		} else {
			v12 = int32(0)
		}
		if v12 == int32(0) {
		} else {
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v15)
		}
	}
	return
}
func F_close_pb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_box_closept_point(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
func F_close_pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_line_closept_point(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
func F_colNameToVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v5 {
		v120 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v120
L2:
	;
	v23 = l0
	v29 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v120 = int32(0)
	goto L1
L5:
	;
	if l2|v103 != 0 {
		v120 = v103
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v103 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 <= v35 {
		v103 = v35
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v45 = v35
	v48 = v35
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	if v57 == int32(0) {
		v74 = v45
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L25
	}
L12:
	;
	goto L11
L13:
	;
	v77 = v48 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v77 < v78 {
		v45 = v74
		v48 = v77
		goto L10
	} else {
		goto L24
	}
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	if v60 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+32)))
	if v63 != int32(1) {
		v74 = v45
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v66 = F_scanNSItemForColumn(m, l0, v56, v29, l1, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	if v66 == int32(0) {
		v74 = v45
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v45 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_check_lateral_ref_ok(m, v23, v56, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v74 = v66
	goto L13
L24:
	;
	v103 = v74
	goto L5
L25:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(_a_F_colNameToVar_0), v15)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_parser_errposition(m, v23, l3)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_colNameToVar_1), int32(932), int32(_a_F_colNameToVar_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v113 != 0 {
		v23 = v113
		v29 = v29 + int32(1)
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L4
}
func F_combo_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module, int32) int32)(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v16 = m.T0[v15].(func(*base.Module, int32) int32)(m, v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v30 = int32(0)
				v31 = F_palloc0(m, v11)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l2) < base.Ui32(v11) {
						v34 = l2
					} else {
						v34 = v11
					}
					if v34 != 0 {
						base.MemoryCopy(m, v31, l1, v34)
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v31, v34, v30)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							F_pfree(m, v30)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v31)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						} else {
							F_pfree(m, v31)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								return v37
							}
						}
					}
				}
			} else {
				v20 = F_palloc0(m, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l4) <= base.Ui32(v16) {
						if l4 == int32(0) {
							v30 = v20
						} else {
							v25 = l4
							if v25 == int32(0) {
								v30 = v20
							} else {
								base.MemoryCopy(m, v20, l3, v25)
								v30 = v20
							}
						}
					} else {
						v25 = v16
						if v25 == int32(0) {
							v30 = v20
						} else {
							base.MemoryCopy(m, v20, l3, v25)
							v30 = v20
						}
					}
					v31 = F_palloc0(m, v11)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if base.Ui32(l2) < base.Ui32(v11) {
							v34 = l2
						} else {
							v34 = v11
						}
						if v34 != 0 {
							base.MemoryCopy(m, v31, l1, v34)
						} else {
						}
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v31, v34, v30)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v30 != 0 {
								F_pfree(m, v30)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v31)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										return v37
									}
								}
							} else {
								F_pfree(m, v31)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_compact(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v3
	v12 = v8
	v13 = v3
	goto L4
L2:
	;
	v28 = v3
	v33 = int32(0)
	goto L3
L3:
	;
	v35 = F_palloc_extended(m, v28, int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v16 = int32(1)
	v17 = v11 + v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v21 = v13 + v18 + v16
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v22 != 0 {
		v11 = v17
		v12 = v22
		v13 = v21
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v28 = v17
	v33 = v21 << (uint(int32(3)) % 32)
	goto L3
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v35
	v38 = int32(2)
	v41 = F_palloc_extended(m, v28<<(uint(v38)%32), v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v41
	v45 = F_palloc_extended(m, v33, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v73
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v75)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v77)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)) = uint16(v79)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v85 != 0 {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v51 = v49
	goto L17
L16:
	;
	v51 = int32(0)
	goto L17
L17:
	;
	if v51 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	F_pfree(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	F_pfree(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v57 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = int32(101)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v66 = v64
	goto L30
L29:
	;
	v66 = int32(12)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v66
	return
L31:
	;
	v90 = int32(0)
	goto L33
L32:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+12)))
	v90 = v87 + int32(1)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = v45
	v105 = v98
	goto L37
L35:
	;
	goto L36
L36:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	if v188 != 0 {
		goto L60
	} else {
		goto L61
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v107))) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(int32(2))%32)))) = v103
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v117 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v120 = v103
	v121 = v117
	goto L42
L40:
	;
	v159 = v103
	goto L41
L41:
	;
	v166 = (v159 - v103) >> (uint(int32(3)) % 32)
	if base.Ui32(int32(2)) <= base.Ui32(v166) {
		goto L55
	} else {
		goto L56
	}
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != int32(76) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v159 = v155
	goto L41
L44:
	;
	v155 = v120 + int32(8)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	if v156 != 0 {
		v120 = v155
		v121 = v156
		goto L42
	} else {
		goto L54
	}
L45:
	;
	if v125 == int32(112) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+4)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v145 = v143 + v144
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v145)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v150 | int32(1)
	goto L44
L48:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v133
	goto L44
L49:
	;
	goto L50
L50:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = int32(101)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v139 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = v139
	goto L53
L52:
	;
	v141 = int32(15)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v141
	return
L54:
	;
	goto L43
L55:
	;
	F_pg_qsort(m, v103, v166, int32(8), int32(971))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = int32(0)
	v175 = int32(_a_F_compact_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v159))) = uint16(v175)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v105)+28))
	if v179 != 0 {
		v103 = v159 + int32(8)
		v105 = v179
		goto L37
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L38
L60:
	;
	v192 = v188
	goto L63
L61:
	;
	v206 = v187
	goto L62
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v211+v212))) = uint8(v214)
	return
L63:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v196+v198))) = uint8(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	if v202 != 0 {
		v192 = v202
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v206 = v203
	goto L62
L65:
	;
	goto L64
}
func F_compareDocR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v8 = int32(_a_F_compareDocR_0)
	v9 = v7 & v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v12 = v10 & v8
	if v9 == v12 {
		v14 = int32(14)
		v15 = int32(base.Ui32(v7) >> (uint(v14) % 32))
		v17 = int32(base.Ui32(v10) >> (uint(v14) % 32))
		if v15 == v17 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v19 == v20 {
				return int32(0)
			} else {
				if base.Ui32(v20) < base.Ui32(v19) {
					v27 = int32(1)
				} else {
					v27 = int32(-1)
				}
				return v27
			}
		} else {
			if base.Ui32(v17) < base.Ui32(v15) {
				v32 = int32(1)
			} else {
				v32 = int32(-1)
			}
			return v32
		}
	} else {
		if base.Ui32(v12) < base.Ui32(v9) {
			v37 = int32(1)
		} else {
			v37 = int32(-1)
		}
		return v37
	}
}
func F_compare_scalars(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v10, v11, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 < int32(0) {
			v21 = int32(1)
		} else {
			v21 = int32(0) - v14
		}
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
		if v22 != 0 {
			v23 = v21
		} else {
			v23 = v14
		}
		if v23 != 0 {
			v42 = v23
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v27 = v24 + v7<<(uint(int32(2))%32)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			if v28 < v6 {
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v6
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v32 = v31
			} else {
				v32 = v24
			}
			v35 = v32 + v6<<(uint(int32(2))%32)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 < v7 {
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v7
			} else {
			}
			v42 = v7 - v6
		}
		return v42
	}
}
func F_computeRegionDelta(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v13 = int32(-1)
	v15 = base.B2i32(base.Ui32(l3) < base.Ui32(l5))
	if base.Ui32(l3) < base.Ui32(l5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l3
	goto L3
L2:
	;
	v16 = v13
	goto L3
L3:
	;
	if base.Ui32(l3) < base.Ui32(l5) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v138 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L5:
	;
	v17 = l5
	goto L7
L6:
	;
	v17 = l3
	goto L7
L7:
	;
	if base.Ui32(l4) < base.Ui32(l6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v19 = l4
	goto L10
L9:
	;
	v19 = l6
	goto L10
L10:
	;
	if base.Ui32(v19) <= base.Ui32(v17) {
		v138 = v16
		v139 = v13
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = v17
	v30 = v16
	goto L12
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
	if v36 != v38 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v138 = v127
	v139 = v128
	goto L4
L14:
	;
	if v30 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v67 = v26
	v71 = v30
	goto L16
L16:
	;
	v77 = v67 + int32(1)
	if v77 < v19 {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	v42 = v26
	goto L19
L18:
	;
	v42 = v30
	goto L19
L19:
	;
	v46 = v26
	goto L20
L20:
	;
	v56 = v46 + int32(1)
	if v19 <= v56 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v67 = v56
	v71 = v42
	goto L16
L22:
	;
	v138 = v42
	v139 = int32(-1)
	goto L4
L23:
	;
	goto L24
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v56))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v56))))
	if v60 != v62 {
		v46 = v56
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v79 = v19
	goto L28
L27:
	;
	v79 = v77
	goto L28
L28:
	;
	v85 = v67
	goto L29
L29:
	;
	if v85 == v79-int32(1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v71 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v102 = v79
	goto L31
L33:
	;
	goto L34
L34:
	;
	v96 = v85 + int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v96))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v96))))
	if v98 == v100 {
		v85 = v96
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v102 = v96
	goto L31
L36:
	;
	if v102 < v19 {
		v26 = v102
		v30 = v127
		goto L12
	} else {
		goto L44
	}
L37:
	;
	v127 = int32(-1)
	v128 = v67
	goto L36
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v102-v67) < base.Ui32(int32(5)) {
		v127 = v71
		v128 = v67
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = l0 + int32(16) + v109
	v111 = v67 - v71
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+2)) = uint16(v111)
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v71)
	v115 = v111 & int32(_a_F_computeRegionDelta_0)
	if v115 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	base.MemoryCopy(m, v110+int32(4), l2+v71, v115)
	goto L43
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115 + v109 + int32(4)
	v124 = int32(-1)
	v127 = v124
	v128 = v124
	goto L36
L44:
	;
	goto L13
L45:
	;
	v145 = v19
	goto L47
L46:
	;
	v145 = v138
	goto L47
L47:
	;
	v146 = base.B2i32(base.Ui32(l6) < base.Ui32(l4))
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v147 = v145
	goto L50
L49:
	;
	v147 = v138
	goto L50
L50:
	;
	if int32(0) <= v147 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = l0 + v150
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+16)) = uint16(v147)
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	return
L54:
	;
	v153 = l4
	goto L56
L55:
	;
	v153 = v139
	goto L56
L56:
	;
	if v153 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v156 = l4
	goto L59
L58:
	;
	v156 = v153
	goto L59
L59:
	;
	v157 = v156 - v147
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+18)) = uint16(v157)
	v160 = v157 & int32(_a_F_computeRegionDelta_0)
	if v160 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v151+int32(20), v147+l2, v160)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v160 + v150 + int32(4)
	goto L53
}
func F_compute_remaining_iovec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v7 = l1
	v8 = l2
	v9 = l3
	goto L2
L1:
	;
	if l0 == v7 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v9) < base.Ui32(v11) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v17 = v8 - int32(1)
	if v17 != 0 {
		v7 = v7 + int32(8)
		v8 = v17
		v9 = v9 - v11
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27 + v9
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30 - v9
	return v8
L7:
	;
	v22 = v8 << (uint(int32(3)) % 32)
	if v22 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, l0, v7, v22)
	goto L6
}
func F_compute_trivial_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 float64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 float64
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+78)))
	if v18 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+76)))
	v24 = int32(_a_F_compute_trivial_stats_0)
	v29 = base.B2i32(v21 < int32(0))
	v30 = base.B2i32(v21&v24 == v24)
	goto L3
L2:
	;
	v29 = v5
	v30 = v5
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v15 + int32(16)
	return
L5:
	;
	v37 = float64(0)
	v41 = v5
	v42 = v5
	v44 = v5
	goto L6
L6:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if int32(0) < v97 {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	return
L9:
	;
	v51 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v44, v15+int32(15))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v53 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v100 = v44 + int32(1)
	if v100 != l2 {
		v37 = v94
		v41 = v96
		v42 = v97
		v44 = v100
		goto L6
	} else {
		goto L28
	}
L12:
	;
	v94 = v37
	v96 = v41 + int32(1)
	v97 = v42
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = v42 + int32(1)
	if v30 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v60 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if v29 == int32(0) {
		v94 = v37
		v96 = v41
		v97 = v59
		goto L11
	} else {
		goto L27
	}
L18:
	;
	v94 = base.F64_add(v37, base.F64_convert_i32_u(v84))
	v96 = v41
	v97 = v59
	goto L11
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v84 = int32(6)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v76 = int32(1)
	if v60&v76 != 0 {
		v84 = int32(base.Ui32(v60) >> (uint(v76) % 32))
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v71 = int32(18)
	if v64 == v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = v71
	goto L25
L24:
	;
	v75 = int32(2)
	goto L25
L25:
	;
	v84 = v75
	goto L18
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v84 = int32(base.Ui32(v80) >> (uint(int32(2)) % 32))
	goto L18
L27:
	;
	v89 = F_strlen(m, v51)
	mBase = m.M
	v94 = base.F64_add(v37, base.F64_convert_i32_u(v89+int32(1)))
	v96 = v41
	v97 = v59
	goto L11
L28:
	;
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v125
	goto L4
L30:
	;
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v104)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v96), base.F64_convert_i32_s(l2)))
	if v29 != 0 {
		v125 = base.I32_trunc_sat_f64_s(base.F64_div(v94, base.F64_convert_i32_u(v97)))
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v96 <= int32(0) {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+76)))
	v125 = v115
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v120)
	if v29 != 0 {
		v125 = int32(0)
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+76)))
	v125 = v124
	goto L29
}
func F_construct_empty_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(64)
		return v4
	}
}
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	v10 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if v10 <= l2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L70
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L66
	}
L3:
	;
	if base.Ui32(int32(7)) <= base.Ui32(l2) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L62
	}
L6:
	;
	v27 = F_ArrayGetNItemsSafe(m, l2, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_ArrayCheckBounds(m, l2, l3, l4)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v27 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	m.G0 = v21 + int32(48)
	return v269
L11:
	;
	v56 = v10
	v58 = v10
	v64 = v10
	goto L19
L12:
	;
	goto L11
L13:
	;
	goto L14
L14:
	;
	v40 = F_palloc0(m, int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(64)
	v269 = v40
	goto L10
L16:
	;
	v235 = v225 + v227
	v236 = F_palloc0(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L54
	}
L17:
	;
	v225 = (l2<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v227 = v172
	v234 = int32(0)
	goto L16
L18:
	;
	v201 = base.I32_div_s(v27+int32(7), int32(8))
	v208 = (v201 + l2<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v225 = v208
	v227 = v191
	v234 = v208
	goto L16
L19:
	;
	if l1 == int32(0) {
		v107 = v56
		v115 = v64
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v115 == int32(0) {
		goto L17
	} else {
		goto L53
	}
L21:
	;
	if base.B2i32(l6 == int32(-1)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v56))))
	if v68 != int32(1) {
		v107 = v56
		v115 = v64
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = v56 + int32(1)
	if v72 == v27 {
		v191 = v58
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v83 = v72
	goto L25
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v83))))
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v107 = v83
	v115 = int32(1)
	goto L21
L27:
	;
	v95 = v83 + int32(1)
	if v27 != v95 {
		v83 = v95
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v191 = v58
	goto L18
L31:
	;
	v159 = v157 + v58
	switch l8 - int32(99) {
	case 0:
		v172 = v159
		goto L47
	case 1:
		goto L49
	default:
		goto L48
	case 6:
		goto L50
	}
L32:
	;
	if int32(0) < l6 {
		v157 = l6
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v129 = l0 + v107<<(uint(int32(2))%32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = F_pg_detoast_datum(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0+v107<<(uint(int32(2))%32))))
	v124 = F_strlen(m, v123)
	mBase = m.M
	v157 = v124 + int32(1)
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v134 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	if base.Ui32((v138-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v157 = int32(6)
		goto L31
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v134&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v145 = int32(18)
	if v138 == v145 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v149 = v145
	goto L43
L42:
	;
	v149 = int32(2)
	goto L43
L43:
	;
	v157 = v149
	goto L31
L44:
	;
	v157 = int32(base.Ui32(v134) >> (uint(int32(1)) % 32))
	goto L31
L45:
	;
	goto L46
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v157 = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
	goto L31
L47:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v172) {
		goto L1
	} else {
		goto L51
	}
L48:
	;
	v172 = (v159 + int32(1)) & int32(-2)
	goto L47
L49:
	;
	v172 = (v159 + int32(7)) & int32(-8)
	goto L47
L50:
	;
	v172 = (v159 + int32(3)) & int32(-4)
	goto L47
L51:
	;
	v176 = v107 + int32(1)
	if v176 != v27 {
		v56 = v176
		v58 = v172
		v64 = v115
		goto L19
	} else {
		goto L52
	}
L52:
	;
	goto L20
L53:
	;
	v191 = v172
	goto L18
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v236)+8)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = l2
	v241 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v235 << (uint(v241) % 32)
	v245 = v236 + int32(16)
	v247 = l2 << (uint(v241) % 32)
	v248 = int32(0)
	v249 = base.B2i32(v247 == v248)
	if v249 == v248 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v245, l3, v247)
	goto L57
L56:
	;
	goto L57
L57:
	;
	if v249 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	base.MemoryCopy(m, v247+v245, l4, v247)
	goto L60
L59:
	;
	goto L60
L60:
	;
	F_CopyArrayEls(m, v236, l0, l1, v27, l6, l7, l8, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v269 = v236
	goto L10
L62:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_0), v21)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3511), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_3), v21+int32(16))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3516), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_construct_md_array_4), v21+int32(32))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3546), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_convert_tuples_by_position(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = F_palloc0(m, int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v419 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L2:
	;
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v21
	v28 = int32(1)
	v31 = F_palloc0(m, v21<<(uint(v28)%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v31
	if v21 <= int32(0) {
		v198 = v4
		v202 = v28
		v205 = v4
		v208 = v31
		v209 = v4
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v211 <= v198 {
		v326 = v202
		v329 = v205
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v39 = v4
	v43 = v28
	v44 = v4
	v46 = v4
	v49 = v31
	v50 = v4
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v58 = l1 + v52<<(uint(int32(4))%32) + v44*int32(100)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+111)))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L27
	}
L9:
	;
	goto L8
L10:
	;
	v63 = v50 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 <= v39 {
		v118 = v39
		v125 = v46
		v128 = v49
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v141 = v39
	v145 = v43
	v148 = v46
	v151 = v49
	v152 = v50
	goto L12
L12:
	;
	v155 = v44 + int32(1)
	if v21 != v155 {
		v39 = v141
		v43 = v145
		v44 = v155
		v46 = v148
		v49 = v151
		v50 = v152
		goto L7
	} else {
		goto L26
	}
L13:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+v44<<(uint(int32(1))%32)))))
	v141 = v118
	v145 = base.B2i32(v134 != int32(0)) & v43
	v148 = v125
	v151 = v128
	v152 = v63
	goto L12
L14:
	;
	v66 = int32(20)
	v67 = v58 + v66
	v76 = v39
	goto L15
L15:
	;
	v91 = l0 + v64<<(uint(int32(4))%32) + v66 + v76*int32(100)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+91)))
	if v92 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v118 = v64
	v125 = v46
	v128 = v49
	goto L13
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	if v95 != v96 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v113 = v76 + int32(1)
	if v113 != v64 {
		v76 = v113
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	if int32(0) <= v98 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	if v98 != v101 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v103 = int32(1)
	v109 = v76 + v103
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v44<<(uint(v103)%32)))) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v118 = v109
	v125 = v46 + v103
	v128 = v111
	goto L13
L24:
	;
	goto L23
L25:
	;
	goto L16
L26:
	;
	v198 = v141
	v202 = v145
	v205 = v148
	v208 = v151
	v209 = v152
	goto L5
L27:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v17+int32(-16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	v173 = F_format_type_with_typemod(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	v177 = F_format_type_with_typemod(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v58 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v173
	F_errdetail(m, int32(_a_F_convert_tuples_by_position_1), v17+int32(-32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(124), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	if v326&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L35:
	;
	v215 = (v211 - v198) & int32(3)
	if v215 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v198-v211) {
		v326 = v257
		v329 = v260
		goto L34
	} else {
		goto L43
	}
L37:
	;
	v252 = v198
	v257 = v202
	v260 = v205
	goto L36
L38:
	;
	goto L39
L39:
	;
	v221 = v198
	v226 = v202
	v227 = int32(0)
	v229 = v205
	goto L40
L40:
	;
	v235 = int32(1)
	v236 = v221 + v235
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v221<<(uint(int32(4))%32))+29)))
	v241 = v240 & v226
	v246 = v229 + (v240^int32(-1))&v235
	v248 = v227 + v235
	if v248 != v215 {
		v221 = v236
		v226 = v241
		v227 = v248
		v229 = v246
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v252 = v236
	v257 = v241
	v260 = v246
	goto L36
L42:
	;
	goto L41
L43:
	;
	v271 = v252
	v276 = v257
	v279 = v260
	goto L44
L44:
	;
	v285 = int32(4)
	v287 = l0 + v271<<(uint(v285)%32)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+77)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+61)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+45)))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+29)))
	v295 = v288 & (v289 & (v290 & v291)) & v276
	v296 = int32(-1)
	v298 = int32(1)
	v315 = v279 + (v291^v296)&v298 + (v290^v296)&v298 + (v289^v296)&v298 + (v288^v296)&v298
	v317 = v271 + v285
	if v317 != v211 {
		v271 = v317
		v276 = v295
		v279 = v315
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v326 = v295
	v329 = v315
	goto L34
L46:
	;
	goto L45
L47:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v211 != v337 {
		v419 = v23
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L2
	} else {
		goto L68
	}
L50:
	;
	m.G0 = v19 - int32(-64)
	goto L1
L51:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v339 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v342 = int32(20)
	v350 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_pfree(m, v208)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L2
	} else {
		goto L66
	}
L55:
	;
	v364 = v350 << (uint(int32(4)) % 32)
	v365 = l0 + v342 + v364
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+8)))
	if v366 != 0 {
		v419 = v23
		goto L50
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v367 = int32(1)
	v368 = v350 + v367
	v372 = int32(*(*int16)(unsafe.Add(mBase, uint32(v208+v350<<(uint(v367)%32)))))
	if v368 != v372 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v372 != 0 {
		v419 = v23
		goto L50
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v368 != v339 {
		v350 = v368
		goto L55
	} else {
		goto L65
	}
L61:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+9)))
	if v374 != int32(1) {
		v419 = v23
		goto L50
	} else {
		goto L62
	}
L62:
	;
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+4)))
	v378 = v364 + (l1 + v342)
	v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378)+4)))
	if v377 != v379 {
		v419 = v23
		goto L50
	} else {
		goto L63
	}
L63:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+12)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+12)))
	if v381 != v382 {
		v419 = v23
		goto L50
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	goto L56
L66:
	;
	F_pfree(m, v23)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v419 = int32(0)
	goto L50
L68:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v17+int32(-48))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v329
	F_errdetail(m, int32(_a_F_convert_tuples_by_position_4), v19)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(149), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	return int32(0)
L74:
	;
	goto L75
L75:
	;
	v454 = F_palloc(m, int32(28))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+8)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v454)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = l0
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v461 = v459 + int32(1)
	v464 = F_palloc(m, v461<<(uint(int32(2))%32))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+20)) = v464
	v467 = F_palloc(m, v461)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+24)) = v467
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v472 = v470 + int32(1)
	v475 = F_palloc(m, v472<<(uint(int32(2))%32))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+12)) = v475
	v478 = F_palloc(m, v472)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+16)) = v478
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v481))) = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v454)+16))
	v485 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v485)
	return v454
}
func F_copy_addr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	v2 = l1
	if v2 != int32(10) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(l4) < base.Ui32(v31) {
		goto L1
	} else {
		goto L12
	}
L3:
	;
	if v2 != int32(2) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v17 = l2 + int32(8)
	v18 = int32(16)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	switch v19 - int32(254) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		v31 = v18
		v32 = v17
		goto L2
	}
L6:
	;
	v13 = int32(4)
	v31 = v13
	v32 = l2 + v13
	goto L2
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
	v31 = v18
	v32 = v17
	goto L2
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v25&int32(15) != int32(2) {
		v31 = v18
		v32 = v17
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v22 < int32(-64) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v31 = v18
	v32 = v17
	goto L2
L11:
	;
	goto L7
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
	if base.Ui32(int32(512)) <= base.Ui32(v31) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
	goto L1
L14:
	;
	if v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v41 = v32 + v31
	if (v32^l3)&int32(3) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	base.MemoryCopy(m, v32, l3, v31)
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L13
L20:
	;
	if base.Ui32(v173) < base.Ui32(v41) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	if v32&int32(3) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		goto L45
	} else {
		goto L46
	}
L24:
	;
	v77 = v41 & int32(-4)
	if base.Ui32(v41) < base.Ui32(int32(64)) {
		v127 = v71
		v128 = v72
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v71 = l3
	v72 = v32
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v31 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v71 = l3
	v72 = v32
	goto L24
L29:
	;
	goto L30
L30:
	;
	v54 = l3
	v55 = v32
	goto L31
L31:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v61 = int32(1)
	v62 = v54 + v61
	v64 = v55 + v61
	if v64&int32(3) == int32(0) {
		v71 = v62
		v72 = v64
		goto L24
	} else {
		goto L33
	}
L32:
	;
	v71 = v62
	v72 = v64
	goto L24
L33:
	;
	if base.Ui32(v64) < base.Ui32(v41) {
		v54 = v62
		v55 = v64
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if base.Ui32(v77) <= base.Ui32(v128) {
		v172 = v127
		v173 = v128
		goto L20
	} else {
		goto L41
	}
L36:
	;
	v81 = v77 + int32(-64)
	if base.Ui32(v81) < base.Ui32(v72) {
		v127 = v71
		v128 = v72
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v84 = v71
	v85 = v72
	goto L38
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+36)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+40)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+44)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+48)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+52)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+56)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+60)) = v119
	v121 = int32(-64)
	v122 = v84 - v121
	v124 = v85 - v121
	if base.Ui32(v124) <= base.Ui32(v81) {
		v84 = v122
		v85 = v124
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v127 = v122
	v128 = v124
	goto L35
L40:
	;
	goto L39
L41:
	;
	v134 = v127
	v135 = v128
	goto L42
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v139
	v141 = int32(4)
	v142 = v134 + v141
	v144 = v135 + v141
	if base.Ui32(v144) < base.Ui32(v77) {
		v134 = v142
		v135 = v144
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v172 = v142
	v173 = v144
	goto L20
L44:
	;
	goto L43
L45:
	;
	v172 = l3
	v173 = v32
	goto L20
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v31) < base.Ui32(int32(4)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v172 = l3
	v173 = v32
	goto L20
L49:
	;
	goto L50
L50:
	;
	v153 = l3
	v154 = v32
	goto L51
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v158)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)) = uint8(v160)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)) = uint8(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)) = uint8(v164)
	v166 = int32(4)
	v167 = v153 + v166
	v169 = v154 + v166
	if base.Ui32(v169) <= base.Ui32(v41-int32(4)) {
		v153 = v167
		v154 = v169
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v172 = v167
	v173 = v169
	goto L20
L53:
	;
	goto L52
L54:
	;
	v179 = v172
	v180 = v173
	goto L57
L55:
	;
	goto L56
L56:
	;
	goto L13
L57:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v184)
	v186 = int32(1)
	v189 = v180 + v186
	if v189 != v41 {
		v179 = v179 + v186
		v180 = v189
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
}
func F_copytup_heap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	v5 = F_minimal_tuple_from_heap_tuple(m, l1, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_GetMemoryChunkSpace(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v11 - base.I64_extend_i32_u(v9)
			return v5
		}
	}
}
func F_cost_material(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 float64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v17 float64
	_ = v17
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_cost_material[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	v13 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[1]))
	v17 = base.F64_add(base.F64_mul(base.F64_add(v13, v13), l4), base.F64_sub(l3, l2))
	v25 = base.F64_mul(l4, base.F64_convert_i32_u((l5+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v25, base.F64_convert_i32_u(v10<<(uint(int32(10))%32))) != 0 {
		v31 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[2]))
		v37 = base.F64_add(base.F64_mul(v31, base.F64_ceil(base.F64_mul(v25, float64(0.0001220703125)))), v17)
	} else {
		v37 = v17
	}
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_material[3])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(l2, v37)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v39 ^ int32(1))
	return
}
func F_cost_samplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v113 float64
	_ = v113
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v160 float64
	_ = v160
	v9 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = F_GetTsmRoutine(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v35 = v21 + v22<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v35 = v28 + v29<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	return
L6:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = l3 + int32(8)
	goto L9
L8:
	;
	v45 = l2 + int32(16)
	goto L9
L9:
	;
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	F_get_tablespace_page_costs(m, v48, v19+int32(8), v19+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v144)+24))
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v148 = *(*float64)(unsafe.Add(mBase, _c_F_cost_samplescan[0]))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v144)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v154 = base.F64_add(v149, base.F64_add(v143, float64(0)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v154
	v156 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	if v56 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v60
	if v59 == int32(0) {
		v113 = v9
		v120 = float64(0)
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v136 = v125
	v143 = v126
	goto L11
L15:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v136 = base.F64_add(v113, v121)
	v143 = base.F64_add(v120, v123)
	goto L11
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v69 <= int32(0) {
		v113 = v9
		v120 = float64(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v76 = int32(0)
	goto L18
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v76<<(uint(int32(2))%32))))
	v96 = F_cost_qual_eval_walker(m, v93, v19+int32(24))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v19)+40))
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	v113 = v102
	v120 = v103
	goto L15
L20:
	;
	v99 = v76 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v99 < v100 {
		v76 = v99
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v160 = v58
	goto L24
L23:
	;
	v160 = v57
	goto L24
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v154, base.F64_add(base.F64_mul(v145, v156), base.F64_add(base.F64_mul(v146, base.F64_add(v136, v148)), base.F64_add(base.F64_mul(v160, base.F64_convert_i32_u(v55)), float64(0)))))
	m.G0 = v19 + int32(48)
	return
}
func F_cost_subqueryscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v138 int32
	_ = v138
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v149 float64
	_ = v149
	var v150 int32
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	v9 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v28)+32))
	v30 = int32(0)
	v33 = F_clauselist_selectivity(m, l1, v26, v30, v30, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+184))
	v23 = F_list_concat_copy(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+184))
	v26 = v25
	goto L1
L5:
	;
	return
L6:
	;
	v26 = v23
	goto L1
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v51
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v50)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v53
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v50)+56))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v55
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v35 = base.F64_mul(v29, v33)
	if base.F64_gt(v35, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807))) != 0 {
		v48 = float64(1e+100)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = float64(1)
	if base.F64_le(v35, v44) != 0 {
		v48 = v44
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = base.F64_nearest(v35)
	goto L7
L11:
	;
	v58 = int32(0)
	goto L13
L12:
	;
	v58 = l4
	goto L13
L13:
	;
	if v58 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	m.G0 = v19 + int32(32)
	return
L17:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v150)+24))
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v138)+32))
	v154 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subqueryscan[0]))
	v155 = *(*float64)(unsafe.Add(mBase, uint32(v150)+16))
	v156 = base.F64_add(v149, v155)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v156, v141)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_add(v156, base.F64_add(base.F64_mul(v151, v142), base.F64_mul(v152, base.F64_add(v144, v154)))), v145)
	goto L16
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v62
	if v61 == int32(0) {
		v115 = v50
		v118 = v53
		v119 = v48
		v121 = v9
		v122 = v55
		v126 = float64(0)
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v138 = v50
	v141 = v53
	v142 = v48
	v144 = v131
	v145 = v55
	v149 = v132
	goto L17
L21:
	;
	v127 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v129 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v138 = v115
	v141 = v118
	v142 = v119
	v144 = base.F64_add(v121, v127)
	v145 = v122
	v149 = base.F64_add(v126, v129)
	goto L17
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v71 <= int32(0) {
		v115 = v50
		v118 = v53
		v119 = v48
		v121 = v9
		v122 = v55
		v126 = float64(0)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v80<<(uint(int32(2))%32))))
	v98 = F_cost_qual_eval_walker(m, v95, v19+int32(8))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v115 = v107
	v118 = v105
	v119 = v106
	v121 = v108
	v122 = v104
	v126 = v109
	goto L21
L26:
	;
	v101 = v80 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v101 < v102 {
		v80 = v101
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_create_empty_pathtarget(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc0(m, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(277)
		return v3
	}
}
func F_create_final_distinct_paths(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 float64
	_ = v114
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v43 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v33 = F_get_sortgrouplist_exprs(m, v31, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v42 = v30
	goto L1
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	if v25 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+36)))
	if v26 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+318)))
	if v27 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	return int32(0)
L9:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v38 = int32(0)
	v40 = F_estimate_num_groups(m, l0, v33, v37, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v42 = v40
	goto L1
L11:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v348 != 0 {
		goto L115
	} else {
		goto L116
	}
L12:
	;
	if v88 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L13:
	;
	v88 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v52 <= int32(0) {
		v80 = int32(1)
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v88 = v80
	goto L12
L17:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v58 = v52
	goto L20
L19:
	;
	v58 = v55
	goto L20
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v61 = int32(0)
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v71 = int32(0)
	v72 = base.B2i32(v70 != v71)
	if v70 == v71 {
		v80 = v72
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v80 = v72
	goto L16
L23:
	;
	v76 = v61 + int32(1)
	if v76 != v58 {
		v61 = v76
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)))
	if v92 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v106 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L27:
	;
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v105 = v91
	goto L26
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v96 = v95
	goto L32
L31:
	;
	v96 = v4
	goto L32
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v97 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v100 = v98
	goto L35
L34:
	;
	v100 = int32(0)
	goto L35
L35:
	;
	if v96 < v100 {
		v105 = v97
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 <= int32(0) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	if v91 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = float64(-1)
	goto L41
L40:
	;
	v114 = float64(1)
	goto L41
L41:
	;
	v128 = v4
	goto L42
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+64))
	v138 = F_get_useful_pathkeys_for_distinct(m, l0, v105, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L11
L44:
	;
	v328 = v128 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v328 < v329 {
		v128 = v328
		goto L42
	} else {
		goto L113
	}
L45:
	;
	if v138 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v143 <= v142 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v147 = v142
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v147<<(uint(int32(2))%32))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v136)+64))
	v170 = v20 + int32(12)
	if v167 == v168 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L44
L50:
	;
	v307 = v147 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v307 < v308 {
		v147 = v307
		goto L48
	} else {
		goto L112
	}
L51:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v276 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L52:
	;
	if v248 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v236
	v248 = int32(1)
	goto L52
L54:
	;
	if v167 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v167 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(0)
	v248 = int32(1)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(0)
	v248 = int32(1)
	goto L52
L59:
	;
	goto L60
L60:
	;
	if v168 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v188
	v248 = v188
	goto L52
L62:
	;
	goto L63
L63:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v192 = int32(0)
	if v192 < v191 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v195 = v191
	goto L66
L65:
	;
	v195 = v192
	goto L66
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v201 = int32(0)
	goto L67
L67:
	;
	if v201 < v196 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v212 = v208 + v201<<(uint(int32(2))%32)
	goto L71
L70:
	;
	v212 = int32(0)
	goto L71
L71:
	;
	if v201 == v195 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v195
	v248 = base.B2i32(v212 == int32(0))
	goto L52
L73:
	;
	goto L74
L74:
	;
	v218 = base.B2i32(v212 == int32(0))
	if v212 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = v218
	goto L52
L76:
	;
	goto L77
L77:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	if v222 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = v218
	goto L52
L79:
	;
	goto L80
L80:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v222+v201<<(uint(int32(2))%32))))
	if v226 != v230 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = int32(0)
	goto L52
L82:
	;
	v201 = v201 + int32(1)
	goto L67
L84:
	;
	v273 = v136
	goto L51
L85:
	;
	goto L86
L86:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[0])))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v136 == v22 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v260&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v260 = v250
	goto L87
L89:
	;
	goto L90
L90:
	;
	if v251 == int32(0) {
		goto L50
	} else {
		goto L91
	}
L91:
	;
	v255 = int32(1)
	if v250&v255 == int32(0) {
		goto L50
	} else {
		goto L92
	}
L92:
	;
	v260 = v255
	goto L87
L93:
	;
	v264 = v251
	goto L95
L94:
	;
	v264 = int32(0)
	goto L95
L95:
	;
	if v264 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v267 = F_create_sort_path(m, l2, v136, v167, v114)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v269 = F_create_incremental_sort_path(m, l0, l2, v136, v167, v251, v114)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	if v267 != 0 {
		v273 = v267
		goto L51
	} else {
		goto L100
	}
L100:
	;
	goto L50
L101:
	;
	if v269 == int32(0) {
		goto L50
	} else {
		goto L102
	}
L102:
	;
	v273 = v269
	goto L51
L103:
	;
	v279 = int32(0)
	v285 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v299 = F_create_upper_unique_path(m, l2, v273, v298, v42)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L110
	}
L106:
	;
	v287 = int32(0)
	v289 = F_makeConst(m, int32(20), int32(-1), v279, int32(8), v285, v287, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	v294 = F_create_limit_path(m, l2, v273, v279, v289, int32(0), int64(0), int64(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_add_path(m, l2, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L50
L110:
	;
	F_add_path(m, l2, v299)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	goto L50
L112:
	;
	goto L49
L113:
	;
	goto L43
L114:
	;
	m.G0 = v20 + int32(16)
	return l2
L115:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)))
	if v349 != 0 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v357 = int32(0)
	if v356 == v357 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[1])))
	if v351&int32(1) == int32(0) {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if v394 == int32(0) {
		goto L114
	} else {
		goto L133
	}
L121:
	;
	v394 = int32(1)
	goto L120
L122:
	;
	goto L123
L123:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v364 <= int32(0) {
		v388 = int32(1)
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v394 = v388
	goto L120
L125:
	;
	v367 = int32(0)
	if v367 < v364 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v370 = v364
	goto L128
L127:
	;
	v370 = v367
	goto L128
L128:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v375 = v357
	goto L129
L129:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v371+v375<<(uint(int32(2))%32))))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+18)))
	if v380 != int32(1) {
		v388 = v380
		goto L124
	} else {
		goto L131
	}
L130:
	;
	v388 = v380
	goto L124
L131:
	;
	v384 = v375 + int32(1)
	if v384 != v370 {
		v375 = v384
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v399 = int32(0)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v403 = F_create_agg_path(m, l0, l2, v22, v397, int32(2), v399, v400, v399, v399, v42)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	F_add_path(m, l2, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L135
	}
L135:
	;
	goto L114
}
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v12 != int32(331) {
		v19 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 == int32(0) {
		v80 = v5
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v15 != 0 {
		v19 = l2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = l2
	goto L6
L5:
	;
	v18 = int32(0)
	goto L6
L6:
	;
	v19 = v18
	goto L1
L7:
	;
	v83 = F_palloc0(m, int32(80))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L23
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v25 <= int32(0) {
		v80 = v5
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v35 = int32(1)
	v37 = v5
	v38 = v5
	goto L10
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v80 = v65
	goto L7
L12:
	;
	v46 = F_replace_nestloop_params_mutator(m, v44, l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v50 = v44
	goto L14
L14:
	;
	v52 = int32(0)
	v54 = F_makeTargetEntry(m, v50, base.I32_extend16_s(v35), v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L17
	}
L15:
	;
	return int32(0)
L16:
	;
	v50 = v46
	goto L14
L17:
	;
	if v28 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v28+v35<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v61
	goto L20
L19:
	;
	goto L20
L20:
	;
	v65 = F_lappend(m, v38, v54)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v68 = v37 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v68 < v69 {
		v35 = v35 + int32(1)
		v37 = v68
		v38 = v65
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L11
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+72)) = l3
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+56)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v83)+52)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v83)+48)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v83)+44)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(331)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v94
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v83)+8)) = v96
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
	v100 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v83)+24)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+36)) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v102
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+37)) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+37)) = uint8(v108)
	return v83
}
func F_create_groupingsets_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 float64
	_ = v167
	var v170 int32
	_ = v170
	var v171 float64
	_ = v171
	var v177 float64
	_ = v177
	var v184 float64
	_ = v184
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 float64
	_ = v199
	var v200 float64
	_ = v200
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v21 = F_palloc0(m, int32(96))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(310)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(365)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v32
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v36 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v41 = v39
	goto L5
L4:
	;
	v41 = int32(0)
	goto L5
L5:
	;
	v42 = int32(1)
	v43 = v41 & v42
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+21)) = uint8(v43)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v45
	switch l4 - v42 {
	case 0:
		goto L9
	default:
		v71 = l4
		v72 = v8
		goto L6
	case 2:
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v72
	if l6 != 0 {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v71 = v53
	v72 = v70
	goto L6
L8:
	;
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if l5 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v71 = int32(1)
	v72 = v8
	goto L6
L11:
	;
	goto L12
L12:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v54 != v53 {
		v71 = v53
		v72 = v8
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v59 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v71 = int32(0)
	v72 = v8
	goto L6
L15:
	;
	v71 = int32(3)
	v72 = v8
	goto L6
L16:
	;
	goto L17
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v66 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(2)
	goto L20
L19:
	;
	v69 = int32(3)
	goto L20
L20:
	;
	v71 = v69
	v72 = v8
	goto L6
L21:
	;
	v77 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l6)+32)))
	v79 = v77
	goto L23
L22:
	;
	v79 = int64(0)
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v79
	if l5 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = base.F64_add(v229, v230)
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	v234 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v233, base.F64_add(base.F64_mul(v234, v235), v237))
	m.G0 = v18 + int32(144)
	return v21
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v84 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = int32(1)
	v99 = int32(1)
	v101 = v8
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v101<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v113 = v112
	goto L31
L30:
	;
	v113 = int32(0)
	goto L31
L31:
	;
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v211 = v101 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v211 < v212 {
		v89 = v207
		v99 = int32(0)
		v101 = v211
		goto L27
	} else {
		goto L47
	}
L33:
	;
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	F_cost_agg(m, v21, l0, v71, l6, v113, v114, l3, v115, v116, v117, v118, base.F64_convert_i32_s(v120))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	if (v126|v89)&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v207 = v124 & v89
	goto L32
L37:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v195 + v196
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
	v200 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v199, v200)
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = base.F64_add(v203, v204)
	v207 = v194
	goto L32
L38:
	;
	v131 = int32(1)
	if v126&v131 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v147 = int32(0)
	v149 = v18 + int32(72)
	v151 = float64(0)
	v152 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_create_groupingsets_path[0]))
	v160 = m.G0
	v161 = int32(16)
	v162 = v160 - v161
	m.G0 = v162
	F_cost_tuplesort(m, v162+int32(8), v162, v152, v154, v151, v157, float64(-1))
	mBase = m.M
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v162)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v149)+32)) = v152
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_groupingsets_path[1])))
	v171 = base.F64_add(v151, v167)
	*(*float64)(unsafe.Add(mBase, uint32(v149)+48)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v149)+40)) = v147 + (v170 ^ int32(1))
	v177 = *(*float64)(unsafe.Add(mBase, uint32(v162)))
	*(*float64)(unsafe.Add(mBase, uint32(v149)+56)) = base.F64_add(v171, v177)
	m.G0 = v162 + v161
	goto L45
L41:
	;
	v134 = int32(2)
	goto L43
L42:
	;
	v134 = v131
	goto L43
L43:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v137 = float64(0)
	v139 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+32))
	F_cost_agg(m, v18, l0, v134, l6, v113, v135, l3, int32(0), v137, v137, v139, base.F64_convert_i32_s(v141))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v194 = v145 & v89
	goto L37
L45:
	;
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v18)+120))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v18)+128))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v18)+104))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+32))
	F_cost_agg(m, v18, l0, int32(1), l6, v113, v184, l3, v185, v186, v187, v188, base.F64_convert_i32_s(v190))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v194 = v147
	goto L37
L47:
	;
	goto L28
}
func F_create_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v251 int32
	_ = v251
	var v254 float64
	_ = v254
	var v256 int32
	_ = v256
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v268 float64
	_ = v268
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v280 float64
	_ = v280
	var v284 float64
	_ = v284
	var v291 float64
	_ = v291
	var v295 float64
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 float64
	_ = v304
	var v306 float64
	_ = v306
	var v309 float64
	_ = v309
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v462 float64
	_ = v462
	var v466 float64
	_ = v466
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v500 float64
	_ = v500
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v509 float64
	_ = v509
	var v535 float64
	_ = v535
	var v538 float64
	_ = v538
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v548 float64
	_ = v548
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 float64
	_ = v595
	var v596 float64
	_ = v596
	var v613 float64
	_ = v613
	var v622 float64
	_ = v622
	var v624 float64
	_ = v624
	var v625 int32
	_ = v625
	var v626 float64
	_ = v626
	var v628 float64
	_ = v628
	var v629 float64
	_ = v629
	var v631 float64
	_ = v631
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = l7
	v32 = F_palloc0(m, int32(96))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(298)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v38 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v41 = v39
	goto L5
L4:
	;
	v41 = int32(0)
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+228))
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v45 = v43
	goto L8
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v45 = v44
	goto L8
L8:
	;
	v46 = int32(0)
	if base.B2i32(v41 == v46)|base.B2i32(v45 == v46) != 0 {
		v91 = v46
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v91 != 0 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	goto L9
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v56 < v57 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v56
	goto L14
L13:
	;
	v59 = v57
	goto L14
L14:
	;
	if v59 <= int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = int32(1)
	goto L17
L16:
	;
	v62 = v59
	goto L17
L17:
	;
	v63 = int32(8)
	v68 = int32(0)
	goto L18
L18:
	;
	v75 = v68 << (uint(int32(2)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45+v63+v75)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41+v63+v75)))
	v80 = v77 & v79
	v82 = base.B2i32(v80 != int32(0))
	if v80 != 0 {
		v91 = v82
		goto L10
	} else {
		goto L20
	}
L19:
	;
	v91 = v82
	goto L10
L20:
	;
	v84 = v68 + int32(1)
	if v84 != v62 {
		v68 = v84
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v92 = F_get_param_path_clause_serials(m, l6)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(356)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v198
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v204 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v201, l9, v28+int32(12))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	if l7 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v156
	goto L24
L27:
	;
	v156 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v98 <= v97 {
		v156 = v97
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v113 = int32(0)
	v114 = v97
	goto L31
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v113<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+56))
	v133 = F_bms_is_member(m, v132, v92)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v156 = v139
	goto L26
L33:
	;
	if v133 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v137 = F_lappend(m, v114, v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v139 = v114
	goto L36
L36:
	;
	v141 = v113 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v141 < v142 {
		v113 = v141
		v114 = v139
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v139 = v137
	goto L36
L38:
	;
	goto L32
L39:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+20)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v204
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v209 != int32(1) {
		v216 = v200
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v218 = v216 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)) = uint8(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v220
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v224)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v228
	v230 = m.G0
	v232 = v230 - int32(32)
	m.G0 = v232
	v234 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v242 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v212 != int32(1) {
		v216 = v200
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v216 = v215
	goto L40
L43:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v248)))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if int32(0) < v251 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v248 = v242 + int32(8)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v248 = v245 + int32(16)
	goto L43
L47:
	;
	v254 = base.F64_convert_i32_u(v251)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_nestloop_path[0])))
	if v256 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if base.F64_le(v239, float64(0)) != 0 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v262 = base.F64_add(base.F64_mul(v254, float64(-0.3)), float64(1))
	if base.F64_gt(v262, float64(0)) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v268 = v254
	goto L52
L52:
	;
	v270 = float64(1e+100)
	v271 = base.F64_div(v249, v268)
	if base.F64_gt(v271, v270)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v271)&int64(9223372036854775807))) != 0 {
		v284 = v270
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v266 = v262
	goto L55
L54:
	;
	v266 = math.Float64frombits(uint64(0x8000000000000000))
	goto L55
L55:
	;
	v268 = base.F64_add(v266, v254)
	goto L52
L56:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v284
	goto L49
L57:
	;
	v280 = float64(1)
	if base.F64_le(v271, v280) != 0 {
		v284 = v280
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v284 = base.F64_nearest(v271)
	goto L56
L59:
	;
	v291 = float64(1)
	goto L61
L60:
	;
	v291 = v239
	goto L61
L61:
	;
	if base.F64_le(v237, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v295 = float64(1)
	goto L64
L63:
	;
	v295 = v237
	goto L64
L64:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v296&int32(-2) != int32(4) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v543 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v232)+24)) = v543
	v548 = float64(0)
	if v542 == int32(0) {
		v613 = v548
		v622 = v548
		goto L123
	} else {
		goto L124
	}
L66:
	;
	v535 = v234
	v538 = base.F64_mul(v295, v291)
	goto L65
L67:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v301 != int32(1) {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v304 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v306 = base.F64_nearest(base.F64_mul(v295, v304))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v312 = base.F64_div(float64(2), base.F64_add(v309, float64(1)))
	v313 = base.F64_mul(base.F64_mul(v291, v306), v312)
	v314 = base.F64_sub(v295, v306)
	v315 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	v316 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	if v317 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v496 = base.F64_ge(v314, float64(1))
	if v496 != 0 {
		goto L113
	} else {
		goto L114
	}
L72:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v318 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	switch v323 - int32(341) {
	case 0, 1:
		v330 = l6
		goto L74
	default:
		goto L71
	case 3:
		goto L75
	}
L74:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	if v331 == int32(0) {
		goto L71
	} else {
		goto L77
	}
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l6)+72))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v327 != int32(280) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v330 = v326
	goto L74
L77:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v334 <= int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330)+76))
	v338 = int32(0)
	v343 = v338
	v347 = v338
	goto L80
L79:
	;
	v462 = base.F64_add(base.F64_mul(v316, v312), v234)
	if base.F64_gt(v306, float64(1)) != 0 {
		goto L110
	} else {
		goto L111
	}
L80:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v343<<(uint(int32(2))%32))))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+8))
	v372 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	v374 = F_bms_is_subset(m, v373, v322)
	mBase = m.M
	if v374 == v372 {
		v385 = v372
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v347 == int32(0) {
		goto L71
	} else {
		goto L109
	}
L82:
	;
	if v385 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L82
L84:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	v378 = F_bms_overlap(m, v371, v377)
	mBase = m.M
	if v378 == int32(0) {
		v385 = v372
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v369)+40))
	v382 = F_bms_overlap(m, v371, v381)
	mBase = m.M
	v385 = v382 ^ int32(1)
	goto L83
L86:
	;
	if v337 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	goto L88
L88:
	;
	v448 = v343 + int32(1)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v448 < v449 {
		v343 = v448
		goto L80
	} else {
		goto L108
	}
L89:
	;
	if v438 == int32(0) {
		goto L71
	} else {
		goto L106
	}
L90:
	;
	goto L89
L91:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v391 <= int32(0) {
		v438 = int32(0)
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v438 = int32(0)
	goto L90
L94:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v369)+60))
	v395 = int32(0)
	if v395 < v391 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v398 = v391
	goto L97
L96:
	;
	v398 = v395
	goto L97
L97:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v402 = int32(0)
	goto L98
L98:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v399+v402<<(uint(int32(2))%32))))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+12)))
	if v412 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L93
L100:
	;
	v423 = v402 + int32(1)
	if v423 != v398 {
		v402 = v423
		goto L98
	} else {
		goto L105
	}
L101:
	;
	v413 = int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	if v369 == v414 {
		v438 = v413
		goto L90
	} else {
		goto L102
	}
L102:
	;
	if v394 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+60))
	if v418 == v394 {
		v438 = v413
		goto L90
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	goto L99
L106:
	;
	v442 = int32(1)
	v444 = v343 + v442
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v444 < v445 {
		v343 = v444
		v347 = v442
		goto L80
	} else {
		goto L107
	}
L107:
	;
	goto L79
L108:
	;
	goto L81
L109:
	;
	goto L79
L110:
	;
	v466 = base.F64_add(base.F64_mul(base.F64_mul(v315, base.F64_add(v306, float64(-1))), v312), v462)
	goto L112
L111:
	;
	v466 = v462
	goto L112
L112:
	;
	v535 = base.F64_add(base.F64_div(base.F64_mul(v315, v314), v291), v466)
	v538 = v313
	goto L65
L113:
	;
	v497 = v306
	goto L115
L114:
	;
	v497 = base.F64_add(v306, float64(-1))
	goto L115
L115:
	;
	v500 = base.F64_add(v234, v316)
	if base.F64_gt(v497, float64(0)) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v504 = base.F64_add(base.F64_mul(base.F64_mul(v315, v497), v312), v500)
	goto L118
L117:
	;
	v504 = v500
	goto L118
L118:
	;
	v506 = base.F64_add(base.F64_mul(v314, v291), v313)
	if v496 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v509 = base.F64_add(v314, float64(-1))
	goto L121
L120:
	;
	v509 = v314
	goto L121
L121:
	;
	if base.F64_gt(v509, float64(0)) == int32(0) {
		v535 = v504
		v538 = v506
		goto L65
	} else {
		goto L122
	}
L122:
	;
	v535 = base.F64_add(base.F64_mul(v509, v315), v504)
	v538 = v506
	goto L65
L123:
	;
	v624 = *(*float64)(unsafe.Add(mBase, _c_F_create_nestloop_path[1]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v626 = *(*float64)(unsafe.Add(mBase, uint32(v625)+24))
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v625)+16))
	v629 = base.F64_add(base.F64_add(v235, v622), v628)
	*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = v629
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v32)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+56)) = base.F64_add(v629, base.F64_add(base.F64_mul(v626, v631), base.F64_add(base.F64_mul(base.F64_add(v613, v624), v538), v535)))
	m.G0 = v232 + int32(32)
	m.G0 = v28 + int32(16)
	return v32
L124:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v553 <= int32(0) {
		v613 = v548
		v622 = float64(0)
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v560 = int32(0)
	goto L126
L126:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582+v560<<(uint(int32(2))%32))))
	v589 = F_cost_qual_eval_walker(m, v586, v232+int32(8))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v232)+24))
	v596 = *(*float64)(unsafe.Add(mBase, uint32(v232)+16))
	v613 = v595
	v622 = v596
	goto L123
L128:
	;
	v592 = v560 + int32(1)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v592 < v593 {
		v560 = v592
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
}
func F_create_secmsg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
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
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v4
	if v18 <= v4 {
		v101 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v108 = v18 + int32(3)
	v109 = F_palloc(m, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v24 = v18 & int32(3)
	v26 = l0 + int32(132)
	if base.Ui32(int32(4)) <= base.Ui32(v18) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = v4
	v38 = v4
	v42 = v4
	goto L6
L4:
	;
	v63 = v4
	v67 = v4
	goto L5
L5:
	;
	v76 = v63
	v80 = v67
	v85 = v4
	goto L10
L6:
	;
	v44 = v34 + v26
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)))
	v52 = v38 + v45 + v47 + v49 + v51
	v53 = int32(4)
	v54 = v34 + v53
	v56 = v42 + v53
	if v56 != v18&int32(2147483644) {
		v34 = v54
		v38 = v52
		v42 = v56
		goto L6
	} else {
		goto L8
	}
L7:
	;
	if v24 == int32(0) {
		v101 = v52
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v63 = v54
	v67 = v52
	goto L5
L10:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v26))))
	v88 = v80 + v87
	v89 = int32(1)
	v92 = v85 + v89
	if v92 != v24 {
		v76 = v76 + v89
		v80 = v88
		v85 = v92
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v101 = v88
	goto L1
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v113)
	if v18 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, v109+int32(1), l0+int32(132), v18)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v121 = int32(8)
	v127 = v101<<(uint(v121)%32) | int32(base.Ui32(v101&int32(_a_F_create_secmsg_0))>>(uint(v121)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v18+v109)+1)) = uint16(v127)
	v130 = l2 - v108
	v132 = v130 - int32(2)
	if v132 < v121 {
		v304 = int32(-12)
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v108 != 0 {
		goto L73
	} else {
		goto L74
	}
L19:
	;
	v135 = F_palloc(m, l2)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v137 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v137)
	v140 = v135 + int32(1)
	v141 = int32(0)
	v145 = m.G0
	v147 = v145 - int32(16)
	m.G0 = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v141
	v153 = F_open(m, int32(_a_F_create_secmsg_1), v141, v147)
	mBase = m.M
	if v153 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v186 != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v156 = int32(1)
	if v132 == int32(0) {
		v179 = v156
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v186 = v141
	goto L24
L24:
	;
	m.G0 = v147 + int32(16)
	goto L21
L25:
	;
	v181 = F_close(m, v153)
	mBase = m.M
	v186 = v179
	goto L24
L26:
	;
	v159 = v140
	v160 = v132
	goto L27
L27:
	;
	v165 = F_read(m, v153, v159, v160)
	mBase = m.M
	if v165 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v179 = v156
	goto L25
L29:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v169 == int32(27) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v174 = v160 - v165
	if v174 != 0 {
		v159 = v159 + v165
		v160 = v174
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v179 = int32(0)
	goto L25
L33:
	;
	goto L28
L34:
	;
	v195 = v140
	goto L38
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v135)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L71
	}
L37:
	;
	if l2 != 0 {
		goto L68
	} else {
		goto L69
	}
L38:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v205 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v132)+1)) = uint8(v267)
	if v108 != 0 {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v263 = int32(1)
	goto L42
L41:
	;
	v208 = int32(0)
	v212 = m.G0
	v214 = v212 - int32(16)
	m.G0 = v214
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v208
	v220 = F_open(m, int32(_a_F_create_secmsg_1), v208, v214)
	mBase = m.M
	if v220 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v264 = v263 + v195
	if base.Ui32(v264) < base.Ui32(v140+v132) {
		v195 = v264
		goto L38
	} else {
		goto L57
	}
L43:
	;
	if v253 == int32(0) {
		goto L37
	} else {
		goto L56
	}
L44:
	;
	goto L48
L45:
	;
	v253 = v208
	goto L46
L46:
	;
	m.G0 = v214 + int32(16)
	goto L43
L47:
	;
	v248 = F_close(m, v220)
	mBase = m.M
	v253 = v246
	goto L46
L48:
	;
	v226 = v195
	v227 = int32(1)
	goto L49
L49:
	;
	v232 = F_read(m, v220, v226, v227)
	mBase = m.M
	if v232 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v246 = int32(1)
	goto L47
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v236 == int32(27) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v241 = v227 - v232
	if v241 != 0 {
		v226 = v226 + v232
		v227 = v241
		goto L49
	} else {
		goto L55
	}
L54:
	;
	v246 = int32(0)
	goto L47
L55:
	;
	goto L50
L56:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v263 = base.B2i32(v260 != int32(0))
	goto L42
L57:
	;
	goto L39
L58:
	;
	base.MemoryCopy(m, v135+v130, v109, v108)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v277 = F_pgp_mpi_create(m, v135, l2<<(uint(int32(3))%32)-int32(6), v16+int32(12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	if l2 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_pfree(m, v135)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L13
	} else {
		goto L66
	}
L63:
	;
	base.MemoryFill(m, v135, int32(0), l2)
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	v304 = v277
	goto L18
L67:
	;
	goto L36
L68:
	;
	base.MemoryFill(m, v135, int32(0), l2)
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v304 = int32(-17)
	goto L18
L72:
	;
	F_pfree(m, v109)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L76
	}
L73:
	;
	base.MemoryFill(m, v109, int32(0), v108)
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	if int32(0) <= v304 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v320
	goto L79
L78:
	;
	goto L79
L79:
	;
	m.G0 = v16 + int32(16)
	return v304
}
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = l1 + int32(4)
	v7 = F_palloc(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v6 << (uint(int32(2)) % 32)
		if l1 != 0 {
			base.MemoryCopy(m, v7+int32(4), l0, l1)
		} else {
		}
		return v7
	}
}
func F_cursor_to_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_text_to_cstring(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v19 = F_pg_detoast_datum_packed(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_text_to_cstring(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_GetPortalByName(m, v15)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							if v26 != 0 {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
								if v28 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(258))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
											F_errmsg(m, int32(_a_F_cursor_to_xmlschema_0), v8+int32(16))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3111), int32(_a_F_cursor_to_xmlschema_2))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
									v31 = int32(0)
									v34 = F_map_sql_table_to_xmlschema(m, v28, v31, base.B2i32(v17 != v31), v21)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										v36 = F_strlen(m, v34)
										mBase = m.M
										v38 = v36 + int32(1)
										v39 = F_SPI_palloc(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											if v38 != 0 {
												base.MemoryCopy(m, v39, v34, v38)
											} else {
											}
											v42 = F_SPI_finish(m)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int32(0)
											} else {
												v44 = F_cstring_to_text(m, v39)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(32)
													return v44
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(259))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
										F_errmsg(m, int32(_a_F_cursor_to_xmlschema_3), v8)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3107), int32(_a_F_cursor_to_xmlschema_2))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
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
		}
	}
}
func F_cword_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_errstart_cold(m, int32(21), int32(_a_F_cword_is_not_variable_0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = F_NameListToString(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v17
				F_errmsg(m, int32(_a_F_cword_is_not_variable_1), v7)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = F_plpgsql_scanner_errposition(m, l1, l2)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cword_is_not_variable_2), int32(2646), int32(_a_F_cword_is_not_variable_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
}
