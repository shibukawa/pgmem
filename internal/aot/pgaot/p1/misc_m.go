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
func F_makeMdArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_makeMdArrayResult[0])) = v9
		if l5 != 0 {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_MemoryContextDelete(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v18
			}
		} else {
			return v18
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
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
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
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
		v238 = v58
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
	return v238
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
		v238 = v58
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
	v90 = l1 << (uint(int32(2)) % 32)
	v92 = v90 + int32(584)
	v94 = v92 & int32(4092)
	if v94 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v105 = v87
	v106 = v71
	v107 = v83
	goto L27
L27:
	;
	v108 = int32(_a_F_make_new_segment_2)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0])) = v111
	v114 = F_dsm_create(m, v106, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v98 = v90 - v94 + int32(_a_F_make_new_segment_1)
	goto L30
L29:
	;
	v98 = v92
	goto L30
L30:
	;
	v101 = v98 + l1<<(uint(int32(12))%32)
	if base.Ui32(int32(134217728)) < base.Ui32(v101) {
		v238 = v58
		goto L13
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v69) < base.Ui32(v101) {
		v238 = v58
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v105 = l1
	v106 = v101
	v107 = v98
	goto L27
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_new_segment[0])) = v109
	if v114 == int32(0) {
		v238 = v58
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_dsm_pin_segment(m, v114)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v124+v57<<(uint(int32(2))%32))+32)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+1456))
	if base.Ui32(v131) < base.Ui32(v57) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+1456)) = v57
	goto L39
L38:
	;
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v134) < base.Ui32(v57) {
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
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+1448))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+1448)) = v138 + v106
	v143 = l0 + v57*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v114
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = v145 + int32(584)
	v152 = v145 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v152
	v154 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v152)+4)) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v152)+12)) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v152)+20)) = v154
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v152)+28)) = v160
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+32)) = uint8(v162)
	if v152 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	F_FreePageManagerPut(m, v175, int32(base.Ui32(v107)>>(uint(int32(12))%32)), v105)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L33
	} else {
		goto L47
	}
L44:
	;
	v168 = v152 - v145 + v162
	goto L46
L45:
	;
	v168 = v160
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v168
	base.MemoryFill(m, v145+int32(68), int32(0), int32(516))
	goto L43
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v57 ^ v182 ^ int32(216163848)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v105
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v106
	v192 = v143 + int32(8)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	if v105 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v194 = int32(14)
	v197 = base.I32_clz(v105) ^ int32(31)
	if base.Ui32(v194) <= base.Ui32(v197) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v205 = int32(0)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+20)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v208 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v213 = int32(2)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211+v212<<(uint(v213)%32))+160))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+24)) = uint8(v219)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v221+v223<<(uint(v213)%32))+160)) = v57
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	if v229 == v208 {
		v238 = v192
		goto L13
	} else {
		goto L54
	}
L51:
	;
	v200 = v194
	goto L53
L52:
	;
	v200 = v197
	goto L53
L53:
	;
	v205 = v200 + int32(1)
	goto L50
L54:
	;
	v232 = F_get_segment_by_index(m, l0, v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L33
	} else {
		goto L55
	}
