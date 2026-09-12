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
			v28 = F___time(m)
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
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[1])) = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_CheckDeadLockAlert[0])) = v3
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
	var v27 int32
	_ = v27
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
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
	v27 = v3
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v27<<(uint(int32(2))%32))))
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
	v237 = v27 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v237 < v238 {
		v27 = v237
		goto L4
	} else {
		goto L61
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v140 == int32(0) {
		goto L38
	} else {
		goto L39
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
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	F_CheckDuplicateColumnOrPathNames(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L37
	}
L13:
	;
	v123 = F_lappend(m, v40, v39)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L30
	} else {
		goto L36
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
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v65 == int32(0) {
		v84 = v64
		v85 = v65
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	if v85-v84 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L18
L20:
	;
	if v64 != v65 {
		v84 = v64
		v85 = v65
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v69 = v39
	v70 = v61
	goto L22
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v74 == int32(0) {
		v84 = v73
		v85 = v74
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v84 = v73
	v85 = v74
	goto L19
L24:
	;
	v77 = int32(1)
	if v73 == v74 {
		v69 = v69 + v77
		v70 = v70 + v77
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v88 = v52 + int32(1)
	if v88 != v43 {
		v52 = v88
		goto L16
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L17
L29:
	;
	goto L13
L30:
	;
	return
L31:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v98
	F_errmsg(m, int32(_a_F_CheckDuplicateColumnOrPathNames_0), v13)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	F_parser_errposition(m, v103, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_CheckDuplicateColumnOrPathNames_1), int32(190), int32(_a_F_CheckDuplicateColumnOrPathNames_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v123
	goto L12
L37:
	;
	goto L6
L38:
	;
	v223 = F_lappend(m, v140, v139)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L30
	} else {
		goto L60
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v143 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v152 = int32(0)
	goto L41
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v146+v152<<(uint(int32(2))%32))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v165 == int32(0) {
		v184 = v164
		v185 = v165
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L30
	} else {
		goto L55
	}
L43:
	;
	if v185-v184 != 0 {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	goto L43
L45:
	;
	if v164 != v165 {
		v184 = v164
		v185 = v165
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v169 = v139
	v170 = v161
	goto L47
L47:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v174 == int32(0) {
		v184 = v173
		v185 = v174
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v184 = v173
	v185 = v174
	goto L44
L49:
	;
	v177 = int32(1)
	if v173 == v174 {
		v169 = v169 + v177
		v170 = v170 + v177
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v188 = v152 + int32(1)
	if v188 != v143 {
		v152 = v188
		goto L41
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L42
L54:
	;
	goto L38
L55:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L30
	} else {
		goto L56
	}
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v197
	F_errmsg(m, int32(_a_F_CheckDuplicateColumnOrPathNames_0), v13+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L30
	} else {
		goto L57
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	F_parser_errposition(m, v204, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L30
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CheckDuplicateColumnOrPathNames_1), int32(203), int32(_a_F_CheckDuplicateColumnOrPathNames_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L30
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v223
	goto L6
L61:
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
	var v28 int32
	_ = v28
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
	v28 = v5
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22+v28<<(uint(int32(2))%32))))
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
	v44 = v28 + int32(1)
	if v44 != v21 {
		v28 = v44
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
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v11 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v67 = l1 - int32(1)
					if base.Ui32(v67) <= base.Ui32(int32(3)) {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
						v75 = v74
					} else {
						v75 = int32(_a_F_CheckSelectLocking_0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v75
					F_errmsg(m, int32(_a_F_CheckSelectLocking_1), v6+int32(80))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3404), int32(_a_F_CheckSelectLocking_3))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
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
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v97 = l1 - int32(1)
						if base.Ui32(v97) <= base.Ui32(int32(3)) {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
							v105 = v104
						} else {
							v105 = int32(_a_F_CheckSelectLocking_0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v105
						F_errmsg(m, int32(_a_F_CheckSelectLocking_4), v6-int32(-64))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3411), int32(_a_F_CheckSelectLocking_3))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
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
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v97 = l1 - int32(1)
							if base.Ui32(v97) <= base.Ui32(int32(3)) {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
								v105 = v104
							} else {
								v105 = int32(_a_F_CheckSelectLocking_0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v105
							F_errmsg(m, int32(_a_F_CheckSelectLocking_4), v6-int32(-64))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3411), int32(_a_F_CheckSelectLocking_3))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
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
						v120 = m.ExcPending
						if v120 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								v127 = l1 - int32(1)
								if base.Ui32(v127) <= base.Ui32(int32(3)) {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v127<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
									v135 = v134
								} else {
									v135 = int32(_a_F_CheckSelectLocking_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v135
								F_errmsg(m, int32(_a_F_CheckSelectLocking_5), v6+int32(48))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3418), int32(_a_F_CheckSelectLocking_3))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
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
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									v157 = l1 - int32(1)
									if base.Ui32(v157) <= base.Ui32(int32(3)) {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v157<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
										v165 = v164
									} else {
										v165 = int32(_a_F_CheckSelectLocking_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v165
									F_errmsg(m, int32(_a_F_CheckSelectLocking_6), v6)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3425), int32(_a_F_CheckSelectLocking_3))
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
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
								v178 = m.ExcPending
								if v178 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										v185 = l1 - int32(1)
										if base.Ui32(v185) <= base.Ui32(int32(3)) {
											v192 = *(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
											v193 = v192
										} else {
											v193 = int32(_a_F_CheckSelectLocking_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v193
										F_errmsg(m, int32(_a_F_CheckSelectLocking_7), v6+int32(16))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3432), int32(_a_F_CheckSelectLocking_3))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
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
									v208 = m.ExcPending
									if v208 != 0 {
										return
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v211 = m.ExcPending
										if v211 != 0 {
											return
										} else {
											v215 = l1 - int32(1)
											if base.Ui32(v215) <= base.Ui32(int32(3)) {
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v215<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
												v223 = v222
											} else {
												v223 = int32(_a_F_CheckSelectLocking_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v223
											F_errmsg(m, int32(_a_F_CheckSelectLocking_8), v6+int32(32))
											mBase = m.M
											v229 = m.ExcPending
											if v229 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3439), int32(_a_F_CheckSelectLocking_3))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
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
				v37 = l1 - int32(1)
				if base.Ui32(v37) <= base.Ui32(int32(3)) {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_CheckSelectLocking[0])))
					v45 = v44
				} else {
					v45 = int32(_a_F_CheckSelectLocking_0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v45
				F_errmsg(m, int32(_a_F_CheckSelectLocking_9), v6+int32(96))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckSelectLocking_2), int32(3397), int32(_a_F_CheckSelectLocking_3))
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
	}
}
func F_charin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
	if v5&int32(3) == int32(0) {
		v30 = v5
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return base.I32_extend8_s(v99)
L2:
	;
	if v63 != int32(4) {
		v99 = v6
		goto L1
	} else {
		goto L19
	}
L3:
	;
	v63 = v55 - v5
	goto L2
L4:
	;
	v34 = v30
	goto L13
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v14 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v63 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v19 = v5
	goto L9
L9:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v55 = v23
	goto L3
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v49 = v34
	goto L16
L15:
	;
	goto L14
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v55 = v49
	goto L3
L18:
	;
	goto L17
L19:
	;
	if v6&int32(255) != int32(92) {
		v99 = v6
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	if v70&int32(248) != int32(48) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(92)
L22:
	;
	goto L23
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	if v77&int32(248) != int32(48) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(92)
L25:
	;
	goto L26
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
	if v85&int32(248) != int32(48) {
		v99 = int32(92)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v99 = v70<<(uint(int32(6))%32) + v77<<(uint(int32(3))%32) + v85 + int32(80)
	goto L1
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
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
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	v205 = int32(*(*int16)(unsafe.Add(mBase, uint32(v204)+120)))
	if v205 <= int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L60
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L31
	} else {
		goto L55
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L31
	} else {
		goto L50
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L31
	} else {
		goto L45
	}
L8:
	;
	v26 = v4
	v29 = v4
	v30 = v4
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v26<<(uint(int32(2))%32))))
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
	if v99 == int32(0) {
		goto L7
	} else {
		goto L30
	}
L14:
	;
	v99 = v51 + int32(1)
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
	v99 = int32(0)
	goto L13
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v103 = F_bms_is_member(m, v99, v29)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v102 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v118 = F_lappend_int(m, v117, v99)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L31
	} else {
		goto L43
	}
L34:
	;
	if v103 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v103 != 0 {
		goto L5
	} else {
		goto L41
	}
L37:
	;
	v109 = F_bms_is_member(m, v99, v30)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v109 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v111 = F_bms_add_member(m, v29, v99)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v115 = v111
	v116 = v30
	goto L33
L41:
	;
	v113 = F_bms_add_member(m, v30, v99)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v115 = v29
	v116 = v113
	goto L33
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
	v122 = v26 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v122 < v123 {
		v26 = v122
		v29 = v115
		v30 = v116
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
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v145 + int32(4)
	F_errmsg(m, int32(_a_F_checkInsertTargets_0), v15)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1073), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v167 = m.ExcPending
	if v167 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v40
	F_errmsg(m, int32(_a_F_checkInsertTargets_3), v15+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1088), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	v188 = m.ExcPending
	if v188 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v40
	F_errmsg(m, int32(_a_F_checkInsertTargets_3), v15+int32(32))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	F_parser_errposition(m, l0, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_checkInsertTargets_1), int32(1099), int32(_a_F_checkInsertTargets_2))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	v211 = v4
	v213 = v4
	goto L61
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+52))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v228 = v221 + v222<<(uint(int32(4))%32) + v211*int32(100)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+111)))
	if v229 == int32(1) {
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
	if v257 != v205 {
		v211 = v257
		v213 = v259
		goto L61
	} else {
		goto L71
	}
L64:
	;
	v257 = v211 + int32(1)
	v259 = v213
	goto L63
L65:
	;
	goto L66
L66:
	;
	v235 = F_palloc0(m, int32(20))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = int32(81)
	v241 = F_pstrdup(m, v228+int32(24))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v235)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v241
	v248 = F_lappend(m, v213, v235)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v252 = v211 + int32(1)
	v253 = F_lappend_int(m, v250, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v253
	v257 = v252
	v259 = v248
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
	if v4 == int32(-1) {
	} else {
		if int32(63) < v4 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(64)
		}
	}
	return int32(1)
}
func F_check_backtrace_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(3) == v4 {
		v35 = v11
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v69 = int32(_a_F_check_backtrace_functions_0)
	v73 = m.G0
	v75 = v73 - int32(32)
	v76 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+24)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v76
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_backtrace_functions[0])))
	if v84 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v68 = v60 - v11
	goto L1
