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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
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
	v32 = int32(23)
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
	F_errfinish(m, int32(_a_F_FileClose_5), int32(2010), int32(_a_F_FileClose_8))
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
	F_errfinish(m, int32(_a_F_FileClose_5), int32(2055), int32(_a_F_FileClose_8))
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
	F_errfinish(m, int32(_a_F_FileClose_5), int32(1546), int32(_a_F_FileClose_6))
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
	F_errfinish(m, int32(_a_F_FileClose_5), int32(2065), int32(_a_F_FileClose_8))
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
	F_ResourceOwnerForget(m, v189, l0, int32(_a_F_FileClose_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_FileClose[0]))
	v197 = v194 + l0*int32(48)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+32))
	if v198 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	F_emscripten_builtin_free(m, v198)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v197)+32)) = int32(0)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v202 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)) = uint16(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = l0
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
	F_walkdir(m, v38, int32(1091), int32(0), int32(15))
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
	F_errfinish(m, int32(_a_F_FreeFile_1), int32(2864), int32(_a_F_FreeFile_2))
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
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int64
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int64
	_ = v518
	var v519 int64
	_ = v519
	var v520 int64
	_ = v520
	var v525 int32
	_ = v525
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v752 int32
	_ = v752
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v941 int64
	_ = v941
	var v973 int64
	_ = v973
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
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
	v982 = m.ExcPending
	if v982 != 0 {
		goto L40
	} else {
		goto L79
	}
L7:
	;
	m.G0 = v19 + int32(48)
	return v973
L8:
	;
	if v29 == int32(0) {
		v973 = v13
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v973 = v38
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
	v941 = v63 + int64(1)
	if base.Ui64(v941) < base.Ui64(int64(base.Ui64(v38)>>(uint(int64(13))%64))) {
		v63 = v941
		goto L12
	} else {
		goto L78
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+32))
	v76 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70)+4)))
	v77 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui64(v75) <= base.Ui64(v76|v77<<(uint(int64(32))%64)) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v82 = l5 + v66
	v83 = int32(0)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v83)
	v119 = m.G0
	v120 = int32(128)
	v121 = v119 - v120
	base.MemoryCopy(m, v121, int32(_a_F_read_file_data_into_buffer_0), v120)
	v128 = v83
	goto L18
L17:
	;
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	if (v441+int32(1))&v440 == v446 {
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v162 = v70 + v128<<(uint(int32(7))%32)
	v169 = int32(0)
	goto L20
L19:
	;
	v239 = int32(0)
	goto L24
L20:
	;
	v199 = int32(2)
	v200 = v169 << (uint(v199) % 32)
	v201 = v121 + v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v162+v200)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v205 = v203 ^ v204
	v206 = int32(16777619)
	v208 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v205*v206 ^ int32(base.Ui32(v205)>>(uint(v208)%32))
	v213 = v200 | int32(4)
	v214 = v121 + v213
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v162+v213)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v218 = v216 ^ v217
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v218*v206 ^ int32(base.Ui32(v218)>>(uint(v208)%32))
	v226 = v169 + v199
	if v226 != int32(32) {
		v169 = v226
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v230 = v128 + int32(1)
	if v230 != int32(64) {
		v128 = v230
		goto L18
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	goto L19
L24:
	;
	v271 = v121 + v239<<(uint(int32(2))%32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v273 = int32(16777619)
	v275 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v272*v273 ^ int32(base.Ui32(v272)>>(uint(v275)%32))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = v279*v273 ^ int32(base.Ui32(v279)>>(uint(v275)%32))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+8)) = v286*v273 ^ int32(base.Ui32(v286)>>(uint(v275)%32))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = v293*v273 ^ int32(base.Ui32(v293)>>(uint(v275)%32))
	v301 = v239 + int32(4)
	if v301 != int32(32) {
		v239 = v301
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v310 = int32(0)
	goto L27
L26:
	;
	goto L25
L27:
	;
	v342 = v121 + v310<<(uint(int32(2))%32)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v344 = int32(16777619)
	v346 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v343*v344 ^ int32(base.Ui32(v343)>>(uint(v346)%32))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+4)) = v350*v344 ^ int32(base.Ui32(v350)>>(uint(v346)%32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+8)) = v357*v344 ^ int32(base.Ui32(v357)>>(uint(v346)%32))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+12)) = v364*v344 ^ int32(base.Ui32(v364)>>(uint(v346)%32))
	v372 = v310 + int32(4)
	if v372 != int32(32) {
		v310 = v372
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v121)+124))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v121)+120))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v121)+116))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v121)+112))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v121)+108))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v121)+104))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v121)+100))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v121)+96))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v121)+92))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v121)+88))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v121)+84))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v121)+80))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v121)+76))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v121)+72))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v121)+68))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v121)+64))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v121)+60))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v121)+52))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v121)+44))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v121)+40))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v121)+24))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v116)
	v440 = int32(_a_F_read_file_data_into_buffer_1)
	v441 = base.I32_rem_u_s(v375^(v376^(v377^(v378^(v379^(v380^(v381^(v382^(v383^(v384^(v385^(v386^(v387^(v388^(v389^(v390^(v391^(v392^(v393^(v394^(v395^(v396^(v397^(v398^(v399^(v400^(v401^(v402^(v403^(v404^(v405^(v82^v406))))))))))))))))))))))))))))))), v440)
	goto L17
