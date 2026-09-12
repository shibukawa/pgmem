package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_nameeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	goto L3
L3:
	;
	if v5&int32(3) == int32(0) {
		v80 = v5
		goto L20
	} else {
		goto L21
	}
L4:
	;
	return base.B2i32(v45-v46 == int32(0))
L6:
	;
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L12
L9:
	;
	v41 = v4
	v45 = int32(0)
	goto L10
L10:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L4
L11:
	;
	v41 = v36
	v45 = v38
	goto L10
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v36 = v30
	v38 = int32(0)
	goto L11
L14:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	if v4&int32(3) == int32(0) {
		v137 = v4
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v113 = v105 - v5
	goto L18
L20:
	;
	v84 = v80
	goto L29
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v64 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v113 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v69 = v5
	goto L25
L25:
	;
	v73 = v69 + int32(1)
	if v73&int32(3) == int32(0) {
		v80 = v73
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v105 = v73
	goto L19
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v78 != 0 {
		v69 = v73
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 == v93 {
		v84 = v84 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v99 = v84
	goto L32
L31:
	;
	goto L30
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v103 != 0 {
		v99 = v99 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v105 = v99
	goto L19
L34:
	;
	goto L33
L35:
	;
	v171 = F_varstr_cmp(m, v5, v113, v4, v170, v6)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v170 = v162 - v4
	goto L35
L37:
	;
	v141 = v137
	goto L46
L38:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v121 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v170 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v126 = v4
	goto L42
L42:
	;
	v130 = v126 + int32(1)
	if v130&int32(3) == int32(0) {
		v137 = v130
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v162 = v130
	goto L36
L44:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v135 != 0 {
		v126 = v130
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v150 = int32(-2139062144)
	if (int32(16843008)-v147|v147)&v150 == v150 {
		v141 = v141 + int32(4)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v156 = v141
	goto L49
L48:
	;
	goto L47
L49:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v160 != 0 {
		v156 = v156 + int32(1)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v162 = v156
	goto L36
L51:
	;
	goto L50
L52:
	;
	return int32(0)
L53:
	;
	return base.B2i32(v171 == int32(0))
}
func F_nameeqfast(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	goto L3
L1:
	;
	return base.B2i32(v39-v40 == int32(0))
L3:
	;
	goto L4
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = l0
	v11 = l1
	v12 = int32(64)
	v13 = v9
	goto L9
L6:
	;
	v35 = l1
	v39 = int32(0)
	goto L7
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	goto L1
L8:
	;
	v35 = v30
	v39 = v32
	goto L7
L9:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v13 != v15 {
		v30 = v11
		v32 = v13
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v30 = v24
	v32 = int32(0)
	goto L8
L11:
	;
	if v15 == int32(0) {
		v30 = v11
		v32 = v13
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v20 = v12 - int32(1)
	if v20 == int32(0) {
		v30 = v11
		v32 = v13
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v23 = int32(1)
	v24 = v11 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v25 != 0 {
		v10 = v10 + v23
		v11 = v24
		v12 = v20
		v13 = v25
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
}
func F_namege(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(base.Ui32(v172^int32(-1)) >> (uint(int32(31)) % 32))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	if v5&int32(3) == int32(0) {
		v77 = v5
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v172 = v45 - v46
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v41 = v4
	v45 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L5
L12:
	;
	v41 = v36
	v45 = v38
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v36 = v30
	v38 = int32(0)
	goto L12
L15:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v4&int32(3) == int32(0) {
		v134 = v4
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v110 = v102 - v5
	goto L19
L21:
	;
	v81 = v77
	goto L30
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = int32(0)
	goto L19
L24:
	;
	goto L25
L25:
	;
	v66 = v5
	goto L26
L26:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v102 = v70
	goto L20
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v96 = v81
	goto L33
L32:
	;
	goto L31
L33:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v102 = v96
	goto L20
L35:
	;
	goto L34
L36:
	;
	v168 = F_varstr_cmp(m, v5, v110, v4, v167, v6)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v167 = v159 - v4
	goto L36
L38:
	;
	v138 = v134
	goto L47
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v118 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v167 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v123 = v4
	goto L43
L43:
	;
	v127 = v123 + int32(1)
	if v127&int32(3) == int32(0) {
		v134 = v127
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v159 = v127
	goto L37
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		v123 = v127
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v147 = int32(-2139062144)
	if (int32(16843008)-v144|v144)&v147 == v147 {
		v138 = v138 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v153 = v138
	goto L50
L49:
	;
	goto L48
L50:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 != 0 {
		v153 = v153 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v159 = v153
	goto L37
L52:
	;
	goto L51
L53:
	;
	return int32(0)
L54:
	;
	v172 = v168
	goto L1
}
func F_nameicregexne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v5&int32(3) == int32(0) {
		v34 = v5
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = F_RE_compile_and_cache(m, v7, int32(27), v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v67 = v59 - v5
	goto L3
L5:
	;
	v38 = v34
	goto L14
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v67 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v23 = v5
	goto L10
L10:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v59 = v27
	goto L4
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v53 = v38
	goto L17
L16:
	;
	goto L15
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = F_palloc(m, v67<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v78 = F_pg_mb2wchar_with_len(m, v5, v76, v67)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(0)
	v83 = F_RE_wchar_execute(m, v76, v78, v80, v80, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v76)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	return v83 ^ int32(1)
}
func F_namenetext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8&int32(3) == int32(0) {
		v37 = v8
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v71 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v70 = v62 - v8
	goto L3
L5:
	;
	v41 = v37
	goto L14
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v26 = v8
	goto L10
L10:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v62 = v30
	goto L4
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v56 = v41
	goto L17
L16:
	;
	goto L15
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v62 = v56
	goto L4
L19:
	;
	goto L18
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v102 != int32(950) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v74 = int32(4)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v76&int32(254) == int32(2) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v89 = int32(1)
	if v71&v89 != 0 {
		v101 = int32(base.Ui32(v71)>>(uint(v89)%32)) - v89
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v85 = v74
	goto L26
L25:
	;
	v85 = base.B2i32(v76 == int32(18)) << (uint(v74) % 32)
	goto L26
L26:
	;
	if v76 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v74
	goto L29
L28:
	;
	v88 = v85
	goto L29
L29:
	;
	v101 = v88
	goto L20
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v208 != v10 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	v197 = int32(1)
	if v71&v197 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v70 != v101 {
		v207 = int32(1)
		goto L31
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(233060), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(534951), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(476976), int32(1648), int32(98766))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v127 = int32(1)
	if v71&v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v127
	goto L45
L44:
	;
	v131 = int32(4)
	goto L45
L45:
	;
	v132 = v10 + v131
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v207 = base.B2i32(v194 != int32(0))
	goto L31
L47:
	;
	v194 = int32(0)
	goto L46
L48:
	;
	v168 = v163
	v169 = v164
	v170 = v165
	goto L58
L49:
	;
	if (v8|v132)&int32(3) != 0 {
		v163 = v8
		v164 = v132
		v165 = v70
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v156 = v8
	v157 = v132
	v158 = v70
	goto L51
L51:
	;
	if v158 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L52:
	;
	v140 = v8
	v141 = v132
	v142 = v70
	goto L53
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v145 != v146 {
		v163 = v140
		v164 = v141
		v165 = v142
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v156 = v151
	v157 = v149
	v158 = v153
	goto L51
L55:
	;
	v148 = int32(4)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v163 = v156
	v164 = v157
	v165 = v158
	goto L48
L58:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v173 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v194 = v173 - v174
	goto L46
L60:
	;
	v176 = int32(1)
	v181 = v170 - v176
	if v181 != 0 {
		v168 = v168 + v176
		v169 = v169 + v176
		v170 = v181
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L47
L64:
	;
	v201 = v197
	goto L66
L65:
	;
	v201 = int32(4)
	goto L66
L66:
	;
	v203 = F_varstr_cmp(m, v8, v70, v10+v201, v101, v102)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v207 = base.B2i32(v203 != int32(0))
	goto L31
L68:
	;
	F_pfree(m, v10)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	return v207
L71:
	;
	goto L70
}
func F_nameregexne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v5&int32(3) == int32(0) {
		v34 = v5
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = F_RE_compile_and_cache(m, v7, int32(19), v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v67 = v59 - v5
	goto L3
L5:
	;
	v38 = v34
	goto L14
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v67 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v23 = v5
	goto L10
L10:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v59 = v27
	goto L4
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v53 = v38
	goto L17
L16:
	;
	goto L15
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = F_palloc(m, v67<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v78 = F_pg_mb2wchar_with_len(m, v5, v76, v67)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(0)
	v83 = F_RE_wchar_execute(m, v76, v78, v80, v80, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v76)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	return v83 ^ int32(1)
}
func F_namesend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v7&int32(3) == int32(0) {
		v35 = v7
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_pq_sendtext(m, v5, v7, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v68 = v60 - v7
	goto L3
L5:
	;
	v39 = v35
	goto L14
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v68 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v24 = v7
	goto L10
L10:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v60 = v28
	goto L4
L12:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v54 = v39
	goto L17
L16:
	;
	goto L15
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v60 = v54
	goto L4
L19:
	;
	goto L18
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 << (uint(int32(2)) % 32)
	goto L21
L21:
	;
	m.G0 = v5 + int32(16)
	return v72
}
func F_newstate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	v8 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v8 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v13 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
				v94 = v13
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
				v103 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
				*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v107 == v103 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
				} else {
				}
				v111 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
				*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
				*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v117 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v120 = v119
				} else {
					v120 = v103
				}
				*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
				return v94
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v16 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					if base.Ui32(v18) <= base.Ui32(v17) {
						v36 = l0 + int32(76)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
						if base.Ui32(v38) < base.Ui32(int32(98000000)) {
							v56 = v18 << (uint(int32(1)) % 32)
							v58 = v36
							v60 = int32(1024)
							if base.Ui32(v60) <= base.Ui32(v56) {
								v63 = v60
							} else {
								v63 = v56
							}
							v67 = v63*int32(36) + int32(8)
							v69 = F_palloc_extended(m, v67, int32(2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
								if v69 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(101)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									if v77 != 0 {
										v79 = v77
									} else {
										v79 = int32(12)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v79
									return int32(0)
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+136))
									*(*int32)(unsafe.Add(mBase, uint32(v71)+136)) = v83 + v67
									*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v63
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v69
									v94 = v69 + int32(8)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
									v103 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
									v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v107 == v103 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
									} else {
									}
									v111 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
									*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
									*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
									v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v117 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
										v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v120 = v119
									} else {
										v120 = v103
									}
									*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
									return v94
								}
							}
						} else {
							v42 = v37
							v43 = v36
							*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(101)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
							if v48 != 0 {
								v50 = v48
							} else {
								v50 = int32(19)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v50
							return int32(0)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17 + int32(1)
						v94 = v16 + v17*int32(36) + int32(8)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
						v103 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
						*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v107 == v103 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
						} else {
						}
						v111 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
						*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
						*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v117 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v120 = v119
						} else {
							v120 = v103
						}
						*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
						return v94
					}
				} else {
					v29 = l0 + int32(76)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
					if base.Ui32(int32(97999999)) < base.Ui32(v32) {
						v42 = v31
						v43 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(101)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
						if v48 != 0 {
							v50 = v48
						} else {
							v50 = int32(19)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v50
						return int32(0)
					} else {
						v56 = int32(32)
						v58 = v29
						v60 = int32(1024)
						if base.Ui32(v60) <= base.Ui32(v56) {
							v63 = v60
						} else {
							v63 = v56
						}
						v67 = v63*int32(36) + int32(8)
						v69 = F_palloc_extended(m, v67, int32(2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
							if v69 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(101)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
								if v77 != 0 {
									v79 = v77
								} else {
									v79 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v79
								return int32(0)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+136))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+136)) = v83 + v67
								*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v63
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v69
								v94 = v69 + int32(8)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
								*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v107 == v103 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
								} else {
								}
								v111 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
								*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
								*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v120 = v119
								} else {
									v120 = v103
								}
								*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
								return v94
							}
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
			v94 = v13
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
			v103 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
			*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v107 == v103 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
			} else {
			}
			v111 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
			*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
			*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v117 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v120 = v119
			} else {
				v120 = v103
			}
			*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
			return v94
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if base.Ui32(v18) <= base.Ui32(v17) {
					v36 = l0 + int32(76)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
					if base.Ui32(v38) < base.Ui32(int32(98000000)) {
						v56 = v18 << (uint(int32(1)) % 32)
						v58 = v36
						v60 = int32(1024)
						if base.Ui32(v60) <= base.Ui32(v56) {
							v63 = v60
						} else {
							v63 = v56
						}
						v67 = v63*int32(36) + int32(8)
						v69 = F_palloc_extended(m, v67, int32(2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
							if v69 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(101)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
								if v77 != 0 {
									v79 = v77
								} else {
									v79 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v79
								return int32(0)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+136))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+136)) = v83 + v67
								*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v63
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v69
								v94 = v69 + int32(8)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
								*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v107 == v103 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
								} else {
								}
								v111 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
								*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
								*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v120 = v119
								} else {
									v120 = v103
								}
								*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
								return v94
							}
						}
					} else {
						v42 = v37
						v43 = v36
						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(101)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
						if v48 != 0 {
							v50 = v48
						} else {
							v50 = int32(19)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v50
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17 + int32(1)
					v94 = v16 + v17*int32(36) + int32(8)
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
					v103 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v107 == v103 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
					} else {
					}
					v111 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
					*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
					*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v117 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v120 = v119
					} else {
						v120 = v103
					}
					*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
					return v94
				}
			} else {
				v29 = l0 + int32(76)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
				if base.Ui32(int32(97999999)) < base.Ui32(v32) {
					v42 = v31
					v43 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(101)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
					if v48 != 0 {
						v50 = v48
					} else {
						v50 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v50
					return int32(0)
				} else {
					v56 = int32(32)
					v58 = v29
					v60 = int32(1024)
					if base.Ui32(v60) <= base.Ui32(v56) {
						v63 = v60
					} else {
						v63 = v56
					}
					v67 = v63*int32(36) + int32(8)
					v69 = F_palloc_extended(m, v67, int32(2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
						if v69 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(101)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
							if v77 != 0 {
								v79 = v77
							} else {
								v79 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v79
							return int32(0)
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v71)+136)) = v83 + v67
							*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v63
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v69
							v94 = v69 + int32(8)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v99 + int32(1)
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v103)
							*(*int32)(unsafe.Add(mBase, uint32(v94))) = v99
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v107 == v103 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
							} else {
							}
							v111 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v111
							*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = v111
							*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v111
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v117 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v94
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v120 = v119
							} else {
								v120 = v103
							}
							*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
							return v94
						}
					}
				}
			}
		}
	}
}
func F_norwegian_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = v10 + int32(3)
	if v8 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v133 < v136 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v10 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v64 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v26 = v10
	goto L6
L5:
	;
	v26 = v24
	goto L6
L6:
	;
	v33 = v10
	goto L8
L7:
	;
	v64 = v44
	goto L3
L8:
	;
	if v33 == v26 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v64 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v33))))
	if int32(248) < v39 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
	v33 = v56
	goto L8
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_consts[1287]))))
	if int32(base.Ui32(v48)>>(uint(v41&int32(7))%32))&v44 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 < v75 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v119 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v78 = v75
	goto L22
