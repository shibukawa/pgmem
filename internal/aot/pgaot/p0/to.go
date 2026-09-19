package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddToDataDirLockFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(_a_F_AddToDataDirLockFile_0)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v3
	v22 = F_open(m, int32(_a_F_AddToDataDirLockFile_1), int32(2), v14+int32(112))
	mBase = m.M
	if v22 < v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(_a_F_AddToDataDirLockFile_0)
	return
L2:
	;
	v27 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v43 = int32(_a_F_AddToDataDirLockFile_2)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(167772187)
	v50 = F_read(m, v22, v14+int32(_a_F_AddToDataDirLockFile_3), int32(_a_F_AddToDataDirLockFile_4))
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	if v50 < v53 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return
L6:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_AddToDataDirLockFile_1)
	F_errmsg(m, int32(_a_F_AddToDataDirLockFile_5), v14)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_AddToDataDirLockFile_6), int32(1604), int32(_a_F_AddToDataDirLockFile_7))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	v59 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v77 = v14 + int32(_a_F_AddToDataDirLockFile_3)
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v50))) = uint8(v79)
	v81 = int32(1)
	if l0 < int32(2) {
		v123 = v81
		v127 = v77
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v75 = F_close(m, v22)
	mBase = m.M
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_AddToDataDirLockFile_1)
	F_errmsg(m, int32(_a_F_AddToDataDirLockFile_8), v14+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_AddToDataDirLockFile_6), int32(1615), int32(_a_F_AddToDataDirLockFile_7))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = l1
	v256 = F_pg_snprintf(m, v245, int32(_a_F_AddToDataDirLockFile_9)-v242, int32(_a_F_AddToDataDirLockFile_10), v14+int32(96))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L65
	}
L22:
	;
	v131 = v14 + int32(_a_F_AddToDataDirLockFile_3)
	v132 = v127 - v131
	if v132 != 0 {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v88 = v81
	v92 = v77
	goto L24
L24:
	;
	v95 = int32(10)
	v96 = F___strchrnul(m, v92, v95)
	mBase = m.M
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v98 == v95 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v111 = v14 + int32(_a_F_AddToDataDirLockFile_3)
	v112 = v106 - v111
	if v112 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if v102 == int32(0) {
		v123 = v88
		v127 = v92
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v102 = v96
	goto L29
L28:
	;
	v102 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	v105 = int32(1)
	v106 = v102 + v105
	v108 = v88 + v105
	if v108 != l0 {
		v88 = v108
		v92 = v106
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	base.MemoryCopy(m, v14+int32(128), v111, v112)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v242 = v112
	v245 = v14 + int32(128) + v112
	v247 = v106
	goto L21
L35:
	;
	base.MemoryCopy(m, v14+int32(128), v131, v132)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v138 = v14 + int32(128) + v132
	if l0 <= v123 {
		v242 = v132
		v245 = v138
		v247 = v127
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v142 = (l0 - v123) & int32(3)
	if v142 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v123-l0) {
		v242 = v174
		v245 = v177
		v247 = v127
		goto L21
	} else {
		goto L49
	}
L40:
	;
	v174 = v132
	v176 = v123
	v177 = v138
	goto L39
L41:
	;
	goto L42
L42:
	;
	v148 = v132
	v150 = v123
	v151 = v138
	v155 = v3
	goto L43
L43:
	;
	v157 = v150 + int32(1)
	if v148 <= int32(_a_F_AddToDataDirLockFile_4) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v174 = v164
	v176 = v157
	v177 = v167
	goto L39
L45:
	;
	v160 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v160)
	v164 = v148 + int32(1)
	goto L47
L46:
	;
	v164 = v148
	goto L47
L47:
	;
	v167 = v14 + int32(128) + v164
	v169 = v155 + int32(1)
	if v169 != v142 {
		v148 = v164
		v150 = v157
		v151 = v167
		v155 = v169
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	v188 = v174
	v190 = v176
	v191 = v177
	goto L50
L50:
	;
	if v188 <= int32(_a_F_AddToDataDirLockFile_4) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v242 = v232
	v245 = v235
	v247 = v127
	goto L21
L52:
	;
	v198 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v198)
	v202 = v188 + int32(1)
	goto L54
L53:
	;
	v202 = v188
	goto L54
L54:
	;
	if v202 <= int32(_a_F_AddToDataDirLockFile_4) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v208 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(128)+v202))) = uint8(v208)
	v212 = v202 + int32(1)
	goto L57
L56:
	;
	v212 = v202
	goto L57
