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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[239]))
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L52
	}
L16:
	;
	v36 = F_palloc0(m, int32(28))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L48
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v36
	v43 = F_palloc0(m, v17<<(uint(int32(3))%32)+int32(20))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_fmgr_info(m, l3, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
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
	if v60 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+10)))
	v119 = v116 & base.B2i32(int32(0) < v17)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[339]))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)))
	if v121 <= v122 {
		goto L35
	} else {
		goto L36
	}
L24:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v64 <= v63 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v68 = v43 + int32(20)
	v72 = v63
	goto L26
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v72<<(uint(int32(2))%32))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 == int32(7) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L23
L28:
	;
	v103 = v72 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v103 < v104 {
		v72 = v103
		goto L26
	} else {
		goto L33
	}
L29:
	;
	v89 = v68 + v72<<(uint(int32(3))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v90
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)) = uint8(v92)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v96 = v68 + v72<<(uint(int32(3))%32)
	F_ExecInitExprRec(m, v83, l5, v96, v96+int32(4))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	goto L27
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v139
	m.G0 = v13 + int32(16)
	return
L35:
	;
	if v119 == int32(0) {
		v139 = int32(26)
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v119 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	if v17 == int32(2) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v132 = int32(29)
	goto L41
L40:
	;
	v132 = int32(27)
	goto L41
L41:
	;
	if v17 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v135 = int32(28)
	goto L44
L43:
	;
	v135 = v132
	goto L44
L44:
	;
	v139 = v135
	goto L34
L45:
	;
	v138 = int32(31)
	goto L47
L46:
	;
	v138 = int32(30)
	goto L47
L47:
	;
	v139 = v138
	goto L34
L48:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v151 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v151
	F_errmsg_plural(m, int32(249843), int32(249891), v151, v13)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(486174), int32(2732), int32(481271))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(104858), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	if v174 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v176 = F_exprLocation(m, l1)
	mBase = m.M
	F_executor_errposition(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_errfinish(m, int32(486174), int32(2759), int32(481271))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
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
	var v81 int32
	_ = v81
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
	var v447 int32
	_ = v447
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
	var v542 int32
	_ = v542
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(656)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 == v4 {
		v81 = v4
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v12 + int32(656)
	return v542
L2:
	;
	switch l0 - int32(1) {
	case 0:
		goto L135
	default:
		v542 = v447
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
	F_errmsg_plural(m, int32(92990), int32(120150), v369, v12+int32(224))
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
		v81 = v17
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
	F_errmsg_plural(m, int32(92901), int32(120059), v31, v12+int32(240))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(491002), int32(2237), int32(152881))
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
	v81 = v17
	goto L8
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+(v12+int32(256))))) = v65
	if v65 == int32(0) {
		v542 = v4
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
	v88 = v81
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
		v542 = int32(0)
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
		v447 = v143
		goto L2
	} else {
		goto L78
	}
L34:
	;
	if v94 != 0 {
		v447 = v94
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
	v143 = F_LookupFuncNameInternal(m, l0, v137, v81, v12+int32(256), int32(1), l2, v12+int32(252))
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
	v447 = v143
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
		v542 = v157
		goto L1
	}
L52:
	;
	if l2 != 0 {
		v542 = v157
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
	v177 = F_func_signature_string(m, v171, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(69051), v12+int32(48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(491002), int32(2413), int32(152881))
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
	if v81 == int32(0) {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v205 = F_func_signature_string(m, v199, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(69023), v12+int32(96))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(491002), int32(2432), int32(152881))
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
	v231 = F_func_signature_string(m, v225, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(68720), v12+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(491002), int32(2447), int32(152881))
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
	F_errmsg(m, int32(338335), v12+int32(112))
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
	F_errhint(m, int32(550214), int32(0))
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
	F_errfinish(m, int32(491002), int32(2461), int32(152881))
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
	F_errmsg(m, int32(338433), v12+int32(128))
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
	F_errhint(m, int32(550343), int32(0))
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
	F_errfinish(m, int32(491002), int32(2469), int32(152881))
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
	F_errmsg(m, int32(338399), v12+int32(144))
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
	F_errhint(m, int32(550278), int32(0))
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
	F_errfinish(m, int32(491002), int32(2477), int32(152881))
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
	F_errmsg(m, int32(338467), v12+int32(160))
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
	F_errhint(m, int32(550408), int32(0))
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
	F_errfinish(m, int32(491002), int32(2485), int32(152881))
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
	F_errfinish(m, int32(491002), int32(2230), int32(152881))
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
	F_errmsg(m, int32(697414), v12+int32(32))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(491002), int32(2407), int32(152881))
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
	F_errmsg(m, int32(697375), v12-int32(-64))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(491002), int32(2421), int32(152881))
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
	F_errmsg(m, int32(69263), v12+int32(80))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(491002), int32(2426), int32(152881))
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
	F_errmsg(m, int32(697338), v12)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(491002), int32(2441), int32(152881))
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
	v508 = F_get_func_prokind(m, v447)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L13
	} else {
		goto L152
	}
