package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistCompressValues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
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
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v5 = l4
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	if v18 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v5 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v79 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = int32(0)
	goto L5
L5:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+l3))))
	if v40 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v79 = v75
	goto L1
L7:
	;
	v75 = v27 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+10)))
	if v75 < v77 {
		v27 = v75
		goto L5
	} else {
		goto L16
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5+v27<<(uint(int32(2))%32)))) = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v49 = v27 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2+v49)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v5)
	v53 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v51
	v62 = l0 + int32(1812) + v27*int32(28)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v63 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49+(l0+int32(_a_F_gistCompressValues_0)))))
	v66 = F_FunctionCall1Coll(m, v62, v65, v15)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v69 = v51
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5+v49))) = v69
	goto L7
L14:
	;
	return
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v69 = v68
	goto L13
L16:
	;
	goto L6
L17:
	;
	m.G0 = v15 + int32(16)
	return
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= v79 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v96 = v79
	goto L20
L20:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+l3))))
	if v112 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v118 = int32(0)
	goto L24
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2+v96<<(uint(int32(2))%32))))
	v118 = v117
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5+v96<<(uint(int32(2))%32)))) = v118
	v121 = v96 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v121 < v123 {
		v96 = v121
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
}
func F_gistFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	F_gistCompressValues(m, l0, l1, l2, l3, l4, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l4 != 0 {
			v17 = int32(8)
		} else {
			v17 = int32(12)
		}
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+v17)))
		v20 = F_index_form_tuple(m, v19, v9, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(_a_F_gistFormTuple_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)) = uint16(v22)
			m.G0 = v9 + int32(128)
			return v20
		}
	}
}
func F_gistMakeUnionItVec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v6
	v32 = F_palloc(m, l2<<(uint(int32(4))%32)+int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if int32(0) < v35 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = v32 + int32(20)
	v47 = v32 + int32(4)
	v56 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v24 + int32(16)
	return
L6:
	;
	v69 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v69
	v72 = int32(1)
	if l2 <= v69 {
		v172 = v69
		v180 = v72
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v56<<(uint(int32(2))%32)))) = v172
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v56))) = uint8(v180)
	v195 = v56 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v195 < v197 {
		v56 = v195
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v79 = l0 + int32(_a_F_gistMakeUnionItVec_0) + v56<<(uint(int32(2))%32)
	v81 = v56 * int32(28)
	v82 = l0 + int32(2708) + v81
	v95 = int32(0)
	goto L10
L10:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1+v95<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_index_getattr_2(m, v108, v56+int32(1), v109, v24+int32(11))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	switch v153 {
	case 0:
		v172 = v153
		v180 = v72
		goto L8
	case 1:
		goto L22
	default:
		goto L21
	}
L12:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+11)))
	if v114 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v120 = v47 + v117<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v120)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v112
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+11)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v126 == v124 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v151 = v95 + int32(1)
	if v151 != l2 {
		v95 = v151
		goto L10
	} else {
		goto L20
	}
L16:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v144 + int32(1)
	goto L15
L17:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v130 = F_FunctionCall1Coll(m, v82, v129, v120)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v120 == v130 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v137
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+12)) = uint16(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+14)) = uint8(v141)
	goto L16
L20:
	;
	goto L11
L21:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v165 = F_FunctionCall2Coll(m, l0+int32(916)+v81, v162, v32, v24+int32(12))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(2)
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v158
	goto L21
L23:
	;
	v172 = v165
	v180 = int32(0)
	goto L8
L24:
	;
	goto L7
}
func F_gistNewBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v173
L2:
	;
	return int32(0)
L3:
	;
	if v14 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v160
	v162 = int32(8)
	v164 = int32(0)
	v167 = F_ExtendBufferedRel(m, v12+v162, v164, v164, v162)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L51
	}
L7:
	;
	v29 = F_ReadBuffer(m, l0, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v31 = F_ConditionalLockBuffer(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v29 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseBuffer(m, v29)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L48
	}
L14:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)))
	if v51 == int32(0) {
		v173 = v29
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v29^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L14
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[1]))
	v50 = v44 + v29<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	F_gistcheckpage(m, l0, v29)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)))
	if v56 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_LockBuffer(m, v29, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L47
	}
L21:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+16)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v57)+12)))
	if v59&int32(2) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[2]))
	if v76 <= int32(0) {
		v173 = v29
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v65) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	v70 = v68
	goto L27
