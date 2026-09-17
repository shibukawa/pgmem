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
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[0]))
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[1]))
		v32 = int32(3)
		v34 = int32(256)
		v35 = v28*v30<<(uint(v32)%32) + v34
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[2]))
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
			v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[3]))
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
		v50 = int32(_a_F_GetAccessStrategy_0)
		v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[3]))
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
		v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategy[3]))
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
			F_errmsg_internal(m, int32(_a_F_GetAccessStrategy_1), v8)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetAccessStrategy_2), int32(611), int32(_a_F_GetAccessStrategy_3))
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
	var v89 int32
	_ = v89
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
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(1088)
	m.G0 = v14
	v16 = int32(_a_F_GetConfFilesInDir_0)
	v20 = m.G0
	v22 = v20 - int32(32)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v23
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetConfFilesInDir[0])))
	if v31 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(1088)
	return v306
L2:
	;
	v100 = F_strlen(m, l0)
	mBase = m.M
	if v99 == v100 {
		goto L21
	} else {
		goto L22
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetConfFilesInDir[1])))
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
		v89 = l0
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v99 = v89 - l0
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
		v89 = v70
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v89 = v87
	goto L15
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v87 = v70 + int32(1)
	if v85 != 0 {
		v70 = v87
		v71 = v85
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v103 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v121 = F_AbsoluteConfigLocation(m, l0, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L33
	}
L24:
	;
	return int32(0)
L25:
	;
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_GetConfFilesInDir_1)
	v306 = v6
	goto L1
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_GetConfFilesInDir_2), v14)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_GetConfFilesInDir_3), int32(89), int32(_a_F_GetConfFilesInDir_4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	F_pfree(m, v121)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L24
	} else {
		goto L81
	}
L33:
	;
	v123 = F_AllocateDir(m, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	if v123 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v128 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L24
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v151 = F_palloc(m, int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L24
	} else {
		goto L46
	}
L38:
	;
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v121
	v147 = F_psprintf(m, int32(_a_F_GetConfFilesInDir_5), v14+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v121
	F_errmsg(m, int32(_a_F_GetConfFilesInDir_6), v14+int32(32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_GetConfFilesInDir_3), int32(101), int32(_a_F_GetConfFilesInDir_4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L24
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v147
	v293 = v6
	goto L32
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v155 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L24
	} else {
		goto L49
	}
L47:
	;
	F_FreeDir(m, v123)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L24
	} else {
		goto L80
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 - int32(-64)
	v268 = F_psprintf(m, int32(_a_F_GetConfFilesInDir_7), v14+int32(48))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L24
	} else {
		goto L78
	}
L49:
	;
	if v155 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v158 = v155
	v164 = v151
	v167 = int32(32)
	goto L53
L51:
	;
	v250 = v151
	goto L52
L52:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v255 <= int32(0) {
		v280 = v250
		goto L47
	} else {
		goto L76
	}
L53:
	;
	v170 = v158 + int32(19)
	v171 = F_strlen(m, v170)
	mBase = m.M
	if base.Ui32(v171) < base.Ui32(int32(6)) {
		v239 = v164
		v241 = v167
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v250 = v239
	goto L52
L55:
	;
	v242 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L24
	} else {
		goto L74
	}
L56:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 == int32(46) {
		v239 = v164
		v241 = v167
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v179 = v171 + v170 - int32(5)
	v180 = int32(_a_F_GetConfFilesInDir_8)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetConfFilesInDir[2])))
	if base.B2i32(v183 == int32(0))|base.B2i32(v183 != v186) != 0 {
		v204 = v183
		v205 = v186
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v204-v205 != 0 {
		v239 = v164
		v241 = v167
		goto L55
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	v189 = v179
	v190 = v180
	goto L61
L61:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v194 == int32(0) {
		v204 = v194
		v205 = v193
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v204 = v194
	v205 = v193
	goto L59
L63:
	;
	v197 = int32(1)
	if v194 == v193 {
		v189 = v189 + v197
		v190 = v190 + v197
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v208 = v14 - int32(-64)
	F_join_path_components(m, v208, v121, v170)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L24
	} else {
		goto L66
	}
L66:
	;
	F_canonicalize_path_enc(m, v208)
	mBase = m.M
	v213 = F_get_dirent_type(m, v208, v158, int32(1), l2)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L24
	} else {
		goto L68
	}
L67:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v167 <= v215 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	switch v213 {
	case 0:
		goto L48
	default:
		goto L67
	case 3:
		v239 = v164
		v241 = v167
		goto L55
	}
L69:
	;
	v218 = v167 + int32(32)
	v221 = F_repalloc(m, v164, v218<<(uint(int32(2))%32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L24
	} else {
		goto L72
	}
L70:
	;
	v223 = v164
	v224 = v167
	goto L71
L71:
	;
	v227 = F_pstrdup(m, v14-int32(-64))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L24
	} else {
		goto L73
	}
L72:
	;
	v223 = v221
	v224 = v218
	goto L71
L73:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v223+v229<<(uint(int32(2))%32)))) = v227
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v234 + int32(1)
	v239 = v223
	v241 = v224
	goto L55
L74:
	;
	if v242 != 0 {
		v158 = v242
		v164 = v239
		v167 = v241
		goto L53
	} else {
		goto L75
	}
L75:
	;
	goto L54
L76:
	;
	F_pg_qsort(m, v250, v255, int32(4), int32(1170))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v280 = v250
	goto L47
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v268
	F_pfree(m, v164)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v280 = int32(0)
	goto L47
L80:
	;
	v293 = v280
	goto L32
L81:
	;
	v306 = v293
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
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetFdwRoutine[0])))
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
					F_errmsg_internal(m, int32(_a_F_GetFdwRoutine_0), v6)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(345), int32(_a_F_GetFdwRoutine_2))
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
						F_errmsg_internal(m, int32(_a_F_GetFdwRoutine_0), v6)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(345), int32(_a_F_GetFdwRoutine_2))
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
				F_errmsg(m, int32(_a_F_GetFdwRoutine_3), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(337), int32(_a_F_GetFdwRoutine_2))
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
	v4 = *(*int64)(unsafe.Add(mBase, _c_F_GetFullPageWriteInfo[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetFullPageWriteInfo[1])))
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+440)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
		F_s_lock(m, v9+int32(440), int32(_a_F_GetInsertRecPtr_0), int32(_a_F_GetInsertRecPtr_1), int32(_a_F_GetInsertRecPtr_2))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+184))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+184))
		return v23
	}
}
func F_GetLatestSnapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestSnapshot[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	if v5 != 0 {
		v8 = int32(1)
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+76)))
		v8 = v7
	}
	if v8&int32(1) == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetLatestSnapshot[1])))
		if v14 == int32(0) {
			v17 = F_GetTransactionSnapshot(m)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			v24 = F_GetSnapshotData(m, int32(_a_F_GetLatestSnapshot_0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_GetLatestSnapshot[2])) = v24
				return v24
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_GetLatestSnapshot_1), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetLatestSnapshot_2), int32(361), int32(_a_F_GetLatestSnapshot_3))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
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
	var v5 int32
	_ = v5
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_GetLocksMethodTable[0])))
	return v5
}
func F_GetPubPartitionOptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = F_get_rel_relkind(m, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v55 = F_list_concat(m, l0, v18)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L21
	}
L2:
	;
	return v50
L3:
	;
	return int32(0)
