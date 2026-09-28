package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileAccess(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[0]))
	v10 = v7 + l0*int32(48)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+16)) = l0
	goto L1
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[1]))
	if v15 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v87 == l0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v60 = F_BasicOpenFilePerm(m, v57, v58, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L15
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[2]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[3]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[4]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[1]))
	if v40 <= int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[3]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[4]))
	if v44 <= v46+(v48+v40) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v60
	if v60 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(-1)
L17:
	;
	goto L18
L18:
	;
	v67 = int32(_a_F_FileAccess_0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_FileAccess[1])) = v69 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_FileAccess[0]))
	v75 = int32(48)
	v77 = v74 + l0*v75
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = l0
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v113 = v74 + v83*v75
	goto L2
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v90 = int32(48)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7+v89*v90)+16)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v7+v93*v90)+20)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v113 = v7 + v104*v90
	goto L2
}
func F_FileClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[0]))
	v17 = v14 + l0*int32(48)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 != int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pgaio_closing_fd(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v80&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	return
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v24 = F_close(m, v23)
	mBase = m.M
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v54 = int32(_a_F_FileClose_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_FileClose[1])) = v56 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(-1)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[0]))
	v64 = int32(48)
	v66 = v63 + l0*v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v63+v67*v64)+16)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v63+v71*v64)+20)) = v67
	goto L3
L7:
	;
	v27 = int32(15)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FileClose[5])))
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = v27
	goto L10
L9:
	;
	v32 = int32(24)
	goto L10
L10:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v33&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = v27
	goto L13
L12:
	;
	v36 = v32
	goto L13
L13:
	;
	v38 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v38 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v42
	F_errmsg_internal(m, int32(_a_F_FileClose_10), v11+int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_FileClose_5), int32(1994), int32(_a_F_FileClose_8))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	v83 = int32(_a_F_FileClose_1)
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_FileClose[2]))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_FileClose[2])) = v85 - v86
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(0)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v80&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v189 != 0 {
		goto L48
	} else {
		goto L49
	}
L22:
	;
	v96 = v80 & int32(_a_F_FileClose_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)) = uint16(v96)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v103 = F___fstatat(m, int32(-100), v98, v11-int32(-64), int32(0))
	mBase = m.M
	goto L23
L23:
	;
	if v103 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[3]))
	v107 = v105
	goto L26
L25:
	;
	v107 = int32(0)
	goto L26
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v109 = F_unlink(m, v108)
	mBase = m.M
	if v109 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v107 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v114 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v114 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v120
	F_errmsg(m, int32(_a_F_FileClose_9), v11+int32(32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_FileClose_5), int32(2039), int32(_a_F_FileClose_8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+88))
	v136 = base.I32_wrap_i64(v135)
	F_pgstat_report_tempfile(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FileClose[3])) = v107
	v167 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L43
	}
L37:
	;
	v140 = int64(*(*int32)(unsafe.Add(mBase, _c_F_FileClose[4])))
	v144 = base.I64_div_s(v135, int64(1024))
	if base.B2i32(v140 < int64(0))|base.B2i32(v144 < v140) != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v149 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v149 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v134
	F_errmsg(m, int32(_a_F_FileClose_4), v11)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_FileClose_5), int32(1530), int32(_a_F_FileClose_6))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	if v167 == int32(0) {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v173
	F_errmsg(m, int32(_a_F_FileClose_7), v11+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_FileClose_5), int32(2049), int32(_a_F_FileClose_8))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L21
L48:
	;
	F_ResourceOwnerForget(m, v189, base.I64_extend_i32_s(l0), int32(_a_F_FileClose_2))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[0]))
	v198 = v195 + l0*int32(48)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
	if v199 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	F_emscripten_builtin_free(m, v199)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v198)+32)) = int32(0)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v203 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+4)) = uint16(v203)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = l0
	m.G0 = v11 + int32(160)
	return
}
func F_FilePathName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_FilePathName[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3+l0*int32(48))+32))
	return v7
}
func F_FileReadV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v8 = F_FileAccess(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_FileReadV[0]))
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_FileReadV[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l4
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17+l0*int32(48))))
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return v40
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_FileReadV[1]))
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	if v40 < v43 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v38 = F_pread(m, v33, v36, v37, l3)
	mBase = m.M
	v40 = v38
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = F_preadv(m, v33, l1, l2, l3)
	mBase = m.M
	v40 = v39
	goto L8
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_FileReadV[2]))
	if v48 == int32(27) {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L7
L15:
	;
	goto L14
}
func F_FileSetDeleteAll(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(2064)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(2064)
	return
L4:
	;
	v25 = v10 + int32(1040)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v20<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v25, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_FileSetDeleteAll_0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v25
	v38 = v10 + int32(16)
	v41 = F_pg_snprintf(m, v38, int32(1024), int32(_a_F_FileSetDeleteAll_1), v10)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = m.G0
	v45 = v43 - int32(96)
	m.G0 = v45
	v49 = F___fstatat(m, int32(-100), v38, v45, int32(0))
	mBase = m.M
	goto L10
L9:
	;
	m.G0 = v45 + int32(96)
	v63 = v20 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v63 < v64 {
		v20 = v63
		goto L4
	} else {
		goto L16
	}
L10:
	;
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetDeleteAll[0]))
	if v51 == int32(44) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_walkdir(m, v38, int32(1180), int32(0), int32(15))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L9
