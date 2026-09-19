package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinBufferStoreTuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v15 != 0 {
		v28 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v33 = (v29 + int32(17)) & int32(_a_F_GinBufferStoreTuple_0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = F_ginPostingListDecodeAllSegments(m, l1+v33, v35-v33, v12+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v17 = l1 + int32(16)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v18 != int32(1) {
		v28 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	base.MemoryCopy(m, v12+int32(8), v17, v21)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v28 = v25
	goto L1
L7:
	;
	return
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v44)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v46
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v48)
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v50)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v52)
	v54 = int32(0)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v56 == v54 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v67 = v41
	goto L11
L11:
	;
	if v67 <= int32(0) {
		v111 = v67
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v61 = F_datumCopy(m, v28, v52&int32(1), v50)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	v64 = v54
	v65 = v54
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
	v67 = v64
	goto L11
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = v63
	v65 = v61
	goto L14
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v111 <= v112 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v74 = int32(6)
	v78 = v73 + v67*v74 - v74
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v84 = (l1 + v79 + int32(17)) & int32(-2)
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+2)))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	v90 = int32(16)
	v92 = v88 | v89<<(uint(v90)%32)
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+2)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
	v97 = v93 | v94<<(uint(v90)%32)
	if base.Ui32(v92) < base.Ui32(v97) {
		v108 = int32(-1)
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v108 != 0 {
		v111 = v109
		goto L16
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	if base.Ui32(v97) < base.Ui32(v92) {
		v108 = int32(1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+4)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	if base.Ui32(v102) < base.Ui32(v103) {
		v108 = int32(-1)
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v108 = base.B2i32(base.Ui32(v103) < base.Ui32(v102))
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v109
	v111 = v109
	goto L16
L24:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v182 = (v178 + v179) * int32(6)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v183 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v119 = v112
	goto L26
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v128 = v125 + v119*int32(6)
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v134 = (l1 + int32(16) + v129 + int32(1)) & int32(-2)
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+2)))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
	v140 = int32(16)
	v142 = v138 | v139<<(uint(v140)%32)
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134))))
	v147 = v143 | v144<<(uint(v140)%32)
	if base.Ui32(v142) < base.Ui32(v147) {
		v158 = int32(-1)
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L24
L28:
	;
	if int32(0) < v158 {
		goto L24
	} else {
		goto L33
	}
L29:
	;
	goto L28
L30:
	;
	if base.Ui32(v147) < base.Ui32(v142) {
		v158 = int32(1)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+4)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)))
	if base.Ui32(v152) < base.Ui32(v153) {
		v158 = int32(-1)
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v158 = base.B2i32(base.Ui32(v153) < base.Ui32(v152))
	goto L29
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v161 + v162
	v166 = v119 + v162
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v166 < v167 {
		v119 = v166
		goto L26
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v201 = F_ginMergeItemPointers(m, v190+v192*int32(6), v196-v192, v39, v198, v12+int32(4))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L41
	}
L36:
	;
	v186 = F_palloc(m, v182)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v188 = F_repalloc(m, v183, v182)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L40
	}
L39:
	;
	v190 = v186
	goto L35
L40:
	;
	v190 = v188
	goto L35
L41:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v205 = v203 * int32(6)
	if v205 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v206+v207*int32(6), v201, v205)
	goto L44
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v201)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v214 + v215
	F_pfree(m, v39)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v12 + int32(16)
	return
}
func F__gin_build_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	v1 = l0
	v2 = l1
	v4 = l3
	v5 = l4
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if v2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v74 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v74
	if l6 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v21 = v9
	goto L4
L3:
	;
	v21 = int32(4)
	goto L4
L4:
	;
	if v2|v5 != 0 {
		v72 = v21
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v4 {
		v72 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	switch v4 + int32(2) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L8
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v72 = int32(base.Ui32(v68) >> (uint(int32(2)) % 32))
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v50 = F_strlen(m, l2)
	mBase = m.M
	v72 = v50 + int32(1)
	goto L1
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v27 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(18)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v33 == v31 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v27&int32(1) == int32(0) {
		goto L7
	} else {
		goto L20
	}
L14:
	;
	v36 = v31
	goto L16
L15:
	;
	v36 = int32(2)
	goto L16
L16:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v43 = int32(6)
	goto L19
L18:
	;
	v43 = v36
	goto L19
L19:
	;
	v72 = v43
	goto L1
L20:
	;
	v72 = int32(base.Ui32(v27) >> (uint(int32(1)) % 32))
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v4
	F_errmsg_internal(m, int32(_a_F__gin_build_tuple_0), v17)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F__gin_build_tuple_1), int32(2270), int32(_a_F__gin_build_tuple_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v86 = int32(0)
	v91 = v9
	goto L28
L26:
	;
	v142 = v9
	goto L27
L27:
	;
	v146 = (v72 + int32(17)) & int32(-2)
	v147 = v142 + v146
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v147
	v149 = F_palloc0(m, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L37
	}
L28:
	;
	v93 = F_palloc(m, int32(12))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L21
	} else {
		goto L30
	}
L29:
	;
	v142 = v114
	goto L27
L30:
	;
	v102 = F_ginCompressPostingList(m, l5+v86*int32(6), l6-v86, int32(_a_F__gin_build_tuple_3), v17+int32(4))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v106 = v105 + v86
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
	v114 = v91 + (v107+int32(1))&int32(_a_F__gin_build_tuple_4) + int32(8)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v115 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v93
	if base.Ui32(v106) < base.Ui32(l6) {
		v86 = v106
		v91 = v114
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v121 = v116
	goto L32
L34:
	;
	goto L35
L35:
	;
	v118 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v118
	v121 = v118
	goto L32
L36:
	;
	goto L29
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+11)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+4)) = uint16(v1)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = l6
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)) = uint16(v72)
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+10)) = uint8(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+8)) = uint16(v4)
	if v2 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v179 = int32(0)
	if base.B2i32(v178 == v179)|base.B2i32(v178 == v17+int32(8)) == v179 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	if v5 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = l2
	goto L38
L41:
	;
	goto L42
L42:
	;
	if int32(0) < v4 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v4 == int32(0) {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	switch v4 + int32(2) {
	case 0:
		goto L47
	case 1:
		goto L48
	default:
		goto L38
	}
L46:
	;
	base.MemoryCopy(m, v149+int32(16), l2, v4)
	goto L38
L47:
	;
	if v72 == int32(0) {
		goto L38
	} else {
		goto L50
	}
L48:
	;
	if v72 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	base.MemoryCopy(m, v149+int32(16), l2, v72)
	goto L38
L50:
	;
	base.MemoryCopy(m, v149+int32(16), l2, v72)
	goto L38
L51:
	;
	v189 = v149 + v146
	v196 = v178
	goto L54
L52:
	;
	goto L53
L53:
	;
	m.G0 = v17 + int32(16)
	return v149
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+6)))
	v210 = (v204+int32(1))&int32(_a_F__gin_build_tuple_4) + int32(8)
	if v210 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	base.MemoryCopy(m, v189, v203, v210)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+6)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	F_pfree(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v196)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	v229 = int32(8)
	if v202 != v17+v229 {
		v189 = v189 + (v213+int32(1))&int32(_a_F__gin_build_tuple_4) + v229
		v196 = v202
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L55
}
func F_ginGetStats(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v5 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_LockBuffer(m, v5, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v5 < int32(0) {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_ginGetStats[0]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(v5^int32(-1))<<(uint(int32(2))%32))))
				v27 = v19
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ginGetStats[1]))
				v27 = v21 + v5<<(uint(int32(13))%32) + int32(-8192)
			}
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v27)+64))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v38
			F_UnlockReleaseBuffer(m, v5)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ginInitBA(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v8 = F_palloc(m, int32(28))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(33)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(31)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(40)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_ginInitBA_0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
		return
	}
}
func F_ginMergeItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	v12 = l1 + l3
	v15 = F_palloc(m, v12*int32(6))
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
	v19 = int32(0)
	if base.B2i32(l1 == v19)|base.B2i32(l3 == v19) == v19 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v257
	return v15
