package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckValidResultRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
	switch v13 - int32(83) {
	case 0:
		goto L10
	default:
		goto L5
	case 19:
		goto L6
	case 26:
		goto L7
	case 29, 31:
		goto L11
	case 33:
		goto L9
	case 35:
		goto L8
	}
L1:
	;
	m.G0 = v9 + int32(176)
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L20
	} else {
		goto L85
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L20
	} else {
		goto L81
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L20
	} else {
		goto L77
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L20
	} else {
		goto L73
	}
L6:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	switch l1 - int32(2) {
	case 0:
		goto L44
	case 1:
		goto L45
	case 2:
		goto L43
	default:
		goto L42
	}
L7:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[500]))
	if int32(0) < v103 {
		goto L1
	} else {
		goto L37
	}
L8:
	;
	v97 = F_view_has_instead_trigger(m, v11, l1, l3)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L20
	} else {
		goto L34
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L30
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L26
	}
L11:
	;
	if l1 == int32(5) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if l2 != int32(2) {
		goto L1
	} else {
		goto L24
	}
L13:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_CheckCmdReplicaIdentity(m, v11, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L23
	}
L16:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v20 <= int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v24 = int32(0)
	goto L18
L18:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v24<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	F_CheckCmdReplicaIdentity(m, v11, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L12
L20:
	;
	return
L21:
	;
	v39 = v24 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v39 < v40 {
		v24 = v39
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L12
L24:
	;
	F_CheckCmdReplicaIdentity(m, v11, int32(2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L1
L26:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v62 + int32(4)
	F_errmsg(m, int32(749554), v9+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(516339), int32(1090), int32(320498))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v83 + int32(4)
	F_errmsg(m, int32(736783), v9+int32(32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(516339), int32(1096), int32(320498))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L20
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
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_error_view_not_updatable(m, v11, l1, l3, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	goto L1
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v113 + int32(4)
	F_errmsg(m, int32(725118), v9+int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(516339), int32(1115), int32(320498))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L20
	} else {
		goto L70
	}
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v127)+68))
	if v192 == int32(0) {
		goto L2
	} else {
		goto L62
	}
L44:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v127)+64))
	if v161 == int32(0) {
		goto L3
	} else {
		goto L54
	}
L45:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+52))
	if v130 == int32(0) {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)+84))
	if v133 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v136 = m.T0[v133].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	if v136&int32(8) != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v147 + int32(4)
	F_errmsg(m, int32(124873), v9+int32(96))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(516339), int32(1133), int32(320498))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L20
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
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v127)+84))
	if v164 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v167 = m.T0[v164].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	if v167&int32(4) != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v178 + int32(4)
	F_errmsg(m, int32(168783), v9+int32(128))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(516339), int32(1146), int32(320498))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L20
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
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v127)+84))
	if v195 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v198 = m.T0[v195].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	if v198&int32(16) != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v209 + int32(4)
	F_errmsg(m, int32(168400), v9+int32(160))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(516339), int32(1159), int32(320498))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l1
	F_errmsg_internal(m, int32(504646), v9-int32(-64))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(516339), int32(1162), int32(320498))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v245 + int32(4)
	F_errmsg(m, int32(736512), v9)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(516339), int32(1170), int32(320498))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L20
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v264 + int32(4)
	F_errmsg(m, int32(747020), v9+int32(80))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(516339), int32(1127), int32(320498))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v285 + int32(4)
	F_errmsg(m, int32(747200), v9+int32(112))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(516339), int32(1140), int32(320498))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L20
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L20
	} else {
		goto L86
	}
