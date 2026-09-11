package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ListComparatorForWalSummaryFiles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	return base.B2i32(base.Ui64(v8) < base.Ui64(v6)) - base.B2i32(base.Ui64(v6) < base.Ui64(v8))
}
func F_cmp_list_len_contents_asc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = v8
	goto L3
L2:
	;
	v9 = v3
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = v11
	goto L6
L5:
	;
	v12 = v3
	goto L6
L6:
	;
	v15 = base.B2i32(v12 < v9) - base.B2i32(v9 < v12)
	if v12 != v9 {
		v61 = v15
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return v61
L8:
	;
	v20 = int32(0)
	goto L9
L9:
	;
	v24 = int32(0)
	if v7 == v24 {
		v34 = v24
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = int32(-1)
	goto L7
L11:
	;
	if v10 == int32(0) {
		v61 = v15
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v28 <= v20 {
		v34 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v34 = v30 + v20<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v37 <= v20 {
		v61 = v15
		goto L7
	} else {
		goto L15
	}
L15:
	;
	if v34 == int32(0) {
		v61 = v15
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v44 = v41 + v20<<(uint(int32(2))%32)
	if v44 == int32(0) {
		v61 = v15
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v48 < v47 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	if v48 <= v47 {
		v20 = v20 + int32(1)
		goto L9
	} else {
		goto L21
	}
L21:
	;
	goto L10
}
func F_list_append_unique_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v34
L2:
	;
	v30 = F_lappend_oid(m, l0, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = v3
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14<<(uint(int32(2))%32))))
	if v20 == l1 {
		v34 = l0
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v23 = v14 + int32(1)
	if v8 != v23 {
		v14 = v23
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v34 = v30
	goto L1
}
func F_list_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = v14 + int32(4)
		if v16 <= v13 {
			v19 = v13
		} else {
			v19 = v16
		}
		if v19&(v19-int32(1)) != 0 {
			v26 = int32(1) << (uint(int32(32)-base.I32_clz(v19)) % 32)
		} else {
			v26 = v19
		}
		v28 = v26 - int32(4)
		v33 = F_palloc(m, v28<<(uint(int32(2))%32)+int32(16))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = v10
			v41 = v33 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v45 = v14 << (uint(int32(2)) % 32)
			if v45 != 0 {
				v46 = F__emscripten_memcpy_bulkmem(m, v41, v43, v45)
				mBase = m.M
			} else {
			}
			return v33
		}
	}
}
func F_list_delete_first_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v181 int32
	_ = v181
	if int32(0) < l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v181
L2:
	;
	v7 = int32(0)
	if l0 == v7 {
		v181 = v7
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v181 = l0
	goto L1
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= l1 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l0+int32(16) != v10 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v24 = int32(2)
	v26 = v10 + l1<<(uint(v24)%32)
	v29 = (v11 - l1) << (uint(v24) % 32)
	if v10 == v26 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	F_pfree(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_pfree(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v174 - l1
	goto L4
L16:
	;
	goto L15
L17:
	;
	v33 = v10 + v29
	if base.Ui32(v26-v33) <= base.Ui32(int32(0)-v29<<(uint(int32(1))%32)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v40 = F___memcpy(m, v10, v26, v29)
	mBase = m.M
	goto L15
L19:
	;
	goto L20
L20:
	;
	v43 = (v10 ^ v26) & int32(3)
	if base.Ui32(v10) < base.Ui32(v26) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v145 == int32(0) {
		goto L16
	} else {
		goto L57
	}
L22:
	;
	if base.Ui32(v123) <= base.Ui32(int32(3)) {
		v144 = v122
		v145 = v123
		v146 = v124
		goto L21
	} else {
		goto L53
	}
L23:
	;
	if v43 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v43 != 0 {
		v105 = v29
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v144 = v26
	v145 = v29
	v146 = v10
	goto L21
L27:
	;
	goto L28
L28:
	;
	if v10&int32(3) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v122 = v26
	v123 = v29
	v124 = v10
	goto L22
L30:
	;
	goto L31
L31:
	;
	v50 = v26
	v51 = v29
	v52 = v10
	goto L32
L32:
	;
	if v51 == int32(0) {
		goto L16
	} else {
		goto L34
	}
L33:
	;
	v122 = v59
	v123 = v61
	v124 = v63
	goto L22
L34:
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
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if v105 == int32(0) {
		goto L16
	} else {
		goto L49
	}
L37:
	;
	if v33&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v70 = v29
	goto L41
L39:
	;
	v85 = v29
	goto L40
L40:
	;
	if base.Ui32(v85) <= base.Ui32(int32(3)) {
		v105 = v85
		goto L36
	} else {
		goto L45
	}
L41:
	;
	if v70 == int32(0) {
		goto L16
	} else {
		goto L43
	}
L42:
	;
	v85 = v76
	goto L40
L43:
	;
	v76 = v70 - int32(1)
	v77 = v10 + v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v79)
	if v77&int32(3) != 0 {
		v70 = v76
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v92 = v85
	goto L46
L46:
	;
	v96 = v92 - int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v26+v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+v96))) = v99
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v92 = v96
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v105 = v96
	goto L36
L48:
	;
	goto L47
L49:
	;
	v112 = v105
	goto L50
L50:
	;
	v116 = v112 - int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v116))) = uint8(v119)
	if v116 != 0 {
		v112 = v116
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L16
L52:
	;
	goto L51
L53:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L54
L54:
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
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v144 = v136
	v145 = v140
	v146 = v138
	goto L21
L56:
	;
	goto L55
L57:
	;
	v151 = v144
	v152 = v145
	v153 = v146
	goto L58
L58:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	v157 = int32(1)
	v162 = v152 - v157
	if v162 != 0 {
		v151 = v151 + v157
		v152 = v162
		v153 = v153 + v157
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L16
L60:
	;
	goto L59
}
