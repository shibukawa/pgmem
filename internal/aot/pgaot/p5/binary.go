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
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
func F_binary_upgrade_logical_slot_has_caught_up(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	v2 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[0])))
	if v10 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = int32(1)
	F_ReplicationSlotAcquire(m, v31, v32, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_0), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_1), int32(290), int32(_a_F_binary_upgrade_logical_slot_has_caught_up_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	v38 = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_3)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[1]))
	v40 = int64(0)
	v43 = base.AtomicRmwCmpxchg64(m, v39, int32(280), v40, v40)
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[2])) = v43
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[1]))
	v51 = base.AtomicRmwCmpxchg64(m, v47, int32(272), v40, v40)
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[3])) = v51
	goto L13
L10:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L73
	}
L11:
	;
	v59 = m.G0
	v61 = v59 - int32(224)
	m.G0 = v61
	v65 = v2
	v66 = int32(-1)
	v68 = v2
	v69 = v2
	goto L15
L13:
	;
	goto L14
L14:
	;
	v58 = *(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[2]))
	goto L11
L15:
	;
	if v66 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4]))
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5]))
	v80 = v61 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v61 + int32(12)
	goto L20
L18:
	;
	v86 = v65
	v87 = v68
	v88 = v69
	goto L19
L19:
	;
	goto L22
L20:
	;
	v86 = int32(0)
	v87 = v76
	v88 = v78
	goto L19
L21:
	;
	goto L16
L22:
	;
	if v86 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L21
L24:
	;
	v229 = int32(m.ExcTag)
	v230 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v229 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = int32(396)
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v61 + int32(32)
	v102 = int32(0)
	v109 = F_CreateDecodingContext(m, int64(0), v102, int32(1), v61+int32(20), v102, v102, v102)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v87
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v88
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L24
	} else {
		goto L61
	}
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[6]))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)+104))
	F_XLogBeginRead(m, v111, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L31
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)+40))
	if base.Ui64(v128) < base.Ui64(v58) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
	if v163 != 0 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = int32(0)
	v134 = F_XLogReadRecord(m, v127, v61+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L24
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v136 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v134 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v141
	F_errmsg_internal(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_4), v61)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_5), int32(2042), int32(_a_F_binary_upgrade_logical_slot_has_caught_up_6))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	F_LogicalDecodingProcessRecord(m, v109, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L24
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+165)))
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[7]))
	if v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L24
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v154&int32(1) == int32(0) {
		goto L31
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L35
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+212)) = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_7)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+208)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v61)+200)) = int32(993)
	v171 = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_8)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v61 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+196)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v61)+204)) = v61 + int32(208)
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+164)) = uint8(v181)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+147)) = uint8(v181)
	m.T0[v163].(func(*base.Module, int32))(m, v109)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L24
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	F_ReorderBufferFree(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L24
	} else {
		goto L56
	}
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v61)+196))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v188
	goto L54
L56:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	F_FreeSnapshotBuilder(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	F_XLogReaderFree(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	F_MemoryContextDelete(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v87
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v88
	m.G0 = v61 + int32(224)
	goto L10
L61:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	goto L23
L63:
	;
	v234 = int32(v230)
	m.G0 = v61
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v61+int32(12) == v240 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	m.ExcPending = 1
	goto L4
L65:
	;
	if v244 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v244 = v242
	goto L68
L67:
	;
	v244 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	F___wasm_longjmp(m, v237, v236)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v65 = v236
	v66 = v244
	v68 = v87
	v69 = v88
	goto L15
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	return base.B2i32(base.Ui64(v128) < base.Ui64(v58)) ^ int32(1)
}
func F_binary_upgrade_replorigin_advance(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_replorigin_advance[0])))
	if v10 != 0 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v11 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_binary_upgrade_replorigin_advance_0), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_replorigin_advance_1), int32(384), int32(_a_F_binary_upgrade_replorigin_advance_2))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v15 = F_pg_detoast_datum_packed(m, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = F_text_to_cstring(m, v15)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v21 == int32(0) {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
						v26 = v25
					} else {
						v26 = int64(0)
					}
					v29 = F_table_open(m, int32(_a_F_binary_upgrade_replorigin_advance_3), int32(3))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v32 = F_get_subscription_oid(m, v19, int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_ReplicationOriginNameForLogicalRep(m, v32, int32(0), v7)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_LockRelationOid(m, int32(_a_F_binary_upgrade_replorigin_advance_4), int32(3))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v42 = F_replorigin_by_name(m, v7, int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v45 = int32(0)
										F_replorigin_advance(m, v42, v26, int64(0), v45, v45)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											F_UnlockRelationOid(m, int32(_a_F_binary_upgrade_replorigin_advance_4), int32(3))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v29, int32(3))
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 - int32(-64)
													return int32(0)
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
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_replorigin_advance_5), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_replorigin_advance_1), int32(377), int32(_a_F_binary_upgrade_replorigin_advance_2))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
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
func F_binary_upgrade_set_next_toast_pg_class_oid(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13871(m, l0, int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_1), int32(145))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
