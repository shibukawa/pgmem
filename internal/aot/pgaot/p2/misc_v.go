package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_validate_option_array_item(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var __phi23 int32
	_ = __phi23
	var v24 int32
	_ = v24
	var __phi24 int32
	_ = __phi24
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l1 != 0 {
		v176 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v184 = int32(1)
	v189 = F_find_option(m, l0, v184, (l2|v176)&v184, int32(21))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		v176 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	__phi23 = l0
	__phi24 = v16
	__phi27 = int32(1)
	__phi30 = v4
	v23 = __phi23
	v24 = __phi24
	v27 = __phi27
	v30 = __phi30
	goto L4
L4:
	;
	if v24 == int32(46) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v176 = base.B2i32(v24&int32(255) != int32(46)) & v162
	goto L1
L6:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v167 != 0 {
		__phi23 = v23 + int32(1)
		__phi24 = v167
		__phi27 = base.B2i32(v24 == int32(46))
		__phi30 = v162
		v23 = __phi23
		v24 = __phi24
		v27 = __phi27
		v30 = __phi30
		goto L4
	} else {
		goto L38
	}
L7:
	;
	if v27 == int32(0) {
		v162 = int32(1)
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v37 = int32(_a_F_validate_option_array_item_0)
	v38 = base.I32_extend8_s(v24)
	v39 = int32(54)
	goto L14
L10:
	;
	v176 = int32(0)
	goto L1
L11:
	;
	if v144|base.B2i32(v38 < int32(0)) != 0 {
		v162 = v30
		goto L6
	} else {
		goto L36
	}
L12:
	;
	v144 = int32(0)
	goto L11
L13:
	;
	v122 = v115
	v124 = v117
	goto L30
L14:
	;
	goto L21
L21:
	;
	v78 = v38 & int32(255)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_option_array_item[0])))
	if base.B2i32(v78 == v79)|int32(0) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = v37
	v90 = v39
	goto L25
L23:
	;
	v108 = v37
	v110 = v39
	goto L24
L24:
	;
	if v110 == int32(0) {
		goto L12
	} else {
		goto L29
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v95 = v94 ^ v78*int32(16843009)
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 != v98 {
		v115 = v88
		v117 = v90
		goto L13
	} else {
		goto L27
	}
L26:
	;
	v108 = v103
	v110 = v105
	goto L24
L27:
	;
	v102 = int32(4)
	v103 = v88 + v102
	v105 = v90 - v102
	if base.Ui32(int32(3)) < base.Ui32(v105) {
		v88 = v103
		v90 = v105
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v115 = v108
	v117 = v110
	goto L13
L30:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v38&int32(255) == v127 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v144 = v122
	goto L11
L33:
	;
	goto L34
L34:
	;
	v129 = int32(1)
	v132 = v124 - v129
	if v132 != 0 {
		v122 = v122 + v129
		v124 = v132
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v38))%64)&int64(287948969894477825) == int64(0))|(v27|base.B2i32(base.Ui32(int32(63)) < base.Ui32(v24))) != 0 {
		v176 = int32(0)
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v162 = v30
	goto L6
L38:
	;
	goto L5
L39:
	;
	m.G0 = v14 + int32(16)
	return v262
L40:
	;
	return int32(0)
L41:
	;
	if v176|v189 == int32(0) {
		v262 = v4
		goto L39
	} else {
		goto L42
	}
L42:
	;
	if v189 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	switch v230 - int32(5) {
	case 0:
		goto L58
	case 1:
		goto L56
	default:
		goto L57
	}
L44:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+21)))
	if v196&int32(2) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v202 = F_superuser(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	if v202 != 0 {
		v262 = int32(1)
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v207 = F_pg_parameter_aclcheck(m, l0, v205, int64(4096))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L40
	} else {
		goto L50
	}
L50:
	;
	v210 = base.B2i32(v207 == int32(0))
	if l2|v210 != 0 {
		v262 = v210
		goto L39
	} else {
		goto L51
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L40
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L40
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_validate_option_array_item_1), v14)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L40
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_validate_option_array_item_2), int32(_a_F_validate_option_array_item_3), int32(_a_F_validate_option_array_item_4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L40
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v248 = F_superuser(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L40
	} else {
		goto L64
	}
L57:
	;
	if l2 != 0 {
		v262 = v4
		goto L39
	} else {
		goto L63
	}
L58:
	;
	v233 = F_superuser(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L40
	} else {
		goto L59
	}
L59:
	;
	if v233 != 0 {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v240 = F_pg_parameter_aclcheck(m, l0, v238, int64(4096))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L40
	} else {
		goto L61
	}
L61:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(v240 == int32(0)) != 0 {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v262 = v4
	goto L39
L63:
	;
	goto L56
L64:
	;
	if v248 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v250 = int32(5)
	goto L67
L66:
	;
	v250 = int32(6)
	goto L67
L67:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v254 = int32(0)
	v258 = F_set_config_with_handle(m, l0, int32(0), l1, v250, int32(12), v253, v254, v254, v254, v254)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L40
	} else {
		goto L68
	}
L68:
	;
	v262 = int32(1)
	goto L39
}
func F_validate_remote_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
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
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = int64(68719476752)
	v13 = v6 + int32(-24)
	F_initStringInfo(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	v18 = F_quote_literal_cstr(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v18
	F_appendStringInfo(m, v13, int32(_a_F_validate_remote_info_0), v6+int32(-32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[1]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v30 = base.B2i32(v28 == int32(2))
	goto L5
L5:
	;
	if v30 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[2]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+60))
	v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v35, int32(2), v6+int32(-8))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	F_pfree(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v47 == int32(2) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L53
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v52 = F_MakeSingleTupleTableSlot(m, v50, int32(_a_F_validate_remote_info_1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
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
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L49
	}
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v57 = F_tuplestore_gettupleslot(m, v54, int32(1), int32(0), v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v57 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+6)))
	if v61 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	m.T0[v66].(func(*base.Module, int32, int32))(m, v52, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	if v70 != int64(0) {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+6)))
	if v73 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	m.T0[v78].(func(*base.Module, int32, int32))(m, v52, int32(2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v82 = v69
	goto L28
L28:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
	if v83 == int64(0) {
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v82 = v81
	goto L28
L30:
	;
	F_ExecDropSingleTupleTableSlot(m, v52)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v88 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_pfree(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v91 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_tuplestore_end(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v94 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_FreeTupleDesc(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v42)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	if v30 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v8 - int32(-64)
	return
L48:
	;
	goto L47
L49:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v110
	F_errmsg(m, int32(_a_F_validate_remote_info_8), v6+int32(-48))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_validate_remote_info_9), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1152), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_validate_remote_info_2), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1157), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_validate_remote_info_5), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1172), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_validate_remote_info_6)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v168
	F_errmsg(m, int32(_a_F_validate_remote_info_7), v8)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1182), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
