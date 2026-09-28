package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeLogRepWorker(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_0), int32(_a_F_InitializeLogRepWorker_1), int32(5), int32(10))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_BackgroundWorkerInitializeConnectionByOid(m, v16, v17, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_2), int32(_a_F_InitializeLogRepWorker_3), int32(5), int32(10))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[1]))
	v34 = F_AllocSetContextCreateInternal(m, v29, int32(_a_F_InitializeLogRepWorker_4), int32(0), int32(_a_F_InitializeLogRepWorker_5), int32(_a_F_InitializeLogRepWorker_6))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[2])) = v34
	F_StartTransactionCommand(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	F_LockSharedObject(m, int32(_a_F_InitializeLogRepWorker_7), v42, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	v51 = F_GetSubscription(m, v49, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3])) = v51
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v322 == int32(3) {
		goto L91
	} else {
		goto L92
	}
L10:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L89
	}
L11:
	;
	F_errfinish(m, int32(_a_F_InitializeLogRepWorker_9), v307, int32(_a_F_InitializeLogRepWorker_11))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L88
	}
L12:
	;
	v292 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L85
	}
L13:
	;
	v276 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L82
	}
L14:
	;
	v254 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L78
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[2]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v61 != v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v227 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L67
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+33)))
	if v92 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	if v61 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	if v57 != 0 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	if v66 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v65 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v65
	goto L24
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v65
	goto L24
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = v71
	goto L22
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v57
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+28)) = v78
	if v78 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
	goto L21
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v55
	goto L34
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v55
	goto L18
L35:
	;
	v97 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v124 = v91
	goto L37
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[2]))
	v128 = F_SubscriptionConninfo(m, v124)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L45
	}
L38:
	;
	if v97 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v101
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_8), v6+int32(80))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+16)))
	if base.B2i32(v116 != int32(1))|base.B2i32(v115 != int32(4)) != 0 {
		v321 = v114
		v322 = v115
		goto L9
	} else {
		goto L44
	}
L42:
	;
	F_errfinish(m, int32(_a_F_InitializeLogRepWorker_9), int32(_a_F_InitializeLogRepWorker_10), int32(_a_F_InitializeLogRepWorker_11))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v124 = v123
	goto L37
L45:
	;
	v130 = F_MemoryContextStrdup(m, v127, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4])) = uint8(v133)
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[5])) = v130
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v141 != int32(3) {
		v182 = v138
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+56))
	F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_12), v185, int32(4), int32(10))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L59
	}
L48:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+41)))
	if v144 != int32(1) {
		v182 = v138
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+48)))
	if v147 != int32(1) {
		v182 = v138
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+72))
	if v150 != 0 {
		v182 = v138
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v153 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v153 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = int32(_a_F_InitializeLogRepWorker_19)
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_20), v6-int32(-64))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+16)))
	if base.B2i32(v174 != int32(1))|base.B2i32(v173 != int32(4)) != 0 {
		v321 = v172
		v322 = v173
		goto L9
	} else {
		goto L58
	}
L56:
	;
	F_errfinish(m, int32(_a_F_InitializeLogRepWorker_9), int32(_a_F_InitializeLogRepWorker_21), int32(_a_F_InitializeLogRepWorker_11))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v182 = v181
	goto L47