L4:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v7 != int32(112)) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = int32(0)
	v18 = F_find_all_inheritors(m, l2, v16, v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v48 = F_lappend_oid(m, l0, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L20
	}
L8:
	;
	switch l1 - int32(1) {
	case 0:
		goto L9
	case 1:
		goto L1
	default:
		v50 = l0
		goto L2
	}
L9:
	;
	if v18 == int32(0) {
		v50 = l0
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v25 <= v24 {
		v50 = l0
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v28 = l0
	v30 = v24
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
	v37 = F_get_rel_relkind(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v50 = v43
	goto L2
L14:
	;
	if v37 != int32(112) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = F_lappend_oid(m, v28, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v43 = v28
	goto L17
L17:
	;
	v45 = v30 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v45 < v46 {
		v28 = v43
		v30 = v45
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v43 = v41
	goto L17
L19:
	;
	goto L13
L20:
	;
	v50 = v48
	goto L2
L21:
	;
	return v55
}
func F___getopt_msg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = F_fputs(m, l0, int32(_a_F___getopt_msg_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if v6 < int32(0) {
			return
		} else {
			v10 = F_strlen(m, l1)
			v13 = F_fwrite(m, l1, v10, int32(1), int32(_a_F___getopt_msg_0))
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				if v13 == int32(0) {
					return
				} else {
					v19 = F_fwrite(m, l2, int32(1), l3, int32(_a_F___getopt_msg_0))
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						if v19 != l3 {
							return
						} else {
							F_do_putc(m, int32(10), int32(_a_F___getopt_msg_0))
							v25 = m.ExcPending
							if v25 != 0 {
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v673 int32
	_ = v673
	var v686 int32
	_ = v686
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v755 int32
	_ = v755
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v786 int32
	_ = v786
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
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
			v59 = (v55 + int32(_a_F_g_intbig_picksplit_0)) & int32(_a_F_g_intbig_picksplit_1)
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
						v77 = int32(-1)
						v78 = v2
						v86 = int32(1)
						v87 = v2
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
							v105 = v86 + int32(1)
							v106 = v105
							v107 = v77
							v108 = v78
							v115 = v105
							v117 = v87
							for {
								v130 = int32(4)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v115<<(uint(v130)%32))))
								v135 = Fn13923(m, v103, v133, v54, v130)
								mBase = m.M
								v136 = base.B2i32(v107 < v135)
								if v107 < v135 {
									v137 = v135
								} else {
									v137 = v107
								}
								if v107 < v135 {
									v138 = v106
								} else {
									v138 = v117
								}
								if v107 < v135 {
									v139 = v86
								} else {
									v139 = v108
								}
								v141 = v106 + int32(1)
								v143 = v141 & int32(_a_F_g_intbig_picksplit_1)
								if base.Ui32(v143) <= base.Ui32(v59) {
									v106 = v141
									v107 = v137
									v108 = v139
									v115 = v143
									v117 = v138
									continue
								} else {
									break
								}
								break
							}
							if v59 != v105 {
								v77 = v137
								v78 = v139
								v86 = v105
								v87 = v138
								continue
							} else {
								break
							}
							break
						}
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v149 = v139
						v158 = v138
						v160 = v146
					} else {
						v149 = v2
						v158 = v2
						v160 = v67
					}
					v171 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v171
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v171
					v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v176 = int32(8)
					v178 = v54 + v176
					v180 = v26 + int32(4)
					v182 = int32(_a_F_g_intbig_picksplit_1)
					v190 = base.B2i32(v149&v182 == v171) | base.B2i32(v158&v182 == v171)
					if v190 != 0 {
						v191 = int32(1)
					} else {
						v191 = v149
					}
					v194 = int32(4)
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v180+v191&int32(_a_F_g_intbig_picksplit_1)<<(uint(v194)%32))))
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
					v200 = v198 & v194
					if v200 != 0 {
						v201 = v176
					} else {
						v201 = v178
					}
					v202 = F_palloc(m, v201)
					mBase = m.M
					v203 = m.ExcPending
					if v203 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v200
						*(*int32)(unsafe.Add(mBase, uint32(v202))) = v201 << (uint(int32(2)) % 32)
						v208 = int32(0)
						if v200|base.B2i32(v54 == v208) == v208 {
							v213 = int32(8)
							base.MemoryCopy(m, v202+v213, v197+v213, v54)
						} else {
						}
						if v190 != 0 {
							v220 = int32(2)
						} else {
							v220 = v158
						}
						v223 = int32(4)
						v226 = *(*int32)(unsafe.Add(mBase, uint32(v180+v220&int32(_a_F_g_intbig_picksplit_1)<<(uint(v223)%32))))
						v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
						v229 = v227 & v223
						if v229 != 0 {
							v230 = int32(8)
						} else {
							v230 = v178
						}
						v231 = F_palloc(m, v230)
						mBase = m.M
						v232 = m.ExcPending
						if v232 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v229
							*(*int32)(unsafe.Add(mBase, uint32(v231))) = v230 << (uint(int32(2)) % 32)
							v237 = int32(0)
							if v229|base.B2i32(v54 == v237) == v237 {
								v242 = int32(8)
								base.MemoryCopy(m, v231+v242, v226+v242, v54)
							} else {
							}
							v247 = int32(_a_F_g_intbig_picksplit_1)
							v248 = v55 + v247
							v250 = v248 & v247
							v253 = F_palloc(m, v250<<(uint(int32(3))%32))
							mBase = m.M
							v254 = m.ExcPending
							if v254 != 0 {
								return int32(0)
							} else {
								if v55&int32(_a_F_g_intbig_picksplit_1) == int32(1) {
									F_pg_qsort(m, v253, v250, int32(8), int32(_a_F_g_intbig_picksplit_2))
									mBase = m.M
									v262 = m.ExcPending
									if v262 != 0 {
										return int32(0)
									} else {
										v798 = v175
										v801 = v160
										v812 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v798))) = uint16(v812)
										*(*uint16)(unsafe.Add(mBase, uint32(v801))) = uint16(v812)
										F_pfree(m, v253)
										mBase = m.M
										v817 = m.ExcPending
										if v817 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v231
											*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v202
											return v25
										}
									}
								} else {
									v263 = int32(1)
									v265 = v263
									v266 = v263
									for {
										v291 = v253 + v265<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v291-int32(8)))) = uint16(v266)
										v295 = int32(4)
										v300 = *(*int32)(unsafe.Add(mBase, uint32(v180+v265<<(uint(v295)%32))))
										v302 = Fn13923(m, v202, v300, v54, v295)
										mBase = m.M
										v304 = Fn13923(m, v231, v300, v54, int32(4))
										mBase = m.M
										v305 = v302 - v304
										v307 = v305 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v291-v295))) = v305 ^ v307 - v307
										v312 = v266 + int32(1)
										v313 = int32(_a_F_g_intbig_picksplit_1)
										v314 = v312 & v313
										if base.Ui32(v314) <= base.Ui32(v248&v313) {
											v265 = v314
											v266 = v312
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v253, v250, int32(8), int32(_a_F_g_intbig_picksplit_2))
									mBase = m.M
									v321 = m.ExcPending
									if v321 != 0 {
										return int32(0)
									} else {
										v322 = int32(1)
										if base.Ui32(v250) <= base.Ui32(v322) {
											v325 = v322
										} else {
											v325 = v250
										}
										v327 = v54 & int32(2147483644)
										v329 = v54 & int32(3)
										v330 = int32(8)
										v331 = v231 + v330
										v333 = v202 + v330
										v337 = int32(0)
										v347 = v175
										v350 = v160
										for {
											v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253+v337<<(uint(int32(3))%32)))))
											if v191&int32(_a_F_g_intbig_picksplit_1) == v364 {
												*(*uint16)(unsafe.Add(mBase, uint32(v347))) = uint16(v191)
												v367 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v367 + int32(1)
												v771 = v347 + int32(2)
												v774 = v350
											} else {
												if v220&int32(_a_F_g_intbig_picksplit_1) == v364 {
													*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v220)
													v755 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v755 + int32(1)
													v771 = v347
													v774 = v350 + int32(2)
												} else {
													v377 = int32(4)
													v380 = *(*int32)(unsafe.Add(mBase, uint32(v180+v364<<(uint(v377)%32))))
													v382 = Fn13923(m, v202, v380, v54, v377)
													mBase = m.M
													v385 = Fn13923(m, v231, v380, v54, int32(4))
													mBase = m.M
													v387 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													v388 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v389 = v387 - v388
													if base.F64_lt(base.F64_convert_i32_s(v382), base.F64_add(base.F64_convert_i32_s(v385), base.F64_mul(base.F64_convert_i32_s(v389*v389*v389), float64(-1e-05)))) != 0 {
														v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+4)))
														if v397&int32(4) != 0 {
														} else {
															v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+4)))
															if v400&int32(4) != 0 {
																if v54 == int32(0) {
																} else {
																	base.MemoryFill(m, v333, int32(255), v54)
																}
															} else {
																if v54 <= int32(0) {
																} else {
																	v410 = v380 + int32(8)
																	v411 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v417 = v411
																		v418 = v411
																		for {
																			v440 = v417 + v333
																			v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
																			v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v410))))
																			v444 = v441 | v443
																			*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v444)
																			v447 = v417 | int32(1)
																			v448 = v333 + v447
																			v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
																			v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v410))))
																			v452 = v449 | v451
																			*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v452)
																			v455 = v417 | int32(2)
																			v456 = v333 + v455
																			v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
																			v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455+v410))))
																			v460 = v457 | v459
																			*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v460)
																			v463 = v417 | int32(3)
																			v464 = v333 + v463
																			v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
																			v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v410))))
																			v468 = v465 | v467
																			*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v468)
																			v470 = int32(4)
																			v471 = v417 + v470
																			v473 = v418 + v470
																			if v473 != v327 {
																				v417 = v471
																				v418 = v473
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v329 == int32(0) {
																		} else {
																			v479 = v471
																			v503 = v479
																			v516 = v411
																			for {
																				v525 = v503 + v333
																				v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																				v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v410))))
																				v529 = v526 | v528
																				*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																				v531 = int32(1)
																				v534 = v516 + v531
																				if v534 != v329 {
																					v503 = v503 + v531
																					v516 = v534
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v479 = v411
																		v503 = v479
																		v516 = v411
																		for {
																			v525 = v503 + v333
																			v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																			v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v410))))
																			v529 = v526 | v528
																			*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																			v531 = int32(1)
																			v534 = v516 + v531
																			if v534 != v329 {
																				v503 = v503 + v531
																				v516 = v534
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
														*(*uint16)(unsafe.Add(mBase, uint32(v347))) = uint16(v364)
														v561 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v561 + int32(1)
														v771 = v347 + int32(2)
														v774 = v350
													} else {
														v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
														if v567&int32(4) != 0 {
														} else {
															v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+4)))
															if v570&int32(4) != 0 {
																if v54 == int32(0) {
																} else {
																	base.MemoryFill(m, v331, int32(255), v54)
																}
															} else {
																if v54 <= int32(0) {
																} else {
																	v580 = v380 + int32(8)
																	v581 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v587 = v581
																		v588 = v581
																		for {
																			v610 = v587 + v331
																			v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
																			v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587+v580))))
																			v614 = v611 | v613
																			*(*uint8)(unsafe.Add(mBase, uint32(v610))) = uint8(v614)
																			v617 = v587 | int32(1)
																			v618 = v331 + v617
																			v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
																			v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617+v580))))
																			v622 = v619 | v621
																			*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v622)
																			v625 = v587 | int32(2)
																			v626 = v331 + v625
																			v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
																			v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v580))))
																			v630 = v627 | v629
																			*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v630)
																			v633 = v587 | int32(3)
																			v634 = v331 + v633
																			v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
																			v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+v580))))
																			v638 = v635 | v637
																			*(*uint8)(unsafe.Add(mBase, uint32(v634))) = uint8(v638)
																			v640 = int32(4)
																			v641 = v587 + v640
																			v643 = v588 + v640
																			if v643 != v327 {
																				v587 = v641
																				v588 = v643
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v329 == int32(0) {
																		} else {
																			v649 = v641
																			v673 = v649
																			v686 = v581
																			for {
																				v695 = v673 + v331
																				v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
																				v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v580))))
																				v699 = v696 | v698
																				*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v699)
																				v701 = int32(1)
																				v704 = v686 + v701
																				if v704 != v329 {
																					v673 = v673 + v701
																					v686 = v704
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v649 = v581
																		v673 = v649
																		v686 = v581
																		for {
																			v695 = v673 + v331
																			v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
																			v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v580))))
																			v699 = v696 | v698
																			*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v699)
																			v701 = int32(1)
																			v704 = v686 + v701
																			if v704 != v329 {
																				v673 = v673 + v701
																				v686 = v704
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
														*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v364)
														v755 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v755 + int32(1)
														v771 = v347
														v774 = v350 + int32(2)
													}
												}
											}
											v786 = v337 + int32(1)
											if v786 != v325 {
												v337 = v786
												v347 = v771
												v350 = v774
												continue
											} else {
												break
											}
											break
										}
										v798 = v771
										v801 = v774
										v812 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v798))) = uint16(v812)
										*(*uint16)(unsafe.Add(mBase, uint32(v801))) = uint16(v812)
										F_pfree(m, v253)
										mBase = m.M
										v817 = m.ExcPending
										if v817 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v231
											*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v202
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
		v59 = (v55 + int32(_a_F_g_intbig_picksplit_0)) & int32(_a_F_g_intbig_picksplit_1)
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
					v77 = int32(-1)
					v78 = v2
					v86 = int32(1)
					v87 = v2
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
						v105 = v86 + int32(1)
						v106 = v105
						v107 = v77
						v108 = v78
						v115 = v105
						v117 = v87
						for {
							v130 = int32(4)
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v115<<(uint(v130)%32))))
							v135 = Fn13923(m, v103, v133, v54, v130)
							mBase = m.M
							v136 = base.B2i32(v107 < v135)
							if v107 < v135 {
								v137 = v135
							} else {
								v137 = v107
							}
							if v107 < v135 {
								v138 = v106
							} else {
								v138 = v117
							}
							if v107 < v135 {
								v139 = v86
							} else {
								v139 = v108
							}
							v141 = v106 + int32(1)
							v143 = v141 & int32(_a_F_g_intbig_picksplit_1)
							if base.Ui32(v143) <= base.Ui32(v59) {
								v106 = v141
								v107 = v137
								v108 = v139
								v115 = v143
								v117 = v138
								continue
							} else {
								break
							}
							break
						}
						if v59 != v105 {
							v77 = v137
							v78 = v139
							v86 = v105
							v87 = v138
							continue
						} else {
							break
						}
						break
					}
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v149 = v139
					v158 = v138
					v160 = v146
				} else {
					v149 = v2
					v158 = v2
					v160 = v67
				}
				v171 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v171
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v171
				v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v176 = int32(8)
				v178 = v54 + v176
				v180 = v26 + int32(4)
				v182 = int32(_a_F_g_intbig_picksplit_1)
				v190 = base.B2i32(v149&v182 == v171) | base.B2i32(v158&v182 == v171)
				if v190 != 0 {
					v191 = int32(1)
				} else {
					v191 = v149
				}
				v194 = int32(4)
				v197 = *(*int32)(unsafe.Add(mBase, uint32(v180+v191&int32(_a_F_g_intbig_picksplit_1)<<(uint(v194)%32))))
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
				v200 = v198 & v194
				if v200 != 0 {
					v201 = v176
				} else {
					v201 = v178
				}
				v202 = F_palloc(m, v201)
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v200
					*(*int32)(unsafe.Add(mBase, uint32(v202))) = v201 << (uint(int32(2)) % 32)
					v208 = int32(0)
					if v200|base.B2i32(v54 == v208) == v208 {
						v213 = int32(8)
						base.MemoryCopy(m, v202+v213, v197+v213, v54)
					} else {
					}
					if v190 != 0 {
						v220 = int32(2)
					} else {
						v220 = v158
					}
					v223 = int32(4)
					v226 = *(*int32)(unsafe.Add(mBase, uint32(v180+v220&int32(_a_F_g_intbig_picksplit_1)<<(uint(v223)%32))))
					v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
					v229 = v227 & v223
					if v229 != 0 {
						v230 = int32(8)
					} else {
						v230 = v178
					}
					v231 = F_palloc(m, v230)
					mBase = m.M
					v232 = m.ExcPending
					if v232 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v229
						*(*int32)(unsafe.Add(mBase, uint32(v231))) = v230 << (uint(int32(2)) % 32)
						v237 = int32(0)
						if v229|base.B2i32(v54 == v237) == v237 {
							v242 = int32(8)
							base.MemoryCopy(m, v231+v242, v226+v242, v54)
						} else {
						}
						v247 = int32(_a_F_g_intbig_picksplit_1)
						v248 = v55 + v247
						v250 = v248 & v247
						v253 = F_palloc(m, v250<<(uint(int32(3))%32))
						mBase = m.M
						v254 = m.ExcPending
						if v254 != 0 {
							return int32(0)
						} else {
							if v55&int32(_a_F_g_intbig_picksplit_1) == int32(1) {
								F_pg_qsort(m, v253, v250, int32(8), int32(_a_F_g_intbig_picksplit_2))
								mBase = m.M
								v262 = m.ExcPending
								if v262 != 0 {
									return int32(0)
								} else {
									v798 = v175
									v801 = v160
									v812 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v798))) = uint16(v812)
									*(*uint16)(unsafe.Add(mBase, uint32(v801))) = uint16(v812)
									F_pfree(m, v253)
									mBase = m.M
									v817 = m.ExcPending
									if v817 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v231
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v202
										return v25
									}
								}
							} else {
								v263 = int32(1)
								v265 = v263
								v266 = v263
								for {
									v291 = v253 + v265<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v291-int32(8)))) = uint16(v266)
									v295 = int32(4)
									v300 = *(*int32)(unsafe.Add(mBase, uint32(v180+v265<<(uint(v295)%32))))
									v302 = Fn13923(m, v202, v300, v54, v295)
									mBase = m.M
									v304 = Fn13923(m, v231, v300, v54, int32(4))
									mBase = m.M
									v305 = v302 - v304
									v307 = v305 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v291-v295))) = v305 ^ v307 - v307
									v312 = v266 + int32(1)
									v313 = int32(_a_F_g_intbig_picksplit_1)
									v314 = v312 & v313
									if base.Ui32(v314) <= base.Ui32(v248&v313) {
										v265 = v314
										v266 = v312
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v253, v250, int32(8), int32(_a_F_g_intbig_picksplit_2))
								mBase = m.M
								v321 = m.ExcPending
								if v321 != 0 {
									return int32(0)
								} else {
									v322 = int32(1)
									if base.Ui32(v250) <= base.Ui32(v322) {
										v325 = v322
									} else {
										v325 = v250
									}
									v327 = v54 & int32(2147483644)
									v329 = v54 & int32(3)
									v330 = int32(8)
									v331 = v231 + v330
									v333 = v202 + v330
									v337 = int32(0)
									v347 = v175
									v350 = v160
									for {
										v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253+v337<<(uint(int32(3))%32)))))
										if v191&int32(_a_F_g_intbig_picksplit_1) == v364 {
											*(*uint16)(unsafe.Add(mBase, uint32(v347))) = uint16(v191)
											v367 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v367 + int32(1)
											v771 = v347 + int32(2)
											v774 = v350
										} else {
											if v220&int32(_a_F_g_intbig_picksplit_1) == v364 {
												*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v220)
												v755 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v755 + int32(1)
												v771 = v347
												v774 = v350 + int32(2)
											} else {
												v377 = int32(4)
												v380 = *(*int32)(unsafe.Add(mBase, uint32(v180+v364<<(uint(v377)%32))))
												v382 = Fn13923(m, v202, v380, v54, v377)
												mBase = m.M
												v385 = Fn13923(m, v231, v380, v54, int32(4))
												mBase = m.M
												v387 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v388 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												v389 = v387 - v388
												if base.F64_lt(base.F64_convert_i32_s(v382), base.F64_add(base.F64_convert_i32_s(v385), base.F64_mul(base.F64_convert_i32_s(v389*v389*v389), float64(-1e-05)))) != 0 {
													v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+4)))
													if v397&int32(4) != 0 {
													} else {
														v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+4)))
														if v400&int32(4) != 0 {
															if v54 == int32(0) {
															} else {
																base.MemoryFill(m, v333, int32(255), v54)
															}
														} else {
															if v54 <= int32(0) {
															} else {
																v410 = v380 + int32(8)
																v411 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v417 = v411
																	v418 = v411
																	for {
																		v440 = v417 + v333
																		v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
																		v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v410))))
																		v444 = v441 | v443
																		*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v444)
																		v447 = v417 | int32(1)
																		v448 = v333 + v447
																		v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
																		v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v410))))
																		v452 = v449 | v451
																		*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v452)
																		v455 = v417 | int32(2)
																		v456 = v333 + v455
																		v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
																		v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455+v410))))
																		v460 = v457 | v459
																		*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v460)
																		v463 = v417 | int32(3)
																		v464 = v333 + v463
																		v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
																		v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v410))))
																		v468 = v465 | v467
																		*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v468)
																		v470 = int32(4)
																		v471 = v417 + v470
																		v473 = v418 + v470
																		if v473 != v327 {
																			v417 = v471
																			v418 = v473
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v329 == int32(0) {
																	} else {
																		v479 = v471
																		v503 = v479
																		v516 = v411
																		for {
																			v525 = v503 + v333
																			v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																			v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v410))))
																			v529 = v526 | v528
																			*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																			v531 = int32(1)
																			v534 = v516 + v531
																			if v534 != v329 {
																				v503 = v503 + v531
																				v516 = v534
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v479 = v411
																	v503 = v479
																	v516 = v411
																	for {
																		v525 = v503 + v333
																		v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																		v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v410))))
																		v529 = v526 | v528
																		*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																		v531 = int32(1)
																		v534 = v516 + v531
																		if v534 != v329 {
																			v503 = v503 + v531
																			v516 = v534
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
													*(*uint16)(unsafe.Add(mBase, uint32(v347))) = uint16(v364)
													v561 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v561 + int32(1)
													v771 = v347 + int32(2)
													v774 = v350
												} else {
													v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
													if v567&int32(4) != 0 {
													} else {
														v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+4)))
														if v570&int32(4) != 0 {
															if v54 == int32(0) {
															} else {
																base.MemoryFill(m, v331, int32(255), v54)
															}
														} else {
															if v54 <= int32(0) {
															} else {
																v580 = v380 + int32(8)
																v581 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v587 = v581
																	v588 = v581
																	for {
																		v610 = v587 + v331
																		v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
																		v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587+v580))))
																		v614 = v611 | v613
																		*(*uint8)(unsafe.Add(mBase, uint32(v610))) = uint8(v614)
																		v617 = v587 | int32(1)
																		v618 = v331 + v617
																		v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
																		v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617+v580))))
																		v622 = v619 | v621
																		*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v622)
																		v625 = v587 | int32(2)
																		v626 = v331 + v625
																		v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
																		v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v580))))
																		v630 = v627 | v629
																		*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v630)
																		v633 = v587 | int32(3)
																		v634 = v331 + v633
																		v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
																		v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+v580))))
																		v638 = v635 | v637
																		*(*uint8)(unsafe.Add(mBase, uint32(v634))) = uint8(v638)
																		v640 = int32(4)
																		v641 = v587 + v640
																		v643 = v588 + v640
																		if v643 != v327 {
																			v587 = v641
																			v588 = v643
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v329 == int32(0) {
																	} else {
																		v649 = v641
																		v673 = v649
																		v686 = v581
																		for {
																			v695 = v673 + v331
																			v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
																			v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v580))))
																			v699 = v696 | v698
																			*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v699)
																			v701 = int32(1)
																			v704 = v686 + v701
																			if v704 != v329 {
																				v673 = v673 + v701
																				v686 = v704
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v649 = v581
																	v673 = v649
																	v686 = v581
																	for {
																		v695 = v673 + v331
																		v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
																		v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v580))))
																		v699 = v696 | v698
																		*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v699)
																		v701 = int32(1)
																		v704 = v686 + v701
																		if v704 != v329 {
																			v673 = v673 + v701
																			v686 = v704
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
													*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v364)
													v755 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v755 + int32(1)
													v771 = v347
													v774 = v350 + int32(2)
												}
											}
										}
										v786 = v337 + int32(1)
										if v786 != v325 {
											v337 = v786
											v347 = v771
											v350 = v774
											continue
										} else {
											break
										}
										break
									}
									v798 = v771
									v801 = v774
									v812 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v798))) = uint16(v812)
									*(*uint16)(unsafe.Add(mBase, uint32(v801))) = uint16(v812)
									F_pfree(m, v253)
									mBase = m.M
									v817 = m.ExcPending
									if v817 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v231
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v202
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
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
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = int32(0)
	v55 = v49 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v55
	v58 = v50 + int32(8)
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryFill(m, v58, int32(0), v47)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v61 <= int32(0) {
		v256 = v55
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v256) >> (uint(int32(2)) % 32))
	return v50
L17:
	;
	v67 = v47 & int32(3)
	v72 = v61
	v83 = v2
	goto L18
L18:
	;
	v89 = int32(4)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v83<<(uint(v89)%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	if v93&v89 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(17179869216)
	v256 = int32(32)
	goto L16
L20:
	;
	if base.B2i32(v47 <= int32(0)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	v101 = v92 + int32(8)
	v102 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v224 = v72
	goto L25
L25:
	;
	v242 = v83 + int32(1)
	if v242 < v224 {
		v72 = v224
		v83 = v242
		goto L18
	} else {
		goto L37
	}
L26:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v224 = v223
	goto L25
L27:
	;
	v107 = v102
	v114 = v102
	goto L30
L28:
	;
	v161 = v102
	goto L29
L29:
	;
	v178 = v161
	v190 = v102
	goto L34
L30:
	;
	v124 = v107 + v58
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v101))))
	v128 = v125 | v127
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v128)
	v131 = v107 | int32(1)
	v132 = v58 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v131))))
	v136 = v133 | v135
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v136)
	v139 = v107 | int32(2)
	v140 = v58 + v139
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v139))))
	v144 = v141 | v143
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v144)
	v147 = v107 | int32(3)
	v148 = v58 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
	v152 = v149 | v151
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v152)
	v154 = int32(4)
	v155 = v107 + v154
	v157 = v114 + v154
	if v157 != v47&int32(2147483644) {
		v107 = v155
		v114 = v157
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v67 == int32(0) {
		goto L26
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v161 = v155
	goto L29
L34:
	;
	v195 = v178 + v58
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v101))))
	v199 = v196 | v198
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v199)
	v201 = int32(1)
	v204 = v190 + v201
	if v204 != v67 {
		v178 = v178 + v201
		v190 = v204
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L26
L36:
	;
	goto L35
L37:
	;
	v256 = v55
	goto L16
}
func F_gb18030_to_utf8(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(0)
	v7 = Fn13848(m, l0, int32(39), int32(_a_F_gb18030_to_utf8_0), v4, v4, int32(_a_F_gb18030_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v33 int32
	_ = v33
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
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
	v31 = v12
	v33 = v2
	v35 = v2
	goto L6
L4:
	;
	v156 = v19
	v159 = v14
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = (v156 - v159) << (uint(int32(2)) % 32)
	return v159
L6:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v39 = v27 - v30
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 <= v36+(v37+(v38+(v39+int32(base.Ui32(v40)>>(uint(int32(16))%32))))) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v156 = v143
	v159 = v78
	goto L5
L8:
	;
	v52 = v30
	v53 = v31
	goto L11
L9:
	;
	v75 = v27
	v76 = v40
	v78 = v30
	v79 = v31
	v80 = v36
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
	v59 = v53 << (uint(int32(1)) % 32)
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
	v79 = v59
	v80 = v62
	goto L10
L13:
	;
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v59 <= v62+(v63+(v64+(v39+int32(base.Ui32(v65)>>(uint(int32(16))%32))))) {
		v52 = v60
		v53 = v59
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = v29 + int32(16)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if (v149-v150)>>(uint(int32(4))%32) < v147 {
		v27 = v143
		v29 = v149
		v30 = v78
		v31 = v79
		v33 = v145
		v35 = v146
		goto L6
	} else {
		goto L44
	}
L16:
	;
	if v33 != 0 {
		v97 = v75
		v98 = v76
		v99 = v35
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
		v143 = v75
		v145 = v33
		v146 = v35
		goto L15
	} else {
		goto L42
	}
L19:
	;
	if v98&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v89 = v35 + int32(1)
	if v89 < int32(2) {
		v97 = v75
		v98 = v76
		v99 = v89
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if v80 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v75, v92, v80)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v97 = v75 + v94
	v98 = v96
	v99 = v89
	goto L19
L25:
	;
	v102 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v102)
	v104 = int32(1)
	v143 = v97 + v104
	v145 = v104
	v146 = v99
	goto L15
L26:
	;
	goto L27
L27:
	;
	v107 = int32(1)
	if v98&int32(16) != 0 {
		v143 = v97
		v145 = v107
		v146 = v99
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v98&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v112 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v118 = v97
	v119 = v98
	goto L31
L31:
	;
	v121 = int32(base.Ui32(v119) >> (uint(int32(16)) % 32))
	if v121 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	base.MemoryCopy(m, v97, v113, v112)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v118 = v97 + v116
	v119 = v115
	goto L31
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	base.MemoryCopy(m, v118, v122, v121)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v127 = v118 + int32(base.Ui32(v124)>>(uint(int32(16))%32))
	if v124&int32(1) == int32(0) {
		v143 = v127
		v145 = v107
		v146 = v99
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	if v132 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v127, v133, v132)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	v143 = v127 + v135
	v145 = v107
	v146 = v99
	goto L15
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	F_pfree(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v143 = v75
	v145 = int32(0)
	v146 = v35
	goto L15
L44:
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
	var v60 int32
	_ = v60
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
	var v142 int32
	_ = v142
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
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	v142 = v133
	v143 = int32(1)
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
	v60 = v49
	goto L17
L17:
	;
	v66 = v60 << (uint(int32(2)) % 32)
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
	if v71 != v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+v18))) = int32(-1)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v78 = v66 + v18
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v80 = F_exprTypmod(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v82 == v42 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
	goto L19
L26:
	;
	goto L27
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v80 == v85 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(-1)
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
	v102 = v60 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v102 < v103 {
		v57 = v100
		v60 = v102
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
	v188 = v140 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v181+v188)))
	v191 = int32(0)
	v192 = base.I32_extend16_s(v143)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188+v18)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v198 = F_makeVar(m, v191, v192, v193, v195, v196, v191)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L51
	}
