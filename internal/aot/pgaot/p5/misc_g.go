package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAccessStrategy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	switch l0 {
	case 0:
		v69 = v2
		m.G0 = v8 + int32(16)
		return v69
	case 1:
		v28 = *(*int32)(unsafe.Add(mBase, _consts[335]))
		v30 = *(*int32)(unsafe.Add(mBase, _consts[632]))
		v32 = int32(3)
		v34 = int32(256)
		v35 = v28*v30<<(uint(v32)%32) + v34
		v38 = *(*int32)(unsafe.Add(mBase, _consts[600]))
		v40 = v38 << (uint(v32) % 32)
		if v40 <= v34 {
			v43 = v34
		} else {
			v43 = v40
		}
		if v35 < v43 {
			v45 = v35
		} else {
			v45 = v43
		}
		if base.Ui32(v45|int32(7)) < base.Ui32(int32(15)) {
			v69 = v2
			m.G0 = v8 + int32(16)
			return v69
		} else {
			v50 = v45
			v53 = *(*int32)(unsafe.Add(mBase, _consts[35]))
			v54 = int32(8)
			v55 = base.I32_div_s(v53, v54)
			v57 = base.I32_div_s(v50, v54)
			if v55 < v57 {
				v59 = v55
			} else {
				v59 = v57
			}
			v64 = F_palloc0(m, v59<<(uint(int32(2))%32)+int32(12))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = l0
				v69 = v64
				m.G0 = v8 + int32(16)
				return v69
			}
		}
	case 2:
		v50 = int32(16384)
		v53 = *(*int32)(unsafe.Add(mBase, _consts[35]))
		v54 = int32(8)
		v55 = base.I32_div_s(v53, v54)
		v57 = base.I32_div_s(v50, v54)
		if v55 < v57 {
			v59 = v55
		} else {
			v59 = v57
		}
		v64 = F_palloc0(m, v59<<(uint(int32(2))%32)+int32(12))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = l0
			v69 = v64
			m.G0 = v8 + int32(16)
			return v69
		}
	case 3:
		v50 = int32(2048)
		v53 = *(*int32)(unsafe.Add(mBase, _consts[35]))
		v54 = int32(8)
		v55 = base.I32_div_s(v53, v54)
		v57 = base.I32_div_s(v50, v54)
		if v55 < v57 {
			v59 = v55
		} else {
			v59 = v57
		}
		v64 = F_palloc0(m, v59<<(uint(int32(2))%32)+int32(12))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = l0
			v69 = v64
			m.G0 = v8 + int32(16)
			return v69
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(502116), v8)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514767), int32(611), int32(20949))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
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
func F_GetConfFilesInDir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(1088)
	m.G0 = v14
	v16 = int32(789774)
	v20 = m.G0
	v22 = v20 - int32(32)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v23
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1271])))
	if v31 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(1088)
	return v308
L2:
	;
	v100 = F_strlen(m, l0)
	mBase = m.M
	if v99 == v100 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v99 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1272])))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = l0
	goto L9
L7:
	;
	goto L8
L8:
	;
	v49 = v16
	v50 = v31
	goto L12
L9:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == v31 {
		v39 = v39 + int32(1)
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v99 = v39 - l0
	goto L2
L11:
	;
	goto L10
L12:
	;
	v57 = v22 + int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(28)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v58 | v59<<(uint(v50)%32)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v63 != 0 {
		v49 = v49 + v59
		v50 = v63
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v66 == int32(0) {
		v91 = l0
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v99 = v91 - l0
	goto L2
L16:
	;
	v70 = l0
	v71 = v66
	goto L17
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(base.Ui32(v71)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v79)>>(uint(v71)%32))&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v91 = v87
	goto L15
L19:
	;
	v91 = v70
	goto L15
L20:
	;
	goto L21
L21:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v87 = v70 + int32(1)
	if v85 != 0 {
		v70 = v87
		v71 = v85
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v103 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v121 = F_AbsoluteConfigLocation(m, l0, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L26
	} else {
		goto L35
	}
L26:
	;
	return int32(0)
L27:
	;
	if v103 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(398655)
	v308 = v6
	goto L1
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(759607), v14)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(517038), int32(89), int32(224220))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	F_pfree(m, v121)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L26
	} else {
		goto L84
	}
L35:
	;
	v123 = F_AllocateDir(m, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	if v123 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L26
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v151 = F_palloc(m, int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L26
	} else {
		goto L48
	}
L40:
	;
	if v128 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L26
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v121
	v147 = F_psprintf(m, int32(722225), v14+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L26
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v121
	F_errmsg(m, int32(310000), v14+int32(32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(517038), int32(101), int32(224220))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v147
	v295 = v6
	goto L34
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v155 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L26
	} else {
		goto L50
	}
L49:
	;
	F_FreeDir(m, v123)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L26
	} else {
		goto L83
	}
L50:
	;
	if v155 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v158 = v155
	v164 = v151
	v167 = int32(32)
	goto L54
L52:
	;
	v264 = v151
	goto L53
L53:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v269 <= int32(0) {
		v282 = v264
		goto L49
	} else {
		goto L81
	}
L54:
	;
	v170 = v158 + int32(19)
	v171 = F_strlen(m, v170)
	mBase = m.M
	if base.Ui32(v171) < base.Ui32(int32(6)) {
		v254 = v164
		v255 = v167
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v264 = v254
	goto L53
L56:
	;
	v256 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L26
	} else {
		goto L79
	}
L57:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 == int32(46) {
		v254 = v164
		v255 = v167
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v179 = v170 + v171 - int32(5)
	v180 = int32(354952)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1273])))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v184 == int32(0) {
		v203 = v183
		v204 = v184
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v204-v203 != 0 {
		v254 = v164
		v255 = v167
		goto L56
	} else {
		goto L67
	}
L60:
	;
	goto L59
L61:
	;
	if v183 != v184 {
		v203 = v183
		v204 = v184
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v188 = v179
	v189 = v180
	goto L63
L63:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v193 == int32(0) {
		v203 = v192
		v204 = v193
		goto L60
	} else {
		goto L65
	}
L64:
	;
	v203 = v192
	v204 = v193
	goto L60
L65:
	;
	v196 = int32(1)
	if v192 == v193 {
		v188 = v188 + v196
		v189 = v189 + v196
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	F_join_path_components(m, v14-int32(-64), v121, v170)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L68
	}
L68:
	;
	v211 = v14 - int32(-64)
	F_canonicalize_path_enc(m, v211)
	mBase = m.M
	v216 = F_get_dirent_type(m, v211, v158, int32(1), l2)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L26
	} else {
		goto L71
	}
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v167 <= v230 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 - int32(-64)
	v224 = F_psprintf(m, int32(748248), v14+int32(48))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L26
	} else {
		goto L72
	}
L71:
	;
	switch v216 {
	case 0:
		goto L70
	default:
		goto L69
	case 3:
		v254 = v164
		v255 = v167
		goto L56
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v224
	F_pfree(m, v164)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L26
	} else {
		goto L73
	}
L73:
	;
	v282 = int32(0)
	goto L49
L74:
	;
	v233 = v167 + int32(32)
	v236 = F_repalloc(m, v164, v233<<(uint(int32(2))%32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L26
	} else {
		goto L77
	}
L75:
	;
	v238 = v164
	v239 = v167
	goto L76
L76:
	;
	v242 = F_pstrdup(m, v14-int32(-64))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L26
	} else {
		goto L78
	}
L77:
	;
	v238 = v236
	v239 = v233
	goto L76
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v238+v244<<(uint(int32(2))%32)))) = v242
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v249 + int32(1)
	v254 = v238
	v255 = v239
	goto L56
L79:
	;
	if v256 != 0 {
		v158 = v256
		v164 = v254
		v167 = v255
		goto L54
	} else {
		goto L80
	}
L80:
	;
	goto L55
L81:
	;
	F_pg_qsort(m, v264, v269, int32(4), int32(1186))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L26
	} else {
		goto L82
	}
L82:
	;
	v282 = v264
	goto L49
L83:
	;
	v295 = v282
	goto L34
