package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v216 int64
	_ = v216
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v23 = l3 + l1*int32(416)
	v26 = l2 + l1*int32(28)
	v29 = l1
	v30 = l2
	v34 = v6
	v35 = v6
	v36 = v6
	v37 = v6
	v38 = v6
	v39 = v6
	v40 = v19
	v43 = int32(-1)
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v19 + int32(48)
	return v163
L4:
	;
	goto L3
L5:
	;
	if v43 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	v298 = int32(m.ExcTag)
	v299 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v298 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L8:
	;
	v46 = int32(32)
	v47 = v40 - v46
	m.G0 = v47
	v50 = v47 - v46
	m.G0 = v50
	v53 = v50 - int32(160)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v120 = v29
	v121 = v30
	v122 = v34
	v123 = v35
	v124 = v36
	v125 = v37
	v126 = v38
	v127 = v39
	v128 = v40
	goto L10
L10:
	;
	if v127 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v38
	v64 = v37 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v47
	F_fmgr_info(m, v59, v26)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		v297 = v53
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v73 = base.B2i32(l3 == int32(0))
	if l3 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v47
	F_InstrStartNode(m, v23)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		v297 = v53
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v83 = int32(4442992)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = l4
	v87 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+18)) = uint16(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)) = uint8(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v84
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v47
	F_pgstat_init_function_usage(m, v47, v50)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		v297 = v53
		goto L7
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v105 = int32(4340100)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	*(*int32)(unsafe.Add(mBase, _consts[315])) = v107 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v114 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v19 + int32(16)
	goto L23
L21:
	;
	v120 = v47
	v121 = v53
	v122 = v50
	v123 = v112
	v124 = v114
	v125 = v73
	v126 = v84
	v127 = int32(0)
	v128 = v53
	goto L10
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v124
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v123
	v133 = int32(4340100)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	v136 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[315])) = v135 - v136
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	v146 = v125 & v136
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v146)
	F_pg_re_throw(m)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v121
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	v158 = v125 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	v163 = m.T0[v153].(func(*base.Module, int32) int32)(m, v120)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L28
	}
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v123
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v124
	v169 = int32(4340100)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	*(*int32)(unsafe.Add(mBase, _consts[315])) = v171 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	v189 = m.G0
	v191 = v189 - int32(16)
	m.G0 = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v193 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v126
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+16)))
	if v230 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	F___clock_gettime(m, int32(1), v191)
	mBase = m.M
	v196 = int32(4422768)
	v197 = *(*int64)(unsafe.Add(mBase, _consts[316]))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v122)+16))
	v200 = int64(*(*int32)(unsafe.Add(mBase, uint32(v191)+8)))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v122)+24))
	v206 = v200 + v201*int64(1000000000) - v205
	*(*int64)(unsafe.Add(mBase, _consts[316])) = v199 + v206
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v122)+8))
	goto L33
L31:
	;
	goto L32
L32:
	;
	m.G0 = v191 + int32(16)
	goto L29
L33:
	;
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = v211 + int64(1)
	goto L35
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v193)+8)) = v209 + v206
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v193)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+16)) = v216 + (v206 - v197 + v199)
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v158 != 0 {
		goto L4
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	F_errcode(m, int32(16908867))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v255
	F_errmsg(m, int32(329098), v19)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	F_errfinish(m, int32(472628), int32(2389), int32(467599))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L1
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v120
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)) = uint8(v158)
	F_InstrStopNode(m, v23, float64(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		v297 = v128
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L6
L45:
	;
	v303 = int32(v299)
	m.G0 = v297
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v19+int32(16) == v310 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	m.ExcPending = 1
	goto L54
L47:
	;
	if v313 != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v313 = v312
	goto L50
L49:
	;
	v313 = int32(0)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+35)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v29 = v314
	v30 = v316
	v34 = v315
	v35 = v319
	v36 = v320
	v37 = v317
	v38 = v318
	v39 = v305
	v40 = v297
	v43 = v313
	goto L2
L52:
	;
	goto L53
L53:
	;
	F___wasm_longjmp(m, v306, v305)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