L43:
	;
	v171 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v170 == v171)|(base.B2i32(v158 == v171)|base.B2i32(v175 <= v140)) == v171 {
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
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L50
	}
L46:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v181 != 0 {
		goto L42
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v183 = v142
	goto L45
L49:
	;
	goto L48
L50:
	;
	return v183
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v201 = F_pstrdup(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v204 = F_makeTargetEntry(m, v198, v192, v201, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v204)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+16)) = v206
	v208 = int32(1)
	v212 = F_lappend(m, v142, v204)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v140 = v140 + v208
	v142 = v212
	v143 = v143 + v208
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L51
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
	v33 = int32(0)
	v34 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v30 = F_palloc0(m, v14)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30
	v33 = v30
	v34 = v14
	goto L6
L11:
	;
	v38 = l2 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v41 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v41)
	v47 = v33
	v48 = l1
	v50 = v34
	v52 = v39
	v54 = v38
	goto L15
L15:
	;
	v55 = l1 + l2
	v61 = v48
	goto L18
L16:
	;
	F_pfree(m, v52)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L50
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui32(v55) <= base.Ui32(v61) {
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v81 = v61
	goto L24
L20:
	;
	v69 = F_pg_mblen_range(m, v61, v55)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v72 = F_t_isalnum_with_len(m, v61, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v72 == int32(0) {
		v61 = v61 + v69
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v88 = F_pg_mblen_range(m, v81, v55)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	if v61 == int32(0) {
		goto L17
	} else {
		goto L32
	}
L26:
	;
	v90 = F_t_isalnum_with_len(m, v81, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v90 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = v81 + v88
	if base.Ui32(v92) < base.Ui32(v55) {
		v81 = v92
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v94 = v81
	goto L30
L30:
	;
	goto L25
L31:
	;
	v94 = v92
	goto L30
L32:
	;
	v99 = F_str_tolower(m, v61, v94-v61, int32(100))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v101 = F_strlen(m, v99)
	mBase = m.M
	if base.Ui32(v54-int32(4)) < base.Ui32(v101) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v52)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v113 = v52
	v114 = v54
	goto L36
L36:
	;
	if v101 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v108 = v101 + int32(4)
	v109 = F_palloc(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v111 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v109))) = uint16(v111)
	v113 = v109
	v114 = v108
	goto L36
L39:
	;
	base.MemoryCopy(m, v113+int32(2), v99, v101)
	goto L41
L40:
	;
	goto L41
L41:
	;
	F_pfree(m, v99)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v121 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v101+v113)+2)) = uint16(v121)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_make_trigrams(m, l0, v113, v101+int32(3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v128 = int32(0)
	if v47 == v128 {
		v47 = v128
		v48 = v94
		v52 = v113
		v54 = v114
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v131 <= v50 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v140 = v138 + v123
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v142 = int32(1)
	v143 = v141 | v142
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v143)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = v138 + v145 - v142
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v151 = v149 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v151)
	v47 = v138
	v48 = v94
	v50 = v139
	v52 = v113
	v54 = v114
	goto L15
L46:
	;
	v138 = v47
	v139 = v50
	goto L45
L47:
	;
	goto L48
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = F_repalloc0(m, v47, v50, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v138 = v134
	v139 = v137
	goto L45
L50:
	;
	goto L13
L51:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_generate_trgm_only_1), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_generate_trgm_only_2), int32(115), int32(_a_F_generate_trgm_only_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
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
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L80
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v331
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
	v331 = v22
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
	v60 = l1
	goto L11
L11:
	;
	v65 = int32(0)
	v69 = v53
	v71 = v65
	v72 = v65
	goto L13
L12:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_pfree(m, v47)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L68
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
	if v60 <= v114-v53 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	goto L14
L16:
	;
	if v71&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v112 = v69 + v81
	if base.Ui32(v112) < base.Ui32(v44) {
		v69 = v112
		v71 = v109
		v72 = v110
		goto L13
	} else {
		goto L32
	}
L18:
	;
	v114 = v69
	v115 = v106
	v116 = v72
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
		v110 = v72
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
	v124 = int32(_a_F_generate_wildcard_trgm_0)
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
	v227 = F_str_tolower(m, v47, v212-v47, int32(100))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L64
	}
L39:
	;
	v211 = v195
	v212 = v196 + int32(1)
	goto L38
L40:
	;
	v191 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v191)
	v195 = v188
	v196 = v131
	goto L39
L41:
	;
	v130 = v114
	v131 = v126
	v132 = v115
	goto L44
L42:
	;
	v172 = v114
	v173 = v126
	goto L43
L43:
	;
	v184 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v184)
	if v172 == int32(0) {
		goto L33
	} else {
		goto L63
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
	v172 = v168
	v173 = v165
	goto L43
L46:
	;
	if v132&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v168 = v130 + v142
	if base.Ui32(v168) < base.Ui32(v44) {
		v130 = v168
		v131 = v165
		v132 = v166
		goto L44
	} else {
		goto L62
	}
L48:
	;
	if v142 != 0 {
		goto L59
	} else {
		goto L60
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
		v165 = v131
		v166 = int32(1)
		goto L47
	case 1, 2:
		goto L54
	case 3:
		v211 = v130
		v212 = v131
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
	v188 = v130 - int32(1)
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
		v211 = v130
		v212 = v131
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
		v188 = v130
		goto L40
	} else {
		goto L58
	}
L58:
	;
	goto L48
L59:
	;
	base.MemoryCopy(m, v131, v130, v142)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v165 = v131 + v142
	v166 = int32(0)
	goto L47
L62:
	;
	goto L45
L63:
	;
	v195 = v172
	v196 = v173
	goto L39
L64:
	;
	v229 = F_strlen(m, v227)
	mBase = m.M
	F_make_trigrams(m, v17+int32(4), v227, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v227)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v235 = l0 - v211 + l1
	if int32(0) < v235 {
		v53 = v211
		v60 = v235
		goto L11
	} else {
		goto L67
	}
L67:
	;
	goto L33
L68:
	;
	if int32(2) <= v253 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v259 = v252 + int32(5)
	F_pg_qsort(m, v259, v253, int32(3), int32(_a_F_generate_wildcard_trgm_1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	v312 = v253
	goto L71
L71:
	;
	v323 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v252)+4)) = uint8(v323)
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v312*int32(12) + int32(20)
	v331 = v252
	goto L2
L72:
	;
	v268 = int32(1)
	v270 = int32(0)
	goto L73
L73:
	;
	v280 = int32(3)
	v282 = v259 + v268*v280
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_generate_wildcard_trgm[0]))
	v288 = m.T0[v287].(func(*base.Module, int32, int32) int32)(m, v282, v259+v270*v280)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	v312 = v303 + int32(1)
	goto L71
L75:
	;
	v305 = v268 + int32(1)
	if v305 != v253 {
		v268 = v305
		v270 = v303
		goto L73
	} else {
		goto L79
	}
L76:
	;
	if v288 == int32(0) {
		v303 = v270
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v293 = v270 + int32(1)
	if v293 == v268 {
		v303 = v268
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v297 = v259 + v293*int32(3)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+2)) = uint8(v298)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282))))
	*(*uint16)(unsafe.Add(mBase, uint32(v297))) = uint16(v300)
	v303 = v293
	goto L75
L79:
	;
	goto L74
L80:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_generate_wildcard_trgm_2), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_generate_wildcard_trgm_3), int32(115), int32(_a_F_generate_wildcard_trgm_4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
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
	var v11 int32
	_ = v11
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
	var v59 int32
	_ = v59
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
	var v125 int32
	_ = v125
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
	var v193 int32
	_ = v193
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
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v6
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
	v64 = v59
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
	v59 = int32(0)
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
		v59 = v35
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v40 = v38 - int32(97)
	if v40 < int32(0) {
		v59 = v35
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v46)>>(uint(v40&int32(7))%32))&int32(1) == int32(0) {
		v59 = v35
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	goto L1
L18:
	;
	return v1021
L19:
	;
	v1016 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L62
	} else {
		goto L300
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v218 = v6
	goto L65
L21:
	;
	v210 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L62
	} else {
		goto L63
	}
L22:
	;
	if v203 <= v11 {
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
	v130 = v125
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
	v125 = int32(0)
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
		v125 = v101
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v106 = v104 - int32(97)
	if v106 < int32(0) {
		v125 = v101
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) == int32(0) {
		v125 = v101
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
	v198 = v193
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
	v193 = int32(0)
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
		v193 = v169
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v174 = v172 - int32(97)
	if v174 < int32(0) {
		v193 = v169
		goto L50
	} else {
		goto L57
	}
L57:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v174)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v180)>>(uint(v174&int32(7))%32))&int32(1) == int32(0) {
		v193 = v169
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
	v206 = v11 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206
	v11 = v206
	goto L1
L62:
	;
	return int32(0)
L63:
	;
	if v210 < int32(0) {
		v1021 = v210
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
	v225 = F_find_among(m, l0, int32(_a_F_german_ISO_8859_1_stem_2), int32(6))
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
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v255
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v269 = v267 + int32(3)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v270 < v269 {
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
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = v262
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
	v251 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L62
	} else {
		goto L81
	}
L72:
	;
	v245 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_4))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L62
	} else {
		goto L79
	}
L73:
	;
	v239 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_5))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L62
	} else {
		goto L77
	}
L74:
	;
	v233 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_ISO_8859_1_stem_6))
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
	v1021 = v233
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
	v1021 = v239
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
	v1021 = v245
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
	v1021 = v251
	goto L18
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v227 + int32(1)
	goto L69
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v507
	if v507 <= v6 {
		goto L151
	} else {
		goto L152
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v269
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v267
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v282 < v267 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v322 < int32(0) {
		goto L84
	} else {
		goto L101
	}
L87:
	;
	v284 = v267
	goto L89
L88:
	;
	v284 = v282
	goto L89
L89:
	;
	v291 = v267
	goto L91
L90:
	;
	v322 = v302
	goto L86
L91:
	;
	if v291 == v284 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v322 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v291))))
	if int32(252) < v297 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v314 = v291 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314
	v291 = v314
	goto L91
L97:
	;
	v299 = v297 - int32(97)
	if v299 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v302 = int32(1)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v299)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v306)>>(uint(v299&int32(7))%32))&v302 != 0 {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L96
L101:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v326 = v325 + v322
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v326
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v337 < v326 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v380 < int32(0) {
		goto L84
	} else {
		goto L116
	}
L103:
	;
	v339 = v326
	goto L105
L104:
	;
	v339 = v337
	goto L105
L105:
	;
	v345 = v326
	goto L107
L106:
	;
	v380 = int32(1)
	goto L102
L107:
	;
	if v345 == v339 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v380 = int32(-1)
	goto L102
L110:
	;
	goto L111
L111:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352+v345))))
	if int32(252) < v354 {
		goto L106
	} else {
		goto L112
	}
L112:
	;
	v356 = v354 - int32(97)
	if v356 < int32(0) {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v356)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v362)>>(uint(v356&int32(7))%32))&int32(1) == int32(0) {
		goto L106
	} else {
		goto L114
	}
L114:
	;
	v371 = v345 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371
	v345 = v371
	goto L107
L116:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v384 = v383 + v380
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	if v387 < v384 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v389 = v384
	goto L119
L118:
	;
	v389 = v387
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+8)) = v389
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v399 < v398 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v439 < int32(0) {
		goto L84
	} else {
		goto L135
	}
L121:
	;
	v401 = v398
	goto L123
L122:
	;
	v401 = v399
	goto L123
L123:
	;
	v408 = v398
	goto L125
L124:
	;
	v439 = v419
	goto L120
L125:
	;
	if v408 == v401 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v439 = int32(-1)
	goto L120
L128:
	;
	goto L129
L129:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v408))))
	if int32(252) < v414 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v431 = v408 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v431
	v408 = v431
	goto L125
L131:
	;
	v416 = v414 - int32(97)
	if v416 < int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v419 = int32(1)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v423)>>(uint(v416&int32(7))%32))&v419 != 0 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	goto L130
L135:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v443 = v442 + v439
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v443
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v454 < v443 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v497 < int32(0) {
		goto L84
	} else {
		goto L150
	}
L137:
	;
	v456 = v443
	goto L139
L138:
	;
	v456 = v454
	goto L139
L139:
	;
	v462 = v443
	goto L141
L140:
	;
	v497 = int32(1)
	goto L136
L141:
	;
	if v462 == v456 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v497 = int32(-1)
	goto L136
L144:
	;
	goto L145
L145:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469+v462))))
	if int32(252) < v471 {
		goto L140
	} else {
		goto L146
	}
L146:
	;
	v473 = v471 - int32(97)
	if v473 < int32(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v473)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v479)>>(uint(v473&int32(7))%32))&int32(1) == int32(0) {
		goto L140
	} else {
		goto L148
	}
L148:
	;
	v488 = v462 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v488
	v462 = v488
	goto L141
L150:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = v501 + v497
	goto L84
L151:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675
	v679 = v675 - int32(1)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v679 <= v680 {
		goto L198
	} else {
		goto L199
	}
L152:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v513 = int32(1)
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v507-v513))))
	if base.B2i32(v515&int32(224) != int32(96))|base.B2i32(v513<<(uint(v515)%32)&int32(_a_F_german_ISO_8859_1_stem_7) == int32(0)) != 0 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v529 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_8), int32(11))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L62
	} else {
		goto L154
	}
L154:
	;
	if v529 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if v533 < v536 {
		goto L151
	} else {
		goto L156
	}
L156:
	;
	switch v529 - int32(1) {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	case 3:
		goto L158
	case 4:
		goto L157
	default:
		goto L151
	}
L157:
	;
	v669 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_9))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L62
	} else {
		goto L196
	}
L158:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L184
L159:
	;
	v569 = F_slice_del(m, l0)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L62
	} else {
		goto L171
	}
L160:
	;
	v565 = F_slice_del(m, l0)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L62
	} else {
		goto L169
	}
L161:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v541 = int32(4)
	v543 = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v545-v546 < v541 {
		v556 = v543
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if v556 != 0 {
		goto L151
	} else {
		goto L166
	}
L163:
	;
	goto L162
L164:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v552 = F_memcmp(m, v549+v545-v541, int32(_a_F_german_ISO_8859_1_stem_10), v541)
	mBase = m.M
	if v552 != 0 {
		v556 = v543
		goto L163
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v545 - v541
	v556 = int32(1)
	goto L163
L166:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v557 + (v533 - v540)
	v561 = F_slice_del(m, l0)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L62
	} else {
		goto L167
	}
L167:
	;
	if int32(0) <= v561 {
		goto L151
	} else {
		goto L168
	}
L168:
	;
	v1021 = v561
	goto L18
L169:
	;
	if int32(0) <= v565 {
		goto L151
	} else {
		goto L170
	}
L170:
	;
	v1021 = v565
	goto L18
L171:
	;
	if v569 < int32(0) {
		v1021 = v569
		goto L18
	} else {
		goto L172
	}
L172:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v573 <= v575 {
		goto L151
	} else {
		goto L173
	}
L173:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v573-int32(1)))))
	if v581 != int32(115) {
		goto L151
	} else {
		goto L174
	}
L174:
	;
	v585 = v573 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v585
	v588 = int32(3)
	v590 = int32(0)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v585-v593 < v588 {
		v603 = v590
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v603 == int32(0) {
		goto L151
	} else {
		goto L179
	}
L176:
	;
	goto L175
L177:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v599 = F_memcmp(m, v596+v585-v588, int32(_a_F_german_ISO_8859_1_stem_11), v588)
	mBase = m.M
	if v599 != 0 {
		v603 = v590
		goto L176
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v585 - v588
	v603 = int32(1)
	goto L176
L179:
	;
	v606 = F_slice_del(m, l0)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L62
	} else {
		goto L180
	}
L180:
	;
	if int32(0) <= v606 {
		goto L151
	} else {
		goto L181
	}
L181:
	;
	v1021 = v606
	goto L18
L182:
	;
	if v662 != 0 {
		goto L151
	} else {
		goto L193
	}
L183:
	;
	v662 = v658
	goto L182
L184:
	;
	if v618 <= v619 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v658 = int32(0)
	goto L183
L186:
	;
	v662 = int32(-1)
	goto L182
L187:
	;
	goto L188
L188:
	;
	v631 = int32(1)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632+v618-v631))))
	if int32(116) < v636 {
		v658 = v631
		goto L183
	} else {
		goto L189
	}
L189:
	;
	v638 = v636 - int32(98)
	if v638 < int32(0) {
		v658 = v631
		goto L183
	} else {
		goto L190
	}
L190:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v638)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v644)>>(uint(v638&int32(7))%32))&int32(1) == int32(0) {
		v658 = v631
		goto L183
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v618 - int32(1)
	goto L192
L192:
	;
	goto L185
L193:
	;
	v663 = F_slice_del(m, l0)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L62
	} else {
		goto L194
	}
