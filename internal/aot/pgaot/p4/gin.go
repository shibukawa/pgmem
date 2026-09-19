package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinInitBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitBuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitBuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v21 = int32(_a_F_GinInitBuffer_0)
	v23 = int32(0)
	if v23|(v20&int32(3)|int32(1)) == v23 {
		v39 = v20 + v21
		v41 = v20 + int32(4)
		if base.Ui32(v41) < base.Ui32(v39) {
			v43 = v39
		} else {
			v43 = v41
		}
		v48 = (v20^int32(-1)+v43)&int32(-4) + int32(4)
		if v48 == int32(0) {
		} else {
			base.MemoryFill(m, v20, int32(0), v48)
		}
	} else {
		base.MemoryFill(m, v20, int32(0), v21)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_GinInitBuffer_1)
	v62 = int32(_a_F_GinInitBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v62)
	v68 = int32(_a_F_GinInitBuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v72 = v20 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v2)
	return
}
func F_freeGinBtreeStack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ReleaseBuffer(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_pfree(m, v4)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	if v7 != 0 {
		v4 = v7
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
}
func F_ginBuildCallbackParallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	v18 = int32(_a_F_ginBuildCallbackParallel_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v21
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v32 = v28 | v29<<(uint(v15)%32)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[2]))))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[3]))))
	v37 = v33 | v34<<(uint(v15)%32)
	if base.Ui32(v32) < base.Ui32(v37) {
		v48 = int32(-1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v48 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	if base.Ui32(v37) < base.Ui32(v32) {
		v48 = int32(1)
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[4]))))
	if base.Ui32(v42) < base.Ui32(v43) {
		v48 = int32(-1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v48 = base.B2i32(base.Ui32(v43) < base.Ui32(v42))
	goto L2
L6:
	;
	F_ginFlushBuildState(m, l5, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[4]))) = uint16(v53)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[3]))) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if int32(0) < v58 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v68 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[5])))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[6])))
	if base.Ui32(v130<<(uint(int32(10))%32)) <= base.Ui32(v129) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v68))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2+v68<<(uint(int32(2))%32))))
	v83 = int32(_a_F_ginBuildCallbackParallel_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[7])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[8])))
	v90 = v68 + int32(1)
	v92 = v90 & int32(_a_F_ginBuildCallbackParallel_1)
	v97 = F_ginExtractEntries(m, v88, v92, v82, v78, v16+int32(8), v16+int32(12))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v84
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_ginInsertBAEntries(m, l5+int32(_a_F_ginBuildCallbackParallel_2), l1, v92, v97, v101, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[9])))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[9]))) = base.F64_add(v105, base.F64_convert_i32_s(v106))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[7])))
	F_MemoryContextReset(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v90 < v114 {
		v68 = v90
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	F_ginFlushBuildState(m, l5, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v19
	m.G0 = v16 + int32(16)
	return
L23:
	;
	goto L22
}
func F_ginCompressPostingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v118 int32
	_ = v118
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = l2 & int32(-2)
	v21 = F_palloc(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	v30 = v21 + int32(8)
	v31 = int32(1)
	if l1 < int32(2) {
		v188 = v30
		v194 = v31
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v200 = v188 - v30
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v200)
	v203 = v200 & int32(_a_F_ginCompressPostingList_0)
	if v203 != (v203+int32(1))&int32(_a_F_ginCompressPostingList_1) {
		goto L28
	} else {
		goto L29
	}
L4:
	;
	v34 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v35 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
	v38 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v46 = v30
	v48 = v34 | (v35<<(uint(int64(11))%64) | v38<<(uint(int64(27))%64))
	v52 = v31
	goto L5
L5:
	;
	v60 = l0 + v52*int32(6)
	v61 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)))
	v64 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	v69 = v61<<(uint(int64(11))%64) | v64<<(uint(int64(27))%64) | v68
	v70 = v69 - v48
	v71 = v20 + v21 - v46
	if int32(7) <= v71 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v188 = v182
	v194 = l1
	goto L3
L7:
	;
	v184 = v52 + int32(1)
	if v184 != l1 {
		v46 = v182
		v48 = v69
		v52 = v184
		goto L5
	} else {
		goto L27
	}
L8:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v70) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v118 = v17 + int32(9)
	if base.Ui64(int64(128)) <= base.Ui64(v70) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v78 = v46
	v80 = v70
	goto L14
L12:
	;
	v102 = v46
	v104 = v70
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v104)
	v182 = v102 + int32(1)
	goto L7
L14:
	;
	v92 = base.I32_wrap_i64(v80) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v92)
	v95 = v78 + int32(1)
	v99 = int64(base.Ui64(v80) >> (uint(int64(7)) % 64))
	if base.Ui64(int64(16383)) < base.Ui64(v80) {
		v78 = v95
		v80 = v99
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v102 = v95
	v104 = v99
	goto L13
L16:
	;
	goto L15
L17:
	;
	v125 = v70
	v130 = v118
	goto L20
L18:
	;
	v149 = v70
	v154 = v118
	goto L19
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v149)
	v162 = v154 - v17 - int32(8)
	if v71 < v162 {
		v188 = v46
		v194 = v52
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v137 = base.I32_wrap_i64(v125) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v137)
	v140 = v130 + int32(1)
	v144 = int64(base.Ui64(v125) >> (uint(int64(7)) % 64))
	if base.Ui64(int64(16383)) < base.Ui64(v125) {
		v125 = v144
		v130 = v140
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v149 = v144
	v154 = v140
	goto L19
L22:
	;
	goto L21
L23:
	;
	if v162 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	base.MemoryCopy(m, v46, v17+int32(9), v162)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v182 = v46 + v162
	goto L7
L27:
	;
	goto L6
L28:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v30))) = uint8(v210)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v194
	goto L33
L32:
	;
	goto L33
