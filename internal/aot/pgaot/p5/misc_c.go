package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CMPTRGM_SIGNED(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		if base.I32_extend8_s(v5) < base.I32_extend8_s(v6) {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v15 != v16 {
			if base.I32_extend8_s(v15) < base.I32_extend8_s(v16) {
				v23 = int32(-1)
			} else {
				v23 = int32(1)
			}
			return v23
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v26 != v27 {
				if base.I32_extend8_s(v26) < base.I32_extend8_s(v27) {
					v34 = int32(-1)
				} else {
					v34 = int32(1)
				}
				v35 = v34
			} else {
				v35 = int32(0)
			}
			return v35
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v31 = v4
	v32 = v25
	goto L6
L4:
	;
	goto L5
L5:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v31<<(uint(int32(2))%32))))
	v41 = v32 + v38*int32(640)
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
	v56 = v31 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v56 < v57 {
		v31 = v56
		v32 = v53
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	if v42 != l0 {
		v53 = v32
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
		v53 = v32
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
	v52 = *(*int32)(unsafe.Add(mBase, _consts[516]))
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	if v12 <= int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
		if v17 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[113]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+316))
			v25 = base.B2i32(v23 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v25)
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
			v30 = *(*int32)(unsafe.Add(mBase, _consts[416]))
			v32 = *(*int64)(unsafe.Add(mBase, _consts[417]))
			if base.I32_wrap_i64(v28-v32) < v30 {
				m.G0 = v9 + int32(16)
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				v43 = F_LWLockAcquire(m, v39+int32(1024), int32(1))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[113]))
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+248))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v46)+256))
					*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v48
					v51 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					F_LWLockRelease(m, v51+int32(1024))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = int32(4423512)
						v58 = *(*int64)(unsafe.Add(mBase, _consts[417]))
						if v47 < v58 {
							v60 = v58
						} else {
							v60 = v47
						}
						*(*int64)(unsafe.Add(mBase, _consts[417])) = v60
						v63 = *(*int32)(unsafe.Add(mBase, _consts[416]))
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
									*(*int64)(unsafe.Add(mBase, _consts[417])) = v28
									m.G0 = v9 + int32(16)
									return
								} else {
									v72 = F_RequestXLogSwitch(m, int32(1))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _consts[116]))
										if v72&base.I64_extend_i32_s(v75-int32(1)) == int64(0) {
											*(*int64)(unsafe.Add(mBase, _consts[417])) = v28
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
													*(*int64)(unsafe.Add(mBase, _consts[417])) = v28
													m.G0 = v9 + int32(16)
													return
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, _consts[416]))
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v89
													F_errmsg_internal(m, int32(678278), v9)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_errfinish(m, int32(495364), int32(728), int32(66201))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, _consts[417])) = v28
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[719])) = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[86])) = v3
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
	F_errmsg(m, int32(203548), v13)
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
	F_errfinish(m, int32(499182), int32(190), int32(163591))
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
	F_errmsg(m, int32(203548), v13+int32(16))
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
	F_errfinish(m, int32(499182), int32(203), int32(163591))
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
	F_errmsg(m, int32(709247), v11)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	F_errdetail(m, int32(586412), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(494498), int32(809), int32(76225))
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
	F_errmsg(m, int32(709247), v11+int32(32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(84456)
	F_errdetail(m, int32(632366), v11+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(494498), int32(824), int32(76225))
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
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_consts[262])))
						v75 = v74
					} else {
						v75 = int32(374632)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v75
					F_errmsg(m, int32(358538), v6+int32(80))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, int32(498510), int32(3404), int32(335660))
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
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_consts[262])))
							v105 = v104
						} else {
							v105 = int32(374632)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v105
						F_errmsg(m, int32(358499), v6-int32(-64))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							F_errfinish(m, int32(498510), int32(3411), int32(335660))
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
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_consts[262])))
								v105 = v104
							} else {
								v105 = int32(374632)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v105
							F_errmsg(m, int32(358499), v6-int32(-64))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								F_errfinish(m, int32(498510), int32(3411), int32(335660))
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
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v127<<(uint(int32(2))%32))+uint32(_consts[262])))
									v135 = v134
								} else {
									v135 = int32(374632)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v135
								F_errmsg(m, int32(359211), v6+int32(48))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									F_errfinish(m, int32(498510), int32(3418), int32(335660))
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
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v157<<(uint(int32(2))%32))+uint32(_consts[262])))
										v165 = v164
									} else {
										v165 = int32(374632)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v165
									F_errmsg(m, int32(141051), v6)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return
									} else {
										F_errfinish(m, int32(498510), int32(3425), int32(335660))
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
											v192 = *(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32))+uint32(_consts[262])))
											v193 = v192
										} else {
											v193 = int32(374632)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v193
										F_errmsg(m, int32(140810), v6+int32(16))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return
										} else {
											F_errfinish(m, int32(498510), int32(3432), int32(335660))
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
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v215<<(uint(int32(2))%32))+uint32(_consts[262])))
												v223 = v222
											} else {
												v223 = int32(374632)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v223
											F_errmsg(m, int32(74284), v6+int32(32))
											mBase = m.M
											v229 = m.ExcPending
											if v229 != 0 {
												return
											} else {
												F_errfinish(m, int32(498510), int32(3439), int32(335660))
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
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_consts[262])))
					v45 = v44
				} else {
					v45 = int32(374632)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v45
				F_errmsg(m, int32(517804), v6+int32(96))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(498510), int32(3397), int32(335660))
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
func F_ConditionVariableSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v7 != 0 {
		F_s_lock(m, l0, int32(499212), int32(264), int32(313845))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[101]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v25 = v22 + v15*int32(640)
				v27 = v25 + int32(84)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
				if v29 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					v38 = v33
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v29*int32(640))+84)) = v28
					v38 = v29
				}
				if v28 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _consts[101]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					*(*int32)(unsafe.Add(mBase, uint32(v44+v28*int32(640))+88)) = v38
				}
				*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v55 = v22 + v15*int32(640)
				if v55 != 0 {
					F_SetLatch(m, v55+int32(20))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 == int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[101]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v25 = v22 + v15*int32(640)
			v27 = v25 + int32(84)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
			if v29 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v38 = v33
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22+v29*int32(640))+84)) = v28
				v38 = v29
			}
			if v28 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[101]))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				*(*int32)(unsafe.Add(mBase, uint32(v44+v28*int32(640))+88)) = v38
			}
			*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v55 = v22 + v15*int32(640)
			if v55 != 0 {
				F_SetLatch(m, v55+int32(20))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
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
		v7 = *(*int32)(unsafe.Add(mBase, _consts[16]))
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
	var v35 int32
	_ = v35
	var v51 int32
	_ = v51
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v79 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v79 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v79 = int32(0)
				} else {
					v79 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v23) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v79 = v11
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						v79 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v79 = v11
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				} else {
					v79 = v11
				}
			}
		}
	} else {
		if l0 <= int32(5999) {
			v35 = l0 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v35) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v79 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v79 = int32(0)
					} else {
						v79 = v11
					}
				}
			} else {
				if int32(1)<<(uint(v35)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v79 = int32(0)
						} else {
							v79 = v11
						}
					}
				} else {
					v79 = v11
				}
			}
		} else {
			switch l0 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v79 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v79 = int32(0)
			default:
				if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
					v79 = v11
				} else {
					v51 = l0 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v51) {
						v79 = int32(0)
					} else {
						if int32(1)<<(uint(v51)%32)&int32(49153) != 0 {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v85 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v79 != 0 {
		v86 = int32(0)
	} else {
		v86 = v85
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v86
	v90 = int32(0)
	v95 = F_LockAcquireExtended(m, v7+int32(16), l1, v90, int32(1), v7+int32(12), v90)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		return int32(0)
	} else {
		switch v95 {
		case 0, 3:
			m.G0 = v7 + int32(32)
			return base.B2i32(v95 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v102 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v101)+53)) = uint8(v102)
				m.G0 = v7 + int32(32)
				return base.B2i32(v95 != int32(0))
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
	var v79 int32
	_ = v79
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
	F_errmsg_internal(m, int32(46808), v12)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(499672), int32(2536), int32(64462))
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[250]))
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
		v79 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = v79
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
		v79 = v72
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v79 = v72
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
	v93 = *(*int32)(unsafe.Add(mBase, _consts[250]))
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
	v102 = *(*int32)(unsafe.Add(mBase, _consts[249]))
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
	v126 = *(*int32)(unsafe.Add(mBase, _consts[249]))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L18
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L69
	}
L3:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v205 + int32(4)
	F_errmsg(m, int32(71628), v12+int32(-48))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L67
	}
L4:
	;
	m.G0 = v14 - int32(-64)
	return v195
L5:
	;
	v57 = v4
	v62 = v4
	goto L21
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v16 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v195 = v4
	goto L4
L10:
	;
	v195 = v4
	goto L4
L11:
	;
	goto L12
L12:
	;
	v27 = v4
	v29 = v4
	goto L13
L13:
	;
	v37 = l0 + int32(20) + v27<<(uint(int32(4))%32)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v38 != 0 {
		v46 = v29
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v195 = v46
	goto L4
L15:
	;
	v48 = v27 + int32(1)
	if v48 != v19 {
		v27 = v48
		v29 = v46
		goto L13
	} else {
		goto L20
	}
L16:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+10)))
	if v39 != 0 {
		v46 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = F_lappend_int(m, v29, v27+int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v46 = v42
	goto L15
L20:
	;
	goto L14
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v62<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v70 <= v69 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v195 = v184
	goto L4
L23:
	;
	v145 = int32(0)
	if v57 == v145 {
		goto L52
	} else {
		goto L53
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L18
	} else {
		goto L46
	}
L25:
	;
	v76 = v69
	v80 = v70
	goto L26
L26:
	;
	v89 = l0 + int32(20) + v80<<(uint(int32(4))%32) + v76*int32(100)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+91)))
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+90)))
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L28:
	;
	goto L27
L29:
	;
	v94 = v89 + int32(4)
	if v94|v68 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v112 = v80
	goto L31
L31:
	;
	v114 = v76 + int32(1)
	if v114 < v112 {
		v76 = v114
		v80 = v112
		goto L26
	} else {
		goto L43
	}
L32:
	;
	if v108 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L33:
	;
	v100 = int32(-1)
	goto L35
L34:
	;
	v100 = int32(0)
	goto L35
L35:
	;
	if v94 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v101 = int32(1)
	goto L38
L37:
	;
	v101 = v100
	goto L38
L38:
	;
	if v94 == int32(0) {
		v108 = v101
		goto L39
	} else {
		goto L40
	}
L39:
	;
	goto L32
L40:
	;
	if v68 == int32(0) {
		v108 = v101
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v107 = F_strncmp(m, v94, v68, int32(64))
	mBase = m.M
	v108 = v107
	goto L39
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = v111
	goto L31
L43:
	;
	goto L24
L44:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+74)))
	if v117 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	goto L24
L46:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v68
	F_errmsg(m, int32(71934), v14)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(492519), int32(1045), int32(150731))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v183 != 0 {
		goto L2
	} else {
		goto L64
	}