L21:
	;
	v78 = v76
	goto L22
L22:
	;
	v85 = v75
	goto L24
L23:
	;
	v119 = int32(1)
	goto L19
L24:
	;
	if v85 == v78 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v119 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v85))))
	if int32(248) < v93 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v95 = v93 - int32(97)
	if v95 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v95)>>(uint(int32(3))%32)))+uint32(_consts[1287]))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v110 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
	v85 = v110
	goto L24
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v122 + v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126 < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v128 = v123
	goto L36
L35:
	;
	v128 = v126
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v128
	goto L1
L37:
	;
	return v390
L38:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v308 < v312 {
		v346 = v310
		goto L85
	} else {
		goto L86
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v136
	if v133 <= v136 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	goto L38
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v133-int32(1)))))
	if v145&int32(224) != int32(96) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(1)<<(uint(v145)%32)&int32(1851426) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v158 = F_find_among_b(m, l0, int32(4153456), int32(29))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v158 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v165
	switch v158 - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	default:
		goto L38
	}
L47:
	;
	v299 = F_slice_from_s(m, l0, int32(2), int32(2127655))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L44
	} else {
		goto L83
	}
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L54
L49:
	;
	v169 = F_slice_del(m, l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	if int32(0) <= v169 {
		goto L38
	} else {
		goto L51
	}
L51:
	;
	v390 = v169
	goto L37
L52:
	;
	if v226 != 0 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v226 = v222
	goto L52
L54:
	;
	if v182 <= v183 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v222 = int32(0)
	goto L53
L56:
	;
	v226 = int32(-1)
	goto L52
L57:
	;
	goto L58
L58:
	;
	v195 = int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v182-v195))))
	if int32(122) < v200 {
		v222 = v195
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v202 = v200 - int32(98)
	if v202 < int32(0) {
		v222 = v195
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v202)>>(uint(int32(3))%32)))+uint32(_consts[1288]))))
	if int32(base.Ui32(v208)>>(uint(v202&int32(7))%32))&int32(1) == int32(0) {
		v222 = v195
		goto L53
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v182 - int32(1)
	goto L62
L62:
	;
	goto L55
L63:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = v227 + (v165 - v173)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v229 <= v231 {
		goto L38
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v293 = F_slice_del(m, l0)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L44
	} else {
		goto L81
	}
L66:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+v229-int32(1)))))
	if v237 != int32(107) {
		goto L38
	} else {
		goto L67
	}
