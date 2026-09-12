package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessWalSndrMessage(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v167 int64
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	v18 = m.G0
	v19 = int32(16)
	v20 = v18 - v19
	m.G0 = v20
	F___gettimeofday(m, v20)
	mBase = m.M
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	m.G0 = v20 + v19
	v32 = v24 + v23*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1456)) = int32(1)
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	F_s_lock(m, v14+int32(1456), int32(489764), int32(1263), int32(401032))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
	if base.Ui64(v43) < base.Ui64(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	goto L4
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = l1
	goto L9
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1456)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = l0
	v51 = int32(13)
	goto L12
L10:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	goto L10
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	goto L15
L13:
	;
	v70 = int32(0)
	goto L21
L15:
	;
	goto L16
L16:
	;
	goto L18
L18:
	;
	if v58 == int32(15) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v58 <= v51 {
		v85 = int32(1)
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	if v74 != int32(2) {
		v85 = v70
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[163])))
	if v78 != 0 {
		v85 = v70
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v85 = int32(0) | base.B2i32(v82 <= v51)
	goto L11
L24:
	;
	v87 = F_timestamptz_to_str(m, l1)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	m.G0 = v11 + int32(32)
	return
L27:
	;
	v89 = F_pstrdup(m, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v91 = F_timestamptz_to_str(m, v32)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v93 = F_pstrdup(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1456)) = int32(1)
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v182 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L51
	}
L32:
	;
	F_s_lock(m, v96+int32(1456), int32(489201), int32(372), int32(26630))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1456)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
	v111 = F_GetXLogReplayRecPtr(m, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	if v109 != v111 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+96)) = int32(1)
	if v116 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v175 = int32(0)
	goto L39
L39:
	;
	v179 = v175
	goto L31
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	F_s_lock(m, v120+int32(96), int32(486941), int32(4672), int32(371831))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+96)) = int32(0)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v129)+72))
	if v133 == int64(0) {
		v179 = int32(-1)
		goto L31
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v139 = m.G0
	v140 = int32(16)
	v141 = v139 - v140
	m.G0 = v141
	F___gettimeofday(m, v141)
	mBase = m.M
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v141)+8)))
	m.G0 = v141 + v140
	v153 = v145 + v144*int64(1000000) - int64(946684800000000)
	goto L45
L45:
	;
	if v153 <= v133 {
		v170 = int32(0)
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v175 = v170
	goto L39
L47:
	;
	goto L46
L48:
	;
	v156 = int32(2147483647)
	v159 = v153 - v133
	if base.B2i32(int64(0) < v133)^base.B2i32(v159 < v153) != 0 {
		v170 = v156
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int64(2147483646000) < v159 {
		v170 = v156
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v167 = base.I64_div_s(v159+int64(999), int64(1000))
	v170 = base.I32_wrap_i64(v167)
	goto L47
L51:
	;
	if v179 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	F_pfree(m, v89)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L64
	}
L53:
	;
	F_errfinish(m, int32(489764), v212, int32(401032))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	if v182 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v182 == int32(0) {
		goto L52
	} else {
		goto L60
	}
L57:
	;
	v189 = F_GetReplicationTransferLatency(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v89
	F_errmsg_internal(m, int32(150179), v11)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v212 = int32(1287)
	goto L53
L60:
	;
	v200 = F_GetReplicationTransferLatency(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v89
	F_errmsg_internal(m, int32(150099), v11+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v212 = int32(1293)
	goto L53
L63:
	;
	goto L52
L64:
	;
	F_pfree(m, v93)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L26
}
func F_WalRcvFetchTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	if base.Ui32(l0) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L84
	}
L2:
	;
	v13 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(128)
	return
L5:
	;
	if v13 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v285 = v13 + int32(1)
	if base.Ui32(v285) <= base.Ui32(l1) {
		v13 = v285
		goto L5
	} else {
		goto L83
	}
L8:
	;
	v22 = F_existsTimeLineHistory(m, v13)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v22 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v26 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v13
	F_errmsg(m, int32(212342), v10+int32(32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[484]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	m.T0[v49].(func(*base.Module, int32, int32, int32, int32, int32))(m, v40, v13, v10+int32(124), v10+int32(120), v10+int32(116))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(489764), int32(741), int32(163492))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v13
	v59 = F_pg_snprintf(m, v10+int32(48), int32(64), int32(12805), v10+int32(16))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v63 = v10 + int32(48)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v87-v86 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v71 = v61
	v72 = v63
	goto L24
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v86 = v75
	v87 = v76
	goto L21
L26:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v10)+116))
	v91 = m.G0
	v93 = v91 - int32(2144)
	m.G0 = v93
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = int32(42)
	v103 = F_pg_snprintf(m, v93+int32(96), int32(1024), int32(461704), v93+int32(80))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v106 = v93 + int32(96)
	v107 = F_unlink(m, v106)
	mBase = m.M
	v111 = F_OpenTransientFile(m, v106, int32(194))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L33
	}
