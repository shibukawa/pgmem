package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateDecodingContext(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v1 = l0
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[0]))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L23
	} else {
		goto L66
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L23
	} else {
		goto L62
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L23
	} else {
		goto L58
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L23
	} else {
		goto L55
	}
L7:
	;
	v23 = v18 + int32(24)
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[1]))
	if v19 != v27 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	if v1 == int64(0) {
		v83 = v51
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v41 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])) = uint8(v39)
	v41 = v39
	goto L16
L15:
	;
	v41 = int32(0)
	goto L16
L16:
	;
	goto L13
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+201)))
	if v44 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[6])))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v84 = int32(0)
	v87 = F_StartupDecodingContext(m, l1, v83, v84, v84, l2, v84, l3, l4, l5, l6)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L30
	}
L21:
	;
	if base.Ui64(v51) <= base.Ui64(v1) {
		v83 = v1
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v57 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v61)
	v63 = int64(32)
	v64 = int64(base.Ui64(v61) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v64)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v1)
	v68 = int64(base.Ui64(v1) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v68)
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_9), v15+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v83 = v81
	goto L20
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(574), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v89 = int32(_a_F_CreateDecodingContext_4)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(_a_F_CreateDecodingContext_5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(992)
	v102 = int32(_a_F_CreateDecodingContext_6)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v15 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v15 + int32(80)
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+164)) = uint8(v112)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+147)) = uint8(v112)
	m.T0[v94].(func(*base.Module, int32, int32, int32))(m, v87, v87+int32(108), v112)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v90
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v127 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v122
	goto L33
L35:
	;
	v130 = int32(1)
	goto L37
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+146)))
	v130 = v129
	goto L37
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+145)))
	v132 = v130 & v131
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+145)) = uint8(v132)
	if v132 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+116)) = uint8(v158)
	v162 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L23
	} else {
		goto L48
	}
L39:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v136 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_s_lock(m, v18, int32(_a_F_CreateDecodingContext_1), int32(601), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v83
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v83
	goto L47
L47:
	;
	goto L38
L48:
	;
	if v162 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_7), v15+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	m.G0 = v15 + int32(96)
	return v87
L52:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v171)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v170)
	v174 = int64(32)
	v175 = int64(base.Ui64(v171) >> (uint(v174) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v175)
	v178 = int64(base.Ui64(v170) >> (uint(v174) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v178)
	F_errdetail(m, int32(_a_F_CreateDecodingContext_8), v15)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(617), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_10), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(517), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
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
	F_errcode(m, int32(325))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_CreateDecodingContext_0), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(523), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L23
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
	F_errcode(m, int32(325))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_3), v15-int32(-64))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(534), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_11), v15+int32(48))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	F_errdetail(m, int32(_a_F_CreateDecodingContext_12), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	F_errhint(m, int32(_a_F_CreateDecodingContext_13), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(547), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v7 = m.G0
	v9 = v7 - int32(1136)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_CreateDirAndVersionFile_0)
	v18 = F_pg_sprintf(m, v9+int32(96), int32(_a_F_CreateDirAndVersionFile_1), v9+int32(80))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[0]))
	v22 = F_mkdir(m, l0, v21)
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L60
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L56
	}
L5:
	;
	if v22 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(_a_F_CreateDirAndVersionFile_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	v40 = F_pg_snprintf(m, v9+int32(112), int32(1024), int32(_a_F_CreateDirAndVersionFile_3), v9+int32(48))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v28 != int32(20) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v45 = F_OpenTransientFile(m, v9+int32(112), int32(193))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v45 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l3 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v62 = v45
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(167772226)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(0)
	v72 = int32(3)
	v73 = F_write(m, v62, v9+int32(96), v72)
	mBase = m.M
	if v73 != v72 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v52 != int32(20) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v58 = F_OpenTransientFile(m, v9+int32(112), int32(513))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v58 < int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v62 = v58
	goto L15
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v77 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v102 = int32(_a_F_CreateDirAndVersionFile_4)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(167772225)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[3])))
	if v112 != int32(1) {
		v126 = v104
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(51)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_5), v9+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(507), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	F_fsync_fname(m, l0, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L48
	}