L67:
	;
	v241 = v229 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L70
L68:
	;
	if v291 != 0 {
		goto L38
	} else {
		goto L80
	}
L69:
	;
	v291 = v288
	goto L68
L70:
	;
	if v241 <= v251 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v288 = int32(0)
	goto L69
L72:
	;
	v291 = int32(-1)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v241-int32(1)))))
	if int32(248) < v266 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241 - int32(1)
	goto L79
L76:
	;
	v268 = v266 - int32(97)
	if v268 < int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v271 = int32(1)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v268)>>(uint(int32(3))%32)))+uint32(_consts[1287]))))
	if int32(base.Ui32(v275)>>(uint(v268&int32(7))%32))&v271 != 0 {
		v288 = v271
		goto L69
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	goto L71
L80:
	;
	goto L65
L81:
	;
	if int32(0) <= v293 {
		goto L38
	} else {
		goto L82
	}
L82:
	;
	v390 = v293
	goto L37
L83:
	;
	if int32(0) <= v299 {
		goto L38
	} else {
		goto L84
	}
L84:
	;
	v390 = v299
	goto L37
L85:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v348 < v351 {
		v386 = v346
		goto L99
	} else {
		goto L100
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v312
	v317 = v308 - int32(1)
	if v312 < v317 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v327 = F_find_among_b(m, l0, int32(4154048), int32(2))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L44
	} else {
		goto L92
	}
L88:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v317))))
	if v321 == int32(116) {
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v310
	v346 = v310
	goto L85
L91:
	;
	goto L90
L92:
	;
	if v327 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v310
	v346 = v310
	goto L85
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v310
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v335 <= v310 {
		v346 = v310
		goto L85
	} else {
		goto L96
	}
L96:
	;
	v338 = v335 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v338
	v341 = F_slice_del(m, l0)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L44
	} else {
		goto L97
	}
