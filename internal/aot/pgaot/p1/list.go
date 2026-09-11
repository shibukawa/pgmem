package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_list_concat_unique(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v3 = int32(0)
	if l1 == v3 {
		v67 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v67
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 <= int32(0) {
		v67 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = l0
	v14 = v3
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = v18 + v14<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v12 == int32(0) {
		v52 = v22
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v67 = v57
	goto L1
L6:
	;
	v64 = v14 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 < v65 {
		v12 = v57
		v14 = v64
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v55 = F_lappend(m, v12, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v26 <= v25 {
		v52 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = v25
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v33<<(uint(int32(2))%32))))
	v40 = F_equal(m, v39, v22)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v52 = v48
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	if v40 != 0 {
		v57 = v12
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v33 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v45 < v46 {
		v33 = v45
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v57 = v55
	goto L6
L17:
	;
	goto L5
}
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v3 = int32(0)
	if l0 == v3 {
		v57 = v3
		return v57
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = int32(0)
		if v10 < l1 {
			v13 = l1
		} else {
			v13 = v10
		}
		if v9 <= v13 {
			v57 = v3
			return v57
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = int32(8)
			v19 = v9 - v13
			v21 = v19 + int32(4)
			if v21 <= v18 {
				v24 = v18
			} else {
				v24 = v21
			}
			if v24&(v24-int32(1)) != 0 {
				v31 = int32(1) << (uint(int32(32)-base.I32_clz(v24)) % 32)
			} else {
				v31 = v24
			}
			v33 = v31 - int32(4)
			v38 = F_palloc(m, v33<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v15
				v46 = v38 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v46
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v49 = int32(2)
				v53 = v19 << (uint(v49) % 32)
				if v53 != 0 {
					v54 = F__emscripten_memcpy_bulkmem(m, v46, v48+v13<<(uint(v49)%32), v53)
					mBase = m.M
				} else {
				}
				v57 = v38
				return v57
			}
		}
	}
}
func F_list_delete_cell(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0+int32(16) != v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = l1 + int32(4)
	v25 = int32(2)
	v29 = (v6 + int32(base.Ui32(l1-v5^int32(-1))>>(uint(v25)%32))) << (uint(v25) % 32)
	if l1 == v21 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	F_pfree(m, v5)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v174 - int32(1)
	return l0
L11:
	;
	goto L10
L12:
	;
	v33 = l1 + v29
	if base.Ui32(v21-v33) <= base.Ui32(int32(0)-v29<<(uint(int32(1))%32)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = F___memcpy(m, l1, v21, v29)
	mBase = m.M
	goto L10
L14:
	;
	goto L15
L15:
	;
	v43 = (l1 ^ v21) & int32(3)
	if base.Ui32(l1) < base.Ui32(v21) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v145 == int32(0) {
		goto L11
	} else {
		goto L52
	}
L17:
	;
	if base.Ui32(v123) <= base.Ui32(int32(3)) {
		v144 = v122
		v145 = v123
		v146 = v124
		goto L16
	} else {
		goto L48
	}
L18:
	;
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v43 != 0 {
		v105 = v29
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v144 = v21
	v145 = v29
	v146 = l1
	goto L16
L22:
	;
	goto L23
L23:
	;
	if l1&int32(3) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v122 = v21
	v123 = v29
	v124 = l1
	goto L17
L25:
	;
	goto L26
L26:
	;
	v50 = v21
	v51 = v29
	v52 = l1
	goto L27
L27:
	;
	if v51 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v122 = v59
	v123 = v61
	v124 = v63
	goto L17
L29:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v56)
	v58 = int32(1)
	v59 = v50 + v58
	v61 = v51 - v58
	v63 = v52 + v58
	if v63&int32(3) != 0 {
		v50 = v59
		v51 = v61
		v52 = v63
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v105 == int32(0) {
		goto L11
	} else {
		goto L44
	}
L32:
	;
	if v33&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v70 = v29
	goto L36
L34:
	;
	v85 = v29
	goto L35
L35:
	;
	if base.Ui32(v85) <= base.Ui32(int32(3)) {
		v105 = v85
		goto L31
	} else {
		goto L40
	}
L36:
	;
	if v70 == int32(0) {
		goto L11
	} else {
		goto L38
	}
L37:
	;
	v85 = v76
	goto L35
L38:
	;
	v76 = v70 - int32(1)
	v77 = l1 + v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v79)
	if v77&int32(3) != 0 {
		v70 = v76
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v92 = v85
	goto L41
L41:
	;
	v96 = v92 - int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21+v96)))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v96))) = v99
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v92 = v96
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v105 = v96
	goto L31
L43:
	;
	goto L42
L44:
	;
	v112 = v105
	goto L45
L45:
	;
	v116 = v112 - int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v116))) = uint8(v119)
	if v116 != 0 {
		v112 = v116
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L11
L47:
	;
	goto L46
L48:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L49
L49:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
	v135 = int32(4)
	v136 = v129 + v135
	v138 = v131 + v135
	v140 = v130 - v135
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		v129 = v136
		v130 = v140
		v131 = v138
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v144 = v136
	v145 = v140
	v146 = v138
	goto L16
L51:
	;
	goto L50
L52:
	;
	v151 = v144
	v152 = v145
	v153 = v146
	goto L53
L53:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	v157 = int32(1)
	v162 = v152 - v157
	if v162 != 0 {
		v151 = v151 + v157
		v152 = v162
		v153 = v153 + v157
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L11
L55:
	;
	goto L54
}
func F_list_difference_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v13 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = v3
	v22 = v3
	goto L10
L8:
	;
	v79 = v3
	goto L9
L9:
	;
	return v79
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v29 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v79 = v67
	goto L9
L12:
	;
	v71 = v22 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v21 = v67
		v22 = v71
		goto L10
	} else {
		goto L22
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v58 = F_lappend(m, v21, v28)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32+v36<<(uint(int32(2))%32))))
	if v45 == v28 {
		v67 = v21
		goto L12
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v48 = v36 + int32(1)
	if v29 != v48 {
		v36 = v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(0)
L21:
	;
	v67 = v58
	goto L12
L22:
	;
	goto L11
L23:
	;
	return int32(0)
L24:
	;
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(8)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = v91 + int32(4)
	if v93 <= v90 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = v90
	goto L28
L27:
	;
	v96 = v93
	goto L28
L28:
	;
	if v96&(v96-int32(1)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v103 = int32(1) << (uint(int32(32)-base.I32_clz(v96)) % 32)
	goto L31
L30:
	;
	v103 = v96
	goto L31
L31:
	;
	v105 = v103 - int32(4)
	v110 = F_palloc(m, v105<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v87
	v116 = v110 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v120 = v91 << (uint(int32(2)) % 32)
	if v120 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	return v110
L34:
	;
	v121 = F__emscripten_memcpy_bulkmem(m, v116, v118, v120)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
}
func F_list_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v3 != l0+int32(16) {
			F_pfree(m, v3)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(17179869187)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v6 + int32(16)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v17
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v19
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v21
		return v6
	}
}
func F_list_sort(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 < int32(2) {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pg_qsort(m, v9, v6, int32(4), l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_list_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8
		return
	}
}