L29:
	;
	goto L28
L30:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = m.G0
	v452 = v450 - int32(32)
	m.G0 = v452
	v454 = int32(_a_F_read_file_data_into_buffer_2)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(167772163)
	v459 = base.I64_extend_i32_s(v68)
	v461 = F_pread(m, l2, v448+v68, int32(_a_F_read_file_data_into_buffer_3), l3+v459)
	mBase = m.M
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_read_file_data_into_buffer[0]))
	v464 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v464
	if v464 <= v461 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v461 == int32(0) {
		v973 = v459
		goto L7
	} else {
		goto L49
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L40
	} else {
		goto L45
	}
L33:
	;
	if v461 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v471 = base.B2i32(v461 != int32(_a_F_read_file_data_into_buffer_3))
	goto L38
L37:
	;
	v471 = int32(0)
	goto L38
L38:
	;
	if v471 != 0 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v452 + int32(32)
	goto L31
L40:
	;
	return int64(0)
L41:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_4), v452)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_5), int32(2128), int32(_a_F_read_file_data_into_buffer_6))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452)+24)) = int32(_a_F_read_file_data_into_buffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v452)+20)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v452)+16)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_7), v452+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_5), int32(2133), int32(_a_F_read_file_data_into_buffer_6))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L40
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
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+14)))
	if v514 == int32(0) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v518 = *(*int64)(unsafe.Add(mBase, uint32(v517)+32))
	v519 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70)+4)))
	v520 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui64(v518) <= base.Ui64(v519|v520<<(uint(int64(32))%64)) {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	v525 = int32(0)
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v525)
	v561 = m.G0
	v562 = int32(128)
	v563 = v561 - v562
	base.MemoryCopy(m, v563, int32(_a_F_read_file_data_into_buffer_0), v562)
	v570 = v525
	goto L53
L52:
	;
	v888 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	if v887 == v888 {
		goto L14
	} else {
		goto L65
	}
L53:
	;
	v604 = v70 + v570<<(uint(int32(7))%32)
	v611 = int32(0)
	goto L55
L54:
	;
	v681 = int32(0)
	goto L59
