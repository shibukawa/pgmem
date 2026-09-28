package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MarkLockClear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)) = uint8(v2)
	return
}
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v6 int32
	_ = v6
	if l0 != 0 {
		v4 = base.F64_neg(l1)
	} else {
		v4 = l1
	}
	v6 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v6-int32(16))+8)) = v4
	return base.F64_mul(l1, v4)
}
func F_makeMdArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v8 = int32(_a_F_makeMdArrayResult_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_makeMdArrayResult[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_makeMdArrayResult[0])) = l4
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v17 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+27)))
	v18 = F_construct_md_array(m, v12, v13, l1, l2, l3, v14, v15, v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_makeMdArrayResult[0])) = v9
		if l5 != 0 {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_MemoryContextDelete(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v18)
			}
		} else {
			return base.I64_extend_i32_u(v18)
		}
	}
}
func F_makeRecursiveViewSelect(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(84))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(141)
	v23 = F_palloc0(m, int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(110)
	v28 = F_palloc0(m, int32(56))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(115)
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v28
	v46 = F_list_make1_impl(m, v39, v13+int32(4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v46
	if l1 == int32(0) {
		v109 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v23
	v115 = F_makeRangeVar(m, int32(0), l0, int32(-1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v53 <= int32(0) {
		v109 = v4
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v63 = v4
	v64 = v4
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = F_palloc0(m, int32(20))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v109 = v95
	goto L6
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = int64(81)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66+v63<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v80 = F_palloc0(m, int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(69)
	v86 = F_makeString(m, v78)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v89 = F_lcons(m, v86, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v80
	v95 = F_lappend(m, v64, v68)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v98 = v63 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v98 < v99 {
		v63 = v98
		v64 = v95
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v115
	v120 = F_list_make1_impl(m, int32(1), v13)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v120
	m.G0 = v13 + int32(16)
	return v16
}
func F_makeRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_palloc0(m, int32(28))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(27)
		return v8
	}
}
func F_make_andclause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(21)
		return v4
	}
}
func F_make_new_segment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = v9 + int32(32)
	v16 = int32(1)
	goto L1
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11+v16<<(uint(int32(2))%32))))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1452))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1448))
	if base.Ui32(v59) <= base.Ui32(v60) {
		v246 = v58
		goto L13
	} else {
		goto L14
	}
L3:
	;
	goto L2
L4:
	;
	v57 = v16
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v16 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11+v28<<(uint(int32(2))%32))))
	if v32 == int32(0) {
		v57 = v28
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v35 = int32(2)
	v36 = v16 + v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11+v36<<(uint(v35)%32))))
	if v40 == int32(0) {
		v57 = v36
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v44 = v16 + int32(3)
	if v44 == int32(32) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11+v44<<(uint(int32(2))%32))))
	if v52 == int32(0) {
		v57 = v44
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v16 = v16 + int32(4)
	goto L1
L13:
	;
	return v246
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1440))
	v65 = v62 << (uint(int32(base.Ui32(v57)>>(uint(int32(1))%32))) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1444))
	if base.Ui32(v65) < base.Ui32(v66) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = v65
	goto L17
L16:
	;
	v68 = v66
	goto L17
L17:
	;
	v69 = v59 - v60
	if base.Ui32(v68) < base.Ui32(v69) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = v68
	goto L20
L19:
	;
	v71 = v69
	goto L20
L20:
	;
	v75 = int32(base.Ui32(v71)>>(uint(int32(10))%32)) & int32(_a_F_make_new_segment_0)
	v77 = v75 + int32(584)
	v79 = v77 & int32(4092)
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v83 = v75 - v79 + int32(_a_F_make_new_segment_1)
	goto L23
L22:
	;
	v83 = v77
	goto L23
L23:
	;
	if base.Ui32(v71) <= base.Ui32(v83) {
		v246 = v58
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v87 = int32(base.Ui32(v71-v83) >> (uint(int32(12)) % 32))
	if base.Ui32(v87) < base.Ui32(l1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = int32(2)
	v90 = l1 << (uint(v89) % 32)
	v93 = int32(4092)
	v94 = base.I32_div_u_s(v90+int32(_a_F_make_new_segment_2), v93)
	v99 = v90 + v94<<(uint(v89)%32) + int32(584)
	if v99&v93 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v113 = v87
	v114 = v71
	v115 = v83
	goto L27
L27:
	;
	v116 = int32(_a_F_make_new_segment_3)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0])) = v119
	v122 = F_dsm_create(m, v114, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v106 = v99&int32(16773120) + int32(_a_F_make_new_segment_4)
	goto L30
L29:
	;
	v106 = v99
	goto L30
L30:
	;
	v109 = v106 + l1<<(uint(int32(12))%32)
	if base.Ui32(int32(134217728)) < base.Ui32(v109) {
		v246 = v58
		goto L13
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v69) < base.Ui32(v109) {
		v246 = v58
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v113 = l1
	v114 = v109
	v115 = v106
	goto L27
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0])) = v117
	if v122 == int32(0) {
		v246 = v58
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_dsm_pin_segment(m, v122)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v132+v57<<(uint(int32(2))%32))+32)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+1456))
	if base.Ui32(v139) < base.Ui32(v57) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+1456)) = v57
	goto L39
L38:
	;
	goto L39
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v142) < base.Ui32(v57) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+648)) = v57
	goto L42
L41:
	;
	goto L42
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+1448))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+1448)) = v146 + v114
	v151 = l0 + v57*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v122
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+16)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v151)+12)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v151)+24)) = v153 + int32(584)
	v160 = v153 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v160
	v162 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v160)+4)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v160)+12)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v160)+20)) = v162
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+28)) = v168
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+32)) = uint8(v170)
	if v160 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	F_FreePageManagerPut(m, v183, int32(base.Ui32(v115)>>(uint(int32(12))%32)), v113)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L33
	} else {
		goto L47
	}
L44:
	;
	v176 = v160 - v153 + v170
	goto L46
L45:
	;
	v176 = v168
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v176
	base.MemoryFill(m, v153+int32(68), int32(0), int32(516))
	goto L43
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v57 ^ v190 ^ int32(216163848)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v113
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v114
	v200 = v151 + int32(8)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if v113 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v202 = int32(14)
	v205 = base.I32_clz(v113) ^ int32(31)
	if base.Ui32(v202) <= base.Ui32(v205) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v213 = int32(0)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+20)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v216 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	v221 = int32(2)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v219+v220<<(uint(v221)%32))+160))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+16)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v226)+24)) = uint8(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v231<<(uint(v221)%32))+160)) = v57
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
	if v237 == v216 {
		v246 = v200
		goto L13
	} else {
		goto L54
	}
L51:
	;
	v208 = v202
	goto L53
L52:
	;
	v208 = v205
	goto L53
L53:
	;
	v213 = v208 + int32(1)
	goto L50
L54:
	;
	v240 = F_get_segment_by_index(m, l0, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L33
	} else {
		goto L55
	}
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+12)) = v57
	v246 = v200
	goto L13
}
func F_make_pathkey_from_sortinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = F_get_opfamily_member_for_cmptype(m, l2, l3, l3, int32(3))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v23 = F_get_mergejoin_opfamilies(m, v19)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
						F_errmsg_internal(m, int32(_a_F_make_pathkey_from_sortinfo_0), v16+int32(16))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_make_pathkey_from_sortinfo_1), int32(232), int32(_a_F_make_pathkey_from_sortinfo_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = F_get_eclass_for_sort_expr(m, l0, l1, v23, l3, l4, l7, l8, l9)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != 0 {
							if l5 != 0 {
								v31 = int32(5)
							} else {
								v31 = int32(1)
							}
							v32 = F_make_canonical_pathkey(m, l0, v27, l2, v31, l6)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v35 = v32
								m.G0 = v16 + int32(32)
								return v35
							}
						} else {
							v35 = int32(0)
							m.G0 = v16 + int32(32)
							return v35
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(3)
				F_errmsg_internal(m, int32(_a_F_make_pathkey_from_sortinfo_3), v16)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_make_pathkey_from_sortinfo_1), int32(228), int32(_a_F_make_pathkey_from_sortinfo_2))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_make_pathtarget_from_tlist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v7 = F_palloc0(m, int32(40))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(280)
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(0)
	return v7
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v46 = F_palloc(m, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(0)
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v29 = v24 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = F_lappend(m, v27, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37+v29))) = v39
	v42 = v24 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 < v43 {
		v24 = v42
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v46
	goto L3
}
func F_make_relative_path(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v15 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_canonicalize_path_enc(m, l0)
	mBase = m.M
	m.G0 = v13 + int32(16)
	return
L2:
	;
	goto L93
L3:
	;
	v21 = v15
	v24 = v4
	v27 = int64(0)
	goto L4
L4:
	;
	if v27 == int64(33) {
		v60 = v24
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v60 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v37 = v21 & int32(255)
	if base.B2i32(int64(1)<<(uint(v27)%64)&int64(539103233) == int64(0))|base.B2i32(v37 != int32(47)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v54))))
	if v58 != 0 {
		v21 = v58
		v24 = v55
		v27 = v56
		goto L4
	} else {
		goto L13
	}
L9:
	;
	v44 = v27 + int64(1)
	v45 = base.I32_wrap_i64(v44)
	v54 = v45
	v55 = v45
	v56 = v44
	goto L8
L10:
	;
	goto L11
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v27))+uint32(_c_F_make_relative_path[0]))))
	if v37 != v49 {
		v60 = v24
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v52 = v27 + int64(1)
	v54 = base.I32_wrap_i64(v52)
	v55 = v24
	v56 = v52
	goto L8
L13:
	;
	v60 = v55
	goto L6
L14:
	;
	goto L18
L15:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v183 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	v180 = F_strlen(m, v169)
	mBase = m.M
	goto L15
L18:
	;
	goto L19
L19:
	;
	v70 = int32(1023)
	if (l0^l2)&int32(3) != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v173)
	goto L16
L21:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L42
L22:
	;
	if v144 == int32(0) {
		v169 = v142
		v170 = v143
		goto L20
	} else {
		goto L41
	}
L23:
	;
	v142 = l2
	v143 = l0
	v144 = v70
	goto L22
L24:
	;
	goto L25
L25:
	;
	v74 = int32(0)
	if base.B2i32(l2&int32(3) == v74)|int32(0) == v74 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v110 == int32(0) {
		v169 = v107
		v170 = v108
		goto L20
	} else {
		goto L35
	}
L27:
	;
	v86 = l2
	v87 = l0
	v88 = v70
	goto L30
L28:
	;
	goto L29
L29:
	;
	v107 = l2
	v108 = l0
	v109 = v70
	v110 = int32(1)
	goto L26
L30:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v90)
	if v90 == int32(0) {
		v149 = v86
		v150 = v87
		v151 = v88
		goto L21
	} else {
		goto L32
	}
L31:
	;
	v107 = v101
	v108 = v95
	v109 = v97
	v110 = v99
	goto L26
L32:
	;
	v94 = int32(1)
	v95 = v87 + v94
	v97 = v88 - v94
	v98 = int32(0)
	v99 = base.B2i32(v97 != v98)
	v101 = v86 + v94
	if v101&int32(3) == v98 {
		v107 = v101
		v108 = v95
		v109 = v97
		v110 = v99
		goto L26
	} else {
		goto L33
	}
L33:
	;
	if v97 != 0 {
		v86 = v101
		v87 = v95
		v88 = v97
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if base.B2i32(v113 == int32(0))|base.B2i32(base.Ui32(v109) < base.Ui32(int32(4))) != 0 {
		v142 = v107
		v143 = v108
		v144 = v109
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v120 = v107
	v121 = v108
	v122 = v109
	goto L37
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v128 = int32(-2139062144)
	if (int32(16843008)-v125|v125)&v128 != v128 {
		v149 = v120
		v150 = v121
		v151 = v122
		goto L21
	} else {
		goto L39
	}
L38:
	;
	v142 = v136
	v143 = v134
	v144 = v138
	goto L22
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v125
	v133 = int32(4)
	v134 = v121 + v133
	v136 = v120 + v133
	v138 = v122 - v133
	if base.Ui32(int32(3)) < base.Ui32(v138) {
		v120 = v136
		v121 = v134
		v122 = v138
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v149 = v142
	v150 = v143
	v151 = v144
	goto L21
L42:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v158)
	if v158 == int32(0) {
		v169 = v154
		v170 = v155
		goto L20
	} else {
		goto L44
	}
L43:
	;
	v169 = v165
	v170 = v163
	goto L20
L44:
	;
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v154 + v162
	v167 = v156 - v162
	if v167 != 0 {
		v154 = v165
		v155 = v163
		v156 = v167
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v184 = F_strlen(m, l0)
	mBase = m.M
	v189 = v184 + l0
	goto L49
L47:
	;
	goto L48
L48:
	;
	F_canonicalize_path_enc(m, l0)
	mBase = m.M
	v257 = F_strlen(m, l0)
	mBase = m.M
	v258 = v257 + (v60 - int32(33))
	if v258 <= int32(0) {
		goto L2
	} else {
		goto L67
	}
L49:
	;
	v197 = v189 - int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if base.B2i32(v198 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v197)) != 0 {
		v189 = v197
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v206 = v197
	goto L52
L51:
	;
	goto L50
L52:
	;
	if base.Ui32(l0) < base.Ui32(v206) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v222 = v206
	goto L58
L54:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v216 != int32(47) {
		v206 = v206 - int32(1)
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L56
L58:
	;
	if base.Ui32(l0) < base.Ui32(v222) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if l0 == v222 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v232 = v222 - int32(1)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v233 == int32(47) {
		v222 = v232
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L62
L64:
	;
	v241 = l0 + base.B2i32(v183 == int32(47))
	goto L66
L65:
	;
	v241 = v222
	goto L66
L66:
	;
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v242)
	goto L48
L67:
	;
	v261 = l0 + v258
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261-int32(1)))))
	if v264 != int32(47) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v268 = v60 + int32(_a_F_make_relative_path_0)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v269 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v272 = v269
	v273 = v268
	v277 = v261
	goto L72
L70:
	;
	v300 = v268
	goto L71
L71:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v307 != 0 {
		goto L2
	} else {
		goto L77
	}
L72:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v280 == int32(0) {
		goto L2
	} else {
		goto L74
	}
L73:
	;
	v300 = v293
	goto L71
L74:
	;
	v284 = v272 & int32(255)
	v286 = int32(47)
	if base.B2i32(v280 != v284)&(base.B2i32(v284 != v286)|base.B2i32(v280 != v286)) != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v292 = int32(1)
	v293 = v273 + v292
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v294 != 0 {
		v272 = v294
		v273 = v293
		v277 = v277 + v292
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v308)
	v310 = F_strlen(m, l0)
	mBase = m.M
	if v310 < int32(2) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v344 = l1 + v60
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v345 == int32(0) {
		goto L1
	} else {
		goto L84
	}
L79:
	;
	v319 = l0 + v310 - int32(1)
	goto L80
L80:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v326 != int32(47) {
		goto L78
	} else {
		goto L82
	}
L81:
	;
	goto L78