L55:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v57
	v238 = v192
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
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(277)
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
func F_make_scalar_key(m *base.Module, l0 int32, l1 int32) int32 {
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
			return int32(0)
		} else {
			v756 = int32(2)
			*(*uint8)(unsafe.Add(mBase, uint32(v754)+4)) = uint8(v756)
			*(*int32)(unsafe.Add(mBase, uint32(v754))) = int32(20)
			v760 = v754
			m.G0 = v10 + int32(48)
			return v760
		}
	case 1:
		if l1 != 0 {
			v442 = int32(1)
		} else {
			v442 = int32(5)
		}
		v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
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
				return int32(0)
			} else {
				v723 = int32(8)
				v724 = v442 | int32(16)
				v725 = v713
				v727 = v723 + int32(5)
				v728 = F_palloc(m, v727)
				mBase = m.M
				v729 = m.ExcPending
				if v729 != 0 {
					return int32(0)
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
					return v760
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
				return int32(0)
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
				return v760
			}
		}
	case 2:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
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
						return int32(0)
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
								return int32(0)
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
									return int32(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return v760
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
								return int32(0)
							} else {
								v424 = int32(8)
								v425 = int32(20)
								v426 = v414
								v428 = v424 + int32(5)
								v429 = F_palloc(m, v428)
								mBase = m.M
								v430 = m.ExcPending
								if v430 != 0 {
									return int32(0)
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
										return int32(0)
									} else {
										v760 = v429
										m.G0 = v10 + int32(48)
										return v760
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
						return int32(0)
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
								return int32(0)
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
									return int32(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return v760
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
								return int32(0)
							} else {
								v424 = int32(8)
								v425 = int32(20)
								v426 = v414
								v428 = v424 + int32(5)
								v429 = F_palloc(m, v428)
								mBase = m.M
								v430 = m.ExcPending
								if v430 != 0 {
									return int32(0)
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
										return int32(0)
									} else {
										v760 = v429
										m.G0 = v10 + int32(48)
										return v760
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
					return int32(0)
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
							return int32(0)
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
								return int32(0)
							} else {
								v760 = v429
								m.G0 = v10 + int32(48)
								return v760
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
							return int32(0)
						} else {
							v424 = int32(8)
							v425 = int32(20)
							v426 = v414
							v428 = v424 + int32(5)
							v429 = F_palloc(m, v428)
							mBase = m.M
							v430 = m.ExcPending
							if v430 != 0 {
								return int32(0)
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
									return int32(0)
								} else {
									v760 = v429
									m.G0 = v10 + int32(48)
									return v760
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
				return int32(0)
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
						return int32(0)
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
							return int32(0)
						} else {
							v760 = v429
							m.G0 = v10 + int32(48)
							return v760
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
						return int32(0)
					} else {
						v424 = int32(8)
						v425 = int32(20)
						v426 = v414
						v428 = v424 + int32(5)
						v429 = F_palloc(m, v428)
						mBase = m.M
						v430 = m.ExcPending
						if v430 != 0 {
							return int32(0)
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
								return int32(0)
							} else {
								v760 = v429
								m.G0 = v10 + int32(48)
								return v760
							}
						}
					}
				}
			}
		}
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		v15 = F_palloc(m, int32(6))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
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
			return v760
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v742 = m.ExcPending
		if v742 != 0 {
			return int32(0)
		} else {
			v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v743
			F_errmsg_internal(m, int32(_a_F_make_scalar_key_10), v10)
			mBase = m.M
			v747 = m.ExcPending
			if v747 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_make_scalar_key_11), int32(1403), int32(_a_F_make_scalar_key_12))
				mBase = m.M
				v752 = m.ExcPending
				if v752 != 0 {
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
	F_errfinish(m, int32(_a_F_manifest_report_error_1), int32(1043), int32(_a_F_manifest_report_error_2))
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
	F_errdetail(m, int32(_a_F_matchLocks_1), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_matchLocks_2), int32(1694), int32(_a_F_matchLocks_3))
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v19 != int32(7) {
		v168 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 - int32(-64)
	return v168
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v22 != 0 {
		v168 = v7
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
		v168 = v7
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
		v168 = v7
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
		v168 = v7
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
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if base.B2i32(v92 == int32(0))|base.B2i32(l3 != l5) != 0 {
		v168 = int32(0)
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v98 = F_make_opclause(m, v82, l0, v85, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v98
	v105 = F_list_make1_impl(m, int32(1), v15+int32(-56))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v168 = v105
	goto L1
L34:
	;
	if v84 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v109 = F_get_collation_isdeterministic(m, l3)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v109 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v168 = int32(0)
	goto L1
L38:
	;
	if v83 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v112 = F_op_in_opfamily(m, v80, l4)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v112 == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v116 = F_make_opclause(m, v80, l0, v85, l5)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v116
	v123 = F_list_make1_impl(m, int32(1), v15+int32(-48))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v168 = v123
	goto L1
L44:
	;
	v133 = F_op_in_opfamily(m, v81, l4)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L48
	}
L45:
	;
	v127 = F_pg_newlocale_from_collation(m, l5)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	if v129 == int32(1) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v168 = int32(0)
	goto L1
L48:
	;
	if v133 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v168 = int32(0)
	goto L1
L50:
	;
	goto L51
L51:
	;
	v138 = F_make_opclause(m, v81, l0, v85, l5)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v138
	v145 = F_list_make1_impl(m, int32(1), v15+int32(-52))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v147 = F_op_in_opfamily(m, v79, l4)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v147 == int32(0) {
		v168 = v145
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v151 = F_get_opcode(m, v79)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v154 = v15 + int32(-32)
	F_fmgr_info(m, v151, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v157 = F_make_greater_string(m, v85, v154, l5)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v157 == int32(0) {
		v168 = v145
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v161 = F_make_opclause(m, v79, l0, v157, l5)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v163 = F_lappend(m, v145, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v168 = v163
	goto L1
}
func F_maybe_adjust_io_workers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[0]))
	if v6 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[1]))
	if base.Ui32(int32(8)) < base.Ui32(v10) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[2]))
	if int32(2) < v14 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[3])))
	if base.Ui32(int32(4)) < base.Ui32(v10) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = v18
	goto L7
L6:
	;
	v22 = int32(0)
	goto L7
L7:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[4]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[5]))
	if v26 <= v24 {
		v70 = v24
		v72 = v26
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v70 <= v72 {
		goto L1
	} else {
		goto L24
	}
L10:
	;
	v29 = int32(0)
	goto L11
L11:
	;
	v33 = v29 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_maybe_adjust_io_workers[6])))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v70 = v67
	v72 = v56
	goto L9