L55:
	;
	v641 = int32(2)
	v642 = v611 << (uint(v641) % 32)
	v643 = v563 + v642
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v604+v642)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	v647 = v645 ^ v646
	v648 = int32(16777619)
	v650 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v643))) = v647*v648 ^ int32(base.Ui32(v647)>>(uint(v650)%32))
	v655 = v642 | int32(4)
	v656 = v563 + v655
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v604+v655)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v660 = v658 ^ v659
	*(*int32)(unsafe.Add(mBase, uint32(v656))) = v660*v648 ^ int32(base.Ui32(v660)>>(uint(v650)%32))
	v668 = v611 + v641
	if v668 != int32(32) {
		v611 = v668
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v672 = v570 + int32(1)
	if v672 != int32(64) {
		v570 = v672
		goto L53
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	goto L54
L59:
	;
	v713 = v563 + v681<<(uint(int32(2))%32)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	v715 = int32(16777619)
	v717 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = v714*v715 ^ int32(base.Ui32(v714)>>(uint(v717)%32))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v721*v715 ^ int32(base.Ui32(v721)>>(uint(v717)%32))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = v728*v715 ^ int32(base.Ui32(v728)>>(uint(v717)%32))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+12)) = v735*v715 ^ int32(base.Ui32(v735)>>(uint(v717)%32))
	v743 = v681 + int32(4)
	if v743 != int32(32) {
		v681 = v743
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v752 = int32(0)
	goto L62
L61:
	;
	goto L60
L62:
	;
	v784 = v563 + v752<<(uint(int32(2))%32)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v786 = int32(16777619)
	v788 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v784))) = v785*v786 ^ int32(base.Ui32(v785)>>(uint(v788)%32))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v784)+4)) = v792*v786 ^ int32(base.Ui32(v792)>>(uint(v788)%32))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v784)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v784)+8)) = v799*v786 ^ int32(base.Ui32(v799)>>(uint(v788)%32))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v784)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v784)+12)) = v806*v786 ^ int32(base.Ui32(v806)>>(uint(v788)%32))
	v814 = v752 + int32(4)
	if v814 != int32(32) {
		v752 = v814
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v563)+124))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v563)+120))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v563)+116))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v563)+112))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v563)+108))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v563)+104))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v563)+100))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v563)+96))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v563)+92))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v563)+88))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v563)+84))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v563)+80))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v563)+76))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v563)+72))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v563)+68))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v563)+64))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v563)+60))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v563)+56))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v563)+52))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v563)+48))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v563)+44))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v563)+40))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v563)+36))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v563)+32))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v563)+28))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v563)+24))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v563)+20))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v563)+16))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v558)
	v882 = int32(_a_F_read_file_data_into_buffer_1)
	v883 = base.I32_rem_u_s(v817^(v818^(v819^(v820^(v821^(v822^(v823^(v824^(v825^(v826^(v827^(v828^(v829^(v830^(v831^(v832^(v833^(v834^(v835^(v836^(v837^(v838^(v839^(v840^(v841^(v842^(v843^(v844^(v845^(v846^(v847^(v82^v848))))))))))))))))))))))))))))))), v882)
	v887 = (v883 + int32(1)) & v882
	goto L52
L64:
	;
	goto L63
L65:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v892 = v890 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v892
	if int32(5) < v892 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	v898 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L40
	} else {
		goto L67
	}
L67:
	;
	if v898 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v900 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v887
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_8), v19+int32(32))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L40
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	if v915 != int32(5) {
		goto L14
	} else {
		goto L73
	}
L71:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_5), int32(1924), int32(_a_F_read_file_data_into_buffer_9))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L40
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v920 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L40
	} else {
		goto L74
	}
L74:
	;
	if v920 == int32(0) {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_10), v19+int32(16))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L40
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_5), int32(1929), int32(_a_F_read_file_data_into_buffer_9))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L40
	} else {
		goto L77
	}
L77:
	;
	goto L14
L78:
	;
	goto L13
L79:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L40
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	F_errmsg(m, int32(_a_F_read_file_data_into_buffer_4), v19)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L40
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_read_file_data_into_buffer_5), int32(2128), int32(_a_F_read_file_data_into_buffer_6))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L40
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
