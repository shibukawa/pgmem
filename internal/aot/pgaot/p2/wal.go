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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
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
	v23 = v8 + int32(48)
	v28 = F_pg_snprintf(m, v23, int32(1024), int32(_a_F_OpenWalSummaryFile_0), v8+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v33 = F_PathNameOpenFile(m, v23, int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			if v33 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
						F_errmsg(m, int32(_a_F_OpenWalSummaryFile_1), v8)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_OpenWalSummaryFile_2), int32(220), int32(_a_F_OpenWalSummaryFile_3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
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
				return v33
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
	v19 = F_FileReadV(m, v12, v9+int32(8), int32(1), v11, int32(167772238))
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
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReadWalSummary[0]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v31*int32(48))+32))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v37
					F_errmsg(m, int32(_a_F_ReadWalSummary_0), v9)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ReadWalSummary_1), int32(284), int32(_a_F_ReadWalSummary_2))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(16)
	F_initStringInfo(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v15 = F_appendStringInfoVA(m, v11, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v23 = v8 + int32(16)
	F_enlargeStringInfo(m, v23, v20)
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v27 = F_appendStringInfoVA(m, v23, l1, l2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v27 != 0 {
		v20 = v27
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
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
	F_errmsg_internal(m, int32(_a_F_ReportWalSummaryError_0), v8)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_ReportWalSummaryError_1), int32(340), int32(_a_F_ReportWalSummaryError_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int64
	_ = v208
	var v212 int32
	_ = v212
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v233 int64
	_ = v233
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l4
	if l3 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v255
L2:
	;
	v255 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = l0 + int32(1168)
	v23 = l1
	v24 = l2
	v25 = l3
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v37 = base.I32_wrap_i64(v24) & (v34 - int32(1))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1168))
	if int32(0) <= v38 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v255 = v119
	goto L1
L7:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WALRead[0])))
	v72 = m.G0
	v74 = v72 - int32(16)
	m.G0 = v74
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	v43 = base.I64_div_u_s(v24, base.I64_extend_i32_s(v34))
	if v41 == v43 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v54 = v34
	goto L10
L10:
	;
	v56 = base.I64_div_u_s(v24, base.I64_extend_i32_s(v54))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	m.T0[v59].(func(*base.Module, int32, int64, int32))(m, l0, v56, v14+int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L17
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	if v45 == v46 {
		v66 = v34
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v48].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v54 = v53
	goto L10
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1184)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v66 = v65
	goto L7
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_WALRead[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(167772237)
	*(*int32)(unsafe.Add(mBase, _c_F_WALRead[2])) = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v95 = v66 - v37
	if base.Ui32(v25) < base.Ui32(v95) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F___clock_gettime(m, int32(1), v74)
	mBase = m.M
	v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v74)+8)))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	v83 = v78 + v79*int64(1000000000)
	goto L21
L20:
	;
	v83 = int64(0)
	goto L21
L21:
	;
	m.G0 = v74 + int32(16)
	goto L18
L22:
	;
	v97 = v25
	goto L24
L23:
	;
	v97 = v95
	goto L24
L24:
	;
	v99 = F_pread(m, v94, v23, v97, base.I64_extend_i32_u(v37))
	mBase = m.M
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_WALRead[1]))
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102
	if v99 <= v102 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_WALRead[2]))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v37
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+16)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v116
	v255 = int32(0)
	goto L1
L26:
	;
	goto L27
L27:
	;
	v119 = int32(1)
	v124 = base.I64_extend_i32_u(v99)
	v128 = m.G0
	v130 = v128 - int32(16)
	m.G0 = v130
	if v83 != int64(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v248 = v25 - v99
	if v248 != 0 {
		v23 = v23 + v99
		v24 = v24 + v124
		v25 = v248
		goto L5
	} else {
		goto L45
	}
L29:
	;
	F___clock_gettime(m, int32(1), v130)
	mBase = m.M
	v136 = int64(*(*int32)(unsafe.Add(mBase, uint32(v130)+8)))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	v141 = v136 + (v137*int64(1000000000) - v83)
	goto L32
L30:
	;
	goto L31
L31:
	;
	v228 = int32(880)
	v229 = *(*int64)(unsafe.Add(mBase, _c_F_WALRead[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_WALRead[3])) = v229 + base.I64_extend_i32_u(v119)
	v233 = *(*int64)(unsafe.Add(mBase, _c_F_WALRead[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_WALRead[4])) = v233 + v124
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v119, v124)
	mBase = m.M
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WALRead[5])) = uint8(v238)
	*(*uint8)(unsafe.Add(mBase, _c_F_WALRead[6])) = uint8(v238)
	m.G0 = v130 + int32(16)
	goto L28