L13:
	;
	v36 = v29 + int32(1)
	if v36 != int32(32) {
		v29 = v36
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v53 = F_StartChildProcess(m, int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	F_errmsg_internal(m, int32(_a_F_maybe_adjust_io_workers_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_maybe_adjust_io_workers_1), int32(_a_F_maybe_adjust_io_workers_2), int32(_a_F_maybe_adjust_io_workers_3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[5]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[4]))
	if v53 == int32(0) {
		v70 = v58
		v72 = v56
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_maybe_adjust_io_workers[6]))) = v53
	v67 = v58 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_maybe_adjust_io_workers[4])) = v67
	if v67 < v56 {
		v29 = int32(0)
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L12
L24:
	;
	v76 = int32(31)
	goto L25
L25:
	;
	v81 = v76 << (uint(int32(2)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_maybe_adjust_io_workers[6])))
	if v82 != 0 {
		v96 = v82
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v100 = F_pgmem_kill(m, v98, int32(12))
	mBase = m.M
	goto L1
L27:
	;
	goto L26
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_maybe_adjust_io_workers[7])))
	if v85 != 0 {
		v96 = v85
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_maybe_adjust_io_workers[8])))
	if v88 != 0 {
		v96 = v88
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v90 = v76 - int32(3)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_c_F_maybe_adjust_io_workers[6])))
	if v93 != 0 {
		v96 = v93
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v90 != 0 {
		v76 = v76 - int32(4)
		goto L25
	} else {
		goto L32
	}
L32:
	;
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v25 = F_palloc(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	base.MemoryFill(m, v25, l4^int32(1), v29)
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
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L214
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L211
	}
L8:
	;
	m.G0 = v22 - int32(-64)
	return v25
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v39 = l3 + int32(48)
	v54 = int32(0)
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v54<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 == int32(318) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L8
L13:
	;
	v854 = v54 + int32(1)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v854 < v855 {
		v54 = v854
		goto L11
	} else {
		goto L210
	}
L14:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v240 != int32(52) {
		goto L74
	} else {
		goto L75
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v67 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v73 = v63
	v74 = v64
	goto L17
L17:
	;
	if v74 != int32(17) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v238 = int32(1)
	v239 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v73 = v67
	v74 = v72
	goto L17
L21:
	;
	v238 = int32(0)
	v239 = v73
	goto L14
L22:
	;
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v79 = F_get_opcode(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_fmgr_info(m, v79, v20+int32(-28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v87 = v20 + int32(-32)
	v89 = v20 + int32(-36)
	v91 = v20 + int32(-52)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99 == int32(27) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v128 != 0 {
		goto L46
	} else {
		goto L47
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v103 = v102
	goto L29
L28:
	;
	v103 = v98
	goto L29
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v104 == int32(27) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = v107
	v110 = v108
	goto L32
L31:
	;
	v109 = v97
	v110 = v104
	goto L32
L32:
	;
	if v110 == int32(7) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L26
L34:
	;
	if v87 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v116 = v109
	v117 = v103
	goto L34
L36:
	;
	goto L37
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v113 != int32(7) {
		v128 = int32(0)
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v116 = v103
	v117 = v109
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v117
	goto L41
L40:
	;
	goto L41
L41:
	;
	if v89 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v116
	goto L44
L43:
	;
	goto L44
L44:
	;
	v120 = int32(1)
	if v91 == int32(0) {
		v128 = v120
		goto L33
	} else {
		goto L45
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(base.B2i32(v110 == int32(7)))
	v128 = v120
	goto L33
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v132 = F_mcv_match_expression(m, v129, l1, l2, v20+int32(-40))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L70
	}
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v134 == int32(0) {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v144 = int32(0)
	goto L51
L51:
	;
	v161 = v39 + v144*int32(24)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v132))))
	if v164 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	goto L13
L53:
	;
	v222 = v144 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v222) < base.Ui32(v223) {
		v144 = v222
		goto L51
	} else {
		goto L69
	}
L54:
	;
	v216 = v214 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v25))) = uint8(v216)
	goto L53
L55:
	;
	v177 = v144 + v25
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if l4 == v178 {
		goto L53
	} else {
		goto L61
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+24)))
	if v168 != int32(1) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v172 = int32(0)
	if l4 == v172 {
		v214 = v172
		goto L54
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v25))))
	v214 = v176
	goto L54