L3:
	;
	v39 = v35
	goto L12
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v68 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v24 = v11
	goto L8
L8:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v60 = v28
	goto L2
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v54 = v39
	goto L15
L14:
	;
	goto L13
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v60 = v54
	goto L2
L17:
	;
	goto L16
L18:
	;
	if v68 != v152 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v152 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_backtrace_functions[1])))
	if v88 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = v11
	goto L25
L23:
	;
	goto L24
L24:
	;
	v102 = v69
	v103 = v84
	goto L28
L25:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v98 == v84 {
		v92 = v92 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v152 = v92 - v11
	goto L18
L27:
	;
	goto L26
L28:
	;
	v110 = v75 + int32(base.Ui32(v103)>>(uint(int32(3))%32))&int32(28)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v111 | v112<<(uint(v103)%32)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v116 != 0 {
		v102 = v102 + v112
		v103 = v116
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v119 == int32(0) {
		v144 = v11
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v152 = v144 - v11
	goto L18
L32:
	;
	v123 = v11
	v124 = v119
	goto L33
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(base.Ui32(v124)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v132)>>(uint(v124)%32))&int32(1) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v144 = v140
	goto L31
L35:
	;
	v144 = v123
	goto L31
L36:
	;
	goto L37
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	v140 = v123 + int32(1)
	if v138 != 0 {
		v123 = v140
		v124 = v138
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[3])) = v156
	v161 = F_format_elog_string(m, int32(_a_F_check_backtrace_functions_1), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v168 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	return int32(0)
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_backtrace_functions[4])) = v161
	return int32(0)
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(1)
L45:
	;
	goto L46
