package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenWalSummaryFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = m.G0
	v8 = v6 - int32(1072)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+32)) = uint32(v12)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v16 = int64(32)
	v17 = int64(base.Ui64(v12) >> (uint(v16) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+28)) = uint32(v17)
	v20 = int64(base.Ui64(v11) >> (uint(v16) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v20)
	v28 = F_pg_snprintf(m, v8+int32(48), int32(1024), int32(18207), v8+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v35 = F_PathNameOpenFile(m, v8+int32(48), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(48)
						F_errmsg(m, int32(313019), v8)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515903), int32(220), int32(408559))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
				m.G0 = v8 + int32(1072)
				return v35
			}
		}
	}
}
func F_ReadWalSummary(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
	v19 = F_FileReadV(m, v12, v9+int32(8), int32(1), v11, int32(167772236))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, _consts[265]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v31*int32(48))+32))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v37
					F_errmsg(m, int32(313938), v9)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515903), int32(284), int32(18269))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
			v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v47 + base.I64_extend_i32_u(v19)
			m.G0 = v9 + int32(16)
			return v19
		}
	}
}
func F_ReportWalSummaryError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	F_initStringInfo(m, v7+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v16 = F_appendStringInfoVA(m, v7+int32(16), l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v16
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	F_enlargeStringInfo(m, v7+int32(16), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v29 = F_appendStringInfoVA(m, v7+int32(16), l1, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v29 != 0 {
		v21 = v29
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v42
	F_errmsg_internal(m, int32(216894), v7)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(515903), int32(340), int32(223632))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WALRead(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
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
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v256 int32
	_ = v256
	var v280 int32
	_ = v280
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l4
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v280
L2:
	;
	v18 = l0 + int32(1168)
	v20 = l1
	v21 = l2
	v22 = l3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v280 = int32(1)
	goto L1
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v34 = base.I32_wrap_i64(v21) & (v31 - int32(1))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1168))
	if int32(0) <= v35 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[241])))
	v69 = m.G0
	v71 = v69 - int32(16)
	m.G0 = v71
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	v40 = base.I64_div_u_s(v21, base.I64_extend_i32_s(v31))
	if v38 == v40 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v51 = v31
	goto L10
L10:
	;
	v53 = base.I64_div_u_s(v21, base.I64_extend_i32_s(v51))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	m.T0[v56].(func(*base.Module, int32, int64, int32))(m, l0, v53, v14+int32(12))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L17
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	if v42 == v43 {
		v63 = v31
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v45].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v51 = v50
	goto L10
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1184)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v63 = v62
	goto L7
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(167772235)
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v92 = v63 - v34
	if base.Ui32(v22) < base.Ui32(v92) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F___clock_gettime(m, int32(1), v71)
	mBase = m.M
	v75 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+8)))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	v80 = v75 + v76*int64(1000000000)
	goto L21
L20:
	;
	v80 = int64(0)
	goto L21
L21:
	;
	m.G0 = v71 + int32(16)
	goto L18
L22:
	;
	v94 = v22
	goto L24
L23:
	;
	v94 = v92
	goto L24
L24:
	;
	v96 = F_pread(m, v91, v20, v94, base.I64_extend_i32_u(v34))
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = int32(0)
	v104 = int32(1)
	v105 = base.I64_extend_i32_s(v96)
	v109 = m.G0
	v111 = v109 - int32(16)
	m.G0 = v111
	if v80 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v96 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L26:
	;
	F___clock_gettime(m, int32(1), v111)
	mBase = m.M
	v117 = int64(*(*int32)(unsafe.Add(mBase, uint32(v111)+8)))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	v122 = v117 + (v118*int64(1000000000) - v80)
	goto L29
L27:
	;
	goto L28