L57:
	;
	if v212 <= int32(_a_F_AddToDataDirLockFile_4) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v218 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(128)+v212))) = uint8(v218)
	v222 = v212 + int32(1)
	goto L60
L59:
	;
	v222 = v212
	goto L60
L60:
	;
	if v222 <= int32(_a_F_AddToDataDirLockFile_4) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v228 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(128)+v222))) = uint8(v228)
	v232 = v222 + int32(1)
	goto L63
L62:
	;
	v232 = v222
	goto L63
L63:
	;
	v235 = v14 + int32(128) + v232
	v237 = v190 + int32(4)
	if v237 != l0 {
		v188 = v232
		v190 = v237
		v191 = v235
		goto L50
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	v258 = int32(10)
	v259 = F___strchrnul(m, v247, v258)
	mBase = m.M
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v261 == v258 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v265 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v265 = v259
	goto L69
L68:
	;
	v265 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v266 = F_strlen(m, v245)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v265 + int32(1)
	v277 = F_pg_snprintf(m, v266+v245, int32(_a_F_AddToDataDirLockFile_9)-(v266+v242), int32(_a_F_AddToDataDirLockFile_11), v14+int32(80))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v281 = v14 + int32(128)
	v282 = F_strlen(m, v281)
	mBase = m.M
	v284 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[1])) = v284
	v286 = int32(_a_F_AddToDataDirLockFile_2)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = int32(167772189)
	v291 = F_pwrite(m, v22, v281, v282, int64(0))
	mBase = m.M
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v284
	if v282 != v291 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[1]))
	if v298 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(167772188)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[2])))
	if v329 != int32(1) {
		v343 = int32(0)
		goto L89
	} else {
		goto L90
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[1])) = int32(51)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v306 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	if v306 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v322 = F_close(m, v22)
	mBase = m.M
	goto L1
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(_a_F_AddToDataDirLockFile_1)
	F_errmsg(m, int32(_a_F_AddToDataDirLockFile_12), v14-int32(-64))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_AddToDataDirLockFile_6), int32(1679), int32(_a_F_AddToDataDirLockFile_7))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[0]))
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v367))) = v368
	v370 = F_close(m, v22)
	mBase = m.M
	if v370 == v368 {
		goto L1
	} else {
		goto L101
	}
L88:
	;
	if v343 == int32(0) {
		goto L87
	} else {
		goto L95
	}
L89:
	;
	goto L88
L90:
	;
	goto L91
L91:
	;
	v334 = F_fsync(m, v22)
	mBase = m.M
	if v334 != int32(-1) {
		v343 = v334
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v343 = int32(-1)
	goto L89
L93:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_AddToDataDirLockFile[1]))
	if v338 == int32(27) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v348 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	if v348 == int32(0) {
		goto L87
	} else {
		goto L97
	}
L97:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_AddToDataDirLockFile_1)
	F_errmsg(m, int32(_a_F_AddToDataDirLockFile_12), v14+int32(48))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_AddToDataDirLockFile_6), int32(1690), int32(_a_F_AddToDataDirLockFile_7))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	goto L87
L101:
	;
	v375 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	if v375 == int32(0) {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_AddToDataDirLockFile_1)
	F_errmsg(m, int32(_a_F_AddToDataDirLockFile_12), v14+int32(32))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_AddToDataDirLockFile_6), int32(1698), int32(_a_F_AddToDataDirLockFile_7))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	goto L1
}
func F_CopyToBinaryEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(_a_F_CopyToBinaryEnd_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v9, v5+int32(14), int32(2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_CopySendEndOfRow(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_CopyToCSVOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L28
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v49 = int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v50 <= v49 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = F_OutputFunctionCall(m, v18+v22*int32(28), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v45 = F_strlen(m, v44)
	mBase = m.M
	F_appendBinaryStringInfo(m, v43, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v22))))
	F_CopyAttributeOutCSV(m, l0, v36, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	goto L4
L12:
	;
	v55 = v49
	goto L13
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v65 = int32(2)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v55<<(uint(v65)%32))))
	v69 = int32(1)
	v70 = v68 - v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v71))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v70<<(uint(v65)%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v85 <= v82+v69 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L1
L15:
	;
	if v73&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	F_appendStringInfoChar(m, v81, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89+v82))) = uint8(v80)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v95 = v93 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97+v95))) = uint8(v99)
	goto L15
L19:
	;
	goto L15
