package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = v10 + int32(12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v43 = v40
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
	v40 = v36
	goto L2
L4:
	;
	v32 = int32(0)
	if v13 == v32 {
		v40 = v32
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L6:
	;
	if v13 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v13 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v35 = v25
	v36 = int32(1)
	goto L3
L11:
	;
	v43 = int32(2)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+368))
	v35 = v30
	v36 = int32(2)
	goto L3
L14:
	;
	v35 = v32
	v36 = v2
	goto L3
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L31
	} else {
		goto L67
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = v47
	goto L20
L19:
	;
	v48 = v2
	goto L20
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v49 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v10 + int32(16)
	return v212
L22:
	;
	if v48 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v52 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v48 != 0 {
		v212 = v48
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
	v212 = int32(0)
	goto L21
L28:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v63 = F_initArrayResultWithSize(m, v59, v60, int32(0), v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v122 <= int32(0) {
		v212 = v48
		goto L21
	} else {
		goto L46
	}
L31:
	;
	return int32(0)
L32:
	;
	v67 = int32(_a_F_array_agg_combine_0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if int32(0) < v72 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v76 = int32(0)
	goto L36
L34:
	;
	v110 = v72
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v68
	if v110 != 0 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v76))))
	if v86 == v83 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v110 = v106
	goto L35
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v76<<(uint(int32(2))%32))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+26)))
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+24)))
	v96 = F_datumCopy(m, v93, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L31
	} else {
		goto L41
	}
L39:
	;
	v98 = v83
	goto L40
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v76<<(uint(int32(2))%32)))) = v98
	v105 = v76 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v105 < v106 {
		v76 = v105
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v98 = v96
	goto L40
L42:
	;
	goto L37
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	base.MemoryCopy(m, v117, v118, v110)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v120
	v212 = v63
	goto L21
L46:
	;
	v126 = int32(_a_F_array_agg_combine_0)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v130
	v132 = v122 + v128
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v133 < v132 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v198 != 0 {
		goto L64
	} else {
		goto L65
	}
L48:
	;
	v135 = int32(1)
	if v132&(v132-v135) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v160 = int32(0)
	goto L57
L51:
	;
	v143 = v135 << (uint(int32(32)-base.I32_clz(v132)) % 32)
	goto L53
L52:
	;
	v143 = v132
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v148 = F_repalloc(m, v145, v143<<(uint(int32(2))%32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v153 = F_repalloc(m, v151, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v153
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v156 <= int32(0) {
		v198 = v156
		goto L47
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	v167 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v160))))
	if v170 == v167 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v198 = v194
	goto L47
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v160<<(uint(int32(2))%32))))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
	v180 = F_datumCopy(m, v177, v178, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L31
	} else {
		goto L62
	}
L60:
	;
	v182 = v167
	goto L61
L61:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v184 = int32(2)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v183+v160<<(uint(v184)%32)+v187<<(uint(v184)%32)))) = v182
	v193 = v160 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v193 < v194 {
		v160 = v193
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v182 = v180
	goto L61
L63:
	;
	goto L58
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	base.MemoryCopy(m, v203+v204, v206, v198)
	goto L66
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v132
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v127
	v212 = v48
	goto L21
L67:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_combine_1), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_array_agg_combine_2), int32(609), int32(_a_F_array_agg_combine_3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
	} else {
		goto L72
	}
L2:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v43 = int32(0)
	goto L2
L5:
	;
	goto L3
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L5
	}
L7:
	;
	goto L12
L8:
	;
	goto L9
L9:
	;
	v43 = int32(1)
	goto L2
L12:
	;
	v43 = int32(2)
	goto L2
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = F_pg_detoast_datum_packed(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L69
	}
L19:
	;
	return int32(0)
L20:
	;
	v49 = int32(1)
	v50 = v45 + v49
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v55 = v53 & v49
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = v50
	goto L23
L22:
	;
	v56 = v45 + int32(4)
	goto L23
L23:
	;
	if v53 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v56
	v89 = v11 + int32(16)
	v91 = F_pq_getmsgint(m, v89, int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L35
	}
L25:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v62 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v73 = int32(1)
	if v55 != 0 {
		v83 = int32(base.Ui32(v53)>>(uint(v73)%32)) - v73
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v65 = int32(16)
	goto L30
L29:
	;
	v65 = int32(0)
	goto L30
L30:
	;
	if base.Ui32((v62-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v72 = int32(4)
	goto L33
L32:
	;
	v72 = v65
	goto L33
L33:
	;
	v83 = v72
	goto L24
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v83 = int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v93 = F_pq_getmsgint64(m, v89)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_deserialize[0]))
	v98 = base.I32_wrap_i64(v93)
	v99 = F_initArrayResultWithSize(m, v91, v96, int32(0), v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v98
	v103 = F_pq_getmsgint(m, v89, int32(2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+24)) = uint16(v103)
	v106 = F_pq_getmsgbyte(m, v89)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+26)) = uint8(base.B2i32(v106 != int32(0)))
	v111 = F_pq_getmsgbyte(m, v89)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+27)) = uint8(v111)
	v114 = F_pq_getmsgbytes(m, v89, v98)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v98 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	base.MemoryCopy(m, v116, v114, v98)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+26)))
	if v118 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	F_pq_getmsgend(m, v11+int32(16))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L19
	} else {
		goto L68
	}