L33:
	;
	m.G0 = v17 + int32(16)
	return v21
}
func F_ginFindLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	v9 = F_palloc(m, int32(24))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_ReadBuffer(m, v15, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(1)
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_CheckForSerializableConflictIn(m, v21, int32(0), v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = v9
	goto L8
L7:
	;
	goto L6
L8:
	;
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)) = uint16(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 < v33 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v29
L10:
	;
	F_LockBuffer(m, v35, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L10
L12:
	;
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	if v35 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v75 = int32(1)
	if l1 != 0 {
		v102 = v75
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(v35^int32(-1))<<(uint(int32(2))%32))))
	v74 = v66
	goto L15
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v74 = v68 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	if l1 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+16)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+v76)+6)))
	if v78&int32(2) == int32(0) {
		v102 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v86 = int32(2)
	F_LockBuffer(m, v35, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+16)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v90)+6)))
	if v92&int32(2) != 0 {
		v102 = v86
		goto L19
	} else {
		goto L24
	}
L24:
	;
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v98 = int32(1)
	F_LockBuffer(m, v35, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v102 = v98
	goto L19
L27:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v112 != 0 {
		v171 = v53
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v103)+6)))
	if v105&int32(64) == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_ginFinishOldSplit(m, l0, v29, v102)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+16)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v176)+6)))
	if v178&int32(2) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v115 = v53
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v120 == v121 {
		v171 = v115
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v171 = v156
	goto L31
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v124 = m.T0[v123].(func(*base.Module, int32, int32) int32)(m, l0, v115)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v124 == int32(0) {
		v171 = v115
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+16)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v115+v128)))
	if v130 == int32(-1) {
		v171 = v115
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v135 = F_ginStepRight(m, v133, v134, v102)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v135
	if v135 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v135^int32(-1))<<(uint(int32(2))%32))))
	v156 = v148
	goto L40
L42:
	;
	goto L43
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v156 = v150 + v135<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L44:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v166 == int32(0) {
		v115 = v156
		goto L33
	} else {
		goto L48
	}
L45:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v157)+6)))
	if v159&int32(64) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	F_ginFinishOldSplit(m, l0, v29, v102)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L34
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v184 = m.T0[v183].(func(*base.Module, int32, int32) int32)(m, l0, v29)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L9
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	F_LockBuffer(m, v186, int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if l1 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v184
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v193 = F_ReleaseAndReadBuffer(m, v191, v192, v184)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v197 = F_palloc(m, int32(24))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v193
	goto L8
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v197)+20)) = v29
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v202 = F_ReadBuffer(m, v201, v184)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v202
	v29 = v197
	goto L8
}
func F_ginFinishOldSplit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15 + int32(4)
			F_errmsg_internal(m, int32(_a_F_ginFinishOldSplit_0), v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ginFinishOldSplit_1), int32(783), int32(_a_F_ginFinishOldSplit_2))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if l2 == int32(1) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_LockBuffer(m, v32, int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							F_LockBuffer(m, v36, int32(2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v40 < int32(0) {
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[0]))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v40^int32(-1))<<(uint(int32(2))%32))))
									v58 = v50
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[1]))
									v58 = v52 + v40<<(uint(int32(13))%32) + int32(-8192)
								}
								v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v58)+6)))
								if v61&int32(64) == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									v67 = int32(0)
									F_ginFinishSplit(m, l0, l1, v67, v67)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v67 = int32(0)
						F_ginFinishSplit(m, l0, l1, v67, v67)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			if l2 == int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_LockBuffer(m, v32, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_LockBuffer(m, v36, int32(2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v40 < int32(0) {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[0]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v40^int32(-1))<<(uint(int32(2))%32))))
							v58 = v50
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[1]))
							v58 = v52 + v40<<(uint(int32(13))%32) + int32(-8192)
						}
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v58)+6)))
						if v61&int32(64) == int32(0) {
							m.G0 = v8 + int32(16)
							return
						} else {
							v67 = int32(0)
							F_ginFinishSplit(m, l0, l1, v67, v67)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v67 = int32(0)
				F_ginFinishSplit(m, l0, l1, v67, v67)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_ginInsertBAEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v408 int32
	_ = v408
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v498 int32
	_ = v498
	v3 = l2
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	if int32(0) < l5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(1)
	v29 = int32(base.Ui32(l5)>>(uint(v27)%32)) | l5
	v32 = int32(base.Ui32(v29)>>(uint(int32(2))%32)) | v29
	v33 = int32(4)
	v35 = int32(base.Ui32(v32)>>(uint(v33)%32)) | v32
	v38 = int32(base.Ui32(v35)>>(uint(int32(8))%32)) | v35
	v62 = int32(base.Ui32(v38)>>(uint(int32(17))%32)) | int32(base.Ui32(v38)>>(uint(v27)%32)) + v27
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v23 + int32(48)
	return
L4:
	;
	v69 = v62 - int32(1)
	if base.Ui32(v69) < base.Ui32(l5) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v88 = v69
	goto L9
L7:
	;
	goto L8
L8:
	;
	v498 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
	if v498 != 0 {
		v62 = v498
		goto L4
	} else {
		goto L138
	}
L9:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3+v88<<(uint(int32(2))%32))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v88))))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v96
	v104 = v23 + int32(8)
	v106 = v23 + int32(7)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 == int32(_a_F_ginInsertBAEntries_0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L8
L11:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
	if v431 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L12:
	;
	v163 = int32(0)
	v176 = int32(0)
	goto L14
L13:
	;
	v119 = v109
	goto L15
L14:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v181 = m.T0[v180].(func(*base.Module, int32) int32)(m, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L27
	}
L15:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v135 = m.T0[v134].(func(*base.Module, int32, int32, int32) int32)(m, v104, v119, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v163 = int32(base.Ui32(v135) >> (uint(int32(31)) % 32))
	v176 = v119
	goto L14
L17:
	;
	return
L18:
	;
	if v135 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	m.T0[v140].(func(*base.Module, int32, int32, int32))(m, v119, v104, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v135 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v143)
	v430 = v119
	goto L11
L23:
	;
	v149 = int32(4)
	goto L25
L24:
	;
	v149 = int32(8)
	goto L25
L25:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v119+v149)))
	if v151 != int32(_a_F_ginInsertBAEntries_0) {
		v119 = v151
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v176
	v184 = int32(_a_F_ginInsertBAEntries_0)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+8)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v184
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v188)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v192 = v190 - int32(16)
	if v192 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	base.MemoryCopy(m, v181+int32(16), v23+int32(24), v192)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if v176 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v181 == v201 {
		v396 = v201
		goto L38
	} else {
		goto L39
	}
