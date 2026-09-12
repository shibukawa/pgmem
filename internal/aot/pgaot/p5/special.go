package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpecialHyphen(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = v4 - v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9 - v10
	return
}
func F_resolve_special_varno(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != int32(6) {
		v312 = l0
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(144)
	return
L4:
	;
	m.T0[l2].(func(*base.Module, int32, int32, int32))(m, v312, l1, l3)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L108
	}
L5:
	;
	v18 = l0
	goto L10
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L105
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L102
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L99
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L96
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v33 != int32(-3) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v33 <= int32(0) {
		goto L9
	} else {
		goto L93
	}
L12:
	;
	goto L11
L13:
	;
	switch v33 + int32(2) {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		goto L12
	}
L14:
	;
	goto L15
L15:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	if v190 == int32(0) {
		goto L9
	} else {
		goto L76
	}
L16:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	if v120 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	if v38 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v79 == int32(0) {
		goto L8
	} else {
		goto L32
	}
L20:
	;
	goto L19
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v45 <= int32(0) {
		v79 = int32(0)
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = int32(0)
	goto L20
L24:
	;
	v48 = int32(0)
	if v48 < v45 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v51 = v45
	goto L27
L26:
	;
	v51 = v48
	goto L27
L27:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v56 = int32(0)
	goto L28
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52+v56<<(uint(int32(2))%32))))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+8)))
	if v65 == v41&int32(65535) {
		v79 = v64
		goto L20
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v68 = v56 + int32(1)
	if v68 != v51 {
		v56 = v68
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v85&int32(-2) == int32(334) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+72))
	v91 = F_bms_union(m, v83, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	goto L38
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v91
	goto L35
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v102 = F_lcons(m, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	v98 = F__emscripten_memcpy_bulkmem(m, v11-int32(-64), v32, int32(80))
	mBase = m.M
	goto L40
L40:
	;
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v102
	F_set_deparse_plan(m, v32, v94)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	F_resolve_special_varno(m, v107, l1, l2, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v111 = F_list_delete_first(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L46
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+44)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v83
	goto L3
L46:
	;
	v116 = F__emscripten_memcpy_bulkmem(m, v32, v11-int32(-64), int32(80))
	mBase = m.M
	goto L48
L48:
	;
	goto L45
L49:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	if v120 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v161 == int32(0) {
		goto L7
	} else {
		goto L63
	}
L51:
	;
	goto L50
L52:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v127 <= int32(0) {
		v161 = int32(0)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v161 = int32(0)
	goto L51
L55:
	;
	v130 = int32(0)
	if v130 < v127 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v133 = v127
	goto L58
L57:
	;
	v133 = v130
	goto L58
L58:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v138 = int32(0)
	goto L59
L59:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v134+v138<<(uint(int32(2))%32))))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+8)))
	if v147 == v123&int32(65535) {
		v161 = v146
		goto L51
	} else {
		goto L61
	}
L60:
	;
	goto L54
L61:
	;
	v150 = v138 + int32(1)
	if v150 != v133 {
		v138 = v150
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	goto L65
L64:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v173 = F_lcons(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v169 = F__emscripten_memcpy_bulkmem(m, v11-int32(-64), v32, int32(80))
	mBase = m.M
	goto L67
L67:
	;
	goto L64
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v173
	F_set_deparse_plan(m, v32, v165)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	F_resolve_special_varno(m, v178, l1, l2, l3)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v182 = F_list_delete_first(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L73
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+44)) = v182
	goto L3
L73:
	;
	v187 = F__emscripten_memcpy_bulkmem(m, v32, v11-int32(-64), int32(80))
	mBase = m.M
	goto L75
L75:
	;
	goto L72
L76:
	;
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	if v190 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v231 == int32(0) {
		goto L6
	} else {
		goto L90
	}
L78:
	;
	goto L77
L79:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v197 <= int32(0) {
		v231 = int32(0)
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v231 = int32(0)
	goto L78
L82:
	;
	v200 = int32(0)
	if v200 < v197 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v203 = v197
	goto L85
L84:
	;
	v203 = v200
	goto L85
L85:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v208 = int32(0)
	goto L86
L86:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v204+v208<<(uint(int32(2))%32))))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
	if v217 == v193&int32(65535) {
		v231 = v216
		goto L78
	} else {
		goto L88
	}
L87:
	;
	goto L81
L88:
	;
	v220 = v208 + int32(1)
	if v220 != v203 {
		v208 = v220
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v238 == int32(6) {
		v18 = v235
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v312 = v235
	goto L4
L93:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v243 == int32(0) {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v33 <= v246 {
		v312 = v18
		goto L4
	} else {
		goto L95
	}
L95:
	;
	goto L9
L96:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v254
	F_errmsg_internal(m, int32(483959), v11)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(495103), int32(7993), int32(241388))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v268
	F_errmsg_internal(m, int32(483772), v11+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(495103), int32(7942), int32(241388))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v284
	F_errmsg_internal(m, int32(483809), v11+int32(32))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(495103), int32(7972), int32(241388))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v300
	F_errmsg_internal(m, int32(483735), v11+int32(48))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(495103), int32(7986), int32(241388))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	goto L3
}