L52:
	;
	v183 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v151 <= int32(0) {
		v176 = v145
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v176
	goto L51
L56:
	;
	v154 = int32(0)
	if v154 < v151 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v157 = v151
	goto L59
L58:
	;
	v157 = v154
	goto L59
L59:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v160 = int32(0)
	goto L60
L60:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158+v160<<(uint(int32(2))%32))))
	v169 = base.B2i32(v168 == v117)
	if v168 == v117 {
		v176 = v169
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v176 = v169
	goto L55
L62:
	;
	v171 = v160 + int32(1)
	if v171 != v157 {
		v160 = v171
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v184 = F_lappend_int(m, v57, v117)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v187 = v62 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v187 < v188 {
		v57 = v184
		v62 = v187
		goto L21
	} else {
		goto L66
	}
L66:
	;
	goto L22
L67:
	;
	F_errfinish(m, int32(492519), int32(1040), int32(150731))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v68
	F_errmsg(m, int32(415610), v12+int32(-32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(492519), int32(1052), int32(150731))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L18
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
	F_errcode(m, int32(393348))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
	F_errmsg(m, int32(274774), v12+int32(-16))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	F_errdetail(m, int32(655995), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(492519), int32(1029), int32(150731))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	v3 = F_strlen(m, l0)
	mBase = m.M
	if v3 <= int32(100) {
		v6 = F_pstrdup(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v12 = F_pg_mbcliplen(m, l0, v3, int32(100))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v16 = F_palloc(m, v12+int32(4))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v18 = F__emscripten_memcpy_bulkmem(m, v16, l0, v12)
					mBase = m.M
					v19 = v18
				} else {
					v19 = v16
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19+v12))) = int32(3026478)
				return v19
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
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
	return v461
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v23 == int32(0) {
		v461 = v2
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
	F_errmsg(m, int32(274440), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(498840), int32(1581), int32(63865))
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
	v69 = v64
	v70 = v62
	v73 = v2
	goto L16
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v61 = v60
	goto L14
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v83 <= v73 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v270 - v452
	v461 = v449
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v83 << (uint(int32(1)) % 32)
	v90 = F_repalloc(m, v82, v83<<(uint(int32(3))%32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v93 = v82
	goto L20
L20:
	;
	v95 = v73 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v93+v95))) = v69
	v98 = int32(0)
	if base.Ui32(v63) <= base.Ui32(v70) {
		v268 = v70
		v269 = v70
		v270 = v69
		v274 = v98
		v278 = v98
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v90
	v93 = v90
	goto L20
L22:
	;
	v281 = v268 - v70
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v281 != v282 {
		goto L74
	} else {
		goto L75
	}
L23:
	;
	v103 = v70
	v105 = v69
	v109 = v98
	goto L24
L24:
	;
	v117 = v103 + int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v119 = base.B2i32(v118 == v45&int32(255))
	if v118 == v45&int32(255) {
		v268 = v103
		v269 = v117
		v270 = v105
		v274 = v109
		v278 = v119
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v268 = v258
	v269 = v258
	v270 = v264
	v274 = v260
	v278 = v119
	goto L22
L26:
	;
	if v118 != int32(92) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v257)
	v264 = v105 + int32(1)
	if base.Ui32(v258) < base.Ui32(v63) {
		v103 = v258
		v105 = v264
		v109 = v260
		goto L24
	} else {
		goto L72
	}
L28:
	;
	v257 = v118
	v258 = v117
	v260 = v109
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(v63) <= base.Ui32(v117) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v268 = v103
	v269 = v117
	v270 = v105
	v274 = v109
	v278 = v119
	goto L22
L32:
	;
	goto L33
L33:
	;
	v124 = v103 + int32(2)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v127 = v125 - int32(48)
	switch v127 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		goto L41
	default:
		v257 = v125
		v258 = v124
		v260 = v109
		goto L27
	case 50:
		goto L39
	case 54:
		goto L38
	case 62:
		goto L37
	case 66:
		goto L36
	case 68:
		goto L35
	case 70:
		goto L34
	case 72:
		goto L40
	}
L34:
	;
	v257 = int32(11)
	v258 = v124
	v260 = v109
	goto L27
L35:
	;
	v257 = int32(9)
	v258 = v124
	v260 = v109
	goto L27
L36:
	;
	v257 = int32(13)
	v258 = v124
	v260 = v109
	goto L27
L37:
	;
	v257 = int32(10)
	v258 = v124
	v260 = v109
	goto L27
L38:
	;
	v257 = int32(12)
	v258 = v124
	v260 = v109
	goto L27
L39:
	;
	v257 = int32(8)
	v258 = v124
	v260 = v109
	goto L27
L40:
	;
	v167 = int32(120)
	if base.Ui32(v63) <= base.Ui32(v124) {
		v257 = v167
		v258 = v124
		v260 = v109
		goto L27
	} else {
		goto L51
	}
L41:
	;
	if base.Ui32(v63) <= base.Ui32(v124) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v257 = v154
	v258 = v155
	v260 = base.B2i32(v154&int32(255) == int32(0)) | int32(base.Ui32(v154&int32(128))>>(uint(int32(7))%32)) | v109
	goto L27
L43:
	;
	v154 = v127
	v155 = v124
	goto L42
L44:
	;
	goto L45
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129&int32(248) != int32(48) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v127
	v155 = v124
	goto L42
L47:
	;
	goto L48
L48:
	;
	v134 = int32(3)
	v138 = v129 + v127<<(uint(v134)%32) - int32(48)
	v140 = v103 + v134
	if base.Ui32(v63) <= base.Ui32(v140) {
		v154 = v138
		v155 = v140
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v142&int32(248) != int32(48) {
		v154 = v138
		v155 = v140
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v154 = v142 + v138<<(uint(int32(3))%32) - int32(48)
	v155 = v103 + int32(4)
	goto L42
L51:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	goto L52
L52:
	;
	if base.B2i32(base.Ui32(v169-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v169|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v257 = v167
		v258 = v124
		v260 = v109
		goto L27
	} else {
		goto L53
	}
L53:
	;
	v184 = v169 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v184&int32(255)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.Ui32(v169-int32(65)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v198 = v184
	goto L56
L56:
	;
	v200 = v103 + int32(3)
	if base.Ui32(v63) <= base.Ui32(v200) {
		v237 = v198
		v238 = v200
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v198 = v195 - int32(87)
	goto L56
L58:
	;
	v195 = v169 | int32(32)
	goto L60
L59:
	;
	v195 = v169
	goto L60
L60:
	;
	goto L57
L61:
	;
	v257 = v237
	v258 = v238
	v260 = base.B2i32(v237&int32(255) == int32(0)) | int32(base.Ui32(v237&int32(128))>>(uint(int32(7))%32)) | v109
	goto L27
L62:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	goto L63
L63:
	;
	if base.B2i32(base.Ui32(v202-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v202|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v237 = v198
		v238 = v200
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v219 = v202 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v219&int32(255)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v202-int32(65)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v233 = v219
	goto L67
L67:
	;
	v237 = v233 + v198<<(uint(int32(4))%32)
	v238 = v103 + int32(4)
	goto L61
L68:
	;
	v233 = v230 - int32(87)
	goto L67
L69:
	;
	v230 = v202 | int32(32)
	goto L71
L70:
	;
	v230 = v202
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L25
L73:
	;
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v446)
	v448 = int32(1)
	v449 = v73 + v448
	if v278 != 0 {
		v69 = v270 + v448
		v70 = v269
		v73 = v449
		goto L16
	} else {
		goto L123
	}
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v333 != 0 {
		goto L92
	} else {
		goto L93
	}
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v281 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v328 != 0 {
		goto L74
	} else {
		goto L90
	}
L77:
	;
	v328 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v290 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v291 = v70
	v292 = v284
	v293 = v281
	v294 = v290
	goto L84
L81:
	;
	v316 = v284
	v320 = int32(0)
	goto L82
L82:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v328 = v320 - v321
	goto L76
L83:
	;
	v316 = v311
	v320 = v313
	goto L82
L84:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v294 != v296 {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v311 = v305
	v313 = int32(0)
	goto L83
L86:
	;
	if v296 == int32(0) {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v301 = v293 - int32(1)
	if v301 == int32(0) {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(1)
	v305 = v292 + v304
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v306 != 0 {
		v291 = v291 + v304
		v292 = v305
		v293 = v301
		v294 = v306
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v329+v95))) = int32(0)
	goto L73
L91:
	;
	if v274&int32(1) == int32(0) {
		goto L73
	} else {
		goto L121
	}
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v336 = v334
	goto L94
L93:
	;
	v336 = int32(0)
	goto L94
L94:
	;
	if v336 <= v73 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v338 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v281 != v341 {
		goto L91
	} else {
		goto L97
	}
L97:
	;
	if v281 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v386 != 0 {
		goto L91
	} else {
		goto L112
	}
L99:
	;
	v386 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v348 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v349 = v70
	v350 = v338
	v351 = v281
	v352 = v348
	goto L106
L103:
	;
	v374 = v338
	v378 = int32(0)
	goto L104
L104:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	v386 = v378 - v379
	goto L98
L105:
	;
	v374 = v369
	v378 = v371
	goto L104
L106:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v352 != v354 {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L108
	}
L107:
	;
	v369 = v363
	v371 = int32(0)
	goto L105
L108:
	;
	if v354 == int32(0) {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v359 = v351 - int32(1)
	if v359 == int32(0) {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v362 = int32(1)
	v363 = v350 + v362
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	if v364 != 0 {
		v349 = v349 + v362
		v350 = v363
		v351 = v359
		v352 = v364
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v388+v95)))
	v392 = v390 - int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v387+v392<<(uint(int32(2))%32))))
	if v396 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397+v392))) = uint8(v399)
	goto L73
L114:
	;
	goto L115
L115:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+52))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(505237), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v402 + v403<<(uint(int32(4))%32) + v392*int32(100) + int32(24)
	F_errdetail(m, int32(628676), v18)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(498840), int32(1775), int32(63865))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437+v95)))
	F_pg_verifymbstr(m, v439, v270-v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L73
L123:
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
		return int32(1640320)
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
		v107 = int32(1640300)
		return v107
	case 4:
		return int32(1640340)
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
		return int32(1640280)
	}
}
func F_CreateRestartPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v269 float64
	_ = v269
	var v271 int32
	_ = v271
	var v273 float64
	_ = v273
	var v280 float64
	_ = v280
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int64
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v354 int32
	_ = v354
	var v361 float64
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int64
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	v12 = m.G0
	v14 = v12 - int32(1200)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+440)) = int32(1)
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	F_s_lock(m, v22+int32(440), int32(498262), int32(7650), int32(89079))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+352))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v33)+344))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v33)+336))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)+328))
	goto L7
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+440)) = int32(0)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v48 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v43 = F__emscripten_memcpy_bulkmem(m, v14+int32(84), v33+int32(356), int32(76))
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	m.G0 = v14 + int32(1200)
	return v458
L11:
	;
	if v37 != int64(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
	v54 = base.B2i32(v52 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v54)
	if v52 != int32(2) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v57 = int32(0)
	v60 = F_errstart(m, int32(13), v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v60 == int32(0) {
		v458 = v57
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errmsg_internal(m, int32(461704), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(498262), int32(7663), int32(89079))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v458 = v57
	goto L10
L20:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L36
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
	if base.Ui64(v77) < base.Ui64(v35) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = int32(0)
	v82 = F_errstart(m, int32(13), v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v35)
	v86 = int64(base.Ui64(v35) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v86)
	F_errmsg_internal(m, int32(513344), v14)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	F_errfinish(m, int32(498262), int32(7686), int32(89079))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if l0&int32(1) == int32(0) {
		v458 = v79
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v109 = F_LWLockAcquire(m, v105+int32(1152), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+16)) = int32(2)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	F_update_controlfile(m, v116, v112)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v120+int32(1152))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v458 = v79
	goto L10
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+152)) = v35
	*(*int64)(unsafe.Add(mBase, _consts[115])) = v35
	F_WALInsertLockRelease(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+440)) = int32(1)
	if v136 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	F_s_lock(m, v140+int32(440), int32(498262), int32(7713), int32(89079))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v150 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+440)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v149)+200)) = v35
	v157 = F__emscripten_memset_bulkmem(m, int32(4411160), base.I32_extend8_s(v150), int32(80))
	mBase = m.M
	goto L42
L41:
	;
	goto L40
L42:
	;
	v162 = m.G0
	v163 = int32(16)
	v164 = v162 - v163
	m.G0 = v164
	F___gettimeofday(m, v164)
	mBase = m.M
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	m.G0 = v164 + v163
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, _consts[132])) = v168 + v167*int64(1000000) - int64(946684800000000)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[133])))
	if v179 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_LogCheckpointStart(m, l0, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l0&int32(3) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = int32(88041)
	if l0&int32(2) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_CheckPointGuts(m, v35, l0)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L58
	}
L51:
	;
	v193 = int32(731062)
	goto L53
L52:
	;
	v193 = int32(757461)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v193
	if l0&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v199 = int32(739343)
	goto L56
L55:
	;
	v199 = int32(757461)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v199
	v207 = F_pg_snprintf(m, v14+int32(160), int32(128), int32(175654), v14+int32(48))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v211 = F_strlen(m, v14+int32(160))
	mBase = m.M
	goto L50
L58:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v215)+40))
	v218 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v222 = F_LWLockAcquire(m, v218+int32(1152), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v225)+40))
	if base.Ui64(v226) < base.Ui64(v35) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+48)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v225)+40)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v225)+32)) = v37
	goto L64
L61:
	;
	goto L62
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v258+int32(1152))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L74
	}
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	if v238 != int32(5) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v236 = F__emscripten_memcpy_bulkmem(m, v225+int32(52), v14+int32(84), int32(76))
	mBase = m.M
	goto L66
L66:
	;
	goto L63
L67:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	F_update_controlfile(m, v254, v225)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L73
	}
L68:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v225)+136))
	if base.Ui64(v241) < base.Ui64(v36) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+144)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v225)+136)) = v36
	*(*int64)(unsafe.Add(mBase, _consts[134])) = v36
	goto L71
L70:
	;
	goto L71
L71:
	;
	if l0&int32(1) == int32(0) {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = int32(2)
	goto L67
L73:
	;
	goto L62
L74:
	;
	v264 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	if v216 != int64(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v269 = base.F64_convert_i64_u(v264 - v216)
	*(*float64)(unsafe.Add(mBase, _consts[136])) = v269
	v271 = int32(4411272)
	v273 = *(*float64)(unsafe.Add(mBase, _consts[137]))
	if base.F64_gt(v269, v273) != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v285 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v286 = base.I64_div_u_s(v264, v285)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v286
	v288 = int32(0)
	v290 = F_GetWalRcvFlushRecPtr(m, v288, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L81
	}
L78:
	;
	v280 = v269
	goto L80
L79:
	;
	v280 = base.F64_add(base.F64_mul(v273, float64(0.9)), base.F64_mul(v269, float64(0.1)))
	goto L80
L80:
	;
	*(*float64)(unsafe.Add(mBase, _consts[137])) = v280
	goto L77
L81:
	;
	v294 = F_GetXLogReplayRecPtr(m, v14+int32(80))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	if base.Ui64(v294) < base.Ui64(v290) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v297 = v290
	goto L85
L84:
	;
	v297 = v294
	goto L85
L85:
	;
	F_KeepLogSeg(m, v297, v14+int32(72))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	v304 = int32(0)
	v306 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v303, v304, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v306 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v309 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	v311 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v312 = base.I64_div_u_s(v309, v311)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v312
	F_KeepLogSeg(m, v297, v14+int32(72))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v319 = v303
	goto L90
L90:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v325 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	v319 = v318
	goto L90
L92:
	;
	v342 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	F_RemoveOldXlogFiles(m, v319-int64(1), v342, v297, v339)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L96
	}
L93:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v323)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v337
	v339 = v337
	goto L92
L94:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323)+316))
	v330 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(base.B2i32(v329 != v330))
	if v329 == v330 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v339 = v335
	goto L92
L96:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+320)))
	if v347 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, _consts[138])))
	if v400 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L98:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v352 = v297 - int64(1)
	v354 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v361 = base.F64_mul(base.F64_convert_i32_s(v354), float64(0.75))
	if base.F64_lt(v361, float64(4.294967296e+09))&base.F64_ge(v361, float64(0)) != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui64(v352&base.I64_extend_i32_s(v354-int32(1))) < base.Ui64(base.I64_extend_i32_u(v369)) {
		goto L97
	} else {
		goto L103
	}
L100:
	;
	v367 = base.I32_trunc_f64_u(v361)
	v369 = v367
	goto L99
L101:
	;
	goto L102
L102:
	;
	v369 = int32(0)
	goto L99
L103:
	;
	v373 = base.I64_div_u_s(v352, base.I64_extend_i32_s(v354))
	v380 = F_XLogFileInitInternal(m, v373+int64(1), v350, v14+int32(1199), v14+int32(160))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v380 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v384 = F_close(m, v380)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1199)))
	if v385 != int32(1) {
		goto L97
	} else {
		goto L108
	}
L108:
	;
	v388 = int32(4411208)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v390 + int32(1)
	goto L97
L109:
	;
	v403 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_LogCheckpointEnd(m, int32(1))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	F_TruncateSUBTRANS(m, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v410 = F_GetLatestXTime(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, _consts[133])))
	if v415 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v416 = int32(15)
	goto L118
L117:
	;
	v416 = int32(13)
	goto L118
L118:
	;
	v418 = F_errstart(m, v416, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	if v418 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+36)) = uint32(v35)
	v422 = int64(base.Ui64(v35) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+32)) = uint32(v422)
	F_errmsg(m, int32(512113), v14+int32(32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v444 = int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v446 == int32(0) {
		v458 = v444
		goto L10
	} else {
		goto L130
	}
L123:
	;
	if v410 != int64(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v431 = F_timestamptz_to_str(m, v410)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_errfinish(m, int32(498262), int32(7876), int32(89079))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v431
	F_errdetail(m, int32(606106), v14+int32(16))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	goto L122
L130:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v449 == int32(0) {
		v458 = v444
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_ExecuteRecoveryCommand(m, v446, int32(427724), int32(0), int32(134217729))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v458 = v444
	goto L10
}
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1331 int32
	_ = v1331
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1556 int32
	_ = v1556
	var v1570 int32
	_ = v1570
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
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
	var v1656 int32
	_ = v1656
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
	var v1744 int32
	_ = v1744
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1767 int32
	_ = v1767
	var v1795 int32
	_ = v1795
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1840 int64
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(432)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v27 == v4 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L25
	} else {
		goto L441
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L25
	} else {
		goto L437
	}
L3:
	;
	v1840 = *(*int64)(unsafe.Add(mBase, uint32(v1824)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1840
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1842
	m.G0 = v23 + int32(432)
	return
L4:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1097 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L5:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1047 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L262
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L25
	} else {
		goto L258
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = v4
	goto L22
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L25
	} else {
		goto L253
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L25
	} else {
		goto L249
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L25
	} else {
		goto L243
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L25
	} else {
		goto L239
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L25
	} else {
		goto L235
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L25
	} else {
		goto L230
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L25
	} else {
		goto L226
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L25
	} else {
		goto L222
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L25
	} else {
		goto L218
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L25
	} else {
		goto L214
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L25
	} else {
		goto L210
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L25
	} else {
		goto L205
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L25
	} else {
		goto L201
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v41<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != int32(3) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v127 != 0 {
		goto L51
	} else {
		goto L52
	}
L24:
	;
	v66 = F_relation_openrv(m, v61, int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	v71 = v69 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v71) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if int32(1)<<(uint(v71)%32)&int32(5249) == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v82 = F_object_ownercheck(m, int32(1259), v81, v26)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v82 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v87)+119)))
	switch v88 - int32(73) {
	case 0, 32:
		goto L39
	default:
		v98 = int32(41)
		goto L34
	case 10:
		goto L38
	case 29:
		goto L35
	case 36:
		goto L36
	case 45:
		goto L37
	}
L31:
	;
	goto L32
L32:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[222])))
	if v107 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	F_aclcheck_error(m, int32(2), v100, v101+int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L25
	} else {
		goto L40
	}
L34:
	;
	v100 = v98
	goto L33
L35:
	;
	v98 = int32(18)
	goto L34
L36:
	;
	v100 = int32(23)
	goto L33
L37:
	;
	v100 = int32(51)
	goto L33
L38:
	;
	v100 = int32(37)
	goto L33
L39:
	;
	v100 = int32(20)
	goto L33
L40:
	;
	goto L32