L59:
	;
	F_set_wal_receiver_timeout(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_CacheRegisterSyscacheCallback(m, int32(67), int32(1096), int64(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_CacheRegisterSyscacheCallback(m, int32(32), int32(1096), int64(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_CacheRegisterSyscacheCallback(m, int32(83), int32(1096), int64(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_CacheRegisterSyscacheCallback(m, int32(30), int32(1096), int64(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_CacheRegisterSyscacheCallback(m, int32(11), int32(1096), int64(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+16)))
	if v219 != int32(1) {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	switch v222 - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L12
	}
L67:
	;
	if v227 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v231
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_22), v6)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v243 == int32(3) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	F_errfinish(m, int32(_a_F_InitializeLogRepWorker_9), int32(_a_F_InitializeLogRepWorker_23), int32(_a_F_InitializeLogRepWorker_11))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	if v254 == int32(0) {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+36))
	v264 = F_get_rel_name(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v260
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_15), v6+int32(32))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v307 = int32(_a_F_InitializeLogRepWorker_16)
	goto L11
L82:
	;
	if v276 == int32(0) {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v282
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_17), v6+int32(48))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v307 = int32(_a_F_InitializeLogRepWorker_18)
	goto L11
L85:
	;
	if v292 == int32(0) {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v298
	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_13), v6+int32(16))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v307 = int32(_a_F_InitializeLogRepWorker_14)
	goto L11
L88:
	;
	goto L10
L89:
	;
	F_before_shmem_exit(m, int32(1097), int64(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	m.G0 = v6 + int32(96)
	return
L91:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LogRecoveryConflict(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v22 = l2 - l1
	if v22 <= int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v39 = int32(1000)
	v40 = base.I32_div_s(v38, v39)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v38 - v40*v39
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v48 = int32(0)
	if l3 == v48 {
		v115 = v48
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v34 = int32(0)
	v35 = int32(0)
	goto L4
L3:
	;
	v26 = int64(1000000)
	v27 = base.I64_div_u_s(v22, v26)
	v34 = base.I32_wrap_i64(v27)
	v35 = base.I32_wrap_i64(v22 - v27*v26)
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(76)))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(72)))) = v35
	goto L1
L5:
	;
	v120 = v40 + v45*v39
	v123 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L23
	} else {
		goto L27
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v51 == int32(0) {
		v115 = v48
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = l3
	v60 = v48
	goto L8
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = int32(0)
	if v65 < v66 {
		v84 = v66
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v115 = v104
	goto L5
L10:
	;
	if v84 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_LogRecoveryConflict[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if base.Ui32(v73) <= base.Ui32(v65) {
		v84 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v78 = v75 + v65*int32(768)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v80 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v81 = v78
	goto L16
L15:
	;
	v81 = int32(0)
	goto L16
L16:
	;
	v84 = v81
	goto L11
L17:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v104 = v60
	goto L19
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v106 != 0 {
		v57 = v57 + int32(8)
		v60 = v104
		goto L8
	} else {
		goto L26
	}
L20:
	;
	F_initStringInfo(m, v14+int32(56))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v93 = int32(_a_F_LogRecoveryConflict_0)
	goto L22
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v94
	F_appendStringInfo(m, v14+int32(56), v93, v14+int32(48))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L25
	}
L23:
	;
	return
L24:
	;
	v93 = int32(_a_F_LogRecoveryConflict_1)
	goto L22
L25:
	;
	v104 = v60 + int32(1)
	goto L19
L26:
	;
	goto L9
L27:
	;
	if l4 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if int32(0) < v115 {
		goto L46
	} else {
		goto L47
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LogRecoveryConflict_2), v171, int32(_a_F_LogRecoveryConflict_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L23
	} else {
		goto L45
	}
L30:
	;
	if v123 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v123 == int32(0) {
		goto L28
	} else {
		goto L40
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if base.Ui32(l0) <= base.Ui32(int32(7)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_LogRecoveryConflict[1])))
	v134 = v132
	goto L36
L35:
	;
	v134 = int32(_a_F_LogRecoveryConflict_4)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v120
	F_errmsg(m, int32(_a_F_LogRecoveryConflict_5), v14+int32(16))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v143 = int32(335)
	if v115 <= int32(0) {
		v171 = v143
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v146
	F_errdetail_log_plural(m, int32(_a_F_LogRecoveryConflict_6), int32(_a_F_LogRecoveryConflict_7), v115, v14)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v171 = v143
	goto L29
L40:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if base.Ui32(l0) <= base.Ui32(int32(7)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_LogRecoveryConflict[1])))
	v161 = v159
	goto L43
L42:
	;
	v161 = int32(_a_F_LogRecoveryConflict_4)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v120
	F_errmsg(m, int32(_a_F_LogRecoveryConflict_8), v14+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v171 = int32(341)
	goto L29
L45:
	;
	goto L28
L46:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L23
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	m.G0 = v14 + int32(80)
	return
L49:
	;
	goto L48
}
func F_LogStandbySnapshot(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v208 int64
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int64
	_ = v454
	var v458 int32
	_ = v458
	var v462 int64
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v472 int64
	_ = v472
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = F_IsLogicalDecodingEnabled(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v31 = F_LWLockAcquire(m, v27+int32(_a_F_LogStandbySnapshot_0), int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v38 = F_LWLockAcquire(m, v34+int32(_a_F_LogStandbySnapshot_1), int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v45 = F_LWLockAcquire(m, v41+int32(_a_F_LogStandbySnapshot_2), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v52 = F_LWLockAcquire(m, v48+int32(_a_F_LogStandbySnapshot_3), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v59 = F_LWLockAcquire(m, v55+int32(_a_F_LogStandbySnapshot_4), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v66 = F_LWLockAcquire(m, v62+int32(_a_F_LogStandbySnapshot_5), int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v73 = F_LWLockAcquire(m, v69+int32(_a_F_LogStandbySnapshot_6), int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v80 = F_LWLockAcquire(m, v76+int32(_a_F_LogStandbySnapshot_7), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v87 = F_LWLockAcquire(m, v83+int32(_a_F_LogStandbySnapshot_8), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v94 = F_LWLockAcquire(m, v90+int32(_a_F_LogStandbySnapshot_9), int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v101 = F_LWLockAcquire(m, v97+int32(_a_F_LogStandbySnapshot_10), int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v108 = F_LWLockAcquire(m, v104+int32(_a_F_LogStandbySnapshot_11), int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v115 = F_LWLockAcquire(m, v111+int32(_a_F_LogStandbySnapshot_12), int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v122 = F_LWLockAcquire(m, v118+int32(_a_F_LogStandbySnapshot_13), int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v129 = F_LWLockAcquire(m, v125+int32(_a_F_LogStandbySnapshot_14), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v136 = F_LWLockAcquire(m, v132+int32(_a_F_LogStandbySnapshot_15), int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v141)+808))
	if v143 != int64(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v212 = F_palloc(m, base.I32_wrap_i64(v208)*int32(12))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v141)+752))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v141)+728))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v141)+704))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v141)+680))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v141)+656))
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v141)+632))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v141)+608))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v141)+584))
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v141)+560))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v141)+536))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v141)+512))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v141)+488))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v141)+464))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v141)+440))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v141)+416))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v141)+392))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v141)+368))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v141)+344))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v141)+320))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v141)+296))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v141)+272))
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v141)+248))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v141)+224))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v141)+200))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v141)+176))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v141)+152))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v141)+128))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v141)+104))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v141)+80))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v141)+56))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v141)+32))
	v208 = v146 + (v147 + (v148 + (v149 + (v150 + (v151 + (v152 + (v153 + (v154 + (v155 + (v156 + (v157 + (v158 + (v159 + (v160 + (v161 + (v162 + (v163 + (v164 + (v165 + (v166 + (v167 + (v168 + (v169 + (v170 + (v171 + (v172 + (v173 + (v174 + (v175 + (v176 + v142))))))))))))))))))))))))))))))
	goto L22
