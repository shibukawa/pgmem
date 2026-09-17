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
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
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
		v82 = int32(-1)
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
	v96 = v82
	v98 = v51
	v99 = v2
	v100 = v2
	v102 = v2
	goto L20
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v58 = int32(0)
	goto L16
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v58*int32(48))+4)))
	if v73 != int32(1) {
		v82 = v58
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v82 = int32(-1)
	goto L14
L18:
	;
	v77 = v58 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v77)) < base.Ui64(v52) {
		v58 = v77
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v109 = v96
	v115 = v102
	v116 = v102
	goto L24
L22:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v96 = v132
	v98 = v171
	v99 = v169
	v100 = v170
	v102 = v129
	goto L20
L23:
	;
	if int32(2) <= v100 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	if v116&int32(1) != 0 {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+5)))
	if v138 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v123 = int32(1)
	v124 = v109 - v123
	v128 = base.B2i32(v122&(v124^v82) == int32(0))
	v129 = v128 | v115
	v132 = v122 & v124
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	v134 = v109*int32(48) + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+4)))
	if v135 != v123 {
		v109 = v132
		v115 = v129
		v116 = v128
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v99<<(uint(int32(2))%32)))) = v134
	v169 = v99 + int32(1)
	v170 = v100
	goto L22
L29:
	;
	goto L30
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v148+v100<<(uint(int32(2))%32)))) = v134
	v169 = v99
	v170 = v100 + int32(1)
	goto L22
L31:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_pg_qsort(m, v157, v100, int32(4), int32(817))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v99 < int32(2) {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_pg_qsort(m, v164, v99, int32(4), int32(817))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v76 int32
	_ = v76
	var __phi76 int32
	_ = __phi76
	var v77 int32
	_ = v77
	var __phi77 int32
	_ = __phi77
	var v78 int32
	_ = v78
	var __phi78 int32
	_ = __phi78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_tbm_create_pagetable(m, l0)
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
	v19 = l1 & int32(-256)
	v21 = l1 & int32(255)
	if v21 == int32(0) {
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
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_pagetable_insert(m, v147, v19, v11+int32(15))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = int32(16)
	v34 = (int32(base.Ui32(l1)>>(uint(v30)%32)) ^ l1) * int32(-2048144789)
	v39 = (int32(base.Ui32(v34)>>(uint(int32(13))%32)) ^ v34) * int32(-1028477387)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v48 = int32(base.Ui32(v39)>>(uint(v30)%32)) ^ v39
	goto L9
L8:
	;
	if v134 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L9:
	;
	v52 = v48 & v44
	v55 = v43 + v52*int32(48)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	switch v56 {
	case 0:
		v134 = int32(0)
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v48 = v52 + int32(1)
	goto L9
L12:
	;
	goto L8
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v57 != l1 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v59 - v60
	v66 = v44 & (v52 + v60)
	v69 = v43 + v66*int32(48)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v70 != v60 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)) = uint8(v126)
	v134 = v60
	goto L12
L16:
	;
	v121 = v55
	goto L15
L17:
	;
	goto L18
L18:
	;
	__phi74 = v66
	__phi76 = v55
	__phi77 = v69
	__phi78 = v44
	v74 = __phi74
	v76 = __phi76
	v77 = __phi77
	v78 = __phi78
	goto L19
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v81 = int32(16)
	v85 = (int32(base.Ui32(v80)>>(uint(v81)%32)) ^ v80) * int32(-2048144789)
	v90 = (int32(base.Ui32(v85)>>(uint(int32(13))%32)) ^ v85) * int32(-1028477387)
	if v74 == (int32(base.Ui32(v90)>>(uint(v81)%32))^v90)&v78 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v121 = v77
	goto L15
L21:
	;
	v121 = v76
	goto L15
L22:
	;
	goto L23
L23:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v77)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v77)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v110 = int32(1)
	v112 = v109 & (v74 + v110)
	v115 = v108 + v112*int32(48)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
	if v116 == v110 {
		__phi74 = v112
		__phi76 = v77
		__phi77 = v115
		__phi78 = v109
		v74 = __phi74
		v76 = __phi76
		v77 = __phi77
		v78 = __phi78
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v140 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v139 - v140
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v143 - v140
	goto L6
L26:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v152 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v218 = v150 + int32(base.Ui32(v21)>>(uint(int32(3))%32))&int32(28)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v219 | int32(1)<<(uint(l1)%32)
	m.G0 = v11 + int32(16)
	return
L28:
	;
	v201 = l0 + v198
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v202 + int32(1)
	v206 = l0 + v199
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v207 + v200
	goto L27
L29:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)))
	v156 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v150)+4)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v150)+12)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v150)+20)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v150)+28)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v150)+36)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v150)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v19
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)) = uint8(v155)
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+5)) = uint8(v170)
	v198 = int32(16)
	v199 = int32(28)
	v200 = v170
	goto L28
L30:
	;
	goto L31
L31:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+5)))
	if v175 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)))
	v177 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v150)+4)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v150)+12)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v150)+20)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v150)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v150)+36)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v150)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v19
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)) = uint8(v176)
	v191 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v191
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+5)) = uint8(v191)
	v198 = int32(28)
	v199 = int32(24)
	v200 = int32(-1)
	goto L28
}
