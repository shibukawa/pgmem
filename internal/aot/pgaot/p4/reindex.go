package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexIsProcessingIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v5 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v5 != l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v9 = int32(0)
	if v8 == v9 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v48 = int32(1)
	goto L3
L3:
	;
	return v48
L4:
	;
	v48 = v47
	goto L3
L5:
	;
	v47 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v15 <= int32(0) {
		v40 = v9
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = v40
	goto L4
L9:
	;
	v18 = int32(0)
	if v18 < v15 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v21 = v15
	goto L12
L11:
	;
	v21 = v18
	goto L12
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v24 = int32(0)
	goto L13
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22+v24<<(uint(int32(2))%32))))
	v33 = base.B2i32(v32 == l0)
	if v32 == l0 {
		v40 = v33
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v40 = v33
	goto L8
L15:
	;
	v35 = v24 + int32(1)
	if v35 != v21 {
		v24 = v35
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_reindex_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v314 int32
	_ = v314
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18&int32(4) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v22 = F_try_table_open(m, l1, int32(5))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v27 = F_table_open(m, l1, int32(5))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v29 = v22
	goto L1
L7:
	;
	v29 = v27
	goto L1
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
	if v31 != int32(112) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v314 = int32(0)
	goto L10
L10:
	;
	m.G0 = v16 + int32(32)
	return v314
L11:
	;
	F_sequence_close(m, v29, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L90
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L87
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L84
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L81
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+112))
	v35 = F_RelationGetIndexList(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L77
	}
L18:
	;
	v38 = l2 & int32(2)
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v40 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l2&int32(1) == int32(0) {
		v79 = v5
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v48&int32(1) != 0 {
		goto L13
	} else {
		goto L27
	}
L24:
	;
	v48 = int32(1)
	goto L26
L25:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+76)))
	v48 = v47
	goto L26
L26:
	;
	goto L23
L27:
	;
	v52 = F_list_copy(m, v35)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v52
	v57 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[234])) = v58
	F_CommandCounterIncrement(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	if l2&int32(8) != 0 {
		v88 = int32(117)
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if v34 == int32(0) {
		v79 = v5
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v68 & int32(-5)
	v76 = F_reindex_relation(m, l0, v34, l2, v16+int32(24))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v79 = v76
	goto L31
L35:
	;
	v89 = int32(0)
	if v35 == v89 {
		v287 = v89
		goto L11
	} else {
		goto L38
	}
L36:
	;
	if l2&int32(16) != 0 {
		v88 = int32(112)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+118)))
	v88 = v87
	goto L35
L38:
	;
	v92 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v94 <= v92 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v287 = int32(1)
	goto L11
L40:
	;
	goto L41
L41:
	;
	v104 = v92
	v112 = int32(1)
	goto L42
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v104<<(uint(int32(2))%32))))
	v121 = F_get_rel_namespace(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L46
	}
L43:
	;
	v287 = v218
	goto L11
L44:
	;
	v218 = int32(1)
	v220 = v104 + v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v220 < v221 {
		v104 = v220
		v112 = v217
		goto L42
	} else {
		goto L76
	}
L45:
	;
	F_reindex_index(m, l0, v120, base.B2i32(l2&int32(4) == int32(0)), base.I32_extend8_s(v88), l3)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L70
	}
L46:
	;
	if v121 != int32(99) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v128 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L48:
	;
	v127 = F_isTempToastNamespace(m, v121)
	mBase = m.M
	v128 = v127
	goto L50
L49:
	;
	v128 = int32(1)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v131 = F_get_index_isvalid(m, v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v131 != 0 {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v135 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	if v135 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v38 == int32(0) {
		v217 = v112
		goto L44
	} else {
		goto L63
	}
L58:
	;
	v140 = F_get_namespace_name(m, v121)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v142 = F_get_rel_name(m, v120)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v140
	F_errmsg(m, int32(333450), v16+int32(16))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(492295), int32(4061), int32(262857))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+72))
	if v164 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v166&int32(1) != 0 {
		goto L12
	} else {
		goto L68
	}
L65:
	;
	v166 = int32(1)
	goto L67
L66:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+76)))
	v166 = v165
	goto L67
L67:
	;
	goto L64
L68:
	;
	v169 = int32(4411980)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v172 = F_list_delete_ptr(m, v171, v120)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v172
	v217 = v112
	goto L44
L70:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v183 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v217 = v112 + int32(1)
	goto L44
L73:
	;
	goto L72
L74:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v187 != int32(1) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v190 = int32(4509908)
	v192 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v193 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v192 + v193
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v196 + v193
	*(*int64)(unsafe.Add(mBase, uint32(v183+int32(56))+232)) = base.I64_extend_i32_s(v112)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v204 + v193
	v210 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v210 - v193
	goto L73
L76:
	;
	goto L43
L77:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+68))
	v229 = F_get_namespace_name(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v231 + int32(4)
	F_errmsg_internal(m, int32(691638), v16)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(492295), int32(3980), int32(262857))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
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
	F_errmsg_internal(m, int32(328323), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(492295), int32(4188), int32(336502))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errmsg_internal(m, int32(260576), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(492295), int32(4190), int32(336502))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(260576), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(492295), int32(4203), int32(336520))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v314 = v287 | v79
	goto L10
}