func F_varbittypmodin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(_a_F_varbittypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_varstr_levenshtein(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v138 int32
	_ = v138
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v215 int32
	_ = v215
	var v221 int64
	_ = v221
	var v229 int64
	_ = v229
	var v237 int64
	_ = v237
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v258 int64
	_ = v258
	var v284 int64
	_ = v284
	var v291 int64
	_ = v291
	var v309 int64
	_ = v309
	var v312 int64
	_ = v312
	var v343 int64
	_ = v343
	var v353 int32
	_ = v353
	var __phi353 int32
	_ = __phi353
	var v354 int32
	_ = v354
	var __phi354 int32
	_ = __phi354
	var v358 int32
	_ = v358
	var __phi358 int32
	_ = __phi358
	var v364 int64
	_ = v364
	var __phi364 int64
	_ = __phi364
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v417 int32
	_ = v417
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v507 int64
	_ = v507
	var v538 int64
	_ = v538
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int64
	_ = v584
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int64
	_ = v607
	var v608 int64
	_ = v608
	var v609 int64
	_ = v609
	var v611 int64
	_ = v611
	var v614 int64
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v622 int64
	_ = v622
	var v624 int32
	_ = v624
	var v657 int64
	_ = v657
	var v662 int32
	_ = v662
	var v688 int64
	_ = v688
	var v700 int64
	_ = v700
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	v8 = int64(0)
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	v31 = base.I64_extend_i32_s(l4)
	v32 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v36 = F_pg_mbstrlen_with_len(m, l2, l3)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v32 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L93
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L89
	}
L6:
	;
	m.G0 = v29 + int32(16)
	return base.I32_wrap_i64(v700)
L7:
	;
	v41 = base.I64_extend_i32_s(v36) * v31
	if base.Ui64(int64(-4294967297)) < base.Ui64(v41-int64(2147483648)) {
		v700 = v41
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v46 = base.I64_extend_i32_s(l5)
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v50 = base.I64_extend_i32_s(v32) * v46
	if base.Ui64(int64(-4294967297)) < base.Ui64(v50-int64(2147483648)) {
		v700 = v50
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(int32(255) < v32)|base.B2i32(int32(256) <= v36) != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L4
L15:
	;
	if base.B2i32(l1 == v32)&base.B2i32(l3 == v36) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = F_palloc(m, v32<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v162 = int32(0)
	goto L18
L18:
	;
	v171 = v32 + int32(1)
	v174 = F_palloc(m, v171<<(uint(int32(4))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v71 = int32(0)
	if v71 < v32 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v76 = l0
	v79 = v71
	goto L23
L21:
	;
	v138 = int32(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v138<<(uint(int32(2))%32)))) = int32(0)
	v162 = v69
	goto L18
L23:
	;
	v104 = F_pg_mblen_range(m, v76, l0+l1)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v138 = v32
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v79<<(uint(int32(2))%32)))) = v104
	v109 = v79 + int32(1)
	if v109 != v32 {
		v76 = v76 + v104
		v79 = v109
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v177 = v36 + int32(1)
	if base.Ui32(int32(2147483646)) < base.Ui32(v32) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v177 < int32(2) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v180 = base.I64_extend_i32_u(v171)
	v182 = v180 & int64(3)
	if base.Ui32(int32(4)) <= base.Ui32(v171) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v195 = v8
	v198 = int64(0)
	goto L33
L31:
	;
	v258 = v8
	goto L32
L32:
	;
	v284 = v258
	v291 = v8
	goto L37
L33:
	;
	v215 = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v174+base.I32_wrap_i64(v195)<<(uint(v215)%32)))) = v195 * v46
	v221 = v195 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v174+base.I32_wrap_i64(v221)<<(uint(v215)%32)))) = v46 * v221
	v229 = v195 | int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v174+base.I32_wrap_i64(v229)<<(uint(v215)%32)))) = v46 * v229
	v237 = v195 | int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v174+base.I32_wrap_i64(v237)<<(uint(v215)%32)))) = v46 * v237
	v244 = int64(4)
	v245 = v195 + v244
	v247 = v198 + v244
	if v247 != v180&int64(2147483644) {
		v195 = v245
		v198 = v247
		goto L33
	} else {
		goto L35
	}
