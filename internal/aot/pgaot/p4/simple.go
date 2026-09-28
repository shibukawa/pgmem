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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
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
		v76 = v57
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_ExecARInsertTriggers(m, l1, l0, l2, v76, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L30
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = int32(2)
	goto L26
L25:
	;
	v64 = int32(0)
	goto L26
L26:
	;
	v67 = F_ExecInsertIndexTuples(m, l0, l1, v64, l2, v63, v9+int32(15))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v69 != int32(1) {
		v76 = v67
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v72 = int32(0)
	F_CheckAndReportConflict(m, l0, l1, v72, v67, v72, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v76 = v67
	goto L22
L30:
	;
	F_list_free(m, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
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
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v386 int32
	_ = v386
	v15 = m.G0
	v17 = v15 - int32(1072)
	m.G0 = v17
	v19 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v20 = base.I64_rem_s(l1, v19)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	return v386
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
	v386 = v38
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
	v151 = base.I64_div_s(l1, int64(32))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v154 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v180 = F_OpenTransientFile(m, v17+int32(48), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v153
	v165 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(_a_F_SimpleLruReadPage_0), v17+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+36)) = uint32(v151)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v153
	v175 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(_a_F_SimpleLruReadPage_1), v17+int32(32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+40))
	if v275 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L28:
	;
	if v180 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	if v185 == int32(44) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3])) = v223
	v225 = int32(_a_F_SimpleLruReadPage_2)
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = int32(167772213)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229+v112<<(uint(int32(2))%32))))
	v234 = int32(_a_F_SimpleLruReadPage_3)
	v242 = F_pread(m, v180, v233, v234, base.I64_extend_i32_s(base.I32_wrap_i64(l1-v151<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v223
	if v242 != v234 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v200 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[5])))
	if v189&int32(1) != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v185
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[7])) = v192
	v273 = v192
	goto L27
L36:
	;
	goto L35
L37:
	;
	if v200 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(48)
	F_errmsg(m, int32(_a_F_SimpleLruReadPage_4), v17)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v112<<(uint(int32(2))%32))))
	base.MemoryFill(m, v217, int32(0), int32(_a_F_SimpleLruReadPage_3))
	v273 = int32(1)
	goto L27
L41:
	;
	F_errfinish(m, int32(_a_F_SimpleLruReadPage_5), int32(883), int32(_a_F_SimpleLruReadPage_6))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v255
	v257 = F_CloseTransientFile(m, v180)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v259 = F_CloseTransientFile(m, v180)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v273 = int32(0)
	goto L27
L47:
	;
	if v259 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v273 = int32(1)
	goto L27
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[7])) = int32(5)
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[6])) = v270
	v273 = int32(0)
	goto L27
L51:
	;
	v327 = F_LWLockAcquire(m, v147, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L61
	}
L52:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274)+36))
	v279 = v112 * v275
	v280 = int32(3)
	v282 = v278 + v279<<(uint(v280)%32)
	v286 = v275 << (uint(v280) % 32)
	if v282&v280|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v286)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v286 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	v315 = v286
	goto L55
L55:
	;
	if v315 == int32(0) {
		goto L51
	} else {
		goto L60
	}
L56:
	;
	v294 = int32(3)
	v295 = v279 << (uint(v294) % 32)
	v301 = v295 + v278 + int32(4)
	v307 = v275*(v112<<(uint(v294)%32)+int32(8)) + v278
	if base.Ui32(v307) < base.Ui32(v301) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v309 = v301
	goto L59
L58:
	;
	v309 = v307
	goto L59
L59:
	;
	v315 = (v295^int32(-1)-v278+v309)&int32(-4) + int32(4)
	goto L55
L60:
	;
	base.MemoryFill(m, v282, int32(0), v315)
	goto L51
L61:
	;
	v329 = int32(2)
	v330 = v112 << (uint(v329) % 32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v273 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v335 = v329
	goto L64
L63:
	;
	v335 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330+v331))) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	F_LWLockRelease(m, v337+v112<<(uint(int32(7))%32))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v273 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_SlruReportIOError(m, l0, l1, l3)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v352 = v347 + v112>>(uint(int32(4))%32)<<(uint(int32(2))%32)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354+v330)))
	if v353 != v356 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v359 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v112<<(uint(int32(2))%32)))) = v359
	goto L72
L71:
	;
	goto L72
L72:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[0])) = uint8(v369)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage[1])) = uint8(v369)
	v375 = v367 << (uint(int32(6)) % 32)
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v375)+uint32(_c_F_SimpleLruReadPage[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v375)+uint32(_c_F_SimpleLruReadPage[8]))) = v378 + int64(1)
	v386 = v112
	goto L3
}
func F_SimpleLruZeroAndWritePage(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v6 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v7 = base.I64_rem_s(l1, v6)
	v11 = v5 + base.I32_wrap_i64(v7)<<(uint(int32(7))%32)
	v13 = F_LWLockAcquire(m, v11, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_SimpleLruZeroPage(m, l0, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_SlruInternalWritePage(m, l0, v15, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_LWLockRelease(m, v11)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