L97:
	;
	if v341 < int32(0) {
		v390 = v341
		goto L37
	} else {
		goto L98
	}
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v346 = v345
	goto L85
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386
	v390 = int32(1)
	goto L37
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v351
	v356 = v348 - int32(1)
	if v356 <= v351 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v373 = F_find_among_b(m, l0, int32(4154096), int32(11))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L44
	} else {
		goto L106
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v346
	v386 = v346
	goto L99
L103:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358+v356))))
	if v360&int32(224) != int32(96) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	if int32(1)<<(uint(v360)%32)&int32(4718720) != 0 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	if v373 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v346
	v386 = v346
	goto L99
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v346
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v379
	v381 = F_slice_del(m, l0)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L44
	} else {
		goto L110
	}
L110:
	;
	if v381 < int32(0) {
		v390 = v381
		goto L37
	} else {
		goto L111
	}
L111:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v386 = v385
	goto L99
}
func F_norwegian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v318 < v321 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v62 < int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v17 = v10
	v19 = int32(3)
	goto L9
L8:
	;
	v62 = v47
	goto L2
L9:
	;
	if v7 <= v17 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v62 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v24 = v17 + int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v17))))
	if base.Ui32(v26) < base.Ui32(int32(192)) {
		v47 = v24
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(1)
	if v48 < v19 {
		v17 = v47
		v19 = v19 - v48
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v7 <= v24 {
		v47 = v24
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = v24
	goto L17
L17:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9+v33))))
	if int32(-65) < v36 {
		v47 = v33
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v47 = v7
	goto L14
L19:
	;
	v40 = v33 + int32(1)
	if v40 != v7 {
		v33 = v40
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = v10
	goto L25
L23:
	;
	if v184 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v184 = v156
	goto L23
L25:
	;
	if v80 <= v89 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v184 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v96 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v81))))
	if base.Ui32(v98) < base.Ui32(int32(192)) {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(248) < v155 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v102 = v89 + int32(1)
	if v102 == v80 {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v81))))
	v107 = v105 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v98) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v81))))
	v123 = v121 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v98) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v111 = v89 + int32(2)
	if v111 != v80 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v155 = v98<<(uint(int32(6))%32)&int32(1984) | v107
	v156 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v127))))
	v155 = v140&int32(63) | (v98<<(uint(int32(18))%32)&int32(1835008) | v107<<(uint(int32(12))%32) | v123<<(uint(int32(6))%32))
	v156 = int32(4)
	goto L30
