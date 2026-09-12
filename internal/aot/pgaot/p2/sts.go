package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sts_attach(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v6 = F_palloc0(m, int32(68))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
		v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v14
		return v6
	}
}
func F_sts_end_write(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		F_BufFileWrite(m, v3, v4, int32(32768))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v12 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(int32(0)), int32(32768))
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v13 + int32(8)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = v17 + v18*int32(28) + int32(96)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24 + int32(4)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			F_BufFileClose(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_pfree(m, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v41 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v36+v37*int32(28))+100)) = uint8(v41)
					return
				}
			}
		}
	} else {
		return
	}
}
func F_sts_puttuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
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
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	v9 = m.G0
	v11 = v9 - int32(1040)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16 + int32(12)
	v26 = F_pg_snprintf(m, v11+int32(16), int32(1024), int32(465596), v11)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v52 = v50 + v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v28 = int32(4515120)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = F_BufFileCreateFileSet(m, v33, v11+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v42*int32(28))+100)) = uint8(v46)
	goto L3
L7:
	;
	m.G0 = v11 + int32(1040)
	return
L8:
	;
	if v183 != 0 {
		goto L45
	} else {
		goto L46
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui32(v53+v52) <= base.Ui32(v54) {
		v182 = v53
		v183 = v50
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v57 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if base.Ui32(v100+v52) <= base.Ui32(v101) {
		v182 = v100
		v183 = v103
		goto L8
	} else {
		goto L20
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = F_MemoryContextAllocZero(m, v60, int32(32768))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_BufFileWrite(m, v74, v57, int32(32768))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v69 = v67 + int32(32768)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v69
	v72 = v67 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v72
	v100 = v72
	v101 = v69
	goto L13
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v82 = F__emscripten_memset_bulkmem(m, v78, base.I32_extend8_s(int32(0)), int32(32768))
	mBase = m.M
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v83 + int32(8)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = v87 + v88*int32(28) + int32(96)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v94 + int32(4)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v100 = v99
	v101 = v98
	goto L13
L20:
	;
	if v103 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v112 = v100
	v113 = int32(0)
	goto L23
L23:
	;
	v114 = v112 + v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v116 = v115 - v114
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v112 = v108
	v113 = v110
	goto L23
L25:
	;
	v106 = F__emscripten_memcpy_bulkmem(m, v100, l1, v103)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v120 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v127 = v52 - (v116 + v125)
	if v127 == int32(0) {
		goto L7
	} else {
		goto L32
	}
L29:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, v114, l2, v116)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v131 = v116
	v133 = v127
	goto L33
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_BufFileWrite(m, v138, v139, int32(32768))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L7
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v147 = F__emscripten_memset_bulkmem(m, v143, base.I32_extend8_s(int32(0)), int32(32768))
	mBase = m.M
	goto L36
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v148 + int32(8)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = v152 + v153*int32(28) + int32(96)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v159 + int32(4)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v167 = base.I32_div_u_s(v133+int32(32759), int32(32760))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v172 = v171 - v169
	if base.Ui32(v172) < base.Ui32(v133) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v174 = v172
	goto L39
L38:
	;
	v174 = v133
	goto L39
L39:
	;
	if v174 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v177 + v174
	v181 = v133 - v174
	if v181 != 0 {
		v131 = v131 + v174
		v133 = v181
		goto L33
	} else {
		goto L44
	}
L41:
	;
	v175 = F__emscripten_memcpy_bulkmem(m, v169, v131+l2, v174)
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L34
L45:
	;
	if v183 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v191 = v182
	v192 = int32(0)
	goto L47
L47:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v194 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v191 = v187
	v192 = v189
	goto L47
L49:
	;
	v185 = F__emscripten_memcpy_bulkmem(m, v182, l1, v183)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v197 + v52
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v201 + int32(1)
	goto L7
L53:
	;
	v195 = F__emscripten_memcpy_bulkmem(m, v191+v192, l2, v194)
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
}