L194:
	;
	if int32(0) <= v663 {
		goto L151
	} else {
		goto L195
	}
L195:
	;
	v1021 = v663
	goto L18
L196:
	;
	if v669 < int32(0) {
		v1021 = v669
		goto L18
	} else {
		goto L197
	}
L197:
	;
	goto L151
L198:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v778
	v782 = v778 - int32(1)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v782 <= v783 {
		goto L223
	} else {
		goto L224
	}
L199:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682+v679))))
	if base.B2i32(v684&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v684)%32)&int32(_a_F_german_ISO_8859_1_stem_12) == int32(0)) != 0 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v698 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_13), int32(4))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L62
	} else {
		goto L201
	}
L201:
	;
	if v698 == int32(0) {
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v702
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+8))
	if v702 < v705 {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	switch v698 - int32(1) {
	case 0:
		goto L205
	case 1:
		goto L204
	default:
		goto L198
	}
L204:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L210
L205:
	;
	v709 = F_slice_del(m, l0)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L62
	} else {
		goto L206
	}
L206:
	;
	if int32(0) <= v709 {
		goto L198
	} else {
		goto L207
	}
L207:
	;
	v1021 = v709
	goto L18
L208:
	;
	if v765 != 0 {
		goto L198
	} else {
		goto L219
	}
L209:
	;
	v765 = v761
	goto L208
L210:
	;
	if v721 <= v722 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v761 = int32(0)
	goto L209
L212:
	;
	v765 = int32(-1)
	goto L208
L213:
	;
	goto L214
L214:
	;
	v734 = int32(1)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735+v721-v734))))
	if int32(116) < v739 {
		v761 = v734
		goto L209
	} else {
		goto L215
	}
L215:
	;
	v741 = v739 - int32(98)
	if v741 < int32(0) {
		v761 = v734
		goto L209
	} else {
		goto L216
	}
L216:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v741)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v747)>>(uint(v741&int32(7))%32))&int32(1) == int32(0) {
		v761 = v734
		goto L209
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v721 - int32(1)
	goto L218
L218:
	;
	goto L211
L219:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v768 = v766 - int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v768
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v768 < v770 {
		goto L198
	} else {
		goto L220
	}
L220:
	;
	v772 = F_slice_del(m, l0)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L62
	} else {
		goto L221
	}
L221:
	;
	if v772 < int32(0) {
		v1021 = v772
		goto L18
	} else {
		goto L222
	}
L222:
	;
	goto L198
L223:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v965
	v968 = v965
	goto L281
L224:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785+v782))))
	if base.B2i32(v787&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v787)%32)&int32(_a_F_german_ISO_8859_1_stem_14) == int32(0)) != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v801 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_15), int32(8))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L62
	} else {
		goto L226
	}
L226:
	;
	if v801 == int32(0) {
		goto L223
	} else {
		goto L227
	}
L227:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v805
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)+4))
	if v805 < v808 {
		goto L223
	} else {
		goto L228
	}
L228:
	;
	switch v801 - int32(1) {
	case 0:
		goto L232
	case 1:
		goto L231
	case 2:
		goto L230
	case 3:
		goto L229
	default:
		goto L223
	}
L229:
	;
	v923 = F_slice_del(m, l0)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L62
	} else {
		goto L270
	}
L230:
	;
	v867 = F_slice_del(m, l0)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L62
	} else {
		goto L253
	}
L231:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v854 < v805 {
		goto L247
	} else {
		goto L248
	}
L232:
	;
	v812 = F_slice_del(m, l0)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L62
	} else {
		goto L233
	}
L233:
	;
	if v812 < int32(0) {
		v1021 = v812
		goto L18
	} else {
		goto L234
	}
L234:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v816
	v818 = int32(2)
	v820 = int32(0)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v816-v823 < v818 {
		v833 = v820
		goto L236
	} else {
		goto L237
	}
L235:
	;
	if v833 == int32(0) {
		goto L223
	} else {
		goto L239
	}
L236:
	;
	goto L235
L237:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v829 = F_memcmp(m, v826+v816-v818, int32(_a_F_german_ISO_8859_1_stem_16), v818)
	mBase = m.M
	if v829 != 0 {
		v833 = v820
		goto L236
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v816 - v818
	v833 = int32(1)
	goto L236
L239:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v836
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v838 < v836 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v836-int32(1)))))
	if v844 == int32(101) {
		goto L223
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v836 < v848 {
		goto L223
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	v850 = F_slice_del(m, l0)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L62
	} else {
		goto L245
	}
L245:
	;
	if int32(0) <= v850 {
		goto L223
	} else {
		goto L246
	}
L246:
	;
	v1021 = v850
	goto L18
L247:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856+v805-int32(1)))))
	if v860 == int32(101) {
		goto L223
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v863 = F_slice_del(m, l0)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L62
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	if int32(0) <= v863 {
		goto L223
	} else {
		goto L252
	}
L252:
	;
	v1021 = v863
	goto L18
L253:
	;
	if v867 < int32(0) {
		v1021 = v867
		goto L18
	} else {
		goto L254
	}
L254:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v871
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v874 = int32(2)
	v876 = int32(0)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v871-v879 < v874 {
		v889 = v876
		goto L256
	} else {
		goto L257
	}
L255:
	;
	if v889 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	goto L255
L257:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v885 = F_memcmp(m, v882+v871-v874, int32(_a_F_german_ISO_8859_1_stem_17), v874)
	mBase = m.M
	if v885 != 0 {
		v889 = v876
		goto L256
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v871 - v874
	v889 = int32(1)
	goto L256
L259:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v894 = v892 + (v871 - v873)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v894
	v896 = int32(2)
	v898 = int32(0)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v894-v901 < v896 {
		v911 = v898
		goto L263
	} else {
		goto L264
	}
L260:
	;
	goto L261
L261:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v914
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+8))
	if v914 < v917 {
		goto L223
	} else {
		goto L267
	}
L262:
	;
	if v911 == int32(0) {
		goto L223
	} else {
		goto L266
	}
L263:
	;
	goto L262
L264:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v907 = F_memcmp(m, v904+v894-v896, int32(_a_F_german_ISO_8859_1_stem_18), v896)
	mBase = m.M
	if v907 != 0 {
		v911 = v898
		goto L263
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v894 - v896
	v911 = int32(1)
	goto L263
L266:
	;
	goto L261
L267:
	;
	v919 = F_slice_del(m, l0)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L62
	} else {
		goto L268
	}
L268:
	;
	if int32(0) <= v919 {
		goto L223
	} else {
		goto L269
	}
L269:
	;
	v1021 = v919
	goto L18
L270:
	;
	if v923 < int32(0) {
		v1021 = v923
		goto L18
	} else {
		goto L271
	}
L271:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v927
	v930 = v927 - int32(1)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v930 <= v931 {
		goto L223
	} else {
		goto L272
	}
L272:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933+v930))))
	if base.Ui32(int32(1)) < base.Ui32((v935-int32(103))&int32(255)) {
		goto L223
	} else {
		goto L273
	}
L273:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v945 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_19), int32(2))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L62
	} else {
		goto L274
	}
L274:
	;
	if v945 == int32(0) {
		goto L223
	} else {
		goto L275
	}
L275:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v949
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if v949 < v952 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v954 + (v927 - v942)
	goto L223
L277:
	;
	goto L278
L278:
	;
	v958 = F_slice_del(m, l0)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L62
	} else {
		goto L279
	}
L279:
	;
	if v958 < int32(0) {
		v1021 = v958
		goto L18
	} else {
		goto L280
	}
L280:
	;
	goto L223
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v968
	v975 = F_find_among(m, l0, int32(_a_F_german_ISO_8859_1_stem_20), int32(6))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L62
	} else {
		goto L284
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v965
	v1021 = int32(1)
	goto L18
L283:
	;
	goto L282
L284:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v977
	switch v975 - int32(1) {
	case 0:
		goto L290
	case 1:
		goto L289
	case 2:
		goto L288
	case 3:
		goto L287
	case 4:
		goto L286
	default:
		goto L285
	}
L285:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v968 = v1011
	goto L281
L286:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1005 <= v977 {
		goto L283
	} else {
		goto L299
	}
L287:
	;
	v1001 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_21))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L62
	} else {
		goto L297
	}
L288:
	;
	v995 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_22))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L62
	} else {
		goto L295
	}
L289:
	;
	v989 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_23))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L62
	} else {
		goto L293
	}
L290:
	;
	v983 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_24))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L62
	} else {
		goto L291
	}
L291:
	;
	if int32(0) <= v983 {
		goto L285
	} else {
		goto L292
	}
L292:
	;
	v1021 = v983
	goto L18
L293:
	;
	if int32(0) <= v989 {
		goto L285
	} else {
		goto L294
	}
L294:
	;
	v1021 = v989
	goto L18
L295:
	;
	if int32(0) <= v995 {
		goto L285
	} else {
		goto L296
	}
L296:
	;
	v1021 = v995
	goto L18
L297:
	;
	if int32(0) <= v1001 {
		goto L285
	} else {
		goto L298
	}
L298:
	;
	v1021 = v1001
	goto L18
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977 + int32(1)
	goto L285
L300:
	;
	if int32(0) <= v1016 {
		goto L17
	} else {
		goto L301
	}
L301:
	;
	v1021 = v1016
	goto L18
}
func F_getIthJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8&int32(1073741824) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v8 & int32(268435455)
	if base.Ui32(l1) < base.Ui32(v12) {
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
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L41
	}
L4:
	;
	v18 = l0 + v12<<(uint(int32(2))%32) + int32(4)
	v20 = F_palloc(m, int32(20))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v181 = v3
	goto L6
L6:
	;
	return v181
L7:
	;
	return int32(0)
L8:
	;
	v26 = l1
	v28 = v3
	goto L9
L9:
	;
	if int32(0) < v26 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v54 = l0 + l1<<(uint(int32(2))%32) + int32(4)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	switch int32(base.Ui32(v55)>>(uint(int32(28))%32)) & int32(7) {
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+v26<<(uint(int32(2))%32))))
	v39 = v36&int32(268435455) + v28
	if int32(0) <= v36 {
		v26 = v26 - int32(1)
		v28 = v39
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v45 = v28
	goto L13
L13:
	;
	goto L10
L14:
	;
	v45 = v39
	goto L13
L15:
	;
	v181 = v20
	goto L6
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(18)
	v122 = (v45 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v18 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v125 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(3)
	goto L15
L18:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(3)
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v18 + (v45+int32(3))&int32(-4)
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v18 + v45
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v66 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	goto L15
L22:
	;
	v71 = l1
	v72 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v66 & int32(268435455)
	goto L15
L25:
	;
	v79 = v71 - int32(1)
	if int32(0) <= v79 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v66&int32(268435455) - v92
	goto L15
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0+v71<<(uint(int32(2))%32))))
	v88 = v85&int32(268435455) + v72
	if int32(0) <= v85 {
		v71 = v79
		v72 = v88
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v92 = v72
	goto L29
L29:
	;
	goto L26
L30:
	;
	v92 = v88
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v166 + (v45 - v122)
	goto L15
L32:
	;
	v130 = l1
	v131 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v166 = v125 & int32(268435455)
	goto L31
L35:
	;
	v138 = v130 - int32(1)
	if int32(0) <= v138 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v166 = v125&int32(268435455) - v151
	goto L31
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0+v130<<(uint(int32(2))%32))))
	v147 = v144&int32(268435455) + v131
	if int32(0) <= v144 {
		v130 = v138
		v131 = v147
		goto L35
	} else {
		goto L40
	}
L38:
	;
	v151 = v131
	goto L39
L39:
	;
	goto L36
L40:
	;
	v151 = v147
	goto L39
L41:
	;
	F_errmsg_internal(m, int32(_a_F_getIthJsonbValueFromContainer_0), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_getIthJsonbValueFromContainer_1), int32(479), int32(_a_F_getIthJsonbValueFromContainer_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
	F_relation_close(m, v27, int32(1))
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
		v11 = F_ArrayGetNItemsSafe(m, int32(1), l0+int32(16))
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
						F_errmsg(m, int32(_a_F_getWeights_0), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_getWeights_1), int32(420), int32(_a_F_getWeights_2))
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
								F_errmsg(m, int32(_a_F_getWeights_3), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getWeights_1), int32(425), int32(_a_F_getWeights_2))
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
													F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
												F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
														F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
											F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
														F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
													F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
															F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
														F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
													F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
															F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
												F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
															F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
														F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
																F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
										F_errmsg(m, int32(_a_F_getWeights_4), int32(0))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_getWeights_1), int32(434), int32(_a_F_getWeights_2))
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
				F_errmsg(m, int32(_a_F_getWeights_5), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getWeights_1), int32(415), int32(_a_F_getWeights_2))
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
					F_errmsg_internal(m, int32(_a_F_get_attname_0), v9)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_attname_1), int32(937), int32(_a_F_get_attname_2))
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
	var v17 int32
	_ = v17
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
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
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
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
	v17 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v23 = l0 + int32(8) + v17*int32(20)
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
	v40 = v17 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v40) <= base.Ui32(v41) {
		v17 = v40
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
	v51 = int32(14)
	v54 = base.I32_clz(l1) ^ int32(31)
	if base.Ui32(v51) <= base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v62 = int32(0)
	goto L14
L14:
	;
	v65 = v62
	goto L19
L15:
	;
	v57 = v51
	goto L17
L16:
	;
	v57 = v54
	goto L17
L17:
	;
	v62 = v57 + int32(1)
	goto L14
L18:
	;
	return v120
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v65<<(uint(int32(2))%32))+160))
	if v74 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v120 = int32(0)
	goto L18
L21:
	;
	v77 = int32(1)
	v80 = v77 << (uint(v65-v77) % 32)
	v85 = v74
	goto L24
L22:
	;
	goto L23
L23:
	;
	v113 = v65 + int32(1)
	if v113 != int32(16) {
		v65 = v113
		goto L19
	} else {
		goto L36
	}
L24:
	;
	v88 = F_get_segment_by_index(m, l0, v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	if base.B2i32(base.Ui32(v80) <= base.Ui32(v93))&base.B2i32(base.Ui32(v93) < base.Ui32(l1)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui32(v93) < base.Ui32(v80) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v91 != int32(-1) {
		v85 = v91
		goto L24
	} else {
		goto L35
	}
L30:
	;
	F_rebin_segment(m, l0, v88)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(l1) <= base.Ui32(v93) {
		v120 = v88
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
	F_make_relative_path(m, l0, int32(_a_F_get_doc_path_0), int32(_a_F_get_doc_path_1))
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
	F_make_relative_path(m, l1, int32(_a_F_get_etc_path_0), l0)
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
	var v4 float64
	_ = v4
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
	var v22 float64
	_ = v22
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
	var v47 float64
	_ = v47
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
	var v80 float64
	_ = v80
	v4 = float64(0)
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
	return v80
L2:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v80 = base.F64_convert_i32_u(v75)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L19
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v37 == int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v12 == int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v19 = int32(0)
	v22 = v4
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32))))
	v28 = F_get_indexpath_pages(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v80 = v32
	goto L1
L10:
	;
	return float64(0)
L11:
	;
	v32 = base.F64_add(v22, v28)
	v34 = v19 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v34 < v35 {
		v19 = v34
		v22 = v32
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v44 = int32(0)
	v47 = v4
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v44<<(uint(int32(2))%32))))
	v53 = F_get_indexpath_pages(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v80 = v55
	goto L1
L17:
	;
	v55 = base.F64_add(v47, v53)
	v57 = v44 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v57 < v58 {
		v44 = v57
		v47 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
	F_errmsg_internal(m, int32(_a_F_get_indexpath_pages_0), v7)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_get_indexpath_pages_1), int32(1003), int32(_a_F_get_indexpath_pages_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
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
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v403 int32
	_ = v403
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 float64
	_ = v624
	var v625 float64
	_ = v625
	var v626 float64
	_ = v626
	var v627 int32
	_ = v627
	var v628 float64
	_ = v628
	var v630 float64
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v653 int32
	_ = v653
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = F_bms_union(m, v26, l5)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v35 = F_bms_union(m, v33, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v37 = v8
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v37 = v35
	goto L8
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v42 = F_bms_union(m, v40, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v44 = v8
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v45 == int32(0) {
		v147 = v8
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v44 = v42
	goto L12
L14:
	;
	v156 = F_generate_join_implied_equalities(m, l0, v27, l5, l1, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L38
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 <= int32(0) {
		v147 = v8
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(0)
	v65 = v8
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v59<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v81 = F_bms_is_subset(m, v80, v27)
	mBase = m.M
	if v81 == v79 {
		v92 = v79
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v147 = v129
	goto L14
L19:
	;
	v131 = v59 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v131 < v132 {
		v59 = v131
		v65 = v129
		goto L17
	} else {
		goto L36
	}
L20:
	;
	if v92 == int32(0) {
		v129 = v65
		goto L19
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v85 = F_bms_overlap(m, v78, v84)
	mBase = m.M
	if v85 == int32(0) {
		v92 = v79
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	v89 = F_bms_overlap(m, v78, v88)
	mBase = m.M
	v92 = v89 ^ int32(1)
	goto L21
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v99 = F_bms_is_subset(m, v98, v37)
	mBase = m.M
	if v99 == v97 {
		v110 = v97
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v110 != 0 {
		v129 = v65
		goto L19
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v103 = F_bms_overlap(m, v96, v102)
	mBase = m.M
	if v103 == int32(0) {
		v110 = v97
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	v107 = F_bms_overlap(m, v96, v106)
	mBase = m.M
	v110 = v107 ^ int32(1)
	goto L26
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v113 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v115 = F_bms_is_subset(m, v114, v44)
	mBase = m.M
	if v115 == v113 {
		v126 = v113
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v126 != 0 {
		v129 = v65
		goto L19
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v119 = F_bms_overlap(m, v112, v118)
	mBase = m.M
	if v119 == int32(0) {
		v126 = v113
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	v123 = F_bms_overlap(m, v112, v122)
	mBase = m.M
	v126 = v123 ^ int32(1)
	goto L31
L34:
	;
	v127 = F_lappend(m, v65, v77)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v129 = v127
	goto L19
L36:
	;
	goto L18
L37:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v508 = F_list_concat(m, v499, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L114
	}
L38:
	;
	if v156 == int32(0) {
		v499 = v147
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v160 < v161 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v172 = int32(0)
	v173 = v160
	v178 = v147
	goto L43
L41:
	;
	v242 = v160
	v247 = v147
	goto L42
L42:
	;
	if v242 == int32(0) {
		v499 = v247
		goto L37
	} else {
		goto L61
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v172<<(uint(int32(2))%32))))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	v193 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v195 = F_bms_is_subset(m, v194, v37)
	mBase = m.M
	if v195 == v193 {
		v206 = v193
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v242 = v228
	v247 = v229
	goto L42
L45:
	;
	v231 = v172 + int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v231 < v232 {
		v172 = v231
		v173 = v228
		v178 = v229
		goto L43
	} else {
		goto L60
	}
L46:
	;
	if v206 != 0 {
		v228 = v173
		v229 = v178
		goto L45
	} else {
		goto L50
	}
L47:
	;
	goto L46
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v199 = F_bms_overlap(m, v192, v198)
	mBase = m.M
	if v199 == int32(0) {
		v206 = v193
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v190)+40))
	v203 = F_bms_overlap(m, v192, v202)
	mBase = m.M
	v206 = v203 ^ int32(1)
	goto L47
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	v209 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v211 = F_bms_is_subset(m, v210, v44)
	mBase = m.M
	if v211 == v209 {
		v222 = v209
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v222 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v215 = F_bms_overlap(m, v208, v214)
	mBase = m.M
	if v215 == int32(0) {
		v222 = v209
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v190)+40))
	v219 = F_bms_overlap(m, v208, v218)
	mBase = m.M
	v222 = v219 ^ int32(1)
	goto L52
L55:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v190)+100))
	v224 = F_lappend(m, v173, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v226 = F_lappend(m, v178, v190)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L59
	}
L58:
	;
	v228 = v224
	v229 = v178
	goto L45
L59:
	;
	v228 = v173
	v229 = v226
	goto L45
L60:
	;
	goto L44
L61:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v259 = F_bms_union(m, v258, l5)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	if v428 == int32(0) {
		v499 = v247
		goto L37
	} else {
		goto L101
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if base.Ui32(int32(5)) < base.Ui32(v263) {
		v275 = v262
		v276 = v259
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v277 = int32(0)
	if v242 == v277 {
		v428 = v277
		goto L62
	} else {
		goto L68
	}
L65:
	;
	if int32(1)<<(uint(v263)%32)&int32(44) == int32(0) {
		v275 = v262
		v276 = v259
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)+228))
	v273 = F_bms_union(m, l5, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v275 = v272
	v276 = v273
	goto L64
L68:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if int32(0) < v280 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v296 = int32(0)
	v301 = v8
	goto L72
L70:
	;
	v403 = v8
	goto L71
L71:
	;
	v428 = v403
	goto L62
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305+v296<<(uint(int32(2))%32))))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+40)))
	if v310 != 0 {
		v381 = v301
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v403 = v381
	goto L71
L74:
	;
	v383 = v296 + int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v383 < v384 {
		v296 = v383
		v301 = v381
		goto L72
	} else {
		goto L100
	}
L75:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v309)+16))
	if v311 == int32(0) {
		v381 = v301
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v314 < int32(2) {
		v381 = v301
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v309)+36))
	v318 = int32(0)
	if base.B2i32(v317 == v318)|base.B2i32(v276 == v318) != 0 {
		v363 = v318
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v363 == int32(0) {
		v381 = v301
		goto L74
	} else {
		goto L91
	}
L79:
	;
	goto L78
L80:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v328 < v329 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v331 = v328
	goto L83
L82:
	;
	v331 = v329
	goto L83
L83:
	;
	if v331 <= int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v334 = int32(1)
	goto L86
L85:
	;
	v334 = v331
	goto L86
L86:
	;
	v335 = int32(8)
	v340 = int32(0)
	goto L87
L87:
	;
	v347 = v340 << (uint(int32(2)) % 32)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v276+v335+v347)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v317+v335+v347)))
	v352 = v349 & v351
	v354 = base.B2i32(v352 != int32(0))
	if v352 != 0 {
		v363 = v354
		goto L79
	} else {
		goto L89
	}