L39:
	;
	v127 = v89 + int32(3)
	if v127 != v80 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v155 = v98<<(uint(int32(12))%32)&int32(61440) | v107<<(uint(int32(6))%32) | v123
	v156 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v173 = v156 + v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173
	v89 = v173
	goto L25
L44:
	;
	v160 = v155 - int32(97)
	if v160 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v160)>>(uint(int32(3))%32)))+uint32(_consts[1302]))))
	if int32(base.Ui32(v166)>>(uint(v160&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = v198
	goto L51
L49:
	;
	if v304 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v304 = v275
	goto L49
L51:
	;
	if v199 <= v208 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v304 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v215 = int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v200))))
	if base.Ui32(v217) < base.Ui32(int32(192)) {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(248) < v274 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v221 = v208 + int32(1)
	if v221 == v199 {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v200))))
	v226 = v224 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v217) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v200))))
	v242 = v240 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v217) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v230 = v208 + int32(2)
	if v230 != v199 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = v217<<(uint(int32(6))%32)&int32(1984) | v226
	v275 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v246))))
	v274 = v259&int32(63) | (v217<<(uint(int32(18))%32)&int32(1835008) | v226<<(uint(int32(12))%32) | v242<<(uint(int32(6))%32))
	v275 = int32(4)
	goto L56
