package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_cmp(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v112 = int32(_a_F_tsvector_cmp_0)
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
	v144 = int32(_a_F_tsvector_cmp_1)
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
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_s(v220)
L74:
	;
	goto L73
}
func F_tsvector_ge(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v112 = int32(_a_F_tsvector_ge_0)
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
	v144 = int32(_a_F_tsvector_ge_1)
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
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v220))
L74:
	;
	goto L73
}
func F_tsvector_lt(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v112 = int32(_a_F_tsvector_lt_0)
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
	v144 = int32(_a_F_tsvector_lt_1)
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
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(int32(base.Ui32(v220) >> (uint(int32(31)) % 32)))
L74:
	;
	goto L73
}
func F_tsvector_update_trigger(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L20
	} else {
		goto L167
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L20
	} else {
		goto L163
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L20
	} else {
		goto L159
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L20
	} else {
		goto L155
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L20
	} else {
		goto L151
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L20
	} else {
		goto L147
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L20
	} else {
		goto L143
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L20
	} else {
		goto L139
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L20
	} else {
		goto L136
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L20
	} else {
		goto L133
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L20
	} else {
		goto L130
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L20
	} else {
		goto L127
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(448) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v25&int32(4) == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	if v25&int32(24) != int32(8) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	switch v25 & int32(3) {
	case 0:
		v55 = int32(12)
		v56 = int32(1)
		goto L17
	default:
		goto L19
	case 2:
		goto L18
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v58 <= int32(2) {
		goto L9
	} else {
		goto L24
	}
L18:
	;
	v55 = int32(16)
	v56 = int32(0)
	goto L17
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int64(0)
L21:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2783), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55+v19)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v67 < v69 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v109
	if v109 == int32(-9) {
		goto L8
	} else {
		goto L40
	}
L26:
	;
	v109 = v75 + int32(1)
	goto L25
L27:
	;
	v74 = v69
	v75 = v67
	goto L30
L28:
	;
	goto L29
L29:
	;
	v98 = F_SystemAttributeByName(m, v66)
	mBase = m.M
	if v98 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v81 = v64 + v74<<(uint(int32(3))%32) + v75*int32(100)
	v84 = F_namestrcmp(m, v81+int32(32), v66)
	mBase = m.M
	if v84 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
	if v87 != int32(1) {
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v91 = v75 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v91 < v92 {
		v74 = v92
		v75 = v91
		goto L30
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	goto L31
L37:
	;
	v109 = int32(-9)
	goto L25
L38:
	;
	goto L39
L39:
	;
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+74)))
	v109 = v102
	goto L25
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v114 = F_SPI_gettypeid(m, v113, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	v117 = F_IsBinaryCoercible(m, v114, int32(3614))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	if v117 == int32(0) {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if l1 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+144)) = int64(32)
	v207 = F_palloc_mul(m, int32(16), int32(32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L73
	}
L45:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v124 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v124 < v126 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	goto L47
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v189 = F_stringToQualifiedNameList(m, v187, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L20
	} else {
		goto L69
	}
L48:
	;
	if v166 == int32(-9) {
		goto L6
	} else {
		goto L63
	}
L49:
	;
	v166 = v132 + int32(1)
	goto L48
L50:
	;
	v131 = v126
	v132 = v124
	goto L53
L51:
	;
	goto L52
L52:
	;
	v155 = F_SystemAttributeByName(m, v123)
	mBase = m.M
	if v155 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v138 = v121 + v131<<(uint(int32(3))%32) + v132*int32(100)
	v141 = F_namestrcmp(m, v138+int32(32), v123)
	mBase = m.M
	if v141 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+119)))
	if v144 != int32(1) {
		goto L49
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v148 = v132 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v148 < v149 {
		v131 = v149
		v132 = v148
		goto L53
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L54
L60:
	;
	v166 = int32(-9)
	goto L48
L61:
	;
	goto L62
L62:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+74)))
	v166 = v159
	goto L48
