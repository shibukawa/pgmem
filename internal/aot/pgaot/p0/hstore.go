package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstore_cmp(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
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
	v17 = F_hstoreUpgrade(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v20 = int32(268435455)
	v21 = v19 & v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v24 = v22 & v20
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v197 != v12 {
		goto L65
	} else {
		goto L66
	}
L5:
	;
	v26 = v24
	goto L7
L6:
	;
	v26 = int32(0)
	goto L7
L7:
	;
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v34 = int32(8)
	v35 = v12 + v34
	v36 = int32(3)
	v38 = v35 + v24<<(uint(v36)%32)
	v40 = v17 + v34
	v43 = v40 + v21<<(uint(v36)%32)
	v44 = int32(4)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38-v44)))
	v47 = int32(1073741823)
	v48 = v46 & v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43-v44)))
	v53 = v51 & v47
	if base.Ui32(v48) < base.Ui32(v53) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v32 = int32(-1)
	goto L13
L12:
	;
	v32 = int32(0)
	goto L13
L13:
	;
	if v24 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v33 = int32(1)
	goto L16
L15:
	;
	v33 = v32
	goto L16
L16:
	;
	v188 = v33
	goto L4
L17:
	;
	v55 = v48
	goto L19
L18:
	;
	v55 = v53
	goto L19
L19:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v55) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v117 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v117 = int32(0)
	goto L20
L22:
	;
	v91 = v86
	v92 = v87
	v93 = v88
	goto L32
