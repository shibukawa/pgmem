package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
				F_errfinish(m, int32(_a_F_GetAccessStrategy_2), int32(496), int32(_a_F_GetAccessStrategy_3))
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
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
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
	return v307
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
	v307 = v6
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
	v300 = m.ExcPending
	if v300 != 0 {
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
	v150 = int32(32)
	v153 = F_palloc_mul(m, int32(4), v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
	v294 = v6
	goto L32
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v157 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L24
	} else {
		goto L49
	}
L47:
	;
	F_FreeDir(m, v123)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L24
	} else {
		goto L80
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 - int32(-64)
	v269 = F_psprintf(m, int32(_a_F_GetConfFilesInDir_7), v14+int32(48))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L24
	} else {
		goto L78
	}
L49:
	;
	if v157 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v159 = v157
	v165 = v153
	v168 = v150
	goto L53
L51:
	;
	v251 = v153
	goto L52
L52:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v256 <= int32(0) {
		v281 = v251
		goto L47
	} else {
		goto L76
	}
L53:
	;
	v171 = v159 + int32(19)
	v172 = F_strlen(m, v171)
	mBase = m.M
	if base.Ui32(v172) < base.Ui32(int32(6)) {
		v240 = v165
		v242 = v168
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v251 = v240
	goto L52
L55:
	;
	v243 = F_ReadDir(m, v123, v121)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L74
	}
L56:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v175 == int32(46) {
		v240 = v165
		v242 = v168
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v180 = v172 + v171 - int32(5)
	v181 = int32(_a_F_GetConfFilesInDir_8)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetConfFilesInDir[2])))
	if base.B2i32(v184 == int32(0))|base.B2i32(v184 != v187) != 0 {
		v205 = v184
		v206 = v187
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v205-v206 != 0 {
		v240 = v165
		v242 = v168
		goto L55
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	v190 = v180
	v191 = v181
	goto L61
L61:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	if v195 == int32(0) {
		v205 = v195
		v206 = v194
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v205 = v195
	v206 = v194
	goto L59
L63:
	;
	v198 = int32(1)
	if v195 == v194 {
		v190 = v190 + v198
		v191 = v191 + v198
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v209 = v14 - int32(-64)
	F_join_path_components(m, v209, v121, v171)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L24
	} else {
		goto L66
	}
L66:
	;
	F_canonicalize_path_enc(m, v209)
	mBase = m.M
	v214 = F_get_dirent_type(m, v209, v159, int32(1), l2)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L24
	} else {
		goto L68
	}
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v168 <= v216 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	switch v214 {
	case 0:
		goto L48
	default:
		goto L67
	case 3:
		v240 = v165
		v242 = v168
		goto L55
	}
L69:
	;
	v219 = v168 + int32(32)
	v222 = F_repalloc(m, v165, v219<<(uint(int32(2))%32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L24
	} else {
		goto L72
	}
L70:
	;
	v224 = v165
	v225 = v168
	goto L71
L71:
	;
	v228 = F_pstrdup(m, v14-int32(-64))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L24
	} else {
		goto L73
	}
L72:
	;
	v224 = v222
	v225 = v219
	goto L71
L73:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v224+v230<<(uint(int32(2))%32)))) = v228
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v235 + int32(1)
	v240 = v224
	v242 = v225
	goto L55
L74:
	;
	if v243 != 0 {
		v159 = v243
		v165 = v240
		v168 = v242
		goto L53
	} else {
		goto L75
	}
L75:
	;
	goto L54
L76:
	;
	F_pg_qsort(m, v251, v256, int32(4), int32(1285))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v281 = v251
	goto L47
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v269
	F_pfree(m, v165)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v281 = int32(0)
	goto L47
L80:
	;
	v294 = v281
	goto L32