L26:
	;
	v70 = int64(3)
	goto L27
L27:
	;
	v71 = F_GlobalVisCheckRemovableFullXid(m, int32(0), v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v71 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+118)))
	if v80 != int32(112) {
		v173 = v29
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v83) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	v88 = v86
	goto L34
L33:
	;
	v88 = int64(3)
	goto L34
L34:
	;
	v89 = m.G0
	v91 = v89 - int32(32)
	m.G0 = v91
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[2]))
	if v94 < int32(2) {
		v115 = v3
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v115)
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v120
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v26
	F_XLogBeginInsert(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L44
	}
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+118)))
	if v98 != int32(112) {
		v115 = v3
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L38
L38:
	;
	if base.Ui32(v102) < base.Ui32(int32(_a_F_gistNewBuffer_0)) {
		v115 = int32(1)
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v105 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = int32(0)
	goto L35
L41:
	;
	goto L42
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+119)))
	switch v111 - int32(109) {
	case 0, 5:
		goto L43
	default:
		v115 = int32(0)
		goto L35
	}
L43:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+104)))
	v115 = v114
	goto L35
L44:
	;
	F_XLogRegisterData(m, v91, int32(25))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v131 = F_XLogInsert(m, int32(14), int32(32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v91 + int32(32)
	v173 = v29
	goto L1
L47:
	;
	goto L13
L48:
	;
	v142 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v142 != int32(-1) {
		v26 = v142
		goto L7
	} else {
		goto L50
	}
L50:
	;
	goto L8
L51:
	;
	v173 = v167
	goto L1
}
func F_gistRedoClearFollowRight(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_XLogReadBufferForRedo(m, l0, l1, v9+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14&int32(-3) == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v20 < int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_gistRedoClearFollowRight[0]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
				v38 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_gistRedoClearFollowRight[1]))
				v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
			}
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
			v43 = base.I32_wrap_i64(int64(base.Ui64(v11) >> (uint(int64(32)) % 64)))
			*(*int32)(unsafe.Add(mBase, uint32(v39+v38))) = v43
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
			v47 = base.I32_wrap_i64(v11)
			*(*int32)(unsafe.Add(mBase, uint32(v38+v45)+4)) = v47
			v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
			v50 = v38 + v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
			v53 = v51 & int32(_a_F_gistRedoClearFollowRight_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)) = uint16(v53)
			*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v47
			*(*int32)(unsafe.Add(mBase, uint32(v38))) = v43
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			F_MarkBufferDirty(m, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if v64 != 0 {
					F_UnlockReleaseBuffer(m, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v64 != 0 {
				F_UnlockReleaseBuffer(m, v64)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_gistScanPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v332 int32
	_ = v332
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 float64
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v662 int64
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v702 int32
	_ = v702
	var v734 int32
	_ = v734
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = F_ReadBuffer(m, v34, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v36, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v36 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v34, v59, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v36^int32(-1))<<(uint(int32(6))%32))+16))
	v59 = v50
	goto L4
L6:
	;
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[1]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+v36<<(uint(int32(6))%32)+int32(-64))+16))
	v59 = v58
	goto L4
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_gistcheckpage(m, v63, v36)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v36 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v85 == int64(0) {
		v132 = v84
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[2]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+(v36^int32(-1))<<(uint(int32(2))%32))))
	v83 = v75
	goto L10
L12:
	;
	goto L13
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[3]))
	v83 = v77 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v83)+12)))
	if v136&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v88 = v84 + v83
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+12)))
	if v89&int32(8) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88)+4)))
	v95 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88))))
	if base.Ui64(v94|v95<<(uint(int64(32))%64)) <= base.Ui64(v85) {
		v132 = v84
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v100 == int32(-1) {
		v132 = v84
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v103 = int32(_a_F_gistScanPage_0)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v113 = F_palloc(m, v108<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v113)+16)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v121 = v119 << (uint(int32(4)) % 32)
	if v121 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	base.MemoryCopy(m, v113+int32(32), l2, v121)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_pairingheap_add(m, v125, v113)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v104
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v132 = v130
	goto L14
L26:
	;
	F_UnlockReleaseBuffer(m, v36)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L115
	}