L28:
	;
	v219 = int32(4543704)
	v220 = *(*int64)(unsafe.Add(mBase, _consts[242]))
	*(*int64)(unsafe.Add(mBase, _consts[242])) = v220 + base.I64_extend_i32_u(v104)
	v225 = int32(4542744)
	v226 = *(*int64)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, _consts[243])) = v226 + v105
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v104, v105)
	mBase = m.M
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[244])) = uint8(v231)
	*(*uint8)(unsafe.Add(mBase, _consts[245])) = uint8(v231)
	m.G0 = v111 + int32(16)
	goto L25
L29:
	;
	v174 = int32(4544664)
	v175 = *(*int64)(unsafe.Add(mBase, _consts[246]))
	*(*int64)(unsafe.Add(mBase, _consts[246])) = v175 + v122
	v179 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if base.Ui32(int32(16)) < base.Ui32(v179) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	goto L28
L40:
	;
	if int32(1)<<(uint(v179)%32)&int32(115186) == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v197 = int32(4541560)
	v198 = *(*int64)(unsafe.Add(mBase, _consts[247]))
	*(*int64)(unsafe.Add(mBase, _consts[247])) = v198 + v122
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[244])) = uint8(v202)
	*(*uint8)(unsafe.Add(mBase, _consts[248])) = uint8(v202)
	goto L39
L42:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v242
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v34
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+16)) = v247
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v249
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v251
	v280 = int32(0)
	goto L1
L43:
	;
	goto L44
L44:
	;
	v256 = v22 - v96
	if v256 != 0 {
		v20 = v20 + v96
		v21 = v21 + v105
		v22 = v256
		goto L5
	} else {
		goto L45
	}
L45:
	;
	goto L6
}
func F_WALReadRaiseError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v12
	v16 = int64(*(*int32)(unsafe.Add(mBase, _consts[179])))
	v17 = base.I64_div_u_s(int64(4294967296), v16)
	v18 = base.I64_div_u_s(v11, v17)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v18)
	v21 = v11 - v17*v18
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v21)
	v29 = F_pg_snprintf(m, v9+int32(48), int32(64), int32(536148), v9+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if int32(0) <= v31 {
			if v31 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v69
						*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = base.I64_rotl(v68, int64(32))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
						F_errmsg(m, int32(497074), v9+int32(16))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							F_errfinish(m, int32(518033), int32(1032), int32(223794))
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
				m.G0 = v9 + int32(112)
				return
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, _consts[158])) = v40
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
					F_errmsg(m, int32(310152), v9)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errfinish(m, int32(518033), int32(1024), int32(223794))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
