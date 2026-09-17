package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QualifiedNameGetCreationNamespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_DeconstructQualifiedName(m, l0, v6+int32(12), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v6 + int32(16)
	return v94
L4:
	;
	F_AccessTempTableNamespace(m, base.B2i32(v14 == int32(0)))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L5:
	;
	v15 = int32(_a_F_QualifiedNameGetCreationNamespace_0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_QualifiedNameGetCreationNamespace[0])))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	if v39-v40 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v24 = v14
	v25 = v15
	goto L11
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v39 = v29
	v40 = v28
	goto L9
L13:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = int32(0)
	v48 = F_GetSysCacheOid(m, int32(37), v14, v45, v45, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v48 != 0 {
		v94 = v48
		goto L3
	} else {
		goto L17
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
	F_errmsg(m, int32(_a_F_QualifiedNameGetCreationNamespace_1), v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_QualifiedNameGetCreationNamespace_2), int32(3547), int32(_a_F_QualifiedNameGetCreationNamespace_3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
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
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_QualifiedNameGetCreationNamespace[1])))
	if v69 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_QualifiedNameGetCreationNamespace[2]))
	if v71 != 0 {
		v94 = v71
		goto L3
	} else {
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_QualifiedNameGetCreationNamespace_4), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_QualifiedNameGetCreationNamespace_2), int32(3525), int32(_a_F_QualifiedNameGetCreationNamespace_5))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
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
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_QualifiedNameGetCreationNamespace[3]))
	v94 = v93
	goto L3
}
func F_queryin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(64)
	v22 = F_palloc(m, int32(64))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v22
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v28)
	v32 = F_makepol_2(m, v11+int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v11 + int32(48)
	return v138
L4:
	;
	if v32 != 0 {
		v138 = v3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = F_errsave_start(m, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	v61 = base.I32_div_u_s(int32(1073741815)-v58, int32(12))
	if base.Ui32(v61) < base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v37 == int32(0) {
		v138 = v3
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errmsg(m, int32(_a_F_queryin_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errdetail(m, int32(_a_F_queryin_1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, l1, int32(_a_F_queryin_2), int32(367), int32(_a_F_queryin_3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v138 = v3
	goto L3
L15:
	;
	v63 = F_errsave_start(m, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v83 = v34*int32(12) + v58 + int32(8)
	v84 = F_palloc0(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	if v63 == int32(0) {
		v138 = v3
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_queryin_4), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, l1, int32(_a_F_queryin_2), int32(372), int32(_a_F_queryin_3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v138 = v3
	goto L3
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v83 << (uint(int32(2)) % 32)
	v91 = v84 + int32(8)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v93 = v92
	v100 = v3
	goto L24
L24:
	;
	v103 = v91 + v100*int32(12)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*uint16)(unsafe.Add(mBase, uint32(v103))) = uint16(v104)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v106
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+10)) = uint16(v108)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+9)) = uint8(v110)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)) = uint8(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	F_pfree(m, v93)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v58 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v118 = v100 + int32(1)
	if v118 != v34 {
		v93 = v114
		v100 = v118
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	base.MemoryCopy(m, v91+v121*int32(12), v120, v58)
	goto L30
L29:
	;
	goto L30
L30:
	;
	F_pfree(m, v120)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	F_findoprnd_1(m, v91, v11+int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v138 = v84
	goto L3
}
func F_quote_literal_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v10 = F_strlen(m, l0)
	mBase = m.M
	v15 = F_palloc(m, v10<<(uint(int32(1))%32)+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v179 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v179)
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+2)) = uint8(v181)
	return v15
L4:
	;
	v21 = l0
	goto L8
L5:
	;
	goto L6
L6:
	;
	v166 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v166)
	v171 = v15
	v172 = v15 + int32(1)
	goto L3
L7:
	;
	v41 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
	v44 = v40 + int32(1)
	v46 = v10 & int32(3)
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v29 == int32(92) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v40 = v15
	goto L7
L10:
	;
	v32 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v32)
	v40 = v15 + int32(1)
	goto L7
L11:
	;
	goto L12
L12:
	;
	v37 = v21 + int32(1)
	if base.Ui32(v37) < base.Ui32(l0+v10) {
		v21 = v37
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	if base.Ui32(v10) < base.Ui32(int32(4)) {
		v171 = v81
		v172 = v82
		goto L3
	} else {
		goto L24
	}
L15:
	;
	v80 = l0
	v81 = v40
	v82 = v44
	v86 = v10
	goto L14
L16:
	;
	goto L17
L17:
	;
	v49 = l0
	v50 = v40
	v51 = v44
	v55 = v10
	v57 = int32(0)
	goto L18
L18:
	;
	v59 = v55 - int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if base.B2i32(v60 == int32(92))|base.B2i32(v60 == int32(39)) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v80 = v76
	v81 = v71
	v82 = v74
	v86 = v59
	goto L14
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v60)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v70 = v67
	v71 = v50 + int32(2)
	goto L22
L21:
	;
	v70 = v60
	v71 = v51
	goto L22
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v70)
	v73 = int32(1)
	v74 = v71 + v73
	v76 = v49 + v73
	v78 = v57 + v73
	if v78 != v46 {
		v49 = v76
		v50 = v71
		v51 = v74
		v55 = v59
		v57 = v78
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v91 = v80
	v92 = v81
	v93 = v82
	v97 = v86
	goto L25
L25:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if base.B2i32(v100 != int32(92))&base.B2i32(v100 != int32(39)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v171 = v158
	v172 = v161
	goto L3
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v100)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v112 = v92 + int32(2)
	v113 = v109
	goto L29
L28:
	;
	v112 = v93
	v113 = v100
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v113)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if base.B2i32(v115 == int32(92))|base.B2i32(v115 == int32(39)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)) = uint8(v115)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	v127 = v122
	v128 = v112 + int32(2)
	goto L32
L31:
	;
	v127 = v115
	v128 = v112 + int32(1)
	goto L32
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v127)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
	if base.B2i32(v130 == int32(92))|base.B2i32(v130 == int32(39)) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v130)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
	v142 = v137
	v143 = v128 + int32(2)
	goto L35
L34:
	;
	v142 = v130
	v143 = v128 + int32(1)
	goto L35
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v142)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
	if base.B2i32(v145 == int32(92))|base.B2i32(v145 == int32(39)) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)) = uint8(v145)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
	v157 = v152
	v158 = v143 + int32(2)
	goto L38
L37:
	;
	v157 = v145
	v158 = v143 + int32(1)
	goto L38
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v157)
	v161 = v158 + int32(1)
	v162 = int32(4)
	v165 = v97 - v162
	if v165 != 0 {
		v91 = v91 + v162
		v92 = v158
		v93 = v161
		v97 = v165
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L26
}
