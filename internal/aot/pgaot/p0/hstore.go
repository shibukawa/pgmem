package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstore_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_hstoreUpgrade(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_hstoreUpgrade(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v21 = int32(268435455)
	v22 = v20 & v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v25 = v23 & v21
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v200 != v13 {
		goto L65
	} else {
		goto L66
	}
L5:
	;
	v27 = v25
	goto L7
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v35 = int32(8)
	v36 = v13 + v35
	v37 = int32(3)
	v39 = v36 + v25<<(uint(v37)%32)
	v41 = v18 + v35
	v44 = v41 + v22<<(uint(v37)%32)
	v45 = int32(4)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39-v45)))
	v48 = int32(1073741823)
	v49 = v47 & v48
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44-v45)))
	v54 = v52 & v48
	if base.Ui32(v49) < base.Ui32(v54) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v33 = int32(-1)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if v25 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v34 = int32(1)
	goto L16
L15:
	;
	v34 = v33
	goto L16
L16:
	;
	v190 = v34
	goto L4
L17:
	;
	v56 = v49
	goto L19
L18:
	;
	v56 = v54
	goto L19
L19:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v56) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v118 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v118 = int32(0)
	goto L20
L22:
	;
	v92 = v87
	v93 = v88
	v94 = v89
	goto L32