L84:
	;
	v308 = v295
	goto L1
}
func F_GetFdwRoutine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[365])))
	if v9&int32(2) == int32(0) {
		v14 = F_OidFunctionCall0Coll(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(115860), v6)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519836), int32(345), int32(390751))
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
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v20 != int32(444) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg_internal(m, int32(115860), v6)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519836), int32(345), int32(390751))
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
				} else {
					m.G0 = v6 + int32(16)
					return v14
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(467213), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519836), int32(337), int32(390751))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
func F_GetFullPageWriteInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, _consts[116]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[132])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
	return
}
func F_GetInsertRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+440)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[114]))
		F_s_lock(m, v9+int32(440), int32(521668), int32(6526), int32(217009))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[114]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+184))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[114]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+184))
		return v23
	}
}
func F_GetLatestSnapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
	if v7 != 0 {
		v9 = int32(1)
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+76)))
		v9 = v8
	}
	if v9&int32(1) == int32(0) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[888])))
		if v15 == int32(0) {
			v18 = F_GetTransactionSnapshot(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v18
			}
		} else {
			v25 = F_GetSnapshotData(m, int32(4554440))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1291])) = v25
				return v25
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(272274), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(518207), int32(361), int32(93183))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
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
func F_GetLocksMethodTable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_consts[697])))
	return v7
}
func F_GetPubPartitionOptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v5 = F_get_rel_relkind(m, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v52 = F_list_concat(m, l0, v15)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	return v47
L5:
	;
	v45 = F_lappend_oid(m, l0, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L20
	}
L6:
	;
	if v5 != int32(112) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v13 = int32(0)
	v15 = F_find_all_inheritors(m, l2, v13, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	switch l1 - int32(1) {
	case 0:
		goto L9
	case 1:
		goto L3
	default:
		v47 = l0
		goto L4
	}
L9:
	;
	if v15 == int32(0) {
		v47 = l0
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v22 <= v21 {
		v47 = l0
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v25 = l0
	v27 = v21
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v27<<(uint(int32(2))%32))))
	v34 = F_get_rel_relkind(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v47 = v40
	goto L4
L14:
	;
	if v34 != int32(112) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v38 = F_lappend_oid(m, v25, v33)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v40 = v25
	goto L17
L17:
	;
	v42 = v27 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v42 < v43 {
		v25 = v40
		v27 = v42
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v40 = v38
	goto L17
L19:
	;
	goto L13
L20:
	;
	v47 = v45
	goto L4
L21:
	;
	return v52
}
func F___getopt_msg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v8 = F_fputs(m, l0, int32(4448160))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if v8 < int32(0) {
			return
		} else {
			v12 = F_strlen(m, l1)
			v15 = F_fwrite(m, l1, v12, int32(1), int32(4448160))
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 == int32(0) {
					return
				} else {
					v21 = F_fwrite(m, l2, int32(1), l3, int32(4448160))
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						if v21 != l3 {
							return
						} else {
							F_fputc(m, int32(10), int32(4448160))
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_g_intbig_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v495 int32
	_ = v495
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v593 int32
	_ = v593
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v664 int32
	_ = v664
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v724 int32
	_ = v724
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v779 int32
	_ = v779
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	v2 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 == v2 {
		v45 = v2
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
		if v32 == int32(0) {
			v45 = v2
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			if v35 != int32(7) {
				v45 = v2
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				if v38 != int32(17) {
					v45 = v2
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
					v45 = v41 ^ int32(1)
				}
			}
		}
	}
	if v45&int32(1) != 0 {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v49 = F_get_fn_opclass_options(m, v48)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
			v54 = v53
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v59 = (v55 + int32(65534)) & int32(65535)
			v63 = v59<<(uint(int32(1))%32) + int32(4)
			v64 = F_palloc(m, v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = v64
				v67 = F_palloc(m, v63)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v67
					if base.Ui32(int32(2)) <= base.Ui32(v59) {
						v73 = v26 + int32(4)
						v82 = int32(-1)
						v84 = v2
						v85 = v2
						v86 = int32(1)
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
							v105 = v86 + int32(1)
							v106 = v105
							v109 = v105
							v112 = v82
							v114 = v84
							v115 = v85
							for {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v109<<(uint(int32(4))%32))))
								v134 = F_hemdist_3(m, v103, v133, v54)
								mBase = m.M
								v135 = base.B2i32(v112 < v134)
								if v112 < v134 {
									v136 = v134
								} else {
									v136 = v112
								}
								if v112 < v134 {
									v137 = v106
								} else {
									v137 = v115
								}
								if v112 < v134 {
									v138 = v86
								} else {
									v138 = v114
								}
								v140 = v106 + int32(1)
								v142 = v140 & int32(65535)
								if base.Ui32(v142) <= base.Ui32(v59) {
									v106 = v140
									v109 = v142
									v112 = v136
									v114 = v138
									v115 = v137
									continue
								} else {
									break
								}
								break
							}
							if v59 != v105 {
								v82 = v136
								v84 = v138
								v85 = v137
								v86 = v105
								continue
							} else {
								break
							}
							break
						}
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v154 = v138
						v155 = v137
						v162 = v145
					} else {
						v154 = v2
						v155 = v2
						v162 = v67
					}
					v170 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v170
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v175 = int32(8)
					v177 = v54 + v175
					v179 = v26 + int32(4)
					v181 = int32(65535)
					v189 = base.B2i32(v154&v181 == v170) | base.B2i32(v155&v181 == v170)
					if v189 != 0 {
						v190 = int32(1)
					} else {
						v190 = v154
					}
					v193 = int32(4)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(v179+v190&int32(65535)<<(uint(v193)%32))))
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
					v199 = v197 & v193
					if v199 != 0 {
						v200 = v175
					} else {
						v200 = v177
					}
					v201 = F_palloc(m, v200)
					mBase = m.M
					v202 = m.ExcPending
					if v202 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v199
						*(*int32)(unsafe.Add(mBase, uint32(v201))) = v200 << (uint(int32(2)) % 32)
						if v199 == int32(0) {
							v209 = int32(8)
							if v54 != 0 {
								v213 = F__emscripten_memcpy_bulkmem(m, v201+v209, v196+v209, v54)
								mBase = m.M
							} else {
							}
						} else {
						}
						if v189 != 0 {
							v217 = int32(2)
						} else {
							v217 = v155
						}
						v220 = int32(4)
						v223 = *(*int32)(unsafe.Add(mBase, uint32(v179+v217&int32(65535)<<(uint(v220)%32))))
						v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
						v226 = v224 & v220
						if v226 != 0 {
							v227 = int32(8)
						} else {
							v227 = v177
						}
						v228 = F_palloc(m, v227)
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v226
							*(*int32)(unsafe.Add(mBase, uint32(v228))) = v227 << (uint(int32(2)) % 32)
							if v226 == int32(0) {
								v236 = int32(8)
								if v54 != 0 {
									v240 = F__emscripten_memcpy_bulkmem(m, v228+v236, v223+v236, v54)
									mBase = m.M
								} else {
								}
							} else {
							}
							v242 = int32(65535)
							v243 = v55 + v242
							v245 = v243 & v242
							v248 = F_palloc(m, v245<<(uint(int32(3))%32))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								if v55&int32(65535) == int32(1) {
									F_pg_qsort(m, v248, v245, int32(8), int32(7625))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										v793 = v174
										v797 = v162
										v805 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v793))) = uint16(v805)
										*(*uint16)(unsafe.Add(mBase, uint32(v797))) = uint16(v805)
										F_pfree(m, v248)
										mBase = m.M
										v810 = m.ExcPending
										if v810 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v228
											*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v201
											return v25
										}
									}
								} else {
									v258 = int32(1)
									v260 = v258
									v266 = v258
									for {
										v286 = v248 + v260<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v286-int32(8)))) = uint16(v266)
										v290 = int32(4)
										v295 = *(*int32)(unsafe.Add(mBase, uint32(v179+v260<<(uint(v290)%32))))
										v296 = F_hemdist_3(m, v201, v295, v54)
										mBase = m.M
										v297 = F_hemdist_3(m, v228, v295, v54)
										mBase = m.M
										v298 = v296 - v297
										v300 = v298 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v286-v290))) = v298 ^ v300 - v300
										v305 = v266 + int32(1)
										v306 = int32(65535)
										v307 = v305 & v306
										if base.Ui32(v307) <= base.Ui32(v243&v306) {
											v260 = v307
											v266 = v305
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v248, v245, int32(8), int32(7625))
									mBase = m.M
									v314 = m.ExcPending
									if v314 != 0 {
										return int32(0)
									} else {
										v315 = int32(1)
										if base.Ui32(v245) <= base.Ui32(v315) {
											v318 = v315
										} else {
											v318 = v245
										}
										v320 = v54 & int32(2147483644)
										v322 = v54 & int32(3)
										v323 = int32(8)
										v324 = v228 + v323
										v326 = v201 + v323
										v336 = int32(0)
										v342 = v174
										v346 = v162
										for {
											v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248+v336<<(uint(int32(3))%32)))))
											if v190&int32(65535) == v357 {
												*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v190)
												v360 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v360 + int32(1)
												v766 = v342 + int32(2)
												v770 = v346
											} else {
												if v217&int32(65535) == v357 {
													*(*uint16)(unsafe.Add(mBase, uint32(v346))) = uint16(v217)
													v370 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v370 + int32(1)
													v766 = v342
													v770 = v346 + int32(2)
												} else {
													v377 = *(*int32)(unsafe.Add(mBase, uint32(v179+v357<<(uint(int32(4))%32))))
													v378 = F_hemdist_3(m, v201, v377, v54)
													mBase = m.M
													v380 = F_hemdist_3(m, v228, v377, v54)
													mBase = m.M
													v382 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													v383 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v384 = v382 - v383
													if base.F64_lt(base.F64_convert_i32_s(v378), base.F64_add(base.F64_convert_i32_s(v380), base.F64_mul(base.F64_convert_i32_s(v384*v384*v384), float64(-1e-05)))) != 0 {
														v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+4)))
														if v392&int32(4) != 0 {
														} else {
															v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
															if v395&int32(4) != 0 {
																v400 = F__emscripten_memset_bulkmem(m, v326, base.I32_extend8_s(int32(255)), v54)
																mBase = m.M
															} else {
																if v54 <= int32(0) {
																} else {
																	v404 = v377 + int32(8)
																	v405 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v410 = v405
																		v424 = v405
																		for {
																			v434 = v410 + v326
																			v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
																			v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v404))))
																			v438 = v435 | v437
																			*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v438)
																			v441 = v410 | int32(1)
																			v442 = v326 + v441
																			v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
																			v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v441))))
																			v446 = v443 | v445
																			*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v446)
																			v449 = v410 | int32(2)
																			v450 = v326 + v449
																			v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
																			v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v449))))
																			v454 = v451 | v453
																			*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v454)
																			v457 = v410 | int32(3)
																			v458 = v326 + v457
																			v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
																			v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v457))))
																			v462 = v459 | v461
																			*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v462)
																			v464 = int32(4)
																			v465 = v410 + v464
																			v467 = v424 + v464
																			if v467 != v320 {
																				v410 = v465
																				v424 = v467
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v469 = v465
																	} else {
																		v469 = v405
																	}
																	if v322 == int32(0) {
																	} else {
																		v495 = v469
																		v512 = v405
																		for {
																			v519 = v495 + v326
																			v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
																			v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v404))))
																			v523 = v520 | v522
																			*(*uint8)(unsafe.Add(mBase, uint32(v519))) = uint8(v523)
																			v525 = int32(1)
																			v528 = v512 + v525
																			if v528 != v322 {
																				v495 = v495 + v525
																				v512 = v528
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v357)
														v555 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v555 + int32(1)
														v766 = v342 + int32(2)
														v770 = v346
													} else {
														v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
														if v561&int32(4) != 0 {
														} else {
															v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
															if v564&int32(4) != 0 {
																v569 = F__emscripten_memset_bulkmem(m, v324, base.I32_extend8_s(int32(255)), v54)
																mBase = m.M
															} else {
																if v54 <= int32(0) {
																} else {
																	v573 = v377 + int32(8)
																	v574 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v579 = v574
																		v593 = v574
																		for {
																			v603 = v579 + v324
																			v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
																			v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579+v573))))
																			v607 = v604 | v606
																			*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v607)
																			v610 = v579 | int32(1)
																			v611 = v324 + v610
																			v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
																			v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v610))))
																			v615 = v612 | v614
																			*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v615)
																			v618 = v579 | int32(2)
																			v619 = v324 + v618
																			v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
																			v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v618))))
																			v623 = v620 | v622
																			*(*uint8)(unsafe.Add(mBase, uint32(v619))) = uint8(v623)
																			v626 = v579 | int32(3)
																			v627 = v324 + v626
																			v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
																			v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v626))))
																			v631 = v628 | v630
																			*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v631)
																			v633 = int32(4)
																			v634 = v579 + v633
																			v636 = v593 + v633
																			if v636 != v320 {
																				v579 = v634
																				v593 = v636
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v638 = v634
																	} else {
																		v638 = v574
																	}
																	if v322 == int32(0) {
																	} else {
																		v664 = v638
																		v681 = v574
																		for {
																			v688 = v664 + v324
																			v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
																			v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664+v573))))
																			v692 = v689 | v691
																			*(*uint8)(unsafe.Add(mBase, uint32(v688))) = uint8(v692)
																			v694 = int32(1)
																			v697 = v681 + v694
																			if v697 != v322 {
																				v664 = v664 + v694
																				v681 = v697
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v346))) = uint16(v357)
														v724 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v724 + int32(1)
														v766 = v342
														v770 = v346 + int32(2)
													}
												}
											}
											v779 = v336 + int32(1)
											if v779 != v318 {
												v336 = v779
												v342 = v766
												v346 = v770
												continue
											} else {
												break
											}
											break
										}
										v793 = v766
										v797 = v770
										v805 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v793))) = uint16(v805)
										*(*uint16)(unsafe.Add(mBase, uint32(v797))) = uint16(v805)
										F_pfree(m, v248)
										mBase = m.M
										v810 = m.ExcPending
										if v810 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v228
											*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v201
											return v25
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v54 = int32(252)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v59 = (v55 + int32(65534)) & int32(65535)
		v63 = v59<<(uint(int32(1))%32) + int32(4)
		v64 = F_palloc(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v64
			v67 = F_palloc(m, v63)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v67
				if base.Ui32(int32(2)) <= base.Ui32(v59) {
					v73 = v26 + int32(4)
					v82 = int32(-1)
					v84 = v2
					v85 = v2
					v86 = int32(1)
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
						v105 = v86 + int32(1)
						v106 = v105
						v109 = v105
						v112 = v82
						v114 = v84
						v115 = v85
						for {
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v109<<(uint(int32(4))%32))))
							v134 = F_hemdist_3(m, v103, v133, v54)
							mBase = m.M
							v135 = base.B2i32(v112 < v134)
							if v112 < v134 {
								v136 = v134
							} else {
								v136 = v112
							}
							if v112 < v134 {
								v137 = v106
							} else {
								v137 = v115
							}
							if v112 < v134 {
								v138 = v86
							} else {
								v138 = v114
							}
							v140 = v106 + int32(1)
							v142 = v140 & int32(65535)
							if base.Ui32(v142) <= base.Ui32(v59) {
								v106 = v140
								v109 = v142
								v112 = v136
								v114 = v138
								v115 = v137
								continue
							} else {
								break
							}
							break
						}
						if v59 != v105 {
							v82 = v136
							v84 = v138
							v85 = v137
							v86 = v105
							continue
						} else {
							break
						}
						break
					}
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v154 = v138
					v155 = v137
					v162 = v145
				} else {
					v154 = v2
					v155 = v2
					v162 = v67
				}
				v170 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v170
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v170
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v175 = int32(8)
				v177 = v54 + v175
				v179 = v26 + int32(4)
				v181 = int32(65535)
				v189 = base.B2i32(v154&v181 == v170) | base.B2i32(v155&v181 == v170)
				if v189 != 0 {
					v190 = int32(1)
				} else {
					v190 = v154
				}
				v193 = int32(4)
				v196 = *(*int32)(unsafe.Add(mBase, uint32(v179+v190&int32(65535)<<(uint(v193)%32))))
				v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
				v199 = v197 & v193
				if v199 != 0 {
					v200 = v175
				} else {
					v200 = v177
				}
				v201 = F_palloc(m, v200)
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v199
					*(*int32)(unsafe.Add(mBase, uint32(v201))) = v200 << (uint(int32(2)) % 32)
					if v199 == int32(0) {
						v209 = int32(8)
						if v54 != 0 {
							v213 = F__emscripten_memcpy_bulkmem(m, v201+v209, v196+v209, v54)
							mBase = m.M
						} else {
						}
					} else {
					}
					if v189 != 0 {
						v217 = int32(2)
					} else {
						v217 = v155
					}
					v220 = int32(4)
					v223 = *(*int32)(unsafe.Add(mBase, uint32(v179+v217&int32(65535)<<(uint(v220)%32))))
					v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
					v226 = v224 & v220
					if v226 != 0 {
						v227 = int32(8)
					} else {
						v227 = v177
					}
					v228 = F_palloc(m, v227)
					mBase = m.M
					v229 = m.ExcPending
					if v229 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v226
						*(*int32)(unsafe.Add(mBase, uint32(v228))) = v227 << (uint(int32(2)) % 32)
						if v226 == int32(0) {
							v236 = int32(8)
							if v54 != 0 {
								v240 = F__emscripten_memcpy_bulkmem(m, v228+v236, v223+v236, v54)
								mBase = m.M
							} else {
							}
						} else {
						}
						v242 = int32(65535)
						v243 = v55 + v242
						v245 = v243 & v242
						v248 = F_palloc(m, v245<<(uint(int32(3))%32))
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return int32(0)
						} else {
							if v55&int32(65535) == int32(1) {
								F_pg_qsort(m, v248, v245, int32(8), int32(7625))
								mBase = m.M
								v257 = m.ExcPending
								if v257 != 0 {
									return int32(0)
								} else {
									v793 = v174
									v797 = v162
									v805 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v793))) = uint16(v805)
									*(*uint16)(unsafe.Add(mBase, uint32(v797))) = uint16(v805)
									F_pfree(m, v248)
									mBase = m.M
									v810 = m.ExcPending
									if v810 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v228
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v201
										return v25
									}
								}
							} else {
								v258 = int32(1)
								v260 = v258
								v266 = v258
								for {
									v286 = v248 + v260<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v286-int32(8)))) = uint16(v266)
									v290 = int32(4)
									v295 = *(*int32)(unsafe.Add(mBase, uint32(v179+v260<<(uint(v290)%32))))
									v296 = F_hemdist_3(m, v201, v295, v54)
									mBase = m.M
									v297 = F_hemdist_3(m, v228, v295, v54)
									mBase = m.M
									v298 = v296 - v297
									v300 = v298 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v286-v290))) = v298 ^ v300 - v300
									v305 = v266 + int32(1)
									v306 = int32(65535)
									v307 = v305 & v306
									if base.Ui32(v307) <= base.Ui32(v243&v306) {
										v260 = v307
										v266 = v305
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v248, v245, int32(8), int32(7625))
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return int32(0)
								} else {
									v315 = int32(1)
									if base.Ui32(v245) <= base.Ui32(v315) {
										v318 = v315
									} else {
										v318 = v245
									}
									v320 = v54 & int32(2147483644)
									v322 = v54 & int32(3)
									v323 = int32(8)
									v324 = v228 + v323
									v326 = v201 + v323
									v336 = int32(0)
									v342 = v174
									v346 = v162
									for {
										v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248+v336<<(uint(int32(3))%32)))))
										if v190&int32(65535) == v357 {
											*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v190)
											v360 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v360 + int32(1)
											v766 = v342 + int32(2)
											v770 = v346
										} else {
											if v217&int32(65535) == v357 {
												*(*uint16)(unsafe.Add(mBase, uint32(v346))) = uint16(v217)
												v370 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v370 + int32(1)
												v766 = v342
												v770 = v346 + int32(2)
											} else {
												v377 = *(*int32)(unsafe.Add(mBase, uint32(v179+v357<<(uint(int32(4))%32))))
												v378 = F_hemdist_3(m, v201, v377, v54)
												mBase = m.M
												v380 = F_hemdist_3(m, v228, v377, v54)
												mBase = m.M
												v382 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v383 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												v384 = v382 - v383
												if base.F64_lt(base.F64_convert_i32_s(v378), base.F64_add(base.F64_convert_i32_s(v380), base.F64_mul(base.F64_convert_i32_s(v384*v384*v384), float64(-1e-05)))) != 0 {
													v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+4)))
													if v392&int32(4) != 0 {
													} else {
														v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
														if v395&int32(4) != 0 {
															v400 = F__emscripten_memset_bulkmem(m, v326, base.I32_extend8_s(int32(255)), v54)
															mBase = m.M
														} else {
															if v54 <= int32(0) {
															} else {
																v404 = v377 + int32(8)
																v405 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v410 = v405
																	v424 = v405
																	for {
																		v434 = v410 + v326
																		v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
																		v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v404))))
																		v438 = v435 | v437
																		*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v438)
																		v441 = v410 | int32(1)
																		v442 = v326 + v441
																		v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
																		v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v441))))
																		v446 = v443 | v445
																		*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v446)
																		v449 = v410 | int32(2)
																		v450 = v326 + v449
																		v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
																		v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v449))))
																		v454 = v451 | v453
																		*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v454)
																		v457 = v410 | int32(3)
																		v458 = v326 + v457
																		v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
																		v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v457))))
																		v462 = v459 | v461
																		*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v462)
																		v464 = int32(4)
																		v465 = v410 + v464
																		v467 = v424 + v464
																		if v467 != v320 {
																			v410 = v465
																			v424 = v467
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v469 = v465
																} else {
																	v469 = v405
																}
																if v322 == int32(0) {
																} else {
																	v495 = v469
																	v512 = v405
																	for {
																		v519 = v495 + v326
																		v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
																		v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v404))))
																		v523 = v520 | v522
																		*(*uint8)(unsafe.Add(mBase, uint32(v519))) = uint8(v523)
																		v525 = int32(1)
																		v528 = v512 + v525
																		if v528 != v322 {
																			v495 = v495 + v525
																			v512 = v528
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v357)
													v555 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v555 + int32(1)
													v766 = v342 + int32(2)
													v770 = v346
												} else {
													v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
													if v561&int32(4) != 0 {
													} else {
														v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
														if v564&int32(4) != 0 {
															v569 = F__emscripten_memset_bulkmem(m, v324, base.I32_extend8_s(int32(255)), v54)
															mBase = m.M
														} else {
															if v54 <= int32(0) {
															} else {
																v573 = v377 + int32(8)
																v574 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v579 = v574
																	v593 = v574
																	for {
																		v603 = v579 + v324
																		v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
																		v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579+v573))))
																		v607 = v604 | v606
																		*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v607)
																		v610 = v579 | int32(1)
																		v611 = v324 + v610
																		v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
																		v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v610))))
																		v615 = v612 | v614
																		*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v615)
																		v618 = v579 | int32(2)
																		v619 = v324 + v618
																		v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
																		v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v618))))
																		v623 = v620 | v622
																		*(*uint8)(unsafe.Add(mBase, uint32(v619))) = uint8(v623)
																		v626 = v579 | int32(3)
																		v627 = v324 + v626
																		v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
																		v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v626))))
																		v631 = v628 | v630
																		*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v631)
																		v633 = int32(4)
																		v634 = v579 + v633
																		v636 = v593 + v633
																		if v636 != v320 {
																			v579 = v634
																			v593 = v636
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v638 = v634
																} else {
																	v638 = v574
																}
																if v322 == int32(0) {
																} else {
																	v664 = v638
																	v681 = v574
																	for {
																		v688 = v664 + v324
																		v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
																		v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664+v573))))
																		v692 = v689 | v691
																		*(*uint8)(unsafe.Add(mBase, uint32(v688))) = uint8(v692)
																		v694 = int32(1)
																		v697 = v681 + v694
																		if v697 != v322 {
																			v664 = v664 + v694
																			v681 = v697
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v346))) = uint16(v357)
													v724 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v724 + int32(1)
													v766 = v342
													v770 = v346 + int32(2)
												}
											}
										}
										v779 = v336 + int32(1)
										if v779 != v318 {
											v336 = v779
											v342 = v766
											v346 = v770
											continue
										} else {
											break
										}
										break
									}
									v793 = v766
									v797 = v770
									v805 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v793))) = uint16(v805)
									*(*uint16)(unsafe.Add(mBase, uint32(v797))) = uint16(v805)
									F_pfree(m, v248)
									mBase = m.M
									v810 = m.ExcPending
									if v810 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v228
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v201
										return v25
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_g_intbig_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == v2 {
		v38 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v38&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v25 == int32(0) {
		v38 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v28 != int32(7) {
		v38 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v31 != int32(17) {
		v38 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
	v38 = v34 ^ int32(1)
	goto L2
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = F_get_fn_opclass_options(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v47 = int32(252)
	goto L9
L9:
	;
	v49 = v47 + int32(8)
	v50 = F_palloc(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v47 = v46
	goto L9
L12:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
	v55 = v49 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v55
	v61 = F__emscripten_memset_bulkmem(m, v50+int32(8), base.I32_extend8_s(v52), v47)
	mBase = m.M
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v62 <= int32(0) {
		v256 = v55
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v256) >> (uint(int32(2)) % 32))
	return v50
L15:
	;
	v68 = v47 & int32(3)
	v73 = v62
	v82 = v2
	goto L16
L16:
	;
	v90 = int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v82<<(uint(v90)%32))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
	if v94&v90 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(17179869216)
	v256 = int32(32)
	goto L14
L18:
	;
	if base.B2i32(v47 <= int32(0)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v102 = v93 + int32(8)
	v103 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v223 = v73
	goto L23
L23:
	;
	v241 = v82 + int32(1)
	if v241 < v223 {
		v73 = v223
		v82 = v241
		goto L16
	} else {
		goto L36
	}
L24:
	;
	v108 = v103
	v115 = v103
	goto L27
L25:
	;
	v160 = v103
	goto L26
L26:
	;
	if v68 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v125 = v108 + v61
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v102))))
	v129 = v126 | v128
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v129)
	v132 = v108 | int32(1)
	v133 = v61 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v132))))
	v137 = v134 | v136
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v137)
	v140 = v108 | int32(2)
	v141 = v61 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v140))))
	v145 = v142 | v144
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v145)
	v148 = v108 | int32(3)
	v149 = v61 + v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v148))))
	v153 = v150 | v152
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v153)
	v155 = int32(4)
	v156 = v108 + v155
	v158 = v115 + v155
	if v158 != v47&int32(2147483644) {
		v108 = v156
		v115 = v158
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v160 = v156
	goto L26
L29:
	;
	goto L28
L30:
	;
	v177 = v160
	v189 = v103
	goto L33
L31:
	;
	goto L32
L32:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v223 = v222
	goto L23
L33:
	;
	v194 = v177 + v61
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v102))))
	v198 = v195 | v197
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v198)
	v200 = int32(1)
	v203 = v189 + v200
	if v203 != v68 {
		v177 = v177 + v200
		v189 = v203
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	goto L34
L36:
	;
	v256 = v55
	goto L14
}
func F_gb18030_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(39), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4428196), v18, v18, int32(7350), int32(39), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_generateHeadline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(128)
	v14 = F_palloc(m, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int32(4)
	v19 = v14 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if (v11-v21)>>(uint(v18)%32) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v19
	v29 = v11
	v30 = v14
	v31 = v2
	v32 = v12
	v33 = v2
	goto L6
L4:
	;
	v160 = v19
	v163 = v14
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = (v160 - v163) << (uint(int32(2)) % 32)
	return v163
L6:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v39 = v27 - v30
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 <= v36+(v37+(v38+(v39+int32(base.Ui32(v40)>>(uint(int32(16))%32))))) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v160 = v147
	v163 = v78
	goto L5
L8:
	;
	v52 = v30
	v54 = v32
	goto L11
L9:
	;
	v75 = v27
	v76 = v40
	v78 = v30
	v80 = v32
	v82 = v36
	goto L10
L10:
	;
	if v76&int32(10) == int32(2) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v59 = v54 << (uint(int32(1)) % 32)
	v60 = F_repalloc(m, v52, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v75 = v60 + v39
	v76 = v65
	v78 = v60
	v80 = v59
	v82 = v62
	goto L10
L13:
	;
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v59 <= v62+(v63+(v64+(v39+int32(base.Ui32(v65)>>(uint(int32(16))%32))))) {
		v52 = v60
		v54 = v59
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v153 = v29 + int32(16)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if (v153-v154)>>(uint(int32(4))%32) < v151 {
		v27 = v147
		v29 = v153
		v30 = v78
		v31 = v149
		v32 = v80
		v33 = v150
		goto L6
	} else {
		goto L48
	}
L16:
	;
	if v31 != 0 {
		v98 = v75
		v99 = v76
		v100 = v33
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v76&int32(8) != 0 {
		v147 = v75
		v149 = v31
		v150 = v33
		goto L15
	} else {
		goto L46
	}
L19:
	;
	if v99&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v89 = v33 + int32(1)
	if v89 < int32(2) {
		v98 = v75
		v99 = v76
		v100 = v89
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v82 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v98 = v94 + v95
	v99 = v97
	v100 = v89
	goto L19
L23:
	;
	v93 = F__emscripten_memcpy_bulkmem(m, v75, v92, v82)
	mBase = m.M
	v94 = v93
	goto L25
L24:
	;
	v94 = v75
	goto L25
L25:
	;
	goto L22
L26:
	;
	v103 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v103)
	v105 = int32(1)
	v147 = v98 + v105
	v149 = v105
	v150 = v100
	goto L15
L27:
	;
	goto L28
L28:
	;
	v108 = int32(1)
	if v99&int32(16) != 0 {
		v147 = v98
		v149 = v108
		v150 = v100
		goto L15
	} else {
		goto L29
	}
L29:
	;
	if v99&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v120 = v98
	v121 = v99
	goto L32
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v124 = int32(base.Ui32(v121) >> (uint(int32(16)) % 32))
	if v124 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v120 = v116 + v117
	v121 = v119
	goto L32
L34:
	;
	v115 = F__emscripten_memcpy_bulkmem(m, v98, v113, v114)
	mBase = m.M
	v116 = v115
	goto L36
L35:
	;
	v116 = v98
	goto L36
L36:
	;
	goto L33
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v130 = v126 + int32(base.Ui32(v127)>>(uint(int32(16))%32))
	if v127&int32(1) == int32(0) {
		v147 = v130
		v149 = v108
		v150 = v100
		goto L15
	} else {
		goto L41
	}
L38:
	;
	v125 = F__emscripten_memcpy_bulkmem(m, v120, v122, v124)
	mBase = m.M
	v126 = v125
	goto L40
L39:
	;
	v126 = v120
	goto L40
L40:
	;
	goto L37
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	if v136 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v147 = v138 + v139
	v149 = v108
	v150 = v100
	goto L15
L43:
	;
	v137 = F__emscripten_memcpy_bulkmem(m, v130, v135, v136)
	mBase = m.M
	v138 = v137
	goto L45
L44:
	;
	v138 = v130
	goto L45
L45:
	;
	goto L42
L46:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v147 = v75
	v149 = int32(0)
	v150 = v33
	goto L15
L48:
	;
	goto L7
}
func F_generate_append_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = v13 << (uint(int32(2)) % 32)
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = F_palloc(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v133 = int32(0)
	v140 = v133
	v141 = int32(1)
	v143 = v133
	goto L34
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v24 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = int32(0)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v42 = v39 + v36<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = v45
	goto L13
L12:
	;
	v46 = int32(0)
	goto L13
L13:
	;
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v118 = v36 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v118 < v119 {
		v36 = v118
		goto L9
	} else {
		goto L33
	}
L15:
	;
	v49 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v50 <= v49 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = v46
	v59 = v49
	goto L17
L17:
	;
	v66 = v59 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = F_exprType(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L14
L19:
	;
	v92 = v57 + int32(4)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v92) < base.Ui32(v94+v95<<(uint(int32(2))%32)) {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v71 == v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = v66 + v18
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v77 = F_exprTypmod(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+v18))) = int32(-1)
	goto L19
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v79 == v42 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77
	goto L19
L26:
	;
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v77 == v82 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(-1)
	goto L19
L29:
	;
	v100 = v92
	goto L31
L30:
	;
	v100 = int32(0)
	goto L31
L31:
	;
	v102 = v59 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v102 < v103 {
		v57 = v100
		v59 = v102
		goto L17
	} else {
		goto L32
	}
L32:
	;
	goto L18
L33:
	;
	goto L10
L34:
	;
	v148 = int32(0)
	if l0 == v148 {
		v158 = v148
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v159 = int32(0)
	if l1 == v159 {
		v170 = v159
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v152 <= v140 {
		v158 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = v154 + v140<<(uint(int32(2))%32)
	goto L36
L39:
	;
	if l3 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v164 <= v140 {
		v170 = int32(0)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v170 = v166 + v140<<(uint(int32(2))%32)
	goto L39
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v189 = int32(0)
	v190 = base.I32_extend16_s(v141)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v178+v18)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v196 = F_makeVar(m, v189, v190, v191, v193, v194, v189)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L52
	}
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v171 <= v140 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v183 = v159
	goto L45
L45:
	;
	F_pfree(m, v18)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L51
	}
L46:
	;
	v183 = v143
	goto L45
L47:
	;
	if v158 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	if v170 == int32(0) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v178 = v140 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v180 = v178 + v179
	if v180 != 0 {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	return v183
L52:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v199 = F_pstrdup(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v202 = F_makeTargetEntry(m, v196, v190, v199, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v202)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v204
	v206 = int32(1)
	v210 = F_lappend(m, v143, v202)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v140 = v140 + v206
	v141 = v141 + v206
	v143 = v210
	goto L34
}
func F_generate_subscripts_nodir(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_subscripts(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_trgm_only(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
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
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	v14 = l2 + int32(1)
	if base.Ui32(v14) < base.Ui32(int32(357913942)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_palloc(m, v14*int32(3)+int32(5))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L52
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	if l3 == v25 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if int32(0) < l2 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v34 = v25
	v35 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v31 = F_palloc0(m, v14)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v31
	v34 = v31
	v35 = v14
	goto L6
L11:
	;
	v39 = l2 + int32(4)
	v40 = F_palloc(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	return
L14:
	;
	v42 = int32(8224)
	*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v42)
	v48 = v34
	v50 = l1
	v51 = v35
	v52 = v40
	v55 = v39
	goto L15
L15:
	;
	v56 = l1 + l2
	v63 = v50
	goto L18
L16:
	;
	F_pfree(m, v52)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L51
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui32(v56) <= base.Ui32(v63) {
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v83 = v63
	goto L24
L20:
	;
	v70 = F_pg_mblen_range(m, v63, v56)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v73 = F_t_isalnum_with_len(m, v63, v70)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v73 == int32(0) {
		v63 = v63 + v70
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v89 = F_pg_mblen_range(m, v83, v56)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	if v63 == int32(0) {
		goto L17
	} else {
		goto L32
	}
L26:
	;
	v91 = F_t_isalnum_with_len(m, v83, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v93 = v83 + v89
	if base.Ui32(v93) < base.Ui32(v56) {
		v83 = v93
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v95 = v83
	goto L30
L30:
	;
	goto L25
L31:
	;
	v95 = v93
	goto L30
L32:
	;
	v100 = F_str_tolower(m, v63, v95-v63, int32(100))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v102 = F_strlen(m, v100)
	mBase = m.M
	if base.Ui32(v55-int32(4)) < base.Ui32(v102) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v52)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v114 = v52
	v115 = v55
	goto L36
L36:
	;
	if v102 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v109 = v102 + int32(4)
	v110 = F_palloc(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v112 = int32(8224)
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v112)
	v114 = v110
	v115 = v109
	goto L36
L39:
	;
	F_pfree(m, v100)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L43
	}
L40:
	;
	v118 = F__emscripten_memcpy_bulkmem(m, v114+int32(2), v100, v102)
	mBase = m.M
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	v123 = int32(8224)
	*(*uint16)(unsafe.Add(mBase, uint32(v102+v114)+2)) = uint16(v123)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_make_trigrams(m, l0, v114, v102+int32(3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v130 = int32(0)
	if v48 == v130 {
		v48 = v130
		v50 = v95
		v52 = v114
		v55 = v115
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v133 <= v51 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v142 = v140 + v125
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v144 = int32(1)
	v145 = v143 | v144
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v145)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = v140 + v147 - v144
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v153 = v151 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	v48 = v140
	v50 = v95
	v51 = v141
	v52 = v114
	v55 = v115
	goto L15
L47:
	;
	v140 = v48
	v141 = v51
	goto L46
L48:
	;
	goto L49
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v136 = F_repalloc0(m, v48, v51, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v140 = v136
	v141 = v139
	goto L46
L51:
	;
	goto L13
L52:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(14020), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(518898), int32(115), int32(24857))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generate_wildcard_trgm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l1 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L81
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v339
L3:
	;
	v22 = F_palloc(m, int32(5))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if base.Ui32(int32(357913941)) <= base.Ui32(l1) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(20)
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v28)
	v339 = v22
	goto L2
L8:
	;
	v36 = F_palloc(m, l1*int32(3)+int32(8))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l1 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	v44 = l0 + l1
	v47 = F_palloc(m, l1+int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v53 = l0
	v62 = l1
	goto L11
L11:
	;
	v65 = int32(0)
	v69 = v53
	v70 = v65
	v73 = v65
	goto L13
L12:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_pfree(m, v47)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L69
	}
L13:
	;
	v81 = F_pg_mblen_range(m, v69, v44)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	if v62 <= v114-v53 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	goto L14
L16:
	;
	if v70&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v112 = v69 + v81
	if base.Ui32(v112) < base.Ui32(v44) {
		v69 = v112
		v70 = v109
		v73 = v110
		goto L13
	} else {
		goto L32
	}
L18:
	;
	v114 = v69
	v115 = v106
	v116 = v73
	goto L15
L19:
	;
	v85 = F_t_isalnum_with_len(m, v69, v81)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v93 - int32(92) {
	case 0:
		v109 = int32(1)
		v110 = v73
		goto L17
	case 1, 2:
		goto L26
	case 3:
		goto L27
	default:
		goto L28
	}
L22:
	;
	if v85 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v89 = int32(0)
	v109 = v89
	v110 = v89
	goto L17
L24:
	;
	goto L25
L25:
	;
	v106 = int32(1)
	goto L18
L26:
	;
	v100 = int32(0)
	v102 = F_t_isalnum_with_len(m, v69, v81)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L30
	}
L27:
	;
	v109 = int32(0)
	v110 = int32(1)
	goto L17
L28:
	;
	if v93 != int32(37) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v102 == int32(0) {
		v109 = v100
		v110 = v100
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v106 = v100
	goto L18
L32:
	;
	v114 = v112
	v115 = v109
	v116 = v110
	goto L15
L33:
	;
	goto L12
L34:
	;
	if v116&int32(1) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v124 = int32(8224)
	*(*uint16)(unsafe.Add(mBase, uint32(v47))) = uint16(v124)
	v126 = v47 + int32(2)
	goto L37
L36:
	;
	v126 = v47
	goto L37
L37:
	;
	if base.Ui32(v114) < base.Ui32(v44) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v228 = F_str_tolower(m, v47, v214-v47, int32(100))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L65
	}
L39:
	;
	v212 = v196
	v214 = v198 + int32(1)
	goto L38
L40:
	;
	v192 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v192)
	v196 = v189
	v198 = v132
	goto L39
L41:
	;
	v130 = v114
	v131 = v115
	v132 = v126
	goto L44
L42:
	;
	v173 = v114
	v175 = v126
	goto L43
L43:
	;
	v185 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v185)
	if v173 == int32(0) {
		goto L33
	} else {
		goto L64
	}
L44:
	;
	v142 = F_pg_mblen_range(m, v130, v44)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	v173 = v169
	v175 = v167
	goto L43
L46:
	;
	if v131&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v169 = v130 + v142
	if base.Ui32(v169) < base.Ui32(v44) {
		v130 = v169
		v131 = v166
		v132 = v167
		goto L44
	} else {
		goto L63
	}
L48:
	;
	if v142 != 0 {
		goto L60
	} else {
		goto L61
	}
L49:
	;
	v146 = F_t_isalnum_with_len(m, v130, v142)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	switch v151 - int32(92) {
	case 0:
		v166 = int32(1)
		v167 = v132
		goto L47
	case 1, 2:
		goto L54
	case 3:
		v212 = v130
		v214 = v132
		goto L38
	default:
		goto L55
	}
L52:
	;
	if v146 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v189 = v130 - int32(1)
	goto L40
L54:
	;
	v156 = F_t_isalnum_with_len(m, v130, v142)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	if v151 == int32(37) {
		v212 = v130
		v214 = v132
		goto L38
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v156 == int32(0) {
		v189 = v130
		goto L40
	} else {
		goto L58
	}
L58:
	;
	goto L48
L59:
	;
	v166 = int32(0)
	v167 = v163 + v142
	goto L47
L60:
	;
	v162 = F__emscripten_memcpy_bulkmem(m, v132, v130, v142)
	mBase = m.M
	v163 = v162
	goto L62
L61:
	;
	v163 = v132
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L45
L64:
	;
	v196 = v173
	v198 = v175
	goto L39
L65:
	;
	v230 = F_strlen(m, v228)
	mBase = m.M
	F_make_trigrams(m, v17+int32(4), v228, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_pfree(m, v228)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v236 = l0 - v212 + l1
	if int32(0) < v236 {
		v53 = v212
		v62 = v236
		goto L11
	} else {
		goto L68
	}
L68:
	;
	goto L33
L69:
	;
	if int32(2) <= v254 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v260 = v253 + int32(5)
	F_pg_qsort(m, v260, v254, int32(3), int32(7429))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	v314 = v254
	goto L72
L72:
	;
	v324 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+4)) = uint8(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v314*int32(12) + int32(20)
	v339 = v253
	goto L2
L73:
	;
	v269 = int32(1)
	v270 = int32(0)
	goto L74
L74:
	;
	v281 = int32(3)
	v283 = v260 + v269*v281
	v288 = *(*int32)(unsafe.Add(mBase, _consts[1451]))
	v289 = m.T0[v288].(func(*base.Module, int32, int32) int32)(m, v283, v260+v270*v281)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	v314 = v303 + int32(1)
	goto L72
L76:
	;
	v306 = v269 + int32(1)
	if v306 != v254 {
		v269 = v306
		v270 = v303
		goto L74
	} else {
		goto L80
	}
L77:
	;
	if v289 == int32(0) {
		v303 = v270
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v294 = v270 + int32(1)
	if v269 == v294 {
		v303 = v269
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v298 = v260 + v294*int32(3)
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283))))
	*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v299)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v301)
	v303 = v294
	goto L76
L80:
	;
	goto L75
L81:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(14020), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(518898), int32(115), int32(24857))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_german_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v6
	goto L1
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 < v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v64 != 0 {
		v203 = v65
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v21
	goto L6
L6:
	;
	goto L8
L7:
	;
	v64 = v60
	goto L3
L8:
	;
	if v20 == v23 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v60 = int32(0)
	goto L7
L10:
	;
	v64 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v35 = int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v20))))
	if int32(252) < v38 {
		v60 = v35
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v40 = v38 - int32(97)
	if v40 < int32(0) {
		v60 = v35
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v46)>>(uint(v40&int32(7))%32))&int32(1) == int32(0) {
		v60 = v35
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + int32(1)
	goto L16
L16:
	;
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L1
L18:
	;
	return v1018
L19:
	;
	v1013 = F_slice_from_s(m, l0, int32(1), int32(2215941))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L62
	} else {
		goto L303
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v218 = v6
	goto L65
L21:
	;
	v210 = F_slice_from_s(m, l0, int32(1), int32(2215940))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L62
	} else {
		goto L63
	}
L22:
	;
	if v203 <= v10 {
		goto L20
	} else {
		goto L61
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v66
	if v66 == v65 {
		v134 = v65
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
	if v66 == v134 {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v66))))
	if v71 != int32(117) {
		v134 = v65
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v75 = v66 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v87 < v75 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v130 == int32(0) {
		goto L21
	} else {
		goto L41
	}
L28:
	;
	v89 = v75
	goto L30
L29:
	;
	v89 = v87
	goto L30
L30:
	;
	goto L32
L31:
	;
	v130 = v126
	goto L27
L32:
	;
	if v75 == v89 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v126 = int32(0)
	goto L31
L34:
	;
	v130 = int32(-1)
	goto L27
L35:
	;
	goto L36
L36:
	;
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v75))))
	if int32(252) < v104 {
		v126 = v101
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v106 = v104 - int32(97)
	if v106 < int32(0) {
		v126 = v101
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) == int32(0) {
		v126 = v101
		goto L31
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(2)
	goto L40
L40:
	;
	goto L33
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = v133
	goto L24
L42:
	;
	v203 = v66
	goto L22
L43:
	;
	goto L44
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v66))))
	if v139 != int32(121) {
		v203 = v134
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v143 = v66 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v143
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v155 < v143 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v198 == int32(0) {
		goto L19
	} else {
		goto L60
	}
L47:
	;
	v157 = v143
	goto L49
L48:
	;
	v157 = v155
	goto L49
L49:
	;
	goto L51
L50:
	;
	v198 = v194
	goto L46
L51:
	;
	if v143 == v157 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v194 = int32(0)
	goto L50
L53:
	;
	v198 = int32(-1)
	goto L46
L54:
	;
	goto L55
L55:
	;
	v169 = int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v143))))
	if int32(252) < v172 {
		v194 = v169
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v174 = v172 - int32(97)
	if v174 < int32(0) {
		v194 = v169
		goto L50
	} else {
		goto L57
	}
L57:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v174)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v180)>>(uint(v174&int32(7))%32))&int32(1) == int32(0) {
		v194 = v169
		goto L50
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(2)
	goto L59
L59:
	;
	goto L52
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v203 = v201
	goto L22
L61:
	;
	v206 = v10 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206
	v10 = v206
	goto L1
L62:
	;
	return int32(0)
L63:
	;
	if v210 < int32(0) {
		v1018 = v210
		goto L18
	} else {
		goto L64
	}
L64:
	;
	goto L17
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v218
	v225 = F_find_among(m, l0, int32(4243936), int32(6))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L62
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = v255
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = v266 + int32(3)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v269 < v268 {
		goto L84
	} else {
		goto L85
	}
L67:
	;
	goto L66
L68:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v227
	switch v225 - int32(1) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		goto L69
	}