L136:
	;
	v480 = F_get_func_prokind(m, v447)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L145
	}
L137:
	;
	v452 = F_get_func_prokind(m, v447)
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
		v542 = v447
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
	v467 = F_func_signature_string(m, v463, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(249822), v12+int32(176))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(491002), int32(2360), int32(152881))
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
		v542 = v447
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
	v495 = F_func_signature_string(m, v491, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(357489), v12+int32(192))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(491002), int32(2370), int32(152881))
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
		v542 = v447
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
	v523 = F_func_signature_string(m, v519, v81, int32(0), v12+int32(256))
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
	F_errmsg(m, int32(348862), v12+int32(208))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(491002), int32(2380), int32(152881))
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
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
	v37 = F_FuncnameGetCandidates(m, l0, l3, l2, l5, l6, l7, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L154
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L151
	}
L8:
	;
	m.G0 = v20 + int32(32)
	return v591
L9:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	if v282 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L10:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L11:
	;
	if v37 == int32(0) {
		goto L10
	} else {
		goto L36
	}
L12:
	;
	if v37 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v46 = l3 << (uint(int32(2)) % 32)
	v54 = v37
	goto L14
L14:
	;
	v65 = v54 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	goto L10
L16:
	;
	if v127 == int32(0) {
		v272 = v54
		goto L9
	} else {
		goto L34
	}
L17:
	;
	v127 = int32(0)
	goto L16
L18:
	;
	v101 = v96
	v102 = v97
	v103 = v98
	goto L28
L19:
	;
	if (l4|v65)&int32(3) != 0 {
		v96 = l4
		v97 = v65
		v98 = v46
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v89 = l4
	v90 = v65
	v91 = v46
	goto L21
L21:
	;
	if v91 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v73 = l4
	v74 = v65
	v75 = v46
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v78 != v79 {
		v96 = v73
		v97 = v74
		v98 = v75
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v89 = v84
	v90 = v82
	v91 = v86
	goto L21
L25:
	;
	v81 = int32(4)
	v82 = v74 + v81
	v84 = v73 + v81
	v86 = v75 - v81
	if base.Ui32(int32(3)) < base.Ui32(v86) {
		v73 = v84
		v74 = v82
		v75 = v86
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v96 = v89
	v97 = v90
	v98 = v91
	goto L18
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 == v107 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v127 = v106 - v107
	goto L16
L30:
	;
	v109 = int32(1)
	v114 = v103 - v109
	if v114 != 0 {
		v101 = v101 + v109
		v102 = v102 + v109
		v103 = v114
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
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v130 != 0 {
		v54 = v130
		goto L14
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	v272 = v37
	goto L9
L37:
	;
	v224 = int32(0)
	if v37 == v224 {
		v591 = v224
		goto L8
	} else {
		goto L67
	}
L38:
	;
	if l1 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if l3 != int32(1) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v155 = F_makeTypeNameFromNameList(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v157 = int32(0)
	v160 = F_LookupTypeNameExtended(m, int32(0), v155, v157, v157, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v160 == int32(0) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+22)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v165)+82)))
	if v167 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v179 = F_typeTypeId(m, v160)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L50
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+22)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170+v171)+84))
	if v173 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_ReleaseCatCache(m, v160)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	goto L37
