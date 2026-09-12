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
	v8 = int32(4515712)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l4
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
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
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
	var v61 int32
	_ = v61
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
	var v106 int32
	_ = v106
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
		v106 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v23
	v115 = F_makeRangeVar(m, int32(0), l0, int32(-1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v53 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v106 = v4
	goto L6
L9:
	;
	goto L10
L10:
	;
	v61 = v4
	v64 = v4
	goto L11
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = F_palloc0(m, int32(20))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v106 = v95
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = int64(81)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66+v64<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v80 = F_palloc0(m, int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(69)
	v86 = F_makeString(m, v78)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v89 = F_lcons(m, v86, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v80
	v95 = F_lappend(m, v61, v68)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v98 = v64 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v98 < v99 {
		v61 = v95
		v64 = v98
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v115
	v120 = F_list_make1_impl(m, int32(1), v13)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
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
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = v11 + int32(32)
	v19 = int32(1)
	goto L1
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13+v19<<(uint(int32(2))%32))))
	if v28 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1452))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1448))
	if base.Ui32(v63) <= base.Ui32(v64) {
		v244 = v62
		goto L13
	} else {
		goto L14
	}
L3:
	;
	goto L2
L4:
	;
	v61 = v19
	goto L3
L5:
	;
	goto L6
L6:
	;
	v32 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13+v32<<(uint(int32(2))%32))))
	if v36 == int32(0) {
		v61 = v32
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v39 = int32(2)
	v40 = v19 + v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13+v40<<(uint(v39)%32))))
	if v44 == int32(0) {
		v61 = v40
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v48 = v19 + int32(3)
	if v48 == int32(32) {
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v13+v48<<(uint(int32(2))%32))))
	if v56 == int32(0) {
		v61 = v48
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v19 = v19 + int32(4)
	goto L1
L13:
	;
	return v244
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1440))
	v69 = v66 << (uint(int32(base.Ui32(v61)>>(uint(int32(1))%32))) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1444))
	if base.Ui32(v69) < base.Ui32(v70) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = v69
	goto L17
L16:
	;
	v72 = v70
	goto L17
L17:
	;
	v73 = v63 - v64
	if base.Ui32(v72) < base.Ui32(v73) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = v72
	goto L20
L19:
	;
	v75 = v73
	goto L20
L20:
	;
	v79 = int32(base.Ui32(v75)>>(uint(int32(10))%32)) & int32(4194300)
	v81 = v79 + int32(584)
	v83 = v81 & int32(4092)
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = v79 - v83 + int32(4680)
	goto L23
L22:
	;
	v87 = v81
	goto L23
L23:
	;
	if base.Ui32(v75) <= base.Ui32(v87) {
		v244 = v62
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v91 = int32(base.Ui32(v75-v87) >> (uint(int32(12)) % 32))
	if base.Ui32(v91) < base.Ui32(l1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = l1 << (uint(int32(2)) % 32)
	v96 = v94 + int32(584)
	v98 = v96 & int32(4092)
	if v98 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v109 = v75
	v110 = v87
	v111 = v91
	goto L27
L27:
	;
	v112 = int32(4515764)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v115
	v118 = F_dsm_create(m, v109, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v102 = v94 - v98 + int32(4680)
	goto L30
L29:
	;
	v102 = v96
	goto L30
L30:
	;
	v105 = v102 + l1<<(uint(int32(12))%32)
	if base.Ui32(int32(134217728)) < base.Ui32(v105) {
		v244 = v62
		goto L13
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v73) < base.Ui32(v105) {
		v244 = v62
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v109 = v105
	v110 = v102
	v111 = l1
	goto L27
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v113
	if v118 == int32(0) {
		v244 = v62
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_dsm_pin_segment(m, v118)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v61<<(uint(int32(2))%32))+32)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+1456))
	if base.Ui32(v135) < base.Ui32(v61) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+1456)) = v61
	goto L39
L38:
	;
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v138) < base.Ui32(v61) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+648)) = v61
	goto L42
L41:
	;
	goto L42
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+1448))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+1448)) = v142 + v109
	v145 = int32(20)
	v147 = l0 + v61*v145
	v149 = v147 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v118
	v152 = v147 + int32(16)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v153 + int32(584)
	v160 = v147 + v145
	v162 = v153 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v162)+4)) = v164
	v166 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+32)) = uint8(v166)
	*(*int64)(unsafe.Add(mBase, uint32(v162)+12)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v162)+20)) = v164
	v172 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+28)) = v172
	if v162 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	F_FreePageManagerPut(m, v185, int32(base.Ui32(v110)>>(uint(int32(12))%32)), v111)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L33
	} else {
		goto L47
	}
L44:
	;
	v178 = v162 - v153 + v166
	goto L46
L45:
	;
	v178 = v172
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v178
	v184 = F___memset(m, v153+int32(68), int32(0), int32(516))
	mBase = m.M
	goto L43
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v61 ^ v192 ^ int32(216163848)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v111
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v109
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v202 = int32(15)
	v205 = int32(32) - base.I32_clz(v111)
	if base.Ui32(v202) <= base.Ui32(v205) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v208 = v202
	goto L50
L49:
	;
	v208 = v205
	goto L50
L50:
	;
	if v111 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v210 = v208
	goto L53
L52:
	;
	v210 = int32(0)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+20)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v213 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+12)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v218 = int32(2)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216+v217<<(uint(v218)%32))+160))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+24)) = uint8(v224)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v226+v228<<(uint(v218)%32))+160)) = v61
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	if v234 == v213 {
		v244 = v149
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v237 = F_get_segment_by_index(m, l0, v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L33
	} else {
		goto L55
	}
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = v61
	v244 = v149
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
						F_errmsg_internal(m, int32(42795), v16+int32(16))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493961), int32(232), int32(241438))
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
				F_errmsg_internal(m, int32(39703), v16)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493961), int32(228), int32(241438))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	if l0 == int32(0) {
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
	v16 = F_palloc(m, int32(0))
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = F_palloc(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(0)
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v35 = v30 << (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = F_lappend(m, v33, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L3
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43+v35))) = v45
	v48 = v30 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 < v49 {
		v30 = v48
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
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
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
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
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
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
	goto L104
L3:
	;
	v21 = v4
	v22 = v15
	v24 = v4
	goto L4
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[1265]))))
	if v30 == int32(0) {
		v50 = v24
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v50 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v34 = v22 & int32(255)
	if v34 != int32(47) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v44))))
	if v47 != 0 {
		v21 = v44
		v22 = v47
		v24 = v45
		goto L4
	} else {
		goto L13
	}
L9:
	;
	if v34 != v30 {
		v50 = v24
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if v30 != int32(47) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v40 = v21 + int32(1)
	v44 = v40
	v45 = v40
	goto L8
L12:
	;
	v44 = v21 + int32(1)
	v45 = v24
	goto L8
L13:
	;
	v50 = v45
	goto L6
L14:
	;
	goto L18
L15:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v170 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	v165 = F_strlen(m, v154)
	mBase = m.M
	goto L15
L18:
	;
	goto L19
L19:
	;
	v59 = int32(1023)
	if (l0^l2)&int32(3) != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v158)
	goto L16
L21:
	;
	v139 = v134
	v140 = v135
	v141 = v136
	goto L43
L22:
	;
	if v129 == int32(0) {
		v154 = v127
		v155 = v128
		goto L20
	} else {
		goto L42
	}
L23:
	;
	v127 = l2
	v128 = l0
	v129 = v59
	goto L22
L24:
	;
	goto L25
L25:
	;
	if l2&int32(3) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v96 == int32(0) {
		v154 = v93
		v155 = v94
		goto L20
	} else {
		goto L35
	}
L27:
	;
	v93 = l2
	v94 = l0
	v95 = v59
	v96 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v72 = l2
	v73 = l0
	v74 = v59
	goto L30
L30:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v76)
	if v76 == int32(0) {
		v134 = v72
		v135 = v73
		v136 = v74
		goto L21
	} else {
		goto L32
	}
L31:
	;
	v93 = v87
	v94 = v81
	v95 = v83
	v96 = v85
	goto L26
L32:
	;
	v80 = int32(1)
	v81 = v73 + v80
	v83 = v74 - v80
	v84 = int32(0)
	v85 = base.B2i32(v83 != v84)
	v87 = v72 + v80
	if v87&int32(3) == v84 {
		v93 = v87
		v94 = v81
		v95 = v83
		v96 = v85
		goto L26
	} else {
		goto L33
	}