L32:
	;
	v191 = int32(880)
	v192 = *(*int64)(unsafe.Add(mBase, _c_F_WALRead[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_WALRead[7])) = v192 + v141
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_WALRead[8]))
	v203 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v196))|base.B2i32(int32(1)<<(uint(v196)%32)&int32(_a_F_WALRead_0) == v203) == v203 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_WALRead[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_WALRead[9])) = v208 + v141
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WALRead[5])) = uint8(v212)
	*(*uint8)(unsafe.Add(mBase, _c_F_WALRead[10])) = uint8(v212)
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L31
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
	v16 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WALReadRaiseError[0])))
	v17 = base.I64_div_u_s(int64(4294967296), v16)
	v18 = base.I64_div_u_s(v11, v17)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v18)
	v21 = v11 - v17*v18
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v21)
	v29 = F_pg_snprintf(m, v9+int32(48), int32(64), int32(_a_F_WALReadRaiseError_0), v9+int32(32))
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
						F_errmsg(m, int32(_a_F_WALReadRaiseError_1), v9+int32(16))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_WALReadRaiseError_2), int32(1068), int32(_a_F_WALReadRaiseError_3))
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
			*(*int32)(unsafe.Add(mBase, _c_F_WALReadRaiseError[1])) = v40
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
					F_errmsg(m, int32(_a_F_WALReadRaiseError_4), v9)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_WALReadRaiseError_2), int32(1060), int32(_a_F_WALReadRaiseError_3))
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
func F_WalSndCheckShutdownTimeout(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[0]))
	if v4 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[1]))
		if v8 == int32(0) {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[2]))
			if v12 == int32(0) {
				v71 = m.G0
				v73 = v71 - int32(16)
				m.G0 = v73
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[3]))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if base.Ui32(v77-int32(5)) < base.Ui32(int32(-3)) {
					v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
					if v101 == int32(2) {
						*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
					} else {
					}
					v109 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						if v109 != 0 {
							F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
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
							F_proc_exit(m, int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])))
					if v83&int32(1) != 0 {
						v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
						if v101 == int32(2) {
							*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
						} else {
						}
						v109 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							if v109 != 0 {
								F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											F_proc_exit(m, int32(0))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
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
								F_proc_exit(m, int32(0))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(56)
						F_EndCommandExtended(m, v73)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return
						} else {
							v93 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])) = uint8(v93)
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[6]))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							v98 = m.T0[v97].(func(*base.Module) int32)(m)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
								if v101 == int32(2) {
									*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
								} else {
								}
								v109 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									if v109 != 0 {
										F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_proc_exit(m, int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
										F_proc_exit(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
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
			} else {
				v16 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7]))
				if v16 == int64(0) {
					v23 = m.G0
					v24 = int32(16)
					v25 = v23 - v24
					m.G0 = v25
					F_gettimeofday(m, v25)
					mBase = m.M
					v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
					v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
					m.G0 = v25 + v24
					*(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7])) = v29 + v28*int64(1000000) - int64(946684800000000)
					return
				} else {
					if v12 == int32(-1) {
						return
					} else {
						v44 = m.G0
						v45 = int32(16)
						v46 = v44 - v45
						m.G0 = v46
						F_gettimeofday(m, v46)
						mBase = m.M
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
						v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
						m.G0 = v46 + v45
						v60 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7]))
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[2]))
						if base.I64_extend_i32_s(v62)*int64(1000) <= v50+v49*int64(1000000)-int64(946684800000000)-v60 {
							v71 = m.G0
							v73 = v71 - int32(16)
							m.G0 = v73
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[3]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if base.Ui32(v77-int32(5)) < base.Ui32(int32(-3)) {
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
								if v101 == int32(2) {
									*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
								} else {
								}
								v109 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									if v109 != 0 {
										F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_proc_exit(m, int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
										F_proc_exit(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])))
								if v83&int32(1) != 0 {
									v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
									if v101 == int32(2) {
										*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
									} else {
									}
									v109 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										if v109 != 0 {
											F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return
											} else {
												v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														F_proc_exit(m, int32(0))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
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
											F_proc_exit(m, int32(0))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(56)
									F_EndCommandExtended(m, v73)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v93 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])) = uint8(v93)
										v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[6]))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
										v98 = m.T0[v97].(func(*base.Module) int32)(m)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
											if v101 == int32(2) {
												*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
											} else {
											}
											v109 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return
											} else {
												if v109 != 0 {
													F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																F_proc_exit(m, int32(0))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
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
													F_proc_exit(m, int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[2]))
		if v12 == int32(0) {
			v71 = m.G0
			v73 = v71 - int32(16)
			m.G0 = v73
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[3]))
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
			if base.Ui32(v77-int32(5)) < base.Ui32(int32(-3)) {
				v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
				if v101 == int32(2) {
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
				} else {
				}
				v109 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return
				} else {
					if v109 != 0 {
						F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return
						} else {
							v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
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
						F_proc_exit(m, int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])))
				if v83&int32(1) != 0 {
					v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
					if v101 == int32(2) {
						*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
					} else {
					}
					v109 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						if v109 != 0 {
							F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
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
							F_proc_exit(m, int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(56)
					F_EndCommandExtended(m, v73)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						v93 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])) = uint8(v93)
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[6]))
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
						v98 = m.T0[v97].(func(*base.Module) int32)(m)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
							if v101 == int32(2) {
								*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
							} else {
							}
							v109 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								if v109 != 0 {
									F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												F_proc_exit(m, int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
									F_proc_exit(m, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
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
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7]))
			if v16 == int64(0) {
				v23 = m.G0
				v24 = int32(16)
				v25 = v23 - v24
				m.G0 = v25
				F_gettimeofday(m, v25)
				mBase = m.M
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
				m.G0 = v25 + v24
				*(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7])) = v29 + v28*int64(1000000) - int64(946684800000000)
				return
			} else {
				if v12 == int32(-1) {
					return
				} else {
					v44 = m.G0
					v45 = int32(16)
					v46 = v44 - v45
					m.G0 = v46
					F_gettimeofday(m, v46)
					mBase = m.M
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
					v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
					m.G0 = v46 + v45
					v60 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[7]))
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[2]))
					if base.I64_extend_i32_s(v62)*int64(1000) <= v50+v49*int64(1000000)-int64(946684800000000)-v60 {
						v71 = m.G0
						v73 = v71 - int32(16)
						m.G0 = v73
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[3]))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						if base.Ui32(v77-int32(5)) < base.Ui32(int32(-3)) {
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
							if v101 == int32(2) {
								*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
							} else {
							}
							v109 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								if v109 != 0 {
									F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												F_proc_exit(m, int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
									F_proc_exit(m, int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])))
							if v83&int32(1) != 0 {
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
								if v101 == int32(2) {
									*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
								} else {
								}
								v109 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									if v109 != 0 {
										F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_proc_exit(m, int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
										F_proc_exit(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(56)
								F_EndCommandExtended(m, v73)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v93 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[5])) = uint8(v93)
									v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[6]))
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
									v98 = m.T0[v97].(func(*base.Module) int32)(m)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4]))
										if v101 == int32(2) {
											*(*int32)(unsafe.Add(mBase, _c_F_WalSndCheckShutdownTimeout[4])) = int32(0)
										} else {
										}
										v109 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											if v109 != 0 {
												F_errmsg(m, int32(_a_F_WalSndCheckShutdownTimeout_0), int32(0))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return
												} else {
													v117 = F_errdetail(m, int32(_a_F_WalSndCheckShutdownTimeout_1), int32(0))
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_WalSndCheckShutdownTimeout_2), int32(3802), int32(_a_F_WalSndCheckShutdownTimeout_3))
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_proc_exit(m, int32(0))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
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
												F_proc_exit(m, int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
					} else {
						return
					}
				}
			}
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
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[0]))
	F_ModifyWaitEvent(m, v10, v4, l0, v4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if l2 == int32(100663302) {
			v23 = int32(76)
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[1]))
			F_ConditionVariablePrepareToSleep(m, v25+v23)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[0]))
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
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[2]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
			switch v21 {
			case 0:
				v23 = int32(52)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[1]))
				F_ConditionVariablePrepareToSleep(m, v25+v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[0]))
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
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[1]))
				F_ConditionVariablePrepareToSleep(m, v25+v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[0]))
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
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWait[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_assign_wal_consistency_checking[0])) = l1
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[0]))
	if v11 == l0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[1]))
	v110 = v8 + int32(48)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[2]))
	v114 = *(*int64)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[3]))
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[4]))
	F_XLogFileName(m, v110, v112, v114, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L21
	} else {
		goto L34
	}
