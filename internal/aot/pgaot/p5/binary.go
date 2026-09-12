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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v84 int32
	_ = v84
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == v3 {
		v84 = v3
		return v84
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11 != int32(17) {
			v84 = v3
			return v84
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			if v14 == int32(0) {
				v84 = v3
				return v84
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v17 != int32(2) {
					v84 = v3
					return v84
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
							if v31 != int32(65535) {
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
					if v47 != 0 {
						v68 = v47
						v69 = v48
						v70 = int32(0)
						if v68 == v70 {
							v84 = v70
							return v84
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
							v74 = F_bms_is_member(m, v73, v69)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								if v74 != 0 {
									v84 = v70
									return v84
								} else {
									v78 = F_contain_volatile_functions(m, v68)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v84 = v78 ^ int32(1)
										return v84
									}
								}
							}
						}
					} else {
						if v21 == int32(0) {
							v68 = v47
							v69 = v48
							v70 = int32(0)
							if v68 == v70 {
								v84 = v70
								return v84
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
								v74 = F_bms_is_member(m, v73, v69)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									if v74 != 0 {
										v84 = v70
										return v84
									} else {
										v78 = F_contain_volatile_functions(m, v68)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v84 = v78 ^ int32(1)
											return v84
										}
									}
								}
							}
						} else {
							v51 = int32(0)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v52 != int32(6) {
								v84 = v51
								return v84
							} else {
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+8)))
								if v55 != int32(65535) {
									v84 = v51
									return v84
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
									if v58 != int32(27) {
										v84 = v51
										return v84
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
										if v61 != v62 {
											v84 = v51
											return v84
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
											if v64 != 0 {
												v84 = v51
												return v84
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
												if v65 != 0 {
													v84 = v51
													return v84
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													v68 = v23
													v69 = v66
													v70 = int32(0)
													if v68 == v70 {
														v84 = v70
														return v84
													} else {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
														v74 = F_bms_is_member(m, v73, v69)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															if v74 != 0 {
																v84 = v70
																return v84
															} else {
																v78 = F_contain_volatile_functions(m, v68)
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return int32(0)
																} else {
																	v84 = v78 ^ int32(1)
																	return v84
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
	}
}
func F_binary_upgrade_logical_slot_has_caught_up(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v62 int64
	_ = v62
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v297 int32
	_ = v297
	v2 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[80])))
	if v16 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = int32(1)
	F_ReplicationSlotAcquire(m, v37, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(433674), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(517730), int32(290), int32(246411))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v45 = int32(4457632)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+280)) = v47
	*(*int64)(unsafe.Add(mBase, _consts[117])) = v47
	v52 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+272)) = v53
	*(*int64)(unsafe.Add(mBase, _consts[116])) = v53
	goto L13
L10:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L73
	}
L11:
	;
	v64 = m.G0
	v66 = v64 - int32(16)
	m.G0 = v66
	v69 = int32(0)
	v70 = int32(-1)
	v71 = v2
	v72 = v2
	v73 = v2
	v74 = v2
	v76 = v66
	v77 = v2
	v78 = v2
	v79 = v2
	goto L15
L13:
	;
	goto L14
L14:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	goto L11
L15:
	;
	if v70 != int32(1) {
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
	v85 = int32(16)
	v86 = v76 - v85
	m.G0 = v86
	v89 = v86 - v85
	m.G0 = v89
	v92 = v89 - int32(160)
	m.G0 = v92
	v95 = v92 - v85
	m.G0 = v95
	v98 = v95 - v85
	m.G0 = v98
	v102 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v104 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v66 + int32(12)
	goto L20
L18:
	;
	v110 = v69
	v111 = v71
	v112 = v72
	v113 = v73
	v114 = v74
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = v79
	goto L19
L19:
	;
	goto L22
L20:
	;
	v110 = int32(0)
	v111 = v89
	v112 = v86
	v113 = v95
	v114 = v98
	v115 = v98
	v116 = v92
	v117 = v102
	v118 = v104
	goto L19
L21:
	;
	goto L16
L22:
	;
	if v110 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L21
L24:
	;
	v261 = int32(m.ExcTag)
	v262 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v261 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(396)
	v130 = int32(0)
	v135 = F_CreateDecodingContext(m, int64(0), v130, int32(1), v113, v130, v130, v130)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v117
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v118
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L61
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v139 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+104))
	F_XLogBeginRead(m, v137, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L31
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+40))
	if base.Ui64(v160) < base.Ui64(v62) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	if v194 != 0 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(0)
	v164 = F_XLogReadRecord(m, v159, v114)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v166 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v164 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v171
	F_errmsg_internal(m, int32(214071), v66)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(523877), int32(2042), int32(330587))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	F_LogicalDecodingProcessRecord(m, v135, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L24
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+165)))
	v186 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v186 != 0 {
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
	v188 = m.ExcPending
	if v188 != 0 {
		goto L24
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v184&int32(1) == int32(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v112)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = int32(257136)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = int32(993)
	v203 = int32(4555000)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v204
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v111
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+164)) = uint8(v208)
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+147)) = uint8(v208)
	m.T0[v194].(func(*base.Module, int32))(m, v135)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L24
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	F_ReorderBufferFree(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L24
	} else {
		goto L56
	}
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v215
	goto L54
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	F_FreeSnapshotBuilder(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	F_XLogReaderFree(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	F_MemoryContextDelete(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v117
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v118
	m.G0 = v66 + int32(16)
	goto L10
L61:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	goto L23
L63:
	;
	v266 = int32(v262)
	m.G0 = v115
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v66+int32(12) == v273 {
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
	if v276 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v276 = v275
	goto L68
L67:
	;
	v276 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	F___wasm_longjmp(m, v269, v268)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v69 = v268
	v70 = v276
	v71 = v111
	v72 = v112
	v73 = v113
	v74 = v114
	v76 = v115
	v77 = v116
	v78 = v117
	v79 = v118
	goto L15
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	return base.B2i32(base.Ui64(v160) < base.Ui64(v62)) ^ int32(1)
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
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[80])))
	if v10 != 0 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v11 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(461617), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517730), int32(384), int32(438326))
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
					v29 = F_table_open(m, int32(6100), int32(3))
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
								F_LockRelationOid(m, int32(6000), int32(3))
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
											F_UnlockRelationOid(m, int32(6000), int32(3))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v29, int32(3))
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
				F_errmsg(m, int32(433674), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517730), int32(377), int32(438326))
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[80])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(433674), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517730), int32(145), int32(456086))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[257])) = v25
		return int32(0)
	}
}