L81:
	;
	v307 = v294
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
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
			v18 = base.I32_wrap_i64(v14)
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_GetFdwRoutine_0), v6)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(388), int32(_a_F_GetFdwRoutine_2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v21 != int32(450) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg_internal(m, int32(_a_F_GetFdwRoutine_0), v6)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(388), int32(_a_F_GetFdwRoutine_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
					return v18
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_GetFdwRoutine_3), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutine_1), int32(380), int32(_a_F_GetFdwRoutine_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(440), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4+int32(440), int32(_a_F_GetInsertRecPtr_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+184))
			v18 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16)+440)), uint32(v18))
			return v17
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetInsertRecPtr[0]))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+184))
		v18 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16)+440)), uint32(v18))
		return v17
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
				F_errfinish(m, int32(_a_F_GetLatestSnapshot_2), int32(362), int32(_a_F_GetLatestSnapshot_3))
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
func F_GlobalVisTestXidConsideredRunning(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v15 = v11 + base.I64_extend_i32_s(l1-base.I32_wrap_i64(v11))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v15) < base.Ui64(v16) {
		v34 = int32(0)
		m.G0 = v9 + int32(48)
		return v34
	} else {
		v18 = int32(1)
		if base.Ui64(v11) <= base.Ui64(v15) {
			v34 = v18
			m.G0 = v9 + int32(48)
			return v34
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestXidConsideredRunning[0]))
			if v21 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestXidConsideredRunning[1]))
				if v23 == v21 {
					v34 = v18
					m.G0 = v9 + int32(48)
					return v34
				} else {
					F_ComputeXidHorizons(m, v9+int32(8))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = base.B2i32(base.Ui64(v31) <= base.Ui64(v15))
						m.G0 = v9 + int32(48)
						return v34
					}
				}
			} else {
				F_ComputeXidHorizons(m, v9+int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = base.B2i32(base.Ui64(v31) <= base.Ui64(v15))
					m.G0 = v9 + int32(48)
					return v34
				}
			}
		}
	}
}
func F_g_intbig_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
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
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
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
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v772 int32
	_ = v772
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v804 int32
	_ = v804
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	v2 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = base.I32_wrap_i64(v27)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v30 == v2 {
		v47 = v2
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
		if v34 == int32(0) {
			v47 = v2
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			if v37 != int32(7) {
				v47 = v2
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				if v40 != int32(17) {
					v47 = v2
				} else {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
					v47 = v43 ^ int32(1)
				}
			}
		}
	}
	if v47&int32(1) != 0 {
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v51 = F_get_fn_opclass_options(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int64(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
			v56 = v55
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v61 = (v57 + int32(_a_F_g_intbig_picksplit_0)) & int32(_a_F_g_intbig_picksplit_1)
			v65 = v61<<(uint(int32(1))%32) + int32(4)
			v66 = F_palloc(m, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v66
				v69 = F_palloc(m, v65)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v69
					if base.Ui32(int32(2)) <= base.Ui32(v61) {
						v75 = v26 + int32(8)
						v79 = int32(-1)
						v80 = v2
						v88 = int32(1)
						v89 = v2
						for {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v75+v88*int32(24))))
							v108 = v88 + int32(1)
							v109 = v108
							v110 = v79
							v111 = v80
							v118 = v108
							v120 = v89
							for {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v75+v118*int32(24))))
								v139 = Fn14312(m, v106, v137, v56, int32(4))
								mBase = m.M
								v140 = base.B2i32(v110 < v139)
								if v110 < v139 {
									v141 = v139
								} else {
									v141 = v110
								}
								if v110 < v139 {
									v142 = v109
								} else {
									v142 = v120
								}
								if v110 < v139 {
									v143 = v88
								} else {
									v143 = v111
								}
								v145 = v109 + int32(1)
								v147 = v145 & int32(_a_F_g_intbig_picksplit_1)
								if base.Ui32(v147) <= base.Ui32(v61) {
									v109 = v145
									v110 = v141
									v111 = v143
									v118 = v147
									v120 = v142
									continue
								} else {
									break
								}
								break
							}
							if v61 != v108 {
								v79 = v141
								v80 = v143
								v88 = v108
								v89 = v142
								continue
							} else {
								break
							}
							break
						}
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
						v153 = v143
						v162 = v142
						v164 = v150
					} else {
						v153 = v2
						v162 = v2
						v164 = v69
					}
					v176 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v176
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v176
					v180 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v181 = int32(8)
					v183 = v56 + v181
					v185 = v26 + v181
					v187 = int32(_a_F_g_intbig_picksplit_1)
					v195 = base.B2i32(v153&v187 == v176) | base.B2i32(v162&v187 == v176)
					if v195 != 0 {
						v196 = int32(1)
					} else {
						v196 = v153
					}
					v202 = *(*int32)(unsafe.Add(mBase, uint32(v185+v196&int32(_a_F_g_intbig_picksplit_1)*int32(24))))
					v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
					v205 = v203 & int32(4)
					if v205 != 0 {
						v206 = v181
					} else {
						v206 = v183
					}
					v207 = F_palloc(m, v206)
					mBase = m.M
					v208 = m.ExcPending
					if v208 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v205
						*(*int32)(unsafe.Add(mBase, uint32(v207))) = v206 << (uint(int32(2)) % 32)
						v213 = int32(0)
						if v205|base.B2i32(v56 == v213) == v213 {
							v218 = int32(8)
							base.MemoryCopy(m, v207+v218, v202+v218, v56)
						} else {
						}
						if v195 != 0 {
							v225 = int32(2)
						} else {
							v225 = v162
						}
						v231 = *(*int32)(unsafe.Add(mBase, uint32(v185+v225&int32(_a_F_g_intbig_picksplit_1)*int32(24))))
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
						v234 = v232 & int32(4)
						if v234 != 0 {
							v235 = int32(8)
						} else {
							v235 = v183
						}
						v236 = F_palloc(m, v235)
						mBase = m.M
						v237 = m.ExcPending
						if v237 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v234
							*(*int32)(unsafe.Add(mBase, uint32(v236))) = v235 << (uint(int32(2)) % 32)
							v242 = int32(0)
							if v234|base.B2i32(v56 == v242) == v242 {
								v247 = int32(8)
								base.MemoryCopy(m, v236+v247, v231+v247, v56)
							} else {
							}
							v255 = int32(_a_F_g_intbig_picksplit_1)
							v256 = v57 + v255
							v258 = v256 & v255
							v259 = F_palloc_mul(m, int32(8), v258)
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int64(0)
							} else {
								if v57&int32(_a_F_g_intbig_picksplit_1) == int32(1) {
									F_pg_qsort(m, v259, v258, int32(8), int32(_a_F_g_intbig_picksplit_2))
									mBase = m.M
									v268 = m.ExcPending
									if v268 != 0 {
										return int64(0)
									} else {
										v816 = v180
										v819 = v164
										v831 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v816))) = uint16(v831)
										*(*uint16)(unsafe.Add(mBase, uint32(v819))) = uint16(v831)
										F_pfree(m, v259)
										mBase = m.M
										v836 = m.ExcPending
										if v836 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = base.I64_extend_i32_u(v236)
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = base.I64_extend_i32_u(v207)
											return v27 & int64(4294967295)
										}
									}
								} else {
									v269 = int32(1)
									v271 = v269
									v272 = v269
									for {
										v298 = v259 + v271<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v298-int32(8)))) = uint16(v272)
										v302 = int32(4)
										v307 = *(*int32)(unsafe.Add(mBase, uint32(v185+v271*int32(24))))
										v309 = Fn14312(m, v207, v307, v56, v302)
										mBase = m.M
										v311 = Fn14312(m, v236, v307, v56, int32(4))
										mBase = m.M
										v312 = v309 - v311
										v314 = v312 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v298-v302))) = v312 ^ v314 - v314
										v319 = v272 + int32(1)
										v320 = int32(_a_F_g_intbig_picksplit_1)
										v321 = v319 & v320
										if base.Ui32(v321) <= base.Ui32(v256&v320) {
											v271 = v321
											v272 = v319
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v259, v258, int32(8), int32(_a_F_g_intbig_picksplit_2))
									mBase = m.M
									v328 = m.ExcPending
									if v328 != 0 {
										return int64(0)
									} else {
										v329 = int32(1)
										if base.Ui32(v258) <= base.Ui32(v329) {
											v332 = v329
										} else {
											v332 = v258
										}
										v334 = v56 & int32(2147483644)
										v336 = v56 & int32(3)
										v337 = int32(8)
										v338 = v236 + v337
										v340 = v207 + v337
										v344 = int32(0)
										v354 = v180
										v357 = v164
										for {
											v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259+v344<<(uint(int32(3))%32)))))
											if v196&int32(_a_F_g_intbig_picksplit_1) == v372 {
												*(*uint16)(unsafe.Add(mBase, uint32(v354))) = uint16(v196)
												v375 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v375 + int32(1)
												v788 = v354 + int32(2)
												v791 = v357
											} else {
												if v225&int32(_a_F_g_intbig_picksplit_1) == v372 {
													*(*uint16)(unsafe.Add(mBase, uint32(v357))) = uint16(v225)
													v772 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v772 + int32(1)
													v788 = v354
													v791 = v357 + int32(2)
												} else {
													v388 = *(*int32)(unsafe.Add(mBase, uint32(v185+v372*int32(24))))
													v390 = Fn14312(m, v207, v388, v56, int32(4))
													mBase = m.M
													v393 = Fn14312(m, v236, v388, v56, int32(4))
													mBase = m.M
													v395 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
													v396 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
													v397 = v395 - v396
													if base.F64_lt(base.F64_convert_i32_s(v390), base.F64_add(base.F64_convert_i32_s(v393), base.F64_mul(base.F64_convert_i32_s(v397*v397*v397), float64(-1e-05)))) != 0 {
														v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
														if v405&int32(4) != 0 {
														} else {
															v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+4)))
															if v408&int32(4) != 0 {
																if v56 == int32(0) {
																} else {
																	base.MemoryFill(m, v340, int32(255), v56)
																}
															} else {
																if v56 <= int32(0) {
																} else {
																	v418 = v388 + int32(8)
																	v419 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v56) {
																		v425 = v419
																		v426 = v419
																		for {
																			v449 = v425 + v340
																			v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
																			v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v418))))
																			v453 = v450 | v452
																			*(*uint8)(unsafe.Add(mBase, uint32(v449))) = uint8(v453)
																			v456 = v425 | int32(1)
																			v457 = v340 + v456
																			v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
																			v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456+v418))))
																			v461 = v458 | v460
																			*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v461)
																			v464 = v425 | int32(2)
																			v465 = v340 + v464
																			v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
																			v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v418))))
																			v469 = v466 | v468
																			*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v469)
																			v472 = v425 | int32(3)
																			v473 = v340 + v472
																			v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
																			v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v418))))
																			v477 = v474 | v476
																			*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v477)
																			v479 = int32(4)
																			v480 = v425 + v479
																			v482 = v426 + v479
																			if v482 != v334 {
																				v425 = v480
																				v426 = v482
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v336 == int32(0) {
																		} else {
																			v488 = v480
																			v513 = v488
																			v526 = v419
																			for {
																				v536 = v513 + v340
																				v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
																				v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v418))))
																				v540 = v537 | v539
																				*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v540)
																				v542 = int32(1)
																				v545 = v526 + v542
																				if v545 != v336 {
																					v513 = v513 + v542
																					v526 = v545
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v488 = v419
																		v513 = v488
																		v526 = v419
																		for {
																			v536 = v513 + v340
																			v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
																			v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v418))))
																			v540 = v537 | v539
																			*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v540)
																			v542 = int32(1)
																			v545 = v526 + v542
																			if v545 != v336 {
																				v513 = v513 + v542
																				v526 = v545
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
														*(*uint16)(unsafe.Add(mBase, uint32(v354))) = uint16(v372)
														v573 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v573 + int32(1)
														v788 = v354 + int32(2)
														v791 = v357
													} else {
														v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+4)))
														if v579&int32(4) != 0 {
														} else {
															v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+4)))
															if v582&int32(4) != 0 {
																if v56 == int32(0) {
																} else {
																	base.MemoryFill(m, v338, int32(255), v56)
																}
															} else {
																if v56 <= int32(0) {
																} else {
																	v592 = v388 + int32(8)
																	v593 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v56) {
																		v599 = v593
																		v600 = v593
																		for {
																			v623 = v599 + v338
																			v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
																			v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v592))))
																			v627 = v624 | v626
																			*(*uint8)(unsafe.Add(mBase, uint32(v623))) = uint8(v627)
																			v630 = v599 | int32(1)
																			v631 = v338 + v630
																			v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
																			v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630+v592))))
																			v635 = v632 | v634
																			*(*uint8)(unsafe.Add(mBase, uint32(v631))) = uint8(v635)
																			v638 = v599 | int32(2)
																			v639 = v338 + v638
																			v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
																			v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v592))))
																			v643 = v640 | v642
																			*(*uint8)(unsafe.Add(mBase, uint32(v639))) = uint8(v643)
																			v646 = v599 | int32(3)
																			v647 = v338 + v646
																			v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647))))
																			v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v592))))
																			v651 = v648 | v650
																			*(*uint8)(unsafe.Add(mBase, uint32(v647))) = uint8(v651)
																			v653 = int32(4)
																			v654 = v599 + v653
																			v656 = v600 + v653
																			if v656 != v334 {
																				v599 = v654
																				v600 = v656
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v336 == int32(0) {
																		} else {
																			v662 = v654
																			v687 = v662
																			v700 = v593
																			for {
																				v710 = v687 + v338
																				v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
																				v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v592))))
																				v714 = v711 | v713
																				*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v714)
																				v716 = int32(1)
																				v719 = v700 + v716
																				if v719 != v336 {
																					v687 = v687 + v716
																					v700 = v719
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v662 = v593
																		v687 = v662
																		v700 = v593
																		for {
																			v710 = v687 + v338
																			v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
																			v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v592))))
																			v714 = v711 | v713
																			*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v714)
																			v716 = int32(1)
																			v719 = v700 + v716
																			if v719 != v336 {
																				v687 = v687 + v716
																				v700 = v719
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
														*(*uint16)(unsafe.Add(mBase, uint32(v357))) = uint16(v372)
														v772 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v772 + int32(1)
														v788 = v354
														v791 = v357 + int32(2)
													}
												}
											}
											v804 = v344 + int32(1)
											if v804 != v332 {
												v344 = v804
												v354 = v788
												v357 = v791
												continue
											} else {
												break
											}
											break
										}
										v816 = v788
										v819 = v791
										v831 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v816))) = uint16(v831)
										*(*uint16)(unsafe.Add(mBase, uint32(v819))) = uint16(v831)
										F_pfree(m, v259)
										mBase = m.M
										v836 = m.ExcPending
										if v836 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = base.I64_extend_i32_u(v236)
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = base.I64_extend_i32_u(v207)
											return v27 & int64(4294967295)
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
		v56 = int32(252)
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v61 = (v57 + int32(_a_F_g_intbig_picksplit_0)) & int32(_a_F_g_intbig_picksplit_1)
		v65 = v61<<(uint(int32(1))%32) + int32(4)
		v66 = F_palloc(m, v65)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v66
			v69 = F_palloc(m, v65)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v69
				if base.Ui32(int32(2)) <= base.Ui32(v61) {
					v75 = v26 + int32(8)
					v79 = int32(-1)
					v80 = v2
					v88 = int32(1)
					v89 = v2
					for {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v75+v88*int32(24))))
						v108 = v88 + int32(1)
						v109 = v108
						v110 = v79
						v111 = v80
						v118 = v108
						v120 = v89
						for {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v75+v118*int32(24))))
							v139 = Fn14312(m, v106, v137, v56, int32(4))
							mBase = m.M
							v140 = base.B2i32(v110 < v139)
							if v110 < v139 {
								v141 = v139
							} else {
								v141 = v110
							}
							if v110 < v139 {
								v142 = v109
							} else {
								v142 = v120
							}
							if v110 < v139 {
								v143 = v88
							} else {
								v143 = v111
							}
							v145 = v109 + int32(1)
							v147 = v145 & int32(_a_F_g_intbig_picksplit_1)
							if base.Ui32(v147) <= base.Ui32(v61) {
								v109 = v145
								v110 = v141
								v111 = v143
								v118 = v147
								v120 = v142
								continue
							} else {
								break
							}
							break
						}
						if v61 != v108 {
							v79 = v141
							v80 = v143
							v88 = v108
							v89 = v142
							continue
						} else {
							break
						}
						break
					}
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v153 = v143
					v162 = v142
					v164 = v150
				} else {
					v153 = v2
					v162 = v2
					v164 = v69
				}
				v176 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v176
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v176
				v180 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v181 = int32(8)
				v183 = v56 + v181
				v185 = v26 + v181
				v187 = int32(_a_F_g_intbig_picksplit_1)
				v195 = base.B2i32(v153&v187 == v176) | base.B2i32(v162&v187 == v176)
				if v195 != 0 {
					v196 = int32(1)
				} else {
					v196 = v153
				}
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v185+v196&int32(_a_F_g_intbig_picksplit_1)*int32(24))))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
				v205 = v203 & int32(4)
				if v205 != 0 {
					v206 = v181
				} else {
					v206 = v183
				}
				v207 = F_palloc(m, v206)
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v205
					*(*int32)(unsafe.Add(mBase, uint32(v207))) = v206 << (uint(int32(2)) % 32)
					v213 = int32(0)
					if v205|base.B2i32(v56 == v213) == v213 {
						v218 = int32(8)
						base.MemoryCopy(m, v207+v218, v202+v218, v56)
					} else {
					}
					if v195 != 0 {
						v225 = int32(2)
					} else {
						v225 = v162
					}
					v231 = *(*int32)(unsafe.Add(mBase, uint32(v185+v225&int32(_a_F_g_intbig_picksplit_1)*int32(24))))
					v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
					v234 = v232 & int32(4)
					if v234 != 0 {
						v235 = int32(8)
					} else {
						v235 = v183
					}
					v236 = F_palloc(m, v235)
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v234
						*(*int32)(unsafe.Add(mBase, uint32(v236))) = v235 << (uint(int32(2)) % 32)
						v242 = int32(0)
						if v234|base.B2i32(v56 == v242) == v242 {
							v247 = int32(8)
							base.MemoryCopy(m, v236+v247, v231+v247, v56)
						} else {
						}
						v255 = int32(_a_F_g_intbig_picksplit_1)
						v256 = v57 + v255
						v258 = v256 & v255
						v259 = F_palloc_mul(m, int32(8), v258)
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return int64(0)
						} else {
							if v57&int32(_a_F_g_intbig_picksplit_1) == int32(1) {
								F_pg_qsort(m, v259, v258, int32(8), int32(_a_F_g_intbig_picksplit_2))
								mBase = m.M
								v268 = m.ExcPending
								if v268 != 0 {
									return int64(0)
								} else {
									v816 = v180
									v819 = v164
									v831 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v816))) = uint16(v831)
									*(*uint16)(unsafe.Add(mBase, uint32(v819))) = uint16(v831)
									F_pfree(m, v259)
									mBase = m.M
									v836 = m.ExcPending
									if v836 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = base.I64_extend_i32_u(v236)
										*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = base.I64_extend_i32_u(v207)
										return v27 & int64(4294967295)
									}
								}
							} else {
								v269 = int32(1)
								v271 = v269
								v272 = v269
								for {
									v298 = v259 + v271<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v298-int32(8)))) = uint16(v272)
									v302 = int32(4)
									v307 = *(*int32)(unsafe.Add(mBase, uint32(v185+v271*int32(24))))
									v309 = Fn14312(m, v207, v307, v56, v302)
									mBase = m.M
									v311 = Fn14312(m, v236, v307, v56, int32(4))
									mBase = m.M
									v312 = v309 - v311
									v314 = v312 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v298-v302))) = v312 ^ v314 - v314
									v319 = v272 + int32(1)
									v320 = int32(_a_F_g_intbig_picksplit_1)
									v321 = v319 & v320
									if base.Ui32(v321) <= base.Ui32(v256&v320) {
										v271 = v321
										v272 = v319
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v259, v258, int32(8), int32(_a_F_g_intbig_picksplit_2))
								mBase = m.M
								v328 = m.ExcPending
								if v328 != 0 {
									return int64(0)
								} else {
									v329 = int32(1)
									if base.Ui32(v258) <= base.Ui32(v329) {
										v332 = v329
									} else {
										v332 = v258
									}
									v334 = v56 & int32(2147483644)
									v336 = v56 & int32(3)
									v337 = int32(8)
									v338 = v236 + v337
									v340 = v207 + v337
									v344 = int32(0)
									v354 = v180
									v357 = v164
									for {
										v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259+v344<<(uint(int32(3))%32)))))
										if v196&int32(_a_F_g_intbig_picksplit_1) == v372 {
											*(*uint16)(unsafe.Add(mBase, uint32(v354))) = uint16(v196)
											v375 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v375 + int32(1)
											v788 = v354 + int32(2)
											v791 = v357
										} else {
											if v225&int32(_a_F_g_intbig_picksplit_1) == v372 {
												*(*uint16)(unsafe.Add(mBase, uint32(v357))) = uint16(v225)
												v772 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v772 + int32(1)
												v788 = v354
												v791 = v357 + int32(2)
											} else {
												v388 = *(*int32)(unsafe.Add(mBase, uint32(v185+v372*int32(24))))
												v390 = Fn14312(m, v207, v388, v56, int32(4))
												mBase = m.M
												v393 = Fn14312(m, v236, v388, v56, int32(4))
												mBase = m.M
												v395 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
												v396 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
												v397 = v395 - v396
												if base.F64_lt(base.F64_convert_i32_s(v390), base.F64_add(base.F64_convert_i32_s(v393), base.F64_mul(base.F64_convert_i32_s(v397*v397*v397), float64(-1e-05)))) != 0 {
													v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
													if v405&int32(4) != 0 {
													} else {
														v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+4)))
														if v408&int32(4) != 0 {
															if v56 == int32(0) {
															} else {
																base.MemoryFill(m, v340, int32(255), v56)
															}
														} else {
															if v56 <= int32(0) {
															} else {
																v418 = v388 + int32(8)
																v419 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v56) {
																	v425 = v419
																	v426 = v419
																	for {
																		v449 = v425 + v340
																		v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
																		v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v418))))
																		v453 = v450 | v452
																		*(*uint8)(unsafe.Add(mBase, uint32(v449))) = uint8(v453)
																		v456 = v425 | int32(1)
																		v457 = v340 + v456
																		v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
																		v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456+v418))))
																		v461 = v458 | v460
																		*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v461)
																		v464 = v425 | int32(2)
																		v465 = v340 + v464
																		v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
																		v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v418))))
																		v469 = v466 | v468
																		*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v469)
																		v472 = v425 | int32(3)
																		v473 = v340 + v472
																		v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
																		v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v418))))
																		v477 = v474 | v476
																		*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v477)
																		v479 = int32(4)
																		v480 = v425 + v479
																		v482 = v426 + v479
																		if v482 != v334 {
																			v425 = v480
																			v426 = v482
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v336 == int32(0) {
																	} else {
																		v488 = v480
																		v513 = v488
																		v526 = v419
																		for {
																			v536 = v513 + v340
																			v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
																			v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v418))))
																			v540 = v537 | v539
																			*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v540)
																			v542 = int32(1)
																			v545 = v526 + v542
																			if v545 != v336 {
																				v513 = v513 + v542
																				v526 = v545
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v488 = v419
																	v513 = v488
																	v526 = v419
																	for {
																		v536 = v513 + v340
																		v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
																		v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v418))))
																		v540 = v537 | v539
																		*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v540)
																		v542 = int32(1)
																		v545 = v526 + v542
																		if v545 != v336 {
																			v513 = v513 + v542
																			v526 = v545
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
													*(*uint16)(unsafe.Add(mBase, uint32(v354))) = uint16(v372)
													v573 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v573 + int32(1)
													v788 = v354 + int32(2)
													v791 = v357
												} else {
													v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+4)))
													if v579&int32(4) != 0 {
													} else {
														v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+4)))
														if v582&int32(4) != 0 {
															if v56 == int32(0) {
															} else {
																base.MemoryFill(m, v338, int32(255), v56)
															}
														} else {
															if v56 <= int32(0) {
															} else {
																v592 = v388 + int32(8)
																v593 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v56) {
																	v599 = v593
																	v600 = v593
																	for {
																		v623 = v599 + v338
																		v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
																		v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v592))))
																		v627 = v624 | v626
																		*(*uint8)(unsafe.Add(mBase, uint32(v623))) = uint8(v627)
																		v630 = v599 | int32(1)
																		v631 = v338 + v630
																		v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
																		v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630+v592))))
																		v635 = v632 | v634
																		*(*uint8)(unsafe.Add(mBase, uint32(v631))) = uint8(v635)
																		v638 = v599 | int32(2)
																		v639 = v338 + v638
																		v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
																		v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v592))))
																		v643 = v640 | v642
																		*(*uint8)(unsafe.Add(mBase, uint32(v639))) = uint8(v643)
																		v646 = v599 | int32(3)
																		v647 = v338 + v646
																		v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647))))
																		v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v592))))
																		v651 = v648 | v650
																		*(*uint8)(unsafe.Add(mBase, uint32(v647))) = uint8(v651)
																		v653 = int32(4)
																		v654 = v599 + v653
																		v656 = v600 + v653
																		if v656 != v334 {
																			v599 = v654
																			v600 = v656
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v336 == int32(0) {
																	} else {
																		v662 = v654
																		v687 = v662
																		v700 = v593
																		for {
																			v710 = v687 + v338
																			v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
																			v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v592))))
																			v714 = v711 | v713
																			*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v714)
																			v716 = int32(1)
																			v719 = v700 + v716
																			if v719 != v336 {
																				v687 = v687 + v716
																				v700 = v719
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v662 = v593
																	v687 = v662
																	v700 = v593
																	for {
																		v710 = v687 + v338
																		v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
																		v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v592))))
																		v714 = v711 | v713
																		*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v714)
																		v716 = int32(1)
																		v719 = v700 + v716
																		if v719 != v336 {
																			v687 = v687 + v716
																			v700 = v719
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
													*(*uint16)(unsafe.Add(mBase, uint32(v357))) = uint16(v372)
													v772 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v772 + int32(1)
													v788 = v354
													v791 = v357 + int32(2)
												}
											}
										}
										v804 = v344 + int32(1)
										if v804 != v332 {
											v344 = v804
											v354 = v788
											v357 = v791
											continue
										} else {
											break
										}
										break
									}
									v816 = v788
									v819 = v791
									v831 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v816))) = uint16(v831)
									*(*uint16)(unsafe.Add(mBase, uint32(v819))) = uint16(v831)
									F_pfree(m, v259)
									mBase = m.M
									v836 = m.ExcPending
									if v836 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = base.I64_extend_i32_u(v236)
										*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = base.I64_extend_i32_u(v207)
										return v27 & int64(4294967295)
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
func F_g_intbig_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
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
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
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
	return int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v18)))) = int32(base.Ui32(v256) >> (uint(int32(2)) % 32))
	return base.I64_extend_i32_u(v50)