L69:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = v261
	goto L65
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v255 <= v227 {
		goto L67
	} else {
		goto L83
	}
L71:
	;
	v251 = F_slice_from_s(m, l0, int32(1), int32(2215946))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L62
	} else {
		goto L81
	}
L72:
	;
	v245 = F_slice_from_s(m, l0, int32(1), int32(2215945))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L62
	} else {
		goto L79
	}
L73:
	;
	v239 = F_slice_from_s(m, l0, int32(1), int32(2215944))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L62
	} else {
		goto L77
	}
L74:
	;
	v233 = F_slice_from_s(m, l0, int32(2), int32(2215942))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L62
	} else {
		goto L75
	}
L75:
	;
	if int32(0) <= v233 {
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v1018 = v233
	goto L18
L77:
	;
	if int32(0) <= v239 {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v1018 = v239
	goto L18
L79:
	;
	if int32(0) <= v245 {
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v1018 = v245
	goto L18
L81:
	;
	if int32(0) <= v251 {
		goto L69
	} else {
		goto L82
	}
L82:
	;
	v1018 = v251
	goto L18
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v227 + int32(1)
	goto L69
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	if v506 <= v6 {
		goto L151
	} else {
		goto L152
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v268
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v266
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v281 < v266 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v321 < int32(0) {
		goto L84
	} else {
		goto L101
	}
L87:
	;
	v283 = v266
	goto L89
L88:
	;
	v283 = v281
	goto L89
L89:
	;
	v290 = v266
	goto L91
L90:
	;
	v321 = v301
	goto L86
L91:
	;
	if v290 == v283 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v321 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v290))))
	if int32(252) < v296 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v313 = v290 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v290 = v313
	goto L91
L97:
	;
	v298 = v296 - int32(97)
	if v298 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v301 = int32(1)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v298)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v305)>>(uint(v298&int32(7))%32))&v301 != 0 {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L96
L101:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v325 = v324 + v321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v325
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v336 < v325 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v379 < int32(0) {
		goto L84
	} else {
		goto L116
	}
L103:
	;
	v338 = v325
	goto L105
L104:
	;
	v338 = v336
	goto L105
L105:
	;
	v345 = v325
	goto L107
L106:
	;
	v379 = int32(1)
	goto L102
L107:
	;
	if v345 == v338 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v379 = int32(-1)
	goto L102
L110:
	;
	goto L111
L111:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351+v345))))
	if int32(252) < v353 {
		goto L106
	} else {
		goto L112
	}
L112:
	;
	v355 = v353 - int32(97)
	if v355 < int32(0) {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v355)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v361)>>(uint(v355&int32(7))%32))&int32(1) == int32(0) {
		goto L106
	} else {
		goto L114
	}
L114:
	;
	v370 = v345 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v370
	v345 = v370
	goto L107
L116:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v383 = v382 + v379
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v386 < v383 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v388 = v383
	goto L119
L118:
	;
	v388 = v386
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385)+8)) = v388
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v398 < v397 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v438 < int32(0) {
		goto L84
	} else {
		goto L135
	}
L121:
	;
	v400 = v397
	goto L123
L122:
	;
	v400 = v398
	goto L123
L123:
	;
	v407 = v397
	goto L125
L124:
	;
	v438 = v418
	goto L120
L125:
	;
	if v407 == v400 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v438 = int32(-1)
	goto L120
L128:
	;
	goto L129
L129:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v407))))
	if int32(252) < v413 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v430 = v407 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430
	v407 = v430
	goto L125
L131:
	;
	v415 = v413 - int32(97)
	if v415 < int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v418 = int32(1)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v415)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v422)>>(uint(v415&int32(7))%32))&v418 != 0 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	goto L130
L135:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v442 = v441 + v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v442
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v453 < v442 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v496 < int32(0) {
		goto L84
	} else {
		goto L150
	}
L137:
	;
	v455 = v442
	goto L139
L138:
	;
	v455 = v453
	goto L139
L139:
	;
	v462 = v442
	goto L141
L140:
	;
	v496 = int32(1)
	goto L136
L141:
	;
	if v462 == v455 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v496 = int32(-1)
	goto L136
L144:
	;
	goto L145
L145:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v462))))
	if int32(252) < v470 {
		goto L140
	} else {
		goto L146
	}
L146:
	;
	v472 = v470 - int32(97)
	if v472 < int32(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v472)>>(uint(int32(3))%32)))+uint32(_consts[1432]))))
	if int32(base.Ui32(v478)>>(uint(v472&int32(7))%32))&int32(1) == int32(0) {
		goto L140
	} else {
		goto L148
	}
L148:
	;
	v487 = v462 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v487
	v462 = v487
	goto L141
L150:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+4)) = v500 + v496
	goto L84
L151:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v674
	v678 = v674 - int32(1)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v678 <= v679 {
		goto L199
	} else {
		goto L200
	}
L152:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510+v506-int32(1)))))
	if v514&int32(224) != int32(96) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	if int32(1)<<(uint(v514)%32)&int32(811040) == int32(0) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v527 = F_find_among_b(m, l0, int32(4244064), int32(11))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L62
	} else {
		goto L155
	}
L155:
	;
	if v527 == int32(0) {
		goto L151
	} else {
		goto L156
	}
L156:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v531
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	if v531 < v534 {
		goto L151
	} else {
		goto L157
	}
L157:
	;
	switch v527 - int32(1) {
	case 0:
		goto L162
	case 1:
		goto L161
	case 2:
		goto L160
	case 3:
		goto L159
	case 4:
		goto L158
	default:
		goto L151
	}
L158:
	;
	v667 = F_slice_from_s(m, l0, int32(1), int32(2215966))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L62
	} else {
		goto L197
	}
L159:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L185
L160:
	;
	v567 = F_slice_del(m, l0)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L62
	} else {
		goto L172
	}
L161:
	;
	v563 = F_slice_del(m, l0)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L62
	} else {
		goto L170
	}
L162:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v539 = int32(4)
	v541 = int32(0)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v543-v544 < v539 {
		v554 = v541
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v554 != 0 {
		goto L151
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = F_memcmp(m, v547+v543-v539, int32(2215956), v539)
	mBase = m.M
	if v550 != 0 {
		v554 = v541
		goto L164
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - v539
	v554 = int32(1)
	goto L164
L167:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v555 + (v531 - v538)
	v559 = F_slice_del(m, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L62
	} else {
		goto L168
	}
L168:
	;
	if int32(0) <= v559 {
		goto L151
	} else {
		goto L169
	}
L169:
	;
	v1018 = v559
	goto L18
L170:
	;
	if int32(0) <= v563 {
		goto L151
	} else {
		goto L171
	}
L171:
	;
	v1018 = v563
	goto L18
L172:
	;
	if v567 < int32(0) {
		v1018 = v567
		goto L18
	} else {
		goto L173
	}
L173:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v571
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v571 <= v573 {
		goto L151
	} else {
		goto L174
	}
L174:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v571-int32(1)))))
	if v579 != int32(115) {
		goto L151
	} else {
		goto L175
	}
L175:
	;
	v583 = v571 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v583
	v586 = int32(3)
	v588 = int32(0)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v583-v591 < v586 {
		v601 = v588
		goto L177
	} else {
		goto L178
	}
L176:
	;
	if v601 == int32(0) {
		goto L151
	} else {
		goto L180
	}
L177:
	;
	goto L176
L178:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v597 = F_memcmp(m, v594+v583-v586, int32(2215960), v586)
	mBase = m.M
	if v597 != 0 {
		v601 = v588
		goto L177
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v583 - v586
	v601 = int32(1)
	goto L177
L180:
	;
	v604 = F_slice_del(m, l0)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L62
	} else {
		goto L181
	}
