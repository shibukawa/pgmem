package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v44 = v41
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v41 = v37
	goto L2
L4:
	;
	v33 = int32(0)
	if v14 == v33 {
		v41 = v33
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v19 - int32(435) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L6:
	;
	if v14 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v36 = v26
	v37 = int32(1)
	goto L3
L11:
	;
	v44 = int32(2)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
	v36 = v31
	v37 = int32(2)
	goto L3
L14:
	;
	v36 = v33
	v37 = v2
	goto L3
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v45 == int32(0) {
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
	v232 = m.ExcPending
	if v232 != 0 {
		goto L31
	} else {
		goto L67
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = v48
	goto L20
L19:
	;
	v49 = v2
	goto L20
L20:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v50 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v217)
L22:
	;
	if v49 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v53 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v49 != 0 {
		v217 = v49
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
	v217 = int32(0)
	goto L21
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v64 = F_initArrayResultWithSize(m, v60, v61, int32(0), v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v125 <= int32(0) {
		v217 = v49
		goto L21
	} else {
		goto L46
	}
L31:
	;
	return int64(0)
L32:
	;
	v68 = int32(_a_F_array_agg_combine_0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if int32(0) < v73 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v77 = int32(0)
	goto L36
L34:
	;
	v113 = v73
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v69
	if v113 != 0 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v77))))
	if v88 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v113 = v108
	goto L35
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v91+v77<<(uint(int32(3))%32))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+26)))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+24)))
	v98 = F_datumCopy(m, v95, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L31
	} else {
		goto L41
	}
L39:
	;
	v100 = int64(0)
	goto L40
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v101+v77<<(uint(int32(3))%32)))) = v100
	v107 = v77 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v107 < v108 {
		v77 = v107
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v100 = v98
	goto L40
L42:
	;
	goto L37
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	base.MemoryCopy(m, v120, v121, v113)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v123
	v217 = v64
	goto L21
L46:
	;
	v129 = int32(_a_F_array_agg_combine_0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v133
	v135 = v125 + v131
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v136 < v135 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v203 != 0 {
		goto L64
	} else {
		goto L65
	}
L48:
	;
	v138 = int32(1)
	if v135&(v135-v138) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v163 = int32(0)
	goto L57
L51:
	;
	v146 = v138 << (uint(int32(32)-base.I32_clz(v135)) % 32)
	goto L53
L52:
	;
	v146 = v135
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v151 = F_repalloc(m, v148, v146<<(uint(int32(3))%32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v151
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v156 = F_repalloc(m, v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v156
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v159 <= int32(0) {
		v203 = v159
		goto L47
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+v163))))
	if v174 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v203 = v198
	goto L47
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v177+v163<<(uint(int32(3))%32))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+26)))
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+24)))
	v184 = F_datumCopy(m, v181, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L31
	} else {
		goto L62
	}
L60:
	;
	v186 = int64(0)
	goto L61
L61:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v188 = int32(3)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v187+v163<<(uint(v188)%32)+v191<<(uint(v188)%32)))) = v186
	v197 = v163 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v197 < v198 {
		v163 = v197
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v186 = v184
	goto L61
L63:
	;
	goto L58
L64:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	base.MemoryCopy(m, v208+v209, v211, v203)
	goto L66
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v135
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_combine[0])) = v130
	v217 = v49
	goto L21
