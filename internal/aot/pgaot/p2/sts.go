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
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_sts_attach[0]))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		F_BufFileWrite(m, v3, v4, int32(_a_F_sts_end_write_0))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			base.MemoryFill(m, v8, int32(0), int32(_a_F_sts_end_write_0))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v12 + int32(8)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = v16 + v17*int32(28) + int32(96)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 + int32(4)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			F_BufFileClose(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_pfree(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v35+v36*int32(28))+100)) = uint8(v40)
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	v23 = v11 + int32(16)
	v26 = F_pg_snprintf(m, v23, int32(1024), int32(_a_F_sts_puttuple_0), v11)
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v51 = v49 + v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v28 = int32(_a_F_sts_puttuple_1)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_sts_puttuple[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_sts_puttuple[0])) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = F_BufFileCreateFileSet(m, v33, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_sts_puttuple[0])) = v29
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39+v40*int32(28))+100)) = uint8(v44)
	goto L3
L7:
	;
	m.G0 = v11 + int32(1040)
	return
L8:
	;
	if v178 != 0 {
		goto L40
	} else {
		goto L41
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui32(v52+v51) <= base.Ui32(v53) {
		v176 = v52
		v178 = v49
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v56 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if base.Ui32(v98+v51) <= base.Ui32(v99) {
		v176 = v98
		v178 = v101
		goto L8
	} else {
		goto L19
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v61 = F_MemoryContextAllocZero(m, v59, int32(_a_F_sts_puttuple_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_BufFileWrite(m, v73, v56, int32(_a_F_sts_puttuple_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v68 = v66 + int32(_a_F_sts_puttuple_2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v68
	v71 = v66 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v71
	v98 = v71
	v99 = v68
	goto L13
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	base.MemoryFill(m, v77, int32(0), int32(_a_F_sts_puttuple_2))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v81 + int32(8)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = v85 + v86*int32(28) + int32(96)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v92 + int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v98 = v97
	v99 = v96
	goto L13
L19:
	;
	if v101 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v101 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = v98
	v110 = int32(0)
	goto L22
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v112 = v110 + v109
	v113 = v111 - v112
	if v113 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	base.MemoryCopy(m, v98, l1, v101)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v109 = v105
	v110 = v107
	goto L22
L26:
	;
	base.MemoryCopy(m, v112, l2, v113)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v123 = v51 - (v113 + v121)
	if v123 == int32(0) {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v129 = v123
	v130 = v113
	goto L30
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_BufFileWrite(m, v134, v135, int32(_a_F_sts_puttuple_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L7
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	base.MemoryFill(m, v139, int32(0), int32(_a_F_sts_puttuple_2))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v143 + int32(8)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = v147 + v148*int32(28) + int32(96)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v154 + int32(4)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v162 = base.I32_div_u_s(v129+int32(_a_F_sts_puttuple_3), int32(_a_F_sts_puttuple_4))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+4)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v166 = v164 - v165
	if base.Ui32(v166) < base.Ui32(v129) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v168 = v166
	goto L35
L34:
	;
	v168 = v129
	goto L35
L35:
	;
	if v168 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	base.MemoryCopy(m, v165, l2+v130, v168)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v171 + v168
	v175 = v129 - v168
	if v175 != 0 {
		v129 = v175
		v130 = v168 + v130
		goto L30
	} else {
		goto L39
	}
L39:
	;
	goto L31
L40:
	;
	if v178 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v184 = v176
	v185 = int32(0)
	goto L42
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v186 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	base.MemoryCopy(m, v176, l1, v178)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v184 = v180
	v185 = v182
	goto L42
L46:
	;
	base.MemoryCopy(m, v185+v184, l2, v186)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v189 + v51
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193 + int32(1)
	goto L7
}