L30:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v267 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if v267 != int32(2) {
		goto L76
	} else {
		goto L77
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L71
	}
L32:
	;
	v218 = int32(4640180)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v222 = F_unlink(m, v93+int32(96))
	mBase = m.M
	if v219 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	if int32(0) <= v111 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(167772218)
	v122 = F_write(m, v111, v89, v90)
	mBase = m.M
	if v122 != v90 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L60
	}
L37:
	;
	v124 = int32(4094684)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126
	v129 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = int32(167772217)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v134 != int32(1) {
		v148 = v126
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = int32(0)
	v180 = F_CloseTransientFile(m, v111)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L56
	}
L39:
	;
	if v148 == int32(0) {
		goto L38
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	goto L42
L42:
	;
	v139 = F_fsync(m, v111)
	mBase = m.M
	if v139 != int32(-1) {
		v148 = v139
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v148 = int32(-1)
	goto L40
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v143 == int32(27) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, _consts[42])))
	if v154 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v157 = F_errstart(m, v155, int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L51
	}
L48:
	;
	v155 = int32(21)
	goto L50
L49:
	;
	v155 = int32(23)
	goto L50
L50:
	;
	goto L47
L51:
	;
	if v157 == int32(0) {
		goto L38
	} else {
		goto L52
	}
L52:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = v93 + int32(96)
	F_errmsg(m, int32(296453), v93+int32(48))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(493371), int32(506), int32(385304))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	goto L38
L56:
	;
	if v180 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v13
	v189 = F_pg_snprintf(m, v93+int32(1120), int32(1024), int32(12798), v93+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v196 = F_durable_rename(m, v93+int32(96), v93+int32(1120), int32(21))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	m.G0 = v93 + int32(2144)
	goto L30
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v93 + int32(96)
	F_errmsg(m, int32(296184), v93)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(493371), int32(481), int32(385304))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v225 = v219
	goto L66
L65:
	;
	v225 = int32(51)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v225
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v93 + int32(96)
	F_errmsg(m, int32(295188), v93-int32(-64))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(493371), int32(498), int32(385304))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v93 + int32(96)
	F_errmsg(m, int32(296248), v93+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(493371), int32(512), int32(385304))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	F_pfree(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L81
	}
L76:
	;
	F_XLogArchiveForceDone(m, v265)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_XLogArchiveNotify(m, v265)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L9
	} else {
		goto L80
	}
L79:
	;
	goto L75
L80:
	;
	goto L75
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	F_pfree(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	goto L7
L83:
	;
	goto L6
L84:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg_internal(m, int32(51134), v10)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(489764), int32(755), int32(163492))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WalSummarizerShutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v8 = F_LWLockAcquire(m, v4+int32(6272), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[414]))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		F_LWLockRelease(m, v15+int32(6272))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_WalUsageAccumDiff(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int64)(unsafe.Add(mBase, _consts[68]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3 + (v5 - v6)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int64)(unsafe.Add(mBase, _consts[69]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int64)(unsafe.Add(mBase, _consts[70]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int64)(unsafe.Add(mBase, _consts[71]))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24 + (v26 - v27)
	return
}
func F_WriteWalSummary(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
	v19 = F_FileWriteV(m, v12, v9+int32(40), int32(1), v11, int32(167772237))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v19 {
			if v19 != l2 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v64 = *(*int32)(unsafe.Add(mBase, _consts[413]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v62*int32(48))+32))
						v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)) = uint32(v69)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v68
						F_errmsg(m, int32(41019), v9+int32(16))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(622918), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486975), int32(312), int32(17943))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v26 + base.I64_extend_i32_u(l2)
				m.G0 = v9 + int32(48)
				return l2
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40*int32(48))+32))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
					F_errmsg(m, int32(296090), v9)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(486975), int32(305), int32(17943))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
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
}
func F_check_wal_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(-1) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[207]))
		if v8 == int32(-1) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[208]))
			v15 = base.I32_div_s(v13, int32(32))
			v17 = *(*int32)(unsafe.Add(mBase, _consts[180]))
			v19 = base.I32_div_s(v17, int32(8192))
			if v15 < v19 {
				v21 = v15
			} else {
				v21 = v19
			}
			if v21 <= int32(8) {
				v24 = int32(8)
			} else {
				v24 = v21
			}
			v29 = v24
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
		}
	} else {
		if int32(3) < v4 {
		} else {
			v29 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
		}
	}
	return int32(1)
}
func F_compareWalFileNames(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(8)
	v5 = v3 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = v6 + v4
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v12 == int32(0) {
		v31 = v11
		v32 = v12
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v32 - v31
L2:
	;
	goto L1
L3:
	;
	if v11 != v12 {
		v31 = v11
		v32 = v12
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = v5
	v17 = v8
	goto L5
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v20
		v32 = v21
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v31 = v20
	v32 = v21
	goto L2
L7:
	;
	v24 = int32(1)
	if v20 == v21 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
