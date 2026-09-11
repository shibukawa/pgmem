package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmp_list_len_asc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v7 = v6
	} else {
		v7 = v3
	}
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v10 = v9
	} else {
		v10 = v3
	}
	return base.B2i32(v10 < v7) - base.B2i32(v7 < v10)
}
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	if l0 == int32(0) {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v16 = int32(8)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = v17 + int32(4)
			if v19 <= v16 {
				v22 = v16
			} else {
				v22 = v19
			}
			if v22&(v22-int32(1)) != 0 {
				v29 = int32(1) << (uint(int32(32)-base.I32_clz(v22)) % 32)
			} else {
				v29 = v22
			}
			v31 = v29 - int32(4)
			v36 = F_palloc(m, v31<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v13
				v44 = v36 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v44
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v48 = v17 << (uint(int32(2)) % 32)
				if v48 != 0 {
					v49 = F__emscripten_memcpy_bulkmem(m, v44, v46, v48)
					mBase = m.M
				} else {
				}
				return v36
			}
		}
	} else {
		if l1 == int32(0) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v57 = int32(8)
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v60 = v58 + int32(4)
			if v60 <= v57 {
				v63 = v57
			} else {
				v63 = v60
			}
			if v63&(v63-int32(1)) != 0 {
				v70 = int32(1) << (uint(int32(32)-base.I32_clz(v63)) % 32)
			} else {
				v70 = v63
			}
			v72 = v70 - int32(4)
			v77 = F_palloc(m, v72<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v72
				*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v77))) = v54
				v83 = v77 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v83
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v87 = v58 << (uint(int32(2)) % 32)
				if v87 != 0 {
					v88 = F__emscripten_memcpy_bulkmem(m, v83, v85, v87)
					mBase = m.M
				} else {
				}
				return v77
			}
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v94 = int32(8)
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v97 = v95 + v96
			v99 = v97 + int32(4)
			if v99 <= v94 {
				v102 = v94
			} else {
				v102 = v99
			}
			if v102&(v102-int32(1)) != 0 {
				v109 = int32(1) << (uint(int32(32)-base.I32_clz(v102)) % 32)
			} else {
				v109 = v102
			}
			v111 = v109 - int32(4)
			v116 = F_palloc(m, v111<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v111
				*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v97
				*(*int32)(unsafe.Add(mBase, uint32(v116))) = v91
				v122 = v116 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v122
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v127 = v125 << (uint(int32(2)) % 32)
				if v127 != 0 {
					v128 = F__emscripten_memcpy_bulkmem(m, v122, v124, v127)
					mBase = m.M
					v129 = v128
				} else {
					v129 = v122
				}
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v131 = int32(2)
				v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v137 = v135 << (uint(v131) % 32)
				if v137 != 0 {
					v138 = F__emscripten_memcpy_bulkmem(m, v129+v130<<(uint(v131)%32), v134, v137)
					mBase = m.M
				} else {
				}
				return v116
			}
		}
	}
}
func F_list_delete_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v175 = int32(0)
	goto L3
L3:
	;
	return v175
L4:
	;
	if l0+int32(16) != v4 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v19 = int32(4)
	v20 = v4 + v19
	v24 = v5<<(uint(int32(2))%32) - v19
	if v4 == v20 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_pfree(m, v4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169 - int32(1)
	v175 = l0
	goto L3
L14:
	;
	goto L13
L15:
	;
	v28 = v4 + v24
	if base.Ui32(v20-v28) <= base.Ui32(int32(0)-v24<<(uint(int32(1))%32)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v35 = F___memcpy(m, v4, v20, v24)
	mBase = m.M
	goto L13
L17:
	;
	goto L18
L18:
	;
	v38 = (v4 ^ v20) & int32(3)
	if base.Ui32(v4) < base.Ui32(v20) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v140 == int32(0) {
		goto L14
	} else {
		goto L55
	}
L20:
	;
	if base.Ui32(v118) <= base.Ui32(int32(3)) {
		v139 = v117
		v140 = v118
		v141 = v119
		goto L19
	} else {
		goto L51
	}
L21:
	;
	if v38 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v38 != 0 {
		v100 = v24
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v139 = v20
	v140 = v24
	v141 = v4
	goto L19
L25:
	;
	goto L26
L26:
	;
	if v4&int32(3) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v117 = v20
	v118 = v24
	v119 = v4
	goto L20
L28:
	;
	goto L29
L29:
	;
	v45 = v20
	v46 = v24
	v47 = v4
	goto L30
L30:
	;
	if v46 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L31:
	;
	v117 = v54
	v118 = v56
	v119 = v58
	goto L20
L32:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v51)
	v53 = int32(1)
	v54 = v45 + v53
	v56 = v46 - v53
	v58 = v47 + v53
	if v58&int32(3) != 0 {
		v45 = v54
		v46 = v56
		v47 = v58
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v100 == int32(0) {
		goto L14
	} else {
		goto L47
	}
L35:
	;
	if v28&int32(3) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v65 = v24
	goto L39
L37:
	;
	v80 = v24
	goto L38
L38:
	;
	if base.Ui32(v80) <= base.Ui32(int32(3)) {
		v100 = v80
		goto L34
	} else {
		goto L43
	}
L39:
	;
	if v65 == int32(0) {
		goto L14
	} else {
		goto L41
	}
L40:
	;
	v80 = v71
	goto L38
L41:
	;
	v71 = v65 - int32(1)
	v72 = v4 + v71
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v71))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v74)
	if v72&int32(3) != 0 {
		v65 = v71
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v87 = v80
	goto L44
L44:
	;
	v91 = v87 - int32(4)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20+v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v4+v91))) = v94
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v87 = v91
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v100 = v91
	goto L34
L46:
	;
	goto L45
L47:
	;
	v107 = v100
	goto L48
L48:
	;
	v111 = v107 - int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4+v111))) = uint8(v114)
	if v111 != 0 {
		v107 = v111
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L14
L50:
	;
	goto L49
L51:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L52
L52:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	v130 = int32(4)
	v131 = v124 + v130
	v133 = v126 + v130
	v135 = v125 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v124 = v131
		v125 = v135
		v126 = v133
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v139 = v131
	v140 = v135
	v141 = v133
	goto L19
L54:
	;
	goto L53
L55:
	;
	v146 = v139
	v147 = v140
	v148 = v141
	goto L56
L56:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v150)
	v152 = int32(1)
	v157 = v147 - v152
	if v157 != 0 {
		v146 = v146 + v152
		v147 = v157
		v148 = v148 + v152
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L14
L58:
	;
	goto L57
}
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(17179869185)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v5 + int32(16)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v15
		return v5
	}
}
func F_list_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v16 = v3
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = F_equal(m, v22, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v23
L9:
	;
	return int32(0)
L10:
	;
	if v23 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = v16 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 < v31 {
		v16 = v30
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L8
L14:
	;
	goto L13
}
func F_list_next_fn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = v4 + int32(4)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if base.Ui32(v11) < base.Ui32(v14+v15<<(uint(int32(2))%32)) {
			v20 = v11
		} else {
			v20 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
		return v9
	}
}