L46:
	;
	v124 = v98 << (uint(int32(2)) % 32)
	v125 = F_pq_getmsgbytes(m, v11+int32(16), v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L19
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	if v132 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v124 == int32(0) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	base.MemoryCopy(m, v129, v125, v124)
	goto L45
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	v137 = F_MemoryContextAlloc(m, v135, int32(32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L19
	} else {
		goto L54
	}
L52:
	;
	v150 = v132
	goto L53
L53:
	;
	if v93 <= int64(0) {
		goto L45
	} else {
		goto L57
	}
L54:
	;
	F_getTypeBinaryInputInfo(m, v91, v11, v137+int32(28))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	F_fmgr_info_cxt(m, v143, v137, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v137
	v150 = v137
	goto L53
L57:
	;
	v159 = int32(0)
	goto L58
L58:
	;
	v162 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v159))))
	if v165 == v162 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L45
L60:
	;
	v171 = F_pq_getmsgint(m, v11+int32(16), int32(4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L63
	}
L61:
	;
	v191 = v162
	goto L62
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v159<<(uint(int32(2))%32)))) = v191
	v199 = v159 + int32(1)
	if base.I64_extend_i32_s(v199) < v93 {
		v159 = v199
		goto L58
	} else {
		goto L67
	}
L63:
	;
	if v171 < int32(0) {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v175-v176 < v171 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v181 + v176
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v171 + v176
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v150)+28))
	v189 = F_ReceiveFunctionCall(m, v150, v11, v187, int32(-1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	v191 = v189
	goto L62
L67:
	;
	goto L59
L68:
	;
	m.G0 = v11 + int32(32)
	return v99
L69:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_deserialize_0), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_array_agg_deserialize_1), int32(797), int32(_a_F_array_agg_deserialize_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
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
	F_errcode(m, int32(50462850))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_array_agg_deserialize_3), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_array_agg_deserialize_1), int32(873), int32(_a_F_array_agg_deserialize_2))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_bitmap_copy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	if l4 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = int32(1) << (uint(l1&int32(7)) % 32)
	v17 = base.I32_div_s(l1, int32(8))
	v18 = l0 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v113)
	goto L1
L4:
	;
	v22 = v18
	v23 = v19
	v26 = l4
	v27 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	v57 = base.I32_div_s(l3, int32(8))
	v58 = l2 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v60 = v18
	v61 = v19
	v63 = v58
	v64 = l4
	v65 = v15
	v66 = int32(1) << (uint(l3&int32(7)) % 32)
	v67 = v59
	goto L15
L7:
	;
	v31 = v23 | v27
	v32 = int32(1)
	v33 = v26 - v32
	v35 = v27 << (uint(v32) % 32)
	if v35 == int32(256) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v47 != int32(1) {
		v112 = v45
		v113 = v46
		goto L3
	} else {
		goto L14
	}
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v31)
	if v33 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v45 = v22
	v46 = v31
	v47 = v35
	goto L11
L11:
	;
	if base.Ui32(int32(1)) < base.Ui32(v26) {
		v22 = v45
		v23 = v46
		v26 = v33
		v27 = v47
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v45 = v22 + v41
	v46 = v42
	v47 = v41
	goto L11
L13:
	;
	goto L8
L14:
	;
	goto L1
L15:
	;
	if v66&v67 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v90 == int32(1) {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v74 = v61 | v65
	goto L19
L18:
	;
	v74 = v61 & (v65 ^ int32(-1))
	goto L19
L19:
	;
	v75 = int32(1)
	v76 = v64 - v75
	v78 = v65 << (uint(v75) % 32)
	if v78 == int32(256) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v74)
	if v76 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v88 = v60
	v89 = v74
	v90 = v78
	goto L22
L22:
	;
	v92 = v66 << (uint(int32(1)) % 32)
	if v92 == int32(256) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v84 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v88 = v60 + v84
	v89 = v85
	v90 = v84
	goto L22
L24:
	;
	goto L16
L25:
	;
	if v76 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v101 = v63
	v102 = v92
	v103 = v67
	goto L27
L27:
	;
	if base.Ui32(int32(1)) < base.Ui32(v64) {
		v60 = v88
		v61 = v89
		v63 = v101
		v64 = v76
		v65 = v90
		v66 = v102
		v67 = v103
		goto L15
	} else {
		goto L29
	}
L28:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v98 = int32(1)
	v101 = v63 + v98
	v102 = v98
	v103 = v97
	goto L27
L29:
	;
	goto L24
L30:
	;
	v112 = v88
	v113 = v89
	goto L3
}
func F_array_fill(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 != int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v14 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v18 = v17
			} else {
				v18 = int32(0)
			}
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = F_get_fn_expr_argtype(m, v19, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_array_fill_0), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_fill_1), int32(_a_F_array_fill_2), int32(_a_F_array_fill_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = F_array_fill_internal(m, v10, int32(0), v18, v14, v21, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_array_fill_4), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_array_fill_1), int32(_a_F_array_fill_5), int32(_a_F_array_fill_3))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
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
func F_array_get_slice(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v248 int32
	_ = v248
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v836 int32
	_ = v836
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1014 int32
	_ = v1014
	v22 = m.G0
	v24 = v22 - int32(160)
	m.G0 = v24
	if l6 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v24 + int32(160)
	return v1014
L2:
	;
	v1003 = F_palloc0(m, int32(16))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L7
	} else {
		goto L170
	}
L3:
	;
	if v33 <= v164 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v28 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L30
	}