L88:
	;
	v363 = v354
	goto L79
L89:
	;
	v356 = v340 + int32(1)
	if v356 != v334 {
		v340 = v356
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+42)))
	if v366 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v378 = F_list_concat(m, v301, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L99
	}
L93:
	;
	v369 = F_generate_join_implied_equalities_normal(m, l0, v309, v259, l5, v262)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v375 = F_generate_join_implied_equalities_broken(m, l0, v309, v276, l5, v275, v261)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+42)))
	if v371 != int32(1) {
		v377 = v369
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v377 = v375
	goto L92
L99:
	;
	v381 = v378
	goto L74
L100:
	;
	goto L73
L101:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	if v431 <= int32(0) {
		v499 = v247
		goto L37
	} else {
		goto L102
	}
L102:
	;
	v442 = int32(0)
	v448 = v247
	goto L103
L103:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v428)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v442<<(uint(int32(2))%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v463 = int32(0)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v460)+28))
	v465 = F_bms_is_subset(m, v464, v37)
	mBase = m.M
	if v465 == v463 {
		v476 = v463
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v499 = v481
	goto L37
L105:
	;
	if v476 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L105
L107:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v460)+28))
	v469 = F_bms_overlap(m, v462, v468)
	mBase = m.M
	if v469 == int32(0) {
		v476 = v463
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v460)+40))
	v473 = F_bms_overlap(m, v462, v472)
	mBase = m.M
	v476 = v473 ^ int32(1)
	goto L106
L109:
	;
	v479 = F_lappend(m, v448, v460)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	v481 = v448
	goto L111
L111:
	;
	v483 = v442 + int32(1)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	if v483 < v484 {
		v442 = v483
		v448 = v481
		goto L103
	} else {
		goto L113
	}
L112:
	;
	v481 = v479
	goto L111
L113:
	;
	goto L104
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v508
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v511 == int32(0) {
		v610 = v508
		goto L116
	} else {
		goto L117
	}
L115:
	;
	return v653
L116:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v624 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v625 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v626 = F_calc_joinrel_size_estimate(m, l0, l1, v622, v623, v624, v625, l4, v610)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L134
	}
L117:
	;
	v514 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v515 <= v514 {
		v610 = v508
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v525 = v514
	goto L119
L119:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v511)+12))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v539+v525<<(uint(int32(2))%32))))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v545 = int32(0)
	if base.B2i32(v544 == v545)|base.B2i32(l5 == v545) != 0 {
		v591 = base.B2i32(v544|l5 == v545)
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v610 = v600
	goto L116
L121:
	;
	if v591 != 0 {
		v653 = v543
		goto L115
	} else {
		goto L132
	}
L122:
	;
	goto L121
L123:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v559 != v560 {
		v591 = int32(0)
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v562 = int32(1)
	if v559 <= v562 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v565 = v562
	goto L127
L126:
	;
	v565 = v559
	goto L127
L127:
	;
	v566 = int32(8)
	v571 = int32(0)
	goto L128
L128:
	;
	v579 = v571 << (uint(int32(2)) % 32)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v544+v566+v579)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l5+v566+v579)))
	v584 = base.B2i32(v581 == v583)
	if v581 != v583 {
		v591 = v584
		goto L122
	} else {
		goto L130
	}
L129:
	;
	v591 = v584
	goto L122
L130:
	;
	v587 = v571 + int32(1)
	if v587 != v565 {
		v571 = v587
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v597 = v525 + int32(1)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v597 < v598 {
		v525 = v597
		goto L119
	} else {
		goto L133
	}
L133:
	;
	goto L120
L134:
	;
	v628 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v626, v628) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v630 = v628
	goto L137
L136:
	;
	v630 = v626
	goto L137
L137:
	;
	v632 = F_palloc0(m, int32(24))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v632)+16)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v632)+8)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v632)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = int32(278)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v641 = F_lappend(m, v640, v632)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v641
	v653 = v632
	goto L115
}
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
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
	var v81 int32
	_ = v81
	var v89 float64
	_ = v89
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
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 float64
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
	var v230 int32
	_ = v230
	var v236 float64
	_ = v236
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
	var v270 int32
	_ = v270
	var v271 float64
	_ = v271
	var v274 float64
	_ = v274
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
	var v344 float64
	_ = v344
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
	var v363 float64
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v380 float64
	_ = v380
	var v384 float64
	_ = v384
	var v387 float64
	_ = v387
	var v400 float64
	_ = v400
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
	var v471 float64
	_ = v471
	var v475 float64
	_ = v475
	v11 = float64(0)
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
	v81 = v74
	v89 = v11
	goto L18
L16:
	;
	v471 = v11
	goto L17
L17:
	;
	if base.F64_gt(v471, float64(0)) != 0 {
		goto L111
	} else {
		goto L112
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v90 <= v81 {
		v400 = v89
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v471 = v400
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v81<<(uint(int32(2))%32))))
	if v96 == int32(0) {
		v400 = v89
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v99 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	if v101 == v99 {
		v122 = v99
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v122 != 0 {
		v400 = v89
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
	v122 = int32(1)
	goto L24
L28:
	;
	if v109 != int32(290) {
		v122 = v99
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
		v122 = v99
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
		v380 = v123
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.F64_lt(v380, v89) != 0 {
		goto L93
	} else {
		goto L94
	}
L35:
	;
	v127 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v128 <= v127 {
		v380 = v123
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v138 = v127
	v141 = v123
	goto L37
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v138<<(uint(int32(2))%32))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	if v149 != int32(4) {
		v363 = v141
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v380 = v363
	goto L34
L39:
	;
	v367 = v138 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v367 < v368 {
		v138 = v367
		v141 = v363
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
		v363 = v141
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v160 = F_bms_is_member(m, v81, v159)
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
		v363 = v141
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
	v230 = v222
	v236 = v164
	goto L60
L58:
	;
	v344 = v164
	goto L59
L59:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
	v347 = int32(0)
	v349 = F_estimate_num_groups(m, l0, v346, v344, v347, v347)
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
	if v238 <= v230 {
		v274 = v236
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v344 = v274
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
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v230<<(uint(int32(2))%32))))
	if v244 == int32(0) {
		v274 = v236
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v247 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	if v249 == v247 {
		v270 = v247
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v270 != 0 {
		v274 = v236
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
	v270 = int32(1)
	goto L66
L70:
	;
	if v257 != int32(290) {
		v270 = v247
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
		v270 = v247
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
	v274 = base.F64_mul(v236, v271)
	goto L62
L76:
	;
	if int32(0) <= v330 {
		v230 = v330
		v236 = v274
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
	v281 = v230 + int32(1)
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
	if base.F64_gt(v141, v349) != 0 {
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
	v352 = v141
	goto L91
L91:
	;
	v363 = v352
	goto L39
L92:
	;
	goto L38
L93:
	;
	v384 = v380
	goto L95
L94:
	;
	v384 = v89
	goto L95
L95:
	;
	if base.F64_eq(v89, float64(0)) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v387 = v380
	goto L98
L97:
	;
	v387 = v384
	goto L98
L98:
	;
	v400 = v387
	goto L20
L99:
	;
	if int32(0) <= v456 {
		v81 = v456
		v89 = v400
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
	v407 = v81 + int32(1)
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
	v475 = v471
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	if base.B2i32(v14 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v13)) != 0 {
		v9 = v13
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v20 = v13
	goto L7
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(l0) < base.Ui32(v20) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v32 = v20
	goto L13
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 != int32(47) {
		v20 = v20 - int32(1)
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
	goto L11
L13:
	;
	if base.Ui32(l0) < base.Ui32(v32) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l0 == v32 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v36 = v32 - int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 == int32(47) {
		v32 = v36
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L17
L19:
	;
	v45 = l0 + base.B2i32(v5 == int32(47))
	goto L21
L20:
	;
	v45 = v32
	goto L21
L21:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v46)
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
	var v98 int32
	_ = v98
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
		goto L27
	} else {
		goto L28
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
	v25 = int32(_a_F_get_password_type_0)
	v29 = m.G0
	v31 = v29 - int32(32)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v32
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_password_type[0])))
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
		goto L26
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
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_password_type[1])))
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
		v98 = v24
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v108 = v98 - v24
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
		v98 = v79
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v98 = v96
	goto L20
L24:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v96 = v79 + int32(1)
	if v94 != 0 {
		v79 = v96
		v80 = v94
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L2
L27:
	;
	return int32(0)
L28:
	;
	if v124 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v128 = int32(2)
	goto L31
L30:
	;
	v128 = int32(0)
	goto L31
L31:
	;
	v129 = v128
	goto L1
}
func F_get_ps_display(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return int32(_a_F_get_ps_display_0)
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
			if l1|v15 != 0 {
				v75 = v15
				m.G0 = v7 + int32(48)
				return v75
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v11
						F_errmsg(m, int32(_a_F_get_rolespec_oid_0), v7+int32(16))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_2), int32(_a_F_get_rolespec_oid_3))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
	case 1, 2:
		v74 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_oid[0]))
		v75 = v74
		m.G0 = v7 + int32(48)
		return v75
	case 3:
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_oid[1]))
		v75 = v39
		m.G0 = v7 + int32(48)
		return v75
	case 4:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_get_rolespec_oid_4)
				F_errmsg(m, int32(_a_F_get_rolespec_oid_0), v7+int32(32))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_5), int32(_a_F_get_rolespec_oid_6))
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v63
			F_errmsg_internal(m, int32(_a_F_get_rolespec_oid_7), v7)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_8), int32(_a_F_get_rolespec_oid_6))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
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
				F_errmsg_internal(m, int32(_a_F_get_segment_by_index_0), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_segment_by_index_1), int32(1781), int32(_a_F_get_segment_by_index_2))
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
			v19 = int32(_a_F_get_segment_by_index_3)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_get_segment_by_index[0]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_get_segment_by_index[0])) = v22
			v24 = F_dsm_attach(m, v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_get_segment_by_index[0])) = v20
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_get_segment_by_index_4), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_segment_by_index_1), int32(1788), int32(_a_F_get_segment_by_index_2))
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v3 = int32(0)
	if l0 == v3 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L21
	}
L5:
	;
	v20 = v3
	v21 = v3
	goto L8
L6:
	;
	v70 = v3
	goto L7
L7:
	;
	return v70