L31:
	;
	if v126 == int32(0) {
		goto L30
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	goto L34
L34:
	;
	v117 = F_fsync(m, v62)
	mBase = m.M
	if v117 != int32(-1) {
		v126 = v117
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(-1)
	goto L32
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v121 == int32(27) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[4])))
	if v132 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v135 = F_errstart(m, v133, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v133 = int32(21)
	goto L42
L41:
	;
	v133 = int32(23)
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v135 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_8), v9+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(515), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L30
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = int32(0)
	v161 = F_CloseTransientFile(m, v62)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if l3 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v165 = int32(_a_F_CreateDirAndVersionFile_9)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v167 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	m.G0 = v9 + int32(1136)
	return
L53:
	;
	F_XLogRegisterData(m, v9+int32(88), int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v182 = F_XLogInsert(m, int32(4), int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v184 = int32(_a_F_CreateDirAndVersionFile_9)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v186 - int32(1)
	goto L52
L56:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_10), v9-int32(-64))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(478), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_11), v9)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(495), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v3 = l2
	v4 = l3
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
	v14 = F_MemoryContextAlloc(m, v12, int32(28))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v10 {
				v39 = v17
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
				v49 = v39 + int32(8)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v50 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
				v61 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
				v23 = F_MemoryContextAlloc(m, v21, int32(24))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
					v29 = v23 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
					v32 = int32(_a_F_create_drop_transactional_internal_0)
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
					*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
					v39 = v23
					*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
					v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
					v49 = v39 + int32(8)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					if v50 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
					v61 = v14 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
			v23 = F_MemoryContextAlloc(m, v21, int32(24))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
				v29 = v23 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
				v32 = int32(_a_F_create_drop_transactional_internal_0)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
				*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
				v39 = v23
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
				v49 = v39 + int32(8)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v50 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
				v61 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
				return
			}
		}
	}
}
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = int64(0)
	v6 = F_create_plan_recurse(m, l0, l1, int32(1))
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v10 != int32(333) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v21 = int32(0)
	goto L7
L4:
	;
	goto L5
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L5
L7:
	;
	v23 = int32(0)
	if v13 == v23 {
		v33 = v23
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v14 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= v21 {
		v33 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = v29 + v21<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v51
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+24)) = uint16(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+26)) = uint8(v55)
	v21 = v21 + int32(1)
	goto L7
L13:
	;
	goto L6
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v36 <= v21 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v33 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v43 = v40 + v21<<(uint(int32(2))%32)
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	return v6
L21:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_0), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_create_plan_1), int32(372), int32(_a_F_create_plan_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_syncrep_config(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
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
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	v3 = l2
	v4 = int32(0)
	v8 = int32(16)
	if l1 == v4 {
		v98 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v100 = F_palloc(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		v98 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v11
	goto L6
L5:
	;
	v17 = v14
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = v4
	v24 = v8
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v22<<(uint(int32(2))%32))))
	if v29&int32(3) == int32(0) {
		v53 = v29
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v98 = v89
	goto L1
L9:
	;
	v88 = int32(1)
	v89 = v86 + v24 + v88
	v91 = v22 + v88
	if v91 != v17 {
		v22 = v91
		v24 = v89
		goto L7
	} else {
		goto L26
	}
L10:
	;
	v86 = v78 - v29
	goto L9
L11:
	;
	v57 = v53
	goto L20
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v86 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v42 = v29
	goto L16
L16:
	;
	v46 = v42 + int32(1)
	if v46&int32(3) == int32(0) {
		v53 = v46
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v78 = v46
	goto L10
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v51 != 0 {
		v42 = v46
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 == v66 {
		v57 = v57 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v72 = v57
	goto L23
L22:
	;
	goto L21
L23:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 != 0 {
		v72 = v72 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v78 = v72
	goto L10
L25:
	;
	goto L24
L26:
	;
	goto L8
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v98
	v108 = l0
	goto L30
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+8)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v152
	if l1 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v113 = v108 + int32(1)
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v108))))
	v115 = F___isspace(m, v114)
	mBase = m.M
	if v115 != 0 {
		v108 = v113
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v116 = int32(1)
	switch v114&int32(255) - int32(43) {
	case 0:
		v122 = v116
		goto L34
	default:
		v124 = v114
		v125 = v108
		v126 = v116
		goto L33
	case 2:
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v127 = int32(0)
	v129 = v124 - int32(48)
	if base.Ui32(v129) <= base.Ui32(int32(9)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113))))
	v124 = v123
	v125 = v113
	v126 = v122
	goto L33
