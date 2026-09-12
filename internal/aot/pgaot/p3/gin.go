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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v68 int32
	_ = v68
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
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
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
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
	v33 = (v29 + int32(17)) & int32(131070)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = F_ginPostingListDecodeAllSegments(m, l1+v33, v35-v33, v12+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L9
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
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v28 = v26
	goto L1
L5:
	;
	v24 = F__emscripten_memcpy_bulkmem(m, v12+int32(8), v17, v23)
	mBase = m.M
	goto L7
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
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
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v68 = v41
	goto L12
L12:
	;
	if v68 <= int32(0) {
		v112 = v68
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v61 = F_datumCopy(m, v28, v52&int32(1), v50)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	v64 = v54
	v65 = v54
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	v68 = v65
	goto L12
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = v61
	v65 = v63
	goto L15
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v112 <= v113 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v74 = int32(6)
	v78 = v73 + v68*v74 - v74
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
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v108 != 0 {
		v112 = v109
		goto L17
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	if base.Ui32(v97) < base.Ui32(v92) {
		v108 = int32(1)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+4)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	if base.Ui32(v102) < base.Ui32(v103) {
		v108 = int32(-1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v108 = base.B2i32(base.Ui32(v103) < base.Ui32(v102))
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v109
	v112 = v109
	goto L17
L25:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v183 = (v179 + v180) * int32(6)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v184 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v119 = v113
	goto L27
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v129 = v126 + v119*int32(6)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v135 = (l1 + int32(16) + v130 + int32(1)) & int32(-2)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129))))
	v141 = int32(16)
	v143 = v139 | v140<<(uint(v141)%32)
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+2)))
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135))))
	v148 = v144 | v145<<(uint(v141)%32)
	if base.Ui32(v143) < base.Ui32(v148) {
		v159 = int32(-1)
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L25
L29:
	;
	if int32(0) < v159 {
		goto L25
	} else {
		goto L34
	}
L30:
	;
	goto L29
L31:
	;
	if base.Ui32(v148) < base.Ui32(v143) {
		v159 = int32(1)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+4)))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	if base.Ui32(v153) < base.Ui32(v154) {
		v159 = int32(-1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v159 = base.B2i32(base.Ui32(v154) < base.Ui32(v153))
	goto L30
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v163 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v162 + v163
	v167 = v119 + v163
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v167 < v168 {
		v119 = v167
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v202 = F_ginMergeItemPointers(m, v191+v193*int32(6), v197-v193, v39, v199, v12+int32(4))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L42
	}
L37:
	;
	v187 = F_palloc(m, v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v189 = F_repalloc(m, v184, v183)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L41
	}
L40:
	;
	v191 = v187
	goto L36
L41:
	;
	v191 = v189
	goto L36
L42:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v206 = int32(6)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v211 = v209 * v206
	if v211 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_pfree(m, v202)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L47
	}
L44:
	;
	v212 = F__emscripten_memcpy_bulkmem(m, v204+v205*v206, v202, v211)
	mBase = m.M
	goto L46
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v216 + v217
	F_pfree(m, v39)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
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
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	v1 = l0
	v2 = l1
	v4 = l3
	v5 = l4
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v22 = base.B2i32(v2 == v9) << (uint(int32(2)) % 32)
	if v2 != 0 {
		v74 = v22
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v76 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v76
	v81 = int32(0)
	if l6 == v81 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	if v5 != 0 {
		v74 = v22
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v4 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v74 = v4
	goto L1
L5:
	;
	goto L6
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = int32(base.Ui32(v70) >> (uint(int32(2)) % 32))
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v52 = F_strlen(m, l2)
	mBase = m.M
	v74 = v52 + int32(1)
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
	v30 = int32(6)
	v32 = int32(18)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v34 == v32 {
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
		goto L23
	}
L14:
	;
	v37 = v32
	goto L16
L15:
	;
	v37 = int32(2)
	goto L16
L16:
	;
	if v34&int32(254) == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = v30
	goto L19
L18:
	;
	v42 = v37
	goto L19
L19:
	;
	if v34 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v45 = v30
	goto L22
L21:
	;
	v45 = v42
	goto L22
L22:
	;
	v74 = v45
	goto L1
L23:
	;
	v74 = int32(base.Ui32(v27) >> (uint(int32(1)) % 32))
	goto L1
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v4
	F_errmsg_internal(m, int32(710892), v17)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(516648), int32(2270), int32(402371))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v153 = (v74 + int32(17)) & int32(-2)
	v154 = v149 + v153
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v154
	v156 = F_palloc0(m, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L24
	} else {
		goto L41
	}
L29:
	;
	v149 = v9
	goto L28
L30:
	;
	goto L31
L31:
	;
	v92 = v81
	v97 = v9
	goto L32
L32:
	;
	v99 = F_palloc(m, int32(12))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v149 = v120
	goto L28
L34:
	;
	v108 = F_ginCompressPostingList(m, l5+v92*int32(6), l6-v92, int32(65535), v17+int32(4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v112 = v111 + v92
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+6)))
	v120 = v97 + (v113+int32(1))&int32(131070) + int32(8)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v121 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v99
	if base.Ui32(v112) < base.Ui32(l6) {
		v92 = v112
		v97 = v120
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v128 = v122
	goto L36
L38:
	;
	goto L39
L39:
	;
	v124 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v124
	v128 = v124
	goto L36
L40:
	;
	goto L33
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+11)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v1)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v156)+12)) = l6
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)) = uint16(v74)
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+10)) = uint8(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)) = uint16(v4)
	if v2 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v182 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	if v5 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = l2
	goto L42
L45:
	;
	goto L46
L46:
	;
	if int32(0) < v4 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v4 != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	switch v4 + int32(2) {
	case 0:
		goto L54
	case 1:
		goto L55
	default:
		goto L42
	}
L50:
	;
	goto L42
L51:
	;
	v170 = F__emscripten_memcpy_bulkmem(m, v156+int32(16), l2, v4)
	mBase = m.M
	goto L53
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	if v74 != 0 {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	if v74 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L42
L57:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, v156+int32(16), l2, v74)
	mBase = m.M
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L42
L61:
	;
	v180 = F__emscripten_memcpy_bulkmem(m, v156+int32(16), l2, v74)
	mBase = m.M
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	m.G0 = v17 + int32(16)
	return v156
L65:
	;
	if v182 == v17+int32(8) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v197 = v182
	v201 = v156 + v153
	goto L67
L67:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+6)))
	v211 = (v205+int32(1))&int32(131070) + int32(8)
	if v211 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L64
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	F_pfree(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L24
	} else {
		goto L73
	}
L70:
	;
	v212 = F__emscripten_memcpy_bulkmem(m, v201, v204, v211)
	mBase = m.M
	v213 = v212
	goto L72
L71:
	;
	v213 = v201
	goto L72
L72:
	;
	goto L69
L73:
	;
	F_pfree(m, v197)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	v231 = int32(8)
	if v203 != v17+v231 {
		v197 = v203
		v201 = v213 + (v215+int32(1))&int32(131070) + v231
		goto L67
	} else {
		goto L75
	}
L75:
	;
	goto L68
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
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
				v13 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(v5^int32(-1))<<(uint(int32(2))%32))))
				v27 = v19
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v27-int32(-64))))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v40
			F_UnlockReleaseBuffer(m, v5)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
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
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(4163720)
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
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
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v63 = int32(6)
	v65 = l2 + l3*v63
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v65-int32(4)))))
	v69 = int64(32)
	v73 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v65-v63))))
	v74 = int64(48)
	v79 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v65-int32(2)))))
	v81 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v82 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui64(v81|(v82<<(uint(v69)%64)|v85<<(uint(v74)%64))) <= base.Ui64(v68<<(uint(v69)%64)|v73<<(uint(v74)%64)|v79) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v53 = l1 * int32(6)
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(6)
	v25 = l0 + l1*v23
	v28 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25-int32(4)))))
	v29 = int64(32)
	v33 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25-v23))))
	v34 = int64(48)
	v39 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25-int32(2)))))
	v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v42 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v45 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui64(v41|(v42<<(uint(v29)%64)|v45<<(uint(v34)%64))) <= base.Ui64(v28<<(uint(v29)%64)|v33<<(uint(v34)%64)|v39) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v58 = l3 * int32(6)
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v15, l0, v53)
	mBase = m.M
	v55 = v54
	goto L11