L7:
	;
	return int32(0)
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if base.B2i32(v33 < l1)|base.B2i32(base.Ui32(v33-int32(7)) < base.Ui32(int32(-6))) != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v42 = v28 + int32(16)
	v44 = v33 << (uint(int32(2)) % 32)
	v45 = v42 + v44
	v46 = int32(0)
	if l1 <= v46 {
		v164 = v46
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v55 = v46
	goto L11
L11:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v55))))
	if v71 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v164 = l1
	goto L3
L13:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v55))))
	if v94 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v55<<(uint(int32(2))%32)))) = v86
	v91 = v86
	goto L13
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45+v55<<(uint(int32(2))%32))))
	v86 = v77
	goto L14
L16:
	;
	goto L17
L17:
	;
	v79 = v55 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3+v79)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v45)))
	if v83 <= v81 {
		v91 = v81
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v86 = v83
	goto L14
L19:
	;
	if v127 < v126 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v118 = v55 << (uint(int32(2)) % 32)
	v122 = v115 + v116 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2+v118))) = v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118+l3)))
	v126 = v125
	v127 = v122
	goto L19
L21:
	;
	v98 = v55 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v45+v98)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v42)))
	v115 = v102
	v116 = v100
	goto L20
L22:
	;
	goto L23
L23:
	;
	v104 = v55 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2+v104)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v45+v104)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v42+v104)))
	if v106 < v108+v110 {
		v126 = v91
		v127 = v106
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v115 = v110
	v116 = v108
	goto L20
L25:
	;
	v132 = F_palloc0(m, int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v140 = v55 + int32(1)
	if v140 != l1 {
		v55 = v140
		goto L11
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v132))) = int64(64)
	v1014 = v132
	goto L1
L29:
	;
	goto L12
L30:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_array_get_slice_0), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_array_get_slice_1), int32(2067), int32(_a_F_array_get_slice_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v248 = int32(0)
	if v33 <= v248 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v186 = v164
	goto L36
L36:
	;
	v202 = v186 << (uint(int32(2)) % 32)
	v203 = l3 + v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202+v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v205
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v202+v42)))
	v212 = v205 + v209 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v202+l2))) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v212 < v214 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v220 = F_palloc0(m, int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v217 = v186 + int32(1)
	if v33 != v217 {
		v186 = v217
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L34
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v220)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v220))) = int64(64)
	v1014 = v220
	goto L1
L42:
	;
	v325 = v33 << (uint(int32(3)) % 32)
	v327 = v325 + int32(23)
	if v40 != 0 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	goto L42
L44:
	;
	if v33 != int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v264 = v248
	v267 = v248
	goto L48
L46:
	;
	v301 = v248
	goto L47
L47:
	;
	v306 = v301 << (uint(int32(2)) % 32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306+l2)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v24+v306))) = v309 - v311 + int32(1)
	goto L43
L48:
	;
	v268 = int32(2)
	v269 = v264 << (uint(v268) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v269+l2)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v269+l3)))
	v276 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v269))) = v272 - v274 + v276
	v280 = v269 | int32(4)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280+l2)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v280+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v24+v280))) = v283 - v285 + v276
	v291 = v264 + v268
	v293 = v267 + v268
	if v293 != v33&int32(2147483646) {
		v264 = v291
		v267 = v293
		goto L48
	} else {
		goto L50
	}
L49:
	;
	if v33&int32(1) == int32(0) {
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v301 = v291
	goto L47
L52:
	;
	v330 = v40
	goto L54
L53:
	;
	v330 = v327 & int32(-8)
	goto L54
L54:
	;
	v331 = v28 + v330
	v332 = v325 + v42
	if v40 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v334 = v332
	goto L57
L56:
	;
	v334 = int32(0)
	goto L57
L57:
	;
	v335 = F_array_slice_size(m, v331, v334, v33, v42, v45, l3, l2, l7, l8)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	if v40 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v351 = v350 + v335
	v352 = F_palloc0(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L64
	}
L60:
	;
	v337 = F_ArrayGetNItemsSafe(m, v33, v24)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v349 = int32(0)
	v350 = v327 & int32(120)
	goto L59
L63:
	;
	v342 = base.I32_div_s(v337+int32(7), int32(8))
	v345 = (v327 + v342) & int32(-8)
	v349 = v345
	v350 = v345
	goto L59
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v351 << (uint(int32(2)) % 32)
	v361 = v352 + int32(16)
	v362 = int32(0)
	v363 = base.B2i32(v44 == v362)
	if v363 == v362 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	base.MemoryCopy(m, v361, v24, v44)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if int32(0) < v33 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v369 = v361 + v44
	v370 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v33) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v465 = v349
	goto L70
L70:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v486 = int32(0)
	v497 = v33 - int32(1)
	if v497 < v486 {
		v575 = v486
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v376 = int32(0)
	v381 = v370
	goto L74
L72:
	;
	v417 = v370
	goto L73
L73:
	;
	v432 = v370
	v438 = v417
	goto L77
L74:
	;
	v398 = v369 + v381<<(uint(int32(2))%32)
	v399 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v398)+24)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v398)+16)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v398)+8)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v398))) = v399
	v407 = int32(8)
	v408 = v381 + v407
	v410 = v376 + v407
	if v410 != 0 {
		v376 = v410
		v381 = v408
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v417 = v408
	goto L73
L76:
	;
	goto L75
L77:
	;
	v456 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v369+v438<<(uint(int32(2))%32)))) = v456
	v461 = v432 + v456
	if v461 != v33 {
		v432 = v461
		v438 = v438 + v456
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v352)+8))
	v465 = v463
	goto L70