L27:
	;
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))) = v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v139
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[6])))
	if v143 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_MemoryContextReset(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v146 = F_BufferGetLSNAtomic(m, v36)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v146
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+12)))
	if base.Ui32(v149) < base.Ui32(int32(25)) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v157 = int32(base.Ui32(v149+int32(_a_F_gistScanPage_1))>>(uint(int32(2))%32)) & int32(_a_F_gistScanPage_2)
	if v157 == int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v161 = v32 + int32(48)
	v179 = int32(1)
	goto L35
L35:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v196 = v179 & int32(_a_F_gistScanPage_2)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(20)+v196<<(uint(int32(2))%32))))
	v201 = int32(_a_F_gistScanPage_3)
	if base.B2i32(v192 == int32(1))&base.B2i32(v200&v201 == v201) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L26
L37:
	;
	v702 = v179 + int32(1)
	if base.Ui32(v702&int32(_a_F_gistScanPage_2)) <= base.Ui32(v157) {
		v179 = v702
		goto L35
	} else {
		goto L114
	}
L38:
	;
	v206 = int32(_a_F_gistScanPage_0)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v215 = v83 + v200&int32(_a_F_gistScanPage_4)
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+4)))
	if v216 != int32(_a_F_gistScanPage_5) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v207
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	F_MemoryContextReset(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L90
	}
L40:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v221 = int32(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v221 < v223 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v436)+12)))
	if v438&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L43:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v228 = v226
	v229 = v223
	v239 = v221
	goto L46
L44:
	;
	v332 = v221
	goto L45
L45:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v347 <= int32(0) {
		v542 = v332
		v546 = v221
		goto L39
	} else {
		goto L65
	}
L46:
	;
	v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v228)+4)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v258 = F_index_getattr_2(m, v215, v254, v255, v30+int32(31))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v332 = v312
	goto L45
L48:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v260&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v314 = int32(1)
	if v314 < v229 {
		v228 = v228 + int32(48)
		v229 = v229 - v314
		v239 = v312
		goto L46
	} else {
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v207
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	F_MemoryContextReset(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L63
	}
L51:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v302 == int32(0) {
		v312 = v239
		goto L49
	} else {
		goto L62
	}
L52:
	;
	if v260&int32(64) == int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v277 != 0 {
		goto L50
	} else {
		goto L58
	}
L55:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v267)+12)))
	if v269&int32(1) == int32(0) {
		v312 = v239
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v274&int32(1) != 0 {
		v312 = v239
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v228)+4)))
	v282 = v30 + int32(12)
	F_gistdentryinit(m, v220, v278-int32(1), v282, v258, v219, v83, v196, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v286)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v228)+44))
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v228)+6)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	v296 = F_FunctionCall5Coll(m, v228+int32(16), v290, v282, v291, v292, v293, v30+int32(30))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v296 == int32(0) {
		goto L50
	} else {
		goto L61
	}
L61:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)))
	v312 = v300 | v239
	goto L49
L62:
	;
	goto L50
L63:
	;
	goto L37
L64:
	;
	goto L47
L65:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	v353 = v350
	v354 = v347
	v357 = v351
	v368 = v221
	goto L66
L66:
	;
	v379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v353)+4)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v383 = F_index_getattr_2(m, v215, v379, v380, v30+int32(31))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v542 = v332
	v546 = v426
	goto L39
L68:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v385&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v432 = int32(1)
	if v432 < v354 {
		v353 = v353 + int32(48)
		v354 = v354 - v432
		v357 = v357 + int32(16)
		v368 = v426
		goto L66
	} else {
		goto L77
	}
L70:
	;
	v399 = int32(*(*int16)(unsafe.Add(mBase, uint32(v353)+4)))
	v403 = v30 + int32(12)
	F_gistdentryinit(m, v220, v399-int32(1), v403, v383, v219, v83, v196, int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v390&int32(1) == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v357)+8)) = uint8(v395)
	*(*int64)(unsafe.Add(mBase, uint32(v357))) = int64(0)
	v426 = v368
	goto L69
L74:
	;
	goto L73
L75:
	;
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v407)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v353)+44))
	v413 = int32(*(*int16)(unsafe.Add(mBase, uint32(v353)+6)))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	v417 = F_FunctionCall5Coll(m, v353+int32(16), v411, v403, v412, v413, v414, v30+int32(30))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v419 = *(*float64)(unsafe.Add(mBase, uint32(v417)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)))
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357)+8)) = uint8(v421)
	*(*float64)(unsafe.Add(mBase, uint32(v357))) = v419
	v426 = v420 | v368
	goto L69