L35:
	;
	v122 = int32(0)
	goto L34
L36:
	;
	v132 = v127
	v133 = v129
	v134 = v125
	goto L39
L37:
	;
	v146 = v127
	goto L38
L38:
	;
	if v126 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v136 = int32(10)
	v138 = v132*v136 - v133
	v139 = int32(*(*int8)(unsafe.Add(mBase, uint32(v134)+1)))
	v143 = v139 - int32(48)
	if base.Ui32(v143) < base.Ui32(v136) {
		v132 = v138
		v133 = v143
		v134 = v134 + int32(1)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v146 = v138
	goto L38
L41:
	;
	goto L40
L42:
	;
	v152 = int32(0) - v146
	goto L44
L43:
	;
	v152 = v146
	goto L44
L44:
	;
	goto L29
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = int32(0)
	return v100
L46:
	;
	goto L47
L47:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v162 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v171 = int32(0)
	v173 = v100 + int32(16)
	goto L51
L49:
	;
	goto L50
L50:
	;
	return v100
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32))))
	if (v179^v173)&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	goto L50
L53:
	;
	if v179&int32(3) == int32(0) {
		v277 = v179
		goto L76
	} else {
		goto L77
	}
L54:
	;
	goto L53
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v233)
	if v233&int32(255) == int32(0) {
		goto L54
	} else {
		goto L70
	}
L56:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v232 = v179
	v233 = v185
	v234 = v173
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v179&int32(3) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v189 = v179
	v191 = v173
	goto L62
L60:
	;
	v203 = v179
	v205 = v173
	goto L61
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v210 = int32(-2139062144)
	if (int32(16843008)-v207|v207)&v210 != v210 {
		v232 = v203
		v233 = v207
		v234 = v205
		goto L55
	} else {
		goto L66
	}
L62:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
	if v192 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L63:
	;
	v203 = v199
	v205 = v197
	goto L61
L64:
	;
	v196 = int32(1)
	v197 = v191 + v196
	v199 = v189 + v196
	if v199&int32(3) != 0 {
		v189 = v199
		v191 = v197
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v215 = v203
	v216 = v207
	v217 = v205
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v216
	v219 = int32(4)
	v220 = v217 + v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v223 = v215 + v219
	v227 = int32(-2139062144)
	if (v221|(int32(16843008)-v221))&v227 == v227 {
		v215 = v223
		v216 = v221
		v217 = v220
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v232 = v223
	v233 = v221
	v234 = v220
	goto L55
L69:
	;
	goto L68
L70:
	;
	v241 = v232
	v243 = v234
	goto L71
L71:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)) = uint8(v244)
	v246 = int32(1)
	if v244 != 0 {
		v241 = v241 + v246
		v243 = v243 + v246
		goto L71
	} else {
		goto L73
	}
L72:
	;
	goto L54
L73:
	;
	goto L72
L74:
	;
	v312 = int32(1)
	v315 = v171 + v312
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v315 < v316 {
		v171 = v315
		v173 = v173 + v310 + v312
		goto L51
	} else {
		goto L91
	}
L75:
	;
	v310 = v302 - v179
	goto L74
L76:
	;
	v281 = v277
	goto L85
L77:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v261 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v310 = int32(0)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v266 = v179
	goto L81
L81:
	;
	v270 = v266 + int32(1)
	if v270&int32(3) == int32(0) {
		v277 = v270
		goto L76
	} else {
		goto L83
	}
L82:
	;
	v302 = v270
	goto L75