L23:
	;
	if (v38|v43)&int32(3) != 0 {
		v86 = v38
		v87 = v43
		v88 = v55
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v79 = v38
	v80 = v43
	v81 = v55
	goto L25
L25:
	;
	if v81 == int32(0) {
		goto L21
	} else {
		goto L31
	}
L26:
	;
	v63 = v38
	v64 = v43
	v65 = v55
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v68 != v69 {
		v86 = v63
		v87 = v64
		v88 = v65
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v79 = v74
	v80 = v72
	v81 = v76
	goto L25
L29:
	;
	v71 = int32(4)
	v72 = v64 + v71
	v74 = v63 + v71
	v76 = v65 - v71
	if base.Ui32(int32(3)) < base.Ui32(v76) {
		v63 = v74
		v64 = v72
		v65 = v76
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v86 = v79
	v87 = v80
	v88 = v81
	goto L22
L32:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 == v97 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = v96 - v97
	goto L20
L34:
	;
	v99 = int32(1)
	v104 = v93 - v99
	if v104 != 0 {
		v91 = v91 + v99
		v92 = v92 + v99
		v93 = v104
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
	if v48 != v53 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if v117 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	if base.Ui32(v53) < base.Ui32(v48) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v124 = int32(1)
	goto L46
L45:
	;
	v124 = int32(-1)
	goto L46
L46:
	;
	v188 = v124
	goto L4
L47:
	;
	v188 = int32(1)
	goto L4
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v24) < base.Ui32(v21) {
		v188 = int32(-1)
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v134 = int32(0)
	goto L51
L51:
	;
	v143 = v134 << (uint(int32(2)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v40+v143)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143+v35)))
	if (v145^v147)&int32(2147483647) != 0 {
		v165 = v145
		v166 = v147
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v168 = int32(1073741823)
	v169 = v166 & v168
	v171 = v165 & v168
	if base.Ui32(v169) < base.Ui32(v171) {
		v188 = int32(-1)
		goto L4
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	v152 = v143 | int32(4)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v40+v152)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v35)))
	if (v154^v156)&int32(2147483647) != 0 {
		v165 = v154
		v166 = v156
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v162 = v134 + int32(2)
	if v24<<(uint(int32(1))%32) != v162 {
		v134 = v162
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v188 = int32(0)
	goto L4
L57:
	;
	if v166&int32(1073741824) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v188 = int32(1)
	goto L4
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(v171) < base.Ui32(v169) {
		v188 = int32(1)
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v188 = v165 << (uint(int32(1)) % 32) >> (uint(int32(31)) % 32)
	goto L4
L62:
	;
	v186 = int32(-1)
	goto L64
L63:
	;
	v186 = int32(1)
	goto L64
L64:
	;
	v188 = v186
	goto L4
L65:
	;
	F_pfree(m, v12)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v201 != v17 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	F_pfree(m, v17)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	return v188
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v54 = v52 & int32(268435455)
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v29 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v40 = int32(1)
	if v23 != 0 {
		v50 = int32(base.Ui32(v21)>>(uint(v40)%32)) - v40
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v32 = int32(16)
	goto L10
L9:
	;
	v32 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(4)
	goto L13
L12:
	;
	v39 = v32
	goto L13
L13:
	;
	v50 = v39
	goto L4
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v56 = v12 + int32(8)
	if v23 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v197 = int32(0)
	goto L17
L17:
	;
	return v197
L18:
	;
	v62 = v20
	goto L20
L19:
	;
	v62 = v17 + int32(4)
	goto L20
L20:
	;
	v63 = int32(0)
	v69 = v54
	goto L22
L21:
	;
	v197 = int32(base.Ui32(v179^int32(-1)) >> (uint(int32(31)) % 32))
	goto L17
L22:
	;
	v76 = int32(base.Ui32(v69-v63)>>(uint(int32(1))%32)) + v63
	v79 = v56 + v76<<(uint(int32(3))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v82 = v80 & int32(1073741823)
	if int32(0) <= v80 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v179 = int32(-1)
	goto L21
L24:
	;
	v172 = base.B2i32(v167 < int32(0))
	if v167 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v102 = v101 + (v56 + v54<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	if base.Ui32(v50) < base.Ui32(v94) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79-int32(4))))
	v89 = v87 & int32(1073741823)
	v90 = v82 - v89
	if v90 != v50 {
		v94 = v90
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v82 == v50 {
		v101 = int32(0)
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v101 = v89
	goto L25
L31:
	;
	v94 = v82
	goto L26
L32:
	;
	v99 = int32(1)
	goto L34
L33:
	;
	v99 = int32(-1)
	goto L34
L34:
	;
	v167 = v99
	goto L24
L35:
	;
	if v164 == int32(0) {
		v179 = v76
		goto L21
	} else {
		goto L53
	}
L36:
	;
	v164 = int32(0)
	goto L35
L37:
	;
	v138 = v133
	v139 = v134
	v140 = v135
	goto L47
L38:
	;
	if (v102|v62)&int32(3) != 0 {
		v133 = v102
		v134 = v62
		v135 = v50
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v126 = v102
	v127 = v62
	v128 = v50
	goto L40
L40:
	;
	if v128 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v110 = v102
	v111 = v62
	v112 = v50
	goto L42
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v115 != v116 {
		v133 = v110
		v134 = v111
		v135 = v112
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v126 = v121
	v127 = v119
	v128 = v123
	goto L40
L44:
	;
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v110 + v118
	v123 = v112 - v118
	if base.Ui32(int32(3)) < base.Ui32(v123) {
		v110 = v121
		v111 = v119
		v112 = v123
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v133 = v126
	v134 = v127
	v135 = v128
	goto L37
L47:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v164 = v143 - v144
	goto L35
L49:
	;
	v146 = int32(1)
	v151 = v140 - v146
	if v151 != 0 {
		v138 = v138 + v146
		v139 = v139 + v146
		v140 = v151
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
	v167 = v164
	goto L24
L54:
	;
	v173 = v76 + int32(1)
	goto L56
L55:
	;
	v173 = v63
	goto L56
L56:
	;
	if v167 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v174 = v69
	goto L59
L58:
	;
	v174 = v76
	goto L59
L59:
	;
	if v173 < v174 {
		v63 = v173
		v69 = v174
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L11
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L54
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L11
	} else {
		goto L50
	}
L5:
	;
	m.G0 = v14 + int32(32)
	return v207
L6:
	;
	F_deconstruct_array_builtin(m, v17, int32(25), v14+int32(24), v14+int32(20), v14+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L11
	} else {
		goto L28
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L24
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v48 == int32(2) {
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
	if v27&int32(1) == int32(0) {
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v23 = F_palloc(m, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L11
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	switch v21 {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L7
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(-9223372036854775776)
	v207 = v23
	goto L5
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_0), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(744), int32(_a_F_hstore_from_array_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_3), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(751), int32(_a_F_hstore_from_array_2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L11
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_4), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(757), int32(_a_F_hstore_from_array_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v94 = base.I32_div_s(v92, int32(2))
	if base.Ui32(int32(53687092)) <= base.Ui32(v94) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v99 = F_palloc(m, v94*int32(20))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v92+int32(1)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v105 = int32(1)
	if base.Ui32(v94) <= base.Ui32(v105) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v202 = F_hstoreUniquePairs(m, v99, v94, v14+int32(28))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L48
	}
L34:
	;
	v108 = v105
	goto L36
L35:
	;
	v108 = v94
	goto L36
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v117 = int32(0)
	goto L37
L37:
	;
	v122 = int32(1)
	v123 = v117 << (uint(v122) % 32)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v123))))
	if v125 == v122 {
		goto L3
	} else {
		goto L39
	}
L38:
	;
	goto L33
L39:
	;
	v128 = int32(1)
	v129 = v123 | v128
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v129))))
	v134 = v99 + v117*int32(20)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v109+v123<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138 + int32(4)
	if v131 == v128 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+17)) = uint8(v182)
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+16)) = uint8(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v181
	v187 = v117 + int32(1)
	if v187 != v108 {
		v117 = v187
		goto L37
	} else {
		goto L47
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v150 = int32(base.Ui32(v146)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v150) {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v156 = int32(2)
	v158 = v109 + v129<<(uint(v156)%32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v159 + v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v167 = int32(base.Ui32(v163)>>(uint(v156)%32)) - v160
	if base.Ui32(int32(1073741824)) <= base.Ui32(v167) {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v150
	v180 = int32(1)
	v181 = int32(4)
	goto L40
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v167
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v177 = int32(base.Ui32(v173)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v177) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v180 = int32(0)
	v181 = v177
	goto L40
L47:
	;
	goto L38
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v205 = F_hstorePairs(m, v99, v202, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v207 = v205
	goto L5
L50:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(53687091)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v94
	F_errmsg(m, int32(_a_F_hstore_from_array_5), v14)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(769), int32(_a_F_hstore_from_array_2))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_6), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(778), int32(_a_F_hstore_from_array_2))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_7), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(433), int32(_a_F_hstore_from_array_8))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_hstore_from_array_9), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_hstore_from_array_1), int32(413), int32(_a_F_hstore_from_array_10))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
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
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v56 = v41
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v57 != 0 {
					v98 = v2
					v101 = v2
					v102 = v56 + v98
					v104 = v102 + int32(16)
					v105 = F_palloc(m, v104)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
						*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
						v113 = v105 + int32(16)
						if v56 != 0 {
							if v25&int32(1) != 0 {
								v118 = v24
							} else {
								v118 = v19 + int32(4)
							}
							base.MemoryCopy(m, v113, v118, v56)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
						v121 = v56 + v113
						if v57 != 0 {
							v128 = v121
							v129 = v56 | int32(1073741824)
						} else {
							if v98 != 0 {
								base.MemoryCopy(m, v121, v101, v98)
							} else {
							}
							v128 = v121 + v98
							v129 = v102 & int32(1073741823)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
						return v105
					}
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v59 = F_pg_detoast_datum_packed(m, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = int32(1)
						v62 = v59 + v61
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
						v67 = v65 & v61
						if v67 != 0 {
							v68 = v62
						} else {
							v68 = v59 + int32(4)
						}
						if v65 == int32(1) {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
							if v74 == int32(18) {
								v77 = int32(16)
							} else {
								v77 = int32(0)
							}
							if base.Ui32((v74-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v84 = int32(4)
							} else {
								v84 = v77
							}
							v98 = v84
							v101 = v68
							v102 = v56 + v98
							v104 = v102 + int32(16)
							v105 = F_palloc(m, v104)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
								*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
								v113 = v105 + int32(16)
								if v56 != 0 {
									if v25&int32(1) != 0 {
										v118 = v24
									} else {
										v118 = v19 + int32(4)
									}
									base.MemoryCopy(m, v113, v118, v56)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
								v121 = v56 + v113
								if v57 != 0 {
									v128 = v121
									v129 = v56 | int32(1073741824)
								} else {
									if v98 != 0 {
										base.MemoryCopy(m, v121, v101, v98)
									} else {
									}
									v128 = v121 + v98
									v129 = v102 & int32(1073741823)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
								*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
								*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
								return v105
							}
						} else {
							if v67 != 0 {
								v85 = int32(1)
								v94 = int32(base.Ui32(v65)>>(uint(v85)%32)) - v85
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								v94 = int32(base.Ui32(v89)>>(uint(int32(2))%32)) - int32(4)
							}
							if base.Ui32(int32(1073741824)) <= base.Ui32(v94) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16777346))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_hstore_from_text_0), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_hstore_from_text_1), int32(433), int32(_a_F_hstore_from_text_2))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
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
								v98 = v94
								v101 = v68
								v102 = v56 + v98
								v104 = v102 + int32(16)
								v105 = F_palloc(m, v104)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
									*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
									v113 = v105 + int32(16)
									if v56 != 0 {
										if v25&int32(1) != 0 {
											v118 = v24
										} else {
											v118 = v19 + int32(4)
										}
										base.MemoryCopy(m, v113, v118, v56)
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
									v121 = v56 + v113
									if v57 != 0 {
										v128 = v121
										v129 = v56 | int32(1073741824)
									} else {
										if v98 != 0 {
											base.MemoryCopy(m, v121, v101, v98)
										} else {
										}
										v128 = v121 + v98
										v129 = v102 & int32(1073741823)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
									*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
									*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
									return v105
								}
							}
						}
					}
				}
			} else {
				if v25&int32(1) != 0 {
					v44 = int32(1)
					v53 = int32(base.Ui32(v25)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v53 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
				if base.Ui32(int32(1073741824)) <= base.Ui32(v53) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16777346))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_hstore_from_text_3), int32(0))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_hstore_from_text_1), int32(413), int32(_a_F_hstore_from_text_4))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
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
					v56 = v53
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v57 != 0 {
						v98 = v2
						v101 = v2
						v102 = v56 + v98
						v104 = v102 + int32(16)
						v105 = F_palloc(m, v104)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
							*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
							v113 = v105 + int32(16)
							if v56 != 0 {
								if v25&int32(1) != 0 {
									v118 = v24
								} else {
									v118 = v19 + int32(4)
								}
								base.MemoryCopy(m, v113, v118, v56)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
							v121 = v56 + v113
							if v57 != 0 {
								v128 = v121
								v129 = v56 | int32(1073741824)
							} else {
								if v98 != 0 {
									base.MemoryCopy(m, v121, v101, v98)
								} else {
								}
								v128 = v121 + v98
								v129 = v102 & int32(1073741823)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
							*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
							*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
							return v105
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v59 = F_pg_detoast_datum_packed(m, v58)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = int32(1)
							v62 = v59 + v61
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v67 = v65 & v61
							if v67 != 0 {
								v68 = v62
							} else {
								v68 = v59 + int32(4)
							}
							if v65 == int32(1) {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
								if v74 == int32(18) {
									v77 = int32(16)
								} else {
									v77 = int32(0)
								}
								if base.Ui32((v74-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v84 = int32(4)
								} else {
									v84 = v77
								}
								v98 = v84
								v101 = v68
								v102 = v56 + v98
								v104 = v102 + int32(16)
								v105 = F_palloc(m, v104)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
									*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
									v113 = v105 + int32(16)
									if v56 != 0 {
										if v25&int32(1) != 0 {
											v118 = v24
										} else {
											v118 = v19 + int32(4)
										}
										base.MemoryCopy(m, v113, v118, v56)
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
									v121 = v56 + v113
									if v57 != 0 {
										v128 = v121
										v129 = v56 | int32(1073741824)
									} else {
										if v98 != 0 {
											base.MemoryCopy(m, v121, v101, v98)
										} else {
										}
										v128 = v121 + v98
										v129 = v102 & int32(1073741823)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
									*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
									*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
									return v105
								}
							} else {
								if v67 != 0 {
									v85 = int32(1)
									v94 = int32(base.Ui32(v65)>>(uint(v85)%32)) - v85
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
									v94 = int32(base.Ui32(v89)>>(uint(int32(2))%32)) - int32(4)
								}
								if base.Ui32(int32(1073741824)) <= base.Ui32(v94) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16777346))
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_hstore_from_text_0), int32(0))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_hstore_from_text_1), int32(433), int32(_a_F_hstore_from_text_2))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
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
									v98 = v94
									v101 = v68
									v102 = v56 + v98
									v104 = v102 + int32(16)
									v105 = F_palloc(m, v104)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(-2147483647)
										*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104 << (uint(int32(2)) % 32)
										v113 = v105 + int32(16)
										if v56 != 0 {
											if v25&int32(1) != 0 {
												v118 = v24
											} else {
												v118 = v19 + int32(4)
											}
											base.MemoryCopy(m, v113, v118, v56)
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56
										v121 = v56 + v113
										if v57 != 0 {
											v128 = v121
											v129 = v56 | int32(1073741824)
										} else {
											if v98 != 0 {
												base.MemoryCopy(m, v121, v101, v98)
											} else {
											}
											v128 = v121 + v98
											v129 = v102 & int32(1073741823)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v129
										*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v56 | int32(-2147483648)
										*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v128-v113)<<(uint(int32(2))%32) - int32(-64)
										return v105
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
	return int32(_a_F_hstore_subscript_handler_0)
}