L41:
	;
	v111 = int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	if base.Ui32(v112) < base.Ui32(int32(12000)) {
		v121 = v111
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v123 = v41 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v123 < v124 {
		v41 = v123
		goto L22
	} else {
		goto L49
	}
L44:
	;
	if v121 != 0 {
		goto L19
	} else {
		goto L48
	}
L45:
	;
	goto L44
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
	if v116 == int32(99) {
		v121 = v111
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v119 = F_isTempToastNamespace(m, v116)
	mBase = m.M
	v121 = v119
	goto L45
L48:
	;
	goto L43
L49:
	;
	goto L23
L50:
	;
	v445 = F_strncpy(m, v23+int32(304), v426, int32(64))
	mBase = m.M
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v445)+63)) = uint8(v446)
	goto L116
L51:
	;
	v130 = F_QualifiedNameGetCreationNamespace(m, v127, v23+int32(284))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L25
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+68))
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+304)) = uint8(v136)
	v140 = v134 + int32(4)
	if v133 == v136 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	v426 = v132
	v434 = v130
	goto L50
L55:
	;
	v339 = F_pstrdup(m, v23+int32(304))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L25
	} else {
		goto L103
	}
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v143 <= int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v151 = v136
	v152 = int32(0)
	v155 = v143
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v151<<(uint(int32(2))%32))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v172 == int32(206) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L55
L60:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if int32(0) < v152 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v311 = v152
	v312 = v155
	goto L62
L62:
	;
	v315 = v151 + int32(1)
	if v315 < v312 {
		v151 = v315
		v152 = v311
		v155 = v312
		goto L58
	} else {
		goto L102
	}
L63:
	;
	v181 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(304)+v152))) = uint8(v181)
	v185 = v152 + int32(1)
	goto L65
L64:
	;
	v185 = v152
	goto L65
L65:
	;
	v188 = v23 + int32(304) + v185
	if v175 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v190 = v175
	goto L68
L67:
	;
	v190 = int32(207114)
	goto L68
L68:
	;
	goto L72
L69:
	;
	v306 = F_strlen(m, v188)
	mBase = m.M
	v307 = v306 + v185
	if int32(63) < v307 {
		goto L55
	} else {
		goto L101
	}
L70:
	;
	v303 = F_strlen(m, v292)
	mBase = m.M
	goto L69
L72:
	;
	goto L73
L73:
	;
	v197 = int32(63)
	if (v188^v190)&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v293))) = uint8(v296)
	goto L70
L75:
	;
	v277 = v272
	v278 = v273
	v279 = v274
	goto L97
L76:
	;
	if v267 == int32(0) {
		v292 = v265
		v293 = v266
		goto L74
	} else {
		goto L96
	}
L77:
	;
	v265 = v190
	v266 = v188
	v267 = v197
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v190&int32(3) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v234 == int32(0) {
		v292 = v231
		v293 = v232
		goto L74
	} else {
		goto L89
	}
L81:
	;
	v231 = v190
	v232 = v188
	v233 = v197
	v234 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v210 = v190
	v211 = v188
	v212 = v197
	goto L84
L84:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v214)
	if v214 == int32(0) {
		v272 = v210
		v273 = v211
		v274 = v212
		goto L75
	} else {
		goto L86
	}
L85:
	;
	v231 = v225
	v232 = v219
	v233 = v221
	v234 = v223
	goto L80
L86:
	;
	v218 = int32(1)
	v219 = v211 + v218
	v221 = v212 - v218
	v222 = int32(0)
	v223 = base.B2i32(v221 != v222)
	v225 = v210 + v218
	if v225&int32(3) == v222 {
		v231 = v225
		v232 = v219
		v233 = v221
		v234 = v223
		goto L80
	} else {
		goto L87
	}
L87:
	;
	if v221 != 0 {
		v210 = v225
		v211 = v219
		v212 = v221
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v237 == int32(0) {
		v265 = v231
		v266 = v232
		v267 = v233
		goto L76
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v233) < base.Ui32(int32(4)) {
		v265 = v231
		v266 = v232
		v267 = v233
		goto L76
	} else {
		goto L91
	}
L91:
	;
	v243 = v231
	v244 = v232
	v245 = v233
	goto L92
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v251 = int32(-2139062144)
	if (int32(16843008)-v248|v248)&v251 != v251 {
		v272 = v243
		v273 = v244
		v274 = v245
		goto L75
	} else {
		goto L94
	}
L93:
	;
	v265 = v259
	v266 = v257
	v267 = v261
	goto L76
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v248
	v256 = int32(4)
	v257 = v244 + v256
	v259 = v243 + v256
	v261 = v245 - v256
	if base.Ui32(int32(3)) < base.Ui32(v261) {
		v243 = v259
		v244 = v257
		v245 = v261
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v272 = v265
	v273 = v266
	v274 = v267
	goto L75
L97:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	*(*uint8)(unsafe.Add(mBase, uint32(v278))) = uint8(v281)
	if v281 == int32(0) {
		v292 = v277
		v293 = v278
		goto L74
	} else {
		goto L99
	}
L98:
	;
	v292 = v288
	v293 = v286
	goto L74
L99:
	;
	v285 = int32(1)
	v286 = v278 + v285
	v288 = v277 + v285
	v290 = v279 - v285
	if v290 != 0 {
		v277 = v288
		v278 = v286
		v279 = v290
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v311 = v307
	v312 = v310
	goto L62
L102:
	;
	goto L59
L103:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+308)) = uint8(v343)
	v346 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+304)) = v346
	v351 = F_makeObjectName(m, v140, v339, v23+int32(304))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L25
	} else {
		goto L104
	}
L104:
	;
	v353 = int32(0)
	v355 = F_GetSysCacheOid(m, int32(63), v351, v135, v353, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L25
	} else {
		goto L105
	}
L105:
	;
	if v355 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v361 = v351
	v362 = int32(0)
	goto L109
L107:
	;
	v405 = v351
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+284)) = v405
	v426 = v405
	v434 = v135
	goto L50
L109:
	;
	F_pfree(m, v361)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L111
	}
L110:
	;
	v405 = v395
	goto L108
L111:
	;
	v380 = v362 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = int32(112015)
	v390 = F_pg_snprintf(m, v23+int32(304), int32(64), int32(465989), v23+int32(144))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	v395 = F_makeObjectName(m, v140, v339, v23+int32(304))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	v397 = int32(0)
	v399 = F_GetSysCacheOid(m, int32(63), v395, v135, v397, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L25
	} else {
		goto L114
	}
L114:
	;
	if v399 != 0 {
		v361 = v395
		v362 = v380
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	if l2 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v464 = int32(0)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	v469 = F_SearchSysCacheExists(m, int32(63), v466, v434, v464, v464)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L25
	} else {
		goto L123
	}
L118:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v454 = F_object_aclcheck(m, int32(2615), v434, v452, int64(512))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L25
	} else {
		goto L119
	}
L119:
	;
	if v454 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v459 = F_get_namespace_name(m, v434)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L25
	} else {
		goto L121
	}
L121:
	;
	F_aclcheck_error(m, v454, int32(36), v459)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L25
	} else {
		goto L122
	}
L122:
	;
	goto L117
L123:
	;
	if v469 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+25)))
	if v471 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v516 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L127:
	;
	v476 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L25
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L25
	} else {
		goto L138
	}
L130:
	;
	if v476 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L25
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L25
	} else {
		goto L137
	}
L134:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v481
	F_errmsg(m, int32(333207), v23+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L25
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(494396), int32(209), int32(174043))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L25
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	v1824 = int32(766152)
	goto L3
L138:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L25
	} else {
		goto L139
	}
L139:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v504
	F_errmsg(m, int32(116325), v23+int32(32))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L25
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(494396), int32(216), int32(174043))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L25
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
	v1082 = int32(0)
	v1086 = v464
	v1096 = v4
	goto L4
L143:
	;
	goto L144
L144:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if int32(9) <= v520 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	v523 = int32(0)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v524 <= v523 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v1032 = int32(0)
	v1036 = v464
	goto L5
L147:
	;
	goto L148
L148:
	;
	v534 = int32(0)
	v536 = v523
	v538 = v464
	goto L149
L149:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v536<<(uint(int32(2))%32))))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	if v554 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v1032 = v747
	v1036 = v751
	goto L5
L151:
	;
	v763 = v536 + int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v763 < v764 {
		v534 = v747
		v536 = v763
		v538 = v751
		goto L149
	} else {
		goto L200
	}
L152:
	;
	v555 = F_SearchSysCacheAttName(m, v126, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L25
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	if v587 == int32(6) {
		goto L162
	} else {
		goto L163
	}
L155:
	;
	if v555 == int32(0) {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555)+16))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+22)))
	v561 = v559 + v560
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v561)+74)))
	if v562 <= int32(0) {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+90)))
	if v565 == int32(118) {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v561)+68))
	v570 = F_lookup_type_cache(m, v568, int32(2))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L25
	} else {
		goto L159
	}
L159:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v570)+56))
	if v572 == int32(0) {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	v580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v561)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v538<<(uint(int32(1))%32)))) = uint16(v580)
	F_ReleaseCatCache(m, v555)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L25
	} else {
		goto L161
	}
L161:
	;
	v747 = v534
	v751 = v538 + int32(1)
	goto L151
L162:
	;
	v590 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586)+8)))
	if v590 <= int32(0) {
		goto L13
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = int32(0)
	F_pull_varattnos(m, v586, int32(1), v23+int32(240))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L25
	} else {
		goto L170
	}
L165:
	;
	v593 = F_get_attgenerated(m, v126, v590)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L25
	} else {
		goto L166
	}
L166:
	;
	if v593 == int32(118) {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v599 = F_lookup_type_cache(m, v597, int32(2))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L25
	} else {
		goto L168
	}
L168:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v599)+56))
	if v601 == int32(0) {
		goto L11
	} else {
		goto L169
	}
L169:
	;
	v606 = int32(1)
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v586)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v538<<(uint(v606)%32)))) = uint16(v609)
	v747 = v534
	v751 = v538 + v606
	goto L151
L170:
	;
	v625 = int32(-1)
	goto L172
L171:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v725 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L172:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v23)+240))
	if v641 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L25
	} else {
		goto L189
	}
L174:
	;
	if v697 < int32(0) {
		goto L171
	} else {
		goto L185
	}
L175:
	;
	v697 = base.I32_ctz(v683) | v684<<(uint(int32(5))%32)
	goto L174
L176:
	;
	v697 = int32(-2)
	goto L174
L177:
	;
	v648 = v625 + int32(1)
	v650 = base.I32_div_s(v648, int32(32))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v651 <= v650 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v654 = v641 + int32(8)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v650<<(uint(int32(2))%32))))
	v661 = v658 & (int32(-1) << (uint(v648) % 32))
	if v661 != 0 {
		v683 = v661
		v684 = v650
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v663 = v650 + int32(1)
	if v663 == v651 {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v666 = v663
	goto L181
L181:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v654+v666<<(uint(int32(2))%32))))
	if v673 != 0 {
		v683 = v673
		v684 = v666
		goto L175
	} else {
		goto L183
	}
L182:
	;
	goto L176
L183:
	;
	v675 = v666 + int32(1)
	if v675 != v651 {
		v666 = v675
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v702 = base.I32_extend16_s(v697 - int32(7))
	if v702 <= int32(0) {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	v705 = F_get_attgenerated(m, v126, v702)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L25
	} else {
		goto L187
	}
L187:
	;
	if v705 != int32(118) {
		v625 = v697
		goto L172
	} else {
		goto L188
	}
L188:
	;
	goto L173
L189:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L25
	} else {
		goto L190
	}
L190:
	;
	F_errmsg(m, int32(442288), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L25
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(494396), int32(343), int32(174043))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L25
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v740 = F_lappend(m, v534, v586)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L25
	} else {
		goto L199
	}
L194:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	if v728 < int32(2) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v731 = F_exprType(m, v586)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L25
	} else {
		goto L196
	}
L196:
	;
	v734 = F_lookup_type_cache(m, v731, int32(2))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L25
	} else {
		goto L197
	}
L197:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v734)+56))
	if v736 == int32(0) {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	goto L193
L199:
	;
	v747 = v740
	v751 = v538
	goto L151
L200:
	;
	goto L150
L201:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L25
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(524155), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L25
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(494396), int32(115), int32(174043))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L25
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L25
	} else {
		goto L206
	}
L206:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v789 + int32(4)
	F_errmsg(m, int32(704705), v23)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L25
	} else {
		goto L207
	}
L207:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v797 = int32(*(*int8)(unsafe.Add(mBase, uint32(v796)+119)))
	F_errdetail_relkind_not_supported(m, v797)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L25
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(494396), int32(135), int32(174043))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L25
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L25
	} else {
		goto L211
	}
L211:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v812 + int32(4)
	F_errmsg(m, int32(327401), v23+int32(160))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L25
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(494396), int32(153), int32(174043))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L25
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L25
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(8)
	F_errmsg(m, int32(173913), v23-int32(-64))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L25
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(494396), int32(228), int32(174043))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L25
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L25
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v554
	F_errmsg(m, int32(71934), v23+int32(112))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L25
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(494396), int32(261), int32(174043))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L25
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L25
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(442233), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L25
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(494396), int32(268), int32(174043))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L25
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L25
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(442288), int32(0))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L25
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(494396), int32(274), int32(174043))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L25
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L25
	} else {
		goto L231
	}
L231:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v561)+68))
	v903 = F_format_type_be(m, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L25
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v554
	F_errmsg(m, int32(130741), v23+int32(128))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L25
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(494396), int32(282), int32(174043))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L25
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L25
	} else {
		goto L236
	}
L236:
	;
	F_errmsg(m, int32(442233), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L25
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(494396), int32(297), int32(174043))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L25
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L25
	} else {
		goto L240
	}
L240:
	;
	F_errmsg(m, int32(442288), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L25
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(494396), int32(303), int32(174043))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L25
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L25
	} else {
		goto L244
	}
L244:
	;
	v956 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586)+8)))
	v958 = F_get_attname(m, v126, v956, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L25
	} else {
		goto L245
	}
L245:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v961 = F_format_type_be(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L25
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v961
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v958
	F_errmsg(m, int32(130741), v23+int32(80))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L25
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(494396), int32(311), int32(174043))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L25
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L25
	} else {
		goto L250
	}
L250:
	;
	F_errmsg(m, int32(442233), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L25
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(494396), int32(337), int32(174043))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L25
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L25
	} else {
		goto L254
	}
L254:
	;
	v998 = F_format_type_be(m, v731)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L25
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v998
	F_errmsg(m, int32(130838), v23+int32(96))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L25
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(494396), int32(361), int32(174043))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L25
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L25
	} else {
		goto L259
	}
L259:
	;
	F_errmsg(m, int32(524155), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L25
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(494396), int32(106), int32(174043))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L25
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1050 != int32(1) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L263
	}
L263:
	;
	if v1032 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+4))
	if v1055 != int32(1) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L265
	}
L265:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1058 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L266
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L25
	} else {
		goto L267
	}
L267:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L25
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(458597), int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L25
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(494396), int32(381), int32(174043))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L25
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	v1384 = base.B2i32(v1082 == int32(0))
	if v1096 < int32(2) {
		goto L353
	} else {
		goto L354
	}
L272:
	;
	v1110 = int32(0)
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1114 = int32(109191)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, _consts[299])))
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1118 == v1110 {
		v1137 = v1117
		v1138 = v1118
		goto L281
	} else {
		goto L282
	}
L273:
	;
	v1365 = v1107
	v1376 = v1109
	v1379 = v4
	v1380 = v4
	goto L271
L274:
	;
	v1107 = int32(1)
	v1109 = int32(0)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if int32(0) < v1103 {
		goto L272
	} else {
		goto L277
	}
L277:
	;
	v1107 = int32(1)
	v1109 = int32(0)
	goto L273
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L25
	} else {
		goto L348
	}
L279:
	;
	v1201 = base.B2i32(v1139 == int32(0))
	v1202 = int32(1)
	if v1103 == v1202 {
		v1365 = v1110
		v1376 = v1198
		v1379 = v1199
		v1380 = v1201
		goto L271
	} else {
		goto L309
	}
