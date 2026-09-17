package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSimpleRelationInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v12 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v22 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v25 == v22 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
	if v15 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = F_ExecBRInsertTriggers(m, l1, l0, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v18 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+131)))
	if v41 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+17)))
	if v28 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ExecComputeStoredGenerated(m, l0, l1, l2, int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_ExecConstraints(m, l0, l2, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L8
L16:
	;
	v45 = F_ExecPartitionCheck(m, l0, l2, l1, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+188))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	m.T0[v54].(func(*base.Module, int32, int32, int32, int32, int32))(m, v47, l2, v49, v51, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v58 <= v57 {
		v77 = v57
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_ExecARInsertTriggers(m, l1, l0, l2, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L27
	}
L23:
	;
	v61 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v68 = F_ExecInsertIndexTuples(m, l0, l2, l1, v61, base.B2i32(v62 != v61), v9+int32(15), v62, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v70 != int32(1) {
		v77 = v68
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v73 = int32(0)
	F_CheckAndReportConflict(m, l0, l1, v73, v68, v73, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v77 = v68
	goto L22
L27:
	;
	F_list_free(m, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L1
}
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v379 int64
	_ = v379
	var v387 int32
	_ = v387
	v15 = m.G0
	v17 = v15 - int32(1072)
	m.G0 = v17
	v19 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v20 = base.I64_rem_s(l1, v19)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v23 = F_SlruSelectLRUPage(m, l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	if v31 == int32(0) {
		v112 = v23
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(1072)
	return v387
L4:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v122+v112<<(uint(int32(3))%32)))) = l1
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v112<<(uint(int32(2))%32)))) = int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v112))) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v142 = F_LWLockAcquire(m, v137+v112<<(uint(int32(7))%32), v135)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v38 = v23
	v40 = v31
	goto L6
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v48+v38<<(uint(int32(3))%32))))
	if v52 != l1 {
		v112 = v38
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v112 = v101
	goto L4
L8:
	;
	v57 = int32(0)
	if base.B2i32(l2|base.B2i32(v40 != int32(3)) == v57)|base.B2i32(v40 == int32(1)) == v57 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v67 = int32(2)
	v69 = v64 + v38>>(uint(int32(4))%32)<<(uint(v67)%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v72 = v38 << (uint(v67) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72+v73)))
	if v70 != v75 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_SimpleLruWaitIO(m, l0, v38)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v78 = v70 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v72))) = v78
	goto L14
L13:
	;
	goto L14
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[0])) = uint8(v86)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[1])) = uint8(v86)
	v92 = v84 << (uint(int32(6)) % 32)
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_SimpleLruReadPage[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_SimpleLruReadPage[2]))) = v95 + int64(1)
	goto L15
L15:
	;
	v387 = v38
	goto L3
L16:
	;
	v101 = F_SlruSelectLRUPage(m, l0, l1)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v101<<(uint(int32(2))%32))))
	if v107 != 0 {
		v38 = v101
		v40 = v107
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	v147 = v22 + base.I32_wrap_i64(v20)<<(uint(int32(7))%32)
	F_LWLockRelease(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v151 = l0 + int32(16)
	v153 = base.I64_div_s(l1, int64(32))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v155 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v181 = F_OpenTransientFile(m, v17+int32(48), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v151
	v166 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(_a_F_SimpleLruReadPage_0), v17+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+36)) = uint32(v153)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v151
	v176 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(_a_F_SimpleLruReadPage_1), v17+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	goto L21
L27:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	if v276 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L28:
	;
	if v181 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	if v186 == int32(44) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v224 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3])) = v224
	v226 = int32(_a_F_SimpleLruReadPage_2)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = int32(167772211)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230+v112<<(uint(int32(2))%32))))
	v235 = int32(_a_F_SimpleLruReadPage_3)
	v243 = F_pread(m, v181, v234, v235, base.I64_extend_i32_s(base.I32_wrap_i64(l1-v153<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v224
	if v243 != v235 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v201 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[5])))
	if v190&int32(1) != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v186
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[7])) = v193
	v274 = v193
	goto L27
L36:
	;
	goto L35
L37:
	;
	if v201 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(48)
	F_errmsg(m, int32(_a_F_SimpleLruReadPage_4), v17)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v112<<(uint(int32(2))%32))))
	base.MemoryFill(m, v218, int32(0), int32(_a_F_SimpleLruReadPage_3))
	v274 = int32(1)
	goto L27
L41:
	;
	F_errfinish(m, int32(_a_F_SimpleLruReadPage_5), int32(834), int32(_a_F_SimpleLruReadPage_6))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[7])) = int32(2)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v256
	v258 = F_CloseTransientFile(m, v181)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v260 = F_CloseTransientFile(m, v181)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v274 = int32(0)
	goto L27
L47:
	;
	if v260 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v274 = int32(1)
	goto L27
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[7])) = int32(5)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v271
	v274 = int32(0)
	goto L27
L51:
	;
	v328 = F_LWLockAcquire(m, v147, int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L61
	}
L52:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	v280 = v112 * v276
	v281 = int32(3)
	v283 = v279 + v280<<(uint(v281)%32)
	v287 = v276 << (uint(v281) % 32)
	if v283&v281|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v287)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v287 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	v316 = v287
	goto L55
L55:
	;
	if v316 == int32(0) {
		goto L51
	} else {
		goto L60
	}
L56:
	;
	v295 = int32(3)
	v296 = v280 << (uint(v295) % 32)
	v302 = v296 + v279 + int32(4)
	v308 = v276*(v112<<(uint(v295)%32)+int32(8)) + v279
	if base.Ui32(v308) < base.Ui32(v302) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v310 = v302
	goto L59
L58:
	;
	v310 = v308
	goto L59
L59:
	;
	v316 = (v296^int32(-1)-v279+v310)&int32(-4) + int32(4)
	goto L55
L60:
	;
	base.MemoryFill(m, v283, int32(0), v316)
	goto L51
L61:
	;
	v330 = int32(2)
	v331 = v112 << (uint(v330) % 32)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v274 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v336 = v330
	goto L64
L63:
	;
	v336 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331+v332))) = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	F_LWLockRelease(m, v338+v112<<(uint(int32(7))%32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v274 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_SlruReportIOError(m, l0, l1, l3)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v353 = v348 + v112>>(uint(int32(4))%32)<<(uint(int32(2))%32)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355+v331)))
	if v354 != v357 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v360 = v354 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v362+v112<<(uint(int32(2))%32)))) = v360
	goto L72
L71:
	;
	goto L72
L72:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[0])) = uint8(v370)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[1])) = uint8(v370)
	v376 = v368 << (uint(int32(6)) % 32)
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v376)+uint32(_c_F_SimpleLruReadPage[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v376)+uint32(_c_F_SimpleLruReadPage[8]))) = v379 + int64(1)
	v387 = v112
	goto L3
}
func F_SimpleLruShmemSize(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	v3 = int32(0)
	v8 = int32(7)
	v10 = int32(-8)
	v11 = (l0<<(uint(int32(2))%32) + v8) & v10
	v13 = l0 << (uint(int32(3)) % 32)
	v21 = base.I32_div_s(l0, int32(16))
	if v3 < l1 {
		v30 = l1 * v13
	} else {
		v30 = v3
	}
	return (v11+(v13+(l0+v8)&v10)+(v21+l0)<<(uint(v8)%32)+v30+v11<<(uint(int32(1))%32)+(v21<<(uint(int32(2))%32)+int32(7))&int32(-8)+int32(95))&int32(-32) + l0<<(uint(int32(13))%32)
}