L33:
	;
	if v83 != 0 {
		v72 = v87
		v73 = v81
		v74 = v83
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v99 == int32(0) {
		v127 = v93
		v128 = v94
		v129 = v95
		goto L22
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v95) < base.Ui32(int32(4)) {
		v127 = v93
		v128 = v94
		v129 = v95
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v105 = v93
	v106 = v94
	v107 = v95
	goto L38
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v113 = int32(-2139062144)
	if (int32(16843008)-v110|v110)&v113 != v113 {
		v134 = v105
		v135 = v106
		v136 = v107
		goto L21
	} else {
		goto L40
	}
L39:
	;
	v127 = v121
	v128 = v119
	v129 = v123
	goto L22
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v110
	v118 = int32(4)
	v119 = v106 + v118
	v121 = v105 + v118
	v123 = v107 - v118
	if base.Ui32(int32(3)) < base.Ui32(v123) {
		v105 = v121
		v106 = v119
		v107 = v123
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v134 = v127
	v135 = v128
	v136 = v129
	goto L21
L43:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v143)
	if v143 == int32(0) {
		v154 = v139
		v155 = v140
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v154 = v150
	v155 = v148
	goto L20
L45:
	;
	v147 = int32(1)
	v148 = v140 + v147
	v150 = v139 + v147
	v152 = v141 - v147
	if v152 != 0 {
		v139 = v150
		v140 = v148
		v141 = v152
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v171 = F_strlen(m, l0)
	mBase = m.M
	v176 = v171 + l0
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_canonicalize_path_enc(m, l0)
	mBase = m.M
	v241 = F_strlen(m, l0)
	mBase = m.M
	v242 = v241 + (v50 - int32(109))
	if v242 <= int32(0) {
		goto L2
	} else {
		goto L71
	}
L50:
	;
	v184 = v176 - int32(1)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v185 == int32(47) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v192 = v184
	goto L56
L52:
	;
	if base.Ui32(l0) < base.Ui32(v184) {
		v176 = v184
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L54
L56:
	;
	if base.Ui32(l0) < base.Ui32(v192) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v211 = v192
	goto L62
L58:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v202 != int32(47) {
		v192 = v192 - int32(1)
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L60
L62:
	;
	if base.Ui32(l0) < base.Ui32(v211) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if l0 == v211 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v218 = v211 - int32(1)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v219 == int32(47) {
		v211 = v218
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L66
L68:
	;
	v227 = l0 + base.B2i32(v170 == int32(47))
	goto L70
L69:
	;
	v227 = v211
	goto L70
L70:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v228)
	goto L49
L71:
	;
	v245 = l0 + v242
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245-int32(1)))))
	if v248 != int32(47) {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v252 = v50 + int32(277759)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v253 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v307 != 0 {
		goto L2
	} else {
		goto L87
	}
L74:
	;
	v256 = v245
	v257 = v252
	v258 = v253
	goto L77
L75:
	;
	v286 = v252
	goto L76
L76:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v295 != 0 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if v264 == int32(0) {
		goto L2
	} else {
		goto L79
	}
L78:
	;
	v286 = v279
	goto L76
L79:
	;
	v268 = v258 & int32(255)
	if v268 == v264 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v278 = int32(1)
	v279 = v257 + v278
	v281 = v256 + v278
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v282 != 0 {
		v256 = v281
		v257 = v279
		v258 = v282
		goto L77
	} else {
		goto L83
	}
L81:
	;
	v270 = int32(47)
	if base.B2i32(v268 == v270)&base.B2i32(v264 == v270) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v307 = base.I32_extend8_s(v258) - base.I32_extend8_s(v264)
	goto L73
L83:
	;
	goto L78
L84:
	;
	v296 = int32(-1)
	goto L86
L85:
	;
	v296 = int32(0)
	goto L86
L86:
	;
	v307 = v296
	goto L73
L87:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v308)
	v310 = F_strlen(m, l0)
	mBase = m.M
	v311 = v310 + l0
	if base.Ui32(v311) <= base.Ui32(l0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v344 = l1 + v50
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v345 == int32(0) {
		goto L1
	} else {
		goto L95
	}
L89:
	;
	v314 = v311 - int32(1)
	if base.Ui32(v314) <= base.Ui32(l0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v319 = v314
	goto L91
L91:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v326 != int32(47) {
		goto L88
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v319))) = uint8(v329)
	v332 = v319 - int32(1)
	if base.Ui32(l0) < base.Ui32(v332) {
		v319 = v332
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v348 = F_strlen(m, l0)
	mBase = m.M
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v344
	if v349 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v353 = int32(571217)
	goto L98
L97:
	;
	v353 = int32(757756)
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v353
	v359 = F_pg_snprintf(m, l0+v348, int32(1024)-v348, int32(175865), v13)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	return
L100:
	;
	goto L1
L101:
	;
	goto L1
L102:
	;
	v483 = F_strlen(m, v472)
	mBase = m.M
	goto L101
L104:
	;
	goto L105
L105:
	;
	v377 = int32(1023)
	if (l0^l1)&int32(3) != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v476 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v476)
	goto L102
L107:
	;
	v457 = v452
	v458 = v453
	v459 = v454
	goto L129
L108:
	;
	if v447 == int32(0) {
		v472 = v445
		v473 = v446
		goto L106
	} else {
		goto L128
	}
L109:
	;
	v445 = l1
	v446 = l0
	v447 = v377
	goto L108
L110:
	;
	goto L111
L111:
	;
	if l1&int32(3) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v414 == int32(0) {
		v472 = v411
		v473 = v412
		goto L106
	} else {
		goto L121
	}
L113:
	;
	v411 = l1
	v412 = l0
	v413 = v377
	v414 = int32(1)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v390 = l1
	v391 = l0
	v392 = v377
	goto L116
L116:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v394)
	if v394 == int32(0) {
		v452 = v390
		v453 = v391
		v454 = v392
		goto L107
	} else {
		goto L118
	}
L117:
	;
	v411 = v405
	v412 = v399
	v413 = v401
	v414 = v403
	goto L112
L118:
	;
	v398 = int32(1)
	v399 = v391 + v398
	v401 = v392 - v398
	v402 = int32(0)
	v403 = base.B2i32(v401 != v402)
	v405 = v390 + v398
	if v405&int32(3) == v402 {
		v411 = v405
		v412 = v399
		v413 = v401
		v414 = v403
		goto L112
	} else {
		goto L119
	}
L119:
	;
	if v401 != 0 {
		v390 = v405
		v391 = v399
		v392 = v401
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v417 == int32(0) {
		v445 = v411
		v446 = v412
		v447 = v413
		goto L108
	} else {
		goto L122
	}
L122:
	;
	if base.Ui32(v413) < base.Ui32(int32(4)) {
		v445 = v411
		v446 = v412
		v447 = v413
		goto L108
	} else {
		goto L123
	}
L123:
	;
	v423 = v411
	v424 = v412
	v425 = v413
	goto L124
L124:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v431 = int32(-2139062144)
	if (int32(16843008)-v428|v428)&v431 != v431 {
		v452 = v423
		v453 = v424
		v454 = v425
		goto L107
	} else {
		goto L126
	}
L125:
	;
	v445 = v439
	v446 = v437
	v447 = v441
	goto L108
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v428
	v436 = int32(4)
	v437 = v424 + v436
	v439 = v423 + v436
	v441 = v425 - v436
	if base.Ui32(int32(3)) < base.Ui32(v441) {
		v423 = v439
		v424 = v437
		v425 = v441
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v452 = v445
	v453 = v446
	v454 = v447
	goto L107
L129:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v461)
	if v461 == int32(0) {
		v472 = v457
		v473 = v458
		goto L106
	} else {
		goto L131
	}
L130:
	;
	v472 = v468
	v473 = v466
	goto L106
L131:
	;
	v465 = int32(1)
	v466 = v458 + v465
	v468 = v457 + v465
	v470 = v459 - v465
	if v470 != 0 {
		v457 = v468
		v458 = v466
		v459 = v470
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
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
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
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
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 {
	case 0:
		v757 = F_palloc(m, int32(5))
		mBase = m.M
		v758 = m.ExcPending
		if v758 != 0 {
			return int32(0)
		} else {
			v759 = int32(2)
			*(*uint8)(unsafe.Add(mBase, uint32(v757)+4)) = uint8(v759)
			*(*int32)(unsafe.Add(mBase, uint32(v757))) = int32(20)
			v763 = v757
			m.G0 = v10 + int32(48)
			return v763
		}
	case 1:
		if l1 != 0 {
			v444 = int32(1)
		} else {
			v444 = int32(5)
		}
		v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if int32(126) <= v446 {
			v454 = v446 - int32(1636608432)
			if v445&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v446) {
					v563 = v445
					v564 = v446
					v565 = v454
					v566 = v454
					v567 = v454
					for {
						v569 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
						v570 = v569 + v566
						v571 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
						v573 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
						v574 = v573 + v567
						v576 = int32(4)
						v578 = v571 + v565 - v574 ^ base.I32_rotl(v574, v576)
						v582 = v570 - v578 ^ base.I32_rotl(v578, int32(6))
						v583 = v574 + v570
						v584 = v578 + v583
						v585 = v582 + v584
						v589 = v583 - v582 ^ base.I32_rotl(v582, int32(8))
						v593 = v584 - v589 ^ base.I32_rotl(v589, int32(16))
						v597 = v585 - v593 ^ base.I32_rotl(v593, int32(19))
						v598 = v589 + v585
						v599 = v593 + v598
						v600 = v597 + v599
						v604 = v598 - v597 ^ base.I32_rotl(v597, v576)
						v605 = int32(12)
						v606 = v563 + v605
						v608 = v564 - v605
						if base.Ui32(int32(11)) < base.Ui32(v608) {
							v563 = v606
							v564 = v608
							v565 = v599
							v566 = v600
							v567 = v604
							continue
						} else {
							break
						}
						break
					}
					v611 = v606
					v612 = v608
					v613 = v599
					v614 = v600
					v615 = v604
				} else {
					v611 = v445
					v612 = v446
					v613 = v454
					v614 = v454
					v615 = v454
				}
				switch v612 - int32(1) {
				case 0:
					v674 = v613
					v675 = v614
					v676 = v615
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 1:
					v667 = v613
					v668 = v614
					v669 = v615
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 2:
					v660 = v613
					v661 = v614
					v662 = v615
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 3:
					v654 = v614
					v655 = v615
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 4:
					v650 = v614
					v651 = v615
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 5:
					v644 = v614
					v645 = v615
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 6:
					v638 = v614
					v639 = v615
					v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+6)))
					v644 = v640<<(uint(int32(16))%32) + v638
					v645 = v639
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 7:
					v633 = v615
					v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+7)))
					v638 = v634<<(uint(int32(24))%32) + v614
					v639 = v633
					v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+6)))
					v644 = v640<<(uint(int32(16))%32) + v638
					v645 = v639
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 8:
					v628 = v615
					v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+8)))
					v633 = v629<<(uint(int32(8))%32) + v628
					v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+7)))
					v638 = v634<<(uint(int32(24))%32) + v614
					v639 = v633
					v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+6)))
					v644 = v640<<(uint(int32(16))%32) + v638
					v645 = v639
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 9:
					v623 = v615
					v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+9)))
					v628 = v624<<(uint(int32(16))%32) + v623
					v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+8)))
					v633 = v629<<(uint(int32(8))%32) + v628
					v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+7)))
					v638 = v634<<(uint(int32(24))%32) + v614
					v639 = v633
					v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+6)))
					v644 = v640<<(uint(int32(16))%32) + v638
					v645 = v639
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				case 10:
					v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+10)))
					v623 = v619<<(uint(int32(24))%32) + v615
					v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+9)))
					v628 = v624<<(uint(int32(16))%32) + v623
					v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+8)))
					v633 = v629<<(uint(int32(8))%32) + v628
					v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+7)))
					v638 = v634<<(uint(int32(24))%32) + v614
					v639 = v633
					v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+6)))
					v644 = v640<<(uint(int32(16))%32) + v638
					v645 = v639
					v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+5)))
					v650 = v646<<(uint(int32(8))%32) + v644
					v651 = v645
					v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+4)))
					v654 = v650 + v652
					v655 = v651
					v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+3)))
					v660 = v656<<(uint(int32(24))%32) + v613
					v661 = v654
					v662 = v655
					v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+2)))
					v667 = v663<<(uint(int32(16))%32) + v660
					v668 = v661
					v669 = v662
					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
					v674 = v670<<(uint(int32(8))%32) + v667
					v675 = v668
					v676 = v669
					v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
					v681 = v674 + v677
					v682 = v675
					v683 = v676
				default:
					v681 = v613
					v682 = v614
					v683 = v615
				}
			} else {
				if base.Ui32(v446) < base.Ui32(int32(12)) {
					v509 = v445
					v510 = v446
					v511 = v454
					v512 = v454
					v513 = v454
				} else {
					v461 = v445
					v462 = v446
					v463 = v454
					v464 = v454
					v465 = v454
					for {
						v467 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
						v468 = v467 + v464
						v469 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
						v471 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
						v472 = v471 + v465
						v474 = int32(4)
						v476 = v469 + v463 - v472 ^ base.I32_rotl(v472, v474)
						v480 = v468 - v476 ^ base.I32_rotl(v476, int32(6))
						v481 = v472 + v468
						v482 = v476 + v481
						v483 = v480 + v482
						v487 = v481 - v480 ^ base.I32_rotl(v480, int32(8))
						v491 = v482 - v487 ^ base.I32_rotl(v487, int32(16))
						v495 = v483 - v491 ^ base.I32_rotl(v491, int32(19))
						v496 = v487 + v483
						v497 = v491 + v496
						v498 = v495 + v497
						v502 = v496 - v495 ^ base.I32_rotl(v495, v474)
						v503 = int32(12)
						v504 = v461 + v503
						v506 = v462 - v503
						if base.Ui32(int32(11)) < base.Ui32(v506) {
							v461 = v504
							v462 = v506
							v463 = v497
							v464 = v498
							v465 = v502
							continue
						} else {
							break
						}
						break
					}
					v509 = v504
					v510 = v506
					v511 = v497
					v512 = v498
					v513 = v502
				}
				switch v510 - int32(1) {
				case 0:
					v560 = v511
					v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
					v681 = v560 + v561
					v682 = v512
					v683 = v513
				case 1:
					v555 = v511
					v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
					v560 = v556<<(uint(int32(8))%32) + v555
					v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
					v681 = v560 + v561
					v682 = v512
					v683 = v513
				case 2:
					v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+2)))
					v555 = v551<<(uint(int32(16))%32) + v511
					v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
					v560 = v556<<(uint(int32(8))%32) + v555
					v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
					v681 = v560 + v561
					v682 = v512
					v683 = v513
				case 3:
					v548 = v512
					v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v681 = v549 + v511
					v682 = v548
					v683 = v513
				case 4:
					v545 = v512
					v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+4)))
					v548 = v545 + v546
					v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v681 = v549 + v511
					v682 = v548
					v683 = v513
				case 5:
					v540 = v512
					v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+5)))
					v545 = v541<<(uint(int32(8))%32) + v540
					v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+4)))
					v548 = v545 + v546
					v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v681 = v549 + v511
					v682 = v548
					v683 = v513
				case 6:
					v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+6)))
					v540 = v536<<(uint(int32(16))%32) + v512
					v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+5)))
					v545 = v541<<(uint(int32(8))%32) + v540
					v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+4)))
					v548 = v545 + v546
					v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v681 = v549 + v511
					v682 = v548
					v683 = v513
				case 7:
					v531 = v513
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v534 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
					v681 = v532 + v511
					v682 = v534 + v512
					v683 = v531
				case 8:
					v526 = v513
					v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+8)))
					v531 = v527<<(uint(int32(8))%32) + v526
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v534 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
					v681 = v532 + v511
					v682 = v534 + v512
					v683 = v531
				case 9:
					v521 = v513
					v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+9)))
					v526 = v522<<(uint(int32(16))%32) + v521
					v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+8)))
					v531 = v527<<(uint(int32(8))%32) + v526
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v534 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
					v681 = v532 + v511
					v682 = v534 + v512
					v683 = v531
				case 10:
					v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+10)))
					v521 = v517<<(uint(int32(24))%32) + v513
					v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+9)))
					v526 = v522<<(uint(int32(16))%32) + v521
					v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+8)))
					v531 = v527<<(uint(int32(8))%32) + v526
					v532 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
					v534 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
					v681 = v532 + v511
					v682 = v534 + v512
					v683 = v531
				default:
					v681 = v511
					v682 = v512
					v683 = v513
				}
			}
			v686 = int32(14)
			v688 = v682 ^ v683 - base.I32_rotl(v682, v686)
			v692 = v688 ^ v681 - base.I32_rotl(v688, int32(11))
			v696 = v692 ^ v682 - base.I32_rotl(v692, int32(25))
			v700 = v696 ^ v688 - base.I32_rotl(v696, int32(16))
			v704 = v700 ^ v692 - base.I32_rotl(v700, int32(4))
			v708 = v704 ^ v696 - base.I32_rotl(v704, v686)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v708 ^ v700 - base.I32_rotl(v708, int32(24))
			v720 = F_pg_snprintf(m, v10+int32(38), int32(10), int32(29625), v10+int32(32))
			mBase = m.M
			v721 = m.ExcPending
			if v721 != 0 {
				return int32(0)
			} else {
				v727 = int32(8)
				v728 = v444 | int32(16)
				v729 = v10 + int32(38)
				v731 = v727 + int32(5)
				v732 = F_palloc(m, v731)
				mBase = m.M
				v733 = m.ExcPending
				if v733 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v732)+4)) = uint8(v728)
					*(*int32)(unsafe.Add(mBase, uint32(v732))) = v731 << (uint(int32(2)) % 32)
					if v727 != 0 {
						v740 = F__emscripten_memcpy_bulkmem(m, v732+int32(5), v729, v727)
						mBase = m.M
					} else {
					}
					v763 = v732
					m.G0 = v10 + int32(48)
					return v763
				}
			}
		} else {
			v727 = v446
			v728 = v444
			v729 = v445
			v731 = v727 + int32(5)
			v732 = F_palloc(m, v731)
			mBase = m.M
			v733 = m.ExcPending
			if v733 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v732)+4)) = uint8(v728)
				*(*int32)(unsafe.Add(mBase, uint32(v732))) = v731 << (uint(int32(2)) % 32)
				if v727 != 0 {
					v740 = F__emscripten_memcpy_bulkmem(m, v732+int32(5), v729, v727)
					mBase = m.M
				} else {
				}
				v763 = v732
				m.G0 = v10 + int32(48)
				return v763
			}
		}
	case 2:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v29 = m.G0
		v31 = v29 - int32(32)
		m.G0 = v31
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
		if base.Ui32(int32(49152)) <= base.Ui32(v33) {
			if v33 != int32(61440) {
				if v33 != int32(53248) {
					v47 = F_pstrdup(m, int32(528221))
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
							v425 = v145
							v426 = int32(4)
							v427 = v137
							v429 = v425 + int32(5)
							v430 = F_palloc(m, v429)
							mBase = m.M
							v431 = m.ExcPending
							if v431 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
								*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
								if v425 != 0 {
									v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
									mBase = m.M
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v441 = m.ExcPending
								if v441 != 0 {
									return int32(0)
								} else {
									v763 = v430
									m.G0 = v10 + int32(48)
									return v763
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
							v419 = F_pg_snprintf(m, v10+int32(38), int32(10), int32(29625), v10+int32(16))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int32(0)
							} else {
								v425 = int32(8)
								v426 = int32(20)
								v427 = v10 + int32(38)
								v429 = v425 + int32(5)
								v430 = F_palloc(m, v429)
								mBase = m.M
								v431 = m.ExcPending
								if v431 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
									*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
									if v425 != 0 {
										v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
										mBase = m.M
									} else {
									}
									F_pfree(m, v137)
									mBase = m.M
									v441 = m.ExcPending
									if v441 != 0 {
										return int32(0)
									} else {
										v763 = v430
										m.G0 = v10 + int32(48)
										return v763
									}
								}
							}
						}
					}
				} else {
					v41 = F_pstrdup(m, int32(11457))
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
							v425 = v145
							v426 = int32(4)
							v427 = v137
							v429 = v425 + int32(5)
							v430 = F_palloc(m, v429)
							mBase = m.M
							v431 = m.ExcPending
							if v431 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
								*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
								if v425 != 0 {
									v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
									mBase = m.M
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v441 = m.ExcPending
								if v441 != 0 {
									return int32(0)
								} else {
									v763 = v430
									m.G0 = v10 + int32(48)
									return v763
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
							v419 = F_pg_snprintf(m, v10+int32(38), int32(10), int32(29625), v10+int32(16))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int32(0)
							} else {
								v425 = int32(8)
								v426 = int32(20)
								v427 = v10 + int32(38)
								v429 = v425 + int32(5)
								v430 = F_palloc(m, v429)
								mBase = m.M
								v431 = m.ExcPending
								if v431 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
									*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
									if v425 != 0 {
										v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
										mBase = m.M
									} else {
									}
									F_pfree(m, v137)
									mBase = m.M
									v441 = m.ExcPending
									if v441 != 0 {
										return int32(0)
									} else {
										v763 = v430
										m.G0 = v10 + int32(48)
										return v763
									}
								}
							}
						}
					}
				}
			} else {
				v44 = F_pstrdup(m, int32(11446))
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
						v425 = v145
						v426 = int32(4)
						v427 = v137
						v429 = v425 + int32(5)
						v430 = F_palloc(m, v429)
						mBase = m.M
						v431 = m.ExcPending
						if v431 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
							*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
							if v425 != 0 {
								v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
								mBase = m.M
							} else {
							}
							F_pfree(m, v137)
							mBase = m.M
							v441 = m.ExcPending
							if v441 != 0 {
								return int32(0)
							} else {
								v763 = v430
								m.G0 = v10 + int32(48)
								return v763
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
						v419 = F_pg_snprintf(m, v10+int32(38), int32(10), int32(29625), v10+int32(16))
						mBase = m.M
						v420 = m.ExcPending
						if v420 != 0 {
							return int32(0)
						} else {
							v425 = int32(8)
							v426 = int32(20)
							v427 = v10 + int32(38)
							v429 = v425 + int32(5)
							v430 = F_palloc(m, v429)
							mBase = m.M
							v431 = m.ExcPending
							if v431 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
								*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
								if v425 != 0 {
									v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
									mBase = m.M
								} else {
								}
								F_pfree(m, v137)
								mBase = m.M
								v441 = m.ExcPending
								if v441 != 0 {
									return int32(0)
								} else {
									v763 = v430
									m.G0 = v10 + int32(48)
									return v763
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
				v84 = v33 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v84
			v91 = v33 & int32(49152)
			if v91 == int32(32768) {
				v94 = v33 << (uint(int32(1)) % 32) & int32(16384)
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
					v425 = v145
					v426 = int32(4)
					v427 = v137
					v429 = v425 + int32(5)
					v430 = F_palloc(m, v429)
					mBase = m.M
					v431 = m.ExcPending
					if v431 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
						*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
						if v425 != 0 {
							v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
							mBase = m.M
						} else {
						}
						F_pfree(m, v137)
						mBase = m.M
						v441 = m.ExcPending
						if v441 != 0 {
							return int32(0)
						} else {
							v763 = v430
							m.G0 = v10 + int32(48)
							return v763
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
					v419 = F_pg_snprintf(m, v10+int32(38), int32(10), int32(29625), v10+int32(16))
					mBase = m.M
					v420 = m.ExcPending
					if v420 != 0 {
						return int32(0)
					} else {
						v425 = int32(8)
						v426 = int32(20)
						v427 = v10 + int32(38)
						v429 = v425 + int32(5)
						v430 = F_palloc(m, v429)
						mBase = m.M
						v431 = m.ExcPending
						if v431 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v426)
							*(*int32)(unsafe.Add(mBase, uint32(v430))) = v429 << (uint(int32(2)) % 32)
							if v425 != 0 {
								v438 = F__emscripten_memcpy_bulkmem(m, v430+int32(5), v427, v425)
								mBase = m.M
							} else {
							}
							F_pfree(m, v137)
							mBase = m.M
							v441 = m.ExcPending
							if v441 != 0 {
								return int32(0)
							} else {
								v763 = v430
								m.G0 = v10 + int32(48)
								return v763
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
			v19 = int32(3)
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v19)
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(24)
			if v13 != 0 {
				v25 = int32(116)
			} else {
				v25 = int32(102)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(v25)
			v763 = v15
			m.G0 = v10 + int32(48)
			return v763
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v745 = m.ExcPending
		if v745 != 0 {
			return int32(0)
		} else {
			v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v746
			F_errmsg_internal(m, int32(485445), v10)
			mBase = m.M
			v750 = m.ExcPending
			if v750 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497025), int32(1403), int32(21233))
				mBase = m.M
				v755 = m.ExcPending
				if v755 != 0 {
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
	var v35 int32
	_ = v35
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
	var v73 int32
	_ = v73
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
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
	return v352
L4:
	;
	v25 = int32(0)
	goto L6
L5:
	;
	v332 = int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v334 = F_errsave_start(m, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L77
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = int32(0)
	v35 = v31
	goto L8
L7:
	;
	goto L5
L8:
	;
	switch v35 - int32(1) {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v321 + int32(1)
	if base.Ui32(v319) < base.Ui32(int32(16)) {
		v34 = v319
		v35 = v320
		goto L8
	} else {
		goto L76
	}
L12:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if base.Ui32((v237-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	if v25 == int32(16) {
		goto L52
	} else {
		goto L53
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v106 == int32(32) {
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
	v99 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+64)) = uint8(v102)
	v319 = int32(1)
	v320 = v99
	v321 = v42
	goto L11
L17:
	;
	v50 = int32(0)
	v51 = int32(1)
	switch v43 - int32(32) {
	case 0:
		v319 = v50
		v320 = v51
		v321 = v42
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
	v206 = int32(33)
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
		v352 = v51
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
	v73 = v68 - int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10+v73<<(uint(int32(2))%32))))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v89 + int32(1)
	if v73 != 0 {
		v68 = v73
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v319 = v34
	v320 = int32(3)
	v321 = v105
	goto L11
L29:
	;
	goto L30
L30:
	;
	switch v106 - int32(38) {
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
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v176 != 0 {
		goto L10
	} else {
		goto L46
	}
L32:
	;
	v138 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v105 + v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v141 - v138
	if v141 <= int32(0) {
		goto L10
	} else {
		goto L40
	}
L33:
	;
	v116 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v105 + v116
	if v25 == int32(0) {
		v206 = v118
		goto L13
	} else {
		goto L37
	}
L34:
	;
	if v106 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v106 != int32(124) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	if v118 != int32(124) {
		v206 = v118
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v127 = F_palloc(m, int32(12))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = int64(532575944707)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v134 + int32(1)
	goto L6
L40:
	;
	v147 = int32(0)
	if v25 == v147 {
		v352 = v147
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v153 = v25
	goto L42
L42:
	;
	v158 = v153 - int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10+v158<<(uint(int32(2))%32))))
	v164 = F_palloc(m, int32(12))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v352 = v147
	goto L3
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = int32(3)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v164
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v172 + int32(1)
	if v158 != 0 {
		v153 = v158
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v177 = int32(0)
	if v25 == v177 {
		v352 = v177
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v183 = v25
	goto L48
L48:
	;
	v188 = v183 - int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v10+v188<<(uint(int32(2))%32))))
	v194 = F_palloc(m, int32(12))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v352 = v177
	goto L3
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v194))) = int32(3)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v194
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v202 + int32(1)
	if v188 != 0 {
		v183 = v188
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v212 = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v214 = F_errsave_start(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+v25<<(uint(int32(2))%32)))) = v206
	v25 = v25 + int32(1)
	goto L6
L55:
	;
	if v214 == int32(0) {
		v352 = v212
		goto L3
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(16777477))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(27672), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errsave_finish(m, v213, int32(497761), int32(184), int32(301793))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v352 = v212
	goto L3
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(-64)+v34))) = uint8(v237)
	v319 = v34 + int32(1)
	v320 = int32(2)
	v321 = v236
	goto L11
L61:
	;
	goto L62
L62:
	;
	v252 = v10 - int32(-64)
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v252+v34))) = uint8(v254)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v254
	v264 = F_strtox_2(m, v252, v254, v254, int64(2147483648))
	mBase = m.M
	goto L63