L181:
	;
	if int32(0) <= v604 {
		goto L151
	} else {
		goto L182
	}
L182:
	;
	v1018 = v604
	goto L18
L183:
	;
	if v660 != 0 {
		goto L151
	} else {
		goto L194
	}
L184:
	;
	v660 = v656
	goto L183
L185:
	;
	if v616 <= v617 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v656 = int32(0)
	goto L184
L187:
	;
	v660 = int32(-1)
	goto L183
L188:
	;
	goto L189
L189:
	;
	v629 = int32(1)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630+v616-v629))))
	if int32(116) < v634 {
		v656 = v629
		goto L184
	} else {
		goto L190
	}
L190:
	;
	v636 = v634 - int32(98)
	if v636 < int32(0) {
		v656 = v629
		goto L184
	} else {
		goto L191
	}
L191:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v636)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v642)>>(uint(v636&int32(7))%32))&int32(1) == int32(0) {
		v656 = v629
		goto L184
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616 - int32(1)
	goto L193
L193:
	;
	goto L186
L194:
	;
	v661 = F_slice_del(m, l0)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L62
	} else {
		goto L195
	}
L195:
	;
	if int32(0) <= v661 {
		goto L151
	} else {
		goto L196
	}
L196:
	;
	v1018 = v661
	goto L18
L197:
	;
	if v667 < int32(0) {
		v1018 = v667
		goto L18
	} else {
		goto L198
	}
L198:
	;
	goto L151
L199:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v776
	v780 = v776 - int32(1)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v780 <= v781 {
		goto L225
	} else {
		goto L226
	}
L200:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681+v678))))
	if v683&int32(224) != int32(96) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	if int32(1)<<(uint(v683)%32)&int32(1327104) == int32(0) {
		goto L199
	} else {
		goto L202
	}
L202:
	;
	v696 = F_find_among_b(m, l0, int32(4244288), int32(4))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L62
	} else {
		goto L203
	}
L203:
	;
	if v696 == int32(0) {
		goto L199
	} else {
		goto L204
	}
L204:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v700
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+8))
	if v700 < v703 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	switch v696 - int32(1) {
	case 0:
		goto L207
	case 1:
		goto L206
	default:
		goto L199
	}
L206:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L212
L207:
	;
	v707 = F_slice_del(m, l0)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L62
	} else {
		goto L208
	}
L208:
	;
	if int32(0) <= v707 {
		goto L199
	} else {
		goto L209
	}
L209:
	;
	v1018 = v707
	goto L18
L210:
	;
	if v763 != 0 {
		goto L199
	} else {
		goto L221
	}
L211:
	;
	v763 = v759
	goto L210
L212:
	;
	if v719 <= v720 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v759 = int32(0)
	goto L211
L214:
	;
	v763 = int32(-1)
	goto L210
L215:
	;
	goto L216
L216:
	;
	v732 = int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733+v719-v732))))
	if int32(116) < v737 {
		v759 = v732
		goto L211
	} else {
		goto L217
	}
L217:
	;
	v739 = v737 - int32(98)
	if v739 < int32(0) {
		v759 = v732
		goto L211
	} else {
		goto L218
	}
L218:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v739)>>(uint(int32(3))%32)))+uint32(_consts[1434]))))
	if int32(base.Ui32(v745)>>(uint(v739&int32(7))%32))&int32(1) == int32(0) {
		v759 = v732
		goto L211
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v719 - int32(1)
	goto L220
L220:
	;
	goto L213
L221:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v766 = v764 - int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v766
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v766 < v768 {
		goto L199
	} else {
		goto L222
	}
L222:
	;
	v770 = F_slice_del(m, l0)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L62
	} else {
		goto L223
	}
L223:
	;
	if v770 < int32(0) {
		v1018 = v770
		goto L18
	} else {
		goto L224
	}
L224:
	;
	goto L199
L225:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v962
	v965 = v962
	goto L284
L226:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v780))))
	if v785&int32(224) != int32(96) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	if int32(1)<<(uint(v785)%32)&int32(1051024) == int32(0) {
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v798 = F_find_among_b(m, l0, int32(4244368), int32(8))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L62
	} else {
		goto L229
	}
L229:
	;
	if v798 == int32(0) {
		goto L225
	} else {
		goto L230
	}
L230:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v802
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v802 < v805 {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	switch v798 - int32(1) {
	case 0:
		goto L235
	case 1:
		goto L234
	case 2:
		goto L233
	case 3:
		goto L232
	default:
		goto L225
	}
L232:
	;
	v920 = F_slice_del(m, l0)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L62
	} else {
		goto L273
	}
L233:
	;
	v864 = F_slice_del(m, l0)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L62
	} else {
		goto L256
	}
L234:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v851 < v802 {
		goto L250
	} else {
		goto L251
	}
L235:
	;
	v809 = F_slice_del(m, l0)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L62
	} else {
		goto L236
	}
L236:
	;
	if v809 < int32(0) {
		v1018 = v809
		goto L18
	} else {
		goto L237
	}
L237:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v813
	v815 = int32(2)
	v817 = int32(0)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v813-v820 < v815 {
		v830 = v817
		goto L239
	} else {
		goto L240
	}
L238:
	;
	if v830 == int32(0) {
		goto L225
	} else {
		goto L242
	}
L239:
	;
	goto L238
L240:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v826 = F_memcmp(m, v823+v813-v815, int32(2215970), v815)
	mBase = m.M
	if v826 != 0 {
		v830 = v817
		goto L239
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - v815
	v830 = int32(1)
	goto L239
L242:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v835 < v833 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837+v833-int32(1)))))
	if v841 == int32(101) {
		goto L225
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
	if v833 < v845 {
		goto L225
	} else {
		goto L247
	}
L246:
	;
	goto L245
L247:
	;
	v847 = F_slice_del(m, l0)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L62
	} else {
		goto L248
	}
L248:
	;
	if int32(0) <= v847 {
		goto L225
	} else {
		goto L249
	}
L249:
	;
	v1018 = v847
	goto L18
L250:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853+v802-int32(1)))))
	if v857 == int32(101) {
		goto L225
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v860 = F_slice_del(m, l0)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L62
	} else {
		goto L254
	}
L253:
	;
	goto L252
L254:
	;
	if int32(0) <= v860 {
		goto L225
	} else {
		goto L255
	}
L255:
	;
	v1018 = v860
	goto L18
L256:
	;
	if v864 < int32(0) {
		v1018 = v864
		goto L18
	} else {
		goto L257
	}
L257:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v868
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v871 = int32(2)
	v873 = int32(0)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v868-v876 < v871 {
		v886 = v873
		goto L259
	} else {
		goto L260
	}
L258:
	;
	if v886 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	goto L258
L260:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v882 = F_memcmp(m, v879+v868-v871, int32(2215972), v871)
	mBase = m.M
	if v882 != 0 {
		v886 = v873
		goto L259
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v868 - v871
	v886 = int32(1)
	goto L259
L262:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v891 = v889 + (v868 - v870)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v891
	v893 = int32(2)
	v895 = int32(0)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v891-v898 < v893 {
		v908 = v895
		goto L266
	} else {
		goto L267
	}
L263:
	;
	goto L264
L264:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v911
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+8))
	if v911 < v914 {
		goto L225
	} else {
		goto L270
	}
L265:
	;
	if v908 == int32(0) {
		goto L225
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v904 = F_memcmp(m, v901+v891-v893, int32(2215974), v893)
	mBase = m.M
	if v904 != 0 {
		v908 = v895
		goto L266
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v891 - v893
	v908 = int32(1)
	goto L266
L269:
	;
	goto L264
L270:
	;
	v916 = F_slice_del(m, l0)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L62
	} else {
		goto L271
	}
L271:
	;
	if int32(0) <= v916 {
		goto L225
	} else {
		goto L272
	}
L272:
	;
	v1018 = v916
	goto L18
L273:
	;
	if v920 < int32(0) {
		v1018 = v920
		goto L18
	} else {
		goto L274
	}
L274:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v924
	v927 = v924 - int32(1)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v927 <= v928 {
		goto L225
	} else {
		goto L275
	}
L275:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930+v927))))
	if base.Ui32(int32(1)) < base.Ui32((v932-int32(103))&int32(255)) {
		goto L225
	} else {
		goto L276
	}
L276:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v942 = F_find_among_b(m, l0, int32(4244528), int32(2))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L62
	} else {
		goto L277
	}
L277:
	;
	if v942 == int32(0) {
		goto L225
	} else {
		goto L278
	}
L278:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v946
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v946 < v949 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951 + (v924 - v939)
	goto L225
L280:
	;
	goto L281
L281:
	;
	v955 = F_slice_del(m, l0)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L62
	} else {
		goto L282
	}
L282:
	;
	if v955 < int32(0) {
		v1018 = v955
		goto L18
	} else {
		goto L283
	}
L283:
	;
	goto L225
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v965
	v972 = F_find_among(m, l0, int32(4244576), int32(6))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L62
	} else {
		goto L287
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v962
	v1018 = int32(1)
	goto L18
L286:
	;
	goto L285
L287:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v974
	switch v972 - int32(1) {
	case 0:
		goto L293
	case 1:
		goto L292
	case 2:
		goto L291
	case 3:
		goto L290
	case 4:
		goto L289
	default:
		goto L288
	}
L288:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v965 = v1008
	goto L284
L289:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1002 <= v974 {
		goto L286
	} else {
		goto L302
	}
L290:
	;
	v998 = F_slice_from_s(m, l0, int32(1), int32(2216049))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L62
	} else {
		goto L300
	}
L291:
	;
	v992 = F_slice_from_s(m, l0, int32(1), int32(2216048))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L62
	} else {
		goto L298
	}
L292:
	;
	v986 = F_slice_from_s(m, l0, int32(1), int32(2216047))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L62
	} else {
		goto L296
	}
L293:
	;
	v980 = F_slice_from_s(m, l0, int32(1), int32(2216046))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L62
	} else {
		goto L294
	}
L294:
	;
	if int32(0) <= v980 {
		goto L288
	} else {
		goto L295
	}
L295:
	;
	v1018 = v980
	goto L18
L296:
	;
	if int32(0) <= v986 {
		goto L288
	} else {
		goto L297
	}
L297:
	;
	v1018 = v986
	goto L18
L298:
	;
	if int32(0) <= v992 {
		goto L288
	} else {
		goto L299
	}
L299:
	;
	v1018 = v992
	goto L18
L300:
	;
	if int32(0) <= v998 {
		goto L288
	} else {
		goto L301
	}
L301:
	;
	v1018 = v998
	goto L18
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v974 + int32(1)
	goto L288
L303:
	;
	if int32(0) <= v1013 {
		goto L17
	} else {
		goto L304
	}
L304:
	;
	v1018 = v1013
	goto L18
}
func F_getIthJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9&int32(1073741824) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v9 & int32(268435455)
	if base.Ui32(l1) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L41
	}
L4:
	;
	v16 = l0 + int32(4)
	v19 = v16 + v13<<(uint(int32(2))%32)
	v21 = F_palloc(m, int32(20))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v176 = v3
	goto L6
L6:
	;
	return v176
L7:
	;
	return int32(0)
L8:
	;
	v27 = l1
	v29 = v3
	goto L9
L9:
	;
	if int32(0) < v27 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v52 = l0 + int32(4)
	v55 = v52 + l1<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	switch int32(base.Ui32(v56)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L17
	case 3:
		goto L18
	case 4:
		goto L21
	default:
		goto L16
	}
L11:
	;
	v36 = v27 - int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16+v36<<(uint(int32(2))%32))))
	v43 = v40&int32(268435455) + v29
	if int32(0) <= v40 {
		v27 = v36
		v29 = v43
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v47 = v29
	goto L13
L13:
	;
	goto L10
L14:
	;
	v47 = v43
	goto L13
L15:
	;
	v176 = v21
	goto L6
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(18)
	v121 = (v47 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v19 + v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v124 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(3)
	goto L15
L18:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(3)
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v19 + (v47+int32(3))&int32(-4)
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v19 + v47
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v67 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	goto L15
L22:
	;
	v72 = l1
	v73 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v67 & int32(268435455)
	goto L15
L25:
	;
	v79 = v72 - int32(1)
	if int32(0) <= v79 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v67&int32(268435455) - v91
	goto L15
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v52+v79<<(uint(int32(2))%32))))
	v88 = v85&int32(268435455) + v73
	if int32(0) <= v85 {
		v72 = v79
		v73 = v88
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v91 = v73
	goto L29
L29:
	;
	goto L26
L30:
	;
	v91 = v88
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v162 + (v47 - v121)
	goto L15
L32:
	;
	v129 = l1
	v130 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v162 = v124 & int32(268435455)
	goto L31
L35:
	;
	v136 = v129 - int32(1)
	if int32(0) <= v136 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v162 = v124&int32(268435455) - v148
	goto L31
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v52+v136<<(uint(int32(2))%32))))
	v145 = v142&int32(268435455) + v130
	if int32(0) <= v142 {
		v129 = v136
		v130 = v145
		goto L35
	} else {
		goto L40
	}
L38:
	;
	v148 = v130
	goto L39
L39:
	;
	goto L36
L40:
	;
	v148 = v145
	goto L39
L41:
	;
	F_errmsg_internal(m, int32(26572), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(520985), int32(479), int32(229635))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getOwnedSequences(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_getOwnedSequences_internal(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_getRelationsInNamespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = int32(3)
	F_ScanKeyInit(m, v8, v10, v10, int32(184), l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v8+int32(48), int32(18), int32(3), int32(61), l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = F_table_beginscan_catalog(m, v27, int32(2), v8)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(0)
	goto L6
L6:
	;
	v37 = F_heap_getnext(m, v30)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+188))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	m.T0[v47].(func(*base.Module, int32))(m, v30)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40)))
	v43 = F_lappend_oid(m, v32, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L7
L12:
	;
	v32 = v43
	goto L6
L13:
	;
	F_sequence_close(m, v27, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	m.G0 = v8 + int32(96)
	return v32
}
func F_getWeights(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 float32
	_ = v27
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v97 float32
	_ = v97
	var v107 float32
	_ = v107
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(1) {
		v11 = F_ArrayGetNItems(m, int32(1), l0+int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 <= int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errmsg(m, int32(87101), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(521170), int32(420), int32(131972))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v15 = F_array_contains_nulls(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if v15 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errmsg(m, int32(162011), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_errfinish(m, int32(521170), int32(425), int32(131972))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v17 != 0 {
							v25 = v17
						} else {
							v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v25 = (v18<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						v26 = v25 + l0
						v27 = *(*float32)(unsafe.Add(mBase, uint32(v26)))
						if base.F32_ge(v27, float32(0)) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1036831949)
							v87 = *(*float32)(unsafe.Add(mBase, uint32(v26)+4))
							if base.F32_ge(v87, float32(0)) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1045220557)
								v97 = *(*float32)(unsafe.Add(mBase, uint32(v26)+8))
								if base.F32_ge(v97, float32(0)) == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1053609165)
									v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
									if base.F32_ge(v107, float32(0)) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
										return
									} else {
										*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
										if base.F32_gt(v107, float32(1)) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													F_errmsg(m, int32(420456), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(521170), int32(434), int32(131972))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											return
										}
									}
								} else {
									*(*float32)(unsafe.Add(mBase, uint32(l1)+8)) = v97
									if base.F32_gt(v97, float32(1)) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												F_errmsg(m, int32(420456), int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													F_errfinish(m, int32(521170), int32(434), int32(131972))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
										if base.F32_ge(v107, float32(0)) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
											return
										} else {
											*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
											if base.F32_gt(v107, float32(1)) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														F_errmsg(m, int32(420456), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(521170), int32(434), int32(131972))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return
											}
										}
									}
								}
							} else {
								*(*float32)(unsafe.Add(mBase, uint32(l1)+4)) = v87
								if base.F32_gt(v87, float32(1)) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return
										} else {
											F_errmsg(m, int32(420456), int32(0))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return
											} else {
												F_errfinish(m, int32(521170), int32(434), int32(131972))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v97 = *(*float32)(unsafe.Add(mBase, uint32(v26)+8))
									if base.F32_ge(v97, float32(0)) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1053609165)
										v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
										if base.F32_ge(v107, float32(0)) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
											return
										} else {
											*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
											if base.F32_gt(v107, float32(1)) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														F_errmsg(m, int32(420456), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(521170), int32(434), int32(131972))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return
											}
										}
									} else {
										*(*float32)(unsafe.Add(mBase, uint32(l1)+8)) = v97
										if base.F32_gt(v97, float32(1)) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													F_errmsg(m, int32(420456), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(521170), int32(434), int32(131972))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
											if base.F32_ge(v107, float32(0)) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
												return
											} else {
												*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
												if base.F32_gt(v107, float32(1)) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															F_errmsg(m, int32(420456), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(521170), int32(434), int32(131972))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return
												}
											}
										}
									}
								}
							}
						} else {
							*(*float32)(unsafe.Add(mBase, uint32(l1))) = v27
							if base.F32_gt(v27, float32(1)) == int32(0) {
								v87 = *(*float32)(unsafe.Add(mBase, uint32(v26)+4))
								if base.F32_ge(v87, float32(0)) == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1045220557)
									v97 = *(*float32)(unsafe.Add(mBase, uint32(v26)+8))
									if base.F32_ge(v97, float32(0)) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1053609165)
										v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
										if base.F32_ge(v107, float32(0)) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
											return
										} else {
											*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
											if base.F32_gt(v107, float32(1)) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														F_errmsg(m, int32(420456), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(521170), int32(434), int32(131972))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return
											}
										}
									} else {
										*(*float32)(unsafe.Add(mBase, uint32(l1)+8)) = v97
										if base.F32_gt(v97, float32(1)) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													F_errmsg(m, int32(420456), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(521170), int32(434), int32(131972))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
											if base.F32_ge(v107, float32(0)) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
												return
											} else {
												*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
												if base.F32_gt(v107, float32(1)) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															F_errmsg(m, int32(420456), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(521170), int32(434), int32(131972))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return
												}
											}
										}
									}
								} else {
									*(*float32)(unsafe.Add(mBase, uint32(l1)+4)) = v87
									if base.F32_gt(v87, float32(1)) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												F_errmsg(m, int32(420456), int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													F_errfinish(m, int32(521170), int32(434), int32(131972))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v97 = *(*float32)(unsafe.Add(mBase, uint32(v26)+8))
										if base.F32_ge(v97, float32(0)) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1053609165)
											v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
											if base.F32_ge(v107, float32(0)) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
												return
											} else {
												*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
												if base.F32_gt(v107, float32(1)) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															F_errmsg(m, int32(420456), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(521170), int32(434), int32(131972))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return
												}
											}
										} else {
											*(*float32)(unsafe.Add(mBase, uint32(l1)+8)) = v97
											if base.F32_gt(v97, float32(1)) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														F_errmsg(m, int32(420456), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(521170), int32(434), int32(131972))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v107 = *(*float32)(unsafe.Add(mBase, uint32(v26)+12))
												if base.F32_ge(v107, float32(0)) == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(1065353216)
													return
												} else {
													*(*float32)(unsafe.Add(mBase, uint32(l1)+12)) = v107
													if base.F32_gt(v107, float32(1)) != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																F_errmsg(m, int32(420456), int32(0))
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521170), int32(434), int32(131972))
																	mBase = m.M
																	v133 = m.ExcPending
																	if v133 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														return
													}
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return
									} else {
										F_errmsg(m, int32(420456), int32(0))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											F_errfinish(m, int32(521170), int32(434), int32(131972))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
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
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			F_errcode(m, int32(352845954))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_errmsg(m, int32(327604), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_errfinish(m, int32(521170), int32(415), int32(131972))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
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
}
func F_get_attname(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v21 = F_pstrdup(m, v16+v17+int32(4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v40 = v21
					m.G0 = v9 + int32(16)
					return v40
				}
			}
		} else {
			if l2 != 0 {
				v40 = int32(0)
				m.G0 = v9 + int32(16)
				return v40
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(50140), v9)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(522775), int32(937), int32(395282))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
func F_get_best_segment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+1468))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v9 != v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v51 = int32(15)
	v54 = int32(32) - base.I32_clz(l1)
	if base.Ui32(v51) <= base.Ui32(v54) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v23 = l0 + int32(8) + v18*int32(20)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v9
	goto L3
L6:
	;
	v40 = v18 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v40) <= base.Ui32(v41) {
		v18 = v40
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
	if v27 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	F_dsm_detach(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(0)
	goto L6
L11:
	;
	goto L5
L12:
	;
	v57 = v51
	goto L14
L13:
	;
	v57 = v54
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v59 = v57
	goto L17
L16:
	;
	v59 = int32(0)
	goto L17
L17:
	;
	v65 = v59
	goto L19
L18:
	;
	return v116
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v65<<(uint(int32(2))%32))+160))
	if v71 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v116 = int32(0)
	goto L18