L83:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v275 != 0 {
		v266 = v270
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v290 = int32(-2139062144)
	if (int32(16843008)-v287|v287)&v290 == v290 {
		v281 = v281 + int32(4)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v296 = v281
	goto L88
L87:
	;
	goto L86
L88:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v300 != 0 {
		v296 = v296 + int32(1)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v302 = v296
	goto L75
L90:
	;
	goto L89
L91:
	;
	goto L52
}
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	v3 = l2
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v12
		v90 = v11
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
		if v96 != 0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
			*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if v105 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			if v112 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v116 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
			v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
			if v123 < int32(0) {
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v128 = v126 - int32(97)
				if base.Ui32(int32(17)) < base.Ui32(v128) {
				} else {
					if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						if v137 != 0 {
						} else {
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
							v145 = v140 + v123*int32(24) + int32(12)
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
							if v146 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
								v149 = v148
							} else {
								v149 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
						}
					}
				}
			}
		}
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if base.Ui32(v16) <= base.Ui32(v15) {
				v34 = l0 + int32(76)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+136))
				if base.Ui32(v36) < base.Ui32(int32(98000000)) {
					v53 = v16 << (uint(int32(1)) % 32)
					v54 = v34
					v57 = int32(1024)
					if base.Ui32(v57) <= base.Ui32(v53) {
						v60 = v57
					} else {
						v60 = v53
					}
					v64 = v60*int32(40) | int32(8)
					v66 = F_palloc_extended(m, v64, int32(2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						if v66 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							if v74 != 0 {
								v76 = v74
							} else {
								v76 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v76
							v90 = int32(0)
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+136)) = v79 + v64
							*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v60
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66))) = v83
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v66
							v90 = v66 + int32(8)
						}
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
						if v96 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
							*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if v105 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
							v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							if v112 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v116 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
							v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
							if v123 < int32(0) {
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v128 = v126 - int32(97)
								if base.Ui32(int32(17)) < base.Ui32(v128) {
								} else {
									if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
										if v137 != 0 {
										} else {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
											v145 = v140 + v123*int32(24) + int32(12)
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											if v146 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
												v149 = v148
											} else {
												v149 = int32(0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
											*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
										}
									}
								}
							}
						}
						return
					}
				} else {
					v40 = v34
					v41 = v35
					*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
					if v46 != 0 {
						v48 = v46
					} else {
						v48 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v48
					v90 = int32(0)
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
					if v96 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v105 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v112 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v116 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
						v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
						if v123 < int32(0) {
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							v128 = v126 - int32(97)
							if base.Ui32(int32(17)) < base.Ui32(v128) {
							} else {
								if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v137 != 0 {
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
										v145 = v140 + v123*int32(24) + int32(12)
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										if v146 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											v149 = v148
										} else {
											v149 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
									}
								}
							}
						}
					}
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				v90 = v14 + v15*int32(40) + int32(8)
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
				if v96 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v105 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v112 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v116 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
					v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
					if v123 < int32(0) {
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v128 = v126 - int32(97)
						if base.Ui32(int32(17)) < base.Ui32(v128) {
						} else {
							if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v137 != 0 {
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
									v145 = v140 + v123*int32(24) + int32(12)
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									if v146 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										v149 = v148
									} else {
										v149 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
									*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
								}
							}
						}
					}
				}
				return
			}
		} else {
			v27 = l0 + int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
			if base.Ui32(int32(97999999)) < base.Ui32(v30) {
				v40 = v27
				v41 = v29
				*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
				if v46 != 0 {
					v48 = v46
				} else {
					v48 = int32(19)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v48
				v90 = int32(0)
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
				if v96 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v105 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v112 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v116 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
					v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
					if v123 < int32(0) {
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v128 = v126 - int32(97)
						if base.Ui32(int32(17)) < base.Ui32(v128) {
						} else {
							if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v137 != 0 {
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
									v145 = v140 + v123*int32(24) + int32(12)
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									if v146 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										v149 = v148
									} else {
										v149 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
									*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
								}
							}
						}
					}
				}
				return
			} else {
				v53 = int32(64)
				v54 = v27
				v57 = int32(1024)
				if base.Ui32(v57) <= base.Ui32(v53) {
					v60 = v57
				} else {
					v60 = v53
				}
				v64 = v60*int32(40) | int32(8)
				v66 = F_palloc_extended(m, v64, int32(2))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					if v66 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
						if v74 != 0 {
							v76 = v74
						} else {
							v76 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v76
						v90 = int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+136))
						*(*int32)(unsafe.Add(mBase, uint32(v68)+136)) = v79 + v64
						*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v60
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66))) = v83
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v66
						v90 = v66 + int32(8)
					}
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
					if v96 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v105 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v112 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v116 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
						v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
						if v123 < int32(0) {
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							v128 = v126 - int32(97)
							if base.Ui32(int32(17)) < base.Ui32(v128) {
							} else {
								if int32(1)<<(uint(v128)%32)&int32(_a_F_createarc_0) == int32(0) {
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v137 != 0 {
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
										v145 = v140 + v123*int32(24) + int32(12)
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										if v146 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											v149 = v148
										} else {
											v149 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
									}
								}
							}
						}
					}
					return
				}
			}
		}
	}
}