L67:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_combine_1), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_array_agg_combine_2), int32(610), int32(_a_F_array_agg_combine_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
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
func F_array_agg_deserialize(m *base.Module, l0 int32) int64 {
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	v236 = m.ExcPending
	if v236 != 0 {
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
	switch v18 - int32(435) {
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v222 = m.ExcPending
	if v222 != 0 {
		goto L19
	} else {
		goto L69
	}
L19:
	;
	return int64(0)
L20:
	;
	v49 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v53 = v51 & v49
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v54 = v49
	goto L23
L22:
	;
	v54 = int32(4)
	goto L23
L23:
	;
	if v51 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v45 + v54
	v88 = v11 + int32(16)
	v90 = F_pq_getmsgint(m, v88, int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L19
	} else {
		goto L35
	}
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	if v61 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v72 = int32(1)
	if v53 != 0 {
		v82 = int32(base.Ui32(v51)>>(uint(v72)%32)) - v72
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v64 = int32(16)
	goto L30
L29:
	;
	v64 = int32(0)
	goto L30
L30:
	;
	if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v71 = int32(4)
	goto L33
L32:
	;
	v71 = v64
	goto L33
L33:
	;
	v82 = v71
	goto L24
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v92 = F_pq_getmsgint64(m, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_deserialize[0]))
	v97 = base.I32_wrap_i64(v92)
	v98 = F_initArrayResultWithSize(m, v90, v95, int32(0), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v97
	v102 = F_pq_getmsgint(m, v88, int32(2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+24)) = uint16(v102)
	v105 = F_pq_getmsgbyte(m, v88)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+26)) = uint8(base.B2i32(v105 != int32(0)))
	v110 = F_pq_getmsgbyte(m, v88)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+27)) = uint8(v110)
	v113 = F_pq_getmsgbytes(m, v88, v97)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v97 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	base.MemoryCopy(m, v115, v113, v97)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+26)))
	if v117 == int32(1) {
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
	v123 = v97 << (uint(int32(3)) % 32)
	v124 = F_pq_getmsgbytes(m, v11+int32(16), v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v131 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v123 == int32(0) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	base.MemoryCopy(m, v128, v124, v123)
	goto L45
L51:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	v136 = F_MemoryContextAlloc(m, v134, int32(32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L19
	} else {
		goto L54
	}
L52:
	;
	v149 = v131
	goto L53
L53:
	;
	if v92 <= int64(0) {
		goto L45
	} else {
		goto L57
	}
L54:
	;
	F_getTypeBinaryInputInfo(m, v90, v11, v136+int32(28))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	F_fmgr_info_cxt(m, v142, v136, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+16)) = v136
	v149 = v136
	goto L53
L57:
	;
	v153 = int32(0)
	goto L58
L58:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v153))))
	if v164 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L45
L60:
	;
	v170 = F_pq_getmsgint(m, v11+int32(16), int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L63
	}
L61:
	;
	v192 = int64(0)
	goto L62
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v193+v153<<(uint(int32(3))%32)))) = v192
	v199 = v153 + int32(1)
	if base.I64_extend_i32_s(v199) < v92 {
		v153 = v199
		goto L58
	} else {
		goto L67
	}
L63:
	;
	if v170 < int32(0) {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v174-v175 < v170 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v180 + v175
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v175 + v170
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v149)+28))
	v188 = F_ReceiveFunctionCall(m, v149, v11, v186, int32(-1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	v192 = v188
	goto L62
L67:
	;
	goto L59
L68:
	;
	m.G0 = v11 + int32(32)
	return base.I64_extend_i32_u(v98)
L69:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_deserialize_0), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_array_agg_deserialize_1), int32(798), int32(_a_F_array_agg_deserialize_2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
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
	v239 = m.ExcPending
	if v239 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_array_agg_deserialize_3), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_array_agg_deserialize_1), int32(874), int32(_a_F_array_agg_deserialize_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
func F_array_fill(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v18 int64
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v6 != int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v14 == int32(0) {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				v18 = v17
			} else {
				v18 = int64(0)
			}
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = F_get_fn_expr_argtype(m, v19, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_array_fill_0), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_array_fill_1), int32(_a_F_array_fill_2), int32(_a_F_array_fill_3))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
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
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v26)
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_array_fill_4), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_array_fill_1), int32(_a_F_array_fill_5), int32(_a_F_array_fill_3))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
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
func F_array_get_slice(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v190 int32
	_ = v190
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v847 int32
	_ = v847
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	v23 = m.G0
	v25 = v23 - int32(160)
	m.G0 = v25
	if l6 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v25 + int32(160)
	return base.I64_extend_i32_u(v1027)
L2:
	;
	v1016 = F_palloc0(m, int32(16))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L7
	} else {
		goto L172
	}
L3:
	;
	if v35 <= v167 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v30 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L30
	}
L7:
	;
	return int64(0)
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.B2i32(v35 < l1)|base.B2i32(base.Ui32(v35-int32(7)) < base.Ui32(int32(-6))) != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v44 = v30 + int32(16)
	v46 = v35 << (uint(int32(2)) % 32)
	v47 = v44 + v46
	v48 = int32(0)
	if l1 <= v48 {
		v167 = v48
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v57 = v48
	goto L11
L11:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v57))))
	if v74 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v167 = l1
	goto L3
