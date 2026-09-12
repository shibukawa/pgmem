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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
		v79 = v57
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_ExecARInsertTriggers(m, l1, l0, l2, v79, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	if v70&int32(1) == int32(0) {
		v79 = v68
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v75 = int32(0)
	F_CheckAndReportConflict(m, l0, l1, v75, v68, v75, l2)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v79 = v68
	goto L22
L27:
	;
	F_list_free(m, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
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
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int64
	_ = v368
	var v377 int32
	_ = v377
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
		v107 = v23
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(1072)
	return v377
L4:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v116+v107<<(uint(int32(3))%32)))) = l1
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v107<<(uint(int32(2))%32)))) = int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v107))) = uint8(v129)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v136 = F_LWLockAcquire(m, v131+v107<<(uint(int32(7))%32), v129)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v38 = v31
	v39 = v23
	goto L6
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v48+v39<<(uint(int32(3))%32))))
	if v52 != l1 {
		v107 = v39
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v107 = v95
	goto L4
L8:
	;
	switch v38 - int32(1) {
	case 0:
		goto L9
	default:
		goto L10
	case 2:
		goto L11
	}
L9:
	;
	F_SimpleLruWaitIO(m, l0, v39)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v61 = int32(2)
	v63 = v58 + v39>>(uint(int32(4))%32)<<(uint(v61)%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v66 = v39 << (uint(v61) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)))
	if v64 != v69 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v72 = v64 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v66))) = v72
	goto L15
L14:
	;
	goto L15
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, _consts[155])) = uint8(v80)
	v86 = v78 << (uint(int32(6)) % 32)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v86)+uint32(_consts[156])))
	*(*int64)(unsafe.Add(mBase, uint32(v86)+uint32(_consts[156]))) = v89 + int64(1)
	goto L16
L16:
	;
	v377 = v39
	goto L3
L17:
	;
	v95 = F_SlruSelectLRUPage(m, l0, l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v95<<(uint(int32(2))%32))))
	if v101 != 0 {
		v38 = v101
		v39 = v95
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	v141 = v22 + base.I32_wrap_i64(v20)<<(uint(int32(7))%32)
	F_LWLockRelease(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v145 = l0 + int32(16)
	v147 = base.I64_div_s(l1, int64(32))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v149 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v175 = F_OpenTransientFile(m, v17+int32(48), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v145
	v160 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(511422), v17+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+36)) = uint32(v147)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v145
	v170 = F_pg_snprintf(m, v17+int32(48), int32(1024), int32(511809), v17+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	goto L22
L28:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+40))
	if v269 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L29:
	;
	if v175 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v180 == int32(44) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v217 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v217
	v219 = int32(4126988)
	v220 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = int32(167772211)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223+v107<<(uint(int32(2))%32))))
	v228 = int32(8192)
	v236 = F_pread(m, v175, v227, v228, base.I64_extend_i32_s(base.I32_wrap_i64(l1-v147<<(uint(int64(5))%64))<<(uint(int32(13))%32)))
	mBase = m.M
	v238 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v217
	if v236 != v228 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v193 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L38
	}
L34:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
	if v184 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v180
	*(*int32)(unsafe.Add(mBase, _consts[160])) = v185
	v267 = v185
	goto L28
L37:
	;
	goto L36
L38:
	;
	if v193 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(48)
	F_errmsg(m, int32(163315), v17)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v107<<(uint(int32(2))%32))))
	v214 = F__emscripten_memset_bulkmem(m, v210, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L44
L42:
	;
	F_errfinish(m, int32(493962), int32(834), int32(410275))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v267 = int32(1)
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[160])) = int32(2)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v249
	v251 = F_CloseTransientFile(m, v175)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v253 = F_CloseTransientFile(m, v175)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v267 = int32(0)
	goto L28
L49:
	;
	if v253 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v267 = int32(1)
	goto L28
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[160])) = int32(5)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v264
	v267 = int32(0)
	goto L28
L53:
	;
	v317 = F_LWLockAcquire(m, v141, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L63
	}
L54:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+36))
	v273 = v269 * v107
	v274 = int32(3)
	v276 = v272 + v273<<(uint(v274)%32)
	v278 = v269 << (uint(v274) % 32)
	if base.Ui32(int32(1024)) < base.Ui32(v278) {
		v307 = v278
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v310 = F__emscripten_memset_bulkmem(m, v276, base.I32_extend8_s(int32(0)), v307)
	mBase = m.M
	goto L62
L56:
	;
	if v276&int32(3) != 0 {
		v307 = v278
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v276+v278) <= base.Ui32(v276) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v285 = int32(3)
	v286 = v273 << (uint(v285) % 32)
	v292 = v286 + v272 + int32(4)
	v298 = v269*(v107<<(uint(v285)%32)+int32(8)) + v272
	if base.Ui32(v298) < base.Ui32(v292) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v300 = v292
	goto L61
L60:
	;
	v300 = v298
	goto L61
L61:
	;
	v307 = (v286^int32(-1)-v272+v300)&int32(-4) + int32(4)
	goto L55
L62:
	;
	goto L53
L63:
	;
	v319 = int32(2)
	v320 = v107 << (uint(v319) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v267 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v325 = v319
	goto L66
L65:
	;
	v325 = int32(0)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320+v321))) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	F_LWLockRelease(m, v327+v107<<(uint(int32(7))%32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v267 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_SlruReportIOError(m, l0, l1, l3)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v342 = v337 + v107>>(uint(int32(4))%32)<<(uint(int32(2))%32)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v344+v320)))
	if v343 != v346 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v349 = v343 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v349
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v351+v107<<(uint(int32(2))%32)))) = v349
	goto L74
L73:
	;
	goto L74
L74:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v359)
	*(*uint8)(unsafe.Add(mBase, _consts[155])) = uint8(v359)
	v365 = v357 << (uint(int32(6)) % 32)
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v365)+uint32(_consts[161])))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+uint32(_consts[161]))) = v368 + int64(1)
	v377 = v107
	goto L3
}
func F_SimpleLruShmemSize(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = int32(0)
	v7 = int32(7)
	v9 = int32(-8)
	v10 = (l0<<(uint(int32(2))%32) + v7) & v9
	v12 = l0 << (uint(int32(3)) % 32)
	if v3 < l1 {
		v23 = l1 * v12
	} else {
		v23 = v3
	}
	v26 = base.I32_div_s(l0, int32(16))
	v28 = int32(7)
	return (v10+(v12+(l0+v7)&v9)+v23+(v26+l0)<<(uint(v28)%32)+v10<<(uint(int32(1))%32)+(v26<<(uint(int32(2))%32)+v28)&int32(-8)+int32(95))&int32(-32) + l0<<(uint(int32(13))%32)
}