L20:
	;
	v122 = v55 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v122 < v123 {
		v55 = v122
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v107 = F_strlen(m, v106)
	mBase = m.M
	F_appendBinaryStringInfo(m, v105, v106, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v113 = F_OutputFunctionCall(m, v18+v70*int32(28), v78)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v70))))
	F_CopyAttributeOutCSV(m, l0, v113, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L14
L28:
	;
	return
}
func F_CopyToTextLikeOutFunc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_getTypeOutputInfo(m, l1, v6+int32(12), v6+int32(11))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		F_fmgr_info(m, v14, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_CopyToTextLikeStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = F_pg_server_to_any(m, v12, v13, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v15
	goto L3
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L33
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v35 = l1 + v25<<(uint(int32(4))%32) + v30*int32(100) - int32(76)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v44 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v45 <= v44 {
		goto L9
	} else {
		goto L18
	}
L13:
	;
	F_CopyAttributeOutText(m, l0, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_CopyAttributeOutCSV(m, l0, v35, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	goto L12
L18:
	;
	v51 = v44
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v51<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v67 <= v64+int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L9
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = l1 + v85<<(uint(int32(4))%32) + v60*int32(100) - int32(76)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v94 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_appendStringInfoChar(m, v63, v62)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v64))) = uint8(v62)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v77 = v75 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v77))) = uint8(v81)
	goto L21
L25:
	;
	goto L21
L26:
	;
	v103 = v51 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v103 < v104 {
		v51 = v103
		goto L19
	} else {
		goto L32
	}