L82:
	;
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v319))) = uint8(v329)
	v332 = v319 - int32(1)
	if base.Ui32(l0) < base.Ui32(v332) {
		v319 = v332
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v348 = F_strlen(m, l0)
	mBase = m.M
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v344
	if v349 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v353 = int32(_a_F_make_relative_path_1)
	goto L87
L86:
	;
	v353 = int32(_a_F_make_relative_path_2)
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v353
	v359 = F_pg_snprintf(m, l0+v348, int32(1024)-v348, int32(_a_F_make_relative_path_3), v13)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	return
L89:
	;
	goto L1
L90:
	;
	goto L1
L91:
	;
	v487 = F_strlen(m, v476)
	mBase = m.M
	goto L90
L93:
	;
	goto L94
L94:
	;
	v377 = int32(1023)
	if (l0^l1)&int32(3) != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v477))) = uint8(v480)
	goto L91
L96:
	;
	v461 = v456
	v462 = v457
	v463 = v458
	goto L117
L97:
	;
	if v451 == int32(0) {
		v476 = v449
		v477 = v450
		goto L95
	} else {
		goto L116
	}
L98:
	;
	v449 = l1
	v450 = l0
	v451 = v377
	goto L97
L99:
	;
	goto L100
L100:
	;
	v381 = int32(0)
	if base.B2i32(l1&int32(3) == v381)|int32(0) == v381 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v417 == int32(0) {
		v476 = v414
		v477 = v415
		goto L95
	} else {
		goto L110
	}
L102:
	;
	v393 = l1
	v394 = l0
	v395 = v377
	goto L105
L103:
	;
	goto L104
L104:
	;
	v414 = l1
	v415 = l0
	v416 = v377
	v417 = int32(1)
	goto L101
L105:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	*(*uint8)(unsafe.Add(mBase, uint32(v394))) = uint8(v397)
	if v397 == int32(0) {
		v456 = v393
		v457 = v394
		v458 = v395
		goto L96
	} else {
		goto L107
	}
L106:
	;
	v414 = v408
	v415 = v402
	v416 = v404
	v417 = v406
	goto L101
L107:
	;
	v401 = int32(1)
	v402 = v394 + v401
	v404 = v395 - v401
	v405 = int32(0)
	v406 = base.B2i32(v404 != v405)
	v408 = v393 + v401
	if v408&int32(3) == v405 {
		v414 = v408
		v415 = v402
		v416 = v404
		v417 = v406
		goto L101
	} else {
		goto L108
	}