L17:
	;
	v67 = v47 & int32(3)
	v72 = v61
	v83 = v2
	goto L18
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)+v83*int32(24))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	if v93&int32(4) == int32(0) {
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
func F_gb18030_to_utf8(m *base.Module, l0 int32) int64 {
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(0)
	v7 = Fn14231(m, l0, int32(39), int32(_a_F_gb18030_to_utf8_0), v4, v4, int32(_a_F_gb18030_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = v14
	goto L3
L2:
	;
	v16 = int32(0)
	goto L3
L3:
	;
	v17 = F_palloc_mul(m, int32(4), v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
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
	v132 = int32(0)
	v139 = v132
	v141 = v132
	v142 = int32(1)
	goto L34
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v23 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = int32(0)
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v41 = v38 + v35<<(uint(int32(2))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = v44
	goto L13
L12:
	;
	v45 = int32(0)
	goto L13
L13:
	;
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v117 = v35 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v117 < v118 {
		v35 = v117
		goto L9
	} else {
		goto L33
	}
L15:
	;
	v48 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v49 <= v48 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = v45
	v59 = v48
	goto L17
L17:
	;
	v65 = v59 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v66)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = F_exprType(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L14
L19:
	;
	v91 = v56 + int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v91) < base.Ui32(v93+v94<<(uint(int32(2))%32)) {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v70 != v72 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v17))) = int32(-1)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v77 = v65 + v17
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = F_exprTypmod(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v81 == v41 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v79
	goto L19
L26:
	;
	goto L27
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v79 == v84 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(-1)
	goto L19
L29:
	;
	v99 = v91
	goto L31
L30:
	;
	v99 = int32(0)
	goto L31
L31:
	;
	v101 = v59 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v101 < v102 {
		v56 = v99
		v59 = v101
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
	v147 = int32(0)
	if l0 == v147 {
		v157 = v147
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v158 = int32(0)
	if l1 == v158 {
		v169 = v158
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v151 <= v139 {
		v157 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v157 = v153 + v139<<(uint(int32(2))%32)
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
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v163 <= v139 {
		v169 = int32(0)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v169 = v165 + v139<<(uint(int32(2))%32)
	goto L39
L42:
	;
	v187 = v139 << (uint(int32(2)) % 32)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v180+v187)))
	v190 = int32(0)
	v191 = base.I32_extend16_s(v142)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v187+v17)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v197 = F_makeVar(m, v190, v191, v192, v194, v195, v190)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L51
	}
L43:
	;
	v170 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v169 == v170)|(base.B2i32(v157 == v170)|base.B2i32(v174 <= v139)) == v170 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v182 = v158
	goto L45
L45:
	;
	F_pfree(m, v17)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L50
	}
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v180 != 0 {
		goto L42
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v182 = v141
	goto L45
L49:
	;
	goto L48
L50:
	;
	return v182
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v200 = F_pstrdup(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v203 = F_makeTargetEntry(m, v197, v191, v200, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v205 = int32(*(*int16)(unsafe.Add(mBase, uint32(v203)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+16)) = v205
	v207 = int32(1)
	v211 = F_lappend(m, v141, v203)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v139 = v139 + v207
	v141 = v211
	v142 = v142 + v207
	goto L34
}
func F_generate_subscripts_nodir(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_subscripts(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
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
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L53
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
	v34 = int32(0)
	v35 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v31 = F_palloc0_mul(m, int32(1), v14)
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
	v42 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v42)
	v48 = v34
	v49 = l1
	v51 = v35
	v53 = v40
	v55 = v39
	goto L15
L15:
	;
	v56 = l1 + l2
	v62 = v49
	goto L18
L16:
	;
	F_pfree(m, v53)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L52
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui32(v56) <= base.Ui32(v62) {
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v82 = v62
	goto L24
L20:
	;
	v70 = F_pg_mblen_range(m, v62, v56)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v73 = F_t_isalnum_with_len(m, v62, v70)
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
		v62 = v62 + v70
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v89 = F_pg_mblen_range(m, v82, v56)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	if v62 == int32(0) {
		goto L17
	} else {
		goto L32
	}
L26:
	;
	v91 = F_t_isalnum_with_len(m, v82, v89)
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
	v93 = v82 + v89
	if base.Ui32(v93) < base.Ui32(v56) {
		v82 = v93
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v95 = v82
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
	v100 = F_str_tolower(m, v62, v95-v62, int32(100))
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
	F_pfree(m, v53)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v114 = v53
	v115 = v55
	goto L36
L36:
	;
	if v102 != 0 {
		goto L39
	} else {
		goto L40
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
	v112 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v112)
	v114 = v110
	v115 = v109
	goto L36
L39:
	;
	base.MemoryCopy(m, v114+int32(2), v100, v102)
	goto L41
L40:
	;
	goto L41
L41:
	;
	F_pfree(m, v100)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v122 = int32(_a_F_generate_trgm_only_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v102+v114)+2)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_make_trigrams(m, l0, v114, v102+int32(3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v129 = int32(0)
	if v48 == v129 {
		v48 = v129
		v49 = v95
		v53 = v114
		v55 = v115
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v132 <= v51 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v147 = v145 + v124
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(1)
	v150 = v148 | v149
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v155 = v145 + v152 - v149
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v158 = v156 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v158)
	v48 = v145
	v49 = v95
	v51 = v146
	v53 = v114
	v55 = v115
	goto L15
L46:
	;
	v145 = v48
	v146 = v51
	goto L45
L47:
	;
	goto L48
L48:
	;
	v135 = F_mul_size(m, int32(1), v51)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v139 = F_mul_size(m, int32(1), v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v141 = F_repalloc0(m, v48, v135, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v141
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = v141
	v146 = v144
	goto L45
L52:
	;
	goto L13
L53:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_generate_trgm_only_1), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_generate_trgm_only_2), int32(115), int32(_a_F_generate_trgm_only_3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
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
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
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
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L86
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v336
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
	v336 = v22
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
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l1 + v38
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	v44 = l0 + l1
	v48 = F_palloc_mul(m, v38, l1+int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v54 = l0
	v61 = l1
	goto L11
L11:
	;
	v66 = int32(0)
	v70 = v54
	v72 = v66
	v73 = v66
	goto L13
L12:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_pfree(m, v48)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L68
	}
L13:
	;
	v82 = F_pg_mblen_range(m, v70, v44)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	if v61 <= v115-v54 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	goto L14
L16:
	;
	if v73&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v113 = v70 + v82
	if base.Ui32(v113) < base.Ui32(v44) {
		v70 = v113
		v72 = v110
		v73 = v111
		goto L13
	} else {
		goto L32
	}
L18:
	;
	v115 = v70
	v116 = v72
	v117 = v108
	goto L15
L19:
	;
	v86 = F_t_isalnum_with_len(m, v70, v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	switch v94 - int32(92) {
	case 0:
		v110 = v72
		v111 = int32(1)
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
	if v86 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v90 = int32(0)
	v110 = v90
	v111 = v90
	goto L17
L24:
	;
	goto L25
L25:
	;
	v108 = int32(1)
	goto L18
L26:
	;
	v101 = int32(0)
	v103 = F_t_isalnum_with_len(m, v70, v82)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L30
	}
L27:
	;
	v110 = int32(1)
	v111 = int32(0)
	goto L17
L28:
	;
	if v94 != int32(37) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v103 == int32(0) {
		v110 = v101
		v111 = v101
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v108 = v101
	goto L18
L32:
	;
	v115 = v113
	v116 = v110
	v117 = v111
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
	v125 = int32(_a_F_generate_wildcard_trgm_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v125)
	v127 = v48 + int32(2)
	goto L37
L36:
	;
	v127 = v48
	goto L37
L37:
	;
	if base.Ui32(v115) < base.Ui32(v44) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v228 = F_str_tolower(m, v48, v213-v48, int32(100))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L64
	}
L39:
	;
	v212 = v196
	v213 = v197 + int32(1)
	goto L38
L40:
	;
	v192 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v192)
	v196 = v189
	v197 = v132
	goto L39
L41:
	;
	v131 = v115
	v132 = v127
	v134 = v117
	goto L44
L42:
	;
	v173 = v115
	v174 = v127
	goto L43
L43:
	;
	v185 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v185)
	if v173 == int32(0) {
		goto L33
	} else {
		goto L63
	}
L44:
	;
	v143 = F_pg_mblen_range(m, v131, v44)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	v173 = v169
	v174 = v166
	goto L43
L46:
	;
	if v134&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v169 = v131 + v143
	if base.Ui32(v169) < base.Ui32(v44) {
		v131 = v169
		v132 = v166
		v134 = v167
		goto L44
	} else {
		goto L62
	}
L48:
	;
	if v143 != 0 {
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v147 = F_t_isalnum_with_len(m, v131, v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	switch v152 - int32(92) {
	case 0:
		v166 = v132
		v167 = int32(1)
		goto L47
	case 1, 2:
		goto L54
	case 3:
		v212 = v131
		v213 = v132
		goto L38
	default:
		goto L55
	}
L52:
	;
	if v147 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v189 = v131 - int32(1)
	goto L40
L54:
	;
	v157 = F_t_isalnum_with_len(m, v131, v143)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	if v152 == int32(37) {
		v212 = v131
		v213 = v132
		goto L38
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v157 == int32(0) {
		v189 = v131
		goto L40
	} else {
		goto L58
	}
L58:
	;
	goto L48
L59:
	;
	base.MemoryCopy(m, v132, v131, v143)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v166 = v132 + v143
	v167 = int32(0)
	goto L47
L62:
	;
	goto L45
L63:
	;
	v196 = v173
	v197 = v174
	goto L39
L64:
	;
	v230 = F_strlen(m, v228)
	mBase = m.M
	F_make_trigrams(m, v17+int32(4), v228, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v228)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v236 = l0 - v212 + l1
	if int32(0) < v236 {
		v54 = v212
		v61 = v236
		goto L11
	} else {
		goto L67
	}
L67:
	;
	goto L33
L68:
	;
	if int32(2) <= v254 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v260 = v253 + int32(5)
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_generate_wildcard_trgm[0]))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+272)))
	goto L73
L70:
	;
	v317 = v254
	goto L71
L71:
	;
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+4)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v317*int32(12) + int32(20)
	v336 = v253
	goto L2
L72:
	;
	v270 = int32(1)
	v273 = int32(0)
	goto L77
L73:
	;
	if v263 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_trigram_qsort_signed(m, v260, v254)
	mBase = m.M
	goto L72
L75:
	;
	goto L76
L76:
	;
	F_trigram_qsort_unsigned(m, v260, v254)
	mBase = m.M
	goto L72
L77:
	;
	v282 = int32(3)
	v284 = v260 + v270*v282
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v288 = v260 + v273*v282
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v285 != v289 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v317 = v308 + int32(1)
	goto L71
L79:
	;
	v310 = v270 + int32(1)
	if v310 != v254 {
		v270 = v310
		v273 = v308
		goto L77
	} else {
		goto L85
	}
L80:
	;
	v298 = v273 + int32(1)
	if v298 == v270 {
		v308 = v270
		goto L79
	} else {
		goto L84
	}
L81:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+1)))
	if v291 != v292 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+2)))
	if v294 == v295 {
		v308 = v273
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v302 = v260 + v298*int32(3)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v302)+2)) = uint8(v303)
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	*(*uint16)(unsafe.Add(mBase, uint32(v302))) = uint16(v305)
	v308 = v298
	goto L79
L85:
	;
	goto L78
L86:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_generate_wildcard_trgm_1), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_generate_wildcard_trgm_2), int32(115), int32(_a_F_generate_wildcard_trgm_3))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
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
	return v1088
L19:
	;
	v1083 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L62
	} else {
		goto L305
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
		v1088 = v210
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
	v226 = F_find_among(m, l0, int32(_a_F_german_ISO_8859_1_stem_2), int32(6), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L62
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v256
	v267 = v6 + int32(3)
	if v256 < v267 {
		goto L84
	} else {
		goto L85
	}
L67:
	;
	goto L66
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v228
	switch v226 - int32(1) {
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
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v256 <= v228 {
		goto L67
	} else {
		goto L83
	}
L71:
	;
	v252 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_3))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L62
	} else {
		goto L81
	}
L72:
	;
	v246 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L62
	} else {
		goto L79
	}
L73:
	;
	v240 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_5))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L62
	} else {
		goto L77
	}
L74:
	;
	v234 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_ISO_8859_1_stem_6))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L62
	} else {
		goto L75
	}
L75:
	;
	if int32(0) <= v234 {
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v1088 = v234
	goto L18
L77:
	;
	if int32(0) <= v240 {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v1088 = v240
	goto L18
L79:
	;
	if int32(0) <= v246 {
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v1088 = v246
	goto L18
L81:
	;
	if int32(0) <= v252 {
		goto L69
	} else {
		goto L82
	}
L82:
	;
	v1088 = v252
	goto L18
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v228 + int32(1)
	goto L69
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v499
	if v499 <= v6 {
		goto L151
	} else {
		goto L152
	}
L85:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v277 < v276 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v317 < int32(0) {
		goto L84
	} else {
		goto L101
	}
L87:
	;
	v279 = v276
	goto L89
L88:
	;
	v279 = v277
	goto L89
L89:
	;
	v286 = v276
	goto L91
L90:
	;
	v317 = v297
	goto L86
L91:
	;
	if v286 == v279 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v317 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v286))))
	if int32(252) < v292 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v309 = v286 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v309
	v286 = v309
	goto L91
L97:
	;
	v294 = v292 - int32(97)
	if v294 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v297 = int32(1)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v294)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v301)>>(uint(v294&int32(7))%32))&v297 != 0 {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L96
L101:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v321 = v320 + v317
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v332 < v321 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v375 < int32(0) {
		goto L84
	} else {
		goto L116
	}
L103:
	;
	v334 = v321
	goto L105
L104:
	;
	v334 = v332
	goto L105
L105:
	;
	v340 = v321
	goto L107
L106:
	;
	v375 = int32(1)
	goto L102
L107:
	;
	if v340 == v334 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v375 = int32(-1)
	goto L102
L110:
	;
	goto L111
L111:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v340))))
	if int32(252) < v349 {
		goto L106
	} else {
		goto L112
	}
L112:
	;
	v351 = v349 - int32(97)
	if v351 < int32(0) {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v357)>>(uint(v351&int32(7))%32))&int32(1) == int32(0) {
		goto L106
	} else {
		goto L114
	}
L114:
	;
	v366 = v340 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
	v340 = v366
	goto L107
L116:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v378 + v375
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	if v267 < v379 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v382 = v379
	goto L119
L118:
	;
	v382 = v267
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v382
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v392 < v391 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v432 < int32(0) {
		goto L84
	} else {
		goto L135
	}
L121:
	;
	v394 = v391
	goto L123
L122:
	;
	v394 = v392
	goto L123
L123:
	;
	v401 = v391
	goto L125
L124:
	;
	v432 = v412
	goto L120
L125:
	;
	if v401 == v394 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v432 = int32(-1)
	goto L120
L128:
	;
	goto L129
L129:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405+v401))))
	if int32(252) < v407 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v424 = v401 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v424
	v401 = v424
	goto L125
L131:
	;
	v409 = v407 - int32(97)
	if v409 < int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v412 = int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v409)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v416)>>(uint(v409&int32(7))%32))&v412 != 0 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	goto L130
L135:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v436 = v435 + v432
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v436
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v447 < v436 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v490 < int32(0) {
		goto L84
	} else {
		goto L150
	}
L137:
	;
	v449 = v436
	goto L139
L138:
	;
	v449 = v447
	goto L139
L139:
	;
	v455 = v436
	goto L141
L140:
	;
	v490 = int32(1)
	goto L136
L141:
	;
	if v455 == v449 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v490 = int32(-1)
	goto L136
L144:
	;
	goto L145
L145:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+v455))))
	if int32(252) < v464 {
		goto L140
	} else {
		goto L146
	}
L146:
	;
	v466 = v464 - int32(97)
	if v466 < int32(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v466)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v472)>>(uint(v466&int32(7))%32))&int32(1) == int32(0) {
		goto L140
	} else {
		goto L148
	}