L27:
	;
	F_CopyAttributeOutCSV(m, l0, v93, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_CopyAttributeOutText(m, l0, v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	goto L26
L32:
	;
	goto L20
L33:
	;
	goto L8
}
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_exprType(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(16) {
			v26 = l1
			v27 = F_expression_returns_set(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg(m, int32(_a_F_coerce_to_boolean_0), v9)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = F_exprLocation(m, v26)
								mBase = m.M
								F_parser_errposition(m, l0, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_coerce_to_boolean_1), int32(1190), int32(_a_F_coerce_to_boolean_2))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
					m.G0 = v9 + int32(32)
					return v26
				}
			}
		} else {
			v18 = int32(-1)
			v22 = F_coerce_to_target_type(m, l0, l1, v11, int32(16), v18, int32(1), int32(2), v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_format_type_be(m, v11)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_coerce_to_boolean_3)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
								F_errmsg(m, int32(_a_F_coerce_to_boolean_4), v9+int32(16))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = F_exprLocation(m, l1)
									mBase = m.M
									F_parser_errposition(m, l0, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_coerce_to_boolean_1), int32(1180), int32(_a_F_coerce_to_boolean_2))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
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
				} else {
					v26 = v22
					v27 = F_expression_returns_set(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
									F_errmsg(m, int32(_a_F_coerce_to_boolean_0), v9)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = F_exprLocation(m, v26)
										mBase = m.M
										F_parser_errposition(m, l0, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_coerce_to_boolean_1), int32(1190), int32(_a_F_coerce_to_boolean_2))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
							m.G0 = v9 + int32(32)
							return v26
						}
					}
				}
			}
		}
	}
}
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(160)
	switch l1 - int32(8) {
	case 0:
		v48 = v15
		v49 = int32(_a_F_encode_to_ascii_0)
		if base.Ui32(v14) < base.Ui32(int32(20)) {
		} else {
			v52 = int32(4)
			v53 = l0 + v52
			if v14&v52 != 0 {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v57 = base.I32_extend8_s(v56)
				if int32(0) <= v57 {
					v65 = v57
				} else {
					if base.Ui32(v56) < base.Ui32(v48) {
						v65 = int32(32)
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v56-v48)))))
						v65 = v64
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v65)
				v69 = l0 + int32(5)
			} else {
				v69 = v53
			}
			v72 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v72 == int32(5) {
			} else {
				v77 = v69
				for {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
					v86 = base.I32_extend8_s(v85)
					if int32(0) <= v86 {
						v94 = v86
					} else {
						if base.Ui32(v85) < base.Ui32(v48) {
							v94 = int32(32)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v85-v48)))))
							v94 = v93
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v94)
					v97 = v77 + int32(1)
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					v99 = base.I32_extend8_s(v98)
					if int32(0) <= v99 {
						v107 = v99
					} else {
						if base.Ui32(v98) < base.Ui32(v48) {
							v107 = int32(32)
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v98-v48)))))
							v107 = v106
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v107)
					v110 = v77 + int32(2)
					if v110 != l0+v72 {
						v77 = v110
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 1:
		v48 = v15
		v49 = int32(_a_F_encode_to_ascii_1)
		if base.Ui32(v14) < base.Ui32(int32(20)) {
		} else {
			v52 = int32(4)
			v53 = l0 + v52
			if v14&v52 != 0 {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v57 = base.I32_extend8_s(v56)
				if int32(0) <= v57 {
					v65 = v57
				} else {
					if base.Ui32(v56) < base.Ui32(v48) {
						v65 = int32(32)
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v56-v48)))))
						v65 = v64
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v65)
				v69 = l0 + int32(5)
			} else {
				v69 = v53
			}
			v72 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v72 == int32(5) {
			} else {
				v77 = v69
				for {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
					v86 = base.I32_extend8_s(v85)
					if int32(0) <= v86 {
						v94 = v86
					} else {
						if base.Ui32(v85) < base.Ui32(v48) {
							v94 = int32(32)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v85-v48)))))
							v94 = v93
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v94)
					v97 = v77 + int32(1)
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					v99 = base.I32_extend8_s(v98)
					if int32(0) <= v99 {
						v107 = v99
					} else {
						if base.Ui32(v98) < base.Ui32(v48) {
							v107 = int32(32)
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v98-v48)))))
							v107 = v106
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v107)
					v110 = v77 + int32(2)
					if v110 != l0+v72 {
						v77 = v110
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if base.Ui32(l1) <= base.Ui32(int32(41)) {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_c_F_encode_to_ascii[0])))
					v37 = v35
				} else {
					v37 = int32(_a_F_encode_to_ascii_2)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v37
				F_errmsg(m, int32(_a_F_encode_to_ascii_3), v12)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_encode_to_ascii_4), int32(78), int32(_a_F_encode_to_ascii_5))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 8:
		v48 = v15
		v49 = int32(_a_F_encode_to_ascii_6)
		if base.Ui32(v14) < base.Ui32(int32(20)) {
		} else {
			v52 = int32(4)
			v53 = l0 + v52
			if v14&v52 != 0 {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v57 = base.I32_extend8_s(v56)
				if int32(0) <= v57 {
					v65 = v57
				} else {
					if base.Ui32(v56) < base.Ui32(v48) {
						v65 = int32(32)
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v56-v48)))))
						v65 = v64
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v65)
				v69 = l0 + int32(5)
			} else {
				v69 = v53
			}
			v72 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v72 == int32(5) {
			} else {
				v77 = v69
				for {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
					v86 = base.I32_extend8_s(v85)
					if int32(0) <= v86 {
						v94 = v86
					} else {
						if base.Ui32(v85) < base.Ui32(v48) {
							v94 = int32(32)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v85-v48)))))
							v94 = v93
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v94)
					v97 = v77 + int32(1)
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					v99 = base.I32_extend8_s(v98)
					if int32(0) <= v99 {
						v107 = v99
					} else {
						if base.Ui32(v98) < base.Ui32(v48) {
							v107 = int32(32)
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v98-v48)))))
							v107 = v106
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v107)
					v110 = v77 + int32(2)
					if v110 != l0+v72 {
						v77 = v110
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 21:
		v48 = int32(128)
		v49 = int32(_a_F_encode_to_ascii_7)
		if base.Ui32(v14) < base.Ui32(int32(20)) {
		} else {
			v52 = int32(4)
			v53 = l0 + v52
			if v14&v52 != 0 {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v57 = base.I32_extend8_s(v56)
				if int32(0) <= v57 {
					v65 = v57
				} else {
					if base.Ui32(v56) < base.Ui32(v48) {
						v65 = int32(32)
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v56-v48)))))
						v65 = v64
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v65)
				v69 = l0 + int32(5)
			} else {
				v69 = v53
			}
			v72 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v72 == int32(5) {
			} else {
				v77 = v69
				for {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
					v86 = base.I32_extend8_s(v85)
					if int32(0) <= v86 {
						v94 = v86
					} else {
						if base.Ui32(v85) < base.Ui32(v48) {
							v94 = int32(32)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v85-v48)))))
							v94 = v93
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v94)
					v97 = v77 + int32(1)
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					v99 = base.I32_extend8_s(v98)
					if int32(0) <= v99 {
						v107 = v99
					} else {
						if base.Ui32(v98) < base.Ui32(v48) {
							v107 = int32(32)
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v98-v48)))))
							v107 = v106
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v107)
					v110 = v77 + int32(2)
					if v110 != l0+v72 {
						v77 = v110
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	}
}
func F_to_oct64(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14026(m, l0, int64(3), int64(8), int32(7))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_to_regclass(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14027(m, l0, int32(1481))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_to_regtype(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14027(m, l0, int32(1238))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_to_tsquery_byid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14028(m, l0, int32(0), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_to_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1156), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
