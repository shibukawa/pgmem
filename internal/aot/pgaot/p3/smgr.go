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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v437 int32
	_ = v437
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(256)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = l0 + int32(16)
	if int32(2) <= v23 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v21 + int32(256)
	return
L4:
	;
	F_pg_qsort(m, v25, v23, int32(12), int32(1119))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v32 == int32(1) {
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
	v35 = int32(1)
	if v23 <= int32(0) {
		v136 = v35
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if int32(0) < v23 {
		goto L74
	} else {
		goto L75
	}
L12:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[0]))
	if v153 <= int32(31) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	if v23 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v2
	v48 = v35
	v55 = v2
	goto L17
L15:
	;
	v100 = v2
	v103 = v35
	goto L16
L16:
	;
	v118 = v100 << (uint(int32(2)) % 32)
	v124 = v25 + v100*int32(12)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v118+(v21+int32(128))))) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v118+v21))) = v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+8)))
	v136 = v130 & v103
	goto L12
L17:
	;
	v62 = int32(2)
	v63 = v45 << (uint(v62) % 32)
	v65 = v21 + int32(128)
	v67 = int32(12)
	v69 = v25 + v45*v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v63+v65))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v63))) = v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+8)))
	v77 = v45 | int32(1)
	v79 = v77 << (uint(v62) % 32)
	v83 = v25 + v77*v67
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v79))) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v79))) = v87
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
	v91 = v75 & v89 & v48
	v93 = v45 + v62
	v95 = v55 + v62
	if v95 != v23&int32(2147483646) {
		v45 = v93
		v48 = v91
		v55 = v95
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v23&int32(1) == int32(0) {
		v136 = v91
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v100 = v93
	v103 = v91
	goto L16
L21:
	;
	v156 = int32(_a_F_smgr_bulk_flush_0)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[1]))
	v160 = F_repalloc(m, v158, int32(_a_F_smgr_bulk_flush_1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[2]))
	if v213 <= int32(19) {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[1])) = v160
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[0]))
	v167 = int32(_a_F_smgr_bulk_flush_2)
	v168 = (int32(32) - v165) * v167
	v170 = v165 * v167
	v171 = v160 + v170
	if v171&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v168)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[0])) = int32(32)
	goto L23
L26:
	;
	if v165 == int32(32) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	v196 = v168
	goto L28
L28:
	;
	if v196 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L29:
	;
	v185 = v160 + int32(_a_F_smgr_bulk_flush_1)
	v188 = v170 + v160 + int32(4)
	if base.Ui32(v188) < base.Ui32(v185) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v190 = v185
	goto L32
L31:
	;
	v190 = v188
	goto L32
L32:
	;
	v196 = (v160^int32(-1)-v170+v190)&int32(-4) + int32(4)
	goto L28
L33:
	;
	base.MemoryFill(m, v171, int32(0), v196)
	goto L25
L34:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[3]))
	v219 = F_repalloc(m, v217, int32(240))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v227 = int32(0)
	if v227 < v23 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[2])) = int32(20)
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[3])) = v219
	goto L36
L38:
	;
	v233 = int32(1)
	if v136&v233 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	goto L11
L41:
	;
	v236 = int32(9)
	goto L43
L42:
	;
	v236 = v233
	goto L43
L43:
	;
	v238 = v227
	goto L44
L44:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L46
	}
L45:
	;
	goto L40
L46:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[0]))
	v259 = int32(0)
	if v259 < v258 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v262 = v258
	goto L49
L48:
	;
	v262 = v259
	goto L49
L49:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[1]))
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[4]))
	v269 = v238
	v271 = int32(0)
	v275 = v266
	goto L52
L50:
	;
	if v318 < v23 {
		v238 = v318
		goto L44
	} else {
		goto L73
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L70
	}
L52:
	;
	v287 = v269 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v21+v287)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+(v21+int32(128)))))
	if v275 <= v271 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v327 = F_XLogInsert(m, int32(0), int32(176))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L62
	}
L54:
	;
	v295 = v271 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_flush[4])) = v295
	v297 = v295
	goto L56
L55:
	;
	v297 = v275
	goto L56
L56:
	;
	if v271 == v262 {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v301 = v264 + v271*int32(_a_F_smgr_bulk_flush_2)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v301)+12)) = v302
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+4)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v301)+16)) = v151
	*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)) = uint8(v236)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+28)) = int32(0)
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v312)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+36)) = v301 + int32(32)
	v318 = v269 + v312
	if base.Ui32(v271) <= base.Ui32(int32(30)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v318 < v23 {
		v269 = v318
		v271 = v271 + int32(1)
		v275 = v297
		goto L52
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L53
L61:
	;
	goto L60
L62:
	;
	if v269 < v238 {
		goto L50
	} else {
		goto L63
	}
L63:
	;
	v338 = v238
	goto L64
L64:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v21+v338<<(uint(int32(2))%32))))
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355)+14)))
	if v356 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L50
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+4)) = base.I32_wrap_i64(v327)
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = base.I32_wrap_i64(int64(base.Ui64(v327) >> (uint(int64(32)) % 64)))
	goto L68
L67:
	;
	goto L68
L68:
	;
	if v338 < v269 {
		v338 = v338 + int32(1)
		goto L64
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	F_errmsg_internal(m, int32(_a_F_smgr_bulk_flush_3), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_smgr_bulk_flush_4), int32(320), int32(_a_F_smgr_bulk_flush_5))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	goto L45
L74:
	;
	v437 = int32(0)
	goto L77
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L77:
	;
	v453 = v25 + v437*int32(12)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454)+14)))
	if v456 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L76
L79:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if base.Ui32(v464) <= base.Ui32(v455) {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	goto L79
L81:
	;
	v459 = F_DataChecksumsEnabled(m)
	mBase = m.M
	if v459 == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v462 = F_pg_checksum_page(m, v454, v455)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v454)+8)) = uint16(v462)
	goto L80
L83:
	;
	F_pfree(m, v454)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L96
	}
L84:
	;
	if base.Ui32(v464) < base.Ui32(v455) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v454
	F_smgrwritev(m, v524, v523, v455, v21+int32(128), int32(1))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L7
	} else {
		goto L95
	}
L87:
	;
	v468 = v464
	goto L90
L88:
	;
	goto L89
L89:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_smgrextend(m, v514, v515, v455, v454, int32(1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L7
	} else {
		goto L94
	}
L90:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_smgrextend(m, v485, v486, v468, int32(_a_F_smgr_bulk_flush_6), int32(1))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v493 = v491 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+400)) = v493
	if base.Ui32(v493) < base.Ui32(v455) {
		v468 = v493
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+400)) = v519 + int32(1)
	goto L83
L95:
	;
	goto L83
L96:
	;
	v552 = v437 + int32(1)
	if v552 != v23 {
		v437 = v552
		goto L77
	} else {
		goto L97
	}
L97:
	;
	goto L78
}
func F_smgr_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v16 = v12&int32(240) - int32(16)
	if v16 != 0 {
		if v16 == int32(16) {
			v31 = v8 + int32(24)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			F_GetRelationPath(m, v31, v32, v33, v34, int32(-1), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v40
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v31
				F_appendStringInfo(m, l0, int32(_a_F_smgr_desc_0), v8)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v8 + int32(96)
					return
				}
			}
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	} else {
		v20 = v8 + int32(24)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		F_GetRelationPath(m, v20, v21, v22, v23, int32(-1), v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, v20)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v8 + int32(96)
				return
			}
		}
	}
}