L77:
	;
	goto L67
L78:
	;
	v542 = v443
	v546 = int32(0)
	goto L39
L79:
	;
	v462 = int32(0)
	goto L87
L80:
	;
	v443 = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v443 < v444 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L78
L84:
	;
	F_errmsg_internal(m, int32(_a_F_gistScanPage_6), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_gistScanPage_7), int32(152), int32(_a_F_gistScanPage_8))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v489 = v462 << (uint(int32(4)) % 32)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v489+v490))) = int64(-4503599627370496)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494+v489)+8)) = uint8(v496)
	v499 = v462 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v499 < v500 {
		v462 = v499
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L78
L89:
	;
	goto L88
L90:
	;
	if l3 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v581 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v565)+12)))
	if v567&int32(1) == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v572 = int32(1)
	F_tbm_add_tuples(m, l3, v215, v572, v542&v572)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v577 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v577 + int64(1)
	goto L37
L95:
	;
	v626 = int32(_a_F_gistScanPage_0)
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v629
	v632 = v581 << (uint(int32(4)) % 32)
	v635 = F_palloc(m, v632+int32(32))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L102
	}
L96:
	;
	v582 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v582)+12)))
	if v584&int32(1) == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	v590 = int32(4)
	v592 = v161 + v589<<(uint(v590)%32)
	v593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v592)+4)) = uint16(v593)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v595
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	v600 = v161 + v597<<(uint(v590)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v600)+6)) = uint8(v542)
	*(*uint16)(unsafe.Add(mBase, uint32(v600)+12)) = uint16(v179)
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v603 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v606 = int32(_a_F_gistScanPage_0)
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[6])))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v609
	v611 = F_gistFetchTuple(m, v33, v34, v215)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v620 = v597
	goto L100
L100:
	;
	v624 = v620 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))) = uint16(v624)
	goto L37
L101:
	;
	v613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v613<<(uint(int32(4))%32))+8)) = v611
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v607
	v620 = v613
	goto L100
L102:
	;
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v637)+12)))
	if v639&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v632 != 0 {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v635)+12)) = int32(-1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v635)+16)) = v644
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v635)+20)) = uint16(v646)
	*(*uint8)(unsafe.Add(mBase, uint32(v635)+23)) = uint8(v546)
	*(*uint8)(unsafe.Add(mBase, uint32(v635)+22)) = uint8(v542)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v650 != int32(1) {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+2)))
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215))))
	*(*int32)(unsafe.Add(mBase, uint32(v635)+12)) = v656 | v657<<(uint(int32(16))%32)
	v662 = F_BufferGetLSNAtomic(m, v36)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v653 = F_gistFetchTuple(m, v33, v34, v215)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v635)+24)) = v653
	goto L103
L109:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v635)+16)) = v662
	goto L103
L110:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	base.MemoryCopy(m, v635+int32(32), v667, v632)
	goto L112
L111:
	;
	goto L112
L112:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_pairingheap_add(m, v669, v635)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v627
	goto L37
L114:
	;
	goto L36
L115:
	;
	m.G0 = v30 + int32(32)
	return
}
func F_gistSortedBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v13 = int32(_a_F_gistSortedBuildCallback_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0])) = v17
	F_gistCompressValues(m, v16, l0, l2, l3, int32(1), v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
		F_tuplesort_putindextuplevalues(m, v22, v23, l1, v11, l3)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0])) = v14
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			F_MemoryContextReset(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
				*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v32 + int64(1)
				m.G0 = v11 + int32(128)
				return
			}
		}
	}
}
func F_gist_poly_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		v17 = int32(0)
		if base.B2i32(v16 == v17)|base.B2i32(v8 == v17) != 0 {
			v32 = int32(0)
			return v32
		} else {
			v26 = F_rtree_internal_consistent(m, v16, v8+int32(8), v12&int32(_a_F_gist_poly_consistent_0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v8 == v28 {
					v32 = v26
					return v32
				} else {
					F_pfree(m, v8)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = v26
						return v32
					}
				}
			}
		}
	}
}
func F_gist_translate_cmptype_common(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = v2 - int32(1)
	if base.Ui32(v4) <= base.Ui32(int32(7)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_c_F_gist_translate_cmptype_common[0])))
		v11 = v9
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_gist_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_gist_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