L10:
	;
	v55 = v15
	goto L11
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v55
L13:
	;
	v59 = F__emscripten_memcpy_bulkmem(m, v55+v53, l2, v58)
	mBase = m.M
	goto L15
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v96 = l2
	v97 = v15
	v98 = l0
	goto L19
L17:
	;
	goto L18
L18:
	;
	v241 = l3 * int32(6)
	if v241 != 0 {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	v104 = base.I32_div_s(v96-l2, int32(6))
	if base.Ui32(v104) < base.Ui32(l3) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v161 = base.I32_div_s(v156-l0, int32(6))
	if base.Ui32(v161) < base.Ui32(l1) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v106 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)))
	v107 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v98)+2)))
	v108 = int64(32)
	v110 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v98))))
	v111 = int64(48)
	v114 = v106 | (v107<<(uint(v108)%64) | v110<<(uint(v111)%64))
	v115 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)))
	v116 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v96)+2)))
	v119 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
	v123 = v115 | (v116<<(uint(v108)%64) | v119<<(uint(v111)%64))
	if base.Ui64(v123) < base.Ui64(v114) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v154 = v96
	v155 = v97
	v156 = v98
	goto L23
L23:
	;
	goto L20
L24:
	;
	v148 = int32(6)
	v149 = v97 + v148
	v152 = base.I32_div_s(v147-l0, v148)
	if base.Ui32(v152) < base.Ui32(l1) {
		v96 = v146
		v97 = v149
		v98 = v147
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v125
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v127)
	v146 = v96 + int32(6)
	v147 = v98
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v114 == v123 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v132
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v134)
	v136 = int32(6)
	v146 = v96 + v136
	v147 = v98 + v136
	goto L24
L29:
	;
	goto L30
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v140
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v142)
	v146 = v96
	v147 = v98 + int32(6)
	goto L24
L31:
	;
	v154 = v146
	v155 = v149
	v156 = v147
	goto L23
L32:
	;
	v169 = v155
	v170 = v156
	goto L35
L33:
	;
	v192 = v155
	goto L34
L34:
	;
	v199 = base.I32_div_s(v154-l2, int32(6))
	if base.Ui32(v199) < base.Ui32(l3) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v174
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)) = uint16(v176)
	v178 = int32(6)
	v179 = v169 + v178
	v181 = v170 + v178
	v184 = base.I32_div_s(v181-l0, v178)
	if base.Ui32(v184) < base.Ui32(l1) {
		v169 = v179
		v170 = v181
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v192 = v179
	goto L34
L37:
	;
	goto L36
L38:
	;
	v206 = v154
	v207 = v192
	goto L41
L39:
	;
	v230 = v192
	goto L40
L40:
	;
	v237 = base.I32_div_s(v230-v15, int32(6))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v237
	return v15
L41:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v212
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v207)+4)) = uint16(v214)
	v216 = int32(6)
	v217 = v207 + v216
	v219 = v206 + v216
	v222 = base.I32_div_s(v219-l2, v216)
	if base.Ui32(v222) < base.Ui32(l3) {
		v206 = v219
		v207 = v217
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v230 = v217
	goto L40
L43:
	;
	goto L42
L44:
	;
	v246 = l1 * int32(6)
	if v246 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v242 = F__emscripten_memcpy_bulkmem(m, v15, l2, v241)
	mBase = m.M
	v243 = v242
	goto L47
L46:
	;
	v243 = v15
	goto L47
L47:
	;
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v243
L49:
	;
	v247 = F__emscripten_memcpy_bulkmem(m, v243+v241, l0, v246)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
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
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
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
	v69 = l0 + int32(32)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v70 == int32(0) {
		v301 = v69
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v34 = l0 + int32(32)
	v38 = F_ginCompressPostingList(m, v34, v32, int32(8192), v23+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v53 = v25
	v55 = int32(32)
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v55)
	v57 = l0 + v53
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)))
	v60 = v58 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v60)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v64 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v62)+4)) = uint16(v64)
	goto L3
L7:
	;
	return
L8:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	v44 = (v40 + int32(1)) & int32(131070)
	v46 = v44 + int32(8)
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v53 = v49
	v55 = v44 + int32(40)
	goto L6
L10:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v34, v38, v46)
	mBase = m.M
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v320 = v301 - v69 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v320)
	m.G0 = v23 + int32(16)
	return
L14:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v81 = v69
	v82 = v69
	v85 = v3
	v88 = v69 + v75 - int32(32)
	v91 = v3
	v94 = l1 + int32(2)
	v96 = v3
	goto L15
L15:
	;
	v99 = int32(2)
	v100 = v94 + v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v101&int32(254) == v99 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	if v284 == v255 {
		v301 = v285
		goto L13
	} else {
		goto L70
	}
L17:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v85 < v147 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v132 = v94 + int32(4)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+2)))
	v137 = int32(0)
	v140 = v137
	v141 = v137
	v143 = v132 + v133*int32(6)
	v144 = int32(1)
	v145 = v132
	v146 = v133
	goto L17
L19:
	;
	v129 = int32(0)
	v140 = v124
	v141 = v125
	v143 = v128
	v144 = v127
	v145 = v129
	v146 = v129
	goto L17
L20:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+8)))
	v110 = (v106 + int32(1)) & int32(131070)
	v124 = v110 + int32(8)
	v125 = v100
	v127 = int32(0)
	v128 = v100 + (v110+int32(9))&int32(262142)
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v101 == int32(4) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v121 = int32(0)
	v124 = v121
	v125 = v121
	v127 = v121
	v128 = v100
	goto L19
L24:
	;
	v151 = v81
	v152 = v82
	v155 = v85
	goto L27
L25:
	;
	v193 = v81
	v194 = v82
	v197 = v85
	goto L26
L26:
	;
	if v144 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+6)))
	v173 = (v169 + int32(1)) & int32(131070)
	v175 = v173 + int32(8)
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v193 = v187
	v194 = v184
	v197 = v147
	goto L26
L29:
	;
	if v175 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v183 = v173
	goto L31
L31:
	;
	v184 = v152 + v175
	v187 = v151 + v183 + int32(8)
	v189 = v155 + int32(1)
	if v189 != v147 {
		v151 = v187
		v152 = v184
		v155 = v189
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+6)))
	v183 = (v178 + int32(1)) & int32(131070)
	goto L31
L33:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, v152, v151, v175)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L28
L37:
	;
	v213 = F_ginPostingListDecode(m, v193, v23+int32(12))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	v234 = v140
	v235 = v101
	v236 = v141
	goto L39
L39:
	;
	if v193 == v88 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v218 = F_ginMergeItemPointers(m, v145, v146, v213, v215, v23+int32(8))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v224 = F_ginCompressPostingList(m, v218, v220, int32(8192), v23+int32(4))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+6)))
	v234 = (v226+int32(1))&int32(131070) + int32(8)
	v235 = int32(3)
	v236 = v224
	goto L39
L43:
	;
	switch v235 - int32(1) {
	case 0:
		goto L54
	case 1:
		goto L57
	case 2:
		goto L56
	default:
		goto L55
	}
L44:
	;
	v252 = v193
	v253 = int32(0)
	v255 = v88
	v256 = v91
	goto L43
L45:
	;
	goto L46