L63:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v267 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v270 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v272 == int32(0) {
		goto L10
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v276 = F_palloc(m, int32(12))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = base.I32_wrap_i64(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(2)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+8)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v276
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v284 + int32(1)
	v288 = int32(0)
	if v25 == v288 {
		v25 = v288
		goto L6
	} else {
		goto L70
	}
L70:
	;
	v294 = v25
	goto L71
L71:
	;
	v299 = v294 - int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v10+v299<<(uint(int32(2))%32))))
	switch v303 - int32(33) {
	case 0, 5:
		goto L73
	default:
		v25 = v294
		goto L6
	}
L72:
	;
	goto L4
L73:
	;
	v307 = F_palloc(m, int32(12))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+4)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = int32(3)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+8)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v307
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v315 + int32(1)
	if v299 != 0 {
		v294 = v299
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L10
L77:
	;
	if v334 == int32(0) {
		v352 = v332
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(212189), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errsave_finish(m, v333, int32(497761), int32(211), int32(301793))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v352 = v332
	goto L3
}
func F_manifest_report_error(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	F_initStringInfo(m, v7+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v16 = F_appendStringInfoVA(m, v7+int32(16), l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v16
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	F_enlargeStringInfo(m, v7+int32(16), v21)
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v29 = F_appendStringInfoVA(m, v7+int32(16), l1, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v29 != 0 {
		v21 = v29
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39
	F_errmsg_internal(m, int32(206200), v7)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(497962), int32(1043), int32(211650))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v9) < base.Ui32(int32(25)) {
	} else {
		v15 = int32(base.Ui32(v9+int32(262120)) >> (uint(int32(2)) % 32))
		if v15&int32(65535) == int32(0) {
		} else {
			v20 = int32(1)
			v21 = int32(2)
			v25 = (v15 + v20) & int32(65535)
			if base.Ui32(v25) <= base.Ui32(v21) {
				v28 = v21
			} else {
				v28 = v25
			}
			v29 = int32(1)
			v30 = v28 - v29
			v34 = l0 + int32(24)
			v35 = int32(0)
			if base.Ui32(int32(3)) <= base.Ui32(v25) {
				v40 = v35
				v41 = v20
				for {
					v50 = v41<<(uint(int32(2))%32) + v34
					v52 = v50 - int32(4)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					if v53&int32(98304) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53 & int32(-98305)
					} else {
					}
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					if v59&int32(98304) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v59 & int32(-98305)
					} else {
					}
					v65 = int32(2)
					v68 = v40 + v65
					if v68 != v30&int32(-2) {
						v40 = v68
						v41 = v41 + v65
						continue
					} else {
						break
					}
					break
				}
				v72 = v41 + int32(1)
			} else {
				v72 = v35
			}
			if v30&v29 == int32(0) {
			} else {
				v84 = v34 + v72<<(uint(int32(2))%32)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
				if v85&int32(98304) == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v85 & int32(-98305)
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v18 == v6 {
		v91 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L27
	} else {
		goto L32
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v91
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
		v91 = v6
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
		v91 = v6
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	v34 = v6
	v36 = v6
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v43 = int32(2)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v36<<(uint(v43)%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 == v43 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v91 = v81
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
	v84 = v36 + int32(1)
	if v84 != v26 {
		v34 = v81
		v36 = v84
		goto L9
	} else {
		goto L31
	}
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+16)))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[715]))
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
		v81 = v34
		goto L14
	} else {
		goto L23
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v65 == int32(5) {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	switch v56 - int32(68) {
	case 0, 11:
		v81 = v34
		goto L14
	default:
		goto L18
	}
L20:
	;
	goto L21
L21:
	;
	switch v56 - int32(68) {
	case 0, 14:
		v81 = v34
		goto L14
	default:
		goto L18
	}
L22:
	;
	goto L17
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v70 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = F_rangeTableEntry_used(m, l3, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v79 = F_lappend(m, v34, v46)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L27
	} else {
		goto L30
	}
L27:
	;
	return int32(0)
L28:
	;
	if v73 == int32(0) {
		v81 = v34
		goto L14
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v81 = v79
	goto L14
L31:
	;
	goto L10
L32:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v110 + int32(4)
	F_errmsg(m, int32(706180), v16)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	F_errdetail(m, int32(599147), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(495788), int32(1694), int32(154553))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L36
	}
L36:
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
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
	v79 = v76
	v80 = v72
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
		v79 = v7
		v80 = int32(255)
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
	v79 = v7
	v80 = int32(664)
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
	v79 = v7
	v80 = v64
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
		goto L35
	} else {
		goto L36
	}
L30:
	;
	if l3 != l5 {
		v168 = v91
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v92 == int32(0) {
		v168 = v91
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v97 = F_make_opclause(m, v82, l0, v85, l3)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v97
	v104 = F_list_make1_impl(m, int32(1), v15+int32(-56))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v168 = v104
	goto L1
L35:
	;
	if v84 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v108 = F_get_collation_isdeterministic(m, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v108 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v168 = int32(0)
	goto L1
L39:
	;
	if v83 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v111 = F_op_in_opfamily(m, v79, l4)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v111 == int32(0) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v115 = F_make_opclause(m, v79, l0, v85, l5)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v115
	v122 = F_list_make1_impl(m, int32(1), v15+int32(-48))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v168 = v122
	goto L1
L45:
	;
	v132 = F_op_in_opfamily(m, v81, l4)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	v126 = F_pg_newlocale_from_collation(m, l5)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
	if v128 == int32(1) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v168 = int32(0)
	goto L1
L49:
	;
	if v132 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = int32(0)
	goto L1
L51:
	;
	goto L52
L52:
	;
	v137 = F_make_opclause(m, v81, l0, v85, l5)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v137
	v144 = F_list_make1_impl(m, int32(1), v15+int32(-52))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v146 = F_op_in_opfamily(m, v80, l4)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v146 == int32(0) {
		v168 = v144
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v150 = F_get_opcode(m, v80)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_fmgr_info(m, v150, v15+int32(-32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v158 = F_make_greater_string(m, v85, v15+int32(-32), l5)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v158 == int32(0) {
		v168 = v144
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v162 = F_make_opclause(m, v80, l0, v158, l5)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v164 = F_lappend(m, v144, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v168 = v164
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v6 = *(*int32)(unsafe.Add(mBase, _consts[620]))
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	if base.Ui32(int32(8)) < base.Ui32(v10) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	if int32(2) < v14 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[621]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[622]))
	if v26 <= v24 {
		v70 = v24
		v73 = v26
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v70 <= v73 {
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[623])))
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v70 = v67
	v73 = v58
	goto L9
L13:
	;
	v38 = v29 + int32(1)
	if v38 != int32(32) {
		v29 = v38
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v55 = F_StartChildProcess(m, int32(12))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	F_errmsg_internal(m, int32(85203), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(495628), int32(4408), int32(133846))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[622]))
	v60 = *(*int32)(unsafe.Add(mBase, _consts[621]))
	if v55 == int32(0) {
		v70 = v60
		v73 = v58
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[623]))) = v55
	v67 = v60 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[621])) = v67
	if v67 < v58 {
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
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[623])))
	if v84 != 0 {
		v100 = v84
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v104 = F_kill(m, v102, int32(12))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[624])))
	if v87 != 0 {
		v100 = v87
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[625])))
	if v90 != 0 {
		v100 = v90
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v92 = v76 - int32(3)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32))+uint32(_consts[623])))
	if v97 != 0 {
		v100 = v97
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v92 != 0 {
		v76 = v76 - int32(4)
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	goto L1
}
func F_mbms_add_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v3 = int32(0)
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
	v12 = v11
	goto L6
L5:
	;
	v12 = v3
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = v13
	goto L9
L8:
	;
	v15 = int32(0)
	goto L9
L9:
	;
	if v15 <= v12 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v19 = v3
	goto L13
L11:
	;
	goto L12
L12:
	;
	v55 = F_lappend(m, v6, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L22
	} else {
		goto L24
	}
L13:
	;
	v22 = int32(0)
	if v6 == v22 {
		v32 = v22
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v26 <= v19 {
		v32 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v32 = v28 + v19<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= v19 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if v32 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = v39 + v19<<(uint(int32(2))%32)
	if v42 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v47 = F_bms_add_members(m, v45, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v47
	v19 = v19 + int32(1)
	goto L13
L24:
	;
	v6 = v55
	goto L1
}
func F_mcv_get_match_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v770 int32
	_ = v770
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	v23 = m.G0
	v25 = v23 + int32(-64)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v28 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v36 = F__emscripten_memset_bulkmem(m, v28, base.I32_extend8_s(l4^int32(1)), v34)
	mBase = m.M
	goto L3
L3:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L207
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L204
	}
L6:
	;
	m.G0 = v25 - int32(-64)
	return v36
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = l3 + int32(48)
	v58 = int32(0)
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v58<<(uint(int32(2))%32))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 == int32(318) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L6
L11:
	;
	v898 = v58 + int32(1)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v898 < v899 {
		v58 = v898
		goto L9
	} else {
		goto L203
	}
L12:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252 != int32(52) {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v80 = v70
	v81 = v71
	goto L15
L15:
	;
	if v81 != int32(17) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v250 = int32(1)
	v251 = int32(0)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = v74
	v81 = v79
	goto L15
L19:
	;
	v250 = int32(0)
	v251 = v80
	goto L12
L20:
	;
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v86 = F_get_opcode(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_fmgr_info(m, v86, v23+int32(-28))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v94 = v23 + int32(-32)
	v96 = v23 + int32(-36)
	v98 = v23 + int32(-52)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v106 == int32(27) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v135 != 0 {
		goto L44
	} else {
		goto L45
	}
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v110 = v109
	goto L27
L26:
	;
	v110 = v105
	goto L27
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v111 == int32(27) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = v114
	v117 = v115
	goto L30
L29:
	;
	v116 = v104
	v117 = v111
	goto L30
L30:
	;
	if v117 == int32(7) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L24
L32:
	;
	if v94 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v123 = v116
	v124 = v110
	goto L32
L34:
	;
	goto L35
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v120 != int32(7) {
		v135 = int32(0)
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v123 = v110
	v124 = v116
	goto L32
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v124
	goto L39
L38:
	;
	goto L39
L39:
	;
	if v96 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v123
	goto L42
L41:
	;
	goto L42
L42:
	;
	v127 = int32(1)
	if v98 == int32(0) {
		v135 = v127
		goto L31
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(base.B2i32(v117 == int32(7)))
	v135 = v127
	goto L31
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v139 = F_mcv_match_expression(m, v136, l1, l2, v23+int32(-40))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L68
	}
L47:
	;
	v141 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v142 == v141 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v150 = v141
	goto L49
L49:
	;
	v170 = v43 + v150*int32(24)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v139))))
	if v173 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	goto L11
