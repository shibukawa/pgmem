package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsBinaryTidClause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == v3 {
		v85 = v3
		return v85
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11 != int32(17) {
			v85 = v3
			return v85
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			if v14 == int32(0) {
				v85 = v3
				return v85
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v17 != int32(2) {
					v85 = v3
					return v85
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
					v22 = int32(0)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v23 == v22 {
						v47 = v3
						v48 = v22
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						if v27 != int32(6) {
							v47 = v3
							v48 = int32(0)
						} else {
							v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)))
							if v31 != int32(_a_F_IsBinaryTidClause_0) {
								v47 = v3
								v48 = int32(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								if v35 != int32(27) {
									v47 = v3
									v48 = int32(0)
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
									if v39 != v40 {
										v47 = v3
										v48 = int32(0)
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
										if v43 != 0 {
											v47 = v3
											v48 = int32(0)
										} else {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
											if v45 != 0 {
												v47 = v3
												v48 = int32(0)
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v47 = v21
												v48 = v46
											}
										}
									}
								}
							}
						}
					}
					v49 = int32(0)
					if v47|base.B2i32(v21 == v49) == v49 {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v54 != int32(6) {
							v85 = v3
							return v85
						} else {
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+8)))
							if v57 != int32(_a_F_IsBinaryTidClause_0) {
								v85 = v3
								return v85
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
								if v60 != int32(27) {
									v85 = v3
									return v85
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
									if v63 != v64 {
										v85 = v3
										return v85
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
										if v66 != 0 {
											v85 = v3
											return v85
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
											if v67 != 0 {
												v85 = v3
												return v85
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												v69 = v23
												v70 = v68
												if v69 == int32(0) {
													v85 = v3
													return v85
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
													v74 = F_bms_is_member(m, v73, v70)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														if v74 != 0 {
															v85 = v3
															return v85
														} else {
															v78 = F_contain_volatile_functions(m, v69)
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return int32(0)
															} else {
																v85 = v78 ^ int32(1)
																return v85
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
						v69 = v47
						v70 = v48
						if v69 == int32(0) {
							v85 = v3
							return v85
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
							v74 = F_bms_is_member(m, v73, v70)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								if v74 != 0 {
									v85 = v3
									return v85
								} else {
									v78 = F_contain_volatile_functions(m, v69)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v85 = v78 ^ int32(1)
										return v85
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
func F_binary_upgrade_check_logical_slot_pending_wal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v2 = int32(0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[0])))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = int32(1)
	F_ReplicationSlotAcquire(m, v14, v15, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L76
	}
L4:
	;
	return int64(0)
L5:
	;
	v21 = int32(0)
	v23 = int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[1]))
	v25 = int64(0)
	v28 = base.AtomicRmwCmpxchg64(m, v24, int32(272), v25, v25)
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[2])) = v28
	v33 = base.AtomicRmwOr32(m, v21, int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_1), v21)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[1]))
	v40 = base.AtomicRmwCmpxchg64(m, v36, int32(264), v25, v25)
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[3])) = v40
	goto L9
L6:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L72
	}
L7:
	;
	v48 = m.G0
	v50 = v48 - int32(224)
	m.G0 = v50
	v55 = v2
	v56 = int32(-1)
	v58 = v2
	v59 = v2
	goto L11
L9:
	;
	goto L10
L10:
	;
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[2]))
	goto L7
L11:
	;
	if v56 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4]))
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[5]))
	v71 = v50 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v50 + int32(12)
	goto L16
L14:
	;
	v77 = v55
	v78 = v58
	v79 = v59
	goto L15
L15:
	;
	goto L18
L16:
	;
	v77 = int32(0)
	v78 = v67
	v79 = v69
	goto L15
L17:
	;
	goto L12
L18:
	;
	if v77 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L17
