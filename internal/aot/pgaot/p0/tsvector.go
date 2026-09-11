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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
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
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v62 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v36 = int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v38&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v51 = int32(1)
	if v33 != 0 {
		v61 = int32(base.Ui32(v31)>>(uint(v51)%32)) - v51
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v47 = v36
	goto L10
L9:
	;
	v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
	goto L10
L10:
	;
	if v38 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v36
	goto L13
L12:
	;
	v50 = v47
	goto L13
L13:
	;
	v61 = v50
	goto L4
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	m.G0 = v19 + int32(16)
	return v212
L16:
	;
	v212 = v22
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
	v67 = v30
	goto L21
L20:
	;
	v67 = v27 + int32(4)
	goto L21
L21:
	;
	v68 = int32(0)
	v71 = v22 + int32(8)
	v80 = v68
	v83 = v62
	goto L23
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v94
	if base.Ui32(v92+int32(3)) < base.Ui32(int32(2)) {
		goto L63
	} else {
		goto L64
	}
L23:
	;
	v92 = v80 + v83
	v93 = int32(2)
	v94 = base.I32_div_s(v92, v93)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v71+v94<<(uint(v93)%32))))
	v102 = int32(base.Ui32(v98)>>(uint(int32(1))%32)) & int32(2047)
	if v61 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v212 = v22
	goto L15
L25:
	;
	if v189 < v190 {
		v80 = v189
		v83 = v190
		goto L23
	} else {
		goto L62
	}
L26:
	;
	v189 = v94 + int32(1)
	v190 = v83
	goto L25
L27:
	;
	if v182 == int32(0) {
		goto L22
	} else {
		goto L61
	}
L28:
	;
	if int32(0) <= v179 {
		v182 = v179
		goto L27
	} else {
		goto L60
	}
L29:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v102 == int32(0) {
		v182 = base.B2i32(v68 < v61)
		goto L27
	} else {
		goto L35
	}
L32:
	;
	v107 = int32(-1)
	goto L34
L33:
	;
	v107 = int32(0)
	goto L34
L34:
	;
	v179 = v107
	goto L28
L35:
	;
	v112 = v71 + v62<<(uint(int32(2))%32) + int32(base.Ui32(v98)>>(uint(int32(12))%32))
	if base.Ui32(v61) < base.Ui32(v102) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v114 = v61
	goto L38
L37:
	;
	v114 = v102
	goto L38
L38:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v114) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v176 != 0 {
		v179 = v176
		goto L28
	} else {
		goto L57
	}
L40:
	;
	v176 = int32(0)
	goto L39
L41:
	;
	v150 = v145
	v151 = v146
	v152 = v147
	goto L51
L42:
	;
	if (v67|v112)&int32(3) != 0 {
		v145 = v67
		v146 = v112
		v147 = v114
		goto L41
	} else {
		goto L45
	}
L43:
	;
	v138 = v67
	v139 = v112
	v140 = v114
	goto L44
L44:
	;
	if v140 == int32(0) {
		goto L40
	} else {
		goto L50
	}
L45:
	;
	v122 = v67
	v123 = v112
	v124 = v114
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v127 != v128 {
		v145 = v122
		v146 = v123
		v147 = v124
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v138 = v133
	v139 = v131
	v140 = v135
	goto L44
L48:
	;
	v130 = int32(4)
	v131 = v123 + v130
	v133 = v122 + v130
	v135 = v124 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v122 = v133
		v123 = v131
		v124 = v135
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v145 = v138
	v146 = v139
	v147 = v140
	goto L41
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 == v156 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v176 = v155 - v156
	goto L39
L53:
	;
	v158 = int32(1)
	v163 = v152 - v158
	if v163 != 0 {
		v150 = v150 + v158
		v151 = v151 + v158
		v152 = v163
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
	if v102 == v61 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	if v102 <= v61 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v189 = v80
	v190 = v94
	goto L25
L60:
	;
	v189 = v80
	v190 = v94
	goto L25
L61:
	;
	goto L26
L62:
	;
	goto L24
L63:
	;
	v212 = v22
	goto L15
L64:
	;
	goto L65
L65:
	;
	v201 = F_tsvector_delete_by_indices(m, v22, v19+int32(12), int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v203 != v22 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_pfree(m, v22)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 == v207 {
		v212 = v201
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
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v212 = v201
	goto L15
}
func F_tsvector_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
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
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
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
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v223 != int32(0))
L74:
	;
	goto L73
}