L51:
	;
	v234 = v150 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v234) < base.Ui32(v235) {
		v150 = v234
		goto L49
	} else {
		goto L67
	}
L52:
	;
	v227 = v225 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v150+v36))) = uint8(v227)
	goto L51
L53:
	;
	v186 = v150 + v36
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if l4 == v187 {
		goto L51
	} else {
		goto L59
	}
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+24)))
	if v177 != int32(1) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v181 = int32(0)
	if l4 == v181 {
		v225 = v181
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v36))))
	v225 = v185
	goto L52
L59:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	if v190 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if l4 != 0 {
		v225 = v214 | base.B2i32(v213 != int32(0))
		goto L52
	} else {
		goto L66
	}
L61:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v139<<(uint(int32(2))%32))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	v201 = F_FunctionCall2Coll(m, v23+int32(-28), v189, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v139<<(uint(int32(2))%32))))
	v211 = F_FunctionCall2Coll(m, v23+int32(-28), v189, v205, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v213 = v201
	goto L60
L65:
	;
	v213 = v211
	goto L60
L66:
	;
	v225 = v214 & base.B2i32(v213 != int32(0))
	goto L52
L67:
	;
	goto L50
L68:
	;
	F_errmsg_internal(m, int32(358529), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(492918), int32(1647), int32(238290))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	if v250 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L72:
	;
	if v252 != int32(20) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v548 = F_mcv_match_expression(m, v546, l1, l2, int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L131
	}
L75:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v258 = F_get_opcode(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_fmgr_info(m, v258, v23+int32(-28))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v266 = v23 + int32(-32)
	v268 = v23 + int32(-36)
	v270 = v23 + int32(-41)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	if v278 == int32(27) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v307 == int32(0) {
		goto L5
	} else {
		goto L98
	}
L79:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v282 = v281
	goto L81
L80:
	;
	v282 = v277
	goto L81
L81:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v283 == int32(27) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = v286
	v289 = v287
	goto L84
L83:
	;
	v288 = v276
	v289 = v283
	goto L84
L84:
	;
	if v289 == int32(7) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L78
L86:
	;
	if v266 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v295 = v288
	v296 = v282
	goto L86
L88:
	;
	goto L89
L89:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v292 != int32(7) {
		v307 = int32(0)
		goto L85
	} else {
		goto L90
	}
L90:
	;
	v295 = v282
	v296 = v288
	goto L86
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v296
	goto L93
L92:
	;
	goto L93
L93:
	;
	if v268 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v295
	goto L96
L95:
	;
	goto L96
L96:
	;
	v299 = int32(1)
	if v270 == int32(0) {
		v307 = v299
		goto L85
	} else {
		goto L97
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(base.B2i32(v289 == int32(7)))
	v307 = v299
	goto L85
L98:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+23)))
	if v310 == int32(0) {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+24)))
	if v314 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	v318 = F_pg_detoast_datum(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v349 = F_mcv_match_expression(m, v346, l1, l2, v23+int32(-40))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	F_get_typlenbyvalalign(m, v320, v23+int32(-44), v23+int32(-45), v23+int32(-46))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v330 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+20)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+19)))
	v332 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+18)))
	F_deconstruct_array(m, v318, v330, v331, v332, v23+int32(-56), v23+int32(-60), v23+int32(-52))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v351 == int32(0) {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	v361 = int32(0)
	goto L108
L108:
	;
	v380 = v43 + v361*int32(24)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+v349))))
	if v383 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	goto L11
