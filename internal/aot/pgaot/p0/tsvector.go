package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_delete_str(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum_packed(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(1)
	v30 = v27 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v33 = v31 & v29
	if v31 == v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v61 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v39 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = int32(1)
	if v33 != 0 {
		v60 = int32(base.Ui32(v31)>>(uint(v50)%32)) - v50
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v42 = int32(16)
	goto L10
L9:
	;
	v42 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = int32(4)
	goto L13
L12:
	;
	v49 = v42
	goto L13
L13:
	;
	v60 = v49
	goto L4
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	m.G0 = v19 + int32(16)
	return v210
L16:
	;
	v210 = v22
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v33 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = v30
	goto L21
L20:
	;
	v66 = v27 + int32(4)
	goto L21
L21:
	;
	v70 = v22 + int32(8)
	v81 = v61
	v84 = int32(0)
	goto L23
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v92
	if base.Ui32(v90+int32(3)) < base.Ui32(int32(2)) {
		goto L63
	} else {
		goto L64
	}
L23:
	;
	v90 = v81 + v84
	v91 = int32(2)
	v92 = base.I32_div_s(v90, v91)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v70+v92<<(uint(v91)%32))))
	v100 = int32(base.Ui32(v96)>>(uint(int32(1))%32)) & int32(2047)
	if v60 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v210 = v22
	goto L15
L25:
	;
	if v188 < v187 {
		v81 = v187
		v84 = v188
		goto L23
	} else {
		goto L62
	}
L26:
	;
	v187 = v81
	v188 = v92 + int32(1)
	goto L25
L27:
	;
	if v180 == int32(0) {
		goto L22
	} else {
		goto L61
	}
L28:
	;
	if int32(0) <= v177 {
		v180 = v177
		goto L27
	} else {
		goto L60
	}
L29:
	;
	if v100 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v100 == int32(0) {
		v180 = base.B2i32(int32(0) < v60)
		goto L27
	} else {
		goto L35
	}
L32:
	;
	v105 = int32(-1)
	goto L34
L33:
	;
	v105 = int32(0)
	goto L34
L34:
	;
	v177 = v105
	goto L28
L35:
	;
	v110 = v70 + v61<<(uint(int32(2))%32) + int32(base.Ui32(v96)>>(uint(int32(12))%32))
	if base.Ui32(v60) < base.Ui32(v100) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = v60
	goto L38
L37:
	;
	v112 = v100
	goto L38
L38:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v112) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v174 != 0 {
		v177 = v174
		goto L28
	} else {
		goto L57
	}
L40:
	;
	v174 = int32(0)
	goto L39
L41:
	;
	v148 = v143
	v149 = v144
	v150 = v145
	goto L51
L42:
	;
	if (v66|v110)&int32(3) != 0 {
		v143 = v66
		v144 = v110
		v145 = v112
		goto L41
	} else {
		goto L45
	}
L43:
	;
	v136 = v66
	v137 = v110
	v138 = v112
	goto L44
L44:
	;
	if v138 == int32(0) {
		goto L40
	} else {
		goto L50
	}
L45:
	;
	v120 = v66
	v121 = v110
	v122 = v112
	goto L46