L13:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v57))))
	if v97 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v57<<(uint(int32(2))%32)))) = v89
	v94 = v89
	goto L13
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v47+v57<<(uint(int32(2))%32))))
	v89 = v80
	goto L14
L16:
	;
	goto L17
L17:
	;
	v82 = v57 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3+v82)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v47)))
	if v86 <= v84 {
		v94 = v84
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v89 = v86
	goto L14
L19:
	;
	if v130 < v129 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v121 = v57 << (uint(int32(2)) % 32)
	v125 = v118 + v119 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2+v121))) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3+v121)))
	v129 = v128
	v130 = v125
	goto L19
L21:
	;
	v101 = v57 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v47+v101)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v44)))
	v118 = v105
	v119 = v103
	goto L20
L22:
	;
	goto L23
L23:
	;
	v107 = v57 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2+v107)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v47+v107)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v44+v107)))
	if v109 < v111+v113 {
		v129 = v94
		v130 = v109
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v118 = v113
	v119 = v111
	goto L20
L25:
	;
	v135 = F_palloc0(m, int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v143 = v57 + int32(1)
	if v143 != l1 {
		v57 = v143
		goto L11
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = int64(64)
	v1027 = v135
	goto L1
L29:
	;
	goto L12
L30:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_array_get_slice_0), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_array_get_slice_1), int32(2071), int32(_a_F_array_get_slice_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v254 = int32(0)
	if v35 <= v254 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v190 = v167
	goto L36
L36:
	;
	v207 = v190 << (uint(int32(2)) % 32)
	v208 = l3 + v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v207+v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v210
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207+v44)))
	v217 = v210 + v214 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207+l2))) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v217 < v219 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v225 = F_palloc0(m, int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v222 = v190 + int32(1)
	if v35 != v222 {
		v190 = v222
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L34
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v225))) = int64(64)
	v1027 = v225
	goto L1
L42:
	;
	v331 = v35 << (uint(int32(3)) % 32)
	v333 = v331 + int32(23)
	if v42 != 0 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	goto L42
L44:
	;
	if v35 != int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v270 = v254
	v273 = v254
	goto L48
L46:
	;
	v307 = v254
	goto L47
L47:
	;
	v312 = v307 << (uint(int32(2)) % 32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312+l2)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v25+v312))) = v315 - v317 + int32(1)
	goto L43
L48:
	;
	v274 = int32(2)
	v275 = v270 << (uint(v274) % 32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275+l2)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v275+l3)))
	v282 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v275))) = v278 - v280 + v282
	v286 = v275 | int32(4)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286+l2)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v25+v286))) = v289 - v291 + v282
	v297 = v270 + v274
	v299 = v273 + v274
	if v299 != v35&int32(2147483646) {
		v270 = v297
		v273 = v299
		goto L48
	} else {
		goto L50
	}
L49:
	;
	if v35&int32(1) == int32(0) {
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v307 = v297
	goto L47
L52:
	;
	v336 = v42
	goto L54
L53:
	;
	v336 = v333 & int32(-8)
	goto L54
L54:
	;
	v337 = v30 + v336
	v338 = v331 + v44
	if v42 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v340 = v338
	goto L57
L56:
	;
	v340 = int32(0)
	goto L57
L57:
	;
	v341 = F_array_slice_size(m, v337, v340, v35, v44, v47, l3, l2, l7, l8)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	if v42 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v357 = v341 + v356
	v358 = F_palloc0(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L7
	} else {
		goto L64
	}
L60:
	;
	v343 = F_ArrayGetNItemsSafe(m, v35, v25)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v355 = int32(0)
	v356 = v333 & int32(120)
	goto L59
L63:
	;
	v348 = base.I32_div_s(v343+int32(7), int32(8))
	v351 = (v333 + v348) & int32(-8)
	v355 = v351
	v356 = v351
	goto L59
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v358)+8)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v358)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v357 << (uint(int32(2)) % 32)
	v367 = v358 + int32(16)
	v368 = int32(0)
	v369 = base.B2i32(v46 == v368)
	if v369 == v368 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	base.MemoryCopy(m, v367, v25, v46)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if int32(0) < v35 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v375 = v367 + v46
	v376 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v35) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v474 = v355
	goto L70
