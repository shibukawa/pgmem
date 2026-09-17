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
	F_errfinish(m, int32(_a_F_BufFileDeleteFileSet_2), int32(387), int32(_a_F_BufFileDeleteFileSet_3))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v18 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_BufFileDumpBuffer(m, l0)
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
	if l2 == int32(0) {
		v153 = v5
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
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v195
	F_errstart_cold(m, int32(21), v195)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L46
	}
L7:
	;
	v159 = int32(0)
	if l3&base.B2i32(v153 == v159)|base.B2i32(l2 == v153) == v159 {
		goto L34
	} else {
		goto L35
	}
L8:
	;
	v28 = l0 + int32(48)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = l1
	v36 = v29
	v37 = v5
	v38 = l2
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v43 <= v36 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v153 = v143
	goto L7
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v49 = v47 + base.I64_extend_i32_s(v36)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v49 < int64(1073741824) {
		v62 = v51
		v64 = v49
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v130 = v43
	v131 = v36
	goto L13
L13:
	;
	v135 = v130 - v131
	if base.Ui32(v135) < base.Ui32(v38) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v62<<(uint(int32(2))%32))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileReadCommon[0])))
	if v72 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v55 = v51 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v56 <= v55 {
		v62 = v51
		v64 = v49
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v58 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v55
	v62 = v55
	v64 = v58
	goto L14
L17:
	;
	F___clock_gettime(m, int32(1), v14+int32(-16))
	mBase = m.M
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	v82 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+56)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v85 = v84
	v86 = v79*int64(-1000000000) - v82
	goto L19
L18:
	;
	v85 = v64
	v86 = int64(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = int32(_a_F_BufFileReadCommon_0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v28
	v91 = v14 + int32(-16)
	v94 = F_FileReadV(m, v69, v91, int32(1), v85, int32(167772166))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v94
	if v94 < int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileReadCommon[0])))
	if v100 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F___clock_gettime(m, int32(1), v91)
	mBase = m.M
	v105 = int32(_a_F_BufFileReadCommon_1)
	v107 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[1]))
	v108 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+56)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[1])) = v107 + (v108 + (v109*int64(1000000000) + v86))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v117 = v116
	goto L24
L23:
	;
	v117 = v94
	goto L24
L24:
	;
	if v117 <= int32(0) {
		v153 = v37
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v120 = int32(_a_F_BufFileReadCommon_2)
	v122 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileReadCommon[2])) = v122 + int64(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v126 <= int32(0) {
		v153 = v37
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = v126
	v131 = v129
	goto L13
L27:
	;
	v137 = v135
	goto L29
L28:
	;
	v137 = v38
	goto L29
L29:
	;
	if v137 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	base.MemoryCopy(m, v31, v131+v28, v137)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v141 = v140 + v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v141
	v143 = v137 + v37
	v145 = v38 - v137
	if v145 != 0 {
		v31 = v31 + v137
		v36 = v141
		v37 = v143
		v38 = v145
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
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v16 - int32(-64)
	return v153
L37:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v172 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	F_errfinish(m, int32(_a_F_BufFileReadCommon_3), int32(635), int32(_a_F_BufFileReadCommon_4))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v172
	F_errmsg(m, int32(_a_F_BufFileReadCommon_5), v14+int32(-48))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v153
	F_errmsg(m, int32(_a_F_BufFileReadCommon_6), v16)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileReadCommon[3]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v69*int32(48))+32))
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v208
	F_errmsg(m, int32(_a_F_BufFileReadCommon_7), v14+int32(-32))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_BufFileReadCommon_3), int32(471), int32(_a_F_BufFileReadCommon_8))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
