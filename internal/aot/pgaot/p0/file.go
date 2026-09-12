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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	v7 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v10 = v7 + l0*int32(48)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	if v15 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if l0 != v90 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v60 = F_BasicOpenFilePerm(m, v57, v58, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L13
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	v21 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	if v21+(v23+v15) < v19 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return int32(0)
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	if v40 <= int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	if v44 <= v46+(v48+v40) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v60
	if v60 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(-1)
L15:
	;
	goto L16
L16:
	;
	v67 = int32(4470128)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	*(*int32)(unsafe.Add(mBase, _consts[586])) = v69 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v75 = int32(48)
	v77 = v74 + l0*v75
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = l0
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v83*v75)+16)) = l0
	return v78
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v93 = int32(48)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7+v92*v93)+16)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v7+v96*v93)+20)) = v92
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7+v107*v93)+16)) = l0
	v115 = v102
	goto L19
L18:
	;
	v115 = int32(0)
	goto L19
L19:
	;
	return v115
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[191]))
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
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v82&int32(4) != 0 {
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
	v56 = int32(4470128)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	*(*int32)(unsafe.Add(mBase, _consts[586])) = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(-1)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v66 = int32(48)
	v68 = v65 + l0*v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v69*v66)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v65+v73*v66)+20)) = v69
	goto L3
L7:
	;
	v27 = int32(15)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
	if v31&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = v27
	goto L10
L9:
	;
	v34 = int32(23)
	goto L10
L10:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v35&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v27
	goto L13
L12:
	;
	v38 = v34
	goto L13
L13:
	;
	v40 = F_errstart(m, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v44
	F_errmsg_internal(m, int32(312979), v11+int32(48))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(523542), int32(2010), int32(378432))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	v85 = int32(4470136)
	v87 = *(*int64)(unsafe.Add(mBase, _consts[587]))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, _consts[587])) = v87 - v88
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(0)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v82&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v190 != 0 {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	v98 = v82 & int32(65534)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)) = uint16(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v105 = F___fstatat(m, int32(-100), v100, v11-int32(-64), int32(0))
	mBase = m.M
	goto L23
L23:
	;
	if v105 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v109 = v107
	goto L26
L25:
	;
	v109 = int32(0)
	goto L26
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v111 = F_unlink(m, v110)
	mBase = m.M
	if v111 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v109 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v116 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v116 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
	F_errmsg(m, int32(312851), v11+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(523542), int32(2055), int32(378432))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v11)+88))
	v138 = base.I32_wrap_i64(v137)
	F_pgstat_report_tempfile(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v109
	v168 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L44
	}
L37:
	;
	v142 = base.I64_div_s(v137, int64(1024))
	v144 = int64(*(*int32)(unsafe.Add(mBase, _consts[588])))
	if v144 < int64(0) {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	if v142 < v144 {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	v150 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v150 == int32(0) {
		goto L21
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v136
	F_errmsg(m, int32(39907), v11)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(523542), int32(1546), int32(423830))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L21
L44:
	;
	if v168 == int32(0) {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v174
	F_errmsg(m, int32(311338), v11+int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(523542), int32(2065), int32(378432))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L21
L49:
	;
	F_ResourceOwnerForget(m, v190, l0, int32(1661152))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v198 = v195 + l0*int32(48)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
	if v199 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	F_emscripten_builtin_free(m, v199)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v198)+32)) = int32(0)
	goto L55
L54:
	;
	goto L55
L55:
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[191]))
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
	v17 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
	v48 = *(*int32)(unsafe.Add(mBase, _consts[163]))
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
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	v19 = v2
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v19<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v10+int32(1040), v29)
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(246818)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(1040)
	v43 = F_pg_snprintf(m, v10+int32(16), int32(1024), int32(112863), v10)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v45 = m.G0
	v47 = v45 - int32(96)
	m.G0 = v47
	v50 = v10 + int32(16)
	v53 = F___fstatat(m, int32(-100), v50, v47, int32(0))
	mBase = m.M
	goto L10
L9:
	;
	m.G0 = v47 + int32(96)
	v67 = v19 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 < v68 {
		v19 = v67
		goto L4
	} else {
		goto L16
	}
