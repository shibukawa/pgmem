package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multixact_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
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
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
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
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v253 int64
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int64
	_ = v277
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	switch int32(base.Ui32(v15) >> (uint(int32(4)) % 32)) {
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
	m.G0 = v12 + int32(112)
	return
L2:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+28))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v324)))
	v327 = int64(*(*uint16)(unsafe.Add(mBase, _consts[148])))
	v328 = base.I64_rem_s(v325, v327)
	v332 = v323 + base.I32_wrap_i64(v328)<<(uint(int32(7))%32)
	v334 = F_LWLockAcquire(m, v332, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L11
	} else {
		goto L85
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L11
	} else {
		goto L82
	}
L4:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v184 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L53
	}
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v67 = *(*int64)(unsafe.Add(mBase, _consts[149]))
	if v67 != int64(-1) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v21 = *(*int64)(unsafe.Add(mBase, _consts[149]))
	if v19 != v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, _consts[149])) = int64(-1)
	goto L1
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v27 = int64(*(*uint16)(unsafe.Add(mBase, _consts[151])))
	v28 = base.I64_rem_s(v19, v27)
	v32 = v25 + base.I32_wrap_i64(v28)<<(uint(int32(7))%32)
	v34 = F_LWLockAcquire(m, v32, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v46 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L16
	}
L11:
	;
	return
L12:
	;
	v36 = int32(4456772)
	v38 = F_SimpleLruZeroPage(m, v36, v19)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_SimpleLruWritePage(m, v36, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if v46 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v19
	F_errmsg_internal(m, int32(278230), v12+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(518132), int32(3512), int32(255539))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	v72 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v94 = v65 + int32(12)
	F_RecordNewMultiXact(m, v90, v91, v92, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L29
	}
L23:
	;
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = *(*int64)(unsafe.Add(mBase, _consts[149]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v75
	F_errmsg_internal(m, int32(233719), v12+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, _consts[149])) = int64(-1)
	goto L22
L27:
	;
	F_errfinish(m, int32(518132), int32(3551), int32(255539))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v101 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v105 = F_LWLockAcquire(m, v101+int32(1664), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v107 = v98 + v99
	v109 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = v97 + int32(1)
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
	if v117-v107 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v107
	goto L36
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v124+int32(1664))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+36))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if int32(0) < v131 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v134 = v130
	v135 = int32(0)
	goto L41
L39:
	;
	v165 = v130
	goto L40
L40:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v165)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L51
	}
L41:
	;
	v143 = int32(3)
	v145 = v94 + v135<<(uint(v143)%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v146))&base.B2i32(base.Ui32(v143) <= base.Ui32(v134)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v165 = v160
	goto L40
L43:
	;
	if v158 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v158 = base.B2i32(base.Ui32(v134) < base.Ui32(v146))
	goto L43
L45:
	;
	goto L46
L46:
	;
	v158 = int32(base.Ui32(v134-v146) >> (uint(int32(31)) % 32))
	goto L43
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v160 = v159
	goto L49
L48:
	;
	v160 = v134
	goto L49
L49:
	;
	v162 = v135 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v162 < v163 {
		v134 = v160
		v135 = v162
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
	v235 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v239 = F_LWLockAcquire(m, v235+int32(5248), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L59
	}
L53:
	;
	if v184 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v188 = int32(1636)
	v189 = base.I32_div_u_s(v178, v188)
	v190 = int32(5)
	v193 = base.I32_div_u_s(v177, v188)
	v195 = int32(base.Ui32(v193) >> (uint(v190) % 32))
	v231 = int32(base.Ui32(v189) >> (uint(v190) % 32))
	v232 = v195
	v233 = base.I64_extend_i32_u(v195)
	goto L52
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v178
	v199 = int32(1636)
	v200 = base.I32_div_u_s(v177, v199)
	v201 = int32(5)
	v202 = int32(base.Ui32(v200) >> (uint(v201) % 32))
	v203 = base.I64_extend_i32_u(v202)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v203
	v206 = base.I32_div_u_s(v178, v199)
	v208 = int32(base.Ui32(v206) >> (uint(v201) % 32))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = base.I64_extend_i32_u(v208)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v179
	v212 = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = base.I64_extend_i32_u(int32(base.Ui32(v179) >> (uint(v212) % 32)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = base.I64_extend_i32_u(int32(base.Ui32(v180) >> (uint(v212) % 32)))
	F_errmsg_internal(m, int32(702644), v12-int32(-64))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(518132), int32(3592), int32(255539))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v231 = v208
	v232 = v202
	v233 = v203
	goto L52
L59:
	;
	F_SetMultiXactIdLimit(m, v179, v181, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	if v231 != v232 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v253 = base.I64_extend_i32_u(v231)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v290 = int32(1)
	if v179 == v290 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v257 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L11
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	if v257 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v253
	F_errmsg_internal(m, int32(27806), v12+int32(48))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_SlruDeleteSegment(m, v253)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L72
	}
L70:
	;
	F_errfinish(m, int32(518132), int32(3132), int32(279173))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	if v253 != int64(82040) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v277 = v253 + int64(1)
	goto L75
L74:
	;
	v277 = int64(0)
	goto L75
L75:
	;
	if v277 != v233 {
		v253 = v277
		goto L64
	} else {
		goto L76
	}
L76:
	;
	goto L65
L77:
	;
	v296 = int32(2097151)
	goto L79
L78:
	;
	v296 = int32(base.Ui32(v179-v290) >> (uint(int32(11)) % 32))
	goto L79
L79:
	;
	F_SimpleLruTruncate(m, int32(4456772), base.I64_extend_i32_u(v296))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v301+int32(5248))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15 & int32(240)
	F_errmsg_internal(m, int32(57378), v12)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(518132), int32(3609), int32(255539))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
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
	v336 = int32(4456852)
	v338 = F_SimpleLruZeroPage(m, v336, v325)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	F_SimpleLruWritePage(m, v336, v338)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_LWLockRelease(m, v332)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	goto L1
}