L110:
	;
	v543 = v361 + int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v543) < base.Ui32(v544) {
		v361 = v543
		goto L108
	} else {
		goto L130
	}
L111:
	;
	v518 = v516 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v361+v36))) = uint8(v518)
	goto L110
L112:
	;
	v395 = v361 + v36
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if l4 == v396 {
		goto L110
	} else {
		goto L118
	}
L113:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+24)))
	if v387 != int32(1) {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v390 = int32(0)
	if l4 == v390 {
		v516 = v390
		goto L111
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v36))))
	v516 = v394
	goto L111
L118:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+20)))
	v400 = v398 ^ int32(1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v401 <= int32(0) {
		v478 = v400
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if l4 != 0 {
		v516 = v491 | v478
		goto L111
	} else {
		goto L129
	}
L120:
	;
	v410 = int32(0)
	v414 = v400
	v415 = v398
	goto L121
L121:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428+v410))))
	if v430 == int32(1) {
		v462 = v415
		v464 = v414 & v415
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v478 = v464
	goto L119
L123:
	;
	v466 = v410 + int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v466 < v467 {
		v410 = v466
		v414 = v464
		v415 = v462
		goto L121
	} else {
		goto L128
	}
L124:
	;
	v433 = int32(1)
	if v414&v433 == v415&v433 {
		v478 = v414
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v380)+20))
	v442 = int32(2)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441+v349<<(uint(v442)%32))))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v410<<(uint(v442)%32))))
	v451 = F_FunctionCall2Coll(m, v23+int32(-28), v440, v445, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+20)))
	if v456 == int32(1) {
		v462 = v456
		v464 = v414 | base.B2i32(v451 != int32(0))
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v462 = v456
	v464 = v414 & base.B2i32(v451 != int32(0))
	goto L123
L128:
	;
	goto L122
L129:
	;
	v516 = v491 & v478
	goto L111
L130:
	;
	goto L109
L131:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v550 == int32(0) {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v559 = int32(0)
	goto L133
L133:
	;
	v578 = v43 + v559*int32(24)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	switch v580 {
	case 0:
		goto L137
	case 1:
		goto L136
	default:
		v589 = int32(0)
		goto L135
	}
L134:
	;
	goto L11
L135:
	;
	v590 = v559 + v36
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if l4 != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v548))))
	v589 = v586 ^ int32(1)
	goto L135
L137:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581+v548))))
	v589 = v583
	goto L135
L138:
	;
	v594 = v589 | v591
	goto L140
L139:
	;
	v594 = v589 & v591
	goto L140
L140:
	;
	v595 = int32(1)
	v596 = v594 & v595
	*(*uint8)(unsafe.Add(mBase, uint32(v590))) = uint8(v596)
	v599 = v559 + v595
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v599) < base.Ui32(v600) {
		v559 = v599
		goto L133
	} else {
		goto L141
	}
L141:
	;
	goto L134
L142:
	;
	v817 = int32(0)
	v819 = F_mcv_match_expression(m, v251, l1, l2, v817)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L192
	}
L143:
	;
	v758 = int32(*(*int16)(unsafe.Add(mBase, uint32(v251)+8)))
	v759 = F_bms_member_index(m, l1, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L181
	}
L144:
	;
	switch v252 - int32(6) {
	case 0:
		goto L143
	default:
		goto L142
	case 15:
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v252 != int32(6) {
		goto L142
	} else {
		goto L180
	}
L147:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	switch v606 {
	case 0, 1:
		goto L149
	case 2:
		goto L148
	default:
		goto L142
	}
L148:
	;
	v681 = int32(0)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v684 = F_mcv_get_match_bitmap(m, v682, l1, l2, l3, v681)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L165
	}
L149:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v610 = F_mcv_get_match_bitmap(m, v607, l1, l2, l3, base.B2i32(v606 == int32(1)))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v613 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v619 = int32(0)
	goto L154
L152:
	;
	goto L153
L153:
	;
	F_pfree(m, v610)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L164
	}
L154:
	;
	v636 = v619 + v36
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	if l4 != 0 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	goto L153
L156:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v636))) = uint8(v651)
	v654 = v619 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v654) < base.Ui32(v655) {
		v619 = v654
		goto L154
	} else {
		goto L163
	}
L157:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v610))))
	v651 = v650
	goto L156
L158:
	;
	v638 = int32(1)
	if v637&v638 == int32(0) {
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v643 = int32(0)
	if v637&int32(1) == v643 {
		v651 = v643
		goto L156
	} else {
		goto L162
	}
L161:
	;
	v651 = v638
	goto L156
L162:
	;
	goto L157
L163:
	;
	goto L155
L164:
	;
	goto L11
L165:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v686 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v692 = v681
	goto L169
L167:
	;
	goto L168
L168:
	;
	F_pfree(m, v684)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L179
	}
L169:
	;
	v709 = v692 + v36
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	if l4 != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L168
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v709))) = uint8(v726)
	v729 = v692 + int32(1)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v729) < base.Ui32(v730) {
		v692 = v729
		goto L169
	} else {
		goto L178
	}
L172:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v684))))
	v726 = base.B2i32(v723 == int32(0))
	goto L171
L173:
	;
	v711 = int32(1)
	if v710&v711 == int32(0) {
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v716 = int32(0)
	if v710&int32(1) == v716 {
		v726 = v716
		goto L171
	} else {
		goto L177
	}
L176:
	;
	v726 = v711
	goto L171
L177:
	;
	goto L172
L178:
	;
	goto L170
L179:
	;
	goto L11
L180:
	;
	goto L143
L181:
	;
	v761 = int32(0)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v762 == v761 {
		goto L11
	} else {
		goto L182
	}
L182:
	;
	v770 = v761
	goto L183
L183:
	;
	v787 = int32(0)
	v790 = v43 + v770*int32(24)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)+16))
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791+v759))))
	if v793 == v787 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L11
L185:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v790)+20))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796+v759<<(uint(int32(2))%32))))
	v803 = base.B2i32(v800 != int32(0))
	goto L187
L186:
	;
	v803 = v787
	goto L187
L187:
	;
	v804 = v770 + v36
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	if l4 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v808 = v803 | v805
	goto L190
L189:
	;
	v808 = v803 & v805
	goto L190
L190:
	;
	v809 = int32(1)
	v810 = v808 & v809
	*(*uint8)(unsafe.Add(mBase, uint32(v804))) = uint8(v810)
	v813 = v770 + v809
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v813) < base.Ui32(v814) {
		v770 = v813
		goto L183
	} else {
		goto L191
	}
L191:
	;
	goto L184
L192:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v821 == int32(0) {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	v829 = v817
	goto L194
L194:
	;
	v846 = int32(0)
	v849 = v43 + v829*int32(24)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)+16))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+v819))))
	if v852 == v846 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L11
L196:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v855+v819<<(uint(int32(2))%32))))
	v862 = base.B2i32(v859 != int32(0))
	goto L198
L197:
	;
	v862 = v846
	goto L198
L198:
	;
	v863 = v829 + v36
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	if l4 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v867 = v862 | v864
	goto L201
L200:
	;
	v867 = v862 & v864
	goto L201
L201:
	;
	v868 = int32(1)
	v869 = v867 & v868
	*(*uint8)(unsafe.Add(mBase, uint32(v863))) = uint8(v869)
	v872 = v829 + v868
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(v872) < base.Ui32(v873) {
		v829 = v872
		goto L194
	} else {
		goto L202
	}
L202:
	;
	goto L195
L203:
	;
	goto L10
L204:
	;
	F_errmsg_internal(m, int32(358529), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(492918), int32(1735), int32(238290))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errmsg_internal(m, int32(358529), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(492918), int32(1739), int32(238290))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
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
	F_errmsg_internal(m, int32(110739), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(492918), int32(1550), int32(269259))
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
	F_errmsg_internal(m, int32(110697), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(492918), int32(1573), int32(269259))
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
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v5 = int32(0)
	if l0 <= v5 {
	} else {
		v11 = int32(1)
		if l0 != v11 {
			v21 = v5
			v22 = v5
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
				v50 = v22 + v25
				if v50 != l0&int32(2147483646) {
					v21 = v48
					v22 = v50
					continue
				} else {
					break
				}
				break
			}
			v56 = v48
		} else {
			v56 = v5
		}
		if l0&v11 == int32(0) {
		} else {
			v63 = v56 << (uint(int32(2)) % 32)
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[819]))
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
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
func F_mdreadv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v277 int32
	_ = v277
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int64
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1072)
	m.G0 = v20
	if l4 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(1072)
	return
L2:
	;
	v26 = F__mdfd_getseg(m, l0, l1, l2, int32(0), int32(9))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v31 = l2 & int32(131071)
	v32 = int32(131072) - v31
	if base.Ui32(l4) < base.Ui32(v32) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = l4
	goto L7