L50:
	;
	F_ReleaseCatCache(m, v160)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	if v179 == int32(0) {
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v185 == int32(705) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v210 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v179
	*(*uint8)(unsafe.Add(mBase, uint32(l10))) = uint8(v210)
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = l4
	v591 = int32(6)
	goto L8
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v190 == int32(7) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v196 = F_find_coercion_pathway(m, v179, v185, int32(3), v20+int32(28))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L56
L58:
	;
	if v185 != int32(2249) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	switch v196 - int32(2) {
	case 0:
		goto L53
	default:
		goto L37
	case 2:
		goto L58
	}
L60:
	;
	v202 = F_typeOrDomainTypeRelid(m, v185)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v206 = F_TypeCategory(m, v179)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	if v202 == int32(0) {
		goto L53
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	if v206 == int32(83) {
		goto L37
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v228 = int32(0)
	v234 = v37
	v235 = v224
	goto L68
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v249 = F_can_coerce_type(m, l3, l4, v234+int32(32), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	if v254 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	if v249 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v235
	v254 = v228 + int32(1)
	v255 = v234
	goto L73
L72:
	;
	v254 = v228
	v255 = v235
	goto L73
L73:
	;
	if v245 != 0 {
		v228 = v254
		v234 = v245
		v235 = v255
		goto L68
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	if v254 < int32(2) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v255 != 0 {
		v272 = v255
		goto L9
	} else {
		goto L83
	}
L78:
	;
	v591 = int32(0)
	goto L8
L79:
	;
	goto L80
L80:
	;
	v261 = F_func_select_candidate(m, l3, l4, v255)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	if v261 != 0 {
		v272 = v261
		goto L9
	} else {
		goto L82
	}
L82:
	;
	v591 = int32(1)
	goto L8
L83:
	;
	v591 = int32(0)
	goto L8
L84:
	;
	v591 = int32(1)
	goto L8
L85:
	;
	goto L86
L86:
	;
	if l2 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v282
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v301
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v272 + int32(32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	if v306 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	if l3 <= int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if l5 != 0 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	v292 = l3 - int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290+v292<<(uint(int32(2))%32))))
	if v296 == v292 {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v591 = int32(0)
	goto L8
L92:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v367 = F_SearchSysCache1(m, int32(47), v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L102
	}
L93:
	;
	if l1 == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v311 <= int32(0) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v321 = int32(0)
	goto L96
L96:
	;
	v333 = v321 << (uint(int32(2)) % 32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333+v334)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v337 == int32(16) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L92
L98:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v340+v333)))
	*(*int32)(unsafe.Add(mBase, uint32(v336)+12)) = v342
	goto L100
L99:
	;
	goto L100
L100:
	;
	v345 = v321 + int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v345 < v346 {
		v321 = v345
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	if v367 == int32(0) {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+22)))
	v373 = v371 + v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v374
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l10))) = uint8(v376)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v373)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v378
	if l14 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+96)))
	switch v559 - int32(97) {
	case 0:
		v581 = int32(4)
		goto L142
	default:
		goto L144
	case 5:
		goto L143
	case 15:
		goto L146
	case 22:
		goto L145
	}
L105:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v272)+24))
	if v382 <= int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v385 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373)+106)))
	if v385 < v382 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v389 = F_SysCacheGetAttrNotNull(m, int32(47), v367, int32(24))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v391 = F_text_to_cstring(m, v389)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v393 = F_stringToNode(m, v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_pfree(m, v391)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	if v397 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = v525
	goto L104
L113:
	;
	v398 = int32(0)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v272)+24))
	if v398 < v400 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if v393 != 0 {
		goto L135
	} else {
		goto L136
	}
L116:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	v415 = int32(0)
	v424 = v398
	goto L119
L117:
	;
	v451 = v398
	goto L118
L118:
	;
	if v393 == int32(0) {
		v497 = v398
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v397+(v403-v400)<<(uint(int32(2))%32)+v415<<(uint(int32(2))%32))))
	v430 = F_bms_add_member(m, v424, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L121
	}
L120:
	;
	v451 = v430
	goto L118
L121:
	;
	v433 = v415 + int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v272)+24))
	if v433 < v434 {
		v415 = v433
		v424 = v430
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	F_bms_free(m, v451)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L134
	}
L124:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v455 <= int32(0) {
		v497 = v398
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v459 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373)+106)))
	v464 = v398
	v468 = int32(0)
	v469 = v458 - v459
	goto L126
L126:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v480 = F_bms_is_member(m, v469, v451)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	v497 = v488
	goto L123
L128:
	;
	if v480 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479+v468<<(uint(int32(2))%32))))
	v486 = F_lappend(m, v464, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	v488 = v464
	goto L131
L131:
	;
	v489 = int32(1)
	v492 = v468 + v489
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v492 < v493 {
		v464 = v488
		v468 = v492
		v469 = v469 + v489
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v488 = v486
	goto L131
L133:
	;
	goto L127
L134:
	;
	v525 = v497
	goto L112
L135:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v516 = v514
	goto L137
L136:
	;
	v516 = int32(0)
	goto L137
L137:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v272)+24))
	v518 = v516 - v517
	if v518 <= int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v525 = v393
	goto L112
L139:
	;
	goto L140
L140:
	;
	v521 = F_list_delete_first_n(m, v393, v518)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v525 = v521
	goto L112
L142:
	;
	F_ReleaseCatCache(m, v367)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L150
	}
L143:
	;
	v581 = int32(2)
	goto L142
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	v581 = int32(5)
	goto L142
L146:
	;
	v581 = int32(3)
	goto L142
L147:
	;
	v568 = int32(*(*int8)(unsafe.Add(mBase, uint32(v373)+96)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v568
	F_errmsg_internal(m, int32(493493), v20+int32(16))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(491002), int32(1713), int32(300464))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	v591 = v581
	goto L8
L151:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v609
	F_errmsg_internal(m, int32(44089), v20)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(491002), int32(1627), int32(300464))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errmsg_internal(m, int32(118804), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(491002), int32(1641), int32(300464))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
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
				F_errmsg_internal(m, int32(44089), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490208), int32(1984), int32(332449))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
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
				F_errmsg_internal(m, int32(44089), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490208), int32(1965), int32(417941))
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
				F_errmsg_internal(m, int32(44089), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490208), int32(1889), int32(103614))
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
				F_errmsg_internal(m, int32(44089), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490208), int32(1802), int32(359224))
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