L16:
	;
	goto L5
}
func F_FileTruncate(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v61 int32
	_ = v61
	v6 = F_FileAccess(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+l0*int32(48))))
	goto L7
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	if v28 != 0 {
		v61 = v28
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v28 = F_ftruncate(m, v22, l1)
	mBase = m.M
	if v28 != int32(-1) {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	return int32(-1)
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[2]))
	if v32 == int32(27) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	return v61
L12:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_FileTruncate[1]))
	v50 = v47 + l0*int32(48)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	if v51 <= l1 {
		v61 = v45
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v53 = int32(_a_F_FileTruncate_0)
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_FileTruncate[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_FileTruncate[3])) = v55 + (l1 - v51)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = l1
	v61 = v45
	goto L11
}
func F_FreeFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_FreeFile[0]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_FreeFile[1]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v38 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v19 = v12 + v14*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v14 {
		v14 = v14 - int32(1)
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v21 != l0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = F_FreeDesc(m, v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v23
L11:
	;
	goto L5
L12:
	;
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_FreeFile_0), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v49 = F_fclose(m, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(_a_F_FreeFile_1), int32(2848), int32(_a_F_FreeFile_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return v49
}
func F_read_file_data_into_buffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int64
	_ = v215
	var v247 int64
	_ = v247
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v13 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(167772163)
	if base.Ui32(v22) < base.Ui32(l4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v22
	goto L3
L2:
	;
	v28 = l4
	goto L3
L3:
	;
	v29 = F_pread(m, l2, v21, v28, l3)
	mBase = m.M
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32
	if v32 <= v29 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = int32(0)
	v38 = base.I64_extend_i32_u(v29)
	if base.B2i32(l6 == v36)|base.B2i32(v38&int64(8191) != int64(0)) == v36 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L17
	} else {
		goto L55
	}
L7:
	;
	m.G0 = v19 + int32(48)
	return v247
L8:
	;
	if v29 == int32(0) {
		v247 = v13
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v247 = v38
	goto L7
L11:
	;
	v63 = v13
	goto L12
L12:
	;
	v66 = base.I32_wrap_i64(v63)
	v68 = v66 << (uint(int32(13)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v70 = v68 + v69
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+14)))
	if v71 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L10
L14:
	;
	v215 = v63 + int64(1)
	if base.Ui64(v215) < base.Ui64(int64(base.Ui64(v38)>>(uint(int64(13))%64))) {
		v63 = v215
		goto L12
	} else {
		goto L54
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+32))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	if base.Ui64(v75) <= base.Ui64(base.I64_rotl(v76, int64(32))) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v80 = l5 + v66
	v81 = F_pg_checksum_page(m, v70, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int64(0)
L18:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	if v81 == v85 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = m.G0
	v91 = v89 - int32(32)
	m.G0 = v91
	v93 = int32(_a_F_read_file_data_into_buffer_0)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(167772163)
	v98 = base.I64_extend_i32_s(v68)
	v100 = F_pread(m, l2, v87+v68, int32(_a_F_read_file_data_into_buffer_1), l3+v98)
	mBase = m.M
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103
	if v103 <= v100 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v100 == int32(0) {
		v247 = v98
		goto L7
	} else {
		goto L37
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L33
	}
L22:
	;
	if v100 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L17
	} else {
		goto L29
	}
L25:
	;
	v110 = base.B2i32(v100 != int32(_a_F_read_file_data_into_buffer_1))
	goto L27
L26:
	;
	v110 = int32(0)
	goto L27
L27:
	;
	if v110 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v91 + int32(32)
	goto L20
L29:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_2), v91)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_3), int32(2127), int32(_a_F_read_file_data_into_buffer_4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = int32(_a_F_read_file_data_into_buffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_5), v91+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_3), int32(2132), int32(_a_F_read_file_data_into_buffer_4))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+14)))
	if v151 == int32(0) {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)+32))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	if base.Ui64(v155) <= base.Ui64(base.I64_rotl(v156, int64(32))) {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v160 = F_pg_checksum_page(m, v70, v80)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	if v160 == v162 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v166 = v164 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v166
	if int32(5) < v166 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v172 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	if v172 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_6), v19+int32(32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L17
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	if v189 != int32(5) {
		goto L14
	} else {
		goto L49
	}
L47:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_3), int32(1923), int32(_a_F_read_file_data_into_buffer_7))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L17
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v194 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	if v194 == int32(0) {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_8), v19+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_3), int32(1928), int32(_a_F_read_file_data_into_buffer_7))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	goto L14
L54:
	;
	goto L13
L55:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_2), v19)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_3), int32(2127), int32(_a_F_read_file_data_into_buffer_4))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L17
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