L6:
	;
	v34 = v32
	goto L7
L7:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = int32(128)
	goto L10
L9:
	;
	v37 = v34
	goto L10
L10:
	;
	if l4 == v37 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = base.I64_extend_i32_u(v31 << (uint(int32(13)) % 32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v42
	v46 = int32(1)
	if base.Ui32(v34) < base.Ui32(int32(2)) {
		v158 = v46
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L87
	}
L14:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v179 = F_FileReadV(m, v175, v20+int32(48), v158, v41, int32(167772181))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L35
	}
L15:
	;
	v49 = int32(1)
	v51 = l4 - v49
	if l4 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v51&v49 == int32(0) {
		v158 = v124
		goto L14
	} else {
		goto L31
	}
L17:
	;
	v124 = v46
	v125 = v20 + int32(48)
	v130 = v42
	v131 = v49
	goto L16
L18:
	;
	goto L19
L19:
	;
	v64 = v46
	v65 = v20 + int32(48)
	v70 = v42
	v71 = v49
	v75 = v6
	goto L20
L20:
	;
	v82 = v71 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3+v82)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v84 == v70+v85 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v124 = v116
	v125 = v117
	v130 = v118
	v131 = v120
	goto L16
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)+v82)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 != v100+v103 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v85 - int32(-8192)
	v98 = v64
	v99 = v65
	v100 = v70
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(8192)
	v94 = v65 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v84
	v98 = v64 + int32(1)
	v99 = v94
	v100 = v84
	goto L22
L26:
	;
	v119 = int32(2)
	v120 = v71 + v119
	v122 = v75 + v119
	if v122 != v51&int32(-2) {
		v64 = v116
		v65 = v117
		v70 = v118
		v71 = v120
		v75 = v122
		goto L20
	} else {
		goto L30
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(8192)
	v109 = v99 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v102
	v116 = v98 + int32(1)
	v117 = v109
	v118 = v102
	goto L26
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v103 - int32(-8192)
	v116 = v98
	v117 = v99
	v118 = v100
	goto L26
L30:
	;
	goto L21
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l3+v131<<(uint(int32(2))%32))))
	if v130+v143 != v148 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v148
	v158 = v124 + int32(1)
	goto L14
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v143 - int32(-8192)
	v158 = v124
	goto L14
L35:
	;
	if int32(0) <= v179 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v184 = l4 << (uint(int32(13)) % 32)
	v186 = v158
	v187 = v179
	v192 = int32(0)
	v202 = v41
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L3
	} else {
		goto L82
	}
L39:
	;
	if v187 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, _consts[729])))
	if v206 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v354 = v187 + v192
	if v354 == v184 {
		goto L1
	} else {
		goto L70
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L65
	}
L45:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v210 != int32(1) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v214 = int32(base.Ui32(v192) >> (uint(int32(13)) % 32))
	if base.Ui32(l4) <= base.Ui32(v214) {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v218 = (l4 - v214) & int32(3)
	if v218 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v220 = int32(0)
	v221 = v214
	goto L53
L51:
	;
	v251 = v214
	goto L52
L52:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v214-l4) {
		goto L1
	} else {
		goto L57
	}
L53:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3+v221<<(uint(int32(2))%32))))
	v244 = F__emscripten_memset_bulkmem(m, v240, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L55
L54:
	;
	v251 = v246
	goto L52
L55:
	;
	v245 = int32(1)
	v246 = v221 + v245
	v248 = v220 + v245
	if v248 != v218 {
		v220 = v248
		v221 = v246
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v277 = v251
	goto L58
L58:
	;
	v294 = v277 << (uint(int32(2)) % 32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l3+v294)))
	v300 = F__emscripten_memset_bulkmem(m, v296, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L60
L59:
	;
	goto L1
L60:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v294+(l3+int32(4)))))
	v306 = F__emscripten_memset_bulkmem(m, v302, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L61
L61:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v294+(l3+int32(8)))))
	v312 = F__emscripten_memset_bulkmem(m, v308, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L62
L62:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v294+(l3+int32(12)))))
	v318 = F__emscripten_memset_bulkmem(m, v314, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L63
L63:
	;
	v320 = v277 + int32(4)
	if v320 != l4 {
		v277 = v320
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v331 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331+v329*int32(48))+32))
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = l2 + l4 - int32(1)
	F_errmsg(m, int32(159063), v20+int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(500109), int32(962), int32(35799))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v357 = v20 + int32(48)
	v362 = v357
	v363 = v186
	v364 = v187
	goto L73
L71:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v392 = v202 + base.I64_extend_i32_u(v187)
	v394 = F_FileReadV(m, v388, v20+int32(48), v387, v392, int32(167772181))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L80
	}
L72:
	;
	if v357 != v362 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if base.Ui32(v364) < base.Ui32(v366) {
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v387 = int32(0)
	goto L71
L75:
	;
	v372 = v363 - int32(1)
	if v372 != 0 {
		v362 = v362 + int32(8)
		v363 = v372
		v364 = v364 - v366
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v377 = F_memmove(m, v357, v362, v363<<(uint(int32(3))%32))
	mBase = m.M
	goto L79
L78:
	;
	goto L79
L79:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v378 + v364
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v357)+4)) = v381 - v364
	v387 = v363
	goto L71
L80:
	;
	if int32(0) <= v394 {
		v186 = v387
		v187 = v394
		v192 = v354
		v202 = v392
		goto L39
	} else {
		goto L81
	}
L81:
	;
	goto L40
L82:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v423 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423+v421*int32(48))+32))
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2 + l4 - int32(1)
	F_errmsg(m, int32(298841), v20)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(500109), int32(913), int32(35799))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L3
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
	F_errmsg_internal(m, int32(18247), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(500109), int32(875), int32(35799))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mdtruncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L52
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L47
	}
L3:
	;
	m.G0 = v15 + int32(112)
	return
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v19 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l2 == l3 {
		goto L3
	} else {
		goto L13
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v15+int32(40), v26, v27, v28, v29, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v15 + int32(40)
	F_errmsg(m, int32(31076), v15)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(500109), int32(1305), int32(357741))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v48 = l0 + l1<<(uint(int32(2))%32)
	v50 = v48 + int32(40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 <= int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v55 = v48 + int32(56)
	v67 = base.I64_extend_i32_u(v51)
	goto L15
L15:
	;
	v70 = v67 - int64(1)
	v71 = base.I32_wrap_i64(v70)
	v73 = v71 << (uint(int32(3)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v75 = v73 + v74
	v77 = v71 << (uint(int32(17)) % 32)
	if base.Ui32(l3) < base.Ui32(v77) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L3
L17:
	;
	if base.Ui64(int64(1)) < base.Ui64(v67) {
		v67 = v70
		goto L15
	} else {
		goto L46
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v82 = F_FileTruncate(m, v79, int64(0), int32(167772183))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(v77+int32(131072)) <= base.Ui32(l3) {
		goto L3
	} else {
		goto L41
	}
L21:
	;
	if v82 < int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v86 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_register_dirty_segment(m, l0, l1, v75)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	F_FileClose(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v70 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
	goto L17
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v114
	goto L28
L30:
	;
	if v94 <= int32(0) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v94 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	F_pfree(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v114 = int32(0)
	goto L29
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[819]))
	v107 = F_MemoryContextAlloc(m, v106, v73)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v70 <= base.I64_extend_i32_s(v94) {
		goto L28
	} else {
		goto L39
	}
L38:
	;
	v114 = v107
	goto L29
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v112 = F_repalloc(m, v111, v73)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v114 = v112
	goto L29
L41:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v126 = F_FileTruncate(m, v120, base.I64_extend_i32_u(l3-v77)<<(uint(int64(13))%64), int32(167772183))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	if v126 < int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v130 != int32(-1) {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	F_register_dirty_segment(m, l0, l1, v75)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	goto L17
L46:
	;
	goto L16
L47:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v161 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v159*int32(48))+32))
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v165
	F_errmsg(m, int32(299629), v15+int32(16))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(500109), int32(1333), int32(357741))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v185 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v183*int32(48))+32))
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v189
	F_errmsg(m, int32(293459), v15+int32(32))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(500109), int32(1360), int32(357741))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mdzeroextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v248 int32
	_ = v248
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int64
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int64
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v490 int32
	_ = v490
	var v505 int32
	_ = v505
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	if base.Ui64(base.I64_extend_i32_s(l3)+base.I64_extend_i32_u(l2)) <= base.Ui64(int64(4294967294)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L11
	} else {
		goto L141
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L11
	} else {
		goto L138
	}
L3:
	;
	if int32(0) < l3 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L133
	}
L6:
	;
	v27 = l2
	v28 = l3
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v16 + int32(128)
	return
L9:
	;
	v39 = v27 & int32(131071)
	v42 = base.I64_extend_i32_u(v39 << (uint(int32(13)) % 32))
	v44 = F__mdfd_getseg(m, l0, l1, v27, l4, int32(4))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v46 = int32(131072)
	if base.Ui32(v46) < base.Ui32(v28+v39) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if l4 != 0 {
		goto L128
	} else {
		goto L129
	}
L14:
	;
	v348 = base.I64_extend_i32_u(v51 << (uint(int32(13)) % 32))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v350 = F_FileAccess(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L83
	}
L15:
	;
	v51 = v46 - v39
	goto L17
L16:
	;
	v51 = v28
	goto L17
L17:
	;
	if base.Ui32(v51) < base.Ui32(int32(9)) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[820]))
	if v55 == int32(1) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v55 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v60 = base.I64_extend_i32_u(v51 << (uint(int32(13)) % 32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v62 = F_FileAccess(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	if v307 == int32(0) {
		goto L13
	} else {
		goto L74
	}
L22:
	;
	v307 = int32(-1)
	goto L21
L23:
	;
	if v62 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	v81 = int32(4122268)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(167772177)
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v61*int32(48))))
	v91 = m.Env.X__syscall_fallocate(m, v89, v85, v42, v60)
	mBase = m.M
	v92 = v85 - v91
	v94 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v85
	if v92 == int32(27) {
		goto L25
	} else {
		goto L27
	}
L26:
	;
	if v92 == int32(0) {
		v307 = v92
		goto L21
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v92
	if base.B2i32(v92 != int32(138))&base.B2i32(v92 != int32(28)) != 0 {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v108 = F_FileAccess(m, v61)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	if v108 < int32(0) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(167772177)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v61*int32(48))))
	v131 = m.G0
	v133 = v131 - int32(1024)
	m.G0 = v133
	v136 = base.I32_wrap_i64(v60)
	v137 = v42
	v139 = int32(0)
	goto L33
L32:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v276
	if v263 < v276 {
		goto L22
	} else {
		goto L71
	}
L33:
	;
	v146 = int32(0)
	if v136 == v146 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	m.G0 = v133 + int32(1024)
	goto L32
L35:
	;
	goto L34
L36:
	;
	v263 = v139
	goto L35
L37:
	;
	goto L38
L38:
	;
	v150 = v136
	v152 = v146
	goto L39
L39:
	;
	v162 = v133 + v152<<(uint(int32(3))%32)
	v163 = int32(8192)
	if base.Ui32(v163) <= base.Ui32(v150) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v177 = m.G0
	v179 = v177 - int32(1024)
	m.G0 = v179
	if v171 <= int32(128) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	goto L40
L42:
	;
	v166 = v163
	goto L44