L4:
	;
	v66 = int32(6)
	v68 = l2 + l3*v66
	v71 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-int32(4)))))
	v72 = int64(32)
	v76 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-v66))))
	v77 = int64(48)
	v82 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-int32(2)))))
	v84 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v88 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui64(v84|(v85<<(uint(v72)%64)|v88<<(uint(v77)%64))) <= base.Ui64(v71<<(uint(v72)%64)|v76<<(uint(v77)%64)|v82) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v26 = int32(6)
	v28 = l0 + l1*v26
	v31 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-int32(4)))))
	v32 = int64(32)
	v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-v26))))
	v37 = int64(48)
	v42 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-int32(2)))))
	v44 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v45 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v48 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui64(v44|(v45<<(uint(v32)%64)|v48<<(uint(v37)%64))) <= base.Ui64(v31<<(uint(v32)%64)|v36<<(uint(v37)%64)|v42) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v56 = l1 * int32(6)
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	base.MemoryCopy(m, v15, l0, v56)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v59 = l3 * int32(6)
	if v59 == int32(0) {
		v257 = v12
		goto L3
	} else {
		goto L12
	}
L12:
	;
	base.MemoryCopy(m, v56+v15, l2, v59)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v15
L13:
	;
	v99 = v15
	v100 = l0
	v101 = l2
	goto L16
L14:
	;
	goto L15
L15:
	;
	v242 = l3 * int32(6)
	if v242 != 0 {
		goto L41
	} else {
		goto L42
	}
L16:
	;
	v107 = base.I32_div_s(v101-l2, int32(6))
	if base.Ui32(v107) < base.Ui32(l3) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v164 = base.I32_div_s(v158-l0, int32(6))
	if base.Ui32(v164) < base.Ui32(l1) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v109 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	v110 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
	v111 = int64(32)
	v113 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	v114 = int64(48)
	v117 = v109 | (v110<<(uint(v111)%64) | v113<<(uint(v114)%64))
	v118 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	v119 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101)+2)))
	v122 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	v126 = v118 | (v119<<(uint(v111)%64) | v122<<(uint(v114)%64))
	if base.Ui64(v126) < base.Ui64(v117) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v157 = v99
	v158 = v100
	v159 = v101
	goto L20
L20:
	;
	goto L17
L21:
	;
	v151 = int32(6)
	v152 = v99 + v151
	v155 = base.I32_div_s(v149-l0, v151)
	if base.Ui32(v155) < base.Ui32(l1) {
		v99 = v152
		v100 = v149
		v101 = v150
		goto L16
	} else {
		goto L28
	}
L22:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v130
	v149 = v100
	v150 = v101 + int32(6)
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v117 == v126 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v137
	v139 = int32(6)
	v149 = v100 + v139
	v150 = v101 + v139
	goto L21
L26:
	;
	goto L27
L27:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v143)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v145
	v149 = v100 + int32(6)
	v150 = v101
	goto L21
L28:
	;
	v157 = v152
	v158 = v149
	v159 = v150
	goto L20
L29:
	;
	v171 = v157
	v172 = v158
	goto L32
L30:
	;
	v194 = v157
	goto L31
L31:
	;
	v202 = base.I32_div_s(v159-l2, int32(6))
	if base.Ui32(v202) < base.Ui32(l3) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)) = uint16(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v179
	v181 = int32(6)
	v182 = v171 + v181
	v184 = v172 + v181
	v187 = base.I32_div_s(v184-l0, v181)
	if base.Ui32(v187) < base.Ui32(l1) {
		v171 = v182
		v172 = v184
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v194 = v182
	goto L31
L34:
	;
	goto L33
L35:
	;
	v209 = v194
	v211 = v159
	goto L38
L36:
	;
	v232 = v194
	goto L37
L37:
	;
	v240 = base.I32_div_s(v232-v15, int32(6))
	v257 = v240
	goto L3
L38:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)) = uint16(v215)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v217
	v219 = int32(6)
	v220 = v209 + v219
	v222 = v211 + v219
	v225 = base.I32_div_s(v222-l2, v219)
	if base.Ui32(v225) < base.Ui32(l3) {
		v209 = v220
		v211 = v222
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v232 = v220
	goto L37
L40:
	;
	goto L39
L41:
	;
	base.MemoryCopy(m, v15, l2, v242)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v245 = l1 * int32(6)
	if v245 == int32(0) {
		v257 = v12
		goto L3
	} else {
		goto L44
	}
L44:
	;
	base.MemoryCopy(m, v242+v15, l0, v245)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v15
}
func F_ginRedoRecompress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	v3 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v26 = l0 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
	if v27&int32(128) == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v72 = l0 + int32(32)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v73 == int32(0) {
		v302 = v72
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v34 = l0 + int32(32)
	v38 = F_ginCompressPostingList(m, v34, v32, int32(_a_F_ginRedoRecompress_0), v23+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v52 = v25
	v56 = int32(32)
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v56)
	v58 = l0 + v52
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
	v61 = v59 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)) = uint16(v61)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v65 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v63)+4)) = uint16(v65)
	goto L3
L7:
	;
	return
L8:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	v44 = (v40 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v46 = v44 + int32(8)
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	base.MemoryCopy(m, v34, v38, v46)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v52 = v48
	v56 = v44 + int32(40)
	goto L6
L12:
	;
	v320 = v302 - v72 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v320)
	m.G0 = v23 + int32(16)
	return
L13:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v84 = v72
	v86 = v72
	v91 = v3
	v92 = v3
	v93 = l1 + int32(2)
	v95 = v72 + v78 - int32(32)
	v101 = v3
	goto L14
L14:
	;
	v102 = int32(2)
	v103 = v93 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v104&int32(254) == v102 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	if base.B2i32(v257 == int32(0))|base.B2i32(v284 == v258) != 0 {
		v302 = v285
		goto L12
	} else {
		goto L67
	}
L16:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if base.B2i32(v150 <= v91) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v135 = v93 + int32(4)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
	v140 = int32(0)
	v144 = v140
	v145 = v140
	v146 = v135 + v136*int32(6)
	v147 = int32(1)
	v148 = v135
	v149 = v136
	goto L16
L18:
	;
	v132 = int32(0)
	v144 = v128
	v145 = v129
	v146 = v131
	v147 = v130
	v148 = v132
	v149 = v132
	goto L16