L32:
	;
	if v163 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v181
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v181
	goto L31
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v181
	goto L31
L38:
	;
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v408)
	v430 = v181
	goto L11
L39:
	;
	v204 = v181 + int32(12)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v206 != int32(1) {
		v396 = v201
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v215 = v205
	v216 = v204
	v218 = v181
	goto L41
L41:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v230 == v215 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v396 = v380
	goto L38
L43:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v379 == v380 {
		v396 = v380
		goto L38
	} else {
		goto L122
	}
L44:
	;
	v379 = v369
	goto L43
L45:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v233 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v301 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L48:
	;
	v236 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v236)
	*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v236)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v242)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v379 = v245
	goto L43
L49:
	;
	goto L50
L50:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	if v246 != v218 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v274)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v278)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v283
	if v283 != int32(_a_F_ginInsertBAEntries_0) {
		goto L71
	} else {
		goto L72
	}
L52:
	;
	v270 = v218
	v271 = v215
	goto L51
L53:
	;
	goto L54
L54:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v248
	if v248 != int32(_a_F_ginInsertBAEntries_0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v215
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v254 = v253
	goto L57
L56:
	;
	v254 = v229
	goto L57
L57:
	;
	if v218 != int32(_a_F_ginInsertBAEntries_0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v254
	goto L60
L59:
	;
	goto L60
L60:
	;
	if v254 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v215
	if v215 == int32(_a_F_ginInsertBAEntries_0) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v258 == v215 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v246
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v246
	goto L61
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v246
	goto L61
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertBAEntries[0]))
	v270 = int32(_a_F_ginInsertBAEntries_0)
	v271 = v267
	goto L51
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v246
	v270 = v215
	v271 = v246
	goto L51
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = v281
	goto L73
L72:
	;
	goto L73
L73:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	if v282 != int32(_a_F_ginInsertBAEntries_0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v288
	goto L76
L75:
	;
	goto L76
L76:
	;
	if v288 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v281
	if v281 == int32(_a_F_ginInsertBAEntries_0) {
		v369 = v270
		goto L44
	} else {
		goto L84
	}
L78:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	if v292 == v281 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v282
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+8)) = v282
	goto L77
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v282
	goto L77
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = v282
	v369 = v270
	goto L44
L85:
	;
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v304)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v304)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v310 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v310)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v379 = v313
	goto L43
L86:
	;
	goto L87
L87:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v314 != v218 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v342)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v346 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v346)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v351
	if v351 != int32(_a_F_ginInsertBAEntries_0) {
		goto L108
	} else {
		goto L109
	}
L89:
	;
	v338 = v218
	v339 = v215
	goto L88
L90:
	;
	goto L91
L91:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+4)) = v316
	if v316 != int32(_a_F_ginInsertBAEntries_0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = v215
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v322 = v321
	goto L94
L93:
	;
	v322 = v229
	goto L94
L94:
	;
	if v218 != int32(_a_F_ginInsertBAEntries_0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = v322
	goto L97
L96:
	;
	goto L97
L97:
	;
	if v322 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+8)) = v215
	if v215 == int32(_a_F_ginInsertBAEntries_0) {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	if v326 == v215 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v314
	goto L98
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = v314
	goto L98
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v314
	goto L98
L105:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertBAEntries[0]))
	v338 = int32(_a_F_ginInsertBAEntries_0)
	v339 = v335
	goto L88
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v314
	v338 = v215
	v339 = v314
	goto L88
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v349
	goto L110
L109:
	;
	goto L110
L110:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	if v350 != int32(_a_F_ginInsertBAEntries_0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+12)) = v356
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v356 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v349
	if v349 == int32(_a_F_ginInsertBAEntries_0) {
		v369 = v338
		goto L44
	} else {
		goto L121
	}
L115:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v360 == v349 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v350
	goto L114
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = v350
	goto L114
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v350
	goto L114
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+12)) = v350
	v369 = v338
	goto L44
L122:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v385 == int32(1) {
		v215 = v384
		v216 = v379 + int32(12)
		v218 = v379
		goto L41
	} else {
		goto L123
	}
L123:
	;
	goto L42
L124:
	;
	if v98 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v475 = v88 + v62<<(uint(int32(1))%32)
	if base.Ui32(v475) < base.Ui32(l5) {
		v88 = v475
		goto L9
	} else {
		goto L137
	}
L127:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	v438 = v437 + v3<<(uint(v33)%32)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+10)))
	if v439 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v455 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+24)) = uint8(v455)
	*(*int64)(unsafe.Add(mBase, uint32(v430)+32)) = int64(4294967301)
	v460 = F_palloc(m, int32(30))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L17
	} else {
		goto L135
	}
L130:
	;
	v443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v438)+8)))
	v444 = F_datumCopy(m, v96, int32(0), v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L17
	} else {
		goto L133
	}
L131:
	;
	v451 = v96
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+16)) = v451
	goto L129
L133:
	;
	v446 = F_GetMemoryChunkSpace(m, v444)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L17
	} else {
		goto L134
	}