L65:
	;
	v246 = v208 + int32(3)
	if v246 != v199 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v274 = v217<<(uint(int32(12))%32)&int32(61440) | v226<<(uint(int32(6))%32) | v242
	v275 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v279 = v274 - int32(97)
	if v279 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v279)>>(uint(int32(3))%32)))+uint32(_consts[1302]))))
	if int32(base.Ui32(v285)>>(uint(v279&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v293 = v275 + v208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	v208 = v293
	goto L51
L73:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = v307 + v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	if v311 < v308 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v313 = v308
	goto L76
L75:
	;
	v313 = v311
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	goto L1
L77:
	;
	return v783
L78:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	if v650 < v653 {
		goto L145
	} else {
		goto L146
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v321
	if v318 <= v321 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	goto L78
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-int32(1)))))
	if v330&int32(224) != int32(96) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	if int32(1)<<(uint(v330)%32)&int32(1851426) == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v343 = F_find_among_b(m, l0, int32(4254832), int32(29))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	return int32(0)
L85:
	;
	if v343 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v350
	switch v343 - int32(1) {
	case 0:
		goto L89
	case 1:
		goto L88
	case 2:
		goto L87
	default:
		goto L78
	}
L87:
	;
	v641 = F_slice_from_s(m, l0, int32(2), int32(2155447))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L84
	} else {
		goto L143
	}
L88:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L94
L89:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if int32(0) <= v354 {
		goto L78
	} else {
		goto L91
	}
L91:
	;
	v783 = v354
	goto L77
L92:
	;
	if v487 != 0 {
		goto L116
	} else {
		goto L117
	}
L93:
	;
	v487 = v480
	goto L92
L94:
	;
	if v375 <= v376 {
		v480 = int32(-1)
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v480 = int32(0)
	goto L93
L96:
	;
	v393 = int32(1)
	v394 = v375 - v393
	v396 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372+v394))))
	v398 = v396 & int32(255)
	if v394 == v376 {
		v453 = v398
		v454 = v393
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if int32(122) < v453 {
		goto L106
	} else {
		goto L107
	}
L98:
	;
	if int32(0) <= v396 {
		v453 = v398
		v454 = v393
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v404 = v398 & int32(63)
	v406 = v375 - int32(2)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+v406))))
	v410 = v408 << (uint(int32(6)) % 32)
	if base.B2i32(v406 != v376)&base.B2i32(base.Ui32(v408) < base.Ui32(int32(192))) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v453 = v410&int32(1984) | v404
	v454 = int32(2)
	goto L97
