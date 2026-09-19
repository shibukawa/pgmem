package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = v15
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitFunc[0]))
	v22 = F_object_aclcheck(m, int32(1255), l3, v20, int64(128))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = F_get_func_name(m, l3)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitFunc[1]))
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_aclcheck_error(m, v22, int32(19), v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_RunFunctionExecuteHook(m, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v17 < int32(101) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v36 = F_palloc0(m, int32(28))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L58
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v36
	v43 = F_palloc0(m, v17<<(uint(int32(3))%32)+int32(20))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_fmgr_info(m, l3, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = l1
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+18)) = uint16(v17)
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+16)) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v43)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v46
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v57
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+11)))
	if v60 == v51 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+10)))
	v143 = v140 & base.B2i32(int32(0) < v17)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitFunc[2]))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)))
	if v145 <= v146 {
		goto L45
	} else {
		goto L46
	}
L22:
	;
	v95 = v65
	goto L36
L23:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v66 <= v65 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v70 = v43 + int32(20)
	goto L22
L28:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_ExecInitFunc_0), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v84 = F_exprLocation(m, l1)
	mBase = m.M
	F_executor_errposition(m, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_errfinish(m, int32(_a_F_ExecInitFunc_1), int32(2759), int32(_a_F_ExecInitFunc_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v95<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v107 != int32(7) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L21
L38:
	;
	v127 = v95 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v127 < v128 {
		v95 = v127
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v112 = v70 + v95<<(uint(int32(3))%32)
	F_ExecInitExprRec(m, v106, l5, v112, v112+int32(4))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v119 = v70 + v95<<(uint(int32(3))%32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v120
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)) = uint8(v122)
	goto L38
L42:
	;
	goto L38
L43:
	;
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v163
	m.G0 = v13 + int32(16)
	return
L45:
	;
	if v143 == int32(0) {
		v163 = int32(26)
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v143 != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	if v17 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v156 = int32(29)
	goto L51
L50:
	;
	v156 = int32(27)
	goto L51
L51:
	;
	if v17 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = int32(28)
	goto L54
L53:
	;
	v159 = v156
	goto L54
L54:
	;
	v163 = v159
	goto L44
L55:
	;
	v162 = int32(31)
	goto L57
L56:
	;
	v162 = int32(30)
	goto L57
L57:
	;
	v163 = v162
	goto L44
L58:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v175 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v175
	F_errmsg_plural(m, int32(_a_F_ExecInitFunc_3), int32(_a_F_ExecInitFunc_4), v175, v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ExecInitFunc_1), int32(2732), int32(_a_F_ExecInitFunc_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LookupFuncWithArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(656)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 == v4 {
		v80 = v4
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v12 + int32(656)
	return v541
L2:
	;
	switch l0 - int32(1) {
	case 0:
		goto L135
	default:
		v541 = v446
		goto L1
	case 18:
		goto L137
	case 28:
		goto L136
	}
L3:
	;
	v430 = F_NameListToString(m, v225)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L13
	} else {
		goto L132
	}
L4:
	;
	v417 = F_NameListToString(m, v199)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L13
	} else {
		goto L129
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L124
	}
L6:
	;
	v383 = F_NameListToString(m, v171)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L13
	} else {
		goto L121
	}
L7:
	;
	v369 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v369
	F_errmsg_plural(m, int32(_a_F_LookupFuncWithArgs_0), int32(_a_F_LookupFuncWithArgs_1), v369, v12+int32(224))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L119
	}
L8:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v84 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if int32(101) <= v17 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v45 <= int32(0) {
		v80 = v17
		goto L8
	} else {
		goto L19
	}
L13:
	;
	return int32(0)
L14:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if l0 == int32(29) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v31 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v31
	F_errmsg_plural(m, int32(_a_F_LookupFuncWithArgs_2), int32(_a_F_LookupFuncWithArgs_3), v31, v12+int32(240))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2237), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v52 = v4
	goto L20
L20:
	;
	v58 = v52 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62+v58)))
	v65 = F_LookupTypeNameOid(m, v64, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L22
	}