L46:
	;
	v177 = F_guc_malloc(m, v68+int32(2))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v177 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	goto L50
L50:
	;
	if v68 <= int32(0) {
		v263 = v4
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v269 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v263+v177))) = uint16(v269)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v177
	return int32(1)
L52:
	;
	v185 = int32(1)
	if v68 == v185 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v68&v185 == int32(0) {
		v263 = v240
		goto L51
	} else {
		goto L66
	}
L54:
	;
	v238 = int32(0)
	v240 = v4
	goto L53
L55:
	;
	goto L56
L56:
	;
	v192 = int32(0)
	v197 = v192
	v198 = v192
	v199 = v4
	goto L57
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v197))))
	switch v207 - int32(9) {
	case 0, 1, 23:
		v215 = v199
		goto L59
	default:
		goto L61
	case 35:
		v210 = int32(0)
		goto L60
	}
L58:
	;
	v238 = v231
	v240 = v228
	goto L53
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218+v197)+1)))
	switch v220 - int32(9) {
	case 0, 1, 23:
		v228 = v215
		goto L62
	default:
		goto L64
	case 35:
		v223 = int32(0)
		goto L63
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v199+v177))) = uint8(v210)
	v215 = v199 + int32(1)
	goto L59
L61:
	;
	v210 = v207
	goto L60