L101:
	;
	goto L102
L102:
	;
	v423 = v410&int32(4032) | v404
	v425 = v375 - int32(3)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+v425))))
	if base.B2i32(v425 != v376)&base.B2i32(base.Ui32(v427) < base.Ui32(int32(224))) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v453 = v427<<(uint(int32(12))%32)&int32(61440) | v423
	v454 = int32(3)
	goto L97
L104:
	;
	goto L105
L105:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+(v372-int32(4))))))
	v453 = v427<<(uint(int32(12))%32)&int32(258048) | v445&int32(7)<<(uint(int32(18))%32) | v423
	v454 = int32(4)
	goto L97
L106:
	;
	v487 = v454
	goto L92
L107:
	;
	goto L108
L108:
	;
	v458 = v453 - int32(98)
	if v458 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v487 = v454
	goto L92
L110:
	;
	goto L111
L111:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v458)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v464)>>(uint(v458&int32(7))%32))&int32(1) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v487 = v454
	goto L92
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v375 - v454
	goto L115
L115:
	;
	goto L95
L116:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v490 = v488 + (v350 - v358)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v490 <= v492 {
		goto L78
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v635 = F_slice_del(m, l0)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L84
	} else {
		goto L141
	}
L119:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494+v490-int32(1)))))
	if v498 != int32(107) {
		goto L78
	} else {
		goto L120
	}
L120:
	;
	v502 = v490 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L123
L121:
	;
	if v633 != 0 {
		goto L78
	} else {
		goto L140
	}
L122:
	;
	v633 = v626
	goto L121
L123:
	;
	if v502 <= v521 {
		v626 = int32(-1)
		goto L122
	} else {
		goto L125
	}
L124:
	;
	v626 = int32(0)
	goto L122
L125:
	;
	v538 = int32(1)
	v539 = v502 - v538
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v517+v539))))
	v543 = v541 & int32(255)
	if v539 == v521 {
		v598 = v543
		v599 = v538
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if int32(248) < v598 {
		goto L135
	} else {
		goto L136
	}
L127:
	;
	if int32(0) <= v541 {
		v598 = v543
		v599 = v538
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v549 = v543 & int32(63)
	v551 = v502 - int32(2)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517+v551))))
	v555 = v553 << (uint(int32(6)) % 32)
	if base.B2i32(v551 != v521)&base.B2i32(base.Ui32(v553) < base.Ui32(int32(192))) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v598 = v555&int32(1984) | v549
	v599 = int32(2)
	goto L126
L130:
	;
	goto L131
L131:
	;
	v568 = v555&int32(4032) | v549
	v570 = v502 - int32(3)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517+v570))))
	if base.B2i32(v570 != v521)&base.B2i32(base.Ui32(v572) < base.Ui32(int32(224))) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v598 = v572<<(uint(int32(12))%32)&int32(61440) | v568
	v599 = int32(3)
	goto L126
L133:
	;
	goto L134
L134:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+(v517-int32(4))))))
	v598 = v572<<(uint(int32(12))%32)&int32(258048) | v590&int32(7)<<(uint(int32(18))%32) | v568
	v599 = int32(4)
	goto L126
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502 - v599
	goto L139
L136:
	;
	v603 = v598 - int32(97)
	if v603 < int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v603)>>(uint(int32(3))%32)))+uint32(_consts[1302]))))
	if int32(base.Ui32(v609)>>(uint(v603&int32(7))%32))&int32(1) == int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v633 = v599
	goto L121
L139:
	;
	goto L124
L140:
	;
	goto L118
L141:
	;
	if int32(0) <= v635 {
		goto L78
	} else {
		goto L142
	}
L142:
	;
	v783 = v635
	goto L77
L143:
	;
	if int32(0) <= v641 {
		goto L78
	} else {
		goto L144
	}
L144:
	;
	v783 = v641
	goto L77
L145:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v740
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	if v740 < v744 {
		v780 = v742
		goto L175
	} else {
		goto L176
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v650
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v653
	v659 = v650 - int32(1)
	if v659 <= v653 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	goto L145
L148:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661+v659))))
	if v663 != int32(116) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v668 = F_find_among_b(m, l0, int32(4255424), int32(2))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L84
	} else {
		goto L150
	}