L21:
	;
	v208 = v142
	goto L22
L22:
	;
	goto L19
L23:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[1]))
	F_hash_seq_init(m, v24+int32(12), v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v224 = int32(0)
	goto L25
L25:
	;
	v233 = F_hash_seq_search(m, v24+int32(12))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v257+int32(_a_F_LogStandbySnapshot_15))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L34
	}
L27:
	;
	if v233 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+13)))
	if v235&int32(1) == int32(0) {
		goto L25
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L26
L31:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+14)))
	if v241 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	if v243 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v248 = v212 + v224*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v243
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v252
	v224 = v224 + int32(1)
	goto L25
L34:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v263+int32(_a_F_LogStandbySnapshot_14))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v269+int32(_a_F_LogStandbySnapshot_13))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v275+int32(_a_F_LogStandbySnapshot_12))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v281+int32(_a_F_LogStandbySnapshot_11))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v287+int32(_a_F_LogStandbySnapshot_10))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v293+int32(_a_F_LogStandbySnapshot_9))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v299+int32(_a_F_LogStandbySnapshot_8))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v305+int32(_a_F_LogStandbySnapshot_7))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v311+int32(_a_F_LogStandbySnapshot_6))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v317+int32(_a_F_LogStandbySnapshot_5))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v323+int32(_a_F_LogStandbySnapshot_4))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v329+int32(_a_F_LogStandbySnapshot_3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v335+int32(_a_F_LogStandbySnapshot_2))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v341+int32(_a_F_LogStandbySnapshot_1))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v347+int32(_a_F_LogStandbySnapshot_0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(68)))) = v224
	m.G0 = v24 + int32(32)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if int32(0) < v356 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v356
	F_XLogBeginInsert(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_pfree(m, v212)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L58
	}
L53:
	;
	F_XLogRegisterData(m, v14+int32(72), int32(4))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_XLogRegisterData(m, v212, v356*int32(12))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v372 = int32(_a_F_LogStandbySnapshot_16)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])))
	v375 = v374 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])) = uint8(v375)
	goto L56
L56:
	;
	v379 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v383 = F_GetRunningTransactionData(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v16 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v388+int32(512))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)) = uint8(base.B2i32(v397 != int32(0)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = v405
	F_XLogBeginInsert(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v410 = int32(_a_F_LogStandbySnapshot_16)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])))
	v413 = v412 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])) = uint8(v413)
	goto L65