L86:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v306 + int32(4)
	F_errmsg(m, int32(747133), v9+int32(144))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(516339), int32(1153), int32(320498))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ValidJsonBehaviorDefaultExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = int32(0)
	if l0 == v3 {
		v17 = v3
		return v17
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(7) {
		case 0, 8, 10:
			v17 = int32(1)
			return v17
		default:
			v17 = int32(0)
			return v17
		case 20, 21, 22, 23, 24, 48:
			v11 = F_expression_tree_walker_impl(m, l0, int32(487), l1)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_ValidXLogRecordHeader(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int64
	_ = v34
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	v2 = l1
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui32(v13) <= base.Ui32(int32(23)) {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(24)
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v2)
		v21 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v21)
		F_report_invalid_record(m, l0, int32(42821), v11)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v77 = int32(0)
			m.G0 = v11 - int32(-64)
			return v77
		}
	} else {
		v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+17)))
		if int32(22) <= v28 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v28
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+56)) = uint32(v2)
			v34 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+52)) = uint32(v34)
			F_report_invalid_record(m, l0, int32(532836), v9+int32(-16))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v77 = int32(0)
				m.G0 = v11 - int32(-64)
				return v77
			}
		} else {
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
			if l4 != 0 {
				if base.Ui64(v41) < base.Ui64(v2) {
					v77 = int32(1)
					m.G0 = v11 - int32(-64)
					return v77
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v2)
					v45 = int64(32)
					v46 = int64(base.Ui64(v2) >> (uint(v45) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v46)
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v41)
					v50 = int64(base.Ui64(v41) >> (uint(v45) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v50)
					F_report_invalid_record(m, l0, int32(534719), v9+int32(-48))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v77 = int32(0)
						m.G0 = v11 - int32(-64)
						return v77
					}
				}
			} else {
				if l2 == v41 {
					v77 = int32(1)
					m.G0 = v11 - int32(-64)
					return v77
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+44)) = uint32(v2)
					v60 = int64(32)
					v61 = int64(base.Ui64(v2) >> (uint(v60) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+40)) = uint32(v61)
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+36)) = uint32(v41)
					v65 = int64(base.Ui64(v41) >> (uint(v60) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)) = uint32(v65)
					F_report_invalid_record(m, l0, int32(534719), v9+int32(-32))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v77 = int32(0)
						m.G0 = v11 - int32(-64)
						return v77
					}
				}
			}
		}
	}
}
func F_read_valid_char(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v150 int32
	_ = v150
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = v7
	goto L2
L1:
	;
	return base.I32_extend8_s(v150)
L2:
	;
	v14 = l0 + v13
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14))))
	v17 = v15 & int32(255)
	if int32(0) <= v15 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v150 = int32(92)
	goto L1
L4:
	;
	goto L3
L5:
	;
	if v131&int32(255) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14))))
	if int32(0) <= v79 {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v75 = v17
	goto L9
L8:
	;
	if v17&int32(224) == int32(192) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v75 != 0 {
		v78 = v75
		goto L6
	} else {
		goto L18
	}
L10:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v14))))
	v75 = v67 | v69&int32(63)
	goto L9
L11:
	;
	v66 = int32(1)
	v67 = v17 << (uint(int32(6)) % 32) & int32(1984)
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v17&int32(240) == int32(224) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v66 = int32(2)
	v67 = v17<<(uint(int32(12))%32)&int32(61440) | v38&int32(63)<<(uint(int32(6))%32)
	goto L10
L15:
	;
	goto L16
L16:
	;
	if v17&int32(248) != int32(240) {
		v78 = int32(-1)
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v55 = int32(63)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	v66 = int32(3)
	v67 = v17<<(uint(int32(18))%32)&int32(1835008) | v54&v55<<(uint(int32(12))%32) | v60&v55<<(uint(int32(6))%32)
	goto L10
L18:
	;
	v131 = int32(0)
	v134 = v13
	goto L5
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v105 = v103 + v104
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v105
	v107 = int32(26)
	if base.Ui32(v78-int32(91)) < base.Ui32(int32(3)) {
		v131 = v107
		v134 = v105
		goto L5
	} else {
		goto L32
	}
L20:
	;
	v103 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v84 = v79 & int32(255)
	if v84&int32(224) == int32(192) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v103 = int32(2)
	goto L19
L24:
	;
	goto L25
L25:
	;
	if v84&int32(240) == int32(224) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v103 = int32(3)
	goto L19
L27:
	;
	goto L28
L28:
	;
	if v84&int32(248) == int32(240) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = int32(4)
	goto L31
L30:
	;
	v101 = int32(1)
	goto L31
L31:
	;
	v103 = v101
	goto L19
L32:
	;
	if base.Ui32(v78) <= base.Ui32(int32(95)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v131 = v78
	v134 = v105
	goto L5
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(v78) <= base.Ui32(int32(255)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_consts[1566]))))
	v131 = v118
	v134 = v105
	goto L5
L37:
	;
	goto L38
L38:
	;
	switch v78 - int32(260) {
	case 0, 1:
		v150 = int32(91)
		goto L1
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		v131 = v107
		v134 = v105
		goto L5
	case 20, 21:
		goto L4
	default:
		goto L39
	}
L39:
	;
	v122 = int32(93)
	if base.Ui32(v78-int32(354)) < base.Ui32(int32(2)) {
		v150 = v122
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v78-int32(538)) < base.Ui32(int32(2)) {
		v150 = v122
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v131 = v107
	v134 = v105
	goto L5
L42:
	;
	return base.I32_extend8_s(v131)
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(28)) < base.Ui32((v131-int32(65))&int32(255)) {
		v13 = v134
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v150 = v131
	goto L1
}
