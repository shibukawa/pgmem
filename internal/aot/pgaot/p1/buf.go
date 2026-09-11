package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v3 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v8 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v12 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pfree(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v12<<(uint(int32(2))%32))))
	F_FileClose(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 < v22 {
		v12 = v21
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_pfree(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	return
}
func F_BufFileCreateFileSet(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
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
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v11 = F_palloc(m, int32(8240))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v21 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v20
		*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
		v29 = F_pstrdup(m, l1)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v29
			v33 = F_palloc(m, int32(4))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v36
				v46 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(438200), v8+int32(16))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					v51 = F_FileSetDelete(m, v48, v8+int32(32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v53
						v61 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(438200), v8)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v66 = F_FileSetCreate(m, v63, v8+int32(32))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v68))) = v66
								v70 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)) = uint8(v70)
								m.G0 = v8 + int32(1056)
								return v11
							}
						}
					}
				}
			}
		}
	}
}
func F_BufFileReadMaybeEOF(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_BufFileReadCommon(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BufFileSeek(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v199 int32
	_ = v199
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	switch l3 {
	case 0:
		goto L7
	case 1:
		goto L4
	case 2:
		goto L6
	default:
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v199
L2:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v113 != v117 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	if int64(0) <= v83 {
		v113 = v82
		v114 = v83
		goto L2
	} else {
		goto L20
	}
L4:
	;
	v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v82 = v81
	v83 = v77 + (v78 + l2)
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L17
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v18 - int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17+v20<<(uint(int32(2))%32))))
	v25 = F_FileSize(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if int32(0) <= l1 {
		v82 = l1
		v83 = l2
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v199 = int32(-1)
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	if int64(0) <= v25 {
		v113 = v20
		v114 = v25
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32)-int32(4))))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44*int32(48))+32))
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v50
	F_errmsg(m, int32(281557), v12+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(468330), int32(776), int32(296713))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
	F_errmsg_internal(m, int32(457575), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(468330), int32(779), int32(296713))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v87 = int64(-1073741824)
	v88 = v87 - v83
	if base.Ui64(v88) <= base.Ui64(v87) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v92 = v88
	goto L23
L22:
	;
	v92 = int64(0)
	goto L23
L23:
	;
	v94 = v92 + int64(1073741823)
	v96 = int64(base.Ui64(v94) >> (uint(int64(30)) % 64))
	v98 = v82 - int32(1)
	v99 = int32(-1)
	if v99 <= v98 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v102 = v99
	goto L26
L25:
	;
	v102 = v98
	goto L26
L26:
	;
	if base.Ui64(base.I64_extend_i32_u(v98-v102)) <= base.Ui64(v96) {
		v199 = int32(-1)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v113 = v98 - base.I32_wrap_i64(v96)
	v114 = v83 + v94&int64(-1073741824) + int64(1073741824)
	goto L2
L28:
	;
	v199 = int32(0)
	goto L1
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v127 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if v114 < v119 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+44)))
	if v119+v121 < v114 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v124 = v114 - v119
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)) = uint32(v124)
	goto L28
L33:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = base.B2i32(v114 == int64(0)) & base.B2i32(v113 == v134)
	v137 = v113 - v136
	if v136 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v139 = int64(1073741824)
	goto L39
L38:
	;
	v139 = v114
	goto L39
L39:
	;
	if int64(1073741825) <= v139 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v144 = v139 - int64(2147483648)
	if base.Ui64(v144) <= base.Ui64(v139) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v174 = v139
	v175 = v137
	goto L42
L42:
	;
	if v134 <= v175 {
		v199 = int32(-1)
		goto L1
	} else {
		goto L50
	}
L43:
	;
	v147 = v144
	goto L45
L44:
	;
	v147 = int64(0)
	goto L45
L45:
	;
	v149 = v147 + int64(1073741823)
	v151 = int64(base.Ui64(v149) >> (uint(int64(30)) % 64))
	v155 = v137 + int32(1)
	if v155 < v134 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v157 = v134
	goto L48
L47:
	;
	v157 = v155
	goto L48
L48:
	;
	if base.Ui64(base.I64_extend_i32_u(v113^int32(-1)+v157-(int32(0)-v136))) <= base.Ui64(v151) {
		v199 = int32(-1)
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v174 = v139 - v149&int64(-1073741824) - int64(1073741824)
	v175 = v137 + base.I32_wrap_i64(v151) + int32(1)
	goto L42
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v175
	goto L28
}
func F_BufTableLookup(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, _consts[722]))
	v5 = int32(0)
	v7 = F_hash_search_with_hash_value(m, v4, l0, l1, v5, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			return int32(-1)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			return v15
		}
	}
}