L8:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v70 = v58
	goto L7
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 <= int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v20<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = int32(0)
	goto L12
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37+v41<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v36 != v52 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v58 = F_lappend(m, v21, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v55 = v41 + int32(1)
	if v55 != v28 {
		v41 = v55
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
	goto L4
L18:
	;
	return int32(0)
L19:
	;
	v63 = v20 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 < v64 {
		v20 = v63
		v21 = v58
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	F_errmsg_internal(m, int32(_a_F_get_sortgrouplist_exprs_0), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_sortgrouplist_exprs_1), int32(366), int32(_a_F_get_sortgrouplist_exprs_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_get_sortgroupref_tle_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_get_sortgroupref_tle_1), int32(366), int32(_a_F_get_sortgroupref_tle_2))
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
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	v3 = int32(0)
	if l0 == v3 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v3
	v27 = v3
	goto L7
L5:
	;
	v201 = v3
	goto L6
L6:
	;
	return v201
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v36 = int32(0)
	if v35 == v36 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v201 = v185
	goto L6
L9:
	;
	v185 = F_lappend(m, v27, v180)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L27
	} else {
		goto L45
	}
L10:
	;
	if v89 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v89 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	if l1 == int32(0) {
		v82 = v36
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v89 = v82
	goto L10
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v46 < v45 {
		v82 = v36
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v48 = int32(1)
	if v45 <= v48 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v51 = v48
	goto L19
L18:
	;
	v51 = v45
	goto L19
L19:
	;
	v52 = int32(8)
	v57 = int32(0)
	goto L20
L20:
	;
	v64 = v57 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35+v52+v64)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1+v52+v64)))
	v71 = v66 & (v68 ^ int32(-1))
	v73 = base.B2i32(v71 == int32(0))
	if v71 != 0 {
		v82 = v73
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v82 = v73
	goto L14
L22:
	;
	v75 = v57 + int32(1)
	if v75 != v51 {
		v57 = v75
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v180 = v33
	v184 = int32(1)
	goto L9
L25:
	;
	goto L26
L26:
	;
	v93 = F_palloc0(m, int32(36))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(17)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v104
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+16)) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v113 = F_list_copy(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v116
	v118 = m.G0
	v120 = v118 - int32(16)
	m.G0 = v120
	if v93 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v180 = v93
	v184 = v100
	goto L9
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L27
	} else {
		goto L42
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L27
	} else {
		goto L39
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v124 != int32(17) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	if v127 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v130 != int32(2) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v134 = F_get_commutator(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	if v134 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v134
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v143
	m.G0 = v120 + int32(16)
	goto L30
L39:
	;
	F_errmsg_internal(m, int32(_a_F_get_switched_clauses_0), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_get_switched_clauses_1), int32(2156), int32(_a_F_get_switched_clauses_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v170
	F_errmsg_internal(m, int32(_a_F_get_switched_clauses_3), v120)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_get_switched_clauses_1), int32(2162), int32(_a_F_get_switched_clauses_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+120)) = uint8(v184)
	v189 = v24 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v189 < v190 {
		v24 = v189
		v27 = v185
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
			v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12+v13)+76)))
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
			v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+v13)+129)))
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v8, int32(_a_F_get_values_def_0))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(1)
	v24 = int32(0)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v24<<(uint(int32(2))%32))))
	if v20 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoString(m, v8, int32(_a_F_get_values_def_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v37 = m.ExcPending
	if v37 != 0 {
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
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v40 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v56 = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v57 <= v56 {
		goto L13
	} else {
		goto L22
	}
L17:
	;
	F_get_rule_expr(m, v44, l1, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v47 != int32(6) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = F_get_variable(m, v44, int32(1), l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v65 = v56
	goto L23
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v65<<(uint(int32(2))%32))))
	F_appendStringInfoChar(m, v8, int32(44))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L13
L25:
	;
	if v71 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v87 = v65 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v87 < v88 {
		v65 = v87
		goto L23
	} else {
		goto L32
	}
L27:
	;
	F_get_rule_expr(m, v71, l1, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v77 != int32(6) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v81 = F_get_variable(m, v71, int32(1), l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v102 = v24 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 < v103 {
		v20 = int32(0)
		v24 = v102
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
	F_errmsg_internal(m, int32(_a_F_get_view_query_0), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_get_view_query_1), int32(2498), int32(_a_F_get_view_query_2))
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
	F_errmsg_internal(m, int32(_a_F_get_view_query_3), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_get_view_query_1), int32(2504), int32(_a_F_get_view_query_2))
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v3 = int32(0)
	if l0 == v3 {
		v60 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v60
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(0) {
		v60 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		v60 = v3
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
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v28 == int32(0))|base.B2i32(v28 != v31) != 0 {
		v49 = v28
		v50 = v31
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v60 = int32(0)
	goto L1
L7:
	;
	if v49-v50 == int32(0) {
		v60 = v24
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v34 = v25
	v35 = l1
	goto L10
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v39 == int32(0) {
		v49 = v39
		v50 = v38
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v49 = v39
	v50 = v38
	goto L8
L12:
	;
	v42 = int32(1)
	if v39 == v38 {
		v34 = v34 + v42
		v35 = v35 + v42
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v55 = v16 + int32(1)
	if v11 != v55 {
		v16 = v55
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
}
func F_getenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v182 int32
	_ = v182
	v2 = int32(0)
	goto L5
L1:
	;
	if l0 == v91 {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	goto L1
L3:
	;
	v81 = v76
	goto L20
L4:
	;
	v76 = v68
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
	v14 = l0
	goto L11
L9:
	;
	v28 = l0
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 != v37 {
		v68 = v28
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(v19 == int32(0))|base.B2i32(int32(61) == v19) != 0 {
		v91 = v14
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v28 = v25
	goto L10
L13:
	;
	v25 = v14 + int32(1)
	if v25&int32(3) != 0 {
		v14 = v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v43 = v28
	v45 = v34
	goto L16
L16:
	;
	v49 = v45 ^ int32(1027423549)
	v52 = int32(-2139062144)
	if (int32(16843008)-v49|v49)&v52 != v52 {
		v68 = v43
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v76 = v58
	goto L3
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v58 = v43 + int32(4)
	v62 = int32(-2139062144)
	if (v56|(int32(16843008)-v56))&v62 == v62 {
		v43 = v58
		v45 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v83 == int32(0) {
		v91 = v81
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v91 = v81
	goto L2
L22:
	;
	if v83 != int32(61) {
		v81 = v81 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	goto L26
L26:
	;
	v105 = v91 - l0
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v105))))
	if v107 != 0 {
		v182 = v2
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return v182
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_getenv[0]))
	if v109 == int32(0) {
		v182 = v2
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v112 == int32(0) {
		v182 = v2
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v116 = v109
	v117 = v112
	goto L31
L31:
	;
	if v105 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v182 = v168 + int32(1)
	goto L27
L33:
	;
	goto L32
L34:
	;
	if v164 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L35:
	;
	v164 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v125 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = l0
	v127 = v117
	v128 = v105
	v129 = v125
	goto L42
L39:
	;
	v152 = v117
	v156 = int32(0)
	goto L40
L40:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v164 = v156 - v157
	goto L34
L41:
	;
	v152 = v147
	v156 = v149
	goto L40
L42:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if base.B2i32(v129 != v131)|base.B2i32(v131 == int32(0)) != 0 {
		v147 = v127
		v149 = v129
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v147 = v141
	v149 = int32(0)
	goto L41
L44:
	;
	v137 = v128 - int32(1)
	if v137 == int32(0) {
		v147 = v127
		v149 = v129
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v140 = int32(1)
	v141 = v127 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v142 != 0 {
		v126 = v126 + v140
		v127 = v141
		v128 = v137
		v129 = v142
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v168 = v167 + v105
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169 == int32(61) {
		goto L33
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v173 != 0 {
		v116 = v116 + int32(4)
		v117 = v173
		goto L31
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v182 = v2
	goto L27
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
		v16 = v10
		v17 = int32(0)
		v18 = v7
		for {
			if base.Ui32(v17) <= base.Ui32(int32(214748364)) {
				v26 = v17 * int32(10)
				if base.Ui32(v26^int32(2147483647)) < base.Ui32(v16) {
					v31 = int32(-1)
				} else {
					v31 = v16 + v26
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
				v16 = v39
				v17 = v33
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
func F_getrlimit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	if l1 != 0 {
		switch l0 - int32(3) {
		case 0:
			v10 = m.G2
			v11 = m.G1
			v13 = base.I64_extend_i32_u(v10 - v11)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v13
		default:
			v16 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v16
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16
		case 4:
			v6 = int64(4096)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v6
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v6
		}
	} else {
	}
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
	var v22 int32
	_ = v22
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
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v263 int32
	_ = v263
	v7 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v13 = l0 + int32(8)
	goto L3
L1:
	;
	return v263
L2:
	;
	v263 = int32(3)
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v22 - int32(1) {
	case 0, 2:
		goto L8
	case 1:
		goto L7
	default:
		v233 = v21
		goto L6
	}
L4:
	;
	v241 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21 + v241
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v241)
	goto L2
L5:
	;
	goto L4
L6:
	;
	v236 = F_pg_mblen_cstr(m, v233)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L20
	} else {
		goto L63
	}
L7:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v111 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	switch v25 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v233 = v21
		goto L6
	default:
		goto L9
	case 24, 29, 31, 32, 51, 115:
		goto L10
	case 25:
		goto L11
	case 36:
		goto L5
	}
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v21
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v75 = int32(0)
	v77 = F_gettoken_tsvector(m, v74, l3, l2, v75, v75, v13)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21 + v67
	goto L3
L11:
	;
	v29 = v21 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v37 = v32
	goto L12
L12:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v41 = int32(0)
	if base.B2i32(v40 == v41)|base.B2i32(v40 == int32(34)) == v41 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37 - v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v49 = v37 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v49
	v37 = v49
	goto L12
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v54 + int32(1)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v59 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v61 + int32(1)
	return v59
L20:
	;
	return int32(0)
L21:
	;
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v81 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v81
	return v81
L23:
	;
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v85 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v97 == int32(3) {
		v263 = int32(0)
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v88 != int32(447) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v91 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	return int32(1)
L29:
	;
	v101 = F_palloc0(m, int32(12))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v103 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v103)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v106 = F_lcons(m, v101, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v106
	return int32(0)
L32:
	;
	return int32(0)
L33:
	;
	goto L34
L34:
	;
	v120 = v21
	v121 = int32(_a_F_gettoken_query_websearch_0)
	v122 = int32(2)
	goto L37
L35:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	switch v219 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v233 = v218
		goto L6
	default:
		goto L61
	case 24, 29, 31, 32, 51, 115:
		goto L62
	}
L36:
	;
	if v167 != 0 {
		goto L35
	} else {
		goto L52
	}
L37:
	;
	if v122 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v167 = int32(0)
	goto L36
L39:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v125 == v126 {
		v148 = v125
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v150 = int32(1)
	if v148 != 0 {
		v120 = v120 + v150
		v121 = v121 + v150
		v122 = v122 - v150
		goto L37
	} else {
		goto L51
	}
L43:
	;
	if base.Ui32((v125-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v136 = v125 | int32(32)
	goto L46
L45:
	;
	v136 = v125
	goto L46
L46:
	;
	if base.Ui32((v126-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v145 = v126 | int32(32)
	goto L49
L48:
	;
	v145 = v126
	goto L49
L49:
	;
	if v136 == v145 {
		v148 = v136
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v167 = v136 - v145
	goto L36
L51:
	;
	goto L41
L52:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
	if base.B2i32(v168 == int32(0))|base.B2i32(v168 == int32(45))|base.B2i32(v168 == int32(95)) != 0 {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	v178 = v21 + int32(2)
	v179 = F_t_isalnum_cstr(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	if v179 != 0 {
		goto L35
	} else {
		goto L55
	}
L55:
	;
	v185 = v178
	goto L56
L56:
	;
	v188 = F_pg_mblen_cstr(m, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L58
	}
L57:
	;
	if v191 == int32(0) {
		goto L35
	} else {
		goto L60
	}
L58:
	;
	v190 = v188 + v185
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if base.B2i32(base.Ui32(v191-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v191 == int32(32)) != 0 {
		v185 = v190
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v203 + int32(2)
	v207 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v207)
	return v207
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v227 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v227)
	goto L2
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v218 + int32(1)
	goto L3
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v236 + v238
	goto L3
}
func F_ginbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v312 int32
	_ = v312
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 float64
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v561 int32
	_ = v561
	var v574 int32
	_ = v574
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v664 int32
	_ = v664
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int64
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v789 int32
	_ = v789
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v860 int32
	_ = v860
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v971 int32
	_ = v971
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1111 int32
	_ = v1111
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 float64
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1140 float64
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1513 int64
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int64
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1724 int32
	_ = v1724
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	v32 = m.G0
	v34 = v32 - int32(_a_F_ginbulkdelete_0)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0]))
	v43 = F_AllocSetContextCreateInternal(m, v38, int32(_a_F_ginbulkdelete_1), int32(0), int32(_a_F_ginbulkdelete_2), int32(_a_F_ginbulkdelete_3))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2764)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2760)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2752)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[1]))) = v43
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[2]))) = v51
	v54 = v34 + int32(2768)
	F_initGinState(m, v54, v36)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	v60 = F_palloc0(m, int32(40))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v70 = l1
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2756)) = v70
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v78 = F_ReadBufferExtended(m, v36, v74, int32(1), v74, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[3]))
	F_ginInsertCleanup(m, v54, base.B2i32(v63 != int32(4)), int32(0), int32(1), v60)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v70 = v60
	goto L6
L9:
	;
	v84 = int32(1)
	v86 = v78
	goto L10
L10:
	;
	if int32(0) <= v86 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v312 = v86
	goto L39
L12:
	;
	F_LockBuffer(m, v86, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v129 = v115 + v86<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(v86^int32(-1))<<(uint(int32(2))%32))))
	v129 = v128
	goto L12
L16:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+16)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v133)+6)))
	if v135&int32(2) == int32(0) {
		v230 = v129
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L11
L18:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v230)+24))
	v263 = v230 + v260&int32(_a_F_ginbulkdelete_4)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+2)))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263))))
	F_UnlockReleaseBuffer(m, v86)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L37
	}
L19:
	;
	F_LockBuffer(m, v86, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_LockBuffer(m, v86, int32(2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v84 != int32(1) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+16)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v148)+6)))
	if v150&int32(2) != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_LockBuffer(m, v86, int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	if v86 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBuffer(m, v86, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196+(v86^int32(-1))<<(uint(int32(2))%32))))
	v204 = v198
	goto L27
L29:
	;
	goto L30
L30:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v204 = v200 + v86<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+16)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v208)+6)))
	if v210&int32(2) == int32(0) {
		v230 = v204
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_LockBuffer(m, v86, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_LockBuffer(m, v86, int32(2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+16)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v221)+6)))
	if v223&int32(2) != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	F_LockBuffer(m, v86, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v268 = int32(0)
	v271 = v264 | v265<<(uint(int32(16))%32)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v274 = F_ReadBufferExtended(m, v36, v268, v271, v268, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v84 = v271
	v86 = v274
	goto L10
L39:
	;
	v338 = int32(0)
	v339 = base.B2i32(v338 <= v312)
	if v339 == v338 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[1])))
	F_MemoryContextDelete(m, v1803)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L293
	}
L41:
	;
	v839 = int32(0)
	F_vacuum_delay_point(m, v839)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L127
	}
L42:
	;
	F_UnlockReleaseBuffer(m, v312)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L126
	}
L43:
	;
	v771 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+16)))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v357+v771)))
	v789 = int32(0)
	v805 = v773
	goto L42
L44:
	;
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+12)))
	if base.Ui32(v358) < base.Ui32(int32(25)) {
		goto L43
	} else {
		goto L48
	}
L45:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v343+(v312^int32(-1))<<(uint(int32(2))%32))))
	v357 = v349
	goto L44
L46:
	;
	goto L47
L47:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v357 = v351 + v312<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L48:
	;
	v366 = int32(base.Ui32(v358+int32(_a_F_ginbulkdelete_5))>>(uint(int32(2))%32)) & int32(_a_F_ginbulkdelete_6)
	if v366 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v373 = v357
	v378 = int32(1)
	v386 = int32(0)
	goto L51
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L123
	}
L51:
	;
	v403 = v378 & int32(_a_F_ginbulkdelete_6)
	v405 = v403 << (uint(int32(2)) % 32)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v373+v405)+20))
	v410 = v373 + v407&int32(_a_F_ginbulkdelete_4)
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410)+4)))
	if v411 == int32(0) {
		v651 = v373
		v664 = v386
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+16)))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v357+v685)))
	if base.B2i32(v651 == int32(0))|base.B2i32(v651 == v357) != 0 {
		v789 = v664
		v805 = v687
		goto L42
	} else {
		goto L105
	}
L53:
	;
	v681 = v378 + int32(1)
	if base.Ui32(v681&int32(_a_F_ginbulkdelete_6)) <= base.Ui32(v366) {
		v373 = v651
		v378 = v681
		v386 = v664
		goto L51
	} else {
		goto L104
	}
L54:
	;
	if v411 == int32(_a_F_ginbulkdelete_6) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v416 = int32(16)
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410)+2)))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410))))
	*(*int32)(unsafe.Add(mBase, uint32(v34+v416+v386<<(uint(int32(2))%32)))) = v421 | v422<<(uint(v416)%32)
	v651 = v373
	v664 = v386 + int32(1)
	goto L53
L56:
	;
	goto L57
L57:
	;
	v429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410))))
	v431 = v429 << (uint(int32(16)) % 32)
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410)+2)))
	v433 = v431 | v432
	if int32(0) <= v431 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v574 == int32(0) {
		v651 = v373
		v664 = v386
		goto L53
	} else {
		goto L82
	}
L59:
	;
	F_pfree(m, v537)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L81
	}
L60:
	;
	v455 = int32(0)
	v459 = v455
	v470 = v455
	v471 = v455
	goto L66
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6]))) = v411
	v451 = v411
	v452 = v433 + v410
	goto L60
L62:
	;
	goto L63
L63:
	;
	v443 = F_ginPostingListDecode(m, v410+v433&int32(2147483647), v34+int32(_a_F_ginbulkdelete_7))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6])))
	if int32(0) < v445 {
		v451 = v445
		v452 = v443
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v448 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6]))) = v448
	v537 = v443
	v541 = v448
	goto L59
L66:
	;
	v490 = v459 * int32(6)
	v491 = v452 + v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2764))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2760))
	v494 = m.T0[v493].(func(*base.Module, int32, int32) int32)(m, v491, v492)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6]))) = v522
	if int32(0) <= v431 {
		v574 = v521
		goto L58
	} else {
		goto L80
	}
L68:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2756))
	if v494 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v524 = v459 + int32(1)
	if v524 != v451 {
		v459 = v524
		v470 = v521
		v471 = v522
		goto L66
	} else {
		goto L79
	}
L70:
	;
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v496)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v496)+16)) = base.F64_add(v497, float64(1))
	if v470 != 0 {
		v521 = v470
		v522 = v471
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v496)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v496)+8)) = base.F64_add(v506, float64(1))
	if v470 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v501 = F_palloc(m, v451*int32(6))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v490 == int32(0) {
		v521 = v501
		v522 = v471
		goto L69
	} else {
		goto L75
	}
L75:
	;
	base.MemoryCopy(m, v501, v452, v490)
	v521 = v501
	v522 = v471
	goto L69
L76:
	;
	v512 = v470 + v471*int32(6)
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v512)+4)) = uint16(v513)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v515
	goto L78
L77:
	;
	goto L78
L78:
	;
	v521 = v470
	v522 = v471 + int32(1)
	goto L69
L79:
	;
	goto L67
L80:
	;
	v537 = v452
	v541 = v521
	goto L59
L81:
	;
	v574 = v541
	goto L58
L82:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6])))
	if v595 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v373 == v357 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v598 = int32(0)
	v611 = v598
	v612 = v598
	goto L83
L85:
	;
	goto L86
L86:
	;
	v602 = F_ginCompressPostingList(m, v574, v595, int32(2712), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602)+6)))
	v611 = v602
	v612 = (v604+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	goto L83
L88:
	;
	v614 = F_PageGetTempPageCopy(m, v357)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v621 = v373
	v622 = v410
	goto L90
L90:
	;
	v623 = F_gintuple_get_attrnum(m, v54, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614+v405)+20))
	v621 = v614
	v622 = v614 + v617&int32(_a_F_ginbulkdelete_4)
	goto L90
L92:
	;
	v627 = F_gintuple_get_key(m, v54, v622, v34+int32(_a_F_ginbulkdelete_9))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v629 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[7]))))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6])))
	v632 = F_GinFormTuple(m, v54, v623, v627, v629, v611, v612, v630, int32(1))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v611 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_pfree(m, v611)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	F_PageIndexTupleDelete(m, v621, v403)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v632)+6)))
	v642 = F_PageAddItemExtended(m, v621, v632, v638&int32(_a_F_ginbulkdelete_10), v403, int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v642 != v403 {
		goto L50
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v632)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v574)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v651 = v621
	v664 = v386
	goto L53
L104:
	;
	goto L52
L105:
	;
	v692 = int32(_a_F_ginbulkdelete_11)
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v694 + int32(1)
	F_PageRestoreTempPage(m, v651, v357)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_MarkBufferDirty(m, v312)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v339 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+48))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721)+118)))
	if v722 != int32(112) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v705+(v312^int32(-1))<<(uint(int32(2))%32))))
	v719 = v711
	goto L108
L110:
	;
	goto L111
L111:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v719 = v713 + v312<<(uint(int32(13))%32) + int32(-8192)
	goto L108
L112:
	;
	F_UnlockReleaseBuffer(m, v312)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L122
	}
L113:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v726 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v720)+32))
	if v729 != 0 {
		goto L112
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v720)+40))
	if v730 != 0 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	F_XLogRegisterBuffer(m, int32(0), v312, int32(9))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v739 = F_XLogInsert(m, int32(13), int32(64))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v719))) = base.I64_rotr(v739, int64(32))
	goto L112