L70:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v496 = int32(0)
	v507 = v35 - int32(1)
	if v507 < v496 {
		v585 = v496
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v382 = int32(0)
	v387 = v376
	goto L74
L72:
	;
	v424 = v376
	goto L73
L73:
	;
	v446 = v424
	v449 = v376
	goto L77
L74:
	;
	v405 = v375 + v387<<(uint(int32(2))%32)
	v406 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v405)+24)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v405)+16)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v405)+8)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v405))) = v406
	v414 = int32(8)
	v415 = v387 + v414
	v417 = v382 + v414
	if v417 != 0 {
		v382 = v417
		v387 = v415
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v424 = v415
	goto L73
L76:
	;
	goto L75
L77:
	;
	v465 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v375+v446<<(uint(int32(2))%32)))) = v465
	v470 = v449 + v465
	if v470 != v35 {
		v446 = v446 + v465
		v449 = v470
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v474 = v472
	goto L70
L79:
	;
	goto L78
L80:
	;
	v591 = F_array_seek(m, v337, v496, v340, v585, l7, l8)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L7
	} else {
		goto L90
	}
L81:
	;
	goto L80
L82:
	;
	v510 = int32(1)
	if v507 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v519 = v507
	v520 = v510
	v521 = v496
	v526 = v496
	goto L86
L84:
	;
	v562 = v507
	v563 = v510
	v564 = v496
	goto L85
L85:
	;
	v571 = v562 << (uint(int32(2)) % 32)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l3+v571)))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v571+v47)))
	v585 = (v573-v575)*v563 + v564
	goto L81
L86:
	;
	v527 = int32(2)
	v528 = v519 << (uint(v527) % 32)
	v530 = v528 - int32(4)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l3+v530)))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v47+v530)))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v528+v44)))
	v538 = v537 * v520
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v528+l3)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v528+v47)))
	v547 = (v532-v534)*v538 + ((v541-v543)*v520 + v521)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v44+v530)))
	v550 = v549 * v538
	v552 = v519 - v527
	v554 = v526 + v527
	if v554 != v35&int32(-2) {
		v519 = v552
		v520 = v550
		v521 = v547
		v526 = v554
		goto L86
	} else {
		goto L88
	}
L87:
	;
	if v35&int32(1) == int32(0) {
		v585 = v547
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v562 = v552
	v563 = v550
	v564 = v547
	goto L85
L90:
	;
	v594 = v25 + int32(128)
	v598 = int32(2)
	v599 = v35 << (uint(v598) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v594+v599-int32(4)))) = int32(1)
	v606 = v35 - v598
	if v606 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v664 = v25 + int32(96)
	v665 = int32(0)
	if v35 <= v665 {
		goto L102
	} else {
		goto L103
	}
L92:
	;
	goto L91
L93:
	;
	if v35&int32(1) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v617 = v599 - int32(4)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v44+v617)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v594+v617)))
	*(*int32)(unsafe.Add(mBase, uint32(v594+v606<<(uint(int32(2))%32)))) = v619 * v621
	v626 = v35 - int32(3)
	goto L96
L95:
	;
	v626 = v606
	goto L96
L96:
	;
	if v606 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v632 = v626
	goto L98
L98:
	;
	v635 = int32(2)
	v636 = v632 << (uint(v635) % 32)
	v639 = v636 + int32(4)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v44+v639)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v639+v594)))
	v644 = v641 * v643
	*(*int32)(unsafe.Add(mBase, uint32(v594+v636))) = v644
	v647 = v632 - int32(1)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v44+v636)))
	*(*int32)(unsafe.Add(mBase, uint32(v594+v647<<(uint(v635)%32)))) = v652 * v644
	if v647 != 0 {
		v632 = v632 - v635
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L92
L100:
	;
	goto L99
L101:
	;
	v742 = v25 - int32(-64)
	v743 = int32(0)
	v749 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v742+v35<<(uint(v749)%32)-int32(4)))) = v743
	v757 = v35 - v749
	if v743 <= v757 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	goto L101
L103:
	;
	if v35 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v681 = v665
	v684 = v665
	goto L107
L105:
	;
	v718 = v665
	goto L106
L106:
	;
	v723 = v718 << (uint(int32(2)) % 32)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v723+l2)))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v723+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v664+v723))) = v726 - v728 + int32(1)
	goto L102