func F_WalRcvForceReply(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, _consts[693]))
	v5 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+1472)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = v5
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[693]))
		F_s_lock(m, v11+int32(1456), int32(519652), int32(1356), int32(19760))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[693]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+1456)) = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v23 != int32(-1) {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[156]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				F_SetLatch(m, v28+v23*int32(640)+int32(20))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[693]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+1456)) = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		if v23 != int32(-1) {
			v27 = *(*int32)(unsafe.Add(mBase, _consts[156]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			F_SetLatch(m, v28+v23*int32(640)+int32(20))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_WalSndWait(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[728]))
	F_ModifyWaitEvent(m, v10, v4, l0, v4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if l2 == int32(100663302) {
			v23 = int32(76)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[687]))
			F_ConditionVariablePrepareToSleep(m, v25+v23)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[728]))
				v33 = F_WaitEventSetWait(m, v31, l1, v7, int32(1), l2)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v33 != int32(1) {
						F_ConditionVariableCancelSleep(m)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
						if v37&int32(16) == int32(0) {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_proc_exit(m, int32(1))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
			v20 = *(*int32)(unsafe.Add(mBase, _consts[689]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
			switch v21 {
			case 0:
				v23 = int32(52)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[687]))
				F_ConditionVariablePrepareToSleep(m, v25+v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[728]))
					v33 = F_WaitEventSetWait(m, v31, l1, v7, int32(1), l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v33 != int32(1) {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
							if v37&int32(16) == int32(0) {
								F_ConditionVariableCancelSleep(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								F_ConditionVariableCancelSleep(m)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_proc_exit(m, int32(1))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
			case 1:
				v23 = int32(64)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[687]))
				F_ConditionVariablePrepareToSleep(m, v25+v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[728]))
					v33 = F_WaitEventSetWait(m, v31, l1, v7, int32(1), l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v33 != int32(1) {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
							if v37&int32(16) == int32(0) {
								F_ConditionVariableCancelSleep(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								F_ConditionVariableCancelSleep(m)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_proc_exit(m, int32(1))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
			default:
				v31 = *(*int32)(unsafe.Add(mBase, _consts[728]))
				v33 = F_WaitEventSetWait(m, v31, l1, v7, int32(1), l2)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v33 != int32(1) {
						F_ConditionVariableCancelSleep(m)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
						if v37&int32(16) == int32(0) {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							F_ConditionVariableCancelSleep(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_proc_exit(m, int32(1))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
func F_assign_wal_consistency_checking(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[204])) = l1
	return
}
func F_assign_wal_sync_method(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	if v11 == l0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v111 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v113 = *(*int64)(unsafe.Add(mBase, _consts[227]))
	v115 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	F_XLogFileName(m, v8+int32(48), v111, v113, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v8 + int32(112)
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	if v14 < int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(167772239)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[194])))
	if v23 != int32(1) {
		v37 = int32(0)
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	goto L8
L8:
	;
	v28 = F_fsync(m, v14)
	mBase = m.M
	if v28 != int32(-1) {
		v37 = v28
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v37 = int32(-1)
	goto L6
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v32 == int32(27) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(0)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[194])))
	if v43 != int32(1) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	v51 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	v52 = int32(14)
	v56 = v47 << (uint(int32(13)) % 32) & (base.B2i32(v51 != v52) << (uint(v52) % 32))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	switch v58 {
	case 0, 1, 3:
		v76 = v56
		goto L14
	case 2:
		goto L15
	case 4:
		goto L17
	default:
		goto L16
	}
L14:
	;
	switch l0 {
	case 0, 1, 3:
		v96 = v56
		goto L22
	case 2:
		goto L23
	case 4:
		goto L25
	default:
		goto L24
	}
L15:
	;
	v76 = v56 | int32(1052672)
	goto L14
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v76 = v56 | int32(4096)
	goto L14
L18:
	;
	return
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v58
	F_errmsg_internal(m, int32(511063), v8)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(523077), int32(8695), int32(110181))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	if v96 == v76 {
		goto L2
	} else {
		goto L29
	}
L23:
	;
	v96 = v56 | int32(1052672)
	goto L22
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v96 = v56 | int32(4096)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg_internal(m, int32(511063), v8+int32(16))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(523077), int32(8695), int32(110181))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	goto L2
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v107
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(48)
	F_errmsg(m, int32(313967), v8+int32(32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(523077), int32(8728), int32(443138))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_show_wal_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 - int32(-64)
	return
L2:
	;
	if int64(0) < v10 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_ExplainPropertyInteger(m, int32(182558), int32(0), v10, l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L30
	}
L5:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v16 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v19 != int64(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if v22 <= int64(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	return
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v27, int32(574266))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if int64(0) < v31 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v31
	F_appendStringInfo(m, v34, int32(450613), v6+int32(-16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v42 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v42
	F_appendStringInfo(m, v45, int32(450660), v6+int32(-32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v53 != int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v53
	F_appendStringInfo(m, v56, int32(39425), v6+int32(-48))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if int64(0) < v64 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v64
	F_appendStringInfo(m, v67, int32(450641), v8)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v73, int32(10))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L1
L30:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ExplainPropertyInteger(m, int32(560624), int32(0), v83, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_ExplainPropertyUInteger(m, int32(169840), int32(0), v88, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	F_ExplainPropertyInteger(m, int32(318613), int32(0), v93, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L1
}
