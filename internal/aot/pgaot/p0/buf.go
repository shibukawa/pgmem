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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v4
	v20 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(466808), v9+int32(32))
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
	v24 = F_FileSetDelete(m, l0, v9+int32(48))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v4
	goto L7
L5:
	;
	goto L6
L6:
	;
	if l2|v24 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v37 = v30 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v46 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(466808), v9+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v50 = F_FileSetDelete(m, l0, v9+int32(48))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v50 != 0 {
		v30 = v37
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	m.G0 = v9 + int32(1072)
	return
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(717303), v9)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(499138), int32(387), int32(108698))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v17 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v191
	F_errstart_cold(m, int32(21), v191)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L47
	}
L7:
	;
	if l2 == v153 {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	v27 = l0 + int32(48)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v30 = l1
	v34 = v28
	v36 = v5
	v37 = l2
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v41 <= v34 {
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
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v47 = v45 + base.I64_extend_i32_s(v34)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v47 < int64(1073741824) {
		v60 = v49
		v62 = v47
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v130 = v34
	v131 = v41
	goto L13
L13:
	;
	v135 = v131 - v130
	if base.Ui32(v135) < base.Ui32(v37) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60<<(uint(int32(2))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	if v70 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v53 = v49 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v54 <= v53 {
		v60 = v49
		v62 = v47
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v53
	v60 = v53
	v62 = v56
	goto L14
L17:
	;
	F___clock_gettime(m, int32(1), v13+int32(-16))
	mBase = m.M
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	v81 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+56)))
	v83 = v77
	v84 = v78*int64(-1000000000) - v81
	goto L19
L18:
	;
	v83 = v62
	v84 = int64(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v27
	v92 = F_FileReadV(m, v67, v13+int32(-16), int32(1), v83, int32(167772166))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v92
	if v92 < int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	if v98 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F___clock_gettime(m, int32(1), v13+int32(-16))
	mBase = m.M
	v105 = int32(4413864)
	v107 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v108 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+56)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v107 + (v108 + (v109*int64(1000000000) + v84))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v117 = v116
	goto L24
L23:
	;
	v117 = v92
	goto L24
L24:
	;
	if v117 <= int32(0) {
		v153 = v36
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v120 = int32(4413816)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[337]))
	*(*int64)(unsafe.Add(mBase, _consts[337])) = v122 + int64(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v126 <= int32(0) {
		v153 = v36
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = v129
	v131 = v126
	goto L13
L27:
	;
	v137 = v135
	goto L29
L28:
	;
	v137 = v37
	goto L29
L29:
	;
	if v137 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v141 = v140 + v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v141
	v143 = v137 + v36
	v145 = v37 - v137
	if v145 != 0 {
		v30 = v139 + v137
		v34 = v141
		v36 = v143
		v37 = v145
		goto L9
	} else {
		goto L34
	}
L31:
	;
	v138 = F__emscripten_memcpy_bulkmem(m, v30, v130+v27, v137)
	mBase = m.M
	v139 = v138
	goto L33
L32:
	;
	v139 = v30
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L10
L35:
	;
	m.G0 = v15 - int32(-64)
	return v153
L36:
	;
	if l3&base.B2i32(v153 == int32(0)) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v168 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_errfinish(m, int32(499138), int32(635), int32(246067))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L46
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v168
	F_errmsg(m, int32(158971), v13+int32(-48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v153
	F_errmsg(m, int32(158908), v15)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	goto L40
L45:
	;
	goto L40
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v67*int32(48))+32))
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v204
	F_errmsg(m, int32(299686), v13+int32(-32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(499138), int32(471), int32(226783))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
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
					v38 = *(*int32)(unsafe.Add(mBase, _consts[194]))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36*int32(48))+32))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v42
					F_errmsg(m, int32(299836), v7)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(499138), int32(877), int32(342487))
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
