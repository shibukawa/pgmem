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
	var v113 int32
	_ = v113
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49+(l0+int32(8084)))))
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
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+l3))))
	if v113 != 0 {
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
			v22 = int32(65535)
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
	var v55 int32
	_ = v55
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
	v55 = v6
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
	*(*int32)(unsafe.Add(mBase, uint32(l3+v55<<(uint(int32(2))%32)))) = v172
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v55))) = uint8(v180)
	v195 = v55 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v195 < v197 {
		v55 = v195
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v79 = l0 + int32(8084) + v55<<(uint(int32(2))%32)
	v81 = v55 * int32(28)
	v82 = l0 + int32(2708) + v81
	v95 = int32(0)
	goto L10
L10:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1+v95<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_index_getattr_2(m, v108, v55+int32(1), v109, v24+int32(11))
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
	var v118 int32
	_ = v118
	var v120 int64
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
	var v157 int32
	_ = v157
	var v160 int64
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
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v160
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v29^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L14
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v76 = *(*int32)(unsafe.Add(mBase, _consts[18]))
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
	v94 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v94 < int32(2) {
		v115 = v3
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v115)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v120
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
	if base.Ui32(v102) < base.Ui32(int32(12000)) {
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
				v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
				v38 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
			v53 = v51 & int32(65527)
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v328 int32
	_ = v328
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 float64
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int64
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int64
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v683 int32
	_ = v683
	var v716 int32
	_ = v716
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = F_ReadBuffer(m, v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v37, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v37 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v35, v60, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+(v37^int32(-1))<<(uint(int32(6))%32))+16))
	v60 = v51
	goto L4
L6:
	;
	goto L7
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+v37<<(uint(int32(6))%32)+int32(-64))+16))
	v60 = v59
	goto L4
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_gistcheckpage(m, v64, v37)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v37 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v86 == int64(0) {
		v133 = v85
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v37^int32(-1))<<(uint(int32(2))%32))))
	v84 = v76
	goto L10
L12:
	;
	goto L13
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v84 = v78 + v37<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v84)+12)))
	if v137&int32(2) != 0 {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v89 = v85 + v84
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v90&int32(8) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v95 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v89)+4)))
	v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v89))))
	if base.Ui64(v95|v96<<(uint(int64(32))%64)) <= base.Ui64(v86) {
		v133 = v85
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v101 == int32(-1) {
		v133 = v85
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v104 = int32(4562096)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v114 = F_palloc(m, v109<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v114)+16)) = v118
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v124 = v122 << (uint(int32(4)) % 32)
	if v124 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_pairingheap_add(m, v127, v114)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v125 = F__emscripten_memcpy_bulkmem(m, v114+int32(32), l2, v124)
	mBase = m.M
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v105
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v133 = v132
	goto L14
L27:
	;
	F_UnlockReleaseBuffer(m, v37)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L117
	}
L28:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[21]))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v140
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[22])))
	if v144 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_MemoryContextReset(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v147 = F_BufferGetLSNAtomic(m, v37)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v147
	v151 = v84 + int32(12)
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151))))
	if base.Ui32(v152) < base.Ui32(int32(25)) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v160 = int32(base.Ui32(v152+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v160 == int32(0) {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v164 = v33 + int32(48)
	v184 = int32(1)
	goto L36
L36:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v200 = v184 & int32(65535)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v200<<(uint(int32(2))%32)+(v84+int32(24))-int32(4))))
	v207 = int32(98304)
	if base.B2i32(v196 == int32(1))&base.B2i32(v206&v207 == v207) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L27
L38:
	;
	v683 = v184 + int32(1)
	if base.Ui32(v683&int32(65535)) <= base.Ui32(v160) {
		v184 = v683
		goto L36
	} else {
		goto L116
	}
L39:
	;
	v212 = int32(4562096)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v221 = v84 + v206&int32(32767)
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	if v222 != int32(65534) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v213
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	F_MemoryContextReset(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L90
	}
L41:
	;
	v508 = int32(0)
	v518 = v227
	v519 = v246
	goto L40
L42:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v227 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v227 < v229 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v435))))
	if v437&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L45:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v234 = v232
	v235 = v229
	v246 = v227
	goto L48
L46:
	;
	v328 = v227
	goto L47
L47:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v344 <= int32(0) {
		v508 = int32(1)
		v518 = v227
		v519 = v328
		goto L40
	} else {
		goto L66
	}
L48:
	;
	v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+4)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	v265 = F_index_getattr_2(m, v221, v261, v262, v31+int32(31))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v328 = v308
	goto L47
L50:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v267&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v309 = int32(1)
	if v309 < v235 {
		v234 = v234 + int32(48)
		v235 = v235 - v309
		v246 = v308
		goto L48
	} else {
		goto L65
	}
L52:
	;
	if v267&int32(64) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+31)))
	if v281 != 0 {
		goto L41
	} else {
		goto L61
	}
L55:
	;
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v272))))
	if v274&int32(1) == int32(0) {
		v308 = v246
		goto L51
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+31)))
	if v280 != 0 {
		goto L41
	} else {
		goto L60
	}
L58:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+31)))
	if v279 != 0 {
		v308 = v246
		goto L51
	} else {
		goto L59
	}
L59:
	;
	goto L41
L60:
	;
	v308 = v246
	goto L51
L61:
	;
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+4)))
	F_gistdentryinit(m, v226, v282-int32(1), v31+int32(12), v265, v225, v84, v200, int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+30)) = uint8(v290)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v234)+44))
	v298 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+6)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v302 = F_FunctionCall5Coll(m, v234+int32(16), v294, v31+int32(12), v297, v298, v299, v31+int32(30))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v302 == int32(0) {
		goto L41
	} else {
		goto L64
	}