L62:
	;
	v230 = int32(2)
	v231 = v197 + v230
	v233 = v198 + v230
	if v233 != v68&int32(2147483646) {
		v197 = v231
		v198 = v233
		v199 = v228
		goto L57
	} else {
		goto L65
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v215+v177))) = uint8(v223)
	v228 = v215 + int32(1)
	goto L62
L64:
	;
	v223 = v220
	goto L63
L65:
	;
	goto L58
L66:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+v238))))
	switch v250 - int32(9) {
	case 0, 1, 23:
		v263 = v240
		goto L51
	default:
		goto L68
	case 35:
		v253 = int32(0)
		goto L67
	}
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v240+v177))) = uint8(v253)
	v263 = v240 + int32(1)
	goto L51
L68:
	;
	v253 = v250
	goto L67
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	v22 = v3
	v23 = v3
	v26 = v17
	goto L4
L4:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v26 <= v30 {
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
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v43 = v30
	goto L8
L8:
	;
	if v22 == v43 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31+v43<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v58 == int32(0) {
		v77 = v57
		v78 = v58
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v78-v77 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	goto L11
L13:
	;
	if v57 != v58 {
		v77 = v57
		v78 = v58
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v62 = v36
	v63 = v54
	goto L15
L15:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v66
		v78 = v67
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v77 = v66
	v78 = v67
	goto L12
L17:
	;
	v70 = int32(1)
	if v66 == v67 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v81 = v43 + int32(1)
	if v81 == v26 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L9
L22:
	;
	v43 = v81
	goto L8
L23:
	;
	return
L24:
	;
	F_errcode(m, int32(_a_F_check_duplicates_in_publist_0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v54
	F_errmsg(m, int32(_a_F_check_duplicates_in_publist_1), v13)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_check_duplicates_in_publist_2), int32(2383), int32(_a_F_check_duplicates_in_publist_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v112 = F_cstring_to_text(m, v36)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L23
	} else {
		goto L31
	}
L29:
	;
	v117 = v23
	goto L30
L30:
	;
	v119 = v22 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v119 < v120 {
		v22 = v119
		v23 = v117
		v26 = v120
		goto L4
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v23<<(uint(int32(2))%32)))) = v112
	v117 = v23 + int32(1)
	goto L30
L32:
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
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
	if base.Ui32(l0) < base.Ui32(int32(42)) {
		if base.B2i32(l0 != l3)&base.B2i32(int32(0) <= l3) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v47 = int32(3)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v47)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v51
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v47)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
				F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_0), v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1806), int32(_a_F_check_encoding_conversion_args_2))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
					F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_3), v7+int32(-32))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1808), int32(_a_F_check_encoding_conversion_args_2))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
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
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = int32(3)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v86)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v90
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l4<<(uint(v86)%32))+uint32(_c_F_check_encoding_conversion_args[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v96
						F_errmsg_internal(m, int32(_a_F_check_encoding_conversion_args_4), v7+int32(-48))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1812), int32(_a_F_check_encoding_conversion_args_2))
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
								F_errfinish(m, int32(_a_F_check_encoding_conversion_args_1), int32(1814), int32(_a_F_check_encoding_conversion_args_2))
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
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v204 int32
	_ = v204
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v631 int32
	_ = v631
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v659 int32
	_ = v659
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
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
	v678 = m.ExcPending
	if v678 != 0 {
		goto L11
	} else {
		goto L145
	}
