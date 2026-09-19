package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WarnNoTransactionBlock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_WarnNoTransactionBlock[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if base.B2i32(l0 == int32(0))|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v12)) != 0 {
		m.G0 = v6 + int32(16)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		if int32(1) < v16 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v21 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					F_errcode(m, int32(16908610))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
						F_errmsg(m, int32(_a_F_WarnNoTransactionBlock_0), v6)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_WarnNoTransactionBlock_1), int32(3749), int32(_a_F_WarnNoTransactionBlock_2))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_win866_to_win1251(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13938(m, l0, int32(_a_F_win866_to_win1251_0), int32(23), int32(20))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	var v7 int32
	_ = v7
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_wrapper_handler[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_wrapper_handler[1]))
	v7 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	if v6 != v7 {
		v9 = int32(0)
		v11 = m.G0
		v13 = v11 - int32(32)
		m.G0 = v13
		switch int32(2) {
		case 0, 2:
			v23 = v9
		default:
			*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_wrapper_handler[2]))) = v9
			v23 = int32(_a_F_wrapper_handler_0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v23
		F_sigemptyset(m, v13+int32(16))
		mBase = m.M
		if l0 == int32(17) {
			v32 = int32(268435457)
		} else {
			v32 = int32(268435456)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
		v37 = F___sigaction(m, l0, v13+int32(12), int32(0))
		mBase = m.M
		m.G0 = v13 + int32(32)
		F_raise(m, l0)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			return
		}
	} else {
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_wrapper_handler[2])))
		m.T0[v47].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_wrapper_handler[0])) = v4
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v9 = m.G0
	v11 = v9 - int32(2160)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_write_relmap_file_0)
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
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(_a_F_write_relmap_file_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l6
	v35 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(_a_F_write_relmap_file_2), v11+int32(80))
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
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L46
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(_a_F_write_relmap_file_3)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l6
	v41 = v11 + int32(112)
	v46 = F_pg_snprintf(m, v41, int32(1024), int32(_a_F_write_relmap_file_2), v11-int32(-64))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v49 = F_OpenTransientFile(m, v41, int32(577))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v49 < int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(167772202)
	v57 = int32(524)
	v58 = F_write(m, v49, l0, v57)
	mBase = m.M
	if v58 != v57 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[1]))
	if v62 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(0)
	v91 = F_CloseTransientFile(m, v49)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L21
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[1])) = int32(51)
	goto L16
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(112)
	F_errmsg(m, int32(_a_F_write_relmap_file_4), v11+int32(48))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_write_relmap_file_5), int32(948), int32(_a_F_write_relmap_file_6))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	if v91 != 0 {
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
	v93 = int32(_a_F_write_relmap_file_7)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[2])) = v95 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(167772201)
	v128 = F_durable_rename(m, v11+int32(112), v11+int32(1136), int32(21))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L31
	}
L26:
	;
	F_XLogRegisterData(m, v11+int32(100), int32(12))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_XLogRegisterData(m, l0, int32(524))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v115 = F_XLogInsert(m, int32(7), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_XLogFlush(m, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(0)
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v134 = m.G0
	v136 = v134 - int32(16)
	m.G0 = v136
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = l4
	v139 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v139)
	F_SIInsertDataEntries(m, v136, int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	m.G0 = v136 + int32(16)
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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v150 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v160 = int32(0)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l5
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0+v160<<(uint(int32(3))%32))+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v167
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v11)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v167
	F_RelationPreserveStorage(m, v11+int32(16), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v178 = v160 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v178 < v179 {
		v160 = v178
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v189 = int32(_a_F_write_relmap_file_7)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_write_relmap_file[2])) = v191 - int32(1)
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
	F_errmsg_internal(m, int32(_a_F_write_relmap_file_8), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_write_relmap_file_5), int32(910), int32(_a_F_write_relmap_file_6))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
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
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(112)
	F_errmsg(m, int32(_a_F_write_relmap_file_9), v11)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_write_relmap_file_5), int32(936), int32(_a_F_write_relmap_file_6))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(112)
	F_errmsg(m, int32(_a_F_write_relmap_file_10), v11+int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_write_relmap_file_5), int32(957), int32(_a_F_write_relmap_file_6))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