L108:
	;
	if v404 != 0 {
		v393 = v408
		v394 = v402
		v395 = v404
		goto L105
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if base.B2i32(v420 == int32(0))|base.B2i32(base.Ui32(v416) < base.Ui32(int32(4))) != 0 {
		v449 = v414
		v450 = v415
		v451 = v416
		goto L97
	} else {
		goto L111
	}
L111:
	;
	v427 = v414
	v428 = v415
	v429 = v416
	goto L112
L112:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v435 = int32(-2139062144)
	if (int32(16843008)-v432|v432)&v435 != v435 {
		v456 = v427
		v457 = v428
		v458 = v429
		goto L96
	} else {
		goto L114
	}
L113:
	;
	v449 = v443
	v450 = v441
	v451 = v445
	goto L97
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v432
	v440 = int32(4)
	v441 = v428 + v440
	v443 = v427 + v440
	v445 = v429 - v440
	if base.Ui32(int32(3)) < base.Ui32(v445) {
		v427 = v443
		v428 = v441
		v429 = v445
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v456 = v449
	v457 = v450
	v458 = v451
	goto L96
L117:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v465)
	if v465 == int32(0) {
		v476 = v461
		v477 = v462
		goto L95
	} else {
		goto L119
	}
L118:
	;
	v476 = v472
	v477 = v470
	goto L95
L119:
	;
	v469 = int32(1)
	v470 = v462 + v469
	v472 = v461 + v469
	v474 = v463 - v469
	if v474 != 0 {
		v461 = v472
		v462 = v470
		v463 = v474
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
}
func F_make_scalar_key(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
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
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
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
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 {
	case 0:
		v754 = F_palloc(m, int32(5))
		mBase = m.M
		v755 = m.ExcPending
		if v755 != 0 {
			return int64(0)
		} else {
			v756 = int32(2)
			*(*uint8)(unsafe.Add(mBase, uint32(v754)+4)) = uint8(v756)
			*(*int32)(unsafe.Add(mBase, uint32(v754))) = int32(20)
			v760 = v754
			m.G0 = v10 + int32(48)
			return base.I64_extend_i32_u(v760)
		}
	case 1:
		if l1 != 0 {
			v442 = int32(1)
		} else {
			v442 = int32(5)
		}
		v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if int32(126) <= v444 {
			v452 = v444 - int32(1636608432)
			if v443&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v444) {
					v561 = v443
					v562 = v444
					v563 = v452
					v564 = v452
					v565 = v452
					for {
						v567 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
						v568 = v567 + v564
						v569 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
						v571 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
						v572 = v571 + v565
						v574 = int32(4)
						v576 = v569 + v563 - v572 ^ base.I32_rotl(v572, v574)
						v580 = v568 - v576 ^ base.I32_rotl(v576, int32(6))
						v581 = v572 + v568
						v582 = v576 + v581
						v583 = v580 + v582
						v587 = v581 - v580 ^ base.I32_rotl(v580, int32(8))
						v591 = v582 - v587 ^ base.I32_rotl(v587, int32(16))
						v595 = v583 - v591 ^ base.I32_rotl(v591, int32(19))
						v596 = v587 + v583
						v597 = v591 + v596
						v598 = v595 + v597
						v602 = v596 - v595 ^ base.I32_rotl(v595, v574)
						v603 = int32(12)
						v604 = v561 + v603
						v606 = v562 - v603
						if base.Ui32(int32(11)) < base.Ui32(v606) {
							v561 = v604
							v562 = v606
							v563 = v597
							v564 = v598
							v565 = v602
							continue
						} else {
							break
						}
						break
					}
					v609 = v604
					v610 = v606
					v611 = v597
					v612 = v598
					v613 = v602
				} else {
					v609 = v443
					v610 = v444
					v611 = v452
					v612 = v452
					v613 = v452
				}
				switch v610 - int32(1) {
				case 0:
					v672 = v611
					v673 = v612
					v674 = v613
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 1:
					v665 = v611
					v666 = v612
					v667 = v613
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 2:
					v658 = v611
					v659 = v612
					v660 = v613
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 3:
					v652 = v612
					v653 = v613
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 4:
					v648 = v612
					v649 = v613
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 5:
					v642 = v612
					v643 = v613
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 6:
					v636 = v612
					v637 = v613
					v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+6)))
					v642 = v638<<(uint(int32(16))%32) + v636
					v643 = v637
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 7:
					v631 = v613
					v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+7)))
					v636 = v632<<(uint(int32(24))%32) + v612
					v637 = v631
					v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+6)))
					v642 = v638<<(uint(int32(16))%32) + v636
					v643 = v637
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 8:
					v626 = v613
					v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+8)))
					v631 = v627<<(uint(int32(8))%32) + v626
					v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+7)))
					v636 = v632<<(uint(int32(24))%32) + v612
					v637 = v631
					v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+6)))
					v642 = v638<<(uint(int32(16))%32) + v636
					v643 = v637
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 9:
					v621 = v613
					v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+9)))
					v626 = v622<<(uint(int32(16))%32) + v621
					v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+8)))
					v631 = v627<<(uint(int32(8))%32) + v626
					v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+7)))
					v636 = v632<<(uint(int32(24))%32) + v612
					v637 = v631
					v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+6)))
					v642 = v638<<(uint(int32(16))%32) + v636
					v643 = v637
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				case 10:
					v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+10)))
					v621 = v617<<(uint(int32(24))%32) + v613
					v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+9)))
					v626 = v622<<(uint(int32(16))%32) + v621
					v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+8)))
					v631 = v627<<(uint(int32(8))%32) + v626
					v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+7)))
					v636 = v632<<(uint(int32(24))%32) + v612
					v637 = v631
					v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+6)))
					v642 = v638<<(uint(int32(16))%32) + v636
					v643 = v637
					v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+5)))
					v648 = v644<<(uint(int32(8))%32) + v642
					v649 = v643
					v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+4)))
					v652 = v648 + v650
					v653 = v649
					v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+3)))
					v658 = v654<<(uint(int32(24))%32) + v611
					v659 = v652
					v660 = v653
					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+2)))
					v665 = v661<<(uint(int32(16))%32) + v658
					v666 = v659
					v667 = v660
					v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
					v672 = v668<<(uint(int32(8))%32) + v665
					v673 = v666
					v674 = v667
					v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
					v679 = v672 + v675
					v680 = v673
					v681 = v674
				default:
					v679 = v611
					v680 = v612
					v681 = v613
				}
			} else {
				if base.Ui32(v444) < base.Ui32(int32(12)) {
					v507 = v443
					v508 = v444
					v509 = v452
					v510 = v452
					v511 = v452
				} else {
					v459 = v443
					v460 = v444
					v461 = v452
					v462 = v452
					v463 = v452
					for {
						v465 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
						v466 = v465 + v462
						v467 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
						v469 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
						v470 = v469 + v463
						v472 = int32(4)
						v474 = v467 + v461 - v470 ^ base.I32_rotl(v470, v472)
						v478 = v466 - v474 ^ base.I32_rotl(v474, int32(6))
						v479 = v470 + v466
						v480 = v474 + v479
						v481 = v478 + v480
						v485 = v479 - v478 ^ base.I32_rotl(v478, int32(8))
						v489 = v480 - v485 ^ base.I32_rotl(v485, int32(16))
						v493 = v481 - v489 ^ base.I32_rotl(v489, int32(19))
						v494 = v485 + v481
						v495 = v489 + v494
						v496 = v493 + v495
						v500 = v494 - v493 ^ base.I32_rotl(v493, v472)
						v501 = int32(12)
						v502 = v459 + v501
						v504 = v460 - v501
						if base.Ui32(int32(11)) < base.Ui32(v504) {
							v459 = v502
							v460 = v504
							v461 = v495
							v462 = v496
							v463 = v500
							continue
						} else {
							break
						}
						break
					}
					v507 = v502
					v508 = v504
					v509 = v495
					v510 = v496
					v511 = v500
				}
				switch v508 - int32(1) {
				case 0:
					v558 = v509
					v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
					v679 = v558 + v559
					v680 = v510
					v681 = v511
				case 1:
					v553 = v509
					v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
					v558 = v554<<(uint(int32(8))%32) + v553
					v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
					v679 = v558 + v559
					v680 = v510
					v681 = v511
				case 2:
					v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+2)))
					v553 = v549<<(uint(int32(16))%32) + v509
					v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
					v558 = v554<<(uint(int32(8))%32) + v553
					v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
					v679 = v558 + v559
					v680 = v510
					v681 = v511
				case 3:
					v546 = v510
					v547 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v679 = v547 + v509
					v680 = v546
					v681 = v511
				case 4:
					v543 = v510
					v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+4)))
					v546 = v543 + v544
					v547 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v679 = v547 + v509
					v680 = v546
					v681 = v511
				case 5:
					v538 = v510
					v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+5)))
					v543 = v539<<(uint(int32(8))%32) + v538
					v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+4)))
					v546 = v543 + v544
					v547 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v679 = v547 + v509
					v680 = v546
					v681 = v511
				case 6:
					v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+6)))
					v538 = v534<<(uint(int32(16))%32) + v510
					v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+5)))
					v543 = v539<<(uint(int32(8))%32) + v538
					v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+4)))
					v546 = v543 + v544
					v547 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v679 = v547 + v509
					v680 = v546
					v681 = v511
				case 7:
					v529 = v511
					v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
					v679 = v530 + v509
					v680 = v532 + v510
					v681 = v529
				case 8:
					v524 = v511
					v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+8)))
					v529 = v525<<(uint(int32(8))%32) + v524
					v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
					v679 = v530 + v509
					v680 = v532 + v510
					v681 = v529
				case 9:
					v519 = v511
					v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+9)))
					v524 = v520<<(uint(int32(16))%32) + v519
					v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+8)))
					v529 = v525<<(uint(int32(8))%32) + v524
					v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
					v679 = v530 + v509
					v680 = v532 + v510
					v681 = v529
				case 10:
					v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+10)))
					v519 = v515<<(uint(int32(24))%32) + v511
					v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+9)))
					v524 = v520<<(uint(int32(16))%32) + v519
					v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+8)))
					v529 = v525<<(uint(int32(8))%32) + v524
					v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
					v679 = v530 + v509
					v680 = v532 + v510
					v681 = v529
				default:
					v679 = v509
					v680 = v510
					v681 = v511
				}
			}
			v684 = int32(14)
			v686 = v680 ^ v681 - base.I32_rotl(v680, v684)
			v690 = v686 ^ v679 - base.I32_rotl(v686, int32(11))
			v694 = v690 ^ v680 - base.I32_rotl(v690, int32(25))
			v698 = v694 ^ v686 - base.I32_rotl(v694, int32(16))
			v702 = v698 ^ v690 - base.I32_rotl(v698, int32(4))
			v706 = v702 ^ v694 - base.I32_rotl(v702, v684)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v706 ^ v698 - base.I32_rotl(v706, int32(24))
			v713 = v10 + int32(38)
			v718 = F_pg_snprintf(m, v713, int32(10), int32(_a_F_make_scalar_key_0), v10+int32(32))
			mBase = m.M
			v719 = m.ExcPending
			if v719 != 0 {
				return int64(0)
			} else {
				v723 = int32(8)
				v724 = v442 | int32(16)
				v725 = v713
				v727 = v723 + int32(5)
				v728 = F_palloc(m, v727)
				mBase = m.M
				v729 = m.ExcPending
				if v729 != 0 {
					return int64(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v728)+4)) = uint8(v724)
					*(*int32)(unsafe.Add(mBase, uint32(v728))) = v727 << (uint(int32(2)) % 32)
					if v723 == int32(0) {
						v760 = v728
					} else {
						base.MemoryCopy(m, v728+int32(5), v725, v723)
						v760 = v728
					}
					m.G0 = v10 + int32(48)
					return base.I64_extend_i32_u(v760)
				}
			}
		} else {
			v723 = v444
			v724 = v442
			v725 = v443
			v727 = v723 + int32(5)
			v728 = F_palloc(m, v727)
			mBase = m.M
			v729 = m.ExcPending
			if v729 != 0 {
				return int64(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v728)+4)) = uint8(v724)
				*(*int32)(unsafe.Add(mBase, uint32(v728))) = v727 << (uint(int32(2)) % 32)
				if v723 == int32(0) {
					v760 = v728
				} else {
					base.MemoryCopy(m, v728+int32(5), v725, v723)
					v760 = v728
				}
				m.G0 = v10 + int32(48)
				return base.I64_extend_i32_u(v760)
			}
		}
	case 2:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = m.G0
		v31 = v29 - int32(32)
		m.G0 = v31
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
		if base.Ui32(int32(_a_F_make_scalar_key_1)) <= base.Ui32(v33) {
			if v33 != int32(_a_F_make_scalar_key_2) {
				if v33 != int32(_a_F_make_scalar_key_3) {
					v47 = F_pstrdup(m, int32(_a_F_make_scalar_key_4))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v137 = v47
						m.G0 = v31 + int32(32)
						v145 = F_strlen(m, v137)
						mBase = m.M
						if v145 < int32(126) {
							v424 = v145
							v425 = int32(4)
							v426 = v137
							v428 = v424 + int32(5)
							v429 = F_palloc(m, v428)
							mBase = m.M
							v430 = m.ExcPending
							if v430 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
								*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
								if v424 != 0 {
									base.MemoryCopy(m, v429+int32(5), v426, v424)
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return int64(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return base.I64_extend_i32_u(v760)
								}
							}
						} else {
							v153 = v145 - int32(1636608432)
							if v137&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v145) {
									v262 = v137
									v263 = v145
									v264 = v153
									v265 = v153
									v266 = v153
									for {
										v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
										v269 = v268 + v265
										v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
										v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
										v273 = v272 + v266
										v275 = int32(4)
										v277 = v270 + v264 - v273 ^ base.I32_rotl(v273, v275)
										v281 = v269 - v277 ^ base.I32_rotl(v277, int32(6))
										v282 = v273 + v269
										v283 = v277 + v282
										v284 = v281 + v283
										v288 = v282 - v281 ^ base.I32_rotl(v281, int32(8))
										v292 = v283 - v288 ^ base.I32_rotl(v288, int32(16))
										v296 = v284 - v292 ^ base.I32_rotl(v292, int32(19))
										v297 = v288 + v284
										v298 = v292 + v297
										v299 = v296 + v298
										v303 = v297 - v296 ^ base.I32_rotl(v296, v275)
										v304 = int32(12)
										v305 = v262 + v304
										v307 = v263 - v304
										if base.Ui32(int32(11)) < base.Ui32(v307) {
											v262 = v305
											v263 = v307
											v264 = v298
											v265 = v299
											v266 = v303
											continue
										} else {
											break
										}
										break
									}
									v310 = v305
									v311 = v307
									v312 = v298
									v313 = v299
									v314 = v303
								} else {
									v310 = v137
									v311 = v145
									v312 = v153
									v313 = v153
									v314 = v153
								}
								switch v311 - int32(1) {
								case 0:
									v373 = v312
									v374 = v313
									v375 = v314
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 1:
									v366 = v312
									v367 = v313
									v368 = v314
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 2:
									v359 = v312
									v360 = v313
									v361 = v314
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 3:
									v353 = v313
									v354 = v314
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 4:
									v349 = v313
									v350 = v314
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 5:
									v343 = v313
									v344 = v314
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 6:
									v337 = v313
									v338 = v314
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 7:
									v332 = v314
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 8:
									v327 = v314
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 9:
									v322 = v314
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
									v327 = v323<<(uint(int32(16))%32) + v322
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 10:
									v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+10)))
									v322 = v318<<(uint(int32(24))%32) + v314
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
									v327 = v323<<(uint(int32(16))%32) + v322
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								default:
									v380 = v312
									v381 = v313
									v382 = v314
								}
							} else {
								if base.Ui32(v145) < base.Ui32(int32(12)) {
									v208 = v137
									v209 = v145
									v210 = v153
									v211 = v153
									v212 = v153
								} else {
									v160 = v137
									v161 = v145
									v162 = v153
									v163 = v153
									v164 = v153
									for {
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
										v167 = v166 + v163
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
										v171 = v170 + v164
										v173 = int32(4)
										v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
										v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
										v180 = v171 + v167
										v181 = v175 + v180
										v182 = v179 + v181
										v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
										v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
										v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
										v195 = v186 + v182
										v196 = v190 + v195
										v197 = v194 + v196
										v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
										v202 = int32(12)
										v203 = v160 + v202
										v205 = v161 - v202
										if base.Ui32(int32(11)) < base.Ui32(v205) {
											v160 = v203
											v161 = v205
											v162 = v196
											v163 = v197
											v164 = v201
											continue
										} else {
											break
										}
										break
									}
									v208 = v203
									v209 = v205
									v210 = v196
									v211 = v197
									v212 = v201
								}
								switch v209 - int32(1) {
								case 0:
									v259 = v210
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 1:
									v254 = v210
									v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
									v259 = v255<<(uint(int32(8))%32) + v254
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 2:
									v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
									v254 = v250<<(uint(int32(16))%32) + v210
									v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
									v259 = v255<<(uint(int32(8))%32) + v254
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 3:
									v247 = v211
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 4:
									v244 = v211
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 5:
									v239 = v211
									v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
									v244 = v240<<(uint(int32(8))%32) + v239
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 6:
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
									v239 = v235<<(uint(int32(16))%32) + v211
									v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
									v244 = v240<<(uint(int32(8))%32) + v239
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 7:
									v230 = v212
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 8:
									v225 = v212
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 9:
									v220 = v212
									v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
									v225 = v221<<(uint(int32(16))%32) + v220
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 10:
									v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
									v220 = v216<<(uint(int32(24))%32) + v212
									v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
									v225 = v221<<(uint(int32(16))%32) + v220
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								default:
									v380 = v210
									v381 = v211
									v382 = v212
								}
							}
							v385 = int32(14)
							v387 = v381 ^ v382 - base.I32_rotl(v381, v385)
							v391 = v387 ^ v380 - base.I32_rotl(v387, int32(11))
							v395 = v391 ^ v381 - base.I32_rotl(v391, int32(25))
							v399 = v395 ^ v387 - base.I32_rotl(v395, int32(16))
							v403 = v399 ^ v391 - base.I32_rotl(v399, int32(4))
							v407 = v403 ^ v395 - base.I32_rotl(v403, v385)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v407 ^ v399 - base.I32_rotl(v407, int32(24))
							v414 = v10 + int32(38)
							v419 = F_pg_snprintf(m, v414, int32(10), int32(_a_F_make_scalar_key_0), v10+int32(16))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int64(0)
							} else {
								v424 = int32(8)
								v425 = int32(20)
								v426 = v414
								v428 = v424 + int32(5)
								v429 = F_palloc(m, v428)
								mBase = m.M
								v430 = m.ExcPending
								if v430 != 0 {
									return int64(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
									*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
									if v424 != 0 {
										base.MemoryCopy(m, v429+int32(5), v426, v424)
									} else {
									}
									F_pfree(m, v137)
									mBase = m.M
									v439 = m.ExcPending
									if v439 != 0 {
										return int64(0)
									} else {
										v760 = v429
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v760)
									}
								}
							}
						}
					}
				} else {
					v41 = F_pstrdup(m, int32(_a_F_make_scalar_key_5))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v137 = v41
						m.G0 = v31 + int32(32)
						v145 = F_strlen(m, v137)
						mBase = m.M
						if v145 < int32(126) {
							v424 = v145
							v425 = int32(4)
							v426 = v137
							v428 = v424 + int32(5)
							v429 = F_palloc(m, v428)
							mBase = m.M
							v430 = m.ExcPending
							if v430 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
								*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
								if v424 != 0 {
									base.MemoryCopy(m, v429+int32(5), v426, v424)
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return int64(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return base.I64_extend_i32_u(v760)
								}
							}
						} else {
							v153 = v145 - int32(1636608432)
							if v137&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v145) {
									v262 = v137
									v263 = v145
									v264 = v153
									v265 = v153
									v266 = v153
									for {
										v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
										v269 = v268 + v265
										v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
										v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
										v273 = v272 + v266
										v275 = int32(4)
										v277 = v270 + v264 - v273 ^ base.I32_rotl(v273, v275)
										v281 = v269 - v277 ^ base.I32_rotl(v277, int32(6))
										v282 = v273 + v269
										v283 = v277 + v282
										v284 = v281 + v283
										v288 = v282 - v281 ^ base.I32_rotl(v281, int32(8))
										v292 = v283 - v288 ^ base.I32_rotl(v288, int32(16))
										v296 = v284 - v292 ^ base.I32_rotl(v292, int32(19))
										v297 = v288 + v284
										v298 = v292 + v297
										v299 = v296 + v298
										v303 = v297 - v296 ^ base.I32_rotl(v296, v275)
										v304 = int32(12)
										v305 = v262 + v304
										v307 = v263 - v304
										if base.Ui32(int32(11)) < base.Ui32(v307) {
											v262 = v305
											v263 = v307
											v264 = v298
											v265 = v299
											v266 = v303
											continue
										} else {
											break
										}
										break
									}
									v310 = v305
									v311 = v307
									v312 = v298
									v313 = v299
									v314 = v303
								} else {
									v310 = v137
									v311 = v145
									v312 = v153
									v313 = v153
									v314 = v153
								}
								switch v311 - int32(1) {
								case 0:
									v373 = v312
									v374 = v313
									v375 = v314
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 1:
									v366 = v312
									v367 = v313
									v368 = v314
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 2:
									v359 = v312
									v360 = v313
									v361 = v314
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 3:
									v353 = v313
									v354 = v314
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 4:
									v349 = v313
									v350 = v314
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 5:
									v343 = v313
									v344 = v314
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 6:
									v337 = v313
									v338 = v314
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 7:
									v332 = v314
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 8:
									v327 = v314
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 9:
									v322 = v314
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
									v327 = v323<<(uint(int32(16))%32) + v322
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								case 10:
									v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+10)))
									v322 = v318<<(uint(int32(24))%32) + v314
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
									v327 = v323<<(uint(int32(16))%32) + v322
									v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
									v332 = v328<<(uint(int32(8))%32) + v327
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
									v337 = v333<<(uint(int32(24))%32) + v313
									v338 = v332
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
									v343 = v339<<(uint(int32(16))%32) + v337
									v344 = v338
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
									v349 = v345<<(uint(int32(8))%32) + v343
									v350 = v344
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
									v353 = v349 + v351
									v354 = v350
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
									v359 = v355<<(uint(int32(24))%32) + v312
									v360 = v353
									v361 = v354
									v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
									v366 = v362<<(uint(int32(16))%32) + v359
									v367 = v360
									v368 = v361
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
									v373 = v369<<(uint(int32(8))%32) + v366
									v374 = v367
									v375 = v368
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
									v380 = v373 + v376
									v381 = v374
									v382 = v375
								default:
									v380 = v312
									v381 = v313
									v382 = v314
								}
							} else {
								if base.Ui32(v145) < base.Ui32(int32(12)) {
									v208 = v137
									v209 = v145
									v210 = v153
									v211 = v153
									v212 = v153
								} else {
									v160 = v137
									v161 = v145
									v162 = v153
									v163 = v153
									v164 = v153
									for {
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
										v167 = v166 + v163
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
										v171 = v170 + v164
										v173 = int32(4)
										v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
										v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
										v180 = v171 + v167
										v181 = v175 + v180
										v182 = v179 + v181
										v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
										v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
										v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
										v195 = v186 + v182
										v196 = v190 + v195
										v197 = v194 + v196
										v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
										v202 = int32(12)
										v203 = v160 + v202
										v205 = v161 - v202
										if base.Ui32(int32(11)) < base.Ui32(v205) {
											v160 = v203
											v161 = v205
											v162 = v196
											v163 = v197
											v164 = v201
											continue
										} else {
											break
										}
										break
									}
									v208 = v203
									v209 = v205
									v210 = v196
									v211 = v197
									v212 = v201
								}
								switch v209 - int32(1) {
								case 0:
									v259 = v210
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 1:
									v254 = v210
									v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
									v259 = v255<<(uint(int32(8))%32) + v254
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 2:
									v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
									v254 = v250<<(uint(int32(16))%32) + v210
									v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
									v259 = v255<<(uint(int32(8))%32) + v254
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
									v380 = v259 + v260
									v381 = v211
									v382 = v212
								case 3:
									v247 = v211
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 4:
									v244 = v211
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 5:
									v239 = v211
									v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
									v244 = v240<<(uint(int32(8))%32) + v239
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 6:
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
									v239 = v235<<(uint(int32(16))%32) + v211
									v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
									v244 = v240<<(uint(int32(8))%32) + v239
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
									v247 = v244 + v245
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v380 = v248 + v210
									v381 = v247
									v382 = v212
								case 7:
									v230 = v212
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 8:
									v225 = v212
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 9:
									v220 = v212
									v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
									v225 = v221<<(uint(int32(16))%32) + v220
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								case 10:
									v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
									v220 = v216<<(uint(int32(24))%32) + v212
									v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
									v225 = v221<<(uint(int32(16))%32) + v220
									v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
									v230 = v226<<(uint(int32(8))%32) + v225
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
									v380 = v231 + v210
									v381 = v233 + v211
									v382 = v230
								default:
									v380 = v210
									v381 = v211
									v382 = v212
								}
							}
							v385 = int32(14)
							v387 = v381 ^ v382 - base.I32_rotl(v381, v385)
							v391 = v387 ^ v380 - base.I32_rotl(v387, int32(11))
							v395 = v391 ^ v381 - base.I32_rotl(v391, int32(25))
							v399 = v395 ^ v387 - base.I32_rotl(v395, int32(16))
							v403 = v399 ^ v391 - base.I32_rotl(v399, int32(4))
							v407 = v403 ^ v395 - base.I32_rotl(v403, v385)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v407 ^ v399 - base.I32_rotl(v407, int32(24))
							v414 = v10 + int32(38)
							v419 = F_pg_snprintf(m, v414, int32(10), int32(_a_F_make_scalar_key_0), v10+int32(16))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int64(0)
							} else {
								v424 = int32(8)
								v425 = int32(20)
								v426 = v414
								v428 = v424 + int32(5)
								v429 = F_palloc(m, v428)
								mBase = m.M
								v430 = m.ExcPending
								if v430 != 0 {
									return int64(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
									*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
									if v424 != 0 {
										base.MemoryCopy(m, v429+int32(5), v426, v424)
									} else {
									}
									F_pfree(m, v137)
									mBase = m.M
									v439 = m.ExcPending
									if v439 != 0 {
										return int64(0)
									} else {
										v760 = v429
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v760)
									}
								}
							}
						}
					}
				}
			} else {
				v44 = F_pstrdup(m, int32(_a_F_make_scalar_key_6))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					v137 = v44
					m.G0 = v31 + int32(32)
					v145 = F_strlen(m, v137)
					mBase = m.M
					if v145 < int32(126) {
						v424 = v145
						v425 = int32(4)
						v426 = v137
						v428 = v424 + int32(5)
						v429 = F_palloc(m, v428)
						mBase = m.M
						v430 = m.ExcPending
						if v430 != 0 {
							return int64(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
							*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
							if v424 != 0 {
								base.MemoryCopy(m, v429+int32(5), v426, v424)
							} else {
							}
							F_pfree(m, v137)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return int64(0)
							} else {
								v760 = v429
								m.G0 = v10 + int32(48)
								return base.I64_extend_i32_u(v760)
							}
						}
					} else {
						v153 = v145 - int32(1636608432)
						if v137&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v145) {
								v262 = v137
								v263 = v145
								v264 = v153
								v265 = v153
								v266 = v153
								for {
									v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
									v269 = v268 + v265
									v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
									v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
									v273 = v272 + v266
									v275 = int32(4)
									v277 = v270 + v264 - v273 ^ base.I32_rotl(v273, v275)
									v281 = v269 - v277 ^ base.I32_rotl(v277, int32(6))
									v282 = v273 + v269
									v283 = v277 + v282
									v284 = v281 + v283
									v288 = v282 - v281 ^ base.I32_rotl(v281, int32(8))
									v292 = v283 - v288 ^ base.I32_rotl(v288, int32(16))
									v296 = v284 - v292 ^ base.I32_rotl(v292, int32(19))
									v297 = v288 + v284
									v298 = v292 + v297
									v299 = v296 + v298
									v303 = v297 - v296 ^ base.I32_rotl(v296, v275)
									v304 = int32(12)
									v305 = v262 + v304
									v307 = v263 - v304
									if base.Ui32(int32(11)) < base.Ui32(v307) {
										v262 = v305
										v263 = v307
										v264 = v298
										v265 = v299
										v266 = v303
										continue
									} else {
										break
									}
									break
								}
								v310 = v305
								v311 = v307
								v312 = v298
								v313 = v299
								v314 = v303
							} else {
								v310 = v137
								v311 = v145
								v312 = v153
								v313 = v153
								v314 = v153
							}
							switch v311 - int32(1) {
							case 0:
								v373 = v312
								v374 = v313
								v375 = v314
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 1:
								v366 = v312
								v367 = v313
								v368 = v314
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 2:
								v359 = v312
								v360 = v313
								v361 = v314
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 3:
								v353 = v313
								v354 = v314
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 4:
								v349 = v313
								v350 = v314
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 5:
								v343 = v313
								v344 = v314
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 6:
								v337 = v313
								v338 = v314
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
								v343 = v339<<(uint(int32(16))%32) + v337
								v344 = v338
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 7:
								v332 = v314
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
								v337 = v333<<(uint(int32(24))%32) + v313
								v338 = v332
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
								v343 = v339<<(uint(int32(16))%32) + v337
								v344 = v338
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 8:
								v327 = v314
								v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
								v332 = v328<<(uint(int32(8))%32) + v327
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
								v337 = v333<<(uint(int32(24))%32) + v313
								v338 = v332
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
								v343 = v339<<(uint(int32(16))%32) + v337
								v344 = v338
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 9:
								v322 = v314
								v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
								v327 = v323<<(uint(int32(16))%32) + v322
								v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
								v332 = v328<<(uint(int32(8))%32) + v327
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
								v337 = v333<<(uint(int32(24))%32) + v313
								v338 = v332
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
								v343 = v339<<(uint(int32(16))%32) + v337
								v344 = v338
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							case 10:
								v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+10)))
								v322 = v318<<(uint(int32(24))%32) + v314
								v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
								v327 = v323<<(uint(int32(16))%32) + v322
								v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
								v332 = v328<<(uint(int32(8))%32) + v327
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
								v337 = v333<<(uint(int32(24))%32) + v313
								v338 = v332
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
								v343 = v339<<(uint(int32(16))%32) + v337
								v344 = v338
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
								v349 = v345<<(uint(int32(8))%32) + v343
								v350 = v344
								v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
								v353 = v349 + v351
								v354 = v350
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
								v359 = v355<<(uint(int32(24))%32) + v312
								v360 = v353
								v361 = v354
								v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
								v366 = v362<<(uint(int32(16))%32) + v359
								v367 = v360
								v368 = v361
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
								v373 = v369<<(uint(int32(8))%32) + v366
								v374 = v367
								v375 = v368
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
								v380 = v373 + v376
								v381 = v374
								v382 = v375
							default:
								v380 = v312
								v381 = v313
								v382 = v314
							}
						} else {
							if base.Ui32(v145) < base.Ui32(int32(12)) {
								v208 = v137
								v209 = v145
								v210 = v153
								v211 = v153
								v212 = v153
							} else {
								v160 = v137
								v161 = v145
								v162 = v153
								v163 = v153
								v164 = v153
								for {
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
									v167 = v166 + v163
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
									v171 = v170 + v164
									v173 = int32(4)
									v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
									v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
									v180 = v171 + v167
									v181 = v175 + v180
									v182 = v179 + v181
									v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
									v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
									v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
									v195 = v186 + v182
									v196 = v190 + v195
									v197 = v194 + v196
									v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
									v202 = int32(12)
									v203 = v160 + v202
									v205 = v161 - v202
									if base.Ui32(int32(11)) < base.Ui32(v205) {
										v160 = v203
										v161 = v205
										v162 = v196
										v163 = v197
										v164 = v201
										continue
									} else {
										break
									}
									break
								}
								v208 = v203
								v209 = v205
								v210 = v196
								v211 = v197
								v212 = v201
							}
							switch v209 - int32(1) {
							case 0:
								v259 = v210
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
								v380 = v259 + v260
								v381 = v211
								v382 = v212
							case 1:
								v254 = v210
								v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
								v259 = v255<<(uint(int32(8))%32) + v254
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
								v380 = v259 + v260
								v381 = v211
								v382 = v212
							case 2:
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
								v254 = v250<<(uint(int32(16))%32) + v210
								v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
								v259 = v255<<(uint(int32(8))%32) + v254
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
								v380 = v259 + v260
								v381 = v211
								v382 = v212
							case 3:
								v247 = v211
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v380 = v248 + v210
								v381 = v247
								v382 = v212
							case 4:
								v244 = v211
								v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
								v247 = v244 + v245
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v380 = v248 + v210
								v381 = v247
								v382 = v212
							case 5:
								v239 = v211
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
								v244 = v240<<(uint(int32(8))%32) + v239
								v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
								v247 = v244 + v245
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v380 = v248 + v210
								v381 = v247
								v382 = v212
							case 6:
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
								v239 = v235<<(uint(int32(16))%32) + v211
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
								v244 = v240<<(uint(int32(8))%32) + v239
								v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
								v247 = v244 + v245
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v380 = v248 + v210
								v381 = v247
								v382 = v212
							case 7:
								v230 = v212
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
								v380 = v231 + v210
								v381 = v233 + v211
								v382 = v230
							case 8:
								v225 = v212
								v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
								v230 = v226<<(uint(int32(8))%32) + v225
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
								v380 = v231 + v210
								v381 = v233 + v211
								v382 = v230
							case 9:
								v220 = v212
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
								v225 = v221<<(uint(int32(16))%32) + v220
								v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
								v230 = v226<<(uint(int32(8))%32) + v225
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
								v380 = v231 + v210
								v381 = v233 + v211
								v382 = v230
							case 10:
								v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
								v220 = v216<<(uint(int32(24))%32) + v212
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
								v225 = v221<<(uint(int32(16))%32) + v220
								v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
								v230 = v226<<(uint(int32(8))%32) + v225
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
								v380 = v231 + v210
								v381 = v233 + v211
								v382 = v230
							default:
								v380 = v210
								v381 = v211
								v382 = v212
							}
						}
						v385 = int32(14)
						v387 = v381 ^ v382 - base.I32_rotl(v381, v385)
						v391 = v387 ^ v380 - base.I32_rotl(v387, int32(11))
						v395 = v391 ^ v381 - base.I32_rotl(v391, int32(25))
						v399 = v395 ^ v387 - base.I32_rotl(v395, int32(16))
						v403 = v399 ^ v391 - base.I32_rotl(v399, int32(4))
						v407 = v403 ^ v395 - base.I32_rotl(v403, v385)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v407 ^ v399 - base.I32_rotl(v407, int32(24))
						v414 = v10 + int32(38)
						v419 = F_pg_snprintf(m, v414, int32(10), int32(_a_F_make_scalar_key_0), v10+int32(16))
						mBase = m.M
						v420 = m.ExcPending
						if v420 != 0 {
							return int64(0)
						} else {
							v424 = int32(8)
							v425 = int32(20)
							v426 = v414
							v428 = v424 + int32(5)
							v429 = F_palloc(m, v428)
							mBase = m.M
							v430 = m.ExcPending
							if v430 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
								*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
								if v424 != 0 {
									base.MemoryCopy(m, v429+int32(5), v426, v424)
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return int64(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return base.I64_extend_i32_u(v760)
								}
							}
						}
					}
				}
			}
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v54 = base.I32_extend16_s(v33)
			v56 = base.B2i32(int32(0) <= v54)
			if int32(0) <= v54 {
				v57 = int32(-8)
			} else {
				v57 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(base.Ui32(int32(base.Ui32(v49)>>(uint(int32(2))%32))+v57) >> (uint(int32(1)) % 32))
			if int32(0) <= v54 {
				v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+6)))
				v72 = v62
			} else {
				v72 = v33<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v33&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v72
			v74 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v74
			v83 = base.B2i32(v54 < v74)
			if v54 < v74 {
				v84 = int32(base.Ui32(v33)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v84 = v33 & int32(_a_F_make_scalar_key_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v84
			v91 = v33 & int32(_a_F_make_scalar_key_1)
			if v91 == int32(_a_F_make_scalar_key_8) {
				v94 = v33 << (uint(int32(1)) % 32) & int32(_a_F_make_scalar_key_9)
			} else {
				v94 = v91
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v94
			if v54 < v74 {
				v98 = int32(6)
			} else {
				v98 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v28 + v98
			v103 = F_get_str_from_var(m, v31+int32(8))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int64(0)
			} else {
				v105 = int32(46)
				v106 = F___strchrnul(m, v103, v105)
				mBase = m.M
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
				if v108 == v105 {
					v112 = v106
				} else {
					v112 = int32(0)
				}
				if v112 == int32(0) {
					v137 = v103
				} else {
					v115 = F_strlen(m, v103)
					mBase = m.M
					v116 = v115
					for {
						v124 = v116 - int32(1)
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v124))))
						if v126 == int32(48) {
							v116 = v124
							continue
						} else {
							break
						}
						break
					}
					if v126 != int32(46) {
						v131 = v116
					} else {
						v131 = v124
					}
					v133 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v131+v103))) = uint8(v133)
					v137 = v103
				}
				m.G0 = v31 + int32(32)
				v145 = F_strlen(m, v137)
				mBase = m.M
				if v145 < int32(126) {
					v424 = v145
					v425 = int32(4)
					v426 = v137
					v428 = v424 + int32(5)
					v429 = F_palloc(m, v428)
					mBase = m.M
					v430 = m.ExcPending
					if v430 != 0 {
						return int64(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
						*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
						if v424 != 0 {
							base.MemoryCopy(m, v429+int32(5), v426, v424)
						} else {
						}
						F_pfree(m, v137)
						mBase = m.M
						v439 = m.ExcPending
						if v439 != 0 {
							return int64(0)
						} else {
							v760 = v429
							m.G0 = v10 + int32(48)
							return base.I64_extend_i32_u(v760)
						}
					}
				} else {
					v153 = v145 - int32(1636608432)
					if v137&int32(3) != 0 {
						if base.Ui32(int32(11)) < base.Ui32(v145) {
							v262 = v137
							v263 = v145
							v264 = v153
							v265 = v153
							v266 = v153
							for {
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
								v269 = v268 + v265
								v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
								v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
								v273 = v272 + v266
								v275 = int32(4)
								v277 = v270 + v264 - v273 ^ base.I32_rotl(v273, v275)
								v281 = v269 - v277 ^ base.I32_rotl(v277, int32(6))
								v282 = v273 + v269
								v283 = v277 + v282
								v284 = v281 + v283
								v288 = v282 - v281 ^ base.I32_rotl(v281, int32(8))
								v292 = v283 - v288 ^ base.I32_rotl(v288, int32(16))
								v296 = v284 - v292 ^ base.I32_rotl(v292, int32(19))
								v297 = v288 + v284
								v298 = v292 + v297
								v299 = v296 + v298
								v303 = v297 - v296 ^ base.I32_rotl(v296, v275)
								v304 = int32(12)
								v305 = v262 + v304
								v307 = v263 - v304
								if base.Ui32(int32(11)) < base.Ui32(v307) {
									v262 = v305
									v263 = v307
									v264 = v298
									v265 = v299
									v266 = v303
									continue
								} else {
									break
								}
								break
							}
							v310 = v305
							v311 = v307
							v312 = v298
							v313 = v299
							v314 = v303
						} else {
							v310 = v137
							v311 = v145
							v312 = v153
							v313 = v153
							v314 = v153
						}
						switch v311 - int32(1) {
						case 0:
							v373 = v312
							v374 = v313
							v375 = v314
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 1:
							v366 = v312
							v367 = v313
							v368 = v314
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 2:
							v359 = v312
							v360 = v313
							v361 = v314
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 3:
							v353 = v313
							v354 = v314
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 4:
							v349 = v313
							v350 = v314
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 5:
							v343 = v313
							v344 = v314
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 6:
							v337 = v313
							v338 = v314
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
							v343 = v339<<(uint(int32(16))%32) + v337
							v344 = v338
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 7:
							v332 = v314
							v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
							v337 = v333<<(uint(int32(24))%32) + v313
							v338 = v332
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
							v343 = v339<<(uint(int32(16))%32) + v337
							v344 = v338
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 8:
							v327 = v314
							v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
							v332 = v328<<(uint(int32(8))%32) + v327
							v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
							v337 = v333<<(uint(int32(24))%32) + v313
							v338 = v332
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
							v343 = v339<<(uint(int32(16))%32) + v337
							v344 = v338
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 9:
							v322 = v314
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
							v327 = v323<<(uint(int32(16))%32) + v322
							v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
							v332 = v328<<(uint(int32(8))%32) + v327
							v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
							v337 = v333<<(uint(int32(24))%32) + v313
							v338 = v332
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
							v343 = v339<<(uint(int32(16))%32) + v337
							v344 = v338
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						case 10:
							v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+10)))
							v322 = v318<<(uint(int32(24))%32) + v314
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+9)))
							v327 = v323<<(uint(int32(16))%32) + v322
							v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+8)))
							v332 = v328<<(uint(int32(8))%32) + v327
							v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+7)))
							v337 = v333<<(uint(int32(24))%32) + v313
							v338 = v332
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+6)))
							v343 = v339<<(uint(int32(16))%32) + v337
							v344 = v338
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+5)))
							v349 = v345<<(uint(int32(8))%32) + v343
							v350 = v344
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
							v353 = v349 + v351
							v354 = v350
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+3)))
							v359 = v355<<(uint(int32(24))%32) + v312
							v360 = v353
							v361 = v354
							v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+2)))
							v366 = v362<<(uint(int32(16))%32) + v359
							v367 = v360
							v368 = v361
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
							v373 = v369<<(uint(int32(8))%32) + v366
							v374 = v367
							v375 = v368
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
							v380 = v373 + v376
							v381 = v374
							v382 = v375
						default:
							v380 = v312
							v381 = v313
							v382 = v314
						}
					} else {
						if base.Ui32(v145) < base.Ui32(int32(12)) {
							v208 = v137
							v209 = v145
							v210 = v153
							v211 = v153
							v212 = v153
						} else {
							v160 = v137
							v161 = v145
							v162 = v153
							v163 = v153
							v164 = v153
							for {
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
								v167 = v166 + v163
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
								v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
								v171 = v170 + v164
								v173 = int32(4)
								v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
								v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
								v180 = v171 + v167
								v181 = v175 + v180
								v182 = v179 + v181
								v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
								v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
								v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
								v195 = v186 + v182
								v196 = v190 + v195
								v197 = v194 + v196
								v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
								v202 = int32(12)
								v203 = v160 + v202
								v205 = v161 - v202
								if base.Ui32(int32(11)) < base.Ui32(v205) {
									v160 = v203
									v161 = v205
									v162 = v196
									v163 = v197
									v164 = v201
									continue
								} else {
									break
								}
								break
							}
							v208 = v203
							v209 = v205
							v210 = v196
							v211 = v197
							v212 = v201
						}
						switch v209 - int32(1) {
						case 0:
							v259 = v210
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
							v380 = v259 + v260
							v381 = v211
							v382 = v212
						case 1:
							v254 = v210
							v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
							v259 = v255<<(uint(int32(8))%32) + v254
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
							v380 = v259 + v260
							v381 = v211
							v382 = v212
						case 2:
							v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
							v254 = v250<<(uint(int32(16))%32) + v210
							v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
							v259 = v255<<(uint(int32(8))%32) + v254
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
							v380 = v259 + v260
							v381 = v211
							v382 = v212
						case 3:
							v247 = v211
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v380 = v248 + v210
							v381 = v247
							v382 = v212
						case 4:
							v244 = v211
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							v247 = v244 + v245
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v380 = v248 + v210
							v381 = v247
							v382 = v212
						case 5:
							v239 = v211
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							v247 = v244 + v245
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v380 = v248 + v210
							v381 = v247
							v382 = v212
						case 6:
							v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
							v239 = v235<<(uint(int32(16))%32) + v211
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							v247 = v244 + v245
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v380 = v248 + v210
							v381 = v247
							v382 = v212
						case 7:
							v230 = v212
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
							v380 = v231 + v210
							v381 = v233 + v211
							v382 = v230
						case 8:
							v225 = v212
							v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
							v230 = v226<<(uint(int32(8))%32) + v225
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
							v380 = v231 + v210
							v381 = v233 + v211
							v382 = v230
						case 9:
							v220 = v212
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
							v225 = v221<<(uint(int32(16))%32) + v220
							v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
							v230 = v226<<(uint(int32(8))%32) + v225
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
							v380 = v231 + v210
							v381 = v233 + v211
							v382 = v230
						case 10:
							v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
							v220 = v216<<(uint(int32(24))%32) + v212
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
							v225 = v221<<(uint(int32(16))%32) + v220
							v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
							v230 = v226<<(uint(int32(8))%32) + v225
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
							v380 = v231 + v210
							v381 = v233 + v211
							v382 = v230
						default:
							v380 = v210
							v381 = v211
							v382 = v212
						}
					}
					v385 = int32(14)
					v387 = v381 ^ v382 - base.I32_rotl(v381, v385)
					v391 = v387 ^ v380 - base.I32_rotl(v387, int32(11))
					v395 = v391 ^ v381 - base.I32_rotl(v391, int32(25))
					v399 = v395 ^ v387 - base.I32_rotl(v395, int32(16))
					v403 = v399 ^ v391 - base.I32_rotl(v399, int32(4))
					v407 = v403 ^ v395 - base.I32_rotl(v403, v385)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v407 ^ v399 - base.I32_rotl(v407, int32(24))
					v414 = v10 + int32(38)
					v419 = F_pg_snprintf(m, v414, int32(10), int32(_a_F_make_scalar_key_0), v10+int32(16))
					mBase = m.M
					v420 = m.ExcPending
					if v420 != 0 {
						return int64(0)
					} else {
						v424 = int32(8)
						v425 = int32(20)
						v426 = v414
						v428 = v424 + int32(5)
						v429 = F_palloc(m, v428)
						mBase = m.M
						v430 = m.ExcPending
						if v430 != 0 {
							return int64(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v429)+4)) = uint8(v425)
							*(*int32)(unsafe.Add(mBase, uint32(v429))) = v428 << (uint(int32(2)) % 32)
							if v424 != 0 {
								base.MemoryCopy(m, v429+int32(5), v426, v424)
							} else {
							}
							F_pfree(m, v137)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return int64(0)
							} else {
								v760 = v429
								m.G0 = v10 + int32(48)
								return base.I64_extend_i32_u(v760)
							}
						}
					}
				}
			}
		}
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		v15 = F_palloc(m, int32(6))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v13 != 0 {
				v21 = int32(116)
			} else {
				v21 = int32(102)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(v21)
			v23 = int32(3)
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(24)
			v760 = v15
			m.G0 = v10 + int32(48)
			return base.I64_extend_i32_u(v760)
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v742 = m.ExcPending
		if v742 != 0 {
			return int64(0)
		} else {
			v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v743
			F_errmsg_internal(m, int32(_a_F_make_scalar_key_10), v10)
			mBase = m.M
			v747 = m.ExcPending
			if v747 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_make_scalar_key_11), int32(1404), int32(_a_F_make_scalar_key_12))
				mBase = m.M
				v752 = m.ExcPending
				if v752 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_makepol_3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int64
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L4
L3:
	;
	m.G0 = v10 + int32(80)
	return v357