L10:
	;
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v55 == int32(44) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_walkdir(m, v50, int32(1091), int32(0), int32(15))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	v18 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+l0*int32(48))))
	goto L7
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	if v28 != 0 {
		v62 = v28
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	return int32(-1)
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[163]))
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
	return v62
L12:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v52 = v47 + l0*int32(48) + int32(24)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	if v53 <= l1 {
		v62 = v45
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(4470136)
	v57 = *(*int64)(unsafe.Add(mBase, _consts[587]))
	*(*int64)(unsafe.Add(mBase, _consts[587])) = v57 + (l1 - v53)
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = l1
	v62 = v45
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[584]))
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
	F_errmsg_internal(m, int32(408009), int32(0))
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
	F_errfinish(m, int32(523542), int32(2864), int32(408125))
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
	var v36 int64
	_ = v36
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v214 int64
	_ = v214
	var v246 int64
	_ = v246
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v13 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32
	if v32 <= v29 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = base.I64_extend_i32_u(v29)
	if l6 == int32(0) {
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
	v255 = m.ExcPending
	if v255 != 0 {
		goto L27
	} else {
		goto L53
	}
L7:
	;
	m.G0 = v19 + int32(48)
	return v246
L8:
	;
	v246 = v36
	goto L7
L9:
	;
	if v36&int64(8191) != int64(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v29 == int32(0) {
		v246 = v13
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v60 = v13
	goto L12
L12:
	;
	v63 = base.I32_wrap_i64(v60)
	v65 = v63 << (uint(int32(13)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = v65 + v66
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+14)))
	if v68 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L8
L14:
	;
	v214 = v60 + int64(1)
	if base.Ui64(v214) < base.Ui64(int64(base.Ui64(v36)>>(uint(int64(13))%64))) {
		v60 = v214
		goto L12
	} else {
		goto L52
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
	v73 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+4)))
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67))))
	if base.Ui64(v72) <= base.Ui64(v73|v74<<(uint(int64(32))%64)) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v79 = l5 + v63
	v80 = F_pg_checksum_page(m, v67, v79)
	mBase = m.M
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	if v80 == v81 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v85 = m.G0
	v87 = v85 - int32(32)
	m.G0 = v87
	v89 = int32(4160268)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(167772163)
	v94 = base.I64_extend_i32_s(v65)
	v96 = F_pread(m, l2, v83+v65, int32(8192), l3+v94)
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99
	if v99 <= v96 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v96 == int32(0) {
		v246 = v94
		goto L7
	} else {
		goto L36
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L27
	} else {
		goto L32
	}
L20:
	;
	if v96 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v106 = base.B2i32(v96 != int32(8192))
	goto L25
L24:
	;
	v106 = int32(0)
	goto L25
L25:
	;
	if v106 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v87 + int32(32)
	goto L18
L27:
	;
	return int64(0)
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = l1
	F_errmsg(m, int32(313155), v87)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(518729), int32(2128), int32(404786))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L27
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = l1
	F_errmsg(m, int32(38624), v87+int32(16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(518729), int32(2133), int32(404786))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+14)))
	if v149 == int32(0) {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)+32))
	v154 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+4)))
	v155 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67))))
	if base.Ui64(v153) <= base.Ui64(v154|v155<<(uint(int64(32))%64)) {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v160 = F_pg_checksum_page(m, v67, v79)
	mBase = m.M
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	if v160 == v161 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v165 = v163 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v165
	if int32(5) < v165 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v171 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	if v171 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l1
	F_errmsg(m, int32(540561), v19+int32(32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L27
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	if v188 != int32(5) {
		goto L14
	} else {
		goto L47
	}
L45:
	;
	F_errfinish(m, int32(518729), int32(1924), int32(236298))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v193 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	if v193 == int32(0) {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	F_errmsg(m, int32(464473), v19+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(518729), int32(1929), int32(236298))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L27
	} else {
		goto L51
	}
L51:
	;
	goto L14
L52:
	;
	goto L13
L53:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L27
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	F_errmsg(m, int32(313155), v19)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L27
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(518729), int32(2128), int32(404786))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L27
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