L107:
	;
	v685 = int32(2)
	v686 = v681 << (uint(v685) % 32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v686+l2)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v686+l3)))
	v693 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v664+v686))) = v689 - v691 + v693
	v697 = v686 | int32(4)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v697+l2)))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v697+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v664+v697))) = v700 - v702 + v693
	v708 = v681 + v685
	v710 = v684 + v685
	if v710 != v35&int32(2147483646) {
		v681 = v708
		v684 = v710
		goto L107
	} else {
		goto L109
	}
L108:
	;
	if v35&int32(1) == int32(0) {
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v718 = v708
	goto L106
L111:
	;
	if v369 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L112:
	;
	v764 = v757
	v768 = v743
	goto L115
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	v771 = v764 << (uint(int32(2)) % 32)
	v772 = v742 + v771
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v594+v771)))
	v775 = int32(1)
	v776 = v774 - v775
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v776
	v779 = v764 + v775
	if v35 <= v779 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L114
L117:
	;
	v847 = int32(1)
	if int32(0) < v764 {
		v764 = v764 - v847
		v768 = v768 + v847
		goto L115
	} else {
		goto L126
	}
L118:
	;
	if v768&int32(1) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v785 = int32(2)
	v786 = v779 << (uint(v785) % 32)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v664+v786)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v594+v786)))
	v794 = v776 - (v788-int32(1))*v792
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v794
	v798 = v764 + v785
	v799 = v794
	goto L121
L120:
	;
	v798 = v779
	v799 = v776
	goto L121
L121:
	;
	if v768 == int32(0) {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v806 = v798
	v807 = v799
	goto L123
L123:
	;
	v812 = int32(2)
	v813 = v806 << (uint(v812) % 32)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v664+v813)))
	v816 = int32(1)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v594+v813)))
	v821 = v807 - (v815-v816)*v819
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v821
	v824 = v813 + int32(4)
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v664+v824)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v594+v824)))
	v832 = v821 - (v826-v816)*v830
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v832
	v835 = v806 + v812
	if v835 != v35 {
		v806 = v835
		v807 = v832
		goto L123
	} else {
		goto L125
	}
L124:
	;
	goto L117
L125:
	;
	goto L124
L126:
	;
	goto L116
L127:
	;
	base.MemoryFill(m, v25+int32(32), int32(0), v46)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v870 = v495 << (uint(int32(3)) % 32)
	if v474 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v873 = v367 + v870
	goto L132
L131:
	;
	v873 = int32(0)
	goto L132
L132:
	;
	if v474 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v880 = v474
	goto L135
L134:
	;
	v880 = (v870 + int32(23)) & int32(-8)
	goto L135
L135:
	;
	v884 = v358 + v880
	v885 = v35 - int32(1)
	v888 = v591
	v891 = v585
	v893 = v496
	goto L136
L136:
	;
	v908 = v25 - int32(-64) + v885<<(uint(int32(2))%32)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	if v909 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v1027 = v358
	goto L1
L138:
	;
	v910 = F_array_seek(m, v888, v891, v340, v909, l7, l8)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L7
	} else {
		goto L141
	}
L139:
	;
	v914 = v888
	v915 = v891
	goto L140
L140:
	;
	v917 = F_array_seek(m, v914, v915, v340, int32(1), l7, l8)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L7
	} else {
		goto L142
	}
L141:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	v914 = v910
	v915 = v912 + v891
	goto L140
L142:
	;
	v919 = v917 - v914
	if v919 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	base.MemoryCopy(m, v884, v914, v919)
	goto L145
L144:
	;
	goto L145
L145:
	;
	if v474 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v924 = int32(1) << (uint(v893&int32(7)) % 32)
	v926 = base.I32_div_s(v893, int32(8))
	v927 = v873 + v926
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	if v42 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	goto L148
L148:
	;
	v949 = int32(1)
	v956 = v25 + int32(32)
	v958 = v25 + int32(96)
	if v35 <= int32(0) {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v927))) = uint8(v944)
	goto L148
L150:
	;
	v934 = base.I32_div_s(v915, int32(8))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338+v934))))
	if int32(base.Ui32(v936)>>(uint(v915&int32(7))%32))&int32(1) != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v944 = v924 | v928
	goto L149
L153:
	;
	v942 = v924 | v928
	goto L155
L154:
	;
	v942 = v928 & (v924 ^ int32(-1))
	goto L155
L155:
	;
	v944 = v942
	goto L149
L156:
	;
	if v1012 != int32(-1) {
		v884 = v884 + v919
		v885 = v1012
		v888 = v919 + v914
		v891 = v915 + v949
		v893 = v893 + v949
		goto L136
	} else {
		goto L171
	}