L46:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != v126 {
		v143 = v120
		v144 = v121
		v145 = v122
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v136 = v131
	v137 = v129
	v138 = v133
	goto L44
L48:
	;
	v128 = int32(4)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v143 = v136
	v144 = v137
	v145 = v138
	goto L41
L51:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 == v154 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v174 = v153 - v154
	goto L39
L53:
	;
	v156 = int32(1)
	v161 = v150 - v156
	if v161 != 0 {
		v148 = v148 + v156
		v149 = v149 + v156
		v150 = v161
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L40
L57:
	;
	if v60 == v100 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	if v100 <= v60 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v187 = v92
	v188 = v84
	goto L25
L60:
	;
	v187 = v92
	v188 = v84
	goto L25
L61:
	;
	goto L26
L62:
	;
	goto L24
L63:
	;
	v210 = v22
	goto L15
L64:
	;
	goto L65
L65:
	;
	v199 = F_tsvector_delete_by_indices(m, v22, v19+int32(12), int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v201 != v22 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_pfree(m, v22)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 == v205 {
		v210 = v199
		goto L15
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v27)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v210 = v199
	goto L15
}
func F_tsvector_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v30 = int32(base.Ui32(v28) >> (uint(v26) % 32))
	if base.Ui32(v27) < base.Ui32(v30) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 != v6 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v220 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if base.Ui32(v30) < base.Ui32(v27) {
		v195 = v33
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v220 = v195
	goto L4
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v35 < v36 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v220 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v36 < v35 {
		v195 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v220 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v43 = int32(8)
	v44 = v11 + v43
	v45 = int32(2)
	v47 = v44 + v36<<(uint(v45)%32)
	v49 = v6 + v43
	v52 = v49 + v35<<(uint(v45)%32)
	v61 = v44
	v62 = v49
	v66 = int32(0)
	goto L17
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = int32(1)
	v69 = v67 & v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v72 = v70 & v68
	if v69 != v72 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v195 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v72) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	v80 = int32(2047)
	v81 = int32(base.Ui32(v70)>>(uint(v78)%32)) & v80
	v82 = int32(12)
	v83 = int32(base.Ui32(v70) >> (uint(v82) % 32))
	v85 = int32(base.Ui32(v67) >> (uint(v82) % 32))
	v89 = int32(base.Ui32(v67)>>(uint(v78)%32)) & v80
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v77 = int32(-1)
	goto L24
L23:
	;
	v77 = int32(1)
	goto L24
L24:
	;
	v220 = v77
	goto L4
L25:
	;
	if v69 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v81 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v81 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v220 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v81))
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = v89
	goto L34
L33:
	;
	v96 = v81
	goto L34
L34:
	;
	v97 = F_memcmp(m, v85+v52, v83+v47, v96)
	mBase = m.M
	if v97 != 0 {
		v195 = v97
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v81 == v89 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(-1)
	goto L39
L38:
	;
	v101 = int32(1)
	goto L39
L39:
	;
	v220 = v101
	goto L4
L40:
	;
	v220 = int32(-1)
	goto L4
L41:
	;
	v184 = int32(4)
	v190 = v66 + int32(1)
	if v190 != v35 {
		v61 = v61 + v184
		v62 = v62 + v184
		v66 = v190
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v110 = int32(1)
	v112 = int32(_a_F_tsvector_ne_0)
	v114 = v47 + (v81+v83+v110)&v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v121 = v52 + (v89+v85+v110)&v112
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if v115 == v122 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v129 = v121
	v130 = v114
	v131 = int32(0)
	goto L51
L44:
	;
	if v122 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v115) < base.Ui32(v122) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v128 = int32(-1)
	goto L50
L49:
	;
	v128 = int32(1)
	goto L50
L50:
	;
	v220 = v128
	goto L4
L51:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v144 = int32(_a_F_tsvector_ne_1)
	v145 = v143 & v144
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v148 = v146 & v144
	if v145 != v148 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v157) < base.Ui32(v155) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v148) < base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v154 = int32(14)
	v155 = int32(base.Ui32(v143) >> (uint(v154) % 32))
	v157 = int32(base.Ui32(v146) >> (uint(v154) % 32))
	if v155 == v157 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v153 = int32(-1)
	goto L58
L57:
	;
	v153 = int32(1)
	goto L58
L58:
	;
	v220 = v153
	goto L4
L59:
	;
	v159 = int32(2)
	v164 = v131 + int32(1)
	if v164 == v122 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v129 = v129 + v159
	v130 = v130 + v159
	v131 = v164
	goto L51
L63:
	;
	v169 = int32(-1)
	goto L65
L64:
	;
	v169 = int32(1)
	goto L65
L65:
	;
	v195 = v169
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v6)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v225 != v11 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v220 != int32(0))
L74:
	;
	goto L73
}