L148:
	;
	v481 = v455 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v481
	v455 = v481
	goto L141
L150:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v493 + v490
	goto L84
L151:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v662
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v662
	v666 = v662 - int32(1)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v666 <= v667 {
		goto L193
	} else {
		goto L194
	}
L152:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v505 = int32(1)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v499-v505))))
	if base.B2i32(v507&int32(224) != int32(96))|base.B2i32(v505<<(uint(v507)%32)&int32(_a_F_german_ISO_8859_1_stem_7) == int32(0)) != 0 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v522 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_8), int32(11), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L62
	} else {
		goto L154
	}
L154:
	;
	if v522 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v526
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v526 < v528 {
		goto L151
	} else {
		goto L156
	}
L156:
	;
	switch v522 - int32(1) {
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
	v656 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_9))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L62
	} else {
		goto L191
	}
L158:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L180
L159:
	;
	v559 = F_slice_del(m, l0)
	mBase = m.M
	if v559 < int32(0) {
		v1088 = v559
		goto L18
	} else {
		goto L169
	}
L160:
	;
	v556 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v556 {
		goto L151
	} else {
		goto L168
	}
L161:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v533 = int32(4)
	v535 = int32(0)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v537-v538 < v533 {
		v548 = v535
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if v548 != 0 {
		goto L151
	} else {
		goto L166
	}
L163:
	;
	goto L162
L164:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v544 = F_memcmp(m, v541+v537-v533, int32(_a_F_german_ISO_8859_1_stem_10), v533)
	mBase = m.M
	if v544 != 0 {
		v548 = v535
		goto L163
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v537 - v533
	v548 = int32(1)
	goto L163
L166:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v549 + (v526 - v532)
	v553 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v553 {
		goto L151
	} else {
		goto L167
	}
L167:
	;
	v1088 = v553
	goto L18
L168:
	;
	v1088 = v556
	goto L18
L169:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v562
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v562 <= v564 {
		goto L151
	} else {
		goto L170
	}
L170:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566+v562-int32(1)))))
	if v570 != int32(115) {
		goto L151
	} else {
		goto L171
	}
L171:
	;
	v574 = v562 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v574
	v577 = int32(3)
	v579 = int32(0)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v574-v582 < v577 {
		v592 = v579
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v592 == int32(0) {
		goto L151
	} else {
		goto L176
	}
L173:
	;
	goto L172
L174:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v588 = F_memcmp(m, v585+v574-v577, int32(_a_F_german_ISO_8859_1_stem_11), v577)
	mBase = m.M
	if v588 != 0 {
		v592 = v579
		goto L173
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v574 - v577
	v592 = int32(1)
	goto L173
L176:
	;
	v595 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v595 {
		goto L151
	} else {
		goto L177
	}
L177:
	;
	v1088 = v595
	goto L18
L178:
	;
	if v650 != 0 {
		goto L151
	} else {
		goto L189
	}
L179:
	;
	v650 = v646
	goto L178
L180:
	;
	if v606 <= v607 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v646 = int32(0)
	goto L179
L182:
	;
	v650 = int32(-1)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v619 = int32(1)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620+v606-v619))))
	if int32(116) < v624 {
		v646 = v619
		goto L179
	} else {
		goto L185
	}
L185:
	;
	v626 = v624 - int32(98)
	if v626 < int32(0) {
		v646 = v619
		goto L179
	} else {
		goto L186
	}
L186:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v626)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v632)>>(uint(v626&int32(7))%32))&int32(1) == int32(0) {
		v646 = v619
		goto L179
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v606 - int32(1)
	goto L188
L188:
	;
	goto L181
L189:
	;
	v651 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v651 {
		goto L151
	} else {
		goto L190
	}
L190:
	;
	v1088 = v651
	goto L18
L191:
	;
	if v656 < int32(0) {
		v1088 = v656
		goto L18
	} else {
		goto L192
	}
L192:
	;
	goto L151
L193:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v853
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v853
	v857 = v853 - int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v857 <= v858 {
		goto L235
	} else {
		goto L236
	}
L194:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669+v666))))
	if base.B2i32(v671&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v671)%32)&int32(_a_F_german_ISO_8859_1_stem_12) == int32(0)) != 0 {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v686 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_13), int32(5), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L62
	} else {
		goto L196
	}
L196:
	;
	if v686 == int32(0) {
		goto L193
	} else {
		goto L197
	}
L197:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v690 < v692 {
		goto L193
	} else {
		goto L198
	}
L198:
	;
	switch v686 - int32(1) {
	case 0:
		goto L201
	case 1:
		goto L200
	case 2:
		goto L199
	default:
		goto L193
	}
L199:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L219
L200:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L205
L201:
	;
	v696 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v696 {
		goto L193
	} else {
		goto L202
	}
L202:
	;
	v1088 = v696
	goto L18
L203:
	;
	if v751 != 0 {
		goto L193
	} else {
		goto L214
	}
L204:
	;
	v751 = v747
	goto L203
L205:
	;
	if v707 <= v708 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v747 = int32(0)
	goto L204
L207:
	;
	v751 = int32(-1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v720 = int32(1)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721+v707-v720))))
	if int32(116) < v725 {
		v747 = v720
		goto L204
	} else {
		goto L210
	}
L210:
	;
	v727 = v725 - int32(98)
	if v727 < int32(0) {
		v747 = v720
		goto L204
	} else {
		goto L211
	}
L211:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v727)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v733)>>(uint(v727&int32(7))%32))&int32(1) == int32(0) {
		v747 = v720
		goto L204
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v707 - int32(1)
	goto L213
L213:
	;
	goto L206
L214:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v754 = v752 - int32(3)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v754 < v755 {
		goto L193
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v754
	v758 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v758 {
		goto L193
	} else {
		goto L216
	}
L216:
	;
	v1088 = v758
	goto L18
L217:
	;
	if v814 != 0 {
		goto L193
	} else {
		goto L228
	}
L218:
	;
	v814 = v810
	goto L217
L219:
	;
	if v770 <= v771 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v810 = int32(0)
	goto L218
L221:
	;
	v814 = int32(-1)
	goto L217
L222:
	;
	goto L223
L223:
	;
	v783 = int32(1)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784+v770-v783))))
	if int32(228) < v788 {
		v810 = v783
		goto L218
	} else {
		goto L224
	}
L224:
	;
	v790 = v788 - int32(85)
	if v790 < int32(0) {
		v810 = v783
		goto L218
	} else {
		goto L225
	}
L225:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v790)>>(uint(int32(3))%32)))+uint32(_c_F_german_ISO_8859_1_stem[3]))))
	if int32(base.Ui32(v796)>>(uint(v790&int32(7))%32))&int32(1) == int32(0) {
		v810 = v783
		goto L218
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v770 - int32(1)
	goto L227
L227:
	;
	goto L220
L228:
	;
	v815 = v690 - v761
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v817 = v815 + v816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v817
	v820 = v817 - int32(1)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v820 <= v821 {
		v844 = v817
		goto L229
	} else {
		goto L230
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v844
	v847 = F_slice_del(m, l0)
	mBase = m.M
	if v847 < int32(0) {
		v1088 = v847
		goto L18
	} else {
		goto L234
	}
L230:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823+v820))))
	if base.B2i32(v825&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v825)%32)&int32(_a_F_german_ISO_8859_1_stem_14) == int32(0)) != 0 {
		v844 = v817
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v840 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_15), int32(5), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L62
	} else {
		goto L232
	}
L232:
	;
	if v840 != 0 {
		goto L193
	} else {
		goto L233
	}
L233:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v844 = v842 + v815
	goto L229
L234:
	;
	goto L193
L235:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1031
	v1034 = v1031
	goto L286
L236:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+v857))))
	if base.B2i32(v862&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v862)%32)&int32(_a_F_german_ISO_8859_1_stem_16) == int32(0)) != 0 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v877 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_17), int32(8), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L62
	} else {
		goto L238
	}
L238:
	;
	if v877 == int32(0) {
		goto L235
	} else {
		goto L239
	}
L239:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v881
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v881 < v883 {
		goto L235
	} else {
		goto L240
	}
L240:
	;
	switch v877 - int32(1) {
	case 0:
		goto L244
	case 1:
		goto L243
	case 2:
		goto L242
	case 3:
		goto L241
	default:
		goto L235
	}
L241:
	;
	v991 = F_slice_del(m, l0)
	mBase = m.M
	if v991 < int32(0) {
		v1088 = v991
		goto L18
	} else {
		goto L277
	}
L242:
	;
	v938 = F_slice_del(m, l0)
	mBase = m.M
	if v938 < int32(0) {
		v1088 = v938
		goto L18
	} else {
		goto L262
	}
L243:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v926 < v881 {
		goto L257
	} else {
		goto L258
	}
L244:
	;
	v887 = F_slice_del(m, l0)
	mBase = m.M
	if v887 < int32(0) {
		v1088 = v887
		goto L18
	} else {
		goto L245
	}
L245:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v890
	v892 = int32(2)
	v894 = int32(0)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v890-v897 < v892 {
		v907 = v894
		goto L247
	} else {
		goto L248
	}
L246:
	;
	if v907 == int32(0) {
		goto L235
	} else {
		goto L250
	}
L247:
	;
	goto L246
L248:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v903 = F_memcmp(m, v900+v890-v892, int32(_a_F_german_ISO_8859_1_stem_18), v892)
	mBase = m.M
	if v903 != 0 {
		v907 = v894
		goto L247
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v890 - v892
	v907 = int32(1)
	goto L247
L250:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v910
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v912 < v910 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914+v910-int32(1)))))
	if v918 == int32(101) {
		goto L235
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v910 < v921 {
		goto L235
	} else {
		goto L255
	}
L254:
	;
	goto L253
L255:
	;
	v923 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v923 {
		goto L235
	} else {
		goto L256
	}
L256:
	;
	v1088 = v923
	goto L18
L257:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928+v881-int32(1)))))
	if v932 == int32(101) {
		goto L235
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v935 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v935 {
		goto L235
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	v1088 = v935
	goto L18
L262:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v941
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v944 = int32(2)
	v946 = int32(0)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v941-v949 < v944 {
		v959 = v946
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if v959 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	goto L263
L265:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v955 = F_memcmp(m, v952+v941-v944, int32(_a_F_german_ISO_8859_1_stem_19), v944)
	mBase = m.M
	if v955 != 0 {
		v959 = v946
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v941 - v944
	v959 = int32(1)
	goto L264
L267:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v964 = v962 + (v941 - v943)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964
	v966 = int32(2)
	v968 = int32(0)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v964-v971 < v966 {
		v981 = v968
		goto L271
	} else {
		goto L272
	}
L268:
	;
	goto L269
L269:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v984
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v984 < v986 {
		goto L235
	} else {
		goto L275
	}
L270:
	;
	if v981 == int32(0) {
		goto L235
	} else {
		goto L274
	}
L271:
	;
	goto L270
L272:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v977 = F_memcmp(m, v974+v964-v966, int32(_a_F_german_ISO_8859_1_stem_20), v966)
	mBase = m.M
	if v977 != 0 {
		v981 = v968
		goto L271
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964 - v966
	v981 = int32(1)
	goto L271
L274:
	;
	goto L269
L275:
	;
	v988 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v988 {
		goto L235
	} else {
		goto L276
	}
L276:
	;
	v1088 = v988
	goto L18
L277:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v994
	v997 = v994 - int32(1)
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v997 <= v998 {
		goto L235
	} else {
		goto L278
	}
L278:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000+v997))))
	if base.Ui32(int32(1)) < base.Ui32((v1002-int32(103))&int32(255)) {
		goto L235
	} else {
		goto L279
	}
L279:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1013 = F_find_among_b(m, l0, int32(_a_F_german_ISO_8859_1_stem_21), int32(2), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L62
	} else {
		goto L280
	}
L280:
	;
	if v1013 == int32(0) {
		goto L235
	} else {
		goto L281
	}
L281:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1017
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1017 < v1019 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1021 + (v994 - v1009)
	goto L235
L283:
	;
	goto L284
L284:
	;
	v1025 = F_slice_del(m, l0)
	mBase = m.M
	if v1025 < int32(0) {
		v1088 = v1025
		goto L18
	} else {
		goto L285
	}
L285:
	;
	goto L235
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1034
	v1042 = F_find_among(m, l0, int32(_a_F_german_ISO_8859_1_stem_22), int32(6), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L62
	} else {
		goto L289
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1031
	v1088 = int32(1)
	goto L18
L288:
	;
	goto L287
L289:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1044
	switch v1042 - int32(1) {
	case 0:
		goto L295
	case 1:
		goto L294
	case 2:
		goto L293
	case 3:
		goto L292
	case 4:
		goto L291
	default:
		goto L290
	}
L290:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1034 = v1078
	goto L286
L291:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1072 <= v1044 {
		goto L288
	} else {
		goto L304
	}
L292:
	;
	v1068 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_23))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L62
	} else {
		goto L302
	}
L293:
	;
	v1062 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_24))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L62
	} else {
		goto L300
	}
L294:
	;
	v1056 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_25))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L62
	} else {
		goto L298
	}
L295:
	;
	v1050 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_ISO_8859_1_stem_26))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L62
	} else {
		goto L296
	}
L296:
	;
	if int32(0) <= v1050 {
		goto L290
	} else {
		goto L297
	}
L297:
	;
	v1088 = v1050
	goto L18
L298:
	;
	if int32(0) <= v1056 {
		goto L290
	} else {
		goto L299
	}
L299:
	;
	v1088 = v1056
	goto L18
L300:
	;
	if int32(0) <= v1062 {
		goto L290
	} else {
		goto L301
	}
L301:
	;
	v1088 = v1062
	goto L18
L302:
	;
	if int32(0) <= v1068 {
		goto L290
	} else {
		goto L303
	}
L303:
	;
	v1088 = v1068
	goto L18
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1044 + int32(1)
	goto L290
L305:
	;
	if int32(0) <= v1083 {
		goto L17
	} else {
		goto L306
	}
L306:
	;
	v1088 = v1083
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
	v20 = F_palloc(m, int32(32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v18 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v125 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(3)
	goto L15
L18:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(3)
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v18 + (v45+int32(3))&int32(-4)
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v18 + v45
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v66 & int32(268435455)
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v66&int32(268435455) - v92
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v166 + (v45 - v122)
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
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = int32(3)
	F_ScanKeyInit(m, v8, v10, v10, int32(184), base.I64_extend_i32_u(l0))
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
	F_ScanKeyInit(m, v8+int32(56), int32(18), int32(3), int32(61), base.I64_extend_i32_s(l1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = F_table_beginscan_catalog(m, v29, int32(2), v8)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	v39 = F_heap_getnext(m, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+188))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	m.T0[v49].(func(*base.Module, int32))(m, v32)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v45 = F_lappend_oid(m, v34, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v34 = v45
	goto L6
L13:
	;
	F_relation_close(m, v29, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	m.G0 = v8 + int32(112)
	return v34
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
							F_errfinish(m, int32(_a_F_getWeights_1), int32(443), int32(_a_F_getWeights_2))
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
									F_errfinish(m, int32(_a_F_getWeights_1), int32(448), int32(_a_F_getWeights_2))
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
														F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
													F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
															F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
												F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
															F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
														F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
																F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
															F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
														F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
																F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
													F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
																F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
															F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
																	F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
											F_errfinish(m, int32(_a_F_getWeights_1), int32(457), int32(_a_F_getWeights_2))
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
					F_errfinish(m, int32(_a_F_getWeights_1), int32(438), int32(_a_F_getWeights_2))
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
			v23 = F_pstrdup(m, v18+v19+int32(4))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v14)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v42 = v23
					m.G0 = v9 + int32(16)
					return v42
				}
			}
		} else {
			if l2 != 0 {
				v42 = int32(0)
				m.G0 = v9 + int32(16)
				return v42
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_get_attname_0), v9)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_attname_1), int32(1069), int32(_a_F_get_attname_2))
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
	switch v9 - int32(283) {
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
	F_errfinish(m, int32(_a_F_get_indexpath_pages_1), int32(992), int32(_a_F_get_indexpath_pages_2))
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 float64
	_ = v722
	var v723 float64
	_ = v723
	var v724 float64
	_ = v724
	var v725 int32
	_ = v725
	var v726 float64
	_ = v726
	var v728 float64
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v749 int32
	_ = v749
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v45 == int32(0) {
		v205 = v8
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
	v216 = F_generate_join_implied_equalities(m, l0, v27, l5, l1, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L60
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 <= int32(0) {
		v205 = v8
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = int32(0)
	v62 = v8
	v63 = v8
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v60<<(uint(int32(2))%32))))
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
	v205 = v189
	goto L14
L19:
	;
	v191 = v60 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v191 < v192 {
		v60 = v191
		v62 = v188
		v63 = v189
		goto L17
	} else {
		goto L58
	}
L20:
	;
	if v92 == int32(0) {
		v188 = v62
		v189 = v63
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
		v188 = v62
		v189 = v63
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
		v188 = v62
		v189 = v63
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
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+11)))
	if v127 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v183 = F_lappend(m, v63, v77)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L56
	}