L21:
	;
	v74 = int32(1)
	v77 = v74 << (uint(v65-v74) % 32)
	v81 = v71
	goto L24
L22:
	;
	goto L23
L23:
	;
	v110 = v65 + int32(1)
	if v110 != int32(16) {
		v65 = v110
		goto L19
	} else {
		goto L36
	}
L24:
	;
	v85 = F_get_segment_by_index(m, l0, v81)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if base.B2i32(base.Ui32(v77) <= base.Ui32(v90))&base.B2i32(base.Ui32(v90) < base.Ui32(l1)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui32(v90) < base.Ui32(v77) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v88 != int32(-1) {
		v81 = v88
		goto L24
	} else {
		goto L35
	}
L30:
	;
	F_rebin_segment(m, l0, v85)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(l1) <= base.Ui32(v90) {
		v116 = v85
		goto L18
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L29
L35:
	;
	goto L25
L36:
	;
	goto L20
}
func F_get_doc_path(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_make_relative_path(m, l0, int32(314572), int32(4550096))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_get_etc_path(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	F_make_relative_path(m, l1, int32(314330), l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_get_indexpath_pages(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 float64
	_ = v78
	v2 = float64(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(280) {
	case 0:
		goto L2
	default:
		goto L3
	case 3:
		goto L5
	case 4:
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v78
L2:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v78 = base.F64_convert_i32_u(v75)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L23
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v37 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v12 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v78 = v2
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		v78 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v19 = int32(0)
	v20 = v2
	goto L10
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32))))
	v28 = F_get_indexpath_pages(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v78 = v32
	goto L1
L12:
	;
	return float64(0)
L13:
	;
	v32 = base.F64_add(v20, v28)
	v34 = v19 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v34 < v35 {
		v19 = v34
		v20 = v32
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v78 = v2
	goto L1
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		v78 = v2
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v44 = int32(0)
	v45 = v2
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v44<<(uint(int32(2))%32))))
	v53 = F_get_indexpath_pages(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	v78 = v55
	goto L1
L21:
	;
	v55 = base.F64_add(v45, v53)
	v57 = v44 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v57 < v58 {
		v44 = v57
		v45 = v55
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
	F_errmsg_internal(m, int32(507644), v7)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(521926), int32(1003), int32(180390))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_joinrel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 float64
	_ = v613
	var v614 float64
	_ = v614
	var v615 float64
	_ = v615
	var v616 int32
	_ = v616
	var v617 float64
	_ = v617
	var v619 float64
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v643 int32
	_ = v643
	v8 = int32(0)
	if l5 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = F_bms_union(m, v27, l5)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v36 = F_bms_union(m, v34, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v38 = v8
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v38 = v36
	goto L8
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v43 = F_bms_union(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v45 = v8
	goto L12
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v46 == int32(0) {
		v149 = v8
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v45 = v43
	goto L12
L14:
	;
	v159 = F_generate_join_implied_equalities(m, l0, v28, l5, l1, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L38
	}
L15:
	;
	v49 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v50 <= v49 {
		v149 = v8
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = v49
	v66 = v8
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v60<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v81 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v83 = F_bms_is_subset(m, v82, v28)
	mBase = m.M
	if v83 == v81 {
		v94 = v81
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v149 = v131
	goto L14
L19:
	;
	v133 = v60 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v133 < v134 {
		v60 = v133
		v66 = v131
		goto L17
	} else {
		goto L36
	}
L20:
	;
	if v94 == int32(0) {
		v131 = v66
		goto L19
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v87 = F_bms_overlap(m, v80, v86)
	mBase = m.M
	if v87 == int32(0) {
		v94 = v81
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	v91 = F_bms_overlap(m, v80, v90)
	mBase = m.M
	v94 = v91 ^ int32(1)
	goto L21
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v101 = F_bms_is_subset(m, v100, v38)
	mBase = m.M
	if v101 == v99 {
		v112 = v99
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v112 != 0 {
		v131 = v66
		goto L19
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v105 = F_bms_overlap(m, v98, v104)
	mBase = m.M
	if v105 == int32(0) {
		v112 = v99
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	v109 = F_bms_overlap(m, v98, v108)
	mBase = m.M
	v112 = v109 ^ int32(1)
	goto L26
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v117 = F_bms_is_subset(m, v116, v45)
	mBase = m.M
	if v117 == v115 {
		v128 = v115
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v128 != 0 {
		v131 = v66
		goto L19
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v121 = F_bms_overlap(m, v114, v120)
	mBase = m.M
	if v121 == int32(0) {
		v128 = v115
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	v125 = F_bms_overlap(m, v114, v124)
	mBase = m.M
	v128 = v125 ^ int32(1)
	goto L31
L34:
	;
	v129 = F_lappend(m, v66, v79)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v131 = v129
	goto L19
L36:
	;
	goto L18
L37:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v496 = F_list_concat(m, v486, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L118
	}
L38:
	;
	if v159 == int32(0) {
		v486 = v149
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v163 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v163 < v164 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v175 = int32(0)
	v179 = v163
	v181 = v149
	goto L43
L41:
	;
	v249 = v163
	v251 = v149
	goto L42
L42:
	;
	if v249 == int32(0) {
		v486 = v251
		goto L37
	} else {
		goto L61
	}
L43:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v175<<(uint(int32(2))%32))))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	v197 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v199 = F_bms_is_subset(m, v198, v38)
	mBase = m.M
	if v199 == v197 {
		v210 = v197
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v249 = v232
	v251 = v233
	goto L42
L45:
	;
	v235 = v175 + int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v235 < v236 {
		v175 = v235
		v179 = v232
		v181 = v233
		goto L43
	} else {
		goto L60
	}
L46:
	;
	if v210 != 0 {
		v232 = v179
		v233 = v181
		goto L45
	} else {
		goto L50
	}
L47:
	;
	goto L46
L48:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v203 = F_bms_overlap(m, v196, v202)
	mBase = m.M
	if v203 == int32(0) {
		v210 = v197
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v194)+40))
	v207 = F_bms_overlap(m, v196, v206)
	mBase = m.M
	v210 = v207 ^ int32(1)
	goto L47
L50:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v215 = F_bms_is_subset(m, v214, v45)
	mBase = m.M
	if v215 == v213 {
		v226 = v213
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v226 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v219 = F_bms_overlap(m, v212, v218)
	mBase = m.M
	if v219 == int32(0) {
		v226 = v213
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v194)+40))
	v223 = F_bms_overlap(m, v212, v222)
	mBase = m.M
	v226 = v223 ^ int32(1)
	goto L52
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v194)+100))
	v228 = F_lappend(m, v179, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v230 = F_lappend(m, v181, v194)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L59
	}
L58:
	;
	v232 = v228
	v233 = v181
	goto L45
L59:
	;
	v232 = v179
	v233 = v230
	goto L45
L60:
	;
	goto L44
L61:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v264 = F_bms_union(m, v263, l5)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if base.Ui32(int32(5)) < base.Ui32(v268) {
		v280 = v267
		v281 = v264
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v249 != 0 {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	if int32(1)<<(uint(v268)%32)&int32(44) == int32(0) {
		v280 = v267
		v281 = v264
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v266)+228))
	v278 = F_bms_union(m, l5, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v280 = v277
	v281 = v278
	goto L63
L67:
	;
	if v407 == int32(0) {
		v486 = v251
		goto L37
	} else {
		goto L105
	}
L68:
	;
	v296 = v282
	v301 = v8
	goto L73
L69:
	;
	v282 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v282 < v283 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v407 = v8
	goto L67
L72:
	;
	goto L71
L73:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+v296<<(uint(int32(2))%32))))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+40)))
	if v314 != 0 {
		v387 = v301
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v407 = v387
	goto L67
L75:
	;
	v390 = v296 + int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v390 < v391 {
		v296 = v390
		v301 = v387
		goto L73
	} else {
		goto L104
	}
L76:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v313)+16))
	if v315 == int32(0) {
		v387 = v301
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	if v318 < int32(2) {
		v387 = v301
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v313)+36))
	v322 = int32(0)
	if v321 == v322 {
		v363 = v322
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v363 == int32(0) {
		v387 = v301
		goto L75
	} else {
		goto L93
	}
L80:
	;
	goto L79
L81:
	;
	if v281 == int32(0) {
		v363 = v322
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v331 < v332 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v334 = v331
	goto L85
L84:
	;
	v334 = v332
	goto L85
L85:
	;
	if v334 <= int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v337 = int32(1)
	goto L88
L87:
	;
	v337 = v334
	goto L88
L88:
	;
	v338 = int32(8)
	v343 = int32(0)
	goto L89
L89:
	;
	v350 = v343 << (uint(int32(2)) % 32)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v281+v338+v350)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+(v321+v338))))
	v355 = v352 & v354
	v357 = base.B2i32(v355 != int32(0))
	if v355 != 0 {
		v363 = v357
		goto L80
	} else {
		goto L91
	}
L90:
	;
	v363 = v357
	goto L80
L91:
	;
	v359 = v343 + int32(1)
	if v359 != v337 {
		v343 = v359
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+42)))
	if v369 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v384 = F_list_concat(m, v301, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L103
	}
L95:
	;
	v381 = F_generate_join_implied_equalities_broken(m, l0, v313, v281, l5, v280, v266)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L102
	}
L96:
	;
	v372 = F_generate_join_implied_equalities_normal(m, l0, v313, v264, l5, v267)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v377 = int32(0)
	if v369 == v377 {
		v383 = v377
		goto L94
	} else {
		goto L101
	}
L99:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+42)))
	if v374&int32(1) != 0 {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v383 = v372
	goto L94
L101:
	;
	goto L95
L102:
	;
	v383 = v381
	goto L94
L103:
	;
	v387 = v384
	goto L75
L104:
	;
	goto L74
L105:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v417 <= int32(0) {
		v486 = v251
		goto L37
	} else {
		goto L106
	}
L106:
	;
	v428 = int32(0)
	v434 = v251
	goto L107
L107:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443+v428<<(uint(int32(2))%32))))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v450 = int32(0)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447)+28))
	v452 = F_bms_is_subset(m, v451, v38)
	mBase = m.M
	if v452 == v450 {
		v463 = v450
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v486 = v468
	goto L37
L109:
	;
	if v463 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L109
L111:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v447)+28))
	v456 = F_bms_overlap(m, v449, v455)
	mBase = m.M
	if v456 == int32(0) {
		v463 = v450
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v447)+40))
	v460 = F_bms_overlap(m, v449, v459)
	mBase = m.M
	v463 = v460 ^ int32(1)
	goto L110
L113:
	;
	v466 = F_lappend(m, v434, v447)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	v468 = v434
	goto L115
L115:
	;
	v470 = v428 + int32(1)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v470 < v471 {
		v428 = v470
		v434 = v468
		goto L107
	} else {
		goto L117
	}
L116:
	;
	v468 = v466
	goto L115
L117:
	;
	goto L108
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v496
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v499 == int32(0) {
		v599 = v496
		goto L120
	} else {
		goto L121
	}
L119:
	;
	return v643
L120:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v613 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v614 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v615 = F_calc_joinrel_size_estimate(m, l0, l1, v611, v612, v613, v614, l4, v599)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L4
	} else {
		goto L139
	}
L121:
	;
	v502 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v503 <= v502 {
		v599 = v496
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v513 = v502
	goto L123
L123:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v513<<(uint(int32(2))%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	v534 = int32(0)
	v541 = base.B2i32(v533|l5 == v534)
	if v533 == v534 {
		v580 = v541
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v599 = v588
	goto L120
L125:
	;
	if v580 != 0 {
		v643 = v532
		goto L119
	} else {
		goto L137
	}
L126:
	;
	goto L125
L127:
	;
	if l5 == int32(0) {
		v580 = v541
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v547 != v548 {
		v580 = int32(0)
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v550 = int32(1)
	if v547 <= v550 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v553 = v550
	goto L132
L131:
	;
	v553 = v547
	goto L132
L132:
	;
	v554 = int32(8)
	v559 = int32(0)
	goto L133
L133:
	;
	v567 = v559 << (uint(int32(2)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v533+v554+v567)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567+(l5+v554))))
	v572 = base.B2i32(v569 == v571)
	if v571 != v569 {
		v580 = v572
		goto L126
	} else {
		goto L135
	}
L134:
	;
	v580 = v572
	goto L126
L135:
	;
	v575 = v559 + int32(1)
	if v575 != v553 {
		v559 = v575
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v585 = v513 + int32(1)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v585 < v586 {
		v513 = v585
		goto L123
	} else {
		goto L138
	}
L138:
	;
	goto L124
L139:
	;
	v617 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v615, v617) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v619 = v617
	goto L142
L141:
	;
	v619 = v615
	goto L142
L142:
	;
	v621 = F_palloc0(m, int32(24))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v621)+16)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v621)+8)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v621)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = int32(278)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v630 = F_lappend(m, v629, v621)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v630
	v643 = v621
	goto L119
}
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 float64
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v230 float64
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 float64
	_ = v271
	var v273 float64
	_ = v273
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v338 float64
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 float64
	_ = v349
	var v350 int32
	_ = v350
	var v352 float64
	_ = v352
	var v357 float64
	_ = v357
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 float64
	_ = v374
	var v384 float64
	_ = v384
	var v387 float64
	_ = v387
	var v391 float64
	_ = v391
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v462 float64
	_ = v462
	var v475 float64
	_ = v475
	v4 = float64(0)
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(1)
L2:
	;
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if int32(0) <= v74 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v74 = base.I32_ctz(v60) | v61<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v74 = int32(-2)
	goto L4
L7:
	;
	v27 = base.I32_div_s(int32(0), int32(32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 <= v27 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = l2 + int32(8)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v27<<(uint(int32(2))%32))))
	v38 = v35 & int32(-1)
	if v38 != 0 {
		v60 = v38
		v61 = v27
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v40 = v27 + int32(1)
	if v40 == v28 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v43 = v40
	goto L11
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31+v43<<(uint(int32(2))%32))))
	if v50 != 0 {
		v60 = v50
		v61 = v43
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L6
L13:
	;
	v52 = v43 + int32(1)
	if v52 != v28 {
		v43 = v52
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v80 = v4
	v84 = v74
	goto L18
L16:
	;
	v462 = v4
	goto L17
L17:
	;
	if base.F64_gt(v462, float64(0)) != 0 {
		goto L111
	} else {
		goto L112
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v90 <= v84 {
		v391 = v80
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v462 = v391
	goto L17
L20:
	;
	if l2 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v84<<(uint(int32(2))%32))))
	if v96 == int32(0) {
		v391 = v80
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v99 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	if v101 == v99 {
		v121 = v99
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v121 != 0 {
		v391 = v80
		goto L20
	} else {
		goto L33
	}
L24:
	;
	goto L23
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v105 = v104
	goto L26
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if base.Ui32(int32(2)) <= base.Ui32(v109-int32(301)) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v121 = int32(1)
	goto L24
L28:
	;
	if v109 != int32(290) {
		v121 = v99
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v105 = v108 + int32(72)
	goto L26
L30:
	;
	goto L27
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+72))
	if v116 != 0 {
		v121 = v99
		goto L24
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v96)+16))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v124 == int32(0) {
		v374 = v123
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.F64_gt(v80, v374) != 0 {
		goto L93
	} else {
		goto L94
	}
L35:
	;
	v127 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v128 <= v127 {
		v374 = v123
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v135 = v123
	v141 = v127
	goto L37
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v141<<(uint(int32(2))%32))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	if v149 != int32(4) {
		v357 = v135
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v374 = v357
	goto L34
L39:
	;
	v367 = v141 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v367 < v368 {
		v135 = v357
		v141 = v367
		goto L37
	} else {
		goto L92
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v153 = F_bms_is_member(m, l1, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return float64(0)
L42:
	;
	if v153 == int32(0) {
		v357 = v135
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v160 = F_bms_is_member(m, v84, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	if v160 == int32(0) {
		v357 = v135
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v164 = float64(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	if v165 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if int32(0) <= v222 {
		goto L57
	} else {
		goto L58
	}
L47:
	;
	v222 = base.I32_ctz(v208) | v209<<(uint(int32(5))%32)
	goto L46
L48:
	;
	v222 = int32(-2)
	goto L46
L49:
	;
	v175 = base.I32_div_s(int32(0), int32(32))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v176 <= v175 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v179 = v165 + int32(8)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v175<<(uint(int32(2))%32))))
	v186 = v183 & int32(-1)
	if v186 != 0 {
		v208 = v186
		v209 = v175
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v188 = v175 + int32(1)
	if v188 == v176 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v191 = v188
	goto L53
L53:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v179+v191<<(uint(int32(2))%32))))
	if v198 != 0 {
		v208 = v198
		v209 = v191
		goto L47
	} else {
		goto L55
	}
L54:
	;
	goto L48
L55:
	;
	v200 = v191 + int32(1)
	if v200 != v176 {
		v191 = v200
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v230 = v164
	v231 = v222
	goto L60
L58:
	;
	v338 = v164
	goto L59
L59:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
	v347 = int32(0)
	v349 = F_estimate_num_groups(m, l0, v346, v338, v347, v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L41
	} else {
		goto L88
	}
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v238 <= v231 {
		v273 = v230
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v338 = v273
	goto L59
L62:
	;
	if v165 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L63:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v231<<(uint(int32(2))%32))))
	if v244 == int32(0) {
		v273 = v230
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v247 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	if v249 == v247 {
		v269 = v247
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v269 != 0 {
		v273 = v230
		goto L62
	} else {
		goto L75
	}
L66:
	;
	goto L65
L67:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v253 = v252
	goto L68
L68:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if base.Ui32(int32(2)) <= base.Ui32(v257-int32(301)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v269 = int32(1)
	goto L66
L70:
	;
	if v257 != int32(290) {
		v269 = v247
		goto L66
	} else {
		goto L73
	}
L71:
	;
	v253 = v256 + int32(72)
	goto L68
L72:
	;
	goto L69
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+72))
	if v264 != 0 {
		v269 = v247
		goto L66
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v244)+16))
	v273 = base.F64_mul(v230, v271)
	goto L62
L76:
	;
	if int32(0) <= v330 {
		v230 = v273
		v231 = v330
		goto L60
	} else {
		goto L87
	}
L77:
	;
	v330 = base.I32_ctz(v316) | v317<<(uint(int32(5))%32)
	goto L76
L78:
	;
	v330 = int32(-2)
	goto L76
L79:
	;
	v281 = v231 + int32(1)
	v283 = base.I32_div_s(v281, int32(32))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v284 <= v283 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v287 = v165 + int32(8)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v283<<(uint(int32(2))%32))))
	v294 = v291 & (int32(-1) << (uint(v281) % 32))
	if v294 != 0 {
		v316 = v294
		v317 = v283
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v296 = v283 + int32(1)
	if v296 == v284 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v299 = v296
	goto L83
L83:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v287+v299<<(uint(int32(2))%32))))
	if v306 != 0 {
		v316 = v306
		v317 = v299
		goto L77
	} else {
		goto L85
	}
L84:
	;
	goto L78
L85:
	;
	v308 = v299 + int32(1)
	if v308 != v284 {
		v299 = v308
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	goto L61
L88:
	;
	if base.F64_gt(v135, v349) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v352 = v349
	goto L91
L90:
	;
	v352 = v135
	goto L91
L91:
	;
	v357 = v352
	goto L39
L92:
	;
	goto L38
L93:
	;
	v384 = v374
	goto L95
L94:
	;
	v384 = v80
	goto L95
L95:
	;
	if base.F64_eq(v80, float64(0)) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v387 = v374
	goto L98
L97:
	;
	v387 = v384
	goto L98
L98:
	;
	v391 = v387
	goto L20
L99:
	;
	if int32(0) <= v456 {
		v80 = v391
		v84 = v456
		goto L18
	} else {
		goto L110
	}
L100:
	;
	v456 = base.I32_ctz(v442) | v443<<(uint(int32(5))%32)
	goto L99
L101:
	;
	v456 = int32(-2)
	goto L99
L102:
	;
	v407 = v84 + int32(1)
	v409 = base.I32_div_s(v407, int32(32))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v410 <= v409 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v413 = l2 + int32(8)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413+v409<<(uint(int32(2))%32))))
	v420 = v417 & (int32(-1) << (uint(v407) % 32))
	if v420 != 0 {
		v442 = v420
		v443 = v409
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v422 = v409 + int32(1)
	if v422 == v410 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v425 = v422
	goto L106
L106:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v413+v425<<(uint(int32(2))%32))))
	if v432 != 0 {
		v442 = v432
		v443 = v425
		goto L100
	} else {
		goto L108
	}
L107:
	;
	goto L101
