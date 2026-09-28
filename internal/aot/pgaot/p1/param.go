package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSetParamPlan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
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
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	v3 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = int32(1)
	if base.Ui32(v20) < base.Ui32(v19-v20) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0])) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30
	return
L2:
	;
	v399 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+16)) = uint8(v399)
	*(*int64)(unsafe.Add(mBase, uint32(v384)+8)) = v398
	goto L1
L3:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v19 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L16
	} else {
		goto L78
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L16
	} else {
		goto L75
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L72
	}
L7:
	;
	if v19 == int32(7) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L16
	} else {
		goto L69
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v26 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	if v27 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(1)
	if v19 == int32(6) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0]))
	v38 = F_initArrayResultAny(m, v35, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v40 = v3
	goto L15
L15:
	;
	v41 = int32(_a_F_ExecSetParamPlan_0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0])) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return
L17:
	;
	v40 = v38
	goto L15
L18:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v50 = m.T0[v49].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L16
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v253 = v247 + v250*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = int32(0)
	v384 = v253
	v398 = int64(1)
	goto L2
L23:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v233 = v227 + v230*int32(24)
	v234 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v234 != int64(0) {
		goto L64
	} else {
		goto L65
	}
L24:
	;
	if v50 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v60 = v50
	v64 = v3
	v65 = v40
	goto L28
L26:
	;
	goto L27
L27:
	;
	if v19 != int32(6) {
		goto L3
	} else {
		goto L63
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
	if v75&int32(2) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v19 == int32(6) {
		v217 = v194
		goto L23
	} else {
		goto L61
	}
L30:
	;
	if v19 == int32(0) {
		goto L22
	} else {
		goto L33
	}
L31:
	;
	v193 = v64
	v194 = v65
	goto L32
L32:
	;
	goto L29
L33:
	;
	if base.B2i32(v19 != int32(6)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v180 != 0 {
		goto L55
	} else {
		goto L56
	}
L35:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+6)))
	if v84 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if (base.B2i32(base.Ui32(int32(2)) < base.Ui32(v19-int32(3)))|(v64^int32(-1)))&int32(1) == int32(0) {
		goto L4
	} else {
		goto L43
	}
L38:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	m.T0[v89].(func(*base.Module, int32, int32))(m, v60, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v97 = F_accumArrayResultAny(m, v65, v93, v95, v96, v42)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v170 = v97
	goto L34
L43:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v107 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+44))
	v112 = m.T0[v111].(func(*base.Module, int32) int32)(m, v60)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L16
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v115 == int32(0) {
		v170 = v65
		goto L34
	} else {
		goto L49
	}
L49:
	;
	v119 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v120 <= v119 {
		v170 = v65
		goto L34
	} else {
		goto L50
	}
L50:
	;
	v125 = int32(1)
	v129 = v119
	goto L51
L51:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v129<<(uint(int32(2))%32))))
	v148 = v140 + v145*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = F_heap_getattr_2(m, v151, v125, v106, v148+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L16
	} else {
		goto L53
	}
L52:
	;
	v170 = v65
	goto L34
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v148)+8)) = v154
	v157 = int32(1)
	v160 = v129 + v157
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v160 < v161 {
		v125 = v125 + v157
		v129 = v160
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L16
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v183 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v185 = m.T0[v184].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	if v185 != 0 {
		v60 = v185
		v64 = v183
		v65 = v170
		goto L28
	} else {
		goto L60
	}
L60:
	;
	v193 = v183
	v194 = v170
	goto L32
L61:
	;
	if v193&int32(1) != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L3
L63:
	;
	v217 = v40
	goto L23
L64:
	;
	F_pfree(m, base.I32_wrap_i64(v234))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L16
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v241 = F_makeArrayResultAny(m, v217, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L16
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = int32(0)
	v246 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v384 = v233
	v398 = v246
	goto L2
L69:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_1), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1133), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L16
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
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_4), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1135), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_5), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1137), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_ExecSetParamPlan_6), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1200), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v337 = v332 + v334*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v337))) = int32(0)
	v384 = v337
	v398 = int64(0)
	goto L2