L157:
	;
	v1012 = int32(-1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v964 = int32(1)
	v965 = v35 - v964
	v967 = v965 << (uint(int32(2)) % 32)
	v968 = v956 + v967
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v958+v967)))
	v974 = base.I32_rem_s(v969+v964, v973)
	*(*int32)(unsafe.Add(mBase, uint32(v968))) = v974
	if v965 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v1012 = v1002
	goto L156
L161:
	;
	v976 = v965
	v979 = v974
	goto L164
L162:
	;
	goto L163
L163:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	if v1000 != 0 {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	if v979 != 0 {
		v1002 = v976
		goto L160
	} else {
		goto L166
	}
L165:
	;
	goto L163
L166:
	;
	v981 = int32(1)
	v982 = v976 - v981
	v984 = v982 << (uint(int32(2)) % 32)
	v985 = v956 + v984
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v958+v984)))
	v991 = base.I32_rem_s(v986+v981, v990)
	*(*int32)(unsafe.Add(mBase, uint32(v985))) = v991
	if v982 != 0 {
		v976 = v982
		v979 = v991
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v1001 = int32(0)
	goto L170
L169:
	;
	v1001 = int32(-1)
	goto L170
L170:
	;
	v1002 = v1001
	goto L160
L171:
	;
	goto L137
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1016))) = int64(64)
	v1027 = v1016
	goto L1
}
func F_array_gt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < v2))
	}
}
func F_array_in_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	v8 = int32(0)
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_array_in_safe[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v15
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_array_in_safe[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v18
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v8)
	v26 = F_InputFunctionCallSafe(m, l0, l1, l2, l3, v12+int32(56), v12+int32(72))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int64(0)
	} else {
		if v26 == int32(0) {
			v33 = v12 + int32(40)
			F_initStringInfo(m, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l5
				F_appendStringInfo(m, v33, int32(_a_F_array_in_safe_0), v12+int32(32))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(19)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v47
					F_ThrowErrorData(m, v46)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
						F_pfree(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v109 = v9
							m.G0 = v12 + int32(80)
							return v109
						}
					}
				}
			}
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
			v55 = F_pg_detoast_datum(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int64(0)
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				if v57 != int32(1) {
					v62 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						if v62 == int32(0) {
							v109 = v9
							m.G0 = v12 + int32(80)
							return v109
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l5
								F_errmsg(m, int32(_a_F_array_in_safe_1), v12+int32(16))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_array_in_safe_2), int32(1083), int32(_a_F_array_in_safe_3))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										v109 = v9
										m.G0 = v12 + int32(80)
										return v109
									}
								}
							}
						}
					}
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
					v82 = F_pg_detoast_datum(m, v81)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int64(0)
					} else {
						v84 = F_array_contains_nulls(m, v82)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							if v84 != 0 {
								v88 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int64(0)
								} else {
									if v88 == int32(0) {
										v109 = v9
										m.G0 = v12 + int32(80)
										return v109
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = l5
											F_errmsg(m, int32(_a_F_array_in_safe_4), v12)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_array_in_safe_2), int32(1092), int32(_a_F_array_in_safe_3))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int64(0)
												} else {
													v109 = v9
													m.G0 = v12 + int32(80)
													return v109
												}
											}
										}
									}
								}
							} else {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v105)
								v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
								v109 = v107
								m.G0 = v12 + int32(80)
								return v109
							}
						}
					}
				}
			}
		}
	}
}
func F_array_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	v4 = F_array_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(24)
		} else {
			v10 = int32(40)
		}
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_array_lower(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
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
	var v53 int64
	_ = v53
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
			return int64(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int64(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v47 = v40
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v47 = v7 + v41<<(uint(int32(2))%32) + int32(16)
				}
				v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47+v27<<(uint(int32(2))%32)-int32(4)))))
				return v53
			}
		}
	}
}
func F_array_position(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v192 int64
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v192
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L72
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L68
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L63
	}
L5:
	;
	m.G0 = v15 + int32(16)
	goto L1
L6:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v192 = int64(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int64(0)
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(2) <= v29 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v29 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
	v192 = int64(0)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v37 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v25)+16))
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v56 == int32(3) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v40 = F_array_contains_nulls(m, v25)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v51 = int32(4)
	v52 = v50
	goto L15