L61:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)))
	if v180 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if l4 != 0 {
		v214 = base.B2i32(v203 != int32(0)) | v206
		goto L54
	} else {
		goto L68
	}
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v132<<(uint(int32(2))%32))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v191 = F_FunctionCall2Coll(m, v20+int32(-28), v138, v189, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196+v132<<(uint(int32(2))%32))))
	v201 = F_FunctionCall2Coll(m, v20+int32(-28), v138, v195, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v203 = v191
	goto L62
L67:
	;
	v203 = v201
	goto L62
L68:
	;
	v214 = v206 & base.B2i32(v203 != int32(0))
	goto L54
L69:
	;
	goto L52
L70:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_get_match_bitmap_0), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1647), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
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
	if v238 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L74:
	;
	if v240 != int32(20) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v519 = F_mcv_match_expression(m, v517, l1, l2, int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L135
	}
L77:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v246 = F_get_opcode(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_fmgr_info(m, v246, v20+int32(-28))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v239)+28))
	v254 = v20 + int32(-32)
	v256 = v20 + int32(-36)
	v258 = v20 + int32(-41)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	if v266 == int32(27) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v295 == int32(0) {
		goto L7
	} else {
		goto L100
	}
L81:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v270 = v269
	goto L83
L82:
	;
	v270 = v265
	goto L83
L83:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v271 == int32(27) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v276 = v274
	v277 = v275
	goto L86
L85:
	;
	v276 = v264
	v277 = v271
	goto L86
L86:
	;
	if v277 == int32(7) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L80
L88:
	;
	if v254 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v283 = v276
	v284 = v270
	goto L88
L90:
	;
	goto L91
L91:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v280 != int32(7) {
		v295 = int32(0)
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v283 = v270
	v284 = v276
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v284
	goto L95
L94:
	;
	goto L95
L95:
	;
	if v256 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v283
	goto L98
L97:
	;
	goto L98
L98:
	;
	v287 = int32(1)
	if v258 == int32(0) {
		v295 = v287
		goto L87
	} else {
		goto L99
	}
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(base.B2i32(v277 == int32(7)))
	v295 = v287
	goto L87
L100:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v298 == int32(0) {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+24)))
	if v302 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v306 = F_pg_detoast_datum(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v333 = F_mcv_match_expression(m, v330, l1, l2, v20+int32(-40))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L108
	}
L105:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	F_get_typlenbyvalalign(m, v308, v20+int32(-44), v20+int32(-45), v20+int32(-46))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v318 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+20)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+19)))
	v320 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+18)))
	F_deconstruct_array(m, v306, v318, v319, v320, v20+int32(-56), v20+int32(-60), v20+int32(-52))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v335 == int32(0) {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v347 = int32(0)
	goto L110
L110:
	;
	v362 = v39 + v347*int32(24)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+16))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v333))))
	if v365 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	goto L13
L112:
	;
	v514 = v347 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v514) < base.Ui32(v515) {
		v347 = v514
		goto L110
	} else {
		goto L134
	}
L113:
	;
	v492 = v490 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v347+v25))) = uint8(v492)
	goto L112
L114:
	;
	v377 = v347 + v25
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if l4 == v378 {
		goto L112
	} else {
		goto L120
	}
L115:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+24)))
	if v369 != int32(1) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v372 = int32(0)
	if l4 == v372 {
		v490 = v372
		goto L113
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v25))))
	v490 = v376
	goto L113
L120:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+20)))
	v382 = v380 ^ int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v383 <= int32(0) {
		v458 = v382
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if l4 != 0 {
		v490 = v468 | v458
		goto L113
	} else {
		goto L133
	}
L122:
	;
	v392 = int32(0)
	v396 = v382
	v397 = v380
	goto L123
L123:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407+v392))))
	if v409 == int32(1) {
		v442 = v397
		v444 = v396 & v397
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v458 = v444
	goto L121
L125:
	;
	v446 = v392 + int32(1)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v446 < v447 {
		v392 = v446
		v396 = v444
		v397 = v442
		goto L123
	} else {
		goto L132
	}
L126:
	;
	v412 = int32(1)
	if v396&v412 == v397&v412 {
		v458 = v396
		goto L121
	} else {
		goto L127
	}
L127:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v362)+20))
	v420 = int32(2)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419+v333<<(uint(v420)%32))))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424+v392<<(uint(v420)%32))))
	v429 = F_FunctionCall2Coll(m, v20+int32(-28), v339, v423, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+20)))
	if v431 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v442 = int32(1)
	v444 = v396 | base.B2i32(v429 != int32(0))
	goto L125
L130:
	;
	goto L131
L131:
	;
	v438 = int32(0)
	v442 = v438
	v444 = v396 & base.B2i32(v429 != v438)
	goto L125
L132:
	;
	goto L124
L133:
	;
	v490 = v468 & v458
	goto L113
L134:
	;
	goto L111
L135:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v521 == int32(0) {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v530 = int32(0)
	goto L137
L137:
	;
	v546 = v39 + v530*int32(24)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	switch v548 {
	case 0:
		goto L141
	case 1:
		goto L140
	default:
		v557 = int32(0)
		goto L139
	}
L138:
	;
	goto L13
L139:
	;
	v558 = v530 + v25
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if l4 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552+v519))))
	v557 = v554 ^ int32(1)
	goto L139