L79:
	;
	goto L78
L80:
	;
	v581 = F_array_seek(m, v331, v486, v334, v575, l7, l8)
	mBase = m.M
	v583 = v24 + int32(128)
	v587 = int32(2)
	v588 = v33 << (uint(v587) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v583+v588-int32(4)))) = int32(1)
	v595 = v33 - v587
	if v595 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	goto L80
L82:
	;
	v500 = int32(1)
	if v497 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v509 = v497
	v510 = v500
	v511 = v486
	v516 = v486
	goto L86
L84:
	;
	v552 = v497
	v553 = v500
	v554 = v486
	goto L85
L85:
	;
	v561 = v552 << (uint(int32(2)) % 32)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l3+v561)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v561+v45)))
	v575 = (v563-v565)*v553 + v554
	goto L81
L86:
	;
	v517 = int32(2)
	v518 = v509 << (uint(v517) % 32)
	v520 = v518 - int32(4)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l3+v520)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v45+v520)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v518+v42)))
	v528 = v527 * v510
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v518+l3)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v518+v45)))
	v537 = (v522-v524)*v528 + ((v531-v533)*v510 + v511)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v42+v520)))
	v540 = v539 * v528
	v542 = v509 - v517
	v544 = v516 + v517
	if v544 != v33&int32(-2) {
		v509 = v542
		v510 = v540
		v511 = v537
		v516 = v544
		goto L86
	} else {
		goto L88
	}
L87:
	;
	if v33&int32(1) == int32(0) {
		v575 = v537
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v552 = v542
	v553 = v540
	v554 = v537
	goto L85
L90:
	;
	v653 = v24 + int32(96)
	v654 = int32(0)
	if v33 <= v654 {
		goto L101
	} else {
		goto L102
	}
L91:
	;
	goto L90
L92:
	;
	if v33&int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v606 = v588 - int32(4)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v42+v606)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v583+v606)))
	*(*int32)(unsafe.Add(mBase, uint32(v583+v595<<(uint(int32(2))%32)))) = v608 * v610
	v615 = v33 - int32(3)
	goto L95
L94:
	;
	v615 = v595
	goto L95
L95:
	;
	if v595 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v621 = v615
	goto L97
L97:
	;
	v624 = int32(2)
	v625 = v621 << (uint(v624) % 32)
	v628 = v625 + int32(4)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v42+v628)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v583)))
	v633 = v630 * v632
	*(*int32)(unsafe.Add(mBase, uint32(v583+v625))) = v633
	v636 = v621 - int32(1)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v42+v625)))
	*(*int32)(unsafe.Add(mBase, uint32(v583+v636<<(uint(v624)%32)))) = v641 * v633
	if v636 != 0 {
		v621 = v621 - v624
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L91
L99:
	;
	goto L98
L100:
	;
	v731 = v24 - int32(-64)
	v732 = int32(0)
	v738 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v731+v33<<(uint(v738)%32)-int32(4)))) = v732
	v746 = v33 - v738
	if v732 <= v746 {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	goto L100
L102:
	;
	if v33 != int32(1) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v670 = v654
	v673 = v654
	goto L106
L104:
	;
	v707 = v654
	goto L105
L105:
	;
	v712 = v707 << (uint(int32(2)) % 32)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712+l2)))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v712+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v653+v712))) = v715 - v717 + int32(1)
	goto L101
L106:
	;
	v674 = int32(2)
	v675 = v670 << (uint(v674) % 32)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v675+l2)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v675+l3)))
	v682 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v653+v675))) = v678 - v680 + v682
	v686 = v675 | int32(4)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v686+l2)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v686+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v653+v686))) = v689 - v691 + v682
	v697 = v670 + v674
	v699 = v673 + v674
	if v699 != v33&int32(2147483646) {
		v670 = v697
		v673 = v699
		goto L106
	} else {
		goto L108
	}
L107:
	;
	if v33&int32(1) == int32(0) {
		goto L101
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v707 = v697
	goto L105
L110:
	;
	if v363 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L111:
	;
	v753 = v746
	v757 = v732
	goto L114
L112:
	;
	goto L113
L113:
	;
	goto L110
L114:
	;
	v760 = v753 << (uint(int32(2)) % 32)
	v761 = v731 + v760
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v583+v760)))
	v764 = int32(1)
	v765 = v763 - v764
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v765
	v768 = v753 + v764
	if v33 <= v768 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L113
L116:
	;
	v836 = int32(1)
	if int32(0) < v753 {
		v753 = v753 - v836
		v757 = v757 + v836
		goto L114
	} else {
		goto L125
	}
L117:
	;
	if v757&int32(1) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v774 = int32(2)
	v775 = v768 << (uint(v774) % 32)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v653+v775)))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v583+v775)))
	v783 = v765 - (v777-int32(1))*v781
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v783
	v787 = v753 + v774
	v788 = v783
	goto L120
L119:
	;
	v787 = v768
	v788 = v765
	goto L120
L120:
	;
	if v757 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v795 = v787
	v796 = v788
	goto L122
