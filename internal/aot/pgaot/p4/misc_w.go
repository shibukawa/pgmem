package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WarnNoTransactionBlock(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(int32(1)) < base.Ui32(v11) {
		m.G0 = v7 + int32(16)
		return
	} else {
		if l0 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			if int32(1) < v16 {
				m.G0 = v7 + int32(16)
				return
			} else {
				v21 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					if v21 == int32(0) {
						m.G0 = v7 + int32(16)
						return
					} else {
						F_errcode(m, int32(16908610))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errmsg(m, int32(144758), v7)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(472028), int32(3749), int32(302890))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
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
func F_win866_to_win1251(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(20), int32(23))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(20), int32(23), int32(2177792), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_wordchrs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v6 = F_newstate(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v10 | int32(1024)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v18 = F_cclasscvec(m, l0, int32(13), v15&int32(8))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v20 != 0 {
						return
					} else {
						F_subcolorcvec(m, l0, v18, v6, v6)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v23 != 0 {
								return
							} else {
								v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								F_okcolors(m, v24, v25)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v28 != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v6
									}
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
func F_wrapper_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v4 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	if v6 != int32(42) {
		v9 = int32(0)
		v11 = m.G0
		v13 = v11 - int32(144)
		m.G0 = v13
		switch int32(2) {
		case 0, 2:
			v23 = v9
		default:
			*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1455]))) = v9
			v23 = int32(4729)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
		F_sigemptyset(m, v13+int32(8))
		mBase = m.M
		if l0 == int32(17) {
			v32 = int32(268435457)
		} else {
			v32 = int32(268435456)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v32
		v35 = v13 + int32(4)
		if base.Ui32(int32(65)) <= base.Ui32(l0) {
			*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
		} else {
			if v35 != 0 {
				v41 = int32(140)
				v46 = F___memcpy(m, l0*v41+int32(4608768), v35, v41)
				mBase = m.M
			} else {
			}
		}
		m.G0 = v13 + int32(144)
		v50 = F_raise(m, l0)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			return
		}
	} else {
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1455])))
		m.T0[v56].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
			return
		}
	}
}
func F_write_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v9 = m.G0
	v11 = v9 - int32(2160)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5842711)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v15) < base.Ui32(int32(65)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L49
	}
L3:
	;
	v18 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v18
	v22 = m.Env.Pgmem_crc32c(m, v18, l0, int32(520))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v22 ^ v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(227165)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l6
	v35 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(167509), v11+int32(80))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L46
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(224448)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l6
	v46 = F_pg_snprintf(m, v11+int32(112), int32(1024), int32(167509), v11-int32(-64))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v51 = F_OpenTransientFile(m, v11+int32(112), int32(577))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v51 < int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(167772202)
	v59 = int32(524)
	v60 = F_write(m, v51, l0, v59)
	mBase = m.M
	if v60 != v59 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
	v93 = F_CloseTransientFile(m, v51)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L21
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(51)
	goto L16
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(112)
	F_errmsg(m, int32(285775), v11+int32(48))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(473789), int32(948), int32(368880))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	if v93 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v95 = int32(4438516)
	v97 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v97 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(167772201)
	v130 = F_durable_rename(m, v11+int32(112), v11+int32(1136), int32(21))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L31
	}
L26:
	;
	F_XLogRegisterData(m, v11+int32(100), int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_XLogRegisterData(m, l0, int32(524))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v117 = F_XLogInsert(m, int32(7), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_XLogFlush(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(0)
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v136 = m.G0
	v138 = v136 - int32(16)
	m.G0 = v138
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = l4
	v141 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v141)
	F_SendSharedInvalidMessages(m, v138, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if l3 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	m.G0 = v138 + int32(16)
	goto L34
L36:
	;
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v152 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v164 = int32(0)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l5
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v164<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v171
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v11)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v174
	F_RelationPreserveStorage(m, v11+int32(16), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v182 = v164 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v182 < v183 {
		v164 = v182
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v193 = int32(4438516)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v195 - int32(1)
	goto L45
L44:
	;
	goto L45
L45:
	;
	m.G0 = v11 + int32(2160)
	return
L46:
	;
	F_errmsg_internal(m, int32(318902), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(473789), int32(910), int32(368880))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(112)
	F_errmsg(m, int32(285190), v11)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(473789), int32(936), int32(368880))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(112)
	F_errmsg(m, int32(285933), v11+int32(32))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(473789), int32(957), int32(368880))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