L43:
	;
	v166 = v150
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(1847296)
	v171 = v152 + int32(1)
	v172 = v150 - v166
	if base.Ui32(int32(126)) < base.Ui32(v152) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	if v172 != 0 {
		v150 = v172
		v152 = v171
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	m.G0 = v179 + int32(1024)
	if int32(0) <= v248 {
		v136 = v172
		v137 = v137 + base.I64_extend_i32_u(v248)
		v139 = v139 + v248
		goto L33
	} else {
		goto L70
	}
L48:
	;
	v186 = v133
	v188 = v171
	v190 = int32(0)
	v193 = v137
	goto L51
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	v248 = int32(-1)
	goto L47
L51:
	;
	if v188 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v248 = v204
	goto L47
L53:
	;
	if v200 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v198 = F_pwrite(m, v121, v196, v197, v193)
	mBase = m.M
	v200 = v198
	goto L53
L55:
	;
	goto L56
L56:
	;
	v199 = F_pwritev(m, v121, v186, v188, v193)
	mBase = m.M
	v200 = v199
	goto L53
L57:
	;
	v248 = int32(-1)
	goto L47
L58:
	;
	goto L59
L59:
	;
	v204 = v200 + v190
	v210 = v186
	v212 = v188
	v213 = v200
	goto L60
L60:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if base.Ui32(v218) <= base.Ui32(v213) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v210 != v179 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v224 = v212 - int32(1)
	if v224 != 0 {
		v210 = v210 + int32(8)
		v212 = v224
		v213 = v213 - v218
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L61
L65:
	;
	v248 = v204
	goto L47
L66:
	;
	v228 = F_memmove(m, v179, v210, v212<<(uint(int32(3))%32))
	mBase = m.M
	goto L68
L67:
	;
	goto L68
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v229 + v213
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v232 - v213
	if int32(0) < v212 {
		v186 = v179
		v188 = v212
		v190 = v204
		v193 = v193 + base.I64_extend_i32_u(v200)
		goto L51
	} else {
		goto L69
	}
L69:
	;
	goto L52
L70:
	;
	v263 = v248
	goto L35
L71:
	;
	if v60 == base.I64_extend_i32_u(v263) {
		v307 = int32(0)
		goto L21
	} else {
		goto L72
	}
L72:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v284 != 0 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(51)
	goto L22
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v325 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325+v323*int32(48))+32))
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v329
	F_errmsg(m, int32(296400), v16+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_errhint(m, int32(644780), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(500109), int32(619), int32(426005))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	if v533 < int32(0) {
		goto L1
	} else {
		goto L127
	}
L82:
	;
	v533 = int32(-1)
	goto L81
L83:
	;
	if v350 < int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = int32(167772177)
	v359 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v349*int32(48))))
	v373 = m.G0
	v375 = v373 - int32(1024)
	m.G0 = v375
	v378 = base.I32_wrap_i64(v348)
	v379 = v42
	v381 = int32(0)
	goto L86
L85:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v518 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v518
	if v505 < v518 {
		goto L82
	} else {
		goto L124
	}
L86:
	;
	v388 = int32(0)
	if v378 == v388 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	m.G0 = v375 + int32(1024)
	goto L85
L88:
	;
	goto L87
L89:
	;
	v505 = v381
	goto L88
L90:
	;
	goto L91
L91:
	;
	v392 = v378
	v394 = v388
	goto L92
L92:
	;
	v404 = v375 + v394<<(uint(int32(3))%32)
	v405 = int32(8192)
	if base.Ui32(v405) <= base.Ui32(v392) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v419 = m.G0
	v421 = v419 - int32(1024)
	m.G0 = v421
	if v413 <= int32(128) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	goto L93
L95:
	;
	v408 = v405
	goto L97
L96:
	;
	v408 = v392
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+4)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = int32(1847296)
	v413 = v394 + int32(1)
	v414 = v392 - v408
	if base.Ui32(int32(126)) < base.Ui32(v394) {
		goto L94
	} else {
		goto L98
	}
L98:
	;
	if v414 != 0 {
		v392 = v414
		v394 = v413
		goto L92
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	m.G0 = v421 + int32(1024)
	if int32(0) <= v490 {
		v378 = v414
		v379 = v379 + base.I64_extend_i32_u(v490)
		v381 = v381 + v490
		goto L86
	} else {
		goto L123
	}
L101:
	;
	v428 = v375
	v430 = v413
	v432 = int32(0)
	v435 = v379
	goto L104
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	v490 = int32(-1)
	goto L100
L104:
	;
	if v430 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v490 = v446
	goto L100
L106:
	;
	if v442 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v440 = F_pwrite(m, v363, v438, v439, v435)
	mBase = m.M
	v442 = v440
	goto L106
L108:
	;
	goto L109
L109:
	;
	v441 = F_pwritev(m, v363, v428, v430, v435)
	mBase = m.M
	v442 = v441
	goto L106
L110:
	;
	v490 = int32(-1)
	goto L100
L111:
	;
	goto L112
L112:
	;
	v446 = v442 + v432
	v452 = v428
	v454 = v430
	v455 = v442
	goto L113
L113:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if base.Ui32(v460) <= base.Ui32(v455) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v452 != v421 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v466 = v454 - int32(1)
	if v466 != 0 {
		v452 = v452 + int32(8)
		v454 = v466
		v455 = v455 - v460
		goto L113
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	goto L114
L118:
	;
	v490 = v446
	goto L100
L119:
	;
	v470 = F_memmove(m, v421, v452, v454<<(uint(int32(3))%32))
	mBase = m.M
	goto L121
L120:
	;
	goto L121
L121:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v471 + v455
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+4)) = v474 - v455
	if int32(0) < v454 {
		v428 = v421
		v430 = v454
		v432 = v446
		v435 = v435 + base.I64_extend_i32_u(v442)
		goto L104
	} else {
		goto L122
	}
L122:
	;
	goto L105
L123:
	;
	v505 = v490
	goto L88
L124:
	;
	if v348 == base.I64_extend_i32_u(v505) {
		v533 = int32(0)
		goto L81
	} else {
		goto L125
	}
L125:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v526 != 0 {
		goto L82
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(51)
	goto L82
L127:
	;
	goto L13
L128:
	;
	v555 = v28 - v51
	if int32(0) < v555 {
		v27 = v27 + v51
		v28 = v555
		goto L9
	} else {
		goto L132
	}
L129:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v549 != int32(-1) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	F_register_dirty_segment(m, l0, l1, v44)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	goto L10
L133:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L11
	} else {
		goto L134
	}
L134:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v16+int32(56), v583, v584, v585, v586, l1)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v16 + int32(56)
	F_errmsg(m, int32(153967), v16)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(500109), int32(566), int32(426005))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L11
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _consts[820]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v607
	F_errmsg_internal(m, int32(487469), v16+int32(32))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L11
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(500109), int32(611), int32(426005))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L142
	}
L142:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v627 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v627+v625*int32(48))+32))
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v631
	F_errmsg(m, int32(299807), v16+int32(48))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	F_errhint(m, int32(644780), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(500109), int32(641), int32(426005))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_member_can_set_role(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v4 = int32(1)
	if l0 == l1 {
		v54 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v54
L2:
	;
	v6 = F_superuser_arg(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v6 != 0 {
		v54 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v11 = int32(0)
	v13 = F_roles_is_member_of(m, l0, int32(2), v11, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = int32(0)
	if v13 == v15 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v53
	goto L1
L8:
	;
	v53 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 <= int32(0) {
		v46 = v15
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v46
	goto L7
L12:
	;
	v24 = int32(0)
	if v24 < v21 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v27 = v21
	goto L15
L14:
	;
	v27 = v24
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v30 = int32(0)
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
	v39 = base.B2i32(v38 == l1)
	if v38 == l1 {
		v46 = v39
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v46 = v39
	goto L11
L18:
	;
	v41 = v30 + int32(1)
	if v41 != v27 {
		v30 = v41
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v100)
	return v96 - l0
L2:
	;
	v91 = l1
	v96 = l0
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
		v91 = v14
		v96 = v19
		goto L1
	} else {
		goto L34
	}
L7:
	;
	if l5 != 0 {
		v91 = v14
		v96 = v19
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
	v76 = v32
	v77 = int32(-1)
	v78 = int32(1)
	goto L16
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.Ui32((v39+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v62 = int32(2)
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v76)
	v81 = v14 + int32(1)
	v82 = v19 + v78
	v83 = v15 + v77
	if int32(0) < v83 {
		v14 = v81
		v15 = v83
		v19 = v82
		goto L5
	} else {
		goto L33
	}
L17:
	;
	if v15 < v62 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	goto L17
L19:
	;
	v46 = int32(3)
	v48 = v39 & int32(254)
	if v48 == int32(154) {
		v62 = v46
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v39+int32(112))&int32(255)) < base.Ui32(int32(10)) {
		v62 = v46
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v48 == int32(156) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(4)
	goto L24
L23:
	;
	v61 = int32(1)
	goto L24
L24:
	;
	v62 = v61
	goto L18
L25:
	;
	if l5 != 0 {
		v91 = v14
		v96 = v19
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if l3 != v23 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	F_report_invalid_encoding(m, int32(7), v19, v15)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
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
	if v62 != int32(2) {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+1)))
	if int32(0) <= v71 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v76 = v71
	v77 = int32(-2)
	v78 = int32(2)
	goto L16
L33:
	;
	v91 = v81
	v96 = v82
	goto L1
L34:
	;
	F_report_untranslatable_char(m, int32(7), l4, v19, v15)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mic2latin_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v113)
	return v107 - l0
L2:
	;
	v103 = l1
	v107 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v17 = l1
	v18 = l2
	v21 = l0
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		v103 = v17
		v107 = v21
		goto L1
	} else {
		goto L35
	}
L7:
	;
	if l6 != 0 {
		v103 = v17
		v107 = v21
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v36 = base.I32_extend8_s(v27)
	if int32(0) <= v36 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_report_invalid_encoding(m, int32(7), v21, v18)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v87 = v36
	v88 = int32(-1)
	v89 = int32(1)
	goto L16
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32((v43+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v66 = int32(2)
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v87)
	v92 = v17 + int32(1)
	v93 = v21 + v89
	v94 = v18 + v88
	if int32(0) < v94 {
		v17 = v92
		v18 = v94
		v21 = v93
		goto L5
	} else {
		goto L34
	}
L17:
	;
	if v18 < v66 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	goto L17
L19:
	;
	v50 = int32(3)
	v52 = v43 & int32(254)
	if v52 == int32(154) {
		v66 = v50
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v43+int32(112))&int32(255)) < base.Ui32(int32(10)) {
		v66 = v50
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v52 == int32(156) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(4)
	goto L24
L23:
	;
	v65 = int32(1)
	goto L24
L24:
	;
	v66 = v65
	goto L18
L25:
	;
	if l6 != 0 {
		v103 = v17
		v107 = v21
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if l3 != v27 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	F_report_invalid_encoding(m, int32(7), v21, v18)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
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
	if v66 != int32(2) {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+1)))
	if int32(0) <= v75 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5-int32(128)+v75&int32(255)))))
	if v81 == int32(0) {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v87 = v81
	v88 = int32(-2)
	v89 = int32(2)
	goto L16
L34:
	;
	v103 = v92
	v107 = v93
	goto L1
L35:
	;
	F_report_untranslatable_char(m, int32(7), l4, v21, v18)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
func F_movedb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v192 int32
	_ = v192
	var v216 int32
	_ = v216
	var v243 int32
	_ = v243
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v293 int32
	_ = v293
	var v316 int32
	_ = v316
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v395 int32
	_ = v395
	var v419 int32
	_ = v419
	var v444 int32
	_ = v444
	var v470 int32
	_ = v470
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v517 int32
	_ = v517
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v566 int32
	_ = v566
	var v593 int32
	_ = v593
	var v617 int32
	_ = v617
	var v642 int32
	_ = v642
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v694 int32
	_ = v694
	var v717 int32
	_ = v717
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v765 int32
	_ = v765
	var v789 int32
	_ = v789
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v841 int32
	_ = v841
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v938 int32
	_ = v938
	var v960 int64
	_ = v960
	var v961 int32
	_ = v961
	var v984 int32
	_ = v984
	var v1007 int32
	_ = v1007
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1065 int32
	_ = v1065
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1169 int32
	_ = v1169
	var v1191 int32
	_ = v1191
	var v1218 int32
	_ = v1218
	var v1245 int32
	_ = v1245
	var v1271 int32
	_ = v1271
	var v1296 int32
	_ = v1296
	var v1320 int32
	_ = v1320
	var v1348 int32
	_ = v1348
	var v1373 int32
	_ = v1373
	var v1399 int32
	_ = v1399
	var v1409 int32
	_ = v1409
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int64
	_ = v1506
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1567 int32
	_ = v1567
	var v1591 int32
	_ = v1591
	var v1615 int64
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1642 int32
	_ = v1642
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1719 int32
	_ = v1719
	var v1743 int32
	_ = v1743
	var v1768 int32
	_ = v1768
	var v1794 int32
	_ = v1794
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1870 int32
	_ = v1870
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1946 int32
	_ = v1946
	var v1970 int32
	_ = v1970
	var v1993 int32
	_ = v1993
	var v2018 int32
	_ = v2018
	var v2042 int32
	_ = v2042
	var v2069 int32
	_ = v2069
	var v2092 int32
	_ = v2092
	var v2115 int32
	_ = v2115
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2192 int32
	_ = v2192
	var v2218 int32
	_ = v2218
	var v2244 int32
	_ = v2244
	var v2268 int32
	_ = v2268
	var v2292 int32
	_ = v2292
	var v2316 int64
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2340 int32
	_ = v2340
	var v2363 int32
	_ = v2363
	var v2386 int32
	_ = v2386
	var v2414 int32
	_ = v2414
	var v2437 int32
	_ = v2437
	var v2460 int32
	_ = v2460
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2492 int64
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	v3 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(176)
	m.G0 = v33
	v39 = v3
	v40 = v3
	v41 = int32(-1)
	v42 = v3
	v43 = v3
	v44 = v3
	v45 = v3
	v46 = v3
	v47 = v3
	v48 = v3
	v49 = v3
	v50 = v3
	v51 = v3
	v52 = v3
	v53 = v3
	v54 = v3
	v55 = v3
	v56 = v3
	v57 = v3
	v58 = v3
	v59 = v3
	v60 = v3
	v61 = v3
	v62 = v33
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
	m.G0 = v33 + int32(176)
	return
L4:
	;
	goto L3
L5:
	;
	if v41 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v2491 = int32(m.ExcTag)
	v2492 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2491 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L8:
	;
	if v1489 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L9:
	;
	v1469 = v39
	v1470 = v40
	v1471 = v60
	v1472 = v42
	v1473 = v43
	v1474 = v44
	v1475 = v45
	v1476 = v46
	v1477 = v47
	v1478 = v48
	v1479 = v49
	v1480 = v50
	v1481 = v51
	v1482 = v52
	v1483 = v53
	v1484 = v54
	v1485 = v55
	v1486 = v56
	v1487 = v57
	v1488 = v58
	v1489 = v59
	v1491 = v61
	v1492 = v62
	goto L8
L10:
	;
	goto L11
L11:
	;
	v68 = int32(16)
	v69 = v62 - v68
	m.G0 = v69
	v72 = v69 - v68
	m.G0 = v72
	v75 = v72 - v68
	m.G0 = v75
	v78 = v75 - v68
	m.G0 = v78
	v81 = v78 - int32(48)
	m.G0 = v81
	v84 = v81 - v68
	m.G0 = v84
	v87 = v84 - int32(160)
	m.G0 = v87
	v90 = v87 - int32(80)
	m.G0 = v90
	v92 = int32(32)
	v93 = v90 - v92
	m.G0 = v93
	v96 = v93 - v92
	m.G0 = v96
	v99 = v96 - v68
	m.G0 = v99
	v102 = v99 - v68
	m.G0 = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v127 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v151 = int32(0)
	v164 = F_get_db_info(m, l0, int32(8), v69, v151, v151, v151, v151, v151, v151, v151, v78, v151, v151, v151, v151, v151, v151)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v164 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_LockSharedObjectForSession(m, v270)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errcode(m, int32(1283))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = l0
	F_errmsg(m, int32(72248), v33+int32(80))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2034), int32(503605))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v316 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v339 = F_object_ownercheck(m, int32(1262), v270, v316)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if v339 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_aclcheck_error(m, int32(2), int32(9), l0)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v369 == v270 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v493 = F_get_tablespace_oid(m, l1, int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errcode(m, int32(100663621))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errmsg(m, int32(362417), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2058), int32(503605))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v517 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v541 = F_object_aclcheck(m, int32(1213), v493, v517, int64(512))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v541 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_aclcheck_error(m, v541, int32(42), l1)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v493 == int32(1664) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v493 == v669 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errcode(m, int32(50856066))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errmsg(m, int32(419261), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2080), int32(503605))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L46
	}
