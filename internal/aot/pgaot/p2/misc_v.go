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
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
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
	var v65 int32
	_ = v65
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
	v38 = base.I32_extend8_s(v24)
	goto L15
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
	if base.B2i32(v61 != v62) == int32(0) {
		goto L12
	} else {
		goto L21
	}
L15:
	;
	v53 = int32(_a_F_validate_option_array_item_0)
	v55 = int32(54)
	goto L16
L16:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v58 == v38&int32(255) {
		v115 = v53
		v117 = v55
		goto L13
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v60 = int32(1)
	v61 = v55 - v60
	v62 = int32(0)
	v65 = v53 + v60
	if v65&int32(3) == v62 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v61 != 0 {
		v53 = v65
		v55 = v61
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v78 = v38 & int32(255)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if base.B2i32(v78 == v79)|base.B2i32(base.Ui32(v61) < base.Ui32(int32(4))) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = v65
	v90 = v61
	goto L25
L23:
	;
	v108 = v65
	v110 = v61
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
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[0]))
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
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[0]))
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
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[0]))
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L53
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v52 = F_MakeTupleTableSlot(m, v50, int32(_a_F_validate_remote_info_1))
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
	v105 = m.ExcPending
	if v105 != 0 {
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
	F_slot_getsomeattrs_int(m, v52, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68 != 0 {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+6)))
	if v69 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_slot_getsomeattrs_int(m, v52, int32(2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v76 = v67
	goto L28
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v77 == int32(0) {
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v76 = v75
	goto L28
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	m.T0[v81].(func(*base.Module, int32))(m, v52)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v84 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_pfree(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v87 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_tuplestore_end(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v90 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_FreeTupleDesc(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	v94 = m.ExcPending
	if v94 != 0 {
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
	v98 = m.ExcPending
	if v98 != 0 {
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v106
	F_errmsg(m, int32(_a_F_validate_remote_info_8), v6+int32(-48))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_validate_remote_info_9), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(989), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(994), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_validate_remote_info_5), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1009), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_validate_remote_info_6)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v164
	F_errmsg(m, int32(_a_F_validate_remote_info_7), v8)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1019), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
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
func F_varbittypmodin(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(_a_F_varbittypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_varstr_levenshtein(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v116 int32
	_ = v116
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v306 int32
	_ = v306
	var __phi306 int32
	_ = __phi306
	var v307 int32
	_ = v307
	var __phi307 int32
	_ = __phi307
	var v312 int32
	_ = v312
	var __phi312 int32
	_ = __phi312
	var v317 int32
	_ = v317
	var __phi317 int32
	_ = __phi317
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v448 int32
	_ = v448
	var v474 int32
	_ = v474
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
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
	var v589 int32
	_ = v589
	var v612 int32
	_ = v612
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	v8 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = F_pg_mbstrlen_with_len(m, l2, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L81
	}
L4:
	;
	m.G0 = v26 + int32(16)
	return v636
L5:
	;
	if v28 == int32(0) {
		v636 = v32 * l4
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v32 == int32(0) {
		v636 = l5 * v28
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.B2i32(int32(255) < v28)|base.B2i32(int32(256) <= v32) != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if base.B2i32(l1 == v28)&base.B2i32(l3 == v32) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v54 = F_palloc(m, v28<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v138 = v8
	goto L11
L11:
	;
	v146 = v28 + int32(1)
	v149 = F_palloc(m, v146<<(uint(int32(3))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	if int32(0) < v28 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = l0
	v66 = v8
	goto L16
L14:
	;
	v116 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v116<<(uint(int32(2))%32)))) = int32(0)
	v138 = v54
	goto L11
L16:
	;
	v85 = F_pg_mblen_range(m, v60, l0+l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v116 = v28
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v66<<(uint(int32(2))%32)))) = v85
	v90 = v66 + int32(1)
	if v90 != v28 {
		v60 = v60 + v85
		v66 = v90
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v28) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v32+int32(1) < int32(2) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v155 = int32(3)
	v156 = v146 & v155
	v157 = int32(0)
	if base.Ui32(v155) <= base.Ui32(v28) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v163 = v157
	v173 = v8
	goto L26
L24:
	;
	v219 = v157
	goto L25
L25:
	;
	v242 = v219
	v251 = v8
	goto L30
L26:
	;
	v185 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v149+v163<<(uint(v185)%32)))) = v163 * l5
	v191 = v163 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v149+v191<<(uint(v185)%32)))) = l5 * v191
	v198 = v163 | v185
	*(*int32)(unsafe.Add(mBase, uint32(v149+v198<<(uint(v185)%32)))) = l5 * v198
	v205 = v163 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v149+v205<<(uint(v185)%32)))) = l5 * v205
	v211 = int32(4)
	v212 = v163 + v211
	v214 = v173 + v211
	if v214 != v146&int32(-4) {
		v163 = v212
		v173 = v214
		goto L26
	} else {
		goto L28
	}