L19:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+8)))
	v113 = (v109 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v128 = v113 + int32(8)
	v129 = v103
	v130 = int32(0)
	v131 = v103 + (v113+int32(9))&int32(_a_F_ginRedoRecompress_2)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v104 == int32(4) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v124 = int32(0)
	v128 = v124
	v129 = v124
	v130 = v124
	v131 = v103
	goto L18
L23:
	;
	v156 = v84
	v158 = v86
	v163 = v91
	goto L26
L24:
	;
	v197 = v84
	v199 = v86
	v204 = v91
	goto L25
L25:
	;
	if v147 != 0 {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	v178 = (v174 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v180 = v178 + int32(8)
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v197 = v191
	v199 = v188
	v204 = v150
	goto L25
L28:
	;
	if v180 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v187 = v178
	goto L30
L30:
	;
	v188 = v158 + v180
	v191 = v156 + v187 + int32(8)
	v193 = v163 + int32(1)
	if v193 != v150 {
		v156 = v191
		v158 = v188
		v163 = v193
		goto L26
	} else {
		goto L34
	}
L31:
	;
	base.MemoryCopy(m, v158, v156, v180)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	v187 = (v182 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	goto L30
L34:
	;
	goto L27
L35:
	;
	v217 = F_ginPostingListDecode(m, v197, v23+int32(12))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	v238 = v144
	v239 = v145
	v240 = v104
	goto L37
L37:
	;
	if v197 == v95 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v222 = F_ginMergeItemPointers(m, v148, v149, v217, v219, v23+int32(8))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v228 = F_ginCompressPostingList(m, v222, v224, int32(_a_F_ginRedoRecompress_0), v23+int32(4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+6)))
	v238 = (v230+int32(1))&int32(_a_F_ginRedoRecompress_1) + int32(8)
	v239 = v228
	v240 = int32(3)
	goto L37
L41:
	;
	switch v240 - int32(1) {
	case 0:
		goto L53
	case 1:
		goto L56
	case 2:
		goto L55
	default:
		goto L54
	}
L42:
	;
	v255 = v197
	v256 = int32(0)
	v257 = v92
	v258 = v95
	goto L41
L43:
	;
	goto L44
L44:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+6)))
	v249 = (v243+int32(1))&int32(_a_F_ginRedoRecompress_1) + int32(8)
	if v92 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v255 = v197
	v256 = v249
	v257 = v92
	v258 = v95
	goto L41
L46:
	;
	goto L47
L47:
	;
	v250 = v95 - v197
	v251 = F_palloc(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if v250 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v251, v197, v250)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v255 = v251
	v256 = v249
	v257 = v251
	v258 = v251 + v250
	goto L41
L52:
	;
	v288 = v101 + int32(1)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if base.Ui32(v288) < base.Ui32(v289) {
		v84 = v284
		v86 = v285
		v91 = v286
		v92 = v257
		v93 = v146
		v95 = v258
		v101 = v288
		goto L14
	} else {
		goto L66
	}
L53:
	;
	v284 = v255 + v256
	v285 = v199
	v286 = v204 + int32(1)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L63
	}
L55:
	;
	if v238 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	if v238 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v199, v239, v238)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v284 = v255
	v285 = v199 + v238
	v286 = v204
	goto L52
L60:
	;
	base.MemoryCopy(m, v199, v239, v238)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v284 = v255 + v256
	v285 = v199 + v238
	v286 = v204 + int32(1)
	goto L52
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v240
	F_errmsg_internal(m, int32(_a_F_ginRedoRecompress_3), v23)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_ginRedoRecompress_4), int32(298), int32(_a_F_ginRedoRecompress_5))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L15
L67:
	;
	v295 = v258 - v284
	if v295 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v285, v284, v295)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v302 = v295 + v285
	goto L12
}
func F_gin_enum_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 == int32(0) {
		if v4 != 0 {
			v10 = int32(-1)
		} else {
			v10 = int32(0)
		}
		return v10
	} else {
		if v4 == int32(0) {
			return int32(1)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = F_CallerFInfoFunctionCall2(m, int32(3795), v17, v18, v5, v4)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	}
}
func F_gin_extract_hstore_query(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	switch v19 - int32(7) {
	case 0:
		goto L5
	default:
		goto L3
	case 2:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v206
L3:
	;
	if v19&int32(_a_F_gin_extract_hstore_query_0) != int32(10) {
		goto L1
	} else {
		goto L29
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = F_pg_detoast_datum_packed(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_DirectFunctionCall2Coll(m, int32(_a_F_gin_extract_hstore_query_1), int32(0), v24, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v25 != 0 {
		v206 = v25
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v206 = int32(0)
	goto L2
L9:
	;
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v35
	v38 = v33 + v35
	v40 = F_palloc(m, int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v43 = int32(1)
	v44 = v42 & v43
	if v42 == v43 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v73 = v71 + int32(5)
	v74 = F_palloc(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L22
	}
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v50 == int32(18) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v61 = int32(1)
	if v44 != 0 {
		v71 = int32(base.Ui32(v42)>>(uint(v61)%32)) - v61
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v53 = int32(16)
	goto L17
L16:
	;
	v53 = int32(0)
	goto L17
L17:
	;
	if base.Ui32((v50-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = int32(4)
	goto L20
L19:
	;
	v60 = v53
	goto L20
L20:
	;
	v71 = v60
	goto L11
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L22:
	;
	v76 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v73 << (uint(int32(2)) % 32)
	v81 = int32(0)
	if base.B2i32(v71 == v81)|base.B2i32(v71 <= v81) == v81 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v44 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v74
	v206 = v40
	goto L2
L26:
	;
	v92 = v38
	goto L28
L27:
	;
	v92 = v33 + int32(4)
	goto L28
L28:
	;
	base.MemoryCopy(m, v74+int32(5), v92, v71)
	goto L25
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v100 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_deconstruct_array_builtin(m, v100, int32(25), v15+int32(12), v15+int32(8), v15+int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v114 = F_palloc(m, v111<<(uint(int32(2))%32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) < v116 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = int32(0)
	v121 = v2
	v123 = v116
	goto L36
L34:
	;
	v183 = v2
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v183
	if base.B2i32(v19 != int32(11))|v183 != 0 {
		v206 = v114
		goto L2
	} else {
		goto L46
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v120))))
	if v134 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v183 = v173
	goto L35
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v138 = int32(2)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v120<<(uint(v138)%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v144 = int32(base.Ui32(v142) >> (uint(v138) % 32))
	v146 = v144 + int32(1)
	v147 = F_palloc(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	v173 = v121
	v174 = v123
	goto L40
L40:
	;
	v180 = v120 + int32(1)
	if v180 < v174 {
		v120 = v180
		v121 = v173
		v123 = v174
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v149 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+4)) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v146 << (uint(int32(2)) % 32)
	if base.Ui32(v142) < base.Ui32(int32(20)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+v121<<(uint(int32(2))%32)))) = v147
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v173 = v121 + int32(1)
	v174 = v170
	goto L40
L43:
	;
	v157 = v144 - int32(4)
	if v157 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	base.MemoryCopy(m, v147+int32(5), v141+int32(4), v157)
	goto L42
L45:
	;
	goto L37
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v206 = v114
	goto L2
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	F_errmsg_internal(m, int32(_a_F_gin_extract_hstore_query_2), v15)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_gin_extract_hstore_query_3), int32(141), int32(_a_F_gin_extract_hstore_query_4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_query_anyenum(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(0), int32(_a_F_gin_extract_query_anyenum_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_bytea(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_bytea_0), int32(2847))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_int4(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(-2147483648), int32(2097))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_time(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13908(m, l0, int64(0), int32(1429))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_trgm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	switch v2 - int32(3) {
	case 0:
		v23 = F_gin_extract_value_trgm(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_extract_trgm_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_trgm_1), int32(30), int32(_a_F_gin_extract_trgm_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v5 = F_gin_extract_query_trgm(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_gin_extract_value_trgm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
		v15 = int32(1)
		v16 = v8 + v15
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v21 = v19 & v15
		if v21 != 0 {
			v22 = v16
		} else {
			v22 = v8 + int32(4)
		}
		if v19 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
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
		v50 = F_generate_trgm(m, v22, v49)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			v56 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(5)
			if base.Ui32(int32(3)) <= base.Ui32(v56) {
				v60 = base.I32_div_u_s(v56, int32(3))
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v60
				v62 = int32(1)
				if base.Ui32(v60) <= base.Ui32(v62) {
					v65 = v62
				} else {
					v65 = v60
				}
				v71 = F_palloc(m, v60<<(uint(int32(2))%32))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = int32(0)
					v75 = v50 + int32(5)
					for {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						*(*int32)(unsafe.Add(mBase, uint32(v71+v73<<(uint(int32(2))%32)))) = v82 | (v83<<(uint(int32(8))%32) | v86<<(uint(int32(16))%32))
						v95 = v73 + int32(1)
						if v95 != v65 {
							v73 = v95
							v75 = v75 + int32(3)
							continue
						} else {
							break
						}
						break
					}
					v100 = v71
					return v100
				}
			} else {
				v100 = int32(0)
				return v100
			}
		}
	}
}
func F_gin_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = l0 - int32(16)
	if base.Ui32(v3) <= base.Ui32(int32(143)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_c_F_gin_identify[0])))
		v12 = v10
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_gin_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(_a_F_gin_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v7)+6)))
	if v15&int32(4) != 0 {
		v20 = int32(0)
		base.MemoryFill(m, l0+int32(24), v20, int32(_a_F_gin_mask_1))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
		return
	} else {
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(int32(25)) <= base.Ui32(v25) {
			F_mask_unused_space(m, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_gin_redo(m *base.Module, l0 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int64
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int64
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int64
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int64
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int64
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int64
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int64
	_ = v796
	var v810 int32
	_ = v810
	var v812 int64
	_ = v812
	var v814 int64
	_ = v814
	var v816 int64
	_ = v816
	var v818 int64
	_ = v818
	var v820 int64
	_ = v820
	var v822 int64
	_ = v822
	var v824 int64
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int64
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1230 int64
	_ = v1230
	var v1244 int32
	_ = v1244
	var v1246 int64
	_ = v1246
	var v1248 int64
	_ = v1248
	var v1250 int64
	_ = v1250
	var v1252 int64
	_ = v1252
	var v1254 int64
	_ = v1254
	var v1256 int64
	_ = v1256
	var v1258 int64
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1277 int32
	_ = v1277
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1357 int32
	_ = v1357
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = int32(_a_F_gin_redo_0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0])) = v23
	v26 = v20 & int32(240)
	switch int32(base.Ui32(v26-int32(16)) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L12
	case 5:
		goto L11
	case 6:
		goto L10
	case 7:
		goto L9
	case 8:
		goto L13
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L19
	} else {
		goto L346
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L19
	} else {
		goto L343
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L19
	} else {
		goto L340
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L19
	} else {
		goto L337
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L19
	} else {
		goto L334
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L19
	} else {
		goto L331
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L19
	} else {
		goto L328
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0])) = v18
	v1373 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[1]))
	F_MemoryContextReset(m, v1373)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L19
	} else {
		goto L327
	}
L9:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v1180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1182 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L19
	} else {
		goto L298
	}
