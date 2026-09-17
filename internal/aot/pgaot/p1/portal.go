package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PortalRun(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	v7 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(208)
	m.G0 = v22
	v32 = v7
	v33 = v7
	v34 = v7
	v35 = v7
	v36 = v7
	v37 = v7
	v38 = v7
	v39 = v7
	v40 = v7
	v41 = v7
	v42 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v22 + int32(208)
	return v338 & int32(1)
L3:
	;
	goto L2
L4:
	;
	if v42 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v394 = int32(m.ExcTag)
	v395 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v394 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L7:
	;
	v47 = base.B2i32(l5 == int32(0))
	if l5 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v156 = v32
	v157 = v33
	v158 = v34
	v159 = v35
	v160 = v36
	v161 = v37
	v162 = v38
	v163 = v39
	v164 = v40
	v165 = v41
	goto L9
L9:
	;
	if v165 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PortalRun[0])))
	if v62 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L24
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v65 == int32(4) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	v79 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v79 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	F_errmsg_internal(m, int32(_a_F_PortalRun_0), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	F_getrusage(m, int32(_a_F_PortalRun_1))
	mBase = m.M
	F_gettimeofday(m, int32(_a_F_PortalRun_2))
	mBase = m.M
	goto L23
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v47)
	F_errfinish(m, int32(_a_F_PortalRun_3), int32(708), int32(_a_F_PortalRun_0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L14
L24:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[1]))
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[2]))
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[3]))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[4]))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[5]))
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[6]))
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[7]))
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[8]))
	goto L25
L25:
	;
	v150 = v22 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v22 + int32(12)
	goto L28
L26:
	;
	v156 = v138
	v157 = v142
	v158 = v136
	v159 = v134
	v160 = v140
	v161 = v144
	v162 = v146
	v163 = v148
	v164 = v47
	v165 = int32(0)
	goto L9
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[1])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[2])) = v158
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[6])) = v161
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[4])) = v160
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[7]))
	if v156 == v162 {
		goto L62
	} else {
		goto L63
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[6])) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[2])) = v22 + int32(16)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v174 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[1])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[2])) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v307 = v164 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v307)
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L6
	} else {
		goto L54
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[5])) = v174
	goto L35
L34:
	;
	goto L35
L35:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[3])) = v178
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[4])) = v178
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	switch v182 {
	case 0:
		goto L38
	case 1, 2, 3:
		goto L39
	case 4:
		goto L37
	default:
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v260 = v164 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v260)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L51
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v229 = v164 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v229)
	F_PortalRunMulti(m, l0, l2, int32(0), l3, l4, l5)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L46
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v205 = int32(1)
	v206 = v164 & v205
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v206)
	v209 = F_PortalRunSelect(m, l0, v205, l1, l3)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L42
	}
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v183 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v193 = v164 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v193)
	F_FillPortalStore(m, l0, l2)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	if v206 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v338 = v219
	goto L29
L44:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v211 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v211
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v229)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(4)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v246 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	m.T0[v246].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v338 = int32(1)
	goto L29
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L49
L51:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v260)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v266
	F_errmsg_internal(m, int32(_a_F_PortalRun_4), v22)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v260)
	F_errfinish(m, int32(_a_F_PortalRun_3), int32(800), int32(_a_F_PortalRun_0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[6])) = v161
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[4])) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[7]))
	if v156 == v162 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v321 = v319
	goto L57
L56:
	;
	v321 = v156
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[3])) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[8]))
	if v157 == v163 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v329 = v327
	goto L60
L59:
	;
	v329 = v157
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[5])) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v307)
	F_pg_re_throw(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v353 = v351
	goto L64
L63:
	;
	v353 = v156
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[3])) = v353
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRun[8]))
	if v157 == v163 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v359 = v357
	goto L67
L66:
	;
	v359 = v157
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRun[5])) = v359
	v362 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PortalRun[0])))
	if v362 != int32(1) {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v365 == int32(4) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+172)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v163
	v377 = v164 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)) = uint8(v377)
	F_ShowUsage(m, int32(_a_F_PortalRun_5))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	goto L5
L71:
	;
	v399 = int32(v395)
	m.G0 = v22
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if v22+int32(12) == v405 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	m.ExcPending = 1
	goto L80
L73:
	;
	if v409 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v409 = v407
	goto L76
L75:
	;
	v409 = int32(0)
	goto L76
L76:
	;
	goto L73
L77:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+207)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v22)+196))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v22)+192))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v22)+180))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v22)+176))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v22)+172))
	v32 = v416
	v33 = v414
	v34 = v417
	v35 = v418
	v36 = v415
	v37 = v413
	v38 = v412
	v39 = v411
	v40 = v410
	v41 = v401
	v42 = v409
	goto L1
L78:
	;
	goto L79
L79:
	;
	F___wasm_longjmp(m, v402, v401)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	return int32(0)
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