L141:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+v519))))
	v557 = v551
	goto L139
L142:
	;
	v562 = v557 | v559
	goto L144
L143:
	;
	v562 = v557 & v559
	goto L144
L144:
	;
	v563 = int32(1)
	v564 = v562 & v563
	*(*uint8)(unsafe.Add(mBase, uint32(v558))) = uint8(v564)
	v567 = v530 + v563
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v567) < base.Ui32(v568) {
		v530 = v567
		goto L137
	} else {
		goto L145
	}
L145:
	;
	goto L138
L146:
	;
	v779 = int32(0)
	v781 = F_mcv_match_expression(m, v239, l1, l2, v779)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L199
	}
L147:
	;
	v723 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239)+8)))
	v724 = F_bms_member_index(m, l1, v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L188
	}
L148:
	;
	v573 = v240 - int32(6)
	if v573 == int32(0) {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if v240 != int32(6) {
		goto L146
	} else {
		goto L187
	}
L151:
	;
	if v573 != int32(15) {
		goto L146
	} else {
		goto L152
	}
L152:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if base.Ui32(v578) <= base.Ui32(int32(1)) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v584 = F_mcv_get_match_bitmap(m, v581, l1, l2, l3, base.B2i32(v578 == int32(1)))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if v578 != int32(2) {
		goto L146
	} else {
		goto L171
	}
L156:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v587 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v593 = int32(0)
	goto L160
L158:
	;
	goto L159
L159:
	;
	F_pfree(m, v584)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L170
	}
L160:
	;
	v607 = v593 + v25
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v607))) = uint8(v622)
	v625 = v593 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v625) < base.Ui32(v626) {
		v593 = v625
		goto L160
	} else {
		goto L169
	}
L163:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593+v584))))
	v622 = v621
	goto L162
L164:
	;
	v609 = int32(1)
	if v608&v609 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v614 = int32(0)
	if v608&int32(1) == v614 {
		v622 = v614
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v622 = v609
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
	v651 = int32(0)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v654 = F_mcv_get_match_bitmap(m, v652, l1, l2, l3, v651)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v656 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v662 = v651
	goto L176
L174:
	;
	goto L175
L175:
	;
	F_pfree(m, v654)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L186
	}
L176:
	;
	v676 = v662 + v25
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v693)
	v696 = v662 + int32(1)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v696) < base.Ui32(v697) {
		v662 = v696
		goto L176
	} else {
		goto L185
	}
L179:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662+v654))))
	v693 = base.B2i32(v690 == int32(0))
	goto L178
L180:
	;
	v678 = int32(1)
	if v677&v678 == int32(0) {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v683 = int32(0)
	if v677&int32(1) == v683 {
		v693 = v683
		goto L178
	} else {
		goto L184
	}
L183:
	;
	v693 = v678
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
	v726 = int32(0)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v727 == v726 {
		goto L13
	} else {
		goto L189
	}
L189:
	;
	v735 = v726
	goto L190
L190:
	;
	v749 = int32(0)
	v752 = v39 + v735*int32(24)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+16))
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v724))))
	if v755 == v749 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L13
L192:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v752)+20))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v758+v724<<(uint(int32(2))%32))))
	v765 = base.B2i32(v762 != int32(0))
	goto L194
L193:
	;
	v765 = v749
	goto L194
L194:
	;
	v766 = v735 + v25
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	if l4 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v770 = v765 | v767
	goto L197
L196:
	;
	v770 = v765 & v767
	goto L197
L197:
	;
	v771 = int32(1)
	v772 = v770 & v771
	*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v772)
	v775 = v735 + v771
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v775) < base.Ui32(v776) {
		v735 = v775
		goto L190
	} else {
		goto L198
	}
L198:
	;
	goto L191
L199:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v783 == int32(0) {
		goto L13
	} else {
		goto L200
	}
L200:
	;
	v791 = v779
	goto L201
L201:
	;
	v805 = int32(0)
	v808 = v39 + v791*int32(24)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+16))
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809+v781))))
	if v811 == v805 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	goto L13
L203:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v808)+20))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814+v781<<(uint(int32(2))%32))))
	v821 = base.B2i32(v818 != int32(0))
	goto L205
L204:
	;
	v821 = v805
	goto L205
L205:
	;
	v822 = v791 + v25
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822))))
	if l4 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v826 = v821 | v823
	goto L208
L207:
	;
	v826 = v821 & v823
	goto L208