L10:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v1027 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1029 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L19
	} else {
		goto L264
	}
L11:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v746 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v748 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L19
	} else {
		goto L208
	}
L12:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v587 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v591 = F_XLogReadBufferForRedo(m, l0, int32(2), v13+int32(-8))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L19
	} else {
		goto L163
	}
L13:
	;
	v514 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v518 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L19
	} else {
		goto L140
	}
L14:
	;
	v507 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L19
	} else {
		goto L137
	}
L15:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+24)))
	if v419&int32(2) != 0 {
		goto L112
	} else {
		goto L113
	}
L16:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103))))
	v106 = v104 & int32(2)
	if v106 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v54 = int32(131)
	if v34 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	return
L20:
	;
	if v34 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v34^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L18
L22:
	;
	goto L23
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v53 = v47 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	F_PageInit(m, v72, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+16)))
	v77 = v72 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v54)
	goto L24
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(v34^int32(-1))<<(uint(int32(2))%32))))
	v72 = v64
	goto L25
L27:
	;
	goto L28
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v72 = v66 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+64))
	base.MemoryCopy(m, v53+int32(32), v85+int32(4), v81)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
	*(*uint32)(unsafe.Add(mBase, uint32(v53)+4)) = uint32(v32)
	v92 = int64(base.Ui64(v32) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v53))) = uint32(v92)
	v95 = v89 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)) = uint16(v95)
	F_MarkBufferDirty(m, v34)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v34)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L8
L34:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+6)))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+8)))
	v114 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	v158 = int32(-1)
	goto L36
L36:
	;
	v164 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-24))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L50
	}
L37:
	;
	if v114 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v118 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v152 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v138 = v137 + v136
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+6)))
	v141 = v139 & int32(_a_F_gin_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+6)) = uint16(v141)
	*(*uint32)(unsafe.Add(mBase, uint32(v136)+4)) = uint32(v101)
	v145 = int64(base.Ui64(v101) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v136))) = uint32(v145)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L45
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(v118^int32(-1))<<(uint(int32(2))%32))))
	v136 = v128
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v136 = v130 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L45:
	;
	goto L40
L46:
	;
	F_UnlockReleaseBuffer(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L19
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v158 = v109<<(uint(int32(16))%32) | v110
	goto L36
L49:
	;
	goto L48
L50:
	;
	if v164 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v168 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v413 == int32(0) {
		goto L8
	} else {
		goto L110
	}
L54:
	;
	v187 = int32(0)
	v189 = v13 + int32(-28)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+72))
	if v192 < v187 {
		v214 = v187
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v168^int32(-1))<<(uint(int32(2))%32))))
	v186 = v178
	goto L54
L56:
	;
	goto L57
L57:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v186 = v180 + v168<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L58:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v219&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	v217 = v214
	goto L58
L60:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+int32(0))+76)))
	if v197 != int32(1) {
		v214 = v187
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v201 = v191 + int32(76)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+43)))
	if v202 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v189 == int32(0) {
		v214 = v187
		goto L59
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v189 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v207
	v217 = v207
	goto L58
L66:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v210
	goto L68
L67:
	;
	goto L68
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v201)+44))
	v214 = v212
	goto L59
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186))) = base.I64_rotr(v101, int64(32))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_MarkBufferDirty(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L19
	} else {
		goto L109
	}
L70:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	*(*int32)(unsafe.Add(mBase, uint32(v239+v342*int32(10))+22)) = base.I32_rotr(v158, int32(16))
	v350 = v217 + int32(2)
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+16)))
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239+v351)+4)))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	if v354 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L71:
	;
	if v218 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	if v218 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	if v106 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L75:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v218^int32(-1))<<(uint(int32(2))%32))))
	v239 = v231
	goto L74
L76:
	;
	goto L77
L77:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v239 = v233 + v218<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	F_ginRedoRecompress(m, v239, v217)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	if v158 != int32(-1) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247+(v218^int32(-1))<<(uint(int32(2))%32))))
	v261 = v253
	goto L80