L122:
	;
	v801 = int32(2)
	v802 = v795 << (uint(v801) % 32)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v653+v802)))
	v805 = int32(1)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v583+v802)))
	v810 = v796 - (v804-v805)*v808
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v810
	v813 = v802 + int32(4)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v653+v813)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v583+v813)))
	v821 = v810 - (v815-v805)*v819
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v821
	v824 = v795 + v801
	if v824 != v33 {
		v795 = v824
		v796 = v821
		goto L122
	} else {
		goto L124
	}
L123:
	;
	goto L116
L124:
	;
	goto L123
L125:
	;
	goto L115
L126:
	;
	base.MemoryFill(m, v24+int32(32), int32(0), v44)
	goto L128
L127:
	;
	goto L128
L128:
	;
	v859 = v485 << (uint(int32(3)) % 32)
	if v465 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v862 = v361 + v859
	goto L131
L130:
	;
	v862 = int32(0)
	goto L131
L131:
	;
	if v465 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v869 = v465
	goto L134
L133:
	;
	v869 = (v859 + int32(23)) & int32(-8)
	goto L134
L134:
	;
	v871 = v33 - int32(1)
	v873 = v352 + v869
	v876 = v575
	v877 = v581
	v881 = v486
	goto L135
L135:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(-64)+v871<<(uint(int32(2))%32))))
	if v897 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v1014 = v352
	goto L1
L137:
	;
	v905 = F_array_seek(m, v903, v902, v334, int32(1), l7, l8)
	mBase = m.M
	v906 = v905 - v903
	if v906 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v902 = v876
	v903 = v877
	goto L137
L139:
	;
	goto L140
L140:
	;
	v901 = F_array_seek(m, v877, v876, v334, v897, l7, l8)
	mBase = m.M
	v902 = v897 + v876
	v903 = v901
	goto L137
L141:
	;
	base.MemoryCopy(m, v873, v903, v906)
	goto L143
L142:
	;
	goto L143
L143:
	;
	if v465 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v911 = int32(1) << (uint(v881&int32(7)) % 32)
	v913 = base.I32_div_s(v881, int32(8))
	v914 = v862 + v913
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	if v40 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	v936 = int32(1)
	v943 = v24 + int32(32)
	v945 = v24 + int32(96)
	if v33 <= int32(0) {
		goto L155
	} else {
		goto L156
	}
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v931)
	goto L146
L148:
	;
	v921 = base.I32_div_s(v902, int32(8))
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+v921))))
	if int32(base.Ui32(v923)>>(uint(v902&int32(7))%32))&int32(1) != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v931 = v911 | v915
	goto L147
L151:
	;
	v929 = v911 | v915
	goto L153
L152:
	;
	v929 = v915 & (v911 ^ int32(-1))
	goto L153
L153:
	;
	v931 = v929
	goto L147
L154:
	;
	if v999 != int32(-1) {
		v871 = v999
		v873 = v873 + v906
		v876 = v902 + v936
		v877 = v906 + v903
		v881 = v881 + v936
		goto L135
	} else {
		goto L169
	}
L155:
	;
	v999 = int32(-1)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v951 = int32(1)
	v952 = v33 - v951
	v954 = v952 << (uint(int32(2)) % 32)
	v955 = v943 + v954
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v945+v954)))
	v961 = base.I32_rem_s(v956+v951, v960)
	*(*int32)(unsafe.Add(mBase, uint32(v955))) = v961
	if v952 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v999 = v989
	goto L154
L159:
	;
	v963 = v952
	v966 = v961
	goto L162
L160:
	;
	goto L161
L161:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	if v987 != 0 {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	if v966 != 0 {
		v989 = v963
		goto L158
	} else {
		goto L164
	}
L163:
	;
	goto L161
L164:
	;
	v968 = int32(1)
	v969 = v963 - v968
	v971 = v969 << (uint(int32(2)) % 32)
	v972 = v943 + v971
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v945+v971)))
	v978 = base.I32_rem_s(v973+v968, v977)
	*(*int32)(unsafe.Add(mBase, uint32(v972))) = v978
	if v969 != 0 {
		v963 = v969
		v966 = v978
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v988 = int32(0)
	goto L168
L167:
	;
	v988 = int32(-1)
	goto L168
L168:
	;
	v989 = v988
	goto L158
L169:
	;
	goto L136
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1003))) = int64(64)
	v1014 = v1003
	goto L1
}
func F_array_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_array_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_array_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_array_lower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13 == int32(-1) {
			v16 = int32(28)
		} else {
			v16 = int32(4)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7+v16)))
		if base.Ui32(v18-int32(7)) <= base.Ui32(int32(-7)) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v47 = v40
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v47 = v7 + v41<<(uint(int32(2))%32) + int32(16)
				}
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+v27<<(uint(int32(2))%32)-int32(4))))
				return v53
			}
		}
	}
}
func F_array_position(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v181
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L61
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L56
	}
L5:
	;
	m.G0 = v15 + int32(16)
	goto L1
L6:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v176)
	v181 = int32(0)
	goto L5
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(2) <= v24 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v24 != int32(1) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v29 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v20)+16))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v46 == int32(3) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v32 = F_array_contains_nulls(m, v20)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v41 = v39
	v42 = int32(4)
	goto L12