L208:
	;
	v827 = int32(1)
	v828 = v826 & v827
	*(*uint8)(unsafe.Add(mBase, uint32(v822))) = uint8(v828)
	v831 = v791 + v827
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v831) < base.Ui32(v832) {
		v791 = v831
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
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1735), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
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
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_mcv_get_match_bitmap_1), int32(1739), int32(_a_F_mcv_get_match_bitmap_2))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
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
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 == int32(6) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v110
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
		v110 = v11
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
	F_errfinish(m, int32(_a_F_mcv_match_expression_1), int32(1550), int32(_a_F_mcv_match_expression_2))
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
	v33 = int32(0)
	if l1 == v33 {
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
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v68 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v40 = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 <= v40 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v44 = v40
	goto L24
L23:
	;
	v44 = v41
	goto L24
L24:
	;
	v48 = int32(0)
	v50 = v33
	goto L25
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v48<<(uint(int32(2))%32))))
	if v56 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v68 = v59
	goto L18
L27:
	;
	v59 = v50 + base.I32_popcnt(v56)
	goto L29
L28:
	;
	v59 = v50
	goto L29
L29:
	;
	v61 = v48 + int32(1)
	if v61 != v44 {
		v48 = v61
		v50 = v59
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L39
	}
L32:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v71 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v76 = v68
	v78 = int32(0)
	goto L34
L34:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v78<<(uint(int32(2))%32))))
	v84 = F_equal(m, l0, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L36
	}
L35:
	;
	goto L31