L27:
	;
	if v156 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v219 = v212
	goto L25
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149+v242<<(uint(int32(2))%32)))) = v242 * l5
	v269 = int32(1)
	v272 = v251 + v269
	if v272 != v156 {
		v242 = v242 + v269
		v251 = v272
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L21
L32:
	;
	goto L31
L33:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v589+v28<<(uint(int32(2))%32))))
	v636 = v612
	goto L4
L34:
	;
	v589 = v149
	goto L33
L35:
	;
	goto L36
L36:
	;
	__phi306 = int32(1)
	__phi307 = l2
	__phi312 = v149
	__phi317 = v149 + v146<<(uint(int32(2))%32)
	v306 = __phi306
	v307 = __phi307
	v312 = __phi312
	v317 = __phi317
	goto L37
L37:
	;
	if base.B2i32(l3 == v32) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v589 = v317
	goto L33
L39:
	;
	v331 = F_pg_mblen_range(m, v307, l2+l3)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v333 = int32(1)
	goto L41
L41:
	;
	v334 = v306 * l4
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = v334
	if v138 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v333 = v331
	goto L41
L43:
	;
	if v32 != v306 {
		__phi306 = v306 + int32(1)
		__phi307 = v307 + v333
		__phi312 = v317
		__phi317 = v312
		v306 = __phi306
		v307 = __phi307
		v312 = __phi312
		v317 = __phi317
		goto L37
	} else {
		goto L80
	}
L44:
	;
	if v146 < int32(2) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v507 = int32(1)
	if v146 <= v507 {
		goto L43
	} else {
		goto L67
	}
L47:
	;
	v339 = int32(1)
	v349 = l0
	v352 = v339
	v353 = v334
	goto L48
L48:
	;
	v366 = v352 << (uint(int32(2)) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v312+v366)))
	v369 = v368 + l4
	v371 = l5 + v353
	if v369 < v371 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L43
L50:
	;
	v373 = v369
	goto L52
L51:
	;
	v373 = v371
	goto L52
L52:
	;
	v374 = int32(1)
	v377 = (v352 - v374) << (uint(int32(2)) % 32)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v138+v377)))
	v380 = v349 + v379
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380-v374))))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v333-v339))))
	if base.B2i32(v383 != v384)|base.B2i32(v379 != v333) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v373 < v498 {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v377+v312)))
	v498 = v474
	goto L53
L55:
	;
	if v333 == int32(1) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v377+v312)))
	v498 = v448 + l6
	goto L53
L58:
	;
	v393 = v333
	goto L59
L59:
	;
	if v393 <= int32(0) {
		goto L54
	} else {
		goto L61
	}
L60:
	;
	goto L57
L61:
	;
	v418 = v393 - int32(1)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+v418))))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+v307))))
	if v420 == v422 {
		v393 = v418
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v500 = v373
	goto L65
L64:
	;
	v500 = v498
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317+v366))) = v500
	if base.B2i32(v28 == v352) == int32(0) {
		v349 = v380
		v352 = v352 + int32(1)
		v353 = v500
		goto L48
	} else {
		goto L66
	}
L66:
	;
	goto L49
L67:
	;
	v511 = v507
	v517 = l0
	v521 = v334
	goto L68