L280:
	;
	if v1139 == int32(0) {
		v1198 = v4
		v1199 = v4
		goto L279
	} else {
		goto L288
	}
L281:
	;
	v1139 = v1138 - v1137
	goto L280
L282:
	;
	if v1117 != v1118 {
		v1137 = v1117
		v1138 = v1118
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1122 = v1113
	v1123 = v1114
	goto L284
L284:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123)+1)))
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+1)))
	if v1127 == int32(0) {
		v1137 = v1126
		v1138 = v1127
		goto L281
	} else {
		goto L286
	}
L285:
	;
	v1137 = v1126
	v1138 = v1127
	goto L281
L286:
	;
	v1130 = int32(1)
	if v1126 == v1127 {
		v1122 = v1122 + v1130
		v1123 = v1123 + v1130
		goto L284
	} else {
		goto L287
	}
L287:
	;
	goto L285
L288:
	;
	v1142 = int32(168572)
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, _consts[300])))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1146 == int32(0) {
		v1165 = v1145
		v1166 = v1146
		goto L290
	} else {
		goto L291
	}
L289:
	;
	if v1166-v1165 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L290:
	;
	goto L289
L291:
	;
	if v1145 != v1146 {
		v1165 = v1145
		v1166 = v1146
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1150 = v1113
	v1151 = v1142
	goto L293
L293:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+1)))
	if v1155 == int32(0) {
		v1165 = v1154
		v1166 = v1155
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1165 = v1154
	v1166 = v1155
	goto L290
L295:
	;
	v1158 = int32(1)
	if v1154 == v1155 {
		v1150 = v1150 + v1158
		v1151 = v1151 + v1158
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v1198 = v4
	v1199 = int32(1)
	goto L279
L298:
	;
	goto L299
L299:
	;
	v1171 = int32(35809)
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[301])))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1175 == int32(0) {
		v1194 = v1174
		v1195 = v1175
		goto L301
	} else {
		goto L302
	}
L300:
	;
	if v1195-v1194 != 0 {
		v1331 = v1113
		goto L278
	} else {
		goto L308
	}
L301:
	;
	goto L300
L302:
	;
	if v1174 != v1175 {
		v1194 = v1174
		v1195 = v1175
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1179 = v1113
	v1180 = v1171
	goto L304
L304:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+1)))
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+1)))
	if v1184 == int32(0) {
		v1194 = v1183
		v1195 = v1184
		goto L301
	} else {
		goto L306
	}
L305:
	;
	v1194 = v1183
	v1195 = v1184
	goto L301
L306:
	;
	v1187 = int32(1)
	if v1183 == v1184 {
		v1179 = v1179 + v1187
		v1180 = v1180 + v1187
		goto L304
	} else {
		goto L307
	}
L307:
	;
	goto L305
L308:
	;
	v1198 = int32(1)
	v1199 = v4
	goto L279
L309:
	;
	v1205 = int32(0)
	if v1205 < v1103 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1208 = v1103
	goto L312
L311:
	;
	v1208 = v1205
	goto L312
L312:
	;
	v1213 = v1202
	v1222 = v1198
	v1225 = v1199
	v1226 = v1201
	goto L313
L313:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1111+v1213<<(uint(int32(2))%32))))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+4))
	v1234 = int32(109191)
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, _consts[299])))
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1238 == int32(0) {
		v1257 = v1237
		v1258 = v1238
		goto L317
	} else {
		goto L318
	}
L314:
	;
	v1365 = v1110
	v1376 = v1319
	v1379 = v1320
	v1380 = v1321
	goto L271
L315:
	;
	v1323 = v1213 + int32(1)
	if v1208 != v1323 {
		v1213 = v1323
		v1222 = v1319
		v1225 = v1320
		v1226 = v1321
		goto L313
	} else {
		goto L347
	}
L316:
	;
	if v1258-v1257 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L317:
	;
	goto L316
L318:
	;
	if v1237 != v1238 {
		v1257 = v1237
		v1258 = v1238
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1242 = v1233
	v1243 = v1234
	goto L320
L320:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+1)))
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
	if v1247 == int32(0) {
		v1257 = v1246
		v1258 = v1247
		goto L317
	} else {
		goto L322
	}
L321:
	;
	v1257 = v1246
	v1258 = v1247
	goto L317
L322:
	;
	v1250 = int32(1)
	if v1246 == v1247 {
		v1242 = v1242 + v1250
		v1243 = v1243 + v1250
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v1319 = v1222
	v1320 = v1225
	v1321 = int32(1)
	goto L315
L325:
	;
	goto L326
L326:
	;
	v1263 = int32(168572)
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, _consts[300])))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1267 == int32(0) {
		v1286 = v1266
		v1287 = v1267
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if v1287-v1286 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L328:
	;
	goto L327
L329:
	;
	if v1266 != v1267 {
		v1286 = v1266
		v1287 = v1267
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1271 = v1233
	v1272 = v1263
	goto L331
L331:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+1)))
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+1)))
	if v1276 == int32(0) {
		v1286 = v1275
		v1287 = v1276
		goto L328
	} else {
		goto L333
	}
L332:
	;
	v1286 = v1275
	v1287 = v1276
	goto L328
L333:
	;
	v1279 = int32(1)
	if v1275 == v1276 {
		v1271 = v1271 + v1279
		v1272 = v1272 + v1279
		goto L331
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1319 = v1222
	v1320 = int32(1)
	v1321 = v1226
	goto L315
L336:
	;
	goto L337
L337:
	;
	v1292 = int32(35809)
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, _consts[301])))
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1296 == int32(0) {
		v1315 = v1295
		v1316 = v1296
		goto L339
	} else {
		goto L340
	}
L338:
	;
	if v1316-v1315 != 0 {
		v1331 = v1233
		goto L278
	} else {
		goto L346
	}
L339:
	;
	goto L338
L340:
	;
	if v1295 != v1296 {
		v1315 = v1295
		v1316 = v1296
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v1300 = v1233
	v1301 = v1292
	goto L342
L342:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301)+1)))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+1)))
	if v1305 == int32(0) {
		v1315 = v1304
		v1316 = v1305
		goto L339
	} else {
		goto L344
	}
L343:
	;
	v1315 = v1304
	v1316 = v1305
	goto L339
L344:
	;
	v1308 = int32(1)
	if v1304 == v1305 {
		v1300 = v1300 + v1308
		v1301 = v1301 + v1308
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	v1319 = int32(1)
	v1320 = v1225
	v1321 = v1226
	goto L315
L347:
	;
	goto L314
L348:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L25
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1331
	F_errmsg(m, int32(722529), v23+int32(48))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L25
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(494396), int32(411), int32(174043))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L25
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_pg_qsort(m, v23+int32(288), v1086, int32(2), int32(571))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L25
	} else {
		goto L359
	}
L353:
	;
	if int32(1) < v1096 {
		v1400 = v1376
		v1401 = v1384
		v1402 = v1379
		v1403 = v1380
		goto L352
	} else {
		goto L356
	}
L354:
	;
	if v1365 == int32(0) {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1389 = int32(1)
	v1400 = v1389
	v1401 = v1384
	v1402 = v1389
	v1403 = v1389
	goto L352
L356:
	;
	if v1082 == int32(0) {
		goto L2
	} else {
		goto L357
	}
L357:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1397 != int32(1) {
		goto L2
	} else {
		goto L358
	}
L358:
	;
	v1400 = v1376
	v1401 = int32(0)
	v1402 = v1379
	v1403 = v1380
	goto L352
L359:
	;
	if int32(2) <= v1086 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+288)))
	v1418 = int32(1)
	v1420 = v1412
	goto L363
L361:
	;
	goto L362
L362:
	;
	if v1401 != 0 {
		goto L367
	} else {
		goto L368
	}
L363:
	;
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v1418<<(uint(int32(1))%32)))))
	if v1420&int32(65535) == v1441 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	goto L362
L365:
	;
	v1444 = v1418 + int32(1)
	if v1444 != v1086 {
		v1418 = v1444
		v1420 = v1441
		goto L363
	} else {
		goto L366
	}
L366:
	;
	goto L364
L367:
	;
	v1594 = F_buildint2vector(m, v23+int32(288), v1086)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L25
	} else {
		goto L384
	}
L368:
	;
	v1466 = int32(0)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1467 <= v1466 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1472 = v1466
	v1477 = v1467
	goto L370
L370:
	;
	v1490 = int32(0)
	if v1477 <= v1490 {
		v1556 = v1477
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L367
L372:
	;
	v1570 = v1472 + int32(1)
	if v1570 < v1556 {
		v1472 = v1570
		v1477 = v1556
		goto L370
	} else {
		goto L383
	}
L373:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1493+v1472<<(uint(int32(2))%32))))
	v1503 = v1490
	v1505 = int32(0)
	goto L374
L374:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1519+v1503<<(uint(int32(2))%32))))
	v1524 = F_equal(m, v1497, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L25
	} else {
		goto L376
	}
L375:
	;
	if v1526 <= int32(1) {
		v1556 = v1529
		goto L372
	} else {
		goto L378
	}
L376:
	;
	v1526 = v1524 + v1505
	v1528 = v1503 + int32(1)
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1528 < v1529 {
		v1503 = v1528
		v1505 = v1526
		goto L374
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L25
	} else {
		goto L379
	}
L379:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L25
	} else {
		goto L380
	}
L380:
	;
	F_errmsg(m, int32(250313), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L25
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(494396), int32(492), int32(174043))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L25
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	goto L371
L384:
	;
	if v1403 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	if v1402 != 0 {
		goto L389
	} else {
		goto L390
	}
L386:
	;
	v1604 = v23 + int32(176)
	v1605 = int32(0)
	goto L385
L387:
	;
	goto L388
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = int32(100)
	v1604 = v23 + int32(176) | int32(4)
	v1605 = int32(1)
	goto L385
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1604))) = int32(102)
	v1610 = v1605 + int32(1)
	goto L391
L390:
	;
	v1610 = v1605
	goto L391
L391:
	;
	if v1400 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(176)|v1610<<(uint(int32(2))%32)))) = int32(109)
	v1620 = v1610 + int32(1)
	goto L394
L393:
	;
	v1620 = v1610
	goto L394
L394:
	;
	if v1401 != 0 {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	v1652 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L25
	} else {
		goto L404
	}
L396:
	;
	v1625 = F_construct_array_builtin(m, v23+int32(176), v1620, int32(18))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L25
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1628 = v23 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v1628+v1620<<(uint(int32(2))%32)))) = int32(101)
	v1639 = F_construct_array_builtin(m, v1628, v1620+int32(1), int32(18))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L25
	} else {
		goto L400
	}
L399:
	;
	v1647 = v1625
	v1649 = int32(0)
	goto L395
L400:
	;
	v1641 = F_nodeToString(m, v1082)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L25
	} else {
		goto L401
	}
L401:
	;
	v1643 = F_cstring_to_text(m, v1641)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L25
	} else {
		goto L402
	}
L402:
	;
	F_pfree(m, v1641)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L25
	} else {
		goto L403
	}
L403:
	;
	v1647 = v1639
	v1649 = v1643
	goto L395
L404:
	;
	v1654 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+264)) = v1654
	v1656 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+232)) = uint8(v1656)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+256)) = v1654
	*(*int64)(unsafe.Add(mBase, uint32(v23)+224)) = v1654
	v1664 = F_GetNewOidWithIndex(m, v1652, int32(3380), int32(1))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L25
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+260)) = v1594
	*(*int32)(unsafe.Add(mBase, uint32(v23)+256)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+252)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v23)+244)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = v1664
	v1671 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+230)) = uint8(v1671)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+272)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(v23)+268)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(v23)+248)) = v23 + int32(304)
	if v1649 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1680 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+232)) = uint8(v1680)
	goto L408
L407:
	;
	goto L408
L408:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+52))
	v1687 = F_heap_form_tuple(m, v1682, v23+int32(240), v23+int32(224))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L25
	} else {
		goto L409
	}
L409:
	;
	F_CatalogTupleInsert(m, v1652, v1687)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L25
	} else {
		goto L410
	}
L410:
	;
	F_pfree(m, v1687)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L25
	} else {
		goto L411
	}
L411:
	;
	F_relation_close(m, v1652, int32(3))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L25
	} else {
		goto L412
	}
L412:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v1697 != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1699 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3381), v1664, v1699, v1699)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L25
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	F_CacheInvalidateRelcache(m, v66)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L25
	} else {
		goto L417
	}
L416:
	;
	goto L415
L417:
	;
	v1705 = int32(0)
	F_relation_close(m, v66, v1705)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L25
	} else {
		goto L418
	}
L418:
	;
	v1709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = int32(3381)
	if v1709 < v1086 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if v1401 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L420:
	;
	v1720 = v1705
	goto L423
L421:
	;
	goto L422
L422:
	;
	if v1086 != 0 {
		goto L419
	} else {
		goto L427
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(1259)
	v1744 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+int32(288)+v1720<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = v1744
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(97))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L25
	} else {
		goto L425
	}
L425:
	;
	v1754 = v1720 + int32(1)
	if v1754 != v1086 {
		v1720 = v1754
		goto L423
	} else {
		goto L426
	}
L426:
	;
	goto L419
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(1259)
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(97))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L25
	} else {
		goto L428
	}
L428:
	;
	goto L419
L429:
	;
	F_recordDependencyOnSingleRelExpr(m, v23+int32(200), v1082, v126, int32(97), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L25
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(2615)
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(110))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L25
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	F_recordDependencyOnOwner(m, int32(3381), v1664, v26)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L25
	} else {
		goto L434
	}
L434:
	;
	v1812 = v23 + int32(200)
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1813 == int32(0) {
		v1824 = v1812
		goto L3
	} else {
		goto L435
	}
L435:
	;
	F_CreateComments(m, v1664, int32(3381), int32(0), v1813)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L25
	} else {
		goto L436
	}
L436:
	;
	v1824 = v1812
	goto L3
L437:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L25
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(149024), int32(0))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L25
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(494396), int32(439), int32(174043))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L25
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L25
	} else {
		goto L442
	}
L442:
	;
	F_errmsg(m, int32(250359), int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L25
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(494396), int32(457), int32(174043))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L25
	} else {
		goto L444
	}
L444:
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
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
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
	return v69 + v71
L2:
	;
	v9 = l0
	v10 = v2
	v11 = v5
	goto L5