L82:
	;
	goto L83
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v261 = v255 + v218<<(uint(int32(13))%32) + int32(-8192)
	goto L80
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261+v262<<(uint(int32(2))%32))+20))
	v271 = v261 + v268&int32(_a_F_gin_redo_3)
	v272 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v271)+4)) = uint16(v272)
	*(*uint16)(unsafe.Add(mBase, uint32(v271)+2)) = uint16(v158)
	v276 = int32(base.Ui32(v158) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v271))) = uint16(v276)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+2)))
	if v279 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_PageIndexTupleDelete(m, v261, v262)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L19
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+10)))
	v290 = F_PageAddItemExtended(m, v261, v217+int32(4), v286&int32(_a_F_gin_redo_4), v262, int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L19
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	if v290 != 0 {
		goto L69
	} else {
		goto L92
	}
L92:
	;
	v293 = v13 + int32(-20)
	if v218 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L19
	} else {
		goto L98
	}
L94:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v316
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-4)))) = v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-8)))) = v322
	goto L93
L95:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[4]))
	v315 = v302 + (v218^int32(-1))<<(uint(int32(6))%32)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[5]))
	v315 = v309 + v218<<(uint(int32(6))%32) + int32(-64)
	goto L94
L98:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+20)) = v330
	F_errmsg_internal(m, int32(_a_F_gin_redo_5), v13+int32(-48))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(104), int32(_a_F_gin_redo_7))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L19
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v380)+8)) = uint16(v382)
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v350)))
	*(*int64)(unsafe.Add(mBase, uint32(v380))) = v384
	v388 = v353 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v239+v381)+4)) = uint16(v388)
	v393 = v353*int32(10) + int32(42)
	*(*uint16)(unsafe.Add(mBase, uint32(v239)+12)) = uint16(v393)
	goto L69
L102:
	;
	v380 = v239 + v353*int32(10) + int32(32)
	v381 = v351
	goto L101
L103:
	;
	goto L104
L104:
	;
	v364 = v239 + v354*int32(10)
	v366 = v364 + int32(22)
	if v353+int32(1) == v354 {
		v380 = v366
		v381 = v351
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v371 = int32(10)
	v374 = (v353-v354)*v371 + v371
	if v374 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	base.MemoryCopy(m, v364+int32(32), v366, v374)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+16)))
	v380 = v366
	v381 = v378
	goto L101
L109:
	;
	goto L53
L110:
	;
	F_UnlockReleaseBuffer(m, v413)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	goto L8
L112:
	;
	v475 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L19
	} else {
		goto L125
	}
L113:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v426 = F_XLogReadBufferForRedo(m, l0, int32(3), v13+int32(-20))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	if v426 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v430 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L117
L117:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v464 == int32(0) {
		goto L112
	} else {
		goto L123
	}
L118:
	;
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+16)))
	v450 = v449 + v448
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+6)))
	v453 = v451 & int32(_a_F_gin_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v450)+6)) = uint16(v453)
	*(*uint32)(unsafe.Add(mBase, uint32(v448)+4)) = uint32(v422)
	v457 = int64(base.Ui64(v422) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v448))) = uint32(v457)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L19
	} else {
		goto L122
	}
L119:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v434+(v430^int32(-1))<<(uint(int32(2))%32))))
	v448 = v440
	goto L118
L120:
	;
	goto L121
L121:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v448 = v442 + v430<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	goto L117
L123:
	;
	F_UnlockReleaseBuffer(m, v464)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L19
	} else {
		goto L124
	}
L124:
	;
	goto L112
L125:
	;
	if v475 != int32(2) {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v482 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-4))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	if v482 != int32(2) {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	if v419&int32(4) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v491 = F_XLogReadBufferForRedo(m, l0, int32(2), v13+int32(-8))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L19
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	F_UnlockReleaseBuffer(m, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L19
	} else {
		goto L135
	}
L132:
	;
	if v491 != int32(2) {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	F_UnlockReleaseBuffer(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_UnlockReleaseBuffer(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	goto L8
L137:
	;
	if v507 != int32(2) {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_UnlockReleaseBuffer(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L19
	} else {
		goto L139
	}
L139:
	;
	goto L8
L140:
	;
	if v518 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v522 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v581 == int32(0) {
		goto L8
	} else {
		goto L161
	}
L144:
	;
	v541 = int32(0)
	v543 = v13 + int32(-4)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+72))
	if v546 < v541 {
		v568 = v541
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526+(v522^int32(-1))<<(uint(int32(2))%32))))
	v540 = v532
	goto L144
L146:
	;
	goto L147
L147:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v540 = v534 + v522<<(uint(int32(13))%32) + int32(-8192)
	goto L144
L148:
	;
	F_ginRedoRecompress(m, v540, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L19
	} else {
		goto L159
	}
L149:
	;
	v571 = v568
	goto L148
L150:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+int32(0))+76)))
	if v551 != int32(1) {
		v568 = v541
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v555 = v545 + int32(76)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+43)))
	if v556 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v543 == int32(0) {
		v568 = v541
		goto L149
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v543 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v561 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v561
	v571 = v561
	goto L148
L156:
	;
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v564
	goto L158
L157:
	;
	goto L158
L158:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v555)+44))
	v568 = v566
	goto L149
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v540))) = base.I64_rotr(v514, int64(32))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	goto L143
L161:
	;
	F_UnlockReleaseBuffer(m, v581)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	goto L8
L163:
	;
	if v591 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v595 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v629 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L19
	} else {
		goto L172
	}
L167:
	;
	v614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v613)+16)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v614+v613))) = v616
	*(*uint32)(unsafe.Add(mBase, uint32(v613)+4)) = uint32(v587)
	v620 = int64(base.Ui64(v587) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v613))) = uint32(v620)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	F_MarkBufferDirty(m, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L19
	} else {
		goto L171
	}
L168:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v599+(v595^int32(-1))<<(uint(int32(2))%32))))
	v613 = v605
	goto L167
L169:
	;
	goto L170
L170:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v613 = v607 + v595<<(uint(int32(13))%32) + int32(-8192)
	goto L167
L171:
	;
	goto L166
L172:
	;
	if v629 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v633 < int32(0) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L175
L175:
	;
	v672 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-4))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L19
	} else {
		goto L181
	}
L176:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v651)+16)))
	v653 = v652 + v651
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v653)+6)))
	v656 = v654 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v653)+6)) = uint16(v656)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v651)+4)) = uint32(v587)
	v661 = int64(base.Ui64(v587) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v651))) = uint32(v661)
	*(*int32)(unsafe.Add(mBase, uint32(v651)+20)) = v658
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L19
	} else {
		goto L180
	}
L177:
	;
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v637+(v633^int32(-1))<<(uint(int32(2))%32))))
	v651 = v643
	goto L176
L178:
	;
	goto L179
L179:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v651 = v645 + v633<<(uint(int32(13))%32) + int32(-8192)
	goto L176
L180:
	;
	goto L175
L181:
	;
	if v672 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v676 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	goto L184
L184:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v734 != 0 {
		goto L197
	} else {
		goto L198
	}
L185:
	;
	v695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v586))))
	v698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694)+16)))
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694+v698)+4)))
	if v700 != v695 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v680+(v676^int32(-1))<<(uint(int32(2))%32))))
	v694 = v686
	goto L185
L187:
	;
	goto L188
L188:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v694 = v688 + v676<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v694))) = base.I64_rotr(v587, int64(32))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	F_MarkBufferDirty(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L19
	} else {
		goto L196
	}