L63:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v170 = F_SPI_gettypeid(m, v169, v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	v173 = F_IsBinaryCoercible(m, v170, int32(3734))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	if v173 == int32(0) {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v180 = F_SPI_getbinval(m, v62, v177, v166, v17+int32(127))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+127)))
	if v182 == int32(1) {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v200 = base.I32_wrap_i64(v180)
	goto L44
L69:
	;
	if v189 == int32(0) {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v193 <= int32(1) {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v197 = F_get_ts_config_oid(m, v189, int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	v200 = v197
	goto L44
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v207
	v210 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if int32(3) <= v210 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v215 = int32(2)
	v221 = v56
	goto L77
L75:
	;
	v359 = v56
	goto L76
L76:
	;
	if v359&int32(1) != 0 {
		goto L121
	} else {
		goto L122
	}
L77:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229+v215<<(uint(int32(2))%32))))
	v234 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v234 < v236 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v359 = v347
	goto L76
L79:
	;
	if v276 == int32(-9) {
		goto L2
	} else {
		goto L94
	}
L80:
	;
	v276 = v242 + int32(1)
	goto L79
L81:
	;
	v241 = v236
	v242 = v234
	goto L84
L82:
	;
	goto L83
L83:
	;
	v265 = F_SystemAttributeByName(m, v233)
	mBase = m.M
	if v265 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v248 = v228 + v241<<(uint(int32(3))%32) + v242*int32(100)
	v251 = F_namestrcmp(m, v248+int32(32), v233)
	mBase = m.M
	if v251 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L83
L86:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+119)))
	if v254 != int32(1) {
		goto L80
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v258 = v242 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v258 < v259 {
		v241 = v259
		v242 = v258
		goto L84
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	goto L85
L91:
	;
	v276 = int32(-9)
	goto L79
L92:
	;
	goto L93
L93:
	;
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v265)+74)))
	v276 = v269
	goto L79
L94:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v280 = F_SPI_gettypeid(m, v279, v276)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v283 = F_IsBinaryCoercible(m, v280, int32(25))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L20
	} else {
		goto L96
	}
L96:
	;
	if v283 == int32(0) {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v290 = F_bms_is_member(m, v276+int32(7), v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v295 = F_SPI_getbinval(m, v62, v292, v276, v17+int32(127))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v295
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+127)))
	if v298 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v347 = v221 | v290
	v349 = v215 + int32(1)
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v349 < v350 {
		v215 = v349
		v221 = v347
		goto L77
	} else {
		goto L120
	}
L101:
	;
	v301 = base.I32_wrap_i64(v295)
	v302 = F_pg_detoast_datum_packed(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	v304 = int32(1)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v308 = v306 & v304
	if v308 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v309 = v304
	goto L105
L104:
	;
	v309 = int32(4)
	goto L105
L105:
	;
	if v306 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	F_parsetext(m, v200, v17+int32(140), v302+v309, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L20
	} else {
		goto L117
	}
L107:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	if v316 == int32(18) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v327 = int32(1)
	if v308 != 0 {
		v337 = int32(base.Ui32(v306)>>(uint(v327)%32)) - v327
		goto L106
	} else {
		goto L116
	}
L110:
	;
	v319 = int32(16)
	goto L112
L111:
	;
	v319 = int32(0)
	goto L112
L112:
	;
	if base.Ui32((v316-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v326 = int32(4)
	goto L115
L114:
	;
	v326 = v319
	goto L115
L115:
	;
	v337 = v326
	goto L106
L116:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v337 = int32(base.Ui32(v331)>>(uint(int32(2))%32)) - int32(4)
	goto L106
L117:
	;
	if v302 == v301 {
		goto L100
	} else {
		goto L118
	}
L118:
	;
	F_pfree(m, v302)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	goto L100
L120:
	;
	goto L78
L121:
	;
	v370 = F_make_tsvector(m, v17+int32(140))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L20
	} else {
		goto L124
	}
L122:
	;
	v390 = v62
	goto L123
L123:
	;
	m.G0 = v17 + int32(160)
	return base.I64_extend_i32_u(v390)
L124:
	;
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+127)) = uint8(v372)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = base.I64_extend_i32_u(v370)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v384 = F_heap_modify_tuple_by_cols(m, v62, v376, int32(1), v17+int32(156), v17+int32(128), v17+int32(127))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v17)+128))
	F_pfree(m, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	v390 = v384
	goto L123
L127:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_3), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L20
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2764), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_4), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L20
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2768), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L20
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_5), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2770), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_6), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L20
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2789), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L20
	} else {
		goto L140
	}
L140:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v456
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_7), v17)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L20
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2797), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L20
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L20
	} else {
		goto L144
	}
L144:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v474
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_8), v17+int32(112))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L20
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2804), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L20
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L148
	}
L148:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v494
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_9), v17+int32(16))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L20
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2816), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L20
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L20
	} else {
		goto L152
	}
L152:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v514
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_10), v17+int32(80))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L20
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2822), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L20
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v534
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_11), v17+int32(32))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L20
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2829), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L20
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L20
	} else {
		goto L160
	}
L160:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v554
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_12), v17+int32(96))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L20
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2842), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L20
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v573+v215<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v577
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_13), v17+int32(48))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2862), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L20
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L20
	} else {
		goto L168
	}
L168:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v596+v215<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v600
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_14), v17-int32(-64))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L20
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2867), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L20
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_update_trigger_bycolumn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tsvector_update_trigger(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