L122:
	;
	v746 = int32(_a_F_ginbulkdelete_11)
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v748 - int32(1)
	v823 = v664
	v824 = v687
	goto L41
L123:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v757 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginbulkdelete_12), v34)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_ginbulkdelete_13), int32(552), int32(_a_F_ginbulkdelete_14))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v823 = v789
	v824 = v805
	goto L41
L127:
	;
	if v823 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v860 = v839
	goto L131
L129:
	;
	goto L130
L130:
	;
	if v824 != int32(-1) {
		goto L288
	} else {
		goto L289
	}
L131:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(16)+v860<<(uint(int32(2))%32))))
	v881 = v879
	goto L133
L132:
	;
	goto L130
L133:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v912 = int32(0)
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[2])))
	v915 = F_ReadBufferExtended(m, v911, v912, v881, v912, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	v952 = int32(0)
	F_LockBuffer(m, v915, v952)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L145
	}
L135:
	;
	F_LockBuffer(m, v915, int32(1))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v915 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+16)))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938+v937)+6)))
	if v940&int32(2) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v923+(v915^int32(-1))<<(uint(int32(2))%32))))
	v937 = v929
	goto L137
L139:
	;
	goto L140
L140:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v937 = v931 + v915<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L141:
	;
	v945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+34)))
	v946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+32)))
	F_UnlockReleaseBuffer(m, v915)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	goto L134
L144:
	;
	v881 = v945 | v946<<(uint(int32(16))%32)
	goto L133
L145:
	;
	F_LockBuffer(m, v915, int32(2))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v960 = v915
	v962 = v937
	v971 = v952
	goto L148
L147:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L286
	}
L148:
	;
	v990 = int32(_a_F_ginbulkdelete_15)
	v991 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0]))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0])) = v993
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v997 = v34 + int32(2752)
	v998 = int32(0)
	v999 = m.G0
	v1001 = v999 - int32(16)
	m.G0 = v1001
	if v960 < v998 {
		goto L155
	} else {
		goto L156
	}
L149:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v1634 = int32(0)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[2])))
	v1637 = F_ReadBufferExtended(m, v1633, v1634, v879, v1634, v1636)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L1
	} else {
		goto L275
	}
L150:
	;
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0])) = v991
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[1])))
	F_MemoryContextReset(m, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L256
	}
L152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L1
	} else {
		goto L253
	}
L153:
	;
	m.G0 = v1001 + int32(16)
	goto L151
L154:
	;
	v1021 = F_disassembleLeaf(m, v1020)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L158
	}
L155:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1006+(v960^int32(-1))<<(uint(int32(2))%32))))
	v1020 = v1012
	goto L154
L156:
	;
	goto L157
L157:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v1020 = v1014 + v960<<(uint(int32(13))%32) + int32(-8192)
	goto L154
L158:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	if base.B2i32(v1023 == int32(0))|base.B2i32(v1023 == v1021) != 0 {
		goto L153
	} else {
		goto L159
	}
L159:
	;
	v1030 = v1023
	v1035 = v998
	goto L161
L160:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1229 = int32(0)
	if base.B2i32(v1228 == v1229)|base.B2i32(v1228 == v1021) == v1229 {
		goto L201
	} else {
		goto L202
	}
L161:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+24))
	if v1059 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if v1035&int32(1) == int32(0) {
		goto L153
	} else {
		goto L200
	}
L163:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+20))
	v1065 = F_ginPostingListDecode(m, v1062, v1030+int32(28))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	v1068 = v1059
	goto L165
L165:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+20))
	if v1069 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1030)+24)) = v1065
	v1068 = v1065
	goto L165
L167:
	;
	v1070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1069)+6)))
	v1078 = (v1070+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	goto L169
L168:
	;
	v1078 = int32(_a_F_ginbulkdelete_16)
	goto L169
L169:
	;
	v1080 = v1001 + int32(12)
	v1081 = int32(0)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+28))
	if v1084 <= v1081 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+24))
	F_pfree(m, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L188
	}
L171:
	;
	v1087 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1087
	v1192 = v1087
	goto L170
L172:
	;
	goto L173
L173:
	;
	v1100 = v1081
	v1101 = v1081
	v1111 = v1081
	goto L174
L174:
	;
	v1124 = v1111 * int32(6)
	v1125 = v1068 + v1124
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v997)+8))
	v1128 = m.T0[v1127].(func(*base.Module, int32, int32) int32)(m, v1125, v1126)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1155
	v1192 = v1154
	goto L170
L176:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v997)+4))
	if v1128 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1158 = v1111 + int32(1)
	if v1158 != v1084 {
		v1100 = v1154
		v1101 = v1155
		v1111 = v1158
		goto L174
	} else {
		goto L187
	}
L178:
	;
	v1131 = *(*float64)(unsafe.Add(mBase, uint32(v1130)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1130)+16)) = base.F64_add(v1131, float64(1))
	if v1100 != 0 {
		v1154 = v1100
		v1155 = v1101
		goto L177
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v1140 = *(*float64)(unsafe.Add(mBase, uint32(v1130)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1130)+8)) = base.F64_add(v1140, float64(1))
	if v1100 != 0 {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v1135 = F_palloc(m, v1084*int32(6))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	if v1124 == int32(0) {
		v1154 = v1135
		v1155 = v1101
		goto L177
	} else {
		goto L183
	}
L183:
	;
	base.MemoryCopy(m, v1135, v1068, v1124)
	v1154 = v1135
	v1155 = v1101
	goto L177
L184:
	;
	v1146 = v1100 + v1101*int32(6)
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1125)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1146)+4)) = uint16(v1147)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1125)))
	*(*int32)(unsafe.Add(mBase, uint32(v1146))) = v1149
	goto L186
L185:
	;
	goto L186
L186:
	;
	v1154 = v1100
	v1155 = v1101 + int32(1)
	goto L177
L187:
	;
	goto L175
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1030)+24)) = int64(0)
	if v1192 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	if int32(0) < v1198 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+4))
	if v1220 != v1021 {
		v1030 = v1220
		goto L161
	} else {
		goto L199
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1030)+28)) = v1215
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+4))
	if v1218 != v1021 {
		v1030 = v1218
		v1035 = int32(1)
		goto L161
	} else {
		goto L198
	}
L193:
	;
	v1203 = F_ginCompressPostingList(m, v1192, v1198, v1078, v1001+int32(8))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v1211 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1030)+8)) = uint8(v1211)
	*(*int32)(unsafe.Add(mBase, uint32(v1030)+20)) = int32(0)
	v1215 = v1198
	goto L192
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1030)+20)) = v1203
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+8))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	if v1206 != v1207 {
		goto L152
	} else {
		goto L197
	}
L197:
	;
	v1209 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v1030)+8)) = uint8(v1209)
	v1215 = v1206
	goto L192
L198:
	;
	goto L160
L199:
	;
	goto L162
L200:
	;
	goto L160
L201:
	;
	v1238 = v1228
	v1242 = int32(0)
	goto L204
L202:
	;
	goto L203
L203:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v995)+48))
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328)+118)))
	if v1329 != int32(112) {
		goto L214
	} else {
		goto L215
	}
L204:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1238)+8)))
	v1268 = int32(0)
	v1270 = v1242 | base.B2i32(v1267 != v1268)
	v1271 = int32(1)
	if base.B2i32(v1270&v1271 == v1268)|base.B2i32(v1267 == v1271) == v1268 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L203
L206:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+20))
	v1281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1280)+6)))
	v1287 = (v1281+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	v1288 = F_palloc(m, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+4))
	if v1295 != v1021 {
		v1238 = v1295
		v1242 = v1270
		goto L204
	} else {
		goto L213
	}
L209:
	;
	if v1287 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+20))
	base.MemoryCopy(m, v1288, v1290, v1287)
	goto L212
L211:
	;
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1238)+20)) = v1288
	goto L208
L213:
	;
	goto L205
L214:
	;
	v1340 = int32(_a_F_ginbulkdelete_11)
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v1342 + int32(1)
	if v960 < int32(0) {
		goto L223
	} else {
		goto L224
	}
L215:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v1333 <= int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v995)+32))
	if v1336 != 0 {
		goto L214
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	F_computeLeafRecompressWALData(m, v1021)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v995)+40))
	if v1337 != 0 {
		goto L214
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	goto L214
L222:
	;
	v1364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1363)+16)))
	v1365 = v1364 + v1363
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1365)+6)))
	v1368 = v1366 & int32(128)
	if v1368 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1349+(v960^int32(-1))<<(uint(int32(2))%32))))
	v1363 = v1355
	goto L222
L224:
	;
	goto L225
L225:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v1363 = v1357 + v960<<(uint(int32(13))%32) + int32(-8192)
	goto L222
L226:
	;
	v1372 = v1366 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v1365)+6)) = uint16(v1372)
	v1374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1363)+16)))
	v1376 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1363+v1374)+4)) = uint16(v1376)
	goto L228
L227:
	;
	goto L228
L228:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	if base.B2i32(v1379 == int32(0))|base.B2i32(v1379 == v1021) != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1486 = int32(32)
	goto L231
L230:
	;
	v1385 = int32(0)
	v1392 = v1379
	v1396 = base.B2i32(v1368 == v1385)
	v1397 = v1385
	v1398 = v1363 + int32(32)
	goto L232
L231:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1363)+12)) = uint16(v1486)
	F_MarkBufferDirty(m, v960)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L241
	}
L232:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392)+8)))
	v1424 = base.B2i32(v1421 != int32(0)) | v1396
	if v1421 != int32(1) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v1486 = v1447 + int32(32)
	goto L231
L234:
	;
	v1427 = int32(1)
	v1429 = int32(0)
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+20))
	v1432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1431)+6)))
	v1438 = (v1432+v1427)&int32(_a_F_ginbulkdelete_8) + int32(8)
	if base.B2i32(v1424&v1427 == v1429)|base.B2i32(v1438 == v1429) == v1429 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	v1447 = v1397
	v1448 = v1398
	goto L236
L236:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	if v1451 != v1021 {
		v1392 = v1451
		v1396 = v1424
		v1397 = v1447
		v1398 = v1448
		goto L232
	} else {
		goto L240
	}
L237:
	;
	base.MemoryCopy(m, v1398, v1431, v1438)
	goto L239
L238:
	;
	goto L239
L239:
	;
	v1447 = v1397 + v1438
	v1448 = v1398 + v1438
	goto L236
L240:
	;
	goto L233
L241:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v995)+48))
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1490)+118)))
	if v1491 != int32(112) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1518 = int32(_a_F_ginbulkdelete_11)
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v1520 - int32(1)
	goto L153
L243:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v1495 <= int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v995)+32))
	if v1498 != 0 {
		goto L242
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v995)+40))
	if v1499 != 0 {
		goto L242
	} else {
		goto L248
	}
L248:
	;
	goto L246
L249:
	;
	F_XLogRegisterBuffer(m, int32(0), v960, int32(8))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+24))
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+28))
	F_XLogRegisterBufData(m, int32(0), v1507, v1508)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1513 = F_XLogInsert(m, int32(13), int32(144))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1020))) = base.I64_rotr(v1513, int64(32))
	goto L242
L253:
	;
	F_errmsg_internal(m, int32(_a_F_ginbulkdelete_17), int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_ginbulkdelete_18), int32(782), int32(_a_F_ginbulkdelete_19))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	v1576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962)+16)))
	v1577 = v962 + v1576
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577)+6)))
	if v1578&int32(128) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2752))
	v1605 = int32(0)
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[2])))
	v1608 = F_ReadBufferExtended(m, v1604, v1605, v1602, v1605, v1607)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L1
	} else {
		goto L270
	}
L258:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	F_UnlockReleaseBuffer(m, v960)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L268
	}
L259:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	F_UnlockReleaseBuffer(m, v960)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L1
	} else {
		goto L265
	}
L260:
	;
	v1581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962)+12)))
	if v1581 != int32(32) {
		goto L259
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1577)+4)))
	if v1584 == int32(0) {
		goto L258
	} else {
		goto L264
	}
L263:
	;
	goto L258
L264:
	;
	goto L259
L265:
	;
	if v1587 != int32(-1) {
		v1602 = v1587
		v1603 = v971
		goto L257
	} else {
		goto L266
	}
L266:
	;
	if v971&int32(1) == int32(0) {
		goto L147
	} else {
		goto L267
	}
L267:
	;
	goto L150
L268:
	;
	if v1596 == int32(-1) {
		goto L150
	} else {
		goto L269
	}
L269:
	;
	v1602 = v1596
	v1603 = int32(1)
	goto L257
L270:
	;
	F_LockBuffer(m, v1608, int32(2))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	if v1608 < int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1616+(v1608^int32(-1))<<(uint(int32(2))%32))))
	v1630 = v1622
	goto L274
L273:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v1630 = v1624 + v1608<<(uint(int32(13))%32) + int32(-8192)
	goto L274
L274:
	;
	v960 = v1608
	v962 = v1630
	v971 = v1603
	goto L148
L275:
	;
	F_LockBufferForCleanup(m, v1637)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v1641 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[10]))) = v1641
	v1643 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[11]))) = v1643
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6]))) = v1643
	v1647 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[10]))) = uint8(v1647)
	v1655 = F_ginScanToDelete(m, v34+int32(2752), v879, v1647, v34+int32(_a_F_ginbulkdelete_7), v1641)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_ginbulkdelete[6])))
	if v1657 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1659 = v1657
	goto L281
L279:
	;
	goto L280
L280:
	;
	F_UnlockReleaseBuffer(m, v1637)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L1
	} else {
		goto L285
	}
L281:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1659)))
	F_pfree(m, v1659)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L283
	}
L282:
	;
	goto L280
L283:
	;
	if v1689 != 0 {
		v1659 = v1689
		goto L281
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	goto L147
L286:
	;
	v1760 = v860 + int32(1)
	if v1760 != v823 {
		v860 = v1760
		goto L131
	} else {
		goto L287
	}
L287:
	;
	goto L132
L288:
	;
	v1795 = int32(0)
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1798 = F_ReadBufferExtended(m, v36, v1795, v824, v1795, v1797)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	goto L40
L291:
	;
	F_LockBuffer(m, v1798, int32(2))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v312 = v1798
	goto L39
L293:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2756))
	m.G0 = v34 + int32(_a_F_ginbulkdelete_0)
	return v1806
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
			return v21 & int32(_a_F_gintuple_get_attrnum_0)
		}
	} else {
		v21 = int32(1)
		m.G0 = v7 + int32(16)
		return v21 & int32(_a_F_gintuple_get_attrnum_0)
	}
}
func F_ginvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int64
	_ = v175
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(_a_F_ginvacuumcleanup_0)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v22 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(_a_F_ginvacuumcleanup_0)
	return v244
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[0]))
	if v26 != int32(4) {
		v244 = l1
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
	v30 = v19 + int32(36)
	F_initGinState(m, v30, v21)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v36 = int32(1)
	F_ginInsertCleanup(m, v30, int32(0), v36, v36, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v244 = l1
	goto L1
L9:
	;
	v43 = F_palloc0(m, int32(40))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v57 = l1
	goto L11
L11:
	;
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v59
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v68 = float64(0)
	if base.F64_gt(v67, v68) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v46 = v19 + int32(36)
	F_initGinState(m, v46, v21)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[0]))
	F_ginInsertCleanup(m, v46, base.B2i32(v50 != int32(4)), int32(0), int32(1), v43)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v57 = v43
	goto L11
L15:
	;
	v71 = v67
	goto L17
L16:
	;
	v71 = v68
	goto L17
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v57)+8)) = v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v75 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v94) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	F_LockRelationForExtension(m, v21, int32(7))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L25
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v78 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v82 = F_RelationGetNumberOfBlocksInFork(m, v21, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v94 = v82
	v95 = v3
	goto L18
L25:
	;
	v88 = F_RelationGetNumberOfBlocksInFork(m, v21, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_UnlockRelationForExtension(m, v21, int32(7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v94 = v88
	v95 = int32(1)
	goto L18
L28:
	;
	v106 = int32(1)
	v108 = v3
	v110 = v3
	v111 = v3
	v114 = int64(0)
	goto L31
L29:
	;
	v213 = v3
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v94
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ginUpdateStats(m, v221, v19, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L60
	}
L31:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v213 = v194
	goto L30
L33:
	;
	v118 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v121 = F_ReadBufferExtended(m, v21, v118, v106, v118, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_LockBuffer(m, v121, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if v121 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	F_UnlockReleaseBuffer(m, v121)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L58
	}
L37:
	;
	F_RecordFreeIndexPage(m, v21, v106)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L57
	}
L38:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+14)))
	if v144 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[1]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v121^int32(-1))<<(uint(int32(2))%32))))
	v143 = v135
	goto L38
L40:
	;
	goto L41
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[2]))
	v143 = v137 + v121<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+v147)+6)))
	if v149&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v152 == int32(0) {
		goto L37
	} else {
		goto L46
	}
L44:
	;
	v160 = v149
	goto L45
L45:
	;
	if v160&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v155 = F_GlobalVisCheckRemovableXid(m, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if v155 != 0 {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+v157)+6)))
	v160 = v159
	goto L45
L49:
	;
	v164 = v110 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v164
	v194 = v108
	v195 = v164
	v196 = v111
	v198 = v114
	goto L36
L50:
	;
	goto L51
L51:
	;
	if v160&int32(16) != 0 {
		v194 = v108
		v195 = v110
		v196 = v111
		v198 = v114
		goto L36
	} else {
		goto L52
	}
L52:
	;
	v169 = v111 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v169
	if v160&int32(2) == int32(0) {
		v194 = v108
		v195 = v110
		v196 = v169
		v198 = v114
		goto L36
	} else {
		goto L53
	}