L190:
	;
	v704 = (v700 - v695) * int32(10)
	if v704 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v717 = v698
	goto L192
L192:
	;
	v720 = v700 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v717+v694)+4)) = uint16(v720)
	v725 = v700*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v694)+12)) = uint16(v725)
	goto L189
L193:
	;
	v707 = v694 + v695*int32(10)
	base.MemoryCopy(m, v707+int32(22), v707+int32(32), v704)
	goto L195
L194:
	;
	goto L195
L195:
	;
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694)+16)))
	v717 = v714
	goto L192
L196:
	;
	goto L184
L197:
	;
	F_UnlockReleaseBuffer(m, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L19
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v737 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L199
L201:
	;
	F_UnlockReleaseBuffer(m, v737)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L19
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v740 == int32(0) {
		goto L8
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	F_UnlockReleaseBuffer(m, v740)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	goto L8
L207:
	;
	if v748 < int32(0) {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	if v748 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v753+(v748^int32(-1))<<(uint(int32(2))%32))))
	v767 = v759
	goto L207
L210:
	;
	goto L211
L211:
	;
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v767 = v761 + v748<<(uint(int32(13))%32) + int32(-8192)
	goto L207
L212:
	;
	v812 = *(*int64)(unsafe.Add(mBase, uint32(v745)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+72)) = v812
	v814 = *(*int64)(unsafe.Add(mBase, uint32(v745)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+64)) = v814
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v745)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+56)) = v816
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v745)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+48)) = v818
	v820 = *(*int64)(unsafe.Add(mBase, uint32(v745)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+40)) = v820
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v745)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+32)) = v822
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v745)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+24)) = v824
	*(*int64)(unsafe.Add(mBase, uint32(v767))) = base.I64_rotr(v746, int64(32))
	F_MarkBufferDirty(m, v748)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L19
	} else {
		goto L217
	}
L213:
	;
	v788 = int32(8)
	F_PageInit(m, v786, int32(_a_F_gin_redo_1), v788)
	mBase = m.M
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786)+16)))
	v791 = v786 + v790
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v791)+6)) = uint16(v788)
	v796 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v786)+64)) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v786)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v786)+32)) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v786)+40)) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v786)+48)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v786)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v786)+72)) = int32(2)
	v810 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v786)+12)) = uint16(v810)
	goto L212
L214:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v772+(v748^int32(-1))<<(uint(int32(2))%32))))
	v786 = v778
	goto L213
L215:
	;
	goto L216
L216:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v786 = v780 + v748<<(uint(int32(13))%32) + int32(-8192)
	goto L213
L217:
	;
	v833 = base.I32_wrap_i64(int64(base.Ui64(v746) >> (uint(int64(32)) % 64)))
	v834 = base.I32_wrap_i64(v746)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v745)+80))
	if int32(0) < v835 {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	F_UnlockReleaseBuffer(m, v748)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L19
	} else {
		goto L262
	}
L219:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v1007 == int32(0) {
		goto L218
	} else {
		goto L260
	}
L220:
	;
	v841 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L19
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v745)+72))
	if v959 == int32(-1) {
		goto L218
	} else {
		goto L252
	}
L223:
	;
	if v841 != 0 {
		goto L219
	} else {
		goto L224
	}
L224:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v843 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v864 = v13 + int32(-4)
	v865 = int32(0)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v866)+72))
	if v867 < int32(1) {
		v889 = v865
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v847+(v843^int32(-1))<<(uint(int32(2))%32))))
	v861 = v853
	goto L225
L227:
	;
	goto L228
L228:
	;
	v855 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v861 = v855 + v843<<(uint(int32(13))%32) + int32(-8192)
	goto L225
L229:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v745)+80))
	if int32(0) < v893 {
		goto L240
	} else {
		goto L241
	}
L230:
	;
	v892 = v889
	goto L229
L231:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866+int32(52))+76)))
	if v872 != int32(1) {
		v889 = v865
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v876 = v866 + int32(128)
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+43)))
	if v877 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v864 == int32(0) {
		v889 = v865
		goto L230
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if v864 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v882 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v864))) = v882
	v892 = v882
	goto L229
L237:
	;
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v864))) = v885
	goto L239
L238:
	;
	goto L239
L239:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v876)+44))
	v889 = v887
	goto L230
L240:
	;
	v896 = int32(1)
	v897 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v861)+12)))
	if base.Ui32(v897) < base.Ui32(int32(25)) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v861)+16)))
	v949 = v861 + v948
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v949)+4)))
	v952 = v950 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v949)+4)) = uint16(v952)
	*(*int32)(unsafe.Add(mBase, uint32(v861)+4)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v861))) = v833
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L19
	} else {
		goto L251
	}
L243:
	;
	v906 = v896
	goto L245
L244:
	;
	v906 = int32(base.Ui32(v897+int32(_a_F_gin_redo_8))>>(uint(int32(2))%32)) + v896
	goto L245
L245:
	;
	v907 = v906
	v908 = v892
	v909 = int32(0)
	goto L246
L246:
	;
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v908)+6)))
	v921 = v919 & int32(_a_F_gin_redo_4)
	v925 = F_PageAddItemExtended(m, v861, v908, v921, v907&int32(_a_F_gin_redo_9), int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L19
	} else {
		goto L248
	}
L247:
	;
	goto L242
L248:
	;
	if v925 == int32(0) {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v929 = int32(1)
	v933 = v909 + v929
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v745)+80))
	if v933 < v934 {
		v907 = v907 + v929
		v908 = v908 + v921
		v909 = v933
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	goto L219
L252:
	;
	v965 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L19
	} else {
		goto L253
	}
L253:
	;
	if v965 != 0 {
		goto L219
	} else {
		goto L254
	}
L254:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v967 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v986 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+16)))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v745)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v986+v985))) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v985)+4)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v985))) = v833
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v992)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L19
	} else {
		goto L259
	}
L256:
	;
	v971 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v971+(v967^int32(-1))<<(uint(int32(2))%32))))
	v985 = v977
	goto L255
L257:
	;
	goto L258
L258:
	;
	v979 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v985 = v979 + v967<<(uint(int32(13))%32) + int32(-8192)
	goto L255
L259:
	;
	goto L219
L260:
	;
	F_UnlockReleaseBuffer(m, v1007)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L19
	} else {
		goto L261
	}
L261:
	;
	goto L218
L262:
	;
	goto L8
L263:
	;
	v1049 = int32(16)
	if v1029 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L264:
	;
	if v1029 < int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1034+(v1029^int32(-1))<<(uint(int32(2))%32))))
	v1048 = v1040
	goto L263
L266:
	;
	goto L267
L267:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1048 = v1042 + v1029<<(uint(int32(13))%32) + int32(-8192)
	goto L263
L268:
	;
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1048)+16)))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	*(*int32)(unsafe.Add(mBase, uint32(v1048+v1076))) = v1078
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1048)+16)))
	if v1078 != int32(-1) {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	F_PageInit(m, v1067, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v1071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1067)+16)))
	v1072 = v1067 + v1071
	*(*int32)(unsafe.Add(mBase, uint32(v1072))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1072)+6)) = uint16(v1049)
	goto L268
L270:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1053+(v1029^int32(-1))<<(uint(int32(2))%32))))
	v1067 = v1059
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1067 = v1061 + v1029<<(uint(int32(13))%32) + int32(-8192)
	goto L269
