package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGroup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 float64
	_ = v92
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 float64
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v18 != 0 {
		v198 = v2
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v10 + int32(16)
	return v198
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L33
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v21&int32(2) == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v27 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	F_ExecReScan(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	m.T0[v41].(func(*base.Module, int32, int32))(m, v20, v31)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L23
	}
L18:
	;
	if v31 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	if v33&int32(2) == int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v38)
	v198 = v2
	goto L6
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v20
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v45 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v89 == int32(0) {
		goto L8
	} else {
		goto L32
	}
L25:
	;
	v46 = int32(4449520)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v45, v19, v10+int32(13))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	m.T0[v66].(func(*base.Module, int32))(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v47
	if v54 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v69 = int32(4449520)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	v78 = m.T0[v77].(func(*base.Module, int32, int32, int32) int32)(m, v62+int32(4), v63, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v70
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)))
	v84 = v82 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v87)
	v198 = v64
	goto L6
L32:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v89)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v89)+240)) = base.F64_add(v92, float64(1))
	goto L8
L33:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+52))
	if v106 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_ExecReScan(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v110 = m.T0[v109].(func(*base.Module, int32) int32)(m, v105)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v20
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v122 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if v110 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
	if v112&int32(2) == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v117)
	v198 = int32(0)
	goto L6
L44:
	;
	goto L43
L45:
	;
	v123 = int32(4449520)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	v131 = m.T0[v130].(func(*base.Module, int32, int32, int32) int32)(m, v122, v19, v10+int32(14))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L61
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v131 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+32))
	m.T0[v139].(func(*base.Module, int32, int32))(m, v20, v110)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v20
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v143 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v187 == int32(0) {
		goto L33
	} else {
		goto L60
	}
L53:
	;
	v144 = int32(4449520)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v147
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v152 = m.T0[v151].(func(*base.Module, int32, int32, int32) int32)(m, v143, v19, v10+int32(15))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+72))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	m.T0[v164].(func(*base.Module, int32))(m, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v145
	if v152 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v167 = int32(4449520)
	v168 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v170
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v160)+24))
	v176 = m.T0[v175].(func(*base.Module, int32, int32, int32) int32)(m, v160+int32(4), v161, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v168
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162)+4)))
	v182 = v180 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+4)) = uint16(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+6)) = uint16(v185)
	v198 = v162
	goto L6
L60:
	;
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v187)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v187)+240)) = base.F64_add(v190, float64(1))
	goto L33
L61:
	;
	goto L33
}
func F_create_group_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v121 int32
	_ = v121
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v145 float64
	_ = v145
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v160 float64
	_ = v160
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
	v15 = F_palloc0(m, int32(88))
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(305)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v22 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(364)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v30 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v35 = v33
	goto L5
L4:
	;
	v35 = int32(0)
	goto L5
L5:
	;
	v37 = v35 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v41
	if l3 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v48 = v46
	goto L8
L7:
	;
	v48 = int32(0)
	goto L8
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v53 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(32)
	m.G0 = v56
	v59 = *(*float64)(unsafe.Add(mBase, _consts[482]))
	v63 = base.F64_add(base.F64_mul(base.F64_mul(v52, v59), base.F64_convert_i32_s(v48)), v51)
	if l4 == v53 {
		v145 = l5
		v150 = v63
		v151 = v50
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v150
	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v49
	*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v145
	m.G0 = v56 + int32(32)
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_add(v160, v151)
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(v150, base.F64_add(base.F64_mul(v165, v145), v168))
	return v15
L10:
	;
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+24)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = l0
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v71 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v115 = float64(0)
	v116 = float64(0)
	goto L13
L12:
	;
	v77 = v53
	goto L14