L46:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+6)))
	v245 = (v239+int32(1))&int32(131070) + int32(8)
	if v91 != 0 {
		v252 = v193
		v253 = v245
		v255 = v88
		v256 = v91
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v246 = v88 - v193
	v247 = F_palloc(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if v246 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v252 = v250
	v253 = v245
	v255 = v250 + v246
	v256 = v247
	goto L43
L50:
	;
	v249 = F__emscripten_memcpy_bulkmem(m, v247, v193, v246)
	mBase = m.M
	v250 = v249
	goto L52
L51:
	;
	v250 = v247
	goto L52
L52:
	;
	goto L49
L53:
	;
	v288 = v96 + int32(1)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if base.Ui32(v288) < base.Ui32(v289) {
		v81 = v284
		v82 = v285
		v85 = v286
		v88 = v255
		v91 = v256
		v94 = v143
		v96 = v288
		goto L15
	} else {
		goto L69
	}
L54:
	;
	v284 = v252 + v253
	v285 = v194
	v286 = v197 + int32(1)
	goto L53
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L66
	}
L56:
	;
	if v234 != 0 {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	if v234 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v284 = v252
	v285 = v260 + v234
	v286 = v197
	goto L53
L59:
	;
	v259 = F__emscripten_memcpy_bulkmem(m, v194, v236, v234)
	mBase = m.M
	v260 = v259
	goto L61
L60:
	;
	v260 = v194
	goto L61
L61:
	;
	goto L58
L62:
	;
	v284 = v252 + v253
	v285 = v266 + v234
	v286 = v197 + int32(1)
	goto L53
L63:
	;
	v265 = F__emscripten_memcpy_bulkmem(m, v194, v236, v234)
	mBase = m.M
	v266 = v265
	goto L65
L64:
	;
	v266 = v194
	goto L65
L65:
	;
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v235
	F_errmsg_internal(m, int32(63389), v23)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(523011), int32(298), int32(134863))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	goto L16
L70:
	;
	if v256 == int32(0) {
		v301 = v285
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v294 = v255 - v284
	if v294 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v301 = v296 + v294
	goto L13
L73:
	;
	v295 = F__emscripten_memcpy_bulkmem(m, v285, v284, v294)
	mBase = m.M
	v296 = v295
	goto L75
L74:
	;
	v296 = v285
	goto L75
L75:
	;
	goto L72
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
			v19 = F_CallerFInfoFunctionCall2(m, int32(3814), v17, v18, v5, v4)
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
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
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v196
L3:
	;
	if v19&int32(65534) != int32(10) {
		goto L1
	} else {
		goto L30
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
	v25 = F_DirectFunctionCall2Coll(m, int32(5456), int32(0), v24, v18)
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
		v196 = v25
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v196 = int32(0)
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
	v72 = v70 + int32(5)
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L12:
	;
	v47 = int32(4)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32((v48-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v60 = int32(1)
	if v44 != 0 {
		v70 = int32(base.Ui32(v42)>>(uint(v60)%32)) - v60
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v59 = v47
	goto L17
L16:
	;
	v59 = base.B2i32(v48 == int32(18)) << (uint(v47) % 32)
	goto L17
L17:
	;
	v70 = v59
	goto L11
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L19:
	;
	v75 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v72 << (uint(int32(2)) % 32)
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v44 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v73
	v196 = v40
	goto L2
L23:
	;
	v86 = v38
	goto L25
L24:
	;
	v86 = v33 + int32(4)
	goto L25
L25:
	;
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	v87 = F__emscripten_memcpy_bulkmem(m, v73+int32(5), v86, v70)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = F_pg_detoast_datum(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_deconstruct_array_builtin(m, v95, int32(25), v15+int32(12), v15+int32(8), v15+int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v109 = F_palloc(m, v106<<(uint(int32(2))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) < v111 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = int32(0)
	v116 = v111
	v118 = v2
	goto L37
L35:
	;
	v178 = v2
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v178
	if v19 != int32(11) {
		v196 = v109
		goto L2
	} else {
		goto L51
	}
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v115))))
	if v129 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v178 = v167
	goto L36
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v133 = int32(2)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v115<<(uint(v133)%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v139 = int32(base.Ui32(v137) >> (uint(v133) % 32))
	v141 = v139 + int32(1)
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	v166 = v116
	v167 = v118
	goto L41
L41:
	;
	v173 = v115 + int32(1)
	if v173 < v166 {
		v115 = v173
		v116 = v166
		v118 = v167
		goto L37
	} else {
		goto L50
	}
L42:
	;
	v144 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v142)+4)) = uint8(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v141 << (uint(int32(2)) % 32)
	if base.Ui32(int32(20)) <= base.Ui32(v137) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v153 = int32(4)
	v156 = v139 - v153
	if v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109+v118<<(uint(int32(2))%32)))) = v142
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v166 = v165
	v167 = v118 + int32(1)
	goto L41
L46:
	;
	goto L45
L47:
	;
	v157 = F__emscripten_memcpy_bulkmem(m, v142+int32(5), v136+v153, v156)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L38
L51:
	;
	if v178 != 0 {
		v196 = v109
		goto L2
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v196 = v109
	goto L2
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	F_errmsg_internal(m, int32(505404), v15)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(521107), int32(141), int32(16124))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_query_anyenum(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(6193)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(505404), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521120), int32(97), int32(15901))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
		}
	}
}
func F_gin_extract_query_bytea(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(7630), int32(2866))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_int4(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2116)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(-2147483648)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(505404), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521120), int32(97), int32(15901))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
		}
	}
}
func F_gin_extract_query_time(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(1448)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_Int64GetDatum(m, int64(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(505404), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521120), int32(97), int32(15901))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
		}
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
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	switch v2 - int32(3) {
	case 0:
		v27 = F_gin_extract_value_trgm(m, l0)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(303221), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(521096), int32(30), int32(303255))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
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
			v25 = int32(4)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v27&int32(254) == int32(2) {
				v36 = v25
			} else {
				v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
			}
			if v27 == int32(1) {
				v39 = v25
			} else {
				v39 = v36
			}
			v50 = v39
		} else {
			v40 = int32(1)
			if v21 != 0 {
				v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v52 = F_generate_trgm(m, v22, v50)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
			v58 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(5)
			if base.Ui32(int32(3)) <= base.Ui32(v58) {
				v62 = base.I32_div_u_s(v58, int32(3))
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v62
				v64 = int32(1)
				if base.Ui32(v62) <= base.Ui32(v64) {
					v67 = v64
				} else {
					v67 = v62
				}
				v73 = F_palloc(m, v62<<(uint(int32(2))%32))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = int32(0)
					v76 = v52 + int32(5)
					for {
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
						*(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(int32(2))%32)))) = v84 | (v85<<(uint(int32(8))%32) | v88<<(uint(int32(16))%32))
						v97 = v75 + int32(1)
						if v97 != v67 {
							v75 = v97
							v76 = v76 + int32(3)
							continue
						} else {
							break
						}
						break
					}
					v101 = v73
					return v101
				}
			} else {
				v101 = int32(0)
				return v101
			}
		}
	}
}
func F_gin_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = l0 - int32(16)
	if base.Ui32(v4) <= base.Ui32(int32(143)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v4)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_consts[77])))
		v14 = v13
	} else {
		v14 = int32(0)
	}
	return v14
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
	var v22 int32
	_ = v22
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
	v12 = v10 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v7)+6)))
	if v15&int32(4) != 0 {
		v20 = int32(0)
		v22 = F___memset(m, l0+int32(24), v20, int32(8168))
		mBase = m.M
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int64
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int64
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int64
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int64
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int64
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v944 int64
	_ = v944
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int64
	_ = v962
	var v966 int64
	_ = v966
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v972 int64
	_ = v972
	var v974 int64
	_ = v974
	var v976 int64
	_ = v976
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int64
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1384 int64
	_ = v1384
	var v1398 int32
	_ = v1398
	var v1400 int64
	_ = v1400
	var v1404 int64
	_ = v1404
	var v1406 int64
	_ = v1406
	var v1408 int64
	_ = v1408
	var v1410 int64
	_ = v1410
	var v1412 int64
	_ = v1412
	var v1414 int64
	_ = v1414
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1432 int32
	_ = v1432
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1513 int32
	_ = v1513
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = int32(4562080)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
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
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L19
	} else {
		goto L387
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L19
	} else {
		goto L384
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L19
	} else {
		goto L381
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L19
	} else {
		goto L378
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L19
	} else {
		goto L375
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L19
	} else {
		goto L372
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L19
	} else {
		goto L369
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
	v1529 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	F_MemoryContextReset(m, v1529)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L19
	} else {
		goto L368
	}