L65:
	;
	F_XLogRegisterData(m, v14+int32(72), int32(24))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if int32(0) < v420 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v383)+28))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	F_XLogRegisterData(m, v423, (v424+v420)<<(uint(int32(2))%32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v432 = F_XLogInsert(m, int32(8), int32(16))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	v437 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v434 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	F_XLogSetAsyncXactLSN(m, v432)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L83
	}
L74:
	;
	F_errfinish(m, int32(_a_F_LogStandbySnapshot_17), v484, int32(_a_F_LogStandbySnapshot_18))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L82
	}
L75:
	;
	if v437 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v437 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L78:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v444
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v432)
	v454 = int64(base.Ui64(v432) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v454)
	F_errmsg_internal(m, int32(_a_F_LogStandbySnapshot_19), v14)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v484 = int32(1387)
	goto L74
L80:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v462
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+44)) = uint32(v432)
	v472 = int64(base.Ui64(v432) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+40)) = uint32(v472)
	F_errmsg_internal(m, int32(_a_F_LogStandbySnapshot_20), v14+int32(32))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v484 = int32(1395)
	goto L74
L82:
	;
	goto L73
L83:
	;
	if v16 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v495+int32(512))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v501+int32(384))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	m.G0 = v14 + int32(96)
	return v432
}
func F_check_log_stats(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != v4 {
		v33 = v4
		return v33
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[0])))
		if v9 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
			v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
				v33 = int32(0)
				return v33
			}
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[4])))
			if v11&int32(1) != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
				v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
					v33 = int32(0)
					return v33
				}
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[5])))
				if v15&int32(1) == int32(0) {
					v33 = v4
					return v33
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
					v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
						v33 = int32(0)
						return v33
					}
				}
			}
		}
	}
}
func F_log_min_messages_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = int32(58)
	v6 = F___strchrnul(m, v4, v5)
	mBase = m.M
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v8 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v12 = v6
	goto L4
L3:
	;
	v12 = int32(0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return int32(-1)
L6:
	;
	goto L7
L7:
	;
	v17 = int32(58)
	v18 = F___strchrnul(m, v3, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 == v17 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v24 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v24 = v18
	goto L11
L10:
	;
	v24 = int32(0)
	goto L11
L11:
	;
	goto L8
L12:
	;
	return int32(1)
L13:
	;
	goto L14
L14:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if base.B2i32(v31 == int32(0))|base.B2i32(v31 != v34) != 0 {
		v52 = v31
		v53 = v34
		goto L16
	} else {
		goto L17
	}
L15:
	;
	return v52 - v53
L16:
	;
	goto L15
L17:
	;
	v37 = v4
	v38 = v3
	goto L18
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v42
		v53 = v41
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v52 = v42
	v53 = v41
	goto L16
L20:
	;
	v45 = int32(1)
	if v42 == v41 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
}
func F_log_split_page(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+118)))
	if v5 != int32(112) {
		v24 = F_XLogGetFakeLSN(m, l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = v24
			if l1 < int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[0]))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(2))%32))))
				v44 = v36
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[1]))
				v44 = v38 + l1<<(uint(int32(13))%32) + int32(-8192)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v44))) = base.I64_rotl(v26, int64(32))
			return
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[2]))
		if v9 <= int32(0) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v12 != 0 {
				v24 = F_XLogGetFakeLSN(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = v24
					if l1 < int32(0) {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[0]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(2))%32))))
						v44 = v36
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[1]))
						v44 = v38 + l1<<(uint(int32(13))%32) + int32(-8192)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v44))) = base.I64_rotl(v26, int64(32))
					return
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v13 != 0 {
					v24 = F_XLogGetFakeLSN(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = v24
						if l1 < int32(0) {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[0]))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(2))%32))))
							v44 = v36
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[1]))
							v44 = v38 + l1<<(uint(int32(13))%32) + int32(-8192)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v44))) = base.I64_rotl(v26, int64(32))
						return
					}
				} else {
					F_XLogBeginInsert(m)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_XLogRegisterBuffer(m, int32(0), l1, int32(9))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v22 = F_XLogInsert(m, int32(12), int32(80))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								v26 = v22
								if l1 < int32(0) {
									v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[0]))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(2))%32))))
									v44 = v36
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[1]))
									v44 = v38 + l1<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int64)(unsafe.Add(mBase, uint32(v44))) = base.I64_rotl(v26, int64(32))
								return
							}
						}
					}
				}
			}
		} else {
			F_XLogBeginInsert(m)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_XLogRegisterBuffer(m, int32(0), l1, int32(9))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v22 = F_XLogInsert(m, int32(12), int32(80))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v26 = v22
						if l1 < int32(0) {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[0]))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(2))%32))))
							v44 = v36
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_log_split_page[1]))
							v44 = v38 + l1<<(uint(int32(13))%32) + int32(-8192)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v44))) = base.I64_rotl(v26, int64(32))
						return
					}
				}
			}
		}
	}
}
func F_log_status_format(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int64
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int64
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v388 int64
	_ = v388
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
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v501 int64
	_ = v501
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v836 int32
	_ = v836
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	v10 = m.G0
	v12 = v10 - int32(688)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[1]))
	if v16 == v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2])) = v30
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2]))
	v30 = v21 + int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[1])) = v16
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[3])) = uint8(v27)
	v30 = int32(1)
	goto L1