L13:
	;
	v117 = base.F64_add(v50, v116)
	v120 = base.F64_add(v63, base.F64_add(base.F64_mul(l5, v115), v116))
	v121 = int32(0)
	v124 = F_clauselist_selectivity(m, l0, l4, v121, v121, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v77<<(uint(int32(2))%32))))
	v95 = F_cost_qual_eval_walker(m, v92, v56+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
	v115 = v101
	v116 = v102
	goto L13
L16:
	;
	v98 = v77 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v98 < v99 {
		v77 = v98
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v126 = base.F64_mul(l5, v124)
	if base.F64_gt(v126, float64(1e+100)) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v145 = float64(1e+100)
	v150 = v120
	v151 = v117
	goto L9
L20:
	;
	goto L21
L21:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v145 = float64(1e+100)
	v150 = v120
	v151 = v117
	goto L9
L23:
	;
	goto L24
L24:
	;
	v136 = float64(1)
	if base.F64_le(v126, v136) != 0 {
		v145 = v136
		v150 = v120
		v151 = v117
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v145 = base.F64_nearest(v126)
	v150 = v120
	v151 = v117
	goto L9
}
func F_flatten_group_exprs_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		v140 = v3
		m.G0 = v8 + int32(16)
		return v140
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v12 - int32(6) {
		case 0:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v15 != v16 {
				v140 = l0
				m.G0 = v8 + int32(16)
				return v140
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21<<(uint(int32(2))%32)-int32(4))))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
				if v28 != int32(9) {
					v140 = l0
					m.G0 = v8 + int32(16)
					return v140
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+120))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)-int32(4))))
					v40 = F_copyObjectImpl(m, v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v44 != 0 {
							F_IncrementVarSublevelsUp(m, v40, v44, int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								if v48 == int32(6) {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v51
								} else {
								}
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
								if v53 != int32(1) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v60 == int32(0) {
										v140 = v40
										m.G0 = v8 + int32(16)
										return v140
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v63 == int32(0) {
											v140 = v40
											m.G0 = v8 + int32(16)
											return v140
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
											*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
											v69 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
											v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
												if v77 != 0 {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v79 = F_add_nulling_relids(m, v40, v77, v78)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v140 = v79
														m.G0 = v8 + int32(16)
														return v140
													}
												} else {
													v81 = F_contain_volatile_functions(m, v40)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														if v81 != 0 {
															v140 = v40
															m.G0 = v8 + int32(16)
															return v140
														} else {
															v83 = F_expression_returns_set(m, v40)
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																if v83 != 0 {
																	v140 = v40
																	m.G0 = v8 + int32(16)
																	return v140
																} else {
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																	v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int32(0)
																	} else {
																		v91 = F_make_placeholder_expr(m, v60, v40, v89)
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int32(0)
																		} else {
																			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																			*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																			v96 = F_bms_copy(m, v95)
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																				v140 = v91
																				m.G0 = v8 + int32(16)
																				return v140
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
									}
								} else {
									v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
									if v56 != 0 {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v60 == int32(0) {
											v140 = v40
											m.G0 = v8 + int32(16)
											return v140
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											if v63 == int32(0) {
												v140 = v40
												m.G0 = v8 + int32(16)
												return v140
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
												*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
												v69 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
												v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
													if v77 != 0 {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v79 = F_add_nulling_relids(m, v40, v77, v78)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v140 = v79
															m.G0 = v8 + int32(16)
															return v140
														}
													} else {
														v81 = F_contain_volatile_functions(m, v40)
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
															return int32(0)
														} else {
															if v81 != 0 {
																v140 = v40
																m.G0 = v8 + int32(16)
																return v140
															} else {
																v83 = F_expression_returns_set(m, v40)
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	if v83 != 0 {
																		v140 = v40
																		m.G0 = v8 + int32(16)
																		return v140
																	} else {
																		v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																		v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return int32(0)
																		} else {
																			v91 = F_make_placeholder_expr(m, v60, v40, v89)
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int32(0)
																			} else {
																				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																				*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																				v96 = F_bms_copy(m, v95)
																				mBase = m.M
																				v97 = m.ExcPending
																				if v97 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																					v140 = v91
																					m.G0 = v8 + int32(16)
																					return v140
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
										}
									} else {
										v57 = F_checkExprHasSubLink(m, v40)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v57)
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											if v60 == int32(0) {
												v140 = v40
												m.G0 = v8 + int32(16)
												return v140
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												if v63 == int32(0) {
													v140 = v40
													m.G0 = v8 + int32(16)
													return v140
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
													*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
													v69 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
													v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
														if v77 != 0 {
															v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															v79 = F_add_nulling_relids(m, v40, v77, v78)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																v140 = v79
																m.G0 = v8 + int32(16)
																return v140
															}
														} else {
															v81 = F_contain_volatile_functions(m, v40)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return int32(0)
															} else {
																if v81 != 0 {
																	v140 = v40
																	m.G0 = v8 + int32(16)
																	return v140
																} else {
																	v83 = F_expression_returns_set(m, v40)
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return int32(0)
																	} else {
																		if v83 != 0 {
																			v140 = v40
																			m.G0 = v8 + int32(16)
																			return v140
																		} else {
																			v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																			v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																			v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																			mBase = m.M
																			v90 = m.ExcPending
																			if v90 != 0 {
																				return int32(0)
																			} else {
																				v91 = F_make_placeholder_expr(m, v60, v40, v89)
																				mBase = m.M
																				v92 = m.ExcPending
																				if v92 != 0 {
																					return int32(0)
																				} else {
																					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																					*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																					v96 = F_bms_copy(m, v95)
																					mBase = m.M
																					v97 = m.ExcPending
																					if v97 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																						v140 = v91
																						m.G0 = v8 + int32(16)
																						return v140
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
											}
										}
									}
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							if v48 == int32(6) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v51
							} else {
							}
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
							if v53 != int32(1) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v60 == int32(0) {
									v140 = v40
									m.G0 = v8 + int32(16)
									return v140
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v63 == int32(0) {
										v140 = v40
										m.G0 = v8 + int32(16)
										return v140
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
										v69 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
										v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
											if v77 != 0 {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v79 = F_add_nulling_relids(m, v40, v77, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v140 = v79
													m.G0 = v8 + int32(16)
													return v140
												}
											} else {
												v81 = F_contain_volatile_functions(m, v40)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													if v81 != 0 {
														v140 = v40
														m.G0 = v8 + int32(16)
														return v140
													} else {
														v83 = F_expression_returns_set(m, v40)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															if v83 != 0 {
																v140 = v40
																m.G0 = v8 + int32(16)
																return v140
															} else {
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																mBase = m.M
																v90 = m.ExcPending
																if v90 != 0 {
																	return int32(0)
																} else {
																	v91 = F_make_placeholder_expr(m, v60, v40, v89)
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int32(0)
																	} else {
																		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																		v96 = F_bms_copy(m, v95)
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																			v140 = v91
																			m.G0 = v8 + int32(16)
																			return v140
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
								}
							} else {
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
								if v56 != 0 {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v60 == int32(0) {
										v140 = v40
										m.G0 = v8 + int32(16)
										return v140
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v63 == int32(0) {
											v140 = v40
											m.G0 = v8 + int32(16)
											return v140
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
											*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
											v69 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
											v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
												if v77 != 0 {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v79 = F_add_nulling_relids(m, v40, v77, v78)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v140 = v79
														m.G0 = v8 + int32(16)
														return v140
													}
												} else {
													v81 = F_contain_volatile_functions(m, v40)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														if v81 != 0 {
															v140 = v40
															m.G0 = v8 + int32(16)
															return v140
														} else {
															v83 = F_expression_returns_set(m, v40)
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																if v83 != 0 {
																	v140 = v40
																	m.G0 = v8 + int32(16)
																	return v140
																} else {
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																	v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int32(0)
																	} else {
																		v91 = F_make_placeholder_expr(m, v60, v40, v89)
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int32(0)
																		} else {
																			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																			*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																			v96 = F_bms_copy(m, v95)
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																				v140 = v91
																				m.G0 = v8 + int32(16)
																				return v140
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
									}
								} else {
									v57 = F_checkExprHasSubLink(m, v40)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v57)
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v60 == int32(0) {
											v140 = v40
											m.G0 = v8 + int32(16)
											return v140
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											if v63 == int32(0) {
												v140 = v40
												m.G0 = v8 + int32(16)
												return v140
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v66
												*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v60
												v69 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v69
												v75 = F_query_or_expression_tree_walker_impl(m, v40, int32(895), v8+int32(4), v69)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
													if v77 != 0 {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v79 = F_add_nulling_relids(m, v40, v77, v78)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v140 = v79
															m.G0 = v8 + int32(16)
															return v140
														}
													} else {
														v81 = F_contain_volatile_functions(m, v40)
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
															return int32(0)
														} else {
															if v81 != 0 {
																v140 = v40
																m.G0 = v8 + int32(16)
																return v140
															} else {
																v83 = F_expression_returns_set(m, v40)
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	if v83 != 0 {
																		v140 = v40
																		m.G0 = v8 + int32(16)
																		return v140
																	} else {
																		v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+60))
																		v89 = F_get_relids_in_jointree(m, v86, int32(1), int32(0))
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return int32(0)
																		} else {
																			v91 = F_make_placeholder_expr(m, v60, v40, v89)
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int32(0)
																			} else {
																				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																				*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v93
																				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																				v96 = F_bms_copy(m, v95)
																				mBase = m.M
																				v97 = m.ExcPending
																				if v97 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v96
																					v140 = v91
																					m.G0 = v8 + int32(16)
																					return v140
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
										}
									}
								}
							}
						}
					}
				}
			}
		case 1, 2:
			v111 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				v140 = v111
				m.G0 = v8 + int32(16)
				return v140
			}
		case 3:
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v99 == v100 {
				v133 = F_copyObjectImpl(m, l0)
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
					v136 = F_flatten_group_exprs_mutator(m, v135, l1)
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v136
						v140 = v133
						m.G0 = v8 + int32(16)
						return v140
					}
				}
			} else {
				if v99 <= v100 {
					v111 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						v140 = v111
						m.G0 = v8 + int32(16)
						return v140
					}
				} else {
					v140 = l0
					m.G0 = v8 + int32(16)
					return v140
				}
			}
		case 4:
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v103 < v104 {
				v111 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					v140 = v111
					m.G0 = v8 + int32(16)
					return v140
				}
			} else {
				v140 = l0
				m.G0 = v8 + int32(16)
				return v140
			}
		default:
			if v12 == int32(67) {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v113 + int32(1)
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
				v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v118)
				v122 = F_query_tree_mutator_impl(m, l0, int32(904), l1, int32(256))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+39)))
					v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
					v126 = v124 | v125
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+39)) = uint8(v126)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v117)
					v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v129 - int32(1)
					v140 = v122
					m.G0 = v8 + int32(16)
					return v140
				}
			} else {
				v111 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					v140 = v111
					m.G0 = v8 + int32(16)
					return v140
				}
			}
		}
	}
}
func F_makeSortGroupClauseForSetOp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_palloc0(m, int32(20))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(106)
		v15 = int32(0)
		F_get_sort_group_operators(m, l0, v15, int32(1), v15, v6+int32(12), v6+int32(8), v15, v6+int32(7))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if base.B2i32(l0 != int32(2287))&base.B2i32(l0 != int32(2249)) == int32(0) {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v34)
			} else {
			}
			v36 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)) = uint16(v36)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v40
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+18)) = uint8(v44)
			m.G0 = v6 + int32(16)
			return v9
		}
	}
}