L23:
	;
	if (v39|v44)&int32(3) != 0 {
		v87 = v39
		v88 = v44
		v89 = v56
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v80 = v39
	v81 = v44
	v82 = v56
	goto L25
L25:
	;
	if v82 == int32(0) {
		goto L21
	} else {
		goto L31
	}
L26:
	;
	v64 = v39
	v65 = v44
	v66 = v56
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v69 != v70 {
		v87 = v64
		v88 = v65
		v89 = v66
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v80 = v75
	v81 = v73
	v82 = v77
	goto L25
L29:
	;
	v72 = int32(4)
	v73 = v65 + v72
	v75 = v64 + v72
	v77 = v66 - v72
	if base.Ui32(int32(3)) < base.Ui32(v77) {
		v64 = v75
		v65 = v73
		v66 = v77
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v87 = v80
	v88 = v81
	v89 = v82
	goto L22
L32:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 == v98 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v118 = v97 - v98
	goto L20
L34:
	;
	v100 = int32(1)
	v105 = v94 - v100
	if v105 != 0 {
		v92 = v92 + v100
		v93 = v93 + v100
		v94 = v105
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
	goto L21
L38:
	;
	if v49 != v54 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if v118 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	if base.Ui32(v54) < base.Ui32(v49) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v22) < base.Ui32(v25) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v125 = int32(1)
	goto L46
L45:
	;
	v125 = int32(-1)
	goto L46
L46:
	;
	v190 = v125
	goto L4
L47:
	;
	v190 = int32(1)
	goto L4
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v25) < base.Ui32(v22) {
		v190 = int32(-1)
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v136 = int32(0)
	goto L51
L51:
	;
	v145 = v136 << (uint(int32(2)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v41+v145)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v36)))
	if (v147^v149)&int32(2147483647) != 0 {
		v167 = v149
		v168 = v147
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v170 = int32(1073741823)
	v171 = v167 & v170
	v173 = v168 & v170
	if base.Ui32(v171) < base.Ui32(v173) {
		v190 = int32(-1)
		goto L4
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	v154 = v145 | int32(4)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v41+v154)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v36)))
	if (v156^v158)&int32(2147483647) != 0 {
		v167 = v158
		v168 = v156
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v164 = v136 + int32(2)
	if v25<<(uint(int32(1))%32) != v164 {
		v136 = v164
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v190 = int32(0)
	goto L4
L57:
	;
	if v167&int32(1073741824) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v190 = int32(1)
	goto L4
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(v173) < base.Ui32(v171) {
		v190 = int32(1)
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v190 = v168 << (uint(int32(1)) % 32) >> (uint(int32(31)) % 32)
	goto L4
L62:
	;
	v188 = int32(-1)
	goto L64
L63:
	;
	v188 = int32(1)
	goto L64
L64:
	;
	v190 = v188
	goto L4
L65:
	;
	F_pfree(m, v13)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v204 != v18 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	F_pfree(m, v18)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	return v190
L72:
	;
	goto L71
}
func F_hstore_exists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v197 int32
	_ = v197
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_hstoreUpgrade(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(1)
	v20 = v17 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v23 = v21 & v19
	if v21 == v19 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v55 = v53 & int32(268435455)
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v57 = v12 + int32(8)
	if v23 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v197 = v52
	goto L17
L17:
	;
	return v197
L18:
	;
	v63 = v20
	goto L20
L19:
	;
	v63 = v17 + int32(4)
	goto L20
L20:
	;
	v64 = v52
	v70 = v55
	goto L22
L21:
	;
	v197 = int32(base.Ui32(v180^int32(-1)) >> (uint(int32(31)) % 32))
	goto L17
L22:
	;
	v76 = base.I32_div_s(v70-v64, int32(2))
	v77 = v76 + v64
	v80 = v57 + v77<<(uint(int32(3))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v83 = v81 & int32(1073741823)
	if int32(0) <= v81 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v180 = int32(-1)
	goto L21
L24:
	;
	v173 = base.B2i32(v168 < int32(0))
	if v168 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v103 = v102 + (v57 + v55<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	if base.Ui32(v51) < base.Ui32(v95) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80-int32(4))))
	v90 = v88 & int32(1073741823)
	v91 = v83 - v90
	if v91 != v51 {
		v95 = v91
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v83 == v51 {
		v102 = int32(0)
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v102 = v90
	goto L25
L31:
	;
	v95 = v83
	goto L26
L32:
	;
	v100 = int32(1)
	goto L34
L33:
	;
	v100 = int32(-1)
	goto L34
L34:
	;
	v168 = v100
	goto L24
L35:
	;
	if v165 == int32(0) {
		v180 = v77
		goto L21
	} else {
		goto L53
	}
L36:
	;
	v165 = int32(0)
	goto L35
L37:
	;
	v139 = v134
	v140 = v135
	v141 = v136
	goto L47
L38:
	;
	if (v103|v63)&int32(3) != 0 {
		v134 = v103
		v135 = v63
		v136 = v51
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v127 = v103
	v128 = v63
	v129 = v51
	goto L40
L40:
	;
	if v129 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v111 = v103
	v112 = v63
	v113 = v51
	goto L42
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v116 != v117 {
		v134 = v111
		v135 = v112
		v136 = v113
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v127 = v122
	v128 = v120
	v129 = v124
	goto L40
L44:
	;
	v119 = int32(4)
	v120 = v112 + v119
	v122 = v111 + v119
	v124 = v113 - v119
	if base.Ui32(int32(3)) < base.Ui32(v124) {
		v111 = v122
		v112 = v120
		v113 = v124
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v134 = v127
	v135 = v128
	v136 = v129
	goto L37
L47:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v144 == v145 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v165 = v144 - v145
	goto L35
L49:
	;
	v147 = int32(1)
	v152 = v141 - v147
	if v152 != 0 {
		v139 = v139 + v147
		v140 = v140 + v147
		v141 = v152
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	v168 = v165
	goto L24
L54:
	;
	v174 = v77 + int32(1)
	goto L56
L55:
	;
	v174 = v64
	goto L56
L56:
	;
	if v168 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v175 = v70
	goto L59
L58:
	;
	v175 = v77
	goto L59
L59:
	;
	if v174 < v175 {
		v64 = v174
		v70 = v175
		goto L22
	} else {
		goto L60
	}
L60:
	;
	goto L23
}
func F_hstore_from_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L12
	} else {
		goto L67
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L63
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L12
	} else {
		goto L59
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L55
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L51
	}
L6:
	;
	m.G0 = v14 + int32(32)
	return v219
L7:
	;
	F_deconstruct_array_builtin(m, v17, int32(25), v14+int32(24), v14+int32(20), v14+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L29
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L25
	}
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v52 == int32(2) {
		goto L7
	} else {
		goto L20
	}
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
	if v27&int32(1) == int32(0) {
		goto L7
	} else {
		goto L15
	}
L11:
	;
	v23 = F_palloc(m, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	switch v21 {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	default:
		goto L8
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(-9223372036854775776)
	v219 = v23
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(116286), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(476979), int32(744), int32(22981))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(139695), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(476979), int32(751), int32(22981))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(111552), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(476979), int32(757), int32(22981))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v106 = base.I32_div_s(v104, int32(2))
	if base.Ui32(int32(53687092)) <= base.Ui32(v106) {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v111 = F_palloc(m, v106*int32(20))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v104+int32(1)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = int32(1)
	if base.Ui32(v106) <= base.Ui32(v117) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v214 = F_hstoreUniquePairs(m, v111, v106, v14+int32(28))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L49
	}
L35:
	;
	v120 = v117
	goto L37
L36:
	;
	v120 = v106
	goto L37
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v128 = int32(0)
	goto L38
L38:
	;
	v134 = int32(1)
	v135 = v128 << (uint(v134) % 32)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v135))))
	if v137 == v134 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	goto L34
L40:
	;
	v140 = int32(1)
	v141 = v135 | v140
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v141))))
	v146 = v111 + v128*int32(20)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v121+v135<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v150 + int32(4)
	if v143 == v140 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+17)) = uint8(v194)
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+16)) = uint8(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = v193
	v199 = v128 + int32(1)
	if v199 != v120 {
		v128 = v199
		goto L38
	} else {
		goto L48
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v162 = int32(base.Ui32(v158)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v162) {
		goto L3
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v168 = int32(2)
	v170 = v121 + v141<<(uint(v168)%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v171 + v172
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v179 = int32(base.Ui32(v175)>>(uint(v168)%32)) - v172
	if base.Ui32(int32(1073741824)) <= base.Ui32(v179) {
		goto L2
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v162
	v192 = int32(1)
	v193 = int32(4)
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v179
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v189 = int32(base.Ui32(v185)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v189) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v192 = int32(0)
	v193 = v189
	goto L41
L48:
	;
	goto L39
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v217 = F_hstorePairs(m, v111, v214, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v219 = v217
	goto L6
L51:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(53687091)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v106
	F_errmsg(m, int32(642370), v14)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(476979), int32(769), int32(22981))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(21013), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(476979), int32(778), int32(22981))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(20982), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(476979), int32(413), int32(270713))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(20982), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(476979), int32(413), int32(270713))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(332020), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(476979), int32(433), int32(270731))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_from_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	v2 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 == int32(1) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(1)
			v24 = v19 + v23
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			if v25 == v23 {
				v28 = int32(4)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				if v30&int32(254) == int32(2) {
					v39 = v28
				} else {
					v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
				}
				if v30 == int32(1) {
					v42 = v28
				} else {
					v42 = v39
				}
				v57 = v42
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v58 != 0 {
					v100 = v2
					v103 = v2
					v106 = v57 + v100 + int32(16)
					v107 = F_palloc(m, v106)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
						*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
						v115 = v107 + int32(16)
						if v25&int32(1) != 0 {
							v120 = v24
						} else {
							v120 = v19 + int32(4)
						}
						if v57 != 0 {
							v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
							mBase = m.M
							v122 = v121
						} else {
							v122 = v115
						}
						*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
						v124 = v57 + v122
						if v58 != 0 {
							v133 = v124
							v134 = v57 | int32(1073741824)
						} else {
							if v100 != 0 {
								v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
								mBase = m.M
								v128 = v127
							} else {
								v128 = v124
							}
							v129 = v128 + v100
							v133 = v129
							v134 = (v129 - v122) & int32(1073741823)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
						return v107
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v60 = F_pg_detoast_datum_packed(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(1)
						v63 = v60 + v62
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v68 = v66 & v62
						if v68 != 0 {
							v69 = v63
						} else {
							v69 = v60 + int32(4)
						}
						if v66 == int32(1) {
							v72 = int32(4)
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
							if v74&int32(254) == int32(2) {
								v83 = v72
							} else {
								v83 = base.B2i32(v74 == int32(18)) << (uint(v72) % 32)
							}
							if v74 == int32(1) {
								v86 = v72
							} else {
								v86 = v83
							}
							v100 = v86
							v103 = v69
							v106 = v57 + v100 + int32(16)
							v107 = F_palloc(m, v106)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
								v115 = v107 + int32(16)
								if v25&int32(1) != 0 {
									v120 = v24
								} else {
									v120 = v19 + int32(4)
								}
								if v57 != 0 {
									v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
									mBase = m.M
									v122 = v121
								} else {
									v122 = v115
								}
								*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
								v124 = v57 + v122
								if v58 != 0 {
									v133 = v124
									v134 = v57 | int32(1073741824)
								} else {
									if v100 != 0 {
										v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
										mBase = m.M
										v128 = v127
									} else {
										v128 = v124
									}
									v129 = v128 + v100
									v133 = v129
									v134 = (v129 - v122) & int32(1073741823)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
								return v107
							}
						} else {
							if v68 != 0 {
								v87 = int32(1)
								v96 = int32(base.Ui32(v66)>>(uint(v87)%32)) - v87
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
								v96 = int32(base.Ui32(v91)>>(uint(int32(2))%32)) - int32(4)
							}
							if base.Ui32(int32(1073741824)) <= base.Ui32(v96) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16777346))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(332020), int32(0))
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476979), int32(433), int32(270731))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
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
								v100 = v96
								v103 = v69
								v106 = v57 + v100 + int32(16)
								v107 = F_palloc(m, v106)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
									v115 = v107 + int32(16)
									if v25&int32(1) != 0 {
										v120 = v24
									} else {
										v120 = v19 + int32(4)
									}
									if v57 != 0 {
										v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
										mBase = m.M
										v122 = v121
									} else {
										v122 = v115
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
									v124 = v57 + v122
									if v58 != 0 {
										v133 = v124
										v134 = v57 | int32(1073741824)
									} else {
										if v100 != 0 {
											v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
											mBase = m.M
											v128 = v127
										} else {
											v128 = v124
										}
										v129 = v128 + v100
										v133 = v129
										v134 = (v129 - v122) & int32(1073741823)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
									*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
									return v107
								}
							}
						}
					}
				}
			} else {
				if v25&int32(1) != 0 {
					v45 = int32(1)
					v54 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v54 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
				}
				if base.Ui32(int32(1073741824)) <= base.Ui32(v54) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16777346))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(20982), int32(0))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476979), int32(413), int32(270713))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
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
					v57 = v54
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v58 != 0 {
						v100 = v2
						v103 = v2
						v106 = v57 + v100 + int32(16)
						v107 = F_palloc(m, v106)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
							*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
							v115 = v107 + int32(16)
							if v25&int32(1) != 0 {
								v120 = v24
							} else {
								v120 = v19 + int32(4)
							}
							if v57 != 0 {
								v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
								mBase = m.M
								v122 = v121
							} else {
								v122 = v115
							}
							*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
							v124 = v57 + v122
							if v58 != 0 {
								v133 = v124
								v134 = v57 | int32(1073741824)
							} else {
								if v100 != 0 {
									v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
									mBase = m.M
									v128 = v127
								} else {
									v128 = v124
								}
								v129 = v128 + v100
								v133 = v129
								v134 = (v129 - v122) & int32(1073741823)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
							*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
							return v107
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v60 = F_pg_detoast_datum_packed(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(1)
							v63 = v60 + v62
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v68 = v66 & v62
							if v68 != 0 {
								v69 = v63
							} else {
								v69 = v60 + int32(4)
							}
							if v66 == int32(1) {
								v72 = int32(4)
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
								if v74&int32(254) == int32(2) {
									v83 = v72
								} else {
									v83 = base.B2i32(v74 == int32(18)) << (uint(v72) % 32)
								}
								if v74 == int32(1) {
									v86 = v72
								} else {
									v86 = v83
								}
								v100 = v86
								v103 = v69
								v106 = v57 + v100 + int32(16)
								v107 = F_palloc(m, v106)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
									v115 = v107 + int32(16)
									if v25&int32(1) != 0 {
										v120 = v24
									} else {
										v120 = v19 + int32(4)
									}
									if v57 != 0 {
										v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
										mBase = m.M
										v122 = v121
									} else {
										v122 = v115
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
									v124 = v57 + v122
									if v58 != 0 {
										v133 = v124
										v134 = v57 | int32(1073741824)
									} else {
										if v100 != 0 {
											v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
											mBase = m.M
											v128 = v127
										} else {
											v128 = v124
										}
										v129 = v128 + v100
										v133 = v129
										v134 = (v129 - v122) & int32(1073741823)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
									*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
									return v107
								}
							} else {
								if v68 != 0 {
									v87 = int32(1)
									v96 = int32(base.Ui32(v66)>>(uint(v87)%32)) - v87
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									v96 = int32(base.Ui32(v91)>>(uint(int32(2))%32)) - int32(4)
								}
								if base.Ui32(int32(1073741824)) <= base.Ui32(v96) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16777346))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(332020), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476979), int32(433), int32(270731))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
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
									v100 = v96
									v103 = v69
									v106 = v57 + v100 + int32(16)
									v107 = F_palloc(m, v106)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(-2147483647)
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
										v115 = v107 + int32(16)
										if v25&int32(1) != 0 {
											v120 = v24
										} else {
											v120 = v19 + int32(4)
										}
										if v57 != 0 {
											v121 = F__emscripten_memcpy_bulkmem(m, v115, v120, v57)
											mBase = m.M
											v122 = v121
										} else {
											v122 = v115
										}
										*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57
										v124 = v57 + v122
										if v58 != 0 {
											v133 = v124
											v134 = v57 | int32(1073741824)
										} else {
											if v100 != 0 {
												v127 = F__emscripten_memcpy_bulkmem(m, v124, v103, v100)
												mBase = m.M
												v128 = v127
											} else {
												v128 = v124
											}
											v129 = v128 + v100
											v133 = v129
											v134 = (v129 - v122) & int32(1073741823)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v134
										*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v57 | int32(-2147483648)
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = (v133-v122)<<(uint(int32(2))%32) - int32(-64)
										return v107
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
func F_hstore_subscript_handler(m *base.Module, l0 int32) int32 {
	return int32(4343268)
}