L5:
	;
	m.G0 = v12 + int32(688)
	return
L6:
	;
	v35 = l1
	goto L7
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v43 != int32(37) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v43 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v54 = v35 + int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v55 != int32(37) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	F_appendStringInfoChar(m, l0, base.I32_extend8_s(v43))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v35 = v35 + int32(1)
	goto L7
L15:
	;
	v35 = v122 + int32(1)
	goto L7
L16:
	;
	if v55 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_appendStringInfoChar(m, l0, int32(37))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L13
	} else {
		goto L271
	}
L19:
	;
	v60 = int32(0)
	v61 = base.I32_extend8_s(v55)
	if int32(57) < v61 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	switch v127&int32(255) - int32(76) {
	case 0:
		goto L43
	default:
		goto L15
	case 4:
		goto L50
	case 5:
		goto L37
	case 21:
		goto L56
	case 22:
		goto L55
	case 23:
		goto L52
	case 24:
		goto L53
	case 25:
		goto L38
	case 28:
		goto L41
	case 29:
		goto L44
	case 32:
		goto L49
	case 33:
		goto L48
	case 34:
		goto L46
	case 36:
		goto L51
	case 37:
		goto L36
	case 38:
		goto L42
	case 39:
		goto L45
	case 40:
		goto L47
	case 41:
		goto L54
	case 42:
		goto L40
	case 44:
		goto L39
	}
L21:
	;
	v122 = v54
	v126 = v60
	v127 = v61
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v61 != int32(45) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if base.Ui32((v73-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v73 = v61
	v74 = int32(1)
	v75 = v54
	goto L24
L26:
	;
	goto L27
L27:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
	if v67 == int32(0) {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v73 = v67
	v74 = int32(-1)
	v75 = v35 + int32(2)
	goto L24
L29:
	;
	v83 = v75
	v87 = v60
	v88 = v73
	goto L32
L30:
	;
	v109 = v75
	v113 = v60
	v114 = v73
	goto L31
L31:
	;
	if v114&int32(255) == int32(0) {
		goto L5
	} else {
		goto L35
	}
L32:
	;
	v91 = int32(10)
	v93 = int32(48)
	v95 = int32(255)
	v97 = v87*v91 + (v88-v93)&v95
	v99 = v83 + int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if base.Ui32((v100-v93)&v95) < base.Ui32(v91) {
		v83 = v99
		v87 = v97
		v88 = v100
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v99
	v113 = v97 * v74
	v114 = v100
	goto L31
L34:
	;
	goto L33
L35:
	;
	v122 = v109
	v126 = v113
	v127 = v114
	goto L20
L36:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v846 != 0 {
		goto L15
	} else {
		goto L270
	}
L37:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[5]))
	if v824 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L38:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v767 = int32(63)
	v769 = int32(48)
	v770 = v766&v767 + v769
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[6])) = uint8(v770)
	v778 = int32(base.Ui32(v766)>>(uint(int32(24))%32))&v767 + v769
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[7])) = uint8(v778)
	v786 = int32(base.Ui32(v766)>>(uint(int32(18))%32))&v767 + v769
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[8])) = uint8(v786)
	v794 = int32(base.Ui32(v766)>>(uint(int32(12))%32))&v767 + v769
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[9])) = uint8(v794)
	v802 = int32(base.Ui32(v766)>>(uint(int32(6))%32))&v767 + v769
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[10])) = uint8(v802)
	v805 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[11])) = uint8(v805)
	if v126 != 0 {
		goto L256
	} else {
		goto L257
	}
L39:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[12]))
	if v126 != 0 {
		goto L251
	} else {
		goto L252
	}