L36:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+12)))
	if v130 != int32(1) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	v134 = int32(0)
	if base.B2i32(v133 == v134)|base.B2i32(v27 == v134) != 0 {
		v179 = v134
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	if v179 != 0 {
		v188 = v62
		v189 = v63
		goto L19
	} else {
		goto L53
	}
L41:
	;
	goto L40
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v144 < v145 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v147 = v144
	goto L45
L44:
	;
	v147 = v145
	goto L45
L45:
	;
	if v147 <= int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v150 = int32(1)
	goto L48
L47:
	;
	v150 = v147
	goto L48
L48:
	;
	v151 = int32(8)
	v156 = int32(0)
	goto L49
L49:
	;
	v163 = v156 << (uint(int32(2)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v27+v151+v163)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v133+v151+v163)))
	v168 = v165 & v167
	v170 = base.B2i32(v168 != int32(0))
	if v168 != 0 {
		v179 = v170
		goto L41
	} else {
		goto L51
	}
L50:
	;
	v179 = v170
	goto L41
L51:
	;
	v172 = v156 + int32(1)
	if v172 != v150 {
		v156 = v172
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	v181 = F_bms_is_member(m, v180, v62)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v181 != 0 {
		v188 = v62
		v189 = v63
		goto L19
	} else {
		goto L55
	}
L55:
	;
	goto L35
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	v186 = F_bms_add_member(m, v62, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v188 = v186
	v189 = v183
	goto L19
L58:
	;
	goto L18
L59:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v606 = F_list_concat(m, v595, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L149
	}
L60:
	;
	if v216 == int32(0) {
		v595 = v205
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v220 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v220 < v221 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v232 = v220
	v233 = int32(0)
	v236 = v205
	goto L65
L63:
	;
	v301 = v220
	v305 = v205
	goto L64
L64:
	;
	if v301 == int32(0) {
		v595 = v305
		goto L59
	} else {
		goto L83
	}
L65:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246+v233<<(uint(int32(2))%32))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v255 = F_bms_is_subset(m, v254, v37)
	mBase = m.M
	if v255 == v253 {
		v266 = v253
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v301 = v288
	v305 = v289
	goto L64
L67:
	;
	v291 = v233 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v291 < v292 {
		v232 = v288
		v233 = v291
		v236 = v289
		goto L65
	} else {
		goto L82
	}
L68:
	;
	if v266 != 0 {
		v288 = v232
		v289 = v236
		goto L67
	} else {
		goto L72
	}
L69:
	;
	goto L68
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v259 = F_bms_overlap(m, v252, v258)
	mBase = m.M
	if v259 == int32(0) {
		v266 = v253
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+40))
	v263 = F_bms_overlap(m, v252, v262)
	mBase = m.M
	v266 = v263 ^ int32(1)
	goto L69
L72:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v269 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v271 = F_bms_is_subset(m, v270, v44)
	mBase = m.M
	if v271 == v269 {
		v282 = v269
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v282 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L73
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v250)+28))
	v275 = F_bms_overlap(m, v268, v274)
	mBase = m.M
	if v275 == int32(0) {
		v282 = v269
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v250)+40))
	v279 = F_bms_overlap(m, v268, v278)
	mBase = m.M
	v282 = v279 ^ int32(1)
	goto L74
L77:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v250)+100))
	v284 = F_lappend(m, v232, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v286 = F_lappend(m, v236, v250)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L81
	}
L80:
	;
	v288 = v284
	v289 = v236
	goto L67
L81:
	;
	v288 = v232
	v289 = v286
	goto L67
L82:
	;
	goto L66
L83:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	v319 = F_bms_union(m, v318, l5)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L85
	}
L84:
	;
	if v489 == int32(0) {
		v595 = v305
		goto L59
	} else {
		goto L123
	}
L85:
	;
	v321 = int32(0)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if base.Ui32(int32(5)) < base.Ui32(v324) {
		v336 = v323
		v337 = v319
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v338 = int32(0)
	if v301 == v338 {
		v489 = v338
		goto L84
	} else {
		goto L90
	}
L87:
	;
	if int32(1)<<(uint(v324)%32)&int32(44) == int32(0) {
		v336 = v323
		v337 = v319
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)+252))
	v334 = F_bms_union(m, l5, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v336 = v333
	v337 = v334
	goto L86
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if int32(0) < v341 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v353 = int32(0)
	v355 = v321
	goto L94
L92:
	;
	v457 = v321
	goto L93
L93:
	;
	v489 = v457
	goto L84
L94:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v353<<(uint(int32(2))%32))))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+40)))
	if v371 != 0 {
		v441 = v355
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v457 = v441
	goto L93
L96:
	;
	v444 = v353 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v444 < v445 {
		v353 = v444
		v355 = v441
		goto L94
	} else {
		goto L122
	}
L97:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	if v372 == int32(0) {
		v441 = v355
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	if v375 < int32(2) {
		v441 = v355
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v370)+36))
	v379 = int32(0)
	if base.B2i32(v378 == v379)|base.B2i32(v337 == v379) != 0 {
		v424 = v379
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v424 == int32(0) {
		v441 = v355
		goto L96
	} else {
		goto L113
	}
L101:
	;
	goto L100
L102:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v389 < v390 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v392 = v389
	goto L105
L104:
	;
	v392 = v390
	goto L105
L105:
	;
	if v392 <= int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v395 = int32(1)
	goto L108
L107:
	;
	v395 = v392
	goto L108
L108:
	;
	v396 = int32(8)
	v401 = int32(0)
	goto L109
L109:
	;
	v408 = v401 << (uint(int32(2)) % 32)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v337+v396+v408)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v378+v396+v408)))
	v413 = v410 & v412
	v415 = base.B2i32(v413 != int32(0))
	if v413 != 0 {
		v424 = v415
		goto L101
	} else {
		goto L111
	}
L110:
	;
	v424 = v415
	goto L101
L111:
	;
	v417 = v401 + int32(1)
	if v417 != v395 {
		v401 = v417
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+42)))
	if v427 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v439 = F_list_concat(m, v355, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L121
	}
L115:
	;
	v430 = F_generate_join_implied_equalities_normal(m, l0, v370, v319, l5, v323)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v436 = F_generate_join_implied_equalities_broken(m, l0, v370, v337, l5, v336, v322)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+42)))
	if v432 != int32(1) {
		v438 = v430
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v438 = v436
	goto L114
L121:
	;
	v441 = v439
	goto L96
L122:
	;
	goto L95
L123:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v492 <= int32(0) {
		v595 = v305
		goto L59
	} else {
		goto L124
	}
L124:
	;
	v504 = int32(0)
	v507 = v305
	goto L125
L125:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v504<<(uint(int32(2))%32))))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v524 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v521)+28))
	v526 = F_bms_is_subset(m, v525, v37)
	mBase = m.M
	if v526 == v524 {
		v537 = v524
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v595 = v579
	goto L59
L127:
	;
	v581 = v504 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v581 < v582 {
		v504 = v581
		v507 = v579
		goto L125
	} else {
		goto L148
	}
L128:
	;
	if v537 != 0 {
		v579 = v507
		goto L127
	} else {
		goto L132
	}
L129:
	;
	goto L128
L130:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v521)+28))
	v530 = F_bms_overlap(m, v523, v529)
	mBase = m.M
	if v530 == int32(0) {
		v537 = v524
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v521)+40))
	v534 = F_bms_overlap(m, v523, v533)
	mBase = m.M
	v537 = v534 ^ int32(1)
	goto L129
L132:
	;
	v538 = int32(0)
	if v507 == v538 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v576 != 0 {
		v579 = v507
		goto L127
	} else {
		goto L146
	}
L134:
	;
	v576 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if v544 <= int32(0) {
		v570 = v538
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v576 = v570
	goto L133
L138:
	;
	v547 = int32(0)
	if v547 < v544 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v550 = v544
	goto L141
L140:
	;
	v550 = v547
	goto L141
L141:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v507)+12))
	v553 = int32(0)
	goto L142
L142:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v551+v553<<(uint(int32(2))%32))))
	v562 = base.B2i32(v561 == v521)
	if v561 == v521 {
		v570 = v562
		goto L137
	} else {
		goto L144
	}
L143:
	;
	v570 = v562
	goto L137
L144:
	;
	v564 = v553 + int32(1)
	if v564 != v550 {
		v553 = v564
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v577 = F_lappend(m, v507, v521)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v579 = v577
	goto L127
L148:
	;
	goto L126
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v606
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v609 == int32(0) {
		v706 = v606
		goto L151
	} else {
		goto L152
	}
L150:
	;
	return v749
L151:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v722 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v723 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v724 = F_calc_joinrel_size_estimate(m, l0, l1, v720, v721, v722, v723, l4, v706)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L4
	} else {
		goto L169
	}
L152:
	;
	v612 = int32(0)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v613 <= v612 {
		v706 = v606
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v624 = v612
	goto L154
L154:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v637+v624<<(uint(int32(2))%32))))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	v643 = int32(0)
	if base.B2i32(v642 == v643)|base.B2i32(l5 == v643) != 0 {
		v689 = base.B2i32(v642|l5 == v643)
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v706 = v698
	goto L151
L156:
	;
	if v689 != 0 {
		v749 = v641
		goto L150
	} else {
		goto L167
	}
L157:
	;
	goto L156
L158:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v657 != v658 {
		v689 = int32(0)
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v660 = int32(1)
	if v657 <= v660 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v663 = v660
	goto L162
L161:
	;
	v663 = v657
	goto L162
L162:
	;
	v664 = int32(8)
	v669 = int32(0)
	goto L163
L163:
	;
	v677 = v669 << (uint(int32(2)) % 32)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v642+v664+v677)))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l5+v664+v677)))
	v682 = base.B2i32(v679 == v681)
	if v679 != v681 {
		v689 = v682
		goto L157
	} else {
		goto L165
	}
L164:
	;
	v689 = v682
	goto L157
L165:
	;
	v685 = v669 + int32(1)
	if v685 != v663 {
		v669 = v685
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v695 = v624 + int32(1)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v695 < v696 {
		v624 = v695
		goto L154
	} else {
		goto L168
	}
L168:
	;
	goto L155
L169:
	;
	v726 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v724, v726) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v728 = v726
	goto L172
L171:
	;
	v728 = v724
	goto L172
L172:
	;
	v730 = F_palloc0(m, int32(24))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v730)+16)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v730)+8)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v730)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = int32(281)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v739 = F_lappend(m, v738, v730)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v739
	v749 = v730
	goto L150
}
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	v25 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 <= v25 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = l2 + int32(8)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v38 = v35 & int32(-1)
	if v38 != 0 {
		v60 = v38
		v61 = v25
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v39 = int32(1)
	if v39 == v28 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v43 = v39
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
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
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
	if base.Ui32(int32(2)) <= base.Ui32(v109-int32(303)) {
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
	if v109 != int32(293) {
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
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
	v173 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v176 <= v173 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v179 = v165 + int32(8)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v186 = v183 & int32(-1)
	if v186 != 0 {
		v208 = v186
		v209 = v173
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v187 = int32(1)
	if v187 == v176 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v191 = v187
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
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
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
	if base.Ui32(int32(2)) <= base.Ui32(v257-int32(303)) {
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
	if v257 != int32(293) {
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
	v283 = int32(base.Ui32(v281) >> (uint(int32(5)) % 32))
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
	v409 = int32(base.Ui32(v407) >> (uint(int32(5)) % 32))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+100))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
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
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 {
	case 0:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = int64(0)
		v16 = F_GetSysCacheOid(m, int32(10), base.I64_extend_i32_u(v11), v13, v13, v13)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if l1|v16 != 0 {
				v76 = v16
				m.G0 = v7 + int32(48)
				return v76
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v11
						F_errmsg(m, int32(_a_F_get_rolespec_oid_0), v7+int32(16))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_2), int32(_a_F_get_rolespec_oid_3))
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
	case 1, 2:
		v75 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_oid[0]))
		v76 = v75
		m.G0 = v7 + int32(48)
		return v76
	case 3:
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_oid[1]))
		v76 = v40
		m.G0 = v7 + int32(48)
		return v76
	case 4:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_get_rolespec_oid_4)
				F_errmsg(m, int32(_a_F_get_rolespec_oid_0), v7+int32(32))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_5), int32(_a_F_get_rolespec_oid_6))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
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
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
			F_errmsg_internal(m, int32(_a_F_get_rolespec_oid_7), v7)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_get_rolespec_oid_1), int32(_a_F_get_rolespec_oid_8), int32(_a_F_get_rolespec_oid_6))
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
					F_errfinish(m, int32(_a_F_get_segment_by_index_1), int32(1831), int32(_a_F_get_segment_by_index_2))
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
							F_errfinish(m, int32(_a_F_get_segment_by_index_1), int32(1838), int32(_a_F_get_segment_by_index_2))
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
func F_get_sequences_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v17 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v23 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v8 + int32(16)
	return
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_get_sequences_string[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = int32(2)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v23<<(uint(v29)%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27+v32<<(uint(v29)%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v37 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_get_sequences_string_0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = base.I64_rotl(v43, int64(32))
	F_appendStringInfo(m, l1, int32(_a_F_get_sequences_string_1), v8)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v51 = v23 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 < v52 {
		v23 = v51
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L6
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
	F_errfinish(m, int32(_a_F_get_switched_clauses_1), int32(2497), int32(_a_F_get_switched_clauses_2))
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
	F_errfinish(m, int32(_a_F_get_switched_clauses_1), int32(2503), int32(_a_F_get_switched_clauses_2))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+144))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_get_typlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13+v14)+76)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.I32_extend16_s(v16)
			}
		}
	}
}
func F_get_typstorage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(112)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+v14)+129)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v16)
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
	F_errfinish(m, int32(_a_F_get_view_query_1), int32(2577), int32(_a_F_get_view_query_2))
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
	F_errfinish(m, int32(_a_F_get_view_query_1), int32(2583), int32(_a_F_get_view_query_2))
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
	if v88 != int32(453) {
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v303 int32
	_ = v303
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 float64
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 float64
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int64
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v650 int32
	_ = v650
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v723 int64
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v809 int32
	_ = v809
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1059 int32
	_ = v1059
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 float64
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1088 float64
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1382 int32
	_ = v1382
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1456 int64
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1588 int64
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1666 int32
	_ = v1666
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	v31 = m.G0
	v33 = v31 - int32(_a_F_ginbulkdelete_0)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0]))
	v42 = F_AllocSetContextCreateInternal(m, v37, int32(_a_F_ginbulkdelete_1), int32(0), int32(_a_F_ginbulkdelete_2), int32(_a_F_ginbulkdelete_3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+2764)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+2760)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v33)+2752)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[1]))) = v42
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[2]))) = v50
	v53 = v33 + int32(2768)
	F_initGinState(m, v53, v35)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v59 = F_palloc0(m, int32(40))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v61 = l1
	goto L6
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[3]))
	F_ginInsertCleanup(m, v53, base.B2i32(v63 != int32(4)), int32(0), int32(1), v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v61 = v59
	goto L6
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+2756)) = v61
	v73 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v77 = F_ReadBufferExtended(m, v35, v73, int32(1), v73, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v83 = int32(1)
	v85 = v77
	goto L10
L10:
	;
	if int32(0) <= v85 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v303 = v85
	goto L39
L12:
	;
	F_LockBufferInternal(m, v85, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v127 = v113 + v85<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+(v85^int32(-1))<<(uint(int32(2))%32))))
	v127 = v126
	goto L12