L3:
	;
	v60 = v2
	v61 = v5
	goto L4
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v69 = v60
	v71 = v63&int32(4095) + int32(1)
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
	v60 = v53
	v61 = v55
	goto L4
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v50 == int32(1) {
		v69 = v10
		v71 = v49
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v49 = v46 + v48
	goto L7
L9:
	;
	v21 = v13
	v22 = v14
	v23 = v17
	goto L12
L10:
	;
	v37 = v14
	v38 = v17
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v46 = v37
	v48 = v40&int32(4095) + int32(1)
	goto L8
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_calcstrlen(m, v25)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v27 == int32(1) {
		v46 = v22
		v48 = v26
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v37 = v30
	v38 = v32
	goto L11
L14:
	;
	v30 = v22 + v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 != int32(1) {
		v21 = v31
		v22 = v30
		v23 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v53 = v10 + v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != int32(1) {
		v9 = v54
		v10 = v53
		v11 = v55
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
	var v2 int64
	_ = v2
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
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
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
	var v66 int64
	_ = v66
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	v2 = int64(0)
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
	v76 = v2
	goto L3
L3:
	;
	return v76
L4:
	;
	F_list_free(m, v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L5:
	;
	return int64(0)
L6:
	;
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v66 = v2
	goto L4
L8:
	;
	goto L9
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		v66 = v2
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v23 = v2
	v29 = int32(0)
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v29<<(uint(int32(2))%32))))
	v36 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v66 = v60
	goto L4
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v40 = F_calculate_relation_size(m, v36, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v44 = F_calculate_relation_size(m, v36, v42, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v48 = F_calculate_relation_size(m, v36, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v52 = F_calculate_relation_size(m, v36, v50, int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v60 = v52 + (v48 + (v44 + (v23 + v40)))
	v62 = v29 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v62 < v63 {
		v23 = v60
		v29 = v62
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v76 = v66
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
			v23 = int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v25&int32(254) == int32(2) {
				v34 = v23
			} else {
				v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
			}
			if v25 == int32(1) {
				v37 = v23
			} else {
				v37 = v34
			}
			v48 = v37
		} else {
			v38 = int32(1)
			if v19 != 0 {
				v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v50 = int32(0)
		v51 = m.G0
		v53 = v51 - int32(16)
		m.G0 = v53
		if v20 == v50 {
			v109 = v50
			m.G0 = v53 + int32(16)
			v156 = F_cstring_to_text(m, v109)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v109)
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					return v156
				}
			}
		} else {
			if v49 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(683800)
						F_errmsg(m, int32(252067), v53)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(574955), int32(0))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498322), int32(1847), int32(429899))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
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
				v60 = *(*int32)(unsafe.Add(mBase, _consts[251]))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				if v61 != int32(6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(556139), int32(0))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498322), int32(1853), int32(429899))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
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
					v64 = F_pg_newlocale_from_collation(m, v49)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+3)))
						if v66 == int32(1) {
							v69 = F_pnstrdup(m, v20, v48)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								if v71 == int32(0) {
									v109 = v69
								} else {
									v75 = v71
									v77 = v69
									for {
										v81 = int32(255)
										v82 = v75 & v81
										if base.Ui32((v82-int32(65))&v81) < base.Ui32(int32(26)) {
											v91 = v82 | int32(32)
										} else {
											v91 = v82
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v91)
										v94 = v77 + int32(1)
										v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
										if v95 != 0 {
											v75 = v95
											v77 = v94
											continue
										} else {
											break
										}
										break
									}
									v109 = v69
								}
								m.G0 = v53 + int32(16)
								v156 = F_cstring_to_text(m, v109)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v109)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										return v156
									}
								}
							}
						} else {
							v97 = v48 + int32(1)
							v98 = F_palloc(m, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v100 = F_pg_strfold(m, v98, v97, v20, v48, v64)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v103 = v100 + int32(1)
									if base.Ui32(v103) <= base.Ui32(v97) {
										v109 = v98
										m.G0 = v53 + int32(16)
										v156 = F_cstring_to_text(m, v109)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v109)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												return v156
											}
										}
									} else {
										v105 = F_repalloc(m, v98, v103)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v107 = F_pg_strfold(m, v105, v103, v20, v48, v64)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v109 = v105
												m.G0 = v53 + int32(16)
												v156 = F_cstring_to_text(m, v109)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v109)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
	v7 = F_strlen(m, v5)
	mBase = m.M
	if v7 != int32(4) {
		v43 = v6
		return base.I32_extend8_s(v43)
	} else {
		if v6&int32(255) != int32(92) {
			v43 = v6
			return base.I32_extend8_s(v43)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v14&int32(248) != int32(48) {
				return int32(92)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
				if v21&int32(248) != int32(48) {
					return int32(92)
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
					if v29&int32(248) != int32(48) {
						v43 = int32(92)
					} else {
						v43 = v14<<(uint(int32(6))%32) + v21<<(uint(int32(3))%32) + v29 + int32(80)
					}
					return base.I32_extend8_s(v43)
				}
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
	F_errmsg(m, int32(71628), v15)
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
	F_errfinish(m, int32(493497), int32(1073), int32(124965))
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
	F_errmsg(m, int32(415610), v15+int32(16))
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
	F_errfinish(m, int32(493497), int32(1088), int32(124965))
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
	F_errmsg(m, int32(415610), v15+int32(32))
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
	F_errfinish(m, int32(493497), int32(1099), int32(124965))
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
					F_errmsg(m, int32(313204), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(497659), int32(600), int32(308591))
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
						F_errmsg(m, int32(158372), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(497659), int32(604), int32(308591))
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
				F_errmsg(m, int32(371114), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(497659), int32(596), int32(308591))
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
				F_errmsg_internal(m, int32(43115), v9)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498742), int32(214), int32(362977))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = F_strlen(m, v11)
	mBase = m.M
	v13 = int32(757356)
	v17 = m.G0
	v19 = v17 - int32(32)
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v20
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1210])))
	if v28 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v12 != v96 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v96 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1211])))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v36 = v11
	goto L8
L6:
	;
	goto L7
L7:
	;
	v46 = v13
	v47 = v28
	goto L11
L8:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v42 == v28 {
		v36 = v36 + int32(1)
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v96 = v36 - v11
	goto L1
L10:
	;
	goto L9
L11:
	;
	v54 = v19 + int32(base.Ui32(v47)>>(uint(int32(3))%32))&int32(28)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55 | v56<<(uint(v47)%32)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	if v60 != 0 {
		v46 = v46 + v56
		v47 = v60
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v63 == int32(0) {
		v88 = v11
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v96 = v88 - v11
	goto L1
L15:
	;
	v67 = v11
	v68 = v63
	goto L16
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v68)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v76)>>(uint(v68)%32))&int32(1) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v88 = v84
	goto L14
L18:
	;
	v88 = v67
	goto L14
L19:
	;
	goto L20
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v84 = v67 + int32(1)
	if v82 != 0 {
		v67 = v84
		v68 = v82
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[87])) = v100
	v105 = F_format_elog_string(m, int32(609498), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v112 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	return int32(0)
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v105
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(1)
L28:
	;
	goto L29
L29:
	;
	v121 = F_guc_malloc(m, v12+int32(2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	if v121 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	goto L33
L33:
	;
	if v12 <= int32(0) {
		v207 = v4
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v213 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v207+v121))) = uint16(v213)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v121
	return int32(1)
L35:
	;
	v129 = int32(1)
	if v12 == v129 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v12&v129 == int32(0) {
		v207 = v184
		goto L34
	} else {
		goto L49
	}
L37:
	;
	v182 = int32(0)
	v184 = v4
	goto L36
L38:
	;
	goto L39
L39:
	;
	v136 = int32(0)
	v141 = v136
	v142 = v136
	v143 = v4
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v141))))
	switch v151 - int32(9) {
	case 0, 1, 23:
		v159 = v143
		goto L42
	default:
		goto L44
	case 35:
		v154 = int32(0)
		goto L43
	}
L41:
	;
	v182 = v175
	v184 = v172
	goto L36
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v141)+1)))
	switch v164 - int32(9) {
	case 0, 1, 23:
		v172 = v159
		goto L45
	default:
		goto L47
	case 35:
		v167 = int32(0)
		goto L46
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v121))) = uint8(v154)
	v159 = v143 + int32(1)
	goto L42
L44:
	;
	v154 = v151
	goto L43
L45:
	;
	v174 = int32(2)
	v175 = v141 + v174
	v177 = v142 + v174
	if v177 != v12&int32(2147483646) {
		v141 = v175
		v142 = v177
		v143 = v172
		goto L40
	} else {
		goto L48
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v159+v121))) = uint8(v167)
	v172 = v159 + int32(1)
	goto L45
L47:
	;
	v167 = v164
	goto L46
L48:
	;
	goto L41
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v182))))
	switch v194 - int32(9) {
	case 0, 1, 23:
		v207 = v184
		goto L34
	default:
		goto L51
	case 35:
		v197 = int32(0)
		goto L50
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v184+v121))) = uint8(v197)
	v207 = v184 + int32(1)
	goto L34
L51:
	;
	v197 = v194
	goto L50
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
	F_errcode(m, int32(290948))
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
	F_errmsg(m, int32(415294), v13)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(494448), int32(2383), int32(73848))
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
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v47)%32))+uint32(_consts[257])))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v51
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v47)%32))+uint32(_consts[257])))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
				F_errmsg_internal(m, int32(698587), v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errfinish(m, int32(494088), int32(1806), int32(155343))
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
					F_errmsg_internal(m, int32(487953), v7+int32(-32))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errfinish(m, int32(494088), int32(1808), int32(155343))
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
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v86)%32))+uint32(_consts[257])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v90
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l4<<(uint(v86)%32))+uint32(_consts[257])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v96
						F_errmsg_internal(m, int32(698538), v7+int32(-48))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(494088), int32(1812), int32(155343))
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
							F_errmsg_internal(m, int32(343580), int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								F_errfinish(m, int32(494088), int32(1814), int32(155343))
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
			F_errmsg_internal(m, int32(487989), v7+int32(-16))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(494088), int32(1802), int32(155343))
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
	F_errmsg_internal(m, int32(401251), v26+int32(80))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(498307), int32(1165), int32(8971))
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
	F_errmsg_internal(m, int32(694394), v26-int32(-64))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(498307), int32(838), int32(89354))
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
	F_errmsg(m, int32(698911), v26+int32(16))
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
	F_errmsg(m, int32(698854), v26+int32(48))
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
	F_errdetail(m, int32(603097), v26)
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
	F_errdetail(m, int32(578535), int32(0))
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
	F_errfinish(m, int32(498307), int32(918), int32(89354))
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
	F_errdetail(m, int32(603127), v26+int32(32))
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
	F_errdetail(m, int32(576062), int32(0))
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
	F_errfinish(m, int32(498307), int32(929), int32(89354))
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
	F_errmsg(m, int32(706846), v26+int32(96))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(498307), int32(1173), int32(8971))
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
										F_errmsg(m, int32(269794), v9)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											F_errdetail(m, int32(619307), int32(0))
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
													F_errfinish(m, int32(496174), int32(3207), int32(221765))
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
								F_errmsg(m, int32(269670), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errdetail(m, int32(628870), int32(0))
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
											F_errfinish(m, int32(496174), int32(3214), int32(221765))
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
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[330])))
		if v18 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[331]))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[332])))
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
			v26 = *(*int32)(unsafe.Add(mBase, _consts[39]))
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
										F_errmsg(m, int32(72320), v11)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(499222), int32(864), int32(258636))
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
							v63 = *(*int32)(unsafe.Add(mBase, _consts[86]))
							*(*int32)(unsafe.Add(mBase, _consts[87])) = v63
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v66
							v72 = F_format_elog_string(m, int32(72320), v9+int32(-48))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[333])) = v72
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
							v83 = *(*int32)(unsafe.Add(mBase, _consts[334]))
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
								v86 = *(*int32)(unsafe.Add(mBase, _consts[334]))
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
														F_errmsg(m, int32(704015), v9+int32(-32))
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(499222), int32(890), int32(258636))
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
											*(*int32)(unsafe.Add(mBase, _consts[335])) = int32(16797828)
											v117 = *(*int32)(unsafe.Add(mBase, _consts[86]))
											*(*int32)(unsafe.Add(mBase, _consts[87])) = v117
											v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v120
											v126 = F_format_elog_string(m, int32(703963), v9+int32(-16))
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[333])) = v126
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
		v52 = int32(139543)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[264])))
					v88 = v87
				} else {
					v88 = int32(425660)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v88
				F_errmsg(m, int32(185262), v8+int32(16))
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
						F_errfinish(m, int32(500150), int32(2681), int32(96486))
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
					F_errmsg(m, int32(531718), int32(0))
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
							F_errfinish(m, int32(500150), int32(2555), int32(96486))
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
		v52 = int32(139117)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(140071)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(119701)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(146686)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(144957)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(160819)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(145516)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(146272)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(133099)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(139772)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(144725)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(424901)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(144429)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(122550)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(140321)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v52 = int32(145899)
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
				F_errmsg_internal(m, int32(206170), v8)
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
						F_errfinish(m, int32(500150), int32(2674), int32(96486))
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[86]))
		*(*int32)(unsafe.Add(mBase, _consts[87])) = v8
		v14 = F_format_elog_string(m, int32(431009), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[333])) = v14
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
	v6 = F_uint32in_subr(m, v2, int32(0), int32(436251), v5)
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v5 {
		v118 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v118
L2:
	;
	v24 = l0
	v26 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v118 = int32(0)
	goto L1
L5:
	;
	if l2 != 0 {
		v118 = v102
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 <= v35 {
		v102 = v35
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v44 = v35
	v49 = v35
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v49<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	if v57 != int32(1) {
		v74 = v44
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
	v77 = v49 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v77 < v78 {
		v44 = v74
		v49 = v77
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
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
	if v63 != int32(1) {
		v74 = v44
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v66 = F_scanNSItemForColumn(m, l0, v56, v26, l1, l3)
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
		v74 = v44
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_check_lateral_ref_ok(m, v24, v56, l3)
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
	v102 = v74
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
	F_errmsg(m, int32(114940), v15)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_parser_errposition(m, v24, l3)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(496480), int32(932), int32(230578))
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
	if v102 != 0 {
		v118 = v102
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v112 != 0 {
		v24 = v112
		v26 = v26 + int32(1)
		goto L3
	} else {
		goto L32
	}
L32:
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
				v29 = int32(0)
				v30 = F_palloc0(m, v11)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l2) < base.Ui32(v11) {
						v33 = l2
					} else {
						v33 = v11
					}
					if v33 != 0 {
						v34 = F__emscripten_memcpy_bulkmem(m, v30, l1, v33)
						mBase = m.M
						v35 = v34
					} else {
						v35 = v30
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v35, v33, v29)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							F_pfree(m, v29)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v35)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						} else {
							F_pfree(m, v35)
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
							v29 = v20
						} else {
							v25 = l4
							if v25 != 0 {
								v26 = F__emscripten_memcpy_bulkmem(m, v20, l3, v25)
								mBase = m.M
							} else {
							}
							v29 = v20
						}
					} else {
						v25 = v16
						if v25 != 0 {
							v26 = F__emscripten_memcpy_bulkmem(m, v20, l3, v25)
							mBase = m.M
						} else {
						}
						v29 = v20
					}
					v30 = F_palloc0(m, v11)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if base.Ui32(l2) < base.Ui32(v11) {
							v33 = l2
						} else {
							v33 = v11
						}
						if v33 != 0 {
							v34 = F__emscripten_memcpy_bulkmem(m, v30, l1, v33)
							mBase = m.M
							v35 = v34
						} else {
							v35 = v30
						}
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v35, v33, v29)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								F_pfree(m, v29)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v35)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										return v37
									}
								}
							} else {
								F_pfree(m, v35)
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v104 int32
	_ = v104
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
	var v160 int32
	_ = v160
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
	var v191 int32
	_ = v191
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
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = F_palloc_extended(m, v30, int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v30 = v3
	v31 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = v8
	v14 = v3
	v15 = v3
	goto L5
L5:
	;
	v18 = int32(1)
	v19 = v14 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v23 = v15 + v20 + v18
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v24 != 0 {
		v13 = v24
		v14 = v19
		v15 = v23
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v30 = v19
	v31 = v23 << (uint(int32(3)) % 32)
	goto L1
L7:
	;
	goto L6
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v35
	v38 = int32(2)
	v41 = F_palloc_extended(m, v30<<(uint(v38)%32), v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v41
	v45 = F_palloc_extended(m, v31, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
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
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v51 = v49
	goto L18
L17:
	;
	v51 = int32(0)
	goto L18
L18:
	;
	if v51 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	F_pfree(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = int32(101)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v64 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v66 = v64
	goto L31
L30:
	;
	v66 = int32(12)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v66
	return
L32:
	;
	v90 = int32(0)
	goto L34
L33:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+12)))
	v90 = v87 + int32(1)
	goto L34
L34:
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
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = v45
	v105 = v98
	goto L38
L36:
	;
	goto L37
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	if v188 != 0 {
		goto L61
	} else {
		goto L62
	}
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v107))) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(int32(2))%32)))) = v104
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L37
L40:
	;
	v120 = v117
	v121 = v104
	goto L43
L41:
	;
	v160 = v104
	goto L42