L68:
	;
	v534 = v511 << (uint(int32(2)) % 32)
	v536 = v534 + v312
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v538 = v537 + l4
	v539 = l5 + v521
	if v538 < v539 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L43
L70:
	;
	v541 = v538
	goto L72
L71:
	;
	v541 = v539
	goto L72
L72:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v536-int32(4))))
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v546 != v547 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v549 = l6
	goto L75
L74:
	;
	v549 = int32(0)
	goto L75
L75:
	;
	v550 = v544 + v549
	if v541 < v550 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v552 = v541
	goto L78
L77:
	;
	v552 = v550
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317+v534))) = v552
	v554 = int32(1)
	if v511 != v28 {
		v511 = v511 + v554
		v517 = v517 + v554
		v521 = v552
		goto L68
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	goto L38
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(255)
	F_errmsg(m, int32(_a_F_varstr_levenshtein_0), v26)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_1), int32(135), int32(_a_F_varstr_levenshtein_2))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
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
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int64
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int64
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
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
	v365 = v4
	v366 = v4
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v366
	if l2 != 0 {
		goto L64
	} else {
		goto L65
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
	v365 = v352
	v366 = v353
	goto L5
L8:
	;
	v42 = v40 + int32(24)
	v43 = int32(85)
	v49 = int64(0)
	v50 = int32(_a_F_visibilitymap_count_0)
	if (v40+int32(27))&int32(-4) == v42 {
		goto L13
	} else {
		goto L14
	}
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
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v57 = int32(1431655765)
	v58 = v42
	v60 = v50
	v62 = int32(0)
	v65 = v49
	goto L16
L14:
	;
	v116 = v42
	v118 = v50
	v123 = v49
	goto L15
L15:
	;
	if v118 == int32(0) {
		v195 = v123
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v85 = base.I64_extend_i32_u(base.I32_popcnt(v66&v57)) + (base.I64_extend_i32_u(base.I32_popcnt(v70&v57)) + (base.I64_extend_i32_u(base.I32_popcnt(v74&v57)) + (v65 + base.I64_extend_i32_u(base.I32_popcnt(v78&v57)))))
	v86 = int32(16)
	v87 = v60 - v86
	v89 = v58 + v86
	v91 = v62 + int32(4)
	if v91 != int32(2040) {
		v58 = v89
		v60 = v87
		v62 = v91
		v65 = v85
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v95 = v89
	v97 = v87
	v99 = int32(0)
	v102 = v85
	goto L19
L18:
	;
	goto L17
L19:
	;
	v103 = int32(4)
	v104 = v97 - v103
	v106 = v95 + v103
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v111 = v102 + base.I64_extend_i32_u(base.I32_popcnt(v107&v57))
	v113 = v99 + int32(1)
	if v113 != int32(2) {
		v95 = v106
		v97 = v104
		v99 = v113
		v102 = v111
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v116 = v106
	v118 = v104
	v123 = v111
	goto L15
L21:
	;
	goto L20
L22:
	;
	goto L12
L23:
	;
	v127 = v118 & int32(3)
	if v127 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if base.Ui32(v118) < base.Ui32(int32(4)) {
		v195 = v157
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v150 = v116
	v155 = v118
	v157 = v123
	goto L24
L26:
	;
	goto L27
L27:
	;
	v131 = v116
	v135 = int32(0)
	v136 = v118
	v138 = v123
	goto L28
L28:
	;
	v139 = int32(1)
	v140 = v131 + v139
	v142 = v136 - v139
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v145 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v143&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v146 = v138 + v145
	v148 = v135 + v139
	if v148 != v127 {
		v131 = v140
		v135 = v148
		v136 = v142
		v138 = v146
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v150 = v140
	v155 = v142
	v157 = v146
	goto L24
L30:
	;
	goto L29
L31:
	;
	v160 = v150
	v165 = v155
	v167 = v157
	goto L32
L32:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+3)))
	v170 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v168&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+2)))
	v173 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v171&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	v176 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v174&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v179 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v177&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v183 = v170 + (v173 + (v176 + (v167 + v179)))
	v184 = int32(4)
	v187 = v165 - v184
	if v187 != 0 {
		v160 = v160 + v184
		v165 = v187
		v167 = v183
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v195 = v183
	goto L22
L34:
	;
	goto L33
L35:
	;
	v197 = int32(170)
	v203 = int64(0)
	v204 = int32(_a_F_visibilitymap_count_0)
	if (v40+int32(27))&int32(-4) == v42 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v352 = v18
	goto L37
L37:
	;
	v353 = v19 + base.I32_wrap_i64(v195)
	F_ReleaseBuffer(m, v17)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L61
	}