L16:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+16)))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v131)+6)))
	if v133&int32(2) == int32(0) {
		v223 = v127
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L11
L18:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	v255 = v223 + v252&int32(_a_F_ginbulkdelete_4)
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+2)))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255))))
	F_UnlockReleaseBuffer(m, v85)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L37
	}
L19:
	;
	F_UnlockBuffer(m, v85)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_LockBufferInternal(m, v85, int32(3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v83 != int32(1) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+16)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v145)+6)))
	if v147&int32(2) != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_UnlockBuffer(m, v85)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	if v85 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBufferInternal(m, v85, int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v191+(v85^int32(-1))<<(uint(int32(2))%32))))
	v199 = v193
	goto L27
L29:
	;
	goto L30
L30:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v199 = v195 + v85<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+16)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v203)+6)))
	if v205&int32(2) == int32(0) {
		v223 = v199
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_UnlockBuffer(m, v85)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_LockBufferInternal(m, v85, int32(3))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+16)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v215)+6)))
	if v217&int32(2) != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	F_UnlockBuffer(m, v85)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v260 = int32(0)
	v263 = v256 | v257<<(uint(int32(16))%32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v266 = F_ReadBufferExtended(m, v35, v260, v263, v260, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v83 = v263
	v85 = v266
	goto L10
L39:
	;
	v328 = int32(0)
	v329 = base.B2i32(v328 <= v303)
	if v329 == v328 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[1])))
	F_MemoryContextDelete(m, v1743)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L295
	}
L41:
	;
	F_UnlockReleaseBuffer(m, v303)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L124
	}
L42:
	;
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v347+v752)))
	v770 = v754
	v772 = int32(0)
	goto L41
L43:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+12)))
	if base.Ui32(v348) < base.Ui32(int32(25)) {
		goto L42
	} else {
		goto L47
	}
L44:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333+(v303^int32(-1))<<(uint(int32(2))%32))))
	v347 = v339
	goto L43
L45:
	;
	goto L46
L46:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v347 = v341 + v303<<(uint(int32(13))%32) + int32(-8192)
	goto L43
L47:
	;
	v356 = int32(base.Ui32(v348+int32(_a_F_ginbulkdelete_5))>>(uint(int32(2))%32)) & int32(_a_F_ginbulkdelete_6)
	if v356 == int32(0) {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v363 = v347
	v369 = int32(1)
	v377 = int32(0)
	goto L50
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L121
	}
L50:
	;
	v392 = v369 & int32(_a_F_ginbulkdelete_6)
	v394 = v392 << (uint(int32(2)) % 32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v363+v394)+20))
	v399 = v363 + v396&int32(_a_F_ginbulkdelete_4)
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399)+4)))
	if v400 == int32(0) {
		v636 = v363
		v650 = v377
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v669 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v347+v669)))
	if base.B2i32(v636 == int32(0))|base.B2i32(v636 == v347) != 0 {
		v770 = v671
		v772 = v650
		goto L41
	} else {
		goto L104
	}
L52:
	;
	v665 = v369 + int32(1)
	if base.Ui32(v665&int32(_a_F_ginbulkdelete_6)) <= base.Ui32(v356) {
		v363 = v636
		v369 = v665
		v377 = v650
		goto L50
	} else {
		goto L103
	}
L53:
	;
	if v400 == int32(_a_F_ginbulkdelete_6) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v405 = int32(16)
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399)+2)))
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399))))
	*(*int32)(unsafe.Add(mBase, uint32(v33+v405+v377<<(uint(int32(2))%32)))) = v410 | v411<<(uint(v405)%32)
	v636 = v363
	v650 = v377 + int32(1)
	goto L52
L55:
	;
	goto L56
L56:
	;
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399))))
	v420 = v418 << (uint(int32(16)) % 32)
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399)+2)))
	v422 = v420 | v421
	if int32(0) <= v420 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v560 == int32(0) {
		v636 = v363
		v650 = v377
		goto L52
	} else {
		goto L81
	}
L58:
	;
	F_pfree(m, v529)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L80
	}
L59:
	;
	v442 = int32(0)
	v446 = v442
	v457 = v442
	v459 = v442
	goto L65
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6]))) = v400
	v440 = v400
	v441 = v422 + v399
	goto L59
L61:
	;
	goto L62
L62:
	;
	v432 = F_ginPostingListDecode(m, v399+v422&int32(2147483647), v33+int32(_a_F_ginbulkdelete_7))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6])))
	if int32(0) < v434 {
		v440 = v434
		v441 = v432
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v437 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6]))) = v437
	v528 = v437
	v529 = v432
	goto L58
L65:
	;
	v476 = v446 * int32(6)
	v477 = v441 + v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2764))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2760))
	v480 = m.T0[v479].(func(*base.Module, int32, int32) int32)(m, v477, v478)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6]))) = v509
	if int32(0) <= v420 {
		v560 = v508
		goto L57
	} else {
		goto L79
	}
L67:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2756))
	if v480 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v511 = v446 + int32(1)
	if v511 != v440 {
		v446 = v511
		v457 = v508
		v459 = v509
		goto L65
	} else {
		goto L78
	}
L69:
	;
	v483 = *(*float64)(unsafe.Add(mBase, uint32(v482)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v482)+16)) = base.F64_add(v483, float64(1))
	if v457 != 0 {
		v508 = v457
		v509 = v459
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v493 = *(*float64)(unsafe.Add(mBase, uint32(v482)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v482)+8)) = base.F64_add(v493, float64(1))
	if v457 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v488 = F_palloc_mul(m, int32(6), v440)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v476 == int32(0) {
		v508 = v488
		v509 = v459
		goto L68
	} else {
		goto L74
	}
L74:
	;
	base.MemoryCopy(m, v488, v441, v476)
	v508 = v488
	v509 = v459
	goto L68
L75:
	;
	v499 = v457 + v459*int32(6)
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v477)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v499)+4)) = uint16(v500)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v502
	goto L77
L76:
	;
	goto L77
L77:
	;
	v508 = v457
	v509 = v459 + int32(1)
	goto L68
L78:
	;
	goto L66
L79:
	;
	v528 = v508
	v529 = v441
	goto L58
L80:
	;
	v560 = v528
	goto L57
L81:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6])))
	if v580 <= int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v363 == v347 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v583 = int32(0)
	v596 = v583
	v597 = v583
	goto L82
L84:
	;
	goto L85
L85:
	;
	v587 = F_ginCompressPostingList(m, v560, v580, int32(2712), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587)+6)))
	v596 = v587
	v597 = (v589+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	goto L82
L87:
	;
	v599 = F_PageGetTempPageCopy(m, v347)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	v606 = v363
	v607 = v399
	goto L89
L89:
	;
	v608 = F_gintuple_get_attrnum(m, v53, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v599+v394)+20))
	v606 = v599
	v607 = v599 + v602&int32(_a_F_ginbulkdelete_4)
	goto L89
L91:
	;
	v612 = F_gintuple_get_key(m, v53, v607, v33+int32(_a_F_ginbulkdelete_9))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v614 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[7]))))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6])))
	v617 = F_GinFormTuple(m, v53, v608, v612, v614, v596, v597, v615, int32(1))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v596 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_pfree(m, v596)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	F_PageIndexTupleDelete(m, v606, v392)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v617)+6)))
	v627 = F_PageAddItemExtended(m, v606, v617, v623&int32(_a_F_ginbulkdelete_10), v392, int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v627 != v392 {
		goto L49
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v617)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v560)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v636 = v606
	v650 = v377
	goto L52
L103:
	;
	goto L51
L104:
	;
	v676 = int32(_a_F_ginbulkdelete_11)
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v678 + int32(1)
	F_PageRestoreTempPage(m, v636, v347)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_MarkBufferDirty(m, v303)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v329 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+48))
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705)+118)))
	if v706 != int32(112) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v689+(v303^int32(-1))<<(uint(int32(2))%32))))
	v703 = v695
	goto L107
L109:
	;
	goto L110
L110:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v703 = v697 + v303<<(uint(int32(13))%32) + int32(-8192)
	goto L107
L111:
	;
	v728 = int32(_a_F_ginbulkdelete_11)
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v730 - int32(1)
	v770 = v671
	v772 = v650
	goto L41
L112:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v710 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v704)+32))
	if v713 != 0 {
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v704)+40))
	if v714 != 0 {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	F_XLogRegisterBuffer(m, int32(0), v303, int32(9))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v723 = F_XLogInsert(m, int32(13), int32(64))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v703))) = base.I64_rotl(v723, int64(32))
	goto L111
L121:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v739 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginbulkdelete_12), v33)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_ginbulkdelete_13), int32(619), int32(_a_F_ginbulkdelete_14))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	v788 = int32(0)
	F_vacuum_delay_point(m, v788)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v772 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v809 = v788
	goto L129
L127:
	;
	goto L128
L128:
	;
	if v770 != int32(-1) {
		goto L290
	} else {
		goto L291
	}
L129:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(16)+v809<<(uint(int32(2))%32))))
	v830 = v827
	goto L131
L130:
	;
	goto L128
L131:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v859 = int32(0)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[2])))
	v862 = F_ReadBufferExtended(m, v858, v859, v830, v859, v861)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	v912 = v862
	v914 = v884
	v923 = int32(0)
	goto L149
L133:
	;
	F_LockBufferInternal(m, v862, int32(1))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v862 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L132
L136:
	;
	F_UnlockReleaseBuffer(m, v862)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L147
	}
L137:
	;
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+16)))
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885+v884)+6)))
	if v887&int32(2) != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v870+(v862^int32(-1))<<(uint(int32(2))%32))))
	v884 = v876
	goto L137
L139:
	;
	goto L140
L140:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v884 = v878 + v862<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L141:
	;
	F_UnlockBuffer(m, v862)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v903 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+34)))
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+32)))
	v908 = v903 | v904<<(uint(int32(16))%32)
	goto L136
L144:
	;
	F_LockBufferInternal(m, v862, int32(3))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+16)))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884+v895)+6)))
	if v897&int32(2) == int32(0) {
		v908 = v830
		goto L136
	} else {
		goto L146
	}
L146:
	;
	goto L135
L147:
	;
	v830 = v908
	goto L131
L148:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L288
	}
L149:
	;
	v941 = int32(_a_F_ginbulkdelete_15)
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0]))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0])) = v944
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v948 = v33 + int32(2752)
	v949 = int32(0)
	v950 = m.G0
	v952 = v950 - int32(16)
	m.G0 = v952
	if v912 < v949 {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v1579 = int32(0)
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[2])))
	v1582 = F_ReadBufferExtended(m, v1578, v1579, v827, v1579, v1581)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L1
	} else {
		goto L277
	}
L151:
	;
	goto L150
L152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[0])) = v942
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[1])))
	F_MemoryContextReset(m, v1515)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L257
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L1
	} else {
		goto L254
	}
L154:
	;
	m.G0 = v952 + int32(16)
	goto L152
L155:
	;
	v972 = F_disassembleLeaf(m, v971)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L159
	}
L156:
	;
	v957 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v957+(v912^int32(-1))<<(uint(int32(2))%32))))
	v971 = v963
	goto L155
L157:
	;
	goto L158
L158:
	;
	v965 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v971 = v965 + v912<<(uint(int32(13))%32) + int32(-8192)
	goto L155
L159:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if base.B2i32(v974 == int32(0))|base.B2i32(v974 == v972) != 0 {
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v981 = v974
	v986 = v949
	goto L162
L161:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	v1176 = int32(0)
	if base.B2i32(v1175 == v1176)|base.B2i32(v1175 == v972) == v1176 {
		goto L202
	} else {
		goto L203
	}
L162:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v981)+24))
	if v1009 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v986&int32(1) == int32(0) {
		goto L154
	} else {
		goto L201
	}
L164:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v981)+20))
	v1015 = F_ginPostingListDecode(m, v1012, v981+int32(28))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	v1018 = v1009
	goto L166
L166:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v981)+20))
	if v1019 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v981)+24)) = v1015
	v1018 = v1015
	goto L166
L168:
	;
	v1020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019)+6)))
	v1028 = (v1020+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	goto L170
L169:
	;
	v1028 = int32(_a_F_ginbulkdelete_16)
	goto L170
L170:
	;
	v1030 = v952 + int32(12)
	v1031 = int32(0)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v981)+28))
	if v1034 <= v1031 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v981)+24))
	F_pfree(m, v1140)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L189
	}
L172:
	;
	v1037 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1030))) = v1037
	v1139 = v1037
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1048 = v1031
	v1050 = v1031
	v1059 = v1031
	goto L175
L175:
	;
	v1071 = v1059 * int32(6)
	v1072 = v1018 + v1071
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v948)+12))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v948)+8))
	v1075 = m.T0[v1074].(func(*base.Module, int32, int32) int32)(m, v1072, v1073)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1030))) = v1104
	v1139 = v1102
	goto L171
L177:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v1075 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1106 = v1059 + int32(1)
	if v1106 != v1034 {
		v1048 = v1102
		v1050 = v1104
		v1059 = v1106
		goto L175
	} else {
		goto L188
	}
L179:
	;
	v1078 = *(*float64)(unsafe.Add(mBase, uint32(v1077)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1077)+16)) = base.F64_add(v1078, float64(1))
	if v1048 != 0 {
		v1102 = v1048
		v1104 = v1050
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v1088 = *(*float64)(unsafe.Add(mBase, uint32(v1077)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1077)+8)) = base.F64_add(v1088, float64(1))
	if v1048 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v1083 = F_palloc_mul(m, int32(6), v1034)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	if v1071 == int32(0) {
		v1102 = v1083
		v1104 = v1050
		goto L178
	} else {
		goto L184
	}
L184:
	;
	base.MemoryCopy(m, v1083, v1018, v1071)
	v1102 = v1083
	v1104 = v1050
	goto L178
L185:
	;
	v1094 = v1048 + v1050*int32(6)
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1072)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1094)+4)) = uint16(v1095)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1072)))
	*(*int32)(unsafe.Add(mBase, uint32(v1094))) = v1097
	goto L187
L186:
	;
	goto L187
L187:
	;
	v1102 = v1048
	v1104 = v1050 + int32(1)
	goto L178
L188:
	;
	goto L176
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v981)+24)) = int64(0)
	if v1139 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v952)+12))
	if int32(0) < v1145 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	goto L192
L192:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	if v1167 != v972 {
		v981 = v1167
		goto L162
	} else {
		goto L200
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v981)+28)) = v1162
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	if v1165 != v972 {
		v981 = v1165
		v986 = int32(1)
		goto L162
	} else {
		goto L199
	}