L273:
	;
	v1092 = v1080
	v1093 = int32(0)
	goto L275
L274:
	;
	v1085 = v1080 + v1048
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1085)+6)))
	v1088 = v1086 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1085)+6)) = uint16(v1088)
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1048)+16)))
	v1092 = v1090
	v1093 = int32(1)
	goto L275
L275:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1092+v1048)+4)) = uint16(v1093)
	v1096 = int32(0)
	v1098 = v13 + int32(-20)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+72))
	if v1101 < v1096 {
		v1123 = v1096
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+4))
	if int32(0) < v1128 {
		goto L287
	} else {
		goto L288
	}
L277:
	;
	v1126 = v1123
	goto L276
L278:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100+int32(0))+76)))
	if v1106 != int32(1) {
		v1123 = v1096
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1110 = v1100 + int32(76)
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+43)))
	if v1111 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	if v1098 == int32(0) {
		v1123 = v1096
		goto L277
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v1098 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1098))) = v1116
	v1126 = v1116
	goto L276
L284:
	;
	v1119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1098))) = v1119
	goto L286
L285:
	;
	goto L286
L286:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+44))
	v1123 = v1121
	goto L277
L287:
	;
	v1131 = int32(1)
	v1132 = v1126
	v1135 = int32(0)
	goto L290
L288:
	;
	goto L289
L289:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1048))) = base.I64_rotr(v1027, int64(32))
	F_MarkBufferDirty(m, v1029)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L19
	} else {
		goto L295
	}
L290:
	;
	v1143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132)+6)))
	v1145 = v1143 & int32(_a_F_gin_redo_4)
	v1149 = F_PageAddItemExtended(m, v1048, v1132, v1145, v1131&int32(_a_F_gin_redo_9), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L19
	} else {
		goto L292
	}
L291:
	;
	goto L289
L292:
	;
	if v1149 == int32(0) {
		goto L2
	} else {
		goto L293
	}
L293:
	;
	v1153 = int32(1)
	v1157 = v1135 + v1153
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+4))
	if v1157 < v1158 {
		v1131 = v1131 + v1153
		v1132 = v1132 + v1145
		v1135 = v1157
		goto L290
	} else {
		goto L294
	}
L294:
	;
	goto L291
L295:
	;
	F_UnlockReleaseBuffer(m, v1029)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L19
	} else {
		goto L296
	}
L296:
	;
	goto L8
L297:
	;
	if v1182 < int32(0) {
		goto L304
	} else {
		goto L305
	}
L298:
	;
	if v1182 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1187+(v1182^int32(-1))<<(uint(int32(2))%32))))
	v1201 = v1193
	goto L297
L300:
	;
	goto L301
L301:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1201 = v1195 + v1182<<(uint(int32(13))%32) + int32(-8192)
	goto L297
L302:
	;
	v1246 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+72)) = v1246
	v1248 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+64)) = v1248
	v1250 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+56)) = v1250
	v1252 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+48)) = v1252
	v1254 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+40)) = v1254
	v1256 = *(*int64)(unsafe.Add(mBase, uint32(v1179)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+32)) = v1256
	v1258 = *(*int64)(unsafe.Add(mBase, uint32(v1179)))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+24)) = v1258
	*(*int64)(unsafe.Add(mBase, uint32(v1201))) = base.I64_rotr(v1180, int64(32))
	F_MarkBufferDirty(m, v1182)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L19
	} else {
		goto L307
	}
L303:
	;
	v1222 = int32(8)
	F_PageInit(m, v1220, int32(_a_F_gin_redo_1), v1222)
	mBase = m.M
	v1224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1220)+16)))
	v1225 = v1220 + v1224
	*(*int32)(unsafe.Add(mBase, uint32(v1225))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1225)+6)) = uint16(v1222)
	v1230 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+64)) = v1230
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+32)) = v1230
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+40)) = v1230
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+48)) = v1230
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+72)) = int32(2)
	v1244 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v1220)+12)) = uint16(v1244)
	goto L302
L304:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1206+(v1182^int32(-1))<<(uint(int32(2))%32))))
	v1220 = v1212
	goto L303
L305:
	;
	goto L306
L306:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1220 = v1214 + v1182<<(uint(int32(13))%32) + int32(-8192)
	goto L303
L307:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+56))
	if int32(0) < v1265 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1277 = int32(0)
	goto L311
L309:
	;
	goto L310
L310:
	;
	F_UnlockReleaseBuffer(m, v1182)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L19
	} else {
		goto L326
	}
L311:
	;
	v1286 = v1277 + int32(1)
	v1289 = F_XLogInitBufferForRedo(m, l0, v1286&int32(255))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L19
	} else {
		goto L314
	}
L312:
	;
	goto L310
L313:
	;
	v1309 = int32(4)
	if v1289 < int32(0) {
		goto L320
	} else {
		goto L321
	}
L314:
	;
	if v1289 < int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1294+(v1289^int32(-1))<<(uint(int32(2))%32))))
	v1308 = v1300
	goto L313
L316:
	;
	goto L317
L317:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1308 = v1302 + v1289<<(uint(int32(13))%32) + int32(-8192)
	goto L313
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+4)) = base.I32_wrap_i64(v1180)
	*(*int32)(unsafe.Add(mBase, uint32(v1308))) = base.I32_wrap_i64(int64(base.Ui64(v1180) >> (uint(int64(32)) % 64)))
	F_MarkBufferDirty(m, v1289)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L19
	} else {
		goto L323
	}
L319:
	;
	F_PageInit(m, v1327, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1327)+16)))
	v1332 = v1327 + v1331
	*(*int32)(unsafe.Add(mBase, uint32(v1332))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1332)+6)) = uint16(v1309)
	goto L318
L320:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1313+(v1289^int32(-1))<<(uint(int32(2))%32))))
	v1327 = v1319
	goto L319
L321:
	;
	goto L322
L322:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1327 = v1321 + v1289<<(uint(int32(13))%32) + int32(-8192)
	goto L319
L323:
	;
	F_UnlockReleaseBuffer(m, v1289)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L19
	} else {
		goto L324
	}
L324:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+56))
	if v1286 < v1342 {
		v1277 = v1286
		goto L311
	} else {
		goto L325
	}
L325:
	;
	goto L312
L326:
	;
	goto L8
L327:
	;
	m.G0 = v15 - int32(-64)
	return
L328:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_10), int32(0))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L19
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(419), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L19
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_12), int32(0))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L19
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(422), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L19
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_13), int32(0))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L19
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(427), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L19
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_14), int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L19
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(446), int32(_a_F_gin_redo_15))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_16), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L19
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(579), int32(_a_F_gin_redo_17))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L19
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_16), int32(0))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L19
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(661), int32(_a_F_gin_redo_18))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L19
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v26
	F_errmsg_internal(m, int32(_a_F_gin_redo_19), v15)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L19
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(768), int32(_a_F_gin_redo_20))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L19
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_trgm_triconsistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v36 float64
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v481 float32
	_ = v481
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	if base.Ui32(int32(11)) < base.Ui32(v16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v490
L2:
	;
	if base.F64_le(v36, base.F64_promote_f32(base.F32_div(v481, base.F32_convert_i32_s(v19)))) != 0 {
		goto L91
	} else {
		goto L92
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L13
	} else {
		goto L88
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = int32(1) << (uint(v16) % 32)
	if v22&int32(642) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v154 = v19 & int32(3)
	v155 = F_palloc(m, v19)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L37
	}
L6:
	;
	v131 = int32(0)
	v132 = int32(2)
	if v19 <= v131 {
		v490 = v132
		goto L1
	} else {
		goto L30
	}
L7:
	;
	if v22&int32(2072) != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v36 = F_index_strategy_get_limit(m, v16)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v22&int32(96) == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if int32(0) < v19 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v490 = int32(2)
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	if int32(0) < v19 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = v19 & int32(3)
	v44 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v19) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if v19 != 0 {
		v481 = float32(0)
		goto L2
	} else {
		goto L29
	}