L21:
	;
	v80 = v17
	goto L8
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+(v12+int32(256))))) = v65
	if v65 == int32(0) {
		v541 = v4
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v71 = v52 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v71 < v72 {
		v52 = v71
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v85 = l0
	goto L27
L26:
	;
	v85 = int32(34)
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v84 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = int32(-1)
	goto L30
L29:
	;
	v88 = v80
	goto L30
L30:
	;
	v94 = F_LookupFuncNameInternal(m, v85, v86, v88, v12+int32(256), int32(0), l2, v12+int32(252))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	switch l0 - int32(29) {
	case 0, 5:
		goto L35
	default:
		goto L34
	}
L32:
	;
	switch l0 - int32(1) {
	case 0:
		goto L80
	default:
		v541 = int32(0)
		goto L1
	case 18:
		goto L82
	case 28:
		goto L81
	case 33:
		goto L79
	}
L33:
	;
	if v143 == v94 {
		v446 = v143
		goto L2
	} else {
		goto L78
	}
L34:
	;
	if v94 != 0 {
		v446 = v94
		goto L2
	} else {
		goto L51
	}
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v98 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)+252))
	if v101 == int32(1) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if int32(0) < v104 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v113 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v143 = F_LookupFuncNameInternal(m, l0, v137, v80, v12+int32(256), int32(1), l2, v12+int32(252))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L45
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v107+v113<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	if v122 != int32(100) {
		goto L34
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	v126 = v113 + int32(1)
	if v126 != v104 {
		v113 = v126
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v143 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v94 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+252))
	if v145 == int32(1) {
		goto L32
	} else {
		goto L50
	}
L49:
	;
	v446 = v143
	goto L2
L50:
	;
	goto L34
L51:
	;
	v157 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+252))
	switch v158 {
	case 0:
		goto L52
	case 1:
		goto L32
	default:
		v541 = v157
		goto L1
	}
L52:
	;
	if l2 != 0 {
		v541 = v157
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if l0 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L72
	}
L55:
	;
	if l0 != int32(29) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v159&int32(1) != 0 {
		goto L5
	} else {
		goto L65
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v159&int32(1) != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v177 = F_func_signature_string(m, v171, v80, int32(0), v12+int32(256))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v177
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_6), v12+int32(48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2413), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v80 == int32(0) {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v205 = F_func_signature_string(m, v199, v80, int32(0), v12+int32(256))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v205
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_7), v12+int32(96))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2432), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
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
	F_errcode(m, int32(52461700))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v159&int32(1) != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v231 = F_func_signature_string(m, v225, v80, int32(0), v12+int32(256))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v231
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_8), v12+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2447), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
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
	goto L32
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L13
	} else {
		goto L110
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L13
	} else {
		goto L101
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L13
	} else {
		goto L92
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v265 = F_NameListToString(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v265
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_9), v12+int32(112))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v273 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_errhint(m, int32(_a_F_LookupFuncWithArgs_10), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2461), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L13
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v293 = F_NameListToString(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L13
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v293
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_11), v12+int32(128))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L13
	} else {
		goto L95
	}
L95:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v301 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_errhint(m, int32(_a_F_LookupFuncWithArgs_12), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L13
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2469), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L13
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v321 = F_NameListToString(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v321
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_13), v12+int32(144))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v329 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_errhint(m, int32(_a_F_LookupFuncWithArgs_14), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2477), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v349 = F_NameListToString(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v349
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_15), v12+int32(160))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v357 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_errhint(m, int32(_a_F_LookupFuncWithArgs_16), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2485), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2230), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v383
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_17), v12+int32(32))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2407), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v404 = F_NameListToString(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L13
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v404
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_18), v12-int32(-64))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2421), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L13
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v417
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_19), v12+int32(80))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2426), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v430
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_20), v12)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2441), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v508 = F_get_func_prokind(m, v446)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L13
	} else {
		goto L152
	}
L136:
	;
	v480 = F_get_func_prokind(m, v446)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L145
	}
L137:
	;
	v452 = F_get_func_prokind(m, v446)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	if v452 != int32(112) {
		v541 = v446
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v467 = F_func_signature_string(m, v463, v80, int32(0), v12+int32(256))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v467
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_21), v12+int32(176))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2360), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	if v480 == int32(112) {
		v541 = v446
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v495 = F_func_signature_string(m, v491, v80, int32(0), v12+int32(256))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L13
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v495
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_22), v12+int32(192))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2370), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	if v508 == int32(97) {
		v541 = v446
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v523 = F_func_signature_string(m, v519, v80, int32(0), v12+int32(256))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v523
	F_errmsg(m, int32(_a_F_LookupFuncWithArgs_23), v12+int32(208))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_LookupFuncWithArgs_4), int32(2380), int32(_a_F_LookupFuncWithArgs_5))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_get_detail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	v16 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v16
	*(*uint8)(unsafe.Add(mBase, uint32(l10))) = uint8(v16)
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v16
	if l14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v36 = int32(0)
	v39 = F_FuncnameGetCandidates(m, l0, l3, l2, l5, l6, l7, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L9
	} else {
		goto L149
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L9
	} else {
		goto L146
	}
L6:
	;
	m.G0 = v20 + int32(32)
	return v598
L7:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	if v286 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L8:
	;
	if l2|base.B2i32(l1 == int32(0))|base.B2i32(l3 != int32(1)) != 0 {
		goto L37
	} else {
		goto L38
	}
L9:
	;
	return int32(0)