L134:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v446 + v448
	v451 = v444
	goto L132
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+28)) = v460
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v460)+4)) = uint16(v463)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v460))) = v465
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v430)+28))
	v468 = F_GetMemoryChunkSpace(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468 + v470
	goto L126
L137:
	;
	goto L10
L138:
	;
	goto L5
}
func F_ginInsertItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l0
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+65)) = uint8(base.B2i32(l4 != v6))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = v9 + int32(74)
	v51 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(80)
	return
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v60 = v57 + v51*int32(6)
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+4)) = uint16(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v63
	v66 = v9 + int32(12)
	v69 = F_ginFindLeafPage(m, v66, int32(0), int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
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
	F_ginInsertValue(m, v66, v69, v9, l4)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.Ui32(v73) < base.Ui32(v74) {
		v51 = v73
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_ginReadTuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 << (uint(v8) % 32)
	v18 = v11 + (l0 + v14&int32(2147418112))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v14 < int32(0) {
		if v19 != 0 {
			v24 = F_ginPostingListDecode(m, v18, v9+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if v28 == v19 {
					v55 = v24
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
					m.G0 = v9 + int32(16)
					return v55
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v35
						F_errmsg_internal(m, int32(_a_F_ginReadTuple_0), v9)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ginReadTuple_1), int32(177), int32(_a_F_ginReadTuple_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
		} else {
			v46 = F_palloc(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v55 = v46
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
				m.G0 = v9 + int32(16)
				return v55
			}
		}
	} else {
		v49 = v19 * int32(6)
		v50 = F_palloc(m, v49)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v49 == int32(0) {
				v55 = v50
			} else {
				base.MemoryCopy(m, v50, v18, v49)
				v55 = v50
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
			m.G0 = v9 + int32(16)
			return v55
		}
	}
}
func F_ginScanBeginPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(38)
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v32)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v32)
	v38 = F_ginFindLeafPage(m, l0, v32, v20)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		return v38
	}
}
func F_ginUpdateStats(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_LockBuffer(m, v13, int32(2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v13 < int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v13^int32(-1))<<(uint(int32(2))%32))))
				v35 = v27
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[1]))
				v35 = v29 + v13<<(uint(int32(13))%32) + int32(-8192)
			}
			v36 = int32(_a_F_ginUpdateStats_0)
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v38 + int32(1)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v44
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v46
			v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			v49 = int32(80)
			*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v49)
			*(*int64)(unsafe.Add(mBase, uint32(v35)+64)) = v48
			F_MarkBufferDirty(m, v13)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+118)))
				if v55 != int32(112) {
					F_UnlockReleaseBuffer(m, v13)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						v112 = int32(_a_F_ginUpdateStats_0)
						v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
						m.G0 = v10 + int32(96)
						return
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[3]))
					if v59 <= int32(0) {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v62|l2 != 0 {
							F_UnlockReleaseBuffer(m, v13)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								v112 = int32(_a_F_ginUpdateStats_0)
								v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
								m.G0 = v10 + int32(96)
								return
							}
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v64 == int32(0) {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v67
								v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v69
								*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = int64(-1)
								v76 = v35 + int32(24)
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v77
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v79
								v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v81
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v83
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v85
								v87 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v87
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v89
								F_XLogBeginInsert(m)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_XLogRegisterData(m, v10+int32(8), int32(88))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										F_XLogRegisterBuffer(m, int32(0), v13, int32(14))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											v104 = F_XLogInsert(m, int32(13), int32(96))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v35))) = base.I64_rotr(v104, int64(32))
												F_UnlockReleaseBuffer(m, v13)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v112 = int32(_a_F_ginUpdateStats_0)
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
													m.G0 = v10 + int32(96)
													return
												}
											}
										}
									}
								}
							} else {
								F_UnlockReleaseBuffer(m, v13)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return
								} else {
									v112 = int32(_a_F_ginUpdateStats_0)
									v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
									m.G0 = v10 + int32(96)
									return
								}
							}
						}
					} else {
						if l2 != 0 {
							F_UnlockReleaseBuffer(m, v13)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								v112 = int32(_a_F_ginUpdateStats_0)
								v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
								m.G0 = v10 + int32(96)
								return
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v67
							v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v69
							*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = int64(-1)
							v76 = v35 + int32(24)
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v77
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v79
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v81
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v83
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v85
							v87 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v87
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v89
							F_XLogBeginInsert(m)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_XLogRegisterData(m, v10+int32(8), int32(88))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_XLogRegisterBuffer(m, int32(0), v13, int32(14))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										v104 = F_XLogInsert(m, int32(13), int32(96))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v35))) = base.I64_rotr(v104, int64(32))
											F_UnlockReleaseBuffer(m, v13)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = int32(_a_F_ginUpdateStats_0)
												v114 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v114 - int32(1)
												m.G0 = v10 + int32(96)
												return
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
func F_gin_btree_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2))) = uint8(v3)
	return int32(1)
}
func F_gin_cmp_prefix(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v122 int32
	_ = v122
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v15 = v8 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v21 = v19 & int32(1)
			if v21 != 0 {
				v22 = v15
			} else {
				v22 = v8 + int32(4)
			}
			if v19 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v21 != 0 {
					v49 = int32(base.Ui32(v19)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v17 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v17 + int32(4)
			}
			if v54 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v49 == int32(0) {
				v110 = int32(0)
			} else {
				if v84 == int32(0) {
					v110 = base.B2i32(int32(0) < v49)
				} else {
					if base.Ui32(v49) < base.Ui32(v84) {
						v99 = v49
					} else {
						v99 = v84
					}
					v100 = F_memcmp(m, v22, v57, v99)
					mBase = m.M
					if v100 != 0 {
						v110 = v100
					} else {
						v110 = base.B2i32(v84 < v49)
					}
				}
			}
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v111 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v115 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							if v110 < int32(0) {
								v122 = int32(1)
							} else {
								v122 = v110
							}
							return v122
						}
					} else {
						if v110 < int32(0) {
							v122 = int32(1)
						} else {
							v122 = v110
						}
						return v122
					}
				}
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v115 != v17 {
					F_pfree(m, v17)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						if v110 < int32(0) {
							v122 = int32(1)
						} else {
							v122 = v110
						}
						return v122
					}
				} else {
					if v110 < int32(0) {
						v122 = int32(1)
					} else {
						v122 = v110
					}
					return v122
				}
			}
		}
	}
}
func F_gin_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	switch int32(base.Ui32(v14-int32(32)) >> (uint(int32(4)) % 32)) {
	case 0:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		if v21&int32(2) != 0 {
			v24 = int32(84)
		} else {
			v24 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v24
		if v21&int32(1) != 0 {
			v30 = int32(84)
		} else {
			v30 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v30
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_0), v10-int32(-64))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v37&int32(2) == int32(0) {
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
				v46 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v44 | v45<<(uint(v46)%32)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v42 | v43<<(uint(v46)%32)
				F_appendStringInfo(m, l0, int32(_a_F_gin_desc_1), v10+int32(48))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+105)))
					if v62 == int32(1) {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+106)))
						if v65 == int32(1) {
							F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						}
					} else {
						v74 = int32(0)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
						if v78 < v74 {
							v100 = v74
						} else {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+int32(0))+76)))
							if v83 != int32(1) {
								v100 = v74
							} else {
								v87 = v77 + int32(76)
								v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+43)))
								if v88 == int32(0) {
									v100 = v74
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
									v100 = v98
								}
							}
						}
						v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
						if v104&int32(1) == int32(0) {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
							if v111 != 0 {
								v112 = int32(84)
							} else {
								v112 = int32(70)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v112
							F_appendStringInfo(m, l0, int32(_a_F_gin_desc_4), v10+int32(16))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							if v104&int32(2) != 0 {
								F_desc_recompress_leaf(m, l0, v100)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v10 + int32(112)
									return
								}
							} else {
								v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
								v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
								v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
								v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
								v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+10)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v127
								v129 = int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v125 | v126<<(uint(v129)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123 | v124<<(uint(v129)%32)
								F_appendStringInfo(m, l0, int32(_a_F_gin_desc_5), v10+int32(32))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									m.G0 = v10 + int32(112)
									return
								}
							}
						}
					}
				}
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+105)))
				if v62 == int32(1) {
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+106)))
					if v65 == int32(1) {
						F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					}
				} else {
					v74 = int32(0)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
					if v78 < v74 {
						v100 = v74
					} else {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+int32(0))+76)))
						if v83 != int32(1) {
							v100 = v74
						} else {
							v87 = v77 + int32(76)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+43)))
							if v88 == int32(0) {
								v100 = v74
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
								v100 = v98
							}
						}
					}
					v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
					if v104&int32(1) == int32(0) {
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
						if v111 != 0 {
							v112 = int32(84)
						} else {
							v112 = int32(70)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v112
						F_appendStringInfo(m, l0, int32(_a_F_gin_desc_4), v10+int32(16))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					} else {
						if v104&int32(2) != 0 {
							F_desc_recompress_leaf(m, l0, v100)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
							v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
							v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
							v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+10)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v127
							v129 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v125 | v126<<(uint(v129)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123 | v124<<(uint(v129)%32)
							F_appendStringInfo(m, l0, int32(_a_F_gin_desc_5), v10+int32(32))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	case 1:
		v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+24)))
		if v144&int32(4) != 0 {
			v147 = int32(84)
		} else {
			v147 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v147
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_6), v10+int32(96))
		mBase = m.M
		v153 = m.ExcPending
		if v153 != 0 {
			return
		} else {
			v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+24)))
			if v156&int32(2) != 0 {
				v159 = int32(84)
			} else {
				v159 = int32(70)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v159
			if v156&int32(1) != 0 {
				v165 = int32(84)
			} else {
				v165 = int32(70)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v165
			F_appendStringInfo(m, l0, int32(_a_F_gin_desc_7), v10+int32(80))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return
			} else {
				m.G0 = v10 + int32(112)
				return
			}
		}
	default:
		m.G0 = v10 + int32(112)
		return
	case 6:
		v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v216
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_8), v10)
		mBase = m.M
		v220 = m.ExcPending
		if v220 != 0 {
			return
		} else {
			m.G0 = v10 + int32(112)
			return
		}
	case 7:
		v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+105)))
		if v172 == int32(1) {
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+106)))
			if v175 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return
				} else {
					m.G0 = v10 + int32(112)
					return
				}
			} else {
				F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
					return
				} else {
					m.G0 = v10 + int32(112)
					return
				}
			}
		} else {
			v184 = int32(0)
			v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+72))
			if v188 < v184 {
				v210 = v184
			} else {
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+int32(0))+76)))
				if v193 != int32(1) {
					v210 = v184
				} else {
					v197 = v187 + int32(76)
					v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+43)))
					if v198 == int32(0) {
						v210 = v184
					} else {
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)+44))
						v210 = v208
					}
				}
			}
			F_desc_recompress_leaf(m, l0, v210)
			mBase = m.M
			v215 = m.ExcPending
			if v215 != 0 {
				return
			} else {
				m.G0 = v10 + int32(112)
				return
			}
		}
	}
}
func F_gin_extract_query_bpchar(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_bpchar_0), int32(2407))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_date(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(-2147483648), int32(1428))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_inet(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_inet_0), int32(2328))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_macaddr(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13909(m, l0, int32(6), int32(2266))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_money(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13908(m, l0, int64(-9223372036854775807-1), int32(2119))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_oid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(0), int32(2100))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_trgm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
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
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		if base.Ui32(int32(11)) < base.Ui32(v22) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v290 = m.ExcPending
			if v290 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
				F_errmsg_internal(m, int32(_a_F_gin_extract_query_trgm_0), v15)
				mBase = m.M
				v294 = m.ExcPending
				if v294 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_extract_query_trgm_1), int32(139), int32(_a_F_gin_extract_query_trgm_2))
					mBase = m.M
					v299 = m.ExcPending
					if v299 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(1) << (uint(v22) % 32)
			if v28&int32(2690) == int32(0) {
				if v28&int32(24) != 0 {
					v178 = int32(1)
					v179 = v18 + v178
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					v184 = v182 & v178
					if v184 != 0 {
						v185 = v179
					} else {
						v185 = v18 + int32(4)
					}
					if v182 == int32(1) {
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
						if v191 == int32(18) {
							v194 = int32(16)
						} else {
							v194 = int32(0)
						}
						if base.Ui32((v191-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v201 = int32(4)
						} else {
							v201 = v194
						}
						v212 = v201
					} else {
						v202 = int32(1)
						if v184 != 0 {
							v212 = int32(base.Ui32(v182)>>(uint(v202)%32)) - v202
						} else {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							v212 = int32(base.Ui32(v206)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v213 = F_generate_wildcard_trgm(m, v185, v212)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return int32(0)
					} else {
						v217 = v213
						v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
						v231 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(5)
						v232 = int32(3)
						v233 = base.I32_div_u_s(v231, v232)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v233
						if base.Ui32(v231) < base.Ui32(v232) {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
							v304 = int32(0)
							m.G0 = v15 + int32(16)
							return v304
						} else {
							v237 = int32(1)
							if base.Ui32(v233) <= base.Ui32(v237) {
								v240 = v237
							} else {
								v240 = v233
							}
							v246 = F_palloc(m, v233<<(uint(int32(2))%32))
							mBase = m.M
							v247 = m.ExcPending
							if v247 != 0 {
								return int32(0)
							} else {
								v248 = v217 + int32(5)
								v250 = int32(0)
								for {
									v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
									v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
									v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
									*(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32)))) = v263 | (v264<<(uint(int32(8))%32) | v267<<(uint(int32(16))%32))
									v276 = v250 + int32(1)
									if v276 != v240 {
										v248 = v248 + int32(3)
										v250 = v276
										continue
									} else {
										break
									}
									break
								}
								v304 = v246
								m.G0 = v15 + int32(16)
								return v304
							}
						}
					}
				} else {
					if v28&int32(96) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v290 = m.ExcPending
						if v290 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_trgm_0), v15)
							mBase = m.M
							v294 = m.ExcPending
							if v294 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_trgm_1), int32(139), int32(_a_F_gin_extract_query_trgm_2))
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_gin_extract_query_trgm[0]))
						v45 = F_createTrgmNFA(m, v18, v40, v15+int32(12), v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 == int32(0) {
								v279 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v26))) = v279
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v304 = v279
								m.G0 = v15 + int32(16)
								return v304
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
								v53 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(5)
								if base.Ui32(v53) < base.Ui32(int32(3)) {
									v279 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v26))) = v279
									*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
									v304 = v279
									m.G0 = v15 + int32(16)
									return v304
								} else {
									v57 = base.I32_div_u_s(v53, int32(3))
									v60 = F_palloc(m, v57<<(uint(int32(2))%32))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v39))) = v60
										v63 = int32(1)
										if base.Ui32(v57) <= base.Ui32(v63) {
											v66 = v63
										} else {
											v66 = v57
										}
										v68 = v66 & int32(7)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
										v70 = int32(0)
										if base.Ui32(int32(24)) <= base.Ui32(v53) {
											v77 = v70
											v82 = int32(0)
											for {
												v90 = v60 + v77<<(uint(int32(2))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v69
												*(*int32)(unsafe.Add(mBase, uint32(v90))) = v69
												v99 = int32(8)
												v100 = v77 + v99
												v102 = v82 + v99
												if v102 != v66&int32(2147483640) {
													v77 = v100
													v82 = v102
													continue
												} else {
													break
												}
												break
											}
											if v68 == int32(0) {
												v217 = v45
											} else {
												v107 = v100
												v119 = v107
												v128 = int32(0)
												for {
													*(*int32)(unsafe.Add(mBase, uint32(v60+v119<<(uint(int32(2))%32)))) = v69
													v134 = int32(1)
													v137 = v128 + v134
													if v137 != v68 {
														v119 = v119 + v134
														v128 = v137
														continue
													} else {
														break
													}
													break
												}
												v217 = v45
											}
										} else {
											v107 = v70
											v119 = v107
											v128 = int32(0)
											for {
												*(*int32)(unsafe.Add(mBase, uint32(v60+v119<<(uint(int32(2))%32)))) = v69
												v134 = int32(1)
												v137 = v128 + v134
												if v137 != v68 {
													v119 = v119 + v134
													v128 = v137
													continue
												} else {
													break
												}
												break
											}
											v217 = v45
										}
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
										v231 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(5)
										v232 = int32(3)
										v233 = base.I32_div_u_s(v231, v232)
										*(*int32)(unsafe.Add(mBase, uint32(v26))) = v233
										if base.Ui32(v231) < base.Ui32(v232) {
											*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
											v304 = int32(0)
											m.G0 = v15 + int32(16)
											return v304
										} else {
											v237 = int32(1)
											if base.Ui32(v233) <= base.Ui32(v237) {
												v240 = v237
											} else {
												v240 = v233
											}
											v246 = F_palloc(m, v233<<(uint(int32(2))%32))
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return int32(0)
											} else {
												v248 = v217 + int32(5)
												v250 = int32(0)
												for {
													v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
													v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
													v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
													*(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32)))) = v263 | (v264<<(uint(int32(8))%32) | v267<<(uint(int32(16))%32))
													v276 = v250 + int32(1)
													if v276 != v240 {
														v248 = v248 + int32(3)
														v250 = v276
														continue
													} else {
														break
													}
													break
												}
												v304 = v246
												m.G0 = v15 + int32(16)
												return v304
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v139 = int32(1)
				v140 = v18 + v139
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				v145 = v143 & v139
				if v145 != 0 {
					v146 = v140
				} else {
					v146 = v18 + int32(4)
				}
				if v143 == int32(1) {
					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
					if v152 == int32(18) {
						v155 = int32(16)
					} else {
						v155 = int32(0)
					}
					if base.Ui32((v152-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v162 = int32(4)
					} else {
						v162 = v155
					}
					v163 = F_generate_trgm(m, v146, v162)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						v217 = v163
						v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
						v231 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(5)
						v232 = int32(3)
						v233 = base.I32_div_u_s(v231, v232)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v233
						if base.Ui32(v231) < base.Ui32(v232) {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
							v304 = int32(0)
							m.G0 = v15 + int32(16)
							return v304
						} else {
							v237 = int32(1)
							if base.Ui32(v233) <= base.Ui32(v237) {
								v240 = v237
							} else {
								v240 = v233
							}
							v246 = F_palloc(m, v233<<(uint(int32(2))%32))
							mBase = m.M
							v247 = m.ExcPending
							if v247 != 0 {
								return int32(0)
							} else {
								v248 = v217 + int32(5)
								v250 = int32(0)
								for {
									v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
									v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
									v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
									*(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32)))) = v263 | (v264<<(uint(int32(8))%32) | v267<<(uint(int32(16))%32))
									v276 = v250 + int32(1)
									if v276 != v240 {
										v248 = v248 + int32(3)
										v250 = v276
										continue
									} else {
										break
									}
									break
								}
								v304 = v246
								m.G0 = v15 + int32(16)
								return v304
							}
						}
					}
				} else {
					if v145 != 0 {
						v165 = int32(1)
						v169 = F_generate_trgm(m, v146, int32(base.Ui32(v143)>>(uint(v165)%32))-v165)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							v217 = v169
							v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v231 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(5)
							v232 = int32(3)
							v233 = base.I32_div_u_s(v231, v232)
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v233
							if base.Ui32(v231) < base.Ui32(v232) {
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v304 = int32(0)
								m.G0 = v15 + int32(16)
								return v304
							} else {
								v237 = int32(1)
								if base.Ui32(v233) <= base.Ui32(v237) {
									v240 = v237
								} else {
									v240 = v233
								}
								v246 = F_palloc(m, v233<<(uint(int32(2))%32))
								mBase = m.M
								v247 = m.ExcPending
								if v247 != 0 {
									return int32(0)
								} else {
									v248 = v217 + int32(5)
									v250 = int32(0)
									for {
										v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
										v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
										v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
										*(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32)))) = v263 | (v264<<(uint(int32(8))%32) | v267<<(uint(int32(16))%32))
										v276 = v250 + int32(1)
										if v276 != v240 {
											v248 = v248 + int32(3)
											v250 = v276
											continue
										} else {
											break
										}
										break
									}
									v304 = v246
									m.G0 = v15 + int32(16)
									return v304
								}
							}
						}
					} else {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v176 = F_generate_trgm(m, v146, int32(base.Ui32(v171)>>(uint(int32(2))%32))-int32(4))
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							v217 = v176
							v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v231 = int32(base.Ui32(v227)>>(uint(int32(2))%32)) - int32(5)
							v232 = int32(3)
							v233 = base.I32_div_u_s(v231, v232)
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v233
							if base.Ui32(v231) < base.Ui32(v232) {
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v304 = int32(0)
								m.G0 = v15 + int32(16)
								return v304
							} else {
								v237 = int32(1)
								if base.Ui32(v233) <= base.Ui32(v237) {
									v240 = v237
								} else {
									v240 = v233
								}
								v246 = F_palloc(m, v233<<(uint(int32(2))%32))
								mBase = m.M
								v247 = m.ExcPending
								if v247 != 0 {
									return int32(0)
								} else {
									v248 = v217 + int32(5)
									v250 = int32(0)
									for {
										v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
										v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
										v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
										*(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32)))) = v263 | (v264<<(uint(int32(8))%32) | v267<<(uint(int32(16))%32))
										v276 = v250 + int32(1)
										if v276 != v240 {
											v248 = v248 + int32(3)
											v250 = v276
											continue
										} else {
											break
										}
										break
									}
									v304 = v246
									m.G0 = v15 + int32(16)
									return v304
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gin_extract_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	v2 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v2 {
		v216 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v227 != v18 {
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v28 = v18 + int32(8)
	v29 = F_tsquery_requires_match(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = int32(0)
	goto L7
L6:
	;
	v33 = int32(2)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v35 <= int32(0) {
		v127 = v2
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v127
	v141 = v127 << (uint(int32(2)) % 32)
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L20
	}
L9:
	;
	v39 = v35 & int32(3)
	v40 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v35) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = v40
	v47 = v2
	v57 = v2
	goto L13
L11:
	;
	v86 = v40
	v87 = v2
	goto L12
L12:
	;
	v100 = v86
	v101 = v87
	v112 = v2
	goto L17
L13:
	;
	v61 = v28 + v46*int32(12)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v63 = int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+24)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+36)))
	v77 = v47 + base.B2i32(v62 == v63) + base.B2i32(v66 == v63) + base.B2i32(v70 == v63) + base.B2i32(v74 == v63)
	v78 = int32(4)
	v79 = v46 + v78
	v81 = v57 + v78
	if v81 != v35&int32(2147483644) {
		v46 = v79
		v47 = v77
		v57 = v81
		goto L13
	} else {
		goto L15
	}