L108:
	;
	v434 = v425 + int32(1)
	if v434 != v410 {
		v425 = v434
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	goto L19
L111:
	;
	v475 = v462
	goto L113
L112:
	;
	v475 = float64(1)
	goto L113
L113:
	;
	return v475
}
func F_get_opcode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+100))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_parent_directory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = F_strlen(m, l0)
	mBase = m.M
	v9 = v6 + l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v13 = v9 - int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 == int32(47) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v19 = v13
	goto L10
L6:
	;
	if base.Ui32(l0) < base.Ui32(v13) {
		v9 = v13
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	if base.Ui32(l0) < base.Ui32(v19) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v31 = v19
	goto L16
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v25 != int32(47) {
		v19 = v19 - int32(1)
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	goto L14
L16:
	;
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l0 == v31 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v35 = v31 - int32(1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 == int32(47) {
		v31 = v35
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L20
L22:
	;
	v44 = l0 + base.B2i32(v5 == int32(47))
	goto L24
L23:
	;
	v44 = v31
	goto L24
L24:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
	goto L3
}
func F_get_password_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != int32(109) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(80)
	return v129
L2:
	;
	v124 = F_parse_scram_secret(m, l0, v6+int32(72), v6-int32(-64), v6+int32(68), v6+int32(76), v6+int32(32), v6)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v13 != int32(100) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v16 != int32(53) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = F_strlen(m, l0)
	mBase = m.M
	if v19 != int32(35) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v24 = l0 + int32(3)
	v25 = int32(355855)
	v29 = m.G0
	v31 = v29 - int32(32)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v32
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[369])))
	if v40 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v108 == int32(32) {
		v129 = int32(1)
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v108 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[370])))
	if v44 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v48 = v24
	goto L14
L12:
	;
	goto L13
L13:
	;
	v58 = v25
	v59 = v40
	goto L17
L14:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v54 == v40 {
		v48 = v48 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v108 = v48 - v24
	goto L7
L16:
	;
	goto L15
L17:
	;
	v66 = v31 + int32(base.Ui32(v59)>>(uint(int32(3))%32))&int32(28)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v67 | v68<<(uint(v59)%32)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v72 != 0 {
		v58 = v58 + v68
		v59 = v72
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v75 == int32(0) {
		v100 = v24
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v108 = v100 - v24
	goto L7
L21:
	;
	v79 = v24
	v80 = v75
	goto L22
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(base.Ui32(v80)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v88)>>(uint(v80)%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v100 = v96
	goto L20
L24:
	;
	v100 = v79
	goto L20
L25:
	;
	goto L26
L26:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v96 = v79 + int32(1)
	if v94 != 0 {
		v79 = v96
		v80 = v94
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	goto L2
L29:
	;
	return int32(0)
L30:
	;
	if v124 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v128 = int32(2)
	goto L33
L32:
	;
	v128 = int32(0)
	goto L33
L33:
	;
	v129 = v128
	goto L1
}
func F_get_ps_display(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return int32(790230)
}
func F_get_rolespec_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 {
	case 0:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = int32(0)
		v15 = F_GetSysCacheOid(m, int32(10), v11, v12, v12, v12)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if l1 != 0 {
				v74 = v15
				m.G0 = v7 + int32(48)
				return v74
			} else {
				if v15 != 0 {
					v74 = v15
					m.G0 = v7 + int32(48)
					return v74
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v11
							F_errmsg(m, int32(78201), v7+int32(16))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521065), int32(5562), int32(454908))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
	case 1, 2:
		v73 = *(*int32)(unsafe.Add(mBase, _consts[237]))
		v74 = v73
		m.G0 = v7 + int32(48)
		return v74
	case 3:
		v38 = *(*int32)(unsafe.Add(mBase, _consts[329]))
		v74 = v38
		m.G0 = v7 + int32(48)
		return v74
	case 4:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(513128)
				F_errmsg(m, int32(78201), v7+int32(32))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(521065), int32(5610), int32(455102))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v62
			F_errmsg_internal(m, int32(497757), v7)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(521065), int32(5615), int32(455102))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
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
func F_get_segment_by_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v8 = l0 + l1*int32(20)
	v10 = v8 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v11 != 0 {
		return v10
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+l1<<(uint(int32(2))%32))+32))
		if v16 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(480592), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524377), int32(1781), int32(28695))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = int32(4554292)
			v20 = *(*int32)(unsafe.Add(mBase, _consts[11]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _consts[11])) = v22
			v24 = F_dsm_attach(m, v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[11])) = v20
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(101590), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524377), int32(1788), int32(28695))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v24
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33 + int32(584)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v33 + int32(32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
					if base.Ui32(l1) <= base.Ui32(v42) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+648)) = l1
					}
					return v10
				}
			}
		}
	}
}
func F_get_sortgrouplist_exprs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v3 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v86
L2:
	;
	v15 = v3
	v16 = v3
	goto L7
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v10 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v86 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L18
	} else {
		goto L21
	}
L9:
	;
	goto L8
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v16<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v39 = int32(0)
	goto L12
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33+v39<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v32 != v48 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v54 = F_lappend(m, v15, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v51 = v39 + int32(1)
	if v51 != v24 {
		v39 = v51
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	goto L9
L18:
	;
	return int32(0)
L19:
	;
	v59 = v16 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v59 < v60 {
		v15 = v54
		v16 = v59
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v86 = v54
	goto L1
L21:
	;
	F_errmsg_internal(m, int32(79215), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(514759), int32(366), int32(400870))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_sortgroupref_tle(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v8
	goto L6
L5:
	;
	v14 = v11
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = v3
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if l0 != v25 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v24
L9:
	;
	v28 = v19 + int32(1)
	if v14 != v28 {
		v19 = v28
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	F_errmsg_internal(m, int32(79215), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(514759), int32(366), int32(400870))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_switched_clauses(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	v3 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v196
L2:
	;
	v19 = v3
	v22 = v3
	goto L7
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v11 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v196 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v19<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v31 = int32(0)
	if v30 == v31 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v196 = v181
	goto L1
L9:
	;
	v181 = F_lappend(m, v22, v176)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L27
	} else {
		goto L45
	}
L10:
	;
	if v84 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v84 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	if l1 == int32(0) {
		v75 = v31
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v84 = v75
	goto L10
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 < v40 {
		v75 = v31
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = int32(1)
	if v40 <= v43 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = v43
	goto L19
L18:
	;
	v46 = v40
	goto L19
L19:
	;
	v47 = int32(8)
	v52 = int32(0)
	goto L20
L20:
	;
	v59 = v52 << (uint(int32(2)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30+v47+v59)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+(l1+v47))))
	v66 = v61 & (v63 ^ int32(-1))
	v68 = base.B2i32(v66 == int32(0))
	if v66 != 0 {
		v75 = v68
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v75 = v68
	goto L14
L22:
	;
	v70 = v52 + int32(1)
	if v70 != v46 {
		v52 = v70
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v176 = v29
	v179 = int32(1)
	goto L9
L25:
	;
	goto L26
L26:
	;
	v89 = F_palloc0(m, int32(36))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(17)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v95
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v109 = F_list_copy(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v112
	v114 = m.G0
	v116 = v114 - int32(16)
	m.G0 = v116
	if v89 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v176 = v89
	v179 = v96
	goto L9
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L27
	} else {
		goto L42
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L27
	} else {
		goto L39
	}
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v120 != int32(17) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v123 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v126 != int32(2) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v130 = F_get_commutator(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	if v130 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v130
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v139
	m.G0 = v116 + int32(16)
	goto L30
L39:
	;
	F_errmsg_internal(m, int32(375305), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(516961), int32(2156), int32(217473))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v166
	F_errmsg_internal(m, int32(46559), v116)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(516961), int32(2162), int32(217473))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+120)) = uint8(v179)
	v185 = v19 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v185 < v186 {
		v19 = v185
		v22 = v181
		goto L7
	} else {
		goto L46
	}
L46:
	;
	goto L8
}
func F_get_typcollation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+144))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_typlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v13)+76)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.I32_extend16_s(v15)
			}
		}
	}
}
func F_get_typstorage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(112)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v13)+129)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v15)
			}
		}
	}
}
func F_get_values_def(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v8, int32(777131))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(1)
	v23 = int32(0)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v23<<(uint(int32(2))%32))))
	if v20&int32(1) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoString(m, v8, int32(778962))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_appendStringInfoChar(m, v8, int32(40))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v29 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_appendStringInfoChar(m, v8, int32(41))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v42 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v58 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v59 <= v58 {
		goto L13
	} else {
		goto L22
	}
L17:
	;
	F_get_rule_expr(m, v46, l1, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 != int32(6) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v53 = F_get_variable(m, v46, int32(1), l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L16
L22:
	;
	v68 = v58
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v68<<(uint(int32(2))%32))))
	F_appendStringInfoChar(m, v8, int32(44))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L13
L25:
	;
	if v73 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v89 = v68 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v89 < v90 {
		v68 = v89
		goto L23
	} else {
		goto L32
	}
L27:
	;
	F_get_rule_expr(m, v73, l1, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v79 != int32(6) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v83 = F_get_variable(m, v73, int32(1), l1)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	goto L26
L32:
	;
	goto L24
L33:
	;
	v104 = v23 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v104 < v105 {
		v20 = int32(0)
		v23 = v104
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L7
}
func F_get_view_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	return v62
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L18
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = int32(0)
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9+v11<<(uint(int32(2))%32))))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v23 = v11 + int32(1)
	if v6 != v23 {
		v11 = v23
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 == int32(1) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	F_errmsg_internal(m, int32(280183), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(518454), int32(2498), int32(15839))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	F_errmsg_internal(m, int32(33632), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(518454), int32(2504), int32(15839))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_visible_ENR_metadata(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v3 = int32(0)
	if l0 == v3 {
		v59 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v59
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(0) {
		v59 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		v59 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v16 = int32(0)
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14+v16<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v29 == int32(0) {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v59 = int32(0)
	goto L1
L7:
	;
	if v49-v48 == int32(0) {
		v59 = v24
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v28 != v29 {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = v25
	v34 = l1
	goto L11
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v37
		v49 = v38
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v48 = v37
	v49 = v38
	goto L8
L13:
	;
	v41 = int32(1)
	if v37 == v38 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v54 = v16 + int32(1)
	if v11 != v54 {
		v16 = v54
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L6
}
func F_getenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	v2 = int32(0)
	goto L5
L1:
	;
	if l0 == v92 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	goto L1
L3:
	;
	v82 = v77
	goto L21
L4:
	;
	v77 = v69
	goto L3
L5:
	;
	if l0&int32(3) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = l0
	goto L11
L9:
	;
	v29 = l0
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 != v38 {
		v69 = v29
		goto L4
	} else {
		goto L16
	}
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 == int32(0) {
		v92 = v16
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v29 = v26
	goto L10
L13:
	;
	if int32(61) == v21 {
		v92 = v16
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v26 = v16 + int32(1)
	if v26&int32(3) != 0 {
		v16 = v26
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v44 = v29
	v46 = v35
	goto L17
L17:
	;
	v50 = v46 ^ int32(1027423549)
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 != v53 {
		v69 = v44
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v77 = v59
	goto L3
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v59 = v44 + int32(4)
	v63 = int32(-2139062144)
	if (v57|(int32(16843008)-v57))&v63 == v63 {
		v44 = v59
		v46 = v57
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v84 == int32(0) {
		v92 = v82
		goto L2
	} else {
		goto L23
	}
L22:
	;
	v92 = v82
	goto L2
L23:
	;
	if v84 != int32(61) {
		v82 = v82 + int32(1)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	return int32(0)
L26:
	;
	goto L27
L27:
	;
	v106 = v92 - l0
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v106))))
	if v108 != 0 {
		v181 = v2
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return v181
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v110 == int32(0) {
		v181 = v2
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v113 == int32(0) {
		v181 = v2
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v117 = v110
	v118 = v113
	goto L32
L32:
	;
	if v106 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v181 = v168 + int32(1)
	goto L28
L34:
	;
	goto L33
L35:
	;
	if v164 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v164 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v126 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v127 = l0
	v128 = v118
	v129 = v106
	v130 = v126
	goto L43
L40:
	;
	v152 = v118
	v156 = int32(0)
	goto L41
L41:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v164 = v156 - v157
	goto L35
L42:
	;
	v152 = v147
	v156 = v149
	goto L41
L43:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v130 != v132 {
		v147 = v128
		v149 = v130
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v147 = v141
	v149 = int32(0)
	goto L42
L45:
	;
	if v132 == int32(0) {
		v147 = v128
		v149 = v130
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v137 = v129 - int32(1)
	if v137 == int32(0) {
		v147 = v128
		v149 = v130
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v140 = int32(1)
	v141 = v128 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	if v142 != 0 {
		v127 = v127 + v140
		v128 = v141
		v129 = v137
		v130 = v142
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v168 = v167 + v106
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169 == int32(61) {
		goto L34
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v173 != 0 {
		v117 = v117 + int32(4)
		v118 = v173
		goto L32
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v181 = v2
	goto L28
}
func F_getint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7))))
	v10 = v8 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v10) {
		return int32(0)
	} else {
		v16 = int32(0)
		v17 = v10
		v18 = v7
		for {
			if base.Ui32(v16) <= base.Ui32(int32(214748364)) {
				v26 = v16 * int32(10)
				if base.Ui32(v26^int32(2147483647)) < base.Ui32(v17) {
					v31 = int32(-1)
				} else {
					v31 = v17 + v26
				}
				v33 = v31
			} else {
				v33 = int32(-1)
			}
			v35 = v18 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35
			v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18)+1)))
			v39 = v37 - int32(48)
			if base.Ui32(v39) < base.Ui32(int32(10)) {
				v16 = v33
				v17 = v39
				v18 = v35
				continue
			} else {
				break
			}
			break
		}
		return v33
	}
}
func F_getrlimit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		v10 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	} else {
	}
	m.G0 = v8 + int32(16)
	return int32(0)
}
func F_gettoken_query_websearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v259 int32
	_ = v259
	v7 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v13 = l0 + int32(8)
	goto L3
L1:
	;
	return v259
L2:
	;
	v259 = int32(3)
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v21 - int32(1) {
	case 0, 2:
		goto L9
	case 1:
		goto L8
	default:
		goto L7
	}
L4:
	;
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + v237
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v237)
	goto L2
L5:
	;
	goto L4
L6:
	;
	v232 = F_pg_mblen_cstr(m, v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L22
	} else {
		goto L68
	}
L7:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v229 = v224
	goto L6
L8:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	switch v25 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v229 = v24
		goto L6
	default:
		goto L10
	case 24, 29, 31, 32, 51, 115:
		goto L11
	case 25:
		goto L12
	case 36:
		goto L5
	}
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v24
	goto L21
L11:
	;
	v64 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + v64
	goto L3
L12:
	;
	v29 = v24 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v37 = v32
	goto L13
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37 - v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L14
L16:
	;
	if v40 == int32(34) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v46 = v37 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v46
	v37 = v46
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v51 + int32(1)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v56 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v58 + int32(1)
	return v56
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v72 = int32(0)
	v74 = F_gettoken_tsvector(m, v71, l3, l2, v72, v72, v13)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v78
	return v78
L25:
	;
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v94 == int32(3) {
		v259 = int32(0)
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v85 != int32(447) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v88 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	return int32(1)
L31:
	;
	v98 = F_palloc0(m, int32(12))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v100 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = F_lcons(m, v98, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v103
	return int32(0)
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v118 = v108
	v119 = int32(223771)
	v120 = int32(2)
	goto L39
L37:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	switch v214 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v229 = v213
		goto L6
	default:
		goto L66
	case 24, 29, 31, 32, 51, 115:
		goto L67
	}
L38:
	;
	if v165 != 0 {
		goto L37
	} else {
		goto L54
	}
L39:
	;
	if v120 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v165 = int32(0)
	goto L38
L41:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v123 == v124 {
		v146 = v123
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	v148 = int32(1)
	if v146 != 0 {
		v118 = v118 + v148
		v119 = v119 + v148
		v120 = v120 - v148
		goto L39
	} else {
		goto L53
	}
L45:
	;
	if base.Ui32((v123-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v134 = v123 | int32(32)
	goto L48
L47:
	;
	v134 = v123
	goto L48
L48:
	;
	if base.Ui32((v124-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v143 = v124 | int32(32)
	goto L51
L50:
	;
	v143 = v124
	goto L51
L51:
	;
	if v134 == v143 {
		v146 = v134
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v165 = v134 - v143
	goto L38
L53:
	;
	goto L43
L54:
	;
	v167 = v108 + int32(2)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v168 == int32(0) {
		goto L37
	} else {
		goto L55
	}
L55:
	;
	if v168 == int32(45) {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	if v168 == int32(95) {
		goto L37
	} else {
		goto L57
	}
L57:
	;
	v175 = F_t_isalnum_cstr(m, v167)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	if v175 != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	v181 = v167
	goto L60
L60:
	;
	v184 = F_pg_mblen_cstr(m, v181)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
	} else {
		goto L62
	}
L61:
	;
	if v187 == int32(0) {
		goto L37
	} else {
		goto L65
	}
L62:
	;
	v186 = v184 + v181
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if base.Ui32(v187-int32(9)) < base.Ui32(int32(5)) {
		v181 = v186
		goto L60
	} else {
		goto L63
	}
L63:
	;
	if v187 == int32(32) {
		v181 = v186
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198 + int32(2)
	v202 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v202)
	return v202
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v222 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v222)
	goto L2
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v213 + int32(1)
	goto L3
L68:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v232 + v234
	goto L3
}
func F_ginbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v320 int32
	_ = v320
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
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
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 float64
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 float64
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v573 int32
	_ = v573
	var v585 int32
	_ = v585
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v681 int32
	_ = v681
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v820 int32
	_ = v820
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v868 int32
	_ = v868
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v992 int32
	_ = v992
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1121 int32
	_ = v1121
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 float64
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 float64
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1246 int32
	_ = v1246
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1469 int32
	_ = v1469
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1520 int64
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1652 int64
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1735 int32
	_ = v1735
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	v33 = m.G0
	v35 = v33 - int32(8480)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v44 = F_AllocSetContextCreateInternal(m, v39, int32(65364), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+2764)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v35)+2760)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v35)+2752)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[26]))) = v44
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[27]))) = v52
	v55 = v35 + int32(2768)
	F_initGinState(m, v55, v37)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v61 = F_palloc0(m, int32(40))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v71 = l1
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v71)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+2756)) = v71
	v75 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v79 = F_ReadBufferExtended(m, v37, v75, int32(1), v75, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	F_ginInsertCleanup(m, v55, base.B2i32(v64 != int32(4)), int32(0), int32(1), v61)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v71 = v61
	goto L6
L9:
	;
	v85 = int32(1)
	v88 = v79
	goto L10
L10:
	;
	if int32(0) <= v88 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v320 = v88
	goto L39
L12:
	;
	F_LockBuffer(m, v88, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v131 = v117 + v88<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(v88^int32(-1))<<(uint(int32(2))%32))))
	v131 = v130
	goto L12
L16:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v135)+6)))
	if v137&int32(2) == int32(0) {
		v233 = v131
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L11
L18:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v233)+24))
	v267 = v233 + v264&int32(32767)
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267)+2)))
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267))))
	F_UnlockReleaseBuffer(m, v88)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L37
	}
L19:
	;
	F_LockBuffer(m, v88, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_LockBuffer(m, v88, int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v85 != int32(1) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v150)+6)))
	if v152&int32(2) != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_LockBuffer(m, v88, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	if v88 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBuffer(m, v88, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199+(v88^int32(-1))<<(uint(int32(2))%32))))
	v207 = v201
	goto L27
L29:
	;
	goto L30
L30:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v207 = v203 + v88<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+16)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v211)+6)))
	if v213&int32(2) == int32(0) {
		v233 = v207
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_LockBuffer(m, v88, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_LockBuffer(m, v88, int32(2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+16)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v224)+6)))
	if v226&int32(2) != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	F_LockBuffer(m, v88, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v272 = int32(0)
	v275 = v268 | v269<<(uint(int32(16))%32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v278 = F_ReadBufferExtended(m, v37, v272, v275, v272, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v85 = v275
	v88 = v278
	goto L10
L39:
	;
	v346 = int32(0)
	v347 = base.B2i32(v346 <= v320)
	if v347 == v346 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[26])))
	F_MemoryContextDelete(m, v1816)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L306
	}
L41:
	;
	v853 = int32(0)
	F_vacuum_delay_point(m, v853)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L131
	}
L42:
	;
	F_UnlockReleaseBuffer(m, v320)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L130
	}
L43:
	;
	v783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+16)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v365+v783)))
	v801 = v785
	v806 = int32(0)
	goto L42