L64:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+30)))
	v308 = v306 | v246
	goto L51
L65:
	;
	goto L49
L66:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	v350 = v347
	v354 = v348
	v356 = v344
	v361 = v227
	goto L67
L67:
	;
	v377 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350)+4)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	v381 = F_index_getattr_2(m, v221, v377, v378, v31+int32(31))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v508 = v430
	v518 = v424
	v519 = v328
	goto L40
L69:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v383&int32(1) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v430 = int32(1)
	if v430 < v356 {
		v350 = v350 + int32(48)
		v354 = v354 + int32(16)
		v356 = v356 - v430
		v361 = v424
		goto L67
	} else {
		goto L78
	}
L71:
	;
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350)+4)))
	F_gistdentryinit(m, v226, v395-int32(1), v31+int32(12), v381, v225, v84, v200, int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L76
	}
L72:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+31)))
	if v388 != int32(1) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v391 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v354)+8)) = uint8(v391)
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = int64(0)
	v424 = v361
	goto L70
L75:
	;
	goto L74
L76:
	;
	v403 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+30)) = uint8(v403)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v350)+44))
	v411 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350)+6)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	v415 = F_FunctionCall5Coll(m, v350+int32(16), v407, v31+int32(12), v410, v411, v412, v31+int32(30))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v417 = *(*float64)(unsafe.Add(mBase, uint32(v415)))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+30)))
	v419 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v354)+8)) = uint8(v419)
	*(*float64)(unsafe.Add(mBase, uint32(v354))) = v417
	v424 = v418 | v361
	goto L70
L78:
	;
	goto L68
L79:
	;
	v463 = int32(0)
	goto L87
L80:
	;
	v442 = int32(1)
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
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v508 = v442
	v518 = int32(0)
	v519 = v443
	goto L40
L84:
	;
	F_errmsg_internal(m, int32(428891), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(518454), int32(152), int32(83414))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
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
	v491 = v463 << (uint(int32(4)) % 32)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v491+v492))) = int64(-4503599627370496)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v496+v491)+8)) = uint8(v498)
	v501 = v463 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v501 < v502 {
		v463 = v501
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v508 = v442
	v518 = int32(0)
	v519 = v443
	goto L40
L89:
	;
	goto L88
L90:
	;
	if v508 == int32(0) {
		goto L38
	} else {
		goto L91
	}
L91:
	;
	if l3 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v560 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v544))))
	if v546&int32(1) == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v551 = int32(1)
	F_tbm_add_tuples(m, l3, v221, v551, v519&v551)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v556 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v556 + int64(1)
	goto L38
L96:
	;
	v605 = int32(4562096)
	v606 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v608
	v611 = v560 << (uint(int32(4)) % 32)
	v614 = F_palloc(m, v611+int32(32))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L103
	}
L97:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v561))))
	if v563&int32(1) == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[21]))))
	v569 = int32(4)
	v571 = v164 + v568<<(uint(v569)%32)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v572
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v571)+4)) = uint16(v574)
	v576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[21]))))
	v579 = v164 + v576<<(uint(v569)%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v579)+12)) = uint16(v184)
	*(*uint8)(unsafe.Add(mBase, uint32(v579)+6)) = uint8(v519)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v582 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v585 = int32(4562096)
	v586 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[22])))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v588
	v590 = F_gistFetchTuple(m, v34, v35, v221)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	v599 = v576
	goto L101
L101:
	;
	v603 = v599 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[21]))) = uint16(v603)
	goto L38
L102:
	;
	v592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[21]))))
	*(*int32)(unsafe.Add(mBase, uint32(v164+v592<<(uint(int32(4))%32))+8)) = v590
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v586
	v599 = v592
	goto L101
L103:
	;
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v616))))
	if v618&int32(1) != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	if v611 != 0 {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+12)) = int32(-1)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	*(*int32)(unsafe.Add(mBase, uint32(v614)+16)) = v623
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v614)+20)) = uint16(v625)
	*(*uint8)(unsafe.Add(mBase, uint32(v614)+23)) = uint8(v518)
	*(*uint8)(unsafe.Add(mBase, uint32(v614)+22)) = uint8(v519)
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v629 != int32(1) {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+2)))
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221))))
	*(*int32)(unsafe.Add(mBase, uint32(v614)+12)) = v635 | v636<<(uint(int32(16))%32)
	v641 = F_BufferGetLSNAtomic(m, v37)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	v632 = F_gistFetchTuple(m, v34, v35, v221)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+24)) = v632
	goto L104
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v614)+16)) = v641
	goto L104
L111:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_pairingheap_add(m, v649, v614)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L115
	}
L112:
	;
	v647 = F__emscripten_memcpy_bulkmem(m, v614+int32(32), v646, v611)
	mBase = m.M
	goto L114
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v606
	goto L38
L116:
	;
	goto L37
L117:
	;
	m.G0 = v31 + int32(32)
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
	v13 = int32(4562096)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
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
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
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
	var v2 int32
	_ = v2
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v2 = int32(0)
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
		if v16 == int32(0) {
			v31 = v2
			return v31
		} else {
			if v8 == int32(0) {
				v31 = v2
				return v31
			} else {
				v25 = F_rtree_internal_consistent(m, v16, v8+int32(8), v12&int32(65535))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v8 == v27 {
						v31 = v25
						return v31
					} else {
						F_pfree(m, v8)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = v25
							return v31
						}
					}
				}
			}
		}
	}
}
func F_gist_translate_cmptype_common(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = v3 - int32(1)
	if base.Ui32(v5) <= base.Ui32(int32(7)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[28])))
		v13 = v12
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_gist_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
