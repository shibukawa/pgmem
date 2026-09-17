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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
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
	v39 = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_3)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[1]))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+280)) = v41
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[2])) = v41
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[1]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+272)) = v47
	*(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[3])) = v47
	goto L13
L10:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L73
	}
L11:
	;
	v57 = m.G0
	v59 = v57 - int32(224)
	m.G0 = v59
	v63 = v2
	v64 = int32(-1)
	v66 = v2
	v67 = v2
	goto L15
L13:
	;
	goto L14
L14:
	;
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[2]))
	goto L11
L15:
	;
	if v64 != int32(1) {
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
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4]))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5]))
	v78 = v59 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v59 + int32(12)
	goto L20
L18:
	;
	v84 = v63
	v85 = v66
	v86 = v67
	goto L19
L19:
	;
	goto L22
L20:
	;
	v84 = int32(0)
	v85 = v74
	v86 = v76
	goto L19
L21:
	;
	goto L16
L22:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L21
L24:
	;
	v227 = int32(m.ExcTag)
	v228 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v227 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(396)
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v59 + int32(32)
	v100 = int32(0)
	v107 = F_CreateDecodingContext(m, int64(0), v100, int32(1), v59+int32(20), v100, v100, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v86
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L24
	} else {
		goto L61
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[6]))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+104))
	F_XLogBeginRead(m, v109, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L31
L31:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v125)+40))
	if base.Ui64(v126) < base.Ui64(v56) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	if v161 != 0 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = int32(0)
	v132 = F_XLogReadRecord(m, v125, v59+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v134 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v132 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v139
	F_errmsg_internal(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_4), v59)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_logical_slot_has_caught_up_5), int32(2042), int32(_a_F_binary_upgrade_logical_slot_has_caught_up_6))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	F_LogicalDecodingProcessRecord(m, v107, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L24
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+165)))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[7]))
	if v154 != 0 {
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
	v156 = m.ExcPending
	if v156 != 0 {
		goto L24
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v152&int32(1) == int32(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v59)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+212)) = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_7)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+208)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v59)+200)) = int32(993)
	v169 = int32(_a_F_binary_upgrade_logical_slot_has_caught_up_8)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v59 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+196)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v59)+204)) = v59 + int32(208)
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+164)) = uint8(v179)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+147)) = uint8(v179)
	m.T0[v161].(func(*base.Module, int32))(m, v107)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L24
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	F_ReorderBufferFree(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L24
	} else {
		goto L56
	}
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v59)+196))
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v186
	goto L54
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	F_FreeSnapshotBuilder(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	F_XLogReaderFree(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	F_MemoryContextDelete(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[4])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_logical_slot_has_caught_up[5])) = v86
	m.G0 = v59 + int32(224)
	goto L10
L61:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	goto L23
L63:
	;
	v232 = int32(v228)
	m.G0 = v59
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v59+int32(12) == v238 {
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
	if v242 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v242 = v240
	goto L68
L67:
	;
	v242 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	F___wasm_longjmp(m, v235, v234)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v63 = v234
	v64 = v242
	v66 = v85
	v67 = v86
	goto L15
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	return base.B2i32(base.Ui64(v126) < base.Ui64(v56)) ^ int32(1)
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
	v5 = Fn13849(m, l0, int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_0), int32(_a_F_binary_upgrade_set_next_toast_pg_class_oid_1), int32(145))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