L40:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[13]))
	if v699 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L41:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v668 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L42:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v599 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L43:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v565 != 0 {
		goto L199
	} else {
		goto L200
	}
L44:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v534 != 0 {
		goto L188
	} else {
		goto L189
	}
L45:
	;
	v501 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+544)) = v501
	v504 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[3])))
	if v504 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L46:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])))
	if v461 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L47:
	;
	v431 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+680)) = v431
	v434 = v12 + int32(544)
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v441 = F_pg_localtime(m, v12+int32(680), v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L13
	} else {
		goto L162
	}
L48:
	;
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[17])) = uint8(v376)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])))
	if v379 == v376 {
		goto L151
	} else {
		goto L152
	}
L49:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2]))
	if v126 != 0 {
		goto L146
	} else {
		goto L147
	}
L50:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[13]))
	if v314 != 0 {
		goto L130
	} else {
		goto L131
	}
L51:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	if v126 != 0 {
		goto L125
	} else {
		goto L126
	}
L52:
	;
	if v126 != 0 {
		goto L119
	} else {
		goto L120
	}
L53:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v230 != 0 {
		goto L104
	} else {
		goto L105
	}
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v201 != 0 {
		goto L89
	} else {
		goto L90
	}
L55:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[18]))
	if v164 == v166 {
		v186 = int32(_a_F_log_status_format_0)
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v135 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[19]))
	if v137 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L70
	}
L60:
	;
	if v126 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v138 != 0 {
		v140 = v137
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v140 = int32(_a_F_log_status_format_1)
	goto L60
L64:
	;
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_appendStringInfoString(m, l0, v140)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L69
	}
L68:
	;
	v35 = v122 + int32(1)
	goto L7
L69:
	;
	v35 = v122 + int32(1)
	goto L7
L70:
	;
	v155 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v155-v155)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v35 = v122 + int32(1)
	goto L7
L72:
	;
	if v126 != 0 {
		goto L84
	} else {
		goto L85
	}
L73:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[20]))
	if v169 == int32(5) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[21]))
	if v173 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(v169) <= base.Ui32(int32(17)) {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v177 = v173 + int32(96)
	goto L79
L78:
	;
	v177 = int32(_a_F_log_status_format_3)
	goto L79
L79:
	;
	v186 = v177
	goto L72
L80:
	;
	v186 = v184
	goto L72
L81:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v169<<(uint(int32(2))%32))+uint32(_c_F_log_status_format[22])))
	v184 = v182
	goto L83
L82:
	;
	v184 = int32(_a_F_log_status_format_4)
	goto L83
L83:
	;
	goto L80
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	F_appendStringInfoString(m, l0, v186)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L88
	}
L87:
	;
	v35 = v122 + int32(1)
	goto L7
L88:
	;
	v35 = v122 + int32(1)
	goto L7
L89:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+364))
	if v202 != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L102
	}
L92:
	;
	if v126 != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v203 != 0 {
		v205 = v202
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v205 = int32(_a_F_log_status_format_1)
	goto L92
L96:
	;
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_appendStringInfoString(m, l0, v205)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L13
	} else {
		goto L101
	}
L100:
	;
	v35 = v122 + int32(1)
	goto L7
L101:
	;
	v35 = v122 + int32(1)
	goto L7
L102:
	;
	v222 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v222-v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	v35 = v122 + int32(1)
	goto L7
L104:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+360))
	if v231 != 0 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L106
L106:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L117
	}
L107:
	;
	if v126 != 0 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v232 != 0 {
		v234 = v231
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v234 = int32(_a_F_log_status_format_1)
	goto L107
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_appendStringInfoString(m, l0, v234)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L13
	} else {
		goto L116
	}
L115:
	;
	v35 = v122 + int32(1)
	goto L7
L116:
	;
	v35 = v122 + int32(1)
	goto L7
L117:
	;
	v251 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v251-v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	v35 = v122 + int32(1)
	goto L7
L119:
	;
	v259 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v259
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v262
	v265 = v12 + int32(544)
	v270 = F_pg_snprintf(m, v265, int32(127), int32(_a_F_log_status_format_5), v12+int32(96))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L13
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v282 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v282
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v285
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_5), v12-int32(-64))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v265
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(80))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	v35 = v122 + int32(1)
	goto L7
L124:
	;
	v35 = v122 + int32(1)
	goto L7
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_6), v12+int32(128))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v295
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_7), v12+int32(112))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L129
	}
L128:
	;
	v35 = v122 + int32(1)
	goto L7
L129:
	;
	v35 = v122 + int32(1)
	goto L7