L194:
	;
	v1150 = F_ginCompressPostingList(m, v1139, v1145, v1028, v952+int32(8))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v1158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v981)+8)) = uint8(v1158)
	*(*int32)(unsafe.Add(mBase, uint32(v981)+20)) = int32(0)
	v1162 = v1145
	goto L193
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v981)+20)) = v1150
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v952)+8))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v952)+12))
	if v1153 != v1154 {
		goto L153
	} else {
		goto L198
	}
L198:
	;
	v1156 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v981)+8)) = uint8(v1156)
	v1162 = v1153
	goto L193
L199:
	;
	goto L161
L200:
	;
	goto L163
L201:
	;
	goto L161
L202:
	;
	v1185 = v1175
	v1189 = int32(0)
	goto L205
L203:
	;
	goto L204
L204:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v946)+48))
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1273)+118)))
	if v1274 != int32(112) {
		goto L215
	} else {
		goto L216
	}
L205:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1185)+8)))
	v1214 = int32(0)
	v1216 = v1189 | base.B2i32(v1213 != v1214)
	v1217 = int32(1)
	if base.B2i32(v1216&v1217 == v1214)|base.B2i32(v1213 == v1217) == v1214 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L204
L207:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+20))
	v1227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1226)+6)))
	v1233 = (v1227+int32(1))&int32(_a_F_ginbulkdelete_8) + int32(8)
	v1234 = F_palloc(m, v1233)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L1
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+4))
	if v1241 != v972 {
		v1185 = v1241
		v1189 = v1216
		goto L205
	} else {
		goto L214
	}
L210:
	;
	if v1233 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+20))
	base.MemoryCopy(m, v1234, v1236, v1233)
	goto L213
L212:
	;
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1185)+20)) = v1234
	goto L209
L214:
	;
	goto L206
L215:
	;
	v1285 = int32(_a_F_ginbulkdelete_11)
	v1287 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v1287 + int32(1)
	if v912 < int32(0) {
		goto L224
	} else {
		goto L225
	}
L216:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v1278 <= int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v946)+32))
	if v1281 != 0 {
		goto L215
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	F_computeLeafRecompressWALData(m, v972)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v946)+40))
	if v1282 != 0 {
		goto L215
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	goto L215
L223:
	;
	v1309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1308)+16)))
	v1310 = v1309 + v1308
	v1311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1310)+6)))
	v1313 = v1311 & int32(128)
	if v1313 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1294+(v912^int32(-1))<<(uint(int32(2))%32))))
	v1308 = v1300
	goto L223
L225:
	;
	goto L226
L226:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v1308 = v1302 + v912<<(uint(int32(13))%32) + int32(-8192)
	goto L223
L227:
	;
	v1317 = v1311 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v1310)+6)) = uint16(v1317)
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1308)+16)))
	v1321 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1308+v1319)+4)) = uint16(v1321)
	goto L229
L228:
	;
	goto L229
L229:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if base.B2i32(v1324 == int32(0))|base.B2i32(v1324 == v972) != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1429 = int32(32)
	goto L232
L231:
	;
	v1330 = int32(0)
	v1337 = v1324
	v1341 = base.B2i32(v1313 == v1330)
	v1342 = v1330
	v1343 = v1308 + int32(32)
	goto L233
L232:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1308)+12)) = uint16(v1429)
	F_MarkBufferDirty(m, v912)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L1
	} else {
		goto L242
	}
L233:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1337)+8)))
	v1368 = base.B2i32(v1365 != int32(0)) | v1341
	if v1365 != int32(1) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1429 = v1391 + int32(32)
	goto L232
L235:
	;
	v1371 = int32(1)
	v1373 = int32(0)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+20))
	v1376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1375)+6)))
	v1382 = (v1376+v1371)&int32(_a_F_ginbulkdelete_8) + int32(8)
	if base.B2i32(v1368&v1371 == v1373)|base.B2i32(v1382 == v1373) == v1373 {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v1391 = v1342
	v1392 = v1343
	goto L237
L237:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+4))
	if v1395 != v972 {
		v1337 = v1395
		v1341 = v1368
		v1342 = v1391
		v1343 = v1392
		goto L233
	} else {
		goto L241
	}
L238:
	;
	base.MemoryCopy(m, v1343, v1375, v1382)
	goto L240
L239:
	;
	goto L240
L240:
	;
	v1391 = v1342 + v1382
	v1392 = v1343 + v1382
	goto L237
L241:
	;
	goto L234
L242:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v946)+48))
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433)+118)))
	if v1434 != int32(112) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1461 = int32(_a_F_ginbulkdelete_11)
	v1463 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[8])) = v1463 - int32(1)
	goto L154
L244:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[9]))
	if v1438 <= int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v946)+32))
	if v1441 != 0 {
		goto L243
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L250
	}
L248:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v946)+40))
	if v1442 != 0 {
		goto L243
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	F_XLogRegisterBuffer(m, int32(0), v912, int32(8))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v972)+24))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v972)+28))
	F_XLogRegisterBufData(m, int32(0), v1450, v1451)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v1456 = F_XLogInsert(m, int32(13), int32(144))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v971))) = base.I64_rotl(v1456, int64(32))
	goto L243
L254:
	;
	F_errmsg_internal(m, int32(_a_F_ginbulkdelete_17), int32(0))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_ginbulkdelete_18), int32(782), int32(_a_F_ginbulkdelete_19))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	v1518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v914)+16)))
	v1519 = v914 + v1518
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519)+6)))
	if v1520&int32(128) != 0 {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L1
	} else {
		goto L271
	}
L259:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1519)))
	F_UnlockReleaseBuffer(m, v912)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L269
	}
L260:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1519)))
	F_UnlockReleaseBuffer(m, v912)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L266
	}
L261:
	;
	v1523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v914)+12)))
	if v1523 != int32(32) {
		goto L260
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+4)))
	if v1526 == int32(0) {
		goto L259
	} else {
		goto L265
	}
L264:
	;
	goto L259
L265:
	;
	goto L260
L266:
	;
	if v1529 != int32(-1) {
		v1544 = v1529
		v1545 = v923
		goto L258
	} else {
		goto L267
	}
L267:
	;
	if v923&int32(1) == int32(0) {
		goto L148
	} else {
		goto L268
	}
L268:
	;
	goto L151
L269:
	;
	if v1538 == int32(-1) {
		goto L151
	} else {
		goto L270
	}
L270:
	;
	v1544 = v1538
	v1545 = int32(1)
	goto L258
L271:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2752))
	v1550 = int32(0)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[2])))
	v1553 = F_ReadBufferExtended(m, v1549, v1550, v1544, v1550, v1552)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_LockBufferInternal(m, v1553, int32(3))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	if int32(0) <= v1553 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[4]))
	v1575 = v1561 + v1553<<(uint(int32(13))%32) + int32(-8192)
	goto L276
L275:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, _c_F_ginbulkdelete[5]))
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1568+(v1553^int32(-1))<<(uint(int32(2))%32))))
	v1575 = v1574
	goto L276
L276:
	;
	v912 = v1553
	v914 = v1575
	v923 = v1545
	goto L149
L277:
	;
	F_LockBufferForCleanup(m, v1582)
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[10]))) = int32(0)
	v1588 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[11]))) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6]))) = v1588
	v1592 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[12]))) = uint8(v1592)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[11]))) = v1582
	v1599 = F_ginScanPostingTreeToDelete(m, v33+int32(2752), v33+int32(_a_F_ginbulkdelete_7))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_ginbulkdelete[6])))
	if v1601 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1603 = v1601
	goto L283
L281:
	;
	goto L282
L282:
	;
	F_UnlockReleaseBuffer(m, v1582)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L1
	} else {
		goto L287
	}
L283:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1603)))
	F_pfree(m, v1603)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L1
	} else {
		goto L285
	}
L284:
	;
	goto L282
L285:
	;
	if v1632 != 0 {
		v1603 = v1632
		goto L283
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	goto L148
L288:
	;
	v1701 = v809 + int32(1)
	if v1701 != v772 {
		v809 = v1701
		goto L129
	} else {
		goto L289
	}
L289:
	;
	goto L130
L290:
	;
	v1735 = int32(0)
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1738 = F_ReadBufferExtended(m, v35, v1735, v770, v1735, v1737)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	goto L40
L293:
	;
	F_LockBufferInternal(m, v1738, int32(3))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v303 = v1738
	goto L39
L295:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v33)+2756))
	m.G0 = v33 + int32(_a_F_ginbulkdelete_0)
	return v1746
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
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
			v22 = base.I32_wrap_i64(v17)
			m.G0 = v7 + int32(16)
			return v22 & int32(_a_F_gintuple_get_attrnum_0)
		}
	} else {
		v22 = int32(1)
		m.G0 = v7 + int32(16)
		return v22 & int32(_a_F_gintuple_get_attrnum_0)
	}
}
func F_ginvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 int32
	_ = v74
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int64
	_ = v187
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(_a_F_ginvacuumcleanup_0)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v23 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(_a_F_ginvacuumcleanup_0)
	return v261
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[0]))
	if v27 != int32(4) {
		v261 = l1
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
	v31 = v20 + int32(52)
	F_initGinState(m, v31, v22)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	F_ginInsertCleanup(m, v31, int32(0), v37, v37, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v261 = l1
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
	v58 = l1
	goto L11
L11:
	;
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v60
	v68 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v69 = float64(0)
	if base.F64_gt(v68, v69) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v47 = v20 + int32(52)
	F_initGinState(m, v47, v22)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[0]))
	F_ginInsertCleanup(m, v47, base.B2i32(v51 != int32(4)), int32(0), int32(1), v44)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v58 = v44
	goto L11
L15:
	;
	v72 = v68
	goto L17
L16:
	;
	v72 = v69
	goto L17
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v58)+8)) = v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)) = uint8(v74)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v95
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v102 = int32(0)
	v107 = F_read_stream_begin_relation(m, int32(13), v101, v22, v102, int32(3), v20+int32(8), v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L28
	}
L19:
	;
	F_LockRelationForExtension(m, v22, int32(7))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L25
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if v79 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v83 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v95 = v83
	v96 = v3
	goto L18
L25:
	;
	v89 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_UnlockRelationForExtension(m, v22, int32(7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v95 = v89
	v96 = int32(1)
	goto L18
L28:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v95) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v120 = int32(1)
	v122 = v3
	v124 = v3
	v125 = v3
	v128 = int64(0)
	goto L32
L30:
	;
	v226 = v3
	goto L31
L31:
	;
	F_read_stream_end(m, v107)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L61
	}
L32:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	v226 = v206
	goto L31
L34:
	;
	v133 = F_read_stream_next_buffer(m, v107, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_LockBufferInternal(m, v133, int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if v133 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	F_UnlockReleaseBuffer(m, v133)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L59
	}
L38:
	;
	F_RecordFreeIndexPage(m, v22, v120)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L58
	}
L39:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+14)))
	if v156 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[1]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141+(v133^int32(-1))<<(uint(int32(2))%32))))
	v155 = v147
	goto L39
L41:
	;
	goto L42
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_ginvacuumcleanup[2]))
	v155 = v149 + v133<<(uint(int32(13))%32) + int32(-8192)
	goto L39
L43:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155+v159)+6)))
	if v161&int32(4) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	if v164 == int32(0) {
		goto L38
	} else {
		goto L47
	}
L45:
	;
	v172 = v161
	goto L46
L46:
	;
	if v172&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v167 = F_GlobalVisCheckRemovableXid(m, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v167 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155+v169)+6)))
	v172 = v171
	goto L46
L50:
	;
	v176 = v124 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v176
	v206 = v122
	v207 = v176
	v208 = v125
	v210 = v128
	goto L37
L51:
	;
	goto L52
L52:
	;
	if v172&int32(16) != 0 {
		v206 = v122
		v207 = v124
		v208 = v125
		v210 = v128
		goto L37
	} else {
		goto L53
	}
L53:
	;
	v181 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v181
	if v172&int32(2) == int32(0) {
		v206 = v122
		v207 = v124
		v208 = v181
		v210 = v128
		goto L37
	} else {
		goto L54
	}
L54:
	;
	v187 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v155)+12)))
	if base.Ui64(int64(25)) <= base.Ui64(v187) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v197 = int64(base.Ui64(v187+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
	goto L57
L56:
	;
	v197 = int64(0)
	goto L57
L57:
	;
	v198 = v128 + v197
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v198
	v206 = v122
	v207 = v124
	v208 = v181
	v210 = v198
	goto L37
L58:
	;
	v206 = v122 + int32(1)
	v207 = v124
	v208 = v125
	v210 = v128
	goto L37
L59:
	;
	v214 = v120 + int32(1)
	if v214 != v95 {
		v120 = v214
		v122 = v206
		v124 = v207
		v125 = v208
		v128 = v210
		goto L32
	} else {
		goto L60
	}
L60:
	;
	goto L33
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v95
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ginUpdateStats(m, v236, v20+int32(16), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_FreeSpaceMapVacuum(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v226
	if v96 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_LockRelationForExtension(m, v22, int32(7))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v257 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L70
	}
L67:
	;
	v250 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v250
	F_UnlockRelationForExtension(m, v22, int32(7))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v261 = v58
	goto L1
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v257
	v261 = v58
	goto L1
}
func F_gistadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v13 int32
	_ = v13
	Fn14296(m, l0, l1, l2, l3, int32(_a_F_gistadjustmembers_0), int32(349), int32(_a_F_gistadjustmembers_1), int32(_a_F_gistadjustmembers_2), int32(230), int32(_a_F_gistadjustmembers_3), int32(12))
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
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
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
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
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
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
	v44 = v38
	v46 = v7
	v47 = v7
	goto L1
L1:
	;
	if v47&int32(1) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
	if v114 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v102 = v46
	v103 = v44
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v46&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	F_UnlockBuffer(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	F_ReleaseBuffer(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v70
	v72 = int32(0)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	if v73 != int32(1) {
		v102 = v72
		v103 = v70
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v81 = v70
	goto L14
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_ReleaseBuffer(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v95
	v102 = v72
	v103 = v95
	goto L3
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)))
	if v96 != 0 {
		v81 = v95
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v118 = F_ReadBuffer(m, l0, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v122 = v102 & int32(1)
	if v122 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v118
	goto L20
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_LockBufferInternal(m, v125, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v132 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_gistcheckpage(m, l0, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v150
	if v122 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v132^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L27
L29:
	;
	goto L30
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v150 = v144 + v132<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	v431 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L107
	}
L32:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+24)))
	v44 = v103
	v46 = int32(1)
	v47 = v425
	goto L1
L33:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v190 != 0 {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = base.I64_rotl(v152, int64(32))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+16)))
	v157 = v150 + v156
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+12)))
	if v158&int32(8) == int32(0) {
		v187 = v150
		v188 = v158
		v189 = v157
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v163 = F_BufferGetLSNAtomic(m, v132)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+16)))
	v168 = v166 + v167
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	if v169&int32(8) == int32(0) {
		v187 = v166
		v188 = v169
		v189 = v168
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockBuffer(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_LockBufferInternal(m, v177, int32(3))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v182)+12)))
	if v184&int32(8) != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L102
	}
L44:
	;
	v367 = F_gistinserttuple(m, v19+int32(28), v103, l3, l1, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L10
	} else {
		goto L96
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v352
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+24)))
	v44 = v352
	v46 = int32(0)
	v47 = v360
	goto L1
L46:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v352 = v351
	goto L45
L47:
	;
	if v188&int32(1) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockReleaseBuffer(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L54
	}
