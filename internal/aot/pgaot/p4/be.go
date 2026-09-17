package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_be_lo_export(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(_a_F_be_lo_export_0)
	m.G0 = v12
	v17 = v2
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = int32(-1)
	v22 = v2
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	m.G0 = v12 + int32(_a_F_be_lo_export_0)
	return int32(1)
L4:
	;
	if v21 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v273 = int32(m.ExcTag)
	v274 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v273 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v20
	v32 = F_pg_detoast_datum_packed(m, v27)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v73 = v17
	v74 = v18
	v75 = v19
	v76 = v20
	v78 = v22
	goto L9
L9:
	;
	if v78 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_be_lo_export[4])) = uint8(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v20
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[5]))
	v44 = F_inv_open(m, v26, int32(_a_F_be_lo_export_1), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v44
	F_text_to_cstring_buffer(m, v32, v12+int32(208), int32(1024))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v44
	v60 = F_umask(m, int32(18))
	mBase = m.M
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[6]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[7]))
	goto L13
L13:
	;
	v67 = v12 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v12 + int32(44)
	goto L16
L14:
	;
	v73 = v65
	v74 = v63
	v75 = v60
	v76 = v44
	v78 = int32(0)
	goto L9
L16:
	;
	goto L14
L17:
	;
	goto L29
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[7])) = v12 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v90 = v12 + int32(208)
	v93 = F_OpenTransientFilePerm(m, v90, int32(577), int32(420))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[6])) = v74
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[7])) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v149 = F_umask(m, v75)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_pg_re_throw(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L27
	}
L21:
	;
	v95 = int32(_a_F_be_lo_export_2)
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[6])) = v74
	v97 = int32(_a_F_be_lo_export_3)
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[7])) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v103 = F_umask(m, v75)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[6])) = v74
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_export[7])) = v73
	if int32(0) <= v93 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errcode_for_file_access(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v90
	F_errmsg(m, int32(_a_F_be_lo_export_4), v12)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errfinish(m, int32(_a_F_be_lo_export_5), int32(527), int32(_a_F_be_lo_export_6))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v221 = F_CloseTransientFile(m, v93)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L38
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v170 = v12 + int32(1232)
	v172 = F_inv_read(m, v76, v170, int32(_a_F_be_lo_export_7))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L34
	}
L31:
	;
	if v172 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	v180 = F_write(m, v93, v170, v172)
	mBase = m.M
	if v180 == v172 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errcode_for_file_access(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(208)
	F_errmsg(m, int32(_a_F_be_lo_export_8), v12+int32(16))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errfinish(m, int32(_a_F_be_lo_export_5), int32(539), int32(_a_F_be_lo_export_6))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	if v221 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_pfree(m, v76)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L46
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errcode_for_file_access(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(208)
	F_errmsg(m, int32(_a_F_be_lo_export_9), v12+int32(32))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0]))) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1]))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2]))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3]))) = v76
	F_errfinish(m, int32(_a_F_be_lo_export_5), int32(546), int32(_a_F_be_lo_export_6))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	goto L5
L47:
	;
	v278 = int32(v274)
	m.G0 = v12
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v12+int32(44) == v284 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	m.ExcPending = 1
	goto L56
L49:
	;
	if v288 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v288 = v286
	goto L52
L51:
	;
	v288 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[3])))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[2])))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[0])))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_be_lo_export[1])))
	v17 = v291
	v18 = v292
	v19 = v290
	v20 = v289
	v21 = v288
	v22 = v280
	goto L2
L54:
	;
	goto L55
L55:
	;
	F___wasm_longjmp(m, v281, v280)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(0)
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_be_lo_open(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v101 int32
	_ = v101
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8&int32(_a_F_be_lo_open_0) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_open_1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_be_lo_open[0])) = uint8(v17)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[1]))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[2]))
	v30 = F_AllocSetContextCreateInternal(m, v25, int32(_a_F_be_lo_open_2), int32(0), int32(_a_F_be_lo_open_3), int32(_a_F_be_lo_open_4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v33 = v20
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[3]))
	if int32(0) < v35 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[1])) = v30
	v33 = v30
	goto L8
L10:
	;
	v88 = F_inv_open(m, v7, v8, v83)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[3])) = v69
	*(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[4])) = v75
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[1]))
	v83 = v81
	v84 = v71
	goto L10
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[4]))
	v41 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v67 = F_MemoryContextAllocZero(m, v33, int32(256))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L22
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41<<(uint(int32(2))%32))))
	if v50 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v62 = F_repalloc0(m, v40, v35<<(uint(int32(2))%32), v35<<(uint(int32(3))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L21
	}
L17:
	;
	v83 = v33
	v84 = v41
	goto L10
L18:
	;
	goto L19
L19:
	;
	v54 = v41 + int32(1)
	if v54 != v35 {
		v41 = v54
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v69 = v35 << (uint(int32(1)) % 32)
	v71 = v35
	v75 = v62
	goto L11
L22:
	;
	v69 = int32(64)
	v71 = int32(0)
	v75 = v67
	goto L11
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[5]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v94 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[6]))
	v97 = F_RegisterSnapshotOnOwner(m, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_open[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v101+v84<<(uint(int32(2))%32)))) = v88
	return v84
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v97
	goto L27
}
func F_be_lowrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_PreventCommandIfReadOnly(m, int32(_a_F_be_lowrite_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v21 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v31 = int32(4)
				} else {
					v31 = v24
				}
				v44 = v31
			} else {
				v32 = int32(1)
				if v15&v32 != 0 {
					v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v45 = int32(1)
			if v15&v45 != 0 {
				v49 = v45
			} else {
				v49 = int32(4)
			}
			v51 = m.G0
			v53 = v51 - int32(32)
			m.G0 = v53
			if v6 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = v6
						F_errmsg(m, int32(_a_F_be_lowrite_1), v53)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_be_lowrite_2), int32(190), int32(_a_F_be_lowrite_3))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_be_lowrite[0]))
				if v58 <= v6 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v53))) = v6
							F_errmsg(m, int32(_a_F_be_lowrite_1), v53)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_be_lowrite_2), int32(190), int32(_a_F_be_lowrite_3))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_be_lowrite[1]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v6<<(uint(int32(2))%32))))
					if v65 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v53))) = v6
								F_errmsg(m, int32(_a_F_be_lowrite_1), v53)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_be_lowrite_2), int32(190), int32(_a_F_be_lowrite_3))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+24)))
						if v68&int32(2) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v6
									F_errmsg(m, int32(_a_F_be_lowrite_4), v53+int32(16))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_be_lowrite_2), int32(198), int32(_a_F_be_lowrite_3))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v73 = F_inv_write(m, v65, v8+v49, v44)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								m.G0 = v53 + int32(32)
								return v73
							}
						}
					}
				}
			}
		}
	}
}