L10:
	;
	v43 = int32(0)
	if base.B2i32(l3 == v36)|base.B2i32(v39 == v43) == v43 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = l3 << (uint(int32(2)) % 32)
	v57 = v39
	goto L14
L12:
	;
	goto L13
L13:
	;
	if v39 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L14:
	;
	v68 = v57 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v49) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	goto L8
L16:
	;
	if v130 == int32(0) {
		v276 = v57
		goto L7
	} else {
		goto L34
	}
L17:
	;
	v130 = int32(0)
	goto L16
L18:
	;
	v104 = v99
	v105 = v100
	v106 = v101
	goto L28
L19:
	;
	if (l4|v68)&int32(3) != 0 {
		v99 = l4
		v100 = v68
		v101 = v49
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v92 = l4
	v93 = v68
	v94 = v49
	goto L21
L21:
	;
	if v94 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v76 = l4
	v77 = v68
	v78 = v49
	goto L23
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v81 != v82 {
		v99 = v76
		v100 = v77
		v101 = v78
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v92 = v87
	v93 = v85
	v94 = v89
	goto L21
L25:
	;
	v84 = int32(4)
	v85 = v77 + v84
	v87 = v76 + v84
	v89 = v78 - v84
	if base.Ui32(int32(3)) < base.Ui32(v89) {
		v76 = v87
		v77 = v85
		v78 = v89
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v99 = v92
	v100 = v93
	v101 = v94
	goto L18
L28:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v109 == v110 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v130 = v109 - v110
	goto L16
L30:
	;
	v112 = int32(1)
	v117 = v106 - v112
	if v117 != 0 {
		v104 = v104 + v112
		v105 = v105 + v112
		v106 = v117
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v133 != 0 {
		v57 = v133
		goto L14
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	v276 = v39
	goto L7
L37:
	;
	v228 = int32(0)
	if v39 == v228 {
		v598 = v228
		goto L6
	} else {
		goto L65
	}
L38:
	;
	v160 = F_makeTypeNameFromNameList(m, l0)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v162 = int32(0)
	v165 = F_LookupTypeNameExtended(m, int32(0), v160, v162, v162, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	if v165 == int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+22)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v170)+82)))
	if v172 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v184 = F_typeTypeId(m, v165)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L48
	}
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+22)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175+v176)+84))
	if v178 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_ReleaseCatCache(m, v165)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	goto L37
L48:
	;
	F_ReleaseCatCache(m, v165)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	if v184 == int32(0) {
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v190 == int32(705) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v215
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v184
	*(*uint8)(unsafe.Add(mBase, uint32(l10))) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v215
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v215
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = l4
	v598 = int32(6)
	goto L6
L52:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v195 == int32(7) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v201 = F_find_coercion_pathway(m, v184, v190, int32(3), v20+int32(28))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	if v190 != int32(2249) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	switch v201 - int32(2) {
	case 0:
		goto L51
	default:
		goto L37
	case 2:
		goto L56
	}
L58:
	;
	v207 = F_typeOrDomainTypeRelid(m, v190)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v211 = F_TypeCategory(m, v184)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	if v207 == int32(0) {
		goto L51
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	if v211 == int32(83) {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	v232 = int32(0)
	v238 = v39
	v239 = v228
	goto L66
L66:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v253 = F_can_coerce_type(m, l3, l4, v238+int32(32), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L9
	} else {
		goto L68
	}
L67:
	;
	if v258 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	if v253 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v239
	v258 = v232 + int32(1)
	v259 = v238
	goto L71
L70:
	;
	v258 = v232
	v259 = v239
	goto L71
L71:
	;
	if v249 != 0 {
		v232 = v258
		v238 = v249
		v239 = v259
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	if v258 < int32(2) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if v259 != 0 {
		v276 = v259
		goto L7
	} else {
		goto L81
	}
L76:
	;
	v598 = int32(0)
	goto L6
L77:
	;
	goto L78
L78:
	;
	v265 = F_func_select_candidate(m, l3, l4, v259)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	if v265 != 0 {
		v276 = v265
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v598 = int32(1)
	goto L6
L81:
	;
	v598 = int32(0)
	goto L6
L82:
	;
	v598 = int32(1)
	goto L6
L83:
	;
	goto L84
L84:
	;
	v290 = int32(0)
	if l5|(base.B2i32(l2 == v290)|base.B2i32(l3 <= v290)) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v286
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v276)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v307
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v276 + int32(32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	v313 = int32(0)
	if base.B2i32(v312 == v313)|base.B2i32(l1 == v313) != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	v298 = l3 - int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296+v298<<(uint(int32(2))%32))))
	if v302 == v298 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v598 = int32(0)
	goto L6
L88:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	v374 = F_SearchSysCache1(m, int32(47), v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L97
	}
L89:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v318 <= int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v328 = int32(0)
	goto L91
L91:
	;
	v340 = v328 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v344 == int32(16) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L88
L93:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v347+v340)))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+12)) = v349
	goto L95