L9:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1334 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L19
	} else {
		goto L339
	}
L10:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v1179 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1181 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L19
	} else {
		goto L305
	}
L11:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v892 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v894 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L19
	} else {
		goto L249
	}
L12:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v735 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v739 = F_XLogReadBufferForRedo(m, l0, int32(2), v13+int32(-8))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L19
	} else {
		goto L207
	}
L13:
	;
	v662 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v666 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L19
	} else {
		goto L184
	}
L14:
	;
	v655 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L19
	} else {
		goto L181
	}
L15:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v566)+24)))
	if v567&int32(2) != 0 {
		goto L156
	} else {
		goto L157
	}
L16:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104))))
	v107 = v105 & int32(2)
	if v107 == int32(0) {
		goto L35
	} else {
		goto L36
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v34^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L18
L22:
	;
	goto L23
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v53 = v47 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v87 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	F_PageInit(m, v72, int32(8192), int32(8))
	mBase = m.M
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+16)))
	v77 = v72 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v54)
	goto L24
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(v34^int32(-1))<<(uint(int32(2))%32))))
	v72 = v64
	goto L25
L27:
	;
	goto L28
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v72 = v66 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
	*(*uint32)(unsafe.Add(mBase, uint32(v53)+4)) = uint32(v32)
	v93 = int64(base.Ui64(v32) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v53))) = uint32(v93)
	v96 = v90 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)) = uint16(v96)
	F_MarkBufferDirty(m, v34)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L33
	}
L30:
	;
	v88 = F__emscripten_memcpy_bulkmem(m, v53+int32(32), v84+int32(4), v87)
	mBase = m.M
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	F_UnlockReleaseBuffer(m, v34)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	goto L8
L35:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+6)))
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+8)))
	v115 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L38
	}
L36:
	;
	v159 = int32(-1)
	goto L37
L37:
	;
	v166 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-24))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L51
	}
L38:
	;
	if v115 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v119 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v155 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+16)))
	v139 = v138 + v137
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+6)))
	v142 = v140 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v139)+6)) = uint16(v142)
	*(*uint32)(unsafe.Add(mBase, uint32(v137)+4)) = uint32(v102)
	v146 = int64(base.Ui64(v102) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v137))) = uint32(v146)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L19
	} else {
		goto L46
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+(v119^int32(-1))<<(uint(int32(2))%32))))
	v137 = v129
	goto L42
L44:
	;
	goto L45
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v137 = v131 + v119<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L46:
	;
	goto L41
L47:
	;
	F_UnlockReleaseBuffer(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L19
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = v110<<(uint(int32(16))%32) | v111
	goto L37
L50:
	;
	goto L49
L51:
	;
	if v166 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v170 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v561 == int32(0) {
		goto L8
	} else {
		goto L154
	}
L55:
	;
	v189 = int32(0)
	v191 = v13 + int32(-28)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+72))
	if v194 < v189 {
		v216 = v189
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174+(v170^int32(-1))<<(uint(int32(2))%32))))
	v188 = v180
	goto L55
L57:
	;
	goto L58
L58:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v188 = v182 + v170<<(uint(int32(13))%32) + int32(-8192)
	goto L55
L59:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v221&int32(1) != 0 {
		goto L72
	} else {
		goto L73
	}
L60:
	;
	v219 = v216
	goto L59
L61:
	;
	v200 = v193 + int32(76)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v201 != int32(1) {
		v216 = v189
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+43)))
	if v204 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v191 == int32(0) {
		v216 = v189
		goto L60
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v191 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v209
	v219 = v209
	goto L59
L67:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v212
	goto L69
L68:
	;
	goto L69
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v200)+44))
	v216 = v214
	goto L60
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v188))) = base.I64_rotr(v102, int64(32))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_MarkBufferDirty(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L19
	} else {
		goto L153
	}
L71:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v344*int32(10))+22)) = base.I32_rotr(v159, int32(16))
	v352 = v219 + int32(2)
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+16)))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v353)+4)))
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	if v356 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L72:
	;
	if v220 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	if v220 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	if v107 == int32(0) {
		goto L71
	} else {
		goto L79
	}
L76:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227+(v220^int32(-1))<<(uint(int32(2))%32))))
	v241 = v233
	goto L75
L77:
	;
	goto L78
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v241 = v235 + v220<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	F_ginRedoRecompress(m, v241, v219)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	goto L70
L81:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	if v159 != int32(-1) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v249+(v220^int32(-1))<<(uint(int32(2))%32))))
	v263 = v255
	goto L81
L83:
	;
	goto L84
L84:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v263 = v257 + v220<<(uint(int32(13))%32) + int32(-8192)
	goto L81
L85:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264<<(uint(int32(2))%32)+v263)+20))
	v273 = v263 + v270&int32(32767)
	v274 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v273)+4)) = uint16(v274)
	*(*uint16)(unsafe.Add(mBase, uint32(v273)+2)) = uint16(v159)
	v278 = int32(base.Ui32(v159) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v273))) = uint16(v278)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+2)))
	if v281 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_PageIndexTupleDelete(m, v263, v264)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L19
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+10)))
	v292 = F_PageAddItemExtended(m, v263, v219+int32(4), v288&int32(8191), v264, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L19
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	if v292 != 0 {
		goto L70
	} else {
		goto L93
	}
L93:
	;
	v295 = v13 + int32(-20)
	if v220 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L19
	} else {
		goto L99
	}
L95:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v317)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v319
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v318
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-4)))) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v317)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-8)))) = v324
	goto L94
L96:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v317 = v304 + (v220^int32(-1))<<(uint(int32(6))%32)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v317 = v311 + v220<<(uint(int32(6))%32) + int32(-64)
	goto L95
L99:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+20)) = v332
	F_errmsg_internal(m, int32(41556), v13+int32(-48))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L19
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(523011), int32(104), int32(12673))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L19
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v352)))
	*(*int64)(unsafe.Add(mBase, uint32(v524))) = v526
	v528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v352)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v524)+8)) = uint16(v528)
	v532 = v355 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v241+v525)+4)) = uint16(v532)
	v537 = v355*int32(10) + int32(42)
	*(*uint16)(unsafe.Add(mBase, uint32(v241)+12)) = uint16(v537)
	goto L70
L103:
	;
	v524 = v241 + v355*int32(10) + int32(32)
	v525 = v353
	goto L102
L104:
	;
	goto L105
L105:
	;
	v366 = v241 + v356*int32(10)
	v368 = v366 + int32(22)
	if v355+int32(1) == v356 {
		v524 = v368
		v525 = v353
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v372 = int32(10)
	v373 = v366 + int32(32)
	v378 = (v355-v356)*v372 + v372
	if v373 == v368 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+16)))
	v524 = v368
	v525 = v523
	goto L102
L108:
	;
	goto L107
