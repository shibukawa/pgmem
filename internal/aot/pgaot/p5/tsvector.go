package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_cmp(m *base.Module, l0 int32) int32 {
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
	return v220
L74:
	;
	goto L73
}
func F_tsvector_ge(m *base.Module, l0 int32) int32 {
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
	return int32(base.Ui32(v220^int32(-1)) >> (uint(int32(31)) % 32))
L74:
	;
	goto L73
}
func F_tsvector_lt(m *base.Module, l0 int32) int32 {
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
	return int32(base.Ui32(v220) >> (uint(int32(31)) % 32))
L74:
	;
	goto L73
}
func F_tsvector_update_trigger(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L19
	} else {
		goto L166
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L19
	} else {
		goto L162
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L19
	} else {
		goto L158
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L19
	} else {
		goto L154
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L19
	} else {
		goto L150
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L19
	} else {
		goto L146
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L19
	} else {
		goto L142
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L19
	} else {
		goto L139
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L19
	} else {
		goto L136
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L19
	} else {
		goto L133
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L19
	} else {
		goto L130
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(442) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v25&int32(4) == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	if v25&int32(24) != int32(8) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	switch v25 & int32(3) {
	case 0:
		v55 = int32(12)
		v56 = int32(1)
		goto L16
	default:
		goto L18
	case 2:
		goto L17
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v58 <= int32(2) {
		goto L8
	} else {
		goto L23
	}
L17:
	;
	v55 = int32(16)
	v56 = int32(0)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2779), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55+v19)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v67 < v69 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v109
	if v109 == int32(-9) {
		goto L7
	} else {
		goto L39
	}
L25:
	;
	v109 = v75 + int32(1)
	goto L24
L26:
	;
	v74 = v69
	v75 = v67
	goto L29
L27:
	;
	goto L28
L28:
	;
	v98 = F_SystemAttributeByName(m, v66)
	mBase = m.M
	if v98 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v81 = v64 + v74<<(uint(int32(4))%32) + v75*int32(100)
	v84 = F_namestrcmp(m, v81+int32(24), v66)
	mBase = m.M
	if v84 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+111)))
	if v87 != int32(1) {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v91 = v75 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v91 < v92 {
		v74 = v92
		v75 = v91
		goto L29
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L30
L36:
	;
	v109 = int32(-9)
	goto L24
L37:
	;
	goto L38
L38:
	;
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+74)))
	v109 = v102
	goto L24
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v114 = F_SPI_gettypeid(m, v113, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v117 = F_IsBinaryCoercible(m, v114, int32(3614))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v117 == int32(0) {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = int64(32)
	v225 = F_palloc(m, int32(512))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L76
	}
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v124 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v124 < v126 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	goto L46
L46:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v208 = F_stringToQualifiedNameList(m, v206, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L19
	} else {
		goto L72
	}
L47:
	;
	if v166 == int32(-9) {
		goto L5
	} else {
		goto L62
	}
L48:
	;
	v166 = v132 + int32(1)
	goto L47
L49:
	;
	v131 = v126
	v132 = v124
	goto L52
L50:
	;
	goto L51
L51:
	;
	v155 = F_SystemAttributeByName(m, v123)
	mBase = m.M
	if v155 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	v138 = v121 + v131<<(uint(int32(4))%32) + v132*int32(100)
	v141 = F_namestrcmp(m, v138+int32(24), v123)
	mBase = m.M
	if v141 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+111)))
	if v144 != int32(1) {
		goto L48
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v148 = v132 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v148 < v149 {
		v131 = v149
		v132 = v148
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	goto L53
L59:
	;
	v166 = int32(-9)
	goto L47
L60:
	;
	goto L61
L61:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+74)))
	v166 = v159
	goto L47
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v170 = F_SPI_gettypeid(m, v169, v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	v173 = F_IsBinaryCoercible(m, v170, int32(3734))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	if v173 == int32(0) {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v180 = F_SPI_getbinval(m, v62, v177, v166, v17+int32(119))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v182 != int32(1) {
		v219 = v180
		goto L43
	} else {
		goto L67
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v193
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_3), v17+int32(32))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2825), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	if v208 == int32(0) {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v212 <= int32(1) {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v216 = F_get_ts_config_oid(m, v208, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	v219 = v216
	goto L43
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v225
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if int32(3) <= v228 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v233 = int32(2)
	v239 = v56
	goto L80
L78:
	;
	v377 = v56
	goto L79
L79:
	;
	if v377&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L80:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v233<<(uint(int32(2))%32))))
	v252 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	if v252 < v254 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v377 = v365
	goto L79