L42:
	;
	v166 = (v160 - v104) >> (uint(int32(3)) % 32)
	if base.Ui32(int32(2)) <= base.Ui32(v166) {
		goto L56
	} else {
		goto L57
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v125 != int32(76) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v160 = v155
	goto L42
L45:
	;
	v155 = v121 + int32(8)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v156 != 0 {
		v120 = v156
		v121 = v155
		goto L43
	} else {
		goto L55
	}
L46:
	;
	if v125 == int32(112) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v145 = v143 + v144
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v145)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v150 | int32(1)
	goto L45
L49:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v133
	goto L45
L50:
	;
	goto L51
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = int32(101)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v139 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v141 = v139
	goto L54
L53:
	;
	v141 = int32(15)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v141
	return
L55:
	;
	goto L44
L56:
	;
	F_pg_qsort(m, v104, v166, int32(8), int32(971))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(0)
	v175 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v160))) = uint16(v175)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v105)+28))
	if v179 != 0 {
		v104 = v160 + int32(8)
		v105 = v179
		goto L38
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L39
L61:
	;
	v191 = v188
	goto L64
L62:
	;
	v207 = v187
	goto L63
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v211+v212))) = uint8(v214)
	return
L64:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v196+v198))) = uint8(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v191)+16))
	if v202 != 0 {
		v191 = v202
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = v203
	goto L63
L66:
	;
	goto L65
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
	v8 = int32(16383)
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
	var v22 int32
	_ = v22
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
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	if v137 < int32(0) {
		goto L46
	} else {
		goto L47
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
		v137 = v16
		v139 = v13
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v22 = l0 + int32(16)
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
	v137 = v127
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
	v137 = v42
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
		goto L45
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
	v110 = v22 + v109
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v71)
	v112 = v67 - v71
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+2)) = uint16(v112)
	v115 = v110 + int32(4)
	v118 = v112 & int32(65535)
	if v118 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v120 + v118 - v22
	v124 = int32(-1)
	v127 = v124
	v128 = v124
	goto L36
L42:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v115, l2+v71, v118)
	mBase = m.M
	v120 = v119
	goto L44
L43:
	;
	v120 = v115
	goto L44
L44:
	;
	goto L41
L45:
	;
	goto L13
L46:
	;
	v144 = v19
	goto L48
L47:
	;
	v144 = v137
	goto L48
L48:
	;
	v145 = base.B2i32(base.Ui32(l6) < base.Ui32(l4))
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v146 = v144
	goto L51
L50:
	;
	v146 = v137
	goto L51
L51:
	;
	if int32(0) <= v146 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v150 = l0 + int32(16)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = v150 + v151
	*(*uint16)(unsafe.Add(mBase, uint32(v152))) = uint16(v146)
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	return
L55:
	;
	v154 = l4
	goto L57
L56:
	;
	v154 = v139
	goto L57
L57:
	;
	if v154 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v157 = l4
	goto L60
L59:
	;
	v157 = v154
	goto L60
L60:
	;
	v158 = v157 - v146
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)) = uint16(v158)
	v161 = v152 + int32(4)
	v164 = v158 & int32(65535)
	if v164 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v166 + v164 - v150
	goto L54
L62:
	;
	v165 = F__emscripten_memcpy_bulkmem(m, v161, l2+v146, v164)
	mBase = m.M
	v166 = v165
	goto L64
L63:
	;
	v166 = v161
	goto L64
L64:
	;
	goto L61
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v7 = l1
	v8 = l2
	v9 = l3
	goto L2
L1:
	;
	if l0 != v7 {
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
	v22 = v8 << (uint(int32(3)) % 32)
	if l0 == v7 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v167 + v9
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 - v9
	return v8
L9:
	;
	goto L8
L10:
	;
	goto L9
L11:
	;
	v26 = l0 + v22
	if base.Ui32(v7-v26) <= base.Ui32(int32(0)-v22<<(uint(int32(1))%32)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = F___memcpy(m, l0, v7, v22)
	mBase = m.M
	goto L9
L13:
	;
	goto L14
L14:
	;
	v36 = (l0 ^ v7) & int32(3)
	if base.Ui32(l0) < base.Ui32(v7) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v138 == int32(0) {
		goto L10
	} else {
		goto L51
	}
L16:
	;
	if base.Ui32(v116) <= base.Ui32(int32(3)) {
		v137 = v115
		v138 = v116
		v139 = v117
		goto L15
	} else {
		goto L47
	}
L17:
	;
	if v36 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v36 != 0 {
		v98 = v22
		goto L30
	} else {
		goto L31
	}
L20:
	;
	v137 = v7
	v138 = v22
	v139 = l0
	goto L15
L21:
	;
	goto L22
L22:
	;
	if l0&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v115 = v7
	v116 = v22
	v117 = l0
	goto L16
L24:
	;
	goto L25
L25:
	;
	v43 = v7
	v44 = v22
	v45 = l0
	goto L26
L26:
	;
	if v44 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L27:
	;
	v115 = v52
	v116 = v54
	v117 = v56
	goto L16
L28:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v49)
	v51 = int32(1)
	v52 = v43 + v51
	v54 = v44 - v51
	v56 = v45 + v51
	if v56&int32(3) != 0 {
		v43 = v52
		v44 = v54
		v45 = v56
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v98 == int32(0) {
		goto L10
	} else {
		goto L43
	}
L31:
	;
	if v26&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v63 = v22
	goto L35
L33:
	;
	v78 = v22
	goto L34
L34:
	;
	if base.Ui32(v78) <= base.Ui32(int32(3)) {
		v98 = v78
		goto L30
	} else {
		goto L39
	}
L35:
	;
	if v63 == int32(0) {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	v78 = v69
	goto L34
L37:
	;
	v69 = v63 - int32(1)
	v70 = l0 + v69
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v69))))
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v72)
	if v70&int32(3) != 0 {
		v63 = v69
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v85 = v78
	goto L40
L40:
	;
	v89 = v85 - int32(4)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v7+v89)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v89))) = v92
	if base.Ui32(int32(3)) < base.Ui32(v89) {
		v85 = v89
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v98 = v89
	goto L30
L42:
	;
	goto L41
L43:
	;
	v105 = v98
	goto L44
L44:
	;
	v109 = v105 - int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v109))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v109))) = uint8(v112)
	if v109 != 0 {
		v105 = v109
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L10
L46:
	;
	goto L45
L47:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L48
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v126
	v128 = int32(4)
	v129 = v122 + v128
	v131 = v124 + v128
	v133 = v123 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v122 = v129
		v123 = v133
		v124 = v131
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v137 = v129
	v138 = v133
	v139 = v131
	goto L15
L50:
	;
	goto L49
L51:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L52
L52:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v148)
	v150 = int32(1)
	v155 = v145 - v150
	if v155 != 0 {
		v144 = v144 + v150
		v145 = v155
		v146 = v146 + v150
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L10
L54:
	;
	goto L53
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 float64
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
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
	v24 = int32(65535)
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
	v39 = v5
	v41 = v5
	v42 = v5
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
	if int32(0) < v99 {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	return
L9:
	;
	v51 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v39, v15+int32(15))
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
	v102 = v39 + int32(1)
	if v102 != l2 {
		v37 = v96
		v39 = v102
		v41 = v98
		v42 = v99
		goto L6
	} else {
		goto L28
	}
L12:
	;
	v96 = v37
	v98 = v41 + int32(1)
	v99 = v42
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
		v96 = v37
		v98 = v41
		v99 = v59
		goto L11
	} else {
		goto L27
	}
L18:
	;
	v96 = base.F64_add(v37, base.F64_convert_i32_u(v86))
	v98 = v41
	v99 = v59
	goto L11
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v86 = int32(6)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	if v60&v78 != 0 {
		v86 = int32(base.Ui32(v60) >> (uint(v78) % 32))
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v71 = int32(18)
	if v64&int32(255) == v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = v71
	goto L25
L24:
	;
	v77 = int32(2)
	goto L25
L25:
	;
	v86 = v77
	goto L18
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v86 = int32(base.Ui32(v82) >> (uint(int32(2)) % 32))
	goto L18
L27:
	;
	v91 = F_strlen(m, v51)
	mBase = m.M
	v96 = base.F64_add(v37, base.F64_convert_i32_u(v91+int32(1)))
	v98 = v41
	v99 = v59
	goto L11
L28:
	;
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v132
	goto L4
L30:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v106)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v98), base.F64_convert_i32_s(l2)))
	if v29 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v98 <= int32(0) {
		goto L4
	} else {
		goto L39
	}
L33:
	;
	v114 = base.F64_div(v96, base.F64_convert_i32_u(v99))
	if base.F64_lt(base.F64_abs(v114), float64(2.147483648e+09)) != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v120)+76)))
	v132 = v121
	goto L29
L36:
	;
	v118 = base.I32_trunc_f64_s(v114)
	v132 = v118
	goto L29
L37:
	;
	goto L38
L38:
	;
	v132 = int32(-2147483648)
	goto L29
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v126)
	if v29 != 0 {
		v132 = int32(0)
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+76)))
	v132 = v130
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
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
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
	var v113 int32
	_ = v113
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
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
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
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L68
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
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L64
	}
L6:
	;
	v27 = F_ArrayGetNItems(m, l2, l3)
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
	return v267
L11:
	;
	v56 = v10
	v57 = v10
	v62 = v10
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
	v267 = v40
	goto L10
L16:
	;
	v237 = v227 + v228
	v238 = F_palloc0(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L54
	}
L17:
	;
	v227 = (l2<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v228 = v174
	v236 = int32(0)
	goto L16
L18:
	;
	v203 = base.I32_div_s(v27+int32(7), int32(8))
	v210 = (v203 + l2<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v227 = v210
	v228 = v192
	v236 = v210
	goto L16
L19:
	;
	if l1 == int32(0) {
		v107 = v56
		v113 = v62
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v113 == int32(0) {
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
		v113 = v62
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = v56 + int32(1)
	if v72 == v27 {
		v192 = v57
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
	v113 = int32(1)
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
	v192 = v57
	goto L18
L31:
	;
	v161 = v57 + v159
	switch l8 - int32(99) {
	case 0:
		v174 = v161
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
		v159 = l6
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
	v159 = v124 + int32(1)
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
		v159 = int32(6)
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
	if v138&int32(255) == v145 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v151 = v145
	goto L43
L42:
	;
	v151 = int32(2)
	goto L43
L43:
	;
	v159 = v151
	goto L31
L44:
	;
	v159 = int32(base.Ui32(v134) >> (uint(int32(1)) % 32))
	goto L31
L45:
	;
	goto L46
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v159 = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
	goto L31
L47:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v174) {
		goto L1
	} else {
		goto L51
	}
L48:
	;
	v174 = (v161 + int32(1)) & int32(-2)
	goto L47
L49:
	;
	v174 = (v161 + int32(7)) & int32(-8)
	goto L47
L50:
	;
	v174 = (v161 + int32(3)) & int32(-4)
	goto L47
L51:
	;
	v178 = v107 + int32(1)
	if v178 != v27 {
		v56 = v178
		v57 = v174
		v62 = v113
		goto L19
	} else {
		goto L52
	}
L52:
	;
	goto L20
L53:
	;
	v192 = v174
	goto L18
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = l2
	v243 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v237 << (uint(v243) % 32)
	v247 = v238 + int32(16)
	v249 = l2 << (uint(v243) % 32)
	if v249 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v249 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v250 = F__emscripten_memcpy_bulkmem(m, v247, l3, v249)
	mBase = m.M
	v251 = v250
	goto L58
L57:
	;
	v251 = v247
	goto L58
L58:
	;
	goto L55
L59:
	;
	F_CopyArrayEls(m, v238, l0, l1, v27, l6, l7, l8, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L63
	}
L60:
	;
	v253 = F__emscripten_memcpy_bulkmem(m, v251+v249, l4, v249)
	mBase = m.M
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v267 = v238
	goto L10
L64:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	F_errmsg(m, int32(481943), v21)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(494676), int32(3511), int32(24611))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	F_errmsg(m, int32(679749), v21+int32(16))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(494676), int32(3516), int32(24611))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(1073741823)
	F_errmsg(m, int32(679651), v21+int32(32))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(494676), int32(3546), int32(24611))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = F_palloc0(m, int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v391 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L2:
	;
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v24
	v31 = int32(1)
	v34 = F_palloc0(m, v24<<(uint(v31)%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v34
	v38 = l1 + int32(20)
	if v24 <= int32(0) {
		v208 = v4
		v213 = v31
		v214 = v4
		v218 = v34
		v219 = v4
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v224 <= v208 {
		v295 = v213
		v296 = v214
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v45 = v4
	v48 = v4
	v50 = v31
	v51 = v4
	v55 = v34
	v56 = v4
	goto L7
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = v38 + v61<<(uint(int32(4))%32) + v48*int32(100)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+91)))
	if v68 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L27
	}
L9:
	;
	goto L8
L10:
	;
	v72 = v56 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v73 <= v45 {
		v124 = v45
		v130 = v51
		v134 = v55
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v149 = v45
	v154 = v50
	v155 = v51
	v159 = v55
	v160 = v56
	goto L12
L12:
	;
	v166 = v48 + int32(1)
	if v24 != v166 {
		v45 = v149
		v48 = v166
		v50 = v154
		v51 = v155
		v55 = v159
		v56 = v160
		goto L7
	} else {
		goto L26
	}
L13:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v48<<(uint(int32(1))%32)))))
	v149 = v124
	v154 = base.B2i32(v143 != int32(0)) & v50
	v155 = v130
	v159 = v134
	v160 = v72
	goto L12
L14:
	;
	v80 = v45
	goto L15
L15:
	;
	v98 = l0 + int32(20) + v73<<(uint(int32(4))%32) + v80*int32(100)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+91)))
	if v99 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v124 = v73
	v130 = v51
	v134 = v55
	goto L13
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
	if v102 != v103 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v120 = v80 + int32(1)
	if v120 != v73 {
		v80 = v120
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	if int32(0) <= v105 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	if v105 != v108 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v110 = int32(1)
	v116 = v80 + v110
	*(*uint16)(unsafe.Add(mBase, uint32(v55+v48<<(uint(v110)%32)))) = uint16(v116)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v124 = v116
	v130 = v51 + v110
	v134 = v118
	goto L13
L24:
	;
	goto L23
L25:
	;
	goto L16
L26:
	;
	v208 = v149
	v213 = v154
	v214 = v155
	v218 = v159
	v219 = v160
	goto L5
L27:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = l2
	F_errmsg_internal(m, int32(206170), v20+int32(-16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	v184 = F_format_type_with_typemod(m, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	v188 = F_format_type_with_typemod(m, v186, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v67 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v184
	F_errdetail(m, int32(662198), v20+int32(-32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(496102), int32(124), int32(249505))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
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
	if v295&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v226 = int32(1)
	v227 = v208 + v226
	v229 = l0 + int32(29)
	if (v224-v208)&v226 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v208<<(uint(int32(4))%32)))))
	v243 = v227
	v244 = v236 & v213
	v245 = v214 + (v236^int32(-1))&int32(1)
	goto L38
L37:
	;
	v243 = v208
	v244 = v213
	v245 = v214
	goto L38
L38:
	;
	if v227 == v224 {
		v295 = v244
		v296 = v245
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v251 = v243
	v256 = v244
	v257 = v245
	goto L40
L40:
	;
	v268 = v251 << (uint(int32(4)) % 32)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(45)+v268))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v229))))
	v274 = v270 & v272 & v256
	v275 = int32(-1)
	v277 = int32(1)
	v284 = v257 + (v272^v275)&v277 + (v270^v275)&v277
	v286 = v251 + int32(2)
	if v286 != v224 {
		v251 = v286
		v256 = v274
		v257 = v284
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v295 = v274
	v296 = v284
	goto L34
L42:
	;
	goto L41
L43:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v224 != v308 {
		v391 = v26
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L64
	}
L46:
	;
	m.G0 = v22 - int32(-64)
	goto L1
L47:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if int32(0) < v310 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v318 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v218)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L62
	}