L49:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v193)+16))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	if v188&int32(2)|base.B2i32(base.Ui64(v194) < base.Ui64(base.I64_rotl(v195, int64(32)))) != 0 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v188&int32(2) == int32(0) {
		goto L47
	} else {
		goto L53
	}
L52:
	;
	goto L47
L53:
	;
	goto L48
L54:
	;
	goto L46
L55:
	;
	v211 = F_gistchoose(m, l0, v187, l1, l3)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v122 != 0 {
		goto L44
	} else {
		goto L79
	}
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v211<<(uint(int32(2))%32))+20))
	v220 = v213 + v217&int32(_a_F_gistdoinsert_0)
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	if v221 == int32(_a_F_gistdoinsert_1) {
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220))))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+2)))
	v226 = F_gistgetadjusted(m, l0, v220, l1, l3)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L61
	}
L60:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockBuffer(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L10
	} else {
		goto L77
	}
L61:
	;
	if v226 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v122 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockBuffer(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v267 = F_gistinserttuple(m, v19+int32(28), v103, l3, v226, v211)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L73
	}
L66:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_LockBufferInternal(m, v235, int32(3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v239 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v257)))
	if v259 != base.I64_rotl(v260, int64(32)) {
		goto L32
	} else {
		goto L72
	}
L69:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243+(v239^int32(-1))<<(uint(int32(2))%32))))
	v257 = v249
	goto L68
L70:
	;
	goto L71
L71:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v257 = v251 + v239<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	goto L65
L73:
	;
	if v267 == int32(0) {
		goto L60
	} else {
		goto L74
	}
L74:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v271 == int32(0) {
		goto L32
	} else {
		goto L75
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockReleaseBuffer(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	goto L46
L77:
	;
	v282 = F_palloc0(m, int32(32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+28)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v224<<(uint(int32(16))%32) | v225
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+26)) = uint16(v211)
	v352 = v282
	goto L45
L79:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockBuffer(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_LockBufferInternal(m, v293, int32(3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v297 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = base.I64_rotl(v317, int64(32))
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v315)+16)))
	v322 = v315 + v321
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322)+12)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v324 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301+(v297^int32(-1))<<(uint(int32(2))%32))))
	v315 = v307
	goto L82
L84:
	;
	goto L85
L85:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v315 = v309 + v297<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	if v323&int32(1) != 0 {
		goto L44
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v323&int32(8) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	F_UnlockBuffer(m, v297)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+24)))
	v44 = v103
	v46 = int32(0)
	v47 = v332
	goto L1
L91:
	;
	F_UnlockReleaseBuffer(m, v297)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L95
	}
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)+16))
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
	if v323&int32(2) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if base.Ui64(base.I64_rotl(v337, int64(32))) <= base.Ui64(v336) {
		goto L44
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	goto L46
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockBuffer(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v377 = v103
	goto L98
L98:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	F_ReleaseBuffer(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L10
	} else {
		goto L100
	}
L99:
	;
	m.G0 = v19 + int32(80)
	return
L100:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v377)+28))
	if v391 != 0 {
		v377 = v391
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v399 + int32(4)
	F_errmsg(m, int32(_a_F_gistdoinsert_2), v19)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v408 = F_errdetail(m, int32(_a_F_gistdoinsert_3), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	F_errhint(m, int32(_a_F_gistdoinsert_4), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_gistdoinsert_5), int32(768), int32(_a_F_gistdoinsert_6))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	if v431 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v433 + int32(4)
	F_errmsg(m, int32(_a_F_gistdoinsert_7), v19+int32(16))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L10
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v459 = v450
	v462 = int32(0)
	goto L113
L111:
	;
	F_errfinish(m, int32(_a_F_gistdoinsert_5), int32(1209), int32(_a_F_gistdoinsert_8))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v469 = F_palloc(m, int32(8))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L10
	} else {
		goto L115
	}
L114:
	;
	v636 = int32(0)
	F_gistfinishsplit(m, v19+int32(28), v103, l3, v623, v636)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L151
	}
L115:
	;
	if v459 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	if v459 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L117:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	F_LockBufferInternal(m, v558, int32(3))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L137
	}
L118:
	;
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+12)))
	if base.Ui32(v489) < base.Ui32(int32(25)) {
		goto L117
	} else {
		goto L122
	}
L119:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[0]))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v474+(v459^int32(-1))<<(uint(int32(2))%32))))
	v488 = v480
	goto L118
L120:
	;
	goto L121
L121:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[1]))
	v488 = v482 + v459<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	v493 = v489 + int32(_a_F_gistdoinsert_9)
	if v493&int32(_a_F_gistdoinsert_10) == int32(0) {
		goto L117
	} else {
		goto L123
	}
L123:
	;
	v508 = int32(1)
	v510 = int32(0)
	goto L124
L124:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(20)+v508<<(uint(int32(2))%32))))
	v528 = v488 + v525&int32(_a_F_gistdoinsert_0)
	if v510 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	if v537 != 0 {
		v584 = v537
		goto L116
	} else {
		goto L136
	}
L126:
	;
	if v508 != int32(base.Ui32(v493)>>(uint(int32(2))%32))&int32(_a_F_gistdoinsert_11) {
		v508 = v508 + int32(1)
		v510 = v537
		goto L124
	} else {
		goto L135
	}
L127:
	;
	v531 = F_CopyIndexTuple(m, v528)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v533 = F_gistgetadjusted(m, l0, v510, v528, l3)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L10
	} else {
		goto L131
	}
L130:
	;
	v537 = v531
	goto L126
L131:
	;
	if v533 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v535 = v533
	goto L134
L133:
	;
	v535 = v510
	goto L134
L134:
	;
	v537 = v535
	goto L126
L135:
	;
	goto L125
L136:
	;
	goto L117
L137:
	;
	F_gistFindCorrectParent(m, l0, v103)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L10
	} else {
		goto L138
	}
L138:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+26)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v565+v566<<(uint(int32(2))%32))+20))
	v574 = F_CopyIndexTuple(m, v565+v570&int32(_a_F_gistdoinsert_0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	F_UnlockBuffer(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	v584 = v574
	goto L116
L141:
	;
	v615 = int32(_a_F_gistdoinsert_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v584)+4)) = uint16(v615)
	*(*uint16)(unsafe.Add(mBase, uint32(v584)+2)) = uint16(v614)
	v619 = int32(base.Ui32(v614) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v584))) = uint16(v619)
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v459
	v623 = F_lappend(m, v462, v469)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L10
	} else {
		goto L145
	}
L142:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[2]))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v599+(v459^int32(-1))*int32(56))+16))
	v614 = v605
	goto L141
L143:
	;
	goto L144
L144:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_gistdoinsert[3]))
	v608 = int32(56)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607+v459*v608-v608)+16))
	v614 = v613
	goto L141
L145:
	;
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+16)))
	v626 = v488 + v625
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+12)))
	if v627&int32(8) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	v631 = F_ReadBuffer(m, l0, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L10
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	goto L114
L149:
	;
	F_LockBufferInternal(m, v631, int32(3))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	v459 = v631
	v462 = v623
	goto L113
L151:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	F_UnlockReleaseBuffer(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v645
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+24)))
	v44 = v645
	v46 = v636
	v47 = v647
	goto L1
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v9) {
		v17 = int32(base.Ui32(v9+int32(_a_F_gistextractpage_0)) >> (uint(int32(2)) % 32))
	} else {
		v17 = int32(0)
	}
	v19 = v17 & int32(_a_F_gistextractpage_1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v22 = F_palloc_mul(m, int32(4), v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
		} else {
			v28 = int32(1)
			v30 = l0 + int32(20)
			v34 = (v17 + v28) & int32(_a_F_gistextractpage_1)
			if base.Ui32(int32(3)) <= base.Ui32(v34) {
				v37 = int32(2)
				if base.Ui32(v34) <= base.Ui32(v37) {
					v40 = v37
				} else {
					v40 = v34
				}
				v41 = int32(1)
				v42 = v40 - v41
				v50 = v41
				v51 = int32(0)
				for {
					v57 = int32(2)
					v58 = v50 << (uint(v57) % 32)
					v60 = int32(4)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v58+v30)))
					v64 = int32(_a_F_gistextractpage_2)
					*(*int32)(unsafe.Add(mBase, uint32(v22+v58-v60))) = l0 + v63&v64
					v69 = v58 + v60
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v69+v30)))
					*(*int32)(unsafe.Add(mBase, uint32(v22+v69-v60))) = l0 + v74&v64
					v80 = v50 + v57
					v82 = v51 + v57
					if v82 != v42&int32(-2) {
						v50 = v80
						v51 = v82
						continue
					} else {
						break
					}
					break
				}
				if v42&v41 == int32(0) {
				} else {
					v87 = v80
					v95 = v87 << (uint(int32(2)) % 32)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v95+v30)))
					*(*int32)(unsafe.Add(mBase, uint32(v22+v95-int32(4)))) = l0 + v100&int32(_a_F_gistextractpage_2)
				}
			} else {
				v87 = v28
				v95 = v87 << (uint(int32(2)) % 32)
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v95+v30)))
				*(*int32)(unsafe.Add(mBase, uint32(v22+v95-int32(4)))) = l0 + v100&int32(_a_F_gistextractpage_2)
			}
		}
		return v22
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	v4 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+624))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v25 = F_palloc_mul(m, int32(4), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		if v24 <= int32(0) {
			v131 = v4
		} else {
			if v24 != int32(1) {
				v39 = v4
				v44 = v4
				v50 = v4
				for {
					v53 = v21 + v44<<(uint(int32(1))%32)
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53))))
					if v22 != 0 {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v22))))
						if v56 != 0 {
							v69 = v39
						} else {
							v57 = int32(2)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l1+v54<<(uint(v57)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v25+v39<<(uint(v57)%32)))) = v65
							v69 = v39 + int32(1)
						}
					} else {
						v57 = int32(2)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1+v54<<(uint(v57)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v25+v39<<(uint(v57)%32)))) = v65
						v69 = v39 + int32(1)
					}
					v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+2)))
					if v22 != 0 {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v22))))
						if v72 != 0 {
							v85 = v69
						} else {
							v73 = int32(2)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l1+v70<<(uint(v73)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v25+v69<<(uint(v73)%32)))) = v81
							v85 = v69 + int32(1)
						}
					} else {
						v73 = int32(2)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l1+v70<<(uint(v73)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v25+v69<<(uint(v73)%32)))) = v81
						v85 = v69 + int32(1)
					}
					v86 = int32(2)
					v87 = v44 + v86
					v89 = v50 + v86
					if v89 != v24&int32(2147483646) {
						v39 = v85
						v44 = v87
						v50 = v89
						continue
					} else {
						break
					}
					break
				}
				if v24&int32(1) == int32(0) {
					v131 = v85
				} else {
					v97 = v85
					v102 = v87
					v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v102<<(uint(int32(1))%32)))))
					if v22 != 0 {
						v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v22))))
						if v114 != 0 {
							v131 = v97
						} else {
							v115 = int32(2)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v112<<(uint(v115)%32)-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v25+v97<<(uint(v115)%32)))) = v123
							v131 = v97 + int32(1)
						}
					} else {
						v115 = int32(2)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v112<<(uint(v115)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v25+v97<<(uint(v115)%32)))) = v123
						v131 = v97 + int32(1)
					}
				}
			} else {
				v97 = v4
				v102 = v4
				v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v102<<(uint(int32(1))%32)))))
				if v22 != 0 {
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v22))))
					if v114 != 0 {
						v131 = v97
					} else {
						v115 = int32(2)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v112<<(uint(v115)%32)-int32(4))))
						*(*int32)(unsafe.Add(mBase, uint32(v25+v97<<(uint(v115)%32)))) = v123
						v131 = v97 + int32(1)
					}
				} else {
					v115 = int32(2)
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v112<<(uint(v115)%32)-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v25+v97<<(uint(v115)%32)))) = v123
					v131 = v97 + int32(1)
				}
			}
		}
		F_gistMakeUnionItVec(m, l0, v25, v131, l2+int32(48), l2+int32(304))
		mBase = m.M
		v144 = m.ExcPending
		if v144 != 0 {
			return
		} else {
			F_pfree(m, v25)
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v150 = F_palloc_mul(m, int32(4), v149)
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					if v149 <= int32(0) {
						v261 = v4
					} else {
						v154 = int32(0)
						if v149 != int32(1) {
							v165 = int32(0)
							v166 = v154
							v169 = v4
							for {
								v180 = v147 + v166<<(uint(int32(1))%32)
								v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180))))
								if v22 != 0 {
									v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v181))))
									if v183 != 0 {
										v196 = v169
									} else {
										v184 = int32(2)
										v192 = *(*int32)(unsafe.Add(mBase, uint32(l1+v181<<(uint(v184)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v150+v169<<(uint(v184)%32)))) = v192
										v196 = v169 + int32(1)
									}
								} else {
									v184 = int32(2)
									v192 = *(*int32)(unsafe.Add(mBase, uint32(l1+v181<<(uint(v184)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v169<<(uint(v184)%32)))) = v192
									v196 = v169 + int32(1)
								}
								v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+2)))
								if v22 != 0 {
									v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v197))))
									if v199 != 0 {
										v212 = v196
									} else {
										v200 = int32(2)
										v208 = *(*int32)(unsafe.Add(mBase, uint32(l1+v197<<(uint(v200)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v150+v196<<(uint(v200)%32)))) = v208
										v212 = v196 + int32(1)
									}
								} else {
									v200 = int32(2)
									v208 = *(*int32)(unsafe.Add(mBase, uint32(l1+v197<<(uint(v200)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v196<<(uint(v200)%32)))) = v208
									v212 = v196 + int32(1)
								}
								v213 = int32(2)
								v214 = v166 + v213
								v216 = v165 + v213
								if v216 != v149&int32(2147483646) {
									v165 = v216
									v166 = v214
									v169 = v212
									continue
								} else {
									break
								}
								break
							}
							if v149&int32(1) == int32(0) {
								v261 = v212
							} else {
								v224 = v214
								v227 = v212
								v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147+v224<<(uint(int32(1))%32)))))
								if v22 != 0 {
									v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v22))))
									if v241 != 0 {
										v261 = v227
									} else {
										v242 = int32(2)
										v250 = *(*int32)(unsafe.Add(mBase, uint32(l1+v239<<(uint(v242)%32)-int32(4))))
										*(*int32)(unsafe.Add(mBase, uint32(v150+v227<<(uint(v242)%32)))) = v250
										v261 = v227 + int32(1)
									}
								} else {
									v242 = int32(2)
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l1+v239<<(uint(v242)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v227<<(uint(v242)%32)))) = v250
									v261 = v227 + int32(1)
								}
							}
						} else {
							v224 = v154
							v227 = v4
							v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147+v224<<(uint(int32(1))%32)))))
							if v22 != 0 {
								v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v22))))
								if v241 != 0 {
									v261 = v227
								} else {
									v242 = int32(2)
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l1+v239<<(uint(v242)%32)-int32(4))))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v227<<(uint(v242)%32)))) = v250
									v261 = v227 + int32(1)
								}
							} else {
								v242 = int32(2)
								v250 = *(*int32)(unsafe.Add(mBase, uint32(l1+v239<<(uint(v242)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v227<<(uint(v242)%32)))) = v250
								v261 = v227 + int32(1)
							}
						}
					}
					F_gistMakeUnionItVec(m, l0, v150, v261, l2+int32(336), l2+int32(592))
					mBase = m.M
					v275 = m.ExcPending
					if v275 != 0 {
						return
					} else {
						F_pfree(m, v150)
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
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
func F_gseg_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_same_0), int32(0), v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = uint8(base.B2i32(v9 != int64(0)))
		return v3 & int64(4294967295)
	}
}