L44:
	;
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+12)))
	if base.Ui32(v366) < base.Ui32(int32(25)) {
		goto L43
	} else {
		goto L48
	}
L45:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351+(v320^int32(-1))<<(uint(int32(2))%32))))
	v365 = v357
	goto L44
L46:
	;
	goto L47
L47:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v365 = v359 + v320<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L48:
	;
	v374 = int32(base.Ui32(v366+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v374 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v383 = v365
	v389 = int32(1)
	v398 = int32(0)
	goto L51
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L127
	}
L51:
	;
	v412 = v389 & int32(65535)
	v416 = (v412 - int32(1)) << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v383+v416)+24))
	v421 = v383 + v418&int32(32767)
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+4)))
	if v422 == int32(0) {
		v666 = v383
		v681 = v398
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+16)))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v365+v699)))
	if v666 == v365 {
		v801 = v701
		v806 = v681
		goto L42
	} else {
		goto L108
	}
L53:
	;
	v695 = v389 + int32(1)
	if base.Ui32(v695&int32(65535)) <= base.Ui32(v374) {
		v383 = v666
		v389 = v695
		v398 = v681
		goto L51
	} else {
		goto L107
	}
L54:
	;
	if v422 == int32(65535) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v427 = int32(16)
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+2)))
	v433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421))))
	*(*int32)(unsafe.Add(mBase, uint32(v35+v427+v398<<(uint(int32(2))%32)))) = v432 | v433<<(uint(v427)%32)
	v666 = v383
	v681 = v398 + int32(1)
	goto L53
L56:
	;
	goto L57
L57:
	;
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421))))
	v442 = v440 << (uint(int32(16)) % 32)
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+2)))
	v444 = v442 | v443
	if int32(0) <= v442 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v585 == int32(0) {
		v666 = v383
		v681 = v398
		goto L53
	} else {
		goto L85
	}
L59:
	;
	F_pfree(m, v549)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L84
	}
L60:
	;
	v466 = int32(0)
	v470 = v466
	v480 = v466
	v487 = v466
	goto L66
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28]))) = v422
	v462 = v444 + v421
	v463 = v422
	goto L60
L62:
	;
	goto L63
L63:
	;
	v454 = F_ginPostingListDecode(m, v421+v444&int32(2147483647), v35+int32(8456))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28])))
	if int32(0) < v456 {
		v462 = v454
		v463 = v456
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v459 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28]))) = v459
	v549 = v454
	v551 = v459
	goto L59
L66:
	;
	v502 = v470 * int32(6)
	v503 = v462 + v502
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2764))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2760))
	v506 = m.T0[v505].(func(*base.Module, int32, int32) int32)(m, v503, v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28]))) = v533
	if int32(0) <= v442 {
		v585 = v532
		goto L58
	} else {
		goto L83
	}
L68:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2756))
	if v506 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v535 = v470 + int32(1)
	if v535 != v463 {
		v470 = v535
		v480 = v532
		v487 = v533
		goto L66
	} else {
		goto L82
	}
L70:
	;
	v509 = *(*float64)(unsafe.Add(mBase, uint32(v508)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v508)+16)) = base.F64_add(v509, float64(1))
	if v480 != 0 {
		v532 = v480
		v533 = v487
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v517 = *(*float64)(unsafe.Add(mBase, uint32(v508)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v508)+8)) = base.F64_add(v517, float64(1))
	if v480 != 0 {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v513 = F_palloc(m, v463*int32(6))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v502 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v532 = v513
	v533 = v487
	goto L69
L76:
	;
	v515 = F__emscripten_memcpy_bulkmem(m, v513, v462, v502)
	mBase = m.M
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	v523 = v480 + v487*int32(6)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = v524
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v503)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v523)+4)) = uint16(v526)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v532 = v480
	v533 = v487 + int32(1)
	goto L69
L82:
	;
	goto L67
L83:
	;
	v549 = v462
	v551 = v532
	goto L59
L84:
	;
	v585 = v551
	goto L58
L85:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28])))
	if v608 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v383 == v365 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v611 = int32(0)
	v624 = v611
	v625 = v611
	goto L86
L88:
	;
	goto L89
L89:
	;
	v615 = F_ginCompressPostingList(m, v585, v608, int32(2712), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v615)+6)))
	v624 = v615
	v625 = (v617+int32(1))&int32(131070) + int32(8)
	goto L86
L91:
	;
	v627 = F_PageGetTempPageCopy(m, v365)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	v634 = v383
	v635 = v421
	goto L93
L93:
	;
	v636 = F_gintuple_get_attrnum(m, v55, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v627+v416)+24))
	v634 = v627
	v635 = v627 + v630&int32(32767)
	goto L93
L95:
	;
	v640 = F_gintuple_get_key(m, v55, v635, v35+int32(8455))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[29]))))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28])))
	v645 = F_GinFormTuple(m, v55, v636, v640, v642, v624, v625, v643, int32(1))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v624 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_pfree(m, v624)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_PageIndexTupleDelete(m, v634, v412)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v645)+6)))
	v655 = F_PageAddItemExtended(m, v634, v645, v651&int32(8191), v412, int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v655 != v412 {
		goto L50
	} else {
		goto L104
	}
L104:
	;
	F_pfree(m, v645)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_pfree(m, v585)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v666 = v634
	v681 = v398
	goto L53
L107:
	;
	goto L52
L108:
	;
	if v666 == int32(0) {
		v801 = v701
		v806 = v681
		goto L42
	} else {
		goto L109
	}
L109:
	;
	v705 = int32(4548900)
	v707 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v707 + int32(1)
	F_PageRestoreTempPage(m, v666, v365)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_MarkBufferDirty(m, v320)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v347 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+48))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+118)))
	if v735 != int32(112) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v718+(v320^int32(-1))<<(uint(int32(2))%32))))
	v732 = v724
	goto L112
L114:
	;
	goto L115
L115:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v732 = v726 + v320<<(uint(int32(13))%32) + int32(-8192)
	goto L112
L116:
	;
	F_UnlockReleaseBuffer(m, v320)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L126
	}
L117:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v739 <= int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v733)+32))
	if v742 != 0 {
		goto L116
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v733)+40))
	if v743 != 0 {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	F_XLogRegisterBuffer(m, int32(0), v320, int32(9))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v752 = F_XLogInsert(m, int32(13), int32(64))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v732))) = base.I64_rotr(v752, int64(32))
	goto L116
L126:
	;
	v759 = int32(4548900)
	v761 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v761 - int32(1)
	v835 = v701
	v840 = v681
	goto L41
L127:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v770 + int32(4)
	F_errmsg_internal(m, int32(744667), v35)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(520577), int32(552), int32(427812))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v835 = v801
	v840 = v806
	goto L41
L131:
	;
	if v840 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v868 = v853
	goto L135
L133:
	;
	goto L134
L134:
	;
	if v835 != int32(-1) {
		goto L301
	} else {
		goto L302
	}
L135:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(16)+v868<<(uint(int32(2))%32))))
	v896 = v894
	goto L137
L136:
	;
	goto L134
L137:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v928 = int32(0)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[27])))
	v931 = F_ReadBufferExtended(m, v927, v928, v896, v928, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	v968 = int32(0)
	F_LockBuffer(m, v931, v968)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L149
	}
L139:
	;
	F_LockBuffer(m, v931, int32(1))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v931 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953)+16)))
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v954+v953)+6)))
	if v956&int32(2) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v939 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v939+(v931^int32(-1))<<(uint(int32(2))%32))))
	v953 = v945
	goto L141
L143:
	;
	goto L144
L144:
	;
	v947 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v953 = v947 + v931<<(uint(int32(13))%32) + int32(-8192)
	goto L141
L145:
	;
	v961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953)+34)))
	v962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953)+32)))
	F_UnlockReleaseBuffer(m, v931)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	goto L138
L148:
	;
	v896 = v961 | v962<<(uint(int32(16))%32)
	goto L137
L149:
	;
	F_LockBuffer(m, v931, int32(2))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v976 = v931
	v978 = v953
	v992 = v968
	goto L152
L151:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L299
	}
L152:
	;
	v1007 = int32(4554240)
	v1008 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[26])))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v1010
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v1014 = v35 + int32(2752)
	v1015 = int32(0)
	v1016 = m.G0
	v1018 = v1016 - int32(16)
	m.G0 = v1018
	if v976 < v1015 {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v1641 = int32(0)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[27])))
	v1644 = F_ReadBufferExtended(m, v1640, v1641, v894, v1641, v1643)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L1
	} else {
		goto L288
	}
L154:
	;
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v1008
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[26])))
	F_MemoryContextReset(m, v1581)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L1
	} else {
		goto L269
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L266
	}
L157:
	;
	m.G0 = v1018 + int32(16)
	goto L155
L158:
	;
	v1038 = F_disassembleLeaf(m, v1037)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L162
	}
L159:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1023+(v976^int32(-1))<<(uint(int32(2))%32))))
	v1037 = v1029
	goto L158
L160:
	;
	goto L161
L161:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1037 = v1031 + v976<<(uint(int32(13))%32) + int32(-8192)
	goto L158
L162:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1040 == int32(0) {
		goto L157
	} else {
		goto L163
	}
L163:
	;
	if v1040 == v1038 {
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v1048 = v1040
	v1054 = v1015
	goto L166
L165:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1246 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L166:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+24))
	if v1076 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v1054&int32(1) == int32(0) {
		goto L157
	} else {
		goto L208
	}
L168:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+20))
	v1082 = F_ginPostingListDecode(m, v1079, v1048+int32(28))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	v1085 = v1076
	goto L170
L170:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+20))
	if v1086 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+24)) = v1082
	v1085 = v1082
	goto L170
L172:
	;
	v1087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+6)))
	v1095 = (v1087+int32(1))&int32(131070) + int32(8)
	goto L174
L173:
	;
	v1095 = int32(8152)
	goto L174
L174:
	;
	v1097 = v1018 + int32(12)
	v1098 = int32(0)
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+28))
	if v1101 <= v1098 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+24))
	F_pfree(m, v1211)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L196
	}
L176:
	;
	v1104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1097))) = v1104
	v1210 = v1104
	goto L175
L177:
	;
	goto L178
L178:
	;
	v1121 = v1098
	v1129 = v1098
	v1131 = v1098
	goto L179
L179:
	;
	v1142 = v1131 * int32(6)
	v1143 = v1085 + v1142
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+12))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+8))
	v1146 = m.T0[v1145].(func(*base.Module, int32, int32) int32)(m, v1143, v1144)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1097))) = v1173
	v1210 = v1172
	goto L175
L181:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	if v1146 != 0 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v1175 = v1131 + int32(1)
	if v1175 != v1101 {
		v1121 = v1172
		v1129 = v1173
		v1131 = v1175
		goto L179
	} else {
		goto L195
	}
L183:
	;
	v1149 = *(*float64)(unsafe.Add(mBase, uint32(v1148)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1148)+16)) = base.F64_add(v1149, float64(1))
	if v1121 != 0 {
		v1172 = v1121
		v1173 = v1129
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v1157 = *(*float64)(unsafe.Add(mBase, uint32(v1148)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1148)+8)) = base.F64_add(v1157, float64(1))
	if v1121 != 0 {
		goto L192
	} else {
		goto L193
	}
L186:
	;
	v1153 = F_palloc(m, v1101*int32(6))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v1142 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v1172 = v1153
	v1173 = v1129
	goto L182
L189:
	;
	v1155 = F__emscripten_memcpy_bulkmem(m, v1153, v1085, v1142)
	mBase = m.M
	goto L191
L190:
	;
	goto L191
L191:
	;
	goto L188
L192:
	;
	v1163 = v1121 + v1129*int32(6)
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1143)))
	*(*int32)(unsafe.Add(mBase, uint32(v1163))) = v1164
	v1166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1143)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1163)+4)) = uint16(v1166)
	goto L194
L193:
	;
	goto L194
L194:
	;
	v1172 = v1121
	v1173 = v1129 + int32(1)
	goto L182
L195:
	;
	goto L180
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1048)+24)) = int64(0)
	if v1210 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+12))
	if int32(0) < v1216 {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	goto L199
L199:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1238 != v1038 {
		v1048 = v1238
		goto L166
	} else {
		goto L207
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+28)) = v1233
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1236 != v1038 {
		v1048 = v1236
		v1054 = int32(1)
		goto L166
	} else {
		goto L206
	}
L201:
	;
	v1221 = F_ginCompressPostingList(m, v1210, v1216, v1095, v1018+int32(8))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v1229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1048)+8)) = uint8(v1229)
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+20)) = int32(0)
	v1233 = v1216
	goto L200
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+20)) = v1221
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+8))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+12))
	if v1224 != v1225 {
		goto L156
	} else {
		goto L205
	}
L205:
	;
	v1227 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v1048)+8)) = uint8(v1227)
	v1233 = v1224
	goto L200
L206:
	;
	goto L165
L207:
	;
	goto L167
L208:
	;
	goto L165
L209:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+48))
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+118)))
	if v1343 != int32(112) {
		goto L223
	} else {
		goto L224
	}
L210:
	;
	if v1246 == v1038 {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v1255 = v1246
	v1258 = int32(0)
	goto L212
L212:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1255)+8)))
	v1286 = v1258 | base.B2i32(v1283 != int32(0))
	if v1283 == int32(1) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L209
L214:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+4))
	if v1308 != v1038 {
		v1255 = v1308
		v1258 = v1286
		goto L212
	} else {
		goto L222
	}
L215:
	;
	if v1286&int32(1) == int32(0) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+20))
	v1294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1293)+6)))
	v1300 = (v1294+int32(1))&int32(131070) + int32(8)
	v1301 = F_palloc(m, v1300)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+20))
	if v1300 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+20)) = v1305
	goto L214
L219:
	;
	v1304 = F__emscripten_memcpy_bulkmem(m, v1301, v1303, v1300)
	mBase = m.M
	v1305 = v1304
	goto L221
L220:
	;
	v1305 = v1301
	goto L221
L221:
	;
	goto L218
L222:
	;
	goto L213
L223:
	;
	v1354 = int32(4548900)
	v1356 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1356 + int32(1)
	if v976 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L224:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v1347 <= int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+32))
	if v1350 != 0 {
		goto L223
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	F_computeLeafRecompressWALData(m, v1038)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+40))
	if v1351 != 0 {
		goto L223
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	goto L223
L231:
	;
	v1378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+16)))
	v1379 = v1378 + v1377
	v1380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1379)+6)))
	v1382 = v1380 & int32(128)
	if v1382 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1363+(v976^int32(-1))<<(uint(int32(2))%32))))
	v1377 = v1369
	goto L231
L233:
	;
	goto L234
L234:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1377 = v1371 + v976<<(uint(int32(13))%32) + int32(-8192)
	goto L231
L235:
	;
	v1386 = v1380 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v1379)+6)) = uint16(v1386)
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+16)))
	v1390 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1377+v1388)+4)) = uint16(v1390)
	goto L237
L236:
	;
	goto L237
L237:
	;
	v1392 = int32(32)
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1393 == int32(0) {
		v1469 = v1392
		goto L238
	} else {
		goto L239
	}
L238:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1377)+12)) = uint16(v1469)
	F_MarkBufferDirty(m, v976)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L254
	}
L239:
	;
	if v1393 == v1038 {
		v1469 = v1392
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v1397 = int32(0)
	v1406 = v1393
	v1409 = base.B2i32(v1382 == v1397)
	v1412 = v1397
	v1415 = v1377 + int32(32)
	goto L241
L241:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406)+8)))
	v1437 = base.B2i32(v1434 != int32(0)) | v1409
	if v1434 != int32(1) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1469 = v1455 + int32(32)
	goto L238
L243:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+20))
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1440)+6)))
	v1442 = int32(1)
	v1447 = (v1441+v1442)&int32(131070) + int32(8)
	if v1437&v1442 != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	v1455 = v1412
	v1457 = v1415
	goto L245
L245:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	if v1458 != v1038 {
		v1406 = v1458
		v1409 = v1437
		v1412 = v1455
		v1415 = v1457
		goto L241
	} else {
		goto L253
	}
L246:
	;
	if v1447 != 0 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	goto L248
L248:
	;
	v1455 = v1447 + v1412
	v1457 = v1447 + v1415
	goto L245
L249:
	;
	goto L248
L250:
	;
	v1450 = F__emscripten_memcpy_bulkmem(m, v1415, v1440, v1447)
	mBase = m.M
	goto L252
L251:
	;
	goto L252
L252:
	;
	goto L249
L253:
	;
	goto L242
L254:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+48))
	v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497)+118)))
	if v1498 != int32(112) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1525 = int32(4548900)
	v1527 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1527 - int32(1)
	goto L157
L256:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v1502 <= int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+32))
	if v1505 != 0 {
		goto L255
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+40))
	if v1506 != 0 {
		goto L255
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	F_XLogRegisterBuffer(m, int32(0), v976, int32(8))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+24))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+28))
	F_XLogRegisterBufData(m, int32(0), v1514, v1515)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1520 = F_XLogInsert(m, int32(13), int32(144))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1037))) = base.I64_rotr(v1520, int64(32))
	goto L255
L266:
	;
	F_errmsg_internal(m, int32(81401), int32(0))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(522966), int32(782), int32(356149))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	v1584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v978)+16)))
	v1585 = v978 + v1584
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+6)))
	if v1586&int32(128) != 0 {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2752))
	v1613 = int32(0)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[27])))
	v1616 = F_ReadBufferExtended(m, v1612, v1613, v1610, v1613, v1615)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L283
	}
L271:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	F_UnlockReleaseBuffer(m, v976)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L1
	} else {
		goto L281
	}
L272:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	F_UnlockReleaseBuffer(m, v976)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L1
	} else {
		goto L278
	}
L273:
	;
	v1589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v978)+12)))
	if v1589 != int32(32) {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1585)+4)))
	if v1592 == int32(0) {
		goto L271
	} else {
		goto L277
	}
L276:
	;
	goto L271
L277:
	;
	goto L272
L278:
	;
	if v1595 != int32(-1) {
		v1610 = v1595
		v1611 = v992
		goto L270
	} else {
		goto L279
	}
L279:
	;
	if v992&int32(1) == int32(0) {
		goto L151
	} else {
		goto L280
	}
L280:
	;
	goto L154
L281:
	;
	if v1604 == int32(-1) {
		goto L154
	} else {
		goto L282
	}
L282:
	;
	v1610 = v1604
	v1611 = int32(1)
	goto L270
L283:
	;
	F_LockBuffer(m, v1616, int32(2))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	if v1616 < int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1624+(v1616^int32(-1))<<(uint(int32(2))%32))))
	v976 = v1616
	v978 = v1630
	v992 = v1611
	goto L152
L286:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v976 = v1616
	v978 = v1632 + v1616<<(uint(int32(13))%32) + int32(-8192)
	v992 = v1611
	goto L152
L288:
	;
	F_LockBufferForCleanup(m, v1644)
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1650 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[30]))) = v1650
	v1652 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[31]))) = v1652
	v1654 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[30]))) = uint8(v1654)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28]))) = v1652
	v1664 = F_ginScanToDelete(m, v35+int32(2752), v894, v1654, v35+int32(8456), v1650)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[28])))
	if v1666 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1668 = v1666
	goto L294
L292:
	;
	goto L293
L293:
	;
	F_UnlockReleaseBuffer(m, v1644)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L298
	}
L294:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1668)))
	F_pfree(m, v1668)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L296
	}
L295:
	;
	goto L293