L46:
	;
	goto L1
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_sequence_close(m, v127, int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v739 = F_CountOtherDBBackends(m, v270, v72, v75)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L52
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_UnlockSharedObjectForSession(m, v270)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L51
	}
L51:
	;
	goto L4
L52:
	;
	if v739 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v890 = F_GetDatabasePath(m, v270, v868)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L61
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errcode(m, int32(100663621))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l0
	F_errmsg(m, int32(133687), v33+int32(32))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	F_errdetail_busy_db(m, v818, v817)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2104), int32(503605))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L60
	}
L60:
	;
	goto L1
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v913 = F_GetDatabasePath(m, v270, v493)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_RequestCheckpoint(m, int32(60))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v960 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_WaitForProcSignalBarrier(m, v960)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_DropDatabaseBuffers(m, v270)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v1029 = F_AllocateDir(m, v913)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_before_shmem_exit(m, int32(545), v84)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L95
	}
L68:
	;
	if v1029 == int32(0) {
		v1409 = v45
		goto L67
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v1054 = F_ReadDir(m, v1029, v913)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L90
	}
L71:
	;
	if v1054 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v1065 = v45
	v1079 = v1054
	goto L75
L73:
	;
	v1126 = v45
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_FreeDir(m, v1029)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L85
	}
L75:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+19)))
	if v1086 != int32(46) {
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v1126 = v1115
	goto L74
L77:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+20)))
	if v1089 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+20)))
	if v1090 != int32(46) {
		goto L70
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v1115 = F_ReadDir(m, v1029, v913)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L83
	}
L81:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+21)))
	if v1093 != 0 {
		goto L70
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if v1115 != 0 {
		v1065 = v1115
		v1079 = v1115
		goto L75
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	v1191 = F_rmdir(m, v913)
	mBase = m.M
	if v1191 == int32(0) {
		v1409 = v1126
		goto L67
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v913
	F_errmsg_internal(m, int32(296876), v33+int32(48))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2177), int32(503605))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L89
	}
L89:
	;
	goto L1
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errcode(m, int32(325))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = l0
	F_errmsg(m, int32(722319), v33-int32(-64))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errhint(m, int32(645940), int32(0))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1065
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v72
	F_errfinish(m, int32(494648), int32(2166), int32(503605))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		v2487 = v102
		goto L7
	} else {
		goto L94
	}
L94:
	;
	goto L1
L95:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v1460 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v33 + int32(88)
	goto L99
L97:
	;
	v1469 = v99
	v1470 = v270
	v1471 = v84
	v1472 = v913
	v1473 = v493
	v1474 = v127
	v1475 = v1409
	v1476 = v78
	v1477 = v96
	v1478 = v102
	v1479 = v93
	v1480 = v890
	v1481 = v69
	v1482 = v72
	v1483 = v75
	v1484 = v81
	v1485 = v87
	v1486 = v1460
	v1487 = v1458
	v1488 = v90
	v1489 = int32(0)
	v1491 = v84
	v1492 = v102
	goto L8
L99:
	;
	goto L97
L100:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1485
	v1503 = F__emscripten_memset_bulkmem(m, v1488, base.I32_extend8_s(int32(0)), int32(72))
	mBase = m.M
	goto L103
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1487
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	F_cancel_before_shmem_exit(m, int32(545), v1471)
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L148
	}
L103:
	;
	v1504 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1479)+16)) = uint16(v1504)
	v1506 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1479))) = v1506
	*(*int64)(unsafe.Add(mBase, uint32(v1479)+8)) = v1506
	*(*int64)(unsafe.Add(mBase, uint32(v1477))) = v1506
	*(*int64)(unsafe.Add(mBase, uint32(v1477)+8)) = v1506
	*(*uint16)(unsafe.Add(mBase, uint32(v1477)+16)) = uint16(v1504)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	F_copydir(m, v1480, v1472, v1504)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1469)+8)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v1469)+4)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v1469))) = v1470
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1476)))
	*(*int32)(unsafe.Add(mBase, uint32(v1469)+12)) = v1543
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_XLogBeginInsert(m)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_XLogRegisterData(m, v1469, int32(16))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1615 = F_XLogInsert(m, int32(4), int32(1))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_ScanKeyInit(m, v1484, int32(2), int32(3), int32(62), l0)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1665 = int32(1)
	v1668 = F_systable_beginscan(m, v1474, int32(2671), v1665, int32(0), v1665, v1484)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1691 = F_systable_getnext(m, v1668)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L110
	}
L110:
	;
	if v1691 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1817 = v1691 + int32(4)
	F_LockTuple(m, v1474, v1817, int32(7))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L118
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_errcode(m, int32(1283))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
	F_errmsg(m, int32(72248), v33)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_errfinish(m, int32(494648), int32(2232), int32(503605))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L117
	}
L117:
	;
	goto L1
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+44)) = v1473
	v1822 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1477)+11)) = uint8(v1822)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1846 = F_heap_modify_tuple(m, v1691, v1824, v1503, v1479, v1477)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_CatalogTupleUpdate(m, v1474, v1817, v1846)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_UnlockTuple(m, v1474, v1817, int32(7))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L121
	}
L121:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v1896 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1919 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1262), v1470, v1919, v1919, v1919)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_systable_endscan(m, v1668)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_RequestCheckpoint(m, int32(44))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v1993 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[167])) = uint8(v1993)
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_sequence_close(m, v1474, int32(0))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_cancel_before_shmem_exit(m, int32(545), v1471)
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1487
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_StartTransactionCommand(m)
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v2137 = F_rmtree(m, v1480)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1478)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1478))) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_XLogBeginInsert(m)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L141
	}
L135:
	;
	if v2137 != 0 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v2162 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L137
	}
L137:
	;
	if v2162 == int32(0) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v1480
	F_errmsg(m, int32(692697), v33+int32(16))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_errfinish(m, int32(494648), int32(2296), int32(503605))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L140
	}
L140:
	;
	goto L134
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_XLogRegisterData(m, v1478, int32(8))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_XLogRegisterData(m, v1476, int32(4))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	v2316 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_UnlockSharedObjectForSession(m, v1470)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_pfree(m, v1480)
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_pfree(m, v1472)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L147
	}
L147:
	;
	goto L4
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_movedb_failure_callback(m, v33, v1471)
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v33)+132)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v33)+136)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1484
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v33)+172)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v1482
	F_pg_re_throw(m)
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		v2487 = v1492
		goto L7
	} else {
		goto L150
	}
L150:
	;
	goto L1
L151:
	;
	v2496 = int32(v2492)
	m.G0 = v2487
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v2496)+4))
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2496)))
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2499)))
	if v33+int32(88) == v2503 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	m.ExcPending = 1
	goto L160
L153:
	;
	if v2506 != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2499)+4))
	v2506 = v2505
	goto L156
L155:
	;
	v2506 = int32(0)
	goto L156
L156:
	;
	goto L153
L157:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v33)+172))
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(v33)+164))
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v33)+160))
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v33)+156))
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v33)+152))
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v33)+144))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v33)+140))
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v33)+136))
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v33)+108))
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v33)+104))
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v33)+100))
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
	v39 = v2517
	v40 = v2520
	v41 = v2506
	v42 = v2523
	v43 = v2521
	v44 = v2519
	v45 = v2524
	v46 = v2510
	v47 = v2516
	v48 = v2518
	v49 = v2515
	v50 = v2522
	v51 = v2507
	v52 = v2508
	v53 = v2509
	v54 = v2511
	v55 = v2513
	v56 = v2526
	v57 = v2527
	v58 = v2514
	v59 = v2498
	v60 = v2525
	v61 = v2512
	v62 = v2487
	goto L2
L158:
	;
	goto L159
L159:
	;
	F___wasm_longjmp(m, v2499, v2498)
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	return
L161:
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