L82:
	;
	if v294 == int32(-9) {
		goto L2
	} else {
		goto L97
	}
L83:
	;
	v294 = v260 + int32(1)
	goto L82
L84:
	;
	v259 = v254
	v260 = v252
	goto L87
L85:
	;
	goto L86
L86:
	;
	v283 = F_SystemAttributeByName(m, v251)
	mBase = m.M
	if v283 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v266 = v246 + v259<<(uint(int32(4))%32) + v260*int32(100)
	v269 = F_namestrcmp(m, v266+int32(24), v251)
	mBase = m.M
	if v269 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+111)))
	if v272 != int32(1) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v276 = v260 + int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	if v276 < v277 {
		v259 = v277
		v260 = v276
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L88
L94:
	;
	v294 = int32(-9)
	goto L82
L95:
	;
	goto L96
L96:
	;
	v287 = int32(*(*int16)(unsafe.Add(mBase, uint32(v283)+74)))
	v294 = v287
	goto L82
L97:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v298 = F_SPI_gettypeid(m, v297, v294)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	v301 = F_IsBinaryCoercible(m, v298, int32(25))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	if v301 == int32(0) {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v308 = F_bms_is_member(m, v294+int32(7), v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L19
	} else {
		goto L101
	}
L101:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v313 = F_SPI_getbinval(m, v62, v310, v294, v17+int32(119))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L19
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v313
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v316 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v365 = v239 | v308
	v367 = v233 + int32(1)
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v367 < v368 {
		v233 = v367
		v239 = v365
		goto L80
	} else {
		goto L123
	}
L104:
	;
	v319 = F_pg_detoast_datum_packed(m, v313)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L19
	} else {
		goto L105
	}
L105:
	;
	v321 = int32(1)
	v322 = v319 + v321
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v327 = v325 & v321
	if v327 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v328 = v322
	goto L108
L107:
	;
	v328 = v319 + int32(4)
	goto L108
L108:
	;
	if v325 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	F_parsetext(m, v219, v17+int32(124), v328, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L19
	} else {
		goto L120
	}
L110:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v334 == int32(18) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v345 = int32(1)
	if v327 != 0 {
		v355 = int32(base.Ui32(v325)>>(uint(v345)%32)) - v345
		goto L109
	} else {
		goto L119
	}
L113:
	;
	v337 = int32(16)
	goto L115
L114:
	;
	v337 = int32(0)
	goto L115
L115:
	;
	if base.Ui32((v334-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v344 = int32(4)
	goto L118
L117:
	;
	v344 = v337
	goto L118
L118:
	;
	v355 = v344
	goto L109
L119:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v355 = int32(base.Ui32(v349)>>(uint(int32(2))%32)) - int32(4)
	goto L109
L120:
	;
	if v319 == v313 {
		goto L103
	} else {
		goto L121
	}
L121:
	;
	F_pfree(m, v319)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	goto L103
L123:
	;
	goto L81
L124:
	;
	v388 = F_make_tsvector(m, v17+int32(124))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L19
	} else {
		goto L127
	}
L125:
	;
	v407 = v62
	goto L126
L126:
	;
	m.G0 = v17 + int32(144)
	return v407
L127:
	;
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)) = uint8(v390)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v388
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v401 = F_heap_modify_tuple_by_cols(m, v62, v393, int32(1), v17+int32(140), v17+int32(120), v17+int32(119))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L19
	} else {
		goto L128
	}
L128:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v17)+120))
	F_pfree(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L19
	} else {
		goto L129
	}
L129:
	;
	v407 = v401
	goto L126
L130:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_4), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2760), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L19
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2764), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L19
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
	v445 = m.ExcPending
	if v445 != 0 {
		goto L19
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2766), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L19
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
	F_errmsg_internal(m, int32(_a_F_tsvector_update_trigger_7), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2785), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L19
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v472
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_8), v17)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2793), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L19
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v490
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_9), v17+int32(112))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2800), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L19
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L19
	} else {
		goto L151
	}
L151:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v510
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_10), v17+int32(16))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2812), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L19
	} else {
		goto L155
	}
L155:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v530
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_11), v17+int32(80))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L19
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2818), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v550
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_12), v17+int32(96))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2838), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v569+v233<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v573
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_13), v17+int32(48))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L19
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2858), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L19
	} else {
		goto L167
	}
L167:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v592+v233<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v596
	F_errmsg(m, int32(_a_F_tsvector_update_trigger_14), v17-int32(-64))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_tsvector_update_trigger_1), int32(2863), int32(_a_F_tsvector_update_trigger_2))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_update_trigger_bycolumn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tsvector_update_trigger(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