L130:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+364))
	if v315 != 0 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	goto L132
L132:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L144
	}
L133:
	;
	if v126 != 0 {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	if v316 != v318 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v322 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v322-v322)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L13
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v35 = v122 + int32(1)
	goto L7
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_6), v12+int32(160))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v316
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_7), v12+int32(144))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L143
	}
L142:
	;
	v35 = v122 + int32(1)
	goto L7
L143:
	;
	v35 = v122 + int32(1)
	goto L7
L144:
	;
	v349 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v349-v349)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	v35 = v122 + int32(1)
	goto L7
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_8), v12+int32(192))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L13
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v357
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_9), v12+int32(176))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L13
	} else {
		goto L150
	}
L149:
	;
	v35 = v122 + int32(1)
	goto L7
L150:
	;
	v35 = v122 + int32(1)
	goto L7
L151:
	;
	F_gettimeofday(m, int32(_a_F_log_status_format_10))
	mBase = m.M
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])) = uint8(v385)
	goto L153
L152:
	;
	goto L153
L153:
	;
	v388 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[23]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+680)) = v388
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v397 = F_pg_localtime(m, v12+int32(680), v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	v399 = F_pg_strftime(m, int32(_a_F_log_status_format_11), int32(128), int32(_a_F_log_status_format_12), v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[24]))
	v404 = base.I32_div_s(v402, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v404
	v411 = F_pg_sprintf(m, v12+int32(544), int32(_a_F_log_status_format_13), v12+int32(224))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v12)+544))
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[25])) = v414
	if v126 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+212)) = int32(_a_F_log_status_format_11)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(208))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L13
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_11))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L13
	} else {
		goto L161
	}
L160:
	;
	v35 = v122 + int32(1)
	goto L7
L161:
	;
	v35 = v122 + int32(1)
	goto L7
L162:
	;
	v443 = F_pg_strftime(m, v434, int32(128), int32(_a_F_log_status_format_14), v441)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	if v126 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v434
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(240))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L13
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	F_appendStringInfoString(m, l0, v12+int32(544))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L13
	} else {
		goto L168
	}
L167:
	;
	v35 = v122 + int32(1)
	goto L7
L168:
	;
	v35 = v122 + int32(1)
	goto L7
L169:
	;
	F_gettimeofday(m, int32(_a_F_log_status_format_10))
	mBase = m.M
	v467 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])) = uint8(v467)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v470 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[23]))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+272)) = uint32(v470)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[24]))
	v475 = base.I32_div_s(v473, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+276)) = v475
	v478 = v12 + int32(544)
	v483 = F_pg_snprintf(m, v478, int32(128), int32(_a_F_log_status_format_15), v12+int32(272))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L13
	} else {
		goto L172
	}
L172:
	;
	if v126 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v478
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(256))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L13
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	F_appendStringInfoString(m, l0, v12+int32(544))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L13
	} else {
		goto L177
	}
L176:
	;
	v35 = v122 + int32(1)
	goto L7
L177:
	;
	v35 = v122 + int32(1)
	goto L7
L178:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v514 = F_pg_localtime(m, v12+int32(544), v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L13
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v126 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v516 = F_pg_strftime(m, int32(_a_F_log_status_format_16), int32(128), int32(_a_F_log_status_format_14), v514)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+292)) = int32(_a_F_log_status_format_16)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+288)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(288))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L13
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_16))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L13
	} else {
		goto L187
	}
L186:
	;
	v35 = v122 + int32(1)
	goto L7
L187:
	;
	v35 = v122 + int32(1)
	goto L7
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(544)))) = int32(0)
	v539 = int32(_a_F_log_status_format_17)
	goto L191
L189:
	;
	goto L190
L190:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L197
	}
L191:
	;
	if v126 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+308)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v12)+304)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(304))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L13
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v12)+544))
	F_appendBinaryStringInfo(m, l0, v539, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L196
	}
L195:
	;
	v35 = v122 + int32(1)
	goto L7
L196:
	;
	v35 = v122 + int32(1)
	goto L7
L197:
	;
	v557 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v557-v557)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L13
	} else {
		goto L198
	}
L198:
	;
	v35 = v122 + int32(1)
	goto L7
L199:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+296)))
	if v566 != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v584 = int32(_a_F_log_status_format_18)
	goto L201
L201:
	;
	if v126 != 0 {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	v580 = v565
	goto L204
L203:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v565)+140))
	v573 = int32(0)
	v576 = F_pg_getnameinfo_all(m, v565+int32(12), v569, v565+int32(296), int32(64), v573, v573, int32(3))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L13
	} else {
		goto L205
	}