L38:
	;
	v352 = v18 + base.I32_wrap_i64(v349)
	goto L37
L39:
	;
	v211 = int32(-1431655766)
	v212 = v42
	v214 = v204
	v216 = int32(0)
	v219 = v203
	goto L42
L40:
	;
	v270 = v42
	v272 = v204
	v277 = v203
	goto L41
L41:
	;
	if v272 == int32(0) {
		v349 = v277
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v239 = base.I64_extend_i32_u(base.I32_popcnt(v220&v211)) + (base.I64_extend_i32_u(base.I32_popcnt(v224&v211)) + (base.I64_extend_i32_u(base.I32_popcnt(v228&v211)) + (v219 + base.I64_extend_i32_u(base.I32_popcnt(v232&v211)))))
	v240 = int32(16)
	v241 = v214 - v240
	v243 = v212 + v240
	v245 = v216 + int32(4)
	if v245 != int32(2040) {
		v212 = v243
		v214 = v241
		v216 = v245
		v219 = v239
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v249 = v243
	v251 = v241
	v253 = int32(0)
	v256 = v239
	goto L45
L44:
	;
	goto L43
L45:
	;
	v257 = int32(4)
	v258 = v251 - v257
	v260 = v249 + v257
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v265 = v256 + base.I64_extend_i32_u(base.I32_popcnt(v261&v211))
	v267 = v253 + int32(1)
	if v267 != int32(2) {
		v249 = v260
		v251 = v258
		v253 = v267
		v256 = v265
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v270 = v260
	v272 = v258
	v277 = v265
	goto L41
L47:
	;
	goto L46
L48:
	;
	goto L38
L49:
	;
	v281 = v272 & int32(3)
	if v281 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if base.Ui32(v272) < base.Ui32(int32(4)) {
		v349 = v311
		goto L48
	} else {
		goto L57
	}
L51:
	;
	v304 = v270
	v309 = v272
	v311 = v277
	goto L50
L52:
	;
	goto L53
L53:
	;
	v285 = v270
	v289 = int32(0)
	v290 = v272
	v292 = v277
	goto L54
L54:
	;
	v293 = int32(1)
	v294 = v285 + v293
	v296 = v290 - v293
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v297&v197)+uint32(_c_F_visibilitymap_count[2]))))
	v300 = v292 + v299
	v302 = v289 + v293
	if v302 != v281 {
		v285 = v294
		v289 = v302
		v290 = v296
		v292 = v300
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v304 = v294
	v309 = v296
	v311 = v300
	goto L50
L56:
	;
	goto L55
L57:
	;
	v314 = v304
	v319 = v309
	v321 = v311
	goto L58
L58:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+3)))
	v324 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v322&v197)+uint32(_c_F_visibilitymap_count[2]))))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+2)))
	v327 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v325&v197)+uint32(_c_F_visibilitymap_count[2]))))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	v330 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v328&v197)+uint32(_c_F_visibilitymap_count[2]))))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	v333 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v331&v197)+uint32(_c_F_visibilitymap_count[2]))))
	v337 = v324 + (v327 + (v330 + (v321 + v333)))
	v338 = int32(4)
	v341 = v319 - v338
	if v341 != 0 {
		v314 = v314 + v338
		v319 = v341
		v321 = v337
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v349 = v337
	goto L48
L60:
	;
	goto L59
L61:
	;
	v357 = v22 + int32(1)
	v359 = F_vm_readbuf(m, l0, v357, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v359 != 0 {
		v17 = v359
		v18 = v352
		v19 = v353
		v22 = v357
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L7
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v365
	goto L66
L65:
	;
	goto L66
L66:
	;
	return
}