L51:
	;
	v335 = v318 << (uint(int32(4)) % 32)
	v336 = l0 + int32(20) + v335
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
	if v337 != 0 {
		v391 = v26
		goto L46
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v338 = int32(1)
	v339 = v318 + v338
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v218+v318<<(uint(v338)%32)))))
	if v339 != v343 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v343 != 0 {
		v391 = v26
		goto L46
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v339 != v310 {
		v318 = v339
		goto L51
	} else {
		goto L61
	}
L57:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
	if v345 != int32(1) {
		v391 = v26
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+4)))
	v349 = v335 + v38
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+4)))
	if v348 != v350 {
		v391 = v26
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+12)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+12)))
	if v352 != v353 {
		v391 = v26
		goto L46
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	goto L52
L62:
	;
	F_pfree(m, v26)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v391 = int32(0)
	goto L46
L64:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l2
	F_errmsg_internal(m, int32(206170), v20+int32(-48))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v296
	F_errdetail(m, int32(662070), v22)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(496102), int32(149), int32(249505))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	return int32(0)
L70:
	;
	goto L71
L71:
	;
	v429 = F_palloc(m, int32(28))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = l0
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v436 = v434 + int32(1)
	v439 = F_palloc(m, v436<<(uint(int32(2))%32))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+20)) = v439
	v442 = F_palloc(m, v436)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+24)) = v442
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v447 = v445 + int32(1)
	v450 = F_palloc(m, v447<<(uint(int32(2))%32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+12)) = v450
	v453 = F_palloc(m, v447)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+16)) = v453
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v429)+16))
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v459))) = uint8(v460)
	return v429
}
func F_copy_addr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v2 = l1
	if v2&int32(255) != int32(10) {
		if v2 != int32(2) {
		} else {
			v15 = int32(4)
			v33 = v15
			v34 = l2 + v15
			if base.Ui32(l4) < base.Ui32(v33) {
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
				if v33 != 0 {
					v37 = F__emscripten_memcpy_bulkmem(m, v34, l3, v33)
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
			}
		}
	} else {
		v19 = l2 + int32(8)
		v20 = int32(16)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
		switch v21 - int32(254) {
		case 0:
			v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+1)))
			if v24 < int32(-64) {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
				v33 = v20
				v34 = v19
			} else {
				v33 = v20
				v34 = v19
			}
		case 1:
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
			if v27&int32(15) != int32(2) {
				v33 = v20
				v34 = v19
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
				v33 = v20
				v34 = v19
			}
		default:
			v33 = v20
			v34 = v19
		}
		if base.Ui32(l4) < base.Ui32(v33) {
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
			if v33 != 0 {
				v37 = F__emscripten_memcpy_bulkmem(m, v34, l3, v33)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
		}
	}
	return
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	v13 = *(*float64)(unsafe.Add(mBase, _consts[387]))
	v17 = base.F64_add(base.F64_mul(base.F64_add(v13, v13), l4), base.F64_sub(l3, l2))
	v25 = base.F64_mul(l4, base.F64_convert_i32_u((l5+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v25, base.F64_convert_i32_u(v10<<(uint(int32(10))%32))) != 0 {
		v31 = *(*float64)(unsafe.Add(mBase, _consts[389]))
		v37 = base.F64_add(base.F64_mul(v31, base.F64_ceil(base.F64_mul(v25, float64(0.0001220703125)))), v17)
	} else {
		v37 = v17
	}
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[390])))
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
	var v112 float64
	_ = v112
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
	var v135 float64
	_ = v135
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
	v148 = *(*float64)(unsafe.Add(mBase, _consts[388]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l1
	if v59 == int32(0) {
		v112 = v9
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
	v135 = v125
	v143 = v126
	goto L11
L15:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v135 = base.F64_add(v112, v121)
	v143 = base.F64_add(v120, v123)
	goto L11
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v69 <= int32(0) {
		v112 = v9
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
	v112 = v102
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v154, base.F64_add(base.F64_mul(v145, v156), base.F64_add(base.F64_mul(v146, base.F64_add(v135, v148)), base.F64_add(base.F64_mul(v160, base.F64_convert_i32_u(v55)), float64(0)))))
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
	var v27 float64
	_ = v27
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
	var v43 float64
	_ = v43
	var v47 float64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v114 int32
	_ = v114
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v137 int32
	_ = v137
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
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
	v27 = float64(1e+100)
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v50
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v49)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v52
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v49)+56))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v54
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v35 = base.F64_mul(v29, v33)
	if base.F64_gt(v35, float64(1e+100)) != 0 {
		v47 = v27
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) {
		v47 = v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = float64(1)
	if base.F64_le(v35, v43) != 0 {
		v47 = v43
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v47 = base.F64_nearest(v35)
	goto L7
L12:
	;
	v57 = int32(0)
	goto L14
L13:
	;
	v57 = l4
	goto L14
L14:
	;
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	m.G0 = v19 + int32(32)
	return
L18:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v149)+24))
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v137)+32))
	v153 = *(*float64)(unsafe.Add(mBase, _consts[388]))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v149)+16))
	v155 = base.F64_add(v148, v154)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v155, v140)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_add(v155, base.F64_add(base.F64_mul(v150, v141), base.F64_mul(v151, base.F64_add(v142, v153)))), v144)
	goto L17
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	if v60 == int32(0) {
		v114 = v49
		v117 = v52
		v118 = v47
		v119 = v9
		v121 = v54
		v125 = float64(0)
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v137 = v49
	v140 = v52
	v141 = v47
	v142 = v130
	v144 = v54
	v148 = v131
	goto L18
L22:
	;
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v128 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v137 = v114
	v140 = v117
	v141 = v118
	v142 = base.F64_add(v119, v126)
	v144 = v121
	v148 = base.F64_add(v125, v128)
	goto L18
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v70 <= int32(0) {
		v114 = v49
		v117 = v52
		v118 = v47
		v119 = v9
		v121 = v54
		v125 = float64(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v79 = int32(0)
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v79<<(uint(int32(2))%32))))
	v97 = F_cost_qual_eval_walker(m, v94, v19+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v114 = v106
	v117 = v104
	v118 = v105
	v119 = v107
	v121 = v103
	v125 = v108
	goto L22
L27:
	;
	v100 = v79 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v100 < v101 {
		v79 = v100
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 float64
	_ = v111
	var v124 int32
	_ = v124
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
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v31 = F_get_sortgrouplist_exprs(m, v29, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v40 = v28
	goto L1
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	if v23 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)))
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+318)))
	if v25 != int32(1) {
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
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v36 = int32(0)
	v38 = F_estimate_num_groups(m, l0, v31, v35, v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = v38
	goto L1
L11:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v331 != 0 {
		goto L113
	} else {
		goto L114
	}
L12:
	;
	if v86 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L13:
	;
	v86 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v50 <= int32(0) {
		v78 = int32(1)
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v86 = v78
	goto L12
L17:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = v50
	goto L20
L19:
	;
	v56 = v53
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = int32(0)
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = int32(0)
	v70 = base.B2i32(v68 != v69)
	if v68 == v69 {
		v78 = v70
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v78 = v70
	goto L16
L23:
	;
	v74 = v59 + int32(1)
	if v74 != v56 {
		v59 = v74
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v90 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v103 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L27:
	;
	if v89 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v102 = v89
	goto L26
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v94 = v93
	goto L32
L31:
	;
	v94 = v4
	goto L32
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = v96
	goto L35
L34:
	;
	v97 = v4
	goto L35
L35:
	;
	if v94 < v97 {
		v102 = v95
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v106 <= int32(0) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	if v89 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v111 = float64(-1)
	goto L41
L40:
	;
	v111 = float64(1)
	goto L41
L41:
	;
	v124 = v4
	goto L42
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v124<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+64))
	v133 = F_get_useful_pathkeys_for_distinct(m, l0, v102, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L11
L44:
	;
	v313 = v124 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v313 < v314 {
		v124 = v313
		goto L42
	} else {
		goto L111
	}
L45:
	;
	if v133 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v138 <= v137 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v142 = v137
	goto L48
L48:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v142<<(uint(int32(2))%32))))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v131)+64))
	v163 = v18 + int32(12)
	if v160 == v161 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L44
L50:
	;
	v294 = v142 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v294 < v295 {
		v142 = v294
		goto L48
	} else {
		goto L110
	}
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v264 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L52:
	;
	if v241 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v229
	v241 = int32(1)
	goto L52
L54:
	;
	if v160 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v160 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	v241 = int32(1)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	v241 = int32(1)
	goto L52
L59:
	;
	goto L60
L60:
	;
	if v161 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v181
	v241 = v181
	goto L52
L62:
	;
	goto L63
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v185 = int32(0)
	if v185 < v184 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v188 = v184
	goto L66
L65:
	;
	v188 = v185
	goto L66
L66:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v193 = int32(0)
	goto L67
L67:
	;
	if v193 < v189 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v205 = v201 + v193<<(uint(int32(2))%32)
	goto L71
L70:
	;
	v205 = int32(0)
	goto L71
L71:
	;
	if v193 == v188 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v188
	v241 = base.B2i32(v205 == int32(0))
	goto L52
L73:
	;
	goto L74
L74:
	;
	v211 = base.B2i32(v205 == int32(0))
	if v205 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = v211
	goto L52
L76:
	;
	goto L77
L77:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v218 = v215 + v193<<(uint(int32(2))%32)
	if v218 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = v211
	goto L52
L79:
	;
	goto L80
L80:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if v222 != v223 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = int32(0)
	goto L52
L82:
	;
	v193 = v193 + int32(1)
	goto L67
L84:
	;
	v262 = v131
	goto L51
L85:
	;
	goto L86
L86:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v131 != v20 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v258 = F_create_incremental_sort_path(m, l0, l2, v131, v160, v242, v111)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L99
	}
L88:
	;
	if v242 == int32(0) {
		goto L50
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v242 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	if v247 == int32(0) {
		goto L50
	} else {
		goto L92
	}
L92:
	;
	goto L87
L93:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	if v251&int32(1) != 0 {
		goto L87
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v254 = F_create_sort_path(m, l2, v131, v160, v111)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	if v254 == int32(0) {
		goto L50
	} else {
		goto L98
	}
L98:
	;
	v262 = v254
	goto L51
L99:
	;
	if v258 == int32(0) {
		goto L50
	} else {
		goto L100
	}
L100:
	;
	v262 = v258
	goto L51
L101:
	;
	v267 = int32(0)
	v273 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v287 = F_create_upper_unique_path(m, l2, v262, v286, v40)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L108
	}
L104:
	;
	v275 = int32(0)
	v277 = F_makeConst(m, int32(20), int32(-1), v267, int32(8), v273, v275, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v282 = F_create_limit_path(m, l2, v262, v267, v277, int32(0), int64(0), int64(1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_add_path(m, l2, v282)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	goto L50
L108:
	;
	F_add_path(m, l2, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L50
L110:
	;
	goto L49
L111:
	;
	goto L43
L112:
	;
	m.G0 = v18 + int32(16)
	return l2
L113:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v332 != 0 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v338 = int32(0)
	if v337 == v338 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, _consts[395])))
	if v334 != int32(1) {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if v375 == int32(0) {
		goto L112
	} else {
		goto L131
	}
L119:
	;
	v375 = int32(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v345 <= int32(0) {
		v370 = int32(1)
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v375 = v370
	goto L118
L123:
	;
	v348 = int32(0)
	if v348 < v345 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v351 = v345
	goto L126
L125:
	;
	v351 = v348
	goto L126
L126:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v355 = v338
	goto L127
L127:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v352+v355<<(uint(int32(2))%32))))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+18)))
	if v361 != int32(1) {
		v370 = v361
		goto L122
	} else {
		goto L129
	}
L128:
	;
	v370 = v361
	goto L122
L129:
	;
	v365 = v355 + int32(1)
	if v365 != v351 {
		v355 = v365
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v380 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v384 = F_create_agg_path(m, l0, l2, v20, v378, int32(2), v380, v381, v380, v380, v40)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	F_add_path(m, l2, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	goto L112
}
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v5 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v13 != int32(331) {
		v20 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 == int32(0) {
		v79 = v5
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v16 != 0 {
		v20 = l2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = l2
	goto L6
L5:
	;
	v19 = int32(0)
	goto L6
L6:
	;
	v20 = v19
	goto L1
L7:
	;
	v86 = F_palloc0(m, int32(80))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L25
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v26 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v79 = v5
	goto L7
L10:
	;
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v36 = int32(1)
	v38 = v5
	v40 = v5
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v79 = v67
	goto L7
L14:
	;
	v50 = F_replace_nestloop_params_mutator(m, v48, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v54 = v48
	goto L16
L16:
	;
	v56 = int32(0)
	v58 = F_makeTargetEntry(m, v54, base.I32_extend16_s(v36), v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	v54 = v50
	goto L16
L19:
	;
	if v29 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v29-int32(4)+v36<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v63
	goto L22
L21:
	;
	goto L22
L22:
	;
	v67 = F_lappend(m, v38, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v70 = v40 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v70 < v71 {
		v36 = v36 + int32(1)
		v38 = v67
		v40 = v70
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+72)) = l3
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+44)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(331)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v97
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+8)) = v99
	v101 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+16)) = v101
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+24)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+36)) = uint8(v89)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v105
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+37)) = uint8(v109)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+37)) = uint8(v111)
	return v86
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 float64
	_ = v169
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v179 float64
	_ = v179
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 float64
	_ = v201
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
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
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v37 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v41 = v40
	goto L5
L4:
	;
	v41 = v33
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
	v231 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = base.F64_add(v231, v232)
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	v236 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v235, base.F64_add(base.F64_mul(v236, v237), v239))
	m.G0 = v18 + int32(144)
	return v21
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v83 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(1)
	v98 = v33
	v99 = v86
	v100 = v86
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v98<<(uint(int32(2))%32))))
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
	if v100&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v213 = v98 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v213 < v214 {
		v98 = v213
		v99 = v209
		v100 = int32(0)
		goto L27
	} else {
		goto L47
	}
L33:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	F_cost_agg(m, v21, l0, v71, l6, v113, v116, l3, v117, v118, v119, v120, base.F64_convert_i32_s(v122))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	if (v128|v99)&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v209 = v126 & v99
	goto L32
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v197 + v198
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v201, v202)
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = base.F64_add(v205, v206)
	v209 = v196
	goto L32
L38:
	;
	v133 = int32(1)
	if v128&v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v149 = int32(0)
	v151 = v18 + int32(72)
	v153 = float64(0)
	v154 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	v159 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v162 = m.G0
	v163 = int32(16)
	v164 = v162 - v163
	m.G0 = v164
	F_cost_tuplesort(m, v164+int32(8), v164, v154, v156, v153, v159, float64(-1))
	mBase = m.M
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v164)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v151)+32)) = v154
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _consts[400])))
	v173 = base.F64_add(v153, v169)
	*(*float64)(unsafe.Add(mBase, uint32(v151)+48)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v151)+40)) = v149 + (v172 ^ int32(1))
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v164)))
	*(*float64)(unsafe.Add(mBase, uint32(v151)+56)) = base.F64_add(v173, v179)
	m.G0 = v164 + v163
	goto L45
L41:
	;
	v136 = int32(2)
	goto L43
L42:
	;
	v136 = v133
	goto L43
L43:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v139 = float64(0)
	v141 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	F_cost_agg(m, v18, l0, v136, l6, v113, v137, l3, int32(0), v139, v139, v141, base.F64_convert_i32_s(v143))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v196 = v147 & v99
	goto L37