L109:
	;
	v382 = v373 + v378
	if base.Ui32(v368-v382) <= base.Ui32(int32(0)-v378<<(uint(int32(1))%32)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v389 = F___memcpy(m, v373, v368, v378)
	mBase = m.M
	goto L107
L111:
	;
	goto L112
L112:
	;
	v392 = (v373 ^ v368) & int32(3)
	if base.Ui32(v373) < base.Ui32(v368) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if v494 == int32(0) {
		goto L108
	} else {
		goto L149
	}
L114:
	;
	if base.Ui32(v472) <= base.Ui32(int32(3)) {
		v493 = v471
		v494 = v472
		v495 = v473
		goto L113
	} else {
		goto L145
	}
L115:
	;
	if v392 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	if v392 != 0 {
		v454 = v378
		goto L128
	} else {
		goto L129
	}
L118:
	;
	v493 = v368
	v494 = v378
	v495 = v373
	goto L113
L119:
	;
	goto L120
L120:
	;
	if v373&int32(3) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v471 = v368
	v472 = v378
	v473 = v373
	goto L114
L122:
	;
	goto L123
L123:
	;
	v399 = v368
	v400 = v378
	v401 = v373
	goto L124
L124:
	;
	if v400 == int32(0) {
		goto L108
	} else {
		goto L126
	}
L125:
	;
	v471 = v408
	v472 = v410
	v473 = v412
	goto L114
L126:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v405)
	v407 = int32(1)
	v408 = v399 + v407
	v410 = v400 - v407
	v412 = v401 + v407
	if v412&int32(3) != 0 {
		v399 = v408
		v400 = v410
		v401 = v412
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	if v454 == int32(0) {
		goto L108
	} else {
		goto L141
	}
L129:
	;
	if v382&int32(3) != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v419 = v378
	goto L133
L131:
	;
	v434 = v378
	goto L132
L132:
	;
	if base.Ui32(v434) <= base.Ui32(int32(3)) {
		v454 = v434
		goto L128
	} else {
		goto L137
	}
L133:
	;
	if v419 == int32(0) {
		goto L108
	} else {
		goto L135
	}
L134:
	;
	v434 = v425
	goto L132
L135:
	;
	v425 = v419 - int32(1)
	v426 = v373 + v425
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v425))))
	*(*uint8)(unsafe.Add(mBase, uint32(v426))) = uint8(v428)
	if v426&int32(3) != 0 {
		v419 = v425
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v441 = v434
	goto L138
L138:
	;
	v445 = v441 - int32(4)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v368+v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v373+v445))) = v448
	if base.Ui32(int32(3)) < base.Ui32(v445) {
		v441 = v445
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v454 = v445
	goto L128
L140:
	;
	goto L139
L141:
	;
	v461 = v454
	goto L142
L142:
	;
	v465 = v461 - int32(1)
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v465))))
	*(*uint8)(unsafe.Add(mBase, uint32(v373+v465))) = uint8(v468)
	if v465 != 0 {
		v461 = v465
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L108
L144:
	;
	goto L143
L145:
	;
	v478 = v471
	v479 = v472
	v480 = v473
	goto L146
L146:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = v482
	v484 = int32(4)
	v485 = v478 + v484
	v487 = v480 + v484
	v489 = v479 - v484
	if base.Ui32(int32(3)) < base.Ui32(v489) {
		v478 = v485
		v479 = v489
		v480 = v487
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v493 = v485
	v494 = v489
	v495 = v487
	goto L113
L148:
	;
	goto L147
L149:
	;
	v500 = v493
	v501 = v494
	v502 = v495
	goto L150
L150:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	*(*uint8)(unsafe.Add(mBase, uint32(v502))) = uint8(v504)
	v506 = int32(1)
	v511 = v501 - v506
	if v511 != 0 {
		v500 = v500 + v506
		v501 = v511
		v502 = v502 + v506
		goto L150
	} else {
		goto L152
	}
L151:
	;
	goto L108
L152:
	;
	goto L151
L153:
	;
	goto L54
L154:
	;
	F_UnlockReleaseBuffer(m, v561)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L19
	} else {
		goto L155
	}
L155:
	;
	goto L8
L156:
	;
	v623 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L19
	} else {
		goto L169
	}
L157:
	;
	v570 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v574 = F_XLogReadBufferForRedo(m, l0, int32(3), v13+int32(-20))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	if v574 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v578 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	goto L161
L161:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v612 == int32(0) {
		goto L156
	} else {
		goto L167
	}
L162:
	;
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+16)))
	v598 = v597 + v596
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+6)))
	v601 = v599 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v598)+6)) = uint16(v601)
	*(*uint32)(unsafe.Add(mBase, uint32(v596)+4)) = uint32(v570)
	v605 = int64(base.Ui64(v570) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v596))) = uint32(v605)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L19
	} else {
		goto L166
	}
L163:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582+(v578^int32(-1))<<(uint(int32(2))%32))))
	v596 = v588
	goto L162
L164:
	;
	goto L165
L165:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v596 = v590 + v578<<(uint(int32(13))%32) + int32(-8192)
	goto L162
L166:
	;
	goto L161
L167:
	;
	F_UnlockReleaseBuffer(m, v612)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	goto L156
L169:
	;
	if v623 != int32(2) {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	v630 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-4))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L19
	} else {
		goto L171
	}
L171:
	;
	if v630 != int32(2) {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	if v567&int32(4) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v639 = F_XLogReadBufferForRedo(m, l0, int32(2), v13+int32(-8))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L19
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	F_UnlockReleaseBuffer(m, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L19
	} else {
		goto L179
	}
L176:
	;
	if v639 != int32(2) {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	F_UnlockReleaseBuffer(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L19
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_UnlockReleaseBuffer(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L19
	} else {
		goto L180
	}
L180:
	;
	goto L8
L181:
	;
	if v655 != int32(2) {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_UnlockReleaseBuffer(m, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L19
	} else {
		goto L183
	}
L183:
	;
	goto L8
L184:
	;
	if v666 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v670 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v729 == int32(0) {
		goto L8
	} else {
		goto L205
	}
L188:
	;
	v689 = int32(0)
	v691 = v13 + int32(-4)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+72))
	if v694 < v689 {
		v716 = v689
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v674+(v670^int32(-1))<<(uint(int32(2))%32))))
	v688 = v680
	goto L188
L190:
	;
	goto L191
L191:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v688 = v682 + v670<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L192:
	;
	F_ginRedoRecompress(m, v688, v719)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L19
	} else {
		goto L203
	}
L193:
	;
	v719 = v716
	goto L192
L194:
	;
	v700 = v693 + int32(76)
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v701 != int32(1) {
		v716 = v689
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+43)))
	if v704 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	if v691 == int32(0) {
		v716 = v689
		goto L193
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	if v691 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v709
	v719 = v709
	goto L192
L200:
	;
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v700)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v712
	goto L202
L201:
	;
	goto L202
L202:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v700)+44))
	v716 = v714
	goto L193
L203:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v688))) = base.I64_rotr(v662, int64(32))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L19
	} else {
		goto L204
	}
L204:
	;
	goto L187
L205:
	;
	F_UnlockReleaseBuffer(m, v729)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	goto L8
L207:
	;
	if v739 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v743 < int32(0) {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	goto L210
L210:
	;
	v777 = F_XLogReadBufferForRedo(m, l0, int32(0), v13+int32(-20))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L19
	} else {
		goto L216
	}
L211:
	;
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+16)))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v762+v761))) = v764
	*(*uint32)(unsafe.Add(mBase, uint32(v761)+4)) = uint32(v735)
	v768 = int64(base.Ui64(v735) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v761))) = uint32(v768)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	F_MarkBufferDirty(m, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L19
	} else {
		goto L215
	}
L212:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v747+(v743^int32(-1))<<(uint(int32(2))%32))))
	v761 = v753
	goto L211
L213:
	;
	goto L214
L214:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v761 = v755 + v743<<(uint(int32(13))%32) + int32(-8192)
	goto L211
L215:
	;
	goto L210
L216:
	;
	if v777 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v781 < int32(0) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	v820 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-4))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L19
	} else {
		goto L225
	}
L220:
	;
	v800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v799)+16)))
	v801 = v800 + v799
	v802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v801)+6)))
	v804 = v802 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v801)+6)) = uint16(v804)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v734)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v799)+4)) = uint32(v735)
	v809 = int64(base.Ui64(v735) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v799))) = uint32(v809)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+20)) = v806
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v812)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L19
	} else {
		goto L224
	}
L221:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v785+(v781^int32(-1))<<(uint(int32(2))%32))))
	v799 = v791
	goto L220
L222:
	;
	goto L223
L223:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v799 = v793 + v781<<(uint(int32(13))%32) + int32(-8192)
	goto L220
L224:
	;
	goto L219