L19:
	;
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v51 = v42 << (uint(int32(2)) % 32)
	v52 = int64(0)
	goto L15
L21:
	;
	goto L22
L22:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
	v192 = int64(0)
	goto L5
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v59 == int32(1) {
		goto L3
	} else {
		goto L26
	}
L24:
	;
	v63 = v55
	goto L25
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v65 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v63 = v62
	goto L25
L27:
	;
	v107 = v55 - int32(1)
	v109 = F_array_create_iterator(m, v25, int32(0), v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L39
	}
L28:
	;
	F_get_typlenbyvalalign(m, v53, v81+int32(4), v81+int32(6), v81+int32(7))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L34
	}
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v70 = F_MemoryContextAlloc(m, v68, int32(48))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v79 == v53 {
		v104 = v65
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v70
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v53 ^ int32(-1)
	v81 = v75
	goto L28
L33:
	;
	v81 = v65
	goto L28
L34:
	;
	v91 = F_lookup_type_cache(m, v53, int32(32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	if v93 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v53
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	F_fmgr_info_cxt(m, v97, v81+int32(20), v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v104 = v81
	goto L27
L38:
	;
	F_array_free_iterator(m, v109)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L9
	} else {
		goto L55
	}
L39:
	;
	v115 = F_array_iterate(m, v109, v15+int32(8), v15+int32(7))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	if v115 == int32(0) {
		v158 = v107
		v166 = v2
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v123 = v107
	goto L42
L42:
	;
	v134 = v123 + int32(1)
	if v134 < v63 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v158 = v134
	v166 = v2
	goto L38
L44:
	;
	v154 = F_array_iterate(m, v109, v15+int32(8), v15+int32(7))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L53
	}
L45:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if (v136|v37)&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v158 = v134
	v166 = int32(1)
	goto L38
L47:
	;
	if v136&v37 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v144 = F_FunctionCall2Coll(m, v104+int32(20), v23, v52, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L51
	}
L50:
	;
	goto L46
L51:
	;
	if v144 == int64(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	if v154 != 0 {
		v123 = v134
		goto L42
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v170 != v25 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_pfree(m, v25)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v166 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v176)
	v192 = int64(0)
	goto L5
L61:
	;
	goto L62
L62:
	;
	v192 = base.I64_extend_i32_s(v158)
	goto L5
L63:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v203 = F_format_type_be(m, v53)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v203
	F_errmsg(m, int32(_a_F_array_position_0), v15)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1420), int32(_a_F_array_position_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_array_position_3), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1386), int32(_a_F_array_position_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_array_position_4), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_array_position_1), int32(1357), int32(_a_F_array_position_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_shuffle(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if v11 <= int32(0) {
			v33 = v7
			return base.I64_extend_i32_u(v33)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			if v14 < int32(2) {
				v33 = v7
				return base.I64_extend_i32_u(v33)
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
							return int64(0)
						} else {
							v33 = v31
							return base.I64_extend_i32_u(v33)
						}
					} else {
						v23 = F_lookup_type_cache(m, v17, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
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
								return int64(0)
							} else {
								v33 = v31
								return base.I64_extend_i32_u(v33)
							}
						}
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
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
							return int64(0)
						} else {
							v33 = v31
							return base.I64_extend_i32_u(v33)
						}
					}
				}
			}
		}
	}
}
func F_array_sort_order(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = base.B2i32(v8 != int64(0))
		v11 = F_array_sort_internal(m, v4, v10, v10, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_array_sort_order_nulls_first(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = int64(0)
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v13 = F_array_sort_internal(m, v3, base.B2i32(v7 != v8), base.B2i32(v10 != v8), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v13)
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
	var v10 int64
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if int32(0) < v13 {
		if v8&int32(1) != 0 {
			return
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
			if v18 == int32(0) {
				v33 = v13
				v34 = v10
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
				v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
				v43 = F_array_set_element(m, v34, v35, v12+int32(12), v38, v39, v33, v40, v41, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v45))) = v43
					return
				}
			} else {
				return
			}
		}
	} else {
		if v8&int32(1) == int32(0) {
			v33 = v13
			v34 = v10
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
			v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
			v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
			v43 = F_array_set_element(m, v34, v35, v12+int32(12), v38, v39, v33, v40, v41, v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v45))) = v43
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
				v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
				v33 = v32
				v34 = base.I64_extend_i32_u(v26)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
				v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
				v43 = F_array_set_element(m, v34, v35, v12+int32(12), v38, v39, v33, v40, v41, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v45))) = v43
					return
				}
			}
		}
	}
}
func F_array_subscript_fetch_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+6)))
	v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+9)))
	v22 = F_array_get_slice(m, v14, v8, v9+int32(12), v9+int32(36), v17, v18, v19, v20, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int64)(unsafe.Add(mBase, uint32(v24))) = v22
		return
	}
}
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 float32
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v94 float64
	_ = v94
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v124 float64
	_ = v124
	v6 = float64(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(80)
	return v124
L2:
	;
	v124 = float64(10)
	goto L1
L3:
	;
	v14 = l1
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
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+20)))
	if v44 != 0 {
		goto L4
	} else {
		goto L18
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	switch v19 - int32(7) {
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)))
	if v32 != 0 {
		v124 = v6
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v31 != 0 {
		v14 = v31
		goto L6
	} else {
		goto L13
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 != int32(27) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 != int32(34) {
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return float64(0)
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v41 = F_ArrayGetNItemsSafe(m, v38, v34+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v124 = base.F64_convert_i32_s(v41)
	goto L1
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v45 == int32(0) {
		v124 = v6
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v124 = base.F64_convert_i32_s(v48)
	goto L1
L20:
	;
	if v19 == int32(6) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v55 == int32(0) {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_examine_variable(m, l0, v14, int32(0), v9+int32(48))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	if v63 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v71 = F_get_attstatsslot(m, v9+int32(12), v63, int32(5), int32(0), int32(2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	if v71 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v73 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v105 = v6
	goto L30
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	if v106 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v99 = float64(0)
	goto L33
L32:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v77+v73<<(uint(int32(2))%32)-int32(4))))
	v84 = base.F64_promote_f32(v83)
	v85 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)))|base.F64_gt(v84, v85) != 0 {
		v98 = v85
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_free_attstatsslot(m, v9+int32(12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L38
	}
L34:
	;
	v99 = v98
	goto L33
L35:
	;
	goto L34
L36:
	;
	v94 = float64(1)
	if base.F64_le(v84, v94) != 0 {
		v98 = v94
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v98 = base.F64_nearest(v84)
	goto L35
L38:
	;
	v105 = v99
	goto L30
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	m.T0[v107].(func(*base.Module, int32))(m, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.F64_gt(v105, float64(0)) != 0 {
		v124 = v105
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	goto L2
}
func F_getArrayIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
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
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_getArrayIndex[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v17
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_getArrayIndex[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v20
	v23 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v23
	v25 = int32(2)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v27 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v23, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L43
	}
L2:
	;
	m.G0 = v10 + int32(96)
	return v153
L3:
	;
	return int32(0)
L4:
	;
	if v27 == int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v33 == int32(0) {
		v153 = v25
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v46 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v38 = v33
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	F_pfree(m, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v153 = v25
	goto L2
L11:
	;
	if v43 != 0 {
		v38 = v43
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v153 = int32(2)
	goto L2
L14:
	;
	v96 = int32(0)
	v99 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+40)))
	v101 = F_DirectFunctionCall2Coll(m, int32(1574), v96, v99, int64(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L31
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v49 == int32(2) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v55 = v52
	goto L22
L20:
	;
	goto L21
L21:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v10 + int32(16)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v77 != int32(1) {
		goto L13
	} else {
		goto L26
	}
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	F_pfree(m, v55)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	if v60 != 0 {
		v55 = v60
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_getArrayIndex_0), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3753), int32(_a_F_getArrayIndex_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v104 = F_pg_detoast_datum(m, base.I32_wrap_i64(v101))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v106 = F_numeric_int4_safe(m, v104, v10)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v109 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = v109
	goto L37
L35:
	;
	goto L36
L36:
	;
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v127
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v10 + int32(16)
	if v131 != int32(1) {
		v153 = v96
		goto L2
	} else {
		goto L41
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	F_pfree(m, v112)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	if v117 != 0 {
		v112 = v117
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v137 == int32(1) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L13
L43:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_getArrayIndex_3), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_getArrayIndex_1), int32(3768), int32(_a_F_getArrayIndex_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
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
func F_makeArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	return int64(0)
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
	return base.I64_extend_i32_u(v178)
L50:
	;
	goto L49
}