L204:
	;
	v584 = v580 + int32(296)
	goto L201
L205:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	v580 = v579
	goto L204
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+324)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v12)+320)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(320))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L13
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	F_appendStringInfoString(m, l0, v584)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L13
	} else {
		goto L210
	}
L209:
	;
	v35 = v122 + int32(1)
	goto L7
L210:
	;
	v35 = v122 + int32(1)
	goto L7
L211:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L228
	}
L212:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v599)+276))
	if v602 == int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	if v126 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v599)+292))
	if v605 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	F_appendStringInfoString(m, l0, v602)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L13
	} else {
		goto L224
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+356)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v12)+352)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(352))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L13
	} else {
		goto L223
	}
L218:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	if v608 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+388)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v12)+384)) = v602
	v616 = F_psprintf(m, int32(_a_F_log_status_format_19), v12+int32(384))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L13
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+372)) = v616
	*(*int32)(unsafe.Add(mBase, uint32(v12)+368)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(368))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L221
	}
L221:
	;
	F_pfree(m, v616)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L13
	} else {
		goto L222
	}
L222:
	;
	v35 = v122 + int32(1)
	goto L7
L223:
	;
	v35 = v122 + int32(1)
	goto L7
L224:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+292))
	if v642 == int32(0) {
		goto L15
	} else {
		goto L225
	}
L225:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
	if v645 == int32(0) {
		goto L15
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+336)) = v642
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_20), v12+int32(336))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L13
	} else {
		goto L227
	}
L227:
	;
	v35 = v122 + int32(1)
	goto L7
L228:
	;
	v660 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v660-v660)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L13
	} else {
		goto L229
	}
L229:
	;
	v35 = v122 + int32(1)
	goto L7
L230:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L238
	}
L231:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v668)+276))
	if v671 == int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	if v126 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+404)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v12)+400)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(400))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L13
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	F_appendStringInfoString(m, l0, v671)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L13
	} else {
		goto L237
	}
L236:
	;
	v35 = v122 + int32(1)
	goto L7
L237:
	;
	v35 = v122 + int32(1)
	goto L7
L238:
	;
	v691 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v691-v691)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L13
	} else {
		goto L239
	}
L239:
	;
	v35 = v122 + int32(1)
	goto L7
L240:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L249
	}
L241:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v699)+40))
	if v702 == int32(-1) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	if v126 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v699)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+452)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v12)+448)) = v702
	v709 = v12 + int32(544)
	v714 = F_pg_snprintf(m, v709, int32(127), int32(_a_F_log_status_format_21), v12+int32(448))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L13
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v699)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+420)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(v12)+416)) = v702
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_21), v12+int32(416))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L13
	} else {
		goto L248
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+432)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+436)) = v709
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(432))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L13
	} else {
		goto L247
	}
L247:
	;
	v35 = v122 + int32(1)
	goto L7
L248:
	;
	v35 = v122 + int32(1)
	goto L7
L249:
	;
	v739 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v739-v739)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L13
	} else {
		goto L250
	}
L250:
	;
	v35 = v122 + int32(1)
	goto L7
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+484)) = v747
	*(*int32)(unsafe.Add(mBase, uint32(v12)+480)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_22), v12+int32(480))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L13
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+464)) = v747
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_23), v12+int32(464))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L13
	} else {
		goto L255
	}
L254:
	;
	v35 = v122 + int32(1)
	goto L7
L255:
	;
	v35 = v122 + int32(1)
	goto L7
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+500)) = int32(_a_F_log_status_format_24)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+496)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(496))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L13
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_24))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L13
	} else {
		goto L260
	}
L259:
	;
	v35 = v122 + int32(1)
	goto L7
L260:
	;
	v35 = v122 + int32(1)
	goto L7
L261:
	;
	if v126 != 0 {
		goto L265
	} else {
		goto L266
	}
L262:
	;
	v829 = int64(0)
	goto L261
L263:
	;
	goto L264
L264:
	;
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v824)+392))
	v829 = v828
	goto L261
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+536)) = v829
	*(*int32)(unsafe.Add(mBase, uint32(v12)+528)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_25), v12+int32(528))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L13
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+512)) = v829
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_26), v12+int32(512))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L13
	} else {
		goto L269
	}
L268:
	;
	v35 = v122 + int32(1)
	goto L7
L269:
	;
	goto L15
L270:
	;
	goto L5
L271:
	;
	v35 = v35 + int32(2)
	goto L7
}