L14:
	;
	if v39 == int32(0) {
		v127 = v77
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v86 = v79
	v87 = v77
	goto L12
L17:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v100*int32(12)))))
	v117 = int32(1)
	v119 = v101 + base.B2i32(v116 == v117)
	v123 = v112 + v117
	if v123 != v39 {
		v100 = v100 + v117
		v101 = v119
		v112 = v123
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v127 = v119
	goto L8
L19:
	;
	goto L18
L20:
	;
	v144 = F_palloc(m, v127)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v144
	v147 = F_palloc(m, v141)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v147
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v153 = F_palloc0(m, v150<<(uint(int32(2))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v155 <= int32(0) {
		v216 = v142
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v160 = int32(0)
	v161 = v155
	v169 = v2
	goto L25
L25:
	;
	v175 = v28 + v169*int32(12)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v176 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v216 = v142
	goto L1
L27:
	;
	v180 = v160 << (uint(int32(2)) % 32)
	v182 = int32(12)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v191 = F_cstring_to_text_with_len(m, v28+v161*v182+int32(base.Ui32(v185)>>(uint(v182)%32)), v185&int32(4095))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	v207 = v160
	v208 = v161
	goto L29
L29:
	;
	v211 = v169 + int32(1)
	if v211 < v208 {
		v160 = v207
		v161 = v208
		v169 = v211
		goto L25
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142+v180))) = v191
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v144))) = uint8(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v197+v180))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v153+v169<<(uint(int32(2))%32)))) = v160
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v207 = v160 + int32(1)
	v208 = v204
	goto L29