L16:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v41 = v2
	v42 = v36 << (uint(int32(2)) % 32)
	goto L12
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v49 == int32(1) {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	v53 = v45
	goto L20
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v55 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v53 = v52
	goto L20
L22:
	;
	v97 = v45 - int32(1)
	v99 = F_array_create_iterator(m, v20, int32(0), v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L34
	}
L23:
	;
	F_get_typlenbyvalalign(m, v43, v71+int32(4), v71+int32(6), v71+int32(7))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L29
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v60 = F_MemoryContextAlloc(m, v58, int32(48))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v69 == v43 {
		v94 = v55
		goto L22
	} else {
		goto L28
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v43 ^ int32(-1)
	v71 = v65
	goto L23
L28:
	;
	v71 = v55
	goto L23
L29:
	;
	v81 = F_lookup_type_cache(m, v43, int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	if v83 == int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v43
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	F_fmgr_info_cxt(m, v87, v71+int32(20), v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v94 = v71
	goto L22
L33:
	;
	F_array_free_iterator(m, v99)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L50
	}
L34:
	;
	v105 = F_array_iterate(m, v99, v15+int32(12), v15+int32(11))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v105 == int32(0) {
		v148 = v97
		v157 = v2
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v113 = v97
	goto L37
L37:
	;
	v124 = v113 + int32(1)
	if v124 < v53 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v148 = v124
	v157 = v2
	goto L33
L39:
	;
	v144 = F_array_iterate(m, v99, v15+int32(12), v15+int32(11))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L48
	}
L40:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+11)))
	if (v126|v29)&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v148 = v124
	v157 = int32(1)
	goto L33