L225:
	;
	if v820 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v824 < int32(0) {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	goto L228
L228:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v880 != 0 {
		goto L238
	} else {
		goto L239
	}
L229:
	;
	v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v734))))
	v846 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v842)+16)))
	v848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v842+v846)+4)))
	if v848 != v843 {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v828+(v824^int32(-1))<<(uint(int32(2))%32))))
	v842 = v834
	goto L229
L231:
	;
	goto L232
L232:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v842 = v836 + v824<<(uint(int32(13))%32) + int32(-8192)
	goto L229
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v842))) = base.I64_rotr(v735, int64(32))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	F_MarkBufferDirty(m, v876)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L19
	} else {
		goto L237
	}
L234:
	;
	v850 = int32(10)
	v852 = v842 + v843*v850
	v860 = F_memmove(m, v852+int32(22), v852+int32(32), (v848-v843)*v850)
	mBase = m.M
	v861 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v842)+16)))
	v863 = v861
	goto L236
L235:
	;
	v863 = v846
	goto L236
L236:
	;
	v866 = v848 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v863+v842)+4)) = uint16(v866)
	v871 = v848*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v842)+12)) = uint16(v871)
	goto L233
L237:
	;
	goto L228
L238:
	;
	F_UnlockReleaseBuffer(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L19
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v883 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	F_UnlockReleaseBuffer(m, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L19
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v886 == int32(0) {
		goto L8
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	F_UnlockReleaseBuffer(m, v886)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L19
	} else {
		goto L247
	}
L247:
	;
	goto L8
L248:
	;
	if v894 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L249:
	;
	if v894 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v899+(v894^int32(-1))<<(uint(int32(2))%32))))
	v913 = v905
	goto L248
L251:
	;
	goto L252
L252:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v913 = v907 + v894<<(uint(int32(13))%32) + int32(-8192)
	goto L248
L253:
	;
	v960 = int32(-64)
	v962 = *(*int64)(unsafe.Add(mBase, uint32(v891-v960)))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+72)) = v962
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v891)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v913-v960))) = v966
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v891)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+56)) = v968
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v891)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+48)) = v970
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v891)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+40)) = v972
	v974 = *(*int64)(unsafe.Add(mBase, uint32(v891)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+32)) = v974
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v891)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+24)) = v976
	*(*int64)(unsafe.Add(mBase, uint32(v913))) = base.I64_rotr(v892, int64(32))
	F_MarkBufferDirty(m, v894)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L19
	} else {
		goto L258
	}
L254:
	;
	v934 = int32(8)
	F_PageInit(m, v932, int32(8192), v934)
	mBase = m.M
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v932)+16)))
	v937 = v932 + v936
	*(*int32)(unsafe.Add(mBase, uint32(v937))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v937)+6)) = uint16(v934)
	v944 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v932-int32(-64)))) = v944
	*(*int64)(unsafe.Add(mBase, uint32(v932)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v932)+32)) = v944
	*(*int64)(unsafe.Add(mBase, uint32(v932)+40)) = v944
	*(*int64)(unsafe.Add(mBase, uint32(v932)+48)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v932)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v932)+72)) = int32(2)
	v958 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v932)+12)) = uint16(v958)
	goto L253
L255:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v918+(v894^int32(-1))<<(uint(int32(2))%32))))
	v932 = v924
	goto L254
L256:
	;
	goto L257
L257:
	;
	v926 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v932 = v926 + v894<<(uint(int32(13))%32) + int32(-8192)
	goto L254
L258:
	;
	v985 = base.I32_wrap_i64(int64(base.Ui64(v892) >> (uint(int64(32)) % 64)))
	v986 = base.I32_wrap_i64(v892)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v891)+80))
	if int32(0) < v987 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	F_UnlockReleaseBuffer(m, v894)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L19
	} else {
		goto L303
	}
L260:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v1159 == int32(0) {
		goto L259
	} else {
		goto L301
	}
L261:
	;
	v993 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L19
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v891)+72))
	if v1111 == int32(-1) {
		goto L259
	} else {
		goto L293
	}
L264:
	;
	if v993 != 0 {
		goto L260
	} else {
		goto L265
	}
L265:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v995 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1016 = v13 + int32(-4)
	v1017 = int32(0)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+72))
	if v1019 < int32(1) {
		v1041 = v1017
		goto L271
	} else {
		goto L272
	}
L267:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v999+(v995^int32(-1))<<(uint(int32(2))%32))))
	v1013 = v1005
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1013 = v1007 + v995<<(uint(int32(13))%32) + int32(-8192)
	goto L266
L270:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v891)+80))
	if int32(0) < v1045 {
		goto L281
	} else {
		goto L282
	}
L271:
	;
	v1044 = v1041
	goto L270
L272:
	;
	v1025 = v1018 + int32(128)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025))))
	if v1026 != int32(1) {
		v1041 = v1017
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025)+43)))
	if v1029 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	if v1016 == int32(0) {
		v1041 = v1017
		goto L271
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	if v1016 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1034 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1034
	v1044 = v1034
	goto L270
L278:
	;
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1025)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1037
	goto L280
L279:
	;
	goto L280
L280:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+44))
	v1041 = v1039
	goto L271
L281:
	;
	v1048 = int32(1)
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1013)+12)))
	if base.Ui32(v1049) < base.Ui32(int32(25)) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1013)+16)))
	v1101 = v1013 + v1100
	v1102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1101)+4)))
	v1104 = v1102 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1101)+4)) = uint16(v1104)
	*(*int32)(unsafe.Add(mBase, uint32(v1013)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v1013))) = v985
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L19
	} else {
		goto L292
	}
L284:
	;
	v1058 = v1048
	goto L286
L285:
	;
	v1058 = int32(base.Ui32(v1049+int32(262120))>>(uint(int32(2))%32)) + v1048
	goto L286
L286:
	;
	v1059 = v1058
	v1060 = v1044
	v1063 = int32(0)
	goto L287
L287:
	;
	v1071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1060)+6)))
	v1073 = v1071 & int32(8191)
	v1077 = F_PageAddItemExtended(m, v1013, v1060, v1073, v1059&int32(65535), int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L19
	} else {
		goto L289
	}
L288:
	;
	goto L283
L289:
	;
	if v1077 == int32(0) {
		goto L3
	} else {
		goto L290
	}
L290:
	;
	v1081 = int32(1)
	v1085 = v1063 + v1081
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v891)+80))
	if v1085 < v1086 {
		v1059 = v1059 + v1081
		v1060 = v1060 + v1073
		v1063 = v1085
		goto L287
	} else {
		goto L291
	}
L291:
	;
	goto L288
L292:
	;
	goto L260
L293:
	;
	v1117 = F_XLogReadBufferForRedo(m, l0, int32(1), v13+int32(-20))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L19
	} else {
		goto L294
	}
L294:
	;
	if v1117 != 0 {
		goto L260
	} else {
		goto L295
	}
L295:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v1119 < int32(0) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v1138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137)+16)))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v891)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1138+v1137))) = v1140
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v1137))) = v985
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_MarkBufferDirty(m, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L19
	} else {
		goto L300
	}
L297:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1123+(v1119^int32(-1))<<(uint(int32(2))%32))))
	v1137 = v1129
	goto L296
L298:
	;
	goto L299
L299:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1137 = v1131 + v1119<<(uint(int32(13))%32) + int32(-8192)
	goto L296
L300:
	;
	goto L260
L301:
	;
	F_UnlockReleaseBuffer(m, v1159)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L19
	} else {
		goto L302
	}
L302:
	;
	goto L259
L303:
	;
	goto L8
L304:
	;
	v1201 = int32(16)
	if v1181 < int32(0) {
		goto L311
	} else {
		goto L312
	}
L305:
	;
	if v1181 < int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1186+(v1181^int32(-1))<<(uint(int32(2))%32))))
	v1200 = v1192
	goto L304
L307:
	;
	goto L308
L308:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1200 = v1194 + v1181<<(uint(int32(13))%32) + int32(-8192)
	goto L304