L34:
	;
	if v182 == int64(0) {
		goto L28
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v258 = v245
	goto L32
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v174+base.I32_wrap_i64(v284)<<(uint(int32(3))%32)))) = v284 * v46
	v309 = int64(1)
	v312 = v291 + v309
	if v312 != v182 {
		v284 = v284 + v309
		v291 = v312
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L28
L39:
	;
	goto L38
L40:
	;
	v688 = *(*int64)(unsafe.Add(mBase, uint32(v662+v32<<(uint(int32(3))%32))))
	if base.Ui64(v688-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		goto L4
	} else {
		goto L88
	}
L41:
	;
	v662 = v174
	goto L40
L42:
	;
	goto L43
L43:
	;
	v343 = base.I64_extend_i32_s(l6)
	__phi353 = v174
	__phi354 = l2
	__phi358 = v174 + v171<<(uint(int32(3))%32)
	__phi364 = int64(1)
	v353 = __phi353
	v354 = __phi354
	v358 = __phi358
	v364 = __phi364
	goto L44
L44:
	;
	if base.B2i32(l3 == v36) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v662 = v358
	goto L40
L46:
	;
	v381 = F_pg_mblen_range(m, v354, l2+l3)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	v383 = int32(1)
	goto L48
L48:
	;
	v384 = v364 * v31
	*(*int64)(unsafe.Add(mBase, uint32(v358))) = v384
	if v162 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v383 = v381
	goto L48
L50:
	;
	v657 = v364 + int64(1)
	if v657 != base.I64_extend_i32_u(v177) {
		__phi353 = v358
		__phi354 = v354 + v383
		__phi358 = v353
		__phi364 = v657
		v353 = __phi353
		v354 = __phi354
		v358 = __phi358
		v364 = __phi364
		goto L44
	} else {
		goto L87
	}
L51:
	;
	if v171 < int32(2) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v574 = int32(1)
	if v171 <= v574 {
		goto L50
	} else {
		goto L74
	}
L54:
	;
	v387 = int32(1)
	v394 = l0
	v395 = v387
	v397 = v384
	goto L55
L55:
	;
	v417 = v395 << (uint(int32(3)) % 32)
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v353+v417)))
	v420 = v419 + v31
	v422 = v397 + v46
	if v420 < v422 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L50
L57:
	;
	v424 = v420
	goto L59
L58:
	;
	v424 = v422
	goto L59
L59:
	;
	v425 = int32(1)
	v426 = v395 - v425
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v162+v426<<(uint(int32(2))%32))))
	v431 = v394 + v430
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431-v425))))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+v383-v387))))
	if base.B2i32(v434 != v435)|base.B2i32(v430 != v383) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if v424 < v565 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v353+v426<<(uint(int32(3))%32))))
	v565 = v538
	goto L60
