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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	v3 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = int32(1)
	if base.Ui32(v19) < base.Ui32(v18-v19) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0])) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v29
	return
L2:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v18 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L3:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v300 = v294 + v297*int32(12)
	v301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+8)) = uint8(v301)
	*(*int64)(unsafe.Add(mBase, uint32(v300))) = int64(4294967296)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L16
	} else {
		goto L77
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L16
	} else {
		goto L74
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L16
	} else {
		goto L71
	}
L7:
	;
	if v18 == int32(7) {
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L16
	} else {
		goto L68
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v25 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if v26 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(1)
	if v18 == int32(6) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0]))
	v37 = F_initArrayResultAny(m, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v39 = v3
	goto L15
L15:
	;
	v40 = int32(_a_F_ExecSetParamPlan_0)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSetParamPlan[0])) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return
L17:
	;
	v39 = v37
	goto L15
L18:
	;
	F_ExecReScan(m, v27)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v49 = m.T0[v48].(func(*base.Module, int32) int32)(m, v27)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v225 = v219 + v222*int32(12)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v226 != 0 {
		goto L63
	} else {
		goto L64
	}
L23:
	;
	if v49 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = v49
	v63 = v3
	v64 = v39
	goto L27
L25:
	;
	goto L26
L26:
	;
	if v18 != int32(6) {
		goto L2
	} else {
		goto L62
	}
L27:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
	if v73&int32(2) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v18 == int32(6) {
		v210 = v188
		goto L22
	} else {
		goto L60
	}
L29:
	;
	if v18 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	v187 = v63
	v188 = v64
	goto L31
L31:
	;
	goto L28
L32:
	;
	if base.B2i32(v18 != int32(6)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v174 != 0 {
		goto L54
	} else {
		goto L55
	}
L34:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+6)))
	if v82 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if (base.B2i32(base.Ui32(int32(2)) < base.Ui32(v18-int32(3)))|(v63^int32(-1)))&int32(1) == int32(0) {
		goto L4
	} else {
		goto L42
	}
L37:
	;
	F_slot_getsomeattrs_int(m, v59, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v93 = F_accumArrayResultAny(m, v64, v89, v91, v92, v41)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v165 = v93
	goto L33
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v103 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+44))
	v108 = m.T0[v107].(func(*base.Module, int32) int32)(m, v59)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v111 == int32(0) {
		v165 = v64
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v116 <= v115 {
		v165 = v64
		goto L33
	} else {
		goto L49
	}
L49:
	;
	v121 = int32(1)
	v125 = v115
	goto L50
L50:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v125<<(uint(int32(2))%32))))
	v143 = v135 + v140*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v149 = F_heap_getattr_2(m, v146, v121, v102, v143+int32(8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L16
	} else {
		goto L52
	}
L51:
	;
	v165 = v64
	goto L33
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v149
	v152 = int32(1)
	v155 = v125 + v152
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v155 < v156 {
		v121 = v121 + v152
		v125 = v155
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	F_ExecReScan(m, v27)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v177 = int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v179 = m.T0[v178].(func(*base.Module, int32) int32)(m, v27)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v179 != 0 {
		v59 = v179
		v63 = v177
		v64 = v165
		goto L27
	} else {
		goto L59
	}
L59:
	;
	v187 = v177
	v188 = v165
	goto L31
L60:
	;
	if v187&int32(1) != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L2
L62:
	;
	v210 = v39
	goto L22
L63:
	;
	F_pfree(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v230 = F_makeArrayResultAny(m, v210, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v230
	v233 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+8)) = uint8(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v235
	goto L1
L68:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_1), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1099), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_4), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1101), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSetParamPlan_5), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1103), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L16
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
	F_errcode(m, int32(66))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_ExecSetParamPlan_6), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ExecSetParamPlan_2), int32(1166), int32(_a_F_ExecSetParamPlan_3))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
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
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v329 = v324 + v326*int32(12)
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329)+8)) = uint8(v330)
	*(*int64)(unsafe.Add(mBase, uint32(v329))) = int64(0)
	goto L1
L82:
	;
	goto L83
L83:
	;
	if v321 == int32(0) {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v336 <= int32(0) {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v342 = int32(0)
	goto L86
L86:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v342<<(uint(int32(2))%32))))
	v364 = v356 + v361*int32(12)
	v365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+8)) = uint8(v365)
	*(*int64)(unsafe.Add(mBase, uint32(v364))) = int64(0)
	v370 = v342 + v365
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v370 < v371 {
		v342 = v370
		goto L86
	} else {
		goto L88
	}
L87:
	;
	goto L1
L88:
	;
	goto L87
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
	if v63 != int32(356) {
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
