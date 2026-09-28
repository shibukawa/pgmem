package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecAlterCheckConstrEnforceability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	v20 = m.G0
	v22 = v20 - int32(224)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	v36 = F_table_open(m, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v38 != 0 {
		v47 = v24
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L92
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L87
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L81
	}
L7:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+75)))
	v50 = v47 & int32(1)
	if v48 != v50 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	if l5 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+131)))
	if v40 != 0 {
		v47 = v24
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v43 = F_ATCheckCheckConstrHasEnforcedParent(m, l2, v36, l3, l6, v22+int32(44))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v45 = v43 | v24
	if l5 != 0 {
		v47 = v45
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v43 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v47 = v45
	goto L7
L16:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+119)))
	if v244 != int32(114) {
		goto L59
	} else {
		goto L60
	}
L17:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+106)))
	if v66 != 0 {
		goto L16
	} else {
		goto L25
	}
L18:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v54
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+57)) = uint8(v50)
	F_AlterConstrUpdateConstraintEntry(m, v22+int32(48), l2, l3)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l5 != 0 {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	if l5 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v63&int32(1) != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	v69 = F_find_all_inheritors(m, v67, l7, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v71 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v74
	v80 = F_list_make1_impl(m, int32(480), v22+int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v82 = l6
	goto L29
L29:
	;
	if v69 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L30:
	;
	v82 = v80
	goto L29
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if int32(0) < v85 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = v82
	v99 = int32(0)
	goto L35
L33:
	;
	v135 = v82
	goto L34
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v148 <= int32(0) {
		goto L16
	} else {
		goto L44
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v99<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	if v112 == v113 {
		v124 = v95
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v135 = v124
	goto L34
L37:
	;
	v126 = v99 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v126 < v127 {
		v95 = v124
		v99 = v126
		goto L35
	} else {
		goto L43
	}
L38:
	;
	if l4 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v117 != 0 {
		v124 = v95
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v120 = F_get_relation_constraint_oid(m, v112, v118, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = F_list_append_unique_oid(m, v95, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v124 = v122
	goto L37
L43:
	;
	goto L36
L44:
	;
	v166 = int32(0)
	goto L45
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v166<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	if v179 != v180 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L16
L47:
	;
	v183 = v22 + int32(48)
	F_ScanKeyInit(m, v183, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v179))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v221 = v166 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v221 < v222 {
		v166 = v221
		goto L45
	} else {
		goto L58
	}
L50:
	;
	F_ScanKeyInit(m, v22+int32(104), int32(10), int32(3), int32(184), int64(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v199 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+4)))
	F_ScanKeyInit(m, v22+int32(160), int32(2), int32(3), int32(62), v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v206 = F_systable_beginscan(m, l2, int32(2665), int32(1), int32(0), int32(3), v183)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v208 = F_systable_getnext(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v208 == int32(0) {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v214 = F_ATExecAlterCheckConstrEnforceability(m, l0, l1, l2, v208, int32(0), int32(1), v135, l7)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_systable_endscan(m, v206)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L49
L58:
	;
	goto L46
L59:
	;
	F_relation_close(m, v36, int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L80
	}
L60:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+75)))
	if v47&(v247^int32(-1))&int32(1) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v256 = F_palloc0(m, int32(32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v260 = F_pstrdup(m, v33+int32(4))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v260
	v267 = F_SysCacheGetAttrNotNull(m, int32(19), l3, int32(28))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v270 = F_text_to_cstring(m, base.I32_wrap_i64(v267))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v272 = F_stringToNode(m, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v275 = F_expand_generated_columns_in_expr(m, v272, v36, int32(1))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+24)) = v275
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v279 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v358)+64))
	v377 = F_lappend(m, v376, v256)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L79
	}
L69:
	;
	v335 = F_palloc0(m, int32(144))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L76
	}
L70:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v282 <= int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v292 = int32(0)
	goto L72
L72:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v285+v292<<(uint(int32(2))%32))))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if v310 == v278 {
		v358 = v309
		goto L68
	} else {
		goto L74
	}
L73:
	;
	goto L69
L74:
	;
	v313 = v292 + int32(1)
	if v282 != v313 {
		v292 = v313
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v278
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+4)) = uint8(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v344 = F_CreateTupleDescCopyConstr(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+8)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v335)+88)) = int64(0)
	v349 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+84)) = uint8(v349)
	v351 = int32(_a_F_ATExecAlterCheckConstrEnforceability_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v335)+96)) = uint16(v351)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v354 = F_lappend(m, v353, v335)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v354
	v358 = v335
	goto L68
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+64)) = v377
	goto L59
L80:
	;
	m.G0 = v22 + int32(224)
	return base.B2i32(v50 != v48)
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(_a_F_ATExecAlterCheckConstrEnforceability_1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v33 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_2), v22+int32(32))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v425 = F_get_rel_name(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(_a_F_ATExecAlterCheckConstrEnforceability_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v425
	v433 = F_errdetail(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_4), v22+int32(16))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_5), int32(_a_F_ATExecAlterCheckConstrEnforceability_6), int32(_a_F_ATExecAlterCheckConstrEnforceability_7))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_8), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errhint(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_9), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_5), int32(_a_F_ATExecAlterCheckConstrEnforceability_10), int32(_a_F_ATExecAlterCheckConstrEnforceability_7))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v468 = F_get_rel_name(m, v179)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v467
	F_errmsg(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_11), v22)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ATExecAlterCheckConstrEnforceability_5), int32(_a_F_ATExecAlterCheckConstrEnforceability_12), int32(_a_F_ATExecAlterCheckConstrEnforceability_13))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
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
func F_ATPrepChangeInherit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+76))
	if v3 == int32(0) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+131)))
		if v6 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ATPrepChangeInherit_0), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATPrepChangeInherit_1), int32(_a_F_ATPrepChangeInherit_2), int32(_a_F_ATPrepChangeInherit_3))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_ATPrepChangeInherit_4), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ATPrepChangeInherit_1), int32(_a_F_ATPrepChangeInherit_5), int32(_a_F_ATPrepChangeInherit_3))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_ATTypedTableRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v13 = F_find_typed_table_dependencies(m, v9, v8+int32(4), v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L2:
	;
	return
L3:
	;
	return
L4:
	;
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(0)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v26<<(uint(int32(2))%32))))
	v32 = F_relation_open(m, v31, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L2
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+118)))
	if v35 == int32(116) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
	if v38 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_CheckTableNotInUse(m, v32, int32(_a_F_ATTypedTableRecursion_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v44 = int32(1)
	F_ATPrepCmd(m, l0, v32, l2, v44, v44, l3, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_relation_close(m, v32, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v52 = v26 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v52 < v53 {
		v26 = v52
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_ATTypedTableRecursion_1), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_ATTypedTableRecursion_2), int32(_a_F_ATTypedTableRecursion_3), int32(_a_F_ATTypedTableRecursion_4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