L4:
	;
	v25 = int32(0)
	goto L6
L5:
	;
	v335 = int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v337 = F_errsave_start(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L76
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = int32(0)
	v37 = v31
	goto L8
L7:
	;
	goto L5
L8:
	;
	switch v37 - int32(1) {
	case 0:
		goto L15
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v325 + int32(1)
	if base.Ui32(v322) < base.Ui32(int32(16)) {
		v34 = v322
		v37 = v324
		goto L8
	} else {
		goto L75
	}
L12:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if base.Ui32((v240-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	if v25 == int32(16) {
		goto L51
	} else {
		goto L52
	}
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(32) {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if base.Ui32((v43-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v101 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+64)) = uint8(v104)
	v322 = int32(1)
	v324 = v101
	v325 = v42
	goto L11
L17:
	;
	v50 = int32(0)
	v51 = int32(1)
	switch v43 - int32(32) {
	case 0:
		v322 = v50
		v324 = v51
		v325 = v42
		goto L11
	case 1:
		goto L18
	default:
		goto L10
	case 8:
		goto L19
	case 13:
		goto L16
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 + int32(1)
	v209 = int32(33)
	goto L13
L19:
	;
	v54 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 + v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v57 + v54
	v61 = F_makepol_3(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v61 != 0 {
		v357 = v51
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v25 == int32(0) {
		v25 = v50
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v68 = v25
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10+v68<<(uint(int32(2))%32)-int32(4))))
	switch v77 - int32(33) {
	case 0, 5:
		goto L25
	default:
		v25 = v68
		goto L6
	}
L24:
	;
	goto L4
L25:
	;
	v81 = F_palloc(m, int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = int32(3)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v81
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v90 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v89 + v90
	v94 = v68 - v90
	if v94 != 0 {
		v68 = v94
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v322 = v34
	v324 = int32(3)
	v325 = v107
	goto L11
L29:
	;
	goto L30
L30:
	;
	switch v108 - int32(38) {
	case 0:
		goto L33
	case 1, 2:
		goto L10
	case 3:
		goto L32
	default:
		goto L34
	}
L31:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v179 != 0 {
		goto L10
	} else {
		goto L45
	}
L32:
	;
	v141 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107 + v141
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 - v141
	if v144 <= int32(0) {
		goto L10
	} else {
		goto L39
	}
L33:
	;
	v118 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107 + v118
	if base.B2i32(v25 == int32(0))|base.B2i32(v120 != int32(124)) != 0 {
		v209 = v120
		goto L13
	} else {
		goto L37
	}
L34:
	;
	if v108 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v108 != int32(124) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v130 = F_palloc(m, int32(12))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(532575944707)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v130
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v137 + int32(1)
	goto L6
L39:
	;
	v150 = int32(0)
	if v25 == v150 {
		v357 = v150
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v156 = v25
	goto L41
L41:
	;
	v161 = v156 - int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v10+v161<<(uint(int32(2))%32))))
	v167 = F_palloc(m, int32(12))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v357 = v150
	goto L3
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = int32(3)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v175 + int32(1)
	if v161 != 0 {
		v156 = v161
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v180 = int32(0)
	if v25 == v180 {
		v357 = v180
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v186 = v25
	goto L47
L47:
	;
	v191 = v186 - int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v10+v191<<(uint(int32(2))%32))))
	v197 = F_palloc(m, int32(12))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v357 = v180
	goto L3
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = int32(3)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v197
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v205 + int32(1)
	if v191 != 0 {
		v186 = v191
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v215 = int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v217 = F_errsave_start(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+v25<<(uint(int32(2))%32)))) = v209
	v25 = v25 + int32(1)
	goto L6
L54:
	;
	if v217 == int32(0) {
		v357 = v215
		goto L3
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(16777477))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_makepol_3_0), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, v216, int32(_a_F_makepol_3_1), int32(184), int32(_a_F_makepol_3_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v357 = v215
	goto L3
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(-64)+v34))) = uint8(v240)
	v322 = v34 + int32(1)
	v324 = int32(2)
	v325 = v239
	goto L11
L60:
	;
	goto L61
L61:
	;
	v255 = v10 - int32(-64)
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34+v255))) = uint8(v257)
	*(*int32)(unsafe.Add(mBase, _c_F_makepol_3[0])) = v257
	v265 = F_strtox_2(m, v255, v257, v257, int64(2147483648))
	mBase = m.M
	goto L62
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_makepol_3[0]))
	if v268 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v271 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v273 == int32(0) {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v277 = F_palloc(m, int32(12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = base.I32_wrap_i64(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(2)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v277
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v285 + int32(1)
	v289 = int32(0)
	if v25 == v289 {
		v25 = v289
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v295 = v25
	goto L70
L70:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v10+v295<<(uint(int32(2))%32)-int32(4))))
	switch v304 - int32(33) {
	case 0, 5:
		goto L72
	default:
		v25 = v295
		goto L6
	}
L71:
	;
	goto L4
L72:
	;
	v308 = F_palloc(m, int32(12))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308)+4)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = int32(3)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v308
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v317 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v316 + v317
	v321 = v295 - v317
	if v321 != 0 {
		v295 = v321
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	goto L10
L76:
	;
	if v337 == int32(0) {
		v357 = v335
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_makepol_3_3), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errsave_finish(m, v336, int32(_a_F_makepol_3_1), int32(211), int32(_a_F_makepol_3_2))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v357 = v335
	goto L3
}
func F_manifest_report_error(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(16)
	F_initStringInfo(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v15 = F_appendStringInfoVA(m, v11, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v23 = v8 + int32(16)
	F_enlargeStringInfo(m, v23, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v27 = F_appendStringInfoVA(m, v23, l1, l2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v27 != 0 {
		v20 = v27
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
	F_errmsg_internal(m, int32(_a_F_manifest_report_error_0), v8)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_manifest_report_error_1), int32(1041), int32(_a_F_manifest_report_error_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mask_lp_flags(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v8) < base.Ui32(int32(25)) {
	} else {
		v14 = int32(base.Ui32(v8+int32(_a_F_mask_lp_flags_0)) >> (uint(int32(2)) % 32))
		if v14&int32(_a_F_mask_lp_flags_1) == int32(0) {
		} else {
			v19 = int32(1)
			v21 = l0 + int32(20)
			v25 = (v14 + v19) & int32(_a_F_mask_lp_flags_1)
			if base.Ui32(int32(3)) <= base.Ui32(v25) {
				v28 = int32(2)
				if base.Ui32(v25) <= base.Ui32(v28) {
					v31 = v28
				} else {
					v31 = v25
				}
				v32 = int32(1)
				v33 = v31 - v32
				v41 = v32
				v42 = int32(0)
				for {
					v49 = v21 + v41<<(uint(int32(2))%32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					if v50&int32(_a_F_mask_lp_flags_2) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 & int32(-98305)
					} else {
					}
					v57 = v49 + int32(4)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v58&int32(_a_F_mask_lp_flags_2) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v58 & int32(-98305)
					} else {
					}
					v64 = int32(2)
					v65 = v41 + v64
					v67 = v42 + v64
					if v67 != v33&int32(-2) {
						v41 = v65
						v42 = v67
						continue
					} else {
						break
					}
					break
				}
				if v33&v32 == int32(0) {
				} else {
					v72 = v65
					v80 = v21 + v72<<(uint(int32(2))%32)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
					if v81&int32(_a_F_mask_lp_flags_2) == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 & int32(-98305)
					}
				}
			} else {
				v72 = v19
				v80 = v21 + v72<<(uint(int32(2))%32)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				if v81&int32(_a_F_mask_lp_flags_2) == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 & int32(-98305)
				}
			}
		}
	}
	return
}
func F_matchLocks(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v18 == v6 {
		v103 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L28
	} else {
		goto L33
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v103
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v21 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v24 != l2 {
		v103 = v6
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v26 <= int32(0) {
		v103 = v6
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	v39 = v6
	v40 = v6
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v43 = int32(2)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39<<(uint(v43)%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 == v43 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v103 = v88
	goto L2
L11:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v50)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v53 = v52
	goto L13
L12:
	;
	v53 = v47
	goto L13
L13:
	;
	if v53 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v90 = v39 + int32(1)
	if v90 != v26 {
		v39 = v90
		v40 = v88
		goto L9
	} else {
		goto L32
	}
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+16)))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_matchLocks[0]))
	if v58 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if l0 != v53 {
		v88 = v40
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v71 == int32(5) {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	switch v56 - int32(68) {
	case 0, 11:
		v88 = v40
		goto L14
	default:
		goto L18
	}
L20:
	;
	goto L21
L21:
	;
	v64 = v56 - int32(68)
	if base.B2i32(v64 == int32(0))|base.B2i32(v64 == int32(14)) != 0 {
		v88 = v40
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	goto L17
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v76 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = F_rangeTableEntry_used(m, l3, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v85 = F_lappend(m, v40, v46)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L28
	} else {
		goto L31
	}
L28:
	;
	return int32(0)
L29:
	;
	if v79 == int32(0) {
		v88 = v40
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v88 = v85
	goto L14
L32:
	;
	goto L10
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v116 + int32(4)
	F_errmsg(m, int32(_a_F_matchLocks_0), v16)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v125 = F_errdetail(m, int32(_a_F_matchLocks_1), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_matchLocks_2), int32(1741), int32(_a_F_matchLocks_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_match_pattern_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v19 != int32(7) {
		v176 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 - int32(-64)
	return v176
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v22 != 0 {
		v176 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_pattern_fixed_prefix(m, l1, l2, l3, v15+int32(-4), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v26 == int32(0) {
		v176 = v7
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v32 = int32(1)
	v36 = int32(25)
	v38 = F_exprType(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v78 != v86 {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v78 = v71
	v79 = v72
	v80 = v76
	v81 = v73
	v82 = v74
	v83 = int32(0)
	v84 = v75
	goto L7
L9:
	;
	v71 = int32(17)
	v72 = int32(1957)
	v73 = int32(1960)
	v74 = int32(1955)
	v75 = v32
	v76 = int32(0)
	goto L8
L10:
	;
	v54 = int32(1042)
	if v38 != v54 {
		v176 = v7
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v42 = int32(2317)
	v43 = int32(2314)
	v44 = int32(98)
	if l4 == int32(2095) {
		v71 = v36
		v72 = v43
		v73 = v42
		v74 = v44
		v75 = v32
		v76 = int32(0)
		goto L8
	} else {
		goto L13
	}
L12:
	;
	switch v38 - int32(17) {
	case 0:
		goto L9
	case 1, 3, 4, 5, 6, 7:
		v176 = v7
		goto L1
	case 2:
		v78 = v36
		v79 = int32(255)
		v80 = v7
		v81 = int32(257)
		v82 = int32(254)
		v83 = v32
		v84 = v32
		goto L7
	case 8:
		goto L11
	default:
		goto L10
	}
L13:
	;
	if l4 == int32(4017) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v36
	v72 = v43
	v73 = v42
	v74 = v44
	v75 = int32(0)
	v76 = int32(3877)
	goto L8
L15:
	;
	goto L16
L16:
	;
	v78 = v36
	v79 = int32(664)
	v80 = v7
	v81 = int32(667)
	v82 = v44
	v83 = v32
	v84 = v32
	goto L7
L17:
	;
	v60 = base.B2i32(l4 != int32(2097))
	if l4 != int32(2097) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = int32(1061)
	goto L20
L19:
	;
	v61 = int32(2329)
	goto L20
L20:
	;
	if l4 != int32(2097) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v64 = int32(1058)
	goto L23
L22:
	;
	v64 = int32(2326)
	goto L23
L23:
	;
	v78 = v54
	v79 = v64
	v80 = v7
	v81 = v61
	v82 = int32(1054)
	v83 = v60
	v84 = v32
	goto L7
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v78
	goto L26
L25:
	;
	goto L26
L26:
	;
	if v26 == int32(2) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = int32(0)
	v92 = F_op_in_opfamily(m, v82, l4)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if v92 == int32(0) {
		v176 = v91
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v96 = int32(0)
	if base.B2i32(l3 == v96)|base.B2i32(l3 == l5) == v96 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = F_get_collation_isdeterministic(m, l3)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v106 = F_make_opclause(m, v82, l0, v85, l5)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	if v102 == int32(0) {
		v176 = v91
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v106
	v113 = F_list_make1_impl(m, int32(1), v15+int32(-56))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v176 = v113
	goto L1
L39:
	;
	if v84 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v117 = F_get_collation_isdeterministic(m, l3)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v117 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v176 = int32(0)
	goto L1
L43:
	;
	if v83 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v120 = F_op_in_opfamily(m, v80, l4)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v120 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v124 = F_make_opclause(m, v80, l0, v85, l5)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v124
	v131 = F_list_make1_impl(m, int32(1), v15+int32(-48))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v176 = v131
	goto L1
L49:
	;
	v141 = F_op_in_opfamily(m, v81, l4)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L53
	}
L50:
	;
	v135 = F_pg_newlocale_from_collation(m, l5)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v137 == int32(1) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v176 = int32(0)
	goto L1
L53:
	;
	if v141 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v176 = int32(0)
	goto L1
L55:
	;
	goto L56
L56:
	;
	v146 = F_make_opclause(m, v81, l0, v85, l5)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v146
	v153 = F_list_make1_impl(m, int32(1), v15+int32(-52))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v155 = F_op_in_opfamily(m, v79, l4)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v155 == int32(0) {
		v176 = v153
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v159 = F_get_opcode(m, v79)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v162 = v15 + int32(-32)
	F_fmgr_info(m, v159, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v165 = F_make_greater_string(m, v85, v162, l5)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	if v165 == int32(0) {
		v176 = v153
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v169 = F_make_opclause(m, v79, l0, v165, l5)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v171 = F_lappend(m, v153, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v176 = v171
	goto L1
}
func F_mbms_add_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = l0
	goto L1
L1:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v6
L3:
	;
	goto L2
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v13 = v11
	goto L6
L5:
	;
	v13 = int32(0)
	goto L6
L6:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = v14
	goto L9
L8:
	;
	v16 = int32(0)
	goto L9
L9:
	;
	if v16 <= v13 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v20 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v57 = F_lappend(m, v6, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L21
	} else {
		goto L23
	}
L13:
	;
	v23 = int32(0)
	if v6 == v23 {
		v33 = v23
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v27 <= v20 {
		v33 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v33 = v29 + v20<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v33 == int32(0))|base.B2i32(v38 <= v20) != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v41 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41+v20<<(uint(int32(2))%32))))
	v49 = F_bms_add_members(m, v44, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v49
	v20 = v20 + int32(1)
	goto L13
L23:
	;
	v6 = v57
	goto L1
}
func F_mcv_get_match_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v431 int64
	_ = v431
	var v432 int64
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v765 int64
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int64
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v26 = F_palloc_mul(m, int32(1), v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	base.MemoryFill(m, v26, l4^int32(1), v30)
	goto L5
L4:
	;
	goto L5
L5:
	;
	if l0 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L214
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L211
	}
L8:
	;
	m.G0 = v22 - int32(-64)
	return v26
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v36 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = l3 + int32(48)
	v55 = int32(0)
	goto L11
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v55<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 == int32(320) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L8
L13:
	;
	v857 = v55 + int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v857 < v858 {
		v55 = v857
		goto L11
	} else {
		goto L210
	}
L14:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v243 != int32(52) {
		goto L74
	} else {
		goto L75
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v68 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v74 = v64
	v75 = v65
	goto L17
L17:
	;
	if v75 != int32(17) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v241 = int32(1)
	v242 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v74 = v68
	v75 = v73
	goto L17
L21:
	;
	v241 = int32(0)
	v242 = v74
	goto L14
L22:
	;
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v80 = F_get_opcode(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_fmgr_info(m, v80, v20+int32(-28))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v88 = v20 + int32(-32)
	v90 = v20 + int32(-36)
	v92 = v20 + int32(-52)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 == int32(27) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v129 != 0 {
		goto L46
	} else {
		goto L47
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v104 = v103
	goto L29
L28:
	;
	v104 = v99
	goto L29
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v105 == int32(27) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = v108
	v111 = v109
	goto L32
L31:
	;
	v110 = v98
	v111 = v105
	goto L32
L32:
	;
	if v111 == int32(7) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L26
L34:
	;
	if v88 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v117 = v110
	v118 = v104
	goto L34
L36:
	;
	goto L37
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v114 != int32(7) {
		v129 = int32(0)
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v117 = v104
	v118 = v110
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v118
	goto L41
L40:
	;
	goto L41
L41:
	;
	if v90 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v117
	goto L44
L43:
	;
	goto L44
L44:
	;
	v121 = int32(1)
	if v92 == int32(0) {
		v129 = v121
		goto L33
	} else {
		goto L45
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(base.B2i32(v111 == int32(7)))
	v129 = v121
	goto L33
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v133 = F_mcv_match_expression(m, v130, l1, l2, v20+int32(-40))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L70
	}
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v135 == int32(0) {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v145 = int32(0)
	goto L51
L51:
	;
	v162 = v40 + v145*int32(24)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+16))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v133))))
	if v165 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	goto L13
L53:
	;
	v225 = v145 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v225) < base.Ui32(v226) {
		v145 = v225
		goto L51
	} else {
		goto L69
	}
L54:
	;
	v218 = v216 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v145+v26))) = uint8(v218)
	goto L53
L55:
	;
	v178 = v145 + v26
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if l4 == v179 {
		goto L53
	} else {
		goto L61
	}
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+32)))
	if v169 != int32(1) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v173 = int32(0)
	if l4 == v173 {
		v216 = v173
		goto L54
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v26))))
	v216 = v177
	goto L54
L61:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)))
	if v181 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if l4 != 0 {
		v216 = base.B2i32(v204 != int64(0)) | v207
		goto L54
	} else {
		goto L68
	}
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v186+v133<<(uint(int32(3))%32))))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v168)+24))
	v192 = F_FunctionCall2Coll(m, v20+int32(-28), v139, v190, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v168)+24))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v197+v133<<(uint(int32(3))%32))))
	v202 = F_FunctionCall2Coll(m, v20+int32(-28), v139, v196, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v204 = v192
	goto L62
L67:
	;
	v204 = v202
	goto L62
L68:
	;
	v216 = v207 & base.B2i32(v204 != int64(0))
	goto L54
L69:
	;
	goto L52
L70:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_get_match_bitmap_0), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1645), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	if v241 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L74:
	;
	if v243 != int32(20) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v522 = F_mcv_match_expression(m, v520, l1, l2, int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L135
	}