L83:
	;
	goto L84
L84:
	;
	if v329 == int32(0) {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v342 <= int32(0) {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v348 = int32(0)
	goto L87
L87:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+v348<<(uint(int32(2))%32))))
	v371 = v363 + v368*int32(24)
	v372 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+16)) = uint8(v372)
	*(*int64)(unsafe.Add(mBase, uint32(v371)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(0)
	v379 = v348 + v372
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v379 < v380 {
		v348 = v379
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L1
L89:
	;
	goto L88
}
func F_find_param_referent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 != int32(1) {
		v325 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v325
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v29 == int32(0) {
		v325 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 <= int32(0) {
		v325 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = int32(0)
	if v35 < v32 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v38 = v32
	goto L7
L6:
	;
	v38 = v35
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v52 = v40
	v53 = v5
	goto L8
L8:
	;
	v60 = v53 << (uint(int32(2)) % 32)
	v61 = v39 + v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != int32(23) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	return int32(0)
L10:
	;
	v317 = v53 + int32(1)
	if v317 != v38 {
		v52 = v309
		v53 = v317
		goto L8
	} else {
		goto L59
	}
L11:
	;
	v309 = v62
	goto L10
L12:
	;
	if v63 != int32(360) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v158 = int32(0)
	goto L30
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	if v52 != v68 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+88))
	if v70 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v73 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v63 != int32(23) {
		goto L11
	} else {
		goto L29
	}
L19:
	;
	v76 = int32(0)
	if v76 < v73 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v79 = v73
	goto L22
L21:
	;
	v79 = v76
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v84 = int32(0)
	goto L23
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v81+v84<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v80 != v105 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v61
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v325 = v112
	goto L1
L25:
	;
	v108 = v84 + int32(1)
	if v79 != v108 {
		v84 = v108
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L18
L29:
	;
	goto L14
L30:
	;
	v172 = int32(0)
	if v152 == v172 {
		v182 = v172
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v151 == int32(0) {
		v309 = v52
		goto L10
	} else {
		goto L35
	}
L33:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v176 <= v158 {
		v182 = int32(0)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v182 = v178 + v158<<(uint(int32(2))%32)
	goto L32
L35:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if base.B2i32(v182 == int32(0))|base.B2i32(v187 <= v158) != 0 {
		v309 = v52
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v190 == int32(0) {
		v309 = v52
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v193 == v194 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v196 = int32(4)
	v203 = base.B2i32(v60+v196 < v32<<(uint(int32(2))%32))
	if v60+v196 < v32<<(uint(int32(2))%32) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v158 = v158 + int32(1)
	goto L30
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L55
	} else {
		goto L56
	}
L42:
	;
	v204 = v61 + v196
	goto L44
L43:
	;
	v204 = int32(0)
	goto L44
L44:
	;
	if v60+v196 < v32<<(uint(int32(2))%32) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v208 = (v204 - v39) >> (uint(int32(2)) % 32)
	goto L47
L46:
	;
	v208 = v32
	goto L47
L47:
	;
	if v32 <= v208 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v190+v158<<(uint(int32(2))%32))))
	v215 = v208
	goto L49
L49:
	;
	v234 = v39 + v215<<(uint(int32(2))%32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v236 == int32(23) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v234
	return v213
L51:
	;
	v240 = v215 + int32(1)
	if v32 != v240 {
		v215 = v240
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	goto L41
L55:
	;
	return int32(0)
L56:
	;
	F_errmsg_internal(m, int32(_a_F_find_param_referent_0), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_find_param_referent_1), int32(_a_F_find_param_referent_2), int32(_a_F_find_param_referent_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	goto L9
}