L36:
	;
	if v84 != 0 {
		v110 = v76
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v86 = int32(1)
	v89 = v78 + v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v89 < v90 {
		v76 = v76 + v86
		v78 = v89
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	F_errmsg_internal(m, int32(_a_F_mcv_match_expression_3), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_mcv_match_expression_1), int32(1573), int32(_a_F_mcv_match_expression_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	v4 = Fn13921(m, l0, l1, int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_mic2latin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v96)
	return v92 - l0
L2:
	;
	v87 = l1
	v92 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v14 = l1
	v15 = l2
	v19 = l0
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l5 != 0 {
		v87 = v14
		v92 = v19
		goto L1
	} else {
		goto L35
	}
L7:
	;
	if l5 != 0 {
		v87 = v14
		v92 = v19
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v32 = base.I32_extend8_s(v23)
	if int32(0) <= v32 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_report_invalid_encoding(m, int32(7), v19, v15)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	goto L6
L14:
	;
	v72 = int32(-1)
	v73 = v32
	v74 = int32(1)
	goto L16
L15:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.Ui32((v36+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v73)
	v77 = v14 + int32(1)
	v78 = v19 + v74
	v79 = v15 + v72
	if int32(0) < v79 {
		v14 = v77
		v15 = v79
		v19 = v78
		goto L5
	} else {
		goto L34
	}
L17:
	;
	if v15 < v58 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v58 = int32(2)
	goto L17
L19:
	;
	goto L20
L20:
	;
	if base.Ui32((v36+int32(112))&int32(255)) < base.Ui32(int32(12)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v58 = int32(3)
	goto L17
L22:
	;
	goto L23
L23:
	;
	if v36&int32(254) == int32(156) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = int32(4)
	goto L26
L25:
	;
	v57 = int32(1)
	goto L26
L26:
	;
	v58 = v57
	goto L17
L27:
	;
	if l5 != 0 {
		v87 = v14
		v92 = v19
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.B2i32(l3 != v23)|base.B2i32(v58 != int32(2)) != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	F_report_invalid_encoding(m, int32(7), v19, v15)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
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
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+1)))
	if int32(0) <= v67 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v72 = int32(-2)
	v73 = v67
	v74 = int32(2)
	goto L16
L34:
	;
	v87 = v77
	v92 = v78
	goto L1
L35:
	;
	F_report_untranslatable_char(m, int32(7), l4, v19, v15)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mic2latin_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
	return v105 - l0
L2:
	;
	v99 = l1
	v105 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = l1
	v16 = l2
	v21 = l0
	goto L5
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		v99 = v15
		v105 = v21
		goto L1
	} else {
		goto L36
	}
L7:
	;
	if l6 != 0 {
		v99 = v15
		v105 = v21
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v34 = base.I32_extend8_s(v25)
	if int32(0) <= v34 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_report_invalid_encoding(m, int32(7), v21, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	goto L6
L14:
	;
	v83 = int32(-1)
	v84 = v34
	v85 = int32(1)
	goto L16
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32((v38+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v84)
	v88 = v15 + int32(1)
	v89 = v21 + v85
	v90 = v16 + v83
	if int32(0) < v90 {
		v15 = v88
		v16 = v90
		v21 = v89
		goto L5
	} else {
		goto L35
	}
L17:
	;
	if v16 < v60 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v60 = int32(2)
	goto L17
L19:
	;
	goto L20
L20:
	;
	if base.Ui32((v38+int32(112))&int32(255)) < base.Ui32(int32(12)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v60 = int32(3)
	goto L17
L22:
	;
	goto L23
L23:
	;
	if v38&int32(254) == int32(156) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = int32(4)
	goto L26
L25:
	;
	v59 = int32(1)
	goto L26
L26:
	;
	v60 = v59
	goto L17
L27:
	;
	if l6 != 0 {
		v99 = v15
		v105 = v21
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.B2i32(l3 != v25)|base.B2i32(v60 != int32(2)) != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	F_report_invalid_encoding(m, int32(7), v21, v16)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
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
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+1)))
	if int32(0) <= v69 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v69&int32(255)-int32(128)))))
	if v77 == int32(0) {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v83 = int32(-2)
	v84 = v77
	v85 = int32(2)
	goto L16
L35:
	;
	v99 = v88
	v105 = v89
	goto L1
L36:
	;
	F_report_untranslatable_char(m, int32(7), l4, v21, v16)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v38 int32
	_ = v38
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
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v558 int32
	_ = v558
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v635 int32
	_ = v635
	var v652 int32
	_ = v652
	var v663 int32
	_ = v663
	var v679 int32
	_ = v679
	var v695 int32
	_ = v695
	var v710 int32
	_ = v710
	var v724 int32
	_ = v724
	var v737 int32
	_ = v737
	var v754 int32
	_ = v754
	var v768 int32
	_ = v768
	var v783 int32
	_ = v783
	var v796 int32
	_ = v796
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
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
	var v847 int32
	_ = v847
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
	var v948 int32
	_ = v948
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v992 int32
	_ = v992
	var v1005 int32
	_ = v1005
	var v1019 int32
	_ = v1019
	var v1034 int32
	_ = v1034
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1083 int32
	_ = v1083
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1126 int32
	_ = v1126
	var v1139 int32
	_ = v1139
	var v1151 int32
	_ = v1151
	var v1165 int32
	_ = v1165
	var v1178 int32
	_ = v1178
	var v1194 int32
	_ = v1194
	var v1206 int32
	_ = v1206
	var v1218 int32
	_ = v1218
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1262 int32
	_ = v1262
	var v1277 int32
	_ = v1277
	var v1292 int32
	_ = v1292
	var v1307 int32
	_ = v1307
	var v1322 int32
	_ = v1322
	var v1335 int64
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1348 int32
	_ = v1348
	var v1360 int32
	_ = v1360
	var v1372 int32
	_ = v1372
	var v1389 int32
	_ = v1389
	var v1401 int32
	_ = v1401
	var v1413 int32
	_ = v1413
	var v1432 int32
	_ = v1432
	var v1433 int64
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
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
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(528)
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
	v36 = v3
	v37 = int32(-1)
	v38 = v3
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
	m.G0 = v21 + int32(528)
	return
L4:
	;
	goto L3
L5:
	;
	if v37 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1432 = int32(m.ExcTag)
	v1433 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1432 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v30
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
	v845 = v36
	v847 = v38
	goto L10
L10:
	;
	if v847 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v71 = int32(0)
	v86 = F_get_db_info(m, l0, int32(8), v21+int32(484), v71, v71, v71, v71, v71, v71, v71, v21+int32(472), v71, v71, v71, v71, v71, v71)
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v21)+484))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2034), int32(_a_F_movedb_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2058), int32(_a_F_movedb_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_aclcheck_error(m, v298, int32(42), l1)
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v21)+472))
	if v272 == v371 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2080), int32(_a_F_movedb_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v412 = F_CountOtherDBBackends(m, v156, v21+int32(480), v21+int32(476))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v21)+480))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v21)+476))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2104), int32(_a_F_movedb_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v533 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_WaitForProcSignalBarrier(m, v533)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_DropDatabaseBuffers(m, v156)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v569 = F_AllocateDir(m, v508)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+420)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+416)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v814 = v21 + int32(416)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v814
	F_before_shmem_exit(m, int32(545), v814)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L7
	} else {
		goto L94
	}
L67:
	;
	if v569 == int32(0) {
		v796 = v36
		goto L66
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v583 = F_ReadDir(m, v569, v508)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L7
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L7
	} else {
		goto L89
	}
L70:
	;
	if v583 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v597 = v36
	v598 = v583
	goto L74
L72:
	;
	v635 = v36
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_FreeDir(m, v569)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L7
	} else {
		goto L84
	}