L18:
	;
	v481 = base.F32_convert_i32_u(v119)
	goto L2
L19:
	;
	v50 = v44
	v51 = v44
	v55 = int32(0)
	goto L22
L20:
	;
	v85 = v44
	v86 = v44
	goto L21
L21:
	;
	v97 = v85
	v98 = v86
	v100 = int32(0)
	goto L26
L22:
	;
	v61 = v50 + v20
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v63 = int32(0)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	v77 = v51 + base.B2i32(v62 != v63) + base.B2i32(v66 != v63) + base.B2i32(v70 != v63) + base.B2i32(v74 != v63)
	v78 = int32(4)
	v79 = v50 + v78
	v81 = v55 + v78
	if v81 != v19&int32(2147483644) {
		v50 = v79
		v51 = v77
		v55 = v81
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v43 == int32(0) {
		v119 = v77
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v85 = v79
	v86 = v77
	goto L21
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v20))))
	v112 = v98 + base.B2i32(v109 != int32(0))
	v113 = int32(1)
	v116 = v100 + v113
	if v116 != v43 {
		v97 = v97 + v113
		v98 = v112
		v100 = v116
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v119 = v112
	goto L18
L28:
	;
	goto L27
L29:
	;
	v490 = int32(0)
	goto L1
L30:
	;
	v135 = v131
	goto L31
L31:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v20))))
	if v147 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v490 = int32(0)
	goto L1
L33:
	;
	v149 = v135 + int32(1)
	if v19 != v149 {
		v135 = v149
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v490 = v132
	goto L1
L37:
	;
	v157 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v19) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v257 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v267 != 0 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v163 = v157
	v166 = int32(0)
	goto L42
L40:
	;
	v211 = v157
	goto L41
L41:
	;
	v223 = v211
	v226 = int32(0)
	goto L46
L42:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v20))))
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v155))) = uint8(base.B2i32(v176 != v177))
	v181 = v163 | int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v181))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v181))) = uint8(base.B2i32(v184 != v177))
	v189 = v163 | int32(2)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v189))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v189))) = uint8(base.B2i32(v192 != v177))
	v197 = v163 | int32(3)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v197))) = uint8(base.B2i32(v200 != v177))
	v204 = int32(4)
	v205 = v163 + v204
	v207 = v166 + v204
	if v207 != v19&int32(2147483644) {
		v163 = v205
		v166 = v207
		goto L42
	} else {
		goto L44
	}
L43:
	;
	if v154 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v211 = v205
	goto L41
L46:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223+v20))))
	*(*uint8)(unsafe.Add(mBase, uint32(v223+v155))) = uint8(base.B2i32(v236 != int32(0)))
	v240 = int32(1)
	v243 = v226 + v240
	if v243 != v154 {
		v223 = v223 + v240
		v226 = v243
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L38
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v155)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L13
	} else {
		goto L84
	}
L50:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	base.MemoryFill(m, v268, int32(0), v267)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if v271 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	base.MemoryFill(m, v272, int32(0), v271)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if int32(0) < v275 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v282 = v257
	v283 = v257
	goto L59
L57:
	;
	goto L58
L58:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	v349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v349)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = int32(0)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v361 = v349
	v366 = v257
	goto L71
L59:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v278+v282<<(uint(int32(2))%32))))
	v295 = v294 + v283
	if v294 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v334 = v282 + int32(1)
	if v334 != v275 {
		v282 = v334
		v283 = v295
		goto L59
	} else {
		goto L69
	}
L62:
	;
	v300 = v283
	goto L63
L63:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v300))))
	if v311 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317+v282))) = uint8(v319)
	goto L61
L65:
	;
	v315 = v300 + int32(1)
	if v315 < v295 {
		v300 = v315
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L61
L69:
	;
	goto L60
L70:
	;
	goto L49
L71:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v351+v366<<(uint(int32(2))%32))))
	v374 = v354 + v371<<(uint(int32(3))%32)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if int32(0) < v375 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v449 = int32(0)
	goto L70
L73:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v383 = int32(0)
	v386 = v361
	goto L76
L74:
	;
	v426 = v361
	goto L75
L75:
	;
	v434 = v366 + int32(1)
	if v434 < v426 {
		v361 = v426
		v366 = v434
		goto L71
	} else {
		goto L83
	}
L76:
	;
	v395 = v379 + v383<<(uint(int32(3))%32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378+v396))))
	if v398 != int32(1) {
		v416 = v386
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v426 = v416
	goto L75
L78:
	;
	v419 = v383 + int32(1)
	if v419 != v375 {
		v383 = v419
		v386 = v416
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v401 = int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	if v402 == v401 {
		v449 = v401
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v405 = v348 + v402
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	if v406 != 0 {
		v416 = v386
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v407 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v407)
	*(*int32)(unsafe.Add(mBase, uint32(v351+v386<<(uint(int32(2))%32)))) = v402
	v416 = v386 + v407
	goto L78
L82:
	;
	goto L77
L83:
	;
	goto L72
L84:
	;
	if v449 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v454 = int32(2)
	goto L87
L86:
	;
	v454 = int32(0)
	goto L87
L87:
	;
	v490 = v454
	goto L1
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg_internal(m, int32(_a_F_gin_trgm_triconsistent_0), v14)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_gin_trgm_triconsistent_1), int32(352), int32(_a_F_gin_trgm_triconsistent_2))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v488 = int32(2)
	goto L93
L92:
	;
	v488 = int32(0)
	goto L93
L93:
	;
	v490 = v488
	goto L1
}
func F_gin_tsquery_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= v2 {
		v39 = v2
		m.G0 = v9 + int32(16)
		return v39
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v12
		v22 = v13 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v24
		v28 = F_TS_execute_ternary(m, v22, v9+int32(4))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			switch v28 - int32(1) {
			case 0:
				v39 = int32(1)
			case 1:
				v35 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v35)
				v39 = v35
			default:
				v39 = v2
			}
			m.G0 = v9 + int32(16)
			return v39
		}
	}
}
func F_gin_tsquery_consistent_6args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(7) < v11 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v20 <= v18 {
			v42 = v2
			m.G0 = v9 + int32(16)
			return v42
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
			v25 = v16 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27
			v31 = F_TS_execute_ternary(m, v25, v9+int32(4))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				switch v31 - int32(1) {
				case 0:
					v42 = int32(1)
				case 1:
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v38)
					v42 = v38
				default:
					v42 = v2
				}
				m.G0 = v9 + int32(16)
				return v42
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_tsquery_consistent_6args_0), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_tsquery_consistent_6args_1), int32(331), int32(_a_F_gin_tsquery_consistent_6args_2))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
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
func F_gin_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_startup[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(_a_F_gin_xlog_startup_0), int32(0), int32(_a_F_gin_xlog_startup_1), int32(_a_F_gin_xlog_startup_2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_startup[1])) = v8
		return
	}
}