L20:
	;
	v241 = int32(m.ExcTag)
	v242 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v241 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = int32(416)
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[5])) = v50 + int32(32)
	v93 = int32(0)
	v100 = F_CreateDecodingContext(m, int64(0), v93, int32(1), v50+int32(20), v93, v93, v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4])) = v78
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[5])) = v79
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L20
	} else {
		goto L60
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[6]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+104))
	F_XLogBeginRead(m, v102, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v110 = int64(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+40))
	if base.Ui64(v47) <= base.Ui64(v112) {
		v172 = v110
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	if v174 != 0 {
		goto L51
	} else {
		goto L52
	}
L28:
	;
	v117 = v111
	v122 = v110
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	v128 = F_XLogReadRecord(m, v117, v50+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L31
	}
L30:
	;
	v172 = v156
	goto L27
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v130 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v128 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v135
	F_errmsg_internal(m, int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_2), v50)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_3), int32(2137), int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_4))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	goto L17
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	F_LogicalDecodingProcessRecord(m, v100, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+165)))
	if v148 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v151)+32))
	if base.Ui64(v13) <= base.Ui64(v152) {
		v172 = v152
		goto L27
	} else {
		goto L45
	}
L43:
	;
	v156 = v122
	goto L44
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[7]))
	if v158 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+165)) = uint8(v154)
	v156 = v152
	goto L44
L46:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L20
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v161)+40))
	if base.Ui64(v162) < base.Ui64(v47) {
		v117 = v161
		v122 = v156
		goto L29
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L30
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+212)) = int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_5)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+208)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v50)+200)) = int32(1058)
	v182 = int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_6)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4])) = v50 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+196)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v50)+204)) = v50 + int32(208)
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+164)) = uint8(v192)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+147)) = uint8(v192)
	m.T0[v174].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L20
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	F_ReorderBufferFree(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L20
	} else {
		goto L55
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v50)+196))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4])) = v199
	goto L53
L55:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	F_FreeSnapshotBuilder(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	F_XLogReaderFree(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	F_MemoryContextDelete(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[4])) = v78
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_check_logical_slot_pending_wal[5])) = v79
	m.G0 = v50 + int32(224)
	goto L6
L60:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	goto L19
L62:
	;
	v246 = int32(v242)
	m.G0 = v50
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v50+int32(12) == v252 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	m.ExcPending = 1
	goto L4
L64:
	;
	if v256 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	v256 = v254
	goto L67
L66:
	;
	v256 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	F___wasm_longjmp(m, v249, v248)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v55 = v248
	v56 = v256
	v58 = v78
	v59 = v79
	goto L11
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	if v172 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v276 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v276)
	goto L75
L74:
	;
	goto L75
L75:
	;
	return v172
L76:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_7), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_8), int32(292), int32(_a_F_binary_upgrade_check_logical_slot_pending_wal_9))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_binary_upgrade_replorigin_advance(m *base.Module, l0 int32) int64 {
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
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_replorigin_advance[0])))
	if v10 != 0 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v11 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_binary_upgrade_replorigin_advance_0), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_replorigin_advance_1), int32(391), int32(_a_F_binary_upgrade_replorigin_advance_2))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v15 = F_pg_detoast_datum_packed(m, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = F_text_to_cstring(m, v15)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v21 == int32(0) {
						v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v25 = v24
					} else {
						v25 = int64(0)
					}
					v28 = F_table_open(m, int32(_a_F_binary_upgrade_replorigin_advance_3), int32(3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v31 = F_get_subscription_oid(m, v19, int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_ReplicationOriginNameForLogicalRep(m, v31, int32(0), v7)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								F_LockRelationOid(m, int32(_a_F_binary_upgrade_replorigin_advance_4), int32(3))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									v41 = F_replorigin_by_name(m, v7, int32(0))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int64(0)
									} else {
										v44 = int32(0)
										F_replorigin_advance(m, v41, v25, int64(0), v44, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int64(0)
										} else {
											F_UnlockRelationOid(m, int32(_a_F_binary_upgrade_replorigin_advance_4), int32(3))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int64(0)
											} else {
												F_relation_close(m, v28, int32(3))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int64(0)
												} else {
													m.G0 = v7 - int32(-64)
													return int64(0)
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_replorigin_advance_5), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_replorigin_advance_1), int32(384), int32(_a_F_binary_upgrade_replorigin_advance_2))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
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
func F_binary_upgrade_set_next_toast_pg_class_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_1), int32(146))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