L77:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v249 = F_get_opcode(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_fmgr_info(m, v249, v20+int32(-28))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v257 = v20 + int32(-32)
	v259 = v20 + int32(-36)
	v261 = v20 + int32(-41)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v269 == int32(27) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v298 == int32(0) {
		goto L7
	} else {
		goto L100
	}
L81:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v273 = v272
	goto L83
L82:
	;
	v273 = v268
	goto L83
L83:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v274 == int32(27) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = v277
	v280 = v278
	goto L86
L85:
	;
	v279 = v267
	v280 = v274
	goto L86
L86:
	;
	if v280 == int32(7) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L80
L88:
	;
	if v257 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v286 = v279
	v287 = v273
	goto L88
L90:
	;
	goto L91
L91:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v283 != int32(7) {
		v298 = int32(0)
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v286 = v273
	v287 = v279
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v287
	goto L95
L94:
	;
	goto L95
L95:
	;
	if v259 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v286
	goto L98
L97:
	;
	goto L98
L98:
	;
	v290 = int32(1)
	if v261 == int32(0) {
		v298 = v290
		goto L87
	} else {
		goto L99
	}
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(base.B2i32(v280 == int32(7)))
	v298 = v290
	goto L87
L100:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v301 == int32(0) {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+32)))
	if v305 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
	v309 = F_pg_detoast_datum(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v336 = F_mcv_match_expression(m, v333, l1, l2, v20+int32(-40))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L108
	}