L8:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	if v130 != 0 {
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
		goto L13
	default:
		goto L15
	case 5:
		goto L14
	}
L13:
	;
	v122 = F_pg_detoast_datum(m, v73)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L22
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
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109+int32(base.Ui32(v111)>>(uint(int32(2))%32))-int32(1)))))
	goto L20
L20:
	;
	if v117&int32(1) == int32(0) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L7
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v124 == int32(0) {
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
	return v659
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+280)) = int32(4)
	if int32(0) < v30 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	if v30 <= int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v145 = int32(0)
	goto L28
L28:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v145))))
	if v158 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v659 = int32(1)
	goto L24
L30:
	;
	v162 = v145 + int32(1)
	if v30 != v162 {
		v145 = v162
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L25
L34:
	;
	v204 = int32(0)
	goto L37
L35:
	;
	goto L36
L36:
	;
	v269 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L44
	}
L37:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v204))))
	if v224 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v225 = int32(65)
	goto L41
L40:
	;
	v225 = int32(0)
	goto L41
L41:
	;
	v226 = int32(1)
	v227 = v204 + v226
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+v204<<(uint(v226)%32)))))
	v235 = v204 << (uint(int32(2)) % 32)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v28+v235)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v41)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l4+v235)))
	F_ScanKeyEntryInitialize(m, v26+int32(352)+v204*int32(48), v225, base.I32_extend16_s(v227), v232, int32(0), v237, v239, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	if v227 != v30 {
		v204 = v227
		goto L37
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l6)+152))
	if v271 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v274 = F_MakePerTupleExprContext(m, l6)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L48
	}
L46:
	;
	v276 = v271
	goto L47
L47:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = v269
	v280 = v269 + int32(28)
	goto L50
L48:
	;
	v276 = v274
	goto L47
L49:
	;
	F_index_endscan(m, v309)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L11
	} else {
		goto L143
	}
L50:
	;
	v304 = int32(0)
	v309 = F_index_beginscan(m, l0, l1, v26+int32(280), v304, v30, v304)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L11
	} else {
		goto L52
	}
L51:
	;
	if l9 != 0 {
		goto L112
	} else {
		goto L113
	}
L52:
	;
	v313 = int32(0)
	F_index_rescan(m, v309, v26+int32(352), v30, v313, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v318 = F_index_getnext_slot(m, v309, int32(1), v269)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	if v318 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v631 = int32(1)
	goto L49
L56:
	;
	goto L57
L57:
	;
	v344 = v304
	goto L58
L58:
	;
	if l3 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v26)+284))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v26)+288))
	if v502 != 0 {
		goto L88
	} else {
		goto L89
	}
L60:
	;
	goto L59
L61:
	;
	v475 = int32(1)
	v477 = F_index_getnext_slot(m, v309, v475, v269)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L11
	} else {
		goto L85
	}
L62:
	;
	F_FormIndexDatum(m, l2, v269, l6, v26+int32(144), v26+int32(112))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L76
	}
L63:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v348 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v353 = int32(16)
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+2)))
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280))))
	if v351|v352<<(uint(v353)%32) == v356|v357<<(uint(v353)%32) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v367 == int32(0) {
		goto L62
	} else {
		goto L71
	}
L66:
	;
	goto L65
L67:
	;
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+4)))
	if v363 == v364 {
		v367 = int32(1)
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v367 = int32(0)
	goto L66
L70:
	;
	goto L69
L71:
	;
	if v344 == int32(0) {
		v473 = int32(1)
		goto L61
	} else {
		goto L72
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v377 + int32(4)
	F_errmsg_internal(m, int32(_a_F_check_exclusion_or_unique_constraint_3), v26-int32(-64))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(838), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+72)))
	if v397 != int32(1) {
		goto L60
	} else {
		goto L77
	}
L77:
	;
	v400 = int32(0)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v401)+10)))
	if v402 <= v400 {
		goto L60
	} else {
		goto L78
	}
L78:
	;
	v416 = v400
	goto L79
L79:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(112)+v416))))
	if v431 != 0 {
		v473 = v344
		goto L61
	} else {
		goto L81
	}
L80:
	;
	goto L60