L31:
	;
	goto L26
L32:
	;
	F_pfree(m, v18)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v216
L35:
	;
	goto L34
}
func F_gin_extract_value_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		return v7
	}
}
func F_gin_trgm_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v369 float32
	_ = v369
	var v376 int32
	_ = v376
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v22)
	if base.Ui32(int32(11)) < base.Ui32(v20) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v376
L2:
	;
	v376 = base.F64_le(v235, base.F64_promote_f32(base.F32_div(v369, base.F32_convert_i32_s(v18))))
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L47
	} else {
		goto L69
	}
L4:
	;
	v27 = int32(1) << (uint(v20) % 32)
	if v27&int32(642) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v18 <= int32(0) {
		v376 = v32
		goto L1
	} else {
		goto L64
	}
L6:
	;
	v32 = int32(1)
	if v27&int32(2072) != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v235 = F_index_strategy_get_limit(m, v20)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L47
	} else {
		goto L48
	}
L9:
	;
	if v27&int32(96) == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v18 <= int32(0) {
		v376 = v32
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v42 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v376 = v234
	goto L1
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	base.MemoryFill(m, v53, int32(0), v52)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	base.MemoryFill(m, v57, int32(0), v56)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if int32(0) < v60 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v67 = v42
	v68 = v42
	goto L22
L20:
	;
	goto L21
L21:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v146 = v134
	v151 = v42
	goto L34
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63+v67<<(uint(int32(2))%32))))
	v80 = v79 + v68
	if v79 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v119 = v67 + int32(1)
	if v119 != v60 {
		v67 = v119
		v68 = v80
		goto L22
	} else {
		goto L32
	}