L62:
	;
	if v383 == int32(1) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v353+v426<<(uint(int32(3))%32))))
	v565 = v507 + v343
	goto L60
L65:
	;
	v444 = v383
	goto L66
L66:
	;
	if v444 <= int32(0) {
		goto L61
	} else {
		goto L68
	}
L67:
	;
	goto L64
L68:
	;
	v472 = v444 - int32(1)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+v472))))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v354))))
	if v474 == v476 {
		v444 = v472
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v567 = v424
	goto L72
L71:
	;
	v567 = v565
	goto L72
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v417+v358))) = v567
	if base.B2i32(v395 == v32) == int32(0) {
		v394 = v431
		v395 = v395 + int32(1)
		v397 = v567
		goto L55
	} else {
		goto L73
	}
L73:
	;
	goto L56
L74:
	;
	v578 = v574
	v581 = l0
	v584 = v384
	goto L75
L75:
	;
	v604 = v578 << (uint(int32(3)) % 32)
	v606 = v604 + v353
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v606)))
	v608 = v607 + v31
	v609 = v584 + v46
	if v608 < v609 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L50
L77:
	;
	v611 = v608
	goto L79
L78:
	;
	v611 = v609
	goto L79
L79:
	;
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v606-int32(8))))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v616 != v617 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v619 = v343
	goto L82
L81:
	;
	v619 = int64(0)
	goto L82
L82:
	;
	v620 = v614 + v619
	if v611 < v620 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v622 = v611
	goto L85
L84:
	;
	v622 = v620
	goto L85
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v358+v604))) = v622
	v624 = int32(1)
	if v578 != v32 {
		v578 = v578 + v624
		v581 = v581 + v624
		v584 = v622
		goto L75
	} else {
		goto L86
	}
L86:
	;
	goto L76
L87:
	;
	goto L45
L88:
	;
	v700 = v688
	goto L6
L89:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(255)
	F_errmsg(m, int32(_a_F_varstr_levenshtein_0), v29)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_1), int32(138), int32(_a_F_varstr_levenshtein_2))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_varstr_levenshtein_3), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_4), int32(_a_F_varstr_levenshtein_5), int32(_a_F_varstr_levenshtein_6))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_visibilitymap_count(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	v4 = int32(0)
	v12 = F_vm_readbuf(m, l0, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = v12
	v18 = v4
	v19 = v4
	v22 = v4
	goto L6
L4:
	;
	v117 = v4
	v118 = v4
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v118
	if l2 != 0 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	if v17 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v117 = v104
	v118 = v105
	goto L5
L8:
	;
	v42 = v40 + int32(24)
	v43 = int32(85)
	v46 = v42
	v48 = int64(0)
	v49 = int32(0)
	goto L13
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_count[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v17^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L8
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_count[1]))
	v40 = v34 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)))
	v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v50&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v53&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v56&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v61 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v59&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v65 = v52 + (v55 + (v58 + (v48 + v61)))
	v66 = int32(4)
	v69 = v49 + v66
	if v69 != int32(_a_F_visibilitymap_count_0) {
		v46 = v46 + v66
		v48 = v65
		v49 = v69
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	goto L14
L16:
	;
	v73 = int32(170)
	v76 = v42
	v78 = int64(0)
	v79 = int32(0)
	goto L20
L17:
	;
	v104 = v18
	goto L18
L18:
	;
	v105 = v19 + base.I32_wrap_i64(v65)
	F_ReleaseBuffer(m, v17)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v104 = v18 + base.I32_wrap_i64(v95)
	goto L18
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+3)))
	v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v80&v73)+uint32(_c_F_visibilitymap_count[2]))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
	v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v83&v73)+uint32(_c_F_visibilitymap_count[2]))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v86&v73)+uint32(_c_F_visibilitymap_count[2]))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v91 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89&v73)+uint32(_c_F_visibilitymap_count[2]))))
	v95 = v82 + (v85 + (v88 + (v78 + v91)))
	v96 = int32(4)
	v99 = v79 + v96
	if v99 != int32(_a_F_visibilitymap_count_0) {
		v76 = v76 + v96
		v78 = v95
		v79 = v99
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	goto L21
L23:
	;
	v109 = v22 + int32(1)
	v111 = F_vm_readbuf(m, l0, v109, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v111 != 0 {
		v17 = v111
		v18 = v104
		v19 = v105
		v22 = v109
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L7
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v117
	goto L28
L27:
	;
	goto L28
L28:
	;
	return
}
