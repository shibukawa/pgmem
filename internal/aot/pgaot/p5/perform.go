package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_perform_relmap_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(528)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[0]))
	v22 = F_LWLockAcquire(m, v18+int32(3200), v3)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v47 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	F_read_relmap_file(m, int32(_a_F_perform_relmap_update_0), int32(_a_F_perform_relmap_update_1), int32(1), int32(22))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[1]))
	F_read_relmap_file(m, int32(_a_F_perform_relmap_update_2), v37, int32(1), int32(22))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	base.MemoryCopy(m, v15+int32(4), int32(_a_F_perform_relmap_update_0), int32(524))
	goto L3
L8:
	;
	base.MemoryCopy(m, v15+int32(4), int32(_a_F_perform_relmap_update_2), int32(524))
	goto L3
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L46
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L43
	}
L11:
	;
	v53 = v15 + int32(12)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_relmap_update[2])))
	v62 = v54
	v65 = v3
	goto L14
L12:
	;
	goto L13
L13:
	;
	v153 = v15 + int32(4)
	v154 = int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[3]))
	if l0 != 0 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v71 = int32(0)
	v74 = l1 + int32(8) + v65<<(uint(int32(3))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v62 <= v71 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	v138 = v65 + int32(1)
	if v138 != v47 {
		v62 = v128
		v65 = v138
		goto L14
	} else {
		goto L27
	}
L17:
	;
	if v56&int32(1) == int32(0) {
		goto L10
	} else {
		goto L25
	}
L18:
	;
	v80 = v71
	goto L19
L19:
	;
	v93 = v53 + v80<<(uint(int32(3))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != v76 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v75
	v128 = v62
	goto L16
L21:
	;
	v97 = v80 + int32(1)
	if v62 != v97 {
		v80 = v97
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L17
L25:
	;
	if int32(64) <= v62 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v118 = v53 + v62<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v76
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v123 = v121 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v123
	v128 = v123
	goto L16
L27:
	;
	goto L15
L28:
	;
	v160 = int32(0)
	goto L30
L29:
	;
	v160 = v159
	goto L30
L30:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[4]))
	if l0 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v164 = int32(1664)
	goto L33
L32:
	;
	v164 = v163
	goto L33
L33:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[1]))
	if l0 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v168 = int32(_a_F_perform_relmap_update_1)
	goto L36
L35:
	;
	v168 = v167
	goto L36
L36:
	;
	F_write_relmap_file(m, v153, v154, v154, v154, v160, v164, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if l0 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_perform_relmap_update[0]))
	F_LWLockRelease(m, v180+int32(3200))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L42
	}
L39:
	;
	base.MemoryCopy(m, int32(_a_F_perform_relmap_update_0), v153, int32(524))
	goto L38
L40:
	;
	goto L41
L41:
	;
	base.MemoryCopy(m, int32(_a_F_perform_relmap_update_2), v15+int32(4), int32(524))
	goto L38
L42:
	;
	m.G0 = v15 + int32(528)
	return
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v76
	F_errmsg_internal(m, int32(_a_F_perform_relmap_update_3), v15)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_perform_relmap_update_4), int32(402), int32(_a_F_perform_relmap_update_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_perform_relmap_update_6), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_perform_relmap_update_4), int32(404), int32(_a_F_perform_relmap_update_5))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_perform_work_item(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v204 int64
	_ = v204
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int64
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v316 int32
	_ = v316
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
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
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v532 int32
	_ = v532
	var v547 int32
	_ = v547
	var v548 int64
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(480)
	m.G0 = v21
	v24 = l0 + int32(12)
	v26 = l0 + int32(16)
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = v2
	v38 = v2
	v39 = v2
	v41 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v21 + int32(480)
	return
L3:
	;
	goto L2
L4:
	;
	if v41 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v547 = int32(m.ExcTag)
	v548 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v547 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L7:
	;
	v504 = v496 & int32(1)
	if v504 != 0 {
		goto L59
	} else {
		goto L60
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v462
	v481 = int32(1)
	v482 = v467 & v481
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v482)
	v485 = v468 & v481
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v485)
	F_pfree(m, v463)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L58
	}
L9:
	;
	if v98 == int32(0) {
		v489 = v85
		v490 = v62
		v491 = v24
		v492 = v98
		v493 = v34
		v494 = v35
		v495 = v36
		v496 = v101
		v497 = v103
		goto L7
	} else {
		goto L57
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	v52 = int32(1)
	v53 = v37 & v52
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v53)
	v56 = v38 & v52
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v62 = F_get_rel_name(m, v48)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v242 = v30
	v243 = v31
	v244 = v32
	v245 = v33
	v246 = v34
	v247 = v35
	v248 = v36
	v249 = v37
	v250 = v38
	v251 = v39
	goto L12
