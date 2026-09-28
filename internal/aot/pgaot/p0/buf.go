package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileDeleteFileSet(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v4
	v15 = v9 + int32(48)
	v20 = F_pg_snprintf(m, v15, int32(1024), int32(_a_F_BufFileDeleteFileSet_0), v9+int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = F_FileSetDelete(m, l0, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(1072)
	return
L4:
	;
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	if l2 != 0 {
		goto L3
	} else {
		goto L17
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileDeleteFileSet[0]))
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v35 = v29 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	v39 = v9 + int32(48)
	v42 = F_pg_snprintf(m, v39, int32(1024), int32(_a_F_BufFileDeleteFileSet_0), v9)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v44 = F_FileSetDelete(m, l0, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v44 != 0 {
		v29 = v35
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_BufFileDeleteFileSet_1), v9+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_BufFileDeleteFileSet_2), int32(388), int32(_a_F_BufFileDeleteFileSet_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BufFileReadCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v19 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l2 == int32(0) {
		v161 = v5
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L46
	}
L7:
	;
	v169 = int32(0)
	if l3&base.B2i32(v161 == v169)|base.B2i32(l2 == v161) == v169 {
		goto L34
	} else {
		goto L35
	}
L8:
	;
	v29 = l0 + int32(56)
	v31 = l0 + int32(40)
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = l1
	v39 = v5
	v41 = l2
	v45 = v32
	goto L9
L9:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v47 <= v45 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v161 = v152
	goto L7
L11:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v50 = v49 + v45
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v50
	v52 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v50 < int64(1073741824) {
		v67 = v56
		v69 = v50
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v139 = v45
	v140 = v47
	goto L13
L13:
	;
	v142 = base.I32_wrap_i64(v140 - v139)
	if base.Ui32(v41) < base.Ui32(v142) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v67<<(uint(int32(2))%32))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileReadCommon[0])))
	if v77 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v60 = v56 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v61 <= v60 {
		v67 = v56
		v69 = v50
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v60
	v67 = v60
	v69 = v63
	goto L14
L17:
	;
	F___clock_gettime(m, int32(1), v15+int32(-16))
	mBase = m.M
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+56)))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v90 = v89
	v91 = v84*int64(-1000000000) - v87
	goto L19
L18:
	;
	v90 = v69
	v91 = int64(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(_a_F_BufFileReadCommon_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v29
	v96 = v15 + int32(-16)
	v99 = F_FileReadV(m, v74, v96, int32(1), v90, int32(167772166))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v101 = base.I64_extend_i32_s(v99)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v101
	if v99 < int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileReadCommon[0])))
	if v106 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F___clock_gettime(m, int32(1), v96)
	mBase = m.M
	v111 = int32(_a_F_BufFileReadCommon_1)
	v113 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[1]))
	v114 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+56)))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[1])) = v113 + (v114 + (v115*int64(1000000000) + v91))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v123 = v122
	goto L24
L23:
	;
	v123 = v101
	goto L24
L24:
	;
	if v123 <= int64(0) {
		v161 = v39
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v126 = int32(_a_F_BufFileReadCommon_2)
	v128 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[2])) = v128 + int64(1)
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v132 <= int64(0) {
		v161 = v39
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v139 = v135
	v140 = v132
	goto L13
L27:
	;
	v144 = v41
	goto L29
L28:
	;
	v144 = v142
	goto L29
L29:
	;
	if v144 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	base.MemoryCopy(m, v34, v29+base.I32_wrap_i64(v139), v144)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v150 = v148 + base.I64_extend_i32_u(v144)
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v150
	v152 = v144 + v39
	v154 = v41 - v144
	if v154 != 0 {
		v34 = v34 + v144
		v39 = v152
		v41 = v154
		v45 = v150
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v17 - int32(-64)
	return v161
L37:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v182 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	F_errfinish(m, int32(_a_F_BufFileReadCommon_3), int32(636), int32(_a_F_BufFileReadCommon_4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v182
	F_errmsg(m, int32(_a_F_BufFileReadCommon_5), v15+int32(-48))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v161
	F_errmsg(m, int32(_a_F_BufFileReadCommon_6), v17)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L39
L44:
	;
	goto L39
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileReadCommon[3]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v74*int32(48))+32))
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v218
	F_errmsg(m, int32(_a_F_BufFileReadCommon_7), v15+int32(-32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_BufFileReadCommon_3), int32(472), int32(_a_F_BufFileReadCommon_8))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BufFileReadExact(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = F_BufFileReadCommon(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_BufFileSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32)-int32(4))))
	v17 = F_FileSize(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		if v17 < int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(2))%32)-int32(4))))
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileSize[0]))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36*int32(48))+32))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v42
					F_errmsg(m, int32(_a_F_BufFileSize_0), v7)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_BufFileSize_1), int32(877), int32(_a_F_BufFileSize_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_s(v54-int32(1))<<(uint(int64(30))%64) + v17
		}
	}
}
func F_BufTableShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	v2 = m.G0
	v4 = v2 - int32(96)
	m.G0 = v4
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+48)) = int64(103079215124)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+40)) = int64(128)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(_a_F_BufTableShmemRequest_0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+64)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+72)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+80)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+92)) = int32(_a_F_BufTableShmemRequest_1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = int32(_a_F_BufTableShmemRequest_2)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_BufTableShmemRequest[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = base.I64_extend_i32_s(v31 + int32(128))
	F_ShmemRequestHashWithOpts(m, v4+int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return
	} else {
		m.G0 = v4 + int32(96)
		return
	}
}