L42:
	;
	if v126&v29 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v134 = F_FunctionCall2Coll(m, v94+int32(20), v18, v41, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	if v134 == int32(0) {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L41
L48:
	;
	if v144 != 0 {
		v113 = v124
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v160 != v20 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_pfree(m, v20)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v157 != 0 {
		v181 = v148
		goto L5
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	goto L6
L56:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v201 = F_format_type_be(m, v43)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v201
	F_errmsg(m, int32(_a_F_array_position_0), v15)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1411), int32(_a_F_array_position_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_array_position_3), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1377), int32(_a_F_array_position_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_array_position_4), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1348), int32(_a_F_array_position_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_shuffle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if v11 <= int32(0) {
			v33 = v7
			return v33
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			if v14 < int32(2) {
				v33 = v7
				return v33
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				if v19 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v20 == v17 {
						v28 = v19
						v29 = v14
						v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v31
							return v33
						}
					} else {
						v23 = F_lookup_type_cache(m, v17, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v23
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							v28 = v23
							v29 = v27
							v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v31
								return v33
							}
						}
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v23
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						v28 = v23
						v29 = v27
						v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v31
							return v33
						}
					}
				}
			}
		}
	}
}
func F_array_sort_order(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = base.B2i32(v8 != int32(0))
		v11 = F_array_sort_internal(m, v4, v10, v10, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_array_sort_order_nulls_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_array_sort_internal(m, v3, base.B2i32(v7 != v8), base.B2i32(v10 != v8), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
func F_array_subscript_assign(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if int32(0) < v13 {
		if v8&int32(1) != 0 {
			return
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
			if v18 == int32(0) {
				v32 = v10
				v33 = v13
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
				v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
				v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
					return
				}
			} else {
				return
			}
		}
	} else {
		if v8&int32(1) == int32(0) {
			v32 = v10
			v33 = v13
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
			v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
			v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
			v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				return
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v26 = F_construct_empty_array(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v29)
				v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
				v32 = v26
				v33 = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
				v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
				v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
					return
				}
			}
		}
	}
}
func F_array_subscript_fetch_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+4)))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+6)))
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+9)))
	v21 = F_array_get_slice(m, v13, v7, v8+int32(12), v8+int32(36), v16, v17, v18, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
		return
	}
}
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 float32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v96 int32
	_ = v96
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v115 float64
	_ = v115
	v5 = float64(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return v115
L2:
	;
	v115 = float64(10)
	goto L1
L3:
	;
	v13 = l1
	goto L6
L4:
	;
	if l0 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L5:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)))
	if v42 != 0 {
		goto L4
	} else {
		goto L18
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	switch v17 - int32(7) {
	case 0:
		goto L8
	default:
		goto L4
	case 20:
		goto L9
	case 22:
		goto L10
	case 28:
		goto L5
	}
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
	if v30 != 0 {
		v115 = v5
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v29 != 0 {
		v13 = v29
		goto L6
	} else {
		goto L13
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != int32(27) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 != int32(34) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L2
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return float64(0)
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v39 = F_ArrayGetNItemsSafe(m, v36, v32+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v115 = base.F64_convert_i32_s(v39)
	goto L1
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v43 == int32(0) {
		v115 = v5
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v115 = base.F64_convert_i32_s(v46)
	goto L1
L20:
	;
	F_examine_variable(m, l0, v13, int32(0), v8+int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v56 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v64 = F_get_attstatsslot(m, v8+int32(12), v56, int32(5), int32(0), int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	if v64 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v66 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v98 = v5
	goto L26
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v99 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v92 = float64(0)
	goto L29
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v70+v66<<(uint(int32(2))%32)-int32(4))))
	v77 = base.F64_promote_f32(v76)
	v78 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)))|base.F64_gt(v77, v78) != 0 {
		v91 = v78
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_free_attstatsslot(m, v8+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L34
	}
L30:
	;
	v92 = v91
	goto L29
L31:
	;
	goto L30
L32:
	;
	v87 = float64(1)
	if base.F64_le(v77, v87) != 0 {
		v91 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v91 = base.F64_nearest(v77)
	goto L31
L34:
	;
	v98 = v92
	goto L26
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	m.T0[v100].(func(*base.Module, int32))(m, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L15
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.F64_gt(v98, float64(0)) != 0 {
		v115 = v98
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	goto L2
}
func F_getArrayIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v15 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v8+int32(8), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v19)
		v21 = int32(2)
		if v15 == v21 {
			v97 = v21
			m.G0 = v8 + int32(16)
			return v97
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			if v24 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if v27 == int32(0) {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
					if v40 != int32(1) {
						v97 = v21
						m.G0 = v8 + int32(16)
						return v97
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(51118210))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_getArrayIndex_0), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3474), int32(_a_F_getArrayIndex_2))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					if v30 != int32(1) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						if v40 != int32(1) {
							v97 = v21
							m.G0 = v8 + int32(16)
							return v97
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(51118210))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_getArrayIndex_0), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3474), int32(_a_F_getArrayIndex_2))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
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
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v35 = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v36 == int32(2) {
							v59 = int32(0)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							v64 = F_DirectFunctionCall2Coll(m, int32(1426), v59, v62, v59)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = F_pg_detoast_datum(m, v64)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v70 = F_numeric_int4_opt_error(m, v66, v8+int32(7))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
										if v73 != int32(1) {
											v97 = v59
											m.G0 = v8 + int32(16)
											return v97
										} else {
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v77 != int32(1) {
												v97 = int32(2)
												m.G0 = v8 + int32(16)
												return v97
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(51118210))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_getArrayIndex_3), int32(0))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3486), int32(_a_F_getArrayIndex_2))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
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
							}
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							if v40 != int32(1) {
								v97 = v21
								m.G0 = v8 + int32(16)
								return v97
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(51118210))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_getArrayIndex_0), int32(0))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3474), int32(_a_F_getArrayIndex_2))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
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
			} else {
				v35 = v24
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				if v36 == int32(2) {
					v59 = int32(0)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v64 = F_DirectFunctionCall2Coll(m, int32(1426), v59, v62, v59)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = F_pg_detoast_datum(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v70 = F_numeric_int4_opt_error(m, v66, v8+int32(7))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
								if v73 != int32(1) {
									v97 = v59
									m.G0 = v8 + int32(16)
									return v97
								} else {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v77 != int32(1) {
										v97 = int32(2)
										m.G0 = v8 + int32(16)
										return v97
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(51118210))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_getArrayIndex_3), int32(0))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3486), int32(_a_F_getArrayIndex_2))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
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
					}
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
					if v40 != int32(1) {
						v97 = v21
						m.G0 = v8 + int32(16)
						return v97
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(51118210))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_getArrayIndex_0), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3474), int32(_a_F_getArrayIndex_2))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
	}
}
func F_get_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 < v10 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = int32(1)
		v15 = v10 - v14
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
		if v17 != v14 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v20 == int32(0) {
				return int32(0)
			} else {
				v24 = v15 << (uint(int32(2)) % 32)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v24+v25)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20)))
				if v27 != v29 {
					return int32(0)
				} else {
					if v10 < v11 {
						v33 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v10+v13))) = uint8(v33)
						return v33
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v37 == int32(0) {
							return int32(0)
						} else {
							if l1 != 0 {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								if v41 != 0 {
									v48 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
									return int32(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v44 = F_cstring_to_text_with_len(m, v37, v42-v37)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = v44
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
										return int32(0)
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								v44 = F_cstring_to_text_with_len(m, v37, v42-v37)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
									return int32(0)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v13 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return int32(0)
L2:
	;
	v17 = v13 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v17+v18))) = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v17)))
	if base.Ui32(v24) < base.Ui32(int32(-2147483647)) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if v14|v13 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = m.G0
	v32 = v30 - int32(80)
	m.G0 = v32
	if v27 == int32(_a_F_get_array_start_0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	m.G0 = v32 + int32(80)
	if v114 != 0 {
		goto L42
	} else {
		goto L43
	}
L7:
	;
	v114 = int32(16)
	goto L6
L8:
	;
	goto L9
L9:
	;
	base.MemoryCopy(m, v32+int32(12), v27, int32(68))
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v43 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v47 != int32(5) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v50 = int32(11)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v60 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v54 = int32(6)
	goto L15
L14:
	;
	v54 = v50
	goto L15
L15:
	;
	if v47 == int32(12) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = v50
	goto L18
L17:
	;
	v57 = v54
	goto L18
L18:
	;
	v114 = v57
	goto L6
L19:
	;
	return int32(0)
L20:
	;
	if v60 != 0 {
		v114 = v60
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v64 == int32(6) {
		v105 = v2
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v109 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L19
	} else {
		goto L40
	}
L23:
	;
	v72 = v2
	goto L24
L24:
	;
	v77 = F_parse_array_element(m, v32+int32(12), int32(_a_F_get_array_start_1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v114 = v96
	goto L6
L26:
	;
	if v77 != 0 {
		v114 = v77
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v80 = v72 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v81 != int32(7) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v81 == int32(6) {
		v105 = v80
		goto L22
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v96 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L38
	}
L31:
	;
	v86 = int32(11)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v90 = int32(7)
	goto L34
L33:
	;
	v90 = v86
	goto L34
L34:
	;
	if v81 == int32(12) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v93 = v86
	goto L37
L36:
	;
	v93 = v90
	goto L37
L37:
	;
	v114 = v93
	goto L6
L38:
	;
	if v96 == int32(0) {
		v72 = v80
		goto L24
	} else {
		goto L39
	}
L39:
	;
	goto L25
L40:
	;
	if v109 != 0 {
		v114 = v109
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v105
	v114 = int32(0)
	goto L6
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_json_errsave_error(m, v114, v123, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L19
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v132 = v129 + v13<<(uint(int32(2))%32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v127 < int32(0)-v133 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v133 + v127
	goto L1
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v139
	goto L1
}
func F_initArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == int32(0) {
		v13 = F_get_element_type(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = F_format_type_be(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
							F_errmsg(m, int32(_a_F_initArrayResultArr_0), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_initArrayResultArr_1), int32(_a_F_initArrayResultArr_2), int32(_a_F_initArrayResultArr_3))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
			} else {
				v19 = v13
				if v4 != 0 {
					v24 = F_AllocSetContextCreateInternal(m, l2, int32(_a_F_initArrayResultArr_4), int32(0), int32(_a_F_initArrayResultArr_5), int32(_a_F_initArrayResultArr_6))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v24
						v28 = F_MemoryContextAllocZero(m, v26, int32(92))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
							m.G0 = v9 + int32(16)
							return v28
						}
					}
				} else {
					v26 = l2
					v28 = F_MemoryContextAllocZero(m, v26, int32(92))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
						m.G0 = v9 + int32(16)
						return v28
					}
				}
			}
		}
	} else {
		v19 = l1
		if v4 != 0 {
			v24 = F_AllocSetContextCreateInternal(m, l2, int32(_a_F_initArrayResultArr_4), int32(0), int32(_a_F_initArrayResultArr_5), int32(_a_F_initArrayResultArr_6))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = v24
				v28 = F_MemoryContextAllocZero(m, v26, int32(92))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
					m.G0 = v9 + int32(16)
					return v28
				}
			}
		} else {
			v26 = l2
			v28 = F_MemoryContextAllocZero(m, v26, int32(92))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
				m.G0 = v9 + int32(16)
				return v28
			}
		}
	}
}
func F_makeArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
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
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	v4 = int32(0)
	v14 = int32(_a_F_makeArrayResultArr_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultArr[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultArr[0])) = l1
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v18 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultArr[0])) = v15
	if l2 != 0 {
		goto L47
	} else {
		goto L48
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v23 = F_palloc0(m, int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v33 = l0 + int32(32)
	v34 = F_ArrayGetNItemsSafe(m, v18, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(64)
	v178 = v23
	goto L1
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v38 = l0 + int32(56)
	F_ArrayCheckBounds(m, v36, v33, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = v41 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = v61 + v44
	v63 = F_palloc0(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L13
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v50 = base.I32_div_s(v46+int32(7), int32(8))
	v55 = (v43 + v50 + int32(23)) & int32(-8)
	v60 = v55
	v61 = v55
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = v4
	v61 = (v43 + int32(23)) & int32(-8)
	goto L9
L13:
	;
	v65 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v62 << (uint(v65) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = v63 + int32(16)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v77 = v75 << (uint(v65) % 32)
	if v77 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	base.MemoryCopy(m, v74, v33, v77)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v81 = v79 << (uint(int32(2)) % 32)
	if v81 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	base.MemoryCopy(m, v74+v68<<(uint(int32(2))%32), v38, v81)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v86 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L22
L21:
	;
	v96 = v86
	goto L22
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v97 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	base.MemoryCopy(m, v96+v63, v99, v97)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v101 == int32(0) {
		v178 = v63
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v110 = v74 + v105<<(uint(int32(3))%32)
	goto L29
L28:
	;
	v110 = int32(0)
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v111 <= int32(0) {
		v178 = v63
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v116 = int32(1)
	v119 = v116
	v123 = v111
	v124 = v110
	v125 = v115
	v126 = v101
	v127 = v116
	v128 = v114
	goto L31
L31:
	;
	if v127&v128 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v150 == int32(1) {
		v178 = v63
		goto L1
	} else {
		goto L46
	}
L33:
	;
	v136 = v119 | v125
	goto L35
L34:
	;
	v136 = v125 & (v119 ^ int32(-1))
	goto L35
L35:
	;
	v137 = int32(1)
	v138 = v123 - v137
	v140 = v119 << (uint(v137) % 32)
	if v140 == int32(256) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v136)
	if v138 == int32(0) {
		v178 = v63
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v150 = v140
	v151 = v124
	v152 = v136
	goto L38
L38:
	;
	v154 = v127 << (uint(int32(1)) % 32)
	if v154 == int32(256) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	v147 = int32(1)
	v150 = v147
	v151 = v124 + v147
	v152 = v146
	goto L38
L40:
	;
	goto L32
L41:
	;
	if v138 == int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v163 = v126
	v164 = v154
	v165 = v128
	goto L43
L43:
	;
	if base.Ui32(int32(1)) < base.Ui32(v123) {
		v119 = v150
		v123 = v138
		v124 = v151
		v125 = v152
		v126 = v163
		v127 = v164
		v128 = v165
		goto L31
	} else {
		goto L45
	}
L44:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v160 = int32(1)
	v163 = v126 + v160
	v164 = v160
	v165 = v159
	goto L43
L45:
	;
	goto L40
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v152)
	v178 = v63
	goto L1
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_MemoryContextDelete(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	return v178
L50:
	;
	goto L49
}