L12:
	;
	if v251 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v74 = F_get_rel_namespace(m, v64)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v56)
	v85 = F_get_namespace_name(m, v74)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v56)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[0]))
	v98 = F_get_database_name(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v100 = int32(0)
	v101 = base.B2i32(v85 != v100)
	v103 = base.B2i32(v62 != v100)
	if base.B2i32(v98 == v100)|(base.B2i32(v62 == v100)|base.B2i32(v85 == v100)) != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v112 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v129 = F_pg_snprintf(m, v21+int32(240), int32(184), int32(_a_F_perform_work_item_0), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v142 = F_strlen(m, v21+int32(240))
	mBase = m.M
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v143 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v21 + int32(226)
	v181 = v21 + int32(240)
	v188 = F_pg_snprintf(m, v181+v142, int32(184)-v142, int32(_a_F_perform_work_item_1), v21+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v143
	v162 = F_pg_snprintf(m, v21+int32(226), int32(14), int32(_a_F_perform_work_item_2), v21+int32(48))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+226)) = uint8(v164)
	goto L22
L26:
	;
	goto L22
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[1]))
	if v200 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	F_pgstat_report_activity(m, int32(3), v181)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v101)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v24
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	F_MemoryContextReset(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L32
	}
L29:
	;
	v204 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_perform_work_item[3])) = v204
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4]))
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5]))
	goto L33
L33:
	;
	v236 = v21 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v21 + int32(60)
	goto L36
L34:
	;
	v242 = v85
	v243 = v62
	v244 = v24
	v245 = v98
	v246 = v234
	v247 = v232
	v248 = v26
	v249 = v101
	v250 = v103
	v251 = int32(0)
	goto L12
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4])) = v247
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v246
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[7])) = v456
	v460 = v242
	v461 = v243
	v462 = v244
	v463 = v245
	v464 = v246
	v465 = v247
	v466 = v248
	v467 = v249
	v468 = v250
	goto L8
L38:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[7])) = v256
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v21 - int32(-64)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v262 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4])) = v247
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v246
	v343 = int32(_a_F_perform_work_item_3)
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8]))
	v346 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8])) = v345 + v346
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	v357 = v249 & v346
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	v360 = v250 & v346
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L6
	} else {
		goto L50
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[9])) = int32(0)
	goto L37
L42:
	;
	v265 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v248))))
	v266 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v244))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	v270 = int32(1)
	v271 = v249 & v270
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v271)
	v274 = v250 & v270
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v274)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	v282 = F_DirectFunctionCall2Coll(m, int32(4), int32(0), v266, v265)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	v291 = int32(1)
	v292 = v249 & v291
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v292)
	v295 = v250 & v291
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v295)
	v299 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	if v299 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v292)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v303
	F_errmsg_internal(m, int32(_a_F_perform_work_item_4), v21)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v292)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v295)
	F_errfinish(m, int32(_a_F_perform_work_item_5), int32(2729), int32(_a_F_perform_work_item_6))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v245
	F_errcontext_msg(m, int32(_a_F_perform_work_item_7), v21+int32(16))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	F_EmitErrorReport(m)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	F_FlushErrorState(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	F_MemoryContextReset(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v360)
	F_StartTransactionCommand(m)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v439 = int32(_a_F_perform_work_item_3)
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8])) = v441 - int32(1)
	goto L37
L57:
	;
	v460 = v85
	v461 = v62
	v462 = v24
	v463 = v98
	v464 = v34
	v465 = v35
	v466 = v36
	v467 = v101
	v468 = v103
	goto L8
L58:
	;
	v489 = v460
	v490 = v461
	v491 = v462
	v492 = v463
	v493 = v464
	v494 = v465
	v495 = v466
	v496 = v467
	v497 = v468
	goto L7
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v491
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v504)
	v514 = v497 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v514)
	F_pfree(m, v489)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v519 = v497 & int32(1)
	if v519 == int32(0) {
		goto L3
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+452)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v21)+448)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+456)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+464)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v21)+468)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v21)+472)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v21)+476)) = v491
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)) = uint8(v504)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)) = uint8(v519)
	F_pfree(m, v490)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	goto L5
L65:
	;
	v552 = int32(v548)
	m.G0 = v21
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v21+int32(60) == v558 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	m.ExcPending = 1
	goto L74
L67:
	;
	if v562 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	v562 = v560
	goto L70
L69:
	;
	v562 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v21)+476))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v21)+472))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v21)+468))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v21)+464))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+463)))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+462)))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v21)+456))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v21)+452))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v21)+448))
	v30 = v565
	v31 = v564
	v32 = v563
	v33 = v566
	v34 = v570
	v35 = v571
	v36 = v569
	v37 = v568
	v38 = v567
	v39 = v554
	v41 = v562
	goto L1
L72:
	;
	goto L73
L73:
	;
	F___wasm_longjmp(m, v555, v554)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	return
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
