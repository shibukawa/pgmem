package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSetNamespace(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v5 = F_get_namespace_name(m, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L81
	}
L3:
	;
	F_pfree(m, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L80
	}
L4:
	;
	return
L5:
	;
	if v5 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v7 = int32(495952)
	goto L11
L7:
	;
	goto L8
L8:
	;
	v106 = F_get_namespace_name(m, l0)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L40
	}
L9:
	;
	if v44-v45 == int32(0) {
		v228 = v5
		goto L3
	} else {
		goto L23
	}
L11:
	;
	goto L12
L12:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v14 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v15 = v5
	v16 = v7
	v17 = int32(8)
	v18 = v14
	goto L17
L14:
	;
	v40 = v7
	v44 = int32(0)
	goto L15
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	goto L9
L16:
	;
	v40 = v35
	v44 = v37
	goto L15
L17:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v18 != v20 {
		v35 = v16
		v37 = v18
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v35 = v29
	v37 = int32(0)
	goto L16
L19:
	;
	if v20 == int32(0) {
		v35 = v16
		v37 = v18
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v25 = v17 - int32(1)
	if v25 == int32(0) {
		v35 = v16
		v37 = v18
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v28 = int32(1)
	v29 = v16 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v30 != 0 {
		v15 = v15 + v28
		v16 = v29
		v17 = v25
		v18 = v30
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v55 = int32(495937)
	goto L26
L24:
	;
	F_pfree(m, v5)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L38
	}
L26:
	;
	goto L27
L27:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v62 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v63 = v5
	v64 = v55
	v65 = int32(14)
	v66 = v62
	goto L32
L29:
	;
	v88 = v55
	v92 = int32(0)
	goto L30
L30:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	goto L24
L31:
	;
	v88 = v83
	v92 = v85
	goto L30
L32:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v66 != v68 {
		v83 = v64
		v85 = v66
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v83 = v77
	v85 = int32(0)
	goto L31
L34:
	;
	if v68 == int32(0) {
		v83 = v64
		v85 = v66
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v73 = v65 - int32(1)
	if v73 == int32(0) {
		v83 = v64
		v85 = v66
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v76 = int32(1)
	v77 = v64 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v78 != 0 {
		v63 = v63 + v76
		v64 = v77
		v65 = v73
		v66 = v78
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if v92-v93 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	if v106 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v108 = int32(495952)
	goto L46
L42:
	;
	goto L43
L43:
	;
	v207 = int32(99)
	if base.B2i32(l0 != v207)&base.B2i32(l1 != v207) != 0 {
		goto L1
	} else {
		goto L75
	}
L44:
	;
	if v145-v146 == int32(0) {
		v228 = v106
		goto L3
	} else {
		goto L58
	}
L46:
	;
	goto L47
L47:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v115 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v116 = v106
	v117 = v108
	v118 = int32(8)
	v119 = v115
	goto L52
L49:
	;
	v141 = v108
	v145 = int32(0)
	goto L50
L50:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	goto L44
L51:
	;
	v141 = v136
	v145 = v138
	goto L50
L52:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v119 != v121 {
		v136 = v117
		v138 = v119
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v136 = v130
	v138 = int32(0)
	goto L51
L54:
	;
	if v121 == int32(0) {
		v136 = v117
		v138 = v119
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v126 = v118 - int32(1)
	if v126 == int32(0) {
		v136 = v117
		v138 = v119
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v129 = int32(1)
	v130 = v117 + v129
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v131 != 0 {
		v116 = v116 + v129
		v117 = v130
		v118 = v126
		v119 = v131
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v156 = int32(495937)
	goto L61
L59:
	;
	F_pfree(m, v106)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L73
	}
L61:
	;
	goto L62
L62:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v163 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v164 = v106
	v165 = v156
	v166 = int32(14)
	v167 = v163
	goto L67
L64:
	;
	v189 = v156
	v193 = int32(0)
	goto L65
L65:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	goto L59
L66:
	;
	v189 = v184
	v193 = v186
	goto L65
L67:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v167 != v169 {
		v184 = v165
		v186 = v167
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v184 = v178
	v186 = int32(0)
	goto L66
L69:
	;
	if v169 == int32(0) {
		v184 = v165
		v186 = v167
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v174 = v166 - int32(1)
	if v174 == int32(0) {
		v184 = v165
		v186 = v167
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v177 = int32(1)
	v178 = v165 + v177
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	if v179 != 0 {
		v164 = v164 + v177
		v165 = v178
		v166 = v174
		v167 = v179
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	if v193-v194 == int32(0) {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	goto L43
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(495096), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(488296), int32(3474), int32(409258))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L2
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(170532), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(488296), int32(3468), int32(409258))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if base.B2i32(v5 != int32(0))&base.B2i32(l0 == v5) != 0 {
		v11 = F_pstrdup(m, int32(230553))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		v17 = F_SearchSysCache1(m, int32(38), l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v28 = F_pstrdup(m, v23+v24+int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v17)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v28
					}
				}
			}
		}
	}
}
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = F_GetSysCacheOid(m, int32(37), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			m.G0 = v7 + int32(16)
			return v13
		} else {
			if v13 != 0 {
				m.G0 = v7 + int32(16)
				return v13
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1411))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(71148), v7)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488296), int32(3547), int32(425348))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
func F_report_namespace_conflict(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 <= int32(3599) {
		if l0 != int32(2607) {
			if l0 != int32(3381) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(56802), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(484126), int32(144), int32(106551))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = int32(694743)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(290948))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = F_get_namespace_name(m, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
							F_errmsg(m, v25, v7+int32(16))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errfinish(m, int32(484126), int32(150), int32(106551))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
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
		} else {
			v25 = int32(695000)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(484126), int32(150), int32(106551))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
	} else {
		switch l0 - int32(3600) {
		case 0:
			v25 = int32(694685)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(484126), int32(150), int32(106551))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		case 1:
			v25 = int32(694796)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(484126), int32(150), int32(106551))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		case 2:
			v25 = int32(694850)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(484126), int32(150), int32(106551))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
			if l0 != int32(3764) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(56802), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(484126), int32(144), int32(106551))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = int32(695109)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(290948))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = F_get_namespace_name(m, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
							F_errmsg(m, v25, v7+int32(16))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errfinish(m, int32(484126), int32(150), int32(106551))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
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
	}
}
