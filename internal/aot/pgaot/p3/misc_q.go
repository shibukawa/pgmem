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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	return v93
L4:
	;
	F_AccessTempTableNamespace(m, base.B2i32(v14 == int32(0)))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v15 = int32(235520)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[438])))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
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
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L8:
	;
	if v39-v38 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v23 = v14
	v24 = v15
	goto L12
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v38 = v27
	v39 = v28
	goto L9
L14:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v44 = int32(0)
	v47 = F_GetSysCacheOid(m, int32(37), v14, v44, v44, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v47 != 0 {
		v93 = v47
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
	F_errmsg(m, int32(72674), v6)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(499480), int32(3547), int32(434898))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[439])))
	if v68 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[440]))
	if v70 != 0 {
		v93 = v70
		goto L3
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(280094), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(499480), int32(3525), int32(418490))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v93 = v92
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
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
	return v147
L4:
	;
	if v32 != 0 {
		v147 = v3
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	v66 = base.I32_div_u_s(int32(1073741815)-v63, int32(12))
	if base.Ui32(v66) < base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v37 == int32(0) {
		v147 = v3
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
	F_errmsg(m, int32(212102), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errdetail(m, int32(573270), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, l1, int32(496046), int32(367), int32(275078))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v147 = v3
	goto L3
L15:
	;
	v68 = F_errsave_start(m, l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v92 = v34*int32(12) + v63 + int32(8)
	v93 = F_palloc0(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	if v68 == int32(0) {
		v147 = v3
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(400188), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, l1, int32(496046), int32(372), int32(275078))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v147 = v3
	goto L3
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v92 << (uint(int32(2)) % 32)
	v100 = v93 + int32(8)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v102 = v101
	v108 = v3
	goto L24
L24:
	;
	v112 = v100 + v108*int32(12)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112))) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v115
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+10)) = uint16(v117)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+9)) = uint8(v119)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+8)) = uint8(v121)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	F_pfree(m, v102)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v127 = v108 + int32(1)
	if v127 != v34 {
		v102 = v123
		v108 = v127
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	F_pfree(m, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v134 = F__emscripten_memcpy_bulkmem(m, v100+v129*int32(12), v133, v63)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	F_findoprnd_1(m, v100, v11+int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v147 = v93
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
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
	v19 = l0 + v10
	if base.Ui32(v19) <= base.Ui32(l0) {
		v41 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v49 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v49)
	v52 = v41 + int32(1)
	if v10 == int32(0) {
		v179 = v41
		v180 = v52
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v22 = l0
	goto L6
L5:
	;
	v36 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v36)
	v41 = v15 + int32(1)
	goto L3
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v30 == int32(92) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v41 = v15
	goto L3
L8:
	;
	v34 = v22 + int32(1)
	if v34 != v19 {
		v22 = v34
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v187 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v187)
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)) = uint8(v189)
	return v15
L11:
	;
	v56 = v10 & int32(3)
	if v56 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if base.Ui32(v10) < base.Ui32(int32(4)) {
		v179 = v90
		v180 = v91
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v89 = l0
	v90 = v41
	v91 = v52
	v95 = v10
	goto L12
L14:
	;
	goto L15
L15:
	;
	v59 = l0
	v60 = v41
	v61 = v52
	v64 = int32(0)
	v65 = v10
	goto L16
L16:
	;
	v69 = v65 - int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v70 == int32(92) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v89 = v85
	v90 = v80
	v91 = v83
	v95 = v69
	goto L12
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v79)
	v82 = int32(1)
	v83 = v80 + v82
	v85 = v59 + v82
	v87 = v64 + v82
	if v87 != v56 {
		v59 = v85
		v60 = v80
		v61 = v83
		v64 = v87
		v65 = v69
		goto L16
	} else {
		goto L22
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v70)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v79 = v76
	v80 = v60 + int32(2)
	goto L18
L20:
	;
	if v70 == int32(39) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = v70
	v80 = v61
	goto L18
L22:
	;
	goto L17
L23:
	;
	v100 = v89
	v101 = v90
	v102 = v91
	v106 = v95
	goto L24
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if base.B2i32(v109 != int32(92))&base.B2i32(v109 != int32(39)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v179 = v170
	v180 = v173
	goto L10
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v109)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v121 = v101 + int32(2)
	v122 = v118
	goto L28
L27:
	;
	v121 = v102
	v122 = v109
	goto L28
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v122)
	v125 = v100 + int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v126 == int32(92) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v137)
	v141 = v100 + int32(2)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142 == int32(92) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)) = uint8(v126)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v137 = v134
	v138 = v121 + int32(2)
	goto L29
L31:
	;
	if v126 == int32(39) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v137 = v126
	v138 = v121 + int32(1)
	goto L29
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v153)
	v157 = v100 + int32(3)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v158 == int32(92) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)) = uint8(v142)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v153 = v150
	v154 = v138 + int32(2)
	goto L33
L35:
	;
	if v142 == int32(39) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v153 = v142
	v154 = v138 + int32(1)
	goto L33
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v169)
	v173 = v170 + int32(1)
	v174 = int32(4)
	v177 = v106 - v174
	if v177 != 0 {
		v100 = v100 + v174
		v101 = v170
		v102 = v173
		v106 = v177
		goto L24
	} else {
		goto L41
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)) = uint8(v158)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v169 = v166
	v170 = v154 + int32(2)
	goto L37
L39:
	;
	if v158 == int32(39) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v169 = v158
	v170 = v154 + int32(1)
	goto L37
L41:
	;
	goto L25
}