L45:
	;
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v18)+120))
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v18)+128))
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v18)+104))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+32))
	F_cost_agg(m, v18, l0, int32(1), l6, v113, v186, l3, v187, v188, v189, v190, base.F64_convert_i32_s(v192))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v196 = v149
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v240 float64
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v253 int32
	_ = v253
	var v256 float64
	_ = v256
	var v258 int32
	_ = v258
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v270 float64
	_ = v270
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v281 float64
	_ = v281
	var v285 float64
	_ = v285
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v306 float64
	_ = v306
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v459 float64
	_ = v459
	var v463 float64
	_ = v463
	var v493 int32
	_ = v493
	var v494 float64
	_ = v494
	var v497 float64
	_ = v497
	var v501 float64
	_ = v501
	var v503 float64
	_ = v503
	var v506 float64
	_ = v506
	var v533 float64
	_ = v533
	var v535 float64
	_ = v535
	var v539 int32
	_ = v539
	var v540 int64
	_ = v540
	var v545 float64
	_ = v545
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 float64
	_ = v592
	var v593 float64
	_ = v593
	var v611 float64
	_ = v611
	var v619 float64
	_ = v619
	var v621 float64
	_ = v621
	var v622 int32
	_ = v622
	var v623 float64
	_ = v623
	var v625 float64
	_ = v625
	var v626 float64
	_ = v626
	var v628 float64
	_ = v628
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
	if v41 == v46 {
		v87 = v46
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v87 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	goto L9
L11:
	;
	if v45 == int32(0) {
		v87 = v46
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v55 < v56 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v55
	goto L15
L14:
	;
	v58 = v56
	goto L15
L15:
	;
	if v58 <= int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v61 = int32(1)
	goto L18
L17:
	;
	v61 = v58
	goto L18
L18:
	;
	v62 = int32(8)
	v67 = int32(0)
	goto L19
L19:
	;
	v74 = v67 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v45+v62+v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v41+v62))))
	v79 = v76 & v78
	v81 = base.B2i32(v79 != int32(0))
	if v79 != 0 {
		v87 = v81
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v87 = v81
	goto L10
L21:
	;
	v83 = v67 + int32(1)
	if v83 != v61 {
		v67 = v83
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v91 = F_get_param_path_clause_serials(m, l6)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(356)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v196
	v198 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v202 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v199, l9, v28+int32(12))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L41
	}
L26:
	;
	if l7 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v154
	goto L25
L28:
	;
	v111 = v93
	v112 = int32(0)
	goto L33
L29:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v93 < v94 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v154 = int32(0)
	goto L27
L32:
	;
	goto L31
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v111<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+56))
	v131 = F_bms_is_member(m, v130, v91)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v154 = v137
	goto L27
L35:
	;
	if v131 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v135 = F_lappend(m, v112, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v137 = v112
	goto L38
L38:
	;
	v139 = v111 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v139 < v140 {
		v111 = v139
		v112 = v137
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v137 = v135
	goto L38
L40:
	;
	goto L34
L41:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+20)) = uint8(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v202
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v207 != int32(1) {
		v214 = v198
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v216 = v214 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)) = uint8(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v218
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v222)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v226
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	v232 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v238
	v240 = float64(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v244 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v210 != int32(1) {
		v214 = v198
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v214 = v213
	goto L42
L45:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v250)))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if int32(0) < v253 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v250 = v244 + int32(8)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v250 = v247 + int32(16)
	goto L45
L49:
	;
	v256 = base.F64_convert_i32_u(v253)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _consts[342])))
	if v258 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.F64_le(v237, v240) != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v264 = base.F64_add(base.F64_mul(v256, float64(-0.3)), float64(1))
	if base.F64_gt(v264, float64(0)) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v270 = v256
	goto L54
L54:
	;
	v272 = float64(1e+100)
	v273 = base.F64_div(v251, v270)
	if base.F64_gt(v273, v272) != 0 {
		v285 = v272
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v268 = v264
	goto L57
L56:
	;
	v268 = math.Float64frombits(uint64(0x8000000000000000))
	goto L57
L57:
	;
	v270 = base.F64_add(v268, v256)
	goto L54
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v285
	goto L51
L59:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v273)&int64(9223372036854775807)) {
		v285 = v272
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v281 = float64(1)
	if base.F64_le(v273, v281) != 0 {
		v285 = v281
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v285 = base.F64_nearest(v273)
	goto L58
L62:
	;
	v290 = float64(1)
	goto L64
L63:
	;
	v290 = v237
	goto L64
L64:
	;
	if base.F64_le(v235, v240) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v292 = float64(1)
	goto L67
L66:
	;
	v292 = v235
	goto L67
L67:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v293&int32(-2) != int32(4) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v540 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+24)) = v540
	*(*int64)(unsafe.Add(mBase, uint32(v230)+16)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = l0
	v545 = float64(0)
	if v539 == int32(0) {
		v611 = v545
		v619 = v545
		goto L126
	} else {
		goto L127
	}
L69:
	;
	v533 = v232
	v535 = base.F64_mul(v292, v290)
	goto L68
L70:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v298 != int32(1) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v301 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v303 = base.F64_nearest(base.F64_mul(v292, v301))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v309 = base.F64_div(float64(2), base.F64_add(v306, float64(1)))
	v310 = base.F64_mul(base.F64_mul(v290, v303), v309)
	v311 = base.F64_sub(v292, v303)
	v312 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	v313 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	if v314 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v493 = base.F64_ge(v311, float64(1))
	if v493 != 0 {
		goto L116
	} else {
		goto L117
	}
L75:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v315 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	switch v320 - int32(341) {
	case 0, 1:
		v327 = l6
		goto L77
	default:
		goto L74
	case 3:
		goto L78
	}
L77:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	if v328 == int32(0) {
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l6)+72))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v324 != int32(280) {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v327 = v323
	goto L77
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v331 <= int32(0) {
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v327)+76))
	v335 = int32(0)
	v340 = v335
	v344 = v335
	goto L83
L82:
	;
	v459 = base.F64_add(base.F64_mul(v313, v309), v232)
	if base.F64_gt(v303, float64(1)) != 0 {
		goto L113
	} else {
		goto L114
	}
L83:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v340<<(uint(int32(2))%32))))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v369 = int32(0)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366)+28))
	v371 = F_bms_is_subset(m, v370, v319)
	mBase = m.M
	if v371 == v369 {
		v382 = v369
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v344 == int32(0) {
		goto L74
	} else {
		goto L112
	}
L85:
	;
	if v382 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+28))
	v375 = F_bms_overlap(m, v368, v374)
	mBase = m.M
	if v375 == int32(0) {
		v382 = v369
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v366)+40))
	v379 = F_bms_overlap(m, v368, v378)
	mBase = m.M
	v382 = v379 ^ int32(1)
	goto L86
L89:
	;
	if v334 != 0 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	goto L91
L91:
	;
	v445 = v340 + int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v445 < v446 {
		v340 = v445
		goto L83
	} else {
		goto L111
	}
L92:
	;
	if v433 == int32(0) {
		goto L74
	} else {
		goto L109
	}
L93:
	;
	goto L92
L94:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v388 <= int32(0) {
		v433 = int32(0)
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v433 = int32(0)
	goto L93
L97:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v366)+60))
	v392 = int32(0)
	if v392 < v388 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v395 = v388
	goto L100
L99:
	;
	v395 = v392
	goto L100
L100:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v399 = int32(0)
	goto L101
L101:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v396+v399<<(uint(int32(2))%32))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)))
	if v409 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	v420 = v399 + int32(1)
	if v420 != v395 {
		v399 = v420
		goto L101
	} else {
		goto L108
	}
L104:
	;
	v410 = int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v366 == v411 {
		v433 = v410
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if v391 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411)+60))
	if v415 == v391 {
		v433 = v410
		goto L93
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	goto L102
L109:
	;
	v439 = int32(1)
	v441 = v340 + v439
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v441 < v442 {
		v340 = v441
		v344 = v439
		goto L83
	} else {
		goto L110
	}
L110:
	;
	goto L82
L111:
	;
	goto L84
L112:
	;
	goto L82
L113:
	;
	v463 = base.F64_add(base.F64_mul(base.F64_mul(v312, base.F64_add(v303, float64(-1))), v309), v459)
	goto L115
L114:
	;
	v463 = v459
	goto L115
L115:
	;
	v533 = base.F64_add(base.F64_div(base.F64_mul(v312, v311), v290), v463)
	v535 = v310
	goto L68
L116:
	;
	v494 = v303
	goto L118
L117:
	;
	v494 = base.F64_add(v303, float64(-1))
	goto L118
L118:
	;
	v497 = base.F64_add(v232, v313)
	if base.F64_gt(v494, float64(0)) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v501 = base.F64_add(base.F64_mul(base.F64_mul(v312, v494), v309), v497)
	goto L121
L120:
	;
	v501 = v497
	goto L121
L121:
	;
	v503 = base.F64_add(base.F64_mul(v311, v290), v310)
	if v493 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v506 = base.F64_add(v311, float64(-1))
	goto L124
L123:
	;
	v506 = v311
	goto L124
L124:
	;
	if base.F64_gt(v506, float64(0)) == int32(0) {
		v533 = v501
		v535 = v503
		goto L68
	} else {
		goto L125
	}
L125:
	;
	v533 = base.F64_add(base.F64_mul(v506, v312), v501)
	v535 = v503
	goto L68
L126:
	;
	v621 = *(*float64)(unsafe.Add(mBase, _consts[388]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v623 = *(*float64)(unsafe.Add(mBase, uint32(v622)+24))
	v625 = *(*float64)(unsafe.Add(mBase, uint32(v622)+16))
	v626 = base.F64_add(base.F64_add(v233, v619), v625)
	*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = v626
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v32)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+56)) = base.F64_add(v626, base.F64_add(base.F64_mul(v623, v628), base.F64_add(base.F64_mul(base.F64_add(v611, v621), v535), v533)))
	m.G0 = v230 + int32(32)
	m.G0 = v28 + int32(16)
	return v32
L127:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v550 <= int32(0) {
		v611 = v545
		v619 = float64(0)
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v557 = int32(0)
	goto L129
L129:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579+v557<<(uint(int32(2))%32))))
	v586 = F_cost_qual_eval_walker(m, v583, v230+int32(8))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	v592 = *(*float64)(unsafe.Add(mBase, uint32(v230)+24))
	v593 = *(*float64)(unsafe.Add(mBase, uint32(v230)+16))
	v611 = v592
	v619 = v593
	goto L126
L131:
	;
	v589 = v557 + int32(1)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v589 < v590 {
		v557 = v589
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v4
	if v18 <= v4 {
		v98 = v4
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
	v35 = v4
	v36 = v4
	goto L6
L4:
	;
	v61 = v4
	v62 = v4
	goto L5
L5:
	;
	if v24 == int32(0) {
		v98 = v62
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v44 = v34 + v26
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)))
	v52 = v35 + v45 + v47 + v49 + v51
	v53 = int32(4)
	v54 = v34 + v53
	v56 = v36 + v53
	if v56 != v18&int32(2147483644) {
		v34 = v54
		v35 = v52
		v36 = v56
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v61 = v54
	v62 = v52
	goto L5
L8:
	;
	goto L7
L9:
	;
	v76 = v61
	v77 = v62
	v79 = v4
	goto L10
L10:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v26))))
	v88 = v77 + v87
	v89 = int32(1)
	v92 = v79 + v89
	if v92 != v24 {
		v76 = v76 + v89
		v77 = v88
		v79 = v92
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v98 = v88
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
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v122 = int32(8)
	v128 = v98<<(uint(v122)%32) | int32(base.Ui32(v98&int32(65280))>>(uint(v122)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v18)+1)) = uint16(v128)
	v131 = l2 - v108
	v133 = v131 - int32(2)
	if v133 < v122 {
		v307 = int32(-12)
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v109+int32(1), l0+int32(132), v18)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v318 = F___memset(m, v109, int32(0), v108)
	mBase = m.M
	goto L68
L20:
	;
	v136 = F_palloc(m, l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v138 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v138)
	v141 = v136 + int32(1)
	v142 = int32(0)
	v146 = m.G0
	v148 = v146 - int32(16)
	m.G0 = v148
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v142
	v154 = F_open(m, int32(288599), v142, v148)
	mBase = m.M
	if v154 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v187 != 0 {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v157 = int32(1)
	if v133 == int32(0) {
		v180 = v157
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v187 = v142
	goto L25
L25:
	;
	m.G0 = v148 + int32(16)
	goto L22
L26:
	;
	v182 = F_close(m, v154)
	mBase = m.M
	v187 = v180
	goto L25
L27:
	;
	v160 = v141
	v161 = v133
	goto L28
L28:
	;
	v166 = F_read(m, v154, v160, v161)
	mBase = m.M
	if v166 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v180 = v157
	goto L26
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v170 == int32(27) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v175 = v161 - v166
	if v175 != 0 {
		v160 = v160 + v166
		v161 = v175
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v180 = int32(0)
	goto L26
L34:
	;
	goto L29
L35:
	;
	v192 = v131 + v136
	v198 = v141
	goto L39
L36:
	;
	goto L37
L37:
	;
	F_pfree(m, v136)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L67
	}
L38:
	;
	v287 = F___memset(m, v136, int32(0), l2)
	mBase = m.M
	goto L66
L39:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v209 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136+v133)+1)) = uint8(v270)
	if v108 != 0 {
		goto L60
	} else {
		goto L61
	}
L41:
	;
	v266 = int32(1)
	goto L43
L42:
	;
	v211 = int32(0)
	v215 = m.G0
	v217 = v215 - int32(16)
	m.G0 = v217
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v211
	v223 = F_open(m, int32(288599), v211, v217)
	mBase = m.M
	if v223 != int32(-1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v267 = v266 + v198
	if base.Ui32(v267) < base.Ui32(v192-int32(1)) {
		v198 = v267
		goto L39
	} else {
		goto L58
	}
L44:
	;
	if v256 == int32(0) {
		goto L38
	} else {
		goto L57
	}
L45:
	;
	goto L49
L46:
	;
	v256 = v211
	goto L47
L47:
	;
	m.G0 = v217 + int32(16)
	goto L44
L48:
	;
	v251 = F_close(m, v223)
	mBase = m.M
	v256 = v249
	goto L47
L49:
	;
	v229 = v198
	v230 = int32(1)
	goto L50
L50:
	;
	v235 = F_read(m, v223, v229, v230)
	mBase = m.M
	if v235 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v249 = int32(1)
	goto L48
L52:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v239 == int32(27) {
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v244 = v230 - v235
	if v244 != 0 {
		v229 = v229 + v235
		v230 = v244
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v249 = int32(0)
	goto L48
L56:
	;
	goto L51
L57:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v266 = base.B2i32(v263 != int32(0))
	goto L43
L58:
	;
	goto L40
L59:
	;
	v280 = F_pgp_mpi_create(m, v136, l2<<(uint(int32(3))%32)-int32(6), v16+int32(12))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L63
	}
L60:
	;
	v272 = F__emscripten_memcpy_bulkmem(m, v192, v109, v108)
	mBase = m.M
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v283 = F___memset(m, v136, int32(0), l2)
	mBase = m.M
	goto L64
L64:
	;
	F_pfree(m, v136)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v307 = v280
	goto L19
L66:
	;
	goto L37
L67:
	;
	v307 = int32(-17)
	goto L19
L68:
	;
	F_pfree(m, v109)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	if int32(0) <= v307 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
	goto L72
L71:
	;
	goto L72
L72:
	;
	m.G0 = v16 + int32(16)
	return v307
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
	var v16 int32
	_ = v16
	v6 = l1 + int32(4)
	v7 = F_palloc(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v6 << (uint(int32(2)) % 32)
		if l1 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, v7+int32(4), l0, l1)
			mBase = m.M
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
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
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(258))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
											F_errmsg(m, int32(164534), v8+int32(16))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(497496), int32(3111), int32(506030))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
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
												v41 = F__emscripten_memcpy_bulkmem(m, v39, v34, v38)
												mBase = m.M
												v42 = v41
											} else {
												v42 = v39
											}
											v43 = F_SPI_finish(m)
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												v45 = F_cstring_to_text(m, v42)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(32)
													return v45
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(259))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
										F_errmsg(m, int32(70860), v8)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(497496), int32(3107), int32(506030))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(556366))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = F_NameListToString(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
				F_errmsg(m, int32(396776), v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = F_plpgsql_scanner_errposition(m, l1, l2)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(26975), int32(2646), int32(396155))
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
}