L309:
	;
	v1228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1200)+16)))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1178)))
	*(*int32)(unsafe.Add(mBase, uint32(v1200+v1228))) = v1230
	v1232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1200)+16)))
	if v1230 != int32(-1) {
		goto L314
	} else {
		goto L315
	}
L310:
	;
	F_PageInit(m, v1219, int32(8192), int32(8))
	mBase = m.M
	v1223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1219)+16)))
	v1224 = v1219 + v1223
	*(*int32)(unsafe.Add(mBase, uint32(v1224))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1224)+6)) = uint16(v1201)
	goto L309
L311:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1205+(v1181^int32(-1))<<(uint(int32(2))%32))))
	v1219 = v1211
	goto L310
L312:
	;
	goto L313
L313:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1219 = v1213 + v1181<<(uint(int32(13))%32) + int32(-8192)
	goto L310
L314:
	;
	v1244 = v1232
	v1245 = int32(0)
	goto L316
L315:
	;
	v1237 = v1232 + v1200
	v1238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1237)+6)))
	v1240 = v1238 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1237)+6)) = uint16(v1240)
	v1242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1200)+16)))
	v1244 = v1242
	v1245 = int32(1)
	goto L316
L316:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1244+v1200)+4)) = uint16(v1245)
	v1248 = int32(0)
	v1250 = v13 + int32(-20)
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+72))
	if v1253 < v1248 {
		v1275 = v1248
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+4))
	if int32(0) < v1280 {
		goto L328
	} else {
		goto L329
	}
L318:
	;
	v1278 = v1275
	goto L317
L319:
	;
	v1259 = v1252 + int32(76)
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259))))
	if v1260 != int32(1) {
		v1275 = v1248
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259)+43)))
	if v1263 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if v1250 == int32(0) {
		v1275 = v1248
		goto L318
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	if v1250 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1268 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1268
	v1278 = v1268
	goto L317
L325:
	;
	v1271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1259)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1271
	goto L327
L326:
	;
	goto L327
L327:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+44))
	v1275 = v1273
	goto L318
L328:
	;
	v1283 = int32(1)
	v1284 = v1278
	v1286 = int32(0)
	goto L331
L329:
	;
	goto L330
L330:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1200))) = base.I64_rotr(v1179, int64(32))
	F_MarkBufferDirty(m, v1181)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L19
	} else {
		goto L336
	}
L331:
	;
	v1295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1284)+6)))
	v1297 = v1295 & int32(8191)
	v1301 = F_PageAddItemExtended(m, v1200, v1284, v1297, v1283&int32(65535), int32(0))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L19
	} else {
		goto L333
	}
L332:
	;
	goto L330
L333:
	;
	if v1301 == int32(0) {
		goto L2
	} else {
		goto L334
	}
L334:
	;
	v1305 = int32(1)
	v1309 = v1286 + v1305
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+4))
	if v1309 < v1310 {
		v1283 = v1283 + v1305
		v1284 = v1284 + v1297
		v1286 = v1309
		goto L331
	} else {
		goto L335
	}
L335:
	;
	goto L332
L336:
	;
	F_UnlockReleaseBuffer(m, v1181)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L19
	} else {
		goto L337
	}
L337:
	;
	goto L8
L338:
	;
	if v1334 < int32(0) {
		goto L345
	} else {
		goto L346
	}
L339:
	;
	if v1334 < int32(0) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1339+(v1334^int32(-1))<<(uint(int32(2))%32))))
	v1353 = v1345
	goto L338
L341:
	;
	goto L342
L342:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1353 = v1347 + v1334<<(uint(int32(13))%32) + int32(-8192)
	goto L338
L343:
	;
	v1400 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+72)) = v1400
	v1404 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1353-int32(-64)))) = v1404
	v1406 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+56)) = v1406
	v1408 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+48)) = v1408
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+40)) = v1410
	v1412 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+32)) = v1412
	v1414 = *(*int64)(unsafe.Add(mBase, uint32(v1331)))
	*(*int64)(unsafe.Add(mBase, uint32(v1353)+24)) = v1414
	*(*int64)(unsafe.Add(mBase, uint32(v1353))) = base.I64_rotr(v1332, int64(32))
	F_MarkBufferDirty(m, v1334)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L19
	} else {
		goto L348
	}
L344:
	;
	v1374 = int32(8)
	F_PageInit(m, v1372, int32(8192), v1374)
	mBase = m.M
	v1376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1372)+16)))
	v1377 = v1372 + v1376
	*(*int32)(unsafe.Add(mBase, uint32(v1377))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1377)+6)) = uint16(v1374)
	v1384 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1372-int32(-64)))) = v1384
	*(*int64)(unsafe.Add(mBase, uint32(v1372)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1372)+32)) = v1384
	*(*int64)(unsafe.Add(mBase, uint32(v1372)+40)) = v1384
	*(*int64)(unsafe.Add(mBase, uint32(v1372)+48)) = v1384
	*(*int32)(unsafe.Add(mBase, uint32(v1372)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1372)+72)) = int32(2)
	v1398 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v1372)+12)) = uint16(v1398)
	goto L343
L345:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1358+(v1334^int32(-1))<<(uint(int32(2))%32))))
	v1372 = v1364
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1372 = v1366 + v1334<<(uint(int32(13))%32) + int32(-8192)
	goto L344
L348:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+56))
	if int32(0) < v1421 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1432 = int32(0)
	goto L352
L350:
	;
	goto L351
L351:
	;
	F_UnlockReleaseBuffer(m, v1334)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L19
	} else {
		goto L367
	}
L352:
	;
	v1442 = v1432 + int32(1)
	v1445 = F_XLogInitBufferForRedo(m, l0, v1442&int32(255))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L19
	} else {
		goto L355
	}
L353:
	;
	goto L351
L354:
	;
	v1465 = int32(4)
	if v1445 < int32(0) {
		goto L361
	} else {
		goto L362
	}
L355:
	;
	if v1445 < int32(0) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1450+(v1445^int32(-1))<<(uint(int32(2))%32))))
	v1464 = v1456
	goto L354
L357:
	;
	goto L358
L358:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1464 = v1458 + v1445<<(uint(int32(13))%32) + int32(-8192)
	goto L354
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1464)+4)) = base.I32_wrap_i64(v1332)
	*(*int32)(unsafe.Add(mBase, uint32(v1464))) = base.I32_wrap_i64(int64(base.Ui64(v1332) >> (uint(int64(32)) % 64)))
	F_MarkBufferDirty(m, v1445)
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L19
	} else {
		goto L364
	}
L360:
	;
	F_PageInit(m, v1483, int32(8192), int32(8))
	mBase = m.M
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1483)+16)))
	v1488 = v1483 + v1487
	*(*int32)(unsafe.Add(mBase, uint32(v1488))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1488)+6)) = uint16(v1465)
	goto L359
L361:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1469+(v1445^int32(-1))<<(uint(int32(2))%32))))
	v1483 = v1475
	goto L360
L362:
	;
	goto L363
L363:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1483 = v1477 + v1445<<(uint(int32(13))%32) + int32(-8192)
	goto L360
L364:
	;
	F_UnlockReleaseBuffer(m, v1445)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L19
	} else {
		goto L365
	}
L365:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+56))
	if v1442 < v1498 {
		v1432 = v1442
		goto L352
	} else {
		goto L366
	}
L366:
	;
	goto L353
L367:
	;
	goto L8
L368:
	;
	m.G0 = v15 - int32(-64)
	return