L105:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	F_get_typlenbyvalalign(m, v311, v20+int32(-44), v20+int32(-45), v20+int32(-46))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+20)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+19)))
	v323 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+18)))
	F_deconstruct_array(m, v309, v321, v322, v323, v20+int32(-56), v20+int32(-60), v20+int32(-52))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v338 == int32(0) {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v350 = int32(0)
	goto L110
L110:
	;
	v365 = v40 + v350*int32(24)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366+v336))))
	if v368 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	goto L13
L112:
	;
	v517 = v350 + int32(1)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v517) < base.Ui32(v518) {
		v350 = v517
		goto L110
	} else {
		goto L134
	}
L113:
	;
	v495 = v493 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v350+v26))) = uint8(v495)
	goto L112
L114:
	;
	v380 = v350 + v26
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if l4 == v381 {
		goto L112
	} else {
		goto L120
	}
L115:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+32)))
	if v372 != int32(1) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v375 = int32(0)
	if l4 == v375 {
		v493 = v375
		goto L113
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v26))))
	v493 = v379
	goto L113
L120:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+20)))
	v385 = v383 ^ int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v386 <= int32(0) {
		v461 = v385
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if l4 != 0 {
		v493 = v471 | v461
		goto L113
	} else {
		goto L133
	}
L122:
	;
	v395 = int32(0)
	v399 = v385
	v400 = v383
	goto L123
L123:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v395))))
	if v412 == int32(1) {
		v445 = v400
		v447 = v399 & v400
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v461 = v447
	goto L121
L125:
	;
	v449 = v395 + int32(1)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v449 < v450 {
		v395 = v449
		v399 = v447
		v400 = v445
		goto L123
	} else {
		goto L132
	}
L126:
	;
	v415 = int32(1)
	if v399&v415 == v400&v415 {
		v461 = v399
		goto L121
	} else {
		goto L127
	}
L127:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	v423 = int32(3)
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v422+v336<<(uint(v423)%32))))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v427+v395<<(uint(v423)%32))))
	v432 = F_FunctionCall2Coll(m, v20+int32(-28), v342, v426, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+20)))
	if v434 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v445 = int32(1)
	v447 = v399 | base.B2i32(v432 != int64(0))
	goto L125
L130:
	;
	goto L131
L131:
	;
	v445 = int32(0)
	v447 = v399 & base.B2i32(v432 != int64(0))
	goto L125
L132:
	;
	goto L124
L133:
	;
	v493 = v471 & v461
	goto L113
L134:
	;
	goto L111
L135:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v524 == int32(0) {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v533 = int32(0)
	goto L137
L137:
	;
	v549 = v40 + v533*int32(24)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	switch v551 {
	case 0:
		goto L141
	case 1:
		goto L140
	default:
		v560 = int32(0)
		goto L139
	}
L138:
	;
	goto L13
L139:
	;
	v561 = v533 + v26
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	if l4 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v549)+16))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v522))))
	v560 = v557 ^ int32(1)
	goto L139
L141:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)+16))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552+v522))))
	v560 = v554
	goto L139
L142:
	;
	v565 = v560 | v562
	goto L144
L143:
	;
	v565 = v560 & v562
	goto L144
L144:
	;
	v566 = int32(1)
	v567 = v565 & v566
	*(*uint8)(unsafe.Add(mBase, uint32(v561))) = uint8(v567)
	v570 = v533 + v566
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v570) < base.Ui32(v571) {
		v533 = v570
		goto L137
	} else {
		goto L145
	}
L145:
	;
	goto L138
L146:
	;
	v782 = int32(0)
	v784 = F_mcv_match_expression(m, v242, l1, l2, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L199
	}
L147:
	;
	v726 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+8)))
	v727 = F_bms_member_index(m, l1, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L188
	}
L148:
	;
	v576 = v243 - int32(6)
	if v576 == int32(0) {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if v243 != int32(6) {
		goto L146
	} else {
		goto L187
	}
L151:
	;
	if v576 != int32(15) {
		goto L146
	} else {
		goto L152
	}
L152:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if base.Ui32(v581) <= base.Ui32(int32(1)) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	v587 = F_mcv_get_match_bitmap(m, v584, l1, l2, l3, base.B2i32(v581 == int32(1)))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if v581 != int32(2) {
		goto L146
	} else {
		goto L171
	}
L156:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v590 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v596 = int32(0)
	goto L160
L158:
	;
	goto L159
L159:
	;
	F_pfree(m, v587)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L170
	}
L160:
	;
	v610 = v596 + v26
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if l4 != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	goto L159
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v610))) = uint8(v625)
	v628 = v596 + int32(1)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v628) < base.Ui32(v629) {
		v596 = v628
		goto L160
	} else {
		goto L169
	}
L163:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596+v587))))
	v625 = v624
	goto L162
