package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_begin_private_iterate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v2 = int32(0)
	v15 = F_palloc(m, int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	return v15
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v27 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v29 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = F_MemoryContextAlloc(m, v32, v29<<(uint(int32(2))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v35
	goto L6
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	if v52 == int64(0) {
		v83 = int32(-1)
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v40 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = F_MemoryContextAlloc(m, v43, v40<<(uint(int32(2))%32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v46
	goto L10
L14:
	;
	v95 = v83
	v98 = v2
	v99 = v2
	v100 = v51
	v102 = v2
	goto L20
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v59 = int32(0)
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v59*int32(48))+4)))
	if v73 != int32(1) {
		v83 = v59
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v83 = int32(-1)
	goto L14
L18:
	;
	v77 = v59 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v77)) < base.Ui64(v52) {
		v59 = v77
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v108 = v95
	v115 = v102
	v117 = v102
	goto L24
L22:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v95 = v132
	v98 = v177
	v99 = v178
	v100 = v179
	v102 = v129
	goto L20
L23:
	;
	if int32(2) <= v98 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	if v117&int32(1) != 0 {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	if v134 == int32(0) {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v123 = int32(1)
	v124 = v108 - v123
	v128 = base.B2i32(v122&(v124^v83) == int32(0))
	v129 = v128 | v115
	v132 = v122 & v124
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v134 = v108*int32(48) + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+4)))
	if v135 != v123 {
		v108 = v132
		v115 = v129
		v117 = v128
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+5)))
	if v140 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v143+v99<<(uint(int32(2))%32)))) = v134
	v177 = v98
	v178 = v99 + int32(1)
	goto L22
L30:
	;
	goto L31
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v150+v98<<(uint(int32(2))%32)))) = v134
	v177 = v98 + int32(1)
	v178 = v99
	goto L22
L32:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_pg_qsort(m, v165, v98, int32(4), int32(816))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v99 < int32(2) {
		goto L3
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_pg_qsort(m, v172, v99, int32(4), int32(816))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L3
}
func F_tbm_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	return base.B2i32(base.Ui32(v6) < base.Ui32(v4)) - base.B2i32(base.Ui32(v4) < base.Ui32(v6))
}
func F_tbm_end_iterate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
		return
	}
}
func F_tbm_is_empty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	return base.B2i32(v2 == int32(0))
}
func F_tbm_mark_page_lossy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var __phi75 int32
	_ = __phi75
	var v76 int32
	_ = v76
	var __phi76 int32
	_ = __phi76
	var v78 int32
	_ = v78
	var __phi78 int32
	_ = __phi78
	var v79 int32
	_ = v79
	var __phi79 int32
	_ = __phi79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_tbm_create_pagetable(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = l1 & int32(-256)
	v22 = l1 & int32(255)
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = F_pagetable_insert(m, v148, v20, v12+int32(15))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = int32(16)
	v35 = (int32(base.Ui32(l1)>>(uint(v31)%32)) ^ l1) * int32(-2048144789)
	v40 = (int32(base.Ui32(v35)>>(uint(int32(13))%32)) ^ v35) * int32(-1028477387)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v48 = int32(base.Ui32(v40)>>(uint(v31)%32)) ^ v40
	goto L9
L8:
	;
	if v135 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L9:
	;
	v53 = v48 & v45
	v56 = v44 + v53*int32(48)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	switch v57 {
	case 0:
		v135 = int32(0)
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v48 = v53 + int32(1)
	goto L9
L12:
	;
	goto L8
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v58 != l1 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v60 - v61
	v67 = v45 & (v53 + v61)
	v70 = v44 + v67*int32(48)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	if v71 != v61 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)) = uint8(v127)
	v135 = v61
	goto L12
L16:
	;
	v123 = v56
	goto L15
L17:
	;
	goto L18
L18:
	;
	__phi75 = v67
	__phi76 = v56
	__phi78 = v70
	__phi79 = v45
	v75 = __phi75
	v76 = __phi76
	v78 = __phi78
	v79 = __phi79
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v82 = int32(16)
	v86 = (int32(base.Ui32(v81)>>(uint(v82)%32)) ^ v81) * int32(-2048144789)
	v91 = (int32(base.Ui32(v86)>>(uint(int32(13))%32)) ^ v86) * int32(-1028477387)
	if v75 == (int32(base.Ui32(v91)>>(uint(v82)%32))^v91)&v79 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v123 = v78
	goto L15
L21:
	;
	v123 = v76
	goto L15
L22:
	;
	goto L23
L23:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v78)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v78)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v105
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v111 = int32(1)
	v113 = v110 & (v75 + v111)
	v116 = v109 + v113*int32(48)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
	if v117 == v111 {
		__phi75 = v113
		__phi76 = v78
		__phi78 = v116
		__phi79 = v110
		v75 = __phi75
		v76 = __phi76
		v78 = __phi78
		v79 = __phi79
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v141 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v140 - v141
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v144 - v141
	goto L6
L26:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v153 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v270 = v151 + int32(base.Ui32(v22)>>(uint(int32(3))%32))&int32(28) + int32(8)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v271 | int32(1)<<(uint(l1)%32)
	m.G0 = v12 + int32(16)
	return
L28:
	;
	v250 = v249 + l0
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v251 + int32(1)
	v255 = l0 + v246
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v256 + v245
	goto L27
L29:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
	if v151&int32(3) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
	if v199 != 0 {
		goto L27
	} else {
		goto L41
	}
L32:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)) = uint8(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v20
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)) = uint8(v156)
	v245 = v192
	v246 = int32(28)
	v249 = int32(16)
	goto L28
L33:
	;
	v162 = v151 + int32(48)
	if base.Ui32(v162) <= base.Ui32(v151) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v178 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+6)) = v178
	v180 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+46)) = uint16(v180)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+38)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v151)+30)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v151)+22)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v151)+14)) = v178
	goto L32
L36:
	;
	v168 = v151 + int32(4)
	if base.Ui32(v168) < base.Ui32(v162) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v170 = v162
	goto L39
L38:
	;
	v170 = v168
	goto L39
L39:
	;
	v177 = F__emscripten_memset_bulkmem(m, v151, base.I32_extend8_s(int32(0)), (v151^int32(-1)+v170)&int32(-4)+int32(4))
	mBase = m.M
	goto L40
L40:
	;
	goto L32
L41:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
	if v151&int32(3) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v236 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v236
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)) = uint8(v236)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v20
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)) = uint8(v200)
	v245 = int32(-1)
	v246 = int32(24)
	v249 = int32(28)
	goto L28
L43:
	;
	v206 = v151 + int32(48)
	if base.Ui32(v206) <= base.Ui32(v151) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v222 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+6)) = v222
	v224 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+46)) = uint16(v224)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+38)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v151)+30)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v151)+22)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v151)+14)) = v222
	goto L42
L46:
	;
	v212 = v151 + int32(4)
	if base.Ui32(v212) < base.Ui32(v206) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v214 = v206
	goto L49
L48:
	;
	v214 = v212
	goto L49
L49:
	;
	v221 = F__emscripten_memset_bulkmem(m, v151, base.I32_extend8_s(int32(0)), (v151^int32(-1)+v214)&int32(-4)+int32(4))
	mBase = m.M
	goto L50
L50:
	;
	goto L42
}