L81:
	;
	v433 = v416 << (uint(int32(2)) % 32)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v41+v433)))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v436+v433)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(144)+v433)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l4+v433)))
	v445 = F_OidFunctionCall2Coll(m, v435, v438, v442, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	if v445 == int32(0) {
		v473 = v344
		goto L61
	} else {
		goto L83
	}
L83:
	;
	v450 = v416 + int32(1)
	if v402 != v450 {
		v416 = v450
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	if v477 != 0 {
		v344 = v473
		goto L58
	} else {
		goto L86
	}
L86:
	;
	v631 = v475
	goto L49
L87:
	;
	goto L51
L88:
	;
	v504 = v502
	goto L90
L89:
	;
	v504 = v503
	goto L90
L90:
	;
	if v504 == int32(0) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	if l8 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if l8 != int32(2) {
		goto L87
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	F_index_endscan(m, v309)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L103
	}
L95:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	if v509 == int32(0) {
		goto L87
	} else {
		goto L96
	}
L96:
	;
	v512 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v504))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v512)) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v525 == int32(0) {
		goto L87
	} else {
		goto L102
	}
L99:
	;
	v525 = base.B2i32(base.Ui32(v512) < base.Ui32(v504))
	goto L98
L100:
	;
	goto L101
L101:
	;
	v525 = int32(base.Ui32(v512-v504) >> (uint(int32(31)) % 32))
	goto L98
L102:
	;
	goto L94
L103:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	if v531 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v26)+284))
	F_SpeculativeInsertionWait(m, v532, v531)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L11
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v528 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L50
L108:
	;
	v537 = int32(8)
	goto L110
L109:
	;
	v537 = int32(5)
	goto L110
L110:
	;
	F_XactLockTableWait(m, v504, l0, v280, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	goto L50
L112:
	;
	if l10 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	v545 = F_BuildIndexValueDescription(m, l1, l4, l5)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L118
	}
L115:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v540
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l10)+4)) = uint16(v542)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v631 = int32(0)
	goto L49
L118:
	;
	v551 = F_BuildIndexValueDescription(m, l1, v26+int32(144), v26+int32(112))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(16908482))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	v560 = int32(0)
	v564 = base.B2i32(v545 != v560) & base.B2i32(v551 != v560)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v567 = v565 + int32(4)
	if l7 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v567
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_5), v26+int32(16))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L11
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v567
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_6), v26+int32(48))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L11
	} else {
		goto L134
	}
L125:
	;
	if v564 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v583+int32(4))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L11
	} else {
		goto L132
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v545
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_7), v26)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L11
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_8), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L11
	} else {
		goto L131
	}
L130:
	;
	goto L126
L131:
	;
	goto L126
L132:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(918), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	if v564 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	F_errtableconstraint(m, l0, v610+int32(4))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L11
	} else {
		goto L141
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v545
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_9), v26+int32(32))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_errdetail(m, int32(_a_F_check_exclusion_or_unique_constraint_10), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L11
	} else {
		goto L140
	}
L139:
	;
	goto L135
L140:
	;
	goto L135
L141:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(929), int32(_a_F_check_exclusion_or_unique_constraint_4))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L11
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = v277
	F_ExecDropSingleTupleTableSlot(m, v269)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	v659 = v631
	goto L24
L145:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v682 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(352)
	F_errmsg(m, int32(_a_F_check_exclusion_or_unique_constraint_11), v26+int32(96))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_check_exclusion_or_unique_constraint_1), int32(1173), int32(_a_F_check_exclusion_or_unique_constraint_2))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L11
	} else {
		goto L148
	}
L148:
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
	var v90 int32
	_ = v90
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v90 = v3
		m.G0 = v9 + int32(16)
		return v90
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 == int32(6) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32)-int32(4))))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			if v25 == int32(0) {
				v90 = v3
				m.G0 = v9 + int32(16)
				return v90
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
							v90 = v3
							m.G0 = v9 + int32(16)
							return v90
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
						v90 = v3
						m.G0 = v9 + int32(16)
						return v90
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
				v90 = v87
				m.G0 = v9 + int32(16)
				return v90
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_c_F_check_srf_call_placement[0])))
					v88 = v87
				} else {
					v88 = int32(_a_F_check_srf_call_placement_4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v88
				F_errmsg(m, int32(_a_F_check_srf_call_placement_5), v8+int32(16))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_parser_errposition(m, l0, l2)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_srf_call_placement_2), int32(2681), int32(_a_F_check_srf_call_placement_3))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
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
		v102 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v102)
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