L25:
	;
	v85 = v68
	goto L26
L26:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v85))))
	if v96 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v102+v67))) = uint8(v104)
	goto L24
L28:
	;
	v100 = v85 + int32(1)
	if v100 < v80 {
		v85 = v100
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L24
L32:
	;
	goto L23
L33:
	;
	goto L12
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v136+v151<<(uint(int32(2))%32))))
	v159 = v139 + v156<<(uint(int32(3))%32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if int32(0) < v160 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v234 = int32(0)
	goto L33
L36:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v168 = int32(0)
	v171 = v146
	goto L39
L37:
	;
	v211 = v146
	goto L38
L38:
	;
	v219 = v151 + int32(1)
	if v219 < v211 {
		v146 = v211
		v151 = v219
		goto L34
	} else {
		goto L46
	}
L39:
	;
	v180 = v164 + v168<<(uint(int32(3))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v181))))
	if v183 != int32(1) {
		v201 = v171
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v211 = v201
	goto L38
L41:
	;
	v204 = v168 + int32(1)
	if v204 != v160 {
		v168 = v204
		v171 = v201
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v186 = int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v187 == v186 {
		v234 = v186
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v190 = v133 + v187
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v191 != 0 {
		v201 = v171
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v136+v171<<(uint(int32(2))%32)))) = v187
	v201 = v171 + v192
	goto L41
L45:
	;
	goto L40
L46:
	;
	goto L35
L47:
	;
	return int32(0)
L48:
	;
	if int32(0) < v18 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v242 = v18 & int32(3)
	v243 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v18) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	if v18 != 0 {
		v369 = float32(0)
		goto L2
	} else {
		goto L63
	}