L2:
	;
	m.G0 = v8 + int32(112)
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[5]))
	if v14 < int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(167772241)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[7])))
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
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[1]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(0)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[7])))
	if v43 != int32(1) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[8]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[9]))
	if v53 != int32(14) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = int32(_a_F_assign_wal_sync_method_0)
	goto L16
L15:
	;
	v56 = int32(0)
	goto L16
L16:
	;
	v57 = v47 << (uint(int32(13)) % 32) & v56
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[0]))
	switch v59 {
	case 0, 1, 3:
		v77 = v57
		goto L17
	case 2:
		goto L18
	case 4:
		goto L20
	default:
		goto L19
	}
L17:
	;
	switch l0 {
	case 0, 1, 3:
		v97 = v57
		goto L25
	case 2:
		goto L26
	case 4:
		goto L28
	default:
		goto L27
	}
L18:
	;
	v77 = v57 | int32(_a_F_assign_wal_sync_method_1)
	goto L17
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v77 = v57 | int32(_a_F_assign_wal_sync_method_2)
	goto L17
L21:
	;
	return
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v59
	F_errmsg_internal(m, int32(_a_F_assign_wal_sync_method_3), v8)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_assign_wal_sync_method_4), int32(_a_F_assign_wal_sync_method_5), int32(_a_F_assign_wal_sync_method_6))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	if v97 == v77 {
		goto L2
	} else {
		goto L32
	}