L53:
	;
	v175 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v143)+12)))
	if base.Ui64(int64(25)) <= base.Ui64(v175) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v185 = int64(base.Ui64(v175+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
	goto L56
L55:
	;
	v185 = int64(0)
	goto L56
L56:
	;
	v186 = v114 + v185
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v186
	v194 = v108
	v195 = v110
	v196 = v169
	v198 = v186
	goto L36
L57:
	;
	v194 = v108 + int32(1)
	v195 = v110
	v196 = v111
	v198 = v114
	goto L36
L58:
	;
	v202 = v106 + int32(1)
	if v202 != v94 {
		v106 = v202
		v108 = v194
		v110 = v195
		v111 = v196
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
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_FreeSpaceMapVacuum(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v213
	if v95 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_LockRelationForExtension(m, v21, int32(7))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v240 = F_RelationGetNumberOfBlocksInFork(m, v21, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L68
	}
L65:
	;
	v233 = F_RelationGetNumberOfBlocksInFork(m, v21, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v233
	F_UnlockRelationForExtension(m, v21, int32(7))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v244 = v57
	goto L1
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v240
	v244 = v57
	goto L1
}
func F_gistadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v13 int32
	_ = v13
	Fn13908(m, l0, l1, l2, l3, int32(_a_F_gistadjustmembers_0), int32(349), int32(_a_F_gistadjustmembers_1), int32(_a_F_gistadjustmembers_2), int32(230), int32(_a_F_gistadjustmembers_3), int32(12))
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_gistdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v658 int32
	_ = v658
	v6 = l5
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+72)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v7
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+74)) = uint16(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+40)) = uint8(v6)
	v36 = v17 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v36
	v42 = v36
	v44 = v7
	v45 = v7
	goto L1
L1:
	;
	if v45&int32(1) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v98)+16))
	if v107 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v97 = v44
	v98 = v42
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v44&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	F_LockBuffer(m, v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	F_ReleaseBuffer(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v67
	v69 = int32(0)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+24)))
	if v70 != int32(1) {
		v97 = v69
		v98 = v67
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v78 = v67
	goto L14
L14:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	F_ReleaseBuffer(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v90
	v97 = v69
	v98 = v90
	goto L3
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)))
	if v91 != 0 {
		v78 = v90
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v111 = F_ReadBuffer(m, l0, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v115 = v97 & int32(1)
	if v115 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v111
	goto L20
L22:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v118, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v125 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_gistcheckpage(m, l0, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v143
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v125^int32(-1))<<(uint(int32(2))%32))))
	v143 = v135
	goto L27
L29:
	;
	goto L30
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v143 = v137 + v125<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v645
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+24)))
	v42 = v645
	v44 = v647
	v45 = v658
	goto L1
L32:
	;
	v435 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L10
	} else {
		goto L108
	}
L33:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)))
	v42 = v98
	v44 = int32(1)
	v45 = v429
	goto L1
L34:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v186 != 0 {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v145 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v143)+4)))
	v146 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v143))))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v145 | v146<<(uint(int64(32))%64)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
	v152 = v143 + v151
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+12)))
	if v153&int32(8) == int32(0) {
		v183 = v143
		v184 = v153
		v185 = v152
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v158 = F_BufferGetLSNAtomic(m, v125)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	goto L32
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+16)))
	v163 = v161 + v162
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+12)))
	if v164&int32(8) == int32(0) {
		v183 = v161
		v184 = v164
		v185 = v163
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v169, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v173, int32(2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+16)))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v178)+12)))
	if v180&int32(8) != 0 {
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
	v402 = m.ExcPending
	if v402 != 0 {
		goto L10
	} else {
		goto L103
	}
L45:
	;
	v372 = F_gistinserttuple(m, v17+int32(28), v98, l3, l1, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L97
	}
L46:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v362
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+24)))
	v42 = v362
	v44 = int32(0)
	v45 = v365
	goto L1
L47:
	;
	if v184&int32(1) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_UnlockReleaseBuffer(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L55
	}
L49:
	;
	if v184&int32(2) != 0 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v184&int32(2) == int32(0) {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v189)+16))
	v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v185)+4)))
	v192 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui64(v190) < base.Ui64(v191|v192<<(uint(int64(32))%64)) {
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
	v208 = F_gistchoose(m, l0, v183, l1, l3)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v115 != 0 {
		goto L45
	} else {
		goto L80
	}
L59:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210+v208<<(uint(int32(2))%32))+20))
	v217 = v210 + v214&int32(_a_F_gistdoinsert_0)
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v218 == int32(_a_F_gistdoinsert_1) {
		goto L44
	} else {
		goto L60
	}
L60:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
	v223 = F_gistgetadjusted(m, l0, v217, l1, l3)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L62
	}
L61:
	;
	v278 = int32(0)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v279, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L78
	}
L62:
	;
	if v223 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if v115 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v229, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v267 = F_gistinserttuple(m, v17+int32(28), v98, l3, v223, v208)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L74
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v233, int32(2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v237 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v98)+16))
	v258 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v255)+4)))
	v259 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v255))))
	if v257 != v258|v259<<(uint(int64(32))%64) {
		goto L33
	} else {
		goto L73
	}
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241+(v237^int32(-1))<<(uint(int32(2))%32))))
	v255 = v247
	goto L69
L71:
	;
	goto L72
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v255 = v249 + v237<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	goto L66
L74:
	;
	if v267 == int32(0) {
		goto L61
	} else {
		goto L75
	}
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v271 == int32(0) {
		goto L33
	} else {
		goto L76
	}
L76:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_UnlockReleaseBuffer(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	goto L46
L78:
	;
	v284 = F_palloc0(m, int32(32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+28)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v221<<(uint(int32(16))%32) | v222
	*(*uint16)(unsafe.Add(mBase, uint32(v284)+26)) = uint16(v208)
	v645 = v284
	v647 = v278
	goto L31
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v292, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v296, int32(2))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v300 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v318
	v320 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v318)+4)))
	v321 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v318))))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v320 | v321<<(uint(int64(32))%64)
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+16)))
	v327 = v318 + v326
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+12)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v329 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v304+(v300^int32(-1))<<(uint(int32(2))%32))))
	v318 = v310
	goto L83
L85:
	;
	goto L86
L86:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v318 = v312 + v300<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	if v328&int32(1) != 0 {
		goto L45
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v328&int32(8)|v328&int32(2) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v334 = int32(0)
	F_LockBuffer(m, v300, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)))
	v42 = v98
	v44 = v334
	v45 = v338
	goto L1
L92:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v346)+16))
	v348 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v327)+4)))
	v349 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v327))))
	if base.Ui64(v348|v349<<(uint(int64(32))%64)) <= base.Ui64(v347) {
		goto L45
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_UnlockReleaseBuffer(m, v300)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	goto L46
L97:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_LockBuffer(m, v374, int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v383 = v98
	goto L99
L99:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	F_ReleaseBuffer(m, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L101
	}
L100:
	;
	m.G0 = v17 + int32(80)
	return
L101:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v383)+28))
	if v395 != 0 {
		v383 = v395
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v403 + int32(4)
	F_errmsg(m, int32(_a_F_gistdoinsert_2), v17)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	F_errdetail(m, int32(_a_F_gistdoinsert_3), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errhint(m, int32(_a_F_gistdoinsert_4), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_gistdoinsert_5), int32(768), int32(_a_F_gistdoinsert_6))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
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
	if v435 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v437 + int32(4)
	F_errmsg(m, int32(_a_F_gistdoinsert_7), v17+int32(16))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L10
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v463 = v454
	v466 = int32(0)
	goto L114
L112:
	;
	F_errfinish(m, int32(_a_F_gistdoinsert_5), int32(1209), int32(_a_F_gistdoinsert_8))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v471 = F_palloc(m, int32(8))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L10
	} else {
		goto L116
	}
L115:
	;
	v633 = int32(0)
	F_gistfinishsplit(m, v17+int32(28), v98, l3, v620, v633)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L10
	} else {
		goto L152
	}
L116:
	;
	if v463 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if v463 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L118:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	F_LockBuffer(m, v556, int32(2))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L10
	} else {
		goto L138
	}
L119:
	;
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v490)+12)))
	if base.Ui32(v491) < base.Ui32(int32(25)) {
		goto L118
	} else {
		goto L123
	}
L120:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v476+(v463^int32(-1))<<(uint(int32(2))%32))))
	v490 = v482
	goto L119
L121:
	;
	goto L122
L122:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v490 = v484 + v463<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	v495 = v491 + int32(_a_F_gistdoinsert_9)
	if v495&int32(_a_F_gistdoinsert_10) == int32(0) {
		goto L118
	} else {
		goto L124
	}
L124:
	;
	v510 = int32(1)
	v512 = int32(0)
	goto L125
L125:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v490+int32(20)+v510<<(uint(int32(2))%32))))
	v528 = v490 + v525&int32(_a_F_gistdoinsert_0)
	if v512 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	if v537 != 0 {
		v583 = v537
		goto L117
	} else {
		goto L137
	}
L127:
	;
	if v510 != int32(base.Ui32(v495)>>(uint(int32(2))%32))&int32(_a_F_gistdoinsert_11) {
		v510 = v510 + int32(1)
		v512 = v537
		goto L125
	} else {
		goto L136
	}
L128:
	;
	v531 = F_CopyIndexTuple(m, v528)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v533 = F_gistgetadjusted(m, l0, v512, v528, l3)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L10
	} else {
		goto L132
	}
L131:
	;
	v537 = v531
	goto L127
L132:
	;
	if v533 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v535 = v533
	goto L135
L134:
	;
	v535 = v512
	goto L135
L135:
	;
	v537 = v535
	goto L127
L136:
	;
	goto L126
L137:
	;
	goto L118
L138:
	;
	F_gistFindCorrectParent(m, l0, v98)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+8))
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+26)))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v563+v564<<(uint(int32(2))%32))+20))
	v572 = F_CopyIndexTuple(m, v563+v568&int32(_a_F_gistdoinsert_0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	F_LockBuffer(m, v575, int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v583 = v572
	goto L117
L142:
	;
	v612 = int32(_a_F_gistdoinsert_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v583)+4)) = uint16(v612)
	*(*uint16)(unsafe.Add(mBase, uint32(v583)+2)) = uint16(v611)
	v616 = int32(base.Ui32(v611) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v583))) = uint16(v616)
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v463
	v620 = F_lappend(m, v466, v471)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L10
	} else {
		goto L146
	}
L143:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[2]))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v596+(v463^int32(-1))<<(uint(int32(6))%32))+16))
	v611 = v602
	goto L142
L144:
	;
	goto L145
L145:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[3]))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v604+v463<<(uint(int32(6))%32)+int32(-64))+16))
	v611 = v610
	goto L142
L146:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v490)+16)))
	v623 = v490 + v622
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+12)))
	if v624&int32(8) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	v628 = F_ReadBuffer(m, l0, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
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
	F_LockBuffer(m, v628, int32(2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	v463 = v628
	v466 = v620
	goto L114
L152:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	F_UnlockReleaseBuffer(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v645 = v642
	v647 = v633
	goto L31
}
func F_gistextractpage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v9) {
		v17 = int32(base.Ui32(v9+int32(_a_F_gistextractpage_0)) >> (uint(int32(2)) % 32))
	} else {
		v17 = int32(0)
	}
	v19 = v17 & int32(_a_F_gistextractpage_1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v23 = F_palloc(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
		} else {
			v29 = int32(1)
			v31 = l0 + int32(20)
			v35 = (v17 + v29) & int32(_a_F_gistextractpage_1)
			if base.Ui32(int32(3)) <= base.Ui32(v35) {
				v38 = int32(2)
				if base.Ui32(v35) <= base.Ui32(v38) {
					v41 = v38
				} else {
					v41 = v35
				}
				v42 = int32(1)
				v43 = v41 - v42
				v51 = v42
				v52 = int32(0)
				for {
					v58 = int32(2)
					v59 = v51 << (uint(v58) % 32)
					v61 = int32(4)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v59+v31)))
					v65 = int32(_a_F_gistextractpage_2)
					*(*int32)(unsafe.Add(mBase, uint32(v23+v59-v61))) = l0 + v64&v65
					v70 = v59 + v61
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v70+v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v23+v70-v61))) = l0 + v75&v65
					v81 = v51 + v58
					v83 = v52 + v58
					if v83 != v43&int32(-2) {
						v51 = v81
						v52 = v83
						continue
					} else {
						break
					}
					break
				}
				if v43&v42 == int32(0) {
				} else {
					v88 = v81
					v96 = v88 << (uint(int32(2)) % 32)
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v96+v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v23+v96-int32(4)))) = l0 + v101&int32(_a_F_gistextractpage_2)
				}
			} else {
				v88 = v29
				v96 = v88 << (uint(int32(2)) % 32)
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v96+v31)))
				*(*int32)(unsafe.Add(mBase, uint32(v23+v96-int32(4)))) = l0 + v101&int32(_a_F_gistextractpage_2)
			}
		}
		return v23
	}
}
func F_gistunionsubkey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	v4 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+352))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v26 = F_palloc(m, v23<<(uint(int32(2))%32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		if v23 <= int32(0) {
			v132 = v4
		} else {
			if v23 != int32(1) {
				v40 = v4
				v45 = v4
				v51 = v4
				for {
					v54 = v21 + v45<<(uint(int32(1))%32)
					v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
					if v22 != 0 {
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v22))))
						if v57 != 0 {
							v70 = v40
						} else {
							v58 = int32(2)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(v58)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v26+v40<<(uint(v58)%32)))) = v66
							v70 = v40 + int32(1)
						}
					} else {
						v58 = int32(2)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(v58)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v26+v40<<(uint(v58)%32)))) = v66
						v70 = v40 + int32(1)
					}
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
					if v22 != 0 {
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v22))))
						if v73 != 0 {
							v86 = v70
						} else {
							v74 = int32(2)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l1+v71<<(uint(v74)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v26+v70<<(uint(v74)%32)))) = v82
							v86 = v70 + int32(1)
						}
					} else {
						v74 = int32(2)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1+v71<<(uint(v74)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v26+v70<<(uint(v74)%32)))) = v82
						v86 = v70 + int32(1)
					}
					v87 = int32(2)
					v88 = v45 + v87
					v90 = v51 + v87
					if v90 != v23&int32(2147483646) {
						v40 = v86
						v45 = v88
						v51 = v90
						continue
					} else {
						break
					}
					break
				}
				if v23&int32(1) == int32(0) {
					v132 = v86
				} else {
					v98 = v86
					v103 = v88
					v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v103<<(uint(int32(1))%32)))))
					if v22 != 0 {
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v22))))
						if v115 != 0 {
							v132 = v98
						} else {
							v116 = int32(2)
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(v116)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v26+v98<<(uint(v116)%32)))) = v124
							v132 = v98 + int32(1)
						}
					} else {
						v116 = int32(2)
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(v116)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v26+v98<<(uint(v116)%32)))) = v124
						v132 = v98 + int32(1)
					}
				}
			} else {
				v98 = v4
				v103 = v4
				v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v103<<(uint(int32(1))%32)))))
				if v22 != 0 {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v22))))
					if v115 != 0 {
						v132 = v98
					} else {
						v116 = int32(2)
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(v116)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v26+v98<<(uint(v116)%32)))) = v124
						v132 = v98 + int32(1)
					}
				} else {
					v116 = int32(2)
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(v116)%32)-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v26+v98<<(uint(v116)%32)))) = v124
					v132 = v98 + int32(1)
				}
			}
		}
		F_gistMakeUnionItVec(m, l0, v26, v132, l2+int32(32), l2+int32(160))
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			F_pfree(m, v26)
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return
			} else {
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v152 = F_palloc(m, v149<<(uint(int32(2))%32))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return
				} else {
					if v149 <= int32(0) {
						v263 = v4
					} else {
						v156 = int32(0)
						if v149 != int32(1) {
							v167 = int32(0)
							v168 = v156
							v171 = v4
							for {
								v182 = v148 + v168<<(uint(int32(1))%32)
								v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
								if v22 != 0 {
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v183))))
									if v185 != 0 {
										v198 = v171
									} else {
										v186 = int32(2)
										v194 = *(*int32)(unsafe.Add(mBase, uint32(l1+v183<<(uint(v186)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v152+v171<<(uint(v186)%32)))) = v194
										v198 = v171 + int32(1)
									}
								} else {
									v186 = int32(2)
									v194 = *(*int32)(unsafe.Add(mBase, uint32(l1+v183<<(uint(v186)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v152+v171<<(uint(v186)%32)))) = v194
									v198 = v171 + int32(1)
								}
								v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+2)))
								if v22 != 0 {
									v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v199))))
									if v201 != 0 {
										v214 = v198
									} else {
										v202 = int32(2)
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l1+v199<<(uint(v202)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v152+v198<<(uint(v202)%32)))) = v210
										v214 = v198 + int32(1)
									}
								} else {
									v202 = int32(2)
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l1+v199<<(uint(v202)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v152+v198<<(uint(v202)%32)))) = v210
									v214 = v198 + int32(1)
								}
								v215 = int32(2)
								v216 = v168 + v215
								v218 = v167 + v215
								if v218 != v149&int32(2147483646) {
									v167 = v218
									v168 = v216
									v171 = v214
									continue
								} else {
									break
								}
								break
							}
							if v149&int32(1) == int32(0) {
								v263 = v214
							} else {
								v226 = v216
								v229 = v214
								v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+v226<<(uint(int32(1))%32)))))
								if v22 != 0 {
									v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v22))))
									if v243 != 0 {
										v263 = v229
									} else {
										v244 = int32(2)
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l1+v241<<(uint(v244)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v152+v229<<(uint(v244)%32)))) = v252
										v263 = v229 + int32(1)
									}
								} else {
									v244 = int32(2)
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l1+v241<<(uint(v244)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v152+v229<<(uint(v244)%32)))) = v252
									v263 = v229 + int32(1)
								}
							}
						} else {
							v226 = v156
							v229 = v4
							v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+v226<<(uint(int32(1))%32)))))
							if v22 != 0 {
								v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v22))))
								if v243 != 0 {
									v263 = v229
								} else {
									v244 = int32(2)
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l1+v241<<(uint(v244)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v152+v229<<(uint(v244)%32)))) = v252
									v263 = v229 + int32(1)
								}
							} else {
								v244 = int32(2)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l1+v241<<(uint(v244)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v229<<(uint(v244)%32)))) = v252
								v263 = v229 + int32(1)
							}
						}
					}
					F_gistMakeUnionItVec(m, l0, v152, v263, l2+int32(192), l2+int32(320))
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return
					} else {
						F_pfree(m, v152)
						mBase = m.M
						v279 = m.ExcPending
						if v279 != 0 {
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_same_0), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(base.B2i32(v8 != int32(0)))
		return v3
	}
}