L369:
	;
	F_errmsg_internal(m, int32(427318), int32(0))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L19
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(523011), int32(419), int32(109594))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L19
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	F_errmsg_internal(m, int32(427253), int32(0))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L19
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(523011), int32(422), int32(109594))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L19
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	F_errmsg_internal(m, int32(427189), int32(0))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L19
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(523011), int32(427), int32(109594))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L19
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	F_errmsg_internal(m, int32(427914), int32(0))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L19
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(523011), int32(446), int32(428874))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L19
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	F_errmsg_internal(m, int32(426764), int32(0))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L19
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(523011), int32(579), int32(426367))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L19
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	F_errmsg_internal(m, int32(426764), int32(0))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L19
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(523011), int32(661), int32(428720))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L19
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v26
	F_errmsg_internal(m, int32(57468), v15)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L19
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(523011), int32(768), int32(255199))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L19
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_trgm_triconsistent(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 float64
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v491 float32
	_ = v491
	var v500 int32
	_ = v500
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	if base.Ui32(int32(11)) < base.Ui32(v17) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v500
L2:
	;
	v500 = base.F64_le(v39, base.F64_promote_f32(base.F32_div(v491, base.F32_convert_i32_s(v20)))) << (uint(int32(1)) % 32)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L13
	} else {
		goto L84
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = int32(1) << (uint(v17) % 32)
	if v23&int32(642) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v159 = v20 & int32(3)
	v161 = F_palloc(m, v20)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L38
	}
L6:
	;
	v135 = int32(0)
	v136 = int32(2)
	if v20 <= v135 {
		v500 = v136
		goto L1
	} else {
		goto L31
	}
L7:
	;
	if v23&int32(2072) != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v39 = F_index_strategy_get_limit(m, v17)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if int32(1)<<(uint(v17)%32)&int32(96) == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if int32(0) < v20 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v500 = int32(2)
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	if int32(0) < v20 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = v20 & int32(3)
	v47 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v20) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v20 != 0 {
		v491 = float32(0)
		goto L2
	} else {
		goto L30
	}
L18:
	;
	v53 = v47
	v54 = v47
	v58 = v2
	goto L21
L19:
	;
	v87 = v47
	v88 = v47
	goto L20
L20:
	;
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v65 = v53 + v21
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v67 = int32(0)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+3)))
	v81 = v54 + base.B2i32(v66 != v67) + base.B2i32(v70 != v67) + base.B2i32(v74 != v67) + base.B2i32(v78 != v67)
	v82 = int32(4)
	v83 = v53 + v82
	v85 = v58 + v82
	if v85 != v20&int32(2147483644) {
		v53 = v83
		v54 = v81
		v58 = v85
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v87 = v83
	v88 = v81
	goto L20
L23:
	;
	goto L22
L24:
	;
	v99 = v87
	v100 = v88
	v106 = v2
	goto L27
L25:
	;
	v122 = v88
	goto L26
L26:
	;
	v491 = base.F32_convert_i32_u(v122)
	goto L2
L27:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v21))))
	v115 = v100 + base.B2i32(v112 != int32(0))
	v116 = int32(1)
	v119 = v106 + v116
	if v119 != v46 {
		v99 = v99 + v116
		v100 = v115
		v106 = v119
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v122 = v115
	goto L26
L29:
	;
	goto L28
L30:
	;
	v500 = int32(0)
	goto L1
L31:
	;
	v139 = v135
	goto L32
L32:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+v21))))
	if v152 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v500 = int32(0)
	goto L1
L34:
	;
	v154 = v139 + int32(1)
	if v20 != v154 {
		v139 = v154
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v500 = v136
	goto L1
L38:
	;
	v163 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v20) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v168 = v163
	v175 = v2
	goto L42
L40:
	;
	v215 = v163
	goto L41
L41:
	;
	if v159 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v21))))
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168+v161))) = uint8(base.B2i32(v182 != v183))
	v187 = v168 | int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v187))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v187))) = uint8(base.B2i32(v190 != v183))
	v195 = v168 | int32(2)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v195))) = uint8(base.B2i32(v198 != v183))
	v203 = v168 | int32(3)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v203))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v203))) = uint8(base.B2i32(v206 != v183))
	v210 = int32(4)
	v211 = v168 + v210
	v213 = v175 + v210
	if v213 != v20&int32(2147483644) {
		v168 = v211
		v175 = v213
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v215 = v211
	goto L41
L44:
	;
	goto L43
L45:
	;
	v227 = v215
	v231 = int32(0)
	goto L48
L46:
	;
	goto L47
L47:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v263 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v276 = F___memset(m, v273, v263, v275)
	mBase = m.M
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v280 = F___memset(m, v277, v263, v279)
	mBase = m.M
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v263 < v281 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+v21))))
	*(*uint8)(unsafe.Add(mBase, uint32(v227+v161))) = uint8(base.B2i32(v241 != int32(0)))
	v245 = int32(1)
	v248 = v231 + v245
	if v248 != v159 {
		v227 = v227 + v245
		v231 = v248
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v161)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L13
	} else {
		goto L80
	}
L52:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v288 = v263
	v290 = v263
	goto L55
L53:
	;
	goto L54
L54:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	v355 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v355)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v262)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v366 = v355
	v368 = v263
	goto L67
L55:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v284+v290<<(uint(int32(2))%32))))
	v301 = v300 + v288
	if v300 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v340 = v290 + int32(1)
	if v340 != v281 {
		v288 = v301
		v290 = v340
		goto L55
	} else {
		goto L65
	}
L58:
	;
	v306 = v288
	goto L59
L59:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v306))))
	if v317 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+v290))) = uint8(v325)
	goto L57
L61:
	;
	v321 = v306 + int32(1)
	if v321 < v301 {
		v306 = v321
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L57
L65:
	;
	goto L56
L66:
	;
	goto L51
L67:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v357+v368<<(uint(int32(2))%32))))
	v380 = v360 + v377<<(uint(int32(3))%32)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if int32(0) < v381 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v455 = int32(0)
	goto L66
L69:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v389 = int32(0)
	v391 = v366
	goto L72
L70:
	;
	v431 = v366
	goto L71
L71:
	;
	v440 = v368 + int32(1)
	if v440 < v431 {
		v366 = v431
		v368 = v440
		goto L67
	} else {
		goto L79
	}
L72:
	;
	v401 = v385 + v389<<(uint(int32(3))%32)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v402))))
	if v404 != int32(1) {
		v422 = v391
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v431 = v422
	goto L71
L74:
	;
	v425 = v389 + int32(1)
	if v425 != v381 {
		v389 = v425
		v391 = v422
		goto L72
	} else {
		goto L78
	}
L75:
	;
	v407 = int32(1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	if v408 == v407 {
		v455 = v407
		goto L66
	} else {
		goto L76
	}
L76:
	;
	v411 = v408 + v354
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v412 != 0 {
		v422 = v391
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v411))) = uint8(v413)
	*(*int32)(unsafe.Add(mBase, uint32(v357+v391<<(uint(int32(2))%32)))) = v408
	v422 = v391 + v413
	goto L74
L78:
	;
	goto L73
L79:
	;
	goto L68
L80:
	;
	if v455 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v460 = int32(2)
	goto L83
L82:
	;
	v460 = int32(0)
	goto L83
L83:
	;
	v500 = v460
	goto L1
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	F_errmsg_internal(m, int32(505404), v15)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(521096), int32(352), int32(97939))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_tsquery_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v18 <= v2 {
		v40 = v2
		m.G0 = v10 + int32(16)
		return v40
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v14
		v23 = v13 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v25
		v29 = F_TS_execute_ternary(m, v23, v10+int32(4))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			switch v29 - int32(1) {
			case 0:
				v40 = int32(1)
			case 1:
				v36 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v36)
				v40 = v36
			default:
				v40 = v2
			}
			m.G0 = v10 + int32(16)
			return v40
		}
	}
}
func F_gin_tsquery_consistent_6args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(7) < v12 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v21 <= v19 {
			v43 = v2
			m.G0 = v10 + int32(16)
			return v43
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v17
			v26 = v16 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v28
			v32 = F_TS_execute_ternary(m, v26, v10+int32(4))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				switch v32 - int32(1) {
				case 0:
					v43 = int32(1)
				case 1:
					v39 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v39)
					v43 = v39
				default:
					v43 = v2
				}
				m.G0 = v10 + int32(16)
				return v43
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(128987), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(516035), int32(331), int32(165309))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(65067), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[20])) = v8
		return
	}
}
