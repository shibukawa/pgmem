package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WalSndKill(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = int32(4354268)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	*(*int32)(unsafe.Add(mBase, _consts[212])) = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = int32(1)
	if v8 != 0 {
		F_s_lock(m, v4+int32(76), int32(474026), int32(3023), int32(290606))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18
			return
		}
	} else {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18
		return
	}
}
func F_WalSndLastCycleHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[589])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_WalSndPrepareWrite(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v5
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_enlargeStringInfo(m, v14, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v21 = int32(119)
		*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v18 + int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		F_enlargeStringInfo(m, v26, int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v33 = int64(56)
			v35 = int64(65280)
			v37 = int64(40)
			v40 = int64(16711680)
			v42 = int64(24)
			v44 = int64(4278190080)
			v46 = int64(8)
			if l3 != 0 {
				v69 = l1<<(uint(v33)%64) | l1&v35<<(uint(v37)%64) | (l1&v40<<(uint(v42)%64) | l1&v44<<(uint(v46)%64)) | (int64(base.Ui64(l1)>>(uint(v46)%64))&v44 | int64(base.Ui64(l1)>>(uint(v42)%64))&v40 | (int64(base.Ui64(l1)>>(uint(v37)%64))&v35 | int64(base.Ui64(l1)>>(uint(v33)%64))))
			} else {
				v69 = int64(0)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v30+v31))) = v69
			v71 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v30 + v71
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_enlargeStringInfo(m, v74, v71)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				*(*int64)(unsafe.Add(mBase, uint32(v78+v79))) = v69
				v82 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v78 + v82
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				F_enlargeStringInfo(m, v85, v82)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v89 + int32(8)
					return
				}
			}
		}
	}
}
func F_check_wal_consistency_checking(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(288)
	m.G0 = v13
	v20 = F__emscripten_memset_bulkmem(m, v13+int32(16), base.I32_extend8_s(v4), int32(256))
	mBase = m.M
	goto L1
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = F_pstrdup(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v13 + int32(288)
	return v301
L3:
	;
	return int32(0)
L4:
	;
	v29 = F_SplitIdentifierString(m, v22, int32(44), v13+int32(284))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[87])) = v34
	goto L9
L7:
	;
	goto L8
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v40 = F_format_elog_string(m, int32(606998), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v40
	F_pfree(m, v22)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v301 = v4
	goto L2
L13:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[87])) = v278
	goto L75
L14:
	;
	F_pfree(m, v22)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L3
	} else {
		goto L67
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v61 = v4
	goto L17
L17:
	;
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v61<<(uint(int32(2))%32))))
	v74 = v70
	v75 = int32(291673)
	goto L21
L18:
	;
	goto L14
L19:
	;
	v247 = v61 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v247 < v248 {
		v61 = v247
		goto L17
	} else {
		goto L66
	}
L20:
	;
	if v112 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v78 == v79 {
		v101 = v78
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v112 = int32(0)
	goto L20
L23:
	;
	v103 = int32(1)
	if v101 != 0 {
		v74 = v74 + v103
		v75 = v75 + v103
		goto L21
	} else {
		goto L32
	}
L24:
	;
	if base.Ui32((v78-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = v78 | int32(32)
	goto L27
L26:
	;
	v89 = v78
	goto L27
L27:
	;
	if base.Ui32((v79-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v79 | int32(32)
	goto L30
L29:
	;
	v98 = v79
	goto L30
L30:
	;
	if v89 == v98 {
		v101 = v89
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v112 = v89 - v98
	goto L20
L32:
	;
	goto L22
L33:
	;
	v115 = v64
	goto L36
L34:
	;
	goto L35
L35:
	;
	v164 = v64
	goto L45
L36:
	;
	v126 = v115 << (uint(int32(5)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+uint32(_consts[121])))
	if v129 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v141 = v115 | int32(1)
	v143 = v141 << (uint(int32(5)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[121])))
	if v146 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)+uint32(_consts[122])))
	if v132 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v115))) = uint8(v138)
	goto L38
L41:
	;
	v158 = v115 + int32(2)
	if v158 != int32(256) {
		v115 = v158
		goto L36
	} else {
		goto L44
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[122])))
	if v149 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v141))) = uint8(v155)
	goto L41
L44:
	;
	goto L19
L45:
	;
	v172 = v164 << (uint(int32(5)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_consts[121])))
	if v175 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v232 != 0 {
		goto L13
	} else {
		goto L65
	}
L47:
	;
	v228 = v164 + int32(1)
	if v228 != int32(256) {
		v164 = v228
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_consts[122])))
	if v178 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v183 = v70
	v184 = v175
	goto L51
L50:
	;
	if v221 != 0 {
		goto L47
	} else {
		goto L63
	}
L51:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v187 == v188 {
		v210 = v187
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v221 = int32(0)
	goto L50
L53:
	;
	v212 = int32(1)
	if v210 != 0 {
		v183 = v183 + v212
		v184 = v184 + v212
		goto L51
	} else {
		goto L62
	}
L54:
	;
	if base.Ui32((v187-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v198 = v187 | int32(32)
	goto L57
L56:
	;
	v198 = v187
	goto L57
L57:
	;
	if base.Ui32((v188-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v207 = v188 | int32(32)
	goto L60
L59:
	;
	v207 = v188
	goto L60
L60:
	;
	if v198 == v207 {
		v210 = v198
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v221 = v198 - v207
	goto L50
L62:
	;
	goto L52
L63:
	;
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v164))) = uint8(v225)
	goto L19
L64:
	;
	goto L46
L65:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[124])) = uint8(v234)
	goto L19
L66:
	;
	goto L18
L67:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v266 = F_guc_malloc(m, int32(256))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v266
	if v266 == int32(0) {
		v301 = v4
		goto L2
	} else {
		goto L70
	}
L70:
	;
	goto L72
L71:
	;
	v301 = int32(1)
	goto L2
L72:
	;
	v274 = F__emscripten_memcpy_bulkmem(m, v266, v13+int32(16), int32(256))
	mBase = m.M
	goto L74
L74:
	;
	goto L71
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v70
	v284 = F_format_elog_string(m, int32(628253), v13)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v284
	F_pfree(m, v22)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v301 = v4
	goto L2
}
func F_check_wal_segment_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(0)
	v14 = base.B2i32(base.Ui32(v4+int32(-1048576)) < base.Ui32(int32(1072693249))) & base.B2i32(v4&(v4-int32(1)) == v12)
	if v14 == v12 {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[86]))
		*(*int32)(unsafe.Add(mBase, _consts[87])) = v18
		v24 = F_format_elog_string(m, int32(620089), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[88])) = v24
			return v14
		}
	} else {
		return v14
	}
}
func F_wal_segment_open(m *base.Module, l0 int32, l1 int64, l2 int32) {
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
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v12
	v15 = base.I64_div_u_s(int64(4294967296), v11)
	v16 = base.I64_div_u_s(l1, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v16)
	v19 = l1 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v19)
	v27 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(487923), v9+int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v32 = F_BasicOpenFile(m, v9+int32(48), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = v32
			if v32 < int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, _consts[86]))
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v38 == int32(44) {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
							F_errmsg(m, int32(421017), v9)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_errfinish(m, int32(472461), int32(821), int32(268565))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
							F_errmsg(m, int32(285190), v9+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_errfinish(m, int32(472461), int32(826), int32(268565))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
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
				m.G0 = v9 + int32(1072)
				return
			}
		}
	}
}