L52:
	;
	v369 = base.F32_convert_i32_u(v310)
	goto L2
L53:
	;
	v248 = v243
	v250 = v2
	v257 = v2
	goto L56
L54:
	;
	v276 = v243
	v278 = v2
	goto L55
L55:
	;
	v288 = v276
	v290 = v278
	v295 = v2
	goto L60
L56:
	;
	v260 = v248 + v19
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+2)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+3)))
	v268 = v250 + v261 + v263 + v265 + v267
	v269 = int32(4)
	v270 = v248 + v269
	v272 = v257 + v269
	if v272 != v18&int32(2147483644) {
		v248 = v270
		v250 = v268
		v257 = v272
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v242 == int32(0) {
		v310 = v268
		goto L52
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v276 = v270
	v278 = v268
	goto L55
L60:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288+v19))))
	v302 = v290 + v301
	v303 = int32(1)
	v306 = v295 + v303
	if v306 != v242 {
		v288 = v288 + v303
		v290 = v302
		v295 = v306
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v310 = v302
	goto L52
L62:
	;
	goto L61
L63:
	;
	v376 = v2
	goto L1
L64:
	;
	v324 = int32(0)
	goto L65
L65:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324+v19))))
	if v337 != int32(1) {
		v376 = v337
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v376 = v337
	goto L1
L67:
	;
	v341 = v324 + int32(1)
	if v341 != v18 {
		v324 = v341
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
	F_errmsg_internal(m, int32(_a_F_gin_trgm_consistent_0), v15)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_gin_trgm_consistent_1), int32(254), int32(_a_F_gin_trgm_consistent_2))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L47
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_tsquery_triconsistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 <= int32(0) {
		v31 = int32(0)
		m.G0 = v7 + int32(16)
		return v31
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
		v18 = v9 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
		v24 = F_TS_execute_ternary(m, v18, v7+int32(4))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v31 = base.I32_extend8_s(v24)
			m.G0 = v7 + int32(16)
			return v31
		}
	}
}
func F_gin_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_cleanup[0])) = int32(0)
		return
	}
}