L164:
	;
	v612 = int32(1)
	if v611&v612 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v617 = int32(0)
	if v611&int32(1) == v617 {
		v625 = v617
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v625 = v612
	goto L162
L168:
	;
	goto L163
L169:
	;
	goto L161
L170:
	;
	goto L13
L171:
	;
	v654 = int32(0)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	v657 = F_mcv_get_match_bitmap(m, v655, l1, l2, l3, v654)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v659 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v665 = v654
	goto L176
L174:
	;
	goto L175
L175:
	;
	F_pfree(m, v657)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L186
	}
L176:
	;
	v679 = v665 + v26
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if l4 != 0 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L175
L178:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v679))) = uint8(v696)
	v699 = v665 + int32(1)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v699) < base.Ui32(v700) {
		v665 = v699
		goto L176
	} else {
		goto L185
	}
L179:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665+v657))))
	v696 = base.B2i32(v693 == int32(0))
	goto L178
L180:
	;
	v681 = int32(1)
	if v680&v681 == int32(0) {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v686 = int32(0)
	if v680&int32(1) == v686 {
		v696 = v686
		goto L178
	} else {
		goto L184
	}
L183:
	;
	v696 = v681
	goto L178
L184:
	;
	goto L179
L185:
	;
	goto L177
L186:
	;
	goto L13
L187:
	;
	goto L147
L188:
	;
	v729 = int32(0)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v730 == v729 {
		goto L13
	} else {
		goto L189
	}
L189:
	;
	v738 = v729
	goto L190
L190:
	;
	v752 = int32(0)
	v755 = v40 + v738*int32(24)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+16))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756+v727))))
	if v758 == v752 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L13
L192:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v755)+20))
	v765 = *(*int64)(unsafe.Add(mBase, uint32(v761+v727<<(uint(int32(3))%32))))
	v768 = base.B2i32(v765 != int64(0))
	goto L194
L193:
	;
	v768 = v752
	goto L194
L194:
	;
	v769 = v738 + v26
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769))))
	if l4 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v773 = v768 | v770
	goto L197
L196:
	;
	v773 = v768 & v770
	goto L197
L197:
	;
	v774 = int32(1)
	v775 = v773 & v774
	*(*uint8)(unsafe.Add(mBase, uint32(v769))) = uint8(v775)
	v778 = v738 + v774
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v778) < base.Ui32(v779) {
		v738 = v778
		goto L190
	} else {
		goto L198
	}
L198:
	;
	goto L191
L199:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v786 == int32(0) {
		goto L13
	} else {
		goto L200
	}
L200:
	;
	v794 = v782
	goto L201
L201:
	;
	v808 = int32(0)
	v811 = v40 + v794*int32(24)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)+16))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812+v784))))
	if v814 == v808 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	goto L13
L203:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v811)+20))
	v821 = *(*int64)(unsafe.Add(mBase, uint32(v817+v784<<(uint(int32(3))%32))))
	v824 = base.B2i32(v821 != int64(0))
	goto L205
L204:
	;
	v824 = v808
	goto L205
L205:
	;
	v825 = v794 + v26
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825))))
	if l4 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v829 = v824 | v826
	goto L208
L207:
	;
	v829 = v824 & v826
	goto L208
L208:
	;
	v830 = int32(1)
	v831 = v829 & v830
	*(*uint8)(unsafe.Add(mBase, uint32(v825))) = uint8(v831)
	v834 = v794 + v830
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v834) < base.Ui32(v835) {
		v794 = v834
		goto L201
	} else {
		goto L209
	}
L209:
	;
	goto L202
L210:
	;
	goto L12
L211:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_get_match_bitmap_0), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1733), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_get_match_bitmap_0), int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1737), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mcv_match_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 == int32(6) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v120
L2:
	;
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v8
	goto L7
L6:
	;
	goto L7
L7:
	;
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v11 = F_bms_member_index(m, l1, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if int32(0) <= v11 {
		v120 = v11
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_match_expression_0), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_mcv_match_expression_1), int32(1548), int32(_a_F_mcv_match_expression_2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v30 = F_exprCollation(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v34 = int64(0)
	if l1 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30
	goto L16
L18:
	;
	if l2 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v78 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v39 = l1 + int32(8)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v78 = base.I32_popcnt(v43)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v46 = v40 << (uint(int32(2)) % 32)
	if v46 <= int32(7) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v78 = base.I32_wrap_i64(v73)
	goto L18
L26:
	;
	if v46 == int32(0) {
		v73 = v34
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v70 = F_pg_popcount_optimized(m, v39, v46)
	mBase = m.M
	v73 = v70
	goto L25
L29:
	;
	v51 = v46
	v52 = v39
	v53 = v34
	goto L30
L30:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+3)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_mcv_match_expression[0]))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+2)))
	v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_mcv_match_expression[0]))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v59 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_mcv_match_expression[0]))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v61 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_mcv_match_expression[0]))))
	v65 = v55 + (v57 + (v59 + (v53 + v61)))
	v66 = int32(4)
	v69 = v51 - v66
	if v69 != 0 {
		v51 = v69
		v52 = v52 + v66
		v53 = v65
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v73 = v65
	goto L25
L32:
	;
	goto L31
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L41
	}
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v81 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v86 = v78
	v88 = int32(0)
	goto L36
L36:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v88<<(uint(int32(2))%32))))
	v94 = F_equal(m, l0, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	if v94 != 0 {
		v120 = v86
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v96 = int32(1)
	v99 = v88 + v96
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v99 < v100 {
		v86 = v86 + v96
		v88 = v99
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_match_expression_3), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_mcv_match_expression_1), int32(1571), int32(_a_F_mcv_match_expression_2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_md5_password_warning_enabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_md5_password_warning_enabled[0])))
	return v2
}
func F_mda_get_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v5 = int32(0)
	if l0 <= v5 {
	} else {
		if l0 != int32(1) {
			v21 = v5
			v24 = v5
			for {
				v25 = int32(2)
				v26 = v21 << (uint(v25) % 32)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26+l3)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26+l2)))
				v33 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1+v26))) = v29 - v31 + v33
				v37 = v26 | int32(4)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+l3)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v37+l2)))
				*(*int32)(unsafe.Add(mBase, uint32(l1+v37))) = v40 - v42 + v33
				v48 = v21 + v25
				v50 = v24 + v25
				if v50 != l0&int32(2147483646) {
					v21 = v48
					v24 = v50
					continue
				} else {
					break
				}
				break
			}
			if l0&int32(1) == int32(0) {
			} else {
				v58 = v48
				v63 = v58 << (uint(int32(2)) % 32)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+l3)))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v63+l2)))
				*(*int32)(unsafe.Add(mBase, uint32(l1+v63))) = v66 - v68 + int32(1)
			}
		} else {
			v58 = v5
			v63 = v58 << (uint(int32(2)) % 32)
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+l3)))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v63+l2)))
			*(*int32)(unsafe.Add(mBase, uint32(l1+v63))) = v66 - v68 + int32(1)
		}
	}
	return
}
func F_mdclose(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v7 = l0 + l1<<(uint(int32(2))%32)
	v9 = v7 + int32(40)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if int32(0) < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v7 + int32(56)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v17 = v10 - int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+v17<<(uint(int32(3))%32))))
	F_FileClose(m, v21)
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
	return
L4:
	;
	return
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = v24
	v26 = v17
	goto L9
L7:
	;
	v56 = v24
	v57 = v17
	goto L8
L8:
	;
	if int32(0) < v56 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v56 = v55
	v57 = v48
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v26
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v48 = v26 - int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+v48<<(uint(int32(3))%32))))
	F_FileClose(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L19
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v43
	goto L11
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_mdclose[0]))
	v35 = F_MemoryContextAlloc(m, v32, v26<<(uint(int32(3))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v26 <= v25 {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v43 = v35
	goto L12
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v41 = F_repalloc(m, v38, v26<<(uint(int32(3))%32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v43 = v41
	goto L12
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v48 != 0 {
		v25 = v55
		v26 = v48
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_pfree(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
	goto L3
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
	goto L23
}
func F_mdexists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mdexists[0])))
	if v4 == int32(0) {
		F_mdclose(m, l0, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = F_mdopenfork(m, l0, l1, int32(2))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v12 != int32(0))
			}
		}
	} else {
		v12 = F_mdopenfork(m, l0, l1, int32(2))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v12 != int32(0))
		}
	}
}
func F_member_can_set_role(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14308(m, l0, l1, int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_mmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int64
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v537 int32
	_ = v537
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	if l1&int32(32) == int32(0) {
		if base.Ui32(int32(2147483647)) <= base.Ui32(l0) {
			*(*int32)(unsafe.Add(mBase, _c_F_mmap[0])) = int32(48)
			v365 = int32(-1)
		} else {
			if l1&int32(16) != 0 {
				v21 = int32(-63)
			} else {
				v21 = int32(-48)
			}
			if l1&int32(32) != 0 {
				v24 = int32(_a_F_mmap_0)
				v28 = (l0 + int32(15)) & int32(-16)
				v30 = v28 + int32(40)
				if base.Ui32(int32(-65600)) <= base.Ui32(v30) {
					*(*int32)(unsafe.Add(mBase, _c_F_mmap[0])) = int32(48)
					v194 = int32(0)
				} else {
					if base.Ui32(v30) < base.Ui32(int32(11)) {
						v81 = int32(16)
					} else {
						v81 = (v28 + int32(51)) & int32(-8)
					}
					v85 = F_emscripten_builtin_malloc(m, v81+int32(_a_F_mmap_1))
					mBase = m.M
					if v85 == int32(0) {
						v194 = int32(0)
					} else {
						v89 = v85 - int32(8)
						if int32(_a_F_mmap_2)&v85 == int32(0) {
							v149 = v89
						} else {
							v96 = v85 - int32(4)
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							v107 = (v24+v85-int32(1))&int32(-65536) - int32(8)
							if base.Ui32(v107-v89) <= base.Ui32(int32(15)) {
								v112 = v24
							} else {
								v112 = int32(0)
							}
							v113 = v107 + v112
							v114 = v113 - v89
							v115 = v97&int32(-8) - v114
							if v97&int32(3) == int32(0) {
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v115
								*(*int32)(unsafe.Add(mBase, uint32(v113))) = v120 + v114
								v149 = v113
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
								v125 = int32(1)
								v128 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v115 | v124&v125 | v128
								v131 = v113 + v115
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v132 | v125
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v114 | v136&v125 | v128
								v143 = v89 + v114
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v144 | v125
								F_dispose_chunk(m, v89, v114)
								mBase = m.M
								v149 = v113
							}
						}
						v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
						if v155&int32(3) == int32(0) {
						} else {
							v161 = v155 & int32(-8)
							if base.Ui32(v161) <= base.Ui32(v81+int32(16)) {
							} else {
								v165 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v81 | v155&v165 | int32(2)
								v171 = v149 + v81
								v172 = v161 - v81
								*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v172 | int32(3)
								v176 = v149 + v161
								v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v177 | v165
								F_dispose_chunk(m, v171, v172)
								mBase = m.M
							}
						}
						v194 = v149 + int32(8)
					}
				}
				if v194 != 0 {
					v215 = int32(0)
					if v28 == v215 {
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v215)
						v222 = v194 + v28
						*(*uint8)(unsafe.Add(mBase, uint32(v222-int32(1)))) = uint8(v215)
						if base.Ui32(v28) < base.Ui32(int32(3)) {
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v194)+2)) = uint8(v215)
							*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)) = uint8(v215)
							*(*uint8)(unsafe.Add(mBase, uint32(v222-int32(3)))) = uint8(v215)
							*(*uint8)(unsafe.Add(mBase, uint32(v222-int32(2)))) = uint8(v215)
							if base.Ui32(v28) < base.Ui32(int32(7)) {
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v194)+3)) = uint8(v215)
								*(*uint8)(unsafe.Add(mBase, uint32(v222-int32(4)))) = uint8(v215)
								if base.Ui32(v28) < base.Ui32(int32(9)) {
								} else {
									v244 = int32(0)
									v247 = (v244 - v194) & int32(3)
									v248 = v194 + v247
									*(*int32)(unsafe.Add(mBase, uint32(v248))) = v244
									v256 = (v28 - v247) & int32(-4)
									v257 = v248 + v256
									*(*int32)(unsafe.Add(mBase, uint32(v257-int32(4)))) = v244
									if base.Ui32(v256) < base.Ui32(int32(9)) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v244
										*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = v244
										*(*int32)(unsafe.Add(mBase, uint32(v257-int32(8)))) = v244
										*(*int32)(unsafe.Add(mBase, uint32(v257-int32(12)))) = v244
										if base.Ui32(v256) < base.Ui32(int32(25)) {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v248)+24)) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v248)+20)) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v257-int32(16)))) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v257-int32(20)))) = v244
											v283 = int32(24)
											*(*int32)(unsafe.Add(mBase, uint32(v257-v283))) = v244
											*(*int32)(unsafe.Add(mBase, uint32(v257-int32(28)))) = v244
											v292 = v248&int32(4) | v283
											v293 = v256 - v292
											if base.Ui32(v293) < base.Ui32(int32(32)) {
											} else {
												v298 = base.I64_extend_i32_u(v244) * int64(4294967297)
												v301 = v292 + v248
												v302 = v293
												for {
													*(*int64)(unsafe.Add(mBase, uint32(v301)+24)) = v298
													*(*int64)(unsafe.Add(mBase, uint32(v301)+16)) = v298
													*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v298
													*(*int64)(unsafe.Add(mBase, uint32(v301))) = v298
													v310 = int32(32)
													v313 = v302 - v310
													if base.Ui32(int32(31)) < base.Ui32(v313) {
														v301 = v301 + v310
														v302 = v313
														continue
													} else {
														break
													}
													break
												}
											}
										}
									}
								}
							}
						}
					}
					v322 = v194 + v28
					*(*int32)(unsafe.Add(mBase, uint32(v322))) = v194
					*(*int64)(unsafe.Add(mBase, uint32(v322)+8)) = int64(-4294967295)
					v327 = v322
					*(*int32)(unsafe.Add(mBase, uint32(v327)+32)) = int32(3)
					*(*int64)(unsafe.Add(mBase, uint32(v327)+24)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = l0
					v335 = int32(_a_F_mmap_3)
					v336 = *(*int32)(unsafe.Add(mBase, _c_F_mmap[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v327)+36)) = v336
					*(*int32)(unsafe.Add(mBase, _c_F_mmap[1])) = v327
					v340 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
					v344 = v340
				} else {
					v344 = int32(-48)
				}
			} else {
				v207 = F_emscripten_builtin_malloc(m, int32(40))
				mBase = m.M
				v210 = m.Env.X_mmap_js(m, l0, int32(3), l1, l2, int64(0), v207+int32(8), v207)
				mBase = m.M
				if int32(0) <= v210 {
					*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = l2
					v327 = v207
					*(*int32)(unsafe.Add(mBase, uint32(v327)+32)) = int32(3)
					*(*int64)(unsafe.Add(mBase, uint32(v327)+24)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = l0
					v335 = int32(_a_F_mmap_3)
					v336 = *(*int32)(unsafe.Add(mBase, _c_F_mmap[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v327)+36)) = v336
					*(*int32)(unsafe.Add(mBase, _c_F_mmap[1])) = v327
					v340 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
					v344 = v340
				} else {
					F_emscripten_builtin_free(m, v207)
					mBase = m.M
					v344 = v210
				}
			}
			if l1&int32(32) != 0 {
				v348 = v21
			} else {
				v348 = int32(-63)
			}
			if v344 != int32(-63) {
				v351 = v344
			} else {
				v351 = v348
			}
			if base.Ui32(int32(-4095)) <= base.Ui32(v351) {
				*(*int32)(unsafe.Add(mBase, _c_F_mmap[0])) = int32(0) - v351
				v359 = int32(-1)
			} else {
				v359 = v351
			}
			v365 = v359
		}
		return v365
	} else {
		v368 = F_sbrk(m, int32(0))
		mBase = m.M
		v369 = int32(_a_F_mmap_0)
		v373 = (l0 + int32(_a_F_mmap_2)) & int32(-65536)
		if base.Ui32(int32(-65600)) <= base.Ui32(v373) {
			*(*int32)(unsafe.Add(mBase, _c_F_mmap[0])) = int32(48)
			v537 = int32(0)
		} else {
			v418 = int32(11)
			if base.Ui32(v373) < base.Ui32(v418) {
				v424 = int32(16)
			} else {
				v424 = (v373 + v418) & int32(-8)
			}
			v428 = F_emscripten_builtin_malloc(m, v424+int32(_a_F_mmap_1))
			mBase = m.M
			if v428 == int32(0) {
				v537 = int32(0)
			} else {
				v432 = v428 - int32(8)
				if int32(_a_F_mmap_2)&v428 == int32(0) {
					v492 = v432
				} else {
					v439 = v428 - int32(4)
					v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
					v450 = (v369+v428-int32(1))&int32(-65536) - int32(8)
					if base.Ui32(v450-v432) <= base.Ui32(int32(15)) {
						v455 = v369
					} else {
						v455 = int32(0)
					}
					v456 = v450 + v455
					v457 = v456 - v432
					v458 = v440&int32(-8) - v457
					if v440&int32(3) == int32(0) {
						v463 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
						*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v458
						*(*int32)(unsafe.Add(mBase, uint32(v456))) = v463 + v457
						v492 = v456
					} else {
						v467 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
						v468 = int32(1)
						v471 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v458 | v467&v468 | v471
						v474 = v456 + v458
						v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v474)+4)) = v475 | v468
						v479 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
						*(*int32)(unsafe.Add(mBase, uint32(v439))) = v457 | v479&v468 | v471
						v486 = v432 + v457
						v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v486)+4)) = v487 | v468
						F_dispose_chunk(m, v432, v457)
						mBase = m.M
						v492 = v456
					}
				}
				v498 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
				if v498&int32(3) == int32(0) {
				} else {
					v504 = v498 & int32(-8)
					if base.Ui32(v504) <= base.Ui32(v424+int32(16)) {
					} else {
						v508 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v492)+4)) = v424 | v498&v508 | int32(2)
						v514 = v492 + v424
						v515 = v504 - v424
						*(*int32)(unsafe.Add(mBase, uint32(v514)+4)) = v515 | int32(3)
						v519 = v492 + v504
						v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v520 | v508
						F_dispose_chunk(m, v514, v515)
						mBase = m.M
					}
				}
				v537 = v492 + int32(8)
			}
		}
		if v537 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_mmap[0])) = int32(48)
			return int32(-1)
		} else {
			if base.Ui32(v368) <= base.Ui32(v537) {
			} else {
				v554 = v368 - v537
				if base.Ui32(v554) < base.Ui32(v373) {
					v556 = v554
				} else {
					v556 = v373
				}
				if v556 == int32(0) {
				} else {
					base.MemoryFill(m, v537, int32(0), v556)
				}
			}
			return v537
		}
	}
}
func F_movedb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v355 int32
	_ = v355
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v522 int32
	_ = v522
	var v534 int64
	_ = v534
	var v535 int32
	_ = v535
	var v547 int32
	_ = v547
	var v559 int32
	_ = v559
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v635 int32
	_ = v635
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v680 int32
	_ = v680
	var v696 int32
	_ = v696
	var v711 int32
	_ = v711
	var v725 int32
	_ = v725
	var v738 int32
	_ = v738
	var v755 int32
	_ = v755
	var v769 int32
	_ = v769
	var v784 int32
	_ = v784
	var v796 int32
	_ = v796
	var v809 int64
	_ = v809
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
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
	var v846 int32
	_ = v846
	var v849 int64
	_ = v849
	var v859 int32
	_ = v859
	var v864 int64
	_ = v864
	var v886 int32
	_ = v886
	var v902 int32
	_ = v902
	var v917 int32
	_ = v917
	var v930 int64
	_ = v930
	var v931 int32
	_ = v931
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v993 int32
	_ = v993
	var v1006 int32
	_ = v1006
	var v1020 int32
	_ = v1020
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1085 int32
	_ = v1085
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1128 int32
	_ = v1128
	var v1141 int32
	_ = v1141
	var v1153 int32
	_ = v1153
	var v1167 int32
	_ = v1167
	var v1180 int32
	_ = v1180
	var v1196 int32
	_ = v1196
	var v1208 int32
	_ = v1208
	var v1220 int32
	_ = v1220
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1264 int32
	_ = v1264
	var v1279 int32
	_ = v1279
	var v1294 int32
	_ = v1294
	var v1309 int32
	_ = v1309
	var v1324 int32
	_ = v1324
	var v1337 int64
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1350 int32
	_ = v1350
	var v1362 int32
	_ = v1362
	var v1374 int32
	_ = v1374
	var v1391 int32
	_ = v1391
	var v1403 int32
	_ = v1403
	var v1415 int32
	_ = v1415
	var v1434 int32
	_ = v1434
	var v1435 int64
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int64
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(608)
	m.G0 = v21
	v27 = v3
	v28 = v3
	v29 = v3
	v30 = v3
	v31 = v3
	v32 = v3
	v33 = v3
	v34 = v3
	v35 = v3
	v36 = int32(-1)
	v37 = v3
	v40 = int64(0)
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v21 + int32(608)
	return
