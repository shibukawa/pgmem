package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multixact_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int64
	_ = v281
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	switch int32(base.Ui32(v16) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L6
	case 1:
		goto L2
	case 2:
		goto L5
	case 3:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(112)
	return
L2:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
	v332 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_multixact_redo[1])))
	v333 = base.I64_rem_s(v330, v332)
	v337 = v328 + base.I32_wrap_i64(v333)<<(uint(int32(7))%32)
	v339 = F_LWLockAcquire(m, v337, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L11
	} else {
		goto L85
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L82
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v187 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L53
	}
L5:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v68 = *(*int64)(unsafe.Add(mBase, _c_F_multixact_redo[2]))
	if v68 != int64(-1) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_multixact_redo[2]))
	if v20 != v22 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_multixact_redo[2])) = int64(-1)
	goto L1
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[3]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v28 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_multixact_redo[4])))
	v29 = base.I64_rem_s(v20, v28)
	v33 = v26 + base.I32_wrap_i64(v29)<<(uint(int32(7))%32)
	v35 = F_LWLockAcquire(m, v33, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v47 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L16
	}
L11:
	;
	return
L12:
	;
	v37 = int32(_a_F_multixact_redo_0)
	v39 = F_SimpleLruZeroPage(m, v37, v20)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_SimpleLruWritePage(m, v37, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_LWLockRelease(m, v33)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if v47 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v20
	F_errmsg_internal(m, int32(_a_F_multixact_redo_1), v13+int32(16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_multixact_redo_2), int32(3512), int32(_a_F_multixact_redo_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	v73 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v95 = v66 + int32(12)
	F_RecordNewMultiXact(m, v91, v92, v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L29
	}
L23:
	;
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_multixact_redo[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v76
	F_errmsg_internal(m, int32(_a_F_multixact_redo_4), v13+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_multixact_redo[2])) = int64(-1)
	goto L22
L27:
	;
	F_errfinish(m, int32(_a_F_multixact_redo_2), int32(3551), int32(_a_F_multixact_redo_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[5]))
	v106 = F_LWLockAcquire(m, v102+int32(1664), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[6]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = v98 + int32(1)
	if v110-v112 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v112
	goto L33
L32:
	;
	goto L33
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v118 = v99 + v100
	if v117-v118 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v118
	goto L36
L35:
	;
	goto L36
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[5]))
	F_LWLockRelease(m, v125+int32(1664))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+36))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if int32(0) < v132 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = v131
	v138 = int32(0)
	goto L41
L39:
	;
	v167 = v131
	goto L40
L40:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v167)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L51
	}
L41:
	;
	v145 = int32(3)
	v147 = v95 + v138<<(uint(v145)%32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v148))&base.B2i32(base.Ui32(v145) <= base.Ui32(v135)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v167 = v162
	goto L40
L43:
	;
	if v160 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v160 = base.B2i32(base.Ui32(v135) < base.Ui32(v148))
	goto L43
L45:
	;
	goto L46
L46:
	;
	v160 = int32(base.Ui32(v135-v148) >> (uint(int32(31)) % 32))
	goto L43
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v162 = v161
	goto L49
L48:
	;
	v162 = v135
	goto L49
L49:
	;
	v164 = v138 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v164 < v165 {
		v135 = v162
		v138 = v164
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	goto L1
L52:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[5]))
	v242 = F_LWLockAcquire(m, v238+int32(_a_F_multixact_redo_5), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L59
	}
L53:
	;
	if v187 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v191 = int32(1636)
	v192 = base.I32_div_u_s(v181, v191)
	v193 = int32(5)
	v196 = base.I32_div_u_s(v180, v191)
	v198 = int32(base.Ui32(v196) >> (uint(v193) % 32))
	v234 = v198
	v235 = int32(base.Ui32(v192) >> (uint(v193) % 32))
	v236 = base.I64_extend_i32_u(v198)
	goto L52
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v181
	v202 = int32(1636)
	v203 = base.I32_div_u_s(v180, v202)
	v204 = int32(5)
	v205 = int32(base.Ui32(v203) >> (uint(v204) % 32))
	v206 = base.I64_extend_i32_u(v205)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+104)) = v206
	v209 = base.I32_div_u_s(v181, v202)
	v211 = int32(base.Ui32(v209) >> (uint(v204) % 32))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = base.I64_extend_i32_u(v211)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v182
	v215 = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = base.I64_extend_i32_u(int32(base.Ui32(v182) >> (uint(v215) % 32)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v183
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = base.I64_extend_i32_u(int32(base.Ui32(v183) >> (uint(v215) % 32)))
	F_errmsg_internal(m, int32(_a_F_multixact_redo_6), v13-int32(-64))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_multixact_redo_2), int32(3592), int32(_a_F_multixact_redo_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v234 = v205
	v235 = v211
	v236 = v206
	goto L52
L59:
	;
	F_SetMultiXactIdLimit(m, v182, v184, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	if v234 != v235 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v257 = base.I64_extend_i32_u(v235)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v295 = int32(1)
	if v182 == v295 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v261 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	if v261 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v257
	F_errmsg_internal(m, int32(_a_F_multixact_redo_7), v13+int32(48))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L11
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_SlruDeleteSegment(m, v257)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L72
	}
L70:
	;
	F_errfinish(m, int32(_a_F_multixact_redo_2), int32(3132), int32(_a_F_multixact_redo_8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	if v257 != int64(82040) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v281 = v257 + int64(1)
	goto L75
L74:
	;
	v281 = int64(0)
	goto L75
L75:
	;
	if v281 != v236 {
		v257 = v281
		goto L64
	} else {
		goto L76
	}
L76:
	;
	goto L65
L77:
	;
	v301 = int32(_a_F_multixact_redo_9)
	goto L79
L78:
	;
	v301 = int32(base.Ui32(v182-v295) >> (uint(int32(11)) % 32))
	goto L79
L79:
	;
	F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_0), base.I64_extend_i32_u(v301))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[5]))
	F_LWLockRelease(m, v306+int32(_a_F_multixact_redo_5))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16 & int32(240)
	F_errmsg_internal(m, int32(_a_F_multixact_redo_10), v13)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_multixact_redo_2), int32(3609), int32(_a_F_multixact_redo_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v341 = int32(_a_F_multixact_redo_11)
	v343 = F_SimpleLruZeroPage(m, v341, v330)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	F_SimpleLruWritePage(m, v341, v343)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_LWLockRelease(m, v337)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	goto L1
}