L296:
	;
	if v1699 != 0 {
		v1668 = v1699
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	goto L151
L299:
	;
	v1772 = v868 + int32(1)
	if v1772 != v840 {
		v868 = v1772
		goto L135
	} else {
		goto L300
	}
L300:
	;
	goto L136
L301:
	;
	v1808 = int32(0)
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1811 = F_ReadBufferExtended(m, v37, v1808, v835, v1808, v1810)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	goto L40
L304:
	;
	F_LockBuffer(m, v1811, int32(2))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v320 = v1811
	goto L39
L306:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2756))
	m.G0 = v35 + int32(8480)
	return v1819
}
func F_gintuple_get_attrnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v10 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = F_index_getattr_1(m, l1, int32(1), v14, v7+int32(15))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			m.G0 = v7 + int32(16)
			return v21 & int32(65535)
		}
	} else {
		v21 = int32(1)
		m.G0 = v7 + int32(16)
		return v21 & int32(65535)
	}
}
func F_ginvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int64
	_ = v176
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(5712)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v21 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(5712)
	return v243
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if v25 != int32(4) {
		v243 = l1
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_initGinState(m, v18+int32(36), v20)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v37 = int32(1)
	F_ginInsertCleanup(m, v18+int32(36), int32(0), v37, v37, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v243 = l1
	goto L1
L9:
	;
	v44 = F_palloc0(m, int32(40))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v60 = l1
	goto L11
L11:
	;
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v61
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = float64(0)
	if base.F64_gt(v69, v70) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	F_initGinState(m, v18+int32(36), v20)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	F_ginInsertCleanup(m, v18+int32(36), base.B2i32(v53 != int32(4)), int32(0), int32(1), v44)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v60 = v44
	goto L11
L15:
	;
	v73 = v69
	goto L17
L16:
	;
	v73 = v70
	goto L17
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v60)+8)) = v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v75)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	if v77 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v96) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	F_LockRelationForExtension(m, v20, int32(7))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L25
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v80 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v84 = F_RelationGetNumberOfBlocksInFork(m, v20, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v96 = v84
	v97 = v3
	goto L18
L25:
	;
	v90 = F_RelationGetNumberOfBlocksInFork(m, v20, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_UnlockRelationForExtension(m, v20, int32(7))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v96 = v90
	v97 = int32(1)
	goto L18
L28:
	;
	v108 = int32(1)
	v110 = v3
	v112 = v3
	v113 = v3
	v114 = int64(0)
	goto L31
L29:
	;
	v213 = v3
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v96
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ginUpdateStats(m, v220, v18, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L60
	}
L31:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v213 = v195
	goto L30
L33:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v122 = F_ReadBufferExtended(m, v20, v119, v108, v119, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_LockBuffer(m, v122, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if v122 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L58
	}
L37:
	;
	F_RecordFreeIndexPage(m, v20, v108)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L57
	}
L38:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+14)))
	if v145 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+(v122^int32(-1))<<(uint(int32(2))%32))))
	v144 = v136
	goto L38
L40:
	;
	goto L41
L41:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v144 = v138 + v122<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+16)))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+v148)+6)))
	if v150&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v153 == int32(0) {
		goto L37
	} else {
		goto L46
	}
L44:
	;
	v161 = v150
	goto L45
L45:
	;
	if v161&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v156 = F_GlobalVisCheckRemovableXid(m, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if v156 != 0 {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+16)))
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+v158)+6)))
	v161 = v160
	goto L45
L49:
	;
	v165 = v112 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v165
	v195 = v110
	v196 = v165
	v197 = v113
	v198 = v114
	goto L36
L50:
	;
	goto L51
L51:
	;
	if v161&int32(16) != 0 {
		v195 = v110
		v196 = v112
		v197 = v113
		v198 = v114
		goto L36
	} else {
		goto L52
	}
L52:
	;
	v170 = v113 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v170
	if v161&int32(2) == int32(0) {
		v195 = v110
		v196 = v112
		v197 = v170
		v198 = v114
		goto L36
	} else {
		goto L53
	}
L53:
	;
	v176 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v144)+12)))
	if base.Ui64(int64(25)) <= base.Ui64(v176) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v186 = int64(base.Ui64(v176+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
	goto L56
L55:
	;
	v186 = int64(0)
	goto L56
L56:
	;
	v187 = v114 + v186
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v187
	v195 = v110
	v196 = v112
	v197 = v170
	v198 = v187
	goto L36
L57:
	;
	v195 = v110 + int32(1)
	v196 = v112
	v197 = v113
	v198 = v114
	goto L36
L58:
	;
	v202 = v108 + int32(1)
	if v202 != v96 {
		v108 = v202
		v110 = v195
		v112 = v196
		v113 = v197
		v114 = v198
		goto L31
	} else {
		goto L59
	}
L59:
	;
	goto L32
L60:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_FreeSpaceMapVacuum(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v213
	if v97 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_LockRelationForExtension(m, v20, int32(7))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v239 = F_RelationGetNumberOfBlocksInFork(m, v20, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L68
	}
L65:
	;
	v232 = F_RelationGetNumberOfBlocksInFork(m, v20, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v232
	F_UnlockRelationForExtension(m, v20, int32(7))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v243 = v60
	goto L1
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v239
	v243 = v60
	goto L1
}
func F_gistadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v13 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v14 <= v13 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v13
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	v29 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+24)) = uint16(v29)
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 < v33 {
		v18 = v32
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(16)
	return
L8:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v44 <= v43 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = v43
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v48<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui32(int32(12)) < base.Ui32(v58) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	v95 = v48 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v95 < v96 {
		v48 = v95
		goto L10
	} else {
		goto L23
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = l0
	v92 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)) = uint16(v92)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v62 = int32(1) << (uint(v58) % 32)
	if v62&int32(7960) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v62&int32(230) == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+24)) = uint8(v69)
	goto L12
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(81641)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v79
	F_errmsg(m, int32(206490), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(522105), int32(349), int32(144015))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L11
}
func F_gistdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v686 int32
	_ = v686
	v6 = l5
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v7
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+74)) = uint16(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+40)) = uint8(v6)
	v38 = v19 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v38
	v46 = v38
	v48 = v7
	v49 = v7
	goto L1
L1:
	;
	if v49&int32(1) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
	if v117 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v105 = v48
	v107 = v46
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v48&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	F_LockBuffer(m, v66, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	F_ReleaseBuffer(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v73
	v75 = int32(0)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+24)))
	if v76 != int32(1) {
		v105 = v75
		v107 = v73
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v85 = v73
	goto L14
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	F_ReleaseBuffer(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v98
	v105 = v75
	v107 = v98
	goto L3
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)+28))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)))
	if v99 != 0 {
		v85 = v98
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v122 = F_ReadBuffer(m, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v126 = v105 & int32(1)
	if v126 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v122
	goto L20
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v129, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v137 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_gistcheckpage(m, v133, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v155
	if v126 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141+(v137^int32(-1))<<(uint(int32(2))%32))))
	v155 = v147
	goto L27
L29:
	;
	goto L30
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v155 = v149 + v137<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v671
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+24)))
	v46 = v671
	v48 = v673
	v49 = v686
	goto L1
L32:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v448 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L10
	} else {
		goto L108
	}
L33:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+24)))
	v46 = v107
	v48 = int32(1)
	v49 = v441
	goto L1
L34:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v198 != 0 {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v157 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v155)+4)))
	v158 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v155))))
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v157 | v158<<(uint(int64(32))%64)
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	v164 = v155 + v163
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+12)))
	if v165&int32(8) == int32(0) {
		v195 = v155
		v196 = v165
		v197 = v164
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v170 = F_BufferGetLSNAtomic(m, v137)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	goto L32
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+16)))
	v175 = v173 + v174
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+12)))
	if v176&int32(8) == int32(0) {
		v195 = v173
		v196 = v176
		v197 = v175
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v181, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v185, int32(2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+16)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v190)+12)))
	if v192&int32(8) != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L103
	}
L45:
	;
	v382 = F_gistinserttuple(m, v19+int32(28), v107, l3, l1, int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L10
	} else {
		goto L97
	}
L46:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v372
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+24)))
	v46 = v372
	v48 = int32(0)
	v49 = v375
	goto L1
L47:
	;
	if v196&int32(1) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_UnlockReleaseBuffer(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L55
	}
L49:
	;
	if v196&int32(2) != 0 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v196&int32(2) == int32(0) {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v201)+16))
	v203 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v197)+4)))
	v204 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v197))))
	if base.Ui64(v202) < base.Ui64(v203|v204<<(uint(int64(32))%64)) {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	goto L48
L55:
	;
	goto L46
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v221 = F_gistchoose(m, v220, v195, l1, l3)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v126 != 0 {
		goto L45
	} else {
		goto L80
	}
L59:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223+v221<<(uint(int32(2))%32))+20))
	v230 = v223 + v227&int32(32767)
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)))
	if v231 == int32(65534) {
		goto L44
	} else {
		goto L60
	}
L60:
	;
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230))))
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+2)))
	v236 = F_gistgetadjusted(m, v220, v230, l1, l3)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L62
	}
L61:
	;
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v292, v291)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L10
	} else {
		goto L78
	}
L62:
	;
	if v236 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if v126 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v242, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v280 = F_gistinserttuple(m, v19+int32(28), v107, l3, v236, v221)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L10
	} else {
		goto L74
	}
L67:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v246, int32(2))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v250 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v268
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
	v271 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v268)+4)))
	v272 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v268))))
	if v270 != v271|v272<<(uint(int64(32))%64) {
		goto L33
	} else {
		goto L73
	}
L70:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v254+(v250^int32(-1))<<(uint(int32(2))%32))))
	v268 = v260
	goto L69
L71:
	;
	goto L72
L72:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v268 = v262 + v250<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	goto L66
L74:
	;
	if v280 == int32(0) {
		goto L61
	} else {
		goto L75
	}
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v284 == int32(0) {
		goto L33
	} else {
		goto L76
	}
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_UnlockReleaseBuffer(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	goto L46
L78:
	;
	v297 = F_palloc0(m, int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v234<<(uint(int32(16))%32) | v235
	*(*uint16)(unsafe.Add(mBase, uint32(v297)+26)) = uint16(v221)
	v671 = v297
	v673 = v291
	goto L31
L80:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v305, int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v309, int32(2))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v313 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v331
	v333 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v331)+4)))
	v334 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v331))))
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v333 | v334<<(uint(int64(32))%64)
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+16)))
	v340 = v331 + v339
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+12)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v342 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317+(v313^int32(-1))<<(uint(int32(2))%32))))
	v331 = v323
	goto L83
L85:
	;
	goto L86
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v331 = v325 + v313<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	if v341&int32(1) != 0 {
		goto L45
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v341&int32(8) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v347 = int32(0)
	F_LockBuffer(m, v313, v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+24)))
	v46 = v107
	v48 = v347
	v49 = v351
	goto L1
L92:
	;
	F_UnlockReleaseBuffer(m, v313)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L10
	} else {
		goto L96
	}
L93:
	;
	if v341&int32(2) != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v356)+16))
	v358 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v340)+4)))
	v359 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v340))))
	if base.Ui64(v358|v359<<(uint(int64(32))%64)) <= base.Ui64(v357) {
		goto L45
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L46
L97:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_LockBuffer(m, v384, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v394 = v107
	goto L99
L99:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	F_ReleaseBuffer(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L10
	} else {
		goto L101
	}
L100:
	;
	m.G0 = v19 + int32(80)
	return
L101:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v394)+28))
	if v407 != 0 {
		v394 = v407
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v415 + int32(4)
	F_errmsg(m, int32(455690), v19)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	F_errdetail(m, int32(688742), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errhint(m, int32(608465), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(514917), int32(768), int32(87639))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	if v448 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+48))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v451 + int32(4)
	F_errmsg(m, int32(52197), v19+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L10
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	v480 = v468
	v483 = int32(0)
	goto L114
L112:
	;
	F_errfinish(m, int32(514917), int32(1209), int32(107937))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v487 = F_palloc(m, int32(8))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L116
	}
L115:
	;
	v659 = int32(0)
	F_gistfinishsplit(m, v19+int32(28), v445, l3, v645, v659)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L10
	} else {
		goto L152
	}
L116:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v480 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if v480 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L118:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v445)+28))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	F_LockBuffer(m, v579, int32(2))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L10
	} else {
		goto L138
	}
L119:
	;
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v507)+12)))
	if base.Ui32(v508) < base.Ui32(int32(25)) {
		goto L118
	} else {
		goto L123
	}
L120:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v493+(v480^int32(-1))<<(uint(int32(2))%32))))
	v507 = v499
	goto L119
L121:
	;
	goto L122
L122:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v507 = v501 + v480<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	v512 = v508 + int32(262120)
	if v512&int32(262140) == int32(0) {
		goto L118
	} else {
		goto L124
	}
L124:
	;
	v527 = int32(1)
	v529 = int32(0)
	goto L125
L125:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v527<<(uint(int32(2))%32)+(v507+int32(24))-int32(4))))
	v549 = v507 + v546&int32(32767)
	if v529 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	if v558 != 0 {
		v606 = v558
		goto L117
	} else {
		goto L137
	}
L127:
	;
	if v527 != int32(base.Ui32(v512)>>(uint(int32(2))%32))&int32(65535) {
		v527 = v527 + int32(1)
		v529 = v558
		goto L125
	} else {
		goto L136
	}
L128:
	;
	v552 = F_CopyIndexTuple(m, v549)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L10
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v554 = F_gistgetadjusted(m, v489, v529, v549, l3)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L132
	}
L131:
	;
	v558 = v552
	goto L127
L132:
	;
	if v554 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v556 = v554
	goto L135
L134:
	;
	v556 = v529
	goto L135
L135:
	;
	v558 = v556
	goto L127
L136:
	;
	goto L126
L137:
	;
	goto L118
L138:
	;
	F_gistFindCorrectParent(m, v489, v445)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v445)+28))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+8))
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445)+26)))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v586+v587<<(uint(int32(2))%32))+20))
	v595 = F_CopyIndexTuple(m, v586+v591&int32(32767))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v445)+28))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	F_LockBuffer(m, v598, int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v606 = v595
	goto L117
L142:
	;
	v637 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v606)+4)) = uint16(v637)
	*(*uint16)(unsafe.Add(mBase, uint32(v606)+2)) = uint16(v636)
	v641 = int32(base.Ui32(v636) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v606))) = uint16(v641)
	*(*int32)(unsafe.Add(mBase, uint32(v487)+4)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v480
	v645 = F_lappend(m, v483, v487)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L10
	} else {
		goto L146
	}
L143:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v621+(v480^int32(-1))<<(uint(int32(6))%32))+16))
	v636 = v627
	goto L142
L144:
	;
	goto L145
L145:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v629+v480<<(uint(int32(6))%32)+int32(-64))+16))
	v636 = v635
	goto L142
L146:
	;
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v507)+16)))
	v648 = v507 + v647
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648)+12)))
	if v649&int32(8) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	v654 = F_ReadBuffer(m, v652, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	goto L115
L150:
	;
	F_LockBuffer(m, v654, int32(2))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	v480 = v654
	v483 = v645
	goto L114
L152:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_UnlockReleaseBuffer(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	v671 = v668
	v673 = v659
	goto L31
}
func F_gistextractpage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v11) {
		v19 = int32(base.Ui32(v11+int32(262120)) >> (uint(int32(2)) % 32))
	} else {
		v19 = int32(0)
	}
	v21 = v19 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
	v25 = F_palloc(m, v21<<(uint(int32(2))%32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if v21 == int32(0) {
		} else {
			v32 = v25 - int32(4)
			v33 = int32(1)
			v34 = int32(2)
			v38 = (v19 + v33) & int32(65535)
			if base.Ui32(v38) <= base.Ui32(v34) {
				v41 = v34
			} else {
				v41 = v38
			}
			v42 = int32(1)
			v43 = v41 - v42
			v47 = l0 + int32(24)
			if base.Ui32(int32(3)) <= base.Ui32(v38) {
				v55 = v33
				v56 = int32(0)
				for {
					v63 = int32(2)
					v64 = v55 << (uint(v63) % 32)
					v66 = v64 + v47
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v66-int32(4))))
					v70 = int32(32767)
					*(*int32)(unsafe.Add(mBase, uint32(v32+v64))) = l0 + v69&v70
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					*(*int32)(unsafe.Add(mBase, uint32(v64+v25))) = l0 + v75&v70
					v81 = v55 + v63
					v83 = v56 + v63
					if v83 != v43&int32(-2) {
						v55 = v81
						v56 = v83
						continue
					} else {
						break
					}
					break
				}
				v87 = v81
			} else {
				v87 = v33
			}
			if v43&v42 == int32(0) {
			} else {
				v98 = v87 << (uint(int32(2)) % 32)
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v98+v47-int32(4))))
				*(*int32)(unsafe.Add(mBase, uint32(v32+v98))) = l0 + v103&int32(32767)
			}
		}
		return v25
	}
}
func F_gistunionsubkey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	v4 = int32(0)
	v18 = l1 - int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+352))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v28 = F_palloc(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		if v25 <= int32(0) {
			v129 = v4
		} else {
			v32 = int32(1)
			if v25 == v32 {
				v95 = v4
				v96 = int32(0)
			} else {
				v43 = v4
				v44 = int32(0)
				v51 = v4
				for {
					v58 = v23 + v44<<(uint(int32(1))%32)
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58))))
					if v24 != 0 {
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v59))))
						if v61 != 0 {
							v72 = v43
						} else {
							v62 = int32(2)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v18+v59<<(uint(v62)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v28+v43<<(uint(v62)%32)))) = v68
							v72 = v43 + int32(1)
						}
					} else {
						v62 = int32(2)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v18+v59<<(uint(v62)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v43<<(uint(v62)%32)))) = v68
						v72 = v43 + int32(1)
					}
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+2)))
					if v24 != 0 {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v73))))
						if v75 != 0 {
							v86 = v72
						} else {
							v76 = int32(2)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v18+v73<<(uint(v76)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v28+v72<<(uint(v76)%32)))) = v82
							v86 = v72 + int32(1)
						}
					} else {
						v76 = int32(2)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v18+v73<<(uint(v76)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v72<<(uint(v76)%32)))) = v82
						v86 = v72 + int32(1)
					}
					v87 = int32(2)
					v88 = v44 + v87
					v90 = v51 + v87
					if v90 != v25&int32(2147483646) {
						v43 = v86
						v44 = v88
						v51 = v90
						continue
					} else {
						break
					}
					break
				}
				v95 = v86
				v96 = v88
			}
			if v25&v32 == int32(0) {
				v129 = v95
			} else {
				v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v96<<(uint(int32(1))%32)))))
				if v24 != 0 {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v113))))
					if v115 != 0 {
						v129 = v95
					} else {
						v116 = int32(2)
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v18+v113<<(uint(v116)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v95<<(uint(v116)%32)))) = v122
						v129 = v95 + int32(1)
					}
				} else {
					v116 = int32(2)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v18+v113<<(uint(v116)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v95<<(uint(v116)%32)))) = v122
					v129 = v95 + int32(1)
				}
			}
		}
		F_gistMakeUnionItVec(m, l0, v28, v129, l2+int32(32), l2+int32(160))
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return
		} else {
			F_pfree(m, v28)
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return
			} else {
				v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v150 = F_palloc(m, v147<<(uint(int32(2))%32))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					if v147 <= int32(0) {
						v254 = v4
					} else {
						v154 = int32(1)
						if v147 == v154 {
							v218 = int32(0)
							v220 = v4
						} else {
							v161 = int32(0)
							v166 = v161
							v168 = v4
							v170 = v161
							for {
								v181 = v146 + v166<<(uint(int32(1))%32)
								v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181))))
								if v24 != 0 {
									v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v182))))
									if v184 != 0 {
										v195 = v168
									} else {
										v185 = int32(2)
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v18+v182<<(uint(v185)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v150+v168<<(uint(v185)%32)))) = v191
										v195 = v168 + int32(1)
									}
								} else {
									v185 = int32(2)
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v18+v182<<(uint(v185)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v168<<(uint(v185)%32)))) = v191
									v195 = v168 + int32(1)
								}
								v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+2)))
								if v24 != 0 {
									v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v196))))
									if v198 != 0 {
										v209 = v195
									} else {
										v199 = int32(2)
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v18+v196<<(uint(v199)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v150+v195<<(uint(v199)%32)))) = v205
										v209 = v195 + int32(1)
									}
								} else {
									v199 = int32(2)
									v205 = *(*int32)(unsafe.Add(mBase, uint32(v18+v196<<(uint(v199)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v195<<(uint(v199)%32)))) = v205
									v209 = v195 + int32(1)
								}
								v210 = int32(2)
								v211 = v166 + v210
								v213 = v170 + v210
								if v213 != v147&int32(2147483646) {
									v166 = v211
									v168 = v209
									v170 = v213
									continue
								} else {
									break
								}
								break
							}
							v218 = v211
							v220 = v209
						}
						if v147&v154 == int32(0) {
							v254 = v220
						} else {
							v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146+v218<<(uint(int32(1))%32)))))
							if v24 != 0 {
								v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v236))))
								if v238 != 0 {
									v254 = v220
								} else {
									v239 = int32(2)
									v245 = *(*int32)(unsafe.Add(mBase, uint32(v18+v236<<(uint(v239)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v220<<(uint(v239)%32)))) = v245
									v254 = v220 + int32(1)
								}
							} else {
								v239 = int32(2)
								v245 = *(*int32)(unsafe.Add(mBase, uint32(v18+v236<<(uint(v239)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v220<<(uint(v239)%32)))) = v245
								v254 = v220 + int32(1)
							}
						}
					}
					F_gistMakeUnionItVec(m, l0, v150, v254, l2+int32(192), l2+int32(320))
					mBase = m.M
					v270 = m.ExcPending
					if v270 != 0 {
						return
					} else {
						F_pfree(m, v150)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_gseg_picksplit_item_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F32_ne(v6, v7)
	}
	return v10
}
func F_gseg_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = F_DirectFunctionCall2Coll(m, int32(6687), int32(0), v5, v4)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(base.B2i32(v9 != int32(0)))
		return v6
	}
}