L4:
	;
	goto L3
L5:
	;
	if v36 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1434 = int32(m.ExcTag)
	v1435 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1434 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v30
	v56 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v836 = v27
	v837 = v28
	v838 = v29
	v839 = v30
	v840 = v31
	v841 = v32
	v842 = v33
	v843 = v34
	v844 = v35
	v846 = v37
	v849 = v40
	goto L10
L10:
	;
	if v846 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v71 = int32(0)
	v86 = F_get_db_info(m, l0, int32(8), v21+int32(556), v71, v71, v71, v71, v71, v71, v71, v21+int32(544), v71, v71, v71, v71, v71, v71)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v86 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v21)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_LockSharedObjectForSession(m, v156)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L20
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errcode(m, int32(1283))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = l0
	F_errmsg(m, int32(_a_F_movedb_0), v21+int32(80))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2052), int32(_a_F_movedb_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v184 = F_object_ownercheck(m, int32(1262), v156, v172)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v184 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_aclcheck_error(m, int32(2), int32(9), l0)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[1]))
	if v203 == v156 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v272 = F_get_tablespace_oid(m, l1, int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L33
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errcode(m, int32(100663621))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errmsg(m, int32(_a_F_movedb_3), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2076), int32(_a_F_movedb_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v298 = F_object_aclcheck(m, int32(1213), v272, v285, int64(512))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	if v298 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_aclcheck_error(m, v298, int32(43), l1)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v272 == int32(1664) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v21)+544))
	if v272 == v371 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errcode(m, int32(50856066))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errmsg(m, int32(_a_F_movedb_4), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2098), int32(_a_F_movedb_2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_relation_close(m, v56, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v412 = F_CountOtherDBBackends(m, v156, v21+int32(552), v21+int32(548))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_UnlockSharedObjectForSession(m, v156)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	goto L4
L51:
	;
	if v412 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v496 = F_GetDatabasePath(m, v156, v371)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L60
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errcode(m, int32(100663621))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l0
	F_errmsg(m, int32(_a_F_movedb_5), v21+int32(32))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v21)+552))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v21)+548))
	F_errdetail_busy_db(m, v467, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2122), int32(_a_F_movedb_2))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L1
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v508 = F_GetDatabasePath(m, v156, v272)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_RequestCheckpoint(m, int32(60))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v534 = F_EmitProcSignalBarrier(m, int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_WaitForProcSignalBarrier(m, v534)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_DropDatabaseBuffers(m, v156)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v570 = F_AllocateDir(m, v508)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+484)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+480)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	v809 = base.I64_extend_i32_u(v21 + int32(480))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_before_shmem_exit(m, int32(586), v809)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L7
	} else {
		goto L94
	}
L67:
	;
	if v570 == int32(0) {
		v796 = v35
		goto L66
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v584 = F_ReadDir(m, v570, v508)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L7
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L7
	} else {
		goto L89
	}
L70:
	;
	if v584 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v597 = v35
	v598 = v584
	goto L74
L72:
	;
	v635 = v35
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_FreeDir(m, v570)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L7
	} else {
		goto L84
	}
L74:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+19)))
	if v604 != int32(46) {
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v635 = v622
	goto L73
L76:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+20)))
	if v607 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+20)))
	if v608 != int32(46) {
		goto L69
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v622 = F_ReadDir(m, v570, v508)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+21)))
	if v611 != 0 {
		goto L69
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if v622 != 0 {
		v597 = v622
		v598 = v622
		goto L74
	} else {
		goto L83
	}
L83:
	;
	goto L75
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	v664 = F_rmdir(m, v508)
	mBase = m.M
	if v664 == int32(0) {
		v796 = v635
		goto L66
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v508
	F_errmsg_internal(m, int32(_a_F_movedb_6), v21+int32(48))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2195), int32(_a_F_movedb_2))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errcode(m, int32(325))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l0
	F_errmsg(m, int32(_a_F_movedb_7), v21-int32(-64))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errhint(m, int32(_a_F_movedb_8), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2184), int32(_a_F_movedb_2))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	goto L1
L94:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[2]))
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[3]))
	goto L95
L95:
	;
	v827 = v21 + int32(320)
	*(*int32)(unsafe.Add(mBase, uint32(v827)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v827))) = v21 + int32(84)
	goto L98
L96:
	;
	v836 = v156
	v837 = v508
	v838 = v272
	v839 = v56
	v840 = v496
	v841 = v823
	v842 = v825
	v843 = v371
	v844 = v796
	v846 = int32(0)
	v849 = v809
	goto L10
L98:
	;
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v21 + int32(320)
	v859 = int32(0)
	base.MemoryFill(m, v21+int32(176), v859, int32(144))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+160)) = uint16(v859)
	v864 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+152)) = v864
	*(*int64)(unsafe.Add(mBase, uint32(v21)+144)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+128)) = uint16(v859)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v864
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v864
	F_copydir(m, v840, v837, v859)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[2])) = v841
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_cancel_before_shmem_exit(m, int32(586), v849)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L7
	} else {
		goto L146
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_XLogBeginInsert(m)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_XLogRegisterData(m, v21+int32(96), int32(16))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v930 = F_XLogInsert(m, int32(4), int32(1))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v943 = v21 + int32(488)
	F_ScanKeyInit(m, v943, int32(2), int32(3), int32(62), base.I64_extend_i32_u(l0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v961 = int32(1)
	v964 = F_systable_beginscan(m, v839, int32(2671), v961, int32(0), v961, v943)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v976 = F_systable_getnext(m, v964)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	if v976 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1047 = v976 + int32(4)
	F_LockTuple(m, v839, v1047, int32(7))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L7
	} else {
		goto L116
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_errcode(m, int32(1283))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l0
	F_errmsg(m, int32(_a_F_movedb_0), v21)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_errfinish(m, int32(_a_F_movedb_1), int32(2250), int32(_a_F_movedb_2))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	goto L1
L116:
	;
	v1051 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+123)) = uint8(v1051)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+264)) = base.I64_extend_i32_u(v838)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v839)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1072 = F_heap_modify_tuple(m, v976, v1055, v21+int32(176), v21+int32(144), v21+int32(112))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_CatalogTupleUpdate(m, v839, v1047, v1072)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_UnlockTuple(m, v839, v1047, int32(7))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[4]))
	if v1100 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1112 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1262), v836, v1112, v1112, v1112)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L7
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_systable_endscan(m, v964)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L7
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_RequestCheckpoint(m, int32(44))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_movedb[5])) = uint8(v1153)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_relation_close(m, v839, int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_cancel_before_shmem_exit(m, int32(586), v849)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[2])) = v841
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_StartTransactionCommand(m)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1231 = F_rmtree(m, v840)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L7
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_XLogBeginInsert(m)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L7
	} else {
		goto L139
	}
L133:
	;
	if v1231 != 0 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1245 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	if v1245 == int32(0) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v840
	F_errmsg(m, int32(_a_F_movedb_9), v21+int32(16))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_errfinish(m, int32(_a_F_movedb_1), int32(2314), int32(_a_F_movedb_2))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_XLogRegisterData(m, v21+int32(88), int32(8))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_XLogRegisterData(m, v21+int32(544), int32(4))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	v1337 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_UnlockSharedObjectForSession(m, v836)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_pfree(m, v840)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_pfree(m, v837)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	goto L4
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_movedb_failure_callback(m, v21, v849)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+564)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+560)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v21)+568)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v21)+580)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+584)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+588)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+592)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+596)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+600)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+604)) = v839
	F_pg_re_throw(m)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L7
	} else {
		goto L148
	}
L148:
	;
	goto L1
L149:
	;
	v1439 = int32(v1435)
	m.G0 = v21
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+4))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1439)))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	if v21+int32(84) == v1445 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	m.ExcPending = 1
	goto L158
L151:
	;
	if v1449 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	v1449 = v1447
	goto L154
L153:
	;
	v1449 = int32(0)
	goto L154
L154:
	;
	goto L151
L155:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v21)+604))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v21)+600))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v21)+596))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v21)+592))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v21)+588))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v21)+584))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v21)+580))
	v1457 = *(*int64)(unsafe.Add(mBase, uint32(v21)+568))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v21)+564))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v21)+560))
	v27 = v1451
	v28 = v1455
	v29 = v1452
	v30 = v1450
	v31 = v1454
	v32 = v1459
	v33 = v1458
	v34 = v1453
	v35 = v1456
	v36 = v1449
	v37 = v1441
	v40 = v1457
	goto L2
L156:
	;
	goto L157
L157:
	;
	F___wasm_longjmp(m, v1442, v1441)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	return
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mxactMemberComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v7) < base.Ui32(v6) {
		v21 = int32(1)
	} else {
		if base.Ui32(v6) < base.Ui32(v7) {
			v21 = int32(-1)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v13) < base.Ui32(v12) {
				v21 = int32(1)
			} else {
				if base.Ui32(v12) < base.Ui32(v13) {
					v18 = int32(-1)
				} else {
					v18 = int32(0)
				}
				v21 = v18
			}
		}
	}
	return v21
}