L94:
	;
	goto L95
L95:
	;
	v352 = v328 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v352 < v353 {
		v328 = v352
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	if v374 == int32(0) {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)+16))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+22)))
	v380 = v378 + v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v381
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l10))) = uint8(v383)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v380)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v385
	if l14 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+96)))
	switch v566 - int32(97) {
	case 0:
		v588 = int32(4)
		goto L137
	default:
		goto L139
	case 5:
		goto L138
	case 15:
		goto L141
	case 22:
		goto L140
	}
L100:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	if v389 <= int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v392 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380)+106)))
	if v392 < v389 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v396 = F_SysCacheGetAttrNotNull(m, int32(47), v374, int32(24))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v398 = F_text_to_cstring(m, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v400 = F_stringToNode(m, v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_pfree(m, v398)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	if v404 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = v532
	goto L99
L108:
	;
	v405 = int32(0)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	if v405 < v407 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	if v400 != 0 {
		goto L130
	} else {
		goto L131
	}
L111:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v276)+16))
	v419 = v405
	v422 = int32(0)
	goto L114
L112:
	;
	v446 = v405
	goto L113
L113:
	;
	if v400 == int32(0) {
		v504 = v405
		goto L118
	} else {
		goto L119
	}
L114:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v404+(v410-v407)<<(uint(int32(2))%32)+v422<<(uint(int32(2))%32))))
	v437 = F_bms_add_member(m, v419, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L9
	} else {
		goto L116
	}
L115:
	;
	v446 = v437
	goto L113
L116:
	;
	v440 = v422 + int32(1)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	if v440 < v441 {
		v419 = v437
		v422 = v440
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	F_bms_free(m, v446)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L129
	}
L119:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v462 <= int32(0) {
		v504 = v405
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v466 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380)+106)))
	v471 = v405
	v475 = int32(0)
	v476 = v465 - v466
	goto L121
L121:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v487 = F_bms_is_member(m, v476, v446)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L9
	} else {
		goto L123
	}
L122:
	;
	v504 = v495
	goto L118
L123:
	;
	if v487 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486+v475<<(uint(int32(2))%32))))
	v493 = F_lappend(m, v471, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L9
	} else {
		goto L127
	}
L125:
	;
	v495 = v471
	goto L126
L126:
	;
	v496 = int32(1)
	v499 = v475 + v496
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v499 < v500 {
		v471 = v495
		v475 = v499
		v476 = v476 + v496
		goto L121
	} else {
		goto L128
	}
L127:
	;
	v495 = v493
	goto L126
L128:
	;
	goto L122
L129:
	;
	v532 = v504
	goto L107
L130:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v523 = v521
	goto L132
L131:
	;
	v523 = int32(0)
	goto L132
L132:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v525 = v523 - v524
	if v525 <= int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v532 = v400
	goto L107
L134:
	;
	goto L135
L135:
	;
	v528 = F_list_delete_first_n(m, v400, v525)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	v532 = v528
	goto L107
L137:
	;
	F_ReleaseCatCache(m, v374)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L9
	} else {
		goto L145
	}
L138:
	;
	v588 = int32(2)
	goto L137
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L9
	} else {
		goto L142
	}
L140:
	;
	v588 = int32(5)
	goto L137
L141:
	;
	v588 = int32(3)
	goto L137
L142:
	;
	v575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v380)+96)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v575
	F_errmsg_internal(m, int32(_a_F_func_get_detail_0), v20+int32(16))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_func_get_detail_1), int32(1713), int32(_a_F_func_get_detail_2))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v598 = v588
	goto L6
L146:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v616
	F_errmsg_internal(m, int32(_a_F_func_get_detail_3), v20)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_func_get_detail_1), int32(1627), int32(_a_F_func_get_detail_2))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errmsg_internal(m, int32(_a_F_func_get_detail_4), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_func_get_detail_1), int32(1641), int32(_a_F_func_get_detail_2))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v5 = F_NameListToString(m, l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_funcname_signature_string(m, v5, l1, l2, l3)
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_get_func_leakproof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_leakproof_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_func_leakproof_1), int32(1984), int32(_a_F_get_func_leakproof_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+98)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_func_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13917(m, l0, int32(47))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_func_prokind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_prokind_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_func_prokind_1), int32(1965), int32(_a_F_get_func_prokind_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+96)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_func_retset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_retset_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_func_retset_1), int32(1889), int32(_a_F_get_func_retset_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+100)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_func_rettype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_rettype_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_func_rettype_1), int32(1802), int32(_a_F_get_func_rettype_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+108))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
