package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgr_bulk_flush(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v390 int32
	_ = v390
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(256)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = l0 + int32(16)
	if int32(2) <= v25 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v23 + int32(256)
	return
L4:
	;
	F_pg_qsort(m, v27, v25, int32(12), int32(1116))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v34 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v37 = int32(1)
	if v25 <= int32(0) {
		v144 = v37
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if int32(0) < v25 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = int32(0)
	F_XLogEnsureRecordSpace(m, int32(31), v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L21
	}
L13:
	;
	if v25 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v2
	v48 = v37
	v51 = v2
	goto L17
L15:
	;
	v102 = v2
	v105 = v37
	goto L16
L16:
	;
	if v25&int32(1) == int32(0) {
		v144 = v105
		goto L12
	} else {
		goto L20
	}
L17:
	;
	v64 = int32(2)
	v65 = v45 << (uint(v64) % 32)
	v67 = v23 + int32(128)
	v69 = int32(12)
	v71 = v27 + v45*v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v67))) = v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v23))) = v75
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
	v79 = v45 | int32(1)
	v81 = v79 << (uint(v64) % 32)
	v87 = v27 + v79*v69
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v81+v67))) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v81))) = v91
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+8)))
	v95 = v77 & v93 & v48
	v97 = v45 + v64
	v99 = v51 + v64
	if v99 != v25&int32(2147483646) {
		v45 = v97
		v48 = v95
		v51 = v99
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v102 = v97
	v105 = v95
	goto L16
L19:
	;
	goto L18
L20:
	;
	v126 = v102 << (uint(int32(2)) % 32)
	v132 = v27 + v102*int32(12)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126+(v23+int32(128))))) = v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v23))) = v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+8)))
	v144 = v138 & v105
	goto L12
L21:
	;
	if int32(0) < v25 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v144&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L11
L25:
	;
	v175 = int32(9)
	goto L27
L26:
	;
	v175 = int32(1)
	goto L27
L27:
	;
	v179 = v166
	goto L28
L28:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	goto L24
L30:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[873]))
	v202 = int32(0)
	if v202 < v201 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v205 = v201
	goto L33
L32:
	;
	v205 = v202
	goto L33
L33:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[377]))
	v209 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	v212 = v179
	v213 = int32(0)
	v221 = v209
	goto L36
L34:
	;
	if v263 < v25 {
		v179 = v263
		goto L28
	} else {
		goto L57
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L54
	}
L36:
	;
	v232 = v212 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v23+v232)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232+(v23+int32(128)))))
	if v221 <= v213 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v272 = F_XLogInsert(m, int32(0), int32(176))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L46
	}
L38:
	;
	v240 = v213 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[370])) = v240
	v242 = v240
	goto L40
L39:
	;
	v242 = v221
	goto L40
L40:
	;
	if v213 == v205 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	v246 = v207 + v213*int32(8260)
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	*(*int64)(unsafe.Add(mBase, uint32(v246)+4)) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v160+int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v246)+20)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v161
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)) = uint8(v175)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = int32(0)
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v256)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v246 + int32(32)
	v263 = v212 + v256
	if base.Ui32(v213) <= base.Ui32(int32(30)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v263 < v25 {
		v212 = v263
		v213 = v213 + int32(1)
		v221 = v242
		goto L36
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L37
L45:
	;
	goto L44
L46:
	;
	if v212 < v179 {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v282 = v179
	goto L48
L48:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v23+v282<<(uint(int32(2))%32))))
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+14)))
	if v303 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L34
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+4)) = base.I32_wrap_i64(v272)
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = base.I32_wrap_i64(int64(base.Ui64(v272) >> (uint(int64(32)) % 64)))
	goto L52
L51:
	;
	goto L52
L52:
	;
	if v282 < v212 {
		v282 = v282 + int32(1)
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	F_errmsg_internal(m, int32(135342), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(493403), int32(320), int32(317246))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	goto L29
L58:
	;
	v390 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L61:
	;
	v408 = v27 + v390*int32(12)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+14)))
	if v411 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L60
L63:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if base.Ui32(v419) <= base.Ui32(v410) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	goto L63
L65:
	;
	v414 = F_DataChecksumsEnabled(m)
	mBase = m.M
	if v414 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v417 = F_pg_checksum_page(m, v409, v410)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v409)+8)) = uint16(v417)
	goto L64
L67:
	;
	F_pfree(m, v409)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L7
	} else {
		goto L80
	}
L68:
	;
	if base.Ui32(v419) < base.Ui32(v410) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v409
	F_smgrwritev(m, v483, v482, v410, v23+int32(128), int32(1))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L79
	}
L71:
	;
	v423 = v419
	goto L74
L72:
	;
	goto L73
L73:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_smgrextend(m, v473, v474, v410, v409, int32(1))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L78
	}
L74:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_smgrextend(m, v442, v443, v423, int32(1630208), int32(1))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L76
	}
L75:
	;
	goto L73
L76:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v450 = v448 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+400)) = v450
	if base.Ui32(v450) < base.Ui32(v410) {
		v423 = v450
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+400)) = v478 + int32(1)
	goto L67
L79:
	;
	goto L67
L80:
	;
	v513 = v390 + int32(1)
	if v513 != v25 {
		v390 = v513
		goto L61
	} else {
		goto L81
	}
L81:
	;
	goto L62
}
func F_smgr_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	switch v11&int32(240) - int32(16) {
	case 0:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		F_GetRelationPath(m, v7+int32(88), v18, v19, v20, int32(-1), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, v7+int32(88))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				m.G0 = v7 + int32(160)
				return
			}
		}
	default:
		m.G0 = v7 + int32(160)
		return
	case 16:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		F_GetRelationPath(m, v7+int32(16), v31, v32, v33, int32(-1), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(16)
			F_appendStringInfo(m, l0, int32(470419), v7)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				m.G0 = v7 + int32(160)
				return
			}
		}
	}
}