L74:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+19)))
	if v603 != int32(46) {
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v635 = v621
	goto L73
L76:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+20)))
	if v606 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+20)))
	if v607 != int32(46) {
		goto L69
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v621 = F_ReadDir(m, v569, v508)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+21)))
	if v610 != 0 {
		goto L69
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if v621 != 0 {
		v597 = v621
		v598 = v621
		goto L74
	} else {
		goto L83
	}
L83:
	;
	goto L75
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	v663 = F_rmdir(m, v508)
	mBase = m.M
	if v663 == int32(0) {
		v796 = v635
		goto L66
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v508
	F_errmsg_internal(m, int32(_a_F_movedb_6), v21+int32(48))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2177), int32(_a_F_movedb_2))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errcode(m, int32(325))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l0
	F_errmsg(m, int32(_a_F_movedb_7), v21-int32(-64))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errhint(m, int32(_a_F_movedb_8), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v56
	F_errfinish(m, int32(_a_F_movedb_1), int32(2166), int32(_a_F_movedb_2))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	goto L1
L94:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[2]))
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[3]))
	goto L95
L95:
	;
	v825 = v21 + int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v825)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v825))) = v21 + int32(84)
	goto L98
L96:
	;
	v836 = v156
	v837 = v508
	v838 = v272
	v839 = v56
	v840 = v21 + int32(416)
	v841 = v496
	v842 = v821
	v843 = v823
	v844 = v371
	v845 = v796
	v847 = int32(0)
	goto L10
L98:
	;
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v21 + int32(256)
	v859 = int32(0)
	base.MemoryFill(m, v21+int32(176), v859, int32(72))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+160)) = uint16(v859)
	v864 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+152)) = v864
	*(*int64)(unsafe.Add(mBase, uint32(v21)+144)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+128)) = uint16(v859)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v864
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v864
	F_copydir(m, v841, v837, v859)
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
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[2])) = v842
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_cancel_before_shmem_exit(m, int32(545), v840)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L7
	} else {
		goto L146
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v943 = v21 + int32(424)
	F_ScanKeyInit(m, v943, int32(2), int32(3), int32(62), l0)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v960 = int32(1)
	v963 = F_systable_beginscan(m, v839, int32(2671), v960, int32(0), v960, v943)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v975 = F_systable_getnext(m, v963)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	if v975 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1046 = v975 + int32(4)
	F_LockTuple(m, v839, v1046, int32(7))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L7
	} else {
		goto L116
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_errcode(m, int32(1283))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l0
	F_errmsg(m, int32(_a_F_movedb_0), v21)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_errfinish(m, int32(_a_F_movedb_1), int32(2232), int32(_a_F_movedb_2))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	goto L1
L116:
	;
	v1050 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+123)) = uint8(v1050)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+220)) = v838
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v839)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1070 = F_heap_modify_tuple(m, v975, v1053, v21+int32(176), v21+int32(144), v21+int32(112))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_CatalogTupleUpdate(m, v839, v1046, v1070)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_UnlockTuple(m, v839, v1046, int32(7))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _c_F_movedb[4]))
	if v1098 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1110 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1262), v836, v1110, v1110, v1110)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L7
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_systable_endscan(m, v963)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L7
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_RequestCheckpoint(m, int32(44))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_movedb[5])) = uint8(v1151)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_relation_close(m, v839, int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_cancel_before_shmem_exit(m, int32(545), v840)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[2])) = v842
	*(*int32)(unsafe.Add(mBase, _c_F_movedb[3])) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_StartTransactionCommand(m)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1229 = F_rmtree(m, v841)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L7
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_XLogBeginInsert(m)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L7
	} else {
		goto L139
	}
L133:
	;
	if v1229 != 0 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1243 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	if v1243 == int32(0) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v841
	F_errmsg(m, int32(_a_F_movedb_9), v21+int32(16))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_errfinish(m, int32(_a_F_movedb_1), int32(2296), int32(_a_F_movedb_2))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_XLogRegisterData(m, v21+int32(88), int32(8))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_XLogRegisterData(m, v21+int32(472), int32(4))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	v1335 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_UnlockSharedObjectForSession(m, v836)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_pfree(m, v841)
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_pfree(m, v837)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	goto L4
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_movedb_failure_callback(m, v21, v840)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+492)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v21)+488)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v21)+496)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+500)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v21)+504)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v839
	F_pg_re_throw(m)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L7
	} else {
		goto L148
	}
L148:
	;
	goto L1
L149:
	;
	v1437 = int32(v1433)
	m.G0 = v21
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+4))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	if v21+int32(84) == v1443 {
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
	if v1447 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+4))
	v1447 = v1445
	goto L154
L153:
	;
	v1447 = int32(0)
	goto L154
L154:
	;
	goto L151
L155:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v21)+524))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v21)+520))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v21)+516))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v21)+512))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v21)+508))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v21)+504))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v21)+500))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v21)+496))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v21)+492))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v21)+488))
	v27 = v1449
	v28 = v1453
	v29 = v1450
	v30 = v1448
	v31 = v1455
	v32 = v1452
	v33 = v1457
	v34 = v1456
	v35 = v1451
	v36 = v1454
	v37 = v1447
	v38 = v1439
	goto L2
L156:
	;
	goto L157
L157:
	;
	F___wasm_longjmp(m, v1440, v1439)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
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