L150:
	;
	if v668 == int32(0) {
		goto L147
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L154
L152:
	;
	if v728 < int32(0) {
		goto L145
	} else {
		goto L172
	}
L154:
	;
	goto L155
L155:
	;
	goto L156
L156:
	;
	v684 = v675
	v686 = int32(1)
	goto L159
L158:
	;
	v728 = v710
	goto L152
L159:
	;
	if v684 <= v656 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L158
L161:
	;
	v728 = int32(-1)
	goto L152
L162:
	;
	goto L163
L163:
	;
	v691 = v684 - int32(1)
	v693 = int32(*(*int8)(unsafe.Add(mBase, uint32(v677+v691))))
	if int32(0) <= v693 {
		v710 = v691
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v714 = int32(1)
	if v714 < v686 {
		v684 = v710
		v686 = v686 - v714
		goto L159
	} else {
		goto L171
	}
L165:
	;
	if v691 <= v656 {
		v710 = v691
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v698 = v691
	goto L167
L167:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v698))))
	if base.Ui32(int32(191)) < base.Ui32(v703) {
		v710 = v698
		goto L164
	} else {
		goto L169
	}
L168:
	;
	v710 = v656
	goto L164
L169:
	;
	v707 = v698 - int32(1)
	if v656 < v707 {
		v698 = v707
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	goto L160
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v733 = F_slice_del(m, l0)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L84
	} else {
		goto L173
	}
L173:
	;
	if int32(0) <= v733 {
		goto L145
	} else {
		goto L174
	}
L174:
	;
	v783 = v733
	goto L77
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v780
	v783 = int32(1)
	goto L77
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v740
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v744
	v749 = v740 - int32(1)
	if v749 <= v744 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v766 = F_find_among_b(m, l0, int32(4255472), int32(11))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L84
	} else {
		goto L182
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v742
	v780 = v742
	goto L175
L179:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v749))))
	if v753&int32(224) != int32(96) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	if int32(1)<<(uint(v753)%32)&int32(4718720) != 0 {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	if v766 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v742
	v780 = v742
	goto L175
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v742
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v772
	v774 = F_slice_del(m, l0)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L84
	} else {
		goto L186
	}
L186:
	;
	if v774 < int32(0) {
		v783 = v774
		goto L77
	} else {
		goto L187
	}
L187:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v780 = v778
	goto L175
}
func F_numerictypmodin(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = F_ArrayGetIntegerTypmods(m, v10, v7+int32(44))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
			switch v18 - int32(1) {
			case 0:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if base.Ui32(v40-int32(1001)) <= base.Ui32(int32(-1001)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(1000)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v104
							F_errmsg(m, int32(456278), v7+int32(32))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476879), int32(1353), int32(263790))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v52 = v40<<(uint(int32(16))%32) | int32(4)
					m.G0 = v7 + int32(48)
					return v52
				}
			case 1:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if base.Ui32(v21-int32(1001)) <= base.Ui32(int32(-1001)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(1000)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
							F_errmsg(m, int32(456278), v7)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476879), int32(1339), int32(263790))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v27 = v16 + int32(4)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					if base.Ui32(v28-int32(1001)) <= base.Ui32(int32(-2002)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = int64(4299262262296)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v83
								F_errmsg(m, int32(456235), v7+int32(16))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(476879), int32(1344), int32(263790))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v52 = v28&int32(2047) | v21<<(uint(int32(16))%32) + int32(4)
						m.G0 = v7 + int32(48)
						return v52
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(211786), int32(0))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476879), int32(1361), int32(263790))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
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
func F_numericvar_to_double_no_overflow(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_get_str_from_var(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return float64(0)
	} else {
		v14 = F_strtod(m, v8, v6+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return float64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v17 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return float64(0)
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v8
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(259295)
						F_errmsg(m, int32(683524), v6)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(476879), int32(8476), int32(28968))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_pfree(m, v8)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return float64(0)
				} else {
					m.G0 = v6 + int32(16)
					return v14
				}
			}
		}
	}
}
func F_numrange_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_DirectFunctionCall2Coll(m, int32(18), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(17), v3, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
			v15 = F_Float8GetDatum(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