L26:
	;
	v97 = v57 | int32(_a_F_assign_wal_sync_method_1)
	goto L25
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L21
	} else {
		goto L29
	}
L28:
	;
	v97 = v57 | int32(_a_F_assign_wal_sync_method_2)
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_assign_wal_sync_method_3), v8+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_assign_wal_sync_method_4), int32(_a_F_assign_wal_sync_method_5), int32(_a_F_assign_wal_sync_method_6))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_wal_sync_method[1])) = v108
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v110
	F_errmsg(m, int32(_a_F_assign_wal_sync_method_7), v8+int32(32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_assign_wal_sync_method_4), int32(_a_F_assign_wal_sync_method_8), int32(_a_F_assign_wal_sync_method_9))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
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
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(80)
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
	m.G0 = v8 + int32(80)
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
	F_ExplainPropertyInteger(m, int32(_a_F_show_wal_usage_6), int32(0), v10, l0)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L35
	}
L5:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L12
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
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v22 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if v25 == int64(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	return
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v30, int32(_a_F_show_wal_usage_0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if int64(0) < v34 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v34
	F_appendStringInfo(m, v37, int32(_a_F_show_wal_usage_1), v8-int32(-64))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v44 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v44
	F_appendStringInfo(m, v47, int32(_a_F_show_wal_usage_2), v8+int32(48))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v54 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v54
	F_appendStringInfo(m, v57, int32(_a_F_show_wal_usage_3), v8+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if v64 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v64
	F_appendStringInfo(m, v67, int32(_a_F_show_wal_usage_4), v8+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v74 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
	F_appendStringInfo(m, v77, int32(_a_F_show_wal_usage_5), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v82, int32(10))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L1
L35:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ExplainPropertyInteger(m, int32(_a_F_show_wal_usage_7), int32(0), v92, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_ExplainPropertyUInteger(m, int32(_a_F_show_wal_usage_8), int32(0), v97, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	F_ExplainPropertyUInteger(m, int32(_a_F_show_wal_usage_9), int32(0), v102, l0)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	F_ExplainPropertyInteger(m, int32(_a_F_show_wal_usage_10), int32(0), v107, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	goto L1
}
